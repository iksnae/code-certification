package agent_test

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/code-certification/certify/internal/agent"
)

func TestOpenRouter_SuccessfulChat(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request
		if r.Method != "POST" {
			t.Errorf("method = %q, want POST", r.Method)
		}
		if r.URL.Path != "/chat/completions" {
			t.Errorf("path = %q, want /chat/completions", r.URL.Path)
		}

		// Verify body is valid
		body, _ := io.ReadAll(r.Body)
		var req agent.ChatRequest
		if err := json.Unmarshal(body, &req); err != nil {
			t.Errorf("invalid request body: %v", err)
		}

		resp := agent.ChatResponse{
			ID:    "gen-test",
			Model: req.Model,
			Choices: []agent.Choice{
				{Index: 0, Message: agent.Message{Role: "assistant", Content: "LGTM"}, FinishReason: "stop"},
			},
			Usage: agent.Usage{PromptTokens: 10, CompletionTokens: 5, TotalTokens: 15},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	p := agent.NewOpenRouterProvider(server.URL, "test-key", "", "")

	resp, err := p.Chat(context.Background(), agent.ChatRequest{
		Model:    "test-model",
		Messages: []agent.Message{{Role: "user", Content: "hello"}},
	})
	if err != nil {
		t.Fatalf("Chat() error: %v", err)
	}
	if resp.Content() != "LGTM" {
		t.Errorf("content = %q, want LGTM", resp.Content())
	}
}

func TestOpenRouter_Headers(t *testing.T) {
	var gotHeaders http.Header
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotHeaders = r.Header
		resp := agent.ChatResponse{
			Choices: []agent.Choice{{Message: agent.Message{Content: "ok"}}},
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	p := agent.NewOpenRouterProvider(server.URL, "sk-test-123", "https://example.com", "CertifyTest")
	p.Chat(context.Background(), agent.ChatRequest{
		Model:    "test",
		Messages: []agent.Message{{Role: "user", Content: "hi"}},
	})

	if got := gotHeaders.Get("Authorization"); got != "Bearer sk-test-123" {
		t.Errorf("Authorization = %q", got)
	}
	if got := gotHeaders.Get("HTTP-Referer"); got != "https://example.com" {
		t.Errorf("HTTP-Referer = %q", got)
	}
	if got := gotHeaders.Get("X-Title"); got != "CertifyTest" {
		t.Errorf("X-Title = %q", got)
	}
	if got := gotHeaders.Get("Content-Type"); got != "application/json" {
		t.Errorf("Content-Type = %q", got)
	}
}

func TestOpenRouter_401(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(401)
		w.Write([]byte(`{"error": {"message": "invalid api key"}}`))
	}))
	defer server.Close()

	p := agent.NewOpenRouterProvider(server.URL, "bad-key", "", "")
	_, err := p.Chat(context.Background(), agent.ChatRequest{
		Model:    "test",
		Messages: []agent.Message{{Role: "user", Content: "hi"}},
	})
	if err == nil {
		t.Fatal("401 should return error")
	}
}

func TestOpenRouter_429(t *testing.T) {
	calls := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		if calls <= 2 {
			w.WriteHeader(429)
			w.Write([]byte(`{"error": {"message": "rate limited"}}`))
			return
		}
		resp := agent.ChatResponse{
			Choices: []agent.Choice{{Message: agent.Message{Content: "ok"}}},
		}
		json.NewEncoder(w).Encode(resp)
	}))
	defer server.Close()

	p := agent.NewOpenRouterProvider(server.URL, "key", "", "")
	resp, err := p.Chat(context.Background(), agent.ChatRequest{
		Model:    "test",
		Messages: []agent.Message{{Role: "user", Content: "hi"}},
	})
	if err != nil {
		t.Fatalf("should retry and succeed: %v", err)
	}
	if resp.Content() != "ok" {
		t.Errorf("content = %q, want ok", resp.Content())
	}
	if calls != 3 {
		t.Errorf("expected 3 calls (2 retries + success), got %d", calls)
	}
}

func TestOpenRouter_Name(t *testing.T) {
	p := agent.NewOpenRouterProvider("http://example.com", "key", "", "")
	if p.Name() != "openrouter" {
		t.Errorf("Name() = %q, want openrouter", p.Name())
	}
}

func TestOpenRouter_MissingAPIKey(t *testing.T) {
	p := agent.NewOpenRouterProvider("http://example.com", "", "", "")
	_, err := p.Chat(context.Background(), agent.ChatRequest{
		Model:    "test",
		Messages: []agent.Message{{Role: "user", Content: "hi"}},
	})
	if err == nil {
		t.Fatal("missing API key should return error")
	}
}
