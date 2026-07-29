package github_test

import (
	"strings"
	"testing"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
	gh "github.com/iksnae/code-certification/internal/github"
)

// makeUnsupportedRecord builds the record the pipeline produces for a unit in a
// language the engine cannot analyse: no verdict asserted, Unsupported set.
func makeUnsupportedRecord(lang, path, symbol string) domain.CertificationRecord {
	return domain.CertificationRecord{
		UnitID:      domain.NewUnitID(lang, path, symbol),
		UnitPath:    path,
		Status:      domain.StatusExempt,
		Grade:       domain.GradeNA,
		Score:       0,
		Confidence:  0,
		Unsupported: true,
		ExpiresAt:   time.Now().Add(90 * 24 * time.Hour),
	}
}

// TestFormatPRComment_UnsupportedNotCounted guards the client-facing half of the
// unsupported-language fix. StatusExempt.IsPassing() is true, so an unassessed
// unit inflates the numerator AND the denominator and turns the header green —
// the tool reports "3/3 units passing certification" for code it never opened.
// That quiet false positive is worse than the loud false negative it replaced.
func TestFormatPRComment_UnsupportedNotCounted(t *testing.T) {
	records := []domain.CertificationRecord{
		makeUnsupportedRecord("swift", "App.swift", "App"),
		makeUnsupportedRecord("swift", "View.swift", "View"),
		makeUnsupportedRecord("swift", "Model.swift", "Model"),
	}

	comment := gh.FormatPRComment(records, true)

	if strings.Contains(comment, "✅") {
		t.Errorf("all-unsupported comment claims certification success:\n%s", comment)
	}
	if strings.Contains(comment, "3/3") {
		t.Errorf("unassessed units counted as passing certification:\n%s", comment)
	}
	if !strings.Contains(comment, "not assessed") {
		t.Errorf("unassessed units are not reported separately:\n%s", comment)
	}
}

// TestFormatPRComment_UnsupportedExcludedFromDenominator checks the mixed case:
// two real assessed units, one unassessed. The ratio must describe only what was
// actually certified.
func TestFormatPRComment_UnsupportedExcludedFromDenominator(t *testing.T) {
	records := []domain.CertificationRecord{
		makeTestRecord("main.go", "main", domain.StatusCertified, 0.95),
		makeTestRecord("util.go", "util", domain.StatusCertified, 0.91),
		makeUnsupportedRecord("swift", "App.swift", "App"),
	}

	comment := gh.FormatPRComment(records, false)

	if !strings.Contains(comment, "2/2") {
		t.Errorf("want 2/2 assessed units passing, got:\n%s", comment)
	}
	if strings.Contains(comment, "3/3") {
		t.Errorf("unassessed unit inflated the denominator:\n%s", comment)
	}
	if !strings.Contains(comment, "not assessed") {
		t.Errorf("the 1 unassessed unit is not reported separately:\n%s", comment)
	}
}

// TestFormatPRComment_UnsupportedNotFailing guards the opposite error. Excluding
// unassessed units must not push them into the failing bucket — that would
// reproduce the original "unassessed counted as failing" defect.
func TestFormatPRComment_UnsupportedNotFailing(t *testing.T) {
	records := []domain.CertificationRecord{
		makeTestRecord("main.go", "main", domain.StatusCertified, 0.95),
		makeUnsupportedRecord("swift", "App.swift", "App"),
	}

	comment := gh.FormatPRComment(records, true)

	if strings.Contains(comment, "BLOCKED") {
		t.Errorf("an unassessed unit must not block an enforcing merge:\n%s", comment)
	}
	if strings.Contains(comment, "failed certification") {
		t.Errorf("an unassessed unit must not be reported as failed:\n%s", comment)
	}
}

// TestComputeTrustDelta_UnsupportedNotNewlyCertified guards the trust delta. On
// the first run after the unsupported fix, every unsupported unit moves from
// StatusDecertified (the old wrong verdict) to StatusExempt. Both sides of
// `!or.Status.IsPassing() && nr.Status.IsPassing()` are satisfied, so the tool
// reports the units as newly certified. Moving from wrongly-failed to
// unassessed is not a new certification.
func TestComputeTrustDelta_UnsupportedNotNewlyCertified(t *testing.T) {
	oldRecords := []domain.CertificationRecord{
		makeTestRecord("App.swift", "App", domain.StatusDecertified, 0),
		makeTestRecord("View.swift", "View", domain.StatusDecertified, 0),
		makeTestRecord("Model.swift", "Model", domain.StatusDecertified, 0),
	}
	newRecords := []domain.CertificationRecord{
		makeUnsupportedRecord("go", "App.swift", "App"),
		makeUnsupportedRecord("go", "View.swift", "View"),
		makeUnsupportedRecord("go", "Model.swift", "Model"),
	}

	d := gh.ComputeTrustDelta(oldRecords, newRecords)

	if d.NewlyCertified != 0 {
		t.Errorf("NewlyCertified = %d, want 0 — becoming unassessed is not a certification", d.NewlyCertified)
	}
}

// TestComputeTrustDelta_UnsupportedNotNewlyDecertified guards the mirror case: a
// unit that was genuinely certified and is now recognised as unassessable has
// not been decertified either.
func TestComputeTrustDelta_UnsupportedNotNewlyDecertified(t *testing.T) {
	oldRecords := []domain.CertificationRecord{
		makeTestRecord("App.swift", "App", domain.StatusCertified, 0.95),
	}
	newRecords := []domain.CertificationRecord{
		makeUnsupportedRecord("go", "App.swift", "App"),
	}

	d := gh.ComputeTrustDelta(oldRecords, newRecords)

	if d.NewlyDecertified != 0 {
		t.Errorf("NewlyDecertified = %d, want 0 — becoming unassessed is not a decertification", d.NewlyDecertified)
	}
	if d.ScoreDegraded != 0 {
		t.Errorf("ScoreDegraded = %d, want 0 — an unassessed unit has no score to compare", d.ScoreDegraded)
	}
}

// TestComputeTrustDelta_UnsupportedExcludedFromAverage checks that the zero
// score the pipeline stores for an unassessed unit does not drag the average
// score delta. An absent measurement is not a score of 0.
func TestComputeTrustDelta_UnsupportedExcludedFromAverage(t *testing.T) {
	oldRecords := []domain.CertificationRecord{
		makeTestRecord("main.go", "main", domain.StatusCertified, 0.90),
	}
	newRecords := []domain.CertificationRecord{
		makeTestRecord("main.go", "main", domain.StatusCertified, 0.90),
		makeUnsupportedRecord("swift", "App.swift", "App"),
	}

	d := gh.ComputeTrustDelta(oldRecords, newRecords)

	if d.AverageScoreDelta != 0 {
		t.Errorf("AverageScoreDelta = %f, want 0 — an unassessed unit is not a 0.0 score", d.AverageScoreDelta)
	}
}
