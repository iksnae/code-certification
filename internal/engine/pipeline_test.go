package engine_test

import (
	"testing"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/engine"
	"github.com/iksnae/code-certification/internal/evidence"
)

func TestPipeline_CertifyUnit(t *testing.T) {
	unit := domain.NewUnit(domain.NewUnitID("go", "main.go", "main"), domain.UnitTypeFunction)
	rules := []domain.PolicyRule{
		{ID: "lint-clean", Dimension: domain.DimCorrectness, Severity: domain.SeverityError, Metric: "lint_errors", Threshold: 0},
	}
	ev := []domain.Evidence{
		evidence.LintResult{Tool: "golangci-lint", ErrorCount: 0}.ToEvidence(),
	}
	cfg := domain.ExpiryConfig{
		DefaultWindowDays: 90,
		MinWindowDays:     7,
		MaxWindowDays:     365,
	}

	record := engine.CertifyUnit(unit, rules, ev, cfg, time.Now())

	if record.Status != domain.StatusCertified {
		t.Errorf("Status = %v, want certified", record.Status)
	}
	if record.UnitID.String() != "go://main.go#main" {
		t.Errorf("UnitID = %s, want go://main.go#main", record.UnitID)
	}
	if record.Score < 0.7 {
		t.Errorf("Score = %f, want >= 0.7", record.Score)
	}
	if record.Grade > domain.GradeC {
		t.Errorf("Grade = %v, want B or better", record.Grade)
	}
	if record.ExpiresAt.IsZero() {
		t.Error("ExpiresAt should not be zero")
	}
	if record.Source != "deterministic" {
		t.Errorf("Source = %q, want deterministic", record.Source)
	}
}

func TestPipeline_CertifyUnit_WithViolations(t *testing.T) {
	unit := domain.NewUnit(domain.NewUnitID("go", "bad.go", "broken"), domain.UnitTypeFunction)
	rules := []domain.PolicyRule{
		{ID: "lint-clean", Dimension: domain.DimCorrectness, Severity: domain.SeverityError, Metric: "lint_errors", Threshold: 0},
	}
	ev := []domain.Evidence{
		evidence.LintResult{Tool: "golangci-lint", ErrorCount: 5}.ToEvidence(),
	}
	cfg := domain.ExpiryConfig{
		DefaultWindowDays: 90,
		MinWindowDays:     7,
		MaxWindowDays:     365,
	}

	record := engine.CertifyUnit(unit, rules, ev, cfg, time.Now())

	if record.Status == domain.StatusCertified {
		t.Error("unit with violations should not be certified")
	}
}
