// Package report generates certification reports from records.
package report

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/iksnae/code-certification/internal/analysis"
	"github.com/iksnae/code-certification/internal/domain"
)

// Card is a concise, shareable certification report card.
type Card struct {
	// Header
	Repository  string `json:"repository"`
	GeneratedAt string `json:"generated_at"`
	CommitSHA   string `json:"commit_sha,omitempty"`

	// Overall grade
	OverallGrade string  `json:"overall_grade"`
	OverallScore float64 `json:"overall_score"`
	PassRate     float64 `json:"pass_rate"`

	// Counts
	TotalUnits   int `json:"total_units"`
	Passing      int `json:"passing"`
	Failing      int `json:"failing"`
	Expired      int `json:"expired"`
	Observations int `json:"observations"`

	// Grade distribution
	GradeDistribution map[string]int `json:"grade_distribution"`

	// By-language summary
	Languages []LanguageDetail `json:"languages"`

	// Top issues (up to 10)
	TopIssues []IssueCard `json:"top_issues,omitempty"`

	// Package summaries (populated for report tree navigation)
	Packages []PackageSummary `json:"packages,omitempty"`

	// Unsupported-unit counts. UnsupportedCount units carry no verdict: the
	// engine has no analyzer for their language, so nothing about them was
	// measured. AnalyzableUnits is TotalUnits minus that count and is the
	// denominator of PassRate — an unassessed unit belongs in neither the
	// numerator nor the denominator of a rate that claims to summarise verdicts.
	UnsupportedCount     int      `json:"unsupported_count"`
	UnsupportedLanguages []string `json:"unsupported_languages"`
	AnalyzableUnits      int      `json:"analyzable_units"`
}

// PassRateKnown reports whether PassRate is a measurement rather than a
// placeholder. With no analyzable units the ratio is 0/0 — undefined, which is
// neither 0% nor 100% — so every renderer must gate on this and print a
// non-numeric marker rather than a percentage that implies an assessment ran.
func (c Card) PassRateKnown() bool { return c.AnalyzableUnits > 0 }

// FormatRate renders a pass rate for display. It is the single place that
// decides how an undefined rate looks, so the marker cannot drift between the
// report card, the full report, the site, the report tree and the workspace
// rollup — surfaces that a reader compares against each other.
func FormatRate(known bool, rate float64, decimals int) string {
	if !known {
		return "n/a"
	}
	return fmt.Sprintf("%.*f%%", decimals, rate*100)
}

// formatPassRate renders a Card's pass rate, honouring the undefined case.
func formatPassRate(c Card) string {
	return FormatRate(c.PassRateKnown(), c.PassRate, 1)
}

// distributionGrades is the render order for the grade distribution. N/A is a
// real bucket, not a gap: it is what an unassessed unit records, and omitting
// it printed a Grade Distribution header with no rows over an all-unsupported
// repo's units.
var distributionGrades = []string{"A", "A-", "B+", "B", "C", "D", "F", "N/A"}

// distributionBarWidth is the column budget for the text card's histogram bar.
const distributionBarWidth = 36

// IssueCard describes a single unit needing attention.
type IssueCard struct {
	UnitID string  `json:"unit_id"`
	Grade  string  `json:"grade"`
	Score  float64 `json:"score"`
	Reason string  `json:"reason"`
}

// GenerateCard creates a report card from certification records.
func GenerateCard(records []domain.CertificationRecord, repo, commit string, now time.Time) Card {
	c := Card{
		Repository:        repo,
		GeneratedAt:       now.Format(time.RFC3339),
		CommitSHA:         commit,
		GradeDistribution: make(map[string]int),
	}

	if len(records) == 0 {
		c.OverallGrade = "N/A"
		return c
	}

	c.TotalUnits = len(records)
	var totalScore float64
	unsupportedLangs := make(map[string]bool)

	for _, r := range records {
		totalScore += r.Score
		c.GradeDistribution[r.Grade.String()]++

		// Unsupported units are unassessed, not passed and not failed. They
		// carry StatusExempt, whose IsPassing() is true, so the switch below
		// would otherwise count code the engine never opened as certified.
		// Reading the recorded flag rather than re-deriving the language
		// verdict here keeps this surface in agreement with the run tally and
		// the PR comment, which branch on the same field.
		if r.Unsupported {
			c.UnsupportedCount++
			unsupportedLangs[r.UnitID.Language()] = true
			continue
		}

		switch {
		case r.Status == domain.StatusExpired:
			c.Expired++
		case r.Status.IsPassing() && r.Status == domain.StatusCertifiedWithObservations:
			c.Passing++
			c.Observations++
		case r.Status.IsPassing():
			c.Passing++
		default:
			c.Failing++
		}
	}

	c.AnalyzableUnits = c.TotalUnits - c.UnsupportedCount
	for lang := range unsupportedLangs {
		c.UnsupportedLanguages = append(c.UnsupportedLanguages, lang)
	}
	sort.Strings(c.UnsupportedLanguages)

	c.OverallScore = totalScore / float64(c.TotalUnits)
	c.OverallGrade = domain.GradeFromScore(c.OverallScore).String()
	if c.PassRateKnown() {
		c.PassRate = float64(c.Passing) / float64(c.AnalyzableUnits)
	}
	c.Languages = buildLanguageDetail(records)
	c.TopIssues = buildTopIssues(records)

	return c
}

