package report_test

import (
	"strings"
	"testing"
	"time"

	"github.com/code-certification/certify/internal/domain"
	"github.com/code-certification/certify/internal/report"
)

func makeDetailedRecord(lang, path, symbol string, status domain.Status, score float64, dims domain.DimensionScores, expiresIn time.Duration) domain.CertificationRecord {
	now := time.Now()
	return domain.CertificationRecord{
		UnitID:       domain.NewUnitID(lang, path, symbol),
		UnitType:     domain.UnitTypeFunction,
		UnitPath:     path,
		Status:       status,
		Score:        score,
		Grade:        domain.GradeFromScore(score),
		Dimensions:   dims,
		CertifiedAt:  now,
		ExpiresAt:    now.Add(expiresIn),
		Observations: []string{},
	}
}

func TestDetailedReport_DimensionAverages(t *testing.T) {
	dims1 := domain.DimensionScores{
		domain.DimCorrectness:     0.9,
		domain.DimMaintainability: 0.8,
	}
	dims2 := domain.DimensionScores{
		domain.DimCorrectness:     0.7,
		domain.DimMaintainability: 0.6,
	}

	records := []domain.CertificationRecord{
		makeDetailedRecord("go", "a.go", "A", domain.StatusCertified, 0.9, dims1, 90*24*time.Hour),
		makeDetailedRecord("go", "b.go", "B", domain.StatusCertified, 0.8, dims2, 90*24*time.Hour),
	}

	d := report.Detailed(records, time.Now())

	if v, ok := d.Dimensions["correctness"]; !ok || v < 0.79 || v > 0.81 {
		t.Errorf("correctness avg = %v, want ~0.8", v)
	}
	if v, ok := d.Dimensions["maintainability"]; !ok || v < 0.69 || v > 0.71 {
		t.Errorf("maintainability avg = %v, want ~0.7", v)
	}
}

func TestDetailedReport_ByLanguage(t *testing.T) {
	records := []domain.CertificationRecord{
		makeDetailedRecord("go", "a.go", "A", domain.StatusCertified, 0.9, nil, 90*24*time.Hour),
		makeDetailedRecord("go", "b.go", "B", domain.StatusCertified, 0.8, nil, 90*24*time.Hour),
		makeDetailedRecord("ts", "c.ts", "C", domain.StatusDecertified, 0.3, nil, 90*24*time.Hour),
	}

	d := report.Detailed(records, time.Now())

	goLang := d.ByLanguage["go"]
	if goLang.Total != 2 {
		t.Errorf("go total = %d, want 2", goLang.Total)
	}
	if goLang.Passing != 2 {
		t.Errorf("go passing = %d, want 2", goLang.Passing)
	}
	if goLang.AverageScore < 0.84 || goLang.AverageScore > 0.86 {
		t.Errorf("go avg score = %f, want ~0.85", goLang.AverageScore)
	}

	tsLang := d.ByLanguage["ts"]
	if tsLang.Total != 1 {
		t.Errorf("ts total = %d, want 1", tsLang.Total)
	}
	if tsLang.Passing != 0 {
		t.Errorf("ts passing = %d, want 0", tsLang.Passing)
	}
}

func TestDetailedReport_ExpiringSoon(t *testing.T) {
	now := time.Now()
	records := []domain.CertificationRecord{
		makeDetailedRecord("go", "expiring.go", "A", domain.StatusCertified, 0.9, nil, 5*24*time.Hour),  // 5 days — expiring soon
		makeDetailedRecord("go", "safe.go", "B", domain.StatusCertified, 0.9, nil, 60*24*time.Hour),     // 60 days — not soon
		makeDetailedRecord("go", "also-exp.go", "C", domain.StatusCertified, 0.9, nil, 10*24*time.Hour), // 10 days — expiring soon
	}

	d := report.Detailed(records, now)

	if len(d.ExpiringSoon) != 2 {
		t.Fatalf("expiring soon = %d, want 2", len(d.ExpiringSoon))
	}
	// Should be sorted by expiry date (earliest first)
	if !strings.Contains(d.ExpiringSoon[0].UnitID, "expiring.go") {
		t.Errorf("first expiring = %s, want expiring.go", d.ExpiringSoon[0].UnitID)
	}
}

func TestDetailedReport_HighestRisk(t *testing.T) {
	records := []domain.CertificationRecord{
		makeDetailedRecord("go", "good.go", "A", domain.StatusCertified, 0.95, nil, 90*24*time.Hour),
		makeDetailedRecord("go", "risky.go", "B", domain.StatusProbationary, 0.55, nil, 90*24*time.Hour),
		makeDetailedRecord("go", "bad.go", "C", domain.StatusDecertified, 0.30, nil, 90*24*time.Hour),
	}

	d := report.Detailed(records, time.Now())

	if len(d.HighestRisk) != 3 {
		t.Fatalf("highest risk = %d, want 3", len(d.HighestRisk))
	}
	// Lowest score first
	if d.HighestRisk[0].Score > 0.31 {
		t.Errorf("highest risk[0] score = %f, want ~0.30", d.HighestRisk[0].Score)
	}
}

func TestDetailedReport_Failing(t *testing.T) {
	rec := makeDetailedRecord("go", "bad.go", "C", domain.StatusDecertified, 0.30, nil, 90*24*time.Hour)
	rec.Observations = []string{"lint_errors: 5 exceeds threshold 0"}

	records := []domain.CertificationRecord{
		makeDetailedRecord("go", "good.go", "A", domain.StatusCertified, 0.95, nil, 90*24*time.Hour),
		rec,
	}

	d := report.Detailed(records, time.Now())

	if len(d.Failing) != 1 {
		t.Fatalf("failing = %d, want 1", len(d.Failing))
	}
	if !strings.Contains(d.Failing[0].Explanation, "Decertified") {
		t.Errorf("explanation = %s, want 'Decertified...'", d.Failing[0].Explanation)
	}
	if len(d.Failing[0].Observations) != 1 {
		t.Errorf("observations = %d, want 1", len(d.Failing[0].Observations))
	}
}

func TestDetailedReport_Empty(t *testing.T) {
	d := report.Detailed(nil, time.Now())
	if d.TotalUnits != 0 {
		t.Errorf("empty total = %d, want 0", d.TotalUnits)
	}
	if len(d.Dimensions) != 0 {
		t.Errorf("empty dimensions = %d, want 0", len(d.Dimensions))
	}
}

func TestFormatDetailedText(t *testing.T) {
	dims := domain.DimensionScores{
		domain.DimCorrectness: 0.9,
	}
	records := []domain.CertificationRecord{
		makeDetailedRecord("go", "a.go", "A", domain.StatusCertified, 0.9, dims, 5*24*time.Hour),
		makeDetailedRecord("ts", "b.ts", "B", domain.StatusDecertified, 0.3, nil, 90*24*time.Hour),
	}
	records[1].Observations = []string{"lint failed"}

	d := report.Detailed(records, time.Now())
	text := report.FormatDetailedText(d)

	if !strings.Contains(text, "Dimension Averages") {
		t.Error("should contain 'Dimension Averages'")
	}
	if !strings.Contains(text, "By Language") {
		t.Error("should contain 'By Language'")
	}
	if !strings.Contains(text, "Expiring Soon") {
		t.Error("should contain 'Expiring Soon'")
	}
	if !strings.Contains(text, "Highest Risk") {
		t.Error("should contain 'Highest Risk'")
	}
	if !strings.Contains(text, "Failing") {
		t.Error("should contain 'Failing'")
	}
}
