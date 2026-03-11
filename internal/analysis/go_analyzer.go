package analysis

import (
	"go/ast"
	"go/parser"
	"go/token"
	"regexp"
	"strings"
	"unicode"
)

// GoAnalyzer implements Analyzer for Go source code using go/ast.
type GoAnalyzer struct{}

// NewGoAnalyzer creates a new Go analyzer.
func NewGoAnalyzer() *GoAnalyzer {
	return &GoAnalyzer{}
}

func init() {
	Register("go", func() Analyzer { return NewGoAnalyzer() })
}

func (a *GoAnalyzer) Language() string { return "go" }

// Discover finds all functions, methods, and types in Go source.
func (a *GoAnalyzer) Discover(path string, src []byte) ([]Symbol, error) {
	if len(strings.TrimSpace(string(src))) == 0 {
		return nil, nil
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, src, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	var symbols []Symbol

	for _, decl := range f.Decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			sym := Symbol{
				Name:      d.Name.Name,
				StartLine: fset.Position(d.Pos()).Line,
				EndLine:   fset.Position(d.End()).Line,
				Exported:  isExported(d.Name.Name),
			}
			if d.Recv != nil && len(d.Recv.List) > 0 {
				sym.Kind = SymbolMethod
				sym.Parent = goReceiverTypeName(d.Recv.List[0].Type)
			} else {
				sym.Kind = SymbolFunction
			}
			symbols = append(symbols, sym)

		case *ast.GenDecl:
			if d.Tok != token.TYPE {
				continue
			}
			for _, spec := range d.Specs {
				ts, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}
				symbols = append(symbols, Symbol{
					Name:      ts.Name.Name,
					Kind:      SymbolClass,
					StartLine: fset.Position(d.Pos()).Line,
					EndLine:   fset.Position(d.End()).Line,
					Exported:  isExported(ts.Name.Name),
				})
			}
		}
	}

	return symbols, nil
}

// Analyze returns structural metrics for a specific symbol.
func (a *GoAnalyzer) Analyze(path string, src []byte, symbol string) (Metrics, error) {
	if len(strings.TrimSpace(string(src))) == 0 {
		return Metrics{}, nil
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, src, parser.ParseComments)
	if err != nil {
		return Metrics{}, nil
	}

	// Collect file-level context
	unsafeImports := goDetectUnsafeImports(f)
	fileMeta := goAnalyzeFileDecls(f)

	// Try to find as function/method first
	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Name == nil || fn.Name.Name != symbol {
			continue
		}
		m := goAnalyzeFunc(fset, fn)
		m.UnsafeImports = unsafeImports
		m.HasInitFunc = fileMeta.HasInitFunc
		m.GlobalMutableCount = fileMeta.GlobalMutableCount
		m.HardcodedSecrets = goCountHardcodedSecrets(fn.Body)
		return m, nil
	}

	// Try to find as type
	for _, decl := range f.Decls {
		gd, ok := decl.(*ast.GenDecl)
		if !ok || gd.Tok != token.TYPE {
			continue
		}
		for _, spec := range gd.Specs {
			ts, ok := spec.(*ast.TypeSpec)
			if !ok || ts.Name.Name != symbol {
				continue
			}
			m := Metrics{
				HasDocComment: gd.Doc != nil && len(gd.Doc.List) > 0,
				IsExported:    isExported(symbol),
			}
			// Count methods on this type
			m.MethodCount = goCountMethods(f, symbol)
			m.UnsafeImports = unsafeImports
			m.HasInitFunc = fileMeta.HasInitFunc
			m.GlobalMutableCount = fileMeta.GlobalMutableCount
			return m, nil
		}
	}

	return Metrics{}, nil // not found
}

// AnalyzeFile returns file-level metrics.
func (a *GoAnalyzer) AnalyzeFile(path string, src []byte) (FileMetrics, error) {
	if len(strings.TrimSpace(string(src))) == 0 {
		return FileMetrics{}, nil
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, src, 0)
	if err != nil {
		return FileMetrics{}, nil
	}

	return goAnalyzeFileDecls(f), nil
}

