package report_test

import (
	"strings"
	"testing"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/report"
)

func makeCardRecord(lang, path, symbol string, status domain.Status, score float64) domain.CertificationRecord {
	now := time.Now()
	return domain.CertificationRecord{
		UnitID:      domain.NewUnitID(lang, path, symbol),
		UnitType:    domain.UnitTypeFunction,
		UnitPath:    path,
		Status:      status,
		Score:       score,
		Grade:       domain.GradeFromScore(score),
		CertifiedAt: now,
		ExpiresAt:   now.Add(90 * 24 * time.Hour),
	}
}

func TestGenerateCard_Basic(t *testing.T) {
	records := []domain.CertificationRecord{
		makeCardRecord("go", "a.go", "A", domain.StatusCertified, 0.95),
		makeCardRecord("go", "b.go", "B", domain.StatusCertified, 0.85),
		makeCardRecord("ts", "c.ts", "C", domain.StatusCertified, 0.80),
	}

	c := report.GenerateCard(records, "test/repo", "abc123", time.Now())

	if c.TotalUnits != 3 {
		t.Errorf("total = %d, want 3", c.TotalUnits)
	}
	if c.Passing != 3 {
		t.Errorf("passing = %d, want 3", c.Passing)
	}
	if c.Failing != 0 {
		t.Errorf("failing = %d, want 0", c.Failing)
	}
	if c.OverallGrade == "" {
		t.Error("grade should not be empty")
	}
	if c.Repository != "test/repo" {
		t.Errorf("repo = %s, want test/repo", c.Repository)
	}
	if len(c.Languages) != 2 {
		t.Errorf("languages = %d, want 2", len(c.Languages))
	}
}

func TestGenerateCard_WithFailures(t *testing.T) {
	records := []domain.CertificationRecord{
		makeCardRecord("go", "good.go", "A", domain.StatusCertified, 0.90),
		makeCardRecord("go", "bad.go", "B", domain.StatusDecertified, 0.30),
	}
	records[1].Observations = []string{"lint_errors: 5 exceeds threshold 0"}

	c := report.GenerateCard(records, "", "", time.Now())

	if c.Failing != 1 {
		t.Errorf("failing = %d, want 1", c.Failing)
	}
	if len(c.TopIssues) < 1 {
		t.Fatal("should have at least 1 top issue")
	}
}

func TestGenerateCard_Empty(t *testing.T) {
	c := report.GenerateCard(nil, "", "", time.Now())
	if c.OverallGrade != "N/A" {
		t.Errorf("empty grade = %s, want N/A", c.OverallGrade)
	}
}

func TestGenerateCard_GradeDistribution(t *testing.T) {
	records := []domain.CertificationRecord{
		makeCardRecord("go", "a.go", "A", domain.StatusCertified, 0.95), // A-
		makeCardRecord("go", "b.go", "B", domain.StatusCertified, 0.85), // B
		makeCardRecord("go", "c.go", "C", domain.StatusCertified, 0.85), // B
		makeCardRecord("go", "d.go", "D", domain.StatusCertified, 0.75), // C
	}

	c := report.GenerateCard(records, "", "", time.Now())

	totalGraded := 0
	for _, count := range c.GradeDistribution {
		totalGraded += count
	}
	if totalGraded != 4 {
		t.Errorf("total graded = %d, want 4", totalGraded)
	}
}

func TestFormatCardText(t *testing.T) {
	records := []domain.CertificationRecord{
		makeCardRecord("go", "a.go", "A", domain.StatusCertified, 0.90),
		makeCardRecord("ts", "b.ts", "B", domain.StatusCertified, 0.85),
	}

	c := report.GenerateCard(records, "iksnae/code-certification", "abc1234", time.Now())
	text := report.FormatCardText(c)

	if !strings.Contains(text, "REPORT CARD") {
		t.Error("should contain 'REPORT CARD'")
	}
	if !strings.Contains(text, "iksnae/code-certification") {
		t.Error("should contain repository name")
	}
	if !strings.Contains(text, "abc1234") {
		t.Error("should contain commit SHA")
	}
	if !strings.Contains(text, "Grade Distribution") {
		t.Error("should contain grade distribution")
	}
	if !strings.Contains(text, "By Language") {
		t.Error("should contain language breakdown")
	}
}

func TestFormatCardMarkdown(t *testing.T) {
	records := []domain.CertificationRecord{
		makeCardRecord("go", "a.go", "A", domain.StatusCertified, 0.90),
		makeCardRecord("go", "bad.go", "B", domain.StatusDecertified, 0.30),
	}

	c := report.GenerateCard(records, "test/repo", "def456", time.Now())
	md := report.FormatCardMarkdown(c)

	if !strings.Contains(md, "# ") {
		t.Error("should contain markdown heading")
	}
	if !strings.Contains(md, "| Metric |") {
		t.Error("should contain metrics table")
	}
	if !strings.Contains(md, "Grade Distribution") {
		t.Error("should contain grade distribution")
	}
	if !strings.Contains(md, "certify") {
		t.Error("should contain attribution link")
	}
}
