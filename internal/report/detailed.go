package report

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/code-certification/certify/internal/domain"
)

// DetailedReport extends HealthReport with richer analysis.
type DetailedReport struct {
	HealthReport

	// Dimension-level averages
	Dimensions map[string]float64 `json:"dimensions,omitempty"`

	// By-language breakdown
	ByLanguage map[string]LanguageBreakdown `json:"by_language,omitempty"`

	// Expiring-soon units (within 14 days)
	ExpiringSoon []UnitSummary `json:"expiring_soon,omitempty"`

	// Highest-risk units (lowest scores)
	HighestRisk []UnitSummary `json:"highest_risk,omitempty"`

	// Failing/non-passing units with explanations
	Failing []UnitSummary `json:"failing,omitempty"`

	// Recurrently failing areas (directories with multiple failing units)
	RecurrentlyFailing []AreaSummary `json:"recurrently_failing,omitempty"`
}

// AreaSummary summarizes certification for a directory or package.
type AreaSummary struct {
	Path         string  `json:"path"`
	Total        int     `json:"total"`
	Failing      int     `json:"failing"`
	AverageScore float64 `json:"average_score"`
}

// LanguageBreakdown summarizes certification status for a single language.
type LanguageBreakdown struct {
	Total        int     `json:"total"`
	Passing      int     `json:"passing"`
	AverageScore float64 `json:"average_score"`
}

// UnitSummary is a compact representation of a unit for report lists.
type UnitSummary struct {
	UnitID       string   `json:"unit_id"`
	Path         string   `json:"path"`
	Status       string   `json:"status"`
	Grade        string   `json:"grade"`
	Score        float64  `json:"score"`
	ExpiresAt    string   `json:"expires_at,omitempty"`
	Explanation  string   `json:"explanation,omitempty"`
	Observations []string `json:"observations,omitempty"`
}

// Detailed computes a full detailed report from certification records.
func Detailed(records []domain.CertificationRecord, now time.Time) DetailedReport {
	d := DetailedReport{
		HealthReport: Health(records),
		Dimensions:   make(map[string]float64),
		ByLanguage:   make(map[string]LanguageBreakdown),
	}

	if len(records) == 0 {
		return d
	}

	// Dimension averages
	dimSums := make(map[string]float64)
	dimCounts := make(map[string]int)
	for _, r := range records {
		for dim, score := range r.Dimensions {
			dimSums[dim.String()] += score
			dimCounts[dim.String()]++
		}
	}
	for dim, sum := range dimSums {
		if dimCounts[dim] > 0 {
			d.Dimensions[dim] = sum / float64(dimCounts[dim])
		}
	}

	// By-language breakdown
	langTotals := make(map[string]int)
	langPassing := make(map[string]int)
	langScores := make(map[string]float64)
	for _, r := range records {
		lang := r.UnitID.Language()
		langTotals[lang]++
		langScores[lang] += r.Score
		if r.Status.IsPassing() {
			langPassing[lang]++
		}
	}
	for lang, total := range langTotals {
		d.ByLanguage[lang] = LanguageBreakdown{
			Total:        total,
			Passing:      langPassing[lang],
			AverageScore: langScores[lang] / float64(total),
		}
	}

	// Expiring soon (within 14 days)
	threshold := now.Add(14 * 24 * time.Hour)
	for _, r := range records {
		if r.Status.IsPassing() && !r.ExpiresAt.IsZero() && r.ExpiresAt.Before(threshold) && r.ExpiresAt.After(now) {
			d.ExpiringSoon = append(d.ExpiringSoon, unitSummaryFrom(r))
		}
	}
	sort.Slice(d.ExpiringSoon, func(i, j int) bool {
		return d.ExpiringSoon[i].ExpiresAt < d.ExpiringSoon[j].ExpiresAt
	})

	// Highest risk (bottom 10 by score, non-exempt)
	sorted := make([]domain.CertificationRecord, len(records))
	copy(sorted, records)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Score < sorted[j].Score
	})
	for i, r := range sorted {
		if i >= 10 {
			break
		}
		if r.Status == domain.StatusExempt {
			continue
		}
		d.HighestRisk = append(d.HighestRisk, unitSummaryFrom(r))
	}

	// Failing units with explanations
	for _, r := range records {
		if !r.Status.IsPassing() {
			s := unitSummaryFrom(r)
			s.Explanation = explainStatus(r)
			s.Observations = r.Observations
			d.Failing = append(d.Failing, s)
		}
	}

	// Recurrently failing areas (directories with 2+ failing units)
	dirTotals := make(map[string]int)
	dirFailing := make(map[string]int)
	dirScores := make(map[string]float64)
	for _, r := range records {
		dir := filepath.Dir(r.UnitPath)
		dirTotals[dir]++
		dirScores[dir] += r.Score
		if !r.Status.IsPassing() {
			dirFailing[dir]++
		}
	}
	for dir, failing := range dirFailing {
		if failing >= 2 {
			d.RecurrentlyFailing = append(d.RecurrentlyFailing, AreaSummary{
				Path:         dir,
				Total:        dirTotals[dir],
				Failing:      failing,
				AverageScore: dirScores[dir] / float64(dirTotals[dir]),
			})
		}
	}
	sort.Slice(d.RecurrentlyFailing, func(i, j int) bool {
		return d.RecurrentlyFailing[i].Failing > d.RecurrentlyFailing[j].Failing
	})

	return d
}

