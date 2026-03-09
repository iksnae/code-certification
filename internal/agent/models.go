package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// ModelInfo describes a model available from a provider.
type ModelInfo struct {
	ID            string `json:"id"`
	OwnedBy       string `json:"owned_by,omitempty"`
	ContextWindow int    `json:"context_window,omitempty"`
	Created       int64  `json:"created,omitempty"`
}

// ListModels queries an OpenAI-compatible /models endpoint.
// Falls back to Ollama's /api/tags if the standard endpoint returns 404.
// apiKey can be empty for local providers.
func ListModels(ctx context.Context, baseURL, apiKey string) ([]ModelInfo, error) {
	client := &http.Client{Timeout: 10 * time.Second}

	// Try OpenAI-compatible /models first
	models, err := listOpenAIModels(ctx, client, baseURL, apiKey)
	if err == nil {
		return models, nil
	}

	// If 404 or similar, try Ollama /api/tags
	ollamaBase := strings.TrimSuffix(baseURL, "/v1")
	models, ollamaErr := listOllamaModels(ctx, client, ollamaBase)
	if ollamaErr == nil {
		return models, nil
	}

	// Return the original error
	return nil, fmt.Errorf("listing models: %w", err)
}

// openAIModelsResponse is the OpenAI /v1/models response shape.
type openAIModelsResponse struct {
	Data []openAIModel `json:"data"`
}

type openAIModel struct {
	ID            string `json:"id"`
	OwnedBy       string `json:"owned_by"`
	ContextWindow int    `json:"context_window,omitempty"`
	Created       int64  `json:"created"`
}

func listOpenAIModels(ctx context.Context, client *http.Client, baseURL, apiKey string) ([]ModelInfo, error) {
	url := strings.TrimRight(baseURL, "/") + "/models"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	if apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+apiKey)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(body))
	}

	var result openAIModelsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("parsing response: %w", err)
	}

	models := make([]ModelInfo, 0, len(result.Data))
	for _, m := range result.Data {
		models = append(models, ModelInfo{
			ID:            m.ID,
			OwnedBy:       m.OwnedBy,
			ContextWindow: m.ContextWindow,
			Created:       m.Created,
		})
	}
	return models, nil
}

// ollamaTagsResponse is Ollama's /api/tags response shape.
type ollamaTagsResponse struct {
	Models []ollamaModel `json:"models"`
}

type ollamaModel struct {
	Name       string `json:"name"`
	Size       int64  `json:"size"`
	ModifiedAt string `json:"modified_at"`
}

func listOllamaModels(ctx context.Context, client *http.Client, baseURL string) ([]ModelInfo, error) {
	url := strings.TrimRight(baseURL, "/") + "/api/tags"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(body))
	}

	var result ollamaTagsResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("parsing ollama response: %w", err)
	}

	models := make([]ModelInfo, 0, len(result.Models))
	for _, m := range result.Models {
		models = append(models, ModelInfo{
			ID:      m.Name,
			OwnedBy: "ollama",
		})
	}
	return models, nil
}
