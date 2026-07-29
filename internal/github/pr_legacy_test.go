package github_test

import (
	"math"
	"strings"
	"testing"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
	gh "github.com/iksnae/code-certification/internal/github"
)

// Two of pr.go's unsupported skips are unpinned, and both are unpinned for the
// same reason: the fixtures reach for the record the pipeline writes TODAY, whose
// StatusExempt already satisfies the surrounding condition. Neither skip is
// load-bearing on that shape — so removing it changes nothing, and the tests
// cannot tell whether it is there.
//
// The shape they ARE load-bearing on is the legacy corpus this branch exists to
// handle: a Swift unit recorded as decertified at score 0 before the engine knew
// it had no analyzer for Swift. The store backfills Unsupported from the language
// on read, so the record comes back Unsupported=true with Status=decertified —
// IsPassing() false, and the recorded flag is the only thing that keeps it out of
// a failing table and out of a trust delta.

// makeCert is an assessed record in a named language. pr_test.go's helper hard-
// codes "go", and the transition that matters here is a SWIFT unit that has
// become assessable.
func makeCert(lang, path, symbol string, status domain.Status, score float64) domain.CertificationRecord {
	return domain.CertificationRecord{
		UnitID:    domain.NewUnitID(lang, path, symbol),
		UnitPath:  path,
		Status:    status,
		Score:     score,
		Grade:     domain.GradeFromScore(score),
		ExpiresAt: time.Now().Add(90 * 24 * time.Hour),
	}
}

// makeLegacyUnsupportedRecord is that record: unassessed, but carrying the
// verdict the engine wrongly asserted before the unsupported-language fix.
func makeLegacyUnsupportedRecord(lang, path, symbol string) domain.CertificationRecord {
	return domain.CertificationRecord{
		UnitID:      domain.NewUnitID(lang, path, symbol),
		UnitPath:    path,
		Status:      domain.StatusDecertified,
		Grade:       domain.GradeF,
		Score:       0,
		Unsupported: true,
		ExpiresAt:   time.Now().Add(90 * 24 * time.Hour),
	}
}

// TestFormatPRComment_LegacyUnsupportedNotInFailingTable covers pr.go's
// "Units needing attention" filter.
//
// This is the client-facing table, and the legacy record is exactly what it
// would list: an F at 0.00 for a file the engine cannot open. The count above
// it already excludes the unit, so listing it puts a row in the table that the
// summary line says does not exist.
func TestFormatPRComment_LegacyUnsupportedNotInFailingTable(t *testing.T) {
	records := []domain.CertificationRecord{
		makeCert("go", "engine.go", "Run", domain.StatusDecertified, 0.30),
		makeLegacyUnsupportedRecord("swift", "App.swift", "App"),
	}

	comment := gh.FormatPRComment(records, true)

	if !strings.Contains(comment, "### Units needing attention") {
		t.Fatalf("fixture: the failing table should render for the genuinely failing unit:\n%s", comment)
	}
	if strings.Contains(comment, "App.swift") {
		t.Errorf("failing table lists a unit the engine never assessed — an F at 0.00 for code it "+
			"cannot open, contradicting the count above it:\n%s", comment)
	}
	if !strings.Contains(comment, "engine.go") {
		t.Errorf("failing table dropped the genuinely failing unit:\n%s", comment)
	}
	if !strings.Contains(comment, "**0/1** units passing certification") {
		t.Errorf("comment should count one assessed unit, got:\n%s", comment)
	}
	if !strings.Contains(comment, "**1** unit(s) were not assessed") {
		t.Errorf("comment should report the unassessed unit separately, got:\n%s", comment)
	}
	if !strings.Contains(comment, "**1 units failed certification in enforcing mode.") {
		t.Errorf("enforcing footer should count only the assessed failure, got:\n%s", comment)
	}
}

// TestComputeTrustDelta_OldSideUnsupportedIsNotAPriorVerdict covers the OLD-side
// skip in ComputeTrustDelta.
//
// #45 item 1 reads as closed because three ComputeTrustDelta tests were added,
// but all three vary the NEW side; the old side is still the un-asserted half.
// It matters on exactly one transition — a unit that was unassessed and has now
// been assessed, which is what every repo sees the first time an analyzer for its
// language ships. Keeping the old record in the map gives that unit a prior
// "verdict" of StatusExempt, whose IsPassing() is true, so becoming genuinely
// certified reads as no movement at all and the improvement is reported as zero.
func TestComputeTrustDelta_OldSideUnsupportedIsNotAPriorVerdict(t *testing.T) {
	old := []domain.CertificationRecord{
		makeUnsupportedRecord("swift", "App.swift", "App"),
	}
	updated := []domain.CertificationRecord{
		makeCert("swift", "App.swift", "App", domain.StatusCertified, 0.88),
	}

	d := gh.ComputeTrustDelta(old, updated)

	if d.NewlyCertified != 1 {
		t.Errorf("NewlyCertified = %d, want 1 — a unit that went from unassessed to certified became "+
			"certified. Keeping the old unassessed record as a prior verdict makes StatusExempt.IsPassing() "+
			"stand in for one, and the movement disappears", d.NewlyCertified)
	}
	if d.NewlyDecertified != 0 {
		t.Errorf("NewlyDecertified = %d, want 0", d.NewlyDecertified)
	}
	if d.ScoreImproved != 0 {
		t.Errorf("ScoreImproved = %d, want 0 — there is no prior score to have improved on; the stored 0 "+
			"is the absence of a measurement, not a measurement of zero", d.ScoreImproved)
	}
	if d.AverageScoreDelta != 0 {
		t.Errorf("AverageScoreDelta = %v, want 0 — with no comparable prior record there is no delta. "+
			"Admitting the old placeholder 0 reports a fabricated +0.88 improvement", d.AverageScoreDelta)
	}
}

// TestComputeTrustDelta_LegacyOldSideUnsupportedIsExcluded is the same skip on
// the legacy shape, where the old record's status is decertified rather than
// exempt. Here the un-skipped path fabricates movement in the opposite
// direction — a newly-certified unit that was never decertified.
func TestComputeTrustDelta_LegacyOldSideUnsupportedIsExcluded(t *testing.T) {
	old := []domain.CertificationRecord{
		makeLegacyUnsupportedRecord("swift", "App.swift", "App"),
		makeCert("go", "engine.go", "Run", domain.StatusCertified, 0.80),
	}
	updated := []domain.CertificationRecord{
		makeUnsupportedRecord("swift", "App.swift", "App"),
		makeCert("go", "engine.go", "Run", domain.StatusCertified, 0.90),
	}

	d := gh.ComputeTrustDelta(old, updated)

	if d.NewlyCertified != 0 {
		t.Errorf("NewlyCertified = %d, want 0 — the Swift unit did not become certified, it stopped "+
			"being wrongly decertified. Reporting it as newly certified is the false positive the "+
			"unsupported exclusion exists to prevent", d.NewlyCertified)
	}
	if d.ScoreImproved != 1 {
		t.Errorf("ScoreImproved = %d, want 1 — only the Go unit moved", d.ScoreImproved)
	}
	// Only the Go unit is comparable on both sides: 0.90 now against 0.80 before.
	// Admitting the Swift placeholder on the old side alone drags the old mean to
	// 0.40 and doubles the reported improvement.
	if math.Abs(d.AverageScoreDelta-0.10) > 1e-9 {
		t.Errorf("AverageScoreDelta = %v, want 0.10 — the Swift placeholder must not enter either mean",
			d.AverageScoreDelta)
	}
}
