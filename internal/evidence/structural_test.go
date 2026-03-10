package evidence_test

import (
	"testing"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/evidence"
)

func TestAnalyzeGoFunc_WithDocComment(t *testing.T) {
	src := `package foo

// Greet says hello to the user.
func Greet(name string) string {
	return "hello " + name
}
`
	m := evidence.AnalyzeGoFunc(src, "Greet")
	if !m.HasDocComment {
		t.Error("HasDocComment should be true")
	}
	if !m.ExportedName {
		t.Error("ExportedName should be true for Greet")
	}
}

func TestAnalyzeGoFunc_WithoutDocComment(t *testing.T) {
	src := `package foo

func Greet(name string) string {
	return "hello " + name
}
`
	m := evidence.AnalyzeGoFunc(src, "Greet")
	if m.HasDocComment {
		t.Error("HasDocComment should be false")
	}
}

func TestAnalyzeGoFunc_UnexportedNoDocPenalty(t *testing.T) {
	src := `package foo

func greet(name string) string {
	return "hello " + name
}
`
	m := evidence.AnalyzeGoFunc(src, "greet")
	if m.ExportedName {
		t.Error("ExportedName should be false for lowercase greet")
	}
	if m.HasDocComment {
		t.Error("HasDocComment should be false")
	}
}

func TestAnalyzeGoFunc_ParamCount(t *testing.T) {
	src := `package foo

func Process(a int, b string, c bool, d float64, e []byte) error {
	return nil
}
`
	m := evidence.AnalyzeGoFunc(src, "Process")
	if m.ParamCount != 5 {
		t.Errorf("ParamCount = %d, want 5", m.ParamCount)
	}
}

func TestAnalyzeGoFunc_ZeroParams(t *testing.T) {
	src := `package foo

func NoArgs() {}
`
	m := evidence.AnalyzeGoFunc(src, "NoArgs")
	if m.ParamCount != 0 {
		t.Errorf("ParamCount = %d, want 0", m.ParamCount)
	}
}

func TestAnalyzeGoFunc_ReturnCount(t *testing.T) {
	src := `package foo

func Multi() (int, string, error) {
	return 0, "", nil
}
`
	m := evidence.AnalyzeGoFunc(src, "Multi")
	if m.ReturnCount != 3 {
		t.Errorf("ReturnCount = %d, want 3", m.ReturnCount)
	}
}

func TestAnalyzeGoFunc_NestingDepth(t *testing.T) {
	src := `package foo

func Deep(x int) {
	if x > 0 {
		for i := 0; i < x; i++ {
			if i%2 == 0 {
				if i > 5 {
					println(i)
				}
			}
		}
	}
}
`
	m := evidence.AnalyzeGoFunc(src, "Deep")
	if m.MaxNestingDepth != 4 {
		t.Errorf("MaxNestingDepth = %d, want 4", m.MaxNestingDepth)
	}
}

func TestAnalyzeGoFunc_FlatNesting(t *testing.T) {
	src := `package foo

func Flat() {
	println("hello")
}
`
	m := evidence.AnalyzeGoFunc(src, "Flat")
	if m.MaxNestingDepth != 0 {
		t.Errorf("MaxNestingDepth = %d, want 0", m.MaxNestingDepth)
	}
}

func TestAnalyzeGoFunc_NakedReturns(t *testing.T) {
	src := `package foo

func WithNaked() (result int, err error) {
	result = 42
	return
}
`
	m := evidence.AnalyzeGoFunc(src, "WithNaked")
	if m.NakedReturns != 1 {
		t.Errorf("NakedReturns = %d, want 1", m.NakedReturns)
	}
}

func TestAnalyzeGoFunc_NoNakedReturns(t *testing.T) {
	src := `package foo

func Explicit() (int, error) {
	return 42, nil
}
`
	m := evidence.AnalyzeGoFunc(src, "Explicit")
	if m.NakedReturns != 0 {
		t.Errorf("NakedReturns = %d, want 0", m.NakedReturns)
	}
}

func TestAnalyzeGoFunc_ErrorsIgnored(t *testing.T) {
	src := `package foo

import "os"

func BadError() {
	_, _ = os.Open("file.txt")
}
`
	m := evidence.AnalyzeGoFunc(src, "BadError")
	if m.ErrorsIgnored != 1 {
		t.Errorf("ErrorsIgnored = %d, want 1", m.ErrorsIgnored)
	}
}

func TestAnalyzeGoFunc_ErrorsNotIgnored(t *testing.T) {
	src := `package foo

import "os"

func GoodError() error {
	f, err := os.Open("file.txt")
	if err != nil {
		return err
	}
	_ = f
	return nil
}
`
	m := evidence.AnalyzeGoFunc(src, "GoodError")
	if m.ErrorsIgnored != 0 {
		t.Errorf("ErrorsIgnored = %d, want 0", m.ErrorsIgnored)
	}
}

