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

// makeSiteRecord is a helper to build test records for site tests.
func makeSiteRecord(lang, path, symbol string, utype domain.UnitType, status domain.Status, score float64) domain.CertificationRecord {
	now := time.Now()
	dims := domain.DimensionScores{
		domain.DimCorrectness:     score + 0.05,
		domain.DimMaintainability: score,
		domain.DimReadability:     score - 0.02,
		domain.DimTestability:     score + 0.03,
		domain.DimSecurity:        0.80,
	}
	return domain.CertificationRecord{
		UnitID:      domain.NewUnitID(lang, path, symbol),
		UnitType:    utype,
		UnitPath:    path,
		Status:      status,
		Score:       score,
		Grade:       domain.GradeFromScore(score),
		Confidence:  1.0,
		Dimensions:  dims,
		CertifiedAt: now,
		ExpiresAt:   now.Add(90 * 24 * time.Hour),
		Source:      "deterministic",
	}
}

func makeSiteReport(records []domain.CertificationRecord) report.FullReport {
	return report.GenerateFullReport(records, "test/repo", "abc123", time.Now())
}

func testRecords() []domain.CertificationRecord {
	records := []domain.CertificationRecord{
		makeSiteRecord("go", "internal/engine/scorer.go", "Score", domain.UnitTypeFunction, domain.StatusCertified, 0.92),
		makeSiteRecord("go", "internal/engine/scorer.go", "Normalize", domain.UnitTypeFunction, domain.StatusCertified, 0.88),
		makeSiteRecord("go", "internal/engine/pipeline.go", "CertifyUnit", domain.UnitTypeFunction, domain.StatusCertified, 0.85),
		makeSiteRecord("go", "internal/engine/pipeline.go", "RunPipeline", domain.UnitTypeFunction, domain.StatusCertifiedWithObservations, 0.78),
		makeSiteRecord("go", "internal/report/card.go", "GenerateCard", domain.UnitTypeFunction, domain.StatusCertified, 0.90),
		makeSiteRecord("go", "internal/report/card.go", "FormatText", domain.UnitTypeFunction, domain.StatusCertified, 0.87),
		makeSiteRecord("go", "cmd/certify/main.go", "main", domain.UnitTypeFunction, domain.StatusCertified, 0.80),
		makeSiteRecord("go", "cmd/certify/root.go", "Execute", domain.UnitTypeFunction, domain.StatusProbationary, 0.65),
		makeSiteRecord("ts", "src/parser.ts", "parse", domain.UnitTypeFunction, domain.StatusCertifiedWithObservations, 0.75),
		makeSiteRecord("ts", "src/parser.ts", "tokenize", domain.UnitTypeFunction, domain.StatusCertified, 0.82),
		makeSiteRecord("ts", "src/utils.ts", "format", domain.UnitTypeFunction, domain.StatusDecertified, 0.45),
	}
	records[3].Observations = []string{"🤖 High cyclomatic complexity detected", "💡 Consider breaking into smaller functions"}
	records[10].Observations = []string{"lint errors: 5", "test failures: 2"}
	return records
}

func TestGenerateSite_CreatesOutputStructure(t *testing.T) {
	records := testRecords()
	fr := makeSiteReport(records)
	outDir := t.TempDir()

	cfg := report.SiteConfig{
		OutputDir:     outDir,
		Title:         "test/repo",
		IncludeSearch: true,
	}

	if err := report.GenerateSite(fr, cfg); err != nil {
		t.Fatalf("GenerateSite failed: %v", err)
	}

	// index.html must exist
	indexPath := filepath.Join(outDir, "index.html")
	assertFileExists(t, indexPath)
	assertFileContains(t, indexPath, "<!DOCTYPE html>")

	// Package pages: internal/engine, internal/report, cmd/certify, src
	expectedPackages := []string{
		"internal/engine",
		"internal/report",
		"cmd/certify",
		"src",
	}
	for _, pkg := range expectedPackages {
		pkgPath := filepath.Join(outDir, "packages", pkg, "index.html")
		assertFileExists(t, pkgPath)
		assertFileContains(t, pkgPath, "<!DOCTYPE html>")
	}

	// Unit pages: one per unit
	unitDir := filepath.Join(outDir, "units")
	entries, err := os.ReadDir(unitDir)
	if err != nil {
		t.Fatalf("reading units dir: %v", err)
	}
	if len(entries) != len(fr.Units) {
		t.Errorf("unit pages = %d, want %d", len(entries), len(fr.Units))
	}
	for _, e := range entries {
		fp := filepath.Join(unitDir, e.Name())
		assertFileContains(t, fp, "<!DOCTYPE html>")
	}
}

