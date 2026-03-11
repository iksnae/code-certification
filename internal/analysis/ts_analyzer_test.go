package analysis

import (
	"testing"
)

const tsTestSrc = `import { readFile } from "fs";
import { exec } from "child_process";

/**
 * Greet says hello.
 */
export function greet(name: string): string {
  return "Hello, " + name;
}

function internal() {}

export async function process(items: string[], ctx: any): Promise<void> {
  for (const item of items) {
    if (item === "") {
      continue;
    }
    console.log(item);
  }
}

export function bigFunc(x: number): string {
  if (x > 0) {
    for (let i = 0; i < x; i++) {
      if (i % 2 === 0) {
        switch (true) {
          case i < 10:
            return "small";
        }
      }
    }
  }
  return "done";
}

export class MyService {
  private name: string;

  constructor(name: string) {
    this.name = name;
  }

  run(): void {}
  stop(): void {}
}

export interface Logger {
  log(msg: string): void;
  error(msg: string): void;
}

export type Config = {
  host: string;
  port: number;
};

export const VERSION = "1.0.0";

export function crasher(): never {
  throw new Error("boom");
}

export function badErrors(err: Error): Error {
  try {
    return new Error("something failed");
  } catch (e) {
  }
  return err;
}
`

func TestTSAnalyzer_Discover(t *testing.T) {
	a := ForLanguage("ts")
	if a == nil {
		t.Fatal("expected TS analyzer")
	}

	symbols, err := a.Discover("test.ts", []byte(tsTestSrc))
	if err != nil {
		t.Fatal(err)
	}

	names := map[string]bool{}
	for _, s := range symbols {
		names[s.Name] = true
	}

	expected := []string{"greet", "internal", "process", "bigFunc", "MyService", "Logger", "Config", "VERSION", "crasher", "badErrors"}
	for _, name := range expected {
		if !names[name] {
			t.Errorf("missing symbol %q (found: %v)", name, names)
		}
	}

	// Check kinds
	for _, s := range symbols {
		switch s.Name {
		case "greet":
			if s.Kind != SymbolFunction {
				t.Errorf("greet kind = %v, want function", s.Kind)
			}
			if !s.Exported {
				t.Error("greet should be exported")
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
		case "run":
			if s.Kind != SymbolMethod {
				t.Errorf("run kind = %v, want method", s.Kind)
			}
		}
	}
}

func TestTSAnalyzer_Analyze_Greet(t *testing.T) {
	a := ForLanguage("ts")
	if a == nil {
		t.Fatal("expected TS analyzer")
	}

	m, err := a.Analyze("test.ts", []byte(tsTestSrc), "greet")
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
		t.Error("expected doc comment (JSDoc)")
	}
	if !m.IsExported {
		t.Error("expected exported")
	}
}

func TestTSAnalyzer_Analyze_Process(t *testing.T) {
	a := ForLanguage("ts")
	if a == nil {
		t.Fatal("expected TS analyzer")
	}

	m, err := a.Analyze("test.ts", []byte(tsTestSrc), "process")
	if err != nil {
		t.Fatal(err)
	}

	if m.ParamCount != 2 {
		t.Errorf("param_count = %d, want 2", m.ParamCount)
	}
	if m.MaxNestingDepth < 2 {
		t.Errorf("nesting = %d, want >= 2", m.MaxNestingDepth)
	}
}

func TestTSAnalyzer_Analyze_BigFunc(t *testing.T) {
	a := ForLanguage("ts")
	if a == nil {
		t.Fatal("expected TS analyzer")
	}

	m, err := a.Analyze("test.ts", []byte(tsTestSrc), "bigFunc")
	if err != nil {
		t.Fatal(err)
	}

	if m.MaxNestingDepth < 3 {
		t.Errorf("nesting = %d, want >= 3", m.MaxNestingDepth)
	}
	if m.LoopNestingDepth < 1 {
		t.Errorf("loop_nesting = %d, want >= 1", m.LoopNestingDepth)
	}
	if m.CyclomaticComplexity < 3 {
		t.Errorf("cyclomatic = %d, want >= 3", m.CyclomaticComplexity)
	}
	if m.CognitiveComplexity < 4 {
		t.Errorf("cognitive = %d, want >= 4", m.CognitiveComplexity)
	}
}

