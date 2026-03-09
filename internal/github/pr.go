package github

import (
	"fmt"
	"strings"

	"github.com/code-certification/certify/internal/domain"
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

// BuildPRCommentCommand constructs a `gh pr comment` command.
func BuildPRCommentCommand(prNumber, body string) []string {
	return []string{"gh", "pr", "comment", prNumber, "--body", body}
}
