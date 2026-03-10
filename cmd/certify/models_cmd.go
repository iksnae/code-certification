package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/iksnae/code-certification/internal/agent"
	"github.com/spf13/cobra"
)

var (
	modelsProviderURL string
	modelsAPIKeyEnv   string
)

var modelsCmd = &cobra.Command{
	Use:   "models",
	Short: "List available models from an AI provider",
	Long: `Query an OpenAI-compatible provider for available models.

Supports any provider: OpenRouter, Groq, Together, Fireworks, OpenAI,
Ollama, LM Studio, vLLM, or any custom OpenAI-compatible endpoint.

Examples:
  certify models                                              # auto-detect providers
  certify models --provider-url http://localhost:11434/v1      # Ollama (local)
  certify models --provider-url https://api.groq.com/openai/v1 --api-key-env GROQ_API_KEY`,
	RunE: runModels,
}

func bindModelsFlags() {
	modelsCmd.Flags().StringVar(&modelsProviderURL, "provider-url", "", "Provider API base URL (e.g., https://openrouter.ai/api/v1)")
	modelsCmd.Flags().StringVar(&modelsAPIKeyEnv, "api-key-env", "", "Environment variable containing the API key")
}

func runModels(cmd *cobra.Command, args []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if modelsProviderURL != "" {
		return listFromProvider(ctx, modelsProviderURL, modelsAPIKeyEnv)
	}

	// Auto-detect providers and list from the first available
	providers := agent.DetectProviders()
	if len(providers) == 0 {
		fmt.Fprintln(os.Stderr, "No providers detected. Use --provider-url to specify one.")
		fmt.Println("[]")
		return nil
	}

	for _, p := range providers {
		fmt.Fprintf(os.Stderr, "Provider: %s (%s)\n", p.Name, p.BaseURL)
		if err := listFromProvider(ctx, p.BaseURL, ""); err != nil {
			fmt.Fprintf(os.Stderr, "  error: %v\n", err)
			continue
		}
		return nil
	}
	return fmt.Errorf("no providers returned models")
}

func listFromProvider(ctx context.Context, baseURL, apiKeyEnv string) error {
	apiKey := ""
	if apiKeyEnv != "" {
		apiKey = os.Getenv(apiKeyEnv)
		if apiKey == "" {
			fmt.Fprintf(os.Stderr, "warning: %s not set\n", apiKeyEnv)
		}
	}

	models, err := agent.ListModels(ctx, baseURL, apiKey)
	if err != nil {
		return err
	}

	// JSON to stdout
	data, err := json.MarshalIndent(models, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(data))

	// Summary to stderr
	fmt.Fprintf(os.Stderr, "%d models available\n", len(models))
	return nil
}
