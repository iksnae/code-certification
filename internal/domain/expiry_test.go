package domain_test

import (
	"testing"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
)

func TestExpiryWindow_IsExpired(t *testing.T) {
	now := time.Now()

	active := domain.ExpiryWindow{
		CertifiedAt: now.Add(-24 * time.Hour),
		ExpiresAt:   now.Add(24 * time.Hour),
	}
	if active.IsExpired(now) {
		t.Error("active window should not be expired")
	}

	expired := domain.ExpiryWindow{
		CertifiedAt: now.Add(-48 * time.Hour),
		ExpiresAt:   now.Add(-24 * time.Hour),
	}
	if !expired.IsExpired(now) {
		t.Error("past window should be expired")
	}
}

func TestExpiryWindow_Duration(t *testing.T) {
	now := time.Now()
	w := domain.ExpiryWindow{
		CertifiedAt: now,
		ExpiresAt:   now.Add(90 * 24 * time.Hour),
	}
	d := w.Duration()
	if d != 90*24*time.Hour {
		t.Errorf("Duration() = %v, want 90 days", d)
	}
}

func TestExpiryWindow_RemainingAt(t *testing.T) {
	now := time.Now()
	w := domain.ExpiryWindow{
		CertifiedAt: now,
		ExpiresAt:   now.Add(90 * 24 * time.Hour),
	}

	remaining := w.RemainingAt(now.Add(30 * 24 * time.Hour))
	expected := 60 * 24 * time.Hour
	if remaining != expected {
		t.Errorf("RemainingAt(+30d) = %v, want %v", remaining, expected)
	}

	// Already expired returns 0
	remaining = w.RemainingAt(now.Add(100 * 24 * time.Hour))
	if remaining != 0 {
		t.Errorf("RemainingAt(+100d) = %v, want 0", remaining)
	}
}

func TestExpiryFactors(t *testing.T) {
	f := domain.ExpiryFactors{
		BaseWindowDays:    90,
		ChurnRate:         0.5,
		TestCoverage:      0.8,
		Complexity:        15.0,
		PriorPassCount:    3,
		PriorFailCount:    1,
		SecuritySensitive: false,
	}

	if f.BaseWindowDays != 90 {
		t.Errorf("BaseWindowDays = %d, want 90", f.BaseWindowDays)
	}
	if f.ChurnRate != 0.5 {
		t.Errorf("ChurnRate = %f, want 0.5", f.ChurnRate)
	}
}
