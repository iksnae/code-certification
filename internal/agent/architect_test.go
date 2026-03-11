package agent_test

import (
	"context"
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

func TestFormatForLLM_StructuralMetrics(t *testing.T) {
	records := []domain.CertificationRecord{
		makeRecordWithEvidence("go://pkg/a.go#Foo", 0.85, domain.Evidence{
			Kind:   domain.EvidenceKindStructural,
			Source: "structural",
			Passed: true,
			Metrics: map[string]float64{
				"panic_calls":          0,
				"os_exit_calls":        1,
				"global_mutable_count": 5,
				"defer_in_loop":        3,
			},
		}),
	}

	snap := agent.BuildSnapshot(records, "")
	pc := &agent.ProjectContext{
		RepoName: "test-repo",
		Snapshot: snap,
	}

	output := pc.FormatForLLM(4000)

	if !strings.Contains(output, "Structural Metrics") {
		t.Error("output should contain Structural Metrics section")
	}
	if !strings.Contains(output, "| panic_calls | 0 |") {
		t.Error("output should show panic_calls with exact count 0")
	}
	if !strings.Contains(output, "| os_exit_calls | 1 |") {
		t.Error("output should show os_exit_calls with exact count 1")
	}
	if !strings.Contains(output, "| global_mutable_count | 5 |") {
		t.Error("output should show global_mutable_count with exact count 5")
	}
	if !strings.Contains(output, "| defer_in_loop | 3 |") {
		t.Error("output should show defer_in_loop with exact count 3")
	}
}

func TestFormatForLLM_CoverageMetrics(t *testing.T) {
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
	}

	snap := agent.BuildSnapshot(records, "")
	pc := &agent.ProjectContext{
		RepoName: "test-repo",
		Snapshot: snap,
	}

	output := pc.FormatForLLM(4000)

	if !strings.Contains(output, "Coverage Metrics") {
		t.Error("output should contain Coverage Metrics section")
	}
	if !strings.Contains(output, "Units with coverage data: 2") {
		t.Error("output should show units with coverage")
	}
	if !strings.Contains(output, "Average coverage: 72.5%") {
		t.Error("output should show average coverage")
	}
}

func TestFormatForLLM_CodeMetrics(t *testing.T) {
	records := []domain.CertificationRecord{
		makeRecordWithEvidence("go://pkg/a.go#Foo", 0.85, domain.Evidence{
			Kind:   domain.EvidenceKindMetrics,
			Source: "metrics",
			Passed: true,
			Metrics: map[string]float64{
				"code_lines": 200,
				"complexity": 12,
				"todo_count": 3,
			},
		}),
	}

	snap := agent.BuildSnapshot(records, "")
	pc := &agent.ProjectContext{
		RepoName: "test-repo",
		Snapshot: snap,
	}

	output := pc.FormatForLLM(4000)

	if !strings.Contains(output, "Code Metrics") {
		t.Error("output should contain Code Metrics section")
	}
	if !strings.Contains(output, "| total_code_lines | 200 |") {
		t.Error("output should show total_code_lines")
	}
	if !strings.Contains(output, "| total_complexity | 12 |") {
		t.Error("output should show total_complexity")
	}
	if !strings.Contains(output, "| total_todos | 3 |") {
		t.Error("output should show total_todos")
	}
}

func TestFormatForLLM_SchemaVersion(t *testing.T) {
	snap := agent.BuildSnapshot(nil, "")
	pc := &agent.ProjectContext{
		RepoName: "test-repo",
		Snapshot: snap,
	}

	output := pc.FormatForLLM(4000)
	if !strings.Contains(output, "Snapshot Schema:") {
		t.Error("output should contain schema version")
	}
}

