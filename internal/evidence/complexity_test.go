package evidence_test

import (
	"testing"

	"github.com/code-certification/certify/internal/evidence"
)

func TestComputeGoComplexity_SimpleFunc(t *testing.T) {
	src := `package main

func hello() {
	fmt.Println("hello")
}
`
	result := evidence.ComputeGoComplexity(src)
	if len(result) != 1 {
		t.Fatalf("expected 1 function, got %d", len(result))
	}
	if result["hello"] != 1 {
		t.Errorf("hello complexity = %d, want 1", result["hello"])
	}
}

func TestComputeGoComplexity_WithBranches(t *testing.T) {
	src := `package main

func complex(x int) string {
	if x > 0 {
		if x > 10 {
			return "big"
		}
		return "positive"
	} else if x < 0 {
		return "negative"
	}
	for i := 0; i < x; i++ {
		switch i {
		case 1:
			break
		case 2:
			continue
		default:
			return "default"
		}
	}
	return "zero"
}
`
	result := evidence.ComputeGoComplexity(src)
	// Base 1 + 2 if + 1 else if + 1 for + 3 case = 8
	if result["complex"] < 5 {
		t.Errorf("complex complexity = %d, want >= 5", result["complex"])
	}
}

func TestComputeGoComplexity_MultipleFuncs(t *testing.T) {
	src := `package main

func a() {}
func b(x int) {
	if x > 0 {
		return
	}
}
`
	result := evidence.ComputeGoComplexity(src)
	if len(result) != 2 {
		t.Fatalf("expected 2 functions, got %d", len(result))
	}
	if result["a"] != 1 {
		t.Errorf("a complexity = %d, want 1", result["a"])
	}
	if result["b"] != 2 {
		t.Errorf("b complexity = %d, want 2", result["b"])
	}
}

func TestComputeGoComplexity_Method(t *testing.T) {
	src := `package main

type Foo struct{}

func (f *Foo) Bar(x int) {
	if x > 0 {
		return
	}
}
`
	result := evidence.ComputeGoComplexity(src)
	if result["Foo.Bar"] != 2 {
		t.Errorf("Foo.Bar complexity = %d, want 2", result["Foo.Bar"])
	}
}

func TestComputeGoComplexity_EmptySource(t *testing.T) {
	result := evidence.ComputeGoComplexity("")
	if len(result) != 0 {
		t.Errorf("empty source should return 0 funcs, got %d", len(result))
	}
}

func TestComputeGoComplexity_LogicalOps(t *testing.T) {
	src := `package main

func check(a, b, c bool) bool {
	if a && b || c {
		return true
	}
	return false
}
`
	result := evidence.ComputeGoComplexity(src)
	// Base 1 + 1 if + 1 && + 1 || = 4
	if result["check"] < 3 {
		t.Errorf("check complexity = %d, want >= 3", result["check"])
	}
}

func TestComputeSymbolMetrics(t *testing.T) {
	src := `package main

// hello greets
func hello() {
	fmt.Println("hello")
}

// TODO: fix this
func broken(x int) {
	if x > 0 {
		return
	}
}
`
	metrics := evidence.ComputeSymbolMetrics(src, "hello")
	if metrics.CodeLines < 1 {
		t.Errorf("hello CodeLines = %d, want >= 1", metrics.CodeLines)
	}
	if metrics.TodoCount != 0 {
		t.Errorf("hello TodoCount = %d, want 0", metrics.TodoCount)
	}
	if metrics.Complexity < 1 {
		t.Errorf("hello Complexity = %d, want >= 1", metrics.Complexity)
	}

	brokenMetrics := evidence.ComputeSymbolMetrics(src, "broken")
	if brokenMetrics.TodoCount != 1 {
		t.Errorf("broken TodoCount = %d, want 1", brokenMetrics.TodoCount)
	}
	if brokenMetrics.Complexity < 2 {
		t.Errorf("broken Complexity = %d, want >= 2", brokenMetrics.Complexity)
	}
}

func TestComputeSymbolMetrics_NotFound(t *testing.T) {
	src := `package main

func hello() {}
`
	metrics := evidence.ComputeSymbolMetrics(src, "nonexistent")
	// Falls back to file-level
	if metrics.TotalLines == 0 {
		t.Error("nonexistent symbol should fall back to file-level metrics")
	}
}
