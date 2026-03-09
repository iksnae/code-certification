package domain

import "fmt"

// Dimension represents a quality dimension used for certification scoring.
type Dimension int

const (
	DimCorrectness                Dimension = iota // Code does what it claims
	DimMaintainability                             // Ease of future modification
	DimReadability                                 // Clarity and understandability
	DimTestability                                 // Ease of testing
	DimSecurity                                    // Security posture
	DimArchitecturalFitness                        // Alignment with architecture
	DimOperationalQuality                          // Production readiness
	DimPerformanceAppropriateness                  // Performance fitness
	DimChangeRisk                                  // Risk introduced by changes
)

var dimensionStrings = map[Dimension]string{
	DimCorrectness:                "correctness",
	DimMaintainability:            "maintainability",
	DimReadability:                "readability",
	DimTestability:                "testability",
	DimSecurity:                   "security",
	DimArchitecturalFitness:       "architectural_fitness",
	DimOperationalQuality:         "operational_quality",
	DimPerformanceAppropriateness: "performance_appropriateness",
	DimChangeRisk:                 "change_risk",
}

// String returns the string representation of a Dimension.
func (d Dimension) String() string {
	if s, ok := dimensionStrings[d]; ok {
		return s
	}
	return fmt.Sprintf("Dimension(%d)", d)
}

// AllDimensions returns all 9 certification dimensions in canonical order.
func AllDimensions() []Dimension {
	return []Dimension{
		DimCorrectness,
		DimMaintainability,
		DimReadability,
		DimTestability,
		DimSecurity,
		DimArchitecturalFitness,
		DimOperationalQuality,
		DimPerformanceAppropriateness,
		DimChangeRisk,
	}
}

// DimensionScores maps each dimension to a score between 0.0 and 1.0.
type DimensionScores map[Dimension]float64

// DimensionWeights maps each dimension to its relative weight.
// A nil or empty weights map means equal weighting.
type DimensionWeights map[Dimension]float64

// WeightedAverage computes the weighted average across all dimensions.
// If weights is nil, all dimensions are weighted equally.
func (ds DimensionScores) WeightedAverage(weights DimensionWeights) float64 {
	dims := AllDimensions()

	if len(weights) == 0 {
		// Equal weights
		var sum float64
		var count int
		for _, d := range dims {
			if score, ok := ds[d]; ok {
				sum += score
				count++
			}
		}
		if count == 0 {
			return 0
		}
		return sum / float64(count)
	}

	// Weighted average
	var weightedSum, totalWeight float64
	for _, d := range dims {
		w, hasWeight := weights[d]
		if !hasWeight {
			continue
		}
		score, hasScore := ds[d]
		if !hasScore {
			continue
		}
		weightedSum += score * w
		totalWeight += w
	}
	if totalWeight == 0 {
		return 0
	}
	return weightedSum / totalWeight
}

// Grade represents a letter grade computed from a certification score.
type Grade int

const (
	GradeA      Grade = iota // 0.93+
	GradeAMinus              // 0.90+
	GradeBPlus               // 0.87+
	GradeB                   // 0.80+
	GradeC                   // 0.70+
	GradeD                   // 0.60+
	GradeF                   // < 0.60
)

var gradeStrings = map[Grade]string{
	GradeA:      "A",
	GradeAMinus: "A-",
	GradeBPlus:  "B+",
	GradeB:      "B",
	GradeC:      "C",
	GradeD:      "D",
	GradeF:      "F",
}

// String returns the letter grade string.
func (g Grade) String() string {
	if s, ok := gradeStrings[g]; ok {
		return s
	}
	return fmt.Sprintf("Grade(%d)", g)
}

// GradeFromScore converts a numeric score (0.0–1.0) to a letter grade.
func GradeFromScore(score float64) Grade {
	switch {
	case score >= 0.93:
		return GradeA
	case score >= 0.90:
		return GradeAMinus
	case score >= 0.87:
		return GradeBPlus
	case score >= 0.80:
		return GradeB
	case score >= 0.70:
		return GradeC
	case score >= 0.60:
		return GradeD
	default:
		return GradeF
	}
}