// goAnalyzeFunc extracts all structural metrics from a function declaration.
func goAnalyzeFunc(fset *token.FileSet, fn *ast.FuncDecl) Metrics {
	m := Metrics{
		HasDocComment: fn.Doc != nil && len(fn.Doc.List) > 0,
		IsExported:    isExported(fn.Name.Name),
		IsConstructor: strings.HasPrefix(fn.Name.Name, "New"),
	}

	// Receiver
	if fn.Recv != nil && len(fn.Recv.List) > 0 {
		m.ReceiverName = goReceiverTypeName(fn.Recv.List[0].Type)
	}

	// Params
	if fn.Type.Params != nil {
		for _, field := range fn.Type.Params.List {
			if len(field.Names) == 0 {
				m.ParamCount++
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

	// context.Context position
	m.ContextNotFirst = goCheckContextNotFirst(fn)

	// Body analysis
	if fn.Body != nil {
		m.MaxNestingDepth = goComputeNestingDepth(fn.Body)
		m.NakedReturns = goCountNakedReturns(fn)
		m.ErrorsIgnored = goCountIgnoredErrors(fn.Body)
		m.FuncLines = goCountFuncLines(fset, fn)
		m.PanicCalls = goCountCallExpr(fn.Body, "", "panic")
		m.OsExitCalls = goCountCallExpr(fn.Body, "os", "Exit")
		m.DeferInLoop = goCountDeferInLoop(fn.Body)
		m.CyclomaticComplexity = goComputeCyclomaticComplexity(fn)
		m.CognitiveComplexity = goComputeCognitiveComplexity(fn)
		m.ErrorsNotWrapped = goCountErrorsNotWrapped(fn)

		// Algorithmic complexity
		m.RecursiveCalls = goCountRecursiveCalls(fn.Body, fn.Name.Name)
		m.LoopNestingDepth, m.NestedLoopPairs = goComputeLoopNesting(fn.Body)
		m.QuadraticPatterns = goCountQuadraticPatterns(fn.Body)
		m.AlgoComplexity = classifyAlgo(m.LoopNestingDepth, m.RecursiveCalls)
	}

	if m.AlgoComplexity == "" {
		m.AlgoComplexity = "O(1)"
	}

	return m
}

// goAnalyzeFileDecls extracts file-level metrics from declarations.
func goAnalyzeFileDecls(f *ast.File) FileMetrics {
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
					if goIsMutableVar(vs) {
						fm.GlobalMutableCount++
					}
				}
			}
		}
	}
	return fm
}

// goCountMethods counts methods on a named type.
func goCountMethods(f *ast.File, typeName string) int {
	count := 0
	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Recv == nil || len(fn.Recv.List) == 0 {
			continue
		}
		if goReceiverTypeName(fn.Recv.List[0].Type) == typeName {
			count++
		}
	}
	return count
}

// --- Nesting depth ---

func goComputeNestingDepth(body *ast.BlockStmt) int {
	maxDepth := 0
	goWalkBlock(body.List, 0, &maxDepth)
	return maxDepth
}

func goWalkBlock(stmts []ast.Stmt, depth int, maxDepth *int) {
	for _, stmt := range stmts {
		goWalkStmt(stmt, depth, maxDepth)
	}
}

func goWalkStmt(stmt ast.Stmt, depth int, maxDepth *int) {
	switch s := stmt.(type) {
	case *ast.IfStmt:
		goWalkNested(s.Body, depth, maxDepth)
		if s.Else != nil {
			goWalkStmt(s.Else, depth, maxDepth)
		}
	case *ast.ForStmt:
		goWalkNested(s.Body, depth, maxDepth)
	case *ast.RangeStmt:
		goWalkNested(s.Body, depth, maxDepth)
	case *ast.SwitchStmt:
		goWalkNested(s.Body, depth, maxDepth)
	case *ast.TypeSwitchStmt:
		goWalkNested(s.Body, depth, maxDepth)
	case *ast.SelectStmt:
		goWalkNested(s.Body, depth, maxDepth)
	case *ast.CaseClause:
		goWalkBlock(s.Body, depth, maxDepth)
	case *ast.CommClause:
		goWalkBlock(s.Body, depth, maxDepth)
	case *ast.BlockStmt:
		goWalkBlock(s.List, depth, maxDepth)
	}
}

func goWalkNested(block *ast.BlockStmt, depth int, maxDepth *int) {
	d := depth + 1
	if d > *maxDepth {
		*maxDepth = d
	}
	if block != nil {
		goWalkBlock(block.List, d, maxDepth)
	}
}

// --- Naked returns ---

