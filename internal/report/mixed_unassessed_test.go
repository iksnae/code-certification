package report_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/report"
)

// The all-unassessed fixtures elsewhere in this package cannot observe a
// miscounted numerator: with zero analyzable units every rate short-circuits to
// "n/a" no matter what the passing count is. A package holding BOTH an assessed
// and an unassessed unit is the only shape where the count is visible, and its
// absence is why the false positive survived at three sites after being fixed at
// the others. Every aggregation over a package or a language is asserted here.
//
// The fixture is one package, `app`, with one certified Go unit and two
// unassessed Swift units — 3 units, 1 analyzable, 1 passing, 100% of what was
// measured, and 2 units about which nothing was claimed.
func mixedPackageRecords() []domain.CertificationRecord {
	return []domain.CertificationRecord{
		makeCardRecord("go", "app/a.go", "A", domain.StatusCertified, 0.90),
		unassessedRecord("swift", "app/b.swift", "b"),
		unassessedRecord("swift", "app/c.swift", "c"),
	}
}

func mixedPackageReport() report.FullReport {
	return report.GenerateFullReport(mixedPackageRecords(), "test/repo", "abc", time.Now())
}

// TestBuildPackageSummaries_MixedPackage pins the numerator and the denominator
// independently. An all-unassessed package cannot tell a dropped `continue`
// (passing 3) from a correct exclusion (passing 1), nor a denominator of Units
// (1/3) from one of Analyzable (1/1).
func TestBuildPackageSummaries_MixedPackage(t *testing.T) {
	pkgs := report.BuildPackageSummaries(mixedPackageReport())

	if len(pkgs) != 1 {
		t.Fatalf("package count = %d, want 1", len(pkgs))
	}
	p := pkgs[0]

	if p.Units != 3 {
		t.Errorf("Units = %d, want 3 — the unassessed units still exist", p.Units)
	}
	if p.Unsupported != 2 {
		t.Errorf("Unsupported = %d, want 2", p.Unsupported)
	}
	if p.Analyzable() != 1 {
		t.Errorf("Analyzable() = %d, want 1", p.Analyzable())
	}
	if !p.PassRateKnown() {
		t.Error("PassRateKnown() = false, want true — one unit was assessed")
	}
	if p.PassRate != 1.0 {
		t.Errorf("PassRate = %.3f, want 1.000 — 1 of 1 analyzable unit passed", p.PassRate)
	}
}

// TestFormatPackageIndexMarkdown_MixedPackage covers the seventh copy of the
// count. .certification/reports/<pkg>/index.md is committed and is one click
// from reports/index.md, so a passing count that includes unassessed units puts
// two contradicting claims in the same delivered artifact.
func TestFormatPackageIndexMarkdown_MixedPackage(t *testing.T) {
	dir := t.TempDir()
	if _, err := report.GenerateReportTree(mixedPackageReport(), dir); err != nil {
		t.Fatalf("GenerateReportTree() error: %v", err)
	}

	data, err := os.ReadFile(filepath.Join(dir, "app", "index.md"))
	if err != nil {
		t.Fatalf("app/index.md not found: %v", err)
	}
	md := string(data)

	if strings.Contains(md, "**Passing:** 3 / 3") {
		t.Errorf("package index claims 3 of 3 units passed; only 1 was assessed:\n%s", md)
	}
	if !strings.Contains(md, "**Passing:** 1 / 1") {
		t.Errorf("package index should report 1 of 1 analyzable unit passing, got:\n%s", md)
	}
	if !strings.Contains(md, "**Not Assessed:** 2") {
		t.Errorf("package index should report the unassessed count, got:\n%s", md)
	}
}

// TestFormatReportTreeIndex_MixedPackage covers the packages table. A row
// reading "3 units · 100%" is read as three passing units; the rate is honest
// per analyzable unit only if the row says how many that is.
func TestFormatReportTreeIndex_MixedPackage(t *testing.T) {
	dir := t.TempDir()
	if _, err := report.GenerateReportTree(mixedPackageReport(), dir); err != nil {
		t.Fatalf("GenerateReportTree() error: %v", err)
	}

	data, err := os.ReadFile(filepath.Join(dir, "index.md"))
	if err != nil {
		t.Fatalf("index.md not found: %v", err)
	}
	md := string(data)

	if !strings.Contains(md, "| Not Assessed |") {
		t.Errorf("packages table should carry a Not Assessed column when a package has unassessed units, got:\n%s", md)
	}
	row := findLine(md, "| [app](app/index.md) |")
	if row == "" {
		t.Fatalf("no packages row for app in:\n%s", md)
	}
	if !strings.Contains(row, "| 3 | 2 |") {
		t.Errorf("packages row = %q, want 3 units with 2 not assessed", row)
	}
	if !strings.Contains(row, "100%") {
		t.Errorf("packages row = %q, want the 100%% rate over the single analyzable unit", row)
	}
}

// TestFormatReportTreeIndex_AllAssessedOmitsColumn pins the ordinary path: a
// repo with nothing unassessed keeps the table it had.
func TestFormatReportTreeIndex_AllAssessedOmitsColumn(t *testing.T) {
	fr := report.GenerateFullReport([]domain.CertificationRecord{
		makeCardRecord("go", "app/a.go", "A", domain.StatusCertified, 0.90),
		makeCardRecord("go", "app/b.go", "B", domain.StatusCertified, 0.80),
	}, "test/repo", "abc", time.Now())

	dir := t.TempDir()
	if _, err := report.GenerateReportTree(fr, dir); err != nil {
		t.Fatalf("GenerateReportTree() error: %v", err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "index.md"))
	if err != nil {
		t.Fatalf("index.md not found: %v", err)
	}
	md := string(data)

	if strings.Contains(md, "Not Assessed") {
		t.Errorf("packages table should omit the Not Assessed column when every unit was assessed, got:\n%s", md)
	}
	if !strings.Contains(md, "| Package | Units | Grade | Score | Pass Rate |") {
		t.Errorf("packages table header changed for an all-assessed repo, got:\n%s", md)
	}
}

