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
	if srcData, err := os.ReadFile(srcPath); err == nil {
		srcCode = string(srcData)
		sym := unit.ID.Symbol()
		var metrics evidence.CodeMetrics
		if sym != "" && strings.HasSuffix(unit.ID.Path(), ".go") {
			metrics = evidence.ComputeSymbolMetrics(srcCode, sym)
		} else {
			metrics = evidence.ComputeMetrics(srcCode)
		}
		ev = append(ev, metrics.ToEvidence())
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
func (c *Certifier) CollectRepoEvidence() []domain.Evidence {
	executor := evidence.NewToolExecutor(c.Root)
	return executor.CollectAll()
}

// SaveReportArtifacts generates REPORT_CARD.md, badge.json, and per-unit reports
// from all records in the store. This is a free function usable from both the
// certify and report commands without constructing a full Certifier.
func SaveReportArtifacts(certDir string, store *record.Store, repo, commit string, now time.Time) error {
	records, err := store.ListAll()
	if err != nil {
		return err
	}
	if len(records) == 0 {
		return nil
	}

	// Full report card (markdown)
	fr := report.GenerateFullReport(records, repo, commit, now)
	md := report.FormatFullMarkdown(fr)
	if err := os.WriteFile(filepath.Join(certDir, "REPORT_CARD.md"), []byte(md), 0o644); err != nil {
		return fmt.Errorf("writing REPORT_CARD.md: %w", err)
	}

	// Per-unit reports
	reportsDir := filepath.Join(certDir, "reports")
	if _, err := report.GenerateUnitReports(fr, reportsDir); err != nil {
		return fmt.Errorf("writing unit reports: %w", err)
	}

	// Badge
	card := report.GenerateCard(records, repo, commit, now)
	badge := report.GenerateBadge(card)
	if data, err := report.FormatBadgeJSON(badge); err == nil {
		if writeErr := os.WriteFile(filepath.Join(certDir, "badge.json"), data, 0o644); writeErr != nil {
			return fmt.Errorf("writing badge.json: %w", writeErr)
		}
	}

	return nil
}
