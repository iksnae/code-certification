package config_test

import (
	"testing"

	"github.com/iksnae/code-certification/internal/config"
	"github.com/iksnae/code-certification/internal/domain"
)

func TestValidateConfig_Valid(t *testing.T) {
	cfg := domain.DefaultConfig()
	errs := config.ValidateConfig(cfg)
	if len(errs) != 0 {
		t.Errorf("default config should be valid, got %d errors: %v", len(errs), errs)
	}
}

func TestValidateConfig_InvalidExpiry(t *testing.T) {
	cfg := domain.DefaultConfig()
	cfg.Expiry.DefaultWindowDays = 0
	cfg.Expiry.MaxWindowDays = 3
	cfg.Expiry.MinWindowDays = 10 // min > max

	errs := config.ValidateConfig(cfg)
	if len(errs) < 2 {
		t.Errorf("expected at least 2 errors, got %d: %v", len(errs), errs)
	}
}

func TestValidateConfig_AgentMissingProvider(t *testing.T) {
	cfg := domain.DefaultConfig()
	cfg.Agent.Enabled = true
	// No provider settings

	errs := config.ValidateConfig(cfg)
	if len(errs) != 1 {
		t.Errorf("expected 1 error for missing provider base_url, got %d: %v", len(errs), errs)
	}
}

func TestValidatePolicyPack_Valid(t *testing.T) {
	p := domain.PolicyPack{
		Name:    "test",
		Version: "1.0.0",
		Rules: []domain.PolicyRule{
			{ID: "r1", Metric: "lint_errors", Severity: domain.SeverityError, Dimension: domain.DimCorrectness},
		},
	}
	errs := config.ValidatePolicyPack(p)
	if len(errs) != 0 {
		t.Errorf("valid pack should have no errors, got %d: %v", len(errs), errs)
	}
}

func TestValidatePolicyPack_MissingFields(t *testing.T) {
	p := domain.PolicyPack{} // no name, version, rules
	errs := config.ValidatePolicyPack(p)
	if len(errs) < 2 {
		t.Errorf("expected at least 2 errors, got %d: %v", len(errs), errs)
	}
}

func TestValidatePolicyPack_InvalidRuleMetric(t *testing.T) {
	p := domain.PolicyPack{
		Name:    "test",
		Version: "1.0.0",
		Rules: []domain.PolicyRule{
			{ID: "", Metric: ""},
		},
	}
	errs := config.ValidatePolicyPack(p)
	hasID := false
	hasMetric := false
	for _, e := range errs {
		if e.Field == "rules[0].id" {
			hasID = true
		}
		if e.Field == "rules[0].metric" {
			hasMetric = true
		}
	}
	if !hasID || !hasMetric {
		t.Errorf("expected rule id and metric errors, got %v", errs)
	}
}
