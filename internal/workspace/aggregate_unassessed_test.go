package workspace

import (
	"strings"
	"testing"
)

// TestAggregateCards_UnassessedLeavesTheDenominator is the workspace half of
// #47. The rollup sums per-repo Passing into a multi-repo PassRate, so a
// submodule whose units were never analysed must leave both the numerator and
// the denominator here too.
func TestAggregateCards_UnassessedLeavesTheDenominator(t *testing.T) {
	subs := []SubmoduleSummary{
		{Name: "api", Path: "services/api", Grade: "A", Score: 0.95, Units: 10, Passing: 8, Failing: 2, HasCertify: true},
		{Name: "ios", Path: "apps/ios", Grade: "F", Score: 0.0, Units: 3, Passing: 0, Failing: 0, Unsupported: 3, HasCertify: true},
	}

	wc := AggregateCards(subs)

	if wc.TotalUnits != 13 {
		t.Errorf("TotalUnits = %d, want 13 — the units still exist", wc.TotalUnits)
	}
	if wc.TotalUnsupported != 3 {
		t.Errorf("TotalUnsupported = %d, want 3", wc.TotalUnsupported)
	}
	if wc.AnalyzableUnits != 10 {
		t.Errorf("AnalyzableUnits = %d, want 10", wc.AnalyzableUnits)
	}
	if wc.PassRate != 0.8 {
		t.Errorf("PassRate = %.3f, want 0.800 (8 of 10 analyzable, not 8 of 13)", wc.PassRate)
	}
	if !wc.PassRateKnown() {
		t.Error("PassRateKnown() = false, want true — ten units were assessed")
	}
}

// TestAggregateCards_AllUnassessedHasNoPassRate covers the 0/0 case at the
// workspace level.
func TestAggregateCards_AllUnassessedHasNoPassRate(t *testing.T) {
	subs := []SubmoduleSummary{
		{Name: "ios", Path: "apps/ios", Grade: "F", Score: 0.0, Units: 3, Unsupported: 3, HasCertify: true},
		{Name: "mac", Path: "apps/mac", Grade: "F", Score: 0.0, Units: 2, Unsupported: 2, HasCertify: true},
	}

	wc := AggregateCards(subs)

	if wc.PassRateKnown() {
		t.Error("PassRateKnown() = true, want false — nothing was assessed")
	}
	if wc.PassRate == 1.0 {
		t.Errorf("PassRate = %.1f%%, want not 100%%", wc.PassRate*100)
	}

	md := FormatWorkspaceCardMarkdown(wc)
	if !strings.Contains(md, "| Pass Rate | n/a |") {
		t.Errorf("workspace card should render the pass rate as n/a, got:\n%s", md)
	}
	if strings.Contains(md, "| Pass Rate | 100.0% |") {
		t.Errorf("workspace card claims a 100%% pass rate for an unassessed workspace:\n%s", md)
	}
	if !strings.Contains(md, "| Not Assessed | 5 |") {
		t.Errorf("workspace card should report the unassessed count, got:\n%s", md)
	}
}

// TestSubmoduleSummary_UnassessedRendersNoRate asserts on the per-submodule
// rendering, which has its own pass-rate column.
func TestSubmoduleSummary_UnassessedRendersNoRate(t *testing.T) {
	s := SubmoduleSummary{Name: "ios", Path: "apps/ios", Grade: "F", Units: 3, Unsupported: 3, HasCertify: true}

	out := formatSubmoduleSummary(s)
	if !strings.Contains(out, "| **Pass Rate** | n/a |") {
		t.Errorf("submodule summary should render the pass rate as n/a, got:\n%s", out)
	}
	if !strings.Contains(out, "| **Not Assessed** | 3 |") {
		t.Errorf("submodule summary should report the unassessed count, got:\n%s", out)
	}
}
