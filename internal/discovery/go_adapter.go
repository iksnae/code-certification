package discovery

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/iksnae/code-certification/internal/domain"
)

// GoAdapter discovers Go code units using go/ast.
type GoAdapter struct{}

// NewGoAdapter creates a new Go discovery adapter.
func NewGoAdapter() *GoAdapter {
	return &GoAdapter{}
}

// Scan discovers Go functions, methods, and types in .go files.
func (a *GoAdapter) Scan(root string) ([]domain.Unit, error) {
	var units []domain.Unit

	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			name := d.Name()
			if strings.HasPrefix(name, ".") || name == "vendor" || name == "node_modules" || name == "testdata" {
				return filepath.SkipDir
			}
			return nil
		}
		if filepath.Ext(path) != ".go" || strings.HasSuffix(path, "_test.go") {
			return nil
		}

		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)

		fileUnits, err := a.parseFile(path, rel)
		if err != nil {
			// Graceful fallback: return file-level unit
			id := domain.NewUnitID("go", rel, "")
			units = append(units, domain.NewUnit(id, domain.UnitTypeFile))
			return nil
		}
		units = append(units, fileUnits...)
		return nil
	})
	if err != nil {
		return nil, err
	}

	sort.Slice(units, func(i, j int) bool {
		return units[i].ID.String() < units[j].ID.String()
	})

	return units, nil
}

func (a *GoAdapter) parseFile(path, rel string) ([]domain.Unit, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	var units []domain.Unit

	for _, decl := range file.Decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			name := d.Name.Name
			if d.Recv != nil && len(d.Recv.List) > 0 {
				// Method
				id := domain.NewUnitID("go", rel, name)
				units = append(units, domain.NewUnit(id, domain.UnitTypeMethod))
			} else {
				// Function
				id := domain.NewUnitID("go", rel, name)
				units = append(units, domain.NewUnit(id, domain.UnitTypeFunction))
			}
		case *ast.GenDecl:
			if d.Tok == token.TYPE {
				for _, spec := range d.Specs {
					ts, ok := spec.(*ast.TypeSpec)
					if ok {
						id := domain.NewUnitID("go", rel, ts.Name.Name)
						units = append(units, domain.NewUnit(id, domain.UnitTypeClass))
					}
				}
			}
		}
	}

	return units, nil
}
