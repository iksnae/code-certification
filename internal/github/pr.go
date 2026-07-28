package github

import (
	"fmt"
	"strings"

	"github.com/iksnae/code-certification/internal/domain"
)

// FormatPRComment generates a Markdown PR comment from certification records.
//
// Unsupported units are excluded from both the numerator and the denominator and
// reported on their own line. They carry StatusExempt, for which IsPassing() is
// true, so counting them would tell the client their code passed certification
// when the engine never analysed it. Nor may they be counted as failing — that is
// the defect this exclusion exists to correct. They are a third state: unassessed.
func FormatPRComment(records []domain.CertificationRecord, enforcing bool) string {
	var b strings.Builder
	var passing, failing, unsupported int

	for _, r := range records {
		switch {
		case r.Unsupported:
			unsupported++
		case r.Status.IsPassing():
			passing++
		default:
			failing++
		}
	}
	assessed := passing + failing

	// Header. A run in which nothing was assessed has no certification to
	// report, so it never earns the green check.
	switch {
	case failing > 0 && enforcing:
		b.WriteString("## ❌ Certification BLOCKED\n\n")
	case failing > 0:
		b.WriteString("## ⚠️ Certification Report\n\n")
	case assessed == 0:
		b.WriteString("## ⚠️ Certification Report\n\n")
	default:
		b.WriteString("## ✅ Certification Report\n\n")
	}

	if assessed == 0 {
		b.WriteString("**No units were assessed.**\n\n")
	} else {
		fmt.Fprintf(&b, "**%d/%d** units passing certification\n\n", passing, assessed)
	}

	if unsupported > 0 {
		fmt.Fprintf(&b, "> **%d** unit(s) were not assessed — unsupported language. "+
			"They are counted as neither passing nor failing.\n\n", unsupported)
	}

	// Summary table
	if failing > 0 {
		b.WriteString("### Units needing attention\n\n")
		b.WriteString("| Unit | Status | Grade | Score |\n")
		b.WriteString("|------|--------|-------|-------|\n")
		for _, r := range records {
			if !r.Unsupported && !r.Status.IsPassing() {
				fmt.Fprintf(&b, "| `%s` | %s | %s | %.2f |\n",
					r.UnitID, r.Status, r.Grade, r.Score)
			}
		}
		b.WriteString("\n")
	}

	if enforcing && failing > 0 {
		fmt.Fprintf(&b, "> **%d units failed certification in enforcing mode. Merge is blocked.**\n", failing)
	}

	return b.String()
}

// TrustDelta computes the change in certification metrics between old and new records.
type TrustDelta struct {
	NewlyCertified    int
	NewlyDecertified  int
	ScoreImproved     int
	ScoreDegraded     int
	AverageScoreDelta float64
}

// ComputeTrustDelta compares old vs new record sets.
//
// Unsupported units are excluded from every counter on both sides. A unit that
// moves from a verdict to unassessed has not been certified or decertified — it
// has stopped being measured, which is not a movement in trust. Without this
// exclusion the first run after the unsupported-language fix reports every such
// unit as newly certified, because it moves from StatusDecertified (IsPassing
// false) to StatusExempt (IsPassing true). Their stored score of 0 is likewise
// the absence of a measurement, not a measurement of zero, so it must not enter
// the average.
func ComputeTrustDelta(oldRecords, newRecords []domain.CertificationRecord) TrustDelta {
	oldMap := make(map[string]domain.CertificationRecord)
	for _, r := range oldRecords {
		if r.Unsupported {
			continue
		}
		oldMap[r.UnitID.String()] = r
	}

	var d TrustDelta
	var oldTotal, newTotal float64
	var oldCount, newCount int
	for _, nr := range newRecords {
		if nr.Unsupported {
			continue
		}
		newTotal += nr.Score
		newCount++

		or, ok := oldMap[nr.UnitID.String()]
		if !ok {
			// No prior verdict — either a genuinely new unit, or one that was
			// previously unassessed. Becoming certified from nothing counts.
			if nr.Status.IsPassing() {
				d.NewlyCertified++
			}
			continue
		}
		oldTotal += or.Score
		oldCount++

		if !or.Status.IsPassing() && nr.Status.IsPassing() {
			d.NewlyCertified++
		}
		if or.Status.IsPassing() && !nr.Status.IsPassing() {
			d.NewlyDecertified++
		}
		if nr.Score > or.Score+0.01 {
			d.ScoreImproved++
		}
		if nr.Score < or.Score-0.01 {
			d.ScoreDegraded++
		}
	}
	if oldCount > 0 && newCount > 0 {
		d.AverageScoreDelta = (newTotal / float64(newCount)) - (oldTotal / float64(oldCount))
	}
	return d
}

// BuildPRCommentCommand constructs a `gh pr comment` command.
func BuildPRCommentCommand(prNumber, body string) []string {
	return []string{"gh", "pr", "comment", prNumber, "--body", body}
}
