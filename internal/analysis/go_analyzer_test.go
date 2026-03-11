package analysis

import (
	"testing"
)

const goTestSrc = `package example

import (
	"context"
	"fmt"
	"os"
)

// Greet says hello to the user.
func Greet(name string) string {
	return "Hello, " + name
}

func unexported() {}

// Process handles items with context.
func Process(items []string, ctx context.Context) error {
	for _, item := range items {
		if item == "" {
			continue
		}
		fmt.Println(item)
	}
	return nil
}

// BigFunc has deep nesting.
func BigFunc(x int) string {
	if x > 0 {
		for i := 0; i < x; i++ {
			if i%2 == 0 {
				switch {
				case i < 10:
					return "small"
				}
			}
		}
	}
	return "done"
}

type MyService struct{}

// Run starts the service.
func (s *MyService) Run() error {
	return nil
}

// Stop halts the service.
func (s *MyService) Stop() {}

// NewMyService creates a new service.
func NewMyService() *MyService {
	return &MyService{}
}

func crasher() {
	panic("boom")
}

func exiter() {
	os.Exit(1)
}

func init() {
	fmt.Println("init")
}

var mutableGlobal int
var anotherMutable string
`

func TestGoAnalyzer_Discover(t *testing.T) {
	a := NewGoAnalyzer()

	symbols, err := a.Discover("test.go", []byte(goTestSrc))
	if err != nil {
		t.Fatal(err)
	}

	// Check we found key symbols
	names := map[string]bool{}
	for _, s := range symbols {
		names[s.Name] = true
	}

	expected := []string{"Greet", "unexported", "Process", "BigFunc", "Run", "Stop", "NewMyService", "crasher", "exiter", "init", "MyService"}
	for _, name := range expected {
		if !names[name] {
			t.Errorf("missing symbol %q", name)
		}
	}

	// Check symbol kinds
	for _, s := range symbols {
		switch s.Name {
		case "Greet":
			if s.Kind != SymbolFunction {
				t.Errorf("Greet kind = %v, want function", s.Kind)
			}
			if !s.Exported {
				t.Error("Greet should be exported")
			}
			if s.StartLine == 0 || s.EndLine == 0 {
				t.Error("Greet should have line range")
			}
		case "unexported":
			if s.Exported {
				t.Error("unexported should not be exported")
			}
		case "Run":
			if s.Kind != SymbolMethod {
				t.Errorf("Run kind = %v, want method", s.Kind)
			}
			if s.Parent != "MyService" {
				t.Errorf("Run parent = %q, want MyService", s.Parent)
			}
		case "MyService":
			if s.Kind != SymbolClass {
				t.Errorf("MyService kind = %v, want class", s.Kind)
			}
		}
	}
}

func TestGoAnalyzer_Analyze_Greet(t *testing.T) {
	a := NewGoAnalyzer()

	m, err := a.Analyze("test.go", []byte(goTestSrc), "Greet")
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
		t.Error("expected exported")
	}
	if m.MaxNestingDepth != 0 {
		t.Errorf("nesting = %d, want 0", m.MaxNestingDepth)
	}
}

func TestGoAnalyzer_Analyze_ContextNotFirst(t *testing.T) {
	a := NewGoAnalyzer()

	m, err := a.Analyze("test.go", []byte(goTestSrc), "Process")
	if err != nil {
		t.Fatal(err)
	}

	if !m.ContextNotFirst {
		t.Error("expected context_not_first = true")
	}
	if m.ParamCount != 2 {
		t.Errorf("param_count = %d, want 2", m.ParamCount)
	}
}

func TestGoAnalyzer_Analyze_Nesting(t *testing.T) {
	a := NewGoAnalyzer()

	m, err := a.Analyze("test.go", []byte(goTestSrc), "BigFunc")
	if err != nil {
		t.Fatal(err)
	}

	if m.MaxNestingDepth < 3 {
		t.Errorf("nesting = %d, want >= 3", m.MaxNestingDepth)
	}
	if m.LoopNestingDepth < 1 {
		t.Errorf("loop_nesting = %d, want >= 1", m.LoopNestingDepth)
	}
}

