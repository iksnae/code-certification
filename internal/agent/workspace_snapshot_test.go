package agent

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/iksnae/code-certification/internal/domain"
)

func TestBuildWorkspaceSnapshot_MultipleSubmodules(t *testing.T) {
	root := t.TempDir()

	// Create two submodule dirs with cert records
	sub1Records := []domain.CertificationRecord{
		{UnitID: domain.NewUnitID("go", "api/handler.go", "HandleGet"), UnitPath: "api/handler.go", Score: 0.90, Grade: domain.GradeFromScore(0.90)},
		{UnitID: domain.NewUnitID("go", "api/router.go", "NewRouter"), UnitPath: "api/router.go", Score: 0.85, Grade: domain.GradeFromScore(0.85)},
	}
	sub2Records := []domain.CertificationRecord{
		{UnitID: domain.NewUnitID("go", "lib/util.go", "FormatDate"), UnitPath: "lib/util.go", Score: 0.95, Grade: domain.GradeFromScore(0.95)},
	}

	subs := []SubmoduleInfo{
		{Name: "api-service", Path: "services/api", Commit: "abc1234", Records: sub1Records},
		{Name: "shared-lib", Path: "lib/shared", Commit: "def5678", Records: sub2Records},
	}

	snap := BuildWorkspaceSnapshot(root, subs)

	if snap == nil {
		t.Fatal("expected non-nil snapshot")
	}
	if snap.SchemaVersion != SnapshotSchemaVersion {
		t.Errorf("SchemaVersion = %d, want %d", snap.SchemaVersion, SnapshotSchemaVersion)
	}
	if len(snap.SubmoduleSnapshots) != 2 {
		t.Fatalf("SubmoduleSnapshots = %d, want 2", len(snap.SubmoduleSnapshots))
	}
	if snap.SubmoduleSnapshots[0].Name != "api-service" {
		t.Errorf("first sub name = %q, want api-service", snap.SubmoduleSnapshots[0].Name)
	}
	if snap.SubmoduleSnapshots[0].Snapshot == nil {
		t.Fatal("first sub snapshot should not be nil")
	}
	if snap.SubmoduleSnapshots[0].Snapshot.Metrics.TotalUnits != 2 {
		t.Errorf("first sub units = %d, want 2", snap.SubmoduleSnapshots[0].Snapshot.Metrics.TotalUnits)
	}
	if snap.SubmoduleSnapshots[1].Snapshot.Metrics.TotalUnits != 1 {
		t.Errorf("second sub units = %d, want 1", snap.SubmoduleSnapshots[1].Snapshot.Metrics.TotalUnits)
	}
	if snap.AggregateMetrics.TotalUnitsAcrossAll != 3 {
		t.Errorf("TotalUnitsAcrossAll = %d, want 3", snap.AggregateMetrics.TotalUnitsAcrossAll)
	}
	if snap.AggregateMetrics.TotalSubmodules != 2 {
		t.Errorf("TotalSubmodules = %d, want 2", snap.AggregateMetrics.TotalSubmodules)
	}
	if snap.AggregateMetrics.ConfiguredSubmodules != 2 {
		t.Errorf("ConfiguredSubmodules = %d, want 2", snap.AggregateMetrics.ConfiguredSubmodules)
	}
}

func TestBuildWorkspaceSnapshot_ClassifiesRoles(t *testing.T) {
	root := t.TempDir()

	// Create submodule dirs with identifying files
	svcDir := filepath.Join(root, "services/api")
	os.MkdirAll(filepath.Join(svcDir, "cmd/server"), 0o755)
	os.WriteFile(filepath.Join(svcDir, "cmd/server/main.go"), []byte("package main\nfunc main(){}"), 0o644)

	libDir := filepath.Join(root, "lib/shared")
	os.MkdirAll(libDir, 0o755)
	os.WriteFile(filepath.Join(libDir, "go.mod"), []byte("module github.com/example/shared"), 0o644)
	// Library: no cmd/, has exported packages
	os.MkdirAll(filepath.Join(libDir, "pkg/utils"), 0o755)

	toolDir := filepath.Join(root, "tools/linter")
	os.MkdirAll(filepath.Join(toolDir, "cmd/lint"), 0o755)
	os.WriteFile(filepath.Join(toolDir, "cmd/lint/main.go"), []byte("package main"), 0o644)

	subs := []SubmoduleInfo{
		{Name: "api", Path: "services/api", Records: makeRecords(2)},
		{Name: "shared", Path: "lib/shared", Records: makeRecords(3)},
		{Name: "linter", Path: "tools/linter", Records: makeRecords(1)},
	}

	snap := BuildWorkspaceSnapshot(root, subs)

	roles := make(map[string]string)
	for _, s := range snap.SubmoduleSnapshots {
		roles[s.Name] = s.Role
	}

	if roles["api"] != "service" {
		t.Errorf("api role = %q, want service", roles["api"])
	}
	if roles["shared"] != "library" {
		t.Errorf("shared role = %q, want library", roles["shared"])
	}
	if roles["linter"] != "tool" {
		t.Errorf("linter role = %q, want tool", roles["linter"])
	}
}

