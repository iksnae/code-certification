// Package discovery handles code unit discovery and indexing.
package discovery

import (
	"github.com/iksnae/code-certification/internal/domain"
)

// UnitList is a convenience alias for a slice of units.
type UnitList = []domain.Unit

// Scanner discovers certifiable code units in a repository.
type Scanner interface {
	// Scan discovers all code units under the given root directory.
	Scan(root string) ([]domain.Unit, error)
}

// Scanners returns a registry mapping adapter names to Scanner implementations.
// Use this for polymorphic dispatch instead of switch statements on adapter names.
// GenericScanner is excluded because it requires config params and always runs unconditionally.
//
// Languages with a registered analysis.Analyzer get tree-sitter-backed discovery
// (full AST, symbol-level units). Languages without fall back to regex or file-level.
func Scanners() map[string]Scanner {
	return map[string]Scanner{
		"go": NewGoAdapter(),
		"ts": NewAnalysisAdapter("ts", []string{".ts", ".tsx"}),
		"py": NewAnalysisAdapter("py", []string{".py"}),
		"rs": NewAnalysisAdapter("rs", []string{".rs"}),
	}
}

// Merge combines multiple unit lists, deduplicating by UnitID string.
// When duplicate IDs exist, the more specific type wins (function > file).
func Merge(lists ...UnitList) []domain.Unit {
	seen := make(map[string]domain.Unit)
	for _, list := range lists {
		for _, u := range list {
			key := u.ID.String()
			if existing, ok := seen[key]; ok {
				// Keep the more specific type
				if u.Type > existing.Type {
					seen[key] = u
				}
			} else {
				seen[key] = u
			}
		}
	}

	result := make([]domain.Unit, 0, len(seen))
	for _, u := range seen {
		result = append(result, u)
	}
	return result
}

// DeduplicateFileLevel removes file-level units when symbol-level units exist
// for the same file from a language adapter.
func DeduplicateFileLevel(units []domain.Unit) []domain.Unit {
	// Track which paths have symbol-level units
	hasSymbols := make(map[string]bool)
	for _, u := range units {
		if u.ID.Symbol() != "" {
			hasSymbols[u.ID.Path()] = true
		}
	}

	// Filter out file-level units for paths that have symbol-level entries
	var result []domain.Unit
	for _, u := range units {
		if u.ID.Symbol() == "" && hasSymbols[u.ID.Path()] {
			continue // Skip file-level, we have symbols
		}
		result = append(result, u)
	}
	return result
}
