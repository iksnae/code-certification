package agent_test

import (
	"context"
	"testing"

	"github.com/iksnae/code-certification/internal/agent"
	"github.com/iksnae/code-certification/internal/domain"
)

func TestPrescreenStage_NeedsReview(t *testing.T) {
	mock := &mockProvider{response: `{"needs_review": true, "reason": "borderline", "confidence": 0.7}`}
	stage := agent.NewPrescreenStage(mock, "test-model")

	result, cont, err := stage.Execute(context.Background(), agent.StageInput{
		Unit: domain.NewUnit(domain.NewUnitID("go", "main.go", "main"), domain.UnitTypeFunction),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !cont {
		t.Error("should continue pipeline when review needed")
	}
	if result.TokensUsed == 0 {
		t.Error("should report tokens used")
	}
}

func TestPrescreenStage_SkipsReview(t *testing.T) {
	mock := &mockProvider{response: `{"needs_review": false, "reason": "clean", "confidence": 0.95}`}
	stage := agent.NewPrescreenStage(mock, "test-model")

	_, cont, err := stage.Execute(context.Background(), agent.StageInput{
		Unit: domain.NewUnit(domain.NewUnitID("go", "main.go", "main"), domain.UnitTypeFunction),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cont {
		t.Error("should NOT continue pipeline when review not needed")
	}
}

func TestPrescreenStage_MalformedJSON_FallsThrough(t *testing.T) {
	// Free models often return non-JSON; prescreen should parse loosely
	mock := &mockProvider{response: "Sure! This unit looks fine, no review needed. needs_review: false"}
	stage := agent.NewPrescreenStage(mock, "test-model")

	_, cont, err := stage.Execute(context.Background(), agent.StageInput{
		Unit: domain.NewUnit(domain.NewUnitID("go", "main.go", "main"), domain.UnitTypeFunction),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// Malformed output that mentions "no review" → should skip
	if cont {
		t.Error("loose parse should detect 'no review needed'")
	}
}

func TestPrescreenStage_MalformedJSON_DefaultsToReview(t *testing.T) {
	mock := &mockProvider{response: "I think this code has some issues worth looking at."}
	stage := agent.NewPrescreenStage(mock, "test-model")

	_, cont, err := stage.Execute(context.Background(), agent.StageInput{
		Unit: domain.NewUnit(domain.NewUnitID("go", "main.go", "main"), domain.UnitTypeFunction),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// Ambiguous output → default to review (safer)
	if !cont {
		t.Error("ambiguous output should default to continuing review")
	}
}

func TestReviewStage_ProducesOutput(t *testing.T) {
	mock := &mockProvider{response: "This function has good error handling but lacks input validation."}
	stage := agent.NewReviewStage(mock, "test-model")

	result, cont, err := stage.Execute(context.Background(), agent.StageInput{
		Unit:       domain.NewUnit(domain.NewUnitID("go", "main.go", "main"), domain.UnitTypeFunction),
		SourceCode: "func main() { fmt.Println(\"hello\") }",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !cont {
		t.Error("review stage should always continue")
	}
	if result.ReviewOutput == "" {
		t.Error("review stage should produce output")
	}
}

func TestScoringStage_ParsesScores(t *testing.T) {
	mock := &mockProvider{response: `{"scores":{"correctness":0.9,"maintainability":0.8},"confidence":0.85,"reasoning":"good"}`}
	stage := agent.NewScoringStage(mock, "test-model")

	result, _, err := stage.Execute(context.Background(), agent.StageInput{
		Unit:         domain.NewUnit(domain.NewUnitID("go", "main.go", "main"), domain.UnitTypeFunction),
		ReviewOutput: "good error handling",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Scores["correctness"] != 0.9 {
		t.Errorf("correctness = %f, want 0.9", result.Scores["correctness"])
	}
	if result.Confidence != 0.85 {
		t.Errorf("confidence = %f, want 0.85", result.Confidence)
	}
}

func TestScoringStage_MalformedJSON_FallbackScores(t *testing.T) {
	mock := &mockProvider{response: "The code looks solid, I'd give it about 0.85 overall."}
	stage := agent.NewScoringStage(mock, "test-model")

	result, cont, err := stage.Execute(context.Background(), agent.StageInput{
		Unit: domain.NewUnit(domain.NewUnitID("go", "main.go", "main"), domain.UnitTypeFunction),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !cont {
		t.Error("should continue even with fallback scores")
	}
	// Should produce some default scores, not empty
	if len(result.Scores) == 0 {
		t.Error("fallback should produce default scores")
	}
}

func TestPipeline_QuickStrategy(t *testing.T) {
	mock := &mockProvider{response: `{"needs_review": false, "reason": "clean", "confidence": 0.95}`}
	pipeline := agent.NewPipeline(mock, agent.PipelineConfig{
		Strategy: agent.StrategyQuick,
		Models: domain.ModelAssignments{
			Prescreen: "test-model",
		},
	})

	result, err := pipeline.Run(context.Background(), agent.StageInput{
		Unit: domain.NewUnit(domain.NewUnitID("go", "main.go", "main"), domain.UnitTypeFunction),
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Reviewed {
		t.Error("quick strategy with clean prescreen should not mark as reviewed")
	}
}

func TestPipeline_StandardStrategy(t *testing.T) {
	callCount := 0
	mock := &sequenceProvider{responses: []string{
		`{"needs_review": true, "reason": "borderline", "confidence": 0.6}`,
		"Good error handling, missing input validation on line 5.",
		`{"scores":{"correctness":0.85,"maintainability":0.80},"confidence":0.8,"reasoning":"solid"}`,
	}, callCount: &callCount}

	pipeline := agent.NewPipeline(mock, agent.PipelineConfig{
		Strategy: agent.StrategyStandard,
		Models: domain.ModelAssignments{
			Prescreen: "prescreen-model",
			Review:    "review-model",
			Scoring:   "scoring-model",
		},
	})

	result, err := pipeline.Run(context.Background(), agent.StageInput{
		Unit:       domain.NewUnit(domain.NewUnitID("go", "main.go", "main"), domain.UnitTypeFunction),
		SourceCode: "func main() {}",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.Reviewed {
		t.Error("standard strategy should mark as reviewed")
	}
	if callCount != 3 {
		t.Errorf("standard strategy should make 3 calls, got %d", callCount)
	}
	if result.Scores["correctness"] != 0.85 {
		t.Errorf("correctness = %f, want 0.85", result.Scores["correctness"])
	}
}

func TestPipeline_APIError_GracefulDegradation(t *testing.T) {
	mock := &errorProvider{err: &agent.APIError{StatusCode: 429, Body: "rate limited"}}
	pipeline := agent.NewPipeline(mock, agent.PipelineConfig{
		Strategy: agent.StrategyStandard,
		Models:   domain.ModelAssignments{Prescreen: "test-model"},
	})

	result, err := pipeline.Run(context.Background(), agent.StageInput{
		Unit: domain.NewUnit(domain.NewUnitID("go", "main.go", "main"), domain.UnitTypeFunction),
	})
	// Should not error — graceful degradation
	if err != nil {
		t.Fatalf("pipeline should degrade gracefully, got: %v", err)
	}
	if result.Reviewed {
		t.Error("failed pipeline should not be marked reviewed")
	}
}

func TestCircuitBreaker_OpensAfterFailures(t *testing.T) {
	calls := 0
	mock := &errorProvider{err: &agent.APIError{StatusCode: 500, Body: "server error"}, callCount: &calls}
	cb := agent.NewCircuitBreaker(mock, 3)

	ctx := context.Background()
	req := agent.ChatRequest{Model: "test", Messages: []agent.Message{{Role: "user", Content: "hi"}}}

	// First 3 fail normally
	for i := 0; i < 3; i++ {
		cb.Chat(ctx, req)
	}

	// Circuit should be open now — no more calls to provider
	callsBefore := calls
	_, err := cb.Chat(ctx, req)
	if err == nil {
		t.Error("open circuit should return error")
	}
	if calls != callsBefore {
		t.Error("open circuit should NOT call provider")
	}
}

func TestCircuitBreaker_ClosesAfterSuccess(t *testing.T) {
	// Provider that fails twice then succeeds
	callCount := 0
	mock := &conditionalProvider{
		failUntil: 2,
		response:  "ok",
		callCount: &callCount,
	}
	cb := agent.NewCircuitBreaker(mock, 3)

	ctx := context.Background()
	req := agent.ChatRequest{Model: "test", Messages: []agent.Message{{Role: "user", Content: "hi"}}}

	cb.Chat(ctx, req)              // fail 1
	cb.Chat(ctx, req)              // fail 2
	resp, err := cb.Chat(ctx, req) // success — should reset
	if err != nil {
		t.Fatalf("third call should succeed: %v", err)
	}
	if resp.Content() != "ok" {
		t.Errorf("content = %q, want ok", resp.Content())
	}

	// Circuit should be closed, more calls should work
	_, err = cb.Chat(ctx, req)
	if err != nil {
		t.Errorf("circuit should be closed after success: %v", err)
	}
}

func TestCoordinator_DeduplicatesByFile(t *testing.T) {
	callCount := 0
	mock := &sequenceProvider{responses: []string{
		`{"needs_review": false, "reason": "clean", "confidence": 0.95}`,
		`{"needs_review": false, "reason": "clean", "confidence": 0.95}`,
	}, callCount: &callCount}

	coord := agent.NewCoordinator(mock, agent.CoordinatorConfig{
		Models:      domain.ModelAssignments{Prescreen: "m"},
		Strategy:    agent.StrategyQuick,
		TokenBudget: 10000,
	})

	ctx := context.Background()
	u1 := domain.NewUnit(domain.NewUnitID("go", "main.go", "foo"), domain.UnitTypeFunction)
	u2 := domain.NewUnit(domain.NewUnitID("go", "main.go", "bar"), domain.UnitTypeFunction)

	coord.ReviewUnit(ctx, u1, "func foo() {}", nil)
	coord.ReviewUnit(ctx, u2, "func bar() {}", nil)

	// Same file → only 1 API call
	if callCount != 1 {
		t.Errorf("should deduplicate by file, got %d calls", callCount)
	}
}

func TestCoordinator_RespectsTokenBudget(t *testing.T) {
	mock := &mockProvider{response: `{"needs_review": true}`, tokensUsed: 5000}
	coord := agent.NewCoordinator(mock, agent.CoordinatorConfig{
		Models:      domain.ModelAssignments{Prescreen: "m", Review: "m", Scoring: "m"},
		Strategy:    agent.StrategyStandard,
		TokenBudget: 6000,
	})

	ctx := context.Background()
	u1 := domain.NewUnit(domain.NewUnitID("go", "a.go", "a"), domain.UnitTypeFunction)
	u2 := domain.NewUnit(domain.NewUnitID("go", "b.go", "b"), domain.UnitTypeFunction)

	coord.ReviewUnit(ctx, u1, "func a() {}", nil)
	r2 := coord.ReviewUnit(ctx, u2, "func b() {}", nil)

	// Second file should be skipped — budget exhausted
	if r2.Reviewed {
		t.Error("should skip review when token budget exhausted")
	}
}

func TestPromptRegistry_LoadsAndCaches(t *testing.T) {
	reg := agent.NewPromptRegistry("../../templates/prompts")
	tmpl, err := reg.Get(agent.TaskPrescreen)
	if err != nil {
		t.Fatalf("loading prescreen template: %v", err)
	}

	rendered, err := tmpl.Render(map[string]string{
		"UnitID":          "go://main.go#main",
		"Language":        "go",
		"Path":            "main.go",
		"UnitType":        "function",
		"EvidenceSummary": "lint: clean, tests: 10/10 pass",
	})
	if err != nil {
		t.Fatalf("rendering: %v", err)
	}
	if rendered == "" {
		t.Error("rendered prompt should not be empty")
	}
}

// === Test helpers ===

type mockProvider struct {
	response   string
	tokensUsed int
}

func (m *mockProvider) Chat(_ context.Context, _ agent.ChatRequest) (agent.ChatResponse, error) {
	return agent.ChatResponse{
		Choices: []agent.Choice{{Message: agent.Message{Content: m.response}}},
		Usage:   agent.Usage{TotalTokens: max(m.tokensUsed, 10)},
	}, nil
}
func (m *mockProvider) Name() string { return "mock" }

type sequenceProvider struct {
	responses []string
	callCount *int
}

func (s *sequenceProvider) Chat(_ context.Context, _ agent.ChatRequest) (agent.ChatResponse, error) {
	idx := *s.callCount
	*s.callCount++
	resp := ""
	if idx < len(s.responses) {
		resp = s.responses[idx]
	}
	return agent.ChatResponse{
		Choices: []agent.Choice{{Message: agent.Message{Content: resp}}},
		Usage:   agent.Usage{TotalTokens: 10},
	}, nil
}
func (s *sequenceProvider) Name() string { return "sequence" }

type errorProvider struct {
	err       error
	callCount *int
}

func (e *errorProvider) Chat(_ context.Context, _ agent.ChatRequest) (agent.ChatResponse, error) {
	if e.callCount != nil {
		*e.callCount++
	}
	return agent.ChatResponse{}, e.err
}
func (e *errorProvider) Name() string { return "error" }

type conditionalProvider struct {
	failUntil int
	response  string
	callCount *int
}

func (c *conditionalProvider) Chat(_ context.Context, _ agent.ChatRequest) (agent.ChatResponse, error) {
	call := *c.callCount
	*c.callCount++
	if call < c.failUntil {
		return agent.ChatResponse{}, &agent.APIError{StatusCode: 500, Body: "fail"}
	}
	return agent.ChatResponse{
		Choices: []agent.Choice{{Message: agent.Message{Role: "assistant", Content: c.response}}},
		Usage:   agent.Usage{TotalTokens: 5},
	}, nil
}
func (c *conditionalProvider) Name() string { return "conditional" }
