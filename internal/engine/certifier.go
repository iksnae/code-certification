package engine

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/iksnae/code-certification/internal/agent"
	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/evidence"
	"github.com/iksnae/code-certification/internal/override"
	"github.com/iksnae/code-certification/internal/policy"
	"github.com/iksnae/code-certification/internal/record"
	"github.com/iksnae/code-certification/internal/report"
)

// CertifyResult holds the outcome of certifying a single unit.
type CertifyResult struct {
	Record      domain.CertificationRecord
	AgentReview *agent.ReviewResult // nil if agent not used or skipped
}

// Certifier is the service that owns the full certification pipeline.
// It handles policy matching, evidence collection, agent review,
// record building, override application, and persistence.
type Certifier struct {
	Root           string              // repo root (for reading source files)
	Store          *record.Store       // record persistence (nil = don't save)
	Matcher        *policy.Matcher     // policy matcher (nil = no policy rules)
	Overrides      []domain.Override   // governance overrides
	ExpiryCfg      domain.ExpiryConfig // expiry window config
	Agent          *agent.Coordinator  // optional AI reviewer (nil = skip)
	AgentTimeout   time.Duration       // per-unit timeout for agent calls
	RunID          string              // current run ID (set once per invocation)
	PolicyVersions []string            // active policy pack versions ("name@version")

	// Per-unit attribution data (set by CollectRepoEvidence or manually for tests)
	RepoLintFindings []evidence.LintFinding // raw lint findings for per-unit attribution
	RepoCoverProfile string                 // raw coverage profile for per-unit coverage
}

// Certify runs the full certification pipeline for a single unit.
// repoEvidence is shared across units and should be collected once via CollectRepoEvidence().
func (c *Certifier) Certify(ctx context.Context, unit domain.Unit, repoEvidence []domain.Evidence, now time.Time) (*CertifyResult, error) {
	// 1. Match policies → extract rules
	var rules []domain.PolicyRule
	if c.Matcher != nil {
		for _, p := range c.Matcher.Match(unit) {
			rules = append(rules, p.Rules...)
		}
	}

	// 2. Build evidence: start with repo evidence copy
	ev := make([]domain.Evidence, len(repoEvidence))
	copy(ev, repoEvidence)

	// 3. Read source + compute per-unit metrics
	var srcCode string
	srcPath := filepath.Join(c.Root, unit.ID.Path())
	isGo := strings.HasSuffix(unit.ID.Path(), ".go")
	if srcData, err := os.ReadFile(srcPath); err == nil {
		srcCode = string(srcData)
		sym := unit.ID.Symbol()
		var metrics evidence.CodeMetrics
		if sym != "" && isGo {
			metrics = evidence.ComputeSymbolMetrics(srcCode, sym)
		} else {
			metrics = evidence.ComputeMetrics(srcCode)
		}
		ev = append(ev, metrics.ToEvidence())

		// 3b. Per-unit lint attribution
		if len(c.RepoLintFindings) > 0 {
			unitLint := evidence.AttributeLintToFile(c.RepoLintFindings, unit.ID.Path())
			unitLintEv := unitLint.ToEvidence()
			unitLintEv.Source = "golangci-lint:unit"
			unitLintEv.Metrics["unit_lint_errors"] = float64(unitLint.ErrorCount)
			unitLintEv.Metrics["unit_lint_warnings"] = float64(unitLint.WarnCount)
			ev = append(ev, unitLintEv)
		}

		// 3c. Per-unit coverage attribution
		if c.RepoCoverProfile != "" {
			cm := evidence.ParseCoverProfilePerFunc(c.RepoCoverProfile)
			// Try matching by file path (coverage profiles use module-qualified paths)
			unitCov := float64(-1)
			for filePath := range cm {
				if strings.HasSuffix(filePath, "/"+unit.ID.Path()) || filePath == unit.ID.Path() {
					unitCov = evidence.CoverageForFile(cm, filePath)
					break
				}
			}
			if unitCov >= 0 {
				covEv := domain.Evidence{
					Kind:    domain.EvidenceKindTest,
					Source:  "coverage:unit",
					Passed:  true,
					Summary: fmt.Sprintf("per-unit coverage: %.0f%%", unitCov*100),
					Metrics: map[string]float64{
						"unit_test_coverage": unitCov,
					},
					Confidence: 1.0,
				}
				ev = append(ev, covEv)
			}
		}

		// 3d. Structural analysis (Go only)
		if isGo && sym != "" {
			var structural evidence.StructuralMetrics
			if unit.Type == domain.UnitTypeClass {
				structural = evidence.AnalyzeGoType(srcCode, sym)
			} else {
				structural = evidence.AnalyzeGoFunc(srcCode, sym)
			}

			// Merge file-level metrics into the structural evidence
			fileMeta := evidence.AnalyzeGoFile(srcCode)
			structural.HasInitFunc = fileMeta.HasInitFunc
			structural.GlobalMutableCount = fileMeta.GlobalMutableCount

			ev = append(ev, structural.ToEvidence())
		} else if isGo && sym == "" {
			// File-level Go unit: only file metrics
			fileMeta := evidence.AnalyzeGoFile(srcCode)
			structural := evidence.StructuralMetrics{
				HasInitFunc:        fileMeta.HasInitFunc,
				GlobalMutableCount: fileMeta.GlobalMutableCount,
			}
			ev = append(ev, structural.ToEvidence())
		}
	}

	// 4. Agent review (optional)
	var agentResult *agent.ReviewResult
	var aiObs []string
	if c.Agent != nil {
		timeout := c.AgentTimeout
		if timeout == 0 {
			timeout = 30 * time.Second
		}
		if c.Agent.IsLocal() && timeout < 120*time.Second {
			timeout = 120 * time.Second
		}
		agentCtx, cancel := context.WithTimeout(ctx, timeout)
		result := c.Agent.ReviewUnit(agentCtx, unit, srcCode, ev)
		cancel()

		if result.Reviewed {
			ev = append(ev, result.ToEvidence())
		} else if result.Prescreened {
			ev = append(ev, result.ToPrescreenEvidence())
		}
		aiObs = agent.FormatDeepObservations(result)
		agentResult = &result
	}

	// 5. Build record via existing CertifyUnit
	rec := CertifyUnit(unit, rules, ev, c.ExpiryCfg, now)

	// 5b. Populate run metadata
	rec.RunID = c.RunID
	if len(c.PolicyVersions) > 0 {
		rec.PolicyVersion = strings.Join(c.PolicyVersions, ",")
	}

	// 6. Merge AI observations
	if len(aiObs) > 0 {
		rec.Observations = append(rec.Observations, aiObs...)
	}

	// 7. Apply overrides
	if len(c.Overrides) > 0 {
		rec = override.ApplyAll(rec, c.Overrides)
	}

	// 8. Persist (if store available)
	if c.Store != nil {
		if err := c.Store.Save(rec); err != nil {
			return nil, fmt.Errorf("saving record for %s: %w", unit.ID, err)
		}
		c.Store.AppendHistory(rec)
	}

	return &CertifyResult{
		Record:      rec,
		AgentReview: agentResult,
	}, nil
}

