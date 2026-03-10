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
	"github.com/iksnae/code-certification/internal/workspace"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Discover and index certifiable code units",
	RunE: func(cmd *cobra.Command, args []string) error {
		root, _ := cmd.Flags().GetString("path")
		if root == "" {
			var err error
			root, err = os.Getwd()
			if err != nil {
				return fmt.Errorf("getting working directory: %w", err)
			}
		}

		wsMode, _ := cmd.Flags().GetBool("workspace")
		if wsMode {
			return runWorkspaceScan(root)
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

		// Run language-specific adapters based on detection (polymorphic dispatch)
		scanners := discovery.Scanners()
		for _, adapter := range adapters {
			s, ok := scanners[adapter]
			if !ok {
				continue
			}
			units, err := s.Scan(root)
			if err != nil {
				fmt.Fprintf(os.Stderr, "warning: %s adapter: %v\n", strings.ToUpper(adapter), err)
			} else {
				allUnits = append(allUnits, units)
				fmt.Printf("  %s adapter: %d symbols\n", strings.ToUpper(adapter), len(units))
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

func bindScanFlags() {
	scanCmd.Flags().String("path", "", "Path to repository (default: current directory)")
}

func runWorkspaceScan(root string) error {
	subs, err := workspace.DiscoverSubmodules(root)
	if err != nil {
		return fmt.Errorf("discovering submodules: %w", err)
	}

	configured := workspace.ConfiguredSubmodules(subs)
	if len(configured) == 0 {
		return fmt.Errorf("no configured submodules found — run 'certify init --workspace' first")
	}

	fmt.Printf("🔍 Workspace scan: %d submodule(s)\n\n", len(configured))

	for _, s := range configured {
		fmt.Printf("  → Scanning %s...\n", s.Path)
		subPath := filepath.Join(root, s.Path)
		if err := runSubcommand("scan", "--path", subPath); err != nil {
			fmt.Fprintf(os.Stderr, "    warning: scan failed for %s: %v\n", s.Path, err)
		}
	}

	fmt.Println("\n✓ Workspace scan complete.")
	return nil
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
