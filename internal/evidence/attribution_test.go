package evidence_test

import (
	"testing"

	"github.com/iksnae/code-certification/internal/evidence"
)

func TestAttributeLintToFile_MatchingFindings(t *testing.T) {
	findings := []evidence.LintFinding{
		{File: "internal/foo.go", Line: 10, Message: "unused var", Severity: "error"},
		{File: "internal/foo.go", Line: 20, Message: "shadow", Severity: "warning"},
		{File: "internal/bar.go", Line: 5, Message: "other", Severity: "error"},
	}
	result := evidence.AttributeLintToFile(findings, "internal/foo.go")
	if result.ErrorCount != 1 {
		t.Errorf("ErrorCount = %d, want 1", result.ErrorCount)
	}
	if result.WarnCount != 1 {
		t.Errorf("WarnCount = %d, want 1", result.WarnCount)
	}
	if len(result.Findings) != 2 {
		t.Errorf("Findings = %d, want 2", len(result.Findings))
	}
}

func TestAttributeLintToFile_NoFindings(t *testing.T) {
	findings := []evidence.LintFinding{
		{File: "internal/bar.go", Line: 5, Message: "other", Severity: "error"},
	}
	result := evidence.AttributeLintToFile(findings, "internal/foo.go")
	if result.ErrorCount != 0 {
		t.Errorf("ErrorCount = %d, want 0", result.ErrorCount)
	}
	if len(result.Findings) != 0 {
		t.Errorf("Findings = %d, want 0", len(result.Findings))
	}
}

func TestAttributeLintToFile_EmptyInput(t *testing.T) {
	result := evidence.AttributeLintToFile(nil, "internal/foo.go")
	if result.ErrorCount != 0 || result.WarnCount != 0 {
		t.Error("nil findings should produce clean result")
	}
	if result.Tool != "golangci-lint:unit" {
		t.Errorf("Tool = %q, want golangci-lint:unit", result.Tool)
	}
}

func TestAttributeLintToUnit_LineRange(t *testing.T) {
	findings := []evidence.LintFinding{
		{File: "pkg/handler.go", Line: 5, Message: "outside", Severity: "error"},
		{File: "pkg/handler.go", Line: 15, Message: "inside", Severity: "error"},
		{File: "pkg/handler.go", Line: 25, Message: "inside2", Severity: "warning"},
		{File: "pkg/handler.go", Line: 35, Message: "outside2", Severity: "error"},
		{File: "other.go", Line: 15, Message: "wrong file", Severity: "error"},
	}
	result := evidence.AttributeLintToUnit(findings, "pkg/handler.go", 10, 30)
	if result.ErrorCount != 1 {
		t.Errorf("ErrorCount = %d, want 1", result.ErrorCount)
	}
	if result.WarnCount != 1 {
		t.Errorf("WarnCount = %d, want 1", result.WarnCount)
	}
	if len(result.Findings) != 2 {
		t.Errorf("Findings = %d, want 2", len(result.Findings))
	}
}

func TestAttributeLintToUnit_NoMatch(t *testing.T) {
	findings := []evidence.LintFinding{
		{File: "pkg/handler.go", Line: 50, Message: "outside", Severity: "error"},
	}
	result := evidence.AttributeLintToUnit(findings, "pkg/handler.go", 10, 30)
	if result.ErrorCount != 0 {
		t.Error("should have 0 errors for out-of-range findings")
	}
}
