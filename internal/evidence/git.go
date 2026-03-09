package evidence

import (
	"fmt"
	"strings"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
)

// GitStats holds parsed git history data for a file.
type GitStats struct {
	CommitCount int `json:"commit_count"`
	AuthorCount int `json:"author_count"`
	AgeDays     int `json:"age_days"`
}

// ChurnRate returns commits per day. Returns 0 if age is zero.
func (s GitStats) ChurnRate() float64 {
	if s.AgeDays == 0 {
		return 0
	}
	return float64(s.CommitCount) / float64(s.AgeDays)
}

// ToEvidence converts GitStats to a domain.Evidence.
func (s GitStats) ToEvidence() domain.Evidence {
	return domain.Evidence{
		Kind:       domain.EvidenceKindGitHistory,
		Source:     "git",
		Passed:     true, // Git history is informational, always "passes"
		Summary:    fmt.Sprintf("%d commits by %d authors over %d days", s.CommitCount, s.AuthorCount, s.AgeDays),
		Details:    s,
		Timestamp:  time.Now(),
		Confidence: 1.0, // Git data is deterministic
	}
}

// ParseGitLog parses tab-separated git log output (hash\tauthor\tdate).
func ParseGitLog(output string) GitStats {
	lines := strings.Split(strings.TrimSpace(output), "\n")
	if len(lines) == 1 && lines[0] == "" {
		return GitStats{}
	}

	authors := make(map[string]bool)
	for _, line := range lines {
		parts := strings.Split(line, "\t")
		if len(parts) >= 2 {
			authors[parts[1]] = true
		}
	}

	return GitStats{
		CommitCount: len(lines),
		AuthorCount: len(authors),
	}
}
