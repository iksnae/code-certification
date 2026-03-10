package main

import (
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
	"github.com/iksnae/code-certification/internal/override"
	"github.com/iksnae/code-certification/internal/policy"
	"github.com/iksnae/code-certification/internal/queue"
	"github.com/iksnae/code-certification/internal/record"
	"github.com/iksnae/code-certification/internal/workspace"
	"github.com/spf13/cobra"
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

func bindCertifyFlags() {
	certifyCmd.Flags().String("path", "", "Path to repository (default: current directory)")
	certifyCmd.Flags().Bool("skip-agent", false, "Skip agent-assisted review")
	certifyCmd.Flags().Int("batch", 0, "Max units to process per run (0=all)")
	certifyCmd.Flags().Bool("reset-queue", false, "Rebuild queue from index")
	certifyCmd.Flags().StringSlice("target", nil, "Target specific paths/directories (can specify multiple)")
	certifyCmd.Flags().String("diff-base", "", "Only certify files changed since this git ref")
}

// certifyFlags holds flag values resolved from cobra for the certify command.
type certifyFlags struct {
	path       string
	skipAgent  bool
	batch      int
	resetQueue bool
	target     []string
	diffBase   string
}

func getCertifyFlags(cmd *cobra.Command) certifyFlags {
	path, _ := cmd.Flags().GetString("path")
	skipAgent, _ := cmd.Flags().GetBool("skip-agent")
	batch, _ := cmd.Flags().GetInt("batch")
	resetQueue, _ := cmd.Flags().GetBool("reset-queue")
	target, _ := cmd.Flags().GetStringSlice("target")
	diffBase, _ := cmd.Flags().GetString("diff-base")
	return certifyFlags{path: path, skipAgent: skipAgent, batch: batch, resetQueue: resetQueue, target: target, diffBase: diffBase}
}

// certifyContext holds all loaded state for a certification run.
type certifyContext struct {
	root      string
	certDir   string
	cfg       domain.Config
	unitMap   map[string]domain.Unit
	wq        *queue.Queue
	queuePath string
	certifier *engine.Certifier
	repoEv    []domain.Evidence
}

func runCertify(cmd *cobra.Command, args []string) error {
	flags := getCertifyFlags(cmd)
	wsMode, _ := cmd.Flags().GetBool("workspace")
	if wsMode {
		return runWorkspaceCertify(cmd, flags)
	}

	ctx, err := loadCertifyContext(flags)
	if err != nil {
		return err
	}

	remaining := ctx.printQueueStatus()
	if remaining == 0 {
		ctx.wq.Save(ctx.queuePath)
		return nil
	}

	// Generate run ID and set policy versions
	startedAt := time.Now()
	runID := domain.GenerateRunID(startedAt)
	ctx.certifier.RunID = runID
	ctx.certifier.PolicyVersions = policyVersions(ctx.certifier.Matcher)

	ctx.certifier.Agent = setupAgent(ctx.cfg, flags.skipAgent)

	fmt.Println("  Collecting repo-level evidence...")
	ctx.repoEv = ctx.certifier.CollectRepoEvidence()
	fmt.Printf("  Collected %d repo-level evidence items\n", len(ctx.repoEv))
	fmt.Println()

	certified, observations, failed, processed := ctx.processQueue(cmd, remaining, flags.batch)

	ctx.printSummary(certified, observations, failed, processed)

	now := time.Now()
	repo := detectRepoName(ctx.root)
	commit := detectCommit(ctx.root)
	if err := engine.SaveReportArtifactsFromStore(ctx.certDir, ctx.certifier.Store, repo, commit, now); err == nil {
		if n, _ := os.ReadDir(filepath.Join(ctx.certDir, "reports")); len(n) > 0 {
			fmt.Printf("✓ Unit certificates written to %s\n", filepath.Join(ctx.certDir, "reports"))
		}
	}

	// Persist run record
	run := buildCertificationRun(runID, startedAt, commit, ctx.certifier.PolicyVersions, certified+observations, failed, processed, ctx.certifier.Store)
	if err := ctx.certifier.Store.AppendRun(run); err != nil {
		fmt.Fprintf(os.Stderr, "warning: saving run record: %v\n", err)
	}

	// Save state snapshot for git tracking
	snapshotPath := filepath.Join(ctx.certDir, "state.json")
	if err := ctx.certifier.Store.SaveSnapshot(snapshotPath, commit); err != nil {
		fmt.Fprintf(os.Stderr, "warning: saving state snapshot: %v\n", err)
	}

	ctx.wq.Save(ctx.queuePath)

	if ctx.cfg.Mode == domain.ModeEnforcing && failed > 0 {
		return fmt.Errorf("%d units failed certification in enforcing mode", failed)
	}
	return nil
}

func loadCertifyContext(flags certifyFlags) (*certifyContext, error) {
	root := flags.path
	if root == "" {
		var err error
		root, err = os.Getwd()
		if err != nil {
			return nil, fmt.Errorf("getting working directory: %w", err)
		}
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

	wq := loadQueue(filepath.Join(certDir, "queue.json"), flags.resetQueue)
	units := filterUnits(root, idx, flags.diffBase, flags.target)

	unitMap := make(map[string]domain.Unit, len(units))
	for _, u := range units {
		unitMap[u.ID.String()] = u
		wq.Enqueue(u.ID.String(), u.ID.Path())
	}

	snapshotPath := filepath.Join(certDir, "state.json")
	store := record.NewStoreWithSnapshot(filepath.Join(certDir, "records"), snapshotPath)

	certifier := &engine.Certifier{
		Root:      root,
		Store:     store,
		Matcher:   config.NewPolicyMatcher(packs),
		Overrides: overrides,
		ExpiryCfg: cfg.Expiry,
	}

	return &certifyContext{
		root: root, certDir: certDir, cfg: cfg,
		unitMap: unitMap,
		wq:      wq, queuePath: filepath.Join(certDir, "queue.json"),
		certifier: certifier,
	}, nil
}

func loadQueue(queuePath string, resetQueue bool) *queue.Queue {
	if resetQueue {
		return queue.New()
	}
	wq, err := queue.Load(queuePath)
	if err != nil {
		return queue.New()
	}
	return wq
}

func filterUnits(root string, idx *discovery.Index, diffBase string, target []string) []domain.Unit {
	units := idx.Units()
	if diffBase != "" {
		changedFiles, err := discovery.ChangedFiles(root, diffBase, "HEAD")
		if err != nil {
			fmt.Fprintf(os.Stderr, "warning: git diff failed: %v — certifying all units\n", err)
		} else {
			units = discovery.FilterChanged(units, changedFiles)
			fmt.Printf("  Changed files since %s: %d → %d units\n", diffBase, len(idx.Units()), len(units))
		}
	}
	if len(target) > 0 {
		units = discovery.FilterByPaths(units, target)
		fmt.Printf("  Targeting %v: %d units\n", target, len(units))
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

func (c *certifyContext) processQueue(cmd *cobra.Command, remaining int, batch int) (certified, observations, failed, processed int) {
	now := time.Now()
	batchSize := remaining
	if batch > 0 && batch < batchSize {
		batchSize = batch
	}

	startTime := time.Now()

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

		result, err := c.certifier.Certify(cmd.Context(), unit, c.repoEv, now)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\n  ✗ %s — error: %v\n", unit.ID, err)
			c.wq.Fail(item.UnitID, err.Error())
			failed++
			c.wq.Save(c.queuePath)
			continue
		}

		// Update queue based on agent result
		agentTag := ""
		if result.AgentReview != nil {
			model := ""
			if len(result.AgentReview.ModelsUsed) > 0 {
				model = result.AgentReview.ModelsUsed[0]
			}
			if result.AgentReview.Reviewed {
				c.wq.Complete(item.UnitID, model)
				agentTag = fmt.Sprintf(" 🤖 %s", model)
			} else if result.AgentReview.Prescreened {
				c.wq.Complete(item.UnitID, model)
				agentTag = " 🤖 prescreened"
			} else {
				c.wq.Skip(item.UnitID, "prescreen: no review needed")
				agentTag = " 🤖 skipped"
			}
		} else {
			c.wq.Complete(item.UnitID, "")
		}

		switch {
		case result.Record.Status == domain.StatusCertified:
			certified++
		case result.Record.Status == domain.StatusCertifiedWithObservations || result.Record.Status == domain.StatusExempt:
			observations++
		default:
			failed++
		}

		// Per-unit progress line
		grade := result.Record.Grade.String()
		emoji := gradeEmoji(grade)
		elapsed := time.Since(startTime)
		rate := float64(processed) / elapsed.Seconds()
		eta := time.Duration(float64(batchSize-processed)/rate) * time.Second

		obsTag := ""
		if len(result.Record.Observations) > 0 {
			obsTag = fmt.Sprintf(" ⚠ %d obs", len(result.Record.Observations))
		}
		failTag := ""
		if !result.Record.Status.IsPassing() {
			failTag = " ✗"
		}

		fmt.Printf("  [%d/%d] %s %-2s  %-55s %5.1f%%%s%s%s  (%.1f/s, ~%s)\n",
			processed, batchSize,
			emoji, grade,
			unit.ID,
			result.Record.Score*100,
			agentTag, obsTag, failTag,
			rate, formatETA(eta))

		c.wq.Save(c.queuePath)
	}
	return
}

func formatETA(d time.Duration) string {
	if d < time.Second {
		return "<1s"
	}
	if d < time.Minute {
		return fmt.Sprintf("%ds", int(d.Seconds()))
	}
	return fmt.Sprintf("%dm%ds", int(d.Minutes()), int(d.Seconds())%60)
}

func gradeEmoji(g string) string {
	switch g {
	case "A", "A-":
		return "🟢"
	case "B+", "B":
		return "🟢"
	case "C":
		return "🟡"
	case "D":
		return "🟠"
	case "F":
		return "🔴"
	default:
		return "⚪"
	}
}

func (c *certifyContext) printSummary(certified, observations, failed, processed int) {
	total := certified + observations + failed
	fmt.Println()
	if c.certifier.Agent != nil {
		filesReviewed, totalFiles, tokens := c.certifier.Agent.Stats()
		fmt.Printf("  🤖 Agent: %d/%d files reviewed, %d tokens used\n", filesReviewed, totalFiles, tokens)
	}

	fmt.Printf("  Processed %d units this run\n", processed)

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

func runWorkspaceCertify(cmd *cobra.Command, flags certifyFlags) error {
	root := flags.path
	if root == "" {
		var err error
		root, err = os.Getwd()
		if err != nil {
			return fmt.Errorf("getting working directory: %w", err)
		}
	}

	subs, err := workspace.DiscoverSubmodules(root)
	if err != nil {
		return fmt.Errorf("discovering submodules: %w", err)
	}

	configured := workspace.ConfiguredSubmodules(subs)
	if len(configured) == 0 {
		return fmt.Errorf("no configured submodules found — run 'certify init --workspace' first")
	}

	fmt.Printf("🔍 Workspace certify: %d submodule(s)\n\n", len(configured))

	// Forward relevant flags to subcommand
	for _, s := range configured {
		fmt.Printf("═══ %s ═══\n", s.Path)
		subPath := filepath.Join(root, s.Path)
		args := []string{"certify", "--path", subPath}
		if flags.skipAgent {
			args = append(args, "--skip-agent")
		}
		if flags.batch > 0 {
			args = append(args, "--batch", fmt.Sprintf("%d", flags.batch))
		}
		if flags.resetQueue {
			args = append(args, "--reset-queue")
		}
		if err := runSubcommand(args...); err != nil {
			fmt.Fprintf(os.Stderr, "  warning: certify failed for %s: %v\n", s.Path, err)
		}
		fmt.Println()
	}

	fmt.Println("✓ Workspace certification complete.")
	return nil
}

func defaultConfigObj() domain.Config {
	return domain.DefaultConfig()
}

func setupAgent(cfg domain.Config, skip bool) *agent.Coordinator {
	if skip {
		fmt.Println("  Agent review: skipped (--skip-agent)")
		return nil
	}
	if cfg.Agent.Enabled {
		return setupExplicitAgent(cfg)
	}
	if cfg.Agent.ExplicitlyDisabled {
		return nil
	}
	return setupConservativeAgent()
}

func setupExplicitAgent(cfg domain.Config) *agent.Coordinator {
	baseURL := cfg.Agent.Provider.BaseURL
	apiKey := ""
	isLocal := isLocalURL(baseURL)

	if cfg.Agent.Provider.APIKeyEnv != "" {
		apiKey = os.Getenv(cfg.Agent.Provider.APIKeyEnv)
		if apiKey == "" && !isLocal {
			fmt.Fprintf(os.Stderr, "  Agent review configured but %s not set — skipping\n", cfg.Agent.Provider.APIKeyEnv)
			return nil
		}
	} else if !isLocal {
		apiKey, _ = agent.DetectAPIKey() //nolint: second return is env var name, not error
		if apiKey == "" {
			fmt.Fprintf(os.Stderr, "  Agent review configured but no API key found — skipping\n")
			return nil
		}
	}

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
		strategy = agent.StrategyLocal
		tokenBudget = 0
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

func isLocalURL(u string) bool {
	return strings.Contains(u, "localhost") || strings.Contains(u, "127.0.0.1") || strings.Contains(u, "0.0.0.0")
}

func setupConservativeAgent() *agent.Coordinator {
	providers := agent.DetectProviders()
	if len(providers) == 0 {
		return nil
	}
	summary := agent.FormatProviderSummary(providers)
	fmt.Printf("  Agent review: auto-detected [%s] (conservative mode)\n", summary)
	return agent.NewConservativeCoordinator(providers)
}

// policyVersions extracts "name@version" strings from a policy matcher.
func policyVersions(matcher *policy.Matcher) []string {
	if matcher == nil {
		return nil
	}
	var vers []string
	for _, p := range matcher.Packs() {
		vers = append(vers, p.Name+"@"+p.Version)
	}
	return vers
}

// buildCertificationRun creates a CertificationRun from run results.
// It computes overall grade/score from all records currently in the store.
func buildCertificationRun(runID string, startedAt time.Time, commit string, policyVers []string, certified, failed, processed int, store *record.Store) domain.CertificationRun {
	run := domain.CertificationRun{
		ID:             runID,
		StartedAt:      startedAt,
		CompletedAt:    time.Now(),
		Commit:         commit,
		PolicyVersions: policyVers,
		UnitsProcessed: processed,
		UnitsCertified: certified,
		UnitsFailed:    failed,
	}

	// Compute overall grade/score from all records in store
	if store != nil {
		if records, err := store.ListAll(); err == nil && len(records) > 0 {
			var totalScore float64
			for _, r := range records {
				totalScore += r.Score
			}
			run.OverallScore = totalScore / float64(len(records))
			run.OverallGrade = domain.GradeFromScore(run.OverallScore).String()
		}
	}

	return run
}
