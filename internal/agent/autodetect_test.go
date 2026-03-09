package agent

import (
	"os"
	"testing"
)

func TestDetectAPIKey_NoVars(t *testing.T) {
	// Clear all relevant env vars
	for _, v := range AutoDetectEnvVars {
		os.Unsetenv(v)
	}
	key, envVar := DetectAPIKey()
	if key != "" {
		t.Errorf("DetectAPIKey() key = %q, want empty", key)
	}
	if envVar != "" {
		t.Errorf("DetectAPIKey() envVar = %q, want empty", envVar)
	}
}

func TestDetectAPIKey_OpenRouter(t *testing.T) {
	for _, v := range AutoDetectEnvVars {
		os.Unsetenv(v)
	}
	os.Setenv("OPENROUTER_API_KEY", "sk-or-test-key")
	defer os.Unsetenv("OPENROUTER_API_KEY")

	key, envVar := DetectAPIKey()
	if key != "sk-or-test-key" {
		t.Errorf("DetectAPIKey() key = %q, want sk-or-test-key", key)
	}
	if envVar != "OPENROUTER_API_KEY" {
		t.Errorf("DetectAPIKey() envVar = %q, want OPENROUTER_API_KEY", envVar)
	}
}

func TestDetectAPIKey_CertifyKey(t *testing.T) {
	for _, v := range AutoDetectEnvVars {
		os.Unsetenv(v)
	}
	os.Setenv("CERTIFY_API_KEY", "sk-certify-key")
	defer os.Unsetenv("CERTIFY_API_KEY")

	key, envVar := DetectAPIKey()
	if key != "sk-certify-key" {
		t.Errorf("DetectAPIKey() key = %q, want sk-certify-key", key)
	}
	if envVar != "CERTIFY_API_KEY" {
		t.Errorf("DetectAPIKey() envVar = %q, want CERTIFY_API_KEY", envVar)
	}
}

func TestDetectAPIKey_Priority(t *testing.T) {
	for _, v := range AutoDetectEnvVars {
		os.Unsetenv(v)
	}
	os.Setenv("OPENROUTER_API_KEY", "sk-or-primary")
	os.Setenv("CERTIFY_API_KEY", "sk-certify-secondary")
	defer os.Unsetenv("OPENROUTER_API_KEY")
	defer os.Unsetenv("CERTIFY_API_KEY")

	key, envVar := DetectAPIKey()
	if key != "sk-or-primary" {
		t.Errorf("DetectAPIKey() key = %q, want sk-or-primary (OPENROUTER_API_KEY should win)", key)
	}
	if envVar != "OPENROUTER_API_KEY" {
		t.Errorf("DetectAPIKey() envVar = %q, want OPENROUTER_API_KEY", envVar)
	}
}

func TestNewConservativeCoordinator(t *testing.T) {
	coord := NewConservativeCoordinator("sk-or-test-key")
	if coord == nil {
		t.Fatal("NewConservativeCoordinator() returned nil")
	}

	// Verify coordinator has correct config
	if coord.config.Strategy != StrategyQuick {
		t.Errorf("Strategy = %v, want StrategyQuick", coord.config.Strategy)
	}
	if coord.config.TokenBudget != ConservativeTokenBudget {
		t.Errorf("TokenBudget = %d, want %d", coord.config.TokenBudget, ConservativeTokenBudget)
	}
}

func TestConservativeModels(t *testing.T) {
	if len(ConservativeModels) == 0 {
		t.Fatal("ConservativeModels should not be empty")
	}
	// First model should be free-tier
	first := ConservativeModels[0]
	if first != "qwen/qwen3-coder:free" {
		t.Errorf("ConservativeModels[0] = %q, want qwen/qwen3-coder:free", first)
	}
}
