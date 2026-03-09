// Package engine contains the certification scoring and status logic.
package engine

import (
	"github.com/code-certification/certify/internal/domain"
	"github.com/code-certification/certify/internal/policy"
)

// Score computes dimension scores from evidence and evaluation results.
func Score(ev []domain.Evidence, evalResult policy.EvaluationResult) domain.DimensionScores {
	scores := make(domain.DimensionScores)

	// Base all dimensions at 0.8 (decent default)
	for _, d := range domain.AllDimensions() {
		scores[d] = 0.8
	}

	// Adjust based on evidence
	for _, e := range ev {
		switch e.Kind {
		case domain.EvidenceKindLint:
			if e.Passed {
				scores[domain.DimCorrectness] = max(scores[domain.DimCorrectness], 0.9)
			} else {
				scores[domain.DimCorrectness] = min(scores[domain.DimCorrectness], 0.5)
			}
		case domain.EvidenceKindTest:
			if e.Passed {
				scores[domain.DimTestability] = max(scores[domain.DimTestability], 0.9)
			} else {
				scores[domain.DimTestability] = min(scores[domain.DimTestability], 0.4)
			}
		case domain.EvidenceKindMetrics:
			// Metrics affect maintainability and readability
			scores[domain.DimMaintainability] = 0.8
			scores[domain.DimReadability] = 0.8
		}
	}

	// Penalize for violations
	for _, v := range evalResult.Violations {
		penalty := severityPenalty(v.Severity)
		scores[v.Dimension] = max(0, scores[v.Dimension]-penalty)
	}

	return scores
}

func severityPenalty(s domain.Severity) float64 {
	switch s {
	case domain.SeverityCritical:
		return 0.5
	case domain.SeverityError:
		return 0.3
	case domain.SeverityWarning:
		return 0.1
	default:
		return 0.05
	}
}

// StatusFromScore determines certification status from a weighted score.
func StatusFromScore(score float64, hasBlockingViolations bool) domain.Status {
	if hasBlockingViolations {
		return domain.StatusProbationary
	}

	switch {
	case score >= 0.80:
		return domain.StatusCertified
	case score >= 0.60:
		return domain.StatusCertifiedWithObservations
	case score >= 0.40:
		return domain.StatusProbationary
	default:
		return domain.StatusDecertified
	}
}
