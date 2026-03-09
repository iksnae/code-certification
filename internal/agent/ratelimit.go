package agent

import (
	"sync"
	"time"
)

// RateLimiter implements a simple token bucket rate limiter.
type RateLimiter struct {
	mu       sync.Mutex
	tokens   int
	maxRate  int
	interval time.Duration
	lastTime time.Time
}

// NewRateLimiter creates a rate limiter that allows maxRate requests per interval.
func NewRateLimiter(maxRate int, interval time.Duration) *RateLimiter {
	return &RateLimiter{
		tokens:   maxRate,
		maxRate:  maxRate,
		interval: interval,
		lastTime: time.Now(),
	}
}

// Allow checks if a request is allowed (non-blocking).
func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	rl.refill()

	if rl.tokens > 0 {
		rl.tokens--
		return true
	}
	return false
}

// Wait blocks until a token is available.
func (rl *RateLimiter) Wait() {
	for {
		if rl.Allow() {
			return
		}
		// Sleep for a fraction of the interval
		time.Sleep(rl.interval / time.Duration(rl.maxRate))
	}
}

func (rl *RateLimiter) refill() {
	now := time.Now()
	elapsed := now.Sub(rl.lastTime)
	if elapsed >= rl.interval {
		periods := int(elapsed / rl.interval)
		rl.tokens = min(rl.maxRate, rl.tokens+periods*rl.maxRate)
		rl.lastTime = rl.lastTime.Add(time.Duration(periods) * rl.interval)
	}
}
