// Package discovery handles code unit discovery and indexing.
package discovery

import (
	"github.com/code-certification/certify/internal/domain"
)

// Scanner discovers certifiable code units in a repository.
type Scanner interface {
	// Scan discovers all code units under the given root directory.
	Scan(root string) ([]domain.Unit, error)
}
