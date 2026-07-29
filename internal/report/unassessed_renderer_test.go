package report_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/report"
)

// Three published surfaces render a unit's Score with no *Known() companion
// beside it. Each is the same defect as #47's headline one level down: a
// placeholder zero printed as though it were a measurement, next to a grade of
// N/A that says nothing was measured.
//
// They are grouped here because the shared lesson is about the VERIFICATION, not
// the code — each survived a sweep because the assertion nearest it happened to
// be scoped to a different row, a different section, or a different file format.

// siteIndexHTML generates a site into a temp dir and returns its index.html.
func siteIndexHTML(t *testing.T, fr report.FullReport) string {
	t.Helper()
	dir := t.TempDir()
	if err := report.GenerateSite(fr, report.SiteConfig{OutputDir: dir, Title: "test"}); err != nil {
		t.Fatalf("GenerateSite() error: %v", err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "index.html"))
	if err != nil {
		t.Fatalf("index.html not found: %v", err)
	}
	return string(data)
}

// --- The HTML dashboard's By-Language row (#79) ------------------------------

// TestGenerateSiteIndex_AllUnassessedLanguagesRow guards the Languages table.
//
// TestGenerateSiteIndex_AllUnassessed scopes its assertion to the Packages
// section, deliberately, because the two tables render an identical cell pair and
// an unscoped match is satisfied by the wrong row. That fixed the ambiguity in one
// direction and left the Languages row asserted by nothing at all: reverting it to
// an ungated formatter renders `N/A` and `0.0%` in the same cell pair with the
// suite still green. Both rows now have their own scoped assertion.
func TestGenerateSiteIndex_AllUnassessedLanguagesRow(t *testing.T) {
	html := siteIndexHTML(t, unsupportedPackageReport())

	langSection := sectionAfter(html, "<h2>By Language</h2>")
	if langSection == "" {
		t.Fatalf("no By Language section in:\n%s", html)
	}
	if strings.Contains(langSection, "<td>0.0%</td>") {
		t.Errorf("Languages row prints a 0.0%% score for a language the engine has no analyzer for:\n%s", langSection)
	}
	if !strings.Contains(langSection, "<td>swift</td><td>3</td>\n<td><span class=\"grade grade-na\">N/A</span></td>\n<td>n/a</td>") {
		t.Errorf("Languages row should read N/A and n/a, got:\n%s", langSection)
	}
}

// --- Top Issues (#87) --------------------------------------------------------

// topIssuesRecords has a failing unit — both Top Issues renderers are gated on
// Failing > 0 — plus a passing unit and two unassessed ones. The unassessed units
// carry the lowest score in the corpus (the placeholder 0), so they sort to the
// TOP of a list ordered worst-first: this is the only shape in which the table
// puts unopened code at the head of "units needing attention".
func topIssuesRecords() []domain.CertificationRecord {
	return []domain.CertificationRecord{
		makeCardRecord("go", "app/bad.go", "Bad", domain.StatusDecertified, 0.20),
		makeCardRecord("go", "app/good.go", "Good", domain.StatusCertified, 0.90),
		unassessedRecord("swift", "app/b.swift", "b"),
		unassessedRecord("swift", "app/c.swift", "c"),
	}
}

// TestBuildTopIssues_ExcludesUnassessedUnits pins the producer on the RECORDED
// predicate.
//
// The list is currently protected only by a skip on StatusExempt — a derived
// proxy for "unassessed" rather than the recorded flag every other surface on
// this branch reads. That proxy is exactly what #47 established cannot be trusted
// to carry the claim, and it is the single line standing between the delivered
// card and a row reading `| N/A | 0.0% | lowest score |` about a file nobody
// opened.
func TestBuildTopIssues_ExcludesUnassessedUnits(t *testing.T) {
	c := report.GenerateCard(topIssuesRecords(), "test/repo", "abc", time.Now())

	if len(c.TopIssues) == 0 {
		t.Fatal("no top issues; the fixture must contain a failing unit for the table to render")
	}
	for _, issue := range c.TopIssues {
		if strings.Contains(issue.UnitID, ".swift") {
			t.Errorf("top issues names %q — an unassessed unit is not known to need attention, "+
				"and it heads the list only because its placeholder score sorts lowest", issue.UnitID)
		}
		if issue.Grade == "N/A" {
			t.Errorf("top issues carries a unit graded N/A: %+v", issue)
		}
	}
	if c.TopIssues[0].UnitID == "" || !strings.Contains(c.TopIssues[0].UnitID, "bad.go") {
		t.Errorf("worst assessed unit should head the list, got %q", c.TopIssues[0].UnitID)
	}
}

