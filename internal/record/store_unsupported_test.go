package record_test

import (
	"testing"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/record"
)

// TestStore_UnsupportedRecordRoundTrip guards the persisted half of the
// fabricated-grade regression. Reports are rendered from stored records, not
// from the in-memory pipeline result, so an N/A verdict that does not survive
// save/load is not actually fixed: parseGrade's unknown-grade fallback turns
// "N/A" straight back into F, and a dropped Unsupported flag leaves nothing to
// distinguish an unassessable unit from an operator-exempted one.
func TestStore_UnsupportedRecordRoundTrip(t *testing.T) {
	store := record.NewStore(t.TempDir())

	rec := domain.CertificationRecord{
		UnitID:      domain.NewUnitID("swift", "Sources/App/Main.swift", "main"),
		UnitType:    domain.UnitTypeFunction,
		UnitPath:    "Sources/App/Main.swift",
		Status:      domain.StatusExempt,
		Grade:       domain.GradeNA,
		Score:       0,
		Confidence:  0,
		Unsupported: true,
		CertifiedAt: time.Now(),
		ExpiresAt:   time.Now().Add(24 * time.Hour),
		Source:      "deterministic",
		Version:     1,
	}

	if err := store.Save(rec); err != nil {
		t.Fatalf("Save() error: %v", err)
	}
	loaded, err := store.Load(rec.UnitID)
	if err != nil {
		t.Fatalf("Load() error: %v", err)
	}

	if loaded.Grade != domain.GradeNA {
		t.Errorf("loaded Grade = %v, want N/A — persistence must not re-fabricate a letter grade", loaded.Grade)
	}
	if !loaded.Unsupported {
		t.Error("loaded Unsupported = false, want true — the flag must survive the round trip")
	}
	if loaded.Status != domain.StatusExempt {
		t.Errorf("loaded Status = %v, want exempt", loaded.Status)
	}
}
