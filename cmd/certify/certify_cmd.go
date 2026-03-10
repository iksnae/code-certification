package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/iksnae/code-certification/internal/agent"
	"github.com/iksnae/code-certification/internal/config"
	"github.com/iksnae/code-certification/internal/discovery"
	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/engine"
	"github.com/iksnae/code-certification/internal/evidence"
	"github.com/iksnae/code-certification/internal/override"
	"github.com/iksnae/code-certification/internal/queue"
	"github.com/iksnae/code-certification/internal/record"
	"github.com/iksnae/code-certification/internal/report"
	"github.com/spf13/cobra"
)

var (
	certifyPath       string
	certifySkipAgent  bool
	certifyBatch      int
	certifyResetQueue bool
	certifyTarget     []string
	certifyDiffBase   string
)

var certifyCmd = &cobra.Command{
	Use:   "certify",
	Short: "Evaluate and certify code units",
	Long: `Certify evaluates code units against policies and collects evidence.

On first run, builds a work queue from the index. Subsequent runs resume
from where they left off. Use --batch to control how many units to process
per run (useful for rate-limited agent review).

Examples:
  certify certify                    # process all units (deterministic only)
  certify certify --batch 20         # process 20 units then stop
  certify certify --reset-queue      # rebuild queue and start over
  certify certify --skip-agent       # skip agent review, deterministic only`,
	RunE: runCertify,
}

func init() {
	certifyCmd.Flags().StringVar(&certifyPath, "path", "", "Path to repository (default: current directory)")
	certifyCmd.Flags().BoolVar(&certifySkipAgent, "skip-agent", false, "Skip agent-assisted review")
	certifyCmd.Flags().IntVar(&certifyBatch, "batch", 0, "Max units to process per run (0=all)")
	certifyCmd.Flags().BoolVar(&certifyResetQueue, "reset-queue", false, "Rebuild queue from index")
	certifyCmd.Flags().StringSliceVar(&certifyTarget, "target", nil, "Target specific paths/directories (can specify multiple)")
	certifyCmd.Flags().StringVar(&certifyDiffBase, "diff-base", "", "Only certify files changed since this git ref")
}

// certifyContext holds all loaded state for a certification run.
type certifyContext struct {
	root      string
	certDir   string
	cfg       domain.Config
	units     []domain.Unit
	unitMap   map[string]domain.Unit
	packs     []domain.PolicyPack
	overrides []domain.Override
	wq        *queue.Queue
	queuePath string
	store     *record.Store
	coord     *agent.Coordinator
	repoEv    []domain.Evidence
}

func runCertify(cmd *cobra.Command, args []string) error {
	ctx, err := loadCertifyContext()
	if err != nil {
		return err
	}

	remaining := ctx.printQueueStatus()
	if remaining == 0 {
		ctx.wq.Save(ctx.queuePath)
		return nil
	}

	ctx.coord = setupAgent(ctx.cfg, certifySkipAgent)
	ctx.collectRepoEvidence()

	certified, observations, failed, processed := ctx.processQueue(cmd, remaining)

	ctx.printSummary(certified, observations, failed, processed)
	ctx.saveReportArtifacts()
	ctx.wq.Save(ctx.queuePath)

	if ctx.cfg.Mode == domain.ModeEnforcing && failed > 0 {
		return fmt.Errorf("%d units failed certification in enforcing mode", failed)
	}
	return nil
}

