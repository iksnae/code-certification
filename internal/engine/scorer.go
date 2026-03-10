// Package engine contains the certification scoring and status logic.
package engine

import (
	"fmt"
	"strings"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/policy"
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
				cov := metricOrSummaryFloat(e, "test_coverage", "coverage")
				if cov > 0 {
					// Metrics stores as 0.0-1.0, summary as percentage
					covPct := cov
					if covPct <= 1.0 {
						covPct = cov * 100
					}
					if covPct >= 80 {
						scores[domain.DimTestability] = max(scores[domain.DimTestability], 0.95)
					} else if covPct >= 60 {
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

		case domain.EvidenceKindStructural:
			scoreFromStructural(e, scores)

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
	// Extract complexity — prefer Metrics, fall back to Summary
	complexity := metricOrSummaryInt(e, "complexity", "complexity")
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

	// Extract code lines for readability — prefer Metrics, fall back to Summary
	codeLines := metricOrSummaryInt(e, "code_lines", "code")
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

	// Note: todo_count observation already handled by policy violation penalties
}

func scoreFromStructural(e domain.Evidence, scores domain.DimensionScores) {
	if e.Metrics == nil {
		return
	}
	scoreStructuralReadability(e.Metrics, scores)
	scoreStructuralCorrectness(e.Metrics, scores)
	scoreStructuralArchitecture(e.Metrics, scores)
}

// scoreStructuralReadability adjusts readability and maintainability
// based on code shape metrics: docs, params, nesting, length.
func scoreStructuralReadability(m map[string]float64, scores domain.DimensionScores) {
	if m["has_doc_comment"] == 1.0 {
		scores[domain.DimReadability] = max(scores[domain.DimReadability], 0.90)
	} else if m["exported_name"] == 1.0 {
		scores[domain.DimReadability] = min(scores[domain.DimReadability], 0.70)
	}

	if params := int(m["param_count"]); params > 5 {
		penalty := float64(params-5) * 0.10
		scores[domain.DimMaintainability] = max(0, scores[domain.DimMaintainability]-penalty)
	}

	if nesting := int(m["max_nesting_depth"]); nesting > 3 {
		penalty := float64(nesting-3) * 0.05
		scores[domain.DimReadability] = max(0, scores[domain.DimReadability]-penalty)
	}

	if naked := int(m["naked_returns"]); naked > 0 {
		penalty := float64(naked) * 0.05
		scores[domain.DimReadability] = max(0, scores[domain.DimReadability]-penalty)
	}

	funcLines := int(m["func_lines"])
	if funcLines > 0 {
		switch {
		case funcLines <= 30:
			scores[domain.DimReadability] = max(scores[domain.DimReadability], 0.90)
		case funcLines <= 60:
			// neutral
		case funcLines <= 100:
			scores[domain.DimReadability] = min(scores[domain.DimReadability], 0.70)
		default:
			scores[domain.DimReadability] = min(scores[domain.DimReadability], 0.50)
			scores[domain.DimMaintainability] = min(scores[domain.DimMaintainability], 0.60)
		}
	}

	if methodCount := int(m["method_count"]); methodCount > 15 {
		scores[domain.DimMaintainability] = min(scores[domain.DimMaintainability], 0.50)
	} else if methodCount > 10 {
		scores[domain.DimMaintainability] = min(scores[domain.DimMaintainability], 0.65)
	}
}

// scoreStructuralCorrectness adjusts correctness, testability, security,
// and performance based on error handling, panics, exits, and global state.
func scoreStructuralCorrectness(m map[string]float64, scores domain.DimensionScores) {
	if ignored := int(m["errors_ignored"]); ignored > 0 {
		scores[domain.DimCorrectness] = min(scores[domain.DimCorrectness], 0.60)
	}

	if panicCalls := int(m["panic_calls"]); panicCalls > 0 {
		scores[domain.DimCorrectness] = min(scores[domain.DimCorrectness], 0.50)
	}

	if osExitCalls := int(m["os_exit_calls"]); osExitCalls > 0 {
		scores[domain.DimCorrectness] = min(scores[domain.DimCorrectness], 0.55)
		scores[domain.DimTestability] = min(scores[domain.DimTestability], 0.50)
	}

	if deferInLoop := int(m["defer_in_loop"]); deferInLoop > 0 {
		scores[domain.DimCorrectness] = min(scores[domain.DimCorrectness], 0.55)
		scores[domain.DimPerformanceAppropriateness] = min(scores[domain.DimPerformanceAppropriateness], 0.50)
	}

	if m["has_init_func"] == 1.0 {
		scores[domain.DimTestability] = min(scores[domain.DimTestability], 0.65)
		scores[domain.DimMaintainability] = min(scores[domain.DimMaintainability], 0.70)
	}

	if globalMut := int(m["global_mutable_count"]); globalMut > 0 {
		penalty := float64(globalMut) * 0.05
		scores[domain.DimSecurity] = max(0, scores[domain.DimSecurity]-penalty)
		scores[domain.DimTestability] = min(scores[domain.DimTestability], 0.65)
	}
}

// scoreStructuralArchitecture adjusts architectural fitness based on
// API design patterns: context position, god objects.
func scoreStructuralArchitecture(m map[string]float64, scores domain.DimensionScores) {
	if m["context_not_first"] == 1.0 {
		scores[domain.DimCorrectness] = min(scores[domain.DimCorrectness], 0.70)
		scores[domain.DimArchitecturalFitness] = min(scores[domain.DimArchitecturalFitness], 0.65)
	}

	if methodCount := int(m["method_count"]); methodCount > 15 {
		scores[domain.DimArchitecturalFitness] = min(scores[domain.DimArchitecturalFitness], 0.55)
	}
}

func scoreFromGitHistory(e domain.Evidence, scores domain.DimensionScores) {
	// Multiple authors = lower change risk (bus factor)
	authors := metricOrSummaryInt(e, "author_count", "author")
	if authors > 1 {
		scores[domain.DimChangeRisk] = max(scores[domain.DimChangeRisk], 0.90)
	}

	// More commits = more stable, better operational quality
	commits := metricOrSummaryInt(e, "commit_count", "commit")
	if commits > 10 {
		scores[domain.DimOperationalQuality] = max(scores[domain.DimOperationalQuality], 0.85)
	}
}

// metricOrSummaryInt returns an int metric from e.Metrics[metricKey], falling back to
// extractSummaryInt(e.Summary, summaryKeyword) for backward compatibility.
func metricOrSummaryInt(e domain.Evidence, metricKey, summaryKeyword string) int {
	if e.Metrics != nil {
		if v, ok := e.Metrics[metricKey]; ok {
			return int(v)
		}
	}
	return extractSummaryInt(e.Summary, summaryKeyword)
}

// metricOrSummaryFloat returns a float metric from e.Metrics[metricKey], falling back to
// extractSummaryFloat(e.Summary, summaryKeyword) for backward compatibility.
func metricOrSummaryFloat(e domain.Evidence, metricKey, summaryKeyword string) float64 {
	if e.Metrics != nil {
		if v, ok := e.Metrics[metricKey]; ok {
			return v
		}
	}
	return extractSummaryFloat(e.Summary, summaryKeyword)
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