func buildTopIssues(records []domain.CertificationRecord) []IssueCard {
	// Collect non-passing + lowest-scoring units
	sorted := make([]domain.CertificationRecord, len(records))
	copy(sorted, records)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Score < sorted[j].Score
	})

	var issues []IssueCard
	for _, r := range sorted {
		if len(issues) >= 10 {
			break
		}
		if r.Status == domain.StatusExempt {
			continue
		}
		reason := "lowest score"
		if !r.Status.IsPassing() {
			reason = r.Status.String()
		} else if len(r.Observations) > 0 {
			reason = r.Observations[0]
		}
		issues = append(issues, IssueCard{
			UnitID: r.UnitID.String(),
			Grade:  r.Grade.String(),
			Score:  r.Score,
			Reason: reason,
		})
	}
	return issues
}

// FormatCardText renders the report card as a human-readable text block.
func FormatCardText(c Card) string {
	var b strings.Builder

	// Header
	b.WriteString("╔══════════════════════════════════════════════════════════════╗\n")
	b.WriteString("║              CERTIFY — REPORT CARD                 ║\n")
	b.WriteString("╠══════════════════════════════════════════════════════════════╣\n")

	if c.Repository != "" {
		fmt.Fprintf(&b, "║  Repository:  %-45s║\n", c.Repository)
	}
	if c.CommitSHA != "" {
		fmt.Fprintf(&b, "║  Commit:      %-45s║\n", c.CommitSHA)
	}
	fmt.Fprintf(&b, "║  Generated:   %-45s║\n", c.GeneratedAt[:19])
	b.WriteString("╠══════════════════════════════════════════════════════════════╣\n")

	// Overall grade — big and prominent
	emoji := gradeEmoji(c.OverallGrade)
	fmt.Fprintf(&b, "║                                                              ║\n")
	fmt.Fprintf(&b, "║       Overall Grade:  %s %-5s    Score: %-6.1f%%            ║\n",
		emoji, c.OverallGrade, c.OverallScore*100)
	fmt.Fprintf(&b, "║                                                              ║\n")
	b.WriteString("╠══════════════════════════════════════════════════════════════╣\n")

	// Stats
	fmt.Fprintf(&b, "║  Total Units:     %-6d    Pass Rate:  %7s              ║\n", c.TotalUnits, formatPassRate(c))
	fmt.Fprintf(&b, "║  Passing:         %-6d    Failing:    %6d               ║\n", c.Passing, c.Failing)
	fmt.Fprintf(&b, "║  Observations:    %-6d    Expired:    %6d               ║\n", c.Observations, c.Expired)
	if c.UnsupportedCount > 0 {
		fmt.Fprintf(&b, "║  Not Assessed:    %-6d    Analyzable: %6d               ║\n", c.UnsupportedCount, c.AnalyzableUnits)
	}
	b.WriteString("╠══════════════════════════════════════════════════════════════╣\n")

	// Grade distribution
	b.WriteString("║  Grade Distribution:                                         ║\n")
	grades := distributionGrades
	for _, g := range grades {
		count := c.GradeDistribution[g]
		if count == 0 {
			continue
		}
		pct := float64(count) / float64(c.TotalUnits) * 100
		barLen := int(pct / 2)
		if barLen < 1 && count > 0 {
			barLen = 1
		}
		// The bar shares a 36-column field with the box border; a grade holding
		// 100% of the units would otherwise run straight through the right edge.
		if barLen > distributionBarWidth {
			barLen = distributionBarWidth
		}
		bar := strings.Repeat("█", barLen)
		fmt.Fprintf(&b, "║    %2s: %4d (%5.1f%%) %-36s║\n", g, count, pct, bar)
	}
	b.WriteString("╠══════════════════════════════════════════════════════════════╣\n")

	// Languages
	if len(c.Languages) > 0 {
		b.WriteString("║  By Language:                                                ║\n")
		for _, l := range c.Languages {
			fmt.Fprintf(&b, "║    %-12s %4d units   %s %-5s  (%.1f%%)              ║\n",
				l.Name, l.Units, gradeEmoji(l.Grade), l.Grade, l.AverageScore*100)
		}
		b.WriteString("╠══════════════════════════════════════════════════════════════╣\n")
	}

	// Top issues
	if len(c.TopIssues) > 0 && c.Failing > 0 {
		b.WriteString("║  Top Issues:                                                 ║\n")
		for i, issue := range c.TopIssues {
			if i >= 5 {
				break
			}
			id := issue.UnitID
			if len(id) > 40 {
				id = "..." + id[len(id)-37:]
			}
			fmt.Fprintf(&b, "║    %s %-5s %-42s║\n", gradeEmoji(issue.Grade), issue.Grade, id)
		}
		b.WriteString("╠══════════════════════════════════════════════════════════════╣\n")
	}

	b.WriteString("╚══════════════════════════════════════════════════════════════╝\n")

	return b.String()
}

