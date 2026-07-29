package main

import (
	"math"
	"testing"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/report"
	"github.com/iksnae/code-certification/internal/workspace"
)

// report_cmd.go had no test file at all, and it holds the single wire that
// carries the unassessed count from a submodule's report card into the workspace
// rollup. Every workspace test constructs a SubmoduleSummary by hand and so
// bypasses this line entirely: zero the Unsupported assignment and an
// all-unassessed workspace prints a real grade, a real score and a real pass
// rate, with both packages' suites green.
//
// That is the shape of the gap the reviewer named — the verification followed the
// code that was interesting to write, not the code that carries the claim. A
// mapping this load-bearing has to be a named function that a test can call,
// rather than eight assignments inside a command that needs a git checkout to
// reach.

func unassessedCard(t *testing.T, n int) report.Card {
	t.Helper()
	records := make([]domain.CertificationRecord, 0, n)
	for i := 0; i < n; i++ {
		r := domain.CertificationRecord{
			UnitID:   domain.NewUnitID("swift", "app/f.swift", string(rune('a'+i))),
			UnitType: domain.UnitTypeFunction,
			UnitPath: "app/f.swift",
			Status:   domain.StatusExempt,
			Grade:    domain.GradeNA,
		}
		r.Unsupported = true
		records = append(records, r)
	}
	return report.GenerateCard(records, "ios", "abc", time.Now())
}

// TestSubmoduleSummaryFromCard_CarriesTheUnassessedCount pins the wire itself.
func TestSubmoduleSummaryFromCard_CarriesTheUnassessedCount(t *testing.T) {
	card := unassessedCard(t, 3)
	if card.UnsupportedCount != 3 {
		t.Fatalf("fixture: card.UnsupportedCount = %d, want 3", card.UnsupportedCount)
	}

	s := submoduleSummaryFromCard(workspace.SubmoduleSummary{Name: "ios", Path: "apps/ios", HasCertify: true}, card)

	if s.Unsupported != 3 {
		t.Errorf("Unsupported = %d, want 3 — this is the only path by which the workspace rollup learns "+
			"that a submodule was never analysed", s.Unsupported)
	}
	if s.Units != 3 {
		t.Errorf("Units = %d, want 3", s.Units)
	}
	if s.Analyzable() != 0 {
		t.Errorf("Analyzable() = %d, want 0", s.Analyzable())
	}
	if s.ScoreKnown() {
		t.Error("ScoreKnown() = true, want false")
	}
	if s.Grade != "N/A" {
		t.Errorf("Grade = %q, want %q", s.Grade, "N/A")
	}
	// The remaining fields are asserted so the mapping cannot lose one silently.
	if s.Passing != 0 || s.Failing != 0 {
		t.Errorf("Passing/Failing = %d/%d, want 0/0", s.Passing, s.Failing)
	}
	if s.StateAt == "" {
		t.Error("StateAt should carry the card's generation time")
	}
	if s.Name != "ios" || s.Path != "apps/ios" || !s.HasCertify {
		t.Errorf("mapping dropped an identity field: %+v", s)
	}
}

// TestAggregateCards_AllUnassessedThroughTheRealWire is the end-to-end assertion
// the hand-built fixtures cannot make: records in, workspace card out, with the
// production mapping in between.
func TestAggregateCards_AllUnassessedThroughTheRealWire(t *testing.T) {
	subs := []workspace.SubmoduleSummary{
		submoduleSummaryFromCard(
			workspace.SubmoduleSummary{Name: "ios", Path: "apps/ios", HasCertify: true}, unassessedCard(t, 3)),
		submoduleSummaryFromCard(
			workspace.SubmoduleSummary{Name: "mac", Path: "apps/mac", HasCertify: true}, unassessedCard(t, 2)),
	}

	wc := workspace.AggregateCards(subs)

	if wc.TotalUnsupported != 5 {
		t.Errorf("TotalUnsupported = %d, want 5", wc.TotalUnsupported)
	}
	if wc.AnalyzableUnits != 0 {
		t.Errorf("AnalyzableUnits = %d, want 0", wc.AnalyzableUnits)
	}
	if wc.ScoreKnown() || wc.PassRateKnown() {
		t.Error("workspace claims a known score or pass rate over units nothing analysed")
	}
	if wc.OverallGrade != "N/A" {
		t.Errorf("OverallGrade = %q, want %q — a real grade here is a verdict about five unopened files",
			wc.OverallGrade, "N/A")
	}
	if wc.PassRate != 0 || wc.OverallScore != 0 {
		t.Errorf("PassRate/OverallScore = %v/%v, want the zero values left untouched", wc.PassRate, wc.OverallScore)
	}
}

// TestSubmoduleSummaryFromCard_MixedCarriesTheAnalyzableScore pins the wire on a
// mixed submodule, where a dropped Unsupported assignment is visible in the
// ROLLUP's arithmetic rather than only in its N/A gate.
func TestSubmoduleSummaryFromCard_MixedCarriesTheAnalyzableScore(t *testing.T) {
	records := []domain.CertificationRecord{
		{
			UnitID: domain.NewUnitID("go", "app/a.go", "A"), UnitType: domain.UnitTypeFunction,
			UnitPath: "app/a.go", Status: domain.StatusCertified, Score: 0.90, Grade: domain.GradeAMinus,
		},
		{
			UnitID: domain.NewUnitID("swift", "app/b.swift", "b"), UnitType: domain.UnitTypeFunction,
			UnitPath: "app/b.swift", Status: domain.StatusExempt, Grade: domain.GradeNA, Unsupported: true,
		},
	}
	s := submoduleSummaryFromCard(
		workspace.SubmoduleSummary{Name: "api", Path: "services/api", HasCertify: true},
		report.GenerateCard(records, "api", "abc", time.Now()))

	if s.Unsupported != 1 {
		t.Fatalf("Unsupported = %d, want 1", s.Unsupported)
	}
	if s.Analyzable() != 1 {
		t.Fatalf("Analyzable() = %d, want 1", s.Analyzable())
	}

	wc := workspace.AggregateCards([]workspace.SubmoduleSummary{s})
	if math.Abs(wc.OverallScore-0.90) > 1e-9 {
		t.Errorf("workspace OverallScore = %v, want 0.90. A dropped Unsupported assignment weights the "+
			"submodule by 2 and divides by 2, halving it to 0.45 and grading the workspace F", wc.OverallScore)
	}
	if wc.OverallGrade != "A-" {
		t.Errorf("workspace OverallGrade = %q, want %q", wc.OverallGrade, "A-")
	}
}
