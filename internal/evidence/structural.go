package evidence

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
	"time"
	"unicode"

	"github.com/iksnae/code-certification/internal/domain"
)

// StructuralMetrics holds AST-derived structural analysis results for a code unit.
type StructuralMetrics struct {
	HasDocComment   bool   // Exported func/type has a preceding doc comment
	ParamCount      int    // Number of function parameters (excluding receiver)
	ReturnCount     int    // Number of return values
	MaxNestingDepth int    // Deepest nesting level within the function body
	NakedReturns    int    // Count of bare return statements in named-return functions
	ErrorsIgnored   int    // Count of blank identifier assignments to error-typed returns
	ExportedName    bool   // Symbol starts with uppercase
	ReceiverName    string // Receiver type name (empty for standalone functions)
	IsConstructor   bool   // Function name matches New* pattern

	// Tier 1: new metrics
	FuncLines          int  // Number of lines in function body
	PanicCalls         int  // Count of panic() calls
	OsExitCalls        int  // Count of os.Exit() calls
	DeferInLoop        int  // Count of defer statements inside for/range loops
	ContextNotFirst    bool // context.Context param exists but is not the first param
	MethodCount        int  // Number of methods on a type (type-level only)
	HasInitFunc        bool // File contains an init() function (file-level only)
	GlobalMutableCount int  // Number of package-level var declarations (file-level only)

	// Algorithmic complexity metrics
	AlgoComplexity    string // Estimated complexity class: O(1), O(n), O(n²), O(n³), O(2^n)
	LoopNestingDepth  int    // Maximum depth of nested loops (for/range only)
	RecursiveCalls    int    // Count of direct recursive calls (func calls own name)
	NestedLoopPairs   int    // Count of inner loops nested inside outer loops
	QuadraticPatterns int    // Count of known O(n²) anti-patterns
}

// FileMetrics holds file-level structural analysis results.
type FileMetrics struct {
	HasInitFunc        bool
	GlobalMutableCount int
}

// ToEvidence converts StructuralMetrics to a domain.Evidence.
func (m StructuralMetrics) ToEvidence() domain.Evidence {
	docVal := 0.0
	if m.HasDocComment {
		docVal = 1.0
	}
	exportedVal := 0.0
	if m.ExportedName {
		exportedVal = 1.0
	}
	constructorVal := 0.0
	if m.IsConstructor {
		constructorVal = 1.0
	}

	contextVal := 0.0
	if m.ContextNotFirst {
		contextVal = 1.0
	}
	initVal := 0.0
	if m.HasInitFunc {
		initVal = 1.0
	}

	return domain.Evidence{
		Kind:   domain.EvidenceKindStructural,
		Source: "structural",
		Passed: true,
		Summary: fmt.Sprintf("structural: params=%d returns=%d nesting=%d doc=%v exported=%v",
			m.ParamCount, m.ReturnCount, m.MaxNestingDepth, m.HasDocComment, m.ExportedName),
		Metrics: map[string]float64{
			"has_doc_comment":      docVal,
			"param_count":          float64(m.ParamCount),
			"return_count":         float64(m.ReturnCount),
			"max_nesting_depth":    float64(m.MaxNestingDepth),
			"naked_returns":        float64(m.NakedReturns),
			"errors_ignored":       float64(m.ErrorsIgnored),
			"exported_name":        exportedVal,
			"is_constructor":       constructorVal,
			"func_lines":           float64(m.FuncLines),
			"panic_calls":          float64(m.PanicCalls),
			"os_exit_calls":        float64(m.OsExitCalls),
			"defer_in_loop":        float64(m.DeferInLoop),
			"context_not_first":    contextVal,
			"method_count":         float64(m.MethodCount),
			"has_init_func":        initVal,
			"global_mutable_count": float64(m.GlobalMutableCount),
			"loop_nesting_depth":   float64(m.LoopNestingDepth),
			"recursive_calls":      float64(m.RecursiveCalls),
			"nested_loop_pairs":    float64(m.NestedLoopPairs),
			"quadratic_patterns":   float64(m.QuadraticPatterns),
		},
		Timestamp:  time.Now(),
		Confidence: 1.0,
	}
}

// AnalyzeGoFunc parses Go source and analyzes a function or method by name.
// Returns zero-value StructuralMetrics if the source can't be parsed or symbol not found.
func AnalyzeGoFunc(src string, funcName string) StructuralMetrics {
	if strings.TrimSpace(src) == "" {
		return StructuralMetrics{}
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", src, parser.ParseComments)
	if err != nil {
		return StructuralMetrics{}
	}

	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Name == nil {
			continue
		}
		if fn.Name.Name != funcName {
			continue
		}

		return analyzeFunc(fset, fn)
	}

	return StructuralMetrics{} // not found
}

