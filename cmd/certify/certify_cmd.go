package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
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

var certifyPath string

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

		// Set up agent reviewer if enabled
		var reviewer *agent.Reviewer
		if cfg.Agent.Enabled {
			apiKey := os.Getenv(cfg.Agent.Provider.APIKeyEnv)
			if apiKey != "" {
				provider := agent.NewOpenRouterProvider(
					cfg.Agent.Provider.BaseURL,
					apiKey,
					cfg.Agent.Provider.HTTPReferer,
					cfg.Agent.Provider.XTitle,
				)
				router := agent.NewRouter(cfg.Agent.Models)
				reviewer = agent.NewReviewer(provider, router)
				fmt.Println("  Agent review enabled")
			} else {
				fmt.Fprintf(os.Stderr, "  Agent review configured but %s not set — skipping\n", cfg.Agent.Provider.APIKeyEnv)
			}
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
			if srcData, readErr := os.ReadFile(srcPath); readErr == nil {
				metrics := evidence.ComputeMetrics(string(srcData))
				ev = append(ev, metrics.ToEvidence())
			}

			// Run agent review if available
			if reviewer != nil {
				srcCode := ""
				if data, err := os.ReadFile(srcPath); err == nil {
					srcCode = string(data)
				}
				result, err := reviewer.Review(context.Background(), agent.ReviewInput{
					Unit:       unit,
					SourceCode: srcCode,
					Evidence:   ev,
				})
				if err == nil && result.Reviewed {
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
		fmt.Printf("\n✓ Certified %d/%d units", certified+observations, total)
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
}

func defaultConfigObj() domain.Config {
	return domain.DefaultConfig()
}
