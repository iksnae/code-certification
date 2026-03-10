package engine_test

import (
	"testing"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/engine"
	"github.com/iksnae/code-certification/internal/evidence"
	"github.com/iksnae/code-certification/internal/policy"
)

func TestScorer_AllClean(t *testing.T) {
	ev := []domain.Evidence{
		evidence.LintResult{Tool: "golangci-lint", ErrorCount: 0}.ToEvidence(),
		evidence.TestResult{Tool: "go test", TotalCount: 10, PassedCount: 10, Coverage: 0.9}.ToEvidence(),
		evidence.CodeMetrics{TotalLines: 100, CodeLines: 80, TodoCount: 0}.ToEvidence(),
	}

	evalResult := policy.EvaluationResult{Passed: true}
	scores := engine.Score(ev, evalResult)

	// All evidence is clean, expect good scores
	avg := scores.WeightedAverage(nil)
	if avg < 0.7 {
		t.Errorf("clean evidence average = %f, want >= 0.7", avg)
	}
}

func TestScorer_WithViolations(t *testing.T) {
	ev := []domain.Evidence{
		evidence.LintResult{Tool: "golangci-lint", ErrorCount: 5}.ToEvidence(),
	}

	evalResult := policy.EvaluationResult{
		Passed: false,
		Violations: []domain.Violation{
			{RuleID: "lint-clean", Severity: domain.SeverityError, Dimension: domain.DimCorrectness},
		},
	}

	scores := engine.Score(ev, evalResult)

	// Correctness should be penalized
	if scores[domain.DimCorrectness] >= 0.8 {
		t.Errorf("correctness with lint errors = %f, want < 0.8", scores[domain.DimCorrectness])
	}
}

func TestStatusFromScore_Certified(t *testing.T) {
	status := engine.StatusFromScore(0.85, false)
	if status != domain.StatusCertified {
		t.Errorf("score 0.85 status = %v, want certified", status)
	}
}

func TestStatusFromScore_CertifiedWithObservations(t *testing.T) {
	status := engine.StatusFromScore(0.72, false)
	if status != domain.StatusCertifiedWithObservations {
		t.Errorf("score 0.72 status = %v, want certified_with_observations", status)
	}
}

func TestStatusFromScore_Probationary(t *testing.T) {
	status := engine.StatusFromScore(0.55, false)
	if status != domain.StatusProbationary {
		t.Errorf("score 0.55 status = %v, want probationary", status)
	}
}

func TestStatusFromScore_Decertified(t *testing.T) {
	status := engine.StatusFromScore(0.30, false)
	if status != domain.StatusDecertified {
		t.Errorf("score 0.30 status = %v, want decertified", status)
	}
}

func TestStatusFromScore_HasBlockingViolation(t *testing.T) {
	// Even with high score, blocking violations force probationary
	status := engine.StatusFromScore(0.90, true)
	if status != domain.StatusProbationary {
		t.Errorf("high score with blocking = %v, want probationary", status)
	}
}

func TestScorer_ComplexityBoostsMaintainability(t *testing.T) {
	// Low complexity = high maintainability
	lowCx := evidence.CodeMetrics{TotalLines: 20, CodeLines: 15, Complexity: 2}
	ev := []domain.Evidence{lowCx.ToEvidence()}
	scores := engine.Score(ev, policy.EvaluationResult{Passed: true})
	if scores[domain.DimMaintainability] < 0.90 {
		t.Errorf("low complexity maintainability = %f, want >= 0.90", scores[domain.DimMaintainability])
	}

	// High complexity = lower maintainability
	highCx := evidence.CodeMetrics{TotalLines: 500, CodeLines: 400, Complexity: 25}
	ev2 := []domain.Evidence{highCx.ToEvidence()}
	scores2 := engine.Score(ev2, policy.EvaluationResult{Passed: true})
	if scores2[domain.DimMaintainability] >= 0.60 {
		t.Errorf("high complexity maintainability = %f, want < 0.60", scores2[domain.DimMaintainability])
	}
}

