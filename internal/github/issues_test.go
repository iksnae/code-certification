package github_test

import (
	"strings"
	"testing"

	"github.com/code-certification/certify/internal/domain"
	gh "github.com/code-certification/certify/internal/github"
)

func TestFormatIssueBody(t *testing.T) {
	rec := makeTestRecord("bad.go", "broken", domain.StatusDecertified, 0.3)
	rec.Observations = []string{"lint: 5 errors", "tests: 2 failures"}

	body := gh.FormatIssueBody(rec)

	if !strings.Contains(body, "bad.go") {
		t.Error("should contain file path")
	}
	if !strings.Contains(body, "decertified") {
		t.Error("should contain status")
	}
	if !strings.Contains(body, "lint: 5 errors") {
		t.Error("should contain observations")
	}
}

func TestFormatIssueTitle(t *testing.T) {
	rec := makeTestRecord("bad.go", "broken", domain.StatusDecertified, 0.3)
	title := gh.FormatIssueTitle(rec)

	if !strings.Contains(title, "bad.go") {
		t.Error("title should contain path")
	}
	if !strings.Contains(title, "certification") || !strings.Contains(strings.ToLower(title), "decertified") {
		t.Error("title should mention certification status")
	}
}

func TestBuildIssueCreateCommand(t *testing.T) {
	cmd := gh.BuildIssueCreateCommand("Test Title", "Test Body", []string{"certification", "tech-debt"})

	if cmd[0] != "gh" || cmd[1] != "issue" || cmd[2] != "create" {
		t.Errorf("unexpected command prefix: %v", cmd[:3])
	}

	joined := strings.Join(cmd, " ")
	if !strings.Contains(joined, "--title") {
		t.Error("should have --title flag")
	}
	if !strings.Contains(joined, "--label") {
		t.Error("should have --label flag")
	}
}

func TestBuildIssueCloseCommand(t *testing.T) {
	cmd := gh.BuildIssueCloseCommand("42")
	if cmd[0] != "gh" || cmd[1] != "issue" || cmd[2] != "close" {
		t.Errorf("unexpected command: %v", cmd)
	}
	if cmd[3] != "42" {
		t.Errorf("issue number = %q, want 42", cmd[3])
	}
}