// TestGeneratePackagePages_MixedPackage covers the HTML package page, which
// computes the same numbers a third time.
func TestGeneratePackagePages_MixedPackage(t *testing.T) {
	dir := t.TempDir()
	if err := report.GenerateSite(mixedPackageReport(), report.SiteConfig{
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

	if !strings.Contains(html, `<div class="value">100.0%</div><div class="label">Pass Rate</div>`) {
		t.Errorf("package page should render 100.0%% over the single analyzable unit, got:\n%s", html)
	}
	if !strings.Contains(html, `<div class="value">2</div><div class="label">Not Assessed</div>`) {
		t.Errorf("package page should render the unassessed count, got:\n%s", html)
	}
}

// TestComputeLanguageBreakdowns_MixedPackage covers DetailedReport.ByLanguage.
// This is the producer behind `certify report --detailed`; it counted
// StatusExempt as passing and never set Unsupported, so the renderer's
// "(n not assessed)" suffix could not fire.
func TestComputeLanguageBreakdowns_MixedPackage(t *testing.T) {
	d := report.Detailed(mixedPackageRecords(), time.Now())

	swift, ok := d.ByLanguage["swift"]
	if !ok {
		t.Fatalf("no swift entry in ByLanguage: %v", d.ByLanguage)
	}
	if swift.Units != 2 {
		t.Errorf("swift Units = %d, want 2", swift.Units)
	}
	if swift.Passing != 0 {
		t.Errorf("swift Passing = %d, want 0 — no swift unit was assessed", swift.Passing)
	}
	if swift.Unsupported != 2 {
		t.Errorf("swift Unsupported = %d, want 2 — the renderer reads this field", swift.Unsupported)
	}
	if got := d.ByLanguage["go"].Passing; got != 1 {
		t.Errorf("go Passing = %d, want 1", got)
	}

	out := report.FormatDetailedText(d)
	if !strings.Contains(out, "(2 not assessed)") {
		t.Errorf("detailed text should annotate swift as unassessed, got:\n%s", out)
	}
	if strings.Contains(out, "swift     2 units, 2 passing") {
		t.Errorf("detailed text claims 2 passing swift units:\n%s", out)
	}
}

// TestHealth_AllUnassessedIsNotAZeroPassRate covers the health report. Every
// bucket is zero because nothing was counted, and printing 0.0% from that turns
// "nothing was measured" into "everything failed" — the same false claim as
// #47, inverted.
func TestHealth_AllUnassessedIsNotAZeroPassRate(t *testing.T) {
	h := report.Health([]domain.CertificationRecord{
		unassessedRecord("swift", "a.swift", "a"),
		unassessedRecord("swift", "b.swift", "b"),
		unassessedRecord("swift", "c.swift", "c"),
	})

	if h.TotalUnits != 3 {
		t.Errorf("TotalUnits = %d, want 3", h.TotalUnits)
	}
	if h.AnalyzableUnits != 0 {
		t.Errorf("AnalyzableUnits = %d, want 0", h.AnalyzableUnits)
	}
	if h.Unsupported != 3 {
		t.Errorf("Unsupported = %d, want 3", h.Unsupported)
	}
	if h.PassRateKnown() {
		t.Error("PassRateKnown() = true, want false — 0/0 is undefined")
	}

	out := report.FormatText(h)
	if strings.Contains(out, "Pass Rate:              0.0%") {
		t.Errorf("health report states a 0%% pass rate for a repo that was never assessed:\n%s", out)
	}
	if !strings.Contains(out, "Pass Rate:              n/a") {
		t.Errorf("health report should render the pass rate as n/a, got:\n%s", out)
	}
	if !strings.Contains(out, "Average Score:          n/a") {
		t.Errorf("health report should render the average score as n/a, got:\n%s", out)
	}
	if !strings.Contains(out, "Not Assessed:           3") {
		t.Errorf("health report should report the unassessed count, got:\n%s", out)
	}
}

// TestHealth_MixedPackage pins the rate against what was actually measured.
func TestHealth_MixedPackage(t *testing.T) {
	h := report.Health(mixedPackageRecords())

	if h.TotalUnits != 3 {
		t.Errorf("TotalUnits = %d, want 3", h.TotalUnits)
	}
	if h.AnalyzableUnits != 1 {
		t.Errorf("AnalyzableUnits = %d, want 1", h.AnalyzableUnits)
	}
	if h.Unsupported != 2 {
		t.Errorf("Unsupported = %d, want 2", h.Unsupported)
	}
	if h.PassRate != 1.0 {
		t.Errorf("PassRate = %.3f, want 1.000", h.PassRate)
	}
	if !h.PassRateKnown() {
		t.Error("PassRateKnown() = false, want true")
	}
	if out := report.FormatText(h); !strings.Contains(out, "Pass Rate:              100.0%") {
		t.Errorf("health report should render 100.0%% over the analyzable unit, got:\n%s", out)
	}
}

func findLine(s, substr string) string {
	for _, line := range strings.Split(s, "\n") {
		if strings.Contains(line, substr) {
			return line
		}
	}
	return ""
}