func TestScorer_SmallCodeBoostsReadability(t *testing.T) {
	small := evidence.CodeMetrics{TotalLines: 30, CodeLines: 20, Complexity: 1}
	ev := []domain.Evidence{small.ToEvidence()}
	scores := engine.Score(ev, policy.EvaluationResult{Passed: true})
	if scores[domain.DimReadability] < 0.90 {
		t.Errorf("small code readability = %f, want >= 0.90", scores[domain.DimReadability])
	}
}

func TestScorer_GitHistoryBoostsScores(t *testing.T) {
	git := evidence.GitStats{CommitCount: 15, AuthorCount: 3, AgeDays: 100}
	ev := []domain.Evidence{git.ToEvidence()}
	scores := engine.Score(ev, policy.EvaluationResult{Passed: true})
	if scores[domain.DimChangeRisk] < 0.85 {
		t.Errorf("multi-author change_risk = %f, want >= 0.85", scores[domain.DimChangeRisk])
	}
	if scores[domain.DimOperationalQuality] < 0.85 {
		t.Errorf("many commits op_quality = %f, want >= 0.85", scores[domain.DimOperationalQuality])
	}
}

func TestScorer_MetricsBasedScoring(t *testing.T) {
	// Evidence with only Metrics set — no Summary for parsing
	ev := []domain.Evidence{
		{
			Kind:    domain.EvidenceKindMetrics,
			Source:  "metrics",
			Passed:  true,
			Metrics: map[string]float64{"complexity": 3, "code_lines": 20},
		},
		{
			Kind:    domain.EvidenceKindGitHistory,
			Source:  "git",
			Passed:  true,
			Metrics: map[string]float64{"author_count": 3, "commit_count": 20},
		},
		{
			Kind:    domain.EvidenceKindTest,
			Source:  "go test",
			Passed:  true,
			Metrics: map[string]float64{"test_coverage": 0.90},
		},
	}
	scores := engine.Score(ev, policy.EvaluationResult{Passed: true})

	// Low complexity should give high maintainability
	if scores[domain.DimMaintainability] < 0.90 {
		t.Errorf("Metrics-based maintainability = %f, want >= 0.90", scores[domain.DimMaintainability])
	}
	// Small code lines should give high readability
	if scores[domain.DimReadability] < 0.90 {
		t.Errorf("Metrics-based readability = %f, want >= 0.90", scores[domain.DimReadability])
	}
	// Multi-author should boost change risk
	if scores[domain.DimChangeRisk] < 0.85 {
		t.Errorf("Metrics-based change_risk = %f, want >= 0.85", scores[domain.DimChangeRisk])
	}
	// Many commits should boost operational quality
	if scores[domain.DimOperationalQuality] < 0.85 {
		t.Errorf("Metrics-based op_quality = %f, want >= 0.85", scores[domain.DimOperationalQuality])
	}
	// High coverage should boost testability
	if scores[domain.DimTestability] < 0.90 {
		t.Errorf("Metrics-based testability = %f, want >= 0.90", scores[domain.DimTestability])
	}
}

func TestScore_StructuralDocComment(t *testing.T) {
	ev := []domain.Evidence{
		{
			Kind:   domain.EvidenceKindStructural,
			Source: "structural",
			Passed: true,
			Metrics: map[string]float64{
				"has_doc_comment": 1.0,
				"exported_name":   1.0,
				"param_count":     2,
			},
		},
	}
	scores := engine.Score(ev, policy.EvaluationResult{Passed: true})
	if scores[domain.DimReadability] < 0.90 {
		t.Errorf("documented func readability = %f, want >= 0.90", scores[domain.DimReadability])
	}
}

func TestScore_StructuralMissingDocExported(t *testing.T) {
	ev := []domain.Evidence{
		{
			Kind:   domain.EvidenceKindStructural,
			Source: "structural",
			Passed: true,
			Metrics: map[string]float64{
				"has_doc_comment": 0.0,
				"exported_name":   1.0,
				"param_count":     2,
			},
		},
	}
	scores := engine.Score(ev, policy.EvaluationResult{Passed: true})
	if scores[domain.DimReadability] > 0.75 {
		t.Errorf("undocumented exported func readability = %f, want <= 0.75", scores[domain.DimReadability])
	}
}