// AnalyzeGoType parses Go source and analyzes a type declaration by name.
func AnalyzeGoType(src string, typeName string) StructuralMetrics {
	if strings.TrimSpace(src) == "" {
		return StructuralMetrics{}
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", src, parser.ParseComments)
	if err != nil {
		return StructuralMetrics{}
	}

	found := false
	var result StructuralMetrics

	for _, decl := range f.Decls {
		gd, ok := decl.(*ast.GenDecl)
		if !ok || gd.Tok != token.TYPE {
			continue
		}
		for _, spec := range gd.Specs {
			ts, ok := spec.(*ast.TypeSpec)
			if !ok || ts.Name.Name != typeName {
				continue
			}
			result = StructuralMetrics{
				HasDocComment: gd.Doc != nil && len(gd.Doc.List) > 0,
				ExportedName:  isExported(typeName),
			}
			found = true
		}
	}

	if !found {
		return StructuralMetrics{}
	}

	// Count methods on this type
	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Recv == nil || len(fn.Recv.List) == 0 {
			continue
		}
		recvType := exprTypeName(fn.Recv.List[0].Type)
		if recvType == typeName {
			result.MethodCount++
		}
	}

	return result
}

func analyzeFunc(fset *token.FileSet, fn *ast.FuncDecl) StructuralMetrics {
	m := StructuralMetrics{
		HasDocComment: fn.Doc != nil && len(fn.Doc.List) > 0,
		ExportedName:  isExported(fn.Name.Name),
		IsConstructor: strings.HasPrefix(fn.Name.Name, "New"),
	}

	// Receiver
	if fn.Recv != nil && len(fn.Recv.List) > 0 {
		recv := fn.Recv.List[0].Type
		if star, ok := recv.(*ast.StarExpr); ok {
			recv = star.X
		}
		if ident, ok := recv.(*ast.Ident); ok {
			m.ReceiverName = ident.Name
		}
	}

	// Params (excluding receiver)
	if fn.Type.Params != nil {
		for _, field := range fn.Type.Params.List {
			if len(field.Names) == 0 {
				m.ParamCount++ // unnamed parameter
			} else {
				m.ParamCount += len(field.Names)
			}
		}
	}

	// Returns
	if fn.Type.Results != nil {
		for _, field := range fn.Type.Results.List {
			if len(field.Names) == 0 {
				m.ReturnCount++
			} else {
				m.ReturnCount += len(field.Names)
			}
		}
	}

	// Check context.Context position
	m.ContextNotFirst = checkContextNotFirst(fn)

	// Body analysis
	if fn.Body != nil {
		m.MaxNestingDepth = computeNestingDepth(fn.Body)
		m.NakedReturns = countNakedReturns(fn)
		m.ErrorsIgnored = countIgnoredErrors(fn.Body)
		m.FuncLines = countFuncLines(fset, fn)
		m.PanicCalls = countCallExpr(fn.Body, "", "panic")
		m.OsExitCalls = countCallExpr(fn.Body, "os", "Exit")
		m.DeferInLoop = countDeferInLoop(fn.Body)
		analyzeAlgoComplexity(fn, &m)
	}

	if m.AlgoComplexity == "" {
		m.AlgoComplexity = "O(1)"
	}

	return m
}

// computeNestingDepth walks the AST body and returns the max nesting depth.
// Only counts if/for/range/switch/select as nesting levels.
func computeNestingDepth(body *ast.BlockStmt) int {
	maxDepth := 0
	walkBlock(body.List, 0, &maxDepth)
	return maxDepth
}

func walkBlock(stmts []ast.Stmt, depth int, maxDepth *int) {
	for _, stmt := range stmts {
		walkStmt(stmt, depth, maxDepth)
	}
}

func walkStmt(stmt ast.Stmt, depth int, maxDepth *int) {
	switch s := stmt.(type) {
	case *ast.IfStmt:
		walkNested(s.Body, depth, maxDepth)
		if s.Else != nil {
			walkStmt(s.Else, depth, maxDepth) // else is same level as if
		}
	case *ast.ForStmt:
		walkNested(s.Body, depth, maxDepth)
	case *ast.RangeStmt:
		walkNested(s.Body, depth, maxDepth)
	case *ast.SwitchStmt:
		walkNested(s.Body, depth, maxDepth)
	case *ast.TypeSwitchStmt:
		walkNested(s.Body, depth, maxDepth)
	case *ast.SelectStmt:
		walkNested(s.Body, depth, maxDepth)
	case *ast.CaseClause:
		walkBlock(s.Body, depth, maxDepth)
	case *ast.CommClause:
		walkBlock(s.Body, depth, maxDepth)
	case *ast.BlockStmt:
		walkBlock(s.List, depth, maxDepth)
	}
}

