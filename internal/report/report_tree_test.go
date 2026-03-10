package report_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/report"
)

// testFullReport builds a FullReport with units across 2 packages for tree tests.
func testFullReport() report.FullReport {
	records := []domain.CertificationRecord{
		makeFullRecord("go", "cmd/certify/main.go", "main", domain.UnitTypeFunction, domain.StatusCertified, 0.90),
		makeFullRecord("go", "cmd/certify/root.go", "Execute", domain.UnitTypeFunction, domain.StatusCertified, 0.88),
		makeFullRecord("go", "internal/engine/scorer.go", "Score", domain.UnitTypeFunction, domain.StatusCertified, 0.92),
		makeFullRecord("go", "internal/engine/pipeline.go", "CertifyUnit", domain.UnitTypeFunction, domain.StatusCertifiedWithObservations, 0.78),
	}
	records[3].Observations = []string{"complexity: 25 exceeds threshold 20"}
	return report.GenerateFullReport(records, "test/repo", "abc123", time.Now())
}

func TestGenerateReportTree_CreatesPackageIndexes(t *testing.T) {
	dir := t.TempDir()
	fr := testFullReport()

	_, err := report.GenerateReportTree(fr, dir)
	if err != nil {
		t.Fatalf("GenerateReportTree: %v", err)
	}

	for _, pkg := range []string{"cmd/certify", "internal/engine"} {
		idx := filepath.Join(dir, pkg, "index.md")
		if _, err := os.Stat(idx); os.IsNotExist(err) {
			t.Errorf("missing package index: %s", idx)
		}
	}
}

func TestGenerateReportTree_CreatesTopLevelIndex(t *testing.T) {
	dir := t.TempDir()
	fr := testFullReport()

	_, err := report.GenerateReportTree(fr, dir)
	if err != nil {
		t.Fatalf("GenerateReportTree: %v", err)
	}

	idx := filepath.Join(dir, "index.md")
	data, err := os.ReadFile(idx)
	if err != nil {
		t.Fatalf("reading top-level index: %v", err)
	}
	content := string(data)

	// Must link to both packages
	if !strings.Contains(content, "cmd/certify") {
		t.Error("top-level index should reference cmd/certify")
	}
	if !strings.Contains(content, "internal/engine") {
		t.Error("top-level index should reference internal/engine")
	}
	// Must contain relative links to package indexes
	if !strings.Contains(content, "cmd/certify/index.md") {
		t.Error("top-level index should link to cmd/certify/index.md")
	}
	if !strings.Contains(content, "internal/engine/index.md") {
		t.Error("top-level index should link to internal/engine/index.md")
	}
}

func TestGenerateReportTree_UnitFilesInSubdirs(t *testing.T) {
	dir := t.TempDir()
	fr := testFullReport()

	_, err := report.GenerateReportTree(fr, dir)
	if err != nil {
		t.Fatalf("GenerateReportTree: %v", err)
	}

	// Units with symbols: <pkg>/<filename>/<symbol>.md
	expected := []string{
		"cmd/certify/main.go/main.md",
		"cmd/certify/root.go/Execute.md",
		"internal/engine/scorer.go/Score.md",
		"internal/engine/pipeline.go/CertifyUnit.md",
	}
	for _, rel := range expected {
		path := filepath.Join(dir, rel)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Errorf("missing unit cert: %s", path)
		}
	}
}

func TestGenerateReportTree_FileUnitsNoSymbol(t *testing.T) {
	records := []domain.CertificationRecord{
		makeFullRecord("go", "scripts/release.sh", "", domain.UnitTypeFile, domain.StatusCertified, 0.85),
	}
	fr := report.GenerateFullReport(records, "test/repo", "abc123", time.Now())

	dir := t.TempDir()
	_, err := report.GenerateReportTree(fr, dir)
	if err != nil {
		t.Fatalf("GenerateReportTree: %v", err)
	}

	// File-level units (no symbol) → <pkg>/<filename>.md
	path := filepath.Join(dir, "scripts/release.sh.md")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Errorf("missing file-level unit cert: %s", path)
	}
}