// TestBuildTopIssues_ExcludesLegacyUnassessedUnits is the case the StatusExempt
// skip cannot reach, and therefore the one that proves the recorded flag is the
// predicate that carries the claim.
//
// A legacy record — a Swift unit written as decertified at score 0, before the
// engine knew it had no analyzer for Swift — comes back from the store with
// Unsupported backfilled true and Status still decertified. StatusExempt does not
// match it, IsPassing() is false, and its placeholder 0 sorts below every real
// score, so it lands at the HEAD of "units needing attention" as `🔴 F | 0.0% |
// decertified`. That is the delivered card telling a client its worst code is a
// file the engine cannot open.
//
// Removing `r.Unsupported` from the skip leaves the whole suite green on every
// fixture built from today's pipeline output, because those carry StatusExempt.
// This fixture is the corpus the branch exists to handle.
func TestBuildTopIssues_ExcludesLegacyUnassessedUnits(t *testing.T) {
	legacy := makeCardRecord("swift", "app/App.swift", "App", domain.StatusDecertified, 0)
	legacy.Unsupported = true
	legacy.Grade = domain.GradeF

	c := report.GenerateCard([]domain.CertificationRecord{
		makeCardRecord("go", "app/bad.go", "Bad", domain.StatusDecertified, 0.20),
		makeCardRecord("go", "app/good.go", "Good", domain.StatusCertified, 0.90),
		legacy,
	}, "test/repo", "abc", time.Now())

	for _, issue := range c.TopIssues {
		if strings.Contains(issue.UnitID, "App.swift") {
			t.Errorf("top issues names a legacy unassessed unit %q as %s at %.1f%% — StatusExempt does not "+
				"match this record, so only the recorded Unsupported flag keeps the engine from reporting "+
				"a file it cannot open as the worst code in the repo", issue.UnitID, issue.Grade, issue.Score*100)
		}
	}
	// The list holds every assessed unit worst-first, so the assertion that
	// matters is which unit HEADS it: the legacy record's placeholder 0 sorts
	// below the genuinely failing unit's 0.20 and would displace it.
	if len(c.TopIssues) != 2 {
		t.Fatalf("top issues should hold the two assessed units, got %+v", c.TopIssues)
	}
	if !strings.Contains(c.TopIssues[0].UnitID, "bad.go") {
		t.Errorf("worst ASSESSED unit should head the list, got %q", c.TopIssues[0].UnitID)
	}

	md := report.FormatCardMarkdown(c)
	if strings.Contains(md, "App.swift") {
		t.Errorf("REPORT_CARD.md Top Issues lists a unit the engine never opened:\n%s", md)
	}
}

// TestIssueCard_UnassessedScoreIsNotRendered pins the two RENDERERS independently
// of the producer.
//
// The producer's skip and the renderers' gate are separate protections and must
// be verified separately — a test that only builds the card through GenerateCard
// can never observe the renderer, which is how both surfaces came to print the
// score raw. The IssueCard is constructed directly here on purpose.
func TestIssueCard_UnassessedScoreIsNotRendered(t *testing.T) {
	c := report.GenerateCard(topIssuesRecords(), "test/repo", "abc", time.Now())
	c.TopIssues = append(c.TopIssues, report.IssueCard{
		UnitID:      "app/b.swift:b",
		Grade:       "N/A",
		Score:       0,
		Unsupported: true,
		Reason:      "lowest score",
	})

	if c.TopIssues[len(c.TopIssues)-1].ScoreKnown() {
		t.Error("IssueCard.ScoreKnown() = true for an unassessed unit, want false")
	}

	md := report.FormatCardMarkdown(c)
	if strings.Contains(md, "| `app/b.swift:b` | N/A | 0.0% |") {
		t.Errorf("markdown Top Issues prints a 0.0%% score beside a grade of N/A:\n%s", md)
	}
	if !strings.Contains(md, "| `app/b.swift:b` | N/A | n/a |") {
		t.Errorf("markdown Top Issues should render the score as n/a, got:\n%s", md)
	}
}

