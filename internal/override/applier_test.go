package override_test

import (
	"testing"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/override"
)

func TestApply_Exempt(t *testing.T) {
	rec := domain.CertificationRecord{
		UnitID:    domain.NewUnitID("go", "legacy/old.go", "Deprecated"),
		Status:    domain.StatusDecertified,
		Score:     0.3,
		ExpiresAt: time.Now().Add(90 * 24 * time.Hour),
	}

	o := domain.Override{
		UnitID:    rec.UnitID,
		Action:    domain.OverrideExempt,
		Rationale: "Legacy code",
		Actor:     "kmills",
	}

	result := override.Apply(rec, o)
	if result.Status != domain.StatusExempt {
		t.Errorf("status = %v, want exempt", result.Status)
	}
}

func TestApply_ExtendWindow(t *testing.T) {
	now := time.Now()
	rec := domain.CertificationRecord{
		UnitID:      domain.NewUnitID("go", "service/sync.go", "Apply"),
		Status:      domain.StatusCertified,
		CertifiedAt: now,
		ExpiresAt:   now.Add(90 * 24 * time.Hour),
	}

	o := domain.Override{
		UnitID:    rec.UnitID,
		Action:    domain.OverrideExtendWindow,
		Rationale: "Stable function",
		Actor:     "kmills",
	}

	result := override.Apply(rec, o)
	if !result.ExpiresAt.After(rec.ExpiresAt) {
		t.Error("extended window should expire later")
	}
}

func TestApply_ShortenWindow(t *testing.T) {
	now := time.Now()
	rec := domain.CertificationRecord{
		UnitID:      domain.NewUnitID("go", "crypto/hmac.go", "Sign"),
		Status:      domain.StatusCertified,
		CertifiedAt: now,
		ExpiresAt:   now.Add(90 * 24 * time.Hour),
	}

	o := domain.Override{
		UnitID:    rec.UnitID,
		Action:    domain.OverrideShortenWindow,
		Rationale: "Security sensitive",
		Actor:     "security-team",
	}

	result := override.Apply(rec, o)
	if !result.ExpiresAt.Before(rec.ExpiresAt) {
		t.Error("shortened window should expire sooner")
	}
}

func TestApply_NoMatchingOverride(t *testing.T) {
	rec := domain.CertificationRecord{
		UnitID: domain.NewUnitID("go", "main.go", "main"),
		Status: domain.StatusCertified,
		Score:  0.9,
	}

	// Override for different unit
	o := domain.Override{
		UnitID:    domain.NewUnitID("go", "other.go", "other"),
		Action:    domain.OverrideExempt,
		Rationale: "test",
		Actor:     "test",
	}

	result := override.Apply(rec, o)
	// Should be unchanged
	if result.Status != rec.Status {
		t.Error("non-matching override should not change status")
	}
}

func TestApplyAll(t *testing.T) {
	rec := domain.CertificationRecord{
		UnitID:    domain.NewUnitID("go", "legacy/old.go", "Deprecated"),
		Status:    domain.StatusDecertified,
		ExpiresAt: time.Now().Add(90 * 24 * time.Hour),
	}

	overrides := []domain.Override{
		{
			UnitID:    domain.NewUnitID("go", "other.go", "other"),
			Action:    domain.OverrideExempt,
			Rationale: "not this one",
			Actor:     "test",
		},
		{
			UnitID:    domain.NewUnitID("go", "legacy/old.go", "Deprecated"),
			Action:    domain.OverrideExempt,
			Rationale: "Legacy code",
			Actor:     "kmills",
		},
	}

	result := override.ApplyAll(rec, overrides)
	if result.Status != domain.StatusExempt {
		t.Errorf("status = %v, want exempt", result.Status)
	}
}
