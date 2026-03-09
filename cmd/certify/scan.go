package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/code-certification/certify/internal/config"
	"github.com/code-certification/certify/internal/discovery"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Discover and index certifiable code units",
	RunE: func(cmd *cobra.Command, args []string) error {
		root, _ := os.Getwd()
		certDir := filepath.Join(root, ".certification")

		// Load config
		cfg, err := config.LoadFromDir(certDir)
		if err != nil {
			// Use defaults if no config
			cfg = defaultConfigObj()
		}

		// Run scanners
		var allUnits []discovery.UnitList

		// Generic file scanner
		generic := discovery.NewGenericScanner(cfg.Scope.Include, cfg.Scope.Exclude)
		fileUnits, err := generic.Scan(root)
		if err != nil {
			return fmt.Errorf("generic scan: %w", err)
		}
		allUnits = append(allUnits, fileUnits)

		// Go adapter
		goAdapter := discovery.NewGoAdapter()
		goUnits, err := goAdapter.Scan(root)
		if err != nil {
			fmt.Fprintf(os.Stderr, "warning: Go adapter: %v\n", err)
		} else {
			allUnits = append(allUnits, goUnits)
		}

		// TS adapter
		tsAdapter := discovery.NewTSAdapter()
		tsUnits, err := tsAdapter.Scan(root)
		if err != nil {
			fmt.Fprintf(os.Stderr, "warning: TS adapter: %v\n", err)
		} else {
			allUnits = append(allUnits, tsUnits)
		}

		// Merge and deduplicate
		merged := discovery.Merge(allUnits...)
		idx := discovery.NewIndex(merged)

		// Save index
		indexPath := filepath.Join(certDir, "index.json")
		if err := idx.Save(indexPath); err != nil {
			return fmt.Errorf("saving index: %w", err)
		}

		fmt.Printf("✓ Discovered %d code units (saved to .certification/index.json)\n", len(merged))
		return nil
	},
}