func TestScore_StructuralHighParamCount(t *testing.T) {
	ev := []domain.Evidence{
		{
			Kind:   domain.EvidenceKindStructural,
			Source: "structural",
			Passed: true,
			Metrics: map[string]float64{
				"param_count": 8,
			},
		},
	}
	scores := engine.Score(ev, policy.EvaluationResult{Passed: true})
	if scores[domain.DimMaintainability] > 0.60 {
		t.Errorf("8-param func maintainability = %f, want <= 0.60", scores[domain.DimMaintainability])
	}
}

func TestScore_StructuralDeepNesting(t *testing.T) {
	ev := []domain.Evidence{
		{
			Kind:   domain.EvidenceKindStructural,
			Source: "structural",
			Passed: true,
			Metrics: map[string]float64{
				"max_nesting_depth": 5,
			},
		},
	}
	scores := engine.Score(ev, policy.EvaluationResult{Passed: true})
	if scores[domain.DimReadability] > 0.75 {
		t.Errorf("deep nesting readability = %f, want <= 0.75", scores[domain.DimReadability])
	}
}

func TestScore_StructuralIgnoredErrors(t *testing.T) {
	ev := []domain.Evidence{
		{
			Kind:   domain.EvidenceKindStructural,
			Source: "structural",
			Passed: true,
			Metrics: map[string]float64{
				"errors_ignored": 2,
			},
		},
	}
	scores := engine.Score(ev, policy.EvaluationResult{Passed: true})
	if scores[domain.DimCorrectness] > 0.65 {
		t.Errorf("ignored errors correctness = %f, want <= 0.65", scores[domain.DimCorrectness])
	}
}

func TestScore_PerUnitLintOverride(t *testing.T) {
	ev := []domain.Evidence{
		// Repo-wide lint fails
		{
			Kind:    domain.EvidenceKindLint,
			Source:  "golangci-lint",
			Passed:  false,
			Metrics: map[string]float64{"lint_errors": 5},
		},
		// Per-unit lint is clean
		{
			Kind:    domain.EvidenceKindLint,
			Source:  "golangci-lint:unit",
			Passed:  true,
			Metrics: map[string]float64{"unit_lint_errors": 0, "unit_lint_warnings": 0},
		},
	}
	scores := engine.Score(ev, policy.EvaluationResult{Passed: true})
	// Per-unit clean should override repo-wide failure
	if scores[domain.DimCorrectness] < 0.85 {
		t.Errorf("per-unit clean lint correctness = %f, want >= 0.85", scores[domain.DimCorrectness])
	}
}

func TestScore_PerUnitCoverage(t *testing.T) {
	ev := []domain.Evidence{
		// Repo-wide test passes with low coverage
		{
			Kind:    domain.EvidenceKindTest,
			Source:  "go test",
			Passed:  true,
			Metrics: map[string]float64{"test_coverage": 0.50},
		},
		// Per-unit coverage is high
		{
			Kind:    domain.EvidenceKindTest,
			Source:  "coverage:unit",
			Passed:  true,
			Metrics: map[string]float64{"unit_test_coverage": 0.95},
		},
	}
	scores := engine.Score(ev, policy.EvaluationResult{Passed: true})
	if scores[domain.DimTestability] < 0.90 {
		t.Errorf("high per-unit coverage testability = %f, want >= 0.90", scores[domain.DimTestability])
	}
}

