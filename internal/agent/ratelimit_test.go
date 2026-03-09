package agent_test

import (
	"testing"
	"time"

	"github.com/code-certification/certify/internal/agent"
)

func TestRateLimiter_AllowsWithinLimit(t *testing.T) {
	rl := agent.NewRateLimiter(10, time.Second)

	for i := 0; i < 10; i++ {
		if !rl.Allow() {
			t.Errorf("request %d should be allowed", i)
		}
	}
}

func TestRateLimiter_BlocksOverLimit(t *testing.T) {
	rl := agent.NewRateLimiter(5, time.Second)

	for i := 0; i < 5; i++ {
		rl.Allow()
	}

	if rl.Allow() {
		t.Error("6th request should be blocked")
	}
}

func TestRateLimiter_RefillsOverTime(t *testing.T) {
	rl := agent.NewRateLimiter(5, 50*time.Millisecond)

	// Exhaust all tokens
	for i := 0; i < 5; i++ {
		rl.Allow()
	}

	if rl.Allow() {
		t.Error("should be blocked immediately after exhaustion")
	}

	// Wait for refill
	time.Sleep(60 * time.Millisecond)

	if !rl.Allow() {
		t.Error("should be allowed after refill period")
	}
}

func TestRateLimiter_Wait(t *testing.T) {
	rl := agent.NewRateLimiter(1, 50*time.Millisecond)

	rl.Allow() // Take the one token

	start := time.Now()
	rl.Wait() // Should block until refill
	elapsed := time.Since(start)

	if elapsed < 40*time.Millisecond {
		t.Errorf("Wait() returned too fast: %v", elapsed)
	}
}
