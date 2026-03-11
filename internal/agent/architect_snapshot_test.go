package agent_test

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/iksnae/code-certification/internal/agent"
	"github.com/iksnae/code-certification/internal/domain"
)

func makeRecord(unitID string, score float64, observations []string) domain.CertificationRecord {
	uid, _ := domain.ParseUnitID(unitID)
	return domain.CertificationRecord{
		UnitID:       uid,
		UnitType:     domain.UnitTypeFunction,
		UnitPath:     uid.Path(),
		Score:        score,
		Grade:        domain.GradeFromScore(score),
		Status:       domain.StatusCertified,
		Observations: observations,
		CertifiedAt:  time.Now(),
		ExpiresAt:    time.Now().Add(90 * 24 * time.Hour),
	}
}

func makeRecordWithEvidence(unitID string, score float64, evidence ...domain.Evidence) domain.CertificationRecord {
	r := makeRecord(unitID, score, nil)
	r.Evidence = evidence
	return r
}

func TestBuildSnapshot_StructuralMetrics(t *testing.T) {
	records := []domain.CertificationRecord{
		makeRecordWithEvidence("go://pkg/a.go#Foo", 0.85, domain.Evidence{
			Kind:   domain.EvidenceKindStructural,
			Source: "structural",
			Passed: true,
			Metrics: map[string]float64{
				"panic_calls":          0,
				"os_exit_calls":        1,
				"global_mutable_count": 3,
				"defer_in_loop":        2,
				"has_init_func":        1,
				"context_not_first":    0,
				"errors_ignored":       1,
			},
		}),
		makeRecordWithEvidence("go://pkg/b.go#Bar", 0.80, domain.Evidence{
			Kind:   domain.EvidenceKindStructural,
			Source: "structural",
			Passed: true,
			Metrics: map[string]float64{
				"panic_calls":          2,
				"os_exit_calls":        0,
				"global_mutable_count": 1,
				"defer_in_loop":        0,
				"has_init_func":        0,
				"context_not_first":    1,
				"errors_ignored":       3,
			},
		}),
		// Record with no structural evidence — should not affect aggregates
		makeRecord("go://pkg/c.go#Baz", 0.90, nil),
	}

	snap := agent.BuildSnapshot(records, "")

	s := snap.Metrics.Structural
	if s.PanicCalls != 2 {
		t.Errorf("expected 2 panic_calls, got %d", s.PanicCalls)
	}
	if s.OsExitCalls != 1 {
		t.Errorf("expected 1 os_exit_calls, got %d", s.OsExitCalls)
	}
	if s.GlobalMutableCount != 4 {
		t.Errorf("expected 4 global_mutable_count, got %d", s.GlobalMutableCount)
	}
	if s.DeferInLoop != 2 {
		t.Errorf("expected 2 defer_in_loop, got %d", s.DeferInLoop)
	}
	if s.InitFuncCount != 1 {
		t.Errorf("expected 1 init_func_count, got %d", s.InitFuncCount)
	}
	if s.ContextNotFirst != 1 {
		t.Errorf("expected 1 context_not_first, got %d", s.ContextNotFirst)
	}
	if s.ErrorsIgnored != 4 {
		t.Errorf("expected 4 errors_ignored, got %d", s.ErrorsIgnored)
	}
}

func TestBuildSnapshot_StructuralMetrics_AllZero(t *testing.T) {
	records := []domain.CertificationRecord{
		makeRecordWithEvidence("go://pkg/a.go#Foo", 0.90, domain.Evidence{
			Kind:   domain.EvidenceKindStructural,
			Source: "structural",
			Passed: true,
			Metrics: map[string]float64{
				"panic_calls":          0,
				"os_exit_calls":        0,
				"global_mutable_count": 0,
			},
		}),
	}

	snap := agent.BuildSnapshot(records, "")
	s := snap.Metrics.Structural

	if s.PanicCalls != 0 {
		t.Errorf("expected 0 panic_calls, got %d", s.PanicCalls)
	}
	if s.OsExitCalls != 0 {
		t.Errorf("expected 0 os_exit_calls, got %d", s.OsExitCalls)
	}
}

