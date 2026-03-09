package config

import (
	"github.com/code-certification/certify/internal/domain"
	"github.com/code-certification/certify/internal/policy"
)

// NewPolicyMatcher creates a policy matcher from loaded policy packs.
func NewPolicyMatcher(packs []domain.PolicyPack) *policy.Matcher {
	return policy.NewMatcher(packs)
}

// FilterPolicyPacks applies enable/disable config to policy packs.
func FilterPolicyPacks(packs []domain.PolicyPack, cfg domain.PolicyConfig) []domain.PolicyPack {
	if len(cfg.Enabled) == 0 && len(cfg.Disabled) == 0 {
		return packs
	}

	enabled := make(map[string]bool)
	for _, n := range cfg.Enabled {
		enabled[n] = true
	}
	disabled := make(map[string]bool)
	for _, n := range cfg.Disabled {
		disabled[n] = true
	}

	var result []domain.PolicyPack
	for _, p := range packs {
		if len(enabled) > 0 && !enabled[p.Name] {
			continue
		}
		if disabled[p.Name] {
			continue
		}
		result = append(result, p)
	}
	return result
}
