package agent

import "context"

// Provider is the interface for LLM API providers.
type Provider interface {
	// Chat sends a chat completion request and returns the response.
	Chat(ctx context.Context, req ChatRequest) (ChatResponse, error)

	// Name returns the provider name.
	Name() string
}
