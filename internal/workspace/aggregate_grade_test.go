package workspace

import (
	"math"
	"strings"
	"testing"
)

// The workspace rollup carries the same pair of claims as the report card — a
// pass rate and a weighted score — and #47 fixed only the first. A workspace in
// which nothing was analysable rendered `Pass Rate: n/a` beside `Overall: F
// (0.0%)`; one in which SOME unit was analysable rendered a real grade computed
// over units nobody opened.
//
// Both halves are now taken over AnalyzableUnits, together with the card, the
// package aggregate and the language aggregate. A partial flip is worse than
// none: these four print into the same artifacts, so any one moving alone
// relocates the contradiction instead of removing it.

func allUnassessedWorkspace() []SubmoduleSummary {
	return []SubmoduleSummary{
		{Name: "ios", Path: "apps/ios", Grade: "N/A", Score: 0, Units: 3, Unsupported: 3, HasCertify: true},
		{Name: "mac", Path: "apps/mac", Grade: "N/A", Score: 0, Units: 2, Unsupported: 2, HasCertify: true},
	}
}

// TestAggregateCards_AllUnassessedHasNoOverallGrade covers the rollup struct.
func TestAggregateCards_AllUnassessedHasNoOverallGrade(t *testing.T) {
	wc := AggregateCards(allUnassessedWorkspace())

	if wc.ScoreKnown() {
		t.Error("ScoreKnown() = true, want false — no unit in the workspace was scored")
	}
	if wc.OverallGrade != "N/A" {
		t.Errorf("OverallGrade = %q, want %q", wc.OverallGrade, "N/A")
	}
	if math.IsNaN(wc.OverallScore) {
		t.Error("OverallScore is NaN")
	}
	if wc.OverallScore != 0 {
		t.Errorf("OverallScore = %v, want the zero value left untouched when unknown", wc.OverallScore)
	}
}

// TestFormatWorkspaceCardMarkdown_AllUnassessedGrade asserts on the workspace
// REPORT_CARD.md.
func TestFormatWorkspaceCardMarkdown_AllUnassessedGrade(t *testing.T) {
	md := FormatWorkspaceCardMarkdown(AggregateCards(allUnassessedWorkspace()))

	if strings.Contains(md, "Overall: F") {
		t.Errorf("workspace card grades an unassessed workspace F:\n%s", md)
	}
	if !strings.Contains(md, "## ⚪ Overall: N/A (n/a)") {
		t.Errorf("workspace card should render the overall grade as N/A, got:\n%s", md)
	}
	if !strings.Contains(md, "| [ios](apps/ios/.certification/reports/index.md) | 3 | ⚪ N/A | n/a | n/a |") {
		t.Errorf("submodule row should carry no grade, score or rate, got:\n%s", md)
	}
}

// TestFormatSubmoduleSummary_AllUnassessedScore covers the per-submodule page.
func TestFormatSubmoduleSummary_AllUnassessedScore(t *testing.T) {
	out := formatSubmoduleSummary(SubmoduleSummary{
		Name: "ios", Path: "apps/ios", Grade: "N/A", Units: 3, Unsupported: 3, HasCertify: true,
	})

	if strings.Contains(out, "| **Score** | 0.0% |") {
		t.Errorf("submodule summary prints a 0.0%% score for units that were never opened:\n%s", out)
	}
	if !strings.Contains(out, "| **Score** | n/a |") {
		t.Errorf("submodule summary should render the score as n/a, got:\n%s", out)
	}
}

// TestFormatWorkspaceIndex_AllUnassessedGrade covers the workspace index page.
func TestFormatWorkspaceIndex_AllUnassessedGrade(t *testing.T) {
	out := formatWorkspaceIndex(AggregateCards(allUnassessedWorkspace()))

	if !strings.Contains(out, "**Overall:** ⚪ N/A (n/a)") {
		t.Errorf("workspace index should render the overall grade as N/A, got:\n%s", out)
	}
	if strings.Contains(out, "0.0%") {
		t.Errorf("workspace index prints a 0.0%% figure for an unassessed workspace:\n%s", out)
	}
}

// TestAggregateCards_MixedScoreExcludesUnassessedUnits is the workspace half of
// the score denominator, moved together with Card.OverallScore and the package
// and language aggregates.
//
// It used to pin the whole-unit weighting: 0.95 over 10 assessed units, spread
// across 13, gave 0.731 and graded the workspace C — a two-grade drop caused
// entirely by a submodule nobody analysed. The submodule scores summed here are
// already means over each submodule's ANALYZABLE units, so the only weighting
// that reconstructs a correct overall mean is by that same count.
func TestAggregateCards_MixedScoreExcludesUnassessedUnits(t *testing.T) {
	wc := AggregateCards([]SubmoduleSummary{
		{Name: "api", Path: "services/api", Grade: "A", Score: 0.95, Units: 10, Passing: 8, Failing: 2, HasCertify: true},
		{Name: "ios", Path: "apps/ios", Grade: "N/A", Score: 0, Units: 3, Unsupported: 3, HasCertify: true},
	})

	if !wc.ScoreKnown() {
		t.Fatal("ScoreKnown() = false, want true — ten units were assessed")
	}
	if wc.AnalyzableUnits != 10 {
		t.Fatalf("AnalyzableUnits = %d, want 10", wc.AnalyzableUnits)
	}
	if math.Abs(wc.OverallScore-0.95) > 1e-9 {
		t.Errorf("OverallScore = %v, want 0.95 — the mean over the ten units that were assessed. "+
			"Spreading it over all 13 gives 0.731 and grades the workspace on unopened code", wc.OverallScore)
	}
	if wc.OverallGrade != "A" {
		t.Errorf("OverallGrade = %q, want %q", wc.OverallGrade, "A")
	}
}

// TestAggregateCards_MixedWeightsBySubmoduleAnalyzableCount pins the WEIGHTING
// rather than the denominator. With one submodule per test above, weight and
// denominator cancel and either could be wrong unobserved; two submodules of
// different analyzable size are the only shape that tells them apart.
func TestAggregateCards_MixedWeightsBySubmoduleAnalyzableCount(t *testing.T) {
	wc := AggregateCards([]SubmoduleSummary{
		// 3 analyzable at 1.00, 7 unassessed.
		{Name: "api", Path: "services/api", Grade: "A", Score: 1.00, Units: 10, Passing: 3, Unsupported: 7, HasCertify: true},
		// 1 analyzable at 0.60, 0 unassessed.
		{Name: "web", Path: "apps/web", Grade: "D", Score: 0.60, Units: 1, Passing: 0, Failing: 1, HasCertify: true},
	})

	if wc.AnalyzableUnits != 4 {
		t.Fatalf("AnalyzableUnits = %d, want 4", wc.AnalyzableUnits)
	}
	want := (1.00*3 + 0.60*1) / 4 // 0.90
	if math.Abs(wc.OverallScore-want) > 1e-9 {
		t.Errorf("OverallScore = %v, want %v — each submodule's score weighted by the units it was "+
			"actually taken over. Weighting by Units gives %v, which counts seven unopened files as "+
			"evidence for the 1.00", wc.OverallScore, want, (1.00*10+0.60*1)/4)
	}
}
