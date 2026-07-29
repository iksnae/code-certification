package report_test

import (
	"math"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/report"
)

// A pass rate and an average score are the same kind of claim: a ratio whose
// denominator counts the units about which a verdict was asserted. #47 fixed the
// rate and left the score beside it, so an artifact could carry `Pass Rate: n/a`
// and `🔴 F (0.0%)` at once — one saying nothing was measured, the other stating
// a definite failure over the same units.
//
// The rule these tests pin: an unassessed unit contributes a placeholder zero,
// not a low score, so a mean over nothing but placeholders is undefined. It is
// rendered as N/A and n/a, never as F and 0.0%.
//
// This file covers the all-unassessed case, where every rate short-circuits. The
// MIXED case — where the score and the pass rate are both defined and can
// therefore disagree — is in mixed_denominator_test.go, and is the half in which
// #47's own definition of the defect kept reproducing after this file was green.

// unsupportedPackageRecords is the fixture whose absence let three mutants live.
// The suite had an all-unassessed *repo* and a *mixed* package, but no package
// in which every unit is unassessed — the only shape where a package-level
// aggregate is asked to summarise zero measurements.
func unsupportedPackageRecords() []domain.CertificationRecord {
	return []domain.CertificationRecord{
		unassessedRecord("swift", "app/a.swift", "a"),
		unassessedRecord("swift", "app/b.swift", "b"),
		unassessedRecord("swift", "app/c.swift", "c"),
	}
}

func unsupportedPackageReport() report.FullReport {
	return report.GenerateFullReport(unsupportedPackageRecords(), "test/repo", "abc", time.Now())
}

// --- Card -------------------------------------------------------------------

// TestGenerateCard_AllUnassessedHasNoOverallGrade covers the headline defect:
// `## 🔴 Overall: F (0.0%)` printed directly above `| Pass Rate | n/a |`.
func TestGenerateCard_AllUnassessedHasNoOverallGrade(t *testing.T) {
	c := report.GenerateCard(unsupportedPackageRecords(), "test/repo", "abc", time.Now())

	if c.ScoreKnown() {
		t.Error("ScoreKnown() = true, want false — no unit was scored")
	}
	if c.OverallGrade != "N/A" {
		t.Errorf("OverallGrade = %q, want %q — an F is a verdict about code the engine never opened", c.OverallGrade, "N/A")
	}
	if c.OverallScore != 0 {
		t.Errorf("OverallScore = %v, want the zero value left untouched when unknown", c.OverallScore)
	}
}

// TestFormatCardMarkdown_AllUnassessedGradeAndScore asserts on the artifact
// that lands in .certification/REPORT_CARD.md.
func TestFormatCardMarkdown_AllUnassessedGradeAndScore(t *testing.T) {
	fr := unsupportedPackageReport()
	fr.Card.Packages = report.BuildPackageSummaries(fr)
	out := report.FormatCardMarkdown(fr.Card)

	if strings.Contains(out, "Overall: F") {
		t.Errorf("markdown card grades an unassessed repo F:\n%s", out)
	}
	if !strings.Contains(out, "## ⚪ Overall: N/A (n/a)") {
		t.Errorf("markdown card should render the overall grade as N/A, got:\n%s", out)
	}
	if strings.Contains(out, "🔴") {
		t.Errorf("markdown card renders a failing marker for a repo that was never assessed:\n%s", out)
	}
	if !strings.Contains(out, "| swift | 3 | ⚪ N/A | n/a |") {
		t.Errorf("by-language row should be N/A over an unassessed language, got:\n%s", out)
	}
	if !strings.Contains(out, "| [app](reports/app/index.md) | 3 | ⚪ N/A | n/a |") {
		t.Errorf("packages row should be N/A over an unassessed package, got:\n%s", out)
	}
}