func TestBuildSnapshot_StructuralMetrics_IgnoresNonStructuralEvidence(t *testing.T) {
	records := []domain.CertificationRecord{
		makeRecordWithEvidence("go://pkg/a.go#Foo", 0.85,
			// Lint evidence — should be ignored by structural aggregation
			domain.Evidence{
				Kind:    domain.EvidenceKindLint,
				Source:  "golangci-lint",
				Passed:  true,
				Metrics: map[string]float64{"lint_errors": 3},
			},
			// Structural evidence — should be aggregated
			domain.Evidence{
				Kind:   domain.EvidenceKindStructural,
				Source: "structural",
				Passed: true,
				Metrics: map[string]float64{
					"panic_calls":   1,
					"os_exit_calls": 0,
				},
			},
		),
	}

	snap := agent.BuildSnapshot(records, "")
	s := snap.Metrics.Structural
	if s.PanicCalls != 1 {
		t.Errorf("expected 1 panic_calls, got %d", s.PanicCalls)
	}
	if s.OsExitCalls != 0 {
		t.Errorf("expected 0 os_exit_calls, got %d", s.OsExitCalls)
	}
}

func TestBuildSnapshot_StructuralMetrics_PartialMetricsMap(t *testing.T) {
	// Evidence with only some metrics keys present — missing keys should default to 0
	records := []domain.CertificationRecord{
		makeRecordWithEvidence("go://pkg/a.go#Foo", 0.85, domain.Evidence{
			Kind:    domain.EvidenceKindStructural,
			Source:  "structural",
			Passed:  true,
			Metrics: map[string]float64{"panic_calls": 5},
		}),
	}

	snap := agent.BuildSnapshot(records, "")
	s := snap.Metrics.Structural
	if s.PanicCalls != 5 {
		t.Errorf("expected 5 panic_calls, got %d", s.PanicCalls)
	}
	// Keys not present in evidence should remain 0
	if s.OsExitCalls != 0 {
		t.Errorf("expected 0 os_exit_calls for missing key, got %d", s.OsExitCalls)
	}
	if s.GlobalMutableCount != 0 {
		t.Errorf("expected 0 global_mutable_count for missing key, got %d", s.GlobalMutableCount)
	}
}

func TestBuildSnapshot_StructuralMetrics_Extended(t *testing.T) {
	records := []domain.CertificationRecord{
		makeRecordWithEvidence("go://pkg/a.go#Foo", 0.85, domain.Evidence{
			Kind:   domain.EvidenceKindStructural,
			Source: "structural",
			Passed: true,
			Metrics: map[string]float64{
				"naked_returns":      2,
				"recursive_calls":    1,
				"loop_nesting_depth": 3,
				"nested_loop_pairs":  1,
				"quadratic_patterns": 0,
				"func_lines":         45,
				"param_count":        3,
				"return_count":       2,
				"method_count":       0,
			},
		}),
		makeRecordWithEvidence("go://pkg/b.go#Bar", 0.80, domain.Evidence{
			Kind:   domain.EvidenceKindStructural,
			Source: "structural",
			Passed: true,
			Metrics: map[string]float64{
				"naked_returns":      0,
				"recursive_calls":    2,
				"loop_nesting_depth": 4,
				"nested_loop_pairs":  0,
				"quadratic_patterns": 1,
				"func_lines":         120,
				"param_count":        5,
				"return_count":       1,
				"method_count":       8,
			},
		}),
	}

	snap := agent.BuildSnapshot(records, "")
	s := snap.Metrics.Structural

	if s.NakedReturns != 2 {
		t.Errorf("expected 2 naked_returns, got %d", s.NakedReturns)
	}
	if s.RecursiveCalls != 3 {
		t.Errorf("expected 3 recursive_calls, got %d", s.RecursiveCalls)
	}
	if s.NestedLoopPairs != 1 {
		t.Errorf("expected 1 nested_loop_pairs, got %d", s.NestedLoopPairs)
	}
	if s.QuadraticPatterns != 1 {
		t.Errorf("expected 1 quadratic_patterns, got %d", s.QuadraticPatterns)
	}
	if s.TotalFuncLines != 165 {
		t.Errorf("expected 165 total_func_lines, got %d", s.TotalFuncLines)
	}
	if s.TotalParams != 8 {
		t.Errorf("expected 8 total_params, got %d", s.TotalParams)
	}
	if s.TotalReturns != 3 {
		t.Errorf("expected 3 total_returns, got %d", s.TotalReturns)
	}
	if s.TotalMethods != 8 {
		t.Errorf("expected 8 total_methods, got %d", s.TotalMethods)
	}
	if s.MaxNestingDepth != 4 {
		t.Errorf("expected 4 max_nesting_depth, got %d", s.MaxNestingDepth)
	}
}

