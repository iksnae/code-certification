package service

import "context"

// Syncer coordinates synchronization.
type Syncer struct {
	Name string
}

// Apply runs the synchronization.
func (s *Syncer) Apply(ctx context.Context) error {
	return nil
}

// Reset clears state.
func (s *Syncer) Reset() {
	s.Name = ""
}

// Format returns a formatted string.
func Format(input string) string {
	return "[" + input + "]"
}
