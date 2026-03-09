// Package evidence handles evidence collection from various sources.
package evidence

import (
	"github.com/code-certification/certify/internal/domain"
)

// Collector gathers evidence for a code unit.
type Collector interface {
	// Collect gathers evidence for the given unit in the given repo root.
	Collect(root string, unit domain.Unit) (domain.Evidence, error)

	// Kind returns the kind of evidence this collector produces.
	Kind() domain.EvidenceKind
}
