// Package agent handles agent-assisted review via LLM providers.
package agent

import "fmt"

// TaskType identifies the kind of agent review task.
type TaskType int

const (
	TaskPrescreen   TaskType = iota // Quick pre-filter
	TaskReview                      // Detailed code review
	TaskScoring                     // Dimension scoring
	TaskDecision                    // Status determination with reasoning
	TaskRemediation                 // Remediation suggestion generation
)

var taskTypeStrings = map[TaskType]string{
	TaskPrescreen:   "prescreen",
	TaskReview:      "review",
	TaskScoring:     "scoring",
	TaskDecision:    "decision",
	TaskRemediation: "remediation",
}

// String returns the string representation of a TaskType.
func (t TaskType) String() string {
	if s, ok := taskTypeStrings[t]; ok {
		return s
	}
	return fmt.Sprintf("TaskType(%d)", t)
}

// Message represents a chat message (OpenAI-compatible).
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatRequest is an OpenAI-compatible chat completion request.
type ChatRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature,omitempty"`
	MaxTokens   int       `json:"max_tokens,omitempty"`
	// ResponseFormat can request structured JSON output
	ResponseFormat *ResponseFormat `json:"response_format,omitempty"`
}

// ResponseFormat requests a specific output format.
type ResponseFormat struct {
	Type   string `json:"type"`             // "json_object" or "json_schema"
	Schema any    `json:"json_schema,omitempty"`
}

// ChatResponse is an OpenAI-compatible chat completion response.
type ChatResponse struct {
	ID      string   `json:"id"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
}

// Content returns the first choice's message content, or empty string.
func (r ChatResponse) Content() string {
	if len(r.Choices) == 0 {
		return ""
	}
	return r.Choices[0].Message.Content
}

// Choice represents a single completion choice.
type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

// Usage reports token consumption.
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// ModelConfig describes a model's capabilities.
type ModelConfig struct {
	ID                string `json:"id"`
	ContextWindow     int    `json:"context_window"`
	MaxOutput         int    `json:"max_output"`
	StructuredOutputs bool   `json:"structured_outputs"`
	ToolCalling       bool   `json:"tool_calling"`
	Reasoning         bool   `json:"reasoning"`
}
