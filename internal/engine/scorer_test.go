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
	scores := engine.Score(ev, evalResult, "go")

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
	scores := engine.Score(ev, policy.EvaluationResult{Passed: true}, "go")

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
	cleanScores := engine.Score(cleanEv, policy.EvaluationResult{Passed: true}, "go")
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
	badScores := engine.Score(badEv, policy.EvaluationResult{Passed: true}, "go")
	if v, ok := badScores[domain.DimArchitecturalFitness]; !ok || v >= 0.80 {
		t.Errorf("architectural_fitness with god object = %v (present=%v), want present and < 0.80", v, ok)
	}
	// defer_in_loop should cap performance_appropriateness
	if v, ok := badScores[domain.DimPerformanceAppropriateness]; !ok || v >= 0.80 {
		t.Errorf("performance_appropriateness with defer_in_loop = %v (present=%v), want present and < 0.80", v, ok)
	}
}

func TestScorer_NoEvidenceNoScore(t *testing.T) {
	scores := engine.Score(nil, policy.EvaluationResult{Passed: true}, "go")
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
	cleanScores := engine.Score(cleanGlobals, policy.EvaluationResult{Passed: true}, "go")
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
	dirtyScores := engine.Score(dirtyGlobals, policy.EvaluationResult{Passed: true}, "go")
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

	scores := engine.Score(ev, evalResult, "go")

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
	scores := engine.Score(ev, policy.EvaluationResult{Passed: true}, "go")
	if scores[domain.DimMaintainability] < 0.90 {
		t.Errorf("low complexity maintainability = %f, want >= 0.90", scores[domain.DimMaintainability])
	}

	// High complexity = lower maintainability
	highCx := evidence.CodeMetrics{TotalLines: 500, CodeLines: 400, Complexity: 25}
	ev2 := []domain.Evidence{highCx.ToEvidence()}
	scores2 := engine.Score(ev2, policy.EvaluationResult{Passed: true}, "go")
	if scores2[domain.DimMaintainability] >= 0.60 {
		t.Errorf("high complexity maintainability = %f, want < 0.60", scores2[domain.DimMaintainability])
	}
}

func TestScorer_SmallCodeBoostsReadability(t *testing.T) {
	small := evidence.CodeMetrics{TotalLines: 30, CodeLines: 20, Complexity: 1}
	ev := []domain.Evidence{small.ToEvidence()}
	scores := engine.Score(ev, policy.EvaluationResult{Passed: true}, "go")
	if scores[domain.DimReadability] < 0.90 {
		t.Errorf("small code readability = %f, want >= 0.90", scores[domain.DimReadability])
	}
}

func TestScorer_GitHistoryBoostsScores(t *testing.T) {
	git := evidence.GitStats{CommitCount: 15, AuthorCount: 3, AgeDays: 100}
	ev := []domain.Evidence{git.ToEvidence()}
	scores := engine.Score(ev, policy.EvaluationResult{Passed: true}, "go")
	if scores[domain.DimChangeRisk] < 0.85 {
		t.Errorf("multi-author change_risk = %f, want >= 0.85", scores[domain.DimChangeRisk])
	}
	if scores[domain.DimOperationalQuality] < 0.85 {
		t.Errorf("many commits op_quality = %f, want >= 0.85", scores[domain.DimOperationalQuality])
	}
}

func TestScorer_ViolationsReduceScores(t *testing.T) {
	// violations against a measured dimension should reduce it
	ev := []domain.Evidence{
		evidence.LintResult{Tool: "golangci-lint", ErrorCount: 0}.ToEvidence(),
	}
	evalResult := policy.EvaluationResult{
		Passed: true,
		Violations: []domain.Violation{
			{RuleID: "some-rule", Severity: domain.SeverityCritical, Dimension: domain.DimCorrectness},
		},
	}
	scores := engine.Score(ev, evalResult, "go")
	// lint clean gives 0.95; critical penalty is 0.5; so ~0.45
	if scores[domain.DimCorrectness] > 0.50 {
		t.Errorf("correctness with critical violation = %f, want <= 0.50", scores[domain.DimCorrectness])
	}
}

