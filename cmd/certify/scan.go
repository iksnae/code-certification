package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/iksnae/code-certification/internal/config"
	"github.com/iksnae/code-certification/internal/discovery"
	"github.com/spf13/cobra"
)

var scanPath string

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Discover and index certifiable code units",
	RunE: func(cmd *cobra.Command, args []string) error {
		root := scanPath
		if root == "" {
			root, _ = os.Getwd()
		}
		certDir := filepath.Join(root, ".certification")

		// Load config
		cfg, err := config.LoadFromDir(certDir)
		if err != nil {
			cfg = defaultConfigObj()
		}

		// Detect languages to choose adapters
		langs := discovery.DetectLanguages(root)
		adapters := discovery.DetectedAdapters(langs)

		var allUnits []discovery.UnitList

		// Always run generic scanner for file-level coverage
		generic := discovery.NewGenericScanner(cfg.Scope.Include, cfg.Scope.Exclude)
		fileUnits, err := generic.Scan(root)
		if err != nil {
			return fmt.Errorf("generic scan: %w", err)
		}
		allUnits = append(allUnits, fileUnits)

		// Run language-specific adapters based on detection
		for _, adapter := range adapters {
			switch adapter {
			case "go":
				goAdapter := discovery.NewGoAdapter()
				goUnits, err := goAdapter.Scan(root)
				if err != nil {
					fmt.Fprintf(os.Stderr, "warning: Go adapter: %v\n", err)
				} else {
					allUnits = append(allUnits, goUnits)
					fmt.Printf("  Go adapter: %d symbols\n", len(goUnits))
				}
			case "ts":
				tsAdapter := discovery.NewTSAdapter()
				tsUnits, err := tsAdapter.Scan(root)
				if err != nil {
					fmt.Fprintf(os.Stderr, "warning: TS adapter: %v\n", err)
				} else {
					allUnits = append(allUnits, tsUnits)
					fmt.Printf("  TS adapter: %d symbols\n", len(tsUnits))
				}
			}
		}

		// Merge, deduplicate, and filter
		merged := discovery.Merge(allUnits...)
		merged = discovery.DeduplicateFileLevel(merged)
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

func init() {
	scanCmd.Flags().StringVar(&scanPath, "path", "", "Path to repository (default: current directory)")
}
