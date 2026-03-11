package analysis

import (
	"testing"

	"github.com/iksnae/code-certification/internal/evidence"
)

// TestGoAnalyzer_RegressionVsEvidence verifies that the new Go analyzer
// produces identical core metrics to the existing evidence.AnalyzeGoFunc.
// This ensures zero behavior change during the refactor.
func TestGoAnalyzer_RegressionVsEvidence(t *testing.T) {
	src := `package example

import (
	"context"
	"fmt"
	"os"
)

// Greet says hello.
func Greet(name string) string {
	return "Hello, " + name
}

func Process(items []string, ctx context.Context) error {
	for _, item := range items {
		if item == "" {
			continue
		}
		fmt.Println(item)
	}
	return nil
}

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

func (s *MyService) Run() error { return nil }
func (s *MyService) Stop() {}

func NewMyService() *MyService { return &MyService{} }

func crasher() { panic("boom") }
func exiter() { os.Exit(1) }
`

	funcs := []string{"Greet", "Process", "BigFunc", "NewMyService", "crasher", "exiter"}
	a := NewGoAnalyzer()

	for _, name := range funcs {
		t.Run(name, func(t *testing.T) {
			old := evidence.AnalyzeGoFunc(src, name)
			newM, err := a.Analyze("test.go", []byte(src), name)
			if err != nil {
				t.Fatal(err)
			}

			// Compare all core metrics that exist in both
			check := func(field string, oldVal, newVal int) {
				if oldVal != newVal {
					t.Errorf("%s: %s old=%d new=%d", name, field, oldVal, newVal)
				}
			}
			checkBool := func(field string, oldVal, newVal bool) {
				if oldVal != newVal {
					t.Errorf("%s: %s old=%v new=%v", name, field, oldVal, newVal)
				}
			}

			check("ParamCount", old.ParamCount, newM.ParamCount)
			check("ReturnCount", old.ReturnCount, newM.ReturnCount)
			check("MaxNestingDepth", old.MaxNestingDepth, newM.MaxNestingDepth)
			check("NakedReturns", old.NakedReturns, newM.NakedReturns)
			check("ErrorsIgnored", old.ErrorsIgnored, newM.ErrorsIgnored)
			check("FuncLines", old.FuncLines, newM.FuncLines)
			check("PanicCalls", old.PanicCalls, newM.PanicCalls)
			check("OsExitCalls", old.OsExitCalls, newM.OsExitCalls)
			check("DeferInLoop", old.DeferInLoop, newM.DeferInLoop)
			check("MethodCount", old.MethodCount, newM.MethodCount)
			check("LoopNestingDepth", old.LoopNestingDepth, newM.LoopNestingDepth)
			check("RecursiveCalls", old.RecursiveCalls, newM.RecursiveCalls)
			check("NestedLoopPairs", old.NestedLoopPairs, newM.NestedLoopPairs)
			check("QuadraticPatterns", old.QuadraticPatterns, newM.QuadraticPatterns)
			checkBool("HasDocComment", old.HasDocComment, newM.HasDocComment)
			checkBool("IsExported", old.ExportedName, newM.IsExported)
			checkBool("IsConstructor", old.IsConstructor, newM.IsConstructor)
			checkBool("ContextNotFirst", old.ContextNotFirst, newM.ContextNotFirst)

			if old.ReceiverName != newM.ReceiverName {
				t.Errorf("%s: ReceiverName old=%q new=%q", name, old.ReceiverName, newM.ReceiverName)
			}
			if old.AlgoComplexity != newM.AlgoComplexity {
				t.Errorf("%s: AlgoComplexity old=%q new=%q", name, old.AlgoComplexity, newM.AlgoComplexity)
			}
		})
	}

	// Also test type analysis
	t.Run("MyService_type", func(t *testing.T) {
		old := evidence.AnalyzeGoType(src, "MyService")
		newM, err := a.Analyze("test.go", []byte(src), "MyService")
		if err != nil {
			t.Fatal(err)
		}

		if old.MethodCount != newM.MethodCount {
			t.Errorf("MyService MethodCount old=%d new=%d", old.MethodCount, newM.MethodCount)
		}
		if old.HasDocComment != newM.HasDocComment {
			t.Errorf("MyService HasDocComment old=%v new=%v", old.HasDocComment, newM.HasDocComment)
		}
		if old.ExportedName != newM.IsExported {
			t.Errorf("MyService IsExported old=%v new=%v", old.ExportedName, newM.IsExported)
		}
	})

	// File-level regression
	t.Run("file_level", func(t *testing.T) {
		oldFile := evidence.AnalyzeGoFile(src)
		newFile, err := a.AnalyzeFile("test.go", []byte(src))
		if err != nil {
			t.Fatal(err)
		}

		if oldFile.HasInitFunc != newFile.HasInitFunc {
			t.Errorf("HasInitFunc old=%v new=%v", oldFile.HasInitFunc, newFile.HasInitFunc)
		}
		if oldFile.GlobalMutableCount != newFile.GlobalMutableCount {
			t.Errorf("GlobalMutableCount old=%d new=%d", oldFile.GlobalMutableCount, newFile.GlobalMutableCount)
		}
	})
}