func TestGenerateSite_IndexContent(t *testing.T) {
	records := testRecords()
	fr := makeSiteReport(records)
	outDir := t.TempDir()

	cfg := report.SiteConfig{
		OutputDir:     outDir,
		Title:         "test/repo",
		IncludeSearch: true,
	}

	if err := report.GenerateSite(fr, cfg); err != nil {
		t.Fatalf("GenerateSite failed: %v", err)
	}

	content := readFile(t, filepath.Join(outDir, "index.html"))

	// Must contain repo name
	assertContains(t, content, "test/repo")

	// Must contain overall grade
	assertContains(t, content, fr.Card.OverallGrade)

	// Must contain total units count
	assertContains(t, content, fmt.Sprintf("%d", fr.Card.TotalUnits))

	// Must contain language names
	assertContains(t, content, "go")
	assertContains(t, content, "ts")

	// Must contain links to package pages
	assertContains(t, content, "packages/")

	// Must contain proper HTML structure
	assertContains(t, content, "<!DOCTYPE html>")
	assertContains(t, content, "</html>")

	// Must contain search input since IncludeSearch is true
	assertContains(t, content, "search")
}

func TestGenerateSite_PackagePages(t *testing.T) {
	records := testRecords()
	fr := makeSiteReport(records)
	outDir := t.TempDir()

	cfg := report.SiteConfig{OutputDir: outDir, Title: "test/repo"}

	if err := report.GenerateSite(fr, cfg); err != nil {
		t.Fatalf("GenerateSite failed: %v", err)
	}

	// Check internal/engine package page
	enginePage := readFile(t, filepath.Join(outDir, "packages", "internal", "engine", "index.html"))

	// Must contain unit names
	assertContains(t, enginePage, "Score")
	assertContains(t, enginePage, "CertifyUnit")

	// Must contain links to unit detail pages
	assertContains(t, enginePage, "units/")

	// Must contain grade info
	assertContains(t, enginePage, "Grade")

	// src package page
	srcPage := readFile(t, filepath.Join(outDir, "packages", "src", "index.html"))
	assertContains(t, srcPage, "parse")
	assertContains(t, srcPage, "format")
}

func TestGenerateSite_UnitPages(t *testing.T) {
	records := testRecords()
	fr := makeSiteReport(records)
	outDir := t.TempDir()

	cfg := report.SiteConfig{OutputDir: outDir, Title: "test/repo"}

	if err := report.GenerateSite(fr, cfg); err != nil {
		t.Fatalf("GenerateSite failed: %v", err)
	}

	// Find the unit page for the first unit
	unitDir := filepath.Join(outDir, "units")
	entries, err := os.ReadDir(unitDir)
	if err != nil {
		t.Fatalf("reading units dir: %v", err)
	}
	if len(entries) == 0 {
		t.Fatal("no unit pages generated")
	}

	// Read a unit page and check content
	firstPage := readFile(t, filepath.Join(unitDir, entries[0].Name()))
	assertContains(t, firstPage, "<!DOCTYPE html>")

	// Find the RunPipeline unit page by filename (has observations)
	runPipelinePage := filepath.Join(unitDir, "internal-engine-pipeline-go-runpipeline.html")
	if _, err := os.Stat(runPipelinePage); os.IsNotExist(err) {
		t.Fatal("RunPipeline unit page not found")
	}
	rpContent := readFile(t, runPipelinePage)
	assertContains(t, rpContent, "High cyclomatic complexity")
	assertContains(t, rpContent, "Consider breaking into smaller functions")

	// Verify a unit page has back-link to package
	for _, e := range entries {
		content := readFile(t, filepath.Join(unitDir, e.Name()))
		if strings.Contains(content, "Score") && strings.Contains(content, "internal/engine") {
			assertContains(t, content, "packages/")
			break
		}
	}
}

