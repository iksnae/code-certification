package workspace

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseSubmoduleStatus(t *testing.T) {
	output := ` abc1234 lib/core (v1.0.0)
 def5678 lib/utils (v2.1.0)
 fed9876 services/api (heads/main)
`
	subs := ParseSubmoduleStatus(output)
	if len(subs) != 3 {
		t.Fatalf("expected 3 submodules, got %d", len(subs))
	}

	if subs[0].Path != "lib/core" {
		t.Errorf("subs[0].Path = %q, want lib/core", subs[0].Path)
	}
	if subs[0].Commit != "abc1234" {
		t.Errorf("subs[0].Commit = %q, want abc1234", subs[0].Commit)
	}
	if subs[0].Name != "lib/core" {
		t.Errorf("subs[0].Name = %q, want lib/core", subs[0].Name)
	}

	if subs[1].Path != "lib/utils" {
		t.Errorf("subs[1].Path = %q, want lib/utils", subs[1].Path)
	}
	if subs[2].Path != "services/api" {
		t.Errorf("subs[2].Path = %q, want services/api", subs[2].Path)
	}
}

func TestParseSubmoduleStatus_Empty(t *testing.T) {
	subs := ParseSubmoduleStatus("")
	if len(subs) != 0 {
		t.Errorf("expected 0 submodules, got %d", len(subs))
	}
}

func TestParseSubmoduleStatus_Prefixes(t *testing.T) {
	// + means checked out at different commit, - means not initialized
	output := `+abc1234 modified-mod (v1.0.0)
-def5678 uninitialized-mod (v2.0.0)
 fed9876 normal-mod (heads/main)
`
	subs := ParseSubmoduleStatus(output)
	if len(subs) != 3 {
		t.Fatalf("expected 3 submodules, got %d", len(subs))
	}
	if subs[0].Path != "modified-mod" {
		t.Errorf("subs[0].Path = %q, want modified-mod", subs[0].Path)
	}
	if subs[0].Commit != "abc1234" {
		t.Errorf("subs[0].Commit = %q, want abc1234", subs[0].Commit)
	}
	if subs[1].Path != "uninitialized-mod" {
		t.Errorf("subs[1].Path = %q, want uninitialized-mod", subs[1].Path)
	}
	if subs[2].Path != "normal-mod" {
		t.Errorf("subs[2].Path = %q, want normal-mod", subs[2].Path)
	}
}

func TestParseSubmoduleStatus_NoBranch(t *testing.T) {
	// Some submodules may not have branch info in parens
	output := ` abc1234 simple-mod
`
	subs := ParseSubmoduleStatus(output)
	if len(subs) != 1 {
		t.Fatalf("expected 1 submodule, got %d", len(subs))
	}
	if subs[0].Path != "simple-mod" {
		t.Errorf("subs[0].Path = %q, want simple-mod", subs[0].Path)
	}
}

func TestCheckHasConfig(t *testing.T) {
	root := t.TempDir()

	// Submodule with certify config
	configuredPath := filepath.Join(root, "configured", ".certification")
	if err := os.MkdirAll(configuredPath, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(configuredPath, "config.yml"), []byte("mode: advisory"), 0o644); err != nil {
		t.Fatal(err)
	}

	// Submodule without certify config
	if err := os.MkdirAll(filepath.Join(root, "unconfigured"), 0o755); err != nil {
		t.Fatal(err)
	}

	sub1 := Submodule{Path: "configured"}
	if !CheckHasConfig(root, sub1) {
		t.Error("expected configured submodule to have config")
	}

	sub2 := Submodule{Path: "unconfigured"}
	if CheckHasConfig(root, sub2) {
		t.Error("expected unconfigured submodule to not have config")
	}
}

func TestLoadSubmoduleCard_MissingState(t *testing.T) {
	root := t.TempDir()
	sub := Submodule{Name: "missing", Path: "missing"}

	card, err := LoadSubmoduleCard(root, sub)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if card != nil {
		t.Error("expected nil card for missing state")
	}
}

func TestLoadSubmoduleCard_ValidState(t *testing.T) {
	root := t.TempDir()
	subPath := filepath.Join(root, "myrepo", ".certification")
	if err := os.MkdirAll(subPath, 0o755); err != nil {
		t.Fatal(err)
	}

	// Write a minimal state.json
	stateJSON := `{
		"version": 1,
		"generated_at": "2026-03-10T12:00:00Z",
		"commit": "abc123",
		"unit_count": 3,
		"records": [
			{
				"unit_id": "go://a.go#Func1",
				"unit_type": "function",
				"unit_path": "a.go",
				"status": "certified",
				"grade": "A",
				"score": 0.95,
				"confidence": 1.0,
				"dimensions": {"correctness": 0.95},
				"certified_at": "2026-03-10T12:00:00Z",
				"expires_at": "2026-06-10T12:00:00Z"
			},
			{
				"unit_id": "go://b.go#Func2",
				"unit_type": "function",
				"unit_path": "b.go",
				"status": "certified",
				"grade": "B",
				"score": 0.82,
				"confidence": 1.0,
				"dimensions": {"correctness": 0.82},
				"certified_at": "2026-03-10T12:00:00Z",
				"expires_at": "2026-06-10T12:00:00Z"
			},
			{
				"unit_id": "go://c.go#Func3",
				"unit_type": "function",
				"unit_path": "c.go",
				"status": "decertified",
				"grade": "F",
				"score": 0.30,
				"confidence": 1.0,
				"dimensions": {"correctness": 0.30},
				"certified_at": "2026-03-10T12:00:00Z",
				"expires_at": "2026-06-10T12:00:00Z"
			}
		]
	}`
	if err := os.WriteFile(filepath.Join(subPath, "state.json"), []byte(stateJSON), 0o644); err != nil {
		t.Fatal(err)
	}

	sub := Submodule{Name: "myrepo", Path: "myrepo"}
	card, err := LoadSubmoduleCard(root, sub)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if card == nil {
		t.Fatal("expected non-nil card")
	}
	if card.TotalUnits != 3 {
		t.Errorf("card.TotalUnits = %d, want 3", card.TotalUnits)
	}
	if card.Passing != 2 {
		t.Errorf("card.Passing = %d, want 2", card.Passing)
	}
	if card.Failing != 1 {
		t.Errorf("card.Failing = %d, want 1", card.Failing)
	}
}
