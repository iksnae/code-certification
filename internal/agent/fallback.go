package agent

import (
	"context"
	"fmt"
	"time"
)

// FallbackProvider tries multiple providers in order until one succeeds.
// Only retryable errors (429, 5xx) trigger fallback; auth errors abort immediately.
type FallbackProvider struct {
	providers []Provider
}

// NewFallbackProvider creates a provider that falls back through a list.
func NewFallbackProvider(providers []Provider) *FallbackProvider {
	return &FallbackProvider{providers: providers}
}

func (f *FallbackProvider) Chat(ctx context.Context, req ChatRequest) (ChatResponse, error) {
	var lastErr error
	for _, p := range f.providers {
		resp, err := p.Chat(ctx, req)
		if err == nil {
			return resp, nil
		}
		lastErr = err
		// Auth/budget errors = account-level problem, don't try others
		if isAuthError(err) || isBudgetError(err) {
			return ChatResponse{}, err
		}
		// Fall through on retryable errors (429, 5xx) and connection errors.
		// Connection errors (dial failures, timeouts) are non-APIError — always fall through.
		// Only non-retryable API errors (400, 404) stop the chain.
		if isAPIError(err) && !isRetryable(err) {
			return ChatResponse{}, err
		}
		// Non-API errors (connection refused, DNS, timeout) → try next provider
	}
	return ChatResponse{}, fmt.Errorf("all providers failed: %w", lastErr)
}

func (f *FallbackProvider) Name() string { return "fallback" }

// ModelChain creates a FallbackProvider with one OpenRouterProvider per model.
// Each provider is configured with longer backoff suitable for free-tier rate limits.
type ModelChain struct {
	inner *FallbackProvider
}

// NewModelChain creates a chain of model-specific providers.
// When model A returns 429, automatically tries model B, etc.
func NewModelChain(baseURL, apiKey, referer, title string, models []string) *ModelChain {
	var providers []Provider
	for _, model := range models {
		p := NewOpenRouterProvider(baseURL, apiKey, referer, title)
		// Longer backoff for free tier
		p.backoffBase = 2 * time.Second
		p.maxRetries = 2
		providers = append(providers, &modelPinnedProvider{
			provider: p,
			model:    model,
		})
	}
	return &ModelChain{inner: NewFallbackProvider(providers)}
}

func (mc *ModelChain) Chat(ctx context.Context, req ChatRequest) (ChatResponse, error) {
	return mc.inner.Chat(ctx, req)
}

func (mc *ModelChain) Name() string { return "model-chain" }

// modelPinnedProvider overrides the model in every request.
type modelPinnedProvider struct {
	provider *OpenRouterProvider
	model    string
}

func (mp *modelPinnedProvider) Chat(ctx context.Context, req ChatRequest) (ChatResponse, error) {
	req.Model = mp.model
	return mp.provider.Chat(ctx, req)
}

func (mp *modelPinnedProvider) Name() string { return mp.model }

// AdaptiveMessages builds a message list that works with or without system message support.
// When useSystem=true, sends separate system + user messages.
// When useSystem=false, combines system instruction into the user message.
func AdaptiveMessages(systemPrompt, userContent string, useSystem bool) []Message {
	if useSystem {
		return []Message{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: userContent},
		}
	}
	combined := systemPrompt + "\n\n" + userContent
	return []Message{
		{Role: "user", Content: combined},
	}
}