func TestGenerateSite_EmptyReport(t *testing.T) {
	fr := makeSiteReport(nil)
	outDir := t.TempDir()

	cfg := report.SiteConfig{OutputDir: outDir, Title: "empty/repo"}

	if err := report.GenerateSite(fr, cfg); err != nil {
		t.Fatalf("GenerateSite failed: %v", err)
	}

	indexPath := filepath.Join(outDir, "index.html")
	assertFileExists(t, indexPath)
	content := readFile(t, indexPath)
	assertContains(t, content, "<!DOCTYPE html>")
	// Should indicate no data
	if !strings.Contains(strings.ToLower(content), "no") || (!strings.Contains(strings.ToLower(content), "data") && !strings.Contains(strings.ToLower(content), "unit")) {
		// Accept various "no data", "no units", "0 units" messages
		assertContains(t, content, "0")
	}
}

func TestGenerateSite_SingleUnit(t *testing.T) {
	records := []domain.CertificationRecord{
		makeSiteRecord("go", "main.go", "main", domain.UnitTypeFunction, domain.StatusCertified, 0.90),
	}
	fr := makeSiteReport(records)
	outDir := t.TempDir()

	cfg := report.SiteConfig{OutputDir: outDir, Title: "tiny/repo"}

	if err := report.GenerateSite(fr, cfg); err != nil {
		t.Fatalf("GenerateSite failed: %v", err)
	}

	// Should have index + 1 package page + 1 unit page
	assertFileExists(t, filepath.Join(outDir, "index.html"))

	// Package page for "." directory
	pkgEntries, _ := filepath.Glob(filepath.Join(outDir, "packages", "*", "index.html"))
	if len(pkgEntries) == 0 {
		// Try root-level package
		pkgEntries, _ = filepath.Glob(filepath.Join(outDir, "packages", "index.html"))
	}
	if len(pkgEntries) < 1 {
		t.Error("expected at least 1 package page")
	}

	unitEntries, _ := os.ReadDir(filepath.Join(outDir, "units"))
	if len(unitEntries) != 1 {
		t.Errorf("unit pages = %d, want 1", len(unitEntries))
	}
}

func TestGenerateSite_FileProtocol(t *testing.T) {
	records := testRecords()
	fr := makeSiteReport(records)
	outDir := t.TempDir()

	cfg := report.SiteConfig{OutputDir: outDir, Title: "test/repo"}

	if err := report.GenerateSite(fr, cfg); err != nil {
		t.Fatalf("GenerateSite failed: %v", err)
	}

	// Walk all HTML files and check for absolute URLs
	err := filepath.Walk(outDir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(path, ".html") {
			return err
		}
		content := readFile(t, path)
		// Should not contain absolute URLs (http:// or https://) in href/src attributes
		// Exceptions: links to external resources like GitHub
		for _, line := range strings.Split(content, "\n") {
			if strings.Contains(line, `href="http`) || strings.Contains(line, `src="http`) {
				// Allow external links (github, shields.io) but not local resources
				if strings.Contains(line, "github.com") || strings.Contains(line, "shields.io") {
					continue
				}
				t.Errorf("absolute URL found in %s: %s", path, strings.TrimSpace(line))
			}
		}
		return nil
	})
	if err != nil {
		t.Fatalf("walking site dir: %v", err)
	}
}

