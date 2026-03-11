package discovery

import (
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/iksnae/code-certification/internal/analysis"
	"github.com/iksnae/code-certification/internal/domain"
)

// AnalysisAdapter discovers code units using the analysis.Analyzer interface.
// This replaces language-specific regex adapters with tree-sitter-backed discovery.
type AnalysisAdapter struct {
	lang       string
	extensions []string
	skipDirs   map[string]bool
	analyzer   analysis.Analyzer
}

// NewAnalysisAdapter creates a discovery adapter backed by an Analyzer.
func NewAnalysisAdapter(lang string, extensions []string) *AnalysisAdapter {
	return &AnalysisAdapter{
		lang:       lang,
		extensions: extensions,
		skipDirs: map[string]bool{
			"vendor": true, "node_modules": true, "dist": true,
			"build": true, ".git": true, "testdata": true,
			"__pycache__": true, "target": true, ".next": true,
		},
		analyzer: analysis.ForLanguage(lang),
	}
}

// Scan discovers all code units for this language.
func (a *AnalysisAdapter) Scan(root string) ([]domain.Unit, error) {
	if a.analyzer == nil {
		return nil, nil
	}

	var units []domain.Unit

	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			name := d.Name()
			if strings.HasPrefix(name, ".") || a.skipDirs[name] {
				return filepath.SkipDir
			}
			return nil
		}

		ext := filepath.Ext(path)
		if !a.matchesExt(ext) {
			return nil
		}

		// Skip test files
		if a.isTestFile(d.Name()) {
			return nil
		}

		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)

		src, err := os.ReadFile(path)
		if err != nil {
			return nil // skip unreadable files
		}

		symbols, err := a.analyzer.Discover(path, src)
		if err != nil {
			// Fallback to file-level unit
			id := domain.NewUnitID(a.lang, rel, "")
			units = append(units, domain.NewUnit(id, domain.UnitTypeFile))
			return nil
		}

		for _, sym := range symbols {
			unitType := a.symbolToUnitType(sym.Kind)
			id := domain.NewUnitID(a.lang, rel, sym.Name)
			units = append(units, domain.NewUnit(id, unitType))
		}

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

func (a *AnalysisAdapter) matchesExt(ext string) bool {
	for _, e := range a.extensions {
		if ext == e {
			return true
		}
	}
	return false
}

func (a *AnalysisAdapter) isTestFile(name string) bool {
	switch a.lang {
	case "ts", "js":
		return strings.HasSuffix(name, ".test.ts") || strings.HasSuffix(name, ".spec.ts") ||
			strings.HasSuffix(name, ".test.tsx") || strings.HasSuffix(name, ".spec.tsx") ||
			strings.HasSuffix(name, ".test.js") || strings.HasSuffix(name, ".spec.js") ||
			strings.HasSuffix(name, ".d.ts")
	case "py":
		return strings.HasPrefix(name, "test_") || strings.HasSuffix(name, "_test.py")
	case "rs":
		// Rust tests are inline, but test modules in separate files
		return false
	case "go":
		return strings.HasSuffix(name, "_test.go")
	}
	return false
}

func (a *AnalysisAdapter) symbolToUnitType(kind analysis.SymbolKind) domain.UnitType {
	switch kind {
	case analysis.SymbolFunction:
		return domain.UnitTypeFunction
	case analysis.SymbolMethod:
		return domain.UnitTypeMethod
	case analysis.SymbolClass, analysis.SymbolInterface:
		return domain.UnitTypeClass
	case analysis.SymbolConstant:
		return domain.UnitTypeFunction // Treat constants as function-level
	default:
		return domain.UnitTypeFile
	}
}