// TestFormatCardText_AllUnassessedGradeAndScore asserts on the terminal card.
func TestFormatCardText_AllUnassessedGradeAndScore(t *testing.T) {
	out := report.FormatCardText(report.GenerateCard(unsupportedPackageRecords(), "r", "abc", time.Now()))

	if strings.Contains(out, "Score: 0.0") {
		t.Errorf("text card prints a 0.0%% score for a repo that was never assessed:\n%s", out)
	}
	if !strings.Contains(out, "Overall Grade:  ⚪ N/A      Score: n/a") {
		t.Errorf("text card should render grade N/A and score n/a, got:\n%s", out)
	}
	if !strings.Contains(out, "swift           3 units   ⚪ N/A    (n/a)") {
		t.Errorf("text card by-language row should be N/A, got:\n%s", out)
	}
}

// --- Package ----------------------------------------------------------------

// TestBuildPackageSummaries_AllUnsupportedPackage is the missing fixture. It
// pins the pass rate as a finite non-claim rather than the NaN that
// `analyzable() >= 0` would produce, and the grade as N/A rather than F.
func TestBuildPackageSummaries_AllUnsupportedPackage(t *testing.T) {
	pkgs := report.BuildPackageSummaries(unsupportedPackageReport())
	if len(pkgs) != 1 {
		t.Fatalf("package count = %d, want 1", len(pkgs))
	}
	p := pkgs[0]

	if p.Units != 3 {
		t.Errorf("Units = %d, want 3", p.Units)
	}
	if p.Unsupported != 3 {
		t.Errorf("Unsupported = %d, want 3", p.Unsupported)
	}
	if p.Analyzable() != 0 {
		t.Errorf("Analyzable() = %d, want 0", p.Analyzable())
	}
	if p.PassRateKnown() {
		t.Error("PassRateKnown() = true, want false — 0/0 is undefined")
	}
	if math.IsNaN(p.PassRate) {
		t.Error("PassRate is NaN; the guard on the denominator is gone and every renderer will print NaN%")
	}
	if p.PassRate != 0 {
		t.Errorf("PassRate = %v, want the zero value left untouched when unknown", p.PassRate)
	}
	if p.ScoreKnown() {
		t.Error("ScoreKnown() = true, want false — no unit in this package was scored")
	}
	if p.Grade != "N/A" {
		t.Errorf("Grade = %q, want %q", p.Grade, "N/A")
	}
	if math.IsNaN(p.AvgScore) {
		t.Error("AvgScore is NaN")
	}
	if p.AvgScore != 0 {
		t.Errorf("AvgScore = %v, want the zero value left untouched when unknown", p.AvgScore)
	}
}

// TestFormatPackageIndexMarkdown_AllUnsupportedPackage asserts on the committed
// package page. `**Passing:** 0 / 0` is the exact string the guard exists to
// prevent, and the grade line above it must not read F either.
func TestFormatPackageIndexMarkdown_AllUnsupportedPackage(t *testing.T) {
	dir := t.TempDir()
	if _, err := report.GenerateReportTree(unsupportedPackageReport(), dir); err != nil {
		t.Fatalf("GenerateReportTree() error: %v", err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "app", "index.md"))
	if err != nil {
		t.Fatalf("app/index.md not found: %v", err)
	}
	md := string(data)

	if strings.Contains(md, "**Passing:**") {
		t.Errorf("package page states a passing count over zero analyzable units:\n%s", md)
	}
	if !strings.Contains(md, "**Units:** 3 · **Not Assessed:** 3") {
		t.Errorf("package page should state 3 units, all unassessed, got:\n%s", md)
	}
	if !strings.Contains(md, "**Grade:** ⚪ N/A (n/a)") {
		t.Errorf("package page should render grade N/A and score n/a, got:\n%s", md)
	}
	if strings.Contains(md, "🔴") {
		t.Errorf("package page renders a failing marker for units that were never opened:\n%s", md)
	}
}

