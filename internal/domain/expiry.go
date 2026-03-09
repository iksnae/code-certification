package domain

import "time"

// ExpiryWindow represents the time bounds of a certification.
type ExpiryWindow struct {
	CertifiedAt time.Time `json:"certified_at"`
	ExpiresAt   time.Time `json:"expires_at"`
}

// IsExpired returns true if the window has elapsed at the given time.
func (w ExpiryWindow) IsExpired(at time.Time) bool {
	return at.After(w.ExpiresAt) || at.Equal(w.ExpiresAt)
}

// Duration returns the total duration of the certification window.
func (w ExpiryWindow) Duration() time.Duration {
	return w.ExpiresAt.Sub(w.CertifiedAt)
}

// RemainingAt returns how much time is left at the given moment.
// Returns 0 if already expired.
func (w ExpiryWindow) RemainingAt(at time.Time) time.Duration {
	remaining := w.ExpiresAt.Sub(at)
	if remaining < 0 {
		return 0
	}
	return remaining
}

// ExpiryFactors holds the inputs used to compute an expiry window.
type ExpiryFactors struct {
	BaseWindowDays    int     `json:"base_window_days"`
	ChurnRate         float64 `json:"churn_rate"`          // Changes per time period (0.0–1.0+)
	TestCoverage      float64 `json:"test_coverage"`       // 0.0–1.0
	Complexity        float64 `json:"complexity"`          // Cyclomatic complexity
	PriorPassCount    int     `json:"prior_pass_count"`
	PriorFailCount    int     `json:"prior_fail_count"`
	SecuritySensitive bool    `json:"security_sensitive"`
}