func TestGoAnalyzer_Analyze_Method(t *testing.T) {
	a := NewGoAnalyzer()

	m, err := a.Analyze("test.go", []byte(goTestSrc), "Run")
	if err != nil {
		t.Fatal(err)
	}

	if m.ReceiverName != "MyService" {
		t.Errorf("receiver = %q, want MyService", m.ReceiverName)
	}
	if m.ReturnCount != 1 {
		t.Errorf("return_count = %d, want 1", m.ReturnCount)
	}
}

func TestGoAnalyzer_Analyze_Constructor(t *testing.T) {
	a := NewGoAnalyzer()

	m, err := a.Analyze("test.go", []byte(goTestSrc), "NewMyService")
	if err != nil {
		t.Fatal(err)
	}

	if !m.IsConstructor {
		t.Error("expected constructor")
	}
}

func TestGoAnalyzer_Analyze_Panic(t *testing.T) {
	a := NewGoAnalyzer()

	m, err := a.Analyze("test.go", []byte(goTestSrc), "crasher")
	if err != nil {
		t.Fatal(err)
	}

	if m.PanicCalls != 1 {
		t.Errorf("panic_calls = %d, want 1", m.PanicCalls)
	}
}

func TestGoAnalyzer_Analyze_OsExit(t *testing.T) {
	a := NewGoAnalyzer()

	m, err := a.Analyze("test.go", []byte(goTestSrc), "exiter")
	if err != nil {
		t.Fatal(err)
	}

	if m.OsExitCalls != 1 {
		t.Errorf("os_exit_calls = %d, want 1", m.OsExitCalls)
	}
}

func TestGoAnalyzer_Analyze_Type(t *testing.T) {
	a := NewGoAnalyzer()

	m, err := a.Analyze("test.go", []byte(goTestSrc), "MyService")
	if err != nil {
		t.Fatal(err)
	}

	if m.MethodCount != 2 {
		t.Errorf("method_count = %d, want 2", m.MethodCount)
	}
	if !m.IsExported {
		t.Error("expected exported")
	}
}

func TestGoAnalyzer_AnalyzeFile(t *testing.T) {
	a := NewGoAnalyzer()

	fm, err := a.AnalyzeFile("test.go", []byte(goTestSrc))
	if err != nil {
		t.Fatal(err)
	}

	if !fm.HasInitFunc {
		t.Error("expected init func")
	}
	if fm.GlobalMutableCount != 2 {
		t.Errorf("global_mutable = %d, want 2", fm.GlobalMutableCount)
	}
}

func TestGoAnalyzer_Analyze_NotFound(t *testing.T) {
	a := NewGoAnalyzer()

	m, err := a.Analyze("test.go", []byte(goTestSrc), "DoesNotExist")
	if err != nil {
		t.Fatal(err)
	}

	// Should return zero-value metrics (not an error)
	if m.ParamCount != 0 && m.ReturnCount != 0 {
		t.Error("expected zero metrics for missing symbol")
	}
}

func TestGoAnalyzer_Analyze_EmptySource(t *testing.T) {
	a := NewGoAnalyzer()

	m, err := a.Analyze("test.go", []byte(""), "Foo")
	if err != nil {
		t.Fatal(err)
	}

	if m.ParamCount != 0 {
		t.Error("expected zero metrics for empty source")
	}
}

func TestGoAnalyzer_ToEvidence_HasAllMetrics(t *testing.T) {
	a := NewGoAnalyzer()

	m, err := a.Analyze("test.go", []byte(goTestSrc), "Process")
	if err != nil {
		t.Fatal(err)
	}

	ev := m.ToEvidence()

	requiredKeys := []string{
		"has_doc_comment", "param_count", "return_count", "max_nesting_depth",
		"naked_returns", "errors_ignored", "exported_name", "is_constructor",
		"func_lines", "panic_calls", "os_exit_calls", "defer_in_loop",
		"context_not_first", "method_count", "has_init_func", "global_mutable_count",
		"loop_nesting_depth", "recursive_calls", "nested_loop_pairs", "quadratic_patterns",
		"cognitive_complexity", "errors_not_wrapped", "unsafe_import_count",
		"hardcoded_secrets", "empty_catch_blocks",
	}
	for _, key := range requiredKeys {
		if _, ok := ev.Metrics[key]; !ok {
			t.Errorf("evidence missing metric %q", key)
		}
	}
}