func TestArchitectPrompts_AllContainGrounding(t *testing.T) {
	prompts := agent.ArchitectPhasePrompts()
	for i, p := range prompts {
		hasGrounding := strings.Contains(p, "do not") || strings.Contains(p, "Do not") ||
			strings.Contains(p, "DO NOT") || strings.Contains(p, "Never") ||
			strings.Contains(p, "never")
		if !hasGrounding {
			t.Errorf("Phase %d prompt missing anti-hallucination grounding", i+1)
		}
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

func TestArchitectReview_MockProvider(t *testing.T) {
	records := []domain.CertificationRecord{
		makeRecord("go://internal/engine/scorer.go#Score", 0.85, []string{"errors_ignored: 2"}),
		makeRecord("go://internal/domain/unit.go#NewUnit", 0.95, nil),
	}

	snap := agent.BuildSnapshot(records, "")
	pc := &agent.ProjectContext{
		RepoName: "test-repo",
		Snapshot: snap,
	}

	callCount := 0
	responses := []string{
		// Phase 1
		`{"layers":[{"name":"internal","packages":["engine","domain"],"description":"core logic"}],"data_flows":[],"dependency_assessment":"clean"}`,
		// Phase 2
		`{"findings":[{"package":"internal/engine","issue":"high complexity","current_metrics":{"avg_score":0.85},"severity":"medium"}]}`,
		// Phase 3
		`{"coverage_gaps":[{"package":"internal/engine","current_score":0.85,"issue":"no integration tests"}],"strategy_assessment":"adequate"}`,
		// Phase 4
		`{"concerns":[{"area":"operations","description":"no structured logging","affected_packages":["internal/engine"],"metrics":{}}]}`,
		// Phase 5
		`{"recommendations":[{"title":"Add structured logging","current_state":"No logging in engine","proposed_state":"Add slog-based logging","deltas":[{"metric":"observability","current":"none","projected":"structured"}],"affected_units":["internal/engine/scorer.go#Score"],"effort":"S","justification":"Standard practice"}]}`,
		// Phase 6
		`{"executive_summary":"The project is well-structured.","risk_matrix":[{"risk":"No logging","severity":"medium","likelihood":"high","recommendation_ref":"Add structured logging"}],"roadmap":[{"priority":1,"title":"Add logging","effort":"S","impact":"medium","recommendation_ref":"Add structured logging","delta_summary":"observability: none → structured"}]}`,
	}

	mock := &sequenceProvider{responses: responses, callCount: &callCount}

	reviewer := &agent.ArchitectReviewer{
		Provider: mock,
		Model:    "test-model",
	}

	result, err := reviewer.Review(context.Background(), pc, nil)
	if err != nil {
		t.Fatalf("Review failed: %v", err)
	}

	if result.PhasesComplete != 6 {
		t.Errorf("expected 6 phases complete, got %d", result.PhasesComplete)
	}
	if result.TotalTokens == 0 {
		t.Error("expected non-zero token count")
	}
	if result.Snapshot == nil {
		t.Error("snapshot should be carried through")
	}

	// Verify Phase 1 parsed
	if result.Phase1 == nil {
		t.Fatal("Phase1 should be parsed")
	}
	if len(result.Phase1.Layers) != 1 {
		t.Errorf("expected 1 layer, got %d", len(result.Phase1.Layers))
	}

	// Verify Phase 5 has recommendations with deltas
	if result.Phase5 == nil {
		t.Fatal("Phase5 should be parsed")
	}
	if len(result.Phase5.Recommendations) != 1 {
		t.Fatalf("expected 1 recommendation, got %d", len(result.Phase5.Recommendations))
	}
	rec := result.Phase5.Recommendations[0]
	if len(rec.Deltas) == 0 {
		t.Error("recommendation should have deltas")
	}

	// Verify Phase 6 synthesis
	if result.Phase6 == nil {
		t.Fatal("Phase6 should be parsed")
	}
	if result.Phase6.ExecutiveSummary == "" {
		t.Error("executive summary should not be empty")
	}
}

func TestArchitectReview_PhaseFailure(t *testing.T) {
	records := []domain.CertificationRecord{
		makeRecord("go://internal/engine/scorer.go#Score", 0.85, nil),
	}

	snap := agent.BuildSnapshot(records, "")
	pc := &agent.ProjectContext{
		RepoName: "test-repo",
		Snapshot: snap,
	}

	// Provider that fails on phase 2 (second call)
	callCount := 0
	mock := &conditionalProvider{
		failUntil: 1, // fail first call (phase 1), succeed rest
		response:  `{"findings":[]}`,
		callCount: &callCount,
	}

	reviewer := &agent.ArchitectReviewer{
		Provider: mock,
		Model:    "test-model",
	}

	result, err := reviewer.Review(context.Background(), pc, nil)
	if err != nil {
		t.Fatalf("Review should not fail entirely: %v", err)
	}

	// Should have partial results
	if result.PhasesComplete == 6 {
		t.Error("should not complete all phases when one fails")
	}
	if len(result.Errors) == 0 {
		t.Error("should record errors for failed phases")
	}
	// Snapshot should always be present
	if result.Snapshot == nil {
		t.Error("snapshot should always be present even with failures")
	}
}

func TestArchitectReview_SinglePhase(t *testing.T) {
	snap := agent.BuildSnapshot(nil, "")
	pc := &agent.ProjectContext{
		RepoName: "test-repo",
		Snapshot: snap,
	}

	mock := &mockProvider{response: `{"coverage_gaps":[],"strategy_assessment":"looks good"}`}

	reviewer := &agent.ArchitectReviewer{
		Provider: mock,
		Model:    "test-model",
	}

	// Run only phase 3
	result, err := reviewer.Review(context.Background(), pc, []int{3})
	if err != nil {
		t.Fatalf("Review failed: %v", err)
	}

	if result.PhasesComplete != 1 {
		t.Errorf("expected 1 phase complete, got %d", result.PhasesComplete)
	}
	if result.Phase3 == nil {
		t.Error("Phase3 should be parsed")
	}
	if result.Phase1 != nil {
		t.Error("Phase1 should be nil when only phase 3 runs")
	}
}

func TestArchitectReview_Phase5Validation(t *testing.T) {
	snap := agent.BuildSnapshot(nil, "")
	pc := &agent.ProjectContext{
		RepoName: "test-repo",
		Snapshot: snap,
	}

	// Phase 5 response with a recommendation missing deltas
	callCount := 0
	responses := []string{
		`{"layers":[],"data_flows":[],"dependency_assessment":"ok"}`,
		`{"findings":[]}`,
		`{"coverage_gaps":[],"strategy_assessment":"ok"}`,
		`{"concerns":[]}`,
		`{"recommendations":[{"title":"Fix X","current_state":"bad","proposed_state":"good","deltas":[],"affected_units":[],"effort":"S","justification":"because"}]}`,
		`{"executive_summary":"done","risk_matrix":[],"roadmap":[]}`,
	}
	mock := &sequenceProvider{responses: responses, callCount: &callCount}

	reviewer := &agent.ArchitectReviewer{
		Provider: mock,
		Model:    "test-model",
	}

	result, err := reviewer.Review(context.Background(), pc, nil)
	if err != nil {
		t.Fatalf("Review failed: %v", err)
	}

	// Recommendations with empty deltas should get a placeholder
	if result.Phase5 == nil || len(result.Phase5.Recommendations) == 0 {
		t.Fatal("Phase5 should have recommendations")
	}
	rec := result.Phase5.Recommendations[0]
	if len(rec.Deltas) == 0 {
		t.Error("empty deltas should be filled with placeholder")
	}
}

func TestArchitectPrompts(t *testing.T) {
	prompts := agent.ArchitectPhasePrompts()
	names := agent.ArchitectPhaseNames()

	if len(prompts) != 6 {
		t.Fatalf("expected 6 prompts, got %d", len(prompts))
	}
	if len(names) != 6 {
		t.Fatalf("expected 6 phase names, got %d", len(names))
	}

	// Phase 5 should require comparative format
	if !strings.Contains(prompts[4], "deltas") {
		t.Error("Phase 5 prompt should mention deltas")
	}
	if !strings.Contains(prompts[4], "current_state") {
		t.Error("Phase 5 prompt should mention current_state")
	}
	if !strings.Contains(prompts[4], "proposed_state") {
		t.Error("Phase 5 prompt should mention proposed_state")
	}

	// Phase 1 should not recommend changes
	if !strings.Contains(prompts[0], "DO NOT recommend") {
		t.Error("Phase 1 should say DO NOT recommend")
	}
}

func TestArchitectE2E(t *testing.T) {
	// End-to-end test: GatherContext → Review → verify result
	tmpDir := t.TempDir()
	certDir := filepath.Join(tmpDir, ".certification")
	os.MkdirAll(filepath.Join(certDir, "records"), 0755)
	os.MkdirAll(filepath.Join(certDir, "policies"), 0755)

	// Create sample Go files
	os.MkdirAll(filepath.Join(tmpDir, "internal", "engine"), 0755)
	os.MkdirAll(filepath.Join(tmpDir, "internal", "domain"), 0755)
	os.WriteFile(filepath.Join(tmpDir, "go.mod"), []byte("module github.com/test/repo\n\ngo 1.22\n"), 0644)
	os.WriteFile(filepath.Join(tmpDir, "internal", "engine", "scorer.go"),
		[]byte("package engine\n\nimport \"github.com/test/repo/internal/domain\"\n\nfunc Score(u domain.Unit) float64 { return 0 }\n"), 0644)
	os.WriteFile(filepath.Join(tmpDir, "internal", "domain", "unit.go"),
		[]byte("package domain\n\ntype Unit struct{ Name string }\n"), 0644)
	os.WriteFile(filepath.Join(tmpDir, "README.md"), []byte("# Test Project"), 0644)
	os.WriteFile(filepath.Join(certDir, "policies", "go-standard.yml"), []byte("name: go-standard\n"), 0644)

	records := []domain.CertificationRecord{
		makeRecord("go://internal/engine/scorer.go#Score", 0.80, []string{"errors_ignored: 2"}),
		makeRecord("go://internal/domain/unit.go#Unit", 0.95, nil),
	}

	// 1. Gather context
	pc, err := agent.GatherContext(tmpDir, certDir, records)
	if err != nil {
		t.Fatalf("GatherContext: %v", err)
	}
	pc.RepoName = "test/repo"
	pc.CommitSHA = "abc123"

	if len(pc.Snapshot.DependencyEdges) == 0 {
		t.Error("snapshot should have dependency edges from Go import analysis")
	}
	if pc.Documentation["README.md"] == "" {
		t.Error("should load README")
	}
	if len(pc.PolicyRules) == 0 {
		t.Error("should detect policy files")
	}

	// 2. Run review with mock
	callCount := 0
	responses := []string{
		`{"layers":[{"name":"internal","packages":["engine","domain"],"description":"core"}],"data_flows":[{"from":"engine","to":"domain","description":"uses types"}],"dependency_assessment":"clean"}`,
		`{"findings":[{"package":"internal/engine","issue":"ignored errors","current_metrics":{"errors_ignored":2},"severity":"high"}]}`,
		`{"coverage_gaps":[],"strategy_assessment":"minimal tests"}`,
		`{"concerns":[]}`,
		`{"recommendations":[{"title":"Fix ignored errors","current_state":"2 errors ignored","proposed_state":"0 errors","deltas":[{"metric":"errors_ignored","current":"2","projected":"0"}],"affected_units":["internal/engine/scorer.go#Score"],"effort":"S","justification":"Standard practice"}]}`,
		`{"executive_summary":"Well-structured project.","risk_matrix":[{"risk":"Ignored errors","severity":"high","likelihood":"medium","recommendation_ref":"Fix ignored errors"}],"roadmap":[{"priority":1,"title":"Fix errors","effort":"S","impact":"high","recommendation_ref":"Fix ignored errors","delta_summary":"errors_ignored: 2 → 0"}]}`,
	}
	mock := &sequenceProvider{responses: responses, callCount: &callCount}

	var phaseStarts, phaseDones int
	reviewer := &agent.ArchitectReviewer{
		Provider:     mock,
		Model:        "test-model",
		OnPhaseStart: func(int, string) { phaseStarts++ },
		OnPhaseDone:  func(int, string, int) { phaseDones++ },
	}

	result, err := reviewer.Review(context.Background(), pc, nil)
	if err != nil {
		t.Fatalf("Review: %v", err)
	}

	if result.PhasesComplete != 6 {
		t.Errorf("expected 6 phases, got %d", result.PhasesComplete)
	}
	if phaseStarts != 6 || phaseDones != 6 {
		t.Errorf("callbacks: starts=%d dones=%d, expected 6/6", phaseStarts, phaseDones)
	}

	// 3. Verify LLM context includes snapshot data
	llmContext := pc.FormatForLLM(4000)
	if !strings.Contains(llmContext, "internal/engine") {
		t.Error("LLM context should contain package names")
	}
	if !strings.Contains(llmContext, "errors_ignored") {
		t.Error("LLM context should contain observation types")
	}

	// 4. Verify result integrity
	if result.Snapshot == nil {
		t.Fatal("result should carry snapshot")
	}
	if result.Phase5 == nil || len(result.Phase5.Recommendations) == 0 {
		t.Fatal("Phase5 should have recommendations")
	}
	rec := result.Phase5.Recommendations[0]
	if len(rec.Deltas) == 0 {
		t.Error("recommendation should have deltas")
	}
	if result.Phase6 == nil || result.Phase6.ExecutiveSummary == "" {
		t.Error("Phase6 should have executive summary")
	}
}

func TestArchitectReview_ThinkTags(t *testing.T) {
	snap := agent.BuildSnapshot(nil, "")
	pc := &agent.ProjectContext{
		RepoName: "test-repo",
		Snapshot: snap,
	}

	// Model wraps response in <think> tags (qwen3 pattern)
	callCount := 0
	thinkWrapped := `<think>
Let me analyze the architecture. The project has a clear layered structure
with cmd → internal → domain. There are some coupling concerns between
the engine and evidence packages.
</think>

{"layers":[{"name":"internal","packages":["engine"],"description":"core"}],"data_flows":[],"dependency_assessment":"clean"}`

	responses := []string{
		thinkWrapped,
		`<think>Checking quality metrics...</think>{"findings":[]}`,
		`{"coverage_gaps":[],"strategy_assessment":"ok"}`,
		`{"concerns":[]}`,
		`{"recommendations":[]}`,
		`{"executive_summary":"Done.","risk_matrix":[],"roadmap":[]}`,
	}
	mock := &sequenceProvider{responses: responses, callCount: &callCount}

	reviewer := &agent.ArchitectReviewer{
		Provider: mock,
		Model:    "test-model",
	}

	result, err := reviewer.Review(context.Background(), pc, nil)
	if err != nil {
		t.Fatalf("Review failed: %v", err)
	}

	if result.PhasesComplete != 6 {
		t.Errorf("expected 6 phases, got %d (errors: %v)", result.PhasesComplete, result.Errors)
	}

	// Phase 1 should parse JSON despite <think> tags
	if result.Phase1 == nil {
		t.Fatal("Phase1 should be parsed despite think tags")
	}
	if result.Phase1.DependencyAssessment != "clean" {
		t.Errorf("Phase1 dependency assessment should be 'clean', got %q", result.Phase1.DependencyAssessment)
	}

	// Thinking should be captured
	if result.Thinking[0] == "" {
		t.Error("Phase 1 thinking should be captured")
	}
	if !strings.Contains(result.Thinking[0], "layered structure") {
		t.Error("Phase 1 thinking should contain the reasoning text")
	}

	// Phase 2 thinking should also be captured
	if result.Thinking[1] == "" {
		t.Error("Phase 2 thinking should be captured")
	}

	// Phase 3 had no think tags — thinking should be empty
	if result.Thinking[2] != "" {
		t.Error("Phase 3 thinking should be empty (no think tags)")
	}
}

// The test helpers (mockProvider, sequenceProvider, etc.) are in stage_test.go.
var _ = time.Now // reference time package
