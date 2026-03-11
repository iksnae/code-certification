package analysis

import (
	"go/token"
	"go/types"
	"sync"

	"golang.org/x/tools/go/callgraph"
	"golang.org/x/tools/go/callgraph/vta"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

// FuncKey uniquely identifies a function by package path and name.
type FuncKey struct {
	Pkg  string // full package path
	Name string // function or method name (e.g., "Hello" or "(*T).Method")
}

// DeepResult holds the type-aware analysis results for a single function.
type DeepResult struct {
	FanIn      int  // number of call sites that invoke this function
	FanOut     int  // number of distinct functions called by this function
	IsDeadCode bool // exported with zero external references
}

// UnusedSymbol represents an exported symbol with no external references.
type UnusedSymbol struct {
	Pkg      string // package path
	Name     string // symbol name
	FilePath string // file where it's defined
	Line     int    // line number
}

// DeepGoAnalyzer provides type-aware cross-file analysis for Go using
// go/packages + SSA + VTA call graph analysis.
type DeepGoAnalyzer struct {
	pkgs  []*packages.Package
	fset  *token.FileSet
	graph *callgraph.Graph
	prog  *ssa.Program

	// Lazily computed results
	once    sync.Once
	results map[FuncKey]DeepResult
	unused  []UnusedSymbol
}

// LoadGoProject loads all Go packages under root and builds the call graph.
// patterns is typically []string{"./..."} to load all packages.
func LoadGoProject(root string, patterns ...string) (*DeepGoAnalyzer, error) {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedSyntax |
			packages.NeedTypes | packages.NeedTypesInfo | packages.NeedDeps |
			packages.NeedImports | packages.NeedCompiledGoFiles,
		Dir: root,
	}

	pkgs, err := packages.Load(cfg, patterns...)
	if err != nil {
		return nil, err
	}

	// Check for load errors (but don't fail — partial analysis is still useful)
	for _, pkg := range pkgs {
		if len(pkg.Errors) > 0 {
			// Log but continue — partial results are better than none
			continue
		}
	}

	a := &DeepGoAnalyzer{
		pkgs: pkgs,
		fset: cfg.Fset,
	}

	return a, nil
}

// buildCallGraph constructs the SSA program and VTA call graph.
func (a *DeepGoAnalyzer) buildCallGraph() {
	// Build SSA
	prog, ssaPkgs := ssautil.AllPackages(a.pkgs, ssa.InstantiateGenerics)
	prog.Build()
	a.prog = prog

	// Collect all functions for VTA
	funcs := make(map[*ssa.Function]bool)
	for _, pkg := range ssaPkgs {
		if pkg == nil {
			continue
		}
		for _, member := range pkg.Members {
			if fn, ok := member.(*ssa.Function); ok {
				funcs[fn] = true
				// Include anonymous functions
				for _, anon := range fn.AnonFuncs {
					funcs[anon] = true
				}
			}
			// Include methods
			if typ, ok := member.(*ssa.Type); ok {
				methodSet := prog.MethodSets.MethodSet(types.NewPointer(typ.Type()))
				for i := 0; i < methodSet.Len(); i++ {
					sel := methodSet.At(i)
					fn := prog.MethodValue(sel)
					if fn != nil {
						funcs[fn] = true
					}
				}
			}
		}
	}

	a.graph = vta.CallGraph(funcs, nil)
}

// compute populates results and unused lists from the call graph.
func (a *DeepGoAnalyzer) compute() {
	a.once.Do(func() {
		a.buildCallGraph()
		a.results = make(map[FuncKey]DeepResult)
		a.computeFanMetrics()
		a.computeUnusedExports()
	})
}