func TestGenerateReportTree_PackageIndexLinksToUnits(t *testing.T) {
	dir := t.TempDir()
	fr := testFullReport()

	_, err := report.GenerateReportTree(fr, dir)
	if err != nil {
		t.Fatalf("GenerateReportTree: %v", err)
	}

	data, err := os.ReadFile(filepath.Join(dir, "internal/engine/index.md"))
	if err != nil {
		t.Fatalf("reading package index: %v", err)
	}
	content := string(data)

	// Should contain links to unit cert files
	if !strings.Contains(content, "scorer.go/Score.md") {
		t.Error("package index should link to scorer.go/Score.md")
	}
	if !strings.Contains(content, "pipeline.go/CertifyUnit.md") {
		t.Error("package index should link to pipeline.go/CertifyUnit.md")
	}
}

func TestGenerateReportTree_UnitLinksBackToPackage(t *testing.T) {
	dir := t.TempDir()
	fr := testFullReport()

	_, err := report.GenerateReportTree(fr, dir)
	if err != nil {
		t.Fatalf("GenerateReportTree: %v", err)
	}

	data, err := os.ReadFile(filepath.Join(dir, "internal/engine/scorer.go/Score.md"))
	if err != nil {
		t.Fatalf("reading unit cert: %v", err)
	}
	content := string(data)

	// Must link back to package index
	if !strings.Contains(content, "index.md") {
		t.Error("unit cert should link back to package index.md")
	}
	// Must reference the package name
	if !strings.Contains(content, "internal/engine") {
		t.Error("unit cert should reference its package name")
	}
}

func TestGenerateReportTree_CleansOldFlatFiles(t *testing.T) {
	dir := t.TempDir()

	// Plant an old flat file
	stale := filepath.Join(dir, "old-flat-file.md")
	if err := os.WriteFile(stale, []byte("stale"), 0o644); err != nil {
		t.Fatal(err)
	}

	fr := testFullReport()
	_, err := report.GenerateReportTree(fr, dir)
	if err != nil {
		t.Fatalf("GenerateReportTree: %v", err)
	}

	if _, err := os.Stat(stale); !os.IsNotExist(err) {
		t.Error("stale flat file should have been cleaned up")
	}
}

func TestBuildPackageSummaries(t *testing.T) {
	fr := testFullReport()
	summaries := report.BuildPackageSummaries(fr)

	if len(summaries) != 2 {
		t.Fatalf("expected 2 packages, got %d", len(summaries))
	}

	// Build a map for easier lookup
	byPath := make(map[string]report.PackageSummary)
	for _, s := range summaries {
		byPath[s.Path] = s
	}

	cmd := byPath["cmd/certify"]
	if cmd.Units != 2 {
		t.Errorf("cmd/certify units = %d, want 2", cmd.Units)
	}
	if cmd.Grade == "" {
		t.Error("cmd/certify grade should not be empty")
	}

	eng := byPath["internal/engine"]
	if eng.Units != 2 {
		t.Errorf("internal/engine units = %d, want 2", eng.Units)
	}
}

func TestFormatCardMarkdown_LinksToPackages(t *testing.T) {
	fr := testFullReport()
	fr.Card.Packages = report.BuildPackageSummaries(fr)

	md := report.FormatCardMarkdown(fr.Card)

	if !strings.Contains(md, "reports/cmd/certify/index.md") {
		t.Error("card markdown should link to reports/cmd/certify/index.md")
	}
	if !strings.Contains(md, "reports/internal/engine/index.md") {
		t.Error("card markdown should link to reports/internal/engine/index.md")
	}
	if !strings.Contains(md, "### Packages") {
		t.Error("card markdown should have a Packages section")
	}
}

func TestGenerateReportTree_ReturnsFileCount(t *testing.T) {
	dir := t.TempDir()
	fr := testFullReport()

	count, err := report.GenerateReportTree(fr, dir)
	if err != nil {
		t.Fatalf("GenerateReportTree: %v", err)
	}

	// 1 top-level index + 2 package indexes + 4 unit certs = 7
	if count != 7 {
		t.Errorf("file count = %d, want 7", count)
	}
}
