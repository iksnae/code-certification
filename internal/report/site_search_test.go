package report_test

import (
	"strings"
	"testing"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/report"
)

func TestBuildSearchIndex(t *testing.T) {
	records := []domain.CertificationRecord{
		makeSiteRecord("go", "internal/engine/scorer.go", "Score", domain.UnitTypeFunction, domain.StatusCertified, 0.92),
		makeSiteRecord("go", "internal/engine/pipeline.go", "CertifyUnit", domain.UnitTypeFunction, domain.StatusCertified, 0.85),
		makeSiteRecord("ts", "src/parser.ts", "parse", domain.UnitTypeFunction, domain.StatusProbationary, 0.65),
	}

	fr := report.GenerateFullReport(records, "test/repo", "abc123", time.Now())
	entries := report.BuildSearchIndex(fr)

	if len(entries) != len(fr.Units) {
		t.Errorf("search entries = %d, want %d", len(entries), len(fr.Units))
	}

	// Verify fields are populated
	for _, e := range entries {
		if e.Name == "" {
			t.Error("search entry Name should not be empty")
		}
		if e.Path == "" {
			t.Error("search entry Path should not be empty")
		}
		if e.UnitID == "" {
			t.Error("search entry UnitID should not be empty")
		}
		if e.Grade == "" {
			t.Error("search entry Grade should not be empty")
		}
		if e.Status == "" {
			t.Error("search entry Status should not be empty")
		}
		if e.Language == "" {
			t.Error("search entry Language should not be empty")
		}
		if e.UnitURL == "" {
			t.Error("search entry UnitURL should not be empty")
		}
	}

	// Verify specific entries
	foundScore := false
	foundParse := false
	for _, e := range entries {
		if e.Name == "Score" {
			foundScore = true
			if e.Language != "go" {
				t.Errorf("Score language = %s, want go", e.Language)
			}
			if e.Status != "certified" {
				t.Errorf("Score status = %s, want certified", e.Status)
			}
		}
		if e.Name == "parse" {
			foundParse = true
			if e.Language != "ts" {
				t.Errorf("parse language = %s, want ts", e.Language)
			}
		}
	}
	if !foundScore {
		t.Error("search index should contain Score entry")
	}
	if !foundParse {
		t.Error("search index should contain parse entry")
	}
}

func TestFormatSearchIndexJS(t *testing.T) {
	records := []domain.CertificationRecord{
		makeSiteRecord("go", "internal/engine/scorer.go", "Score", domain.UnitTypeFunction, domain.StatusCertified, 0.92),
		makeSiteRecord("ts", "src/parser.ts", "parse", domain.UnitTypeFunction, domain.StatusProbationary, 0.65),
	}

	fr := report.GenerateFullReport(records, "test/repo", "abc123", time.Now())
	entries := report.BuildSearchIndex(fr)
	js := report.FormatSearchIndexJS(entries)

	// Must start with const declaration
	if !strings.HasPrefix(js, "const SEARCH_INDEX = ") {
		t.Error("JS should start with 'const SEARCH_INDEX = '")
	}

	// Must end with semicolon
	trimmed := strings.TrimSpace(js)
	if !strings.HasSuffix(trimmed, ";") {
		t.Error("JS should end with ';'")
	}

	// Must contain unit names
	if !strings.Contains(js, "Score") {
		t.Error("JS should contain 'Score'")
	}
	if !strings.Contains(js, "parse") {
		t.Error("JS should contain 'parse'")
	}

	// Must be valid-ish JSON array inside
	if !strings.Contains(js, "[") || !strings.Contains(js, "]") {
		t.Error("JS should contain JSON array")
	}
}
