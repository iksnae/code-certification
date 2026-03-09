package evidence_test

import (
	"testing"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/evidence"
)

func TestLintResult_Clean(t *testing.T) {
	r := evidence.LintResult{
		Tool:       "golangci-lint",
		ErrorCount: 0,
		WarnCount:  0,
	}
	ev := r.ToEvidence()
	if !ev.Passed {
		t.Error("clean lint should pass")
	}
	if ev.Kind != domain.EvidenceKindLint {
		t.Errorf("Kind = %v, want lint", ev.Kind)
	}
}

func TestLintResult_WithErrors(t *testing.T) {
	r := evidence.LintResult{
		Tool:       "golangci-lint",
		ErrorCount: 3,
		WarnCount:  1,
		Findings: []evidence.LintFinding{
			{File: "main.go", Line: 10, Message: "unused var", Severity: "error"},
			{File: "main.go", Line: 20, Message: "shadow var", Severity: "error"},
			{File: "main.go", Line: 30, Message: "line too long", Severity: "error"},
			{File: "main.go", Line: 40, Message: "minor thing", Severity: "warning"},
		},
	}
	ev := r.ToEvidence()
	if ev.Passed {
		t.Error("lint with errors should not pass")
	}
}

func TestLintResult_WarningsOnly(t *testing.T) {
	r := evidence.LintResult{
		Tool:       "eslint",
		ErrorCount: 0,
		WarnCount:  2,
	}
	ev := r.ToEvidence()
	if !ev.Passed {
		t.Error("lint with only warnings should pass")
	}
}

func TestTestResult_AllPass(t *testing.T) {
	r := evidence.TestResult{
		Tool:        "go test",
		TotalCount:  10,
		PassedCount: 10,
		FailedCount: 0,
		Coverage:    0.85,
	}
	ev := r.ToEvidence()
	if !ev.Passed {
		t.Error("all-pass test result should pass")
	}
	if ev.Kind != domain.EvidenceKindTest {
		t.Errorf("Kind = %v, want test", ev.Kind)
	}
}

func TestTestResult_WithFailures(t *testing.T) {
	r := evidence.TestResult{
		Tool:        "go test",
		TotalCount:  10,
		PassedCount: 8,
		FailedCount: 2,
		Coverage:    0.75,
	}
	ev := r.ToEvidence()
	if ev.Passed {
		t.Error("test result with failures should not pass")
	}
}
