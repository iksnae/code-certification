package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/iksnae/code-certification/internal/agent"
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

		// AI-powered scan suggestions (conservative, optional)
		tryScanSuggestions(langs, len(merged), adapters)

		return nil
	},
}

func init() {
	scanCmd.Flags().StringVar(&scanPath, "path", "", "Path to repository (default: current directory)")
}

// tryScanSuggestions attempts AI-powered policy/scope suggestions.
// Silently does nothing if no provider is available or if the call fails.
func tryScanSuggestions(langs []discovery.LanguageInfo, unitCount int, adapters []string) {
	providers := agent.DetectProviders()
	if len(providers) == 0 {
		return
	}

	// Build a provider from the first detected source
	dp := providers[0]
	var provider agent.Provider
	if dp.Local {
		provider = agent.NewLocalProvider(dp.BaseURL, dp.Name)
	} else {
		provider = agent.NewModelChain(
			dp.BaseURL, dp.APIKey,
			"https://github.com/iksnae/code-certification", "Certify",
			dp.Models,
		)
	}

	var langNames []string
	for _, l := range langs {
		langNames = append(langNames, l.Name)
	}
	summary := agent.RepoSummary{
		Languages:    langNames,
		UnitCount:    unitCount,
		FilePatterns: adapters,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	result := agent.SuggestForRepo(ctx, provider, summary)
	if result.Suggestions != "" {
		fmt.Printf("\n💡 AI Suggestions (via %s):\n%s\n", dp.Name, result.Suggestions)
	}
}