func goCountNakedReturns(fn *ast.FuncDecl) int {
	if fn.Type.Results == nil {
		return 0
	}
	hasNamed := false
	for _, field := range fn.Type.Results.List {
		if len(field.Names) > 0 {
			hasNamed = true
			break
		}
	}
	if !hasNamed {
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

// --- Ignored errors ---

func goCountIgnoredErrors(body *ast.BlockStmt) int {
	count := 0
	ast.Inspect(body, func(n ast.Node) bool {
		assign, ok := n.(*ast.AssignStmt)
		if !ok || len(assign.Lhs) < 2 {
			return true
		}
		last := assign.Lhs[len(assign.Lhs)-1]
		if ident, ok := last.(*ast.Ident); ok && ident.Name == "_" {
			count++
		}
		return true
	})
	return count
}

// --- Func lines ---

func goCountFuncLines(fset *token.FileSet, fn *ast.FuncDecl) int {
	if fn.Body == nil {
		return 0
	}
	start := fset.Position(fn.Body.Lbrace)
	end := fset.Position(fn.Body.Rbrace)
	lines := end.Line - start.Line - 1
	if lines < 0 {
		return 0
	}
	return lines
}

// --- Call expression counting ---

func goCountCallExpr(body *ast.BlockStmt, pkg, funcName string) int {
	count := 0
	ast.Inspect(body, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}
		if pkg == "" {
			if ident, ok := call.Fun.(*ast.Ident); ok && ident.Name == funcName {
				count++
			}
		} else {
			if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
				if x, ok := sel.X.(*ast.Ident); ok && x.Name == pkg && sel.Sel.Name == funcName {
					count++
				}
			}
		}
		return true
	})
	return count
}

// --- Defer in loop ---

func goCountDeferInLoop(body *ast.BlockStmt) int {
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

// --- Context position ---

func goCheckContextNotFirst(fn *ast.FuncDecl) bool {
	if fn.Type.Params == nil {
		return false
	}
	hasContext := false
	contextIsFirst := false
	for i, field := range fn.Type.Params.List {
		typeName := goExprTypeName(field.Type)
		if typeName == "context.Context" || typeName == "Context" {
			hasContext = true
			if i == 0 {
				contextIsFirst = true
			}
		}
	}
	return hasContext && !contextIsFirst
}

// --- Recursive calls ---

func goCountRecursiveCalls(body *ast.BlockStmt, funcName string) int {
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

// --- Loop nesting ---

func goComputeLoopNesting(body *ast.BlockStmt) (int, int) {
	maxDepth := 0
	pairs := 0
	goWalkLoops(body, 0, &maxDepth, &pairs)
	return maxDepth, pairs
}

func goWalkLoops(node ast.Node, depth int, maxDepth, pairs *int) {
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
			goWalkLoops(s.Body, d, maxDepth, pairs)
			return false
		case *ast.RangeStmt:
			d := depth + 1
			if d > *maxDepth {
				*maxDepth = d
			}
			if d >= 2 {
				*pairs++
			}
			goWalkLoops(s.Body, d, maxDepth, pairs)
			return false
		}
		return true
	})
}

// --- Quadratic patterns ---

func goCountQuadraticPatterns(body *ast.BlockStmt) int {
	count := 0
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
		ast.Inspect(loopBody, func(inner ast.Node) bool {
			assign, ok := inner.(*ast.AssignStmt)
			if !ok || assign.Tok != token.ADD_ASSIGN {
				return true
			}
			for _, rhs := range assign.Rhs {
				if goContainsStringConcat(rhs) {
					count++
				}
			}
			return true
		})
		return false
	})
	return count
}

func goContainsStringConcat(expr ast.Expr) bool {
	found := false
	ast.Inspect(expr, func(n ast.Node) bool {
		if lit, ok := n.(*ast.BasicLit); ok && lit.Kind == token.STRING {
			found = true
			return false
		}
		return !found
	})
	return found
}

// --- Cyclomatic complexity ---

func goComputeCyclomaticComplexity(fn *ast.FuncDecl) int {
	if fn.Body == nil {
		return 1
	}
	complexity := 1
	ast.Inspect(fn.Body, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.IfStmt:
			complexity++
		case *ast.ForStmt:
			complexity++
		case *ast.RangeStmt:
			complexity++
		case *ast.CaseClause:
			if node.List != nil {
				complexity++
			}
		case *ast.CommClause:
			if node.Comm != nil {
				complexity++
			}
		case *ast.BinaryExpr:
			if node.Op == token.LAND || node.Op == token.LOR {
				complexity++
			}
		case *ast.SelectStmt:
			complexity++
		}
		return true
	})
	return complexity
}