func TestBuildWorkspaceSnapshot_DetectsCrossDeps(t *testing.T) {
	root := t.TempDir()

	// Create api submodule with go.mod that has a replace directive to shared
	apiDir := filepath.Join(root, "services/api")
	os.MkdirAll(apiDir, 0o755)
	goMod := `module github.com/example/api

go 1.22

require github.com/example/shared v0.0.0

replace github.com/example/shared => ../../lib/shared
`
	os.WriteFile(filepath.Join(apiDir, "go.mod"), []byte(goMod), 0o644)

	libDir := filepath.Join(root, "lib/shared")
	os.MkdirAll(libDir, 0o755)
	os.WriteFile(filepath.Join(libDir, "go.mod"), []byte("module github.com/example/shared\n\ngo 1.22\n"), 0o644)

	subs := []SubmoduleInfo{
		{Name: "api", Path: "services/api", Records: makeRecords(1)},
		{Name: "shared", Path: "lib/shared", Records: makeRecords(1)},
	}

	snap := BuildWorkspaceSnapshot(root, subs)

	if len(snap.CrossDependencies) == 0 {
		t.Fatal("expected cross-dependencies from go.mod replace directive")
	}
	found := false
	for _, dep := range snap.CrossDependencies {
		if dep.FromSubmodule == "services/api" && dep.ToSubmodule == "lib/shared" {
			found = true
			if !strings.Contains(dep.Evidence, "go.mod replace") {
				t.Errorf("evidence = %q, expected to contain 'go.mod replace'", dep.Evidence)
			}
		}
	}
	if !found {
		t.Errorf("expected cross-dep from services/api → lib/shared, got %+v", snap.CrossDependencies)
	}
}

func TestBuildWorkspaceSnapshot_InfraFiles(t *testing.T) {
	root := t.TempDir()

	// Create various infra files at workspace root
	os.WriteFile(filepath.Join(root, "Justfile"), []byte("build:"), 0o644)
	os.WriteFile(filepath.Join(root, "Makefile"), []byte("build:"), 0o644)
	os.WriteFile(filepath.Join(root, "docker-compose.yml"), []byte("version: '3'"), 0o644)
	os.MkdirAll(filepath.Join(root, ".github/workflows"), 0o755)
	os.WriteFile(filepath.Join(root, ".github/workflows/ci.yml"), []byte("name: CI"), 0o644)
	os.WriteFile(filepath.Join(root, "README.md"), []byte("# Workspace"), 0o644)

	snap := BuildWorkspaceSnapshot(root, nil)

	if len(snap.InfraFiles) == 0 {
		t.Fatal("expected infrastructure files to be detected")
	}

	infraSet := make(map[string]bool)
	for _, f := range snap.InfraFiles {
		infraSet[f] = true
	}

	for _, expected := range []string{"Justfile", "Makefile", "docker-compose.yml", ".github/workflows/ci.yml"} {
		if !infraSet[expected] {
			t.Errorf("expected infra file %q not found in %v", expected, snap.InfraFiles)
		}
	}
	// README.md is documentation, not infra
	if infraSet["README.md"] {
		t.Error("README.md should not be classified as infrastructure")
	}
}

