package domain_test

import (
	"testing"

	"github.com/iksnae/code-certification/internal/domain"
)

func TestDimension_AllNine(t *testing.T) {
	dims := domain.AllDimensions()
	if len(dims) != 9 {
		t.Fatalf("AllDimensions() returned %d, want 9", len(dims))
	}

	expected := []domain.Dimension{
		domain.DimCorrectness,
		domain.DimMaintainability,
		domain.DimReadability,
		domain.DimTestability,
		domain.DimSecurity,
		domain.DimArchitecturalFitness,
		domain.DimOperationalQuality,
		domain.DimPerformanceAppropriateness,
		domain.DimChangeRisk,
	}
	for i, d := range expected {
		if dims[i] != d {
			t.Errorf("AllDimensions()[%d] = %v, want %v", i, dims[i], d)
		}
	}
}

func TestDimension_String(t *testing.T) {
	tests := []struct {
		d    domain.Dimension
		want string
	}{
		{domain.DimCorrectness, "correctness"},
		{domain.DimMaintainability, "maintainability"},
		{domain.DimReadability, "readability"},
		{domain.DimTestability, "testability"},
		{domain.DimSecurity, "security"},
		{domain.DimArchitecturalFitness, "architectural_fitness"},
		{domain.DimOperationalQuality, "operational_quality"},
		{domain.DimPerformanceAppropriateness, "performance_appropriateness"},
		{domain.DimChangeRisk, "change_risk"},
	}
	for _, tt := range tests {
		if got := tt.d.String(); got != tt.want {
			t.Errorf("%v.String() = %q, want %q", tt.d, got, tt.want)
		}
	}
}

func TestDimensionScores_WeightedAverage_EqualWeights(t *testing.T) {
	scores := domain.DimensionScores{
		domain.DimCorrectness:                0.8,
		domain.DimMaintainability:            0.8,
		domain.DimReadability:                0.8,
		domain.DimTestability:                0.8,
		domain.DimSecurity:                   0.8,
		domain.DimArchitecturalFitness:       0.8,
		domain.DimOperationalQuality:         0.8,
		domain.DimPerformanceAppropriateness: 0.8,
		domain.DimChangeRisk:                 0.8,
	}

	avg := scores.WeightedAverage(nil) // nil = equal weights
	if avg < 0.799 || avg > 0.801 {
		t.Errorf("WeightedAverage(nil) = %f, want ~0.8", avg)
	}
}

func TestDimensionScores_WeightedAverage_CustomWeights(t *testing.T) {
	scores := domain.DimensionScores{
		domain.DimCorrectness:                1.0,
		domain.DimMaintainability:            0.0,
		domain.DimReadability:                0.0,
		domain.DimTestability:                0.0,
		domain.DimSecurity:                   0.0,
		domain.DimArchitecturalFitness:       0.0,
		domain.DimOperationalQuality:         0.0,
		domain.DimPerformanceAppropriateness: 0.0,
		domain.DimChangeRisk:                 0.0,
	}

	weights := domain.DimensionWeights{
		domain.DimCorrectness: 1.0, // only correctness matters
	}

	avg := scores.WeightedAverage(weights)
	if avg < 0.999 || avg > 1.001 {
		t.Errorf("WeightedAverage(correctness-only) = %f, want ~1.0", avg)
	}
}

func TestDimensionScores_Grade(t *testing.T) {
	tests := []struct {
		avg  float64
		want domain.Grade
	}{
		{0.95, domain.GradeA},
		{0.91, domain.GradeAMinus},
		{0.88, domain.GradeBPlus},
		{0.83, domain.GradeB},
		{0.72, domain.GradeC},
		{0.62, domain.GradeD},
		{0.50, domain.GradeF},
		{0.0, domain.GradeF},
	}
	for _, tt := range tests {
		got := domain.GradeFromScore(tt.avg)
		if got != tt.want {
			t.Errorf("GradeFromScore(%f) = %v, want %v", tt.avg, got, tt.want)
		}
	}
}

func TestGrade_String(t *testing.T) {
	tests := []struct {
		g    domain.Grade
		want string
	}{
		{domain.GradeA, "A"},
		{domain.GradeAMinus, "A-"},
		{domain.GradeBPlus, "B+"},
		{domain.GradeB, "B"},
		{domain.GradeC, "C"},
		{domain.GradeD, "D"},
		{domain.GradeF, "F"},
	}
	for _, tt := range tests {
		if got := tt.g.String(); got != tt.want {
			t.Errorf("Grade(%d).String() = %q, want %q", tt.g, got, tt.want)
		}
	}
}
