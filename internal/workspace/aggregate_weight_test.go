package workspace

import "testing"

// TestAggregateCards_WeightIsAnalyzableUnits covers the weight in the workspace
// rollup — `s.Score * float64(s.Analyzable())` — which every other test in this
// package leaves unconstrained.
//
// The reason is specific and worth stating: the existing fixtures give their
// unassessed submodule Score 0.0, so weighting by Units instead of Analyzable
// multiplies zero by a different number and lands on the same zero. The weight
// only becomes observable when a submodule is PARTLY unassessed and still has a
// real score, which is the ordinary case for a mixed-language repository and
// the case no fixture had.
//
// Weight and divisor do not cancel. They cancel only when the unsupported
// fraction is identical across submodules; here it is 0/10 and 6/10, so all
// four combinations are distinguishable:
//
//	correct      (weight analyzable, divide analyzable): 11/14 = 0.785714
//	weight-only  (weight units,      divide analyzable): 14/14 = 1.000000
//	divisor-only (weight analyzable, divide all units):  11/20 = 0.550000
//	both         (weight units,      divide all units):  14/20 = 0.700000
func TestAggregateCards_WeightIsAnalyzableUnits(t *testing.T) {
	subs := []SubmoduleSummary{
		{Name: "api", Path: "services/api", Grade: "A", Score: 0.9, Units: 10, Passing: 10, HasCertify: true},
		{Name: "app", Path: "apps/app", Grade: "C", Score: 0.5, Units: 10, Passing: 2, Unsupported: 6, HasCertify: true},
	}

	wc := AggregateCards(subs)

	if wc.AnalyzableUnits != 14 {
		t.Fatalf("AnalyzableUnits = %d, want 14 — fixture drifted", wc.AnalyzableUnits)
	}

	const want = 11.0 / 14.0
	if diff := wc.OverallScore - want; diff > 1e-9 || diff < -1e-9 {
		t.Errorf("OverallScore = %.6f, want %.6f\n"+
			"  1.000000 means the weight counted unassessed units\n"+
			"  0.550000 means the divisor counted unassessed units\n"+
			"  0.700000 means both did",
			wc.OverallScore, want)
	}
}
