package report_test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/report"
)

func unassessedFullReport() report.FullReport {
	records := []domain.CertificationRecord{
		unassessedRecord("swift", "app/a.swift", "a"),
		unassessedRecord("swift", "app/b.swift", "b"),
		unassessedRecord("swift", "app/c.swift", "c"),
	}
	return report.GenerateFullReport(records, "test/repo", "abc", time.Now())
}

// TestFormatFullMarkdown_AllUnassessedRendersNoPassRate covers
// `certify report --format full`, which reads the same Card.
func TestFormatFullMarkdown_AllUnassessedRendersNoPassRate(t *testing.T) {
	out := report.FormatFullMarkdown(unassessedFullReport())

	if strings.Contains(out, "| **Pass Rate** | 100.0% |") || strings.Contains(out, "| **Pass Rate** | 0.0% |") {
		t.Errorf("full report states a pass rate for an unassessed repo:\n%s", out)
	}
	if !strings.Contains(out, "| **Pass Rate** | n/a |") {
		t.Errorf("full report should render the pass rate as n/a, got:\n%s", out)
	}
	if !strings.Contains(out, "| **Not Assessed** | 3 |") {
		t.Errorf("full report should report the unassessed count, got:\n%s", out)
	}
	if !strings.Contains(out, "| N/A |") {
		t.Errorf("full report grade distribution should render the N/A row, got:\n%s", out)
	}
	if strings.Contains(out, "3 units certified") {
		t.Errorf("full report footer claims 3 units were certified when none were assessed:\n%s", out)
	}
}

// TestBuildLanguageDetail_ExcludesUnassessedFromPassing covers the by-language
// breakdown, whose Passing count is printed by `certify report --detailed`.
func TestBuildLanguageDetail_ExcludesUnassessedFromPassing(t *testing.T) {
	fr := report.GenerateFullReport([]domain.CertificationRecord{
		makeCardRecord("go", "a.go", "A", domain.StatusCertified, 0.90),
		unassessedRecord("swift", "b.swift", "b"),
		unassessedRecord("swift", "c.swift", "c"),
	}, "r", "abc", time.Now())

	byName := make(map[string]report.LanguageDetail)
	for _, l := range fr.LanguageDetail {
		byName[l.Name] = l
	}

	if got := byName["swift"].Passing; got != 0 {
		t.Errorf("swift Passing = %d, want 0 — no swift unit was assessed", got)
	}
	if got := byName["swift"].Unsupported; got != 2 {
		t.Errorf("swift Unsupported = %d, want 2", got)
	}
	if got := byName["go"].Passing; got != 1 {
		t.Errorf("go Passing = %d, want 1", got)
	}
}

// TestGenerateFullReport_UnitsCarryUnsupported checks the flag survives the
// record → UnitReport projection. Everything downstream of FullReport counts
// over UnitReport, so a projection that drops the flag makes the three-state
// distinction unrepresentable and every consumer reinvents the false positive.
func TestGenerateFullReport_UnitsCarryUnsupported(t *testing.T) {
	fr := unassessedFullReport()
	for _, u := range fr.Units {
		if !u.Unsupported {
			t.Errorf("UnitReport %s lost the Unsupported flag", u.UnitID)
		}
	}
}

// TestBuildPackageSummaries_ExcludesUnassessed is the package-level copy of the
// same inversion. These summaries render into .certification/reports/<pkg>/,
// which is committed.
func TestBuildPackageSummaries_ExcludesUnassessed(t *testing.T) {
	pkgs := report.BuildPackageSummaries(unassessedFullReport())

	if len(pkgs) != 1 {
		t.Fatalf("package count = %d, want 1", len(pkgs))
	}
	p := pkgs[0]
	if p.PassRate == 1.0 {
		t.Errorf("package PassRate = %.1f%%, want not 100%% — no unit was assessed", p.PassRate*100)
	}
	if p.Unsupported != 3 {
		t.Errorf("package Unsupported = %d, want 3", p.Unsupported)
	}
	if p.PassRateKnown() {
		t.Error("package PassRateKnown() = true, want false")
	}
}

// TestGenerateSite_AllUnassessedRendersNoPassRate covers the HTML site, the
// third rendered surface fed by the same numbers.
func TestGenerateSite_AllUnassessedRendersNoPassRate(t *testing.T) {
	dir := t.TempDir()
	if err := report.GenerateSite(unassessedFullReport(), report.SiteConfig{
		OutputDir: dir,
		Title:     "test",
	}); err != nil {
		t.Fatalf("GenerateSite() error: %v", err)
	}

	data, err := os.ReadFile(filepath.Join(dir, "index.html"))
	if err != nil {
		t.Fatalf("index.html not found: %v", err)
	}
	html := string(data)

	const passRateCard = `<div class="stat-card"><div class="value">%s</div><div class="label">Pass Rate</div></div>`
	if !strings.Contains(html, fmt.Sprintf(passRateCard, "n/a")) {
		t.Errorf("site index should render the pass rate as n/a, got:\n%s", html)
	}
	if strings.Contains(html, fmt.Sprintf(passRateCard, "100.0%")) {
		t.Errorf("site index states a 100%% pass rate for an unassessed repo:\n%s", html)
	}
	if strings.Contains(html, "units certified") {
		t.Errorf("site index footer claims units were certified when none were assessed:\n%s", html)
	}
	if !strings.Contains(html, "3 not assessed") {
		t.Errorf("site index footer should report the unassessed count, got:\n%s", html)
	}
}