func TestGenerateSite_SpecialCharacters(t *testing.T) {
	records := []domain.CertificationRecord{
		makeSiteRecord("go", "internal/engine/scorer.go", "Score$Helper", domain.UnitTypeFunction, domain.StatusCertified, 0.90),
		makeSiteRecord("go", "internal/engine/scorer.go", "Process_Data", domain.UnitTypeFunction, domain.StatusCertified, 0.85),
	}
	fr := makeSiteReport(records)
	outDir := t.TempDir()

	cfg := report.SiteConfig{OutputDir: outDir, Title: "test/repo"}

	if err := report.GenerateSite(fr, cfg); err != nil {
		t.Fatalf("GenerateSite failed: %v", err)
	}

	// Should generate valid unit pages for units with special chars
	unitEntries, err := os.ReadDir(filepath.Join(outDir, "units"))
	if err != nil {
		t.Fatalf("reading units dir: %v", err)
	}
	if len(unitEntries) != 2 {
		t.Errorf("unit pages = %d, want 2", len(unitEntries))
	}

	// All generated files should have valid names (no special chars that break file systems)
	for _, e := range unitEntries {
		if strings.ContainsAny(e.Name(), "$<>|\"?*") {
			t.Errorf("invalid filename: %s", e.Name())
		}
		fp := filepath.Join(outDir, "units", e.Name())
		assertFileContains(t, fp, "<!DOCTYPE html>")
	}
}

func TestGenerateSite_LargeScale(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping large scale test in short mode")
	}

	// Generate 5000 units across 200 packages and 3 languages
	var records []domain.CertificationRecord
	langs := []string{"go", "ts", "py"}
	for i := 0; i < 5000; i++ {
		lang := langs[i%3]
		pkg := fmt.Sprintf("pkg/%s/sub%d", lang, i%200)
		symbol := fmt.Sprintf("Func%d", i)
		ext := ".go"
		if lang == "ts" {
			ext = ".ts"
		} else if lang == "py" {
			ext = ".py"
		}
		path := fmt.Sprintf("%s/file%d%s", pkg, i%10, ext)
		score := 0.5 + float64(i%50)/100.0
		status := domain.StatusCertified
		if i%20 == 0 {
			status = domain.StatusProbationary
		}
		records = append(records, makeSiteRecord(lang, path, symbol, domain.UnitTypeFunction, status, score))
	}

	fr := makeSiteReport(records)
	outDir := t.TempDir()

	cfg := report.SiteConfig{
		OutputDir:     outDir,
		Title:         "large/repo",
		IncludeSearch: true,
	}

	start := time.Now()
	if err := report.GenerateSite(fr, cfg); err != nil {
		t.Fatalf("GenerateSite failed: %v", err)
	}
	elapsed := time.Since(start)

	// Must complete in under 10 seconds
	if elapsed > 10*time.Second {
		t.Errorf("site generation took %v, want < 10s", elapsed)
	}

	// index.html under 500KB
	indexInfo, err := os.Stat(filepath.Join(outDir, "index.html"))
	if err != nil {
		t.Fatalf("stat index.html: %v", err)
	}
	if indexInfo.Size() > 512000 {
		t.Errorf("index.html = %d bytes, want < 512000", indexInfo.Size())
	}

	// Verify unit pages exist
	unitEntries, _ := os.ReadDir(filepath.Join(outDir, "units"))
	if len(unitEntries) < 4900 { // allow some dedup
		t.Errorf("unit pages = %d, want ~5000", len(unitEntries))
	}

	// Spot-check a few pages contain valid HTML
	for i, e := range unitEntries {
		if i > 10 {
			break
		}
		assertFileContains(t, filepath.Join(outDir, "units", e.Name()), "<!DOCTYPE html>")
	}

	// No individual page over 200KB (spot-check packages)
	err = filepath.Walk(filepath.Join(outDir, "packages"), func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}
		if info.Size() > 200*1024 {
			t.Errorf("package page %s = %d bytes, want < 200KB", path, info.Size())
		}
		return nil
	})
	if err != nil {
		t.Fatalf("walking packages: %v", err)
	}
}

// Helper functions

func assertFileExists(t *testing.T, path string) {
	t.Helper()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Errorf("expected file to exist: %s", path)
	}
}

func assertFileContains(t *testing.T, path, substr string) {
	t.Helper()
	content := readFile(t, path)
	if !strings.Contains(content, substr) {
		t.Errorf("file %s should contain %q", path, substr)
	}
}

func assertContains(t *testing.T, content, substr string) {
	t.Helper()
	if !strings.Contains(content, substr) {
		t.Errorf("content should contain %q (content length: %d)", substr, len(content))
	}
}

func readFile(t *testing.T, path string) string {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading %s: %v", path, err)
	}
	return string(data)
}