// --- Cognitive complexity (Sonar-style) ---

func goComputeCognitiveComplexity(fn *ast.FuncDecl) int {
	if fn.Body == nil {
		return 0
	}
	total := 0
	goCognitiveWalk(fn.Body.List, 0, &total)
	return total
}

func goCognitiveWalk(stmts []ast.Stmt, nesting int, total *int) {
	for _, stmt := range stmts {
		goCognitiveStmt(stmt, nesting, total)
	}
}

func goCognitiveStmt(stmt ast.Stmt, nesting int, total *int) {
	switch s := stmt.(type) {
	case *ast.IfStmt:
		// +1 for if, +nesting for depth
		*total += 1 + nesting
		if s.Body != nil {
			goCognitiveWalk(s.Body.List, nesting+1, total)
		}
		if s.Else != nil {
			switch e := s.Else.(type) {
			case *ast.IfStmt:
				// else if: +1 (no nesting increment, it's linear)
				*total += 1
				if e.Body != nil {
					goCognitiveWalk(e.Body.List, nesting+1, total)
				}
				if e.Else != nil {
					goCognitiveStmt(e.Else, nesting, total)
				}
			case *ast.BlockStmt:
				// else: +1
				*total += 1
				goCognitiveWalk(e.List, nesting+1, total)
			}
		}

	case *ast.ForStmt:
		*total += 1 + nesting
		if s.Body != nil {
			goCognitiveWalk(s.Body.List, nesting+1, total)
		}

	case *ast.RangeStmt:
		*total += 1 + nesting
		if s.Body != nil {
			goCognitiveWalk(s.Body.List, nesting+1, total)
		}

	case *ast.SwitchStmt:
		*total += 1 + nesting
		if s.Body != nil {
			goCognitiveWalk(s.Body.List, nesting+1, total)
		}

	case *ast.TypeSwitchStmt:
		*total += 1 + nesting
		if s.Body != nil {
			goCognitiveWalk(s.Body.List, nesting+1, total)
		}

	case *ast.SelectStmt:
		*total += 1 + nesting
		if s.Body != nil {
			goCognitiveWalk(s.Body.List, nesting+1, total)
		}

	case *ast.CaseClause:
		// Case clauses don't increment — they're linear branches
		goCognitiveWalk(s.Body, nesting, total)

	case *ast.CommClause:
		goCognitiveWalk(s.Body, nesting, total)

	case *ast.BlockStmt:
		goCognitiveWalk(s.List, nesting, total)

	case *ast.LabeledStmt:
		goCognitiveStmt(s.Stmt, nesting, total)

	case *ast.BranchStmt:
		// break/continue to label: +1
		if s.Label != nil {
			*total += 1
		}

	default:
		// Check for boolean sequences in expressions within statements
		ast.Inspect(stmt, func(n ast.Node) bool {
			if bin, ok := n.(*ast.BinaryExpr); ok {
				if bin.Op == token.LAND || bin.Op == token.LOR {
					*total += 1
				}
			}
			return true
		})
	}
}

// --- Error wrapping detection ---

func goCountErrorsNotWrapped(fn *ast.FuncDecl) int {
	if fn.Body == nil {
		return 0
	}
	count := 0
	ast.Inspect(fn.Body, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}
		// Check fmt.Errorf calls for missing %w
		if sel, ok := call.Fun.(*ast.SelectorExpr); ok {
			if x, ok := sel.X.(*ast.Ident); ok && x.Name == "fmt" && sel.Sel.Name == "Errorf" {
				if len(call.Args) > 0 {
					if lit, ok := call.Args[0].(*ast.BasicLit); ok && lit.Kind == token.STRING {
						if !strings.Contains(lit.Value, "%w") {
							count++
						}
					}
				}
			}
		}
		return true
	})
	return count
}

// --- Unsafe imports ---

// dangerousGoImports lists imports considered security-sensitive.
var dangerousGoImports = map[string]bool{
	"unsafe":       true,
	"os/exec":      true,
	"net/http/cgi": true,
	"crypto/md5":   true,
	"crypto/sha1":  true,
	"crypto/des":   true,
	"crypto/rc4":   true,
	"encoding/gob": true,
	"plugin":       true,
	"debug/elf":    true,
	"debug/pe":     true,
	"reflect":      false, // tracked but not flagged by default
	"syscall":      false,
}

