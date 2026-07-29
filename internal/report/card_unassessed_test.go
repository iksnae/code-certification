package report_test

import (
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/report"
)

// unassessedRecord is a unit in a language the engine cannot analyse. The
// pipeline gives it StatusExempt (IsPassing() == true) and sets Unsupported,
// which is the only field that distinguishes "no verdict was asserted" from
// "certified".
func unassessedRecord(lang, path, symbol string) domain.CertificationRecord {
	r := makeCardRecord(lang, path, symbol, domain.StatusExempt, 0)
	r.Unsupported = true
	r.Grade = domain.GradeNA
	return r
}

// TestGenerateCard_AllUnassessedIsNotAPassRate is the reddening test for #47.
// A repo the engine never opened must not report a pass rate of 100%.
func TestGenerateCard_AllUnassessedIsNotAPassRate(t *testing.T) {
	records := []domain.CertificationRecord{
		unassessedRecord("swift", "a.swift", "a"),
		unassessedRecord("swift", "b.swift", "b"),
		unassessedRecord("swift", "c.swift", "c"),
	}

	c := report.GenerateCard(records, "test/repo", "abc", time.Now())

	if c.PassRate == 1.0 {
		t.Errorf("PassRate = %.1f%%, want not 100%% — no unit was assessed", c.PassRate*100)
	}
	if c.Passing != 0 {
		t.Errorf("Passing = %d, want 0 — an unassessed unit did not pass", c.Passing)
	}
	if c.Failing != 0 {
		t.Errorf("Failing = %d, want 0 — an unassessed unit did not fail", c.Failing)
	}
	if c.UnsupportedCount != 3 {
		t.Errorf("UnsupportedCount = %d, want 3", c.UnsupportedCount)
	}
	if c.AnalyzableUnits != 0 {
		t.Errorf("AnalyzableUnits = %d, want 0", c.AnalyzableUnits)
	}
	if c.TotalUnits != 3 {
		t.Errorf("TotalUnits = %d, want 3 — the units still exist", c.TotalUnits)
	}
	if got := strings.Join(c.UnsupportedLanguages, ","); got != "swift" {
		t.Errorf("UnsupportedLanguages = %v, want [swift]", c.UnsupportedLanguages)
	}
	if c.PassRateKnown() {
		t.Error("PassRateKnown() = true, want false — 0/0 is undefined, not a rate")
	}
}

// TestGenerateCard_MixedExcludesUnassessedFromBothSides checks that the
// unassessed units leave the numerator AND the denominator, rather than moving
// from the passing bucket to the failing one.
func TestGenerateCard_MixedExcludesUnassessedFromBothSides(t *testing.T) {
	records := []domain.CertificationRecord{
		makeCardRecord("go", "a.go", "A", domain.StatusCertified, 0.90),
		makeCardRecord("go", "b.go", "B", domain.StatusDecertified, 0.20),
		unassessedRecord("swift", "c.swift", "c"),
		unassessedRecord("rb", "d.rb", "d"),
	}

	c := report.GenerateCard(records, "test/repo", "abc", time.Now())

	if c.TotalUnits != 4 {
		t.Errorf("TotalUnits = %d, want 4", c.TotalUnits)
	}
	if c.AnalyzableUnits != 2 {
		t.Errorf("AnalyzableUnits = %d, want 2", c.AnalyzableUnits)
	}
	if c.Passing != 1 {
		t.Errorf("Passing = %d, want 1", c.Passing)
	}
	if c.Failing != 1 {
		t.Errorf("Failing = %d, want 1", c.Failing)
	}
	if c.UnsupportedCount != 2 {
		t.Errorf("UnsupportedCount = %d, want 2", c.UnsupportedCount)
	}
	if c.PassRate != 0.5 {
		t.Errorf("PassRate = %.3f, want 0.500 (1 of 2 analyzable)", c.PassRate)
	}
	if !c.PassRateKnown() {
		t.Error("PassRateKnown() = false, want true — two units were assessed")
	}
	want := []string{"rb", "swift"}
	if strings.Join(c.UnsupportedLanguages, ",") != strings.Join(want, ",") {
		t.Errorf("UnsupportedLanguages = %v, want %v (sorted)", c.UnsupportedLanguages, want)
	}
}

// TestGenerateCard_OperatorExemptStillPasses guards the distinction the
// Unsupported flag exists to carry: a unit exempted by explicit override was
// assessed — a human asserted a verdict about it — and stays in the numerator.
func TestGenerateCard_OperatorExemptStillPasses(t *testing.T) {
	records := []domain.CertificationRecord{
		makeCardRecord("go", "a.go", "A", domain.StatusExempt, 0.0),
		makeCardRecord("go", "b.go", "B", domain.StatusCertified, 0.90),
	}

	c := report.GenerateCard(records, "test/repo", "abc", time.Now())

	if c.UnsupportedCount != 0 {
		t.Errorf("UnsupportedCount = %d, want 0 — neither record is unsupported", c.UnsupportedCount)
	}
	if c.Passing != 2 {
		t.Errorf("Passing = %d, want 2 — an operator exemption is a verdict", c.Passing)
	}
	if c.PassRate != 1.0 {
		t.Errorf("PassRate = %.3f, want 1.000", c.PassRate)
	}
}