func TestScorer_ViolationDoesNotInjectDimension(t *testing.T) {
	// a violation against a dimension with no evidence should NOT inject it
	ev := []domain.Evidence{
		evidence.LintResult{Tool: "golangci-lint", ErrorCount: 0}.ToEvidence(),
	}
	evalResult := policy.EvaluationResult{
		Passed: true,
		Violations: []domain.Violation{
			{RuleID: "security-rule", Severity: domain.SeverityError, Dimension: domain.DimSecurity},
		},
	}
	scores := engine.Score(ev, evalResult, "go")
	// Security was not measured by evidence, so it should not appear
	if _, ok := scores[domain.DimSecurity]; ok {
		t.Error("violation should not inject a dimension with no evidence")
	}
}

func TestScorer_TierGenericReturnsNil(t *testing.T) {
	// An unsupported language should return nil
	scores := engine.Score(nil, policy.EvaluationResult{Passed: true}, "brainfuck")
	if scores != nil {
		t.Error("TierGeneric language should return nil scores")
	}
}

func TestScorer_TierGenericComplexitySkipsMaintainability(t *testing.T) {
	// For a TierGeneric language with complexity=0, maintainability should not be set
	ev := []domain.Evidence{{
		Kind:    domain.EvidenceKindMetrics,
		Source:  "metrics",
		Passed:  true,
		Metrics: map[string]float64{"complexity": 0},
	}}
	// "python" is not a registered language → TierGeneric
	scores := engine.Score(ev, policy.EvaluationResult{Passed: true}, "python")
	if _, ok := scores[domain.DimMaintainability]; ok {
		t.Error("TierGeneric with complexity=0 should not set maintainability")
	}
}

func TestScorer_TierFullComplexityDoesSetMaintainability(t *testing.T) {
	// For a TierFull language with complexity=0, maintainability should be set
	ev := []domain.Evidence{{
		Kind:    domain.EvidenceKindMetrics,
		Source:  "metrics",
		Passed:  true,
		Metrics: map[string]float64{"complexity": 0},
	}}
	// "go" is a registered language → TierFull
	scores := engine.Score(ev, policy.EvaluationResult{Passed: true}, "go")
	if v, ok := scores[domain.DimMaintainability]; !ok || v < 0.90 {
		t.Errorf("TierFull with complexity=0 maintainability = %v (present=%v), want present and >= 0.90", v, ok)
	}
}

func TestScorer_Structural_CognitiveComplexity(t *testing.T) {
	tests := []struct {
		name          string
		cogComplexity float64
		wantMinRead   float64
		wantMaxRead   float64
	}{
		{"low complexity", 3, 0.90, 1.0},
		{"medium complexity", 12, 0.80, 0.90},
		{"very high complexity", 30, 0.0, 0.55},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ev := []domain.Evidence{{
				Kind: domain.EvidenceKindStructural, Source: "structural", Passed: true,
				Metrics: map[string]float64{
					"cognitive_complexity": tt.cogComplexity,
				},
			}}
			scores := engine.Score(ev, policy.EvaluationResult{Passed: true}, "go")
			read := scores[domain.DimReadability]
			if read < tt.wantMinRead || read > tt.wantMaxRead {
				t.Errorf("readability = %f, want [%f, %f]", read, tt.wantMinRead, tt.wantMaxRead)
			}
		})
	}
}

func TestScorer_ErrorsNotWrapped(t *testing.T) {
	tests := []struct {
		name    string
		count   float64
		wantMin float64
		wantMax float64
	}{
		{"all wrapped", 0, 0.85, 1.0},
		{"few unwrapped", 2, 0.70, 0.80},
		{"many unwrapped", 5, 0.0, 0.60},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ev := []domain.Evidence{{
				Kind: domain.EvidenceKindStructural, Source: "structural", Passed: true,
				Metrics: map[string]float64{
					"errors_not_wrapped": tt.count,
				},
			}}
			scores := engine.Score(ev, policy.EvaluationResult{Passed: true}, "go")
			opQ := scores[domain.DimOperationalQuality]
			if opQ < tt.wantMin || opQ > tt.wantMax {
				t.Errorf("operational_quality = %f, want [%f, %f]", opQ, tt.wantMin, tt.wantMax)
			}
		})
	}
}

