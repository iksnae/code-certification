// Package engine contains the certification scoring and status logic.
package engine

import (
	"fmt"
	"strings"

	"github.com/code-certification/certify/internal/domain"
	"github.com/code-certification/certify/internal/policy"
)

// Score computes dimension scores from evidence and evaluation results.
func Score(ev []domain.Evidence, evalResult policy.EvaluationResult) domain.DimensionScores {
	scores := make(domain.DimensionScores)

	// Base all dimensions at 0.80 (neutral — evidence adjusts up or down)
	for _, d := range domain.AllDimensions() {
		scores[d] = 0.80
	}

	// Adjust based on evidence
	for _, e := range ev {
		switch e.Kind {
		case domain.EvidenceKindLint:
			if e.Passed {
				scores[domain.DimCorrectness] = max(scores[domain.DimCorrectness], 0.95)
			} else {
				scores[domain.DimCorrectness] = min(scores[domain.DimCorrectness], 0.4)
			}

		case domain.EvidenceKindTest:
			if e.Passed {
				scores[domain.DimTestability] = max(scores[domain.DimTestability], 0.90)
				// Boost further with coverage
				cov := extractSummaryFloat(e.Summary, "coverage")
				if cov > 0 {
					if cov >= 80 {
						scores[domain.DimTestability] = max(scores[domain.DimTestability], 0.95)
					} else if cov >= 60 {
						scores[domain.DimTestability] = max(scores[domain.DimTestability], 0.85)
					}
				}
			} else {
				scores[domain.DimTestability] = min(scores[domain.DimTestability], 0.3)
			}

		case domain.EvidenceKindMetrics:
			scoreFromMetrics(e, scores)

		case domain.EvidenceKindGitHistory:
			scoreFromGitHistory(e, scores)

		case domain.EvidenceKindAgentReview:
			// Agent review provides direct confidence-weighted scores
			if e.Passed {
				for _, d := range domain.AllDimensions() {
					scores[d] = max(scores[d], 0.85)
				}
			}
		}
	}

	// Penalize for violations
	for _, v := range evalResult.Violations {
		penalty := severityPenalty(v.Severity)
		scores[v.Dimension] = max(0, scores[v.Dimension]-penalty)
	}

	return scores
}

func scoreFromMetrics(e domain.Evidence, scores domain.DimensionScores) {
	summary := e.Summary

	// Extract complexity
	complexity := extractSummaryInt(summary, "complexity")
	if complexity >= 0 {
		switch {
		case complexity <= 5:
			scores[domain.DimMaintainability] = max(scores[domain.DimMaintainability], 0.95)
		case complexity <= 10:
			scores[domain.DimMaintainability] = max(scores[domain.DimMaintainability], 0.85)
		case complexity <= 20:
			scores[domain.DimMaintainability] = max(scores[domain.DimMaintainability], 0.70)
		default:
			scores[domain.DimMaintainability] = min(scores[domain.DimMaintainability], 0.50)
		}
	}

	// Extract code lines for readability
	codeLines := extractSummaryInt(summary, "code")
	if codeLines >= 0 {
		switch {
		case codeLines <= 50:
			scores[domain.DimReadability] = max(scores[domain.DimReadability], 0.95)
		case codeLines <= 150:
			scores[domain.DimReadability] = max(scores[domain.DimReadability], 0.85)
		case codeLines <= 300:
			scores[domain.DimReadability] = max(scores[domain.DimReadability], 0.75)
		default:
			scores[domain.DimReadability] = min(scores[domain.DimReadability], 0.60)
		}
	}

	// TODO count already handled by violation penalties
}

func scoreFromGitHistory(e domain.Evidence, scores domain.DimensionScores) {
	summary := e.Summary

	// Multiple authors = lower change risk (bus factor)
	authors := extractSummaryInt(summary, "author")
	if authors > 1 {
		scores[domain.DimChangeRisk] = max(scores[domain.DimChangeRisk], 0.90)
	}

	// More commits = more stable, better operational quality
	commits := extractSummaryInt(summary, "commit")
	if commits > 10 {
		scores[domain.DimOperationalQuality] = max(scores[domain.DimOperationalQuality], 0.85)
	}
}

// extractSummaryInt pulls an integer near a keyword in a summary string.
// Handles both "42 code" (number before keyword) and "complexity 2" (number after keyword).
func extractSummaryInt(summary, keyword string) int {
	idx := strings.Index(summary, keyword)
	if idx < 0 {
		return -1
	}

	// Try number AFTER keyword (e.g., "complexity 2")
	after := strings.TrimSpace(summary[idx+len(keyword):])
	var n int
	if _, err := fmt.Sscanf(after, "%d", &n); err == nil {
		return n
	}

	// Try number BEFORE keyword (e.g., "42 code")
	if idx > 0 {
		sub := strings.TrimSpace(summary[:idx])
		for i := len(sub) - 1; i >= 0; i-- {
			if sub[i] == ' ' || sub[i] == '(' || sub[i] == ',' {
				if _, err := fmt.Sscanf(strings.TrimSpace(sub[i+1:]), "%d", &n); err == nil {
					return n
				}
				break
			}
			if i == 0 {
				if _, err := fmt.Sscanf(strings.TrimSpace(sub), "%d", &n); err == nil {
					return n
				}
			}
		}
	}
	return -1
}

// extractSummaryFloat pulls a float before a keyword.
func extractSummaryFloat(summary, keyword string) float64 {
	idx := strings.Index(summary, keyword)
	if idx <= 0 {
		return -1
	}
	sub := strings.TrimSpace(summary[:idx])
	for i := len(sub) - 1; i >= 0; i-- {
		if sub[i] == ' ' || sub[i] == '(' || sub[i] == ',' {
			var f float64
			if _, err := fmt.Sscanf(strings.TrimSpace(sub[i+1:]), "%f", &f); err == nil {
				return f
			}
			break
		}
	}
	return -1
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
