package agent_test

import (
	"encoding/json"
	"testing"

	"github.com/code-certification/certify/internal/agent"
)

func TestTaskType_String(t *testing.T) {
	tests := []struct {
		tt   agent.TaskType
		want string
	}{
		{agent.TaskPrescreen, "prescreen"},
		{agent.TaskReview, "review"},
		{agent.TaskScoring, "scoring"},
		{agent.TaskDecision, "decision"},
		{agent.TaskRemediation, "remediation"},
	}
	for _, tt := range tests {
		if got := tt.tt.String(); got != tt.want {
			t.Errorf("TaskType(%d).String() = %q, want %q", tt.tt, got, tt.want)
		}
	}
}

func TestChatRequest_JSON(t *testing.T) {
	req := agent.ChatRequest{
		Model: "qwen/qwen3-coder:free",
		Messages: []agent.Message{
			{Role: "system", Content: "You are a code reviewer."},
			{Role: "user", Content: "Review this function."},
		},
		Temperature: 0.3,
		MaxTokens:   2048,
	}

	data, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("marshal error: %v", err)
	}

	var parsed map[string]any
	if err := json.Unmarshal(data, &parsed); err != nil {
		t.Fatal(err)
	}

	if parsed["model"] != "qwen/qwen3-coder:free" {
		t.Errorf("model = %v", parsed["model"])
	}
	msgs := parsed["messages"].([]any)
	if len(msgs) != 2 {
		t.Errorf("messages count = %d, want 2", len(msgs))
	}
}

func TestChatResponse_Parse(t *testing.T) {
	raw := `{
		"id": "gen-123",
		"model": "qwen/qwen3-coder:free",
		"choices": [
			{
				"index": 0,
				"message": {
					"role": "assistant",
					"content": "This function looks correct."
				},
				"finish_reason": "stop"
			}
		],
		"usage": {
			"prompt_tokens": 100,
			"completion_tokens": 20,
			"total_tokens": 120
		}
	}`

	var resp agent.ChatResponse
	if err := json.Unmarshal([]byte(raw), &resp); err != nil {
		t.Fatalf("unmarshal error: %v", err)
	}

	if resp.ID != "gen-123" {
		t.Errorf("ID = %q", resp.ID)
	}
	if len(resp.Choices) != 1 {
		t.Fatalf("choices = %d, want 1", len(resp.Choices))
	}
	if resp.Choices[0].Message.Content != "This function looks correct." {
		t.Errorf("content = %q", resp.Choices[0].Message.Content)
	}
	if resp.Usage.TotalTokens != 120 {
		t.Errorf("total_tokens = %d", resp.Usage.TotalTokens)
	}
}

func TestChatResponse_Empty(t *testing.T) {
	var resp agent.ChatResponse
	content := resp.Content()
	if content != "" {
		t.Errorf("empty response content = %q, want empty", content)
	}
}

func TestModelConfig_Basics(t *testing.T) {
	mc := agent.ModelConfig{
		ID:                "qwen/qwen3-coder:free",
		ContextWindow:     131072,
		MaxOutput:         8192,
		StructuredOutputs: true,
		ToolCalling:       true,
	}

	if !mc.StructuredOutputs {
		t.Error("should support structured outputs")
	}
	if !mc.ToolCalling {
		t.Error("should support tool calling")
	}
}