func unitSummaryFrom(r domain.CertificationRecord) UnitSummary {
	return UnitSummary{
		UnitID:    r.UnitID.String(),
		Path:      r.UnitPath,
		Status:    r.Status.String(),
		Grade:     r.Grade.String(),
		Score:     r.Score,
		ExpiresAt: r.ExpiresAt.Format(time.RFC3339),
	}
}

func explainStatus(r domain.CertificationRecord) string {
	switch r.Status {
	case domain.StatusDecertified:
		if len(r.Observations) > 0 {
			return fmt.Sprintf("Decertified: %s", r.Observations[0])
		}
		return fmt.Sprintf("Decertified: score %.2f below threshold", r.Score)
	case domain.StatusProbationary:
		return fmt.Sprintf("Probationary: score %.2f needs improvement", r.Score)
	case domain.StatusExpired:
		return fmt.Sprintf("Expired: certification window elapsed at %s", r.ExpiresAt.Format("2006-01-02"))
	default:
		return r.Status.String()
	}
}

// FormatDetailedText produces a human-readable detailed report.
func FormatDetailedText(d DetailedReport) string {
	var b strings.Builder

	// Base health report
	b.WriteString(FormatText(d.HealthReport))

	// Dimension breakdown
	if len(d.Dimensions) > 0 {
		b.WriteString("\n  Dimension Averages:\n")
		// Sort for deterministic output
		dims := make([]string, 0, len(d.Dimensions))
		for k := range d.Dimensions {
			dims = append(dims, k)
		}
		sort.Strings(dims)
		for _, dim := range dims {
			fmt.Fprintf(&b, "    %-25s %.3f\n", dim, d.Dimensions[dim])
		}
	}

	// By language
	if len(d.ByLanguage) > 0 {
		b.WriteString("\n  By Language:\n")
		langs := make([]string, 0, len(d.ByLanguage))
		for k := range d.ByLanguage {
			langs = append(langs, k)
		}
		sort.Strings(langs)
		for _, lang := range langs {
			lb := d.ByLanguage[lang]
			fmt.Fprintf(&b, "    %-10s %d units, %d passing, avg %.3f\n", lang, lb.Total, lb.Passing, lb.AverageScore)
		}
	}

	// Expiring soon
	if len(d.ExpiringSoon) > 0 {
		fmt.Fprintf(&b, "\n  Expiring Soon (%d):\n", len(d.ExpiringSoon))
		for _, u := range d.ExpiringSoon {
			fmt.Fprintf(&b, "    %s — expires %s\n", u.UnitID, u.ExpiresAt[:10])
		}
	}

	// Highest risk
	if len(d.HighestRisk) > 0 {
		fmt.Fprintf(&b, "\n  Highest Risk (bottom %d):\n", len(d.HighestRisk))
		for _, u := range d.HighestRisk {
			fmt.Fprintf(&b, "    %s — %s (%.3f)\n", u.UnitID, u.Grade, u.Score)
		}
	}

	// Failing
	if len(d.Failing) > 0 {
		fmt.Fprintf(&b, "\n  Failing (%d):\n", len(d.Failing))
		for _, u := range d.Failing {
			fmt.Fprintf(&b, "    %s — %s\n", u.UnitID, u.Explanation)
		}
	}

	// Recurrently failing areas
	if len(d.RecurrentlyFailing) > 0 {
		fmt.Fprintf(&b, "\n  Recurrently Failing Areas (%d):\n", len(d.RecurrentlyFailing))
		for _, a := range d.RecurrentlyFailing {
			fmt.Fprintf(&b, "    %s — %d/%d failing, avg %.3f\n", a.Path, a.Failing, a.Total, a.AverageScore)
		}
	}

	return b.String()
}
