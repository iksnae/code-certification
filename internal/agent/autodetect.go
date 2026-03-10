package agent

import (
	"os"
	"strings"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
)

const (
	// ConservativeTokenBudget is the max tokens for auto-detected conservative mode.
	ConservativeTokenBudget = 10_000

	// ConservativeCircuitThreshold is the circuit breaker threshold for conservative mode.
	ConservativeCircuitThreshold = 3
)

// AutoDetectEnvVars lists environment variables checked for cloud API keys, in priority order.
var AutoDetectEnvVars = []string{"OPENROUTER_API_KEY", "CERTIFY_API_KEY", "OPENAI_API_KEY", "GROQ_API_KEY"}

// ConservativeModels lists free-tier OpenRouter models in fallback order.
// These are default suggestions — users can choose any model via config or the extension.
var ConservativeModels = []string{
	"qwen/qwen3-coder:free",
	"qwen/qwen-2.5-coder-32b-instruct:free",
	"mistralai/mistral-small-3.1-24b-instruct:free",
	"meta-llama/llama-3.3-70b-instruct:free",
	"microsoft/phi-4:free",
}

// DetectAPIKey checks environment variables for an API key (cloud providers).
// Returns the key and the env var name it was found in.
// Returns empty strings if no key is found.
func DetectAPIKey() (key string, envVar string) {
	for _, v := range AutoDetectEnvVars {
		if k := os.Getenv(v); k != "" {
			return k, v
		}
	}
	return "", ""
}

// HasAnyProvider returns true if any provider (cloud or local) is available.
func HasAnyProvider() bool {
	providers := DetectProviders()
	return len(providers) > 0
}

// NewConservativeCoordinator builds a Coordinator from all detected providers.
// Uses StrategyQuick (prescreen only), conservative token budget, free-tier models.
// Cloud providers are tried first, then local providers as fallback.
func NewConservativeCoordinator(providers []DetectedProvider) *Coordinator {
	if len(providers) == 0 {
		return nil
	}

	var allProviders []Provider

	for _, dp := range providers {
		if dp.Local {
			// Local provider: no auth, short timeout
			for _, model := range dp.Models {
				p := NewLocalProvider(dp.BaseURL, dp.Name)
				allProviders = append(allProviders, &modelPinnedProvider{
					provider: p,
					model:    model,
				})
			}
		} else {
			// Cloud provider: model chain with longer backoff
			for _, model := range dp.Models {
				p := NewOpenRouterProvider(dp.BaseURL, dp.APIKey,
					"https://github.com/iksnae/code-certification", "Certify")
				p.name = dp.Name
				p.backoffBase = 2 * time.Second
				p.maxRetries = 2
				allProviders = append(allProviders, &modelPinnedProvider{
					provider: p,
					model:    model,
				})
			}
		}
	}

	fb := NewFallbackProvider(allProviders)
	cb := NewCircuitBreaker(fb, ConservativeCircuitThreshold)

	// Use first model from first provider as prescreen assignment
	prescreen := providers[0].Models[0]

	return NewCoordinator(cb, CoordinatorConfig{
		Models: domain.ModelAssignments{
			Prescreen: prescreen,
		},
		Strategy:    StrategyQuick,
		TokenBudget: ConservativeTokenBudget,
	})
}

// FormatProviderSummary returns a human-readable summary of detected providers.
func FormatProviderSummary(providers []DetectedProvider) string {
	var parts []string
	for _, p := range providers {
		label := p.Name
		if p.Local {
			label += " (local)"
		}
		parts = append(parts, label)
	}
	return strings.Join(parts, " → ")
}
