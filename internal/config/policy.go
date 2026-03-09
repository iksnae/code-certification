package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/code-certification/certify/internal/domain"
	"gopkg.in/yaml.v3"
)

// rawPolicyRule is the YAML representation of a policy rule.
type rawPolicyRule struct {
	ID          string  `yaml:"id"`
	Dimension   string  `yaml:"dimension"`
	Description string  `yaml:"description"`
	Severity    string  `yaml:"severity"`
	Threshold   float64 `yaml:"threshold"`
	Metric      string  `yaml:"metric"`
}

// rawPolicyPack is the YAML representation of a policy pack.
type rawPolicyPack struct {
	Name         string          `yaml:"name"`
	Version      string          `yaml:"version"`
	Language     string          `yaml:"language"`
	PathPatterns []string        `yaml:"path_patterns"`
	Rules        []rawPolicyRule `yaml:"rules"`
}

// LoadPolicyPack loads a single policy pack from a YAML file.
func LoadPolicyPack(path string) (domain.PolicyPack, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return domain.PolicyPack{}, fmt.Errorf("reading policy pack: %w", err)
	}
	return parsePolicyPack(data)
}

// LoadPolicyPacks loads all *.yml/*.yaml policy packs from a directory.
func LoadPolicyPacks(dir string) ([]domain.PolicyPack, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("reading policies directory: %w", err)
	}

	var packs []domain.PolicyPack
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if !strings.HasSuffix(name, ".yml") && !strings.HasSuffix(name, ".yaml") {
			continue
		}
		pack, err := LoadPolicyPack(filepath.Join(dir, name))
		if err != nil {
			return nil, fmt.Errorf("loading %s: %w", name, err)
		}
		packs = append(packs, pack)
	}
	return packs, nil
}

func parsePolicyPack(data []byte) (domain.PolicyPack, error) {
	var raw rawPolicyPack
	if err := yaml.Unmarshal(data, &raw); err != nil {
		return domain.PolicyPack{}, fmt.Errorf("parsing policy YAML: %w", err)
	}

	if raw.Name == "" {
		return domain.PolicyPack{}, fmt.Errorf("policy pack missing required field: name")
	}

	pack := domain.PolicyPack{
		Name:         raw.Name,
		Version:      raw.Version,
		Language:     raw.Language,
		PathPatterns: raw.PathPatterns,
	}

	for _, rr := range raw.Rules {
		dim, err := parseDimension(rr.Dimension)
		if err != nil {
			return domain.PolicyPack{}, fmt.Errorf("rule %q: %w", rr.ID, err)
		}
		sev, err := parseSeverity(rr.Severity)
		if err != nil {
			return domain.PolicyPack{}, fmt.Errorf("rule %q: %w", rr.ID, err)
		}
		pack.Rules = append(pack.Rules, domain.PolicyRule{
			ID:          rr.ID,
			Dimension:   dim,
			Description: rr.Description,
			Severity:    sev,
			Threshold:   rr.Threshold,
			Metric:      rr.Metric,
		})
	}

	return pack, nil
}

// parseDimension maps a YAML string to a domain.Dimension.
func parseDimension(s string) (domain.Dimension, error) {
	m := map[string]domain.Dimension{
		"correctness":                 domain.DimCorrectness,
		"maintainability":             domain.DimMaintainability,
		"readability":                 domain.DimReadability,
		"testability":                 domain.DimTestability,
		"security":                    domain.DimSecurity,
		"architectural_fitness":       domain.DimArchitecturalFitness,
		"operational_quality":         domain.DimOperationalQuality,
		"performance_appropriateness": domain.DimPerformanceAppropriateness,
		"change_risk":                 domain.DimChangeRisk,
	}
	if d, ok := m[s]; ok {
		return d, nil
	}
	return 0, fmt.Errorf("unknown dimension: %q", s)
}

// parseSeverity maps a YAML string to a domain.Severity.
func parseSeverity(s string) (domain.Severity, error) {
	return domain.ParseSeverity(s)
}
