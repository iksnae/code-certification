package policy_test

import (
	"testing"

	"github.com/code-certification/certify/internal/domain"
	"github.com/code-certification/certify/internal/evidence"
	"github.com/code-certification/certify/internal/policy"
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

	result := policy.Evaluate(rules, ev)
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

	result := policy.Evaluate(rules, ev)
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

	result := policy.Evaluate(rules, ev)
	// Warnings generate violations but don't fail
	if len(result.Violations) != 1 {
		t.Errorf("expected 1 violation, got %d", len(result.Violations))
	}
	if !result.Passed {
		t.Error("warning-only violations should still pass")
	}
}

func TestEvaluator_MissingEvidence(t *testing.T) {
	rules := []domain.PolicyRule{
		{ID: "lint-clean", Dimension: domain.DimCorrectness, Severity: domain.SeverityError, Metric: "lint_errors", Threshold: 0},
	}

	// No evidence at all
	result := policy.Evaluate(rules, nil)
	if len(result.Violations) != 1 {
		t.Errorf("missing evidence should create violation, got %d", len(result.Violations))
	}
}
