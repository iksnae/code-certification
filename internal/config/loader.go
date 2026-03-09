// Package config handles loading and validation of certification configuration.
package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/code-certification/certify/internal/domain"
	"gopkg.in/yaml.v3"
)

// rawConfig mirrors domain.Config but uses string fields for YAML unmarshaling.
type rawConfig struct {
	Mode     string                `yaml:"mode"`
	Scope    domain.ScopeConfig    `yaml:"scope"`
	Agent    domain.AgentConfig    `yaml:"agent"`
	Schedule domain.ScheduleConfig `yaml:"schedule"`
	Expiry   domain.ExpiryConfig   `yaml:"expiry"`
	Issues   domain.IssueConfig    `yaml:"issues"`
}

// LoadFile reads and parses a config YAML file.
func LoadFile(path string) (domain.Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return domain.Config{}, fmt.Errorf("reading config: %w", err)
	}
	return Load(data)
}

// LoadFromDir loads config.yml from a directory (typically .certification/).
func LoadFromDir(dir string) (domain.Config, error) {
	return LoadFile(filepath.Join(dir, "config.yml"))
}

// Load parses YAML bytes into a validated Config with defaults applied.
func Load(data []byte) (domain.Config, error) {
	cfg := domain.DefaultConfig()

	if len(data) == 0 {
		return cfg, nil
	}

	var raw rawConfig
	if err := yaml.Unmarshal(data, &raw); err != nil {
		return domain.Config{}, fmt.Errorf("parsing config YAML: %w", err)
	}

	// Map mode string
	if raw.Mode != "" {
		switch raw.Mode {
		case "advisory":
			cfg.Mode = domain.ModeAdvisory
		case "enforcing":
			cfg.Mode = domain.ModeEnforcing
		default:
			return domain.Config{}, fmt.Errorf("invalid mode: %q (must be advisory or enforcing)", raw.Mode)
		}
	}

	// Scope
	if len(raw.Scope.Include) > 0 {
		cfg.Scope.Include = raw.Scope.Include
	}
	if len(raw.Scope.Exclude) > 0 {
		cfg.Scope.Exclude = raw.Scope.Exclude
	}

	// Agent
	if raw.Agent.Enabled {
		cfg.Agent = raw.Agent
	}

	// Schedule
	cfg.Schedule = raw.Schedule

	// Expiry (apply non-zero values over defaults)
	if raw.Expiry.DefaultWindowDays > 0 {
		cfg.Expiry.DefaultWindowDays = raw.Expiry.DefaultWindowDays
	}
	if raw.Expiry.MinWindowDays > 0 {
		cfg.Expiry.MinWindowDays = raw.Expiry.MinWindowDays
	}
	if raw.Expiry.MaxWindowDays > 0 {
		cfg.Expiry.MaxWindowDays = raw.Expiry.MaxWindowDays
	}

	// Issues
	if raw.Issues.Enabled {
		cfg.Issues = raw.Issues
	}

	// Validate
	if err := validate(cfg); err != nil {
		return domain.Config{}, err
	}

	return cfg, nil
}

func validate(cfg domain.Config) error {
	if cfg.Expiry.MinWindowDays > cfg.Expiry.MaxWindowDays {
		return fmt.Errorf("expiry.min_window_days (%d) exceeds expiry.max_window_days (%d)",
			cfg.Expiry.MinWindowDays, cfg.Expiry.MaxWindowDays)
	}
	if cfg.Expiry.DefaultWindowDays < cfg.Expiry.MinWindowDays {
		return fmt.Errorf("expiry.default_window_days (%d) is below expiry.min_window_days (%d)",
			cfg.Expiry.DefaultWindowDays, cfg.Expiry.MinWindowDays)
	}
	if cfg.Expiry.DefaultWindowDays > cfg.Expiry.MaxWindowDays {
		return fmt.Errorf("expiry.default_window_days (%d) exceeds expiry.max_window_days (%d)",
			cfg.Expiry.DefaultWindowDays, cfg.Expiry.MaxWindowDays)
	}
	return nil
}
