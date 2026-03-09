package evidence

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

// ComputeGoComplexity parses Go source and returns cyclomatic complexity per function.
// Key: function name (or Receiver.Method for methods). Value: complexity.
func ComputeGoComplexity(src string) map[string]int {
	result := make(map[string]int)
	if strings.TrimSpace(src) == "" {
		return result
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		return result
	}

	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Body == nil {
			continue
		}

		name := funcName(fn)
		complexity := 1 // Base complexity

		ast.Inspect(fn.Body, func(n ast.Node) bool {
			switch node := n.(type) {
			case *ast.IfStmt:
				complexity++
			case *ast.ForStmt:
				complexity++
			case *ast.RangeStmt:
				complexity++
			case *ast.CaseClause:
				if node.List != nil { // Skip default case
					complexity++
				}
			case *ast.CommClause:
				if node.Comm != nil { // Skip default case
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

		result[name] = complexity
	}

	return result
}

func funcName(fn *ast.FuncDecl) string {
	if fn.Recv != nil && len(fn.Recv.List) > 0 {
		recv := fn.Recv.List[0].Type
		// Unwrap pointer
		if star, ok := recv.(*ast.StarExpr); ok {
			recv = star.X
		}
		if ident, ok := recv.(*ast.Ident); ok {
			return ident.Name + "." + fn.Name.Name
		}
	}
	return fn.Name.Name
}

// ComputeSymbolMetrics computes metrics for a specific symbol within Go source.
// If the symbol isn't found, falls back to file-level metrics.
func ComputeSymbolMetrics(src, symbol string) CodeMetrics {
	if strings.TrimSpace(src) == "" {
		return CodeMetrics{}
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", src, parser.ParseComments)
	if err != nil {
		return ComputeMetrics(src)
	}

	// Find the function/method matching symbol
	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Body == nil {
			continue
		}
		name := funcName(fn)
		if name != symbol && fn.Name.Name != symbol {
			continue
		}

		// Found it — extract lines
		startLine := fset.Position(fn.Pos()).Line
		endLine := fset.Position(fn.End()).Line

		// Include preceding comment if attached
		if fn.Doc != nil {
			docStart := fset.Position(fn.Doc.Pos()).Line
			if docStart < startLine {
				startLine = docStart
			}
		}

		// Extract just this function's source lines
		lines := strings.Split(src, "\n")
		if startLine < 1 {
			startLine = 1
		}
		if endLine > len(lines) {
			endLine = len(lines)
		}
		funcSrc := strings.Join(lines[startLine-1:endLine], "\n")

		// Compute metrics on just this symbol's source
		metrics := ComputeMetrics(funcSrc)

		// Get complexity for this specific function
		complexities := ComputeGoComplexity(src)
		if c, ok := complexities[name]; ok {
			metrics.Complexity = c
		} else if c, ok := complexities[fn.Name.Name]; ok {
			metrics.Complexity = c
		} else {
			metrics.Complexity = 1
		}

		return metrics
	}

	// Also check type declarations
	for _, decl := range f.Decls {
		gd, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec := range gd.Specs {
			ts, ok := spec.(*ast.TypeSpec)
			if !ok || ts.Name.Name != symbol {
				continue
			}
			startLine := fset.Position(gd.Pos()).Line
			endLine := fset.Position(gd.End()).Line
			lines := strings.Split(src, "\n")
			if endLine > len(lines) {
				endLine = len(lines)
			}
			funcSrc := strings.Join(lines[startLine-1:endLine], "\n")
			metrics := ComputeMetrics(funcSrc)
			metrics.Complexity = 0 // Types don't have complexity
			return metrics
		}
	}

	// Symbol not found — file-level fallback
	return ComputeMetrics(src)
}
