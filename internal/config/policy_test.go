package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/iksnae/code-certification/internal/config"
	"github.com/iksnae/code-certification/internal/domain"
)

func policiesPath(name string) string {
	return filepath.Join("..", "..", "testdata", "policies", name)
}

func TestLoadPolicyPack_GoStandard(t *testing.T) {
	pack, err := config.LoadPolicyPack(policiesPath("go-standard.yml"))
	if err != nil {
		t.Fatalf("LoadPolicyPack(go-standard.yml) error: %v", err)
	}

	if pack.Name != "go-standard" {
		t.Errorf("Name = %q, want go-standard", pack.Name)
	}
	if pack.Version != "1.0.0" {
		t.Errorf("Version = %q, want 1.0.0", pack.Version)
	}
	if pack.Language != "go" {
		t.Errorf("Language = %q, want go", pack.Language)
	}
	if pack.IsGlobal() {
		t.Error("go-standard should not be global")
	}
	if len(pack.Rules) != 15 {
		t.Fatalf("len(Rules) = %d, want 15", len(pack.Rules))
	}

	// Check first rule
	r := pack.Rules[0]
	if r.ID != "max-complexity" {
		t.Errorf("rule[0].ID = %q, want max-complexity", r.ID)
	}
	if r.Dimension != domain.DimMaintainability {
		t.Errorf("rule[0].Dimension = %v, want maintainability", r.Dimension)
	}
	if r.Severity != domain.SeverityError {
		t.Errorf("rule[0].Severity = %v, want error", r.Severity)
	}
	if r.Threshold != 20.0 {
		t.Errorf("rule[0].Threshold = %f, want 20", r.Threshold)
	}
}

func TestLoadPolicyPack_Security(t *testing.T) {
	pack, err := config.LoadPolicyPack(policiesPath("security.yml"))
	if err != nil {
		t.Fatalf("LoadPolicyPack(security.yml) error: %v", err)
	}

	if pack.Name != "security" {
		t.Errorf("Name = %q, want security", pack.Name)
	}
	// IsGlobal() only checks language — security pack has no language, so it's global
	if !pack.IsGlobal() {
		t.Error("security pack with empty language should be global")
	}
	if pack.Language != "" {
		t.Errorf("Language should be empty, got %q", pack.Language)
	}
	if len(pack.PathPatterns) != 3 {
		t.Errorf("PathPatterns = %d, want 3", len(pack.PathPatterns))
	}
	if len(pack.Rules) != 2 {
		t.Errorf("len(Rules) = %d, want 2", len(pack.Rules))
	}

	// Verify critical severity
	if pack.Rules[0].Severity != domain.SeverityCritical {
		t.Errorf("rule[0].Severity = %v, want critical", pack.Rules[0].Severity)
	}
}

func TestLoadPolicyPacks_Dir(t *testing.T) {
	dir := filepath.Join("..", "..", "testdata", "policies")
	packs, err := config.LoadPolicyPacks(dir)
	if err != nil {
		t.Fatalf("LoadPolicyPacks() error: %v", err)
	}

	if len(packs) != 2 {
		t.Fatalf("len(packs) = %d, want 2", len(packs))
	}

	// Should find both packs
	names := make(map[string]bool)
	for _, p := range packs {
		names[p.Name] = true
	}
	if !names["go-standard"] {
		t.Error("missing go-standard pack")
	}
	if !names["security"] {
		t.Error("missing security pack")
	}
}

func TestLoadPolicyPack_NotFound(t *testing.T) {
	_, err := config.LoadPolicyPack("/nonexistent/policy.yml")
	if err == nil {
		t.Fatal("LoadPolicyPack(nonexistent) should return error")
	}
}

func TestLoadPolicyPack_Invalid(t *testing.T) {
	dir := t.TempDir()
	// Missing name field
	data := []byte("version: '1.0.0'\nrules: []\n")
	path := filepath.Join(dir, "bad.yml")
	if err := os.WriteFile(path, data, 0644); err != nil {
		t.Fatal(err)
	}

	_, err := config.LoadPolicyPack(path)
	if err == nil {
		t.Fatal("LoadPolicyPack with missing name should return error")
	}
}
