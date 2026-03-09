package agent_test

import (
	"context"
	"testing"

	"github.com/code-certification/certify/internal/agent"
	"github.com/code-certification/certify/internal/domain"
)

func TestStageResult_TracksModel(t *testing.T) {
	mock := &mockProvider{response: `{"needs_review": true, "reason": "test", "confidence": 0.8}`}
	stage := agent.NewPrescreenStage(mock, "qwen/qwen3-coder:free")

	result, _, err := stage.Execute(context.Background(), agent.StageInput{
		Unit: domain.NewUnit(domain.NewUnitID("go", "main.go", "main"), domain.UnitTypeFunction),
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(result.ModelsUsed) == 0 {
		t.Error("stage should track which model was used")
	}
}

func TestPipeline_CollectsAllModels(t *testing.T) {
	callCount := 0
	mock := &sequenceProvider{responses: []string{
		`{"needs_review": true, "reason": "test", "confidence": 0.6}`,
		"Good code, minor issues.",
		`{"scores":{"correctness":0.9},"confidence":0.8,"reasoning":"ok"}`,
	}, callCount: &callCount}

	pipeline := agent.NewPipeline(mock, agent.PipelineConfig{
		Strategy: agent.StrategyStandard,
		Models: domain.ModelAssignments{
			Prescreen: "model-a",
			Review:    "model-b",
			Scoring:   "model-c",
		},
	})

	result, err := pipeline.Run(context.Background(), agent.StageInput{
		Unit:       domain.NewUnit(domain.NewUnitID("go", "main.go", "main"), domain.UnitTypeFunction),
		SourceCode: "func main() {}",
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(result.ModelsUsed) != 3 {
		t.Errorf("should track 3 models, got %d: %v", len(result.ModelsUsed), result.ModelsUsed)
	}
}

func TestReviewResult_ToEvidence_IncludesAttribution(t *testing.T) {
	result := agent.ReviewResult{
		Reviewed:   true,
		Status:     "certified",
		Confidence: 0.85,
		ModelsUsed: []string{"qwen/qwen3-coder:free", "mistralai/mistral-nemo"},
	}

	ev := result.ToEvidence()
	if ev.Source != "agent:qwen/qwen3-coder:free,mistralai/mistral-nemo" {
		t.Errorf("source = %q, want agent attribution with model names", ev.Source)
	}
}

func TestReviewResult_ToEvidence_SingleModel(t *testing.T) {
	result := agent.ReviewResult{
		Reviewed:   true,
		Status:     "certified",
		Confidence: 0.85,
		ModelsUsed: []string{"mistralai/mistral-nemo"},
	}

	ev := result.ToEvidence()
	if ev.Source != "agent:mistralai/mistral-nemo" {
		t.Errorf("source = %q, want single model attribution", ev.Source)
	}
}

func TestReviewResult_ToEvidence_NoModels(t *testing.T) {
	result := agent.ReviewResult{
		Reviewed:   true,
		Status:     "certified",
		Confidence: 0.85,
	}

	ev := result.ToEvidence()
	if ev.Source != "agent" {
		t.Errorf("source = %q, want plain agent", ev.Source)
	}
}

func TestCoordinator_PropagatesAttribution(t *testing.T) {
	callCount := 0
	mock := &sequenceProvider{responses: []string{
		`{"needs_review": false, "reason": "clean", "confidence": 0.95}`,
	}, callCount: &callCount}

	coord := agent.NewCoordinator(mock, agent.CoordinatorConfig{
		Models:      domain.ModelAssignments{Prescreen: "mistralai/mistral-nemo"},
		Strategy:    agent.StrategyQuick,
		TokenBudget: 10000,
	})

	result := coord.ReviewUnit(
		context.Background(),
		domain.NewUnit(domain.NewUnitID("go", "main.go", "main"), domain.UnitTypeFunction),
		"func main() {}",
		nil,
	)

	if len(result.ModelsUsed) == 0 {
		t.Error("coordinator should propagate model attribution")
	}
}

func TestModelPinnedProvider_ReportsModelInResponse(t *testing.T) {
	// The ChatResponse.Model field should reflect which model actually served
	chain := agent.NewModelChain(
		"https://example.com/v1", "key", "", "",
		[]string{"model-a", "model-b"},
	)
	// Can't actually call without a server, but verify construction
	if chain.Name() != "model-chain" {
		t.Errorf("name = %q", chain.Name())
	}
}

func TestCertificationRecord_IncludesModelAttribution(t *testing.T) {
	// When agent evidence is present, the record should credit the models
	ev := agent.ReviewResult{
		Reviewed:   true,
		ModelsUsed: []string{"qwen/qwen3-coder:free"},
		Confidence: 0.9,
		Status:     "certified",
	}.ToEvidence()

	// The evidence source carries the attribution
	want := "agent:qwen/qwen3-coder:free"
	if ev.Source != want {
		t.Errorf("evidence source = %q, want %q", ev.Source, want)
	}
}