// computeFanMetrics walks the call graph to determine fan-in and fan-out.
func (a *DeepGoAnalyzer) computeFanMetrics() {
	if a.graph == nil || a.graph.Nodes == nil {
		return
	}

	// Build package path set for project-local filtering
	localPkgs := make(map[string]bool)
	for _, pkg := range a.pkgs {
		localPkgs[pkg.PkgPath] = true
	}

	for fn, node := range a.graph.Nodes {
		if fn == nil || fn.Package() == nil {
			continue
		}
		pkg := fn.Package()
		if !localPkgs[pkg.Pkg.Path()] {
			continue
		}

		key := funcKeyFromSSA(fn)
		result := a.results[key]

		// Fan-in: count incoming edges from distinct call sites
		fanIn := 0
		for _, edge := range node.In {
			if edge.Caller != nil && edge.Caller.Func != nil {
				fanIn++
			}
		}
		result.FanIn = fanIn

		// Fan-out: count outgoing edges to distinct callees
		seen := make(map[FuncKey]bool)
		for _, edge := range node.Out {
			if edge.Callee != nil && edge.Callee.Func != nil {
				calleeKey := funcKeyFromSSA(edge.Callee.Func)
				seen[calleeKey] = true
			}
		}
		result.FanOut = len(seen)

		a.results[key] = result
	}
}

// computeUnusedExports finds exported functions/types with no external callers.
func (a *DeepGoAnalyzer) computeUnusedExports() {
	if a.graph == nil {
		return
	}

	localPkgs := make(map[string]bool)
	for _, pkg := range a.pkgs {
		localPkgs[pkg.PkgPath] = true
	}

	for fn, node := range a.graph.Nodes {
		if fn == nil || fn.Package() == nil {
			continue
		}
		pkg := fn.Package()
		pkgPath := pkg.Pkg.Path()
		if !localPkgs[pkgPath] {
			continue
		}

		// Skip main packages — their exports are entry points
		if pkg.Pkg.Name() == "main" {
			continue
		}

		name := fn.Name()
		if !token.IsExported(name) {
			continue
		}

		// Check if any caller is from a different package
		hasExternalCaller := false
		for _, edge := range node.In {
			if edge.Caller != nil && edge.Caller.Func != nil {
				callerPkg := edge.Caller.Func.Package()
				if callerPkg != nil && callerPkg.Pkg.Path() != pkgPath {
					hasExternalCaller = true
					break
				}
			}
		}

		if !hasExternalCaller {
			sym := UnusedSymbol{
				Pkg:  pkgPath,
				Name: name,
			}
			// Try to get file/line info
			if pos := fn.Pos(); pos.IsValid() {
				position := a.prog.Fset.Position(pos)
				sym.FilePath = position.Filename
				sym.Line = position.Line
			}
			a.unused = append(a.unused, sym)

			// Mark in results
			key := funcKeyFromSSA(fn)
			r := a.results[key]
			r.IsDeadCode = true
			a.results[key] = r
		}
	}
}

// FanIn returns the number of call sites that invoke the named function.
func (a *DeepGoAnalyzer) FanIn(pkg, funcName string) int {
	a.compute()
	key := FuncKey{Pkg: pkg, Name: funcName}
	return a.results[key].FanIn
}

// FanOut returns the number of distinct functions called by the named function.
func (a *DeepGoAnalyzer) FanOut(pkg, funcName string) int {
	a.compute()
	key := FuncKey{Pkg: pkg, Name: funcName}
	return a.results[key].FanOut
}

// UnusedExports returns exported symbols with zero external references.
func (a *DeepGoAnalyzer) UnusedExports() []UnusedSymbol {
	a.compute()
	return a.unused
}

// AllResults returns all computed deep analysis results.
func (a *DeepGoAnalyzer) AllResults() map[FuncKey]DeepResult {
	a.compute()
	return a.results
}

// Lookup returns the deep analysis result for a specific function.
func (a *DeepGoAnalyzer) Lookup(pkg, funcName string) (DeepResult, bool) {
	a.compute()
	key := FuncKey{Pkg: pkg, Name: funcName}
	r, ok := a.results[key]
	return r, ok
}

// funcKeyFromSSA extracts a FuncKey from an SSA function.
func funcKeyFromSSA(fn *ssa.Function) FuncKey {
	pkg := fn.Package()
	if pkg == nil {
		return FuncKey{Name: fn.Name()}
	}
	return FuncKey{
		Pkg:  pkg.Pkg.Path(),
		Name: fn.Name(),
	}
}