func TestScoreFromStructural_AMinusGrade(t *testing.T) {
	// A unit with perfect structural metrics should achieve A- (>=0.90).
	// This requires positive boosts for security, architectural_fitness,
	// performance_appropriateness, and operational_quality.
	ev := []domain.Evidence{
		evidence.LintResult{Tool: "golangci-lint", ErrorCount: 0}.ToEvidence(),
		evidence.TestResult{Tool: "go test", TotalCount: 10, PassedCount: 10, Coverage: 0.90}.ToEvidence(),
		evidence.CodeMetrics{TotalLines: 30, CodeLines: 20, Complexity: 3}.ToEvidence(),
		evidence.GitStats{CommitCount: 20, AuthorCount: 2, AgeDays: 60}.ToEvidence(),
		{
			Kind:   domain.EvidenceKindStructural,
			Source: "structural",
			Passed: true,
			Metrics: map[string]float64{
				"func_lines":           20,
				"errors_ignored":       0,
				"global_mutable_count": 0,
				"panic_calls":          0,
				"os_exit_calls":        0,
				"defer_in_loop":        0,
				"param_count":          2,
				"has_doc_comment":      1,
				"exported_name":        1,
				"context_not_first":    0,
				"method_count":         3,
				"max_nesting_depth":    2,
				"naked_returns":        0,
				"has_init_func":        0,
			},
		},
	}
	scores := engine.Score(ev, policy.EvaluationResult{Passed: true})
	avg := scores.WeightedAverage(nil)
	grade := domain.GradeFromScore(avg)

	if grade > domain.GradeAMinus { // Grade enum: A=0, A-=1, so > means worse
		t.Errorf("clean unit grade = %s (%.3f), want >= A-", grade, avg)
	}
	if scores[domain.DimSecurity] < 0.90 {
		t.Errorf("clean unit security = %.2f, want >= 0.90", scores[domain.DimSecurity])
	}
	if scores[domain.DimArchitecturalFitness] < 0.90 {
		t.Errorf("clean unit arch_fitness = %.2f, want >= 0.90", scores[domain.DimArchitecturalFitness])
	}
	if scores[domain.DimPerformanceAppropriateness] < 0.90 {
		t.Errorf("clean unit perf_appropriateness = %.2f, want >= 0.90", scores[domain.DimPerformanceAppropriateness])
	}
	if scores[domain.DimOperationalQuality] < 0.90 {
		t.Errorf("clean unit op_quality = %.2f, want >= 0.90", scores[domain.DimOperationalQuality])
	}
}

func TestScoreFromStructural_PenaltyOverridesBoost(t *testing.T) {
	// A unit with errors_ignored should NOT get security or operational boosts.
	ev := []domain.Evidence{
		{
			Kind:   domain.EvidenceKindStructural,
			Source: "structural",
			Passed: true,
			Metrics: map[string]float64{
				"errors_ignored":       1,
				"global_mutable_count": 0,
				"panic_calls":          0,
				"os_exit_calls":        0,
				"defer_in_loop":        0,
				"param_count":          2,
				"has_doc_comment":      1,
				"exported_name":        1,
				"context_not_first":    0,
				"method_count":         3,
				"func_lines":           20,
				"max_nesting_depth":    2,
				"naked_returns":        0,
			},
		},
	}
	scores := engine.Score(ev, policy.EvaluationResult{Passed: true})

	// Security should NOT be boosted (errors_ignored > 0)
	if scores[domain.DimSecurity] >= 0.90 {
		t.Errorf("errors_ignored security = %.2f, want < 0.90", scores[domain.DimSecurity])
	}
	// Operational quality should NOT be boosted (errors_ignored > 0)
	if scores[domain.DimOperationalQuality] >= 0.90 {
		t.Errorf("errors_ignored op_quality = %.2f, want < 0.90", scores[domain.DimOperationalQuality])
	}
	// Correctness should be penalized
	if scores[domain.DimCorrectness] > 0.65 {
		t.Errorf("errors_ignored correctness = %.2f, want <= 0.65", scores[domain.DimCorrectness])
	}
	// Architectural fitness SHOULD still be boosted (errors_ignored doesn't affect it)
	if scores[domain.DimArchitecturalFitness] < 0.90 {
		t.Errorf("clean arch with errors_ignored arch_fitness = %.2f, want >= 0.90", scores[domain.DimArchitecturalFitness])
	}
}

