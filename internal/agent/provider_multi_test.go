package agent

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// --- OpenRouterProvider local mode (no auth) ---

func TestOpenRouterProvider_LocalNoAuth(t *testing.T) {
	// Local provider should work without API key
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify no Authorization header sent when key is empty
		auth := r.Header.Get("Authorization")
		if auth != "" {
			t.Errorf("local provider should not send Authorization header, got %q", auth)
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"choices":[{"message":{"content":"ok"}}],"usage":{"total_tokens":5}}`)
	}))
	defer srv.Close()

	p := NewLocalProvider(srv.URL, "test-local")
	resp, err := p.Chat(context.Background(), ChatRequest{
		Model:    "qwen2.5-coder:7b",
		Messages: []Message{{Role: "user", Content: "hello"}},
	})
	if err != nil {
		t.Fatalf("local provider Chat() error: %v", err)
	}
	if resp.Content() != "ok" {
		t.Errorf("Content() = %q, want ok", resp.Content())
	}
}

func TestOpenRouterProvider_LocalProviderName(t *testing.T) {
	p := NewLocalProvider("http://localhost:11434/v1", "ollama")
	if p.Name() != "ollama" {
		t.Errorf("Name() = %q, want ollama", p.Name())
	}
}

// --- FallbackProvider: connection errors should fall through ---

func TestFallbackProvider_ConnectionErrorFallsThrough(t *testing.T) {
	// First provider: connection refused (simulated)
	failProvider := &mockProvider{
		err: fmt.Errorf("dial tcp 127.0.0.1:11434: connection refused"),
	}
	// Second provider: succeeds
	okProvider := &mockProvider{
		resp: ChatResponse{
			Choices: []Choice{{Message: Message{Content: "from fallback"}}},
			Usage:   Usage{TotalTokens: 10},
		},
	}

	fb := NewFallbackProvider([]Provider{failProvider, okProvider})
	resp, err := fb.Chat(context.Background(), ChatRequest{})
	if err != nil {
		t.Fatalf("FallbackProvider should fall through connection error, got: %v", err)
	}
	if resp.Content() != "from fallback" {
		t.Errorf("Content() = %q, want 'from fallback'", resp.Content())
	}
}

// --- DetectProviders ---

func TestDetectProviders_None(t *testing.T) {
	clearProviderEnvVars()
	defer clearProviderEnvVars()

	providers := DetectProviders()
	// Should have no cloud providers (local providers depend on running servers)
	for _, p := range providers {
		if !p.Local {
			t.Errorf("unexpected cloud provider detected: %s", p.Name)
		}
	}
}

func TestDetectProviders_OpenRouter(t *testing.T) {
	clearProviderEnvVars()
	defer clearProviderEnvVars()
	os.Setenv("OPENROUTER_API_KEY", "sk-or-test")

	providers := DetectProviders()
	found := false
	for _, p := range providers {
		if p.Name == "openrouter" {
			found = true
			if p.APIKey != "sk-or-test" {
				t.Errorf("APIKey = %q, want sk-or-test", p.APIKey)
			}
			if len(p.Models) == 0 {
				t.Error("OpenRouter should have default models")
			}
		}
	}
	if !found {
		t.Error("OpenRouter not detected despite OPENROUTER_API_KEY being set")
	}
}

func TestDetectProviders_Groq(t *testing.T) {
	clearProviderEnvVars()
	defer clearProviderEnvVars()
	os.Setenv("GROQ_API_KEY", "gsk-test")

	providers := DetectProviders()
	found := false
	for _, p := range providers {
		if p.Name == "groq" {
			found = true
			if p.APIKey != "gsk-test" {
				t.Errorf("APIKey = %q, want gsk-test", p.APIKey)
			}
			if p.BaseURL != "https://api.groq.com/openai/v1" {
				t.Errorf("BaseURL = %q", p.BaseURL)
			}
		}
	}
	if !found {
		t.Error("Groq not detected despite GROQ_API_KEY being set")
	}
}

func TestDetectProviders_OllamaEnv(t *testing.T) {
	clearProviderEnvVars()
	defer clearProviderEnvVars()
	os.Setenv("OLLAMA_HOST", "http://myhost:11434")

	providers := DetectProviders()
	found := false
	for _, p := range providers {
		if p.Name == "ollama" {
			found = true
			if !p.Local {
				t.Error("Ollama should be marked as local")
			}
			if p.BaseURL != "http://myhost:11434/v1" {
				t.Errorf("BaseURL = %q, want http://myhost:11434/v1", p.BaseURL)
			}
		}
	}
	if !found {
		t.Error("Ollama not detected despite OLLAMA_HOST being set")
	}
}

func TestDetectProviders_LMStudioEnv(t *testing.T) {
	clearProviderEnvVars()
	defer clearProviderEnvVars()
	os.Setenv("LM_STUDIO_URL", "http://localhost:1234")

	providers := DetectProviders()
	found := false
	for _, p := range providers {
		if p.Name == "lmstudio" {
			found = true
			if !p.Local {
				t.Error("LM Studio should be marked as local")
			}
			if p.BaseURL != "http://localhost:1234/v1" {
				t.Errorf("BaseURL = %q, want http://localhost:1234/v1", p.BaseURL)
			}
		}
	}
	if !found {
		t.Error("LM Studio not detected despite LM_STUDIO_URL being set")
	}
}

func TestDetectProviders_Priority(t *testing.T) {
	clearProviderEnvVars()
	defer clearProviderEnvVars()
	os.Setenv("OPENROUTER_API_KEY", "sk-or-test")
	os.Setenv("GROQ_API_KEY", "gsk-test")
	os.Setenv("OLLAMA_HOST", "http://localhost:11434")

	providers := DetectProviders()
	// Cloud providers should come before local
	if len(providers) < 3 {
		t.Fatalf("expected at least 3 providers, got %d", len(providers))
	}

	// Verify ordering: cloud first, local last
	cloudSeen := false
	localSeen := false
	for _, p := range providers {
		if p.Local {
			localSeen = true
		} else {
			if localSeen {
				t.Errorf("cloud provider %s appeared after local provider", p.Name)
			}
			cloudSeen = true
		}
	}
	if !cloudSeen {
		t.Error("no cloud providers detected")
	}
	if !localSeen {
		t.Error("no local providers detected")
	}
}

func TestDetectProviders_OllamaLiveProbe(t *testing.T) {
	clearProviderEnvVars()
	defer clearProviderEnvVars()

	// Start a fake Ollama server
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprint(w, "Ollama is running")
	}))
	defer srv.Close()

	os.Setenv("OLLAMA_HOST", srv.URL)
	providers := DetectProviders()
	found := false
	for _, p := range providers {
		if p.Name == "ollama" {
			found = true
		}
	}
	if !found {
		t.Error("Ollama with live server should be detected")
	}
}

// --- OpenAI detection ---

func TestDetectProviders_OpenAI(t *testing.T) {
	clearProviderEnvVars()
	defer clearProviderEnvVars()
	os.Setenv("OPENAI_API_KEY", "sk-test-openai")

	providers := DetectProviders()
	found := false
	for _, p := range providers {
		if p.Name == "openai" {
			found = true
			if p.APIKey != "sk-test-openai" {
				t.Errorf("APIKey = %q, want sk-test-openai", p.APIKey)
			}
			if p.BaseURL != "https://api.openai.com/v1" {
				t.Errorf("BaseURL = %q, want https://api.openai.com/v1", p.BaseURL)
			}
			if len(p.Models) == 0 {
				t.Error("OpenAI should have default models")
			}
			if p.Local {
				t.Error("OpenAI should not be local")
			}
		}
	}
	if !found {
		t.Error("OpenAI not detected despite OPENAI_API_KEY being set")
	}
}

func TestDetectProviders_OpenAIPriority(t *testing.T) {
	clearProviderEnvVars()
	defer clearProviderEnvVars()
	os.Setenv("OPENROUTER_API_KEY", "sk-or-test")
	os.Setenv("OPENAI_API_KEY", "sk-openai-test")

	providers := DetectProviders()
	// Both should be detected; OpenRouter first (more models, free tier)
	if len(providers) < 2 {
		t.Fatalf("expected at least 2 cloud providers, got %d", len(providers))
	}
	cloudNames := []string{}
	for _, p := range providers {
		if !p.Local {
			cloudNames = append(cloudNames, p.Name)
		}
	}
	if len(cloudNames) < 2 {
		t.Fatalf("expected at least 2 cloud providers, got %v", cloudNames)
	}
	if cloudNames[0] != "openrouter" {
		t.Errorf("OpenRouter should be first cloud provider, got %q", cloudNames[0])
	}
}

func TestOpenAIModels(t *testing.T) {
	if len(DefaultOpenAIModels) == 0 {
		t.Fatal("DefaultOpenAIModels should not be empty")
	}
	// Should contain gpt-4o-mini (cheapest, best for code review)
	found := false
	for _, m := range DefaultOpenAIModels {
		if m == "gpt-4o-mini" {
			found = true
		}
	}
	if !found {
		t.Error("DefaultOpenAIModels should include gpt-4o-mini")
	}
}

// --- Groq models ---

func TestGroqModels(t *testing.T) {
	if len(GroqModels) == 0 {
		t.Fatal("GroqModels should not be empty")
	}
}

func TestOllamaModels(t *testing.T) {
	if len(OllamaModels) == 0 {
		t.Fatal("OllamaModels should not be empty")
	}
}

func TestLMStudioModels(t *testing.T) {
	if len(LMStudioModels) == 0 {
		t.Fatal("LMStudioModels should not be empty")
	}
}

// --- helpers ---

type mockProvider struct {
	resp ChatResponse
	err  error
}

func (m *mockProvider) Chat(_ context.Context, _ ChatRequest) (ChatResponse, error) {
	return m.resp, m.err
}
func (m *mockProvider) Name() string { return "mock" }

func clearProviderEnvVars() {
	for _, v := range AutoDetectEnvVars {
		os.Unsetenv(v)
	}
	os.Unsetenv("GROQ_API_KEY")
	os.Unsetenv("OPENAI_API_KEY")
	os.Unsetenv("OLLAMA_HOST")
	os.Unsetenv("LM_STUDIO_URL")
}
