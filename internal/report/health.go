// Package report generates certification reports from records.
package report

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/iksnae/code-certification/internal/analysis"
	"github.com/iksnae/code-certification/internal/domain"
)

// HealthReport summarizes the certification state of a repository.
type HealthReport struct {
	TotalUnits       int     `json:"total_units"`
	Certified        int     `json:"certified"`
	CertifiedWithObs int     `json:"certified_with_observations"`
	Probationary     int     `json:"probationary"`
	Expired          int     `json:"expired"`
	Decertified      int     `json:"decertified"`
	Exempt           int     `json:"exempt"`
	PassRate         float64 `json:"pass_rate"`
	AverageScore     float64 `json:"average_score"`

	// Unsupported counts units in languages the engine cannot analyse, and
	// AnalyzableUnits is TotalUnits minus that count. Every bucket above is
	// counted over analyzable units only, so without these two the report shows
	// a unit total that no bucket accounts for. See Card.
	Unsupported     int `json:"unsupported"`
	AnalyzableUnits int `json:"analyzable_units"`
}

// PassRateKnown reports whether PassRate and AverageScore are measurements
// rather than placeholders. See Card.PassRateKnown.
func (h HealthReport) PassRateKnown() bool { return h.AnalyzableUnits > 0 }

// Health computes a health report from certification records.
func Health(records []domain.CertificationRecord) HealthReport {
	if len(records) == 0 {
		return HealthReport{}
	}

	var h HealthReport
	h.TotalUnits = len(records)

	var totalScore float64
	var passing int
	var counted int

	for _, r := range records {
		if analysis.ForLanguage(r.UnitID.Language()) == nil {
			continue
		}
		counted++

		totalScore += r.Score

		switch r.Status {
		case domain.StatusCertified:
			h.Certified++
			passing++
		case domain.StatusCertifiedWithObservations:
			h.CertifiedWithObs++
			passing++
		case domain.StatusProbationary:
			h.Probationary++
		case domain.StatusExpired:
			h.Expired++
		case domain.StatusDecertified:
			h.Decertified++
		case domain.StatusExempt:
			h.Exempt++
			passing++
		}
	}

	// Every bucket above is counted over analyzable units only, so TotalUnits
	// alone leaves the difference unexplained: an all-unassessed repo showed
	// three units and seven zeroes.
	h.AnalyzableUnits = counted
	h.Unsupported = h.TotalUnits - counted

	if h.PassRateKnown() {
		h.PassRate = float64(passing) / float64(counted)
		h.AverageScore = totalScore / float64(counted)
	}

	return h
}

// FormatJSON produces JSON output for a report.
func FormatJSON(v any) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}

// FormatText produces a human-readable text report.
func FormatText(h HealthReport) string {
	var b strings.Builder
	b.WriteString("═══════════════════════════════════════════\n")
	b.WriteString("  Certify — Health Report\n")
	b.WriteString("═══════════════════════════════════════════\n\n")

	fmt.Fprintf(&b, "  Total Units:            %d\n", h.TotalUnits)
	fmt.Fprintf(&b, "  Certified:              %d\n", h.Certified)
	fmt.Fprintf(&b, "  Certified w/ Obs:       %d\n", h.CertifiedWithObs)
	fmt.Fprintf(&b, "  Probationary:           %d\n", h.Probationary)
	fmt.Fprintf(&b, "  Expired:                %d\n", h.Expired)
	fmt.Fprintf(&b, "  Decertified:            %d\n", h.Decertified)
	fmt.Fprintf(&b, "  Exempt:                 %d\n", h.Exempt)
	if h.Unsupported > 0 {
		fmt.Fprintf(&b, "  Not Assessed:           %d\n", h.Unsupported)
	}
	b.WriteString("\n")
	// With nothing analyzed both figures are 0/0. Printing "0.0%" here states
	// that an assessment ran and found total failure, which is the same false
	// claim as a 100% pass rate over units the engine never opened — inverted.
	fmt.Fprintf(&b, "  Pass Rate:              %s\n", FormatRate(h.PassRateKnown(), h.PassRate, 1))
	avgScore := "n/a"
	if h.PassRateKnown() {
		avgScore = fmt.Sprintf("%.3f", h.AverageScore)
	}
	fmt.Fprintf(&b, "  Average Score:          %s\n", avgScore)
	b.WriteString("\n═══════════════════════════════════════════\n")

	return b.String()
}
