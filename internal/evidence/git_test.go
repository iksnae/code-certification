package evidence_test

import (
	"testing"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/evidence"
)

func TestGitCollector_ParseLog(t *testing.T) {
	// Simulate git log output (tab-separated: hash, author, date)
	log := `abc1234	alice	2025-01-15
def5678	bob	2025-02-10
ghi9012	alice	2025-03-05
`
	stats := evidence.ParseGitLog(log)

	if stats.CommitCount != 3 {
		t.Errorf("CommitCount = %d, want 3", stats.CommitCount)
	}
	if stats.AuthorCount != 2 {
		t.Errorf("AuthorCount = %d, want 2", stats.AuthorCount)
	}
}

func TestGitCollector_EmptyLog(t *testing.T) {
	stats := evidence.ParseGitLog("")
	if stats.CommitCount != 0 {
		t.Errorf("CommitCount = %d, want 0", stats.CommitCount)
	}
}

func TestGitCollector_ChurnRate(t *testing.T) {
	stats := evidence.GitStats{
		CommitCount: 30,
		AgeDays:     90,
	}
	rate := stats.ChurnRate()
	// 30 commits / 90 days ≈ 0.333
	if rate < 0.33 || rate > 0.34 {
		t.Errorf("ChurnRate() = %f, want ~0.333", rate)
	}
}

func TestGitCollector_ChurnRate_ZeroAge(t *testing.T) {
	stats := evidence.GitStats{
		CommitCount: 5,
		AgeDays:     0,
	}
	rate := stats.ChurnRate()
	if rate != 0 {
		t.Errorf("ChurnRate(age=0) = %f, want 0", rate)
	}
}

func TestGitCollector_ToEvidence(t *testing.T) {
	stats := evidence.GitStats{
		CommitCount: 10,
		AuthorCount: 3,
		AgeDays:     60,
	}

	ev := stats.ToEvidence()
	if ev.Kind != domain.EvidenceKindGitHistory {
		t.Errorf("Kind = %v, want git_history", ev.Kind)
	}
	if ev.Source != "git" {
		t.Errorf("Source = %q, want git", ev.Source)
	}
	if !ev.Passed {
		t.Error("git evidence should always pass (informational)")
	}
}