func TestBuildSnapshot_ContextNotFirst_SumsNotBool(t *testing.T) {
	records := []domain.CertificationRecord{
		makeRecordWithEvidence("go://pkg/a.go#Foo", 0.85, domain.Evidence{
			Kind:    domain.EvidenceKindStructural,
			Source:  "structural",
			Passed:  true,
			Metrics: map[string]float64{"context_not_first": 1},
		}),
		makeRecordWithEvidence("go://pkg/b.go#Bar", 0.80, domain.Evidence{
			Kind:    domain.EvidenceKindStructural,
			Source:  "structural",
			Passed:  true,
			Metrics: map[string]float64{"context_not_first": 1},
		}),
	}
	snap := agent.BuildSnapshot(records, "")
	if snap.Metrics.Structural.ContextNotFirst != 2 {
		t.Errorf("expected 2 context_not_first, got %d", snap.Metrics.Structural.ContextNotFirst)
	}
}

func TestBuildSnapshot_CoverageAggregates(t *testing.T) {
	records := []domain.CertificationRecord{
		makeRecordWithEvidence("go://pkg/a.go#Foo", 0.85, domain.Evidence{
			Kind:    domain.EvidenceKindTest,
			Source:  "coverage:unit",
			Passed:  true,
			Metrics: map[string]float64{"unit_test_coverage": 0.85},
		}),
		makeRecordWithEvidence("go://pkg/b.go#Bar", 0.80, domain.Evidence{
			Kind:    domain.EvidenceKindTest,
			Source:  "coverage:unit",
			Passed:  true,
			Metrics: map[string]float64{"unit_test_coverage": 0.60},
		}),
		// Unit with no coverage evidence
		makeRecord("go://pkg/c.go#Baz", 0.90, nil),
	}

	snap := agent.BuildSnapshot(records, "")
	c := snap.Metrics.Coverage

	if c.UnitsWithCoverage != 2 {
		t.Errorf("expected 2 units with coverage, got %d", c.UnitsWithCoverage)
	}
	if c.UnitsWithoutCoverage != 1 {
		t.Errorf("expected 1 unit without coverage, got %d", c.UnitsWithoutCoverage)
	}
	if diff := c.AvgCoverage - 0.725; diff > 0.001 || diff < -0.001 {
		t.Errorf("expected avg coverage ~0.725, got %.3f", c.AvgCoverage)
	}
	if c.MinCoverage != 0.60 {
		t.Errorf("expected min coverage 0.60, got %.2f", c.MinCoverage)
	}
	if c.MaxCoverage != 0.85 {
		t.Errorf("expected max coverage 0.85, got %.2f", c.MaxCoverage)
	}
}

func TestBuildSnapshot_CodeMetricsAggregates(t *testing.T) {
	records := []domain.CertificationRecord{
		makeRecordWithEvidence("go://pkg/a.go#Foo", 0.85, domain.Evidence{
			Kind:    domain.EvidenceKindMetrics,
			Source:  "metrics",
			Passed:  true,
			Metrics: map[string]float64{
				"code_lines":     80,
				"comment_lines":  10,
				"complexity":     5,
				"todo_count":     1,
			},
		}),
		makeRecordWithEvidence("go://pkg/b.go#Bar", 0.80, domain.Evidence{
			Kind:    domain.EvidenceKindMetrics,
			Source:  "metrics",
			Passed:  true,
			Metrics: map[string]float64{
				"code_lines":     200,
				"comment_lines":  30,
				"complexity":     12,
				"todo_count":     0,
			},
		}),
	}

	snap := agent.BuildSnapshot(records, "")
	cm := snap.Metrics.CodeMetrics

	if cm.TotalCodeLines != 280 {
		t.Errorf("expected 280 total_code_lines, got %d", cm.TotalCodeLines)
	}
	if cm.TotalCommentLines != 40 {
		t.Errorf("expected 40 total_comment_lines, got %d", cm.TotalCommentLines)
	}
	if cm.TotalComplexity != 17 {
		t.Errorf("expected 17 total_complexity, got %d", cm.TotalComplexity)
	}
	if cm.MaxComplexity != 12 {
		t.Errorf("expected 12 max_complexity, got %d", cm.MaxComplexity)
	}
	if cm.TotalTodos != 1 {
		t.Errorf("expected 1 total_todos, got %d", cm.TotalTodos)
	}
	// avg complexity = 17/2 = 8.5
	if diff := cm.AvgComplexity - 8.5; diff > 0.001 || diff < -0.001 {
		t.Errorf("expected avg complexity 8.5, got %.1f", cm.AvgComplexity)
	}
}

