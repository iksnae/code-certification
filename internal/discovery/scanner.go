// Package discovery handles code unit discovery and indexing.
package discovery

import (
	"github.com/code-certification/certify/internal/domain"
)

// UnitList is a convenience alias for a slice of units.
type UnitList = []domain.Unit

// Scanner discovers certifiable code units in a repository.
type Scanner interface {
	// Scan discovers all code units under the given root directory.
	Scan(root string) ([]domain.Unit, error)
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
