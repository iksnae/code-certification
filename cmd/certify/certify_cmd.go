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
	"github.com/iksnae/code-certification/internal/queue"
	"github.com/iksnae/code-certification/internal/record"
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
	unitMap   map[string]domain.Unit
	wq        *queue.Queue
	queuePath string
	certifier *engine.Certifier
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

	ctx.certifier.Agent = setupAgent(ctx.cfg, certifySkipAgent)

	fmt.Println("  Collecting repo-level evidence...")
	ctx.repoEv = ctx.certifier.CollectRepoEvidence()
	fmt.Printf("  Collected %d repo-level evidence items\n", len(ctx.repoEv))

	certified, observations, failed, processed := ctx.processQueue(cmd, remaining)

	ctx.printSummary(certified, observations, failed, processed)

	now := time.Now()
	repo := detectRepoName(ctx.root)
	commit := detectCommit(ctx.root)
	if err := engine.SaveReportArtifacts(ctx.certDir, ctx.certifier.Store, repo, commit, now); err == nil {
		if n, _ := os.ReadDir(filepath.Join(ctx.certDir, "reports")); len(n) > 0 {
			fmt.Printf("✓ %d unit report cards written to %s\n", len(n), filepath.Join(ctx.certDir, "reports"))
		}
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

		result, err := c.certifier.Certify(cmd.Context(), unit, c.repoEv, now)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\nwarning: certifying %s: %v\n", unit.ID, err)
			c.wq.Fail(item.UnitID, err.Error())
			failed++
			c.wq.Save(c.queuePath)
			continue
		}

		// Update queue based on agent result
		if result.AgentReview != nil {
			model := ""
			if len(result.AgentReview.ModelsUsed) > 0 {
				model = result.AgentReview.ModelsUsed[0]
			}
			if result.AgentReview.Reviewed || result.AgentReview.Prescreened {
				c.wq.Complete(item.UnitID, model)
			} else {
				c.wq.Skip(item.UnitID, "prescreen: no review needed")
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
		c.wq.Save(c.queuePath)
	}
	return
}

func (c *certifyContext) printSummary(certified, observations, failed, processed int) {
	total := certified + observations + failed
	if c.certifier.Agent != nil {
		filesReviewed, totalFiles, tokens := c.certifier.Agent.Stats()
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
		apiKey, _ = agent.DetectAPIKey()
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
