package analysis

import (
	"testing"
)

func TestUnusedParams(t *testing.T) {
	a := loadTestProject(t)

	// internalHelper has no params — should report 0 unused
	result := a.UnusedParams("example.com/deepproject/pkg/greet", "internalHelper")
	if result != 0 {
		t.Errorf("internalHelper unused params: got %d, want 0", result)
	}
}

func TestInterfaceSize(t *testing.T) {
	a := loadTestProject(t)

	// Functions don't implement interfaces — should return 0
	size := a.InterfaceSize("example.com/deepproject/pkg/greet", "Hello")
	if size != 0 {
		t.Errorf("Hello interface size: got %d, want 0 (not a method)", size)
	}
}

// TestTypeAwareErrorWrapping tests that error wrapping detection
// considers actual types rather than just string patterns.
func TestTypeAwareErrorWrapping(t *testing.T) {
	// This is a refinement of existing error wrapping detection.
	// The deep analyzer can verify that wrapped values are actually error types.
	a := loadTestProject(t)
	if a == nil {
		t.Skip("deep analyzer not available")
	}
	// Basic smoke test — our test project doesn't have error returns
	count := a.TypeAwareErrorWrapping("example.com/deepproject/pkg/greet", "Hello")
	if count != 0 {
		t.Errorf("Hello error wrapping issues: got %d, want 0 (no error return)", count)
	}
}