// TestFormatReportTreeIndex_AllUnsupportedPackage covers the top-level index —
// both the summary line and the packages table.
func TestFormatReportTreeIndex_AllUnsupportedPackage(t *testing.T) {
	dir := t.TempDir()
	if _, err := report.GenerateReportTree(unsupportedPackageReport(), dir); err != nil {
		t.Fatalf("GenerateReportTree() error: %v", err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "index.md"))
	if err != nil {
		t.Fatalf("index.md not found: %v", err)
	}
	md := string(data)

	if !strings.Contains(md, "**Units:** 3 · **Passing:** 0 · **Failing:** 0 · **Not Assessed:** 3") {
		t.Errorf("index summary line should carry the unassessed count, got:\n%s", md)
	}
	if !strings.Contains(md, "**Overall:** ⚪ N/A (n/a)") {
		t.Errorf("index should render the overall grade as N/A, got:\n%s", md)
	}
	row := findLine(md, "| [app](app/index.md) |")
	if row == "" {
		t.Fatalf("no packages row for app in:\n%s", md)
	}
	if !strings.Contains(row, "| 3 | 3 | ⚪ N/A | n/a | n/a |") {
		t.Errorf("packages row = %q, want 3 units all unassessed with no grade, score or rate", row)
	}
}

// TestGeneratePackagePages_AllUnsupportedPackage covers the HTML package page,
// which reads the same aggregation a third time.
func TestGeneratePackagePages_AllUnsupportedPackage(t *testing.T) {
	dir := t.TempDir()
	if err := report.GenerateSite(unsupportedPackageReport(), report.SiteConfig{
		OutputDir: dir,
		Title:     "test",
	}); err != nil {
		t.Fatalf("GenerateSite() error: %v", err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "packages", "app", "index.html"))
	if err != nil {
		t.Fatalf("package page not found: %v", err)
	}
	html := string(data)

	if !strings.Contains(html, `<div class="value">N/A</div><div class="label">Grade</div>`) {
		t.Errorf("package page should render grade N/A, got:\n%s", html)
	}
	if !strings.Contains(html, `<div class="value">n/a</div><div class="label">Avg Score</div>`) {
		t.Errorf("package page should render the average score as n/a, got:\n%s", html)
	}
	if !strings.Contains(html, `<div class="value">n/a</div><div class="label">Pass Rate</div>`) {
		t.Errorf("package page should render the pass rate as n/a, got:\n%s", html)
	}
	if !strings.Contains(html, `<div class="value">3</div><div class="label">Not Assessed</div>`) {
		t.Errorf("package page should render the unassessed count, got:\n%s", html)
	}
}

// TestUnitScore_UnassessedUnitHasNoScore covers the per-unit score column,
// which is the same placeholder one level below the aggregates. A unit cert
// reading `Grade: N/A` beside `Score: 0.0%` says both "nothing was measured" and
// "the measurement was zero"; the four surfaces that render a unit's score all
// have to agree with the grade printed next to it.
func TestUnitScore_UnassessedUnitHasNoScore(t *testing.T) {
	fr := unsupportedPackageReport()

	treeDir := t.TempDir()
	if _, err := report.GenerateReportTree(fr, treeDir); err != nil {
		t.Fatalf("GenerateReportTree() error: %v", err)
	}
	cert, err := os.ReadFile(filepath.Join(treeDir, "app", "a.swift", "a.md"))
	if err != nil {
		t.Fatalf("unit cert not found: %v", err)
	}
	if !strings.Contains(string(cert), "| **Score** | n/a |") {
		t.Errorf("unit cert should render the score as n/a, got:\n%s", cert)
	}

	pkgIndex, err := os.ReadFile(filepath.Join(treeDir, "app", "index.md"))
	if err != nil {
		t.Fatalf("package index not found: %v", err)
	}
	if !strings.Contains(string(pkgIndex), "| ⚪ N/A | n/a | exempt |") {
		t.Errorf("package units table should render the score as n/a, got:\n%s", pkgIndex)
	}

	if out := report.FormatFullMarkdown(fr); !strings.Contains(out, "| function | N/A | n/a | exempt |") {
		t.Errorf("full report units table should render the score as n/a, got:\n%s", out)
	}

	siteDir := t.TempDir()
	if err := report.GenerateSite(fr, report.SiteConfig{OutputDir: siteDir, Title: "test"}); err != nil {
		t.Fatalf("GenerateSite() error: %v", err)
	}
	unitPage, err := os.ReadFile(filepath.Join(siteDir, "units", "app-a-swift-a.html"))
	if err != nil {
		t.Fatalf("unit page not found: %v", err)
	}
	if !strings.Contains(string(unitPage), "<tr><td>Score</td><td>n/a</td></tr>") {
		t.Errorf("unit page should render the score as n/a, got:\n%s", unitPage)
	}
	pkgPage, err := os.ReadFile(filepath.Join(siteDir, "packages", "app", "index.html"))
	if err != nil {
		t.Fatalf("package page not found: %v", err)
	}
	if !strings.Contains(string(pkgPage), "<span class=\"grade grade-na\">N/A</span></td>\n<td>n/a</td>") {
		t.Errorf("package page units table should render the score as n/a, got:\n%s", pkgPage)
	}
}

