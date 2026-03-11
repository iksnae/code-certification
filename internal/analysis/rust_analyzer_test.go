package analysis

import (
	"testing"
)

const rustTestSrc = `use std::process::Command;
use std::collections::HashMap;

/// Greet says hello.
pub fn greet(name: &str) -> String {
    format!("Hello, {}", name)
}

fn internal() {}

pub fn process(items: Vec<String>, ctx: &str) -> Result<(), String> {
    for item in &items {
        if item.is_empty() {
            continue;
        }
        println!("{}", item);
    }
    Ok(())
}

pub fn big_func(x: i32) -> &'static str {
    if x > 0 {
        for i in 0..x {
            if i % 2 == 0 {
                match i {
                    0..=9 => return "small",
                    _ => {}
                }
            }
        }
    }
    "done"
}

pub struct MyService {
    name: String,
}

impl MyService {
    pub fn new(name: String) -> Self {
        MyService { name }
    }

    pub fn run(&self) {}
    pub fn stop(&self) {}
}

pub trait Logger {
    fn log(&self, msg: &str);
}

pub fn crasher() {
    panic!("boom");
}

pub fn unwrapper(val: Option<i32>) -> i32 {
    val.unwrap()
}
`

func TestRustAnalyzer_Discover(t *testing.T) {
	a := ForLanguage("rs")
	if a == nil {
		t.Fatal("expected Rust analyzer")
	}

	symbols, err := a.Discover("test.rs", []byte(rustTestSrc))
	if err != nil {
		t.Fatal(err)
	}

	names := map[string]bool{}
	for _, s := range symbols {
		names[s.Name] = true
	}

	expected := []string{"greet", "internal", "process", "big_func", "MyService", "new", "run", "stop", "Logger", "crasher", "unwrapper"}
	for _, name := range expected {
		if !names[name] {
			t.Errorf("missing symbol %q (found: %v)", name, names)
		}
	}

	for _, s := range symbols {
		switch s.Name {
		case "greet":
			if s.Kind != SymbolFunction {
				t.Errorf("greet kind = %v, want function", s.Kind)
			}
			if !s.Exported {
				t.Error("greet should be exported (pub)")
			}
		case "internal":
			if s.Exported {
				t.Error("internal should not be exported")
			}
		case "MyService":
			if s.Kind != SymbolClass {
				t.Errorf("MyService kind = %v, want class", s.Kind)
			}
		case "Logger":
			if s.Kind != SymbolInterface {
				t.Errorf("Logger kind = %v, want interface", s.Kind)
			}
		}
	}
}

func TestRustAnalyzer_Analyze_Greet(t *testing.T) {
	a := ForLanguage("rs")
	if a == nil {
		t.Fatal("expected Rust analyzer")
	}

	m, err := a.Analyze("test.rs", []byte(rustTestSrc), "greet")
	if err != nil {
		t.Fatal(err)
	}

	if m.ParamCount != 1 {
		t.Errorf("param_count = %d, want 1", m.ParamCount)
	}
	if m.ReturnCount != 1 {
		t.Errorf("return_count = %d, want 1", m.ReturnCount)
	}
	if !m.HasDocComment {
		t.Error("expected doc comment")
	}
	if !m.IsExported {
		t.Error("expected exported (pub)")
	}
}

func TestRustAnalyzer_Analyze_BigFunc(t *testing.T) {
	a := ForLanguage("rs")
	if a == nil {
		t.Fatal("expected Rust analyzer")
	}

	m, err := a.Analyze("test.rs", []byte(rustTestSrc), "big_func")
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

func TestRustAnalyzer_Analyze_UnsafeImports(t *testing.T) {
	a := ForLanguage("rs")
	if a == nil {
		t.Fatal("expected Rust analyzer")
	}

	m, err := a.Analyze("test.rs", []byte(rustTestSrc), "greet")
	if err != nil {
		t.Fatal(err)
	}

	if len(m.UnsafeImports) < 1 {
		t.Errorf("unsafe_imports = %v, want >= 1 (std::process)", m.UnsafeImports)
	}
}

func TestRustAnalyzer_Analyze_Panic(t *testing.T) {
	a := ForLanguage("rs")
	if a == nil {
		t.Fatal("expected Rust analyzer")
	}

	m, err := a.Analyze("test.rs", []byte(rustTestSrc), "crasher")
	if err != nil {
		t.Fatal(err)
	}

	if m.PanicCalls < 1 {
		t.Errorf("panic_calls = %d, want >= 1", m.PanicCalls)
	}
}

func TestRustAnalyzer_Language(t *testing.T) {
	a := ForLanguage("rs")
	if a == nil {
		t.Fatal("expected Rust analyzer")
	}
	if a.Language() != "rs" {
		t.Errorf("language = %q, want rs", a.Language())
	}
}
