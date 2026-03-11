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

func TestScorer_OnlyMeasuredDimensionsPresent(t *testing.T) {
	// With only lint evidence, only correctness should be scored
	ev := []domain.Evidence{
		evidence.LintResult{Tool: "golangci-lint", ErrorCount: 0}.ToEvidence(),
	}
	scores := engine.Score(ev, policy.EvaluationResult{Passed: true})

	if _, ok := scores[domain.DimCorrectness]; !ok {
		t.Error("lint evidence should set correctness")
	}
	// Dimensions with no evidence should be absent
	for _, dim := range []domain.Dimension{
		domain.DimTestability,
		domain.DimArchitecturalFitness,
		domain.DimPerformanceAppropriateness,
	} {
		if _, ok := scores[dim]; ok {
			t.Errorf("dimension %s should not be set with only lint evidence", dim)
		}
	}
}

func TestScorer_PenaltyOnlyDimsAppearWhenBad(t *testing.T) {
	// architectural_fitness should only appear when there's a violation
	cleanEv := []domain.Evidence{
		{
			Kind:   domain.EvidenceKindStructural,
			Source: "structural",
			Passed: true,
			Metrics: map[string]float64{
				"method_count":      5,
				"context_not_first": 0,
			},
		},
	}
	cleanScores := engine.Score(cleanEv, policy.EvaluationResult{Passed: true})
	if _, ok := cleanScores[domain.DimArchitecturalFitness]; ok {
		t.Error("architectural_fitness should not be set when no violations found")
	}
	// performance_appropriateness IS set for clean code (algo complexity = O(1) → 0.95)
	// because algo complexity is always measured, unlike arch fitness which is penalty-only
	if v, ok := cleanScores[domain.DimPerformanceAppropriateness]; !ok || v < 0.90 {
		t.Errorf("performance_appropriateness for clean code = %v (present=%v), want present and >= 0.90", v, ok)
	}

	// When violations exist, penalty dims should appear with low scores
	badEv := []domain.Evidence{
		{
			Kind:   domain.EvidenceKindStructural,
			Source: "structural",
			Passed: true,
			Metrics: map[string]float64{
				"method_count":      20,
				"context_not_first": 1,
				"defer_in_loop":     1,
			},
		},
	}
	badScores := engine.Score(badEv, policy.EvaluationResult{Passed: true})
	if v, ok := badScores[domain.DimArchitecturalFitness]; !ok || v >= 0.80 {
		t.Errorf("architectural_fitness with god object = %v (present=%v), want present and < 0.80", v, ok)
	}
	// defer_in_loop should cap performance_appropriateness
	if v, ok := badScores[domain.DimPerformanceAppropriateness]; !ok || v >= 0.80 {
		t.Errorf("performance_appropriateness with defer_in_loop = %v (present=%v), want present and < 0.80", v, ok)
	}
}

func TestScorer_NoEvidenceNoScore(t *testing.T) {
	scores := engine.Score(nil, policy.EvaluationResult{Passed: true})
	if len(scores) != 0 {
		t.Errorf("no evidence should produce empty scores, got %d dimensions", len(scores))
	}
	avg := scores.WeightedAverage(nil)
	if avg != 0 {
		t.Errorf("no evidence average = %f, want 0", avg)
	}
}

