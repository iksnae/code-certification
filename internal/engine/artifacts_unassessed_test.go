package engine_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/engine"
	"github.com/iksnae/code-certification/internal/record"
)

// unassessedSampleRecords is a repo entirely in a language the engine cannot
// analyse: StatusExempt (IsPassing() == true) with the Unsupported flag set.
func unassessedSampleRecords(now time.Time) []domain.CertificationRecord {
	var recs []domain.CertificationRecord
	for _, path := range []string{"a.swift", "b.swift", "c.swift"} {
		recs = append(recs, domain.CertificationRecord{
			UnitID:      domain.NewUnitID("swift", path, ""),
			UnitType:    domain.UnitTypeFile,
			UnitPath:    path,
			Status:      domain.StatusExempt,
			Grade:       domain.GradeNA,
			Score:       0,
			Unsupported: true,
			CertifiedAt: now,
			ExpiresAt:   now.Add(90 * 24 * time.Hour),
			Source:      "deterministic",
			Version:     1,
		})
	}
	return recs
}

// TestSaveReportArtifacts_UnassessedBadgeOnDisk is the artifact-level assertion
// for #47. GenerateBadge returning a correct struct is not enough: badge.json
// is the file committed to the repo and served into the public README, so the
// test reads it back off disk exactly as a client would.
func TestSaveReportArtifacts_UnassessedBadgeOnDisk(t *testing.T) {
	certDir := t.TempDir()
	storeDir := filepath.Join(certDir, "records")
	store := record.NewStore(storeDir)

	now := time.Now()
	for _, rec := range unassessedSampleRecords(now) {
		if err := store.Save(rec); err != nil {
			t.Fatalf("Save() error: %v", err)
		}
	}

	if err := engine.SaveReportArtifactsFromStore(certDir, store, "test/repo", "abc123", now); err != nil {
		t.Fatalf("SaveReportArtifactsFromStore() error: %v", err)
	}

	data, err := os.ReadFile(filepath.Join(certDir, "badge.json"))
	if err != nil {
		t.Fatalf("badge.json not found: %v", err)
	}

	var badge struct {
		Message string `json:"message"`
		Color   string `json:"color"`
	}
	if err := json.Unmarshal(data, &badge); err != nil {
		t.Fatalf("badge.json is not valid JSON: %v", err)
	}

	if strings.Contains(badge.Message, "100%") {
		t.Errorf("badge.json message = %q — the public README badge claims a 100%% pass "+
			"rate for a repo the engine never opened", badge.Message)
	}
	if !strings.Contains(badge.Message, "not assessed") {
		t.Errorf("badge.json message = %q, want it to say the repo was not assessed", badge.Message)
	}
	if badge.Color != "9CA3AF" {
		t.Errorf("badge.json color = %q, want neutral 9CA3AF", badge.Color)
	}

	cardData, err := os.ReadFile(filepath.Join(certDir, "REPORT_CARD.md"))
	if err != nil {
		t.Fatalf("REPORT_CARD.md not found: %v", err)
	}
	if strings.Contains(string(cardData), "| Pass Rate | 100.0% |") {
		t.Errorf("REPORT_CARD.md claims a 100%% pass rate for an unassessed repo:\n%s", cardData)
	}
	if !strings.Contains(string(cardData), "| Passing | 0 |") {
		t.Errorf("REPORT_CARD.md should count no unit as passing, got:\n%s", cardData)
	}
	if !strings.Contains(string(cardData), "| Pass Rate | n/a |") {
		t.Errorf("REPORT_CARD.md should render the pass rate as n/a, got:\n%s", cardData)
	}
}
