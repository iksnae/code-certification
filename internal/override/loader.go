// Package override handles loading and applying manual certification overrides.
package override

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
	"gopkg.in/yaml.v3"
)

type rawOverride struct {
	UnitID    string `yaml:"unit_id"`
	Action    string `yaml:"action"`
	Rationale string `yaml:"rationale"`
	Actor     string `yaml:"actor"`
}

type rawOverrideFile struct {
	Overrides []rawOverride `yaml:"overrides"`
}

// LoadFile loads overrides from a YAML file.
func LoadFile(path string) ([]domain.Override, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading overrides: %w", err)
	}

	var raw rawOverrideFile
	if err := yaml.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("parsing overrides: %w", err)
	}

	overrides := make([]domain.Override, 0, len(raw.Overrides))
	for _, ro := range raw.Overrides {
		id, err := domain.ParseUnitID(ro.UnitID)
		if err != nil {
			return nil, fmt.Errorf("override unit_id %q: %w", ro.UnitID, err)
		}

		action, err := parseAction(ro.Action)
		if err != nil {
			return nil, fmt.Errorf("override %q: %w", ro.UnitID, err)
		}

		o := domain.Override{
			UnitID:    id,
			Action:    action,
			Rationale: ro.Rationale,
			Actor:     ro.Actor,
			Timestamp: time.Now(),
		}

		if err := o.Validate(); err != nil {
			return nil, fmt.Errorf("override %q: %w", ro.UnitID, err)
		}

		overrides = append(overrides, o)
	}

	return overrides, nil
}

// LoadDir loads all override YAML files from a directory.
func LoadDir(dir string) ([]domain.Override, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("reading overrides directory: %w", err)
	}

	var all []domain.Override
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if !strings.HasSuffix(name, ".yml") && !strings.HasSuffix(name, ".yaml") {
			continue
		}
		overrides, err := LoadFile(filepath.Join(dir, name))
		if err != nil {
			return nil, fmt.Errorf("loading %s: %w", name, err)
		}
		all = append(all, overrides...)
	}
	return all, nil
}

func parseAction(s string) (domain.OverrideAction, error) {
	m := map[string]domain.OverrideAction{
		"exempt":         domain.OverrideExempt,
		"extend_window":  domain.OverrideExtendWindow,
		"shorten_window": domain.OverrideShortenWindow,
		"force_review":   domain.OverrideForceReview,
	}
	if a, ok := m[s]; ok {
		return a, nil
	}
	return 0, fmt.Errorf("unknown override action: %q", s)
}