func TestTSAnalyzer_Analyze_Class(t *testing.T) {
	a := ForLanguage("ts")
	if a == nil {
		t.Fatal("expected TS analyzer")
	}

	m, err := a.Analyze("test.ts", []byte(tsTestSrc), "MyService")
	if err != nil {
		t.Fatal(err)
	}

	// Constructor + run + stop = 3 methods (or 2 if constructor not counted as method)
	if m.MethodCount < 2 {
		t.Errorf("method_count = %d, want >= 2", m.MethodCount)
	}
	if !m.IsExported {
		t.Error("expected exported")
	}
}

func TestTSAnalyzer_Analyze_EmptyCatch(t *testing.T) {
	a := ForLanguage("ts")
	if a == nil {
		t.Fatal("expected TS analyzer")
	}

	m, err := a.Analyze("test.ts", []byte(tsTestSrc), "badErrors")
	if err != nil {
		t.Fatal(err)
	}

	if m.EmptyCatchBlocks < 1 {
		t.Errorf("empty_catch_blocks = %d, want >= 1", m.EmptyCatchBlocks)
	}
}

func TestTSAnalyzer_Analyze_UnsafeImports(t *testing.T) {
	a := ForLanguage("ts")
	if a == nil {
		t.Fatal("expected TS analyzer")
	}

	m, err := a.Analyze("test.ts", []byte(tsTestSrc), "greet")
	if err != nil {
		t.Fatal(err)
	}

	// child_process is dangerous
	if len(m.UnsafeImports) < 1 {
		t.Errorf("unsafe_imports = %v, want >= 1 (child_process)", m.UnsafeImports)
	}
}

func TestTSAnalyzer_Analyze_NotFound(t *testing.T) {
	a := ForLanguage("ts")
	if a == nil {
		t.Fatal("expected TS analyzer")
	}

	m, err := a.Analyze("test.ts", []byte(tsTestSrc), "nonexistent")
	if err != nil {
		t.Fatal(err)
	}

	if m.ParamCount != 0 {
		t.Error("expected zero metrics for missing symbol")
	}
}

func TestTSAnalyzer_ToEvidence(t *testing.T) {
	a := ForLanguage("ts")
	if a == nil {
		t.Fatal("expected TS analyzer")
	}

	m, err := a.Analyze("test.ts", []byte(tsTestSrc), "bigFunc")
	if err != nil {
		t.Fatal(err)
	}

	ev := m.ToEvidence()

	required := []string{
		"param_count", "return_count", "max_nesting_depth", "has_doc_comment",
		"cognitive_complexity", "unsafe_import_count", "empty_catch_blocks",
	}
	for _, key := range required {
		if _, ok := ev.Metrics[key]; !ok {
			t.Errorf("evidence missing metric %q", key)
		}
	}
}

func TestTSAnalyzer_Language(t *testing.T) {
	a := ForLanguage("ts")
	if a == nil {
		t.Fatal("expected TS analyzer")
	}
	if a.Language() != "ts" {
		t.Errorf("language = %q, want ts", a.Language())
	}
}

func TestTSAnalyzer_AnalyzeFile(t *testing.T) {
	a := ForLanguage("ts")
	if a == nil {
		t.Fatal("expected TS analyzer")
	}

	fm, err := a.AnalyzeFile("test.ts", []byte(tsTestSrc))
	if err != nil {
		t.Fatal(err)
	}

	// TS doesn't have init functions
	if fm.HasInitFunc {
		t.Error("TS should not have init func")
	}
}