// CollectRepoEvidence runs all available tool runners and returns repo-level evidence.
// Also retains raw lint findings and coverage profile for per-unit attribution.
func (c *Certifier) CollectRepoEvidence() []domain.Evidence {
	executor := evidence.NewToolExecutor(c.Root)
	ev := executor.CollectAll()
	c.RepoLintFindings = executor.LintFindings()
	c.RepoCoverProfile = executor.CoverageProfile()
	return ev
}

// SaveReportArtifacts writes REPORT_CARD.md (compact summary), badge.json,
// and the unit certificate tree from a pre-computed FullReport.
//
// REPORT_CARD.md is the Card format — a compact summary with overall grade,
// language breakdown, package links, and top issues.
//
// The reports/ tree contains navigable markdown certificates:
//
//	reports/index.md → reports/<pkg>/index.md → reports/<pkg>/<file>/<symbol>.md
//
// For the full per-unit report, use `certify report --format full`.
// For interactive browsing, use `certify report --site`.
func SaveReportArtifacts(certDir string, fr report.FullReport) error {
	// Populate package summaries for card navigation links
	fr.Card.Packages = report.BuildPackageSummaries(fr)

	// Compact report card with package links
	md := report.FormatCardMarkdown(fr.Card)
	if err := os.WriteFile(filepath.Join(certDir, "REPORT_CARD.md"), []byte(md), 0o644); err != nil {
		return fmt.Errorf("writing REPORT_CARD.md: %w", err)
	}

	// Unit certificate tree (committed to repo)
	reportsDir := filepath.Join(certDir, "reports")
	if _, err := report.GenerateReportTree(fr, reportsDir); err != nil {
		return fmt.Errorf("writing unit certificates: %w", err)
	}

	// Badge (uses the Card already embedded in FullReport)
	badge := report.GenerateBadge(fr.Card)
	if data, err := report.FormatBadgeJSON(badge); err == nil {
		if writeErr := os.WriteFile(filepath.Join(certDir, "badge.json"), data, 0o644); writeErr != nil {
			return fmt.Errorf("writing badge.json: %w", writeErr)
		}
	}

	return nil
}

// SaveReportArtifactsFromStore is a convenience wrapper that loads records
// from the store, generates a FullReport, and writes all artifacts.
// Used by the certify command where no FullReport exists yet.
func SaveReportArtifactsFromStore(certDir string, store *record.Store, repo, commit string, now time.Time) error {
	records, err := store.ListAll()
	if err != nil {
		return err
	}
	if len(records) == 0 {
		return nil
	}
	fr := report.GenerateFullReport(records, repo, commit, now)
	return SaveReportArtifacts(certDir, fr)
}
