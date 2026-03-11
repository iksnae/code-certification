package engine

import (
	"strings"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/expiry"
	"github.com/iksnae/code-certification/internal/policy"
)

// CertifyUnit runs the full certification pipeline for a single unit.
func CertifyUnit(
	unit domain.Unit,
	rules []domain.PolicyRule,
	ev []domain.Evidence,
	expiryCfg domain.ExpiryConfig,
	now time.Time,
) domain.CertificationRecord {
	// 1. Evaluate policy rules against evidence
	evalResult := policy.Evaluate(rules, ev, unit.ID.Path())

	// 2. Score across dimensions
	scores := Score(ev, evalResult)

	// 3. Compute weighted average
	avg := scores.WeightedAverage(nil)

	// 4. Determine status
	status := StatusFromScore(avg, !evalResult.Passed)

	// 5. Compute grade
	grade := domain.GradeFromScore(avg)

	// 6. Compute expiry window
	factors := domain.ExpiryFactors{
		BaseWindowDays: expiryCfg.DefaultWindowDays,
	}
	window := expiry.Calculate(expiryCfg, factors, now)

	// 7. Build observations from violations
	var observations []string
	for _, v := range evalResult.Violations {
		observations = append(observations, v.Description)
	}

	// 8. Determine source attribution
	source := "deterministic"
	for _, e := range ev {
		if e.Kind == domain.EvidenceKindAgentReview && (strings.HasPrefix(e.Source, "agent:") || strings.HasPrefix(e.Source, "agent-prescreen:") || strings.HasPrefix(e.Source, "agent-deep-review:")) {
			source = "deterministic+" + e.Source
			break
		}
	}

	return domain.CertificationRecord{
		UnitID:       unit.ID,
		UnitType:     unit.Type,
		UnitPath:     unit.ID.Path(),
		Status:       status,
		Grade:        grade,
		Score:        avg,
		Confidence:   1.0, // Deterministic evidence = full confidence
		Dimensions:   scores,
		Evidence:     ev,
		Observations: observations,
		CertifiedAt:  now,
		ExpiresAt:    window.ExpiresAt,
		Source:       source,
		Version:      1,
	}
}
