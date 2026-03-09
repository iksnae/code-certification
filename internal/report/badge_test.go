package report_test

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/report"
)

func TestGenerateBadge(t *testing.T) {
	records := []domain.CertificationRecord{
		makeCardRecord("go", "a.go", "A", domain.StatusCertified, 0.90),
		makeCardRecord("go", "b.go", "B", domain.StatusCertified, 0.85),
	}
	c := report.GenerateCard(records, "test/repo", "abc", time.Now())
	badge := report.GenerateBadge(c)

	if badge.SchemaVersion != 1 {
		t.Errorf("schemaVersion = %d, want 1", badge.SchemaVersion)
	}
	if badge.Label != "certification" {
		t.Errorf("label = %s, want certification", badge.Label)
	}
	if !strings.Contains(badge.Message, "2 units") {
		t.Errorf("message = %s, should contain unit count", badge.Message)
	}
	// avg 0.875 → B+ → steel blue
	if badge.Color != "4A6B82" {
		t.Errorf("color = %s, want 4A6B82 for grade B+", badge.Color)
	}
}

func TestBadgeColors(t *testing.T) {
	tests := []struct {
		score float64
		color string
	}{
		{0.98, "2E8B57"}, // A  — certified green
		{0.91, "3DA06A"}, // A- — certified green (lighter)
		{0.88, "4A6B82"}, // B+ — steel blue
		{0.83, "4A6B82"}, // B  — steel blue
		{0.70, "E0A100"}, // C  — observations amber
		{0.60, "F59E0B"}, // D  — probationary warning
		{0.30, "DC2626"}, // F  — decertified red
	}
	for _, tt := range tests {
		records := []domain.CertificationRecord{
			makeCardRecord("go", "a.go", "A", domain.StatusCertified, tt.score),
		}
		c := report.GenerateCard(records, "", "", time.Now())
		badge := report.GenerateBadge(c)
		if badge.Color != tt.color {
			t.Errorf("score=%.2f grade=%s → color=%s, want %s", tt.score, c.OverallGrade, badge.Color, tt.color)
		}
	}
}

func TestFormatBadgeJSON(t *testing.T) {
	badge := report.Badge{
		SchemaVersion: 1,
		Label:         "certification",
		Message:       "B · 100% · 5 units",
		Color:         "blue",
	}
	data, err := report.FormatBadgeJSON(badge)
	if err != nil {
		t.Fatal(err)
	}
	var parsed map[string]interface{}
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if parsed["schemaVersion"].(float64) != 1 {
		t.Error("schemaVersion not preserved")
	}
}

func TestBadgeMarkdown(t *testing.T) {
	md := report.BadgeMarkdown("iksnae/code-certification", "main")

	if !strings.Contains(md, "img.shields.io/endpoint") {
		t.Error("should use shields.io endpoint")
	}
	if !strings.Contains(md, "raw.githubusercontent.com") {
		t.Error("should reference raw content URL")
	}
	if !strings.Contains(md, "badge.json") {
		t.Error("should reference badge.json")
	}
	if !strings.Contains(md, "REPORT_CARD.md") {
		t.Error("badge should link to report card")
	}
	if !strings.HasPrefix(md, "[![") {
		t.Error("should be a markdown image link")
	}
}

func TestBadge_Empty(t *testing.T) {
	c := report.GenerateCard(nil, "", "", time.Now())
	badge := report.GenerateBadge(c)
	if badge.Message != "no data" {
		t.Errorf("message = %s, want 'no data'", badge.Message)
	}
	if badge.Color != "9CA3AF" {
		t.Errorf("color = %s, want 9CA3AF for N/A", badge.Color)
	}
}
