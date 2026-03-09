package config

import (
	"github.com/code-certification/certify/internal/domain"
	"github.com/code-certification/certify/internal/policy"
)

// NewPolicyMatcher creates a policy matcher from loaded policy packs.
func NewPolicyMatcher(packs []domain.PolicyPack) *policy.Matcher {
	return policy.NewMatcher(packs)
}
