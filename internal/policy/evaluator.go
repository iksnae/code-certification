package policy

import (
	"fmt"
	"strings"

	"github.com/iksnae/code-certification/internal/domain"
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
//
// Priority: Metrics map lookup → legacy fallback (Summary parsing, Passed flag).
func extractMetric(metric string, ev []domain.Evidence) float64 {
	// First pass: direct Metrics map lookup across all evidence
	for _, e := range ev {
		if e.Metrics != nil {
			if v, ok := e.Metrics[metric]; ok {
				return v
			}
		}
	}

	// Fallback: legacy extraction for backward compatibility with old records
	for _, e := range ev {
		switch metric {
		case "lint_errors":
			if e.Kind == domain.EvidenceKindLint {
				if !e.Passed {
					return 1
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
				return extractComplexity(e)
			}
		}
	}
	return -1 // Not found
}

// extractTodoCount extracts todo count from summary string (legacy fallback).
func extractTodoCount(e domain.Evidence) float64 {
	// Parse "N TODOs" from summary
	idx := strings.Index(e.Summary, " TODO")
	if idx > 0 {
		// Walk backwards to find the number
		sub := strings.TrimSpace(e.Summary[:idx])
		// Find last separator
		for i := len(sub) - 1; i >= 0; i-- {
			if sub[i] == ' ' || sub[i] == ',' || sub[i] == '(' {
				var n int
				if _, err := fmt.Sscanf(strings.TrimSpace(sub[i+1:]), "%d", &n); err == nil {
					return float64(n)
				}
				break
			}
		}
	}
	return 0
}

// extractComplexity extracts complexity from summary string (legacy fallback).
func extractComplexity(e domain.Evidence) float64 {
	// Parse from summary: "... complexity N"
	var n int
	if _, err := fmt.Sscanf(e.Summary, "%*s complexity %d", &n); err == nil {
		return float64(n)
	}
	// Try scanning from end of summary
	parts := strings.Split(e.Summary, "complexity ")
	if len(parts) >= 2 {
		if _, err := fmt.Sscanf(parts[len(parts)-1], "%d", &n); err == nil {
			return float64(n)
		}
	}
	return 0
}

// extractCoverage pulls coverage from test evidence.
func extractCoverage(e domain.Evidence) float64 {
	var coverage float64
	if _, err := fmt.Sscanf(e.Summary, "%*s %*d/%*d passed (%f%% coverage)", &coverage); err != nil {
		return 0 // summary doesn't match coverage pattern
	}
	return coverage / 100
}
