package report_test

import (
	"strings"
	"testing"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/report"
)

// mixedRecords is the shape every one of these assertions needs and no existing
// badge test had: a corpus that is partly analyzable. badge_test.go covers the
// all-assessed pole and the nothing-assessed pole, and the defect lives strictly
// between them — the branch where a grade is real, a rate is real, and the unit
// count beside them covers a larger population than either.
func mixedRecords() []domain.CertificationRecord {
	var recs []domain.CertificationRecord
	for _, n := range []string{"a", "b", "c", "d", "e"} {
		recs = append(recs, domain.CertificationRecord{
			UnitID:   domain.NewUnitID("go", "src/"+n+".go", n),
			UnitPath: "src/" + n + ".go",
			Status:   domain.StatusCertified,
			Grade:    domain.GradeBPlus,
			Score:    0.8888888888888888,
		})
	}
	for _, n := range []string{"v", "w", "x", "y"} {
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

func mixedCard(t *testing.T) report.Card {
	t.Helper()
	c := report.GenerateCard(mixedRecords(), "test/repo", "abc", time.Now())
	if c.TotalUnits != 9 || c.AnalyzableUnits != 5 {
		t.Fatalf("fixture drifted: TotalUnits=%d AnalyzableUnits=%d, want 9 and 5",
			c.TotalUnits, c.AnalyzableUnits)
	}
	return c
}

// TestBadge_MixedCorpusQualifiesItsPopulation is B2. The real artifact read
// "B+ · 100% · 9 units": a grade over five units, a rate over five units, and
// a count of nine, in the most widely read string the product emits.
func TestBadge_MixedCorpusQualifiesItsPopulation(t *testing.T) {
	msg := report.GenerateBadge(mixedCard(t)).Message

	const bad = "B+ · 100% · 9 units"
	if msg == bad {
		t.Errorf("badge message = %q — a grade and a rate measured over 5 units, beside a count of 9", msg)
	}
	const want = "B+ · 100% · 5 of 9 units analyzable"
	if msg != want {
		t.Errorf("badge message = %q, want %q — the real grade and rate, over a named population", msg, want)
	}
}

// TestBadge_AllAssessedCorpusUnchanged guards the disclosure against
// over-reaching: with nothing unassessed there is no gap to disclose, and the
// badge must keep the string it has always emitted.
func TestBadge_AllAssessedCorpusUnchanged(t *testing.T) {
	recs := mixedRecords()[:5]
	c := report.GenerateCard(recs, "test/repo", "abc", time.Now())
	if got, want := report.GenerateBadge(c).Message, "B+ · 100% · 5 units"; got != want {
		t.Errorf("badge message = %q, want %q", got, want)
	}
}

// TestCardMarkdown_PackagesTableDisclosesUnassessed is B4. The report tree's
// Packages table already carries this column, with a comment saying a row
// reading "3 units · 100%" is read as three passing units. The report card's
// own Packages table is that row, in the artifact clients actually open.
func TestCardMarkdown_PackagesTableDisclosesUnassessed(t *testing.T) {
	c := mixedCard(t)
	fr := report.GenerateFullReport(mixedRecords(), "test/repo", "abc", time.Now())
	c.Packages = report.BuildPackageSummaries(fr)
	if len(c.Packages) == 0 {
		t.Fatal("fixture produced no packages")
	}

	md := report.FormatCardMarkdown(c)
	pkgTable := sectionOf(t, md, "### Packages")

	if !strings.Contains(pkgTable, "Not Assessed") {
		t.Errorf("Packages table has no Not Assessed column:\n%s", pkgTable)
	}
	if strings.Contains(pkgTable, "| 9 | 🟢 B+ |") {
		t.Errorf("Packages table states 9 units beside a B+ measured over 5:\n%s", pkgTable)
	}
}

// TestCardMarkdown_LanguageTableDisclosesUnassessed is B5.
func TestCardMarkdown_LanguageTableDisclosesUnassessed(t *testing.T) {
	md := report.FormatCardMarkdown(mixedCard(t))
	langTable := sectionOf(t, md, "### By Language")

	if !strings.Contains(langTable, "Not Assessed") {
		t.Errorf("By Language table has no Not Assessed column:\n%s", langTable)
	}
	if !strings.Contains(langTable, "swift") {
		t.Fatalf("By Language table lost the swift row:\n%s", langTable)
	}
}

// TestCardMarkdown_AllAssessedTablesUnchanged is the over-reach guard for both
// tables: no gap, no extra column.
func TestCardMarkdown_AllAssessedTablesUnchanged(t *testing.T) {
	c := report.GenerateCard(mixedRecords()[:5], "test/repo", "abc", time.Now())
	md := report.FormatCardMarkdown(c)
	if strings.Contains(md, "Not Assessed") {
		t.Errorf("all-assessed card mentions Not Assessed:\n%s", md)
	}
}

// TestFormatUnitPopulation pins the shared convention directly.
func TestFormatUnitPopulation(t *testing.T) {
	cases := []struct {
		total, unsupported int
		want               string
	}{
		{9, 4, "5 of 9 units analyzable"},
		{5, 0, "5 units"},
		{9, 9, "0 of 9 units analyzable"},
		{1, 0, "1 units"},
	}
	for _, tc := range cases {
		if got := report.FormatUnitPopulation(tc.total, tc.unsupported); got != tc.want {
			t.Errorf("FormatUnitPopulation(%d, %d) = %q, want %q",
				tc.total, tc.unsupported, got, tc.want)
		}
	}
}

// sectionOf returns the markdown from heading up to the next "### " heading.
func sectionOf(t *testing.T, md, heading string) string {
	t.Helper()
	i := strings.Index(md, heading)
	if i < 0 {
		t.Fatalf("heading %q not found in:\n%s", heading, md)
	}
	rest := md[i+len(heading):]
	if j := strings.Index(rest, "\n### "); j >= 0 {
		rest = rest[:j]
	}
	return heading + rest
}