// TestUnitScore_AssessedUnitKeepsItsScore pins the other side, so the gate
// above cannot be widened into hiding real scores.
func TestUnitScore_AssessedUnitKeepsItsScore(t *testing.T) {
	treeDir := t.TempDir()
	if _, err := report.GenerateReportTree(mixedPackageReport(), treeDir); err != nil {
		t.Fatalf("GenerateReportTree() error: %v", err)
	}
	cert, err := os.ReadFile(filepath.Join(treeDir, "app", "a.go", "A.md"))
	if err != nil {
		t.Fatalf("unit cert not found: %v", err)
	}
	if !strings.Contains(string(cert), "| **Score** | 90.0% |") {
		t.Errorf("an assessed unit must keep its measured score, got:\n%s", cert)
	}
}

// TestGenerateSiteIndex_AllUnassessed covers the HTML dashboard.
func TestGenerateSiteIndex_AllUnassessed(t *testing.T) {
	dir := t.TempDir()
	if err := report.GenerateSite(unsupportedPackageReport(), report.SiteConfig{
		OutputDir: dir,
		Title:     "test",
	}); err != nil {
		t.Fatalf("GenerateSite() error: %v", err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "index.html"))
	if err != nil {
		t.Fatalf("index.html not found: %v", err)
	}
	html := string(data)

	if !strings.Contains(html, `<div class="value">N/A</div><div class="label">Overall Grade</div>`) {
		t.Errorf("site index should render grade N/A, got:\n%s", html)
	}
	if !strings.Contains(html, `<div class="value">n/a</div><div class="label">Overall Score</div>`) {
		t.Errorf("site index should render the overall score as n/a, got:\n%s", html)
	}
	// The dashboard's Packages table was a fourth, independent aggregation. It
	// now reads the shared one; asserting its rendered grade is what keeps it
	// from drifting back. Scoped to the Packages section on purpose — the
	// Languages table one heading above renders an identical cell pair, and an
	// unscoped match is satisfied by the wrong row.
	pkgSection := sectionAfter(html, "<h2>Packages</h2>")
	if pkgSection == "" {
		t.Fatalf("no Packages section in:\n%s", html)
	}
	if !strings.Contains(pkgSection, "<td><span class=\"grade grade-na\">N/A</span></td>\n<td>n/a</td>") {
		t.Errorf("site index packages table should render N/A and n/a, got:\n%s", pkgSection)
	}
}

// sectionAfter returns the HTML between marker and the next <h2>, so an
// assertion about one table cannot be satisfied by an identical cell in
// another.
func sectionAfter(html, marker string) string {
	i := strings.Index(html, marker)
	if i < 0 {
		return ""
	}
	rest := html[i+len(marker):]
	if j := strings.Index(rest, "<h2>"); j >= 0 {
		return rest[:j]
	}
	return rest
}

