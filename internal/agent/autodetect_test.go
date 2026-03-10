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
	providers := []DetectedProvider{
		{Name: "openrouter", BaseURL: "https://openrouter.ai/api/v1", APIKey: "sk-or-test", Models: ConservativeModels},
	}
	coord := NewConservativeCoordinator(providers)
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

func TestNewConservativeCoordinator_Empty(t *testing.T) {
	coord := NewConservativeCoordinator(nil)
	if coord != nil {
		t.Error("NewConservativeCoordinator(nil) should return nil")
	}
}

func TestNewConservativeCoordinator_MultiProvider(t *testing.T) {
	providers := []DetectedProvider{
		{Name: "groq", BaseURL: "https://api.groq.com/openai/v1", APIKey: "gsk-test", Models: GroqModels},
		{Name: "ollama", BaseURL: "http://localhost:11434/v1", Models: OllamaModels, Local: true},
	}
	coord := NewConservativeCoordinator(providers)
	if coord == nil {
		t.Fatal("NewConservativeCoordinator(multi) returned nil")
	}
	if coord.config.Strategy != StrategyQuick {
		t.Errorf("Strategy = %v, want StrategyQuick", coord.config.Strategy)
	}
}

func TestNewConservativeCoordinator_LocalOnly(t *testing.T) {
	providers := []DetectedProvider{
		{Name: "ollama", BaseURL: "http://localhost:11434/v1", Models: OllamaModels, Local: true},
	}
	coord := NewConservativeCoordinator(providers)
	if coord == nil {
		t.Fatal("NewConservativeCoordinator(local) returned nil")
	}
}

func TestFormatProviderSummary(t *testing.T) {
	providers := []DetectedProvider{
		{Name: "openrouter"},
		{Name: "groq"},
		{Name: "ollama", Local: true},
	}
	summary := FormatProviderSummary(providers)
	want := "openrouter → groq → ollama (local)"
	if summary != want {
		t.Errorf("FormatProviderSummary() = %q, want %q", summary, want)
	}
}

func TestHasAnyProvider_NoProviders(t *testing.T) {
	clearProviderEnvVars()
	defer clearProviderEnvVars()
	// HasAnyProvider with no env vars — may detect local servers on test machine
	// Just verify it doesn't panic
	_ = HasAnyProvider()
}

func TestDefaultModels_AllProvidersPopulated(t *testing.T) {
	// Regression: after removing init() functions, DefaultModels must still
	// contain all provider entries with non-nil model slices.
	expected := []string{"openrouter", "openai", "groq", "ollama", "lmstudio"}
	for _, name := range expected {
		models, ok := DefaultModels[name]
		if !ok {
			t.Errorf("DefaultModels missing provider %q", name)
			continue
		}
		if models == nil || len(models) == 0 {
			t.Errorf("DefaultModels[%q] is nil or empty", name)
		}
	}
	// openrouter should use ConservativeModels
	if DefaultModels["openrouter"][0] != ConservativeModels[0] {
		t.Errorf("DefaultModels[openrouter][0] = %q, want %q", DefaultModels["openrouter"][0], ConservativeModels[0])
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
