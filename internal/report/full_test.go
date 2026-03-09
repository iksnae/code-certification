package report_test

import (
	"strings"
	"testing"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/report"
)

func makeFullRecord(lang, path, symbol string, utype domain.UnitType, status domain.Status, score float64) domain.CertificationRecord {
	now := time.Now()
	dims := domain.DimensionScores{
		domain.DimCorrectness:     score + 0.05,
		domain.DimMaintainability: score,
		domain.DimReadability:     score - 0.02,
		domain.DimTestability:     score + 0.03,
		domain.DimSecurity:        0.80,
	}
	return domain.CertificationRecord{
		UnitID:      domain.NewUnitID(lang, path, symbol),
		UnitType:    utype,
		UnitPath:    path,
		Status:      status,
		Score:       score,
		Grade:       domain.GradeFromScore(score),
		Confidence:  1.0,
		Dimensions:  dims,
		CertifiedAt: now,
		ExpiresAt:   now.Add(90 * 24 * time.Hour),
		Source:      "deterministic",
	}
}

func TestGenerateFullReport(t *testing.T) {
	records := []domain.CertificationRecord{
		makeFullRecord("go", "internal/engine/scorer.go", "Score", domain.UnitTypeFunction, domain.StatusCertified, 0.90),
		makeFullRecord("go", "internal/engine/pipeline.go", "CertifyUnit", domain.UnitTypeFunction, domain.StatusCertified, 0.85),
		makeFullRecord("go", "cmd/certify/main.go", "main", domain.UnitTypeFunction, domain.StatusCertified, 0.88),
		makeFullRecord("ts", "src/parser.ts", "parse", domain.UnitTypeFunction, domain.StatusCertifiedWithObservations, 0.75),
	}
	records[3].Observations = []string{"complexity: 25 exceeds threshold 20"}

	r := report.GenerateFullReport(records, "test/repo", "abc123", time.Now())

	if len(r.Units) != 4 {
		t.Errorf("units = %d, want 4", len(r.Units))
	}
	if r.Card.TotalUnits != 4 {
		t.Errorf("card total = %d, want 4", r.Card.TotalUnits)
	}
	if len(r.DimensionAverages) == 0 {
		t.Error("dimension averages should not be empty")
	}
	if len(r.LanguageDetail) != 2 {
		t.Errorf("languages = %d, want 2", len(r.LanguageDetail))
	}

	// Check units are sorted by ID
	for i := 1; i < len(r.Units); i++ {
		if r.Units[i].UnitID < r.Units[i-1].UnitID {
			t.Error("units should be sorted by ID")
		}
	}
}

func TestFormatFullMarkdown(t *testing.T) {
	records := []domain.CertificationRecord{
		makeFullRecord("go", "internal/engine/scorer.go", "Score", domain.UnitTypeFunction, domain.StatusCertified, 0.90),
		makeFullRecord("go", "internal/engine/pipeline.go", "CertifyUnit", domain.UnitTypeFunction, domain.StatusCertified, 0.85),
		makeFullRecord("ts", "src/parser.ts", "parse", domain.UnitTypeFunction, domain.StatusCertifiedWithObservations, 0.75),
	}
	records[2].Observations = []string{"complexity too high"}

	r := report.GenerateFullReport(records, "test/repo", "abc123", time.Now())
	md := report.FormatFullMarkdown(r)

	checks := []string{
		"# ",                    // title
		"## Summary",            // summary section
		"## Grade Distribution", // grade dist
		"## Dimension Averages", // dimensions
		"## By Language",        // languages
		"## All Units",          // all units
		"| Unit |",              // table header
		"Score",                 // unit name in table
		"CertifyUnit",           // unit name in table
		"parse",                 // ts unit
		"complexity too high",   // observation detail
		"certify",               // footer link
	}
	for _, check := range checks {
		if !strings.Contains(md, check) {
			t.Errorf("markdown should contain %q", check)
		}
	}
}

func TestFullReport_LanguageDetail(t *testing.T) {
	records := []domain.CertificationRecord{
		makeFullRecord("go", "a.go", "A", domain.UnitTypeFunction, domain.StatusCertified, 0.95),
		makeFullRecord("go", "b.go", "B", domain.UnitTypeFunction, domain.StatusCertified, 0.85),
		makeFullRecord("go", "c.go", "C", domain.UnitTypeFunction, domain.StatusCertified, 0.75),
	}

	r := report.GenerateFullReport(records, "", "", time.Now())

	if len(r.LanguageDetail) != 1 {
		t.Fatalf("expected 1 language, got %d", len(r.LanguageDetail))
	}
	ld := r.LanguageDetail[0]
	if ld.Name != "go" {
		t.Errorf("lang = %s, want go", ld.Name)
	}
	if ld.TopScore < 0.94 || ld.TopScore > 0.96 {
		t.Errorf("top = %f, want ~0.95", ld.TopScore)
	}
	if ld.BottomScore < 0.74 || ld.BottomScore > 0.76 {
		t.Errorf("bottom = %f, want ~0.75", ld.BottomScore)
	}
}
