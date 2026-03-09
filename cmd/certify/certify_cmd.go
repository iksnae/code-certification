package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/code-certification/certify/internal/config"
	"github.com/code-certification/certify/internal/discovery"
	"github.com/code-certification/certify/internal/domain"
	"github.com/code-certification/certify/internal/engine"
	"github.com/code-certification/certify/internal/record"
	"github.com/spf13/cobra"
)

var certifyCmd = &cobra.Command{
	Use:   "certify",
	Short: "Evaluate and certify code units",
	RunE: func(cmd *cobra.Command, args []string) error {
		root, _ := os.Getwd()
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

		// Set up record store
		store := record.NewStore(filepath.Join(certDir, "records"))
		now := time.Now()

		var certified, failed int

		for _, unit := range idx.Units() {
			// Match policies to unit
			matcher := config.NewPolicyMatcher(packs)
			matched := matcher.Match(unit)

			// Collect rules from matched packs
			var rules []domain.PolicyRule
			for _, p := range matched {
				rules = append(rules, p.Rules...)
			}

			// TODO: collect real evidence — for now, pass empty
			var ev []domain.Evidence

			// Run certification pipeline
			rec := engine.CertifyUnit(unit, rules, ev, cfg.Expiry, now)

			// Save record
			if err := store.Save(rec); err != nil {
				fmt.Fprintf(os.Stderr, "warning: saving record for %s: %v\n", unit.ID, err)
			}

			if rec.Status.IsPassing() {
				certified++
			} else {
				failed++
			}
		}

		total := certified + failed
		fmt.Printf("✓ Certified %d/%d units (%d need attention)\n", certified, total, failed)

		if cfg.Mode == domain.ModeEnforcing && failed > 0 {
			return fmt.Errorf("%d units failed certification in enforcing mode", failed)
		}

		return nil
	},
}

func defaultConfigObj() domain.Config {
	return domain.DefaultConfig()
}
