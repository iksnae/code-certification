package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/iksnae/code-certification/internal/config"
	"github.com/iksnae/code-certification/internal/domain"
)

func testdataPath(name string) string {
	return filepath.Join("..", "..", "testdata", "config", name)
}

func TestLoadConfig_Valid(t *testing.T) {
	cfg, err := config.LoadFile(testdataPath("valid.yml"))
	if err != nil {
		t.Fatalf("LoadFile(valid.yml) error: %v", err)
	}

	if cfg.Mode != domain.ModeEnforcing {
		t.Errorf("Mode = %v, want enforcing", cfg.Mode)
	}

	if len(cfg.Scope.Include) != 2 {
		t.Errorf("Scope.Include = %d, want 2", len(cfg.Scope.Include))
	}
	if len(cfg.Scope.Exclude) != 2 {
		t.Errorf("Scope.Exclude = %d, want 2", len(cfg.Scope.Exclude))
	}

	if !cfg.Agent.Enabled {
		t.Error("Agent.Enabled should be true")
	}
	if cfg.Agent.Provider.Type != "openrouter" {
		t.Errorf("Agent.Provider.Type = %q, want openrouter", cfg.Agent.Provider.Type)
	}
	if cfg.Agent.Models.Review != "qwen/qwen3-coder:free" {
		t.Errorf("Agent.Models.Review = %q", cfg.Agent.Models.Review)
	}
	if cfg.Agent.RateLimit.RequestsPerMinute != 20 {
		t.Errorf("RateLimit.RequestsPerMinute = %d, want 20", cfg.Agent.RateLimit.RequestsPerMinute)
	}

	if !cfg.Schedule.Nightly {
		t.Error("Schedule.Nightly should be true")
	}
	if cfg.Schedule.Sweep {
		t.Error("Schedule.Sweep should be false")
	}

	if cfg.Expiry.DefaultWindowDays != 90 {
		t.Errorf("Expiry.DefaultWindowDays = %d, want 90", cfg.Expiry.DefaultWindowDays)
	}

	if !cfg.Issues.Enabled {
		t.Error("Issues.Enabled should be true")
	}
	if cfg.Issues.Grouping != "directory" {
		t.Errorf("Issues.Grouping = %q, want directory", cfg.Issues.Grouping)
	}
}

func TestLoadConfig_Minimal(t *testing.T) {
	cfg, err := config.LoadFile(testdataPath("minimal.yml"))
	if err != nil {
		t.Fatalf("LoadFile(minimal.yml) error: %v", err)
	}

	if cfg.Mode != domain.ModeAdvisory {
		t.Errorf("Mode = %v, want advisory", cfg.Mode)
	}

	// Defaults should be applied
	if cfg.Expiry.DefaultWindowDays != 90 {
		t.Errorf("default Expiry.DefaultWindowDays = %d, want 90", cfg.Expiry.DefaultWindowDays)
	}
	if cfg.Expiry.MinWindowDays != 7 {
		t.Errorf("default Expiry.MinWindowDays = %d, want 7", cfg.Expiry.MinWindowDays)
	}
	if cfg.Expiry.MaxWindowDays != 365 {
		t.Errorf("default Expiry.MaxWindowDays = %d, want 365", cfg.Expiry.MaxWindowDays)
	}
	if cfg.Agent.Enabled {
		t.Error("agent should be disabled by default")
	}
}

func TestLoadConfig_Invalid(t *testing.T) {
	_, err := config.LoadFile(testdataPath("invalid.yml"))
	if err == nil {
		t.Fatal("LoadFile(invalid.yml) should return error for min > max")
	}
}

func TestLoadConfig_NotFound(t *testing.T) {
	_, err := config.LoadFile("/nonexistent/config.yml")
	if err == nil {
		t.Fatal("LoadFile(nonexistent) should return error")
	}
}

func TestLoadConfig_FromBytes(t *testing.T) {
	data := []byte(`
mode: advisory
scope:
  include:
    - "src/**"
`)
	cfg, err := config.Load(data)
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}
	if cfg.Mode != domain.ModeAdvisory {
		t.Errorf("Mode = %v, want advisory", cfg.Mode)
	}
	if len(cfg.Scope.Include) != 1 {
		t.Errorf("Scope.Include = %d, want 1", len(cfg.Scope.Include))
	}
}

func TestLoadConfig_EmptyFile(t *testing.T) {
	// Empty YAML should produce defaults
	cfg, err := config.Load([]byte(""))
	if err != nil {
		t.Fatalf("Load(empty) error: %v", err)
	}
	if cfg.Mode != domain.ModeAdvisory {
		t.Errorf("empty config mode = %v, want advisory", cfg.Mode)
	}
}

func TestConfigExplicitlyDisabled(t *testing.T) {
	data := []byte(`
mode: advisory
agent:
  enabled: false
`)
	cfg, err := config.Load(data)
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}
	if cfg.Agent.Enabled {
		t.Error("Agent.Enabled should be false")
	}
	if !cfg.Agent.ExplicitlyDisabled {
		t.Error("Agent.ExplicitlyDisabled should be true when enabled: false is set")
	}
}

func TestConfigNotExplicitlyDisabled(t *testing.T) {
	data := []byte(`
mode: advisory
`)
	cfg, err := config.Load(data)
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}
	if cfg.Agent.Enabled {
		t.Error("Agent.Enabled should be false by default")
	}
	if cfg.Agent.ExplicitlyDisabled {
		t.Error("Agent.ExplicitlyDisabled should be false when agent section is absent")
	}
}

func TestConfigExplicitlyEnabled(t *testing.T) {
	data := []byte(`
mode: advisory
agent:
  enabled: true
  provider:
    type: openrouter
    base_url: https://openrouter.ai/api/v1
    api_key_env: OPENROUTER_API_KEY
  models:
    prescreen: qwen/qwen3-coder:free
`)
	cfg, err := config.Load(data)
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}
	if !cfg.Agent.Enabled {
		t.Error("Agent.Enabled should be true")
	}
	if cfg.Agent.ExplicitlyDisabled {
		t.Error("Agent.ExplicitlyDisabled should be false when enabled: true")
	}
}

func TestLoadConfig_Dir(t *testing.T) {
	// Create a temp dir with config.yml
	dir := t.TempDir()
	data := []byte("mode: enforcing\n")
	if err := os.WriteFile(filepath.Join(dir, "config.yml"), data, 0644); err != nil {
		t.Fatal(err)
	}

	cfg, err := config.LoadFromDir(dir)
	if err != nil {
		t.Fatalf("LoadFromDir() error: %v", err)
	}
	if cfg.Mode != domain.ModeEnforcing {
		t.Errorf("Mode = %v, want enforcing", cfg.Mode)
	}
}
