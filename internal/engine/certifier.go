package engine

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/iksnae/code-certification/internal/agent"
	"github.com/iksnae/code-certification/internal/analysis"
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

	// Deep analysis (type-aware cross-file analysis for Go)
	DeepAnalyzer *analysis.DeepGoAnalyzer // nil = skip deep analysis
}

// Certify runs the full certification pipeline for a single unit.
// repoEvidence is shared across units and should be collected once via CollectRepoEvidence().
func (c *Certifier) Certify(ctx context.Context, unit domain.Unit, repoEvidence []domain.Evidence, now time.Time) (*CertifyResult, error) {
	// 1. Match policies → extract rules
	rules := c.matchRules(unit)

	// 2. Build evidence: start with repo evidence copy
	ev := make([]domain.Evidence, len(repoEvidence))
	copy(ev, repoEvidence)

	// 3. Read source + compute per-unit evidence
	srcCode := c.collectUnitEvidence(unit, &ev)

	// 4. Agent review (optional)
	agentResult, aiObs := c.runAgentReview(ctx, unit, srcCode, &ev)

	// 5. Build record
	rec := CertifyUnit(unit, rules, ev, c.ExpiryCfg, now)
	rec.RunID = c.RunID
	if len(c.PolicyVersions) > 0 {
		rec.PolicyVersion = strings.Join(c.PolicyVersions, ",")
	}
	if len(aiObs) > 0 {
		rec.Observations = append(rec.Observations, aiObs...)
	}

	// 6. Apply overrides
	if len(c.Overrides) > 0 {
		rec = override.ApplyAll(rec, c.Overrides)
	}

	// 7. Persist (if store available)
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

// matchRules returns all policy rules that apply to the given unit.
func (c *Certifier) matchRules(unit domain.Unit) []domain.PolicyRule {
	var rules []domain.PolicyRule
	if c.Matcher != nil {
		for _, p := range c.Matcher.Match(unit) {
			rules = append(rules, p.Rules...)
		}
	}
	return rules
}

// collectUnitEvidence reads source code and computes per-unit metrics,
// lint attribution, coverage, and structural analysis.
func (c *Certifier) collectUnitEvidence(unit domain.Unit, ev *[]domain.Evidence) string {
	srcPath := filepath.Join(c.Root, unit.ID.Path())
	isGo := strings.HasSuffix(unit.ID.Path(), ".go")

	srcData, err := os.ReadFile(srcPath)
	if err != nil {
		return ""
	}
	srcCode := string(srcData)
	sym := unit.ID.Symbol()

	// Code metrics
	var metrics evidence.CodeMetrics
	if sym != "" && isGo {
		metrics = evidence.ComputeSymbolMetrics(srcCode, sym)
	} else {
		metrics = evidence.ComputeMetrics(srcCode)
	}
	*ev = append(*ev, metrics.ToEvidence())

	// Per-unit lint attribution
	if len(c.RepoLintFindings) > 0 {
		unitLint := evidence.AttributeLintToFile(c.RepoLintFindings, unit.ID.Path())
		unitLintEv := unitLint.ToEvidence()
		unitLintEv.Source = "golangci-lint:unit"
		unitLintEv.Metrics["unit_lint_errors"] = float64(unitLint.ErrorCount)
		unitLintEv.Metrics["unit_lint_warnings"] = float64(unitLint.WarnCount)
		*ev = append(*ev, unitLintEv)
	}

	// Per-unit coverage attribution
	if c.RepoCoverProfile != "" {
		if covEv, ok := c.buildCoverageEvidence(unit); ok {
			*ev = append(*ev, covEv)
		}
	}

	// Structural analysis (all languages with registered analyzer)
	c.collectStructuralEvidence(unit, srcCode, isGo, sym, ev)

	// Deep analysis (type-aware, Go only for now)
	c.collectDeepEvidence(unit, ev)

	return srcCode
}

// buildCoverageEvidence extracts per-unit coverage from the repo cover profile.
func (c *Certifier) buildCoverageEvidence(unit domain.Unit) (domain.Evidence, bool) {
	cm := evidence.ParseCoverProfilePerFunc(c.RepoCoverProfile)
	unitCov := float64(-1)
	for filePath := range cm {
		if strings.HasSuffix(filePath, "/"+unit.ID.Path()) || filePath == unit.ID.Path() {
			unitCov = evidence.CoverageForFile(cm, filePath)
			break
		}
	}
	if unitCov < 0 {
		return domain.Evidence{}, false
	}
	return domain.Evidence{
		Kind:    domain.EvidenceKindTest,
		Source:  "coverage:unit",
		Passed:  true,
		Summary: fmt.Sprintf("per-unit coverage: %.0f%%", unitCov*100),
		Metrics: map[string]float64{
			"unit_test_coverage": unitCov,
		},
		Confidence: 1.0,
	}, true
}

// collectStructuralEvidence runs structural analysis and appends evidence.
// Uses the analysis.Analyzer interface for all languages with a registered analyzer,
// falling back to legacy Go-only analysis for backward compatibility.
func (c *Certifier) collectStructuralEvidence(unit domain.Unit, srcCode string, isGo bool, sym string, ev *[]domain.Evidence) {
	lang := unit.ID.Language()
	analyzer := analysis.ForLanguage(lang)

	if analyzer != nil {
		c.collectStructuralViaAnalyzer(analyzer, unit, srcCode, sym, ev)
		return
	}

	// Legacy fallback for Go (should not be reached since Go analyzer is registered)
	if !isGo {
		return
	}
	c.collectStructuralLegacyGo(srcCode, sym, unit.Type, ev)
}

// collectStructuralViaAnalyzer uses the unified Analyzer interface.
func (c *Certifier) collectStructuralViaAnalyzer(a analysis.Analyzer, unit domain.Unit, srcCode string, sym string, ev *[]domain.Evidence) {
	srcBytes := []byte(srcCode)
	path := unit.ID.Path()

	if sym != "" {
		metrics, err := a.Analyze(path, srcBytes, sym)
		if err != nil {
			return
		}
		// Merge file-level metrics
		fileMeta, _ := a.AnalyzeFile(path, srcBytes)
		metrics.HasInitFunc = fileMeta.HasInitFunc
		metrics.GlobalMutableCount = fileMeta.GlobalMutableCount
		*ev = append(*ev, metrics.ToEvidence())
	} else {
		fileMeta, err := a.AnalyzeFile(path, srcBytes)
		if err != nil {
			return
		}
		metrics := analysis.Metrics{
			HasInitFunc:        fileMeta.HasInitFunc,
			GlobalMutableCount: fileMeta.GlobalMutableCount,
		}
		*ev = append(*ev, metrics.ToEvidence())
	}
}

// collectStructuralLegacyGo is the original Go-only analysis path.
// Retained for backward compatibility; should not normally be reached.
func (c *Certifier) collectStructuralLegacyGo(srcCode string, sym string, unitType domain.UnitType, ev *[]domain.Evidence) {
	if sym != "" {
		var structural evidence.StructuralMetrics
		if unitType == domain.UnitTypeClass {
			structural = evidence.AnalyzeGoType(srcCode, sym)
		} else {
			structural = evidence.AnalyzeGoFunc(srcCode, sym)
		}
		fileMeta := evidence.AnalyzeGoFile(srcCode)
		structural.HasInitFunc = fileMeta.HasInitFunc
		structural.GlobalMutableCount = fileMeta.GlobalMutableCount
		*ev = append(*ev, structural.ToEvidence())
	} else {
		fileMeta := evidence.AnalyzeGoFile(srcCode)
		structural := evidence.StructuralMetrics{
			HasInitFunc:        fileMeta.HasInitFunc,
			GlobalMutableCount: fileMeta.GlobalMutableCount,
		}
		*ev = append(*ev, structural.ToEvidence())
	}
}

// collectDeepEvidence appends type-aware metrics (fan-in, fan-out, dead code)
// from the DeepGoAnalyzer if available.
func (c *Certifier) collectDeepEvidence(unit domain.Unit, ev *[]domain.Evidence) {
	if c.DeepAnalyzer == nil {
		return
	}
	if unit.ID.Language() != "go" {
		return
	}

	sym := unit.ID.Symbol()
	if sym == "" {
		return
	}

	// Determine package path from the unit's file path
	pkgPath := c.resolveGoPackagePath(unit)
	if pkgPath == "" {
		return
	}

	result, ok := c.DeepAnalyzer.Lookup(pkgPath, sym)
	if !ok {
		return
	}

	metrics := map[string]float64{
		"fan_in":       float64(result.FanIn),
		"fan_out":      float64(result.FanOut),
		"is_dead_code": 0,
	}
	if result.IsDeadCode {
		metrics["is_dead_code"] = 1
	}

	*ev = append(*ev, domain.Evidence{
		Kind:       domain.EvidenceKindStructural,
		Source:     "deep-analysis",
		Passed:     true,
		Summary:    fmt.Sprintf("deep: fan_in=%d fan_out=%d dead=%v", result.FanIn, result.FanOut, result.IsDeadCode),
		Metrics:    metrics,
		Confidence: 1.0,
	})
}

// resolveGoPackagePath attempts to determine the Go package path for a unit.
// It looks at the unit's file path relative to the module root.
func (c *Certifier) resolveGoPackagePath(unit domain.Unit) string {
	filePath := unit.ID.Path()
	dir := filepath.Dir(filePath)

	// Try to find the module path from go.mod
	modPath := c.Root
	modName := ""

	// Walk up from the file's directory to find go.mod
	checkDir := filepath.Join(c.Root, dir)
	for {
		gomod := filepath.Join(checkDir, "go.mod")
		if data, err := os.ReadFile(gomod); err == nil {
			// Extract module name from go.mod
			for _, line := range strings.Split(string(data), "\n") {
				line = strings.TrimSpace(line)
				if strings.HasPrefix(line, "module ") {
					modName = strings.TrimSpace(strings.TrimPrefix(line, "module"))
					modPath = checkDir
					break
				}
			}
			break
		}
		parent := filepath.Dir(checkDir)
		if parent == checkDir {
			break
		}
		checkDir = parent
	}

	if modName == "" {
		return ""
	}

	// Compute relative path from module root to the file's directory
	absDir := filepath.Join(c.Root, dir)
	rel, err := filepath.Rel(modPath, absDir)
	if err != nil || rel == "." {
		return modName
	}

	return modName + "/" + filepath.ToSlash(rel)
}

// runAgentReview optionally runs agent-assisted review for the unit.
func (c *Certifier) runAgentReview(ctx context.Context, unit domain.Unit, srcCode string, ev *[]domain.Evidence) (*agent.ReviewResult, []string) {
	if c.Agent == nil {
		return nil, nil
	}
	timeout := c.AgentTimeout
	if timeout == 0 {
		timeout = 30 * time.Second
	}
	if c.Agent.IsLocal() && timeout < 120*time.Second {
		timeout = 120 * time.Second
	}
	agentCtx, cancel := context.WithTimeout(ctx, timeout)
	result := c.Agent.ReviewUnit(agentCtx, unit, srcCode, *ev)
	cancel()

	if result.Reviewed {
		*ev = append(*ev, result.ToEvidence())
	} else if result.Prescreened {
		*ev = append(*ev, result.ToPrescreenEvidence())
	}
	return &result, agent.FormatDeepObservations(result)
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
