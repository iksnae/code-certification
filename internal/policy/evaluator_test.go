package policy_test

import (
	"testing"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/evidence"
	"github.com/iksnae/code-certification/internal/policy"
)

func TestEvaluator_AllPass(t *testing.T) {
	rules := []domain.PolicyRule{
		{ID: "lint-clean", Dimension: domain.DimCorrectness, Severity: domain.SeverityError, Metric: "lint_errors", Threshold: 0},
		{ID: "test-pass", Dimension: domain.DimTestability, Severity: domain.SeverityError, Metric: "test_failures", Threshold: 0},
	}

	ev := []domain.Evidence{
		evidence.LintResult{Tool: "golangci-lint", ErrorCount: 0, WarnCount: 0}.ToEvidence(),
		evidence.TestResult{Tool: "go test", TotalCount: 10, PassedCount: 10, FailedCount: 0}.ToEvidence(),
	}

	result := policy.Evaluate(rules, ev, "")
	if len(result.Violations) != 0 {
		t.Errorf("expected 0 violations, got %d", len(result.Violations))
	}
	if !result.Passed {
		t.Error("all pass should result in Passed=true")
	}
}

func TestEvaluator_LintFailure(t *testing.T) {
	rules := []domain.PolicyRule{
		{ID: "lint-clean", Dimension: domain.DimCorrectness, Severity: domain.SeverityError, Metric: "lint_errors", Threshold: 0},
	}

	ev := []domain.Evidence{
		evidence.LintResult{Tool: "golangci-lint", ErrorCount: 3, WarnCount: 1}.ToEvidence(),
	}

	result := policy.Evaluate(rules, ev, "")
	if len(result.Violations) != 1 {
		t.Errorf("expected 1 violation, got %d", len(result.Violations))
	}
	if result.Passed {
		t.Error("lint failure should result in Passed=false")
	}
	if result.Violations[0].RuleID != "lint-clean" {
		t.Errorf("violation rule = %q, want lint-clean", result.Violations[0].RuleID)
	}
}

func TestEvaluator_WarningDoesntBlock(t *testing.T) {
	rules := []domain.PolicyRule{
		{ID: "todo-check", Dimension: domain.DimReadability, Severity: domain.SeverityWarning, Metric: "todo_count", Threshold: 0},
	}

	ev := []domain.Evidence{
		evidence.CodeMetrics{TodoCount: 3}.ToEvidence(),
	}

	result := policy.Evaluate(rules, ev, "")
	// Warnings generate violations but don't fail
	if len(result.Violations) != 1 {
		t.Errorf("expected 1 violation, got %d", len(result.Violations))
	}
	if !result.Passed {
		t.Error("warning-only violations should still pass")
	}
}

func TestEvaluator_MetricsBasedExtraction(t *testing.T) {
	rules := []domain.PolicyRule{
		{ID: "lint-clean", Dimension: domain.DimCorrectness, Severity: domain.SeverityError, Metric: "lint_errors", Threshold: 0},
	}
	// Evidence with Metrics only — no Details, no Summary parsing needed
	ev := []domain.Evidence{
		{
			Kind:    domain.EvidenceKindLint,
			Source:  "test",
			Passed:  true,
			Metrics: map[string]float64{"lint_errors": 0, "lint_warnings": 0},
		},
	}
	result := policy.Evaluate(rules, ev, "")
	if !result.Passed {
		t.Error("should pass with lint_errors=0 from Metrics")
	}
	if len(result.Violations) != 0 {
		t.Errorf("expected 0 violations, got %d", len(result.Violations))
	}
}

func TestEvaluator_TodoCountFromMetrics(t *testing.T) {
	rules := []domain.PolicyRule{
		{ID: "todo-check", Dimension: domain.DimReadability, Severity: domain.SeverityWarning, Metric: "todo_count", Threshold: 0},
	}
	ev := []domain.Evidence{
		{
			Kind:    domain.EvidenceKindMetrics,
			Source:  "metrics",
			Passed:  true,
			Metrics: map[string]float64{"todo_count": 3, "complexity": 5},
		},
	}
	result := policy.Evaluate(rules, ev, "")
	if len(result.Violations) != 1 {
		t.Errorf("expected 1 violation for todo_count=3 > 0, got %d", len(result.Violations))
	}
}

func TestEvaluator_CoverageFromMetrics(t *testing.T) {
	rules := []domain.PolicyRule{
		{ID: "coverage", Dimension: domain.DimTestability, Severity: domain.SeverityWarning, Metric: "test_coverage", Threshold: 0.8},
	}
	// 85% coverage — should pass
	ev := []domain.Evidence{
		{
			Kind:    domain.EvidenceKindTest,
			Source:  "go test",
			Passed:  true,
			Metrics: map[string]float64{"test_coverage": 0.85},
		},
	}
	result := policy.Evaluate(rules, ev, "")
	if len(result.Violations) != 1 {
		t.Errorf("expected 1 violation for coverage 0.85 > 0.8, got %d", len(result.Violations))
	}
}

func TestEvaluator_ComplexityFromMetrics(t *testing.T) {
	rules := []domain.PolicyRule{
		{ID: "complexity", Dimension: domain.DimMaintainability, Severity: domain.SeverityError, Metric: "complexity", Threshold: 10},
	}
	ev := []domain.Evidence{
		{
			Kind:    domain.EvidenceKindMetrics,
			Source:  "metrics",
			Passed:  true,
			Metrics: map[string]float64{"complexity": 15},
		},
	}
	result := policy.Evaluate(rules, ev, "")
	if len(result.Violations) != 1 {
		t.Errorf("expected 1 violation for complexity=15 > 10, got %d", len(result.Violations))
	}
	if result.Passed {
		t.Error("should fail with blocking complexity violation")
	}
}