func TestScorer_UnsafeImports(t *testing.T) {
	ev := []domain.Evidence{{
		Kind: domain.EvidenceKindStructural, Source: "structural", Passed: true,
		Metrics: map[string]float64{
			"unsafe_import_count": 2,
		},
	}}
	scores := engine.Score(ev, policy.EvaluationResult{Passed: true}, "go")
	sec := scores[domain.DimSecurity]
	if sec > 0.65 {
		t.Errorf("security = %f, want <= 0.65 for unsafe imports", sec)
	}
}

func TestScorer_HardcodedSecrets(t *testing.T) {
	ev := []domain.Evidence{{
		Kind: domain.EvidenceKindStructural, Source: "structural", Passed: true,
		Metrics: map[string]float64{
			"hardcoded_secrets": 1,
		},
	}}
	scores := engine.Score(ev, policy.EvaluationResult{Passed: true}, "go")
	sec := scores[domain.DimSecurity]
	if sec > 0.35 {
		t.Errorf("security = %f, want <= 0.35 for hardcoded secrets", sec)
	}
}

func TestScorer_EmptyCatchBlocks(t *testing.T) {
	ev := []domain.Evidence{{
		Kind: domain.EvidenceKindStructural, Source: "structural", Passed: true,
		Metrics: map[string]float64{
			"empty_catch_blocks": 2,
		},
	}}
	scores := engine.Score(ev, policy.EvaluationResult{Passed: true}, "go")
	corr := scores[domain.DimCorrectness]
	if corr > 0.60 {
		t.Errorf("correctness = %f, want <= 0.60 for empty catch blocks", corr)
	}
}

func TestScorer_QuadraticPatterns(t *testing.T) {
	ev := []domain.Evidence{{
		Kind: domain.EvidenceKindStructural, Source: "structural", Passed: true,
		Metrics: map[string]float64{
			"quadratic_patterns": 1,
			"loop_nesting_depth": 0,
		},
	}}
	scores := engine.Score(ev, policy.EvaluationResult{Passed: true}, "go")
	perf := scores[domain.DimPerformanceAppropriateness]
	if perf > 0.50 {
		t.Errorf("perf = %f, want <= 0.50 for quadratic patterns", perf)
	}
}

func TestScorer_HighReturnCount(t *testing.T) {
	ev := []domain.Evidence{{
		Kind: domain.EvidenceKindStructural, Source: "structural", Passed: true,
		Metrics: map[string]float64{
			"return_count": 8,
		},
	}}
	scores := engine.Score(ev, policy.EvaluationResult{Passed: true}, "go")
	maint := scores[domain.DimMaintainability]
	if maint > 0.70 {
		t.Errorf("maintainability = %f, want <= 0.70 for high return count", maint)
	}
}

func TestScorer_NestedLoopPairs(t *testing.T) {
	ev := []domain.Evidence{{
		Kind: domain.EvidenceKindStructural, Source: "structural", Passed: true,
		Metrics: map[string]float64{
			"nested_loop_pairs":  2,
			"quadratic_patterns": 0,
			"loop_nesting_depth": 0,
		},
	}}
	scores := engine.Score(ev, policy.EvaluationResult{Passed: true}, "go")
	perf := scores[domain.DimPerformanceAppropriateness]
	if perf > 0.65 {
		t.Errorf("perf = %f, want <= 0.65 for nested loop pairs", perf)
	}
}

func TestScorer_FanIn_LowIsGood(t *testing.T) {
	ev := []domain.Evidence{{
		Kind: domain.EvidenceKindStructural, Source: "structural", Passed: true,
		Metrics: map[string]float64{"fan_in": 3},
	}}
	scores := engine.Score(ev, policy.EvaluationResult{Passed: true}, "go")
	cr := scores[domain.DimChangeRisk]
	if cr < 0.90 {
		t.Errorf("change_risk = %f, want >= 0.90 for fan_in=3", cr)
	}
}

func TestScorer_FanIn_HighIsBad(t *testing.T) {
	ev := []domain.Evidence{{
		Kind: domain.EvidenceKindStructural, Source: "structural", Passed: true,
		Metrics: map[string]float64{"fan_in": 25},
	}}
	scores := engine.Score(ev, policy.EvaluationResult{Passed: true}, "go")
	cr := scores[domain.DimChangeRisk]
	if cr > 0.55 {
		t.Errorf("change_risk = %f, want <= 0.55 for fan_in=25", cr)
	}
}

