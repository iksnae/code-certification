package engine_test

import (
	"testing"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/engine"
	"github.com/iksnae/code-certification/internal/evidence"
)

// passingEvidence returns an evidence set that grades well for a supported
// language. Both the unsupported-language test and its control use exactly
// this set, so the only independent variable is the language.
func passingEvidence() []domain.Evidence {
	return []domain.Evidence{
		evidence.LintResult{Tool: "lint", ErrorCount: 0}.ToEvidence(),
		evidence.TestResult{
			Tool:        "test",
			TotalCount:  100,
			PassedCount: 100,
			FailedCount: 0,
			Coverage:    0.95,
		}.ToEvidence(),
	}
}

func certifyExpiryConfig() domain.ExpiryConfig {
	return domain.ExpiryConfig{
		DefaultWindowDays: 90,
		MinWindowDays:     7,
		MaxWindowDays:     365,
	}
}

// TestPipeline_CertifyUnit_UnsupportedLanguage reproduces the field failure
// reported as issue #21: a unit written in a language the engine cannot
// analyse must not receive a fabricated grade. Before the fix, Score()
// returned nil for such a unit, WeightedAverage(nil) collapsed to 0.0, and
// GradeFromScore(0.0) fell through to GradeF — producing a confident F for
// code that was never assessed at all.
func TestPipeline_CertifyUnit_UnsupportedLanguage(t *testing.T) {
	unit := domain.NewUnit(domain.NewUnitID("swift", "Sources/App/Main.swift", "main"), domain.UnitTypeFunction)

	record := engine.CertifyUnit(unit, nil, passingEvidence(), certifyExpiryConfig(), time.Now())

	if !record.Unsupported {
		t.Errorf("Unsupported = false, want true for language %q", unit.ID.Language())
	}
	if record.Grade != domain.GradeNA {
		t.Errorf("Grade = %v, want N/A — an unassessed unit must not carry a letter grade", record.Grade)
	}
	if record.Status == domain.StatusDecertified {
		t.Error("Status = decertified — an unassessed unit must not be judged as failing")
	}
	if record.Status.IsPassing() != domain.StatusExempt.IsPassing() || record.Status != domain.StatusExempt {
		t.Errorf("Status = %v, want exempt — excluded from certification, not judged", record.Status)
	}
	if record.Confidence != 0 {
		t.Errorf("Confidence = %v, want 0 — there is no assessment to be confident about", record.Confidence)
	}
	if len(record.Dimensions) != 0 {
		t.Errorf("Dimensions = %v, want empty — no dimension was scored", record.Dimensions)
	}
}

// TestPipeline_CertifyUnit_SupportedLanguage_Control is the discrimination
// control for the test above: the identical evidence set, on a language the
// engine does support, must still be graded normally. Without this, the fix
// could be inert (e.g. grading everything N/A).
func TestPipeline_CertifyUnit_SupportedLanguage_Control(t *testing.T) {
	unit := domain.NewUnit(domain.NewUnitID("go", "main.go", "main"), domain.UnitTypeFunction)

	record := engine.CertifyUnit(unit, nil, passingEvidence(), certifyExpiryConfig(), time.Now())

	if record.Unsupported {
		t.Error("Unsupported = true, want false for a supported language")
	}
	if record.Grade == domain.GradeNA {
		t.Error("Grade = N/A, want a real letter grade for a supported language")
	}
	if record.Grade > domain.GradeB {
		t.Errorf("Grade = %v, want B or better for clean lint + 95%% coverage", record.Grade)
	}
	if !record.Status.IsPassing() {
		t.Errorf("Status = %v, want a passing status for a supported language with clean evidence", record.Status)
	}
	if record.Confidence != 1.0 {
		t.Errorf("Confidence = %v, want 1.0 for a deterministically assessed unit", record.Confidence)
	}
	if len(record.Dimensions) == 0 {
		t.Error("Dimensions is empty, want scored dimensions for a supported language")
	}
}
