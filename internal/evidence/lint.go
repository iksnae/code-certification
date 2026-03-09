package evidence

import (
	"fmt"
	"time"

	"github.com/code-certification/certify/internal/domain"
)

// LintFinding represents a single lint finding.
type LintFinding struct {
	File     string `json:"file"`
	Line     int    `json:"line"`
	Message  string `json:"message"`
	Severity string `json:"severity"`
	Rule     string `json:"rule,omitempty"`
}

// LintResult holds aggregated lint results.
type LintResult struct {
	Tool       string        `json:"tool"`
	ErrorCount int           `json:"error_count"`
	WarnCount  int           `json:"warn_count"`
	Findings   []LintFinding `json:"findings,omitempty"`
}

// ToEvidence converts LintResult to a domain.Evidence.
func (r LintResult) ToEvidence() domain.Evidence {
	return domain.Evidence{
		Kind:       domain.EvidenceKindLint,
		Source:     r.Tool,
		Passed:     r.ErrorCount == 0,
		Summary:    fmt.Sprintf("%s: %d errors, %d warnings", r.Tool, r.ErrorCount, r.WarnCount),
		Details:    r,
		Timestamp:  time.Now(),
		Confidence: 1.0,
	}
}

// TestResult holds aggregated test execution results.
type TestResult struct {
	Tool        string  `json:"tool"`
	TotalCount  int     `json:"total_count"`
	PassedCount int     `json:"passed_count"`
	FailedCount int     `json:"failed_count"`
	SkipCount   int     `json:"skip_count"`
	Coverage    float64 `json:"coverage"` // 0.0–1.0
}

// ToEvidence converts TestResult to a domain.Evidence.
func (r TestResult) ToEvidence() domain.Evidence {
	return domain.Evidence{
		Kind:       domain.EvidenceKindTest,
		Source:     r.Tool,
		Passed:     r.FailedCount == 0,
		Summary:    fmt.Sprintf("%s: %d/%d passed (%.0f%% coverage)", r.Tool, r.PassedCount, r.TotalCount, r.Coverage*100),
		Details:    r,
		Timestamp:  time.Now(),
		Confidence: 1.0,
	}
}