func TestBuildSnapshot_SchemaVersion(t *testing.T) {
	snap := agent.BuildSnapshot(nil, "")
	if snap.SchemaVersion != agent.SnapshotSchemaVersion {
		t.Errorf("expected schema version %d, got %d", agent.SnapshotSchemaVersion, snap.SchemaVersion)
	}

	records := []domain.CertificationRecord{
		makeRecord("go://pkg/a.go#Foo", 0.85, nil),
	}
	snap2 := agent.BuildSnapshot(records, "")
	if snap2.SchemaVersion != agent.SnapshotSchemaVersion {
		t.Errorf("expected schema version %d, got %d", agent.SnapshotSchemaVersion, snap2.SchemaVersion)
	}
}

func TestBuildSnapshot(t *testing.T) {
	records := []domain.CertificationRecord{
		makeRecord("go://internal/engine/scorer.go#Score", 0.85, []string{"errors_ignored: 2"}),
		makeRecord("go://internal/engine/scorer.go#Grade", 0.90, nil),
		makeRecord("go://internal/engine/pipeline.go#Run", 0.75, []string{"func_lines: 120", "errors_ignored: 3"}),
		makeRecord("go://internal/domain/unit.go#NewUnit", 0.95, nil),
		makeRecord("go://internal/domain/record.go#Status", 0.92, nil),
		makeRecord("go://cmd/certify/main.go#main", 0.70, []string{"init_func: true"}),
	}

	snap := agent.BuildSnapshot(records, "")

	// Should have 3 packages
	if len(snap.Packages) != 3 {
		t.Fatalf("expected 3 packages, got %d", len(snap.Packages))
	}

	// Check aggregate metrics
	if snap.Metrics.TotalUnits != 6 {
		t.Errorf("expected 6 total units, got %d", snap.Metrics.TotalUnits)
	}
	if snap.Metrics.TotalPackages != 3 {
		t.Errorf("expected 3 total packages, got %d", snap.Metrics.TotalPackages)
	}

	// Find engine package
	var enginePkg *agent.PackageNode
	for i := range snap.Packages {
		if snap.Packages[i].Path == "internal/engine" {
			enginePkg = &snap.Packages[i]
			break
		}
	}
	if enginePkg == nil {
		t.Fatal("engine package not found")
	}
	if enginePkg.Units != 3 {
		t.Errorf("engine should have 3 units, got %d", enginePkg.Units)
	}
	// avg of 0.85, 0.90, 0.75 = 0.8333...
	if enginePkg.AvgScore < 0.83 || enginePkg.AvgScore > 0.84 {
		t.Errorf("engine avg score should be ~0.833, got %f", enginePkg.AvgScore)
	}
	if enginePkg.Observations != 3 {
		t.Errorf("engine should have 3 observations, got %d", enginePkg.Observations)
	}

	// Check top observations aggregation
	if snap.Metrics.TopObservations == nil {
		t.Fatal("TopObservations should not be nil")
	}
}

func TestBuildSnapshot_Empty(t *testing.T) {
	snap := agent.BuildSnapshot(nil, "")

	if len(snap.Packages) != 0 {
		t.Errorf("expected 0 packages, got %d", len(snap.Packages))
	}
	if snap.Metrics.TotalUnits != 0 {
		t.Errorf("expected 0 units, got %d", snap.Metrics.TotalUnits)
	}
	if snap.Metrics.TotalPackages != 0 {
		t.Errorf("expected 0 packages, got %d", snap.Metrics.TotalPackages)
	}
	if snap.Metrics.GradeDistribution == nil {
		t.Error("GradeDistribution should be initialized even when empty")
	}
	if snap.Metrics.TopObservations == nil {
		t.Error("TopObservations should be initialized even when empty")
	}
}

