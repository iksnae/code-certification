package domain

import (
	"fmt"
	"time"
)

// EvidenceKind identifies the type of evidence collected.
type EvidenceKind int

const (
	EvidenceKindLint           EvidenceKind = iota // Lint tool results
	EvidenceKindTypeCheck                          // Type checking results
	EvidenceKindTest                               // Test execution results
	EvidenceKindStaticAnalysis                     // Static analysis results
	EvidenceKindMetrics                            // Code metrics (complexity, size)
	EvidenceKindGitHistory                         // Git history analysis
	EvidenceKindStructural                         // AST-derived structural analysis
	EvidenceKindAgentReview                        // Agent-assisted review output
)

var evidenceKindStrings = map[EvidenceKind]string{
	EvidenceKindLint:           "lint",
	EvidenceKindTypeCheck:      "type_check",
	EvidenceKindTest:           "test",
	EvidenceKindStaticAnalysis: "static_analysis",
	EvidenceKindMetrics:        "metrics",
	EvidenceKindGitHistory:     "git_history",
	EvidenceKindStructural:     "structural",
	EvidenceKindAgentReview:    "agent_review",
}

var stringToEvidenceKind = map[string]EvidenceKind{
	"lint":            EvidenceKindLint,
	"type_check":      EvidenceKindTypeCheck,
	"test":            EvidenceKindTest,
	"static_analysis": EvidenceKindStaticAnalysis,
	"metrics":         EvidenceKindMetrics,
	"git_history":     EvidenceKindGitHistory,
	"structural":      EvidenceKindStructural,
	"agent_review":    EvidenceKindAgentReview,
}

// String returns the string representation of an EvidenceKind.
func (ek EvidenceKind) String() string {
	if s, ok := evidenceKindStrings[ek]; ok {
		return s
	}
	return fmt.Sprintf("EvidenceKind(%d)", ek)
}

// ParseEvidenceKind converts a string to an EvidenceKind.
func ParseEvidenceKind(s string) (EvidenceKind, error) {
	if ek, ok := stringToEvidenceKind[s]; ok {
		return ek, nil
	}
	return 0, fmt.Errorf("unknown evidence kind: %q", s)
}

// Evidence represents a piece of evaluation data attached to a certification record.
type Evidence struct {
	Kind       EvidenceKind       `json:"kind"`
	Source     string             `json:"source"`            // Tool or provider name
	Passed     bool               `json:"passed"`            // Whether this evidence represents a pass
	Missing    bool               `json:"missing"`           // True if evidence was expected but not collected
	Summary    string             `json:"summary"`           // Human-readable summary
	Metrics    map[string]float64 `json:"metrics,omitempty"` // Typed metrics for policy evaluation
	Details    any                `json:"details,omitempty"` // Raw or normalized data (kept for backward compat)
	Timestamp  time.Time          `json:"timestamp"`
	Confidence float64            `json:"confidence"` // 0.0–1.0, how reliable is this evidence
}

// Severity indicates how serious a policy violation is.
type Severity int

const (
	SeverityInfo     Severity = iota // Informational
	SeverityWarning                  // Minor issue
	SeverityError                    // Significant issue
	SeverityCritical                 // Must-fix issue
)

var severityStrings = map[Severity]string{
	SeverityInfo:     "info",
	SeverityWarning:  "warning",
	SeverityError:    "error",
	SeverityCritical: "critical",
}

var stringToSeverity = map[string]Severity{
	"info":     SeverityInfo,
	"warning":  SeverityWarning,
	"error":    SeverityError,
	"critical": SeverityCritical,
}

// String returns the string representation of a Severity.
func (s Severity) String() string {
	if str, ok := severityStrings[s]; ok {
		return str
	}
	return fmt.Sprintf("Severity(%d)", s)
}

// ParseSeverity converts a string to a Severity.
func ParseSeverity(s string) (Severity, error) {
	if sev, ok := stringToSeverity[s]; ok {
		return sev, nil
	}
	return 0, fmt.Errorf("unknown severity: %q", s)
}
