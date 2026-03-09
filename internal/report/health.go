// Package report generates certification reports from records.
package report

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/code-certification/certify/internal/domain"
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
}

// Health computes a health report from certification records.
func Health(records []domain.CertificationRecord) HealthReport {
	if len(records) == 0 {
		return HealthReport{}
	}

	var h HealthReport
	h.TotalUnits = len(records)

	var totalScore float64
	var passing int

	for _, r := range records {
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

	h.PassRate = float64(passing) / float64(h.TotalUnits)
	h.AverageScore = totalScore / float64(h.TotalUnits)

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
	b.WriteString("  Code Certification Health Report\n")
	b.WriteString("═══════════════════════════════════════════\n\n")

	fmt.Fprintf(&b, "  Total Units:            %d\n", h.TotalUnits)
	fmt.Fprintf(&b, "  Certified:              %d\n", h.Certified)
	fmt.Fprintf(&b, "  Certified w/ Obs:       %d\n", h.CertifiedWithObs)
	fmt.Fprintf(&b, "  Probationary:           %d\n", h.Probationary)
	fmt.Fprintf(&b, "  Expired:                %d\n", h.Expired)
	fmt.Fprintf(&b, "  Decertified:            %d\n", h.Decertified)
	fmt.Fprintf(&b, "  Exempt:                 %d\n", h.Exempt)
	b.WriteString("\n")
	fmt.Fprintf(&b, "  Pass Rate:              %.1f%%\n", h.PassRate*100)
	fmt.Fprintf(&b, "  Average Score:          %.3f\n", h.AverageScore)
	b.WriteString("\n═══════════════════════════════════════════\n")

	return b.String()
}
