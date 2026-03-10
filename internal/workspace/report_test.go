package workspace

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func testWorkspaceCard() WorkspaceCard {
	return WorkspaceCard{
		GeneratedAt:  "2026-03-10T12:00:00Z",
		TotalUnits:   150,
		TotalPassing: 145,
		TotalFailing: 5,
		OverallGrade: "A-",
		OverallScore: 0.91,
		PassRate:     0.9667,
		Submodules: []SubmoduleSummary{
			{Name: "api", Path: "services/api", Grade: "A", Score: 0.95, Units: 100, Passing: 100, Failing: 0, PassRate: 1.0, HasCertify: true, Commit: "abc123"},
			{Name: "web", Path: "services/web", Grade: "B", Score: 0.82, Units: 50, Passing: 45, Failing: 5, PassRate: 0.90, HasCertify: true, Commit: "def456"},
			{Name: "docs", Path: "docs", Grade: "", Score: 0, Units: 0, HasCertify: false},
		},
	}
}

func TestFormatWorkspaceCardMarkdown_SubmoduleLinks(t *testing.T) {
	wc := testWorkspaceCard()
	md := FormatWorkspaceCardMarkdown(wc)

	// Must link to each configured submodule's report tree
	if !strings.Contains(md, "services/api/.certification/reports/index.md") {
		t.Error("should link to api submodule report tree")
	}
	if !strings.Contains(md, "services/web/.certification/reports/index.md") {
		t.Error("should link to web submodule report tree")
	}
}

func TestFormatWorkspaceCardMarkdown_AggregateStats(t *testing.T) {
	wc := testWorkspaceCard()
	md := FormatWorkspaceCardMarkdown(wc)

	if !strings.Contains(md, "150") {
		t.Error("should show total units (150)")
	}
	if !strings.Contains(md, "145") {
		t.Error("should show total passing (145)")
	}
	if !strings.Contains(md, "Workspace") {
		t.Error("should identify as workspace report")
	}
	if !strings.Contains(md, "A-") {
		t.Error("should show overall grade A-")
	}
}

func TestFormatWorkspaceCardMarkdown_UnconfiguredSubmodule(t *testing.T) {
	wc := testWorkspaceCard()
	md := FormatWorkspaceCardMarkdown(wc)

	// Unconfigured submodule should appear but marked as not configured
	if !strings.Contains(md, "docs") {
		t.Error("should list unconfigured submodule")
	}
	if !strings.Contains(md, "—") || !strings.Contains(md, "not configured") {
		t.Error("should indicate unconfigured submodule status")
	}
}

func TestGenerateWorkspaceReportTree_CreatesFiles(t *testing.T) {
	dir := t.TempDir()
	wc := testWorkspaceCard()

	count, err := GenerateWorkspaceReportTree(wc, dir)
	if err != nil {
		t.Fatalf("GenerateWorkspaceReportTree: %v", err)
	}

	// Should create: index.md + 2 submodule summaries (configured only) = 3
	if count != 3 {
		t.Errorf("file count = %d, want 3", count)
	}

	// Check index exists
	if _, err := os.Stat(filepath.Join(dir, "index.md")); os.IsNotExist(err) {
		t.Error("missing index.md")
	}

	// Check per-submodule summaries
	if _, err := os.Stat(filepath.Join(dir, "api.md")); os.IsNotExist(err) {
		t.Error("missing api.md")
	}
	if _, err := os.Stat(filepath.Join(dir, "web.md")); os.IsNotExist(err) {
		t.Error("missing web.md")
	}
}

func TestGenerateWorkspaceReportTree_RelativePaths(t *testing.T) {
	dir := t.TempDir()
	wc := testWorkspaceCard()

	_, err := GenerateWorkspaceReportTree(wc, dir)
	if err != nil {
		t.Fatalf("GenerateWorkspaceReportTree: %v", err)
	}

	// Check index links to submodule summaries
	indexData, _ := os.ReadFile(filepath.Join(dir, "index.md"))
	index := string(indexData)

	if !strings.Contains(index, "api.md") {
		t.Error("index should link to api.md")
	}
	if !strings.Contains(index, "web.md") {
		t.Error("index should link to web.md")
	}

	// Check submodule summary links to the submodule's own report tree
	apiData, _ := os.ReadFile(filepath.Join(dir, "api.md"))
	api := string(apiData)

	// From reports/api.md, link to ../services/api/.certification/reports/index.md
	if !strings.Contains(api, "services/api/.certification/reports/index.md") {
		t.Error("api summary should link to submodule report tree")
	}

	// Back-link to workspace index
	if !strings.Contains(api, "index.md") {
		t.Error("api summary should link back to workspace index")
	}
}