// TestFormatCardText_AllUnassessedRendersNoPassRate asserts on the rendered
// text card, not the struct.
func TestFormatCardText_AllUnassessedRendersNoPassRate(t *testing.T) {
	records := []domain.CertificationRecord{
		unassessedRecord("swift", "a.swift", "a"),
		unassessedRecord("swift", "b.swift", "b"),
		unassessedRecord("swift", "c.swift", "c"),
	}
	out := report.FormatCardText(report.GenerateCard(records, "r", "abc", time.Now()))

	if strings.Contains(out, "Pass Rate:   100.0%") {
		t.Errorf("text card claims a 100%% pass rate for a repo that was never assessed:\n%s", out)
	}
	if !strings.Contains(out, "Pass Rate:      n/a") {
		t.Errorf("text card should render the pass rate as n/a, got:\n%s", out)
	}
	if !strings.Contains(out, "Not Assessed") {
		t.Errorf("text card should report the unassessed count, got:\n%s", out)
	}
	// The N/A grade rows exist in the distribution and must be rendered.
	if !strings.Contains(out, "N/A:    3") {
		t.Errorf("grade distribution should render the N/A row, got:\n%s", out)
	}
}

// TestFormatCardMarkdown_AllUnassessedRendersNoPassRate asserts on the rendered
// markdown card — this is what lands in .certification/REPORT_CARD.md.
func TestFormatCardMarkdown_AllUnassessedRendersNoPassRate(t *testing.T) {
	records := []domain.CertificationRecord{
		unassessedRecord("swift", "a.swift", "a"),
		unassessedRecord("swift", "b.swift", "b"),
		unassessedRecord("swift", "c.swift", "c"),
	}
	out := report.FormatCardMarkdown(report.GenerateCard(records, "r", "abc", time.Now()))

	if strings.Contains(out, "| Pass Rate | 100.0% |") {
		t.Errorf("markdown card claims a 100%% pass rate for an unassessed repo:\n%s", out)
	}
	if !strings.Contains(out, "| Pass Rate | n/a |") {
		t.Errorf("markdown card should render the pass rate as n/a, got:\n%s", out)
	}
	if !strings.Contains(out, "| Not Assessed | 3 |") {
		t.Errorf("markdown card should report the unassessed count, got:\n%s", out)
	}
	if !strings.Contains(out, "| N/A | 3 |") {
		t.Errorf("markdown grade distribution should render the N/A row, got:\n%s", out)
	}
}

// TestGenerateBadge_AllUnassessedDoesNotClaim100 asserts on the rendered badge
// message. This is the most publicly visible artifact the product has: it is
// committed to .certification/badge.json and served into the README.
func TestGenerateBadge_AllUnassessedDoesNotClaim100(t *testing.T) {
	records := []domain.CertificationRecord{
		unassessedRecord("swift", "a.swift", "a"),
		unassessedRecord("swift", "b.swift", "b"),
		unassessedRecord("swift", "c.swift", "c"),
	}
	b := report.GenerateBadge(report.GenerateCard(records, "r", "abc", time.Now()))

	if strings.Contains(b.Message, "100%") {
		t.Errorf("badge message = %q, must not claim 100%% for an unassessed repo", b.Message)
	}
	if strings.Contains(b.Message, "0%") {
		t.Errorf("badge message = %q, must not claim 0%% either — nothing was measured", b.Message)
	}
	if !strings.Contains(b.Message, "not assessed") {
		t.Errorf("badge message = %q, want it to say the repo was not assessed", b.Message)
	}
	if b.Color != "9CA3AF" {
		t.Errorf("badge color = %q, want the neutral 9CA3AF — a red badge asserts a verdict", b.Color)
	}
}

// TestGenerateBadge_AssessedRepoUnchanged pins the ordinary path so the fix
// above cannot quietly change the badge for a repo that was actually assessed.
func TestGenerateBadge_AssessedRepoUnchanged(t *testing.T) {
	records := []domain.CertificationRecord{
		makeCardRecord("go", "a.go", "A", domain.StatusCertified, 0.90),
		makeCardRecord("go", "b.go", "B", domain.StatusCertified, 0.90),
	}
	b := report.GenerateBadge(report.GenerateCard(records, "r", "abc", time.Now()))

	if b.Message != "A- · 100% · 2 units" {
		t.Errorf("badge message = %q, want %q", b.Message, "A- · 100% · 2 units")
	}
}

// TestFormatCardText_HistogramRowsAlign pins the histogram row width. The bar
// was padded with %-36s against a 40-column field, and %-Ns pads to a width in
// BYTES while █ is three bytes — so the pad never applied at any width and
// every bar row was misaligned. The N/A label is a third column wide in a %2s
// field, which costs the bar one more.
//
// Scoped to the bar rows on purpose: other rows of this box (the title, the
// language rows, the top-issue rows) are independently misaligned, several
// because an emoji is one rune but two display columns. That is pre-existing
// and separate from this change.
func TestFormatCardText_HistogramRowsAlign(t *testing.T) {
	records := []domain.CertificationRecord{
		makeCardRecord("go", "a.go", "A", domain.StatusCertified, 0.95),
		makeCardRecord("go", "b.go", "B", domain.StatusDecertified, 0.20),
		unassessedRecord("swift", "c.swift", "c"),
	}
	out := report.FormatCardText(report.GenerateCard(records, "r", "abc", time.Now()))

	const boxWidth = 64 // the ╔══…╗ border, in display columns
	rows := 0
	for _, line := range strings.Split(out, "\n") {
		if !histogramRow.MatchString(line) {
			continue
		}
		rows++
		if got := len([]rune(line)); got != boxWidth {
			t.Errorf("histogram row is %d columns, want %d:\n%q", got, boxWidth, line)
		}
	}
	if rows != 3 {
		t.Fatalf("matched %d histogram rows, want 3 (A, F and N/A)", rows)
	}
}

// histogramRow matches "║    A:    1 ( 33.3%) ███…" and nothing else in the box.
var histogramRow = regexp.MustCompile(`^║ {4} *[A-Z/+-]{1,3}: +\d+ \( *\d+\.\d%\) `)
