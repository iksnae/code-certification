package config

import (
	"fmt"
	"strings"

	"github.com/code-certification/certify/internal/domain"
)

// ValidationError describes a single validation issue.
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// ValidateConfig checks a Config for structural issues.
func ValidateConfig(cfg domain.Config) []ValidationError {
	var errs []ValidationError

	if cfg.Expiry.DefaultWindowDays <= 0 {
		errs = append(errs, ValidationError{"expiry.default_window_days", "must be positive"})
	}
	if cfg.Expiry.MinWindowDays < 0 {
		errs = append(errs, ValidationError{"expiry.min_window_days", "must be non-negative"})
	}
	if cfg.Expiry.MaxWindowDays < cfg.Expiry.MinWindowDays {
		errs = append(errs, ValidationError{"expiry.max_window_days", "must be >= min_window_days"})
	}

	if cfg.Agent.Enabled {
		if cfg.Agent.Provider.APIKeyEnv == "" {
			errs = append(errs, ValidationError{"agent.provider.api_key_env", "required when agent is enabled"})
		}
		if cfg.Agent.Provider.BaseURL == "" {
			errs = append(errs, ValidationError{"agent.provider.base_url", "required when agent is enabled"})
		}
	}

	return errs
}

// ValidatePolicyPack checks a PolicyPack for structural issues.
func ValidatePolicyPack(p domain.PolicyPack) []ValidationError {
	var errs []ValidationError

	if p.Name == "" {
		errs = append(errs, ValidationError{"name", "required"})
	}
	if p.Version == "" {
		errs = append(errs, ValidationError{"version", "required"})
	}

	validSeverities := map[string]bool{
		"info": true, "warning": true, "error": true, "critical": true,
	}
	validDimensions := map[string]bool{
		"correctness": true, "maintainability": true, "readability": true,
		"testability": true, "security": true, "architectural_fitness": true,
		"operational_quality": true, "performance": true, "change_risk": true,
	}

	for i, rule := range p.Rules {
		prefix := fmt.Sprintf("rules[%d]", i)
		if rule.ID == "" {
			errs = append(errs, ValidationError{prefix + ".id", "required"})
		}
		if rule.Metric == "" {
			errs = append(errs, ValidationError{prefix + ".metric", "required"})
		}
		if !validSeverities[strings.ToLower(rule.Severity.String())] {
			errs = append(errs, ValidationError{prefix + ".severity", fmt.Sprintf("invalid: %q", rule.Severity)})
		}
		dimStr := rule.Dimension.String()
		if dimStr != "" && dimStr != "Dimension(0)" && !validDimensions[dimStr] {
			errs = append(errs, ValidationError{prefix + ".dimension", fmt.Sprintf("unknown: %q", dimStr)})
		}
	}

	return errs
}
