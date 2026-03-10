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
	FuncLines        int  // Number of lines in function body
	PanicCalls       int  // Count of panic() calls
	OsExitCalls      int  // Count of os.Exit() calls
	DeferInLoop      int  // Count of defer statements inside for/range loops
	ContextNotFirst  bool // context.Context param exists but is not the first param
	MethodCount      int  // Number of methods on a type (type-level only)
	HasInitFunc      bool // File contains an init() function (file-level only)
	GlobalMutableCount int // Number of package-level var declarations (file-level only)
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
		d := depth + 1
		if d > *maxDepth {
			*maxDepth = d
		}
		if s.Body != nil {
			walkBlock(s.Body.List, d, maxDepth)
		}
		if s.Else != nil {
			walkStmt(s.Else, depth, maxDepth) // else is same level as if
		}
	case *ast.ForStmt:
		d := depth + 1
		if d > *maxDepth {
			*maxDepth = d
		}
		if s.Body != nil {
			walkBlock(s.Body.List, d, maxDepth)
		}
	case *ast.RangeStmt:
		d := depth + 1
		if d > *maxDepth {
			*maxDepth = d
		}
		if s.Body != nil {
			walkBlock(s.Body.List, d, maxDepth)
		}
	case *ast.SwitchStmt:
		d := depth + 1
		if d > *maxDepth {
			*maxDepth = d
		}
		if s.Body != nil {
			walkBlock(s.Body.List, d, maxDepth)
		}
	case *ast.TypeSwitchStmt:
		d := depth + 1
		if d > *maxDepth {
			*maxDepth = d
		}
		if s.Body != nil {
			walkBlock(s.Body.List, d, maxDepth)
		}
	case *ast.SelectStmt:
		d := depth + 1
		if d > *maxDepth {
			*maxDepth = d
		}
		if s.Body != nil {
			walkBlock(s.Body.List, d, maxDepth)
		}
	case *ast.CaseClause:
		walkBlock(s.Body, depth, maxDepth)
	case *ast.CommClause:
		walkBlock(s.Body, depth, maxDepth)
	case *ast.BlockStmt:
		walkBlock(s.List, depth, maxDepth)
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
		// Check if last LHS is blank identifier (common Go pattern: val, _ = ...)
		// or any blank identifier in the list
		for _, lhs := range assign.Lhs {
			ident, ok := lhs.(*ast.Ident)
			if ok && ident.Name == "_" {
				// This is a heuristic — count blank identifiers in multi-assign
				// that look like error-ignoring patterns
				count++
				return true // count once per statement
			}
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
				for range d.Specs {
					fm.GlobalMutableCount++
				}
			}
		}
	}

	return fm
}

func isExported(name string) bool {
	if name == "" {
		return false
	}
	return unicode.IsUpper(rune(name[0]))
}
