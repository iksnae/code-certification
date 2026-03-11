package analysis

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func testdataDir(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("cannot determine test file location")
	}
	return filepath.Join(filepath.Dir(file), "testdata", "deepproject")
}

func loadTestProject(t *testing.T) *DeepGoAnalyzer {
	t.Helper()
	root := testdataDir(t)
	if _, err := os.Stat(filepath.Join(root, "go.mod")); err != nil {
		t.Skipf("testdata/deepproject not found: %v", err)
	}
	a, err := LoadGoProject(root, "./...")
	if err != nil {
		t.Fatalf("LoadGoProject: %v", err)
	}
	return a
}

func TestLoadGoProject(t *testing.T) {
	a := loadTestProject(t)
	if a == nil {
		t.Fatal("expected non-nil analyzer")
	}
	if len(a.pkgs) == 0 {
		t.Fatal("expected at least one package")
	}
}

func TestFanIn(t *testing.T) {
	a := loadTestProject(t)

	// Format is called by Hello, Goodbye, and internalHelper (3 call sites within the project)
	fanIn := a.FanIn("example.com/deepproject/pkg/greet", "Format")
	if fanIn < 3 {
		t.Errorf("Format fan-in: got %d, want >= 3", fanIn)
	}

	// Hello is called from main and run (2 call sites in cmd/app)
	fanInHello := a.FanIn("example.com/deepproject/pkg/greet", "Hello")
	if fanInHello < 2 {
		t.Errorf("Hello fan-in: got %d, want >= 2", fanInHello)
	}

	// UnusedExport should have fan-in = 0 (no external or internal callers)
	fanInUnused := a.FanIn("example.com/deepproject/pkg/greet", "UnusedExport")
	if fanInUnused != 0 {
		t.Errorf("UnusedExport fan-in: got %d, want 0", fanInUnused)
	}
}

func TestFanOut(t *testing.T) {
	a := loadTestProject(t)

	// Hello calls Format and fmt.Sprintf → fan-out >= 2
	fanOut := a.FanOut("example.com/deepproject/pkg/greet", "Hello")
	if fanOut < 2 {
		t.Errorf("Hello fan-out: got %d, want >= 2", fanOut)
	}

	// Format calls nothing (just returns) → fan-out = 0
	fanOutFormat := a.FanOut("example.com/deepproject/pkg/greet", "Format")
	if fanOutFormat != 0 {
		t.Errorf("Format fan-out: got %d, want 0", fanOutFormat)
	}
}

func TestUnusedExports(t *testing.T) {
	a := loadTestProject(t)

	unused := a.UnusedExports()

	// UnusedExport should be detected
	found := false
	for _, u := range unused {
		if u.Name == "UnusedExport" && u.Pkg == "example.com/deepproject/pkg/greet" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected UnusedExport to be in unused exports, got: %v", unused)
	}

	// Format, Hello, Goodbye should NOT be in the unused list
	for _, u := range unused {
		if u.Pkg == "example.com/deepproject/pkg/greet" {
			if u.Name == "Format" || u.Name == "Hello" || u.Name == "Goodbye" {
				t.Errorf("expected %s to NOT be in unused exports", u.Name)
			}
		}
	}
}

func TestFanInNonexistentFunction(t *testing.T) {
	a := loadTestProject(t)
	fanIn := a.FanIn("example.com/deepproject/pkg/greet", "DoesNotExist")
	if fanIn != 0 {
		t.Errorf("nonexistent function fan-in: got %d, want 0", fanIn)
	}
}

func TestFanOutNonexistentFunction(t *testing.T) {
	a := loadTestProject(t)
	fanOut := a.FanOut("example.com/deepproject/pkg/greet", "DoesNotExist")
	if fanOut != 0 {
		t.Errorf("nonexistent function fan-out: got %d, want 0", fanOut)
	}
}

func TestDeepGoAnalyzerCaching(t *testing.T) {
	a := loadTestProject(t)

	// Call twice — should return same results
	fanIn1 := a.FanIn("example.com/deepproject/pkg/greet", "Format")
	fanIn2 := a.FanIn("example.com/deepproject/pkg/greet", "Format")
	if fanIn1 != fanIn2 {
		t.Errorf("fan-in not consistent: got %d and %d", fanIn1, fanIn2)
	}
}

func TestLookupResults(t *testing.T) {
	a := loadTestProject(t)

	results := a.AllResults()
	if len(results) == 0 {
		t.Fatal("expected non-empty results map")
	}

	// Check that Hello has results
	key := FuncKey{Pkg: "example.com/deepproject/pkg/greet", Name: "Hello"}
	r, ok := results[key]
	if !ok {
		t.Fatalf("expected results for Hello, keys: %v", mapKeys(results))
	}
	if r.FanIn < 2 {
		t.Errorf("Hello fan-in from results: got %d, want >= 2", r.FanIn)
	}
}

func mapKeys(m map[FuncKey]DeepResult) []FuncKey {
	keys := make([]FuncKey, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
