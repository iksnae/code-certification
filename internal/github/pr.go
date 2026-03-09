package github

import (
	"fmt"
	"strings"

	"github.com/iksnae/code-certification/internal/domain"
)

// FormatPRComment generates a Markdown PR comment from certification records.
func FormatPRComment(records []domain.CertificationRecord, enforcing bool) string {
	var b strings.Builder
	total := len(records)
	var passing, failing int

	for _, r := range records {
		if r.Status.IsPassing() {
			passing++
		} else {
			failing++
		}
	}

	// Header
	if failing > 0 && enforcing {
		b.WriteString("## ❌ Certification BLOCKED\n\n")
	} else if failing > 0 {
		b.WriteString("## ⚠️ Certification Report\n\n")
	} else {
		b.WriteString("## ✅ Certification Report\n\n")
	}

	fmt.Fprintf(&b, "**%d/%d** units passing certification\n\n", passing, total)

	// Summary table
	if failing > 0 {
		b.WriteString("### Units needing attention\n\n")
		b.WriteString("| Unit | Status | Grade | Score |\n")
		b.WriteString("|------|--------|-------|-------|\n")
		for _, r := range records {
			if !r.Status.IsPassing() {
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
func ComputeTrustDelta(oldRecords, newRecords []domain.CertificationRecord) TrustDelta {
	oldMap := make(map[string]domain.CertificationRecord)
	for _, r := range oldRecords {
		oldMap[r.UnitID.String()] = r
	}

	var d TrustDelta
	var oldTotal, newTotal float64
	for _, nr := range newRecords {
		newTotal += nr.Score
		if or, ok := oldMap[nr.UnitID.String()]; ok {
			oldTotal += or.Score
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
		} else {
			if nr.Status.IsPassing() {
				d.NewlyCertified++
			}
		}
	}
	if len(oldRecords) > 0 && len(newRecords) > 0 {
		d.AverageScoreDelta = (newTotal / float64(len(newRecords))) - (oldTotal / float64(len(oldRecords)))
	}
	return d
}

// BuildPRCommentCommand constructs a `gh pr comment` command.
func BuildPRCommentCommand(prNumber, body string) []string {
	return []string{"gh", "pr", "comment", prNumber, "--body", body}
}