// FormatCardMarkdown renders the report card as a GitHub-friendly markdown block.
func FormatCardMarkdown(c Card) string {
	var b strings.Builder

	emoji := gradeEmoji(c.OverallGrade)

	fmt.Fprintf(&b, "# %s Certify — Report Card\n\n", emoji)

	if c.Repository != "" {
		fmt.Fprintf(&b, "**Repository:** `%s`\n", c.Repository)
	}
	if c.CommitSHA != "" {
		fmt.Fprintf(&b, "**Commit:** `%s`\n", c.CommitSHA)
	}
	fmt.Fprintf(&b, "**Generated:** %s\n\n", c.GeneratedAt[:19])

	fmt.Fprintf(&b, "## %s Overall: %s (%.1f%%)\n\n", emoji, c.OverallGrade, c.OverallScore*100)

	fmt.Fprintf(&b, "| Metric | Value |\n")
	fmt.Fprintf(&b, "|--------|-------|\n")
	fmt.Fprintf(&b, "| Total Units | %d |\n", c.TotalUnits)
	fmt.Fprintf(&b, "| Passing | %d |\n", c.Passing)
	fmt.Fprintf(&b, "| Failing | %d |\n", c.Failing)
	if c.UnsupportedCount > 0 {
		fmt.Fprintf(&b, "| Not Assessed | %d |\n", c.UnsupportedCount)
		fmt.Fprintf(&b, "| Analyzable Units | %d |\n", c.AnalyzableUnits)
	}
	fmt.Fprintf(&b, "| Pass Rate | %s |\n", formatPassRate(c))
	fmt.Fprintf(&b, "| Observations | %d |\n", c.Observations)
	fmt.Fprintf(&b, "| Expired | %d |\n\n", c.Expired)

	// Grade distribution
	b.WriteString("### Grade Distribution\n\n")
	b.WriteString("| Grade | Count | % |\n")
	b.WriteString("|-------|-------|---|\n")
	grades := distributionGrades
	for _, g := range grades {
		count := c.GradeDistribution[g]
		if count == 0 {
			continue
		}
		pct := float64(count) / float64(c.TotalUnits) * 100
		fmt.Fprintf(&b, "| %s | %d | %.1f%% |\n", g, count, pct)
	}
	b.WriteString("\n")

	// Languages
	if len(c.Languages) > 0 {
		b.WriteString("### By Language\n\n")
		b.WriteString("| Language | Units | Grade | Score |\n")
		b.WriteString("|----------|-------|-------|-------|\n")
		for _, l := range c.Languages {
			fmt.Fprintf(&b, "| %s | %d | %s %s | %.1f%% |\n",
				l.Name, l.Units, gradeEmoji(l.Grade), l.Grade, l.AverageScore*100)
		}
		b.WriteString("\n")
	}

	// Packages (with links into the report tree)
	if len(c.Packages) > 0 {
		b.WriteString("### Packages\n\n")
		b.WriteString("| Package | Units | Grade | Score |\n")
		b.WriteString("|---------|------:|:-----:|------:|\n")
		for _, p := range c.Packages {
			fmt.Fprintf(&b, "| [%s](reports/%s/index.md) | %d | %s %s | %.1f%% |\n",
				p.Path, p.Path, p.Units, gradeEmoji(p.Grade), p.Grade, p.AvgScore*100)
		}
		b.WriteString("\n")
	}

	// Top issues
	if len(c.TopIssues) > 0 && c.Failing > 0 {
		b.WriteString("### Top Issues\n\n")
		b.WriteString("| Unit | Grade | Score | Issue |\n")
		b.WriteString("|------|-------|-------|-------|\n")
		for i, issue := range c.TopIssues {
			if i >= 10 {
				break
			}
			fmt.Fprintf(&b, "| `%s` | %s | %.1f%% | %s |\n",
				issue.UnitID, issue.Grade, issue.Score*100, issue.Reason)
		}
		b.WriteString("\n")
	}

	b.WriteString("---\n\n")
	b.WriteString("**Drill down:** `certify report --format full` · `certify report --site`  \n")
	b.WriteString("*Generated by [Certify](https://github.com/iksnae/code-certification)*\n")

	return b.String()
}

// partitionByLanguage splits records into supported and unsupported slices
// based on whether the language has a registered analyzer.
func partitionByLanguage(records []domain.CertificationRecord) (supported, unsupported []domain.CertificationRecord) {
	for _, r := range records {
		if analysis.ForLanguage(r.UnitID.Language()) == nil {
			unsupported = append(unsupported, r)
		} else {
			supported = append(supported, r)
		}
	}
	return
}

// gradeEmoji maps grades to brand-consistent status emoji.
// 🟢 Certified (A–B), 🟡 Observations (C), 🟠 Probationary (D), 🔴 Decertified (F), ⚪ Expired.
func gradeEmoji(grade string) string {
	switch grade {
	case "A", "A-", "B+", "B":
		return "🟢"
	case "C":
		return "🟡"
	case "D":
		return "🟠"
	case "F":
		return "🔴"
	default:
		return "⚪"
	}
}