func TestScorer_SecurityOnlyWhenMeasured(t *testing.T) {
	// Security should appear when structural evidence checks global state
	cleanGlobals := []domain.Evidence{
		{
			Kind:   domain.EvidenceKindStructural,
			Source: "structural",
			Passed: true,
			Metrics: map[string]float64{
				"global_mutable_count": 0,
			},
		},
	}
	cleanScores := engine.Score(cleanGlobals, policy.EvaluationResult{Passed: true})
	if v, ok := cleanScores[domain.DimSecurity]; !ok || v < 0.85 {
		t.Errorf("clean globals security = %v (present=%v), want present and >= 0.85", v, ok)
	}

	// With globals, security should be penalized
	dirtyGlobals := []domain.Evidence{
		{
			Kind:   domain.EvidenceKindStructural,
			Source: "structural",
			Passed: true,
			Metrics: map[string]float64{
				"global_mutable_count": 5,
			},
		},
	}
	dirtyScores := engine.Score(dirtyGlobals, policy.EvaluationResult{Passed: true})
	if v, ok := dirtyScores[domain.DimSecurity]; !ok || v >= 0.70 {
		t.Errorf("5 globals security = %v (present=%v), want present and < 0.70", v, ok)
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

func TestScorer_GraduatedGitHistory(t *testing.T) {
	tests := []struct {
		name       string
		commits    int
		authors    int
		wantOpQual float64
		wantCR     float64
	}{
		{"high commits + many authors", 60, 3, 0.95, 0.95},
		{"medium commits + 2 authors", 25, 2, 0.90, 0.90},
		{"low commits + 1 author", 12, 1, 0.85, 0.70},
		{"few commits + 1 author", 3, 1, 0.75, 0.70},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			git := evidence.GitStats{CommitCount: tt.commits, AuthorCount: tt.authors, AgeDays: 30}
			ev := []domain.Evidence{git.ToEvidence()}
			scores := engine.Score(ev, policy.EvaluationResult{Passed: true})
			if scores[domain.DimOperationalQuality] < tt.wantOpQual {
				t.Errorf("op_quality = %f, want >= %f (commits=%d)", scores[domain.DimOperationalQuality], tt.wantOpQual, tt.commits)
			}
			if scores[domain.DimChangeRisk] < tt.wantCR {
				t.Errorf("change_risk = %f, want >= %f (authors=%d)", scores[domain.DimChangeRisk], tt.wantCR, tt.authors)
			}
		})
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

func TestScorer_FileCodeLinesThresholds(t *testing.T) {
	// code_lines from metrics evidence represents file-level line count.
	// File-level thresholds should be more generous than function-level:
	// a 400-line file is normal; a 400-line function is not.
	tests := []struct {
		name        string
		codeLines   int
		wantMinRead float64
	}{
		{"small file (50 lines)", 50, 0.95},
		{"medium file (200 lines)", 200, 0.90},
		{"large file (400 lines)", 400, 0.85},
		{"very large file (700 lines)", 700, 0.75},
		{"huge file (1200 lines)", 1200, 0.60},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ev := []domain.Evidence{
				{
					Kind:    domain.EvidenceKindMetrics,
					Source:  "metrics",
					Passed:  true,
					Metrics: map[string]float64{"code_lines": float64(tt.codeLines)},
				},
			}
			scores := engine.Score(ev, policy.EvaluationResult{Passed: true})
			if scores[domain.DimReadability] < tt.wantMinRead {
				t.Errorf("code_lines=%d → readability=%f, want >= %f", tt.codeLines, scores[domain.DimReadability], tt.wantMinRead)
			}
		})
	}
}

func TestScorer_ComplexityGraduated(t *testing.T) {
	// Complexity 11-15 is moderate, not as bad as 16-20.
	tests := []struct {
		name       string
		complexity int
		wantMin    float64
	}{
		{"low complexity (5)", 5, 0.95},
		{"moderate complexity (12)", 12, 0.80},
		{"high complexity (18)", 18, 0.70},
		{"very high complexity (25)", 25, 0.0}, // setMin caps it
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ev := []domain.Evidence{
				{
					Kind:    domain.EvidenceKindMetrics,
					Source:  "metrics",
					Passed:  true,
					Metrics: map[string]float64{"complexity": float64(tt.complexity)},
				},
			}
			scores := engine.Score(ev, policy.EvaluationResult{Passed: true})
			if scores[domain.DimMaintainability] < tt.wantMin {
				t.Errorf("complexity=%d → maintainability=%f, want >= %f", tt.complexity, scores[domain.DimMaintainability], tt.wantMin)
			}
		})
	}
}

func TestScorer_AlgoComplexityScoring(t *testing.T) {
	tests := []struct {
		name        string
		loopDepth   int
		recursive   int
		wantMinPerf float64
		wantMaxPerf float64
	}{
		{"O(1) no loops", 0, 0, 0.95, 1.0},
		{"O(n) single loop", 1, 0, 0.90, 0.95},
		{"O(n²) nested loops", 2, 0, 0.60, 0.75},
		{"O(n³) triple nested", 3, 0, 0.0, 0.55},
		{"O(2^n) recursive", 0, 1, 0.0, 0.45},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ev := []domain.Evidence{
				{
					Kind:   domain.EvidenceKindStructural,
					Source: "structural",
					Passed: true,
					Metrics: map[string]float64{
						"loop_nesting_depth": float64(tt.loopDepth),
						"recursive_calls":    float64(tt.recursive),
					},
				},
			}
			scores := engine.Score(ev, policy.EvaluationResult{Passed: true})
			perf, ok := scores[domain.DimPerformanceAppropriateness]
			if !ok {
				t.Fatal("performance_appropriateness should be set when algo complexity is measured")
			}
			if perf < tt.wantMinPerf {
				t.Errorf("perf = %f, want >= %f", perf, tt.wantMinPerf)
			}
			if perf > tt.wantMaxPerf {
				t.Errorf("perf = %f, want <= %f", perf, tt.wantMaxPerf)
			}
		})
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
