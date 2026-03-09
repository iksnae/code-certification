package engine_test

import (
	"testing"

	"github.com/code-certification/certify/internal/domain"
	"github.com/code-certification/certify/internal/engine"
	"github.com/code-certification/certify/internal/evidence"
	"github.com/code-certification/certify/internal/policy"
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