func TestAnalyzeGoFunc_ErrorsIgnored_DiscardedValueNotError(t *testing.T) {
	// _, err := fmt.Sscanf(...) discards the item count, NOT the error.
	// The error IS checked. This should NOT count as an ignored error.
	src := `package foo

import "fmt"

func ParseCount(s string) (int, error) {
	var n int
	_, err := fmt.Sscanf(s, "%d", &n)
	if err != nil {
		return 0, err
	}
	return n, nil
}
`
	m := evidence.AnalyzeGoFunc(src, "ParseCount")
	if m.ErrorsIgnored != 0 {
		t.Errorf("ErrorsIgnored = %d, want 0 (error is checked, only value discarded)", m.ErrorsIgnored)
	}
}

func TestAnalyzeGoFunc_ErrorsIgnored_StatCheck(t *testing.T) {
	// _, err := os.Stat(...) discards the FileInfo, checks the error.
	// Should NOT count as ignored error.
	src := `package foo

import "os"

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
`
	m := evidence.AnalyzeGoFunc(src, "Exists")
	if m.ErrorsIgnored != 0 {
		t.Errorf("ErrorsIgnored = %d, want 0 (error is checked)", m.ErrorsIgnored)
	}
}

func TestAnalyzeGoFunc_ErrorsIgnored_BothDiscarded(t *testing.T) {
	// _, _ = os.Open(...) discards BOTH value and error.
	// This IS an ignored error.
	src := `package foo

import "os"

func Bad() {
	_, _ = os.Open("file.txt")
}
`
	m := evidence.AnalyzeGoFunc(src, "Bad")
	if m.ErrorsIgnored != 1 {
		t.Errorf("ErrorsIgnored = %d, want 1 (both returns discarded)", m.ErrorsIgnored)
	}
}

func TestAnalyzeGoFunc_Constructor(t *testing.T) {
	src := `package foo

func NewFoo() *Foo {
	return &Foo{}
}
`
	m := evidence.AnalyzeGoFunc(src, "NewFoo")
	if !m.IsConstructor {
		t.Error("IsConstructor should be true for NewFoo")
	}
}

func TestAnalyzeGoFunc_NotConstructor(t *testing.T) {
	src := `package foo

func DoStuff() {}
`
	m := evidence.AnalyzeGoFunc(src, "DoStuff")
	if m.IsConstructor {
		t.Error("IsConstructor should be false for DoStuff")
	}
}

func TestAnalyzeGoFunc_Method(t *testing.T) {
	src := `package foo

type Server struct{}

// Handle processes a request.
func (s *Server) Handle(req string) error {
	return nil
}
`
	m := evidence.AnalyzeGoFunc(src, "Handle")
	if m.ReceiverName != "Server" {
		t.Errorf("ReceiverName = %q, want Server", m.ReceiverName)
	}
	if !m.HasDocComment {
		t.Error("HasDocComment should be true")
	}
	if m.ParamCount != 1 {
		t.Errorf("ParamCount = %d, want 1 (excluding receiver)", m.ParamCount)
	}
}

func TestAnalyzeGoType_WithDoc(t *testing.T) {
	src := `package foo

// Config holds application configuration.
type Config struct {
	Host string
	Port int
}
`
	m := evidence.AnalyzeGoType(src, "Config")
	if !m.HasDocComment {
		t.Error("HasDocComment should be true for Config")
	}
	if !m.ExportedName {
		t.Error("ExportedName should be true")
	}
}

func TestAnalyzeGoType_WithoutDoc(t *testing.T) {
	src := `package foo

type config struct {
	host string
}
`
	m := evidence.AnalyzeGoType(src, "config")
	if m.HasDocComment {
		t.Error("HasDocComment should be false")
	}
	if m.ExportedName {
		t.Error("ExportedName should be false")
	}
}

func TestStructuralMetrics_ToEvidence(t *testing.T) {
	m := evidence.StructuralMetrics{
		HasDocComment:   true,
		ParamCount:      3,
		ReturnCount:     2,
		MaxNestingDepth: 1,
		NakedReturns:    0,
		ErrorsIgnored:   0,
		ExportedName:    true,
		IsConstructor:   false,
	}
	ev := m.ToEvidence()

	if ev.Kind != domain.EvidenceKindStructural {
		t.Errorf("Kind = %v, want structural", ev.Kind)
	}
	if ev.Source != "structural" {
		t.Errorf("Source = %q, want structural", ev.Source)
	}
	if ev.Metrics["has_doc_comment"] != 1.0 {
		t.Errorf("has_doc_comment = %f, want 1.0", ev.Metrics["has_doc_comment"])
	}
	if ev.Metrics["param_count"] != 3 {
		t.Errorf("param_count = %f, want 3", ev.Metrics["param_count"])
	}
	if ev.Metrics["return_count"] != 2 {
		t.Errorf("return_count = %f, want 2", ev.Metrics["return_count"])
	}
	if ev.Metrics["exported_name"] != 1.0 {
		t.Errorf("exported_name = %f, want 1.0", ev.Metrics["exported_name"])
	}
}

