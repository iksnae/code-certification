package report_test

import (
	"testing"
	"time"

	"github.com/code-certification/certify/internal/domain"
	"github.com/code-certification/certify/internal/report"
)

func makeRecord(path, symbol string, status domain.Status, score float64) domain.CertificationRecord {
	now := time.Now()
	return domain.CertificationRecord{
		UnitID:      domain.NewUnitID("go", path, symbol),
		UnitType:    domain.UnitTypeFunction,
		UnitPath:    path,
		Status:      status,
		Score:       score,
		Grade:       domain.GradeFromScore(score),
		CertifiedAt: now,
		ExpiresAt:   now.Add(90 * 24 * time.Hour),
	}
}

func TestHealthReport_Summary(t *testing.T) {
	records := []domain.CertificationRecord{
		makeRecord("main.go", "main", domain.StatusCertified, 0.9),
		makeRecord("main.go", "helper", domain.StatusCertified, 0.85),
		makeRecord("service/sync.go", "Apply", domain.StatusCertifiedWithObservations, 0.72),
		makeRecord("service/sync.go", "Reset", domain.StatusProbationary, 0.55),
		makeRecord("bad.go", "broken", domain.StatusDecertified, 0.30),
	}

	h := report.Health(records)

	if h.TotalUnits != 5 {
		t.Errorf("TotalUnits = %d, want 5", h.TotalUnits)
	}
	if h.Certified != 2 {
		t.Errorf("Certified = %d, want 2", h.Certified)
	}
	if h.CertifiedWithObs != 1 {
		t.Errorf("CertifiedWithObs = %d, want 1", h.CertifiedWithObs)
	}
	if h.Probationary != 1 {
		t.Errorf("Probationary = %d, want 1", h.Probationary)
	}
	if h.Decertified != 1 {
		t.Errorf("Decertified = %d, want 1", h.Decertified)
	}
	if h.PassRate < 0.59 || h.PassRate > 0.61 {
		t.Errorf("PassRate = %f, want ~0.6", h.PassRate)
	}
}

func TestHealthReport_Empty(t *testing.T) {
	h := report.Health(nil)
	if h.TotalUnits != 0 {
		t.Errorf("empty TotalUnits = %d, want 0", h.TotalUnits)
	}
	if h.PassRate != 0 {
		t.Errorf("empty PassRate = %f, want 0", h.PassRate)
	}
}

func TestHealthReport_AllCertified(t *testing.T) {
	records := []domain.CertificationRecord{
		makeRecord("a.go", "A", domain.StatusCertified, 0.95),
		makeRecord("b.go", "B", domain.StatusCertified, 0.90),
	}
	h := report.Health(records)
	if h.PassRate != 1.0 {
		t.Errorf("all certified PassRate = %f, want 1.0", h.PassRate)
	}
}

func TestFormatJSON(t *testing.T) {
	h := report.HealthReport{
		TotalUnits: 5,
		Certified:  3,
		PassRate:   0.6,
	}
	data, err := report.FormatJSON(h)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) == 0 {
		t.Error("FormatJSON should produce output")
	}
}

func TestFormatText(t *testing.T) {
	h := report.HealthReport{
		TotalUnits:      5,
		Certified:       2,
		CertifiedWithObs: 1,
		Probationary:    1,
		Decertified:     1,
		PassRate:        0.6,
		AverageScore:    0.664,
	}
	text := report.FormatText(h)
	if len(text) == 0 {
		t.Error("FormatText should produce output")
	}
}
