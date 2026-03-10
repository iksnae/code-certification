package report

import (
	"strings"
	"testing"
	"time"

	"github.com/iksnae/code-certification/internal/agent"
)

func TestFormatArchitectReport(t *testing.T) {
	snap := &agent.ArchSnapshot{
		Packages: []agent.PackageNode{
			{Path: "internal/engine", Units: 3, AvgScore: 0.833, Grade: "B", Observations: 3, TopIssues: []string{"errors_ignored"}},
			{Path: "internal/domain", Units: 2, AvgScore: 0.95, Grade: "A", Observations: 0},
		},
		Hotspots: []agent.PackageNode{
			{Path: "internal/engine", Units: 3, AvgScore: 0.833, Grade: "B"},
		},
		DependencyEdges: []agent.DepEdge{
			{From: "internal/engine", To: "internal/domain", Weight: 5},
		},
		CouplingPairs: []agent.CouplingPair{
			{PkgA: "internal/domain", PkgB: "internal/engine", EdgeCount: 5},
		},
		Layers: map[string]string{
			"internal/engine": "internal",
			"internal/domain": "domain",
		},
		Metrics: agent.SnapshotMetrics{
			TotalUnits:        5,
			TotalPackages:     2,
			AvgScore:          0.88,
			GradeDistribution: map[string]int{"A": 2, "B": 3},
			TopObservations:   map[string]int{"errors_ignored": 3},
		},
	}

	result := &agent.ArchitectResult{
		Snapshot:       snap,
		PhasesComplete: 6,
		TotalTokens:    5000,
		Duration:       30 * time.Second,
		Model:          "test-model",
		Phase1: &agent.ArchPhase1Result{
			Layers:               []agent.ArchLayer{{Name: "internal", Packages: []string{"engine", "domain"}, Description: "core logic"}},
			DependencyAssessment: "Clean layered architecture.",
		},
		Phase2: &agent.ArchPhase2Result{
			Findings: []agent.ArchFinding{{Package: "internal/engine", Issue: "high complexity", Severity: "medium"}},
		},
		Phase3: &agent.ArchPhase3Result{
			StrategyAssessment: "Test coverage is adequate.",
		},
		Phase4: &agent.ArchPhase4Result{
			Concerns: []agent.ArchConcern{{Area: "operations", Description: "no logging"}},
		},
		Phase5: &agent.ArchPhase5Result{
			Recommendations: []agent.ArchRecommendation{
				{
					Title:         "Reduce engine complexity",
					CurrentState:  "internal/engine: 83.3% avg, 3 observations",
					ProposedState: "Split into engine/scoring + engine/pipeline",
					Deltas: []agent.ArchDelta{
						{Metric: "avg_score", Current: "83.3%", Projected: "89.0%"},
						{Metric: "observations", Current: "3", Projected: "1"},
					},
					AffectedUnits: []string{"internal/engine/scorer.go#Score"},
					Effort:        "M",
					Justification: "Splitting reduces coupling.",
				},
			},
		},
		Phase6: &agent.ArchPhase6Result{
			ExecutiveSummary: "The project is well-structured with room for improvement.",
			RiskMatrix:       []agent.ArchRisk{{Risk: "High complexity", Severity: "medium", Likelihood: "high"}},
			Roadmap:          []agent.ArchRoadItem{{Priority: 1, Title: "Reduce complexity", Effort: "M", Impact: "high", DeltaSummary: "avg_score: 83.3% → 89.0%"}},
		},
	}

	pc := &agent.ProjectContext{
		RepoName:  "test-repo",
		CommitSHA: "abc123",
		Snapshot:  snap,
	}

	output := FormatArchitectReport(result, pc)

	// Part I: Snapshot tables
	if !strings.Contains(output, "Architecture Snapshot") {
		t.Error("should contain Architecture Snapshot section")
	}
	if !strings.Contains(output, "internal/engine") {
		t.Error("should contain package names from snapshot")
	}
	if !strings.Contains(output, "83.3%") {
		t.Error("should contain snapshot scores")
	}

	// Part II: Analysis
	if !strings.Contains(output, "Code Quality") {
		t.Error("should contain Code Quality section")
	}
	if !strings.Contains(output, "Test Strategy") {
		t.Error("should contain Test Strategy section")
	}

	// Part III: Comparative recommendations
	if !strings.Contains(output, "Recommendations") {
		t.Error("should contain Recommendations section")
	}
	if !strings.Contains(output, "Reduce engine complexity") {
		t.Error("should contain recommendation title")
	}
	if !strings.Contains(output, "| Metric | Current | Projected | Delta |") {
		t.Error("should contain delta table header")
	}
	if !strings.Contains(output, "avg_score") {
		t.Error("should contain delta metric name")
	}

	// Executive Summary
	if !strings.Contains(output, "Executive Summary") {
		t.Error("should contain Executive Summary")
	}

	// Risk Matrix
	if !strings.Contains(output, "Risk Matrix") {
		t.Error("should contain Risk Matrix")
	}

	// Roadmap
	if !strings.Contains(output, "Roadmap") {
		t.Error("should contain Roadmap")
	}
}

func TestFormatArchitectReport_SnapshotDeterministic(t *testing.T) {
	snap := &agent.ArchSnapshot{
		Packages: []agent.PackageNode{
			{Path: "pkg/a", Units: 2, AvgScore: 0.90, Grade: "A-"},
			{Path: "pkg/b", Units: 3, AvgScore: 0.80, Grade: "B"},
		},
		Metrics: agent.SnapshotMetrics{
			TotalUnits:        5,
			TotalPackages:     2,
			AvgScore:          0.84,
			GradeDistribution: map[string]int{"A-": 2, "B": 3},
			TopObservations:   map[string]int{},
		},
	}

	result1 := &agent.ArchitectResult{Snapshot: snap, PhasesComplete: 0}
	result2 := &agent.ArchitectResult{Snapshot: snap, PhasesComplete: 0}
	pc := &agent.ProjectContext{RepoName: "test", Snapshot: snap}

	out1 := FormatArchitectReport(result1, pc)
	out2 := FormatArchitectReport(result2, pc)

	if out1 != out2 {
		t.Error("snapshot-only reports should be deterministic")
	}

	// Should contain the package table
	if !strings.Contains(out1, "pkg/a") || !strings.Contains(out1, "pkg/b") {
		t.Error("should contain package names")
	}
}