func TestAnalyzeGoFunc_FuncLines(t *testing.T) {
	src := `package foo

func Big() {
	a := 1
	b := 2
	c := 3
	d := 4
	e := 5
	_ = a + b + c + d + e
}
`
	m := evidence.AnalyzeGoFunc(src, "Big")
	if m.FuncLines != 6 {
		t.Errorf("FuncLines = %d, want 6", m.FuncLines)
	}
}

func TestAnalyzeGoFunc_PanicCalls(t *testing.T) {
	src := `package lib

func Bad() {
	panic("boom")
}
`
	m := evidence.AnalyzeGoFunc(src, "Bad")
	if m.PanicCalls != 1 {
		t.Errorf("PanicCalls = %d, want 1", m.PanicCalls)
	}
}

func TestAnalyzeGoFunc_PanicCalls_Multiple(t *testing.T) {
	src := `package lib

func Worse() {
	if true {
		panic("one")
	}
	panic("two")
}
`
	m := evidence.AnalyzeGoFunc(src, "Worse")
	if m.PanicCalls != 2 {
		t.Errorf("PanicCalls = %d, want 2", m.PanicCalls)
	}
}

func TestAnalyzeGoFunc_OsExitCalls(t *testing.T) {
	src := `package lib

import "os"

func Bail() {
	os.Exit(1)
}
`
	m := evidence.AnalyzeGoFunc(src, "Bail")
	if m.OsExitCalls != 1 {
		t.Errorf("OsExitCalls = %d, want 1", m.OsExitCalls)
	}
}

func TestAnalyzeGoFunc_DeferInLoop(t *testing.T) {
	src := `package lib

import "os"

func Leaky() {
	for i := 0; i < 10; i++ {
		f, _ := os.Open("file")
		defer f.Close()
	}
}
`
	m := evidence.AnalyzeGoFunc(src, "Leaky")
	if m.DeferInLoop != 1 {
		t.Errorf("DeferInLoop = %d, want 1", m.DeferInLoop)
	}
}

func TestAnalyzeGoFunc_DeferInRange(t *testing.T) {
	src := `package lib

import "os"

func LeakyRange() {
	files := []string{"a", "b"}
	for _, name := range files {
		f, _ := os.Open(name)
		defer f.Close()
	}
}
`
	m := evidence.AnalyzeGoFunc(src, "LeakyRange")
	if m.DeferInLoop != 1 {
		t.Errorf("DeferInLoop = %d, want 1", m.DeferInLoop)
	}
}

func TestAnalyzeGoFunc_NoDeferInLoop(t *testing.T) {
	src := `package lib

import "os"

func Clean() {
	f, _ := os.Open("file")
	defer f.Close()
}
`
	m := evidence.AnalyzeGoFunc(src, "Clean")
	if m.DeferInLoop != 0 {
		t.Errorf("DeferInLoop = %d, want 0", m.DeferInLoop)
	}
}

func TestAnalyzeGoFunc_ContextNotFirst(t *testing.T) {
	src := `package lib

import "context"

func Bad(name string, ctx context.Context) {}
`
	m := evidence.AnalyzeGoFunc(src, "Bad")
	if !m.ContextNotFirst {
		t.Error("ContextNotFirst should be true")
	}
}

func TestAnalyzeGoFunc_ContextFirst(t *testing.T) {
	src := `package lib

import "context"

func Good(ctx context.Context, name string) {}
`
	m := evidence.AnalyzeGoFunc(src, "Good")
	if m.ContextNotFirst {
		t.Error("ContextNotFirst should be false when ctx is first")
	}
}

func TestAnalyzeGoFunc_NoContext(t *testing.T) {
	src := `package lib

func Simple(name string) {}
`
	m := evidence.AnalyzeGoFunc(src, "Simple")
	if m.ContextNotFirst {
		t.Error("ContextNotFirst should be false when no context param")
	}
}

func TestAnalyzeGoFile_InitFunc(t *testing.T) {
	src := `package lib

func init() {
	setupGlobals()
}

func Public() {}
`
	m := evidence.AnalyzeGoFile(src)
	if !m.HasInitFunc {
		t.Error("HasInitFunc should be true")
	}
}

