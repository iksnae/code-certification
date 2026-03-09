package domain_test

import (
	"testing"

	"github.com/iksnae/code-certification/internal/domain"
)

func TestConfig_Defaults(t *testing.T) {
	c := domain.DefaultConfig()

	if c.Mode != domain.ModeAdvisory {
		t.Errorf("default mode = %v, want advisory", c.Mode)
	}
	if c.Agent.Enabled {
		t.Error("agent should be disabled by default")
	}
}

func TestCertificationMode_String(t *testing.T) {
	tests := []struct {
		m    domain.CertificationMode
		want string
	}{
		{domain.ModeAdvisory, "advisory"},
		{domain.ModeEnforcing, "enforcing"},
	}
	for _, tt := range tests {
		if got := tt.m.String(); got != tt.want {
			t.Errorf("CertificationMode(%d).String() = %q, want %q", tt.m, got, tt.want)
		}
	}
}

func TestConfig_ScopePatterns(t *testing.T) {
	c := domain.Config{
		Scope: domain.ScopeConfig{
			Include: []string{"internal/**", "cmd/**"},
			Exclude: []string{"vendor/**", "testdata/**"},
		},
	}

	if len(c.Scope.Include) != 2 {
		t.Errorf("include patterns = %d, want 2", len(c.Scope.Include))
	}
	if len(c.Scope.Exclude) != 2 {
		t.Errorf("exclude patterns = %d, want 2", len(c.Scope.Exclude))
	}
}

func TestAgentConfig_ModelAssignments(t *testing.T) {
	ac := domain.AgentConfig{
		Enabled: true,
		Provider: domain.ProviderConfig{
			Type:      "openrouter",
			BaseURL:   "https://openrouter.ai/api/v1",
			APIKeyEnv: "OPENROUTER_API_KEY",
		},
		Models: domain.ModelAssignments{
			Prescreen:   "mistralai/mistral-small-3.1-24b-instruct:free",
			Review:      "qwen/qwen3-coder:free",
			Scoring:     "qwen/qwen3-next-80b-a3b-instruct:free",
			Decision:    "openai/gpt-oss-120b:free",
			Remediation: "qwen/qwen3-coder:free",
			Fallback:    "meta-llama/llama-3.3-70b-instruct:free",
		},
	}

	if ac.Provider.Type != "openrouter" {
		t.Errorf("provider type = %q, want openrouter", ac.Provider.Type)
	}
	if ac.Models.Fallback != "meta-llama/llama-3.3-70b-instruct:free" {
		t.Errorf("fallback model = %q", ac.Models.Fallback)
	}
}
