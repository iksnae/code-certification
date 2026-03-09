package discovery_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/iksnae/code-certification/internal/discovery"
	"github.com/iksnae/code-certification/internal/domain"
)

func TestIndex_SaveAndLoad(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "index.json")

	units := []domain.Unit{
		domain.NewUnit(domain.NewUnitID("go", "main.go", "main"), domain.UnitTypeFunction),
		domain.NewUnit(domain.NewUnitID("go", "main.go", "helper"), domain.UnitTypeFunction),
	}

	idx := discovery.NewIndex(units)
	if err := idx.Save(path); err != nil {
		t.Fatalf("Save() error: %v", err)
	}

	loaded, err := discovery.LoadIndex(path)
	if err != nil {
		t.Fatalf("LoadIndex() error: %v", err)
	}

	if len(loaded.Units()) != 2 {
		t.Errorf("loaded %d units, want 2", len(loaded.Units()))
	}
	if loaded.Units()[0].ID.String() != units[0].ID.String() {
		t.Errorf("first unit = %s, want %s", loaded.Units()[0].ID, units[0].ID)
	}
}

func TestIndex_LoadNotFound(t *testing.T) {
	_, err := discovery.LoadIndex("/nonexistent/index.json")
	if err == nil {
		t.Fatal("LoadIndex(nonexistent) should return error")
	}
}

func TestIndex_Empty(t *testing.T) {
	idx := discovery.NewIndex(nil)
	if len(idx.Units()) != 0 {
		t.Errorf("empty index should have 0 units, got %d", len(idx.Units()))
	}

	dir := t.TempDir()
	path := filepath.Join(dir, "index.json")
	if err := idx.Save(path); err != nil {
		t.Fatal(err)
	}

	loaded, err := discovery.LoadIndex(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(loaded.Units()) != 0 {
		t.Errorf("loaded empty index should have 0 units")
	}
}

func TestDiff_AddedUnits(t *testing.T) {
	old := discovery.NewIndex([]domain.Unit{
		domain.NewUnit(domain.NewUnitID("go", "main.go", "main"), domain.UnitTypeFunction),
	})
	new_ := discovery.NewIndex([]domain.Unit{
		domain.NewUnit(domain.NewUnitID("go", "main.go", "main"), domain.UnitTypeFunction),
		domain.NewUnit(domain.NewUnitID("go", "main.go", "helper"), domain.UnitTypeFunction),
	})

	diff := discovery.Diff(old, new_)
	if len(diff.Added) != 1 {
		t.Errorf("Added = %d, want 1", len(diff.Added))
	}
	if len(diff.Removed) != 0 {
		t.Errorf("Removed = %d, want 0", len(diff.Removed))
	}
	if len(diff.Unchanged) != 1 {
		t.Errorf("Unchanged = %d, want 1", len(diff.Unchanged))
	}
}

func TestDiff_RemovedUnits(t *testing.T) {
	old := discovery.NewIndex([]domain.Unit{
		domain.NewUnit(domain.NewUnitID("go", "main.go", "main"), domain.UnitTypeFunction),
		domain.NewUnit(domain.NewUnitID("go", "main.go", "helper"), domain.UnitTypeFunction),
	})
	new_ := discovery.NewIndex([]domain.Unit{
		domain.NewUnit(domain.NewUnitID("go", "main.go", "main"), domain.UnitTypeFunction),
	})

	diff := discovery.Diff(old, new_)
	if len(diff.Added) != 0 {
		t.Errorf("Added = %d, want 0", len(diff.Added))
	}
	if len(diff.Removed) != 1 {
		t.Errorf("Removed = %d, want 1", len(diff.Removed))
	}
}

func TestDiff_BothEmpty(t *testing.T) {
	old := discovery.NewIndex(nil)
	new_ := discovery.NewIndex(nil)

	diff := discovery.Diff(old, new_)
	if len(diff.Added) != 0 || len(diff.Removed) != 0 || len(diff.Unchanged) != 0 {
		t.Error("diff of two empty indexes should be all zeros")
	}
}

func TestIndex_Roundtrip_PreservesOrder(t *testing.T) {
	units := []domain.Unit{
		domain.NewUnit(domain.NewUnitID("go", "z.go", "Z"), domain.UnitTypeFunction),
		domain.NewUnit(domain.NewUnitID("go", "a.go", "A"), domain.UnitTypeFunction),
	}

	dir := t.TempDir()
	path := filepath.Join(dir, "index.json")

	idx := discovery.NewIndex(units)
	_ = idx.Save(path)

	loaded, _ := discovery.LoadIndex(path)

	// Order should be preserved
	for i, u := range loaded.Units() {
		if u.ID.String() != units[i].ID.String() {
			t.Errorf("unit[%d] = %s, want %s", i, u.ID, units[i].ID)
		}
	}
}

func TestIndex_FileCreatesDirectory(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "sub", "dir", "index.json")

	idx := discovery.NewIndex(nil)
	if err := idx.Save(path); err != nil {
		t.Fatalf("Save() should create directories: %v", err)
	}

	if _, err := os.Stat(path); err != nil {
		t.Fatalf("file should exist: %v", err)
	}
}
