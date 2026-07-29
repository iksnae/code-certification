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

// #47's defect is that one box states two contradicting things about the same
// units. The all-unassessed half was closed by ScoreKnown(); this file closes the
// MIXED half, which is where the contradiction is actually visible:
//
//	## 🔴 Overall: F (53.4%)
//	| Pass Rate | 100.0% |     | Failing | 0 |
//
// Both figures summarise the same units, and they disagree because they use
// different denominators. The pass rate already excludes unassessed units from
// both sides; the score still divides by every unit, so each unopened unit drags
// the mean toward zero as though the engine had measured a failure there. An
// unassessed unit contributes a placeholder, not a data point — it belongs in
// neither side of either ratio.
//
// The three denominators move together or not at all. A partial flip produces the
// mirror contradiction — `Overall: A-` above `| app | 3 | 🔴 F |` — which is the
// same defect one table lower, so the card, the package/language aggregates and
// the workspace rollup are all asserted here.

// mixedCorpus is one Go unit certified at 0.90 and two unassessed Swift units.
// Analyzable mean = 0.90 (A-). Whole-unit mean = 0.30 (F). The pass rate is
// 100.0% either way, so the whole-unit mean is what puts an F beside it.
func mixedCorpus() []domain.CertificationRecord { return mixedPackageRecords() }

// TestGenerateCard_MixedScoreExcludesUnassessedUnits is the reddening test at the
// card level: the grade and the pass rate must summarise the same denominator.
func TestGenerateCard_MixedScoreExcludesUnassessedUnits(t *testing.T) {
	c := report.GenerateCard(mixedCorpus(), "test/repo", "abc", time.Now())

	if !c.ScoreKnown() {
		t.Fatal("ScoreKnown() = false, want true — one unit was assessed")
	}
	if c.AnalyzableUnits != 1 {
		t.Fatalf("AnalyzableUnits = %d, want 1", c.AnalyzableUnits)
	}
	if math.Abs(c.OverallScore-0.90) > 1e-9 {
		t.Errorf("OverallScore = %v, want 0.90 — the mean over the ONE unit that was assessed. "+
			"Dividing by all 3 units gives 0.30 and grades the repo F beside a 100%% pass rate", c.OverallScore)
	}
	if c.OverallGrade != "A-" {
		t.Errorf("OverallGrade = %q, want %q — an F here is a verdict about two units the engine never opened",
			c.OverallGrade, "A-")
	}
	if c.PassRate != 1.0 {
		t.Errorf("PassRate = %v, want 1.0 — unchanged; the score is what moves to meet it", c.PassRate)
	}
}

// TestFormatCardMarkdown_MixedHasNoFBesideAPerfectPassRate asserts on the
// delivered artifact, in the exact terms #47's body uses. The grade line and the
// Pass Rate row sit in one box; a reader compares them.
func TestFormatCardMarkdown_MixedHasNoFBesideAPerfectPassRate(t *testing.T) {
	fr := mixedPackageReport()
	fr.Card.Packages = report.BuildPackageSummaries(fr)
	out := report.FormatCardMarkdown(fr.Card)

	if strings.Contains(out, "Overall: F") && strings.Contains(out, "| Pass Rate | 100.0% |") {
		t.Errorf("REPORT_CARD.md grades the repo F beside a 100.0%% pass rate — false in both directions "+
			"inside the same box:\n%s", out)
	}
	if !strings.Contains(out, "## 🟢 Overall: A- (90.0%)") {
		t.Errorf("overall line should summarise the assessed unit, got:\n%s", out)
	}
	if !strings.Contains(out, "| Pass Rate | 100.0% |") {
		t.Errorf("pass rate should stay 100.0%% over the single analyzable unit, got:\n%s", out)
	}
	if !strings.Contains(out, "| Analyzable Units | 1 |") {
		t.Errorf("card should state the denominator it used, got:\n%s", out)
	}
	// The tables below the headline must agree with it, or the contradiction is
	// simply relocated.
	if !strings.Contains(out, "| go | 1 | 🟢 A- | 90.0% |") {
		t.Errorf("by-language row for the assessed language should read A-/90.0%%, got:\n%s", out)
	}
	if !strings.Contains(out, "| swift | 2 | ⚪ N/A | n/a |") {
		t.Errorf("by-language row for the unassessed language should read N/A, got:\n%s", out)
	}
	if !strings.Contains(out, "| [app](reports/app/index.md) | 3 | 🟢 A- | 90.0% |") {
		t.Errorf("packages row should agree with the overall grade, got:\n%s", out)
	}
}

