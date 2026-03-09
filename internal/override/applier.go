package override

import (
	"time"

	"github.com/code-certification/certify/internal/domain"
)

// Apply applies a single override to a certification record.
// Returns the record unchanged if the override doesn't match.
func Apply(rec domain.CertificationRecord, o domain.Override) domain.CertificationRecord {
	if rec.UnitID.String() != o.UnitID.String() {
		return rec
	}

	switch o.Action {
	case domain.OverrideExempt:
		rec.Status = domain.StatusExempt
		rec.Observations = append(rec.Observations, "Override: exempt — "+o.Rationale)

	case domain.OverrideExtendWindow:
		// Extend by 50% of remaining window
		remaining := rec.ExpiresAt.Sub(rec.CertifiedAt)
		extension := remaining / 2
		rec.ExpiresAt = rec.ExpiresAt.Add(extension)
		rec.Observations = append(rec.Observations, "Override: window extended — "+o.Rationale)

	case domain.OverrideShortenWindow:
		// Shorten to 50% of current window
		remaining := rec.ExpiresAt.Sub(rec.CertifiedAt)
		rec.ExpiresAt = rec.CertifiedAt.Add(remaining / 2)
		rec.Observations = append(rec.Observations, "Override: window shortened — "+o.Rationale)

	case domain.OverrideForceReview:
		// Mark as probationary to force re-evaluation
		if rec.Status == domain.StatusCertified || rec.Status == domain.StatusCertifiedWithObservations {
			rec.Status = domain.StatusProbationary
		}
		// Set expiry to now (force immediate recertification)
		rec.ExpiresAt = time.Now()
		rec.Observations = append(rec.Observations, "Override: forced review — "+o.Rationale)
	}

	return rec
}

// ApplyAll applies all matching overrides to a record.
func ApplyAll(rec domain.CertificationRecord, overrides []domain.Override) domain.CertificationRecord {
	for _, o := range overrides {
		rec = Apply(rec, o)
	}
	return rec
}
