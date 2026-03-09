package agent

import (
	"net/http"
	"os"
	"strings"
	"time"
)

// DetectedProvider describes an auto-detected LLM provider.
type DetectedProvider struct {
	Name    string   // "openrouter", "groq", "ollama", "lmstudio"
	BaseURL string   // API base URL
	APIKey  string   // API key (empty for local providers)
	Models  []string // Preferred models for this provider
	Local   bool     // True for local providers (no auth required)
}

// GroqModels lists preferred models for Groq, in fallback order.
var GroqModels = []string{
	"llama-3.3-70b-versatile",
	"llama-3.1-8b-instant",
	"gemma2-9b-it",
}

// OllamaModels lists common Ollama models, in preference order.
// Users should have at least one of these installed.
var OllamaModels = []string{
	"qwen2.5-coder:7b",
	"qwen2.5-coder:3b",
	"llama3.2:3b",
	"phi4",
	"gemma2:9b",
}

// LMStudioModels lists models for LM Studio.
// LM Studio uses whatever model is loaded — these are hints.
var LMStudioModels = []string{
	"loaded-model", // LM Studio auto-routes to the loaded model
}

// DetectProviders checks for available LLM providers in priority order.
// Cloud providers (requiring API keys) come first, local providers last.
func DetectProviders() []DetectedProvider {
	var cloud, local []DetectedProvider

	// Cloud providers (priority order)
	if key := os.Getenv("OPENROUTER_API_KEY"); key != "" {
		cloud = append(cloud, DetectedProvider{
			Name:    "openrouter",
			BaseURL: "https://openrouter.ai/api/v1",
			APIKey:  key,
			Models:  ConservativeModels,
		})
	} else if key := os.Getenv("CERTIFY_API_KEY"); key != "" {
		cloud = append(cloud, DetectedProvider{
			Name:    "openrouter",
			BaseURL: "https://openrouter.ai/api/v1",
			APIKey:  key,
			Models:  ConservativeModels,
		})
	}

	if key := os.Getenv("GROQ_API_KEY"); key != "" {
		cloud = append(cloud, DetectedProvider{
			Name:    "groq",
			BaseURL: "https://api.groq.com/openai/v1",
			APIKey:  key,
			Models:  GroqModels,
		})
	}

	// Local providers
	if host := os.Getenv("OLLAMA_HOST"); host != "" {
		baseURL := normalizeLocalURL(host)
		local = append(local, DetectedProvider{
			Name:    "ollama",
			BaseURL: baseURL,
			Models:  OllamaModels,
			Local:   true,
		})
	} else if probeLocal("http://localhost:11434") {
		local = append(local, DetectedProvider{
			Name:    "ollama",
			BaseURL: "http://localhost:11434/v1",
			Models:  OllamaModels,
			Local:   true,
		})
	}

	if url := os.Getenv("LM_STUDIO_URL"); url != "" {
		baseURL := normalizeLocalURL(url)
		local = append(local, DetectedProvider{
			Name:    "lmstudio",
			BaseURL: baseURL,
			Models:  LMStudioModels,
			Local:   true,
		})
	} else if probeLocal("http://localhost:1234") {
		local = append(local, DetectedProvider{
			Name:    "lmstudio",
			BaseURL: "http://localhost:1234/v1",
			Models:  LMStudioModels,
			Local:   true,
		})
	}

	// Cloud first, then local
	return append(cloud, local...)
}

// normalizeLocalURL ensures the URL ends with /v1 for OpenAI compatibility.
func normalizeLocalURL(url string) string {
	url = strings.TrimRight(url, "/")
	if !strings.HasSuffix(url, "/v1") {
		url += "/v1"
	}
	return url
}

// probeLocal checks if a local server is listening (quick health check).
func probeLocal(baseURL string) bool {
	client := &http.Client{Timeout: 1 * time.Second}
	resp, err := client.Get(baseURL)
	if err != nil {
		return false
	}
	resp.Body.Close()
	return resp.StatusCode < 500
}

// ProviderNames returns the display names of detected providers.
func ProviderNames(providers []DetectedProvider) []string {
	names := make([]string, len(providers))
	for i, p := range providers {
		names[i] = p.Name
	}
	return names
}
