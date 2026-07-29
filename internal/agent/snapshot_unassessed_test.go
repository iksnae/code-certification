package agent

import (
	"strings"
	"testing"

	"github.com/iksnae/code-certification/internal/domain"
)

// mixedSnapshotRecords: one analyzable Go unit at 0.9 and three Swift units the
// engine cannot read, all in the same package.
//
// Before this pass, internal/agent had zero occurrences of the string
// "Unsupported" anywhere in the package — a complete, parallel implementation
// of the rollup that internal/report and internal/workspace already do
// correctly. It is also the implementation whose output is pasted into the
// architect's LLM prompt, so its errors are laundered into prose the operator
// reads as analysis.
func mixedSnapshotRecords() []domain.CertificationRecord {
	recs := []domain.CertificationRecord{{
		UnitID:   domain.NewUnitID("go", "src/a.go", "A"),
		UnitPath: "src/a.go",
		Status:   domain.StatusCertified,
		Grade:    domain.GradeAMinus,
		Score:    0.9,
	}}
	for _, n := range []string{"x", "y", "z"} {
		recs = append(recs, domain.CertificationRecord{
			UnitID:      domain.NewUnitID("swift", "src/"+n+".swift", n),
			UnitPath:    "src/" + n + ".swift",
			Status:      domain.StatusExempt,
			Grade:       domain.GradeNA,
			Unsupported: true,
		})
	}
	return recs
}

// TestBuildSnapshot_AvgScoreExcludesUnassessedUnits: 0.9 over the one unit that
// was measured, not 0.9/4 = 0.225 over all four.
func TestBuildSnapshot_AvgScoreExcludesUnassessedUnits(t *testing.T) {
	snap := BuildSnapshot(mixedSnapshotRecords(), "")

	if snap.Metrics.TotalUnits != 4 {
		t.Errorf("TotalUnits = %d, want 4", snap.Metrics.TotalUnits)
	}
	if snap.Metrics.UnitsUnsupported != 3 {
		t.Errorf("UnitsUnsupported = %d, want 3", snap.Metrics.UnitsUnsupported)
	}
	if diff := snap.Metrics.AvgScore - 0.9; diff > 1e-9 || diff < -1e-9 {
		t.Errorf("AvgScore = %v, want 0.9 — three unassessed units must leave the denominator (0.225 is the all-four mean)",
			snap.Metrics.AvgScore)
	}
}

// TestBuildSnapshot_PackageNodeReportsItsUnassessedUnits.
func TestBuildSnapshot_PackageNodeReportsItsUnassessedUnits(t *testing.T) {
	snap := BuildSnapshot(mixedSnapshotRecords(), "")
	if len(snap.Packages) != 1 {
		t.Fatalf("packages = %d, want 1", len(snap.Packages))
	}
	p := snap.Packages[0]

	if p.Units != 4 {
		t.Errorf("Units = %d, want 4 — every unit in the package", p.Units)
	}
	if p.Unsupported != 3 {
		t.Errorf("Unsupported = %d, want 3", p.Unsupported)
	}
	if p.Grade != "A-" {
		t.Errorf("Grade = %q, want \"A-\" — the grade of the unit that was scored", p.Grade)
	}
}

// TestBuildSnapshot_AllUnassessedPackageHasNoGrade: with nothing measured there
// is no grade to state, so the package reports N/A rather than a manufactured F.
func TestBuildSnapshot_AllUnassessedPackageHasNoGrade(t *testing.T) {
	snap := BuildSnapshot(mixedSnapshotRecords()[1:], "")
	if len(snap.Packages) != 1 {
		t.Fatalf("packages = %d, want 1", len(snap.Packages))
	}
	if got := snap.Packages[0].Grade; got != domain.GradeNA.String() {
		t.Errorf("Grade = %q, want %q — nothing in this package was scored",
			got, domain.GradeNA.String())
	}
	if got := snap.Metrics.AvgScore; got != 0 {
		t.Errorf("AvgScore = %v, want 0 (unset, not a computed mean)", got)
	}
}

// TestFormatHotspots_UnassessedPackageIsNotTopRisk. Risk is units × (1 − score).
// An all-unassessed package scored 0 over its full unit count, which is the
// maximum risk the formula can produce — so the package the engine cannot read
// was ranked the repository's number one problem, in the prompt.
func TestFormatHotspots_UnassessedPackageIsNotTopRisk(t *testing.T) {
	snap := &ArchSnapshot{Hotspots: []PackageNode{
		{Path: "swiftonly", Units: 8, Unsupported: 8, AvgScore: 0, Grade: "N/A"},
	}}
	var b strings.Builder
	formatHotspots(&b, snap)
	out := b.String()

	if strings.Contains(out, "| 8.00 |") {
		t.Errorf("all-unassessed package carries maximum risk:\n%s", out)
	}
	if !strings.Contains(out, "| 0.00 |") {
		t.Errorf("all-unassessed package should carry no risk, got:\n%s", out)
	}
	if !strings.Contains(out, "n/a") {
		t.Errorf("hotspot row should not state a measured score, got:\n%s", out)
	}
}

// TestFormatHotspots_AssessedPackageRiskUnchanged is the over-reach guard.
func TestFormatHotspots_AssessedPackageRiskUnchanged(t *testing.T) {
	snap := &ArchSnapshot{Hotspots: []PackageNode{
		{Path: "core", Units: 10, Unsupported: 0, AvgScore: 0.5, Grade: "D"},
	}}
	var b strings.Builder
	formatHotspots(&b, snap)
	if got := b.String(); !strings.Contains(got, "| 5.00 |") {
		t.Errorf("risk for a fully assessed package should stay 10 × 0.5 = 5.00, got:\n%s", got)
	}
}
