package policy

import (
	"fmt"

	"github.com/code-certification/certify/internal/domain"
)

// EvaluationResult holds the outcome of evaluating rules against evidence.
type EvaluationResult struct {
	Violations []domain.Violation
	Passed     bool // True if no error/critical violations
}

// Evaluate checks policy rules against collected evidence.
func Evaluate(rules []domain.PolicyRule, ev []domain.Evidence) EvaluationResult {
	var violations []domain.Violation
	hasBlockingViolation := false

	for _, rule := range rules {
		v := evaluateRule(rule, ev)
		if v != nil {
			violations = append(violations, *v)
			if rule.Severity == domain.SeverityError || rule.Severity == domain.SeverityCritical {
				hasBlockingViolation = true
			}
		}
	}

	return EvaluationResult{
		Violations: violations,
		Passed:     !hasBlockingViolation,
	}
}

func evaluateRule(rule domain.PolicyRule, ev []domain.Evidence) *domain.Violation {
	// Find relevant evidence
	metric := extractMetric(rule.Metric, ev)
	if metric < 0 {
		// Missing evidence
		return &domain.Violation{
			RuleID:      rule.ID,
			Severity:    rule.Severity,
			Description: fmt.Sprintf("missing evidence for metric %q", rule.Metric),
			Dimension:   rule.Dimension,
		}
	}

	// Check threshold (metric must be <= threshold)
	if metric > rule.Threshold {
		return &domain.Violation{
			RuleID:      rule.ID,
			Severity:    rule.Severity,
			Description: fmt.Sprintf("%s: %.0f exceeds threshold %.0f", rule.Metric, metric, rule.Threshold),
			Dimension:   rule.Dimension,
		}
	}

	return nil
}

// extractMetric pulls a numeric value from evidence by metric name.
// Returns -1 if not found.
func extractMetric(metric string, ev []domain.Evidence) float64 {
	for _, e := range ev {
		switch metric {
		case "lint_errors":
			if e.Kind == domain.EvidenceKindLint {
				if r, ok := e.Details.(interface{ ErrorCount() int }); ok {
					return float64(r.ErrorCount())
				}
				// Try map access or direct struct field
				if !e.Passed {
					return 1 // Lint failed = at least 1 error
				}
				return 0
			}
		case "test_failures":
			if e.Kind == domain.EvidenceKindTest {
				if !e.Passed {
					return 1
				}
				return 0
			}
		case "todo_count":
			if e.Kind == domain.EvidenceKindMetrics {
				return extractTodoCount(e)
			}
		case "test_coverage":
			if e.Kind == domain.EvidenceKindTest {
				return extractCoverage(e)
			}
		case "complexity":
			if e.Kind == domain.EvidenceKindMetrics {
				return 0 // TODO: implement complexity extraction
			}
		}
	}
	return -1 // Not found
}

// extractTodoCount tries to get TodoCount from metrics evidence details.
func extractTodoCount(e domain.Evidence) float64 {
	// Try to extract from the Details field (CodeMetrics struct)
	if m, ok := e.Details.(interface{ GetTodoCount() int }); ok {
		return float64(m.GetTodoCount())
	}
	// Try map-based access
	if m, ok := e.Details.(map[string]any); ok {
		if v, ok := m["todo_count"]; ok {
			if f, ok := v.(float64); ok {
				return f
			}
		}
	}
	// Parse from summary: "N lines (N code, N comment, N blank), N TODOs"
	var count int
	// Use strings approach since Sscanf can be finicky
	parts := splitSummaryParts(e.Summary)
	if parts >= 0 {
		return float64(parts)
	}
	_ = count
	return 0
}

func splitSummaryParts(summary string) int {
	// Look for "N TODOs" in summary
	var n int
	for i := len(summary) - 1; i >= 0; i-- {
		if summary[i] == ',' {
			sub := summary[i+1:]
			if _, err := fmt.Sscanf(sub, " %d TODO", &n); err == nil {
				return n
			}
			break
		}
	}
	return -1
}

// extractCoverage pulls coverage from test evidence.
func extractCoverage(e domain.Evidence) float64 {
	var coverage float64
	_, _ = fmt.Sscanf(e.Summary, "%*s %*d/%*d passed (%f%% coverage)", &coverage)
	return coverage / 100
}
