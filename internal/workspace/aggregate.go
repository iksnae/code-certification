package workspace

import (
	"time"

	"github.com/iksnae/code-certification/internal/domain"
)

// WorkspaceCard is an aggregated report card across all submodules.
type WorkspaceCard struct {
	GeneratedAt  string             `json:"generated_at"`
	Submodules   []SubmoduleSummary `json:"submodules"`
	TotalUnits   int                `json:"total_units"`
	TotalPassing int                `json:"total_passing"`
	TotalFailing int                `json:"total_failing"`
	// TotalUnsupported counts units the engine could not analyse anywhere in
	// the workspace. AnalyzableUnits is the PassRate denominator.
	TotalUnsupported int     `json:"total_unsupported"`
	AnalyzableUnits  int     `json:"analyzable_units"`
	OverallGrade     string  `json:"overall_grade"`
	OverallScore     float64 `json:"overall_score"`
	PassRate         float64 `json:"pass_rate"`
}

// PassRateKnown reports whether PassRate is a measurement. See Card.PassRateKnown.
func (wc WorkspaceCard) PassRateKnown() bool { return wc.AnalyzableUnits > 0 }

// ScoreKnown reports whether OverallScore and OverallGrade are measurements.
// See Card.ScoreKnown.
func (wc WorkspaceCard) ScoreKnown() bool { return wc.AnalyzableUnits > 0 }

// SubmoduleSummary holds certification stats for a single submodule.
type SubmoduleSummary struct {
	Name    string  `json:"name"`
	Path    string  `json:"path"`
	Grade   string  `json:"grade"`
	Score   float64 `json:"score"`
	Units   int     `json:"units"`
	Passing int     `json:"passing"`
	Failing int     `json:"failing"`
	// Unsupported counts this submodule's unassessed units. Analyzable is
	// derived (Units - Unsupported) rather than stored, so a summary can never
	// carry two counts that disagree.
	Unsupported int     `json:"unsupported"`
	PassRate    float64 `json:"pass_rate"`
	HasCertify  bool    `json:"has_certify"`
	Commit      string  `json:"commit,omitempty"`
	StateAt     string  `json:"state_at,omitempty"` // when the submodule was last certified
}

// Analyzable is the number of units in this submodule about which a verdict
// was asserted — the denominator of PassRate.
func (s SubmoduleSummary) Analyzable() int { return s.Units - s.Unsupported }

// PassRateKnown reports whether PassRate is a measurement. See Card.PassRateKnown.
func (s SubmoduleSummary) PassRateKnown() bool { return s.Analyzable() > 0 }

// ScoreKnown reports whether Score and Grade are measurements. See
// Card.ScoreKnown.
func (s SubmoduleSummary) ScoreKnown() bool { return s.Analyzable() > 0 }

// AggregateCards builds a WorkspaceCard from submodule summaries.
// Submodules without certify setup (HasCertify=false) or with zero units
// are included in the listing but excluded from score aggregation.
func AggregateCards(subs []SubmoduleSummary) WorkspaceCard {
	wc := WorkspaceCard{
		GeneratedAt: time.Now().Format(time.RFC3339),
		Submodules:  subs,
	}

	if len(subs) == 0 {
		wc.OverallGrade = "N/A"
		return wc
	}

	var totalWeightedScore float64
	var totalUnits int

	for _, s := range subs {
		if !s.HasCertify || s.Units == 0 {
			continue
		}
		totalUnits += s.Units
		wc.TotalPassing += s.Passing
		wc.TotalFailing += s.Failing
		wc.TotalUnsupported += s.Unsupported
		totalWeightedScore += s.Score * float64(s.Units)
	}

	wc.TotalUnits = totalUnits
	wc.AnalyzableUnits = totalUnits - wc.TotalUnsupported

	if totalUnits == 0 {
		wc.OverallGrade = "N/A"
		return wc
	}

	// A grade is a claim about assessed code. With nothing analyzable anywhere
	// in the workspace the weighted mean is taken over placeholder zeroes only,
	// and printing "F (0.0%)" from it states a definite failure beside a pass
	// rate that already admits nothing was measured. The denominator when SOME
	// unit was analyzable is issue #32 and is left alone — it must move together
	// with Card.OverallScore, which the submodule scores summed here come from.
	if wc.ScoreKnown() {
		wc.OverallScore = totalWeightedScore / float64(totalUnits)
		wc.OverallGrade = domain.GradeFromScore(wc.OverallScore).String()
	} else {
		wc.OverallGrade = domain.GradeNA.String()
	}
	// Unassessed units leave both sides of the ratio. With none analyzable the
	// rate is 0/0 — undefined — and PassRateKnown() tells the renderers to say
	// so instead of printing the zero value as a measurement.
	if wc.PassRateKnown() {
		wc.PassRate = float64(wc.TotalPassing) / float64(wc.AnalyzableUnits)
	}

	return wc
}