// TestGenerateSiteIndex_MixedPackageAgreesWithTheOtherSurfaces pins the site
// dashboard's Packages row against the markdown one. The two used to be
// computed separately, and the assertion that matters is that they agree.
func TestGenerateSiteIndex_MixedPackageAgreesWithTheOtherSurfaces(t *testing.T) {
	fr := mixedPackageReport()
	dir := t.TempDir()
	if err := report.GenerateSite(fr, report.SiteConfig{OutputDir: dir, Title: "test"}); err != nil {
		t.Fatalf("GenerateSite() error: %v", err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "index.html"))
	if err != nil {
		t.Fatalf("index.html not found: %v", err)
	}
	html := string(data)

	pkgs := report.BuildPackageSummaries(fr)
	if len(pkgs) != 1 {
		t.Fatalf("package count = %d, want 1", len(pkgs))
	}
	want := "<td><span class=\"grade grade-" +
		strings.ReplaceAll(strings.ToLower(pkgs[0].Grade), "+", "plus") +
		"\">" + pkgs[0].Grade + "</span></td>\n<td>" +
		report.FormatRate(pkgs[0].ScoreKnown(), pkgs[0].AvgScore, 1) + "</td>"
	pkgSection := sectionAfter(html, "<h2>Packages</h2>")
	if pkgSection == "" {
		t.Fatalf("no Packages section in:\n%s", html)
	}
	if !strings.Contains(pkgSection, want) {
		t.Errorf("site packages row disagrees with BuildPackageSummaries (%s / %.1f%%), got:\n%s",
			pkgs[0].Grade, pkgs[0].AvgScore*100, pkgSection)
	}
}

// --- Language ---------------------------------------------------------------

// TestBuildLanguageDetail_UnassessedLanguageHasNoGrade covers the `swift | 3 |
// 🔴 F | 0.0%` row. The language accumulator summed every record's score before
// the unsupported branch, so three unopened units averaged to a confident F.
func TestBuildLanguageDetail_UnassessedLanguageHasNoGrade(t *testing.T) {
	fr := unsupportedPackageReport()
	if len(fr.LanguageDetail) != 1 {
		t.Fatalf("language count = %d, want 1", len(fr.LanguageDetail))
	}
	l := fr.LanguageDetail[0]

	if l.Name != "swift" {
		t.Fatalf("language = %q, want swift", l.Name)
	}
	if l.Units != 3 {
		t.Errorf("Units = %d, want 3", l.Units)
	}
	if l.Unsupported != 3 {
		t.Errorf("Unsupported = %d, want 3", l.Unsupported)
	}
	if l.Analyzable() != 0 {
		t.Errorf("Analyzable() = %d, want 0", l.Analyzable())
	}
	if l.ScoreKnown() {
		t.Error("ScoreKnown() = true, want false")
	}
	if l.Grade != "N/A" {
		t.Errorf("Grade = %q, want %q", l.Grade, "N/A")
	}
	if l.AverageScore != 0 {
		t.Errorf("AverageScore = %v, want the zero value left untouched when unknown", l.AverageScore)
	}
}

// TestBuildLanguageDetail_MixedKeepsTheAssessedLanguage guards the other side:
// a language with assessed units keeps its real grade, and only the unassessed
// language loses one.
func TestBuildLanguageDetail_MixedKeepsTheAssessedLanguage(t *testing.T) {
	d := report.Detailed(mixedPackageRecords(), time.Now())

	swift := d.ByLanguage["swift"]
	if swift.Grade != "N/A" || swift.ScoreKnown() {
		t.Errorf("swift Grade = %q ScoreKnown = %v, want N/A and false", swift.Grade, swift.ScoreKnown())
	}
	goLang := d.ByLanguage["go"]
	if goLang.Grade != "A-" {
		t.Errorf("go Grade = %q, want A- — an assessed language keeps its measured grade", goLang.Grade)
	}
	if goLang.AverageScore != 0.90 {
		t.Errorf("go AverageScore = %v, want 0.90", goLang.AverageScore)
	}
	if !goLang.ScoreKnown() {
		t.Error("go ScoreKnown() = false, want true")
	}
}