// TestFormatCardText_MixedHasNoFBesideAPerfectPassRate is the terminal card, the
// surface an operator actually reads during a run.
func TestFormatCardText_MixedHasNoFBesideAPerfectPassRate(t *testing.T) {
	out := report.FormatCardText(report.GenerateCard(mixedCorpus(), "r", "abc", time.Now()))

	if strings.Contains(out, "🔴 F") {
		t.Errorf("text card renders a failing marker for a repo whose only assessed unit scored 0.90:\n%s", out)
	}
	if !strings.Contains(out, "Overall Grade:  🟢 A-       Score: 90.0%") {
		t.Errorf("text card should render A- / 90.0%%, got:\n%s", out)
	}
	if !strings.Contains(out, "Pass Rate:   100.0%") {
		t.Errorf("text card should keep the 100.0%% pass rate, got:\n%s", out)
	}
	if !strings.Contains(out, "Not Assessed:    2         Analyzable:      1") {
		t.Errorf("text card should state the denominator it used, got:\n%s", out)
	}
}

// TestBuildPackageSummaries_MixedScoreExcludesUnassessedUnits is the same flip at
// the package aggregate, which feeds the packages table and both package pages.
func TestBuildPackageSummaries_MixedScoreExcludesUnassessedUnits(t *testing.T) {
	pkgs := report.BuildPackageSummaries(mixedPackageReport())
	if len(pkgs) != 1 {
		t.Fatalf("package count = %d, want 1", len(pkgs))
	}
	p := pkgs[0]

	if !p.ScoreKnown() {
		t.Fatal("ScoreKnown() = false, want true — one unit was assessed")
	}
	if math.Abs(p.AvgScore-0.90) > 1e-9 {
		t.Errorf("AvgScore = %v, want 0.90 — the mean over the analyzable unit, matching Card.OverallScore. "+
			"0.30 would put 🔴 F in the packages table under a 🟢 A- headline", p.AvgScore)
	}
	if p.Grade != "A-" {
		t.Errorf("Grade = %q, want %q", p.Grade, "A-")
	}
	if p.PassRate != 1.0 {
		t.Errorf("PassRate = %v, want 1.0 — unchanged", p.PassRate)
	}
}

// TestBuildLanguageDetail_MixedLanguageScoreExcludesUnassessedUnits pins the
// language aggregate, the third denominator. It is asserted on a language that
// holds BOTH kinds of unit, because a language whose units are uniformly
// unassessed cannot tell the two denominators apart.
func TestBuildLanguageDetail_MixedLanguageScoreExcludesUnassessedUnits(t *testing.T) {
	// One assessed Go unit at 0.80 and one unassessed Go unit. Whole-unit mean
	// 0.40 (F); analyzable mean 0.80 (B).
	unassessedGo := unassessedRecord("go", "app/skipped.go", "Skipped")
	fr := report.GenerateFullReport([]domain.CertificationRecord{
		makeCardRecord("go", "app/a.go", "A", domain.StatusCertified, 0.80),
		unassessedGo,
	}, "test/repo", "abc", time.Now())

	if len(fr.LanguageDetail) != 1 {
		t.Fatalf("language count = %d, want 1", len(fr.LanguageDetail))
	}
	l := fr.LanguageDetail[0]

	if l.Units != 2 {
		t.Errorf("Units = %d, want 2 — the unassessed unit still exists", l.Units)
	}
	if l.Analyzable() != 1 {
		t.Fatalf("Analyzable() = %d, want 1", l.Analyzable())
	}
	if math.Abs(l.AverageScore-0.80) > 1e-9 {
		t.Errorf("AverageScore = %v, want 0.80 — the mean over the analyzable unit. "+
			"0.40 grades the language F on the strength of a unit that was never opened", l.AverageScore)
	}
	if l.Grade != "B" {
		t.Errorf("Grade = %q, want %q", l.Grade, "B")
	}
	// BottomScore is read as "the worst unit in this language". A placeholder
	// zero is not the worst measurement, it is the absence of one.
	if math.Abs(l.BottomScore-0.80) > 1e-9 {
		t.Errorf("BottomScore = %v, want 0.80 — an unassessed unit is not this language's worst score", l.BottomScore)
	}
}