func TestEvaluator_ParamCountFromStructural(t *testing.T) {
	rules := []domain.PolicyRule{
		{ID: "max-params", Dimension: domain.DimMaintainability, Severity: domain.SeverityWarning, Metric: "param_count", Threshold: 5},
	}
	ev := []domain.Evidence{
		{
			Kind:    domain.EvidenceKindStructural,
			Source:  "structural",
			Passed:  true,
			Metrics: map[string]float64{"param_count": 7},
		},
	}
	result := policy.Evaluate(rules, ev, "")
	if len(result.Violations) != 1 {
		t.Errorf("expected 1 violation for param_count=7 > 5, got %d", len(result.Violations))
	}
	if !result.Passed {
		t.Error("warning violation should not block")
	}
}

func TestEvaluator_NestingDepthFromStructural(t *testing.T) {
	rules := []domain.PolicyRule{
		{ID: "max-nesting", Dimension: domain.DimReadability, Severity: domain.SeverityWarning, Metric: "max_nesting_depth", Threshold: 4},
	}
	ev := []domain.Evidence{
		{
			Kind:    domain.EvidenceKindStructural,
			Source:  "structural",
			Passed:  true,
			Metrics: map[string]float64{"max_nesting_depth": 3},
		},
	}
	result := policy.Evaluate(rules, ev, "")
	if len(result.Violations) != 0 {
		t.Errorf("nesting=3 should not violate threshold=4, got %d violations", len(result.Violations))
	}
}

func TestEvaluator_ErrorsIgnoredFromStructural(t *testing.T) {
	rules := []domain.PolicyRule{
		{ID: "no-ignored-errors", Dimension: domain.DimCorrectness, Severity: domain.SeverityWarning, Metric: "errors_ignored", Threshold: 0},
	}
	ev := []domain.Evidence{
		{
			Kind:    domain.EvidenceKindStructural,
			Source:  "structural",
			Passed:  true,
			Metrics: map[string]float64{"errors_ignored": 2},
		},
	}
	result := policy.Evaluate(rules, ev, "")
	if len(result.Violations) != 1 {
		t.Errorf("expected 1 violation for errors_ignored=2 > 0, got %d", len(result.Violations))
	}
}

func TestEvaluator_MissingEvidence(t *testing.T) {
	rules := []domain.PolicyRule{
		{ID: "lint-clean", Dimension: domain.DimCorrectness, Severity: domain.SeverityError, Metric: "lint_errors", Threshold: 0},
	}

	// No evidence at all
	result := policy.Evaluate(rules, nil, "")
	if len(result.Violations) != 1 {
		t.Errorf("missing evidence should create violation, got %d", len(result.Violations))
	}
}

func TestEvaluate_RuleExcludePatterns(t *testing.T) {
	rules := []domain.PolicyRule{
		{
			ID:              "no-todos",
			Dimension:       domain.DimReadability,
			Severity:        domain.SeverityWarning,
			Metric:          "todo_count",
			Threshold:       0,
			ExcludePatterns: []string{"*_test.go"},
		},
	}
	ev := []domain.Evidence{
		{Kind: domain.EvidenceKindMetrics, Source: "metrics", Passed: true, Metrics: map[string]float64{"todo_count": 3}},
	}

	// Test file — rule should be skipped
	result := policy.Evaluate(rules, ev, "internal/evidence/metrics_test.go")
	if len(result.Violations) != 0 {
		t.Errorf("rule should be excluded for test files, got %d violations", len(result.Violations))
	}

	// Non-test file — rule should fire
	result = policy.Evaluate(rules, ev, "internal/evidence/metrics.go")
	if len(result.Violations) != 1 {
		t.Errorf("rule should fire for non-test files, got %d violations", len(result.Violations))
	}
}

func TestEvaluate_RulePathPatterns(t *testing.T) {
	rules := []domain.PolicyRule{
		{
			ID:           "no-os-exit",
			Dimension:    domain.DimCorrectness,
			Severity:     domain.SeverityWarning,
			Metric:       "os_exit_calls",
			Threshold:    0,
			PathPatterns: []string{"internal/*"},
		},
	}
	ev := []domain.Evidence{
		{Kind: domain.EvidenceKindStructural, Source: "structural", Passed: true, Metrics: map[string]float64{"os_exit_calls": 1}},
	}

	// Internal path — rule should fire
	result := policy.Evaluate(rules, ev, "internal/agent/providers.go")
	if len(result.Violations) != 1 {
		t.Errorf("rule should fire for internal/ path, got %d violations", len(result.Violations))
	}

	// Cmd path — rule should be skipped
	result = policy.Evaluate(rules, ev, "cmd/certify/main.go")
	if len(result.Violations) != 0 {
		t.Errorf("rule should be skipped for cmd/ path, got %d violations", len(result.Violations))
	}
}

func TestEvaluate_NoPatterns_AppliesToAll(t *testing.T) {
	rules := []domain.PolicyRule{
		{ID: "lint-clean", Dimension: domain.DimCorrectness, Severity: domain.SeverityError, Metric: "lint_errors", Threshold: 0},
	}
	ev := []domain.Evidence{
		{Kind: domain.EvidenceKindLint, Source: "lint", Passed: true, Metrics: map[string]float64{"lint_errors": 1}},
	}

	// Any path — rule should always fire when no patterns set
	for _, path := range []string{"cmd/certify/main.go", "internal/agent/providers.go", "foo_test.go"} {
		result := policy.Evaluate(rules, ev, path)
		if len(result.Violations) != 1 {
			t.Errorf("rule with no patterns should fire for %q, got %d violations", path, len(result.Violations))
		}
	}
}