func TestScorer_FanOut_LowIsGood(t *testing.T) {
	ev := []domain.Evidence{{
		Kind: domain.EvidenceKindStructural, Source: "structural", Passed: true,
		Metrics: map[string]float64{"fan_out": 3},
	}}
	scores := engine.Score(ev, policy.EvaluationResult{Passed: true}, "go")
	maint := scores[domain.DimMaintainability]
	if maint < 0.90 {
		t.Errorf("maintainability = %f, want >= 0.90 for fan_out=3", maint)
	}
}

func TestScorer_FanOut_HighIsBad(t *testing.T) {
	ev := []domain.Evidence{{
		Kind: domain.EvidenceKindStructural, Source: "structural", Passed: true,
		Metrics: map[string]float64{"fan_out": 20},
	}}
	scores := engine.Score(ev, policy.EvaluationResult{Passed: true}, "go")
	maint := scores[domain.DimMaintainability]
	if maint > 0.60 {
		t.Errorf("maintainability = %f, want <= 0.60 for fan_out=20", maint)
	}
}

func TestScorer_DeadCode(t *testing.T) {
	ev := []domain.Evidence{{
		Kind: domain.EvidenceKindStructural, Source: "structural", Passed: true,
		Metrics: map[string]float64{"is_dead_code": 1},
	}}
	scores := engine.Score(ev, policy.EvaluationResult{Passed: true}, "go")
	maint := scores[domain.DimMaintainability]
	if maint > 0.65 {
		t.Errorf("maintainability = %f, want <= 0.65 for dead code", maint)
	}
}

func TestScorer_DepDepth_ShallowIsGood(t *testing.T) {
	ev := []domain.Evidence{{
		Kind: domain.EvidenceKindStructural, Source: "structural", Passed: true,
		Metrics: map[string]float64{"dep_depth": 2},
	}}
	scores := engine.Score(ev, policy.EvaluationResult{Passed: true}, "go")
	af := scores[domain.DimArchitecturalFitness]
	if af < 0.90 {
		t.Errorf("arch_fitness = %f, want >= 0.90 for dep_depth=2", af)
	}
}

func TestScorer_DepDepth_DeepIsBad(t *testing.T) {
	ev := []domain.Evidence{{
		Kind: domain.EvidenceKindStructural, Source: "structural", Passed: true,
		Metrics: map[string]float64{"dep_depth": 10},
	}}
	scores := engine.Score(ev, policy.EvaluationResult{Passed: true}, "go")
	af := scores[domain.DimArchitecturalFitness]
	if af > 0.60 {
		t.Errorf("arch_fitness = %f, want <= 0.60 for dep_depth=10", af)
	}
}

func TestScorer_ConcreteDeps(t *testing.T) {
	ev := []domain.Evidence{{
		Kind: domain.EvidenceKindStructural, Source: "structural", Passed: true,
		Metrics: map[string]float64{"concrete_deps": 2},
	}}
	scores := engine.Score(ev, policy.EvaluationResult{Passed: true}, "go")
	test := scores[domain.DimTestability]
	if test > 0.70 {
		t.Errorf("testability = %f, want <= 0.70 for concrete_deps=2", test)
	}
	af := scores[domain.DimArchitecturalFitness]
	if af > 0.70 {
		t.Errorf("arch_fitness = %f, want <= 0.70 for concrete_deps=2", af)
	}
}

func TestScorer_DeepAnalysis_AllClean(t *testing.T) {
	// A function with all deep analysis metrics clean should score well
	ev := []domain.Evidence{{
		Kind: domain.EvidenceKindStructural, Source: "structural", Passed: true,
		Metrics: map[string]float64{
			"has_doc_comment":      1,
			"param_count":          2,
			"max_nesting_depth":    1,
			"func_lines":           20,
			"cognitive_complexity": 3,
			"errors_not_wrapped":   0,
			"unsafe_import_count":  0,
			"hardcoded_secrets":    0,
			"empty_catch_blocks":   0,
			"loop_nesting_depth":   0,
			"quadratic_patterns":   0,
			"return_count":         1,
			"exported_name":        1,
		},
	}}
	scores := engine.Score(ev, policy.EvaluationResult{Passed: true}, "go")

	for dim, score := range scores {
		if score < 0.80 {
			t.Errorf("dim %s = %f, want >= 0.80 for clean code", dim, score)
		}
	}
}
