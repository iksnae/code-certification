package report_test

import (
	"testing"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/report"
)

// TestDetailed_RecurrentAreasExcludeUnassessedUnits.
//
// findRecurrentlyFailing counted every record in dirTotals and summed every
// Score into dirScores, while dirFailing could never include an unassessed unit
// — StatusExempt is passing. The two sides of the row therefore covered
// different populations: a directory of two failing Go units beside one
// unassessed Swift file rendered "2/3 failing, avg 0.300" where the measurement
// over what was actually assessed is "2/2, avg 0.450".
func TestDetailed_RecurrentAreasExcludeUnassessedUnits(t *testing.T) {
	recs := []domain.CertificationRecord{
		{
			UnitID:   domain.NewUnitID("go", "src/a.go", "A"),
			UnitPath: "src/a.go",
			Status:   domain.StatusDecertified,
			Grade:    domain.GradeF,
			Score:    0.4,
		},
		{
			UnitID:   domain.NewUnitID("go", "src/b.go", "B"),
			UnitPath: "src/b.go",
			Status:   domain.StatusDecertified,
			Grade:    domain.GradeF,
			Score:    0.5,
		},
		{
			UnitID:      domain.NewUnitID("swift", "src/c.swift", "C"),
			UnitPath:    "src/c.swift",
			Status:      domain.StatusExempt,
			Grade:       domain.GradeNA,
			Unsupported: true,
		},
	}

	d := report.Detailed(recs, time.Now())
	if len(d.RecurrentlyFailing) != 1 {
		t.Fatalf("RecurrentlyFailing = %d, want 1", len(d.RecurrentlyFailing))
	}
	a := d.RecurrentlyFailing[0]

	if a.Total != 2 {
		t.Errorf("Total = %d, want 2 — the unassessed unit is not part of the area's assessed population", a.Total)
	}
	if a.Failing != 2 {
		t.Errorf("Failing = %d, want 2", a.Failing)
	}
	const want = 0.45
	if diff := a.AverageScore - want; diff > 1e-9 || diff < -1e-9 {
		t.Errorf("AverageScore = %v, want %v — 0.300 is the mean with a placeholder zero in the sum",
			a.AverageScore, want)
	}
}
