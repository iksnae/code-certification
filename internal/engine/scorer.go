// Package engine contains the certification scoring and status logic.
package engine

import (
	"fmt"
	"strings"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/policy"
)

// Score computes dimension scores from evidence and evaluation results.
// Only dimensions with actual evidence are included in the returned map.
// Dimensions without evidence are absent — they don't dilute the average.
func Score(ev []domain.Evidence, evalResult policy.EvaluationResult) domain.DimensionScores {
	scores := make(domain.DimensionScores)

	for _, e := range ev {
		switch e.Kind {
		case domain.EvidenceKindLint:
			if e.Passed {
				setMax(scores, domain.DimCorrectness, 0.95)
			} else {
				setMin(scores, domain.DimCorrectness, 0.4)
			}

		case domain.EvidenceKindTest:
			scoreFromTest(e, scores)

		case domain.EvidenceKindMetrics:
			scoreFromMetrics(e, scores)

		case domain.EvidenceKindGitHistory:
			scoreFromGitHistory(e, scores)

		case domain.EvidenceKindStructural:
			scoreFromStructural(e, scores)

		case domain.EvidenceKindAgentReview:
			// Agent review boosts only dimensions it can actually assess from code
			if e.Passed {
				for _, d := range []domain.Dimension{
					domain.DimCorrectness,
					domain.DimMaintainability,
					domain.DimReadability,
					domain.DimTestability,
				} {
					setMax(scores, d, 0.85)
				}
			}
		}
	}

	// Penalize for violations — only affects dimensions already in the map
	for _, v := range evalResult.Violations {
		penalty := severityPenalty(v.Severity)
		if _, ok := scores[v.Dimension]; ok {
			scores[v.Dimension] = max(0, scores[v.Dimension]-penalty)
		} else {
			// Violation introduces the dimension at a penalized value
			scores[v.Dimension] = max(0, 0.80-penalty)
		}
	}

	return scores
}

// setMax sets scores[dim] = max(existing, value).
// If the dimension hasn't been set yet, it's initialized to value.
func setMax(scores domain.DimensionScores, dim domain.Dimension, value float64) {
	if existing, ok := scores[dim]; ok {
		scores[dim] = max(existing, value)
	} else {
		scores[dim] = value
	}
}

// setMin sets scores[dim] = min(existing, value).
// If the dimension hasn't been set yet, it's initialized to value.
func setMin(scores domain.DimensionScores, dim domain.Dimension, value float64) {
	if existing, ok := scores[dim]; ok {
		scores[dim] = min(existing, value)
	} else {
		scores[dim] = value
	}
}

func scoreFromTest(e domain.Evidence, scores domain.DimensionScores) {
	if !e.Passed {
		setMin(scores, domain.DimTestability, 0.3)
		return
	}
	setMax(scores, domain.DimTestability, 0.90)
	cov := metricOrSummaryFloat(e, "test_coverage", "coverage")
	if cov <= 0 {
		return
	}
	covPct := cov
	if covPct <= 1.0 {
		covPct = cov * 100
	}
	if covPct >= 80 {
		setMax(scores, domain.DimTestability, 0.95)
	} else if covPct >= 60 {
		setMax(scores, domain.DimTestability, 0.85)
	}
}

func scoreFromMetrics(e domain.Evidence, scores domain.DimensionScores) {
	complexity := metricOrSummaryInt(e, "complexity", "complexity")
	if complexity >= 0 {
		switch {
		case complexity <= 5:
			setMax(scores, domain.DimMaintainability, 0.95)
		case complexity <= 10:
			setMax(scores, domain.DimMaintainability, 0.85)
		case complexity <= 15:
			setMax(scores, domain.DimMaintainability, 0.80)
		case complexity <= 20:
			setMax(scores, domain.DimMaintainability, 0.70)
		default:
			setMin(scores, domain.DimMaintainability, 0.50)
		}
	}

	// code_lines from metrics evidence is a file-level metric.
	// Thresholds are generous because files naturally grow larger than functions.
	codeLines := metricOrSummaryInt(e, "code_lines", "code")
	if codeLines >= 0 {
		switch {
		case codeLines <= 100:
			setMax(scores, domain.DimReadability, 0.95)
		case codeLines <= 300:
			setMax(scores, domain.DimReadability, 0.90)
		case codeLines <= 500:
			setMax(scores, domain.DimReadability, 0.85)
		case codeLines <= 800:
			setMax(scores, domain.DimReadability, 0.75)
		default:
			setMin(scores, domain.DimReadability, 0.60)
		}
	}
}

func scoreFromStructural(e domain.Evidence, scores domain.DimensionScores) {
	if e.Metrics == nil {
		return
	}
	scoreStructuralReadability(e.Metrics, scores)
	scoreAlgoComplexity(e.Metrics, scores)
	scoreStructuralCorrectness(e.Metrics, scores)
	scoreStructuralArchitecture(e.Metrics, scores)
}

