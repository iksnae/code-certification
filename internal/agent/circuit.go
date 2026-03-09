package agent

import (
	"context"
	"fmt"
	"sync"
)

// CircuitBreaker wraps a Provider and stops calling it after consecutive failures.
type CircuitBreaker struct {
	provider  Provider
	threshold int
	failures  int
	open      bool
	mu        sync.Mutex
}

// NewCircuitBreaker wraps a provider with a circuit breaker.
// After `threshold` consecutive failures, the circuit opens and all calls
// return immediately with an error until a successful call closes it.
func NewCircuitBreaker(provider Provider, threshold int) *CircuitBreaker {
	return &CircuitBreaker{
		provider:  provider,
		threshold: threshold,
	}
}

// Chat delegates to the wrapped provider unless the circuit is open.
func (cb *CircuitBreaker) Chat(ctx context.Context, req ChatRequest) (ChatResponse, error) {
	cb.mu.Lock()
	if cb.open {
		cb.mu.Unlock()
		return ChatResponse{}, fmt.Errorf("circuit breaker open: %d consecutive failures", cb.threshold)
	}
	cb.mu.Unlock()

	resp, err := cb.provider.Chat(ctx, req)

	cb.mu.Lock()
	defer cb.mu.Unlock()
	if err != nil {
		cb.failures++
		if cb.failures >= cb.threshold {
			cb.open = true
		}
		return resp, err
	}

	// Success — reset
	cb.failures = 0
	cb.open = false
	return resp, nil
}

// Name returns the wrapped provider's name.
func (cb *CircuitBreaker) Name() string {
	return cb.provider.Name()
}

// IsOpen returns true if the circuit is currently open.
func (cb *CircuitBreaker) IsOpen() bool {
	cb.mu.Lock()
	defer cb.mu.Unlock()
	return cb.open
}
