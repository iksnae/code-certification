package domain_test

import (
	"testing"

	"github.com/iksnae/code-certification/internal/domain"
)

func TestPolicyRule_Basics(t *testing.T) {
	rule := domain.PolicyRule{
		ID:          "max-complexity",
		Dimension:   domain.DimMaintainability,
		Description: "Cyclomatic complexity must be below 20",
		Severity:    domain.SeverityError,
		Threshold:   20.0,
		Metric:      "complexity",
	}

	if rule.ID != "max-complexity" {
		t.Errorf("rule.ID = %q, want %q", rule.ID, "max-complexity")
	}
	if rule.Dimension != domain.DimMaintainability {
		t.Errorf("rule.Dimension = %v, want %v", rule.Dimension, domain.DimMaintainability)
	}
	if rule.Severity != domain.SeverityError {
		t.Errorf("rule.Severity = %v, want %v", rule.Severity, domain.SeverityError)
	}
}

func TestPolicyPack_Basics(t *testing.T) {
	pack := domain.PolicyPack{
		Name:     "go-standard",
		Version:  "1.0.0",
		Language: "go",
		Rules: []domain.PolicyRule{
			{
				ID:        "test-coverage",
				Dimension: domain.DimTestability,
				Severity:  domain.SeverityWarning,
				Threshold: 0.8,
				Metric:    "test_coverage",
			},
		},
	}

	if pack.Name != "go-standard" {
		t.Errorf("pack.Name = %q, want %q", pack.Name, "go-standard")
	}
	if pack.Version != "1.0.0" {
		t.Errorf("pack.Version = %q, want %q", pack.Version, "1.0.0")
	}
	if pack.Language != "go" {
		t.Errorf("pack.Language = %q, want %q", pack.Language, "go")
	}
	if len(pack.Rules) != 1 {
		t.Fatalf("len(pack.Rules) = %d, want 1", len(pack.Rules))
	}
}

func TestPolicyPack_GlobalPolicy(t *testing.T) {
	pack := domain.PolicyPack{
		Name:    "global",
		Version: "1.0.0",
		// Language empty = applies to all
	}

	if pack.Language != "" {
		t.Errorf("global policy should have empty language, got %q", pack.Language)
	}
	if !pack.IsGlobal() {
		t.Error("pack.IsGlobal() should be true for empty language")
	}
}

func TestPolicyPack_PathPatterns(t *testing.T) {
	pack := domain.PolicyPack{
		Name:         "security",
		Version:      "1.0.0",
		PathPatterns: []string{"**/auth/**", "**/security/**"},
	}

	if len(pack.PathPatterns) != 2 {
		t.Fatalf("len(pack.PathPatterns) = %d, want 2", len(pack.PathPatterns))
	}
}

func TestViolation_Basics(t *testing.T) {
	v := domain.Violation{
		RuleID:      "max-complexity",
		PolicyName:  "go-standard",
		Severity:    domain.SeverityError,
		Description: "Complexity 25 exceeds threshold 20",
		Dimension:   domain.DimMaintainability,
	}

	if v.RuleID != "max-complexity" {
		t.Errorf("v.RuleID = %q, want %q", v.RuleID, "max-complexity")
	}
	if v.Severity != domain.SeverityError {
		t.Errorf("v.Severity = %v, want %v", v.Severity, domain.SeverityError)
	}
}
