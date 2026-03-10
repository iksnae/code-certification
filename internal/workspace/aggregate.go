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
	OverallGrade string             `json:"overall_grade"`
	OverallScore float64            `json:"overall_score"`
	PassRate     float64            `json:"pass_rate"`
}

// SubmoduleSummary holds certification stats for a single submodule.
type SubmoduleSummary struct {
	Name       string  `json:"name"`
	Path       string  `json:"path"`
	Grade      string  `json:"grade"`
	Score      float64 `json:"score"`
	Units      int     `json:"units"`
	Passing    int     `json:"passing"`
	Failing    int     `json:"failing"`
	PassRate   float64 `json:"pass_rate"`
	HasCertify bool    `json:"has_certify"`
	Commit     string  `json:"commit,omitempty"`
	StateAt    string  `json:"state_at,omitempty"` // when the submodule was last certified
}

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
		totalWeightedScore += s.Score * float64(s.Units)
	}

	wc.TotalUnits = totalUnits

	if totalUnits == 0 {
		wc.OverallGrade = "N/A"
		return wc
	}

	wc.OverallScore = totalWeightedScore / float64(totalUnits)
	wc.OverallGrade = domain.GradeFromScore(wc.OverallScore).String()
	wc.PassRate = float64(wc.TotalPassing) / float64(totalUnits)

	return wc
}
