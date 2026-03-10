package workspace

import (
	"testing"
)

func TestAggregateCards_MultipleSubmodules(t *testing.T) {
	subs := []SubmoduleSummary{
		{Name: "api", Path: "services/api", Grade: "A", Score: 0.95, Units: 100, Passing: 100, Failing: 0, HasCertify: true},
		{Name: "web", Path: "services/web", Grade: "B", Score: 0.82, Units: 50, Passing: 45, Failing: 5, HasCertify: true},
	}

	wc := AggregateCards(subs)

	if wc.TotalUnits != 150 {
		t.Errorf("TotalUnits = %d, want 150", wc.TotalUnits)
	}
	if wc.TotalPassing != 145 {
		t.Errorf("TotalPassing = %d, want 145", wc.TotalPassing)
	}
	if wc.TotalFailing != 5 {
		t.Errorf("TotalFailing = %d, want 5", wc.TotalFailing)
	}

	// Weighted average: (0.95*100 + 0.82*50) / 150 = (95+41)/150 = 0.9066...
	expectedScore := (0.95*100 + 0.82*50) / 150.0
	if wc.OverallScore < expectedScore-0.01 || wc.OverallScore > expectedScore+0.01 {
		t.Errorf("OverallScore = %f, want ~%f", wc.OverallScore, expectedScore)
	}

	if wc.OverallGrade == "" {
		t.Error("OverallGrade should not be empty")
	}

	expectedPassRate := 145.0 / 150.0
	if wc.PassRate < expectedPassRate-0.01 || wc.PassRate > expectedPassRate+0.01 {
		t.Errorf("PassRate = %f, want ~%f", wc.PassRate, expectedPassRate)
	}

	if len(wc.Submodules) != 2 {
		t.Errorf("Submodules count = %d, want 2", len(wc.Submodules))
	}
}

func TestAggregateCards_SingleSubmodule(t *testing.T) {
	subs := []SubmoduleSummary{
		{Name: "only", Path: "only", Grade: "B+", Score: 0.88, Units: 200, Passing: 190, Failing: 10, HasCertify: true},
	}

	wc := AggregateCards(subs)

	if wc.TotalUnits != 200 {
		t.Errorf("TotalUnits = %d, want 200", wc.TotalUnits)
	}
	if wc.OverallScore < 0.87 || wc.OverallScore > 0.89 {
		t.Errorf("OverallScore = %f, want ~0.88", wc.OverallScore)
	}
}

func TestAggregateCards_MixedGrades(t *testing.T) {
	subs := []SubmoduleSummary{
		{Name: "excellent", Path: "a", Grade: "A", Score: 0.96, Units: 50, Passing: 50, Failing: 0, HasCertify: true},
		{Name: "poor", Path: "b", Grade: "D", Score: 0.55, Units: 50, Passing: 30, Failing: 20, HasCertify: true},
	}

	wc := AggregateCards(subs)

	// Weighted average: (0.96*50 + 0.55*50) / 100 = 0.755
	expectedScore := (0.96*50 + 0.55*50) / 100.0
	if wc.OverallScore < expectedScore-0.01 || wc.OverallScore > expectedScore+0.01 {
		t.Errorf("OverallScore = %f, want ~%f", wc.OverallScore, expectedScore)
	}
	if wc.TotalPassing != 80 {
		t.Errorf("TotalPassing = %d, want 80", wc.TotalPassing)
	}
	if wc.TotalFailing != 20 {
		t.Errorf("TotalFailing = %d, want 20", wc.TotalFailing)
	}
}

func TestAggregateCards_Empty(t *testing.T) {
	wc := AggregateCards(nil)

	if wc.TotalUnits != 0 {
		t.Errorf("TotalUnits = %d, want 0", wc.TotalUnits)
	}
	if wc.OverallGrade != "N/A" {
		t.Errorf("OverallGrade = %q, want N/A", wc.OverallGrade)
	}
}

func TestAggregateCards_SkipsUnconfigured(t *testing.T) {
	subs := []SubmoduleSummary{
		{Name: "configured", Path: "a", Grade: "A", Score: 0.95, Units: 100, Passing: 100, HasCertify: true},
		{Name: "unconfigured", Path: "b", Grade: "", Score: 0, Units: 0, HasCertify: false},
	}

	wc := AggregateCards(subs)

	// Only configured submodule should count
	if wc.TotalUnits != 100 {
		t.Errorf("TotalUnits = %d, want 100 (unconfigured excluded)", wc.TotalUnits)
	}
}