// TestGenerateReportTree_MixedAgreesAcrossTheTree covers the committed markdown
// tree: the top-level index, its packages row and the package page must all state
// the same grade the report card does.
func TestGenerateReportTree_MixedAgreesAcrossTheTree(t *testing.T) {
	dir := t.TempDir()
	if _, err := report.GenerateReportTree(mixedPackageReport(), dir); err != nil {
		t.Fatalf("GenerateReportTree() error: %v", err)
	}

	index, err := os.ReadFile(filepath.Join(dir, "index.md"))
	if err != nil {
		t.Fatalf("index.md not found: %v", err)
	}
	if !strings.Contains(string(index), "**Overall:** 🟢 A- (90.0%)") {
		t.Errorf("tree index overall should agree with the report card, got:\n%s", index)
	}
	row := findLine(string(index), "| [app](app/index.md) |")
	if row == "" {
		t.Fatalf("no packages row in:\n%s", index)
	}
	if !strings.Contains(row, "| 3 | 2 | 🟢 A- | 90.0% | 100% |") {
		t.Errorf("packages row = %q, want 3 units / 2 unassessed / A- / 90.0%% / 100%%", row)
	}

	pkg, err := os.ReadFile(filepath.Join(dir, "app", "index.md"))
	if err != nil {
		t.Fatalf("app/index.md not found: %v", err)
	}
	if !strings.Contains(string(pkg), "**Grade:** 🟢 A- (90.0%)") {
		t.Errorf("package page grade should agree with the packages row, got:\n%s", pkg)
	}
}

// TestGenerateSite_MixedAgreesAcrossTheDashboard covers the HTML flagship. Both
// the Languages row and the Packages row are asserted, each scoped to its own
// section — see the note on sectionAfter.
func TestGenerateSite_MixedAgreesAcrossTheDashboard(t *testing.T) {
	dir := t.TempDir()
	if err := report.GenerateSite(mixedPackageReport(), report.SiteConfig{OutputDir: dir, Title: "test"}); err != nil {
		t.Fatalf("GenerateSite() error: %v", err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "index.html"))
	if err != nil {
		t.Fatalf("index.html not found: %v", err)
	}
	html := string(data)

	if !strings.Contains(html, `<div class="value">A-</div><div class="label">Overall Grade</div>`) {
		t.Errorf("dashboard overall grade should be A-, got:\n%s", html)
	}
	if !strings.Contains(html, `<div class="value">90.0%</div><div class="label">Overall Score</div>`) {
		t.Errorf("dashboard overall score should be 90.0%%, got:\n%s", html)
	}

	langSection := sectionAfter(html, "<h2>By Language</h2>")
	if langSection == "" {
		t.Fatalf("no By Language section in:\n%s", html)
	}
	if !strings.Contains(langSection, "<td>go</td><td>1</td>\n<td><span class=\"grade grade-a-\">A-</span></td>\n<td>90.0%</td>") {
		t.Errorf("Languages row for go should read A- / 90.0%%, got:\n%s", langSection)
	}

	pkgSection := sectionAfter(html, "<h2>Packages</h2>")
	if pkgSection == "" {
		t.Fatalf("no Packages section in:\n%s", html)
	}
	if !strings.Contains(pkgSection, "<td><span class=\"grade grade-a-\">A-</span></td>\n<td>90.0%</td>") {
		t.Errorf("Packages row should read A- / 90.0%%, got:\n%s", pkgSection)
	}
}