// TestSiteTopIssues_UnassessedScoreIsNotRendered is the same gate on the HTML
// dashboard, which formats the score through its own template function.
func TestSiteTopIssues_UnassessedScoreIsNotRendered(t *testing.T) {
	fr := report.GenerateFullReport(topIssuesRecords(), "test/repo", "abc", time.Now())
	fr.Card.TopIssues = append(fr.Card.TopIssues, report.IssueCard{
		UnitID:      "app/b.swift:b",
		Grade:       "N/A",
		Score:       0,
		Unsupported: true,
		Reason:      "lowest score",
	})

	issues := sectionAfter(siteIndexHTML(t, fr), "<h2>Top Issues</h2>")
	if issues == "" {
		t.Fatalf("no Top Issues section rendered")
	}
	if strings.Contains(issues, "<td>0.0%</td>") {
		t.Errorf("dashboard Top Issues prints a 0.0%% score for an unassessed unit:\n%s", issues)
	}
	if !strings.Contains(issues, "<td>n/a</td>") {
		t.Errorf("dashboard Top Issues should render the score as n/a, got:\n%s", issues)
	}
}

// --- The search index (#94/#95) ----------------------------------------------

// TestBuildSearchIndex_UnassessedUnitCarriesNoScore covers search-index.js.
//
// It is a published artifact carrying per-unit Grade and Score, and it had no
// unsupported flag and no gate — so even UNMUTATED it shipped `"sc":0` for a unit
// the engine never opened. Earlier sweeps missed it because the file is minified
// JS rather than markdown or HTML, so no rendered-cell assertion could see it.
//
// The absent field is the fix, matching this branch's thesis: a score that was
// never measured is not a claim of zero, and JSON has a way to say so.
func TestBuildSearchIndex_UnassessedUnitCarriesNoScore(t *testing.T) {
	fr := report.GenerateFullReport(mixedPackageRecords(), "test/repo", "abc", time.Now())
	entries := report.BuildSearchIndex(fr)

	byName := make(map[string]report.SearchEntry, len(entries))
	for _, e := range entries {
		byName[e.Name] = e
	}

	unassessed, ok := byName["b"]
	if !ok {
		t.Fatalf("no search entry for the unassessed unit, got %v", byName)
	}
	if !unassessed.Unsupported {
		t.Error("search entry Unsupported = false for a unit the engine never opened")
	}
	if unassessed.ScoreKnown() {
		t.Error("search entry ScoreKnown() = true for an unassessed unit, want false")
	}
	if unassessed.Score != nil {
		t.Errorf("search entry Score = %v, want nil — an absent field is not a claim of zero", *unassessed.Score)
	}

	assessed, ok := byName["A"]
	if !ok {
		t.Fatalf("no search entry for the assessed unit, got %v", byName)
	}
	if assessed.Score == nil || *assessed.Score != 0.90 {
		t.Errorf("assessed unit lost its measured score: %v", assessed.Score)
	}
	if assessed.Unsupported {
		t.Error("assessed unit flagged unsupported")
	}
}

// TestFormatSearchIndexJS_UnassessedUnitEmitsNoZeroScore asserts on the bytes
// that ship, because the defect lives in the serialisation and a struct-level
// assertion cannot see an omitempty tag that is missing.
func TestFormatSearchIndexJS_UnassessedUnitEmitsNoZeroScore(t *testing.T) {
	fr := report.GenerateFullReport(mixedPackageRecords(), "test/repo", "abc", time.Now())
	js := report.FormatSearchIndexJS(report.BuildSearchIndex(fr))

	if strings.Contains(js, `"sc":0,`) || strings.HasSuffix(strings.TrimSpace(js), `"sc":0};`) {
		t.Errorf("search-index.js ships a fabricated zero score for an unassessed unit:\n%s", js)
	}

	payload := strings.TrimSuffix(strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(js), "const SEARCH_INDEX = ")), ";")
	var decoded []map[string]any
	if err := json.Unmarshal([]byte(payload), &decoded); err != nil {
		t.Fatalf("search index is not valid JSON: %v\n%s", err, js)
	}
	for _, e := range decoded {
		if e["u"] == true {
			if _, present := e["sc"]; present {
				t.Errorf("unassessed entry carries a score key: %v", e)
			}
		} else if _, present := e["sc"]; !present {
			t.Errorf("assessed entry lost its score key: %v", e)
		}
	}
}
