package agent_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/iksnae/code-certification/internal/agent"
	"github.com/iksnae/code-certification/internal/domain"
)

func mockServer(responses map[int]string) *httptest.Server {
	call := 0
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		content, ok := responses[call]
		if !ok {
			content = `{"needs_review": false, "reason": "default", "confidence": 0.9}`
		}
		call++
		resp := agent.ChatResponse{
			Choices: []agent.Choice{{Message: agent.Message{Content: content}}},
			Usage:   agent.Usage{TotalTokens: 50},
		}
		json.NewEncoder(w).Encode(resp)
	}))
}

func TestReviewer_PrescreenSkips(t *testing.T) {
	server := mockServer(map[int]string{
		0: `{"needs_review": false, "reason": "all checks pass", "confidence": 0.95}`,
	})
	defer server.Close()

	provider := agent.NewOpenRouterProvider(server.URL, "test-key", "", "")
	reviewer := agent.NewReviewer(provider, agent.NewRouter(domain.ModelAssignments{
		Prescreen: "test-model",
	}))

	result, err := reviewer.Review(context.Background(), agent.ReviewInput{
		Unit:     domain.NewUnit(domain.NewUnitID("go", "main.go", "main"), domain.UnitTypeFunction),
		Evidence: nil,
	})
	if err != nil {
		t.Fatalf("Review() error: %v", err)
	}

	if result.Reviewed {
		t.Error("prescreen said no review needed, should not be marked as reviewed")
	}
}

func TestReviewer_FullPipeline(t *testing.T) {
	server := mockServer(map[int]string{
		0: `{"needs_review": true, "reason": "borderline scores", "confidence": 0.7}`,
		1: `Code review: The function has good structure but lacks error handling for edge cases.`,
		2: `{"scores": {"correctness": 0.8, "maintainability": 0.75}, "confidence": 0.7, "reasoning": "decent"}`,
		3: `{"status": "certified_with_observations", "reasoning": "minor issues", "actions": ["add error handling"]}`,
		4: `{"steps": [{"priority": 1, "dimension": "correctness", "description": "add nil check", "effort": "low"}]}`,
	})
	defer server.Close()

	provider := agent.NewOpenRouterProvider(server.URL, "test-key", "", "")
	reviewer := agent.NewReviewer(provider, agent.NewRouter(domain.ModelAssignments{
		Prescreen:   "test",
		Review:      "test",
		Scoring:     "test",
		Decision:    "test",
		Remediation: "test",
	}))

	result, err := reviewer.Review(context.Background(), agent.ReviewInput{
		Unit: domain.NewUnit(domain.NewUnitID("go", "main.go", "main"), domain.UnitTypeFunction),
	})
	if err != nil {
		t.Fatalf("Review() error: %v", err)
	}

	if !result.Reviewed {
		t.Error("should be marked as reviewed")
	}
	if result.ReviewOutput == "" {
		t.Error("should have review output")
	}
	if len(result.Scores) == 0 {
		t.Error("should have scores")
	}
}

func TestReviewer_NoProvider(t *testing.T) {
	reviewer := agent.NewReviewer(nil, nil)
	result, err := reviewer.Review(context.Background(), agent.ReviewInput{
		Unit: domain.NewUnit(domain.NewUnitID("go", "main.go", "main"), domain.UnitTypeFunction),
	})
	if err != nil {
		t.Fatalf("nil provider should not error: %v", err)
	}
	if result.Reviewed {
		t.Error("nil provider should not produce review")
	}
}

func TestReviewResult_ToEvidence(t *testing.T) {
	result := agent.ReviewResult{
		Reviewed:     true,
		ReviewOutput: "LGTM",
		Scores:       map[string]float64{"correctness": 0.9},
		Confidence:   0.8,
	}

	ev := result.ToEvidence()
	if ev.Kind != domain.EvidenceKindAgentReview {
		t.Errorf("Kind = %v, want agent_review", ev.Kind)
	}
	if ev.Source != "agent" {
		t.Errorf("Source = %q, want agent", ev.Source)
	}
}

func TestReviewResult_ToEvidence_Metrics(t *testing.T) {
	result := agent.ReviewResult{
		Reviewed:   true,
		Scores:     map[string]float64{"correctness": 0.9, "maintainability": 0.75},
		Confidence: 0.85,
		TokensUsed: 500,
	}

	ev := result.ToEvidence()
	if ev.Metrics == nil {
		t.Fatal("Metrics should not be nil")
	}
	if ev.Metrics["correctness"] != 0.9 {
		t.Errorf("correctness = %f, want 0.9", ev.Metrics["correctness"])
	}
	if ev.Metrics["maintainability"] != 0.75 {
		t.Errorf("maintainability = %f, want 0.75", ev.Metrics["maintainability"])
	}
	if ev.Metrics["confidence"] != 0.85 {
		t.Errorf("confidence = %f, want 0.85", ev.Metrics["confidence"])
	}
	if ev.Metrics["tokens_used"] != 500 {
		t.Errorf("tokens_used = %f, want 500", ev.Metrics["tokens_used"])
	}
}

func TestReviewResult_ToPrescreenEvidence_Metrics(t *testing.T) {
	result := agent.ReviewResult{
		Prescreened: true,
		Confidence:  0.95,
	}

	ev := result.ToPrescreenEvidence()
	if ev.Metrics == nil {
		t.Fatal("Metrics should not be nil")
	}
	if ev.Metrics["confidence"] != 0.95 {
		t.Errorf("confidence = %f, want 0.95", ev.Metrics["confidence"])
	}
}