// TestFormatDetailedText_UnassessedLanguageHasNoAverage covers `certify report
// --detailed`, which printed `avg 0.000` beside `(3 not assessed)`.
func TestFormatDetailedText_UnassessedLanguageHasNoAverage(t *testing.T) {
	out := report.FormatDetailedText(report.Detailed(unsupportedPackageRecords(), time.Now()))

	if strings.Contains(out, "avg 0.000") {
		t.Errorf("detailed text prints an average of 0.000 over units that were never scored:\n%s", out)
	}
	if !strings.Contains(out, "avg n/a (3 not assessed)") {
		t.Errorf("detailed text should render the average as n/a, got:\n%s", out)
	}
}

// TestFormatFullMarkdown_AllUnassessed covers `certify report --format full`.
func TestFormatFullMarkdown_AllUnassessed(t *testing.T) {
	out := report.FormatFullMarkdown(unsupportedPackageReport())

	if strings.Contains(out, "**Overall Grade** | 🔴 **F**") {
		t.Errorf("full report grades an unassessed repo F:\n%s", out)
	}
	if !strings.Contains(out, "| **Overall Grade** | ⚪ **N/A** |") {
		t.Errorf("full report should render the overall grade as N/A, got:\n%s", out)
	}
	if !strings.Contains(out, "| **Overall Score** | n/a |") {
		t.Errorf("full report should render the overall score as n/a, got:\n%s", out)
	}
}

// --- The score denominator, moved (was the #32 fence) ------------------------

// These two tests used to PIN the whole-unit denominator, on the reading that
// changing it was issue #32 and out of scope here. That reading was wrong, and
// the reason is worth keeping: #47's defect is a card that contradicts itself,
// and over a mixed corpus it still did — `## 🔴 Overall: F (53.4%)` above
// `| Pass Rate | 100.0% |`. Leaving the score denominator alone closed the
// all-unassessed case and left the defect's own definition reproducing. So the
// flip is the remainder of #47, not a follow-up to it.
//
// What the fence got right is that a PARTIAL flip is worse than none: the card,
// the package aggregate, the language aggregate and the workspace rollup all
// print into the same artifacts, and any one of them moving alone relocates the
// contradiction rather than removing it. They are asserted together in
// mixed_denominator_test.go; these two remain as the record of the reversal, so
// a future reader finds the reasoning rather than an unexplained flip.

// TestBuildPackageSummaries_MixedPackageScoreExcludesUnassessed pins the package
// aggregate at the analyzable mean.
func TestBuildPackageSummaries_MixedPackageScoreExcludesUnassessed(t *testing.T) {
	pkgs := report.BuildPackageSummaries(mixedPackageReport())
	if len(pkgs) != 1 {
		t.Fatalf("package count = %d, want 1", len(pkgs))
	}
	p := pkgs[0]

	if !p.ScoreKnown() {
		t.Fatal("ScoreKnown() = false, want true — one unit was assessed")
	}
	if math.Abs(p.AvgScore-0.90) > 1e-9 {
		t.Errorf("AvgScore = %v, want 0.90 — the mean over the analyzable unit. 0.30 divides by all three, "+
			"grading a package F on the strength of two units that were never opened", p.AvgScore)
	}
	if p.Grade != "A-" {
		t.Errorf("Grade = %q, want %q", p.Grade, "A-")
	}
}

// TestGenerateCard_MixedOverallScoreExcludesUnassessed is the same at the card
// level — the figure printed directly above the pass rate it must agree with.
func TestGenerateCard_MixedOverallScoreExcludesUnassessed(t *testing.T) {
	c := report.GenerateCard(mixedPackageRecords(), "test/repo", "abc", time.Now())

	if !c.ScoreKnown() {
		t.Fatal("ScoreKnown() = false, want true — one unit was assessed")
	}
	if math.Abs(c.OverallScore-0.90) > 1e-9 {
		t.Errorf("OverallScore = %v, want 0.90 — the same denominator PassRate uses", c.OverallScore)
	}
	if c.OverallGrade != "A-" {
		t.Errorf("OverallGrade = %q, want %q", c.OverallGrade, "A-")
	}
}
