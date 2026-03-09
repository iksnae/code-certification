package expiry_test

import (
	"testing"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/expiry"
)

func TestCalculate_DefaultWindow(t *testing.T) {
	cfg := domain.ExpiryConfig{
		DefaultWindowDays: 90,
		MinWindowDays:     7,
		MaxWindowDays:     365,
	}
	factors := domain.ExpiryFactors{
		BaseWindowDays: 90,
		ChurnRate:      0.1, // Low churn
		TestCoverage:   0.8,
		Complexity:     10,
	}

	now := time.Now()
	window := expiry.Calculate(cfg, factors, now)

	if window.CertifiedAt != now {
		t.Error("CertifiedAt should be now")
	}

	days := int(window.Duration().Hours() / 24)
	if days < cfg.MinWindowDays || days > cfg.MaxWindowDays {
		t.Errorf("window %d days outside bounds [%d, %d]", days, cfg.MinWindowDays, cfg.MaxWindowDays)
	}
}

func TestCalculate_HighChurnShortensWindow(t *testing.T) {
	cfg := domain.ExpiryConfig{
		DefaultWindowDays: 90,
		MinWindowDays:     7,
		MaxWindowDays:     365,
	}

	low := domain.ExpiryFactors{ChurnRate: 0.1, TestCoverage: 0.8, Complexity: 10, BaseWindowDays: 90}
	high := domain.ExpiryFactors{ChurnRate: 1.0, TestCoverage: 0.8, Complexity: 10, BaseWindowDays: 90}

	now := time.Now()
	lowWindow := expiry.Calculate(cfg, low, now)
	highWindow := expiry.Calculate(cfg, high, now)

	if highWindow.Duration() >= lowWindow.Duration() {
		t.Error("high churn should produce shorter window")
	}
}

func TestCalculate_LowCoverageShortensWindow(t *testing.T) {
	cfg := domain.ExpiryConfig{
		DefaultWindowDays: 90,
		MinWindowDays:     7,
		MaxWindowDays:     365,
	}

	high := domain.ExpiryFactors{ChurnRate: 0.1, TestCoverage: 0.9, Complexity: 10, BaseWindowDays: 90}
	low := domain.ExpiryFactors{ChurnRate: 0.1, TestCoverage: 0.2, Complexity: 10, BaseWindowDays: 90}

	now := time.Now()
	highWindow := expiry.Calculate(cfg, high, now)
	lowWindow := expiry.Calculate(cfg, low, now)

	if lowWindow.Duration() >= highWindow.Duration() {
		t.Error("low coverage should produce shorter window")
	}
}

func TestCalculate_SecuritySensitiveShortensWindow(t *testing.T) {
	cfg := domain.ExpiryConfig{
		DefaultWindowDays: 90,
		MinWindowDays:     7,
		MaxWindowDays:     365,
	}

	normal := domain.ExpiryFactors{ChurnRate: 0.1, TestCoverage: 0.8, BaseWindowDays: 90}
	secure := domain.ExpiryFactors{ChurnRate: 0.1, TestCoverage: 0.8, BaseWindowDays: 90, SecuritySensitive: true}

	now := time.Now()
	normalW := expiry.Calculate(cfg, normal, now)
	secureW := expiry.Calculate(cfg, secure, now)

	if secureW.Duration() >= normalW.Duration() {
		t.Error("security-sensitive should produce shorter window")
	}
}

func TestCalculate_ClampsToMinMax(t *testing.T) {
	cfg := domain.ExpiryConfig{
		DefaultWindowDays: 90,
		MinWindowDays:     30,
		MaxWindowDays:     180,
	}

	// Extreme high churn should still be >= min
	extreme := domain.ExpiryFactors{ChurnRate: 10.0, TestCoverage: 0.0, Complexity: 100, BaseWindowDays: 90, SecuritySensitive: true}
	now := time.Now()
	window := expiry.Calculate(cfg, extreme, now)

	days := int(window.Duration().Hours() / 24)
	if days < cfg.MinWindowDays {
		t.Errorf("window %d days below min %d", days, cfg.MinWindowDays)
	}

	// Very stable should be <= max
	stable := domain.ExpiryFactors{ChurnRate: 0.0, TestCoverage: 1.0, Complexity: 1, BaseWindowDays: 90, PriorPassCount: 100}
	window = expiry.Calculate(cfg, stable, now)
	days = int(window.Duration().Hours() / 24)
	if days > cfg.MaxWindowDays {
		t.Errorf("window %d days above max %d", days, cfg.MaxWindowDays)
	}
}