func goDetectUnsafeImports(f *ast.File) []string {
	var unsafe []string
	for _, imp := range f.Imports {
		path := strings.Trim(imp.Path.Value, `"`)
		if flagged, exists := dangerousGoImports[path]; exists && flagged {
			unsafe = append(unsafe, path)
		}
	}
	return unsafe
}

// --- Hardcoded secrets ---

var secretPatterns = []*regexp.Regexp{
	regexp.MustCompile(`(?i)password\s*[:=]\s*["']`),
	regexp.MustCompile(`(?i)passwd\s*[:=]\s*["']`),
	regexp.MustCompile(`(?i)secret\s*[:=]\s*["']`),
	regexp.MustCompile(`(?i)api[_-]?key\s*[:=]\s*["']`),
	regexp.MustCompile(`(?i)token\s*[:=]\s*["']`),
	regexp.MustCompile(`(?i)auth\s*[:=]\s*["']`),
	regexp.MustCompile(`AKIA[0-9A-Z]{16}`), // AWS access key
	regexp.MustCompile(`(?i)private[_-]?key`),
}

// Variable name patterns that suggest secrets.
var secretVarPatterns = []*regexp.Regexp{
	regexp.MustCompile(`(?i)^(password|passwd|secret|apiKey|api_key|token|auth_token|private_key|access_key)$`),
}

func goCountHardcodedSecrets(body *ast.BlockStmt) int {
	if body == nil {
		return 0
	}
	count := 0
	ast.Inspect(body, func(n ast.Node) bool {
		assign, ok := n.(*ast.AssignStmt)
		if !ok {
			return true
		}
		for i, lhs := range assign.Lhs {
			ident, ok := lhs.(*ast.Ident)
			if !ok {
				continue
			}
			// Check if variable name matches secret patterns
			isSecretVar := false
			for _, pat := range secretVarPatterns {
				if pat.MatchString(ident.Name) {
					isSecretVar = true
					break
				}
			}
			if !isSecretVar {
				continue
			}
			// Check if RHS is a string literal
			if i < len(assign.Rhs) {
				if lit, ok := assign.Rhs[i].(*ast.BasicLit); ok && lit.Kind == token.STRING {
					val := strings.Trim(lit.Value, `"'`+"`")
					// Exclude empty strings, env var references, placeholders
					if len(val) > 3 && !strings.HasPrefix(val, "${") && !strings.HasPrefix(val, "$(") && val != "TODO" && val != "CHANGEME" {
						count++
					}
				}
			}
		}
		return true
	})
	return count
}

// --- Mutable var detection ---

func goIsMutableVar(vs *ast.ValueSpec) bool {
	if len(vs.Values) == 0 {
		return true
	}
	for _, val := range vs.Values {
		if !goIsConstLikeExpr(val) {
			return true
		}
	}
	return false
}

func goIsConstLikeExpr(expr ast.Expr) bool {
	switch e := expr.(type) {
	case *ast.BasicLit:
		return true
	case *ast.CompositeLit:
		return true
	case *ast.UnaryExpr:
		if e.Op == token.AND {
			_, isLit := e.X.(*ast.CompositeLit)
			return isLit
		}
		return false
	case *ast.CallExpr:
		return goIsConstLikeCall(e)
	default:
		return false
	}
}

func goIsConstLikeCall(call *ast.CallExpr) bool {
	name := goCallFuncName(call)
	if name == "make" || name == "new" {
		return false
	}
	constLike := map[string]bool{
		"errors.New": true, "fmt.Errorf": true,
		"regexp.MustCompile": true, "regexp.Compile": true,
	}
	if constLike[name] {
		return true
	}
	return true
}

func goCallFuncName(call *ast.CallExpr) string {
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

// --- Helpers ---

func goReceiverTypeName(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return goReceiverTypeName(t.X)
	case *ast.SelectorExpr:
		if x, ok := t.X.(*ast.Ident); ok {
			return x.Name + "." + t.Sel.Name
		}
	}
	return ""
}

func goExprTypeName(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		if x, ok := t.X.(*ast.Ident); ok {
			return x.Name + "." + t.Sel.Name
		}
	case *ast.StarExpr:
		return goExprTypeName(t.X)
	}
	return ""
}

func classifyAlgo(loopDepth, recursiveCalls int) string {
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

func isExported(name string) bool {
	if name == "" {
		return false
	}
	return unicode.IsUpper(rune(name[0]))
}
