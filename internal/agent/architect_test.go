package agent_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/iksnae/code-certification/internal/agent"
	"github.com/iksnae/code-certification/internal/domain"
)

func TestGatherContext(t *testing.T) {
	tmpDir := t.TempDir()
	certDir := filepath.Join(tmpDir, ".certification")
	recordsDir := filepath.Join(certDir, "records")
	os.MkdirAll(recordsDir, 0755)

	// Create a README
	os.WriteFile(filepath.Join(tmpDir, "README.md"), []byte("# Test Project\nA test repo."), 0644)

	// Create some Go files
	os.MkdirAll(filepath.Join(tmpDir, "internal", "engine"), 0755)
	os.WriteFile(filepath.Join(tmpDir, "internal", "engine", "main.go"), []byte("package engine\n\nfunc Run() {}"), 0644)

	records := []domain.CertificationRecord{
		makeRecord("go://internal/engine/main.go#Run", 0.85, []string{"errors_ignored: 2"}),
	}

	pc, err := agent.GatherContext(tmpDir, certDir, records)
	if err != nil {
		t.Fatalf("GatherContext failed: %v", err)
	}

	if pc.Snapshot == nil {
		t.Fatal("snapshot should not be nil")
	}
	if pc.Snapshot.Metrics.TotalUnits != 1 {
		t.Errorf("expected 1 unit, got %d", pc.Snapshot.Metrics.TotalUnits)
	}
	if pc.Documentation["README.md"] == "" {
		t.Error("README should be loaded")
	}
	if !strings.Contains(pc.FileTree, "internal") {
		t.Error("file tree should contain 'internal'")
	}
}

func TestGatherContext_MissingData(t *testing.T) {
	tmpDir := t.TempDir()
	certDir := filepath.Join(tmpDir, ".certification")

	// No records, no docs, no cert dir
	pc, err := agent.GatherContext(tmpDir, certDir, nil)
	if err != nil {
		t.Fatalf("GatherContext should not fail with missing data: %v", err)
	}
	if pc.Snapshot == nil {
		t.Fatal("snapshot should not be nil even with no data")
	}
	if pc.Snapshot.Metrics.TotalUnits != 0 {
		t.Errorf("expected 0 units, got %d", pc.Snapshot.Metrics.TotalUnits)
	}
}

func TestFormatForLLM(t *testing.T) {
	records := []domain.CertificationRecord{
		makeRecord("go://internal/engine/scorer.go#Score", 0.85, []string{"errors_ignored: 2"}),
		makeRecord("go://internal/engine/pipeline.go#Run", 0.75, []string{"func_lines: 120"}),
		makeRecord("go://internal/domain/unit.go#NewUnit", 0.95, nil),
	}

	snap := agent.BuildSnapshot(records, "")
	pc := &agent.ProjectContext{
		RepoName:  "test-repo",
		CommitSHA: "abc123",
		Languages: []string{"go"},
		Snapshot:  snap,
		Documentation: map[string]string{
			"README.md": "# Test\nSome description.",
		},
	}

	output := pc.FormatForLLM(4000)

	// Should contain snapshot tables
	if !strings.Contains(output, "Package Map") {
		t.Error("should contain Package Map section")
	}
	if !strings.Contains(output, "internal/engine") {
		t.Error("should contain package name")
	}
	if !strings.Contains(output, "Hotspots") {
		t.Error("should contain Hotspots section")
	}
	if !strings.Contains(output, "Grade Distribution") {
		t.Error("should contain Grade Distribution")
	}

	// Should be within token budget hint (roughly 4 chars/token)
	maxChars := 4000 * 4
	if len(output) > maxChars {
		t.Errorf("output too long: %d chars (budget hint: %d tokens)", len(output), 4000)
	}
}

func TestFormatForLLM_SnapshotTables(t *testing.T) {
	records := []domain.CertificationRecord{
		makeRecord("go://internal/engine/scorer.go#Score", 0.85, []string{"errors_ignored: 2"}),
		makeRecord("go://internal/domain/unit.go#NewUnit", 0.95, nil),
	}

	snap := agent.BuildSnapshot(records, "")
	pc := &agent.ProjectContext{
		RepoName: "test-repo",
		Snapshot: snap,
	}

	output := pc.FormatForLLM(4000)

	// Verify table headers
	if !strings.Contains(output, "| Package |") {
		t.Error("package table should have header row")
	}
	if !strings.Contains(output, "85.0%") || !strings.Contains(output, "95.0%") {
		t.Error("package table should contain score percentages")
	}
}

func TestFormatForLLM_Empty(t *testing.T) {
	pc := &agent.ProjectContext{
		Snapshot: agent.BuildSnapshot(nil, ""),
	}

	output := pc.FormatForLLM(4000)
	if output == "" {
		t.Error("should produce output even with empty context")
	}
	// Should not panic
	_ = output
}

// makeRecord is defined in architect_snapshot_test.go — using the same helper
// This file and architect_snapshot_test.go are in the same test package (agent_test)
// so makeRecord is shared.
var _ = time.Now // silence unused import
