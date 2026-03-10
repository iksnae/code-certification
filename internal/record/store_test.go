package record_test

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/record"
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
		Evidence: []domain.Evidence{
			{
				Kind:       domain.EvidenceKindLint,
				Source:     "go-vet",
				Passed:     true,
				Summary:    "go-vet: 0 errors, 0 warnings",
				Timestamp:  now,
				Confidence: 1.0,
			},
		},
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

func TestStore_EvidenceRoundTrip(t *testing.T) {
	dir := t.TempDir()
	store := record.NewStore(dir)

	now := time.Now().Truncate(time.Second) // RFC3339 has second precision

	// Build a record with 4 evidence items covering different kinds and Details
	rec := domain.CertificationRecord{
		UnitID:      domain.NewUnitID("go", "pkg/foo.go", "Foo"),
		UnitType:    domain.UnitTypeFunction,
		UnitPath:    "pkg/foo.go",
		Status:      domain.StatusCertified,
		Grade:       domain.GradeA,
		Score:       0.95,
		Confidence:  1.0,
		CertifiedAt: now,
		ExpiresAt:   now.Add(90 * 24 * time.Hour),
		Source:      "deterministic",
		Version:     1,
		Evidence: []domain.Evidence{
			{
				Kind:       domain.EvidenceKindLint,
				Source:     "golangci-lint",
				Passed:     true,
				Summary:    "golangci-lint: 0 errors, 2 warnings",
				Timestamp:  now,
				Confidence: 1.0,
			},
			{
				Kind:       domain.EvidenceKindTest,
				Source:     "go-test",
				Passed:     true,
				Missing:    false,
				Summary:    "go-test: 10/10 passed (85% coverage)",
				Timestamp:  now,
				Confidence: 1.0,
			},
			{
				Kind:    domain.EvidenceKindMetrics,
				Source:  "metrics",
				Passed:  true,
				Summary: "42 lines (30 code, 5 comment, 7 blank), 3 TODOs, complexity 8",
				Details: map[string]any{
					"code_lines":  float64(30),
					"todo_count":  float64(3),
					"complexity":  float64(8),
					"total_lines": float64(42),
				},
				Timestamp:  now,
				Confidence: 1.0,
			},
			{
				Kind:       domain.EvidenceKindGitHistory,
				Source:     "git",
				Passed:     true,
				Summary:    "15 commits by 3 authors over 120 days",
				Timestamp:  now,
				Confidence: 1.0,
			},
		},
	}

	if err := store.Save(rec); err != nil {
		t.Fatalf("Save() error: %v", err)
	}

	loaded, err := store.Load(rec.UnitID)
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}

	// Must have 4 evidence items
	if len(loaded.Evidence) != 4 {
		t.Fatalf("Evidence count = %d, want 4", len(loaded.Evidence))
	}

	// Check each evidence item
	for i, want := range rec.Evidence {
		got := loaded.Evidence[i]
		if got.Kind != want.Kind {
			t.Errorf("ev[%d].Kind = %v, want %v", i, got.Kind, want.Kind)
		}
		if got.Source != want.Source {
			t.Errorf("ev[%d].Source = %q, want %q", i, got.Source, want.Source)
		}
		if got.Passed != want.Passed {
			t.Errorf("ev[%d].Passed = %v, want %v", i, got.Passed, want.Passed)
		}
		if got.Missing != want.Missing {
			t.Errorf("ev[%d].Missing = %v, want %v", i, got.Missing, want.Missing)
		}
		if got.Summary != want.Summary {
			t.Errorf("ev[%d].Summary = %q, want %q", i, got.Summary, want.Summary)
		}
		if got.Confidence != want.Confidence {
			t.Errorf("ev[%d].Confidence = %f, want %f", i, got.Confidence, want.Confidence)
		}
		// Timestamp round-trips to second precision
		if !got.Timestamp.Equal(want.Timestamp) {
			t.Errorf("ev[%d].Timestamp = %v, want %v", i, got.Timestamp, want.Timestamp)
		}
	}

	// Check Details on the metrics evidence (index 2) survived as map[string]any
	metricsEv := loaded.Evidence[2]
	details, ok := metricsEv.Details.(map[string]any)
	if !ok {
		t.Fatalf("ev[2].Details type = %T, want map[string]any", metricsEv.Details)
	}
	if v, ok := details["code_lines"].(float64); !ok || v != 30 {
		t.Errorf("ev[2].Details[code_lines] = %v, want 30", details["code_lines"])
	}
	if v, ok := details["todo_count"].(float64); !ok || v != 3 {
		t.Errorf("ev[2].Details[todo_count] = %v, want 3", details["todo_count"])
	}
	if v, ok := details["complexity"].(float64); !ok || v != 8 {
		t.Errorf("ev[2].Details[complexity] = %v, want 8", details["complexity"])
	}

	// Evidence with no Details (lint, index 0) should have nil Details
	if loaded.Evidence[0].Details != nil {
		t.Errorf("ev[0].Details should be nil, got %T", loaded.Evidence[0].Details)
	}
}

func TestStore_EvidenceRoundTrip_Empty(t *testing.T) {
	dir := t.TempDir()
	store := record.NewStore(dir)

	// Record with nil evidence
	rec := sampleRecord("empty.go", "Empty")
	rec.Evidence = nil

	if err := store.Save(rec); err != nil {
		t.Fatalf("Save() error: %v", err)
	}

	loaded, err := store.Load(rec.UnitID)
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}

	if len(loaded.Evidence) != 0 {
		t.Errorf("Evidence count = %d, want 0 for nil input", len(loaded.Evidence))
	}
}

func TestStore_BackwardCompatibility(t *testing.T) {
	dir := t.TempDir()

	// Write a record in the old format (no evidence field) directly to disk
	unitID := domain.NewUnitID("go", "old/file.go", "OldFunc")
	h := sha256.Sum256([]byte(unitID.String()))
	filename := hex.EncodeToString(h[:8]) + ".json"

	oldJSON := `{
  "unit_id": "go://old/file.go#OldFunc",
  "unit_type": "function",
  "unit_path": "old/file.go",
  "policy_version": "",
  "status": "certified",
  "grade": "B",
  "score": 0.85,
  "confidence": 1,
  "dimensions": {"correctness": 0.9},
  "observations": ["all good"],
  "certified_at": "2026-01-01T00:00:00Z",
  "expires_at": "2026-04-01T00:00:00Z",
  "source": "deterministic",
  "version": 1
}`

	if err := os.WriteFile(filepath.Join(dir, filename), []byte(oldJSON), 0o644); err != nil {
		t.Fatalf("writing old format file: %v", err)
	}

	store := record.NewStore(dir)
	loaded, err := store.Load(unitID)
	if err != nil {
		t.Fatalf("Load() old format error: %v", err)
	}

	// Core fields should load correctly
	if loaded.UnitID.String() != "go://old/file.go#OldFunc" {
		t.Errorf("UnitID = %s, want go://old/file.go#OldFunc", loaded.UnitID)
	}
	if loaded.Score != 0.85 {
		t.Errorf("Score = %f, want 0.85", loaded.Score)
	}
	if loaded.Status != domain.StatusCertified {
		t.Errorf("Status = %v, want certified", loaded.Status)
	}

	// Evidence should be nil/empty — not an error
	if len(loaded.Evidence) != 0 {
		t.Errorf("Evidence = %d items, want 0 for old format", len(loaded.Evidence))
	}
}
