package agent

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// OpenRouterProvider implements Provider for the OpenRouter API.
type OpenRouterProvider struct {
	baseURL     string
	apiKey      string
	httpReferer string
	xTitle      string
	client      *http.Client
	maxRetries  int
	backoffBase time.Duration
}

// NewOpenRouterProvider creates a new OpenRouter provider.
func NewOpenRouterProvider(baseURL, apiKey, httpReferer, xTitle string) *OpenRouterProvider {
	return &OpenRouterProvider{
		baseURL:     baseURL,
		apiKey:      apiKey,
		httpReferer: httpReferer,
		xTitle:      xTitle,
		client:      &http.Client{Timeout: 60 * time.Second},
		maxRetries:  3,
		backoffBase: 100 * time.Millisecond, // Short for tests, configurable for prod
	}
}

// Name returns the provider name.
func (p *OpenRouterProvider) Name() string {
	return "openrouter"
}

// Chat sends a chat completion request with retry logic.
func (p *OpenRouterProvider) Chat(ctx context.Context, req ChatRequest) (ChatResponse, error) {
	if p.apiKey == "" {
		return ChatResponse{}, fmt.Errorf("openrouter: API key not configured")
	}

	var lastErr error
	for attempt := 0; attempt <= p.maxRetries; attempt++ {
		if attempt > 0 {
			backoff := p.backoffBase * time.Duration(1<<(attempt-1))
			select {
			case <-ctx.Done():
				return ChatResponse{}, ctx.Err()
			case <-time.After(backoff):
			}
		}

		resp, err := p.doRequest(ctx, req)
		if err == nil {
			return resp, nil
		}

		lastErr = err

		// Don't retry on auth errors or context cancellation
		if isAuthError(err) || ctx.Err() != nil {
			return ChatResponse{}, err
		}
		// Only retry on rate limit or server errors
		if !isRetryable(err) {
			return ChatResponse{}, err
		}
	}

	return ChatResponse{}, fmt.Errorf("openrouter: all retries exhausted: %w", lastErr)
}

func (p *OpenRouterProvider) doRequest(ctx context.Context, req ChatRequest) (ChatResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return ChatResponse{}, fmt.Errorf("openrouter: marshaling request: %w", err)
	}

	url := p.baseURL + "/chat/completions"
	httpReq, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewReader(body))
	if err != nil {
		return ChatResponse{}, fmt.Errorf("openrouter: creating request: %w", err)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+p.apiKey)
	if p.httpReferer != "" {
		httpReq.Header.Set("HTTP-Referer", p.httpReferer)
	}
	if p.xTitle != "" {
		httpReq.Header.Set("X-Title", p.xTitle)
	}

	httpResp, err := p.client.Do(httpReq)
	if err != nil {
		return ChatResponse{}, fmt.Errorf("openrouter: HTTP request: %w", err)
	}
	defer httpResp.Body.Close()

	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return ChatResponse{}, fmt.Errorf("openrouter: reading response: %w", err)
	}

	if httpResp.StatusCode != 200 {
		return ChatResponse{}, &APIError{
			StatusCode: httpResp.StatusCode,
			Body:       string(respBody),
		}
	}

	var resp ChatResponse
	if err := json.Unmarshal(respBody, &resp); err != nil {
		return ChatResponse{}, fmt.Errorf("openrouter: parsing response: %w", err)
	}

	return resp, nil
}

// APIError represents an HTTP API error.
type APIError struct {
	StatusCode int
	Body       string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("openrouter API error %d: %s", e.StatusCode, e.Body)
}

func isAuthError(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.StatusCode == 401 || apiErr.StatusCode == 403
	}
	return false
}

func isRetryable(err error) bool {
	if apiErr, ok := err.(*APIError); ok {
		return apiErr.StatusCode == 429 || apiErr.StatusCode >= 500
	}
	return false
}
