// Package expiry computes certification expiry windows based on risk factors.
package expiry

import (
	"math"
	"time"

	"github.com/code-certification/certify/internal/domain"
)

// Calculate computes an expiry window based on config, risk factors, and current time.
func Calculate(cfg domain.ExpiryConfig, factors domain.ExpiryFactors, now time.Time) domain.ExpiryWindow {
	base := float64(factors.BaseWindowDays)
	if base <= 0 {
		base = float64(cfg.DefaultWindowDays)
	}

	multiplier := 1.0

	// High churn shortens the window
	if factors.ChurnRate > 0 {
		// At churn rate 1.0, reduce by ~40%
		multiplier *= 1.0 / (1.0 + factors.ChurnRate*0.6)
	}

	// Low test coverage shortens the window
	if factors.TestCoverage < 0.8 {
		// Scale: 0% coverage → 0.5x, 80%+ → 1.0x
		multiplier *= 0.5 + (factors.TestCoverage * 0.625)
	}

	// High complexity shortens the window
	if factors.Complexity > 15 {
		multiplier *= math.Max(0.5, 1.0-((factors.Complexity-15)*0.02))
	}

	// Security-sensitive code gets shorter windows
	if factors.SecuritySensitive {
		multiplier *= 0.5
	}

	// Prior pass history extends the window slightly
	if factors.PriorPassCount > 0 {
		bonus := math.Min(0.3, float64(factors.PriorPassCount)*0.05)
		multiplier *= (1.0 + bonus)
	}

	// Prior failures shorten the window
	if factors.PriorFailCount > 0 {
		penalty := math.Min(0.3, float64(factors.PriorFailCount)*0.1)
		multiplier *= (1.0 - penalty)
	}

	days := int(math.Round(base * multiplier))

	// Clamp to configured bounds
	if days < cfg.MinWindowDays {
		days = cfg.MinWindowDays
	}
	if days > cfg.MaxWindowDays {
		days = cfg.MaxWindowDays
	}

	return domain.ExpiryWindow{
		CertifiedAt: now,
		ExpiresAt:   now.Add(time.Duration(days) * 24 * time.Hour),
	}
}