// walkNested increments depth for a nested block (if/for/switch/select).
func walkNested(block *ast.BlockStmt, depth int, maxDepth *int) {
	d := depth + 1
	if d > *maxDepth {
		*maxDepth = d
	}
	if block != nil {
		walkBlock(block.List, d, maxDepth)
	}
}

// countNakedReturns counts bare return statements in a function with named results.
func countNakedReturns(fn *ast.FuncDecl) int {
	if fn.Type.Results == nil {
		return 0
	}
	// Check if results are named
	hasNamedResults := false
	for _, field := range fn.Type.Results.List {
		if len(field.Names) > 0 {
			hasNamedResults = true
			break
		}
	}
	if !hasNamedResults {
		return 0
	}

	count := 0
	ast.Inspect(fn.Body, func(n ast.Node) bool {
		ret, ok := n.(*ast.ReturnStmt)
		if ok && len(ret.Results) == 0 {
			count++
		}
		return true
	})
	return count
}

// countIgnoredErrors detects assignments like `_, _ = someFunc()` where
// a blank identifier is used in a position that could be an error.
// Uses a heuristic: blank identifiers in multi-value assign statements.
func countIgnoredErrors(body *ast.BlockStmt) int {
	count := 0
	ast.Inspect(body, func(n ast.Node) bool {
		assign, ok := n.(*ast.AssignStmt)
		if !ok {
			return true
		}
		// Only multi-value assignments (likely function calls returning error)
		if len(assign.Lhs) < 2 {
			return true
		}
		// Go convention: error is the last return value.
		// Only count as "ignored error" when the last LHS is blank identifier.
		// _, err := foo() is fine — the non-error value is discarded, not the error.
		// _, _ = foo() is bad — the error IS discarded.
		last := assign.Lhs[len(assign.Lhs)-1]
		if ident, ok := last.(*ast.Ident); ok && ident.Name == "_" {
			count++
		}
		return true
	})
	return count
}

// countFuncLines returns the number of lines in a function body.
func countFuncLines(fset *token.FileSet, fn *ast.FuncDecl) int {
	if fn.Body == nil {
		return 0
	}
	start := fset.Position(fn.Body.Lbrace)
	end := fset.Position(fn.Body.Rbrace)
	lines := end.Line - start.Line - 1 // exclude braces
	if lines < 0 {
		return 0
	}
	return lines
}

// countCallExpr counts calls to pkg.funcName (or just funcName if pkg is "").
func countCallExpr(body *ast.BlockStmt, pkg, funcName string) int {
	count := 0
	ast.Inspect(body, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}
		if pkg == "" {
			// Bare function call: panic(), etc.
			ident, ok := call.Fun.(*ast.Ident)
			if ok && ident.Name == funcName {
				count++
			}
		} else {
			// Package-qualified: os.Exit(), etc.
			sel, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}
			x, ok := sel.X.(*ast.Ident)
			if ok && x.Name == pkg && sel.Sel.Name == funcName {
				count++
			}
		}
		return true
	})
	return count
}

// countDeferInLoop counts defer statements inside for/range loops.
func countDeferInLoop(body *ast.BlockStmt) int {
	count := 0
	var walk func(node ast.Node, inLoop bool)
	walk = func(node ast.Node, inLoop bool) {
		if node == nil {
			return
		}
		ast.Inspect(node, func(n ast.Node) bool {
			switch s := n.(type) {
			case *ast.ForStmt:
				walk(s.Body, true)
				return false
			case *ast.RangeStmt:
				walk(s.Body, true)
				return false
			case *ast.DeferStmt:
				if inLoop {
					count++
				}
				return false
			}
			return true
		})
	}
	walk(body, false)
	return count
}

// checkContextNotFirst returns true if context.Context appears as a parameter
// but is not the first parameter.
func checkContextNotFirst(fn *ast.FuncDecl) bool {
	if fn.Type.Params == nil {
		return false
	}
	params := fn.Type.Params.List
	hasContext := false
	contextIsFirst := false

	for i, field := range params {
		typeName := exprTypeName(field.Type)
		if typeName == "context.Context" || typeName == "Context" {
			hasContext = true
			if i == 0 {
				contextIsFirst = true
			}
		}
	}
	return hasContext && !contextIsFirst
}

