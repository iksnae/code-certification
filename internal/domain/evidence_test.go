package domain_test

import (
	"testing"

	"github.com/code-certification/certify/internal/domain"
)

func TestEvidenceKind_String(t *testing.T) {
	tests := []struct {
		ek   domain.EvidenceKind
		want string
	}{
		{domain.EvidenceKindLint, "lint"},
		{domain.EvidenceKindTypeCheck, "type_check"},
		{domain.EvidenceKindTest, "test"},
		{domain.EvidenceKindStaticAnalysis, "static_analysis"},
		{domain.EvidenceKindMetrics, "metrics"},
		{domain.EvidenceKindGitHistory, "git_history"},
		{domain.EvidenceKindAgentReview, "agent_review"},
	}
	for _, tt := range tests {
		if got := tt.ek.String(); got != tt.want {
			t.Errorf("EvidenceKind(%d).String() = %q, want %q", tt.ek, got, tt.want)
		}
	}
}

func TestSeverity_String(t *testing.T) {
	tests := []struct {
		s    domain.Severity
		want string
	}{
		{domain.SeverityInfo, "info"},
		{domain.SeverityWarning, "warning"},
		{domain.SeverityError, "error"},
		{domain.SeverityCritical, "critical"},
	}
	for _, tt := range tests {
		if got := tt.s.String(); got != tt.want {
			t.Errorf("Severity(%d).String() = %q, want %q", tt.s, got, tt.want)
		}
	}
}

func TestParseSeverity(t *testing.T) {
	tests := []struct {
		input string
		want  domain.Severity
		ok    bool
	}{
		{"info", domain.SeverityInfo, true},
		{"warning", domain.SeverityWarning, true},
		{"error", domain.SeverityError, true},
		{"critical", domain.SeverityCritical, true},
		{"unknown", 0, false},
		{"", 0, false},
	}
	for _, tt := range tests {
		got, err := domain.ParseSeverity(tt.input)
		if tt.ok {
			if err != nil {
				t.Errorf("ParseSeverity(%q) unexpected error: %v", tt.input, err)
			}
			if got != tt.want {
				t.Errorf("ParseSeverity(%q) = %v, want %v", tt.input, got, tt.want)
			}
		} else {
			if err == nil {
				t.Errorf("ParseSeverity(%q) expected error, got nil", tt.input)
			}
		}
	}
}

func TestEvidence_MissingFlag(t *testing.T) {
	e := domain.Evidence{
		Kind:    domain.EvidenceKindTest,
		Missing: true,
		Summary: "Test results not available",
	}
	if !e.Missing {
		t.Error("Evidence.Missing should be true")
	}
	if e.Passed {
		t.Error("Missing evidence should not be marked as passed")
	}
}
