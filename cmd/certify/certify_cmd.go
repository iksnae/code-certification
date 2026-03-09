package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/code-certification/certify/internal/agent"
	"github.com/code-certification/certify/internal/config"
	"github.com/code-certification/certify/internal/discovery"
	"github.com/code-certification/certify/internal/domain"
	"github.com/code-certification/certify/internal/engine"
	"github.com/code-certification/certify/internal/evidence"
	"github.com/code-certification/certify/internal/override"
	"github.com/code-certification/certify/internal/queue"
	"github.com/code-certification/certify/internal/record"
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

func runCertify(cmd *cobra.Command, args []string) error {
	root := certifyPath
	if root == "" {
		root, _ = os.Getwd()
	}
	certDir := filepath.Join(root, ".certification")

	// Load config
	cfg, err := config.LoadFromDir(certDir)
	if err != nil {
		cfg = defaultConfigObj()
	}

	// Load index
	indexPath := filepath.Join(certDir, "index.json")
	idx, err := discovery.LoadIndex(indexPath)
	if err != nil {
		return fmt.Errorf("loading index (run 'certify scan' first): %w", err)
	}

	// Load policies
	policyDir := filepath.Join(certDir, "policies")
	packs, err := config.LoadPolicyPacks(policyDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: loading policies: %v\n", err)
	}

	// Load overrides
	overrideDir := filepath.Join(certDir, "overrides")
	overrides, err := override.LoadDir(overrideDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: loading overrides: %v\n", err)
	}

	// --- Work Queue ---
	queuePath := filepath.Join(certDir, "queue.json")
	var wq *queue.Queue

	if certifyResetQueue {
		wq = queue.New()
	} else {
		wq, err = queue.Load(queuePath)
		if err != nil {
			wq = queue.New() // First run
		}
	}

	// Filter units by --target or --diff-base
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

	// Populate queue from index (only adds new units)
	for _, u := range units {
		wq.Enqueue(u.ID.String(), u.ID.Path())
	}

	stats := wq.Stats()
	remaining := stats.Pending + stats.InProgress
	if remaining == 0 {
		fmt.Printf("  Queue complete: %d/%d processed (%d skipped)\n",
			stats.Completed, stats.Total, stats.Skipped)
		fmt.Println("  Use --reset-queue to re-process all units.")
		wq.Save(queuePath)
		return nil
	}
	fmt.Printf("  Queue: %d pending, %d completed, %d skipped, %d failed of %d total\n",
		remaining, stats.Completed, stats.Skipped, stats.Failed, stats.Total)

	// --- Agent Setup ---
	var coordinator *agent.Coordinator
	if cfg.Agent.Enabled && !certifySkipAgent {
		apiKey := os.Getenv(cfg.Agent.Provider.APIKeyEnv)
		if apiKey != "" {
			chain := agent.NewModelChain(
				cfg.Agent.Provider.BaseURL,
				apiKey,
				cfg.Agent.Provider.HTTPReferer,
				cfg.Agent.Provider.XTitle,
				[]string{
					cfg.Agent.Models.Prescreen,
					cfg.Agent.Models.Fallback,
					"qwen/qwen-turbo",
					"qwen/qwen3-coder:free",
					"mistralai/mistral-small-3.1-24b-instruct:free",
					"google/gemma-3-12b-it:free",
				},
			)
			cb := agent.NewCircuitBreaker(chain, 5)
			coordinator = agent.NewCoordinator(cb, agent.CoordinatorConfig{
				Models:      cfg.Agent.Models,
				Strategy:    agent.StrategyStandard,
				TokenBudget: 50000,
			})
			fmt.Println("  Agent review: enabled (model chain with fallback)")
		} else {
			fmt.Fprintf(os.Stderr, "  Agent review configured but %s not set — skipping\n", cfg.Agent.Provider.APIKeyEnv)
		}
	} else if certifySkipAgent {
		fmt.Println("  Agent review: skipped (--skip-agent)")
	}

	// --- Evidence Collection ---
	fmt.Println("  Collecting repo-level evidence...")
	executor := evidence.NewToolExecutor(root)
	repoEv := executor.CollectAll()
	fmt.Printf("  Collected %d repo-level evidence items\n", len(repoEv))

	// --- Build unit lookup ---
	unitMap := make(map[string]domain.Unit)
	for _, u := range units {
		unitMap[u.ID.String()] = u
	}

	// --- Process Queue ---
	store := record.NewStore(filepath.Join(certDir, "records"))
	now := time.Now()

	batchSize := remaining
	if certifyBatch > 0 && certifyBatch < batchSize {
		batchSize = certifyBatch
	}

	var certified, observations, failed, processed int

	for processed < batchSize {
		item, ok := wq.Next()
		if !ok {
			break
		}

		unit, exists := unitMap[item.UnitID]
		if !exists {
			wq.Skip(item.UnitID, "unit not in index")
			wq.Save(queuePath)
			continue
		}

		processed++
		if processed%20 == 0 || processed == batchSize {
			fmt.Printf("\r  Processing... %d/%d", processed, batchSize)
		}

		// Match policies
		matcher := config.NewPolicyMatcher(packs)
		matched := matcher.Match(unit)
		var rules []domain.PolicyRule
		for _, p := range matched {
			rules = append(rules, p.Rules...)
		}

		// Collect per-unit evidence
		ev := make([]domain.Evidence, len(repoEv))
		copy(ev, repoEv)

		srcPath := filepath.Join(root, unit.ID.Path())
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

		// Agent review
		if coordinator != nil {
			ctx, cancel := context.WithTimeout(cmd.Context(), 30*time.Second)
			result := coordinator.ReviewUnit(ctx, unit, srcCode, ev)
			cancel()

			if result.Reviewed {
				ev = append(ev, result.ToEvidence())
				model := ""
				if len(result.ModelsUsed) > 0 {
					model = result.ModelsUsed[0]
				}
				wq.Complete(item.UnitID, model)
			} else {
				wq.Skip(item.UnitID, "prescreen: no review needed")
			}
		} else {
			wq.Complete(item.UnitID, "")
		}

		// Certify
		rec := engine.CertifyUnit(unit, rules, ev, cfg.Expiry, now)
		if len(overrides) > 0 {
			rec = override.ApplyAll(rec, overrides)
		}
		if err := store.Save(rec); err != nil {
			fmt.Fprintf(os.Stderr, "\nwarning: saving record for %s: %v\n", unit.ID, err)
		}
		store.AppendHistory(rec) // best-effort history tracking

		switch {
		case rec.Status == domain.StatusCertified:
			certified++
		case rec.Status == domain.StatusCertifiedWithObservations || rec.Status == domain.StatusExempt:
			observations++
		default:
			failed++
		}

		// Save queue after each item (crash-safe)
		wq.Save(queuePath)
	}

	// --- Summary ---
	total := certified + observations + failed
	if coordinator != nil {
		filesReviewed, totalFiles, tokens := coordinator.Stats()
		fmt.Printf("\n  Agent: %d/%d files reviewed, %d tokens used\n", filesReviewed, totalFiles, tokens)
	}

	fmt.Printf("\n  Processed %d units this run\n", processed)

	finalStats := wq.Stats()
	pendingRemaining := finalStats.Pending
	if pendingRemaining > 0 {
		fmt.Printf("  %d units remaining — run again to continue\n", pendingRemaining)
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

	wq.Save(queuePath)

	if cfg.Mode == domain.ModeEnforcing && failed > 0 {
		return fmt.Errorf("%d units failed certification in enforcing mode", failed)
	}

	return nil
}

func defaultConfigObj() domain.Config {
	return domain.DefaultConfig()
}