// exprTypeName extracts a readable type name from an AST expression.
func exprTypeName(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		if x, ok := t.X.(*ast.Ident); ok {
			return x.Name + "." + t.Sel.Name
		}
	case *ast.StarExpr:
		return exprTypeName(t.X)
	}
	return ""
}

// AnalyzeGoFile performs file-level structural analysis.
// Detects init() functions and package-level mutable vars.
func AnalyzeGoFile(src string) FileMetrics {
	if strings.TrimSpace(src) == "" {
		return FileMetrics{}
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		return FileMetrics{}
	}

	var fm FileMetrics

	for _, decl := range f.Decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			if d.Name.Name == "init" && d.Recv == nil {
				fm.HasInitFunc = true
			}
		case *ast.GenDecl:
			if d.Tok == token.VAR {
				for _, spec := range d.Specs {
					vs, ok := spec.(*ast.ValueSpec)
					if !ok {
						fm.GlobalMutableCount++
						continue
					}
					if isMutableVar(vs) {
						fm.GlobalMutableCount++
					}
				}
			}
		}
	}

	return fm
}

// analyzeAlgoComplexity computes algorithmic complexity metrics for a function.
// It measures loop nesting depth, recursive calls, nested loop pairs,
// and quadratic anti-patterns to estimate Big-O complexity class.
func analyzeAlgoComplexity(fn *ast.FuncDecl, m *StructuralMetrics) {
	funcName := fn.Name.Name

	// Count recursive calls
	m.RecursiveCalls = countRecursiveCalls(fn.Body, funcName)

	// Compute loop nesting depth and nested loop pairs
	m.LoopNestingDepth, m.NestedLoopPairs = computeLoopNesting(fn.Body)

	// Detect quadratic anti-patterns
	m.QuadraticPatterns = countQuadraticPatterns(fn.Body)

	// Classify
	m.AlgoComplexity = classifyAlgoComplexity(m.LoopNestingDepth, m.RecursiveCalls)
}

// classifyAlgoComplexity returns the estimated Big-O class based on
// loop nesting depth and recursive call count.
func classifyAlgoComplexity(loopDepth, recursiveCalls int) string {
	if recursiveCalls > 0 {
		return "O(2^n)"
	}
	switch {
	case loopDepth >= 3:
		return "O(n³)"
	case loopDepth == 2:
		return "O(n²)"
	case loopDepth == 1:
		return "O(n)"
	default:
		return "O(1)"
	}
}

// countRecursiveCalls counts direct calls to funcName within the body.
func countRecursiveCalls(body *ast.BlockStmt, funcName string) int {
	count := 0
	ast.Inspect(body, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}
		if ident, ok := call.Fun.(*ast.Ident); ok && ident.Name == funcName {
			count++
		}
		return true
	})
	return count
}

// computeLoopNesting returns (maxLoopDepth, nestedPairCount) by walking
// the AST and tracking only for/range loop nesting — not if/switch/select.
func computeLoopNesting(body *ast.BlockStmt) (int, int) {
	maxDepth := 0
	pairs := 0
	walkLoops(body, 0, &maxDepth, &pairs)
	return maxDepth, pairs
}

// walkLoops recursively walks statements, incrementing depth only for loops.
func walkLoops(node ast.Node, depth int, maxDepth, pairs *int) {
	if node == nil {
		return
	}
	ast.Inspect(node, func(n ast.Node) bool {
		switch s := n.(type) {
		case *ast.ForStmt:
			d := depth + 1
			if d > *maxDepth {
				*maxDepth = d
			}
			if d >= 2 {
				*pairs++
			}
			walkLoops(s.Body, d, maxDepth, pairs)
			return false // don't recurse into body again
		case *ast.RangeStmt:
			d := depth + 1
			if d > *maxDepth {
				*maxDepth = d
			}
			if d >= 2 {
				*pairs++
			}
			walkLoops(s.Body, d, maxDepth, pairs)
			return false
		}
		return true
	})
}

