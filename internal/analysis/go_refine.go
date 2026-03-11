package analysis

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/packages"
)

// UnusedParams returns the count of parameters that are never referenced
// in the function body. Does not count receiver or blank identifier (_).
func (a *DeepGoAnalyzer) UnusedParams(pkg, funcName string) int {
	a.compute()

	p, fn := a.findFunc(pkg, funcName)
	if fn == nil || p == nil {
		return 0
	}

	// Get the types info for reference checking
	info := p.TypesInfo
	if info == nil {
		return 0
	}

	sig := fn.Type
	if sig == nil || sig.Params == nil {
		return 0
	}

	unused := 0
	for _, field := range sig.Params.List {
		for _, name := range field.Names {
			if name.Name == "_" {
				continue
			}
			obj := info.ObjectOf(name)
			if obj == nil {
				continue
			}
			// Check if this param is used anywhere in the function body
			if fn.Body == nil {
				continue
			}
			if !isIdentUsedInBlock(fn.Body, name.Name, info, obj) {
				unused++
			}
		}
	}
	return unused
}

// isIdentUsedInBlock checks if an identifier (by object identity) is referenced
// within a block statement.
func isIdentUsedInBlock(block *ast.BlockStmt, name string, info *types.Info, obj types.Object) bool {
	found := false
	ast.Inspect(block, func(n ast.Node) bool {
		if found {
			return false
		}
		ident, ok := n.(*ast.Ident)
		if !ok {
			return true
		}
		if ident.Name != name {
			return true
		}
		// Check object identity — same variable, not just same name
		identObj := info.ObjectOf(ident)
		if identObj == obj {
			found = true
			return false
		}
		return true
	})
	return found
}

// InterfaceSize returns the number of methods in the interface that the
// named type's method belongs to (0 if not a method or not implementing an interface).
func (a *DeepGoAnalyzer) InterfaceSize(pkg, funcName string) int {
	a.compute()

	// For this to be meaningful, funcName should be a method.
	// Look up the receiver type and check which interfaces it satisfies.
	p := a.findPkg(pkg)
	if p == nil || p.Types == nil {
		return 0
	}

	// Find the function object
	scope := p.Types.Scope()
	obj := scope.Lookup(funcName)
	if obj == nil {
		return 0
	}

	fn, ok := obj.(*types.Func)
	if !ok {
		return 0
	}

	sig := fn.Type().(*types.Signature)
	recv := sig.Recv()
	if recv == nil {
		return 0 // not a method
	}

	// Check all interfaces in the package for satisfaction
	maxSize := 0
	for _, p2 := range a.pkgs {
		scope2 := p2.Types.Scope()
		for _, name := range scope2.Names() {
			obj2 := scope2.Lookup(name)
			if obj2 == nil {
				continue
			}
			named, ok := obj2.Type().(*types.Named)
			if !ok {
				continue
			}
			iface, ok := named.Underlying().(*types.Interface)
			if !ok {
				continue
			}
			if types.Implements(recv.Type(), iface) || types.Implements(types.NewPointer(recv.Type()), iface) {
				if iface.NumMethods() > maxSize {
					maxSize = iface.NumMethods()
				}
			}
		}
	}
	return maxSize
}

// TypeAwareErrorWrapping returns the count of error return statements that
// don't wrap the error with context. Uses type information to verify
// the value is actually of type error, reducing false positives.
func (a *DeepGoAnalyzer) TypeAwareErrorWrapping(pkg, funcName string) int {
	a.compute()

	p, fn := a.findFunc(pkg, funcName)
	if fn == nil || p == nil {
		return 0
	}

	info := p.TypesInfo
	if info == nil {
		return 0
	}

	// Check if function returns error
	sig := fn.Type
	if sig == nil || sig.Results == nil {
		return 0
	}

	returnsError := false
	for _, field := range sig.Results.List {
		if info.TypeOf(field.Type) != nil {
			t := info.TypeOf(field.Type)
			if t != nil && t.String() == "error" {
				returnsError = true
				break
			}
		}
	}

	if !returnsError {
		return 0
	}

	// Count return statements that return an error without wrapping
	unwrapped := 0
	ast.Inspect(fn.Body, func(n ast.Node) bool {
		ret, ok := n.(*ast.ReturnStmt)
		if !ok {
			return true
		}
		for _, result := range ret.Results {
			t := info.TypeOf(result)
			if t == nil {
				continue
			}
			// Check if it's returning an error type
			if t.String() == "error" || types.Implements(t, errorInterface()) {
				// Check if it's a bare identifier (not wrapped)
				if ident, ok := result.(*ast.Ident); ok {
					if ident.Name != "nil" {
						unwrapped++
					}
				}
			}
		}
		return true
	})
	return unwrapped
}

// errorInterface returns the error interface type.
func errorInterface() *types.Interface {
	return types.Universe.Lookup("error").Type().Underlying().(*types.Interface)
}

// findFunc finds the AST FuncDecl and its containing package.
func (a *DeepGoAnalyzer) findFunc(pkg, funcName string) (*packages.Package, *ast.FuncDecl) {
	p := a.findPkg(pkg)
	if p == nil {
		return nil, nil
	}
	for _, file := range p.Syntax {
		for _, decl := range file.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok {
				continue
			}
			if fn.Name.Name == funcName {
				return p, fn
			}
		}
	}
	return p, nil
}

// findPkg finds a package by path.
func (a *DeepGoAnalyzer) findPkg(pkg string) *packages.Package {
	for _, p := range a.pkgs {
		if p.PkgPath == pkg {
			return p
		}
	}
	return nil
}
