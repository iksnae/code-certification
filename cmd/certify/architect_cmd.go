package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/iksnae/code-certification/internal/agent"
	"github.com/iksnae/code-certification/internal/config"
	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/record"
	"github.com/iksnae/code-certification/internal/report"
	"github.com/spf13/cobra"
)

var (
	architectPath    string
	architectModel   string
	architectOutput  string
	architectPhase   int
	architectVerbose bool
)

var architectCmd = &cobra.Command{
	Use:   "architect",
	Short: "AI-powered architectural review",
	Long: `Perform a comprehensive, AI-driven architectural review of the project.

This command builds a complete architecture snapshot from certification data
(package graph, dependency analysis, hotspots, coupling) then runs a 6-phase
AI analysis producing comparative recommendations with current→projected deltas.

Output: .certification/ARCHITECT_REVIEW.md

Requires an AI provider (cloud or local). Run 'certify certify' first to
generate the certification data this command needs.

Examples:
  certify architect                    # full 6-phase review
  certify architect --model gpt-4o     # use specific model
  certify architect --phase 1          # run only architecture narration
  certify architect --verbose          # print full LLM responses`,
	RunE: runArchitect,
}

func init() {
	architectCmd.Flags().StringVar(&architectPath, "path", "", "Path to repository (default: current directory)")
	architectCmd.Flags().StringVar(&architectModel, "model", "", "Override model for all phases")
	architectCmd.Flags().StringVar(&architectOutput, "output", "", "Output file path (default: .certification/ARCHITECT_REVIEW.md)")
	architectCmd.Flags().IntVar(&architectPhase, "phase", 0, "Run specific phase (1-6, default: all)")
	architectCmd.Flags().BoolVar(&architectVerbose, "verbose", false, "Print full LLM responses")
	rootCmd.AddCommand(architectCmd)
}

func runArchitect(cmd *cobra.Command, args []string) error {
	root := architectPath
	if root == "" {
		root, _ = os.Getwd()
	}

	certDir := filepath.Join(root, ".certification")

	// Load config
	cfg, err := config.LoadFromDir(certDir)
	if err != nil {
		cfg = defaultConfigObj()
	}

	// Setup agent — required for architect
	provider, model := setupArchitectProvider(cfg)
	if provider == nil {
		return fmt.Errorf("architect requires an AI provider. Set OPENROUTER_API_KEY, configure agent in .certification/config.yml, or run a local model (Ollama)")
	}

	// Override model if specified
	if architectModel != "" {
		model = architectModel
	}

	// Load certification records
	store := record.NewStore(filepath.Join(certDir, "records"))
	records, err := store.ListAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "⚠️  No certification records found — snapshot will be minimal. Run 'certify certify' first.\n")
		records = nil
	}

	// Gather context
	fmt.Println("🏗  Gathering project context...")
	pc, err := agent.GatherContext(root, certDir, records)
	if err != nil {
		return fmt.Errorf("gathering context: %w", err)
	}
	pc.RepoName = detectRepoName(root)
	pc.CommitSHA = detectCommit(root)

	// Print header
	unitCount := 0
	if pc.Snapshot != nil {
		unitCount = pc.Snapshot.Metrics.TotalUnits
	}
	fmt.Printf("  Repository: %s\n", pc.RepoName)
	fmt.Printf("  Units: %d across %d packages\n", unitCount, len(pc.Snapshot.Packages))
	fmt.Printf("  Model: %s\n", model)
	if pc.Snapshot.Metrics.AvgScore > 0 {
		fmt.Printf("  Current Score: %.1f%%\n", pc.Snapshot.Metrics.AvgScore*100)
	}
	fmt.Println()

	// Build reviewer
	reviewer := &agent.ArchitectReviewer{
		Provider: provider,
		Model:    model,
		Verbose:  architectVerbose,
		OnPhaseStart: func(phase int, name string) {
			fmt.Printf("  [%d/6] 🔄 %s...\n", phase, name)
		},
		OnPhaseDone: func(phase int, name string, tokens int) {
			fmt.Printf("  [%d/6] ✅ %s (%d tokens)\n", phase, name, tokens)
		},
	}

	// Determine phases
	var phases []int
	if architectPhase > 0 {
		phases = []int{architectPhase}
	}

	// Run review
	start := time.Now()
	result, err := reviewer.Review(context.Background(), pc, phases)
	if err != nil {
		return fmt.Errorf("architect review: %w", err)
	}

	// Print verbose output
	if architectVerbose {
		for i, raw := range result.RawOutputs {
			if raw != "" {
				fmt.Printf("\n--- Phase %d Raw Output ---\n%s\n", i+1, raw)
			}
		}
	}

	// Format report
	output := report.FormatArchitectReport(result, pc)

	// Write output
	outputPath := architectOutput
	if outputPath == "" {
		outputPath = filepath.Join(certDir, "ARCHITECT_REVIEW.md")
	}
	os.MkdirAll(filepath.Dir(outputPath), 0755)
	if err := os.WriteFile(outputPath, []byte(output), 0644); err != nil {
		return fmt.Errorf("writing report: %w", err)
	}

	// Print summary
	duration := time.Since(start)
	fmt.Println()
	fmt.Printf("🏗  Architectural Review Complete\n")
	fmt.Printf("  Phases: %d/6 completed\n", result.PhasesComplete)
	fmt.Printf("  Tokens: %d\n", result.TotalTokens)
	fmt.Printf("  Duration: %s\n", duration.Round(time.Second))
	if result.Phase5 != nil {
		fmt.Printf("  Recommendations: %d\n", len(result.Phase5.Recommendations))
	}
	if len(result.Errors) > 0 {
		fmt.Printf("  ⚠️  %d phase(s) failed\n", len(result.Errors))
	}
	fmt.Printf("  Output: %s\n", outputPath)

	return nil
}