func TestGoAnalyzer_Language(t *testing.T) {
	a := NewGoAnalyzer()
	if a.Language() != "go" {
		t.Errorf("language = %q, want go", a.Language())
	}
}

func TestForLanguage_Go(t *testing.T) {
	a := ForLanguage("go")
	if a == nil {
		t.Fatal("expected Go analyzer")
	}
	if a.Language() != "go" {
		t.Errorf("language = %q, want go", a.Language())
	}
}

func TestForLanguage_Unknown(t *testing.T) {
	a := ForLanguage("brainfuck")
	if a != nil {
		t.Error("expected nil for unknown language")
	}
}

// TestGoAnalyzer_CognitiveComplexity tests the new cognitive complexity metric.
func TestGoAnalyzer_CognitiveComplexity(t *testing.T) {
	src := `package example

func simple() int {
	return 1
}

func nested(x int) string {
	if x > 0 {                    // +1
		for i := 0; i < x; i++ {  // +1, +1 nesting
			if i%2 == 0 {          // +1, +2 nesting
				return "even"
			}
		}
	}
	return "done"
}
`
	a := NewGoAnalyzer()

	m1, _ := a.Analyze("test.go", []byte(src), "simple")
	if m1.CognitiveComplexity != 0 {
		t.Errorf("simple cognitive = %d, want 0", m1.CognitiveComplexity)
	}

	m2, _ := a.Analyze("test.go", []byte(src), "nested")
	// if +1 (nesting 0) for outer if, for +1 (nesting 1) = +2, inner if +1 (nesting 2) = +3 → total 6
	if m2.CognitiveComplexity < 4 {
		t.Errorf("nested cognitive = %d, want >= 4", m2.CognitiveComplexity)
	}
}

// TestGoAnalyzer_ErrorsNotWrapped tests detection of unwrapped errors.
func TestGoAnalyzer_ErrorsNotWrapped(t *testing.T) {
	src := `package example

import (
	"errors"
	"fmt"
)

func good(err error) error {
	return fmt.Errorf("context: %w", err)
}

func bad(err error) error {
	return fmt.Errorf("lost: %v", err)
}

func bare(err error) error {
	return errors.New("something failed")
}

func nowrap(err error) error {
	return err
}
`
	a := NewGoAnalyzer()

	m1, _ := a.Analyze("test.go", []byte(src), "good")
	if m1.ErrorsNotWrapped != 0 {
		t.Errorf("good: errors_not_wrapped = %d, want 0", m1.ErrorsNotWrapped)
	}

	m2, _ := a.Analyze("test.go", []byte(src), "bad")
	if m2.ErrorsNotWrapped != 1 {
		t.Errorf("bad: errors_not_wrapped = %d, want 1", m2.ErrorsNotWrapped)
	}
}

// TestGoAnalyzer_UnsafeImports tests detection of dangerous imports.
func TestGoAnalyzer_UnsafeImports(t *testing.T) {
	src := `package example

import (
	"fmt"
	"os/exec"
	"unsafe"
)

func run() {
	_ = fmt.Sprintf("hi")
}
`
	a := NewGoAnalyzer()

	m, _ := a.AnalyzeFile("test.go", []byte(src))
	_ = m // file-level doesn't return unsafe imports

	// Analyze function — unsafe imports are file-level context
	metrics, _ := a.Analyze("test.go", []byte(src), "run")
	if len(metrics.UnsafeImports) < 2 {
		t.Errorf("unsafe_imports = %v, want >= 2 (os/exec, unsafe)", metrics.UnsafeImports)
	}
}

// TestGoAnalyzer_HardcodedSecrets tests detection of hardcoded secrets.
func TestGoAnalyzer_HardcodedSecrets(t *testing.T) {
	src := `package example

func connect() {
	password := "super_secret_123"
	apiKey := "AKIAIOSFODNN7EXAMPLE"
	normal := "hello world"
	_ = password
	_ = apiKey
	_ = normal
}
`
	a := NewGoAnalyzer()

	m, _ := a.Analyze("test.go", []byte(src), "connect")
	if m.HardcodedSecrets < 1 {
		t.Errorf("hardcoded_secrets = %d, want >= 1", m.HardcodedSecrets)
	}
}
