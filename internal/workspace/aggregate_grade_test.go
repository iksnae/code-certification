package workspace

import (
	"math"
	"strings"
	"testing"
)

// The workspace rollup carries the same pair of claims as the report card — a
// pass rate and a weighted score — and #47 fixed only the first. A workspace in
// which nothing was analysable rendered `Pass Rate: n/a` beside `Overall: F
// (0.0%)`.
//
// As at the card and package levels, the denominator used when SOME unit was
// analysable is issue #32 and is pinned rather than changed.

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

// TestAggregateCards_MixedScoreDenominatorIsPinned states what this branch
// deliberately leaves alone. The weighted mean spans every unit, including the
// unassessed ones — 0.95 over 10 assessed units and 0 over 3 unassessed ones
// gives 0.731 rather than 0.95. Moving that denominator is issue #32 and must
// move together with Card.OverallScore.
func TestAggregateCards_MixedScoreDenominatorIsPinned(t *testing.T) {
	wc := AggregateCards([]SubmoduleSummary{
		{Name: "api", Path: "services/api", Grade: "A", Score: 0.95, Units: 10, Passing: 8, Failing: 2, HasCertify: true},
		{Name: "ios", Path: "apps/ios", Grade: "N/A", Score: 0, Units: 3, Unsupported: 3, HasCertify: true},
	})

	if !wc.ScoreKnown() {
		t.Fatal("ScoreKnown() = false, want true — ten units were assessed")
	}
	want := 0.95 * 10 / 13
	if math.Abs(wc.OverallScore-want) > 1e-9 {
		t.Errorf("OverallScore = %v, want %v (weighted over all 13 units) — changing this denominator is issue #32",
			wc.OverallScore, want)
	}
}
