package agent_test

import (
	"context"
	"testing"

	"github.com/code-certification/certify/internal/agent"
)

func TestFallbackProvider_TriesModelsInOrder(t *testing.T) {
	calls := []string{}
	p := agent.NewFallbackProvider([]agent.Provider{
		&trackingProvider{name: "p1", err: &agent.APIError{StatusCode: 429}, calls: &calls},
		&trackingProvider{name: "p2", response: "hello", calls: &calls},
		&trackingProvider{name: "p3", response: "unused", calls: &calls},
	})

	resp, err := p.Chat(context.Background(), agent.ChatRequest{
		Model:    "any",
		Messages: []agent.Message{{Role: "user", Content: "hi"}},
	})
	if err != nil {
		t.Fatalf("should succeed with fallback: %v", err)
	}
	if resp.Content() != "hello" {
		t.Errorf("content = %q, want hello", resp.Content())
	}
	if len(calls) != 2 {
		t.Errorf("should try 2 providers, got %d: %v", len(calls), calls)
	}
}

func TestFallbackProvider_AllFail(t *testing.T) {
	p := agent.NewFallbackProvider([]agent.Provider{
		&trackingProvider{name: "p1", err: &agent.APIError{StatusCode: 429}, calls: new([]string)},
		&trackingProvider{name: "p2", err: &agent.APIError{StatusCode: 429}, calls: new([]string)},
	})

	_, err := p.Chat(context.Background(), agent.ChatRequest{
		Model:    "any",
		Messages: []agent.Message{{Role: "user", Content: "hi"}},
	})
	if err == nil {
		t.Error("should fail when all providers fail")
	}
}

func TestFallbackProvider_SkipsAuthErrors(t *testing.T) {
	calls := []string{}
	p := agent.NewFallbackProvider([]agent.Provider{
		&trackingProvider{name: "p1", err: &agent.APIError{StatusCode: 401}, calls: &calls},
		&trackingProvider{name: "p2", response: "ok", calls: &calls},
	})

	_, err := p.Chat(context.Background(), agent.ChatRequest{
		Model:    "any",
		Messages: []agent.Message{{Role: "user", Content: "hi"}},
	})
	// Auth error = don't fallback, it's a config problem
	if err == nil {
		t.Error("auth errors should not trigger fallback")
	}
	if len(calls) != 1 {
		t.Errorf("should stop at auth error, got %d calls", len(calls))
	}
}

func TestModelChain_BuildsProviderPerModel(t *testing.T) {
	chain := agent.NewModelChain(
		"https://openrouter.ai/api/v1",
		"test-key",
		"http://test",
		"test",
		[]string{
			"qwen/qwen3-coder:free",
			"meta-llama/llama-3.3-70b-instruct:free",
			"mistralai/mistral-small-3.1-24b-instruct:free",
		},
	)

	if chain.Name() != "model-chain" {
		t.Errorf("name = %q", chain.Name())
	}
}

func TestAdaptiveMessage_SystemSupported(t *testing.T) {
	msgs := agent.AdaptiveMessages("You are a reviewer.", "Review this code.", true)
	if len(msgs) != 2 {
		t.Fatalf("expected 2 messages, got %d", len(msgs))
	}
	if msgs[0].Role != "system" {
		t.Errorf("first message role = %q, want system", msgs[0].Role)
	}
}

func TestAdaptiveMessage_SystemNotSupported(t *testing.T) {
	msgs := agent.AdaptiveMessages("You are a reviewer.", "Review this code.", false)
	if len(msgs) != 1 {
		t.Fatalf("expected 1 message, got %d", len(msgs))
	}
	if msgs[0].Role != "user" {
		t.Errorf("role = %q, want user", msgs[0].Role)
	}
	if !contains(msgs[0].Content, "You are a reviewer.") {
		t.Error("should contain system instruction in user message")
	}
	if !contains(msgs[0].Content, "Review this code.") {
		t.Error("should contain user content")
	}
}

func contains(s, sub string) bool {
	return len(s) >= len(sub) && (s == sub || len(s) > 0 && containsStr(s, sub))
}

func containsStr(s, sub string) bool {
	for i := 0; i <= len(s)-len(sub); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}

// === Helpers ===

type trackingProvider struct {
	name     string
	response string
	err      error
	calls    *[]string
}

func (p *trackingProvider) Chat(_ context.Context, req agent.ChatRequest) (agent.ChatResponse, error) {
	*p.calls = append(*p.calls, p.name)
	if p.err != nil {
		return agent.ChatResponse{}, p.err
	}
	return agent.ChatResponse{
		Choices: []agent.Choice{{Message: agent.Message{Content: p.response}}},
		Usage:   agent.Usage{TotalTokens: 10},
	}, nil
}

func (p *trackingProvider) Name() string { return p.name }