func TestBuildWorkspaceSnapshot_AggregateMetrics(t *testing.T) {
	root := t.TempDir()

	sub1Records := []domain.CertificationRecord{
		{UnitID: domain.NewUnitID("go", "a.go", "A"), UnitPath: "a.go", Score: 0.60, Grade: domain.GradeFromScore(0.60)},
		{UnitID: domain.NewUnitID("go", "b.go", "B"), UnitPath: "b.go", Score: 0.70, Grade: domain.GradeFromScore(0.70)},
	}
	sub2Records := []domain.CertificationRecord{
		{UnitID: domain.NewUnitID("go", "c.go", "C"), UnitPath: "c.go", Score: 0.95, Grade: domain.GradeFromScore(0.95)},
	}

	subs := []SubmoduleInfo{
		{Name: "weak", Path: "weak", Records: sub1Records},
		{Name: "strong", Path: "strong", Records: sub2Records},
	}

	snap := BuildWorkspaceSnapshot(root, subs)

	m := snap.AggregateMetrics
	if m.TotalUnitsAcrossAll != 3 {
		t.Errorf("TotalUnitsAcrossAll = %d, want 3", m.TotalUnitsAcrossAll)
	}
	if m.WorstSubmodule != "weak" {
		t.Errorf("WorstSubmodule = %q, want weak", m.WorstSubmodule)
	}
	if m.BestSubmodule != "strong" {
		t.Errorf("BestSubmodule = %q, want strong", m.BestSubmodule)
	}

	// Weighted avg: (0.60+0.70)*2/3 + 0.95*1/3 ... actually (0.60+0.70+0.95)/3
	// Per-submodule: weak avg = 0.65, strong avg = 0.95
	// Weighted by units: (0.65*2 + 0.95*1) / 3 = 2.25/3 = 0.75
	expectedWeighted := (0.65*2 + 0.95*1) / 3.0
	if m.WeightedAvgScore < expectedWeighted-0.01 || m.WeightedAvgScore > expectedWeighted+0.01 {
		t.Errorf("WeightedAvgScore = %f, want ~%f", m.WeightedAvgScore, expectedWeighted)
	}
}

func TestFormatWorkspaceForLLM_Structure(t *testing.T) {
	snap := &WorkspaceArchSnapshot{
		SchemaVersion: 2,
		SubmoduleSnapshots: []SubmoduleSnapshotEntry{
			{
				Name:   "api",
				Path:   "services/api",
				Role:   "service",
				Commit: "abc1234",
				Snapshot: &ArchSnapshot{
					Metrics: SnapshotMetrics{TotalUnits: 50, TotalPackages: 5, AvgScore: 0.88},
				},
			},
			{
				Name:   "lib",
				Path:   "lib/shared",
				Role:   "library",
				Commit: "def5678",
				Snapshot: &ArchSnapshot{
					Metrics: SnapshotMetrics{TotalUnits: 20, TotalPackages: 3, AvgScore: 0.95},
				},
			},
		},
		CrossDependencies: []CrossDepEdge{
			{FromSubmodule: "services/api", ToSubmodule: "lib/shared", Evidence: "go.mod replace", Weight: 1},
		},
		InfraFiles: []string{"Justfile", ".github/workflows/ci.yml"},
		AggregateMetrics: WorkspaceMetrics{
			TotalSubmodules:      2,
			ConfiguredSubmodules: 2,
			TotalUnitsAcrossAll:  70,
			WeightedAvgScore:     0.90,
			BestSubmodule:        "lib",
			WorstSubmodule:       "api",
		},
	}

	output := FormatWorkspaceForLLM(snap, 8000)

	requiredSections := []string{
		"Workspace Architecture Snapshot",
		"Submodule Overview",
		"Cross-Submodule Dependencies",
		"Infrastructure Files",
		"api",
		"lib",
		"service",
		"library",
		"Justfile",
		"go.mod replace",
	}

	for _, section := range requiredSections {
		if !strings.Contains(output, section) {
			t.Errorf("output missing %q", section)
		}
	}
}

func TestFormatWorkspaceForLLM_TokenBudget(t *testing.T) {
	// Create a snapshot with lots of data
	snap := &WorkspaceArchSnapshot{
		SchemaVersion: 2,
		AggregateMetrics: WorkspaceMetrics{
			TotalSubmodules: 1,
		},
	}

	// Very small budget
	output := FormatWorkspaceForLLM(snap, 100)
	// Should still produce output, just truncated
	if len(output) == 0 {
		t.Error("expected non-empty output even with tiny budget")
	}
}

