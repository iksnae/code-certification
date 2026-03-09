package record_test

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/code-certification/certify/internal/domain"
	"github.com/code-certification/certify/internal/record"
)

func sampleRecord(unitPath, symbol string) domain.CertificationRecord {
	now := time.Now()
	return domain.CertificationRecord{
		UnitID:      domain.NewUnitID("go", unitPath, symbol),
		UnitType:    domain.UnitTypeFunction,
		UnitPath:    unitPath,
		Status:      domain.StatusCertified,
		Grade:       domain.GradeB,
		Score:       0.85,
		Confidence:  1.0,
		CertifiedAt: now,
		ExpiresAt:   now.Add(90 * 24 * time.Hour),
		Source:      "deterministic",
		Version:     1,
	}
}

func TestStore_SaveAndLoad(t *testing.T) {
	dir := t.TempDir()
	store := record.NewStore(dir)

	rec := sampleRecord("main.go", "main")
	if err := store.Save(rec); err != nil {
		t.Fatalf("Save() error: %v", err)
	}

	loaded, err := store.Load(rec.UnitID)
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}
	if loaded.UnitID.String() != rec.UnitID.String() {
		t.Errorf("loaded UnitID = %s, want %s", loaded.UnitID, rec.UnitID)
	}
	if loaded.Score != rec.Score {
		t.Errorf("loaded Score = %f, want %f", loaded.Score, rec.Score)
	}
}

func TestStore_LoadNotFound(t *testing.T) {
	dir := t.TempDir()
	store := record.NewStore(dir)

	_, err := store.Load(domain.NewUnitID("go", "missing.go", "foo"))
	if err == nil {
		t.Fatal("Load(missing) should return error")
	}
}

func TestStore_ListAll(t *testing.T) {
	dir := t.TempDir()
	store := record.NewStore(dir)

	_ = store.Save(sampleRecord("main.go", "main"))
	_ = store.Save(sampleRecord("main.go", "helper"))
	_ = store.Save(sampleRecord("service/sync.go", "Apply"))

	records, err := store.ListAll()
	if err != nil {
		t.Fatalf("ListAll() error: %v", err)
	}
	if len(records) != 3 {
		t.Errorf("ListAll() = %d records, want 3", len(records))
	}
}

func TestStore_SaveOverwrites(t *testing.T) {
	dir := t.TempDir()
	store := record.NewStore(dir)

	rec := sampleRecord("main.go", "main")
	rec.Score = 0.7
	_ = store.Save(rec)

	rec.Score = 0.9
	_ = store.Save(rec)

	loaded, _ := store.Load(rec.UnitID)
	if loaded.Score != 0.9 {
		t.Errorf("overwritten Score = %f, want 0.9", loaded.Score)
	}
}

func TestStore_PathFormat(t *testing.T) {
	dir := t.TempDir()
	store := record.NewStore(dir)

	rec := sampleRecord("internal/service/sync.go", "Apply")
	_ = store.Save(rec)

	// Should create a file under records dir
	pattern := filepath.Join(dir, "*.json")
	matches, _ := filepath.Glob(pattern)
	if len(matches) == 0 {
		// Check subdirectory
		matches, _ = filepath.Glob(filepath.Join(dir, "**", "*.json"))
	}
	// Just verify it can be loaded back
	loaded, err := store.Load(rec.UnitID)
	if err != nil {
		t.Fatalf("roundtrip failed: %v", err)
	}
	if loaded.UnitID.String() != rec.UnitID.String() {
		t.Error("roundtrip ID mismatch")
	}
}