func TestFormatArchitectReport_PartialResults(t *testing.T) {
	snap := &agent.ArchSnapshot{
		Packages: []agent.PackageNode{
			{Path: "internal/engine", Units: 3, AvgScore: 0.85, Grade: "B"},
		},
		Metrics: agent.SnapshotMetrics{
			TotalUnits:        3,
			TotalPackages:     1,
			AvgScore:          0.85,
			GradeDistribution: map[string]int{"B": 3},
			TopObservations:   map[string]int{},
		},
	}

	// Only Phase 1 completed, rest nil
	result := &agent.ArchitectResult{
		Snapshot:       snap,
		PhasesComplete: 1,
		Phase1: &agent.ArchPhase1Result{
			DependencyAssessment: "Simple structure.",
		},
		Errors: []string{"Phase 2 failed: timeout"},
	}
	pc := &agent.ProjectContext{RepoName: "test", Snapshot: snap}

	output := FormatArchitectReport(result, pc)

	// Part I should always be present
	if !strings.Contains(output, "Architecture Snapshot") {
		t.Error("snapshot should always be present")
	}
	if !strings.Contains(output, "internal/engine") {
		t.Error("package data should be present")
	}

	// Should note incomplete phases
	if !strings.Contains(output, "Phase 2 failed") {
		t.Error("should note failed phases")
	}
}

func TestFormatArchitectReport_ThinkingSection(t *testing.T) {
	snap := &agent.ArchSnapshot{
		Metrics: agent.SnapshotMetrics{
			GradeDistribution: map[string]int{},
			TopObservations:   map[string]int{},
		},
	}

	result := &agent.ArchitectResult{
		Snapshot:       snap,
		PhasesComplete: 2,
		Thinking: []string{
			"Let me analyze the architecture carefully...",
			"The code quality metrics show some concerns...",
			"", "", "", "",
		},
	}
	pc := &agent.ProjectContext{RepoName: "test", Snapshot: snap}

	output := FormatArchitectReport(result, pc)

	if !strings.Contains(output, "Agent Reasoning") {
		t.Error("should contain Agent Reasoning section")
	}
	if !strings.Contains(output, "<details>") {
		t.Error("thinking should be in collapsible details tags")
	}
	if !strings.Contains(output, "Phase 1: Architecture Narration") {
		t.Error("should label thinking with phase name")
	}
	if !strings.Contains(output, "analyze the architecture carefully") {
		t.Error("should contain the thinking text")
	}
}

func TestFormatArchitectReport_NoThinking(t *testing.T) {
	snap := &agent.ArchSnapshot{
		Metrics: agent.SnapshotMetrics{
			GradeDistribution: map[string]int{},
			TopObservations:   map[string]int{},
		},
	}

	result := &agent.ArchitectResult{
		Snapshot:       snap,
		PhasesComplete: 1,
		Thinking:       []string{"", "", "", "", "", ""},
	}
	pc := &agent.ProjectContext{RepoName: "test", Snapshot: snap}

	output := FormatArchitectReport(result, pc)

	if strings.Contains(output, "Agent Reasoning") {
		t.Error("should NOT contain Agent Reasoning when all thinking is empty")
	}
}

func TestFormatArchitectReport_ComparativeFormat(t *testing.T) {
	snap := &agent.ArchSnapshot{
		Packages: []agent.PackageNode{
			{Path: "internal/engine", Units: 3, AvgScore: 0.78, Grade: "C"},
		},
		Metrics: agent.SnapshotMetrics{
			TotalUnits: 3, TotalPackages: 1, AvgScore: 0.78,
			GradeDistribution: map[string]int{"C": 3},
			TopObservations:   map[string]int{},
		},
	}

	result := &agent.ArchitectResult{
		Snapshot:       snap,
		PhasesComplete: 5,
		Phase5: &agent.ArchPhase5Result{
			Recommendations: []agent.ArchRecommendation{
				{
					Title:         "Refactor engine",
					CurrentState:  "78% avg",
					ProposedState: "86% avg",
					Deltas: []agent.ArchDelta{
						{Metric: "avg_score", Current: "78.0%", Projected: "86.0%"},
						{Metric: "observations", Current: "5", Projected: "1"},
					},
					AffectedUnits: []string{"engine/scorer.go#Score", "engine/pipeline.go#Run"},
					Effort:        "L",
					Justification: "Based on splitting the package.",
				},
			},
		},
	}
	pc := &agent.ProjectContext{RepoName: "test", Snapshot: snap}

	output := FormatArchitectReport(result, pc)

	// Each recommendation should have a delta table
	if !strings.Contains(output, "| Metric | Current | Projected | Delta |") {
		t.Error("should have delta table header")
	}
	// Should have the delta values
	if !strings.Contains(output, "78.0%") {
		t.Error("should contain current value")
	}
	if !strings.Contains(output, "86.0%") {
		t.Error("should contain projected value")
	}
	// Should have affected units
	if !strings.Contains(output, "engine/scorer.go#Score") {
		t.Error("should list affected units")
	}
	// Should have effort
	if !strings.Contains(output, "**Effort:** L") {
		t.Error("should show effort")
	}
}