func TestScoreFromStructural_PartialBoosts(t *testing.T) {
	// A unit with long functions should NOT get performance boost
	// but should still get architectural boost if API is clean.
	ev := []domain.Evidence{
		{
			Kind:   domain.EvidenceKindStructural,
			Source: "structural",
			Passed: true,
			Metrics: map[string]float64{
				"func_lines":           50, // too long for perf boost
				"errors_ignored":       0,
				"global_mutable_count": 0,
				"panic_calls":          0,
				"os_exit_calls":        0,
				"defer_in_loop":        0,
				"param_count":          1,
				"has_doc_comment":      1,
				"exported_name":        1,
				"context_not_first":    0,
				"method_count":         2,
				"max_nesting_depth":    2,
				"naked_returns":        0,
			},
		},
	}
	scores := engine.Score(ev, policy.EvaluationResult{Passed: true})

	// Architectural fitness should be boosted
	if scores[domain.DimArchitecturalFitness] < 0.90 {
		t.Errorf("clean API arch_fitness = %.2f, want >= 0.90", scores[domain.DimArchitecturalFitness])
	}
	// Performance should NOT be boosted (func_lines > 30)
	if scores[domain.DimPerformanceAppropriateness] >= 0.90 {
		t.Errorf("long func perf = %.2f, want < 0.90", scores[domain.DimPerformanceAppropriateness])
	}
	// Security should be boosted (all clean)
	if scores[domain.DimSecurity] < 0.90 {
		t.Errorf("clean security = %.2f, want >= 0.90", scores[domain.DimSecurity])
	}
	// Operational quality should be boosted (all clean)
	if scores[domain.DimOperationalQuality] < 0.90 {
		t.Errorf("clean op_quality = %.2f, want >= 0.90", scores[domain.DimOperationalQuality])
	}
}

func TestScoreFromStructural_FileLevelUnit(t *testing.T) {
	// File-level units (types, vars) have func_lines: 0.
	// They should still qualify for performance boost since they aren't functions.
	ev := []domain.Evidence{
		{
			Kind:   domain.EvidenceKindStructural,
			Source: "structural",
			Passed: true,
			Metrics: map[string]float64{
				"func_lines":           0, // file-level unit, no function body
				"errors_ignored":       0,
				"global_mutable_count": 0,
				"panic_calls":          0,
				"os_exit_calls":        0,
				"defer_in_loop":        0,
				"param_count":          0,
				"has_doc_comment":      1,
				"exported_name":        1,
				"context_not_first":    0,
				"method_count":         0,
				"max_nesting_depth":    0,
				"naked_returns":        0,
			},
		},
	}
	scores := engine.Score(ev, policy.EvaluationResult{Passed: true})

	// All boosts should apply for a clean file-level unit
	if scores[domain.DimPerformanceAppropriateness] < 0.90 {
		t.Errorf("file-level perf = %.2f, want >= 0.90", scores[domain.DimPerformanceAppropriateness])
	}
	if scores[domain.DimSecurity] < 0.90 {
		t.Errorf("file-level security = %.2f, want >= 0.90", scores[domain.DimSecurity])
	}
	if scores[domain.DimArchitecturalFitness] < 0.90 {
		t.Errorf("file-level arch_fitness = %.2f, want >= 0.90", scores[domain.DimArchitecturalFitness])
	}
}

func TestScorer_RichEvidence_HighScore(t *testing.T) {
	ev := []domain.Evidence{
		evidence.LintResult{Tool: "golangci-lint", ErrorCount: 0}.ToEvidence(),
		evidence.TestResult{Tool: "go test", TotalCount: 10, PassedCount: 10, Coverage: 0.85}.ToEvidence(),
		evidence.CodeMetrics{TotalLines: 30, CodeLines: 20, Complexity: 3}.ToEvidence(),
		evidence.GitStats{CommitCount: 20, AuthorCount: 2, AgeDays: 60}.ToEvidence(),
	}
	scores := engine.Score(ev, policy.EvaluationResult{Passed: true})
	avg := scores.WeightedAverage(nil)
	if avg < 0.85 {
		t.Errorf("rich clean evidence avg = %f, want >= 0.85", avg)
	}
}