// scoreStructuralReadability adjusts readability and maintainability
// based on code shape metrics: docs, params, nesting, length.
func scoreStructuralReadability(m map[string]float64, scores domain.DimensionScores) {
	if m["has_doc_comment"] == 1.0 {
		setMax(scores, domain.DimReadability, 0.90)
	} else if m["exported_name"] == 1.0 {
		setMin(scores, domain.DimReadability, 0.70)
	}

	if params := int(m["param_count"]); params > 5 {
		penalty := float64(params-5) * 0.10
		if v, ok := scores[domain.DimMaintainability]; ok {
			scores[domain.DimMaintainability] = max(0, v-penalty)
		} else {
			scores[domain.DimMaintainability] = max(0, 0.80-penalty)
		}
	}

	if nesting := int(m["max_nesting_depth"]); nesting > 3 {
		penalty := float64(nesting-3) * 0.05
		if v, ok := scores[domain.DimReadability]; ok {
			scores[domain.DimReadability] = max(0, v-penalty)
		} else {
			scores[domain.DimReadability] = max(0, 0.80-penalty)
		}
	}

	if naked := int(m["naked_returns"]); naked > 0 {
		penalty := float64(naked) * 0.05
		if v, ok := scores[domain.DimReadability]; ok {
			scores[domain.DimReadability] = max(0, v-penalty)
		} else {
			scores[domain.DimReadability] = max(0, 0.80-penalty)
		}
	}

	funcLines := int(m["func_lines"])
	if funcLines > 0 {
		switch {
		case funcLines <= 30:
			setMax(scores, domain.DimReadability, 0.90)
		case funcLines <= 60:
			// neutral — don't set if not already set
		case funcLines <= 100:
			setMin(scores, domain.DimReadability, 0.70)
		default:
			setMin(scores, domain.DimReadability, 0.50)
			setMin(scores, domain.DimMaintainability, 0.60)
		}
	}

	if methodCount := int(m["method_count"]); methodCount > 15 {
		setMin(scores, domain.DimMaintainability, 0.50)
	} else if methodCount > 10 {
		setMin(scores, domain.DimMaintainability, 0.65)
	}
}

// scoreStructuralCorrectness adjusts correctness, testability, security,
// and performance based on error handling, panics, exits, and global state.
func scoreStructuralCorrectness(m map[string]float64, scores domain.DimensionScores) {
	if ignored := int(m["errors_ignored"]); ignored > 0 {
		setMin(scores, domain.DimCorrectness, 0.60)
	}

	if panicCalls := int(m["panic_calls"]); panicCalls > 0 {
		setMin(scores, domain.DimCorrectness, 0.50)
	}

	if osExitCalls := int(m["os_exit_calls"]); osExitCalls > 0 {
		setMin(scores, domain.DimCorrectness, 0.55)
		setMin(scores, domain.DimTestability, 0.50)
	}

	// defer_in_loop: penalty-only — introduces performance_appropriateness
	// only when the problem exists (steers you to fix it)
	if deferInLoop := int(m["defer_in_loop"]); deferInLoop > 0 {
		setMin(scores, domain.DimCorrectness, 0.55)
		setMin(scores, domain.DimPerformanceAppropriateness, 0.50)
	}

	if m["has_init_func"] == 1.0 {
		setMin(scores, domain.DimTestability, 0.65)
		setMin(scores, domain.DimMaintainability, 0.70)
	}

	// Security: measured via global mutable state analysis.
	// Clean = 0.85 (checked, no issues). Dirty = penalized below 0.80.
	if _, checked := m["global_mutable_count"]; checked {
		globalMut := int(m["global_mutable_count"])
		if globalMut > 0 {
			penalty := float64(globalMut) * 0.05
			setMin(scores, domain.DimSecurity, max(0, 0.85-penalty))
			setMin(scores, domain.DimTestability, 0.65)
		} else {
			setMax(scores, domain.DimSecurity, 0.85)
		}
	}
}

// scoreStructuralArchitecture adjusts architectural fitness based on
// API design patterns: context position, god objects.
// Penalty-only — architectural_fitness only enters the score when
// violations are found, steering you to fix them.
func scoreStructuralArchitecture(m map[string]float64, scores domain.DimensionScores) {
	if m["context_not_first"] == 1.0 {
		setMin(scores, domain.DimCorrectness, 0.70)
		setMin(scores, domain.DimArchitecturalFitness, 0.65)
	}

	if methodCount := int(m["method_count"]); methodCount > 15 {
		setMin(scores, domain.DimArchitecturalFitness, 0.55)
	}
}

// scoreAlgoComplexity sets performance_appropriateness based on algorithmic
// complexity metrics: loop nesting depth and recursive calls.
// Unlike architectural_fitness, this is always measured when structural
// evidence exists — O(1) code earns a high score, O(n²) earns a penalty.
func scoreAlgoComplexity(m map[string]float64, scores domain.DimensionScores) {
	loopDepth := int(m["loop_nesting_depth"])
	recursive := int(m["recursive_calls"])

	// Classify and score
	if recursive > 0 {
		setMin(scores, domain.DimPerformanceAppropriateness, 0.40)
		return
	}
	switch {
	case loopDepth >= 3:
		setMin(scores, domain.DimPerformanceAppropriateness, 0.50)
	case loopDepth == 2:
		setMax(scores, domain.DimPerformanceAppropriateness, 0.70)
	case loopDepth == 1:
		setMax(scores, domain.DimPerformanceAppropriateness, 0.90)
	default:
		setMax(scores, domain.DimPerformanceAppropriateness, 0.95)
	}
}

func scoreFromGitHistory(e domain.Evidence, scores domain.DimensionScores) {
	authors := metricOrSummaryInt(e, "author_count", "author")
	switch {
	case authors >= 3:
		setMax(scores, domain.DimChangeRisk, 0.95)
	case authors >= 2:
		setMax(scores, domain.DimChangeRisk, 0.90)
	case authors == 1:
		setMax(scores, domain.DimChangeRisk, 0.70)
	}

	commits := metricOrSummaryInt(e, "commit_count", "commit")
	switch {
	case commits > 50:
		setMax(scores, domain.DimOperationalQuality, 0.95)
	case commits > 20:
		setMax(scores, domain.DimOperationalQuality, 0.90)
	case commits > 10:
		setMax(scores, domain.DimOperationalQuality, 0.85)
	case commits > 0:
		setMax(scores, domain.DimOperationalQuality, 0.75)
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
