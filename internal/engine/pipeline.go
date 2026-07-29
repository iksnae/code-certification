package engine

import (
	"strings"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/expiry"
	"github.com/iksnae/code-certification/internal/language_tiers"
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
	scores := Score(ev, evalResult, unit.ID.Language())

	// 3-5. Derive the verdict.
	//
	// "Scored well", "scored badly" and "cannot be scored" are three distinct
	// states. Only the first two are expressible as a score, so the third is
	// decided before any arithmetic happens: a unit in a language the engine
	// cannot analyse yields no dimension evidence, and running an empty score
	// set through WeightedAverage/GradeFromScore would silently manufacture a
	// 0.0 and an F for code that was never examined.
	//
	// The tier is asked directly rather than inferred from `scores` being nil.
	// Nil-vs-empty would work today but is not a contract: a future change that
	// returns an empty map for an unsupported language, or nil for "supported
	// but no evidence gathered", would reintroduce the fabricated grade with no
	// test failing.
	var (
		avg        float64
		status     domain.Status
		grade      domain.Grade
		confidence float64
	)
	unsupported := !language_tiers.IsSupported(unit.ID.Language())
	if unsupported {
		// No judgement is asserted: not passing, not failing, not assessed.
		// Exempt is the only status that withholds a quality verdict; the
		// Unsupported flag is what distinguishes "outside the engine's reach"
		// from "excluded by an operator override".
		//
		// The four values come from domain.UnassessedVerdict rather than being
		// spelled out here, because the store's legacy backfill has to produce
		// exactly the same verdict. Two literals would be two answers waiting
		// to drift; this way the fresh path and the migrated path are one path.
		v := domain.UnassessedVerdict()
		avg, grade, status, confidence = v.Score, v.Grade, v.Status, v.Confidence
		scores = nil
	} else {
		avg = scores.WeightedAverage(nil)
		status = StatusFromScore(avg, !evalResult.Passed)
		grade = domain.GradeFromScore(avg)
		confidence = 1.0 // Deterministic evidence = full confidence
	}

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
		Confidence:   confidence,
		Unsupported:  unsupported,
		Dimensions:   scores,
		Evidence:     ev,
		Observations: observations,
		CertifiedAt:  now,
		ExpiresAt:    window.ExpiresAt,
		Source:       source,
		Version:      1,
	}
}
