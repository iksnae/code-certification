// Package analysis provides language-agnostic code analysis through a unified interface.
// Each language implements the Analyzer interface, producing consistent structural
// metrics that feed into the certification scoring pipeline.
package analysis

import (
	"fmt"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
)

// SymbolKind identifies the type of a discovered code symbol.
type SymbolKind int

const (
	SymbolFunction  SymbolKind = iota // Standalone function
	SymbolMethod                      // Method on a type/class
	SymbolClass                       // Type, class, struct, interface
	SymbolConstant                    // Exported constant or variable
	SymbolInterface                   // Interface (separate from class for languages that distinguish)
)

var symbolKindStrings = map[SymbolKind]string{
	SymbolFunction:  "function",
	SymbolMethod:    "method",
	SymbolClass:     "class",
	SymbolConstant:  "constant",
	SymbolInterface: "interface",
}

func (k SymbolKind) String() string {
	if s, ok := symbolKindStrings[k]; ok {
		return s
	}
	return fmt.Sprintf("SymbolKind(%d)", k)
}

// Symbol represents a discovered code unit within a source file.
type Symbol struct {
	Name      string     // Symbol name
	Kind      SymbolKind // Function, method, class, etc.
	StartLine int        // 1-indexed start line
	EndLine   int        // 1-indexed end line (inclusive)
	Parent    string     // Enclosing type/class name (empty for top-level)
	Exported  bool       // Whether the symbol is exported/public
}

// Metrics holds language-agnostic structural analysis results for a code unit.
type Metrics struct {
	// Shape
	ParamCount      int // Number of function parameters
	ReturnCount     int // Number of return values
	FuncLines       int // Lines in function body
	MaxNestingDepth int // Deepest nesting level (if/for/switch/etc.)

	// Documentation
	HasDocComment bool   // Has a doc comment
	IsExported    bool   // Exported / public symbol
	ReceiverName  string // Receiver type name (empty for standalone)
	IsConstructor bool   // Matches constructor pattern (New*, __init__, constructor)

	// Complexity
	CyclomaticComplexity int    // McCabe cyclomatic complexity
	CognitiveComplexity  int    // Sonar-style cognitive complexity
	LoopNestingDepth     int    // Max nested loop depth (for/while only)
	RecursiveCalls       int    // Direct recursive calls
	AlgoComplexity       string // Estimated Big-O: O(1), O(n), O(n²), etc.

	// Error handling
	ErrorsIgnored    int // Discarded error returns
	ErrorsNotWrapped int // Errors returned without wrapping/context
	NakedReturns     int // Bare return in named-return function
	PanicCalls       int // panic(), throw, unwrap() — unrecoverable exits
	EmptyCatchBlocks int // catch/except/recover with empty body

	// Security
	UnsafeImports    []string // Dangerous imports (os/exec, unsafe, eval, etc.)
	HardcodedSecrets int      // String literals matching secret patterns

	// Design
	MethodCount     int  // Methods on a type (type-level only)
	DeferInLoop     int  // Defer/finally inside loops
	ContextNotFirst bool // context.Context not first param (Go-specific)

	// Performance
	NestedLoopPairs   int // Inner loops nested in outer loops
	QuadraticPatterns int // Known O(n²) anti-patterns

	// File-level
	HasInitFunc        bool // File contains init() or equivalent
	GlobalMutableCount int  // Package-level mutable variables
	OsExitCalls        int  // os.Exit() or sys.exit() calls

	// Language-specific extras
	Extra map[string]float64
}

// FileMetrics holds file-level analysis results.
type FileMetrics struct {
	HasInitFunc        bool
	GlobalMutableCount int
}

// ToEvidence converts Metrics to a domain.Evidence for the scoring pipeline.
func (m Metrics) ToEvidence() domain.Evidence {
	boolToFloat := func(b bool) float64 {
		if b {
			return 1.0
		}
		return 0.0
	}

	metrics := map[string]float64{
		"has_doc_comment":      boolToFloat(m.HasDocComment),
		"param_count":          float64(m.ParamCount),
		"return_count":         float64(m.ReturnCount),
		"max_nesting_depth":    float64(m.MaxNestingDepth),
		"naked_returns":        float64(m.NakedReturns),
		"errors_ignored":       float64(m.ErrorsIgnored),
		"exported_name":        boolToFloat(m.IsExported),
		"is_constructor":       boolToFloat(m.IsConstructor),
		"func_lines":           float64(m.FuncLines),
		"panic_calls":          float64(m.PanicCalls),
		"os_exit_calls":        float64(m.OsExitCalls),
		"defer_in_loop":        float64(m.DeferInLoop),
		"context_not_first":    boolToFloat(m.ContextNotFirst),
		"method_count":         float64(m.MethodCount),
		"has_init_func":        boolToFloat(m.HasInitFunc),
		"global_mutable_count": float64(m.GlobalMutableCount),
		"loop_nesting_depth":   float64(m.LoopNestingDepth),
		"recursive_calls":      float64(m.RecursiveCalls),
		"nested_loop_pairs":    float64(m.NestedLoopPairs),
		"quadratic_patterns":   float64(m.QuadraticPatterns),
		"cognitive_complexity": float64(m.CognitiveComplexity),
		"errors_not_wrapped":   float64(m.ErrorsNotWrapped),
		"unsafe_import_count":  float64(len(m.UnsafeImports)),
		"hardcoded_secrets":    float64(m.HardcodedSecrets),
		"empty_catch_blocks":   float64(m.EmptyCatchBlocks),
	}

	// Merge extras
	for k, v := range m.Extra {
		metrics[k] = v
	}

	return domain.Evidence{
		Kind:   domain.EvidenceKindStructural,
		Source: "structural",
		Passed: true,
		Summary: fmt.Sprintf("structural: params=%d returns=%d nesting=%d doc=%v exported=%v cognitive=%d",
			m.ParamCount, m.ReturnCount, m.MaxNestingDepth, m.HasDocComment, m.IsExported, m.CognitiveComplexity),
		Metrics:    metrics,
		Timestamp:  time.Now(),
		Confidence: 1.0,
	}
}

// Analyzer provides language-agnostic structural analysis.
type Analyzer interface {
	// Language returns the language identifier (e.g., "go", "ts", "py").
	Language() string

	// Discover finds all symbols in a source file.
	Discover(path string, src []byte) ([]Symbol, error)

	// Analyze returns structural metrics for a specific symbol.
	Analyze(path string, src []byte, symbol string) (Metrics, error)

	// AnalyzeFile returns file-level metrics.
	AnalyzeFile(path string, src []byte) (FileMetrics, error)
}

// registry maps language identifiers to analyzer constructors.
var registry = map[string]func() Analyzer{}

// Register adds an analyzer for a language.
func Register(lang string, factory func() Analyzer) {
	registry[lang] = factory
}

// ForLanguage returns an Analyzer for the given language, or nil if unsupported.
func ForLanguage(lang string) Analyzer {
	if factory, ok := registry[lang]; ok {
		return factory()
	}
	return nil
}

// Languages returns all registered language identifiers.
func Languages() []string {
	langs := make([]string, 0, len(registry))
	for lang := range registry {
		langs = append(langs, lang)
	}
	return langs
}