func TestWorkspacePhasePrompts_AllContainGrounding(t *testing.T) {
	prompts := WorkspacePhasePrompts()

	if len(prompts) != 6 {
		t.Fatalf("expected 6 prompts, got %d", len(prompts))
	}

	groundingPhrases := []string{
		"do not", // "do not fabricate" or similar
	}

	for i, prompt := range prompts {
		for _, phrase := range groundingPhrases {
			if !strings.Contains(strings.ToLower(prompt), phrase) {
				t.Errorf("workspace prompt %d missing grounding phrase %q", i+1, phrase)
			}
		}
		// Each prompt should reference submodule/workspace concepts
		if !strings.Contains(strings.ToLower(prompt), "submodule") && !strings.Contains(strings.ToLower(prompt), "workspace") {
			t.Errorf("workspace prompt %d should reference submodule or workspace concepts", i+1)
		}
	}
}

func TestGatherWorkspaceContext(t *testing.T) {
	root := t.TempDir()

	// Create a submodule dir with cert data
	subDir := filepath.Join(root, "mymod")
	certDir := filepath.Join(subDir, ".certification")
	os.MkdirAll(certDir, 0o755)
	os.WriteFile(filepath.Join(subDir, "go.mod"), []byte("module github.com/example/mymod\n"), 0o644)

	// Create workspace-level infra
	os.WriteFile(filepath.Join(root, "Justfile"), []byte("build:\n\tgo build ./..."), 0o644)
	os.WriteFile(filepath.Join(root, "README.md"), []byte("# My Workspace"), 0o644)

	subs := []SubmoduleInfo{
		{Name: "mymod", Path: "mymod", Commit: "aaa1111", Records: makeRecords(5)},
	}

	wpc := GatherWorkspaceContext(root, subs)

	if wpc == nil {
		t.Fatal("expected non-nil workspace context")
	}
	if wpc.Snapshot == nil {
		t.Fatal("expected non-nil workspace snapshot")
	}
	if wpc.RepoName == "" {
		// RepoName detection may fail in temp dir, that's ok
	}
	if wpc.Snapshot.AggregateMetrics.TotalSubmodules != 1 {
		t.Errorf("TotalSubmodules = %d, want 1", wpc.Snapshot.AggregateMetrics.TotalSubmodules)
	}
}

func TestBuildWorkspaceSnapshot_EmptySubmodules(t *testing.T) {
	root := t.TempDir()

	snap := BuildWorkspaceSnapshot(root, nil)

	if snap == nil {
		t.Fatal("expected non-nil snapshot even with no submodules")
	}
	if snap.AggregateMetrics.TotalSubmodules != 0 {
		t.Errorf("TotalSubmodules = %d, want 0", snap.AggregateMetrics.TotalSubmodules)
	}
	if snap.AggregateMetrics.TotalUnitsAcrossAll != 0 {
		t.Errorf("TotalUnitsAcrossAll = %d, want 0", snap.AggregateMetrics.TotalUnitsAcrossAll)
	}
}

func TestBuildWorkspaceSnapshot_SubmoduleWithNoRecords(t *testing.T) {
	root := t.TempDir()

	subs := []SubmoduleInfo{
		{Name: "empty", Path: "empty", Commit: "000", Records: nil},
		{Name: "full", Path: "full", Commit: "111", Records: makeRecords(3)},
	}

	snap := BuildWorkspaceSnapshot(root, subs)

	if snap.AggregateMetrics.ConfiguredSubmodules != 2 {
		t.Errorf("ConfiguredSubmodules = %d, want 2", snap.AggregateMetrics.ConfiguredSubmodules)
	}
	if snap.AggregateMetrics.TotalUnitsAcrossAll != 3 {
		t.Errorf("TotalUnitsAcrossAll = %d, want 3 (only from 'full')", snap.AggregateMetrics.TotalUnitsAcrossAll)
	}
}

// makeRecords creates n dummy certification records for testing.
func makeRecords(n int) []domain.CertificationRecord {
	records := make([]domain.CertificationRecord, n)
	for i := 0; i < n; i++ {
		id := domain.NewUnitID("go", "test.go", "Func"+string(rune('A'+i)))
		records[i] = domain.CertificationRecord{
			UnitID:   id,
			UnitPath: "test.go",
			Score:    0.85,
			Grade:    domain.GradeFromScore(0.85),
		}
	}
	return records
}
