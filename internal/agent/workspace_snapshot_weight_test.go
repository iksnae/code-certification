package agent

import (
	"fmt"
	"testing"

	"github.com/iksnae/code-certification/internal/domain"
)

func goRecords(prefix string, n int, score float64) []domain.CertificationRecord {
	var recs []domain.CertificationRecord
	for i := 0; i < n; i++ {
		p := fmt.Sprintf("src/%s%d.go", prefix, i)
		recs = append(recs, domain.CertificationRecord{
			UnitID:   domain.NewUnitID("go", p, "F"),
			UnitPath: p,
			Status:   domain.StatusCertified,
			Grade:    domain.GradeFromScore(score),
			Score:    score,
		})
	}
	return recs
}

func swiftRecords(prefix string, n int) []domain.CertificationRecord {
	var recs []domain.CertificationRecord
	for i := 0; i < n; i++ {
		p := fmt.Sprintf("src/%s%d.swift", prefix, i)
		recs = append(recs, domain.CertificationRecord{
			UnitID:      domain.NewUnitID("swift", p, "S"),
			UnitPath:    p,
			Status:      domain.StatusExempt,
			Grade:       domain.GradeNA,
			Unsupported: true,
		})
	}
	return recs
}

// TestBuildWorkspaceSnapshot_WeightIsAnalyzableUnits covers the workspace
// weighting in internal/agent, which the mutation battery caught as the one
// unconstrained line after this branch's fixes: the harness mutated the weight
// back to TotalUnits and no test noticed.
//
// The fixture makes all four combinations distinguishable, because a weight and
// a divisor that both count the same wrong population cancel and prove nothing.
// api is 10 analyzable at 0.9; app is 4 analyzable at 0.5 plus 6 the engine
// cannot read:
//
//	correct      (weight analyzable, divide analyzable): 11/14 = 0.785714
//	weight-only  (weight all units,  divide analyzable): 14/14 = 1.000000
//	divisor-only (weight analyzable, divide all units):  11/20 = 0.550000
//	both         (weight all units,  divide all units):  14/20 = 0.700000
func TestBuildWorkspaceSnapshot_WeightIsAnalyzableUnits(t *testing.T) {
	root := t.TempDir()
	subs := []SubmoduleInfo{
		{Name: "api", Path: "services/api", Records: goRecords("a", 10, 0.9)},
		{Name: "app", Path: "apps/app", Records: append(
			goRecords("b", 4, 0.5), swiftRecords("v", 6)...)},
	}

	snap := BuildWorkspaceSnapshot(root, subs)

	if got := snap.AggregateMetrics.TotalUnitsAcrossAll; got != 14 {
		t.Fatalf("TotalUnitsAcrossAll = %d, want 14 (10 + 4 analyzable) — fixture drifted", got)
	}

	const want = 11.0 / 14.0
	if diff := snap.AggregateMetrics.WeightedAvgScore - want; diff > 1e-9 || diff < -1e-9 {
		t.Errorf("WeightedAvgScore = %.6f, want %.6f\n"+
			"  1.000000 means the weight counted unassessed units\n"+
			"  0.550000 means the divisor counted them\n"+
			"  0.700000 means both did",
			snap.AggregateMetrics.WeightedAvgScore, want)
	}
}

// TestBuildWorkspaceSnapshot_AllUnassessedSubmoduleIsNotTheWorst. With no
// analyzable units a submodule has no score to compare, so it must not win the
// worst-submodule label with a manufactured zero — that label goes into the
// prompt as the thing most in need of attention.
func TestBuildWorkspaceSnapshot_AllUnassessedSubmoduleIsNotTheWorst(t *testing.T) {
	root := t.TempDir()
	subs := []SubmoduleInfo{
		{Name: "api", Path: "services/api", Records: goRecords("a", 5, 0.9)},
		{Name: "ios", Path: "apps/ios", Records: swiftRecords("v", 5)},
	}

	snap := BuildWorkspaceSnapshot(root, subs)

	if got := snap.AggregateMetrics.WorstSubmodule; got == "ios" {
		t.Errorf("WorstSubmodule = %q — nothing in it was scored", got)
	}
	if got := snap.AggregateMetrics.TotalUnitsAcrossAll; got != 5 {
		t.Errorf("TotalUnitsAcrossAll = %d, want 5", got)
	}
}
