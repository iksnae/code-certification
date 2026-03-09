package agent_test

import (
	"encoding/json"
	"testing"

	"github.com/iksnae/code-certification/internal/agent"
)

func TestPrescreenResponse_Valid(t *testing.T) {
	raw := `{"needs_review": true, "reason": "missing test coverage", "confidence": 0.8}`
	var resp agent.PrescreenResponse
	if err := json.Unmarshal([]byte(raw), &resp); err != nil {
		t.Fatalf("parse error: %v", err)
	}
	if !resp.NeedsReview {
		t.Error("needs_review should be true")
	}
	if resp.Confidence != 0.8 {
		t.Errorf("confidence = %f, want 0.8", resp.Confidence)
	}
}

func TestScoringResponse_Valid(t *testing.T) {
	raw := `{
		"scores": {
			"correctness": 0.9,
			"maintainability": 0.85,
			"readability": 0.8,
			"testability": 0.75,
			"security": 0.95,
			"architectural_fitness": 0.8,
			"operational_quality": 0.7,
			"performance_appropriateness": 0.85,
			"change_risk": 0.6
		},
		"confidence": 0.7,
		"reasoning": "Well-tested but high change risk"
	}`
	var resp agent.ScoringResponse
	if err := json.Unmarshal([]byte(raw), &resp); err != nil {
		t.Fatalf("parse error: %v", err)
	}
	if resp.Scores["correctness"] != 0.9 {
		t.Errorf("correctness = %f, want 0.9", resp.Scores["correctness"])
	}
	if resp.Confidence != 0.7 {
		t.Errorf("confidence = %f, want 0.7", resp.Confidence)
	}
}

func TestDecisionResponse_Valid(t *testing.T) {
	raw := `{
		"status": "certified_with_observations",
		"reasoning": "Minor readability issues",
		"actions": ["add doc comments", "reduce function length"]
	}`
	var resp agent.DecisionResponse
	if err := json.Unmarshal([]byte(raw), &resp); err != nil {
		t.Fatalf("parse error: %v", err)
	}
	if resp.Status != "certified_with_observations" {
		t.Errorf("status = %q", resp.Status)
	}
	if len(resp.Actions) != 2 {
		t.Errorf("actions = %d, want 2", len(resp.Actions))
	}
}

func TestRemediationResponse_Valid(t *testing.T) {
	raw := `{
		"steps": [
			{"priority": 1, "dimension": "correctness", "description": "fix nil check", "effort": "low"},
			{"priority": 2, "dimension": "testability", "description": "add unit test", "effort": "medium"}
		]
	}`
	var resp agent.RemediationResponse
	if err := json.Unmarshal([]byte(raw), &resp); err != nil {
		t.Fatalf("parse error: %v", err)
	}
	if len(resp.Steps) != 2 {
		t.Errorf("steps = %d, want 2", len(resp.Steps))
	}
	if resp.Steps[0].Priority != 1 {
		t.Errorf("step[0].priority = %d, want 1", resp.Steps[0].Priority)
	}
}
