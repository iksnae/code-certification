package analysis

import (
	"testing"
)

const pyTestSrc = `import os
import subprocess
from typing import List, Optional

def greet(name: str) -> str:
    """Say hello to the user."""
    return "Hello, " + name

def _internal():
    pass

def process(items: List[str], ctx) -> None:
    for item in items:
        if item == "":
            continue
        print(item)

def big_func(x: int) -> str:
    if x > 0:
        for i in range(x):
            if i % 2 == 0:
                if i < 10:
                    return "small"
    return "done"

class MyService:
    """A service class."""
    
    def __init__(self, name: str):
        self.name = name

    def run(self) -> None:
        pass

    def stop(self) -> None:
        pass

def crasher():
    raise Exception("boom")

def bad_errors():
    try:
        return 1 / 0
    except:
        pass
`

func TestPyAnalyzer_Discover(t *testing.T) {
	a := ForLanguage("py")
	if a == nil {
		t.Fatal("expected Python analyzer")
	}

	symbols, err := a.Discover("test.py", []byte(pyTestSrc))
	if err != nil {
		t.Fatal(err)
	}

	names := map[string]bool{}
	for _, s := range symbols {
		names[s.Name] = true
	}

	expected := []string{"greet", "_internal", "process", "big_func", "MyService", "__init__", "run", "stop", "crasher", "bad_errors"}
	for _, name := range expected {
		if !names[name] {
			t.Errorf("missing symbol %q", name)
		}
	}

	for _, s := range symbols {
		switch s.Name {
		case "greet":
			if s.Kind != SymbolFunction {
				t.Errorf("greet kind = %v, want function", s.Kind)
			}
			if !s.Exported {
				t.Error("greet should be exported (no underscore)")
			}
		case "_internal":
			if s.Exported {
				t.Error("_internal should not be exported")
			}
		case "run":
			if s.Kind != SymbolMethod {
				t.Errorf("run kind = %v, want method", s.Kind)
			}
			if s.Parent != "MyService" {
				t.Errorf("run parent = %q, want MyService", s.Parent)
			}
		case "MyService":
			if s.Kind != SymbolClass {
				t.Errorf("MyService kind = %v, want class", s.Kind)
			}
		}
	}
}

func TestPyAnalyzer_Analyze_Greet(t *testing.T) {
	a := ForLanguage("py")
	if a == nil {
		t.Fatal("expected Python analyzer")
	}

	m, err := a.Analyze("test.py", []byte(pyTestSrc), "greet")
	if err != nil {
		t.Fatal(err)
	}

	if m.ParamCount != 1 {
		t.Errorf("param_count = %d, want 1", m.ParamCount)
	}
	if !m.HasDocComment {
		t.Error("expected docstring")
	}
	if !m.IsExported {
		t.Error("expected exported")
	}
}

func TestPyAnalyzer_Analyze_BigFunc(t *testing.T) {
	a := ForLanguage("py")
	if a == nil {
		t.Fatal("expected Python analyzer")
	}

	m, err := a.Analyze("test.py", []byte(pyTestSrc), "big_func")
	if err != nil {
		t.Fatal(err)
	}

	if m.MaxNestingDepth < 3 {
		t.Errorf("nesting = %d, want >= 3", m.MaxNestingDepth)
	}
	if m.CyclomaticComplexity < 3 {
		t.Errorf("cyclomatic = %d, want >= 3", m.CyclomaticComplexity)
	}
}

func TestPyAnalyzer_Analyze_Class(t *testing.T) {
	a := ForLanguage("py")
	if a == nil {
		t.Fatal("expected Python analyzer")
	}

	m, err := a.Analyze("test.py", []byte(pyTestSrc), "MyService")
	if err != nil {
		t.Fatal(err)
	}

	if m.MethodCount < 3 {
		t.Errorf("method_count = %d, want >= 3", m.MethodCount)
	}
	if !m.HasDocComment {
		t.Error("expected class docstring")
	}
}

func TestPyAnalyzer_Analyze_EmptyCatch(t *testing.T) {
	a := ForLanguage("py")
	if a == nil {
		t.Fatal("expected Python analyzer")
	}

	m, err := a.Analyze("test.py", []byte(pyTestSrc), "bad_errors")
	if err != nil {
		t.Fatal(err)
	}

	if m.EmptyCatchBlocks < 1 {
		t.Errorf("empty_catch = %d, want >= 1", m.EmptyCatchBlocks)
	}
}

func TestPyAnalyzer_Analyze_UnsafeImports(t *testing.T) {
	a := ForLanguage("py")
	if a == nil {
		t.Fatal("expected Python analyzer")
	}

	m, err := a.Analyze("test.py", []byte(pyTestSrc), "greet")
	if err != nil {
		t.Fatal(err)
	}

	if len(m.UnsafeImports) < 1 {
		t.Errorf("unsafe_imports = %v, want >= 1 (subprocess)", m.UnsafeImports)
	}
}

func TestPyAnalyzer_Language(t *testing.T) {
	a := ForLanguage("py")
	if a == nil {
		t.Fatal("expected Python analyzer")
	}
	if a.Language() != "py" {
		t.Errorf("language = %q, want py", a.Language())
	}
}
