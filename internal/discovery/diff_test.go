package discovery_test

import (
	"testing"

	"github.com/code-certification/certify/internal/discovery"
	"github.com/code-certification/certify/internal/domain"
)

func TestFilterChanged(t *testing.T) {
	units := []domain.Unit{
		domain.NewUnit(domain.NewUnitID("go", "internal/engine/scorer.go", "Score"), domain.UnitTypeFunction),
		domain.NewUnit(domain.NewUnitID("go", "internal/agent/circuit.go", "Call"), domain.UnitTypeFunction),
		domain.NewUnit(domain.NewUnitID("go", "cmd/certify/main.go", "main"), domain.UnitTypeFunction),
	}

	changed := []string{"internal/engine/scorer.go", "README.md"}
	result := discovery.FilterChanged(units, changed)

	if len(result) != 1 {
		t.Fatalf("filtered = %d, want 1", len(result))
	}
	if result[0].ID.Path() != "internal/engine/scorer.go" {
		t.Errorf("got %s, want internal/engine/scorer.go", result[0].ID.Path())
	}
}

func TestFilterChanged_Empty(t *testing.T) {
	units := []domain.Unit{
		domain.NewUnit(domain.NewUnitID("go", "a.go", "A"), domain.UnitTypeFunction),
	}
	result := discovery.FilterChanged(units, nil)
	if len(result) != 0 {
		t.Errorf("filtered = %d, want 0", len(result))
	}
}

func TestFilterByPaths(t *testing.T) {
	units := []domain.Unit{
		domain.NewUnit(domain.NewUnitID("go", "internal/engine/scorer.go", "Score"), domain.UnitTypeFunction),
		domain.NewUnit(domain.NewUnitID("go", "internal/agent/circuit.go", "Call"), domain.UnitTypeFunction),
		domain.NewUnit(domain.NewUnitID("go", "cmd/certify/main.go", "main"), domain.UnitTypeFunction),
	}

	result := discovery.FilterByPaths(units, []string{"internal/engine"})
	if len(result) != 1 {
		t.Fatalf("filtered = %d, want 1", len(result))
	}

	result = discovery.FilterByPaths(units, []string{"internal/"})
	if len(result) != 2 {
		t.Fatalf("filtered = %d, want 2", len(result))
	}

	// Empty paths = no filter
	result = discovery.FilterByPaths(units, nil)
	if len(result) != 3 {
		t.Fatalf("empty filter = %d, want 3", len(result))
	}
}
