package main

import (
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
	"github.com/code-certification/certify/internal/record"
	"github.com/spf13/cobra"
)

var (
	certifyPath      string
	certifySkipAgent bool
	certifyLimit     int
)

var certifyCmd = &cobra.Command{
	Use:   "certify",
	Short: "Evaluate and certify code units",
	RunE: func(cmd *cobra.Command, args []string) error {
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

		// Set up agent coordinator if enabled
		var coordinator *agent.Coordinator
		if cfg.Agent.Enabled && !certifySkipAgent {
			apiKey := os.Getenv(cfg.Agent.Provider.APIKeyEnv)
			if apiKey != "" {
				// Build model chain: tries each model in order on 429/402
				// All open-weight, fine-tune ready. Models credited in records.
				chain := agent.NewModelChain(
					cfg.Agent.Provider.BaseURL,
					apiKey,
					cfg.Agent.Provider.HTTPReferer,
					cfg.Agent.Provider.XTitle,
					[]string{
						cfg.Agent.Models.Prescreen, // primary: qwen3-coder (free, Apache 2.0)
						cfg.Agent.Models.Fallback,  // paid fallback: mistral-nemo ($0.002/run)
						"qwen/qwen-turbo",          // paid backup ($0.004/run)
						"qwen/qwen3-coder:free",    // retry free tier
						"mistralai/mistral-small-3.1-24b-instruct:free",
						"google/gemma-3-12b-it:free", // last resort
					},
				)
				// Wrap with circuit breaker: stop after 5 consecutive all-model failures
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

		// Set up record store
		store := record.NewStore(filepath.Join(certDir, "records"))
		now := time.Now()

		// Collect repo-level evidence once
		fmt.Println("  Collecting evidence...")
		executor := evidence.NewToolExecutor(root)
		repoEv := executor.CollectAll()
		fmt.Printf("  Collected %d repo-level evidence items\n", len(repoEv))

		var certified, observations, failed int
		units := idx.Units()
		if certifyLimit > 0 && certifyLimit < len(units) {
			units = units[:certifyLimit]
			fmt.Printf("  Limiting to %d units (--limit)\n", certifyLimit)
		}

		for i, unit := range units {
			if (i+1)%50 == 0 || i == len(units)-1 {
				fmt.Printf("\r  Certifying... %d/%d", i+1, len(units))
			}

			// Match policies to unit
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

			// Agent review via coordinator (handles dedup, budget, circuit breaking)
			if coordinator != nil {
				result := coordinator.ReviewUnit(cmd.Context(), unit, srcCode, ev)
				if result.Reviewed {
					ev = append(ev, result.ToEvidence())
				}
			}

			// Run certification pipeline
			rec := engine.CertifyUnit(unit, rules, ev, cfg.Expiry, now)

			// Apply overrides
			if len(overrides) > 0 {
				rec = override.ApplyAll(rec, overrides)
			}

			// Save record
			if err := store.Save(rec); err != nil {
				fmt.Fprintf(os.Stderr, "\nwarning: saving record for %s: %v\n", unit.ID, err)
			}

			switch {
			case rec.Status == domain.StatusCertified:
				certified++
			case rec.Status == domain.StatusCertifiedWithObservations || rec.Status == domain.StatusExempt:
				observations++
			default:
				failed++
			}
		}

		total := certified + observations + failed
		if coordinator != nil {
			filesReviewed, totalFiles, tokens := coordinator.Stats()
			fmt.Printf("\n  Agent: %d/%d files reviewed, %d tokens used\n", filesReviewed, totalFiles, tokens)
		}
		fmt.Printf("✓ Certified %d/%d units", certified+observations, total)
		if observations > 0 {
			fmt.Printf(" (%d with observations)", observations)
		}
		if failed > 0 {
			fmt.Printf(" (%d need attention)", failed)
		}
		fmt.Println()

		if cfg.Mode == domain.ModeEnforcing && failed > 0 {
			return fmt.Errorf("%d units failed certification in enforcing mode", failed)
		}

		return nil
	},
}

func init() {
	certifyCmd.Flags().StringVar(&certifyPath, "path", "", "Path to repository (default: current directory)")
	certifyCmd.Flags().BoolVar(&certifySkipAgent, "skip-agent", false, "Skip agent-assisted review")
	certifyCmd.Flags().IntVar(&certifyLimit, "limit", 0, "Limit number of units to certify (0=all)")
}

func defaultConfigObj() domain.Config {
	return domain.DefaultConfig()
}
