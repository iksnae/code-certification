package analysis

import (
	"go/types"
	"strings"

	"golang.org/x/tools/go/packages"
)

// AbstractionResult holds the param abstraction analysis for a function.
type AbstractionResult struct {
	TotalParams     int // total number of params
	ConcreteDeps    int // params accepting concrete struct types from other packages
	InterfaceParams int // params accepting interface types
}

// DepDepth returns the maximum transitive local import depth for a package.
// Only counts imports within the project (same module prefix), not stdlib.
func (a *DeepGoAnalyzer) DepDepth(pkgPath string) int {
	a.compute()
	pkgMap := a.buildPkgMap()
	visited := make(map[string]bool)
	return a.depDepthRec(pkgPath, pkgMap, visited)
}

func (a *DeepGoAnalyzer) depDepthRec(pkgPath string, pkgMap map[string]*packages.Package, visited map[string]bool) int {
	if visited[pkgPath] {
		return 0 // cycle protection
	}
	visited[pkgPath] = true

	pkg, ok := pkgMap[pkgPath]
	if !ok {
		return 0
	}

	maxDepth := 0
	for impPath := range pkg.Imports {
		if !a.isLocalPkg(impPath) {
			continue
		}
		d := 1 + a.depDepthRec(impPath, pkgMap, visited)
		if d > maxDepth {
			maxDepth = d
		}
	}
	return maxDepth
}

// Instability computes Robert C. Martin's instability metric: Ce / (Ca + Ce)
// Ce = efferent coupling (local imports this package makes)
// Ca = afferent coupling (local packages that import this one)
// Returns 0.0 (fully stable) to 1.0 (fully unstable).
func (a *DeepGoAnalyzer) Instability(pkgPath string) float64 {
	a.compute()
	pkgMap := a.buildPkgMap()

	pkg, ok := pkgMap[pkgPath]
	if !ok {
		return 0
	}

	// Ce: count local imports
	ce := 0
	for impPath := range pkg.Imports {
		if a.isLocalPkg(impPath) {
			ce++
		}
	}

	// Ca: count local packages that import this one
	ca := 0
	for _, p := range a.pkgs {
		if p.PkgPath == pkgPath {
			continue
		}
		for impPath := range p.Imports {
			if impPath == pkgPath {
				ca++
				break
			}
		}
	}

	total := ca + ce
	if total == 0 {
		return 0
	}
	return float64(ce) / float64(total)
}

// ParamAbstraction checks whether function params use interfaces or concrete types
// from other packages.
func (a *DeepGoAnalyzer) ParamAbstraction(pkg, funcName string) AbstractionResult {
	a.compute()

	result := AbstractionResult{}

	// Search through types info
	for _, p := range a.pkgs {
		if p.PkgPath != pkg {
			continue
		}
		scope := p.Types.Scope()
		obj := scope.Lookup(funcName)
		if obj == nil {
			continue
		}
		fn, ok := obj.(*types.Func)
		if !ok {
			continue
		}
		sig, ok := fn.Type().(*types.Signature)
		if !ok {
			continue
		}
		return analyzeSignature(sig, pkg)
	}

	return result
}

// analyzeSignature examines function parameters for interface vs concrete types.
func analyzeSignature(sig *types.Signature, selfPkg string) AbstractionResult {
	params := sig.Params()
	result := AbstractionResult{TotalParams: params.Len()}

	for i := 0; i < params.Len(); i++ {
		paramType := params.At(i).Type()
		underlying := paramType.Underlying()

		switch underlying.(type) {
		case *types.Interface:
			result.InterfaceParams++
		case *types.Struct:
			// Check if it's from an external package
			if named, ok := paramType.(*types.Named); ok {
				obj := named.Obj()
				if obj != nil && obj.Pkg() != nil && obj.Pkg().Path() != selfPkg {
					result.ConcreteDeps++
				}
			}
		case *types.Pointer:
			// Check pointer-to-struct from external package
			ptr := underlying.(*types.Pointer)
			if st, ok := ptr.Elem().Underlying().(*types.Struct); ok {
				_ = st
				if named, ok := ptr.Elem().(*types.Named); ok {
					obj := named.Obj()
					if obj != nil && obj.Pkg() != nil && obj.Pkg().Path() != selfPkg {
						result.ConcreteDeps++
					}
				}
			}
		}
		// Primitives (string, int, etc.) and slices/maps of primitives are fine
	}

	return result
}

// CouplingScore computes a normalized coupling score for a function.
// coupling = fan_in * fan_out, normalized to [0, 1] range.
func (a *DeepGoAnalyzer) CouplingScore(pkg, funcName string) float64 {
	fanIn := a.FanIn(pkg, funcName)
	fanOut := a.FanOut(pkg, funcName)
	coupling := float64(fanIn * fanOut)

	// Normalize: 0-25 → good, 25-100 → moderate, 100+ → high
	if coupling <= 0 {
		return 0
	}
	if coupling >= 100 {
		return 1.0
	}
	return coupling / 100.0
}

// buildPkgMap creates a lookup from package path to package.
func (a *DeepGoAnalyzer) buildPkgMap() map[string]*packages.Package {
	m := make(map[string]*packages.Package, len(a.pkgs))
	for _, p := range a.pkgs {
		m[p.PkgPath] = p
	}
	return m
}

// isLocalPkg checks if a package path belongs to the project (same module prefix).
func (a *DeepGoAnalyzer) isLocalPkg(path string) bool {
	if len(a.pkgs) == 0 {
		return false
	}
	// Use the first package's module path as prefix
	for _, p := range a.pkgs {
		if p.Module != nil {
			return strings.HasPrefix(path, p.Module.Path)
		}
	}
	// Fallback: check if it's in our package list
	for _, p := range a.pkgs {
		if p.PkgPath == path {
			return true
		}
	}
	return false
}