// countQuadraticPatterns detects known O(n²) anti-patterns inside loops:
// - String concatenation with += inside a loop
func countQuadraticPatterns(body *ast.BlockStmt) int {
	count := 0
	// Walk looking for loops, then inspect their bodies for patterns
	ast.Inspect(body, func(n ast.Node) bool {
		var loopBody *ast.BlockStmt
		switch s := n.(type) {
		case *ast.ForStmt:
			loopBody = s.Body
		case *ast.RangeStmt:
			loopBody = s.Body
		}
		if loopBody == nil {
			return true
		}
		// Check for string += inside this loop body.
		// Heuristic: flag += only when the RHS contains a string literal or
		// a binary + expression (likely string concatenation). Pure variable
		// += (like total += n) is numeric and fine.
		ast.Inspect(loopBody, func(inner ast.Node) bool {
			assign, ok := inner.(*ast.AssignStmt)
			if !ok || assign.Tok != token.ADD_ASSIGN {
				return true
			}
			for _, rhs := range assign.Rhs {
				if containsStringConcat(rhs) {
					count++
				}
			}
			return true
		})
		return false // don't double-count nested loops
	})
	return count
}

// containsStringConcat returns true if the expression likely involves string
// concatenation — contains a string literal or a binary + expression
// (which in Go is either numeric add or string concat; string literals confirm it).
func containsStringConcat(expr ast.Expr) bool {
	found := false
	ast.Inspect(expr, func(n ast.Node) bool {
		switch n.(type) {
		case *ast.BasicLit:
			lit := n.(*ast.BasicLit)
			if lit.Kind == token.STRING {
				found = true
				return false
			}
		case *ast.BinaryExpr:
			bin := n.(*ast.BinaryExpr)
			if bin.Op == token.ADD {
				// binary + with any string literal descendant → string concat
				if hasStringLit(bin) {
					found = true
					return false
				}
			}
		}
		return true
	})
	return found
}

// hasStringLit returns true if any descendant of n is a string literal.
func hasStringLit(n ast.Node) bool {
	found := false
	ast.Inspect(n, func(node ast.Node) bool {
		if lit, ok := node.(*ast.BasicLit); ok && lit.Kind == token.STRING {
			found = true
			return false
		}
		return !found
	})
	return found
}

// isMutableVar returns true if a package-level var declaration represents
// truly mutable state. Const-like vars (composite literals, error sentinels,
// compiled regexes, etc.) return false.
func isMutableVar(vs *ast.ValueSpec) bool {
	// No initializer → mutable (e.g., `var counter int`)
	if len(vs.Values) == 0 {
		return true
	}
	// Check each initializer expression
	for _, val := range vs.Values {
		if !isConstLikeExpr(val) {
			return true
		}
	}
	return false
}

// isConstLikeExpr returns true if the expression produces an effectively
// immutable value when used as a package-level var initializer.
func isConstLikeExpr(expr ast.Expr) bool {
	switch e := expr.(type) {
	case *ast.BasicLit:
		// Basic literals: var x = "dev", var x = 42
		// These are effectively const — Go just requires var for ldflags.
		_ = e
		return true
	case *ast.CompositeLit:
		// map/slice/struct literals: var x = map[K]V{...}, var x = []T{...}
		return true
	case *ast.UnaryExpr:
		// &T{...} (pointer to composite literal, e.g. &cobra.Command{})
		if e.Op == token.AND {
			_, isLit := e.X.(*ast.CompositeLit)
			return isLit
		}
		return false
	case *ast.CallExpr:
		return isConstLikeCall(e)
	default:
		return false
	}
}

// isConstLikeCall returns true for calls to known constructors that produce
// effectively immutable values: errors.New, fmt.Errorf, regexp.MustCompile, etc.
func isConstLikeCall(call *ast.CallExpr) bool {
	name := callFuncName(call)
	switch name {
	case "make", "new":
		// make(map/slice/chan) and new(T) produce mutable containers
		return false
	}
	// Known immutable constructors
	constLikeConstructors := map[string]bool{
		"errors.New":         true,
		"fmt.Errorf":         true,
		"regexp.MustCompile": true,
		"regexp.Compile":     true,
	}
	if constLikeConstructors[name] {
		return true
	}
	// Any other function call — treat as const-like.
	// Rationale: most package-level var = someFunc() patterns produce
	// values that are never reassigned (loggers, compiled templates, etc.).
	// The truly mutable patterns (make, new, no-init) are caught above.
	return true
}

// callFuncName extracts a readable name from a call expression.
// Returns "pkg.Func" for selector calls, "Func" for ident calls.
func callFuncName(call *ast.CallExpr) string {
	switch fn := call.Fun.(type) {
	case *ast.SelectorExpr:
		if ident, ok := fn.X.(*ast.Ident); ok {
			return ident.Name + "." + fn.Sel.Name
		}
		return fn.Sel.Name
	case *ast.Ident:
		return fn.Name
	}
	return ""
}

func isExported(name string) bool {
	if name == "" {
		return false
	}
	return unicode.IsUpper(rune(name[0]))
}
