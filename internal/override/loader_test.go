package override_test

import (
	"path/filepath"
	"testing"

	"github.com/code-certification/certify/internal/domain"
	"github.com/code-certification/certify/internal/override"
)

func testdataPath(name string) string {
	return filepath.Join("..", "..", "testdata", "config", name)
}

func TestLoadOverrides(t *testing.T) {
	overrides, err := override.LoadFile(testdataPath("overrides.yml"))
	if err != nil {
		t.Fatalf("LoadFile() error: %v", err)
	}

	if len(overrides) != 4 {
		t.Fatalf("loaded %d overrides, want 4", len(overrides))
	}

	// Check first override
	o := overrides[0]
	if o.UnitID.String() != "go://legacy/old.go#Deprecated" {
		t.Errorf("override[0].UnitID = %s", o.UnitID)
	}
	if o.Action != domain.OverrideExempt {
		t.Errorf("override[0].Action = %v, want exempt", o.Action)
	}
	if o.Actor != "kmills" {
		t.Errorf("override[0].Actor = %q", o.Actor)
	}
}

func TestLoadOverrides_Actions(t *testing.T) {
	overrides, err := override.LoadFile(testdataPath("overrides.yml"))
	if err != nil {
		t.Fatal(err)
	}

	expected := []domain.OverrideAction{
		domain.OverrideExempt,
		domain.OverrideShortenWindow,
		domain.OverrideExtendWindow,
		domain.OverrideForceReview,
	}
	for i, want := range expected {
		if overrides[i].Action != want {
			t.Errorf("override[%d].Action = %v, want %v", i, overrides[i].Action, want)
		}
	}
}

func TestLoadOverrides_NotFound(t *testing.T) {
	_, err := override.LoadFile("/nonexistent/overrides.yml")
	if err == nil {
		t.Fatal("should error on missing file")
	}
}

func TestLoadOverrides_Empty(t *testing.T) {
	dir := t.TempDir()
	overrides, err := override.LoadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(overrides) != 0 {
		t.Errorf("empty dir should return 0 overrides, got %d", len(overrides))
	}
}
