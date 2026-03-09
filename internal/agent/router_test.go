package agent_test

import (
	"testing"

	"github.com/code-certification/certify/internal/agent"
	"github.com/code-certification/certify/internal/domain"
)

func TestRouter_SelectModel(t *testing.T) {
	assignments := domain.ModelAssignments{
		Prescreen:   "mistralai/mistral-small-3.1-24b-instruct:free",
		Review:      "qwen/qwen3-coder:free",
		Scoring:     "qwen/qwen3-next-80b-a3b-instruct:free",
		Decision:    "openai/gpt-oss-120b:free",
		Remediation: "qwen/qwen3-coder:free",
		Fallback:    "meta-llama/llama-3.3-70b-instruct:free",
	}

	r := agent.NewRouter(assignments)

	tests := []struct {
		task agent.TaskType
		want string
	}{
		{agent.TaskPrescreen, "mistralai/mistral-small-3.1-24b-instruct:free"},
		{agent.TaskReview, "qwen/qwen3-coder:free"},
		{agent.TaskScoring, "qwen/qwen3-next-80b-a3b-instruct:free"},
		{agent.TaskDecision, "openai/gpt-oss-120b:free"},
		{agent.TaskRemediation, "qwen/qwen3-coder:free"},
	}

	for _, tt := range tests {
		model := r.ModelFor(tt.task)
		if model != tt.want {
			t.Errorf("ModelFor(%s) = %q, want %q", tt.task, model, tt.want)
		}
	}
}

func TestRouter_Fallback(t *testing.T) {
	assignments := domain.ModelAssignments{
		Fallback: "meta-llama/llama-3.3-70b-instruct:free",
	}

	r := agent.NewRouter(assignments)

	// All task types with empty assignment should fall back
	model := r.ModelFor(agent.TaskReview)
	if model != "meta-llama/llama-3.3-70b-instruct:free" {
		t.Errorf("empty review should fallback, got %q", model)
	}
}

func TestRouter_EmptyFallback(t *testing.T) {
	r := agent.NewRouter(domain.ModelAssignments{})

	model := r.ModelFor(agent.TaskReview)
	if model != "" {
		t.Errorf("no assignments should return empty, got %q", model)
	}
}
