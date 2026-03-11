package analysis

import (
	"testing"
)

func TestDepDepth(t *testing.T) {
	a := loadTestProject(t)

	// cmd/app imports pkg/greet → depth = 1 from greet
	// greet imports only fmt (stdlib) → depth = 0 for local deps
	depth := a.DepDepth("example.com/deepproject/pkg/greet")
	if depth != 0 {
		t.Errorf("greet dep depth: got %d, want 0 (no local deps)", depth)
	}

	depthApp := a.DepDepth("example.com/deepproject/cmd/app")
	if depthApp < 1 {
		t.Errorf("cmd/app dep depth: got %d, want >= 1 (imports greet)", depthApp)
	}
}

func TestInstability(t *testing.T) {
	a := loadTestProject(t)

	// greet: imported by cmd/app (Ca=1), imports nothing local (Ce=0)
	// Instability = Ce / (Ca + Ce) = 0 / (1 + 0) = 0.0
	inst := a.Instability("example.com/deepproject/pkg/greet")
	if inst > 0.1 {
		t.Errorf("greet instability: got %f, want ~0.0 (stable)", inst)
	}

	// cmd/app: imported by nobody (Ca=0), imports greet (Ce=1)
	// Instability = 1 / (0 + 1) = 1.0
	instApp := a.Instability("example.com/deepproject/cmd/app")
	if instApp < 0.9 {
		t.Errorf("cmd/app instability: got %f, want ~1.0 (unstable)", instApp)
	}
}

func TestParamAbstraction(t *testing.T) {
	a := loadTestProject(t)

	// Hello takes string — primitive, not concrete struct. Should score well.
	result := a.ParamAbstraction("example.com/deepproject/pkg/greet", "Hello")
	if result.ConcreteDeps > 0 {
		t.Errorf("Hello concrete deps: got %d, want 0 (string is primitive)", result.ConcreteDeps)
	}
}

func TestDepDepthNonexistent(t *testing.T) {
	a := loadTestProject(t)
	depth := a.DepDepth("example.com/nonexistent")
	if depth != 0 {
		t.Errorf("nonexistent pkg dep depth: got %d, want 0", depth)
	}
}

func TestInstabilityNonexistent(t *testing.T) {
	a := loadTestProject(t)
	inst := a.Instability("example.com/nonexistent")
	if inst != 0 {
		t.Errorf("nonexistent pkg instability: got %f, want 0", inst)
	}
}