func TestAnalyzeGoFile_NoInitFunc(t *testing.T) {
	src := `package lib

func Public() {}
`
	m := evidence.AnalyzeGoFile(src)
	if m.HasInitFunc {
		t.Error("HasInitFunc should be false")
	}
}

func TestAnalyzeGoFile_GlobalMutable(t *testing.T) {
	src := `package lib

var globalCounter int
var anotherGlobal = make(map[string]int)

const maxSize = 100

func Foo() {}
`
	m := evidence.AnalyzeGoFile(src)
	if m.GlobalMutableCount != 2 {
		t.Errorf("GlobalMutableCount = %d, want 2", m.GlobalMutableCount)
	}
}

func TestAnalyzeGoFile_NoGlobalMutable(t *testing.T) {
	src := `package lib

const maxSize = 100

func Foo() {}
`
	m := evidence.AnalyzeGoFile(src)
	if m.GlobalMutableCount != 0 {
		t.Errorf("GlobalMutableCount = %d, want 0", m.GlobalMutableCount)
	}
}

func TestAnalyzeGoFile_ConstLikeVars(t *testing.T) {
	tests := []struct {
		name    string
		src     string
		wantGMC int
	}{
		{
			name: "map literal is const-like",
			src: `package lib
var lookupTable = map[string]string{"a": "b", "c": "d"}
`,
			wantGMC: 0,
		},
		{
			name: "slice literal is const-like",
			src: `package lib
var names = []string{"alice", "bob"}
`,
			wantGMC: 0,
		},
		{
			name: "uninitialized var is mutable",
			src: `package lib
var counter int
`,
			wantGMC: 1,
		},
		{
			name: "make() is mutable",
			src: `package lib
var cache = make(map[string]int)
`,
			wantGMC: 1,
		},
		{
			name: "make chan is mutable",
			src: `package lib
var ch = make(chan int)
`,
			wantGMC: 1,
		},
		{
			name: "new() is mutable",
			src: `package lib
import "bytes"
var buf = new(bytes.Buffer)
`,
			wantGMC: 1,
		},
		{
			name: "errors.New is const-like",
			src: `package lib
import "errors"
var ErrNotFound = errors.New("not found")
`,
			wantGMC: 0,
		},
		{
			name: "regexp.MustCompile is const-like",
			src: `package lib
import "regexp"
var re = regexp.MustCompile("^[a-z]+$")
`,
			wantGMC: 0,
		},
		{
			name: "struct literal is const-like",
			src: `package lib
type config struct{ name string }
var defaultCfg = config{name: "default"}
`,
			wantGMC: 0,
		},
		{
			name: "mixed const-like and mutable",
			src: `package lib
import "errors"
var lookupTable = map[string]int{"a": 1}
var ErrBad = errors.New("bad")
var counter int
var cache = make(map[string]string)
`,
			wantGMC: 2,
		},
		{
			name: "var block with mixed",
			src: `package lib
var (
	names = []string{"a", "b"}
	counter int
)
`,
			wantGMC: 1,
		},
		{
			name: "cobra command is const-like",
			src: `package lib
var rootCmd = &cobra.Command{Use: "certify"}
`,
			wantGMC: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := evidence.AnalyzeGoFile(tt.src)
			if m.GlobalMutableCount != tt.wantGMC {
				t.Errorf("GlobalMutableCount = %d, want %d", m.GlobalMutableCount, tt.wantGMC)
			}
		})
	}
}

func TestAnalyzeGoType_MethodCount(t *testing.T) {
	src := `package lib

type Server struct{}

func NewServer() *Server { return &Server{} }
func (s *Server) Start() error { return nil }
func (s *Server) Stop() error { return nil }
func (s *Server) Handle(path string) {}
func (s Server) Name() string { return "" }
`
	m := evidence.AnalyzeGoType(src, "Server")
	if m.MethodCount != 4 {
		t.Errorf("MethodCount = %d, want 4 (excluding NewServer constructor)", m.MethodCount)
	}
}

func TestAnalyzeGoType_NoMethods(t *testing.T) {
	src := `package lib

type Config struct {
	Name string
}
`
	m := evidence.AnalyzeGoType(src, "Config")
	if m.MethodCount != 0 {
		t.Errorf("MethodCount = %d, want 0", m.MethodCount)
	}
}

func TestAnalyzeGoFunc_ParseError(t *testing.T) {
	m := evidence.AnalyzeGoFunc("not valid go code {{{", "Foo")
	if m.ParamCount != 0 && m.ReturnCount != 0 {
		t.Error("parse error should return zero metrics")
	}
}

func TestAnalyzeGoFunc_NotFound(t *testing.T) {
	src := `package foo

func Bar() {}
`
	m := evidence.AnalyzeGoFunc(src, "NonExistent")
	if m.ParamCount != 0 {
		t.Error("not-found symbol should return zero metrics")
	}
}
