package agent

import (
	"os"

	"github.com/iksnae/code-certification/internal/domain"
)

const (
	// ConservativeTokenBudget is the max tokens for auto-detected conservative mode.
	ConservativeTokenBudget = 10_000

	// ConservativeCircuitThreshold is the circuit breaker threshold for conservative mode.
	ConservativeCircuitThreshold = 3

	// conservativeBaseURL is the default OpenRouter API endpoint.
	conservativeBaseURL = "https://openrouter.ai/api/v1"
)

// AutoDetectEnvVars lists environment variables checked for API keys, in priority order.
var AutoDetectEnvVars = []string{"OPENROUTER_API_KEY", "CERTIFY_API_KEY"}

// ConservativeModels lists free-tier models used in conservative mode, in fallback order.
var ConservativeModels = []string{
	"qwen/qwen3-coder:free",
	"mistralai/mistral-small-3.1-24b-instruct:free",
	"google/gemma-3-12b-it:free",
}

// DetectAPIKey checks environment variables for an API key.
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

// NewConservativeCoordinator builds a Coordinator for auto-detected conservative mode.
// Uses StrategyQuick (prescreen only), conservative token budget, and free-tier models.
func NewConservativeCoordinator(apiKey string) *Coordinator {
	chain := NewModelChain(
		conservativeBaseURL, apiKey,
		"https://github.com/iksnae/code-certification",
		"Certify",
		ConservativeModels,
	)
	cb := NewCircuitBreaker(chain, ConservativeCircuitThreshold)

	return NewCoordinator(cb, CoordinatorConfig{
		Models: domain.ModelAssignments{
			Prescreen: ConservativeModels[0],
		},
		Strategy:    StrategyQuick,
		TokenBudget: ConservativeTokenBudget,
	})
}
