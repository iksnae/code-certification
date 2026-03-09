package github_test

import (
	"strings"
	"testing"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
	gh "github.com/iksnae/code-certification/internal/github"
)

func makeTestRecord(path, symbol string, status domain.Status, score float64) domain.CertificationRecord {
	return domain.CertificationRecord{
		UnitID:    domain.NewUnitID("go", path, symbol),
		UnitPath:  path,
		Status:    status,
		Score:     score,
		Grade:     domain.GradeFromScore(score),
		ExpiresAt: time.Now().Add(90 * 24 * time.Hour),
	}
}

func TestFormatPRComment(t *testing.T) {
	records := []domain.CertificationRecord{
		makeTestRecord("main.go", "main", domain.StatusCertified, 0.9),
		makeTestRecord("main.go", "helper", domain.StatusCertifiedWithObservations, 0.75),
		makeTestRecord("bad.go", "broken", domain.StatusDecertified, 0.3),
	}

	comment := gh.FormatPRComment(records, false)

	if !strings.Contains(comment, "Certification") {
		t.Error("should contain Certification header")
	}
	if !strings.Contains(comment, "2/3") || !strings.Contains(comment, "passing") {
		t.Error("should show passing count")
	}
	if !strings.Contains(comment, "decertified") {
		t.Error("should mention decertified units")
	}
}

func TestFormatPRComment_AllPass(t *testing.T) {
	records := []domain.CertificationRecord{
		makeTestRecord("main.go", "main", domain.StatusCertified, 0.95),
	}

	comment := gh.FormatPRComment(records, false)
	if !strings.Contains(comment, "✅") || !strings.Contains(comment, "1/1") {
		t.Error("all-pass should show check mark")
	}
}

func TestFormatPRComment_Enforcing(t *testing.T) {
	records := []domain.CertificationRecord{
		makeTestRecord("bad.go", "broken", domain.StatusDecertified, 0.3),
	}

	comment := gh.FormatPRComment(records, true)
	if !strings.Contains(comment, "❌") || !strings.Contains(comment, "BLOCKED") {
		t.Error("enforcing mode with failures should show blocked")
	}
}

func TestBuildGHCommand(t *testing.T) {
	cmd := gh.BuildPRCommentCommand("123", "test comment body")

	if len(cmd) < 4 {
		t.Fatalf("command too short: %v", cmd)
	}
	if cmd[0] != "gh" {
		t.Errorf("cmd[0] = %q, want gh", cmd[0])
	}
	if cmd[1] != "pr" {
		t.Errorf("cmd[1] = %q, want pr", cmd[1])
	}
}