func TestBuildSnapshot_Hotspots(t *testing.T) {
	records := []domain.CertificationRecord{
		// Large low-quality package — should rank highest
		makeRecord("go://internal/big/a.go#A", 0.60, []string{"issue1"}),
		makeRecord("go://internal/big/b.go#B", 0.55, []string{"issue2"}),
		makeRecord("go://internal/big/c.go#C", 0.50, []string{"issue3"}),
		makeRecord("go://internal/big/d.go#D", 0.65, []string{"issue4"}),
		// Small low-quality package — should rank lower despite low scores
		makeRecord("go://internal/tiny/x.go#X", 0.40, []string{"bad"}),
		// Large high-quality package — should rank lowest
		makeRecord("go://internal/good/a.go#A", 0.95, nil),
		makeRecord("go://internal/good/b.go#B", 0.92, nil),
		makeRecord("go://internal/good/c.go#C", 0.90, nil),
		makeRecord("go://internal/good/d.go#D", 0.88, nil),
	}

	snap := agent.BuildSnapshot(records, "")

	if len(snap.Hotspots) == 0 {
		t.Fatal("expected hotspots")
	}

	// internal/big should be #1 hotspot: 4 units × (1 - 0.575) = 4 × 0.425 = 1.7
	// internal/tiny: 1 unit × (1 - 0.40) = 0.6
	// internal/good: 4 units × (1 - 0.9125) = 4 × 0.0875 = 0.35
	if snap.Hotspots[0].Path != "internal/big" {
		t.Errorf("expected internal/big as top hotspot, got %s", snap.Hotspots[0].Path)
	}
}

func TestBuildSnapshot_DependencyGraph(t *testing.T) {
	// Create temp Go files with imports
	tmpDir := t.TempDir()

	// Create go.mod
	goMod := `module github.com/test/repo

go 1.21
`
	if err := os.MkdirAll(filepath.Join(tmpDir, "internal", "engine"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(tmpDir, "internal", "domain"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(tmpDir, "go.mod"), []byte(goMod), 0644); err != nil {
		t.Fatal(err)
	}

	// engine imports domain
	engineSrc := `package engine

import (
	"fmt"

	"github.com/test/repo/internal/domain"
)

func Score(u domain.Unit) float64 {
	fmt.Println(u)
	return 0.5
}
`
	if err := os.WriteFile(filepath.Join(tmpDir, "internal", "engine", "scorer.go"), []byte(engineSrc), 0644); err != nil {
		t.Fatal(err)
	}

	domainSrc := `package domain

type Unit struct {
	Name string
}
`
	if err := os.WriteFile(filepath.Join(tmpDir, "internal", "domain", "unit.go"), []byte(domainSrc), 0644); err != nil {
		t.Fatal(err)
	}

	records := []domain.CertificationRecord{
		makeRecord("go://internal/engine/scorer.go#Score", 0.80, nil),
		makeRecord("go://internal/domain/unit.go#Unit", 0.90, nil),
	}

	snap := agent.BuildSnapshot(records, tmpDir)

	// Should have dependency edge: engine → domain
	if len(snap.DependencyEdges) == 0 {
		t.Fatal("expected dependency edges")
	}

	found := false
	for _, e := range snap.DependencyEdges {
		if e.From == "internal/engine" && e.To == "internal/domain" {
			found = true
			if e.Weight < 1 {
				t.Errorf("expected weight >= 1, got %d", e.Weight)
			}
		}
	}
	if !found {
		t.Error("expected edge from internal/engine to internal/domain")
	}

	// Check importedBy
	var domainNode *agent.PackageNode
	for i := range snap.Packages {
		if snap.Packages[i].Path == "internal/domain" {
			domainNode = &snap.Packages[i]
		}
	}
	if domainNode == nil {
		t.Fatal("domain package not found")
	}
	if len(domainNode.ImportedBy) == 0 {
		t.Error("domain should be imported by engine")
	}
}

func TestBuildSnapshot_Deterministic(t *testing.T) {
	records1 := []domain.CertificationRecord{
		makeRecord("go://internal/b/x.go#X", 0.80, nil),
		makeRecord("go://internal/a/y.go#Y", 0.90, nil),
		makeRecord("go://internal/c/z.go#Z", 0.70, nil),
	}
	records2 := []domain.CertificationRecord{
		makeRecord("go://internal/c/z.go#Z", 0.70, nil),
		makeRecord("go://internal/a/y.go#Y", 0.90, nil),
		makeRecord("go://internal/b/x.go#X", 0.80, nil),
	}

	snap1 := agent.BuildSnapshot(records1, "")
	snap2 := agent.BuildSnapshot(records2, "")

	if len(snap1.Packages) != len(snap2.Packages) {
		t.Fatal("package count differs")
	}
	for i := range snap1.Packages {
		if snap1.Packages[i].Path != snap2.Packages[i].Path {
			t.Errorf("package order differs at %d: %s vs %s", i, snap1.Packages[i].Path, snap2.Packages[i].Path)
		}
		if snap1.Packages[i].AvgScore != snap2.Packages[i].AvgScore {
			t.Errorf("score differs for %s", snap1.Packages[i].Path)
		}
	}
}