func loadCertifyContext() (*certifyContext, error) {
	root := certifyPath
	if root == "" {
		root, _ = os.Getwd()
	}
	certDir := filepath.Join(root, ".certification")

	cfg, err := config.LoadFromDir(certDir)
	if err != nil {
		cfg = domain.DefaultConfig()
	}

	idx, err := discovery.LoadIndex(filepath.Join(certDir, "index.json"))
	if err != nil {
		return nil, fmt.Errorf("loading index (run 'certify scan' first): %w", err)
	}

	packs, err := config.LoadPolicyPacks(filepath.Join(certDir, "policies"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: loading policies: %v\n", err)
	}

	overrides, err := override.LoadDir(filepath.Join(certDir, "overrides"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: loading overrides: %v\n", err)
	}

	wq := loadQueue(filepath.Join(certDir, "queue.json"))
	units := filterUnits(root, idx)

	unitMap := make(map[string]domain.Unit, len(units))
	for _, u := range units {
		unitMap[u.ID.String()] = u
		wq.Enqueue(u.ID.String(), u.ID.Path())
	}

	return &certifyContext{
		root: root, certDir: certDir, cfg: cfg,
		units: units, unitMap: unitMap,
		packs: packs, overrides: overrides,
		wq: wq, queuePath: filepath.Join(certDir, "queue.json"),
		store: record.NewStore(filepath.Join(certDir, "records")),
	}, nil
}

func loadQueue(queuePath string) *queue.Queue {
	if certifyResetQueue {
		return queue.New()
	}
	wq, err := queue.Load(queuePath)
	if err != nil {
		return queue.New()
	}
	return wq
}

func filterUnits(root string, idx *discovery.Index) []domain.Unit {
	units := idx.Units()
	if certifyDiffBase != "" {
		changedFiles, err := discovery.ChangedFiles(root, certifyDiffBase, "HEAD")
		if err != nil {
			fmt.Fprintf(os.Stderr, "warning: git diff failed: %v — certifying all units\n", err)
		} else {
			units = discovery.FilterChanged(units, changedFiles)
			fmt.Printf("  Changed files since %s: %d → %d units\n", certifyDiffBase, len(idx.Units()), len(units))
		}
	}
	if len(certifyTarget) > 0 {
		units = discovery.FilterByPaths(units, certifyTarget)
		fmt.Printf("  Targeting %v: %d units\n", certifyTarget, len(units))
	}
	return units
}

func (c *certifyContext) printQueueStatus() int {
	stats := c.wq.Stats()
	remaining := stats.Pending + stats.InProgress
	if remaining == 0 {
		fmt.Printf("  Queue complete: %d/%d processed (%d skipped)\n",
			stats.Completed, stats.Total, stats.Skipped)
		fmt.Println("  Use --reset-queue to re-process all units.")
		return 0
	}
	fmt.Printf("  Queue: %d pending, %d completed, %d skipped, %d failed of %d total\n",
		remaining, stats.Completed, stats.Skipped, stats.Failed, stats.Total)
	return remaining
}

func setupAgent(cfg domain.Config, skip bool) *agent.Coordinator {
	if skip {
		fmt.Println("  Agent review: skipped (--skip-agent)")
		return nil
	}
	// Explicit config takes priority
	if cfg.Agent.Enabled {
		return setupExplicitAgent(cfg)
	}
	// Respect explicit disable even when key is present
	if cfg.Agent.ExplicitlyDisabled {
		return nil
	}
	// Auto-detect: check for API key in environment
	return setupConservativeAgent()
}

// setupExplicitAgent builds the full agent pipeline from explicit config.
func setupExplicitAgent(cfg domain.Config) *agent.Coordinator {
	baseURL := cfg.Agent.Provider.BaseURL
	apiKey := ""

	// Local providers (no API key needed)
	isLocal := isLocalURL(baseURL)

	if cfg.Agent.Provider.APIKeyEnv != "" {
		apiKey = os.Getenv(cfg.Agent.Provider.APIKeyEnv)
		if apiKey == "" && !isLocal {
			fmt.Fprintf(os.Stderr, "  Agent review configured but %s not set — skipping\n", cfg.Agent.Provider.APIKeyEnv)
			return nil
		}
	} else if !isLocal {
		// No api_key_env and not local — check common env vars
		apiKey, _ = agent.DetectAPIKey()
		if apiKey == "" {
			fmt.Fprintf(os.Stderr, "  Agent review configured but no API key found — skipping\n")
			return nil
		}
	}

	// Build model list from config
	models := []string{}
	if cfg.Agent.Models.Prescreen != "" {
		models = append(models, cfg.Agent.Models.Prescreen)
	}
	if cfg.Agent.Models.Review != "" && cfg.Agent.Models.Review != cfg.Agent.Models.Prescreen {
		models = append(models, cfg.Agent.Models.Review)
	}
	if cfg.Agent.Models.Fallback != "" {
		models = append(models, cfg.Agent.Models.Fallback)
	}
	if len(models) == 0 {
		models = []string{"qwen3:latest"}
	}

	var provider agent.Provider
	strategy := agent.StrategyStandard
	tokenBudget := 50000

	if isLocal {
		provider = agent.NewLocalProvider(baseURL, "local")
		strategy = agent.StrategyLocal // deep review, no prescreen gate
		tokenBudget = 0                // unlimited — local tokens are free
		fmt.Printf("  Agent review: enabled (local %s, deep review, models: %v)\n", baseURL, models)
	} else {
		provider = agent.NewModelChain(
			baseURL, apiKey,
			cfg.Agent.Provider.HTTPReferer, cfg.Agent.Provider.XTitle,
			models,
		)
		fmt.Println("  Agent review: enabled (model chain with fallback)")
	}

	cb := agent.NewCircuitBreaker(provider, 5)
	return agent.NewCoordinator(cb, agent.CoordinatorConfig{
		Models: cfg.Agent.Models, Strategy: strategy, TokenBudget: tokenBudget,
	})
}

// isLocalURL returns true if the URL points to a local server.
func isLocalURL(u string) bool {
	return strings.Contains(u, "localhost") || strings.Contains(u, "127.0.0.1") || strings.Contains(u, "0.0.0.0")
}

// setupConservativeAgent auto-detects available providers and builds a conservative coordinator.
func setupConservativeAgent() *agent.Coordinator {
	providers := agent.DetectProviders()
	if len(providers) == 0 {
		return nil
	}
	summary := agent.FormatProviderSummary(providers)
	fmt.Printf("  Agent review: auto-detected [%s] (conservative mode)\n", summary)
	return agent.NewConservativeCoordinator(providers)
}

func (c *certifyContext) collectRepoEvidence() {
	fmt.Println("  Collecting repo-level evidence...")
	executor := evidence.NewToolExecutor(c.root)
	c.repoEv = executor.CollectAll()
	fmt.Printf("  Collected %d repo-level evidence items\n", len(c.repoEv))
}

func (c *certifyContext) processQueue(cmd *cobra.Command, remaining int) (certified, observations, failed, processed int) {
	now := time.Now()
	batchSize := remaining
	if certifyBatch > 0 && certifyBatch < batchSize {
		batchSize = certifyBatch
	}

	for processed < batchSize {
		item, ok := c.wq.Next()
		if !ok {
			break
		}
		unit, exists := c.unitMap[item.UnitID]
		if !exists {
			c.wq.Skip(item.UnitID, "unit not in index")
			c.wq.Save(c.queuePath)
			continue
		}

		processed++
		if processed%20 == 0 || processed == batchSize {
			fmt.Printf("\r  Processing... %d/%d", processed, batchSize)
		}

		rec := c.certifyUnit(cmd, unit, item.UnitID, now)

		switch {
		case rec.Status == domain.StatusCertified:
			certified++
		case rec.Status == domain.StatusCertifiedWithObservations || rec.Status == domain.StatusExempt:
			observations++
		default:
			failed++
		}
		c.wq.Save(c.queuePath)
	}
	return
}

func (c *certifyContext) certifyUnit(cmd *cobra.Command, unit domain.Unit, unitID string, now time.Time) domain.CertificationRecord {
	matcher := config.NewPolicyMatcher(c.packs)
	matched := matcher.Match(unit)
	var rules []domain.PolicyRule
	for _, p := range matched {
		rules = append(rules, p.Rules...)
	}

	ev := make([]domain.Evidence, len(c.repoEv))
	copy(ev, c.repoEv)
	var aiObs []string

	srcPath := filepath.Join(c.root, unit.ID.Path())
	var srcCode string
	if srcData, readErr := os.ReadFile(srcPath); readErr == nil {
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

	if c.coord != nil {
		// Local models are slower — give them more time
		timeout := 30 * time.Second
		if c.coord.IsLocal() {
			timeout = 120 * time.Second
		}
		ctx, cancel := context.WithTimeout(cmd.Context(), timeout)
		result := c.coord.ReviewUnit(ctx, unit, srcCode, ev)
		cancel()
		model := ""
		if len(result.ModelsUsed) > 0 {
			model = result.ModelsUsed[0]
		}
		if result.Reviewed {
			ev = append(ev, result.ToEvidence())
			c.wq.Complete(unitID, model)
		} else if result.Prescreened {
			ev = append(ev, result.ToPrescreenEvidence())
			c.wq.Complete(unitID, model)
		} else {
			c.wq.Skip(unitID, "prescreen: no review needed")
		}
		// Collect AI observations for the record
		aiObs = append(aiObs, agent.FormatDeepObservations(result)...)
	} else {
		c.wq.Complete(unitID, "")
	}

	rec := engine.CertifyUnit(unit, rules, ev, c.cfg.Expiry, now)
	if len(aiObs) > 0 {
		rec.Observations = append(rec.Observations, aiObs...)
	}
	if len(c.overrides) > 0 {
		rec = override.ApplyAll(rec, c.overrides)
	}
	if err := c.store.Save(rec); err != nil {
		fmt.Fprintf(os.Stderr, "\nwarning: saving record for %s: %v\n", unit.ID, err)
	}
	c.store.AppendHistory(rec)
	return rec
}

func defaultConfigObj() domain.Config {
	return domain.DefaultConfig()
}

func (c *certifyContext) printSummary(certified, observations, failed, processed int) {
	total := certified + observations + failed
	if c.coord != nil {
		filesReviewed, totalFiles, tokens := c.coord.Stats()
		fmt.Printf("\n  Agent: %d/%d files reviewed, %d tokens used\n", filesReviewed, totalFiles, tokens)
	}

	fmt.Printf("\n  Processed %d units this run\n", processed)

	finalStats := c.wq.Stats()
	if finalStats.Pending > 0 {
		fmt.Printf("  %d units remaining — run again to continue\n", finalStats.Pending)
	} else {
		fmt.Printf("  Queue complete!\n")
	}

	fmt.Printf("✓ Certified %d/%d units", certified+observations, total)
	if observations > 0 {
		fmt.Printf(" (%d with observations)", observations)
	}
	if failed > 0 {
		fmt.Printf(" (%d need attention)", failed)
	}
	fmt.Println()
}

// saveReportArtifacts generates REPORT_CARD.md and badge.json after certification.
func (c *certifyContext) saveReportArtifacts() {
	records, err := c.store.ListAll()
	if err != nil || len(records) == 0 {
		return
	}
	now := time.Now()
	repo := detectRepoName(c.root)
	commit := detectCommit(c.root)

	// Save report card
	fr := report.GenerateFullReport(records, repo, commit, now)
	md := report.FormatFullMarkdown(fr)
	os.WriteFile(filepath.Join(c.certDir, "REPORT_CARD.md"), []byte(md), 0o644)

	// Save per-unit reports
	reportsDir := filepath.Join(c.certDir, "reports")
	if n, err := report.GenerateUnitReports(fr, reportsDir); err == nil {
		fmt.Printf("✓ %d unit report cards written to %s\n", n, reportsDir)
	}

	// Save badge
	card := report.GenerateCard(records, repo, commit, now)
	badge := report.GenerateBadge(card)
	if data, err := report.FormatBadgeJSON(badge); err == nil {
		os.WriteFile(filepath.Join(c.certDir, "badge.json"), data, 0o644)
	}
}