// setupArchitectProvider creates a Provider for the architect command.
// Unlike setupAgent which returns a Coordinator, this returns the raw Provider.
func setupArchitectProvider(cfg domain.Config) (agent.Provider, string) {
	// Architect needs generous timeouts — local 30B+ models can take several minutes per phase
	architectTimeout := 10 * time.Minute

	// Try explicit config first
	if cfg.Agent.Enabled {
		baseURL := cfg.Agent.Provider.BaseURL
		isLocal := isLocalURL(baseURL)

		apiKey := ""
		if cfg.Agent.Provider.APIKeyEnv != "" {
			apiKey = os.Getenv(cfg.Agent.Provider.APIKeyEnv)
		}
		if apiKey == "" && !isLocal {
			apiKey, _ = agent.DetectAPIKey()
		}
		if apiKey == "" && !isLocal {
			return nil, ""
		}

		model := cfg.Agent.Models.Review
		if model == "" {
			model = cfg.Agent.Models.Prescreen
		}
		if model == "" {
			model = "qwen3:latest"
		}

		if isLocal {
			p := agent.NewLocalProvider(baseURL, "local")
			p.SetTimeout(architectTimeout)
			return p, model
		}
		p := agent.NewOpenRouterProvider(
			baseURL, apiKey,
			cfg.Agent.Provider.HTTPReferer, cfg.Agent.Provider.XTitle,
		)
		p.SetTimeout(architectTimeout)
		return p, model
	}

	// Auto-detect providers
	providers := agent.DetectProviders()
	if len(providers) == 0 {
		return nil, ""
	}

	dp := providers[0]
	model := ""
	if len(dp.Models) > 0 {
		model = dp.Models[0]
	}

	if dp.Local {
		p := agent.NewLocalProvider(dp.BaseURL, dp.Name)
		p.SetTimeout(architectTimeout)
		return p, model
	}

	p := agent.NewOpenRouterProvider(
		dp.BaseURL, dp.APIKey,
		"https://github.com/iksnae/code-certification", "Certify",
	)
	p.SetTimeout(architectTimeout)
	return p, model
}
