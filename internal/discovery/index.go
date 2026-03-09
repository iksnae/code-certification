package discovery

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/iksnae/code-certification/internal/domain"
)

// indexEntry is the JSON-serializable form of a unit.
type indexEntry struct {
	ID       string `json:"id"`
	Language string `json:"language"`
	Path     string `json:"path"`
	Symbol   string `json:"symbol,omitempty"`
	Type     string `json:"type"`
}

// Index holds a snapshot of discovered code units.
type Index struct {
	units []domain.Unit
}

// NewIndex creates an Index from a slice of units.
func NewIndex(units []domain.Unit) *Index {
	if units == nil {
		units = []domain.Unit{}
	}
	return &Index{units: units}
}

// Units returns the units in this index.
func (idx *Index) Units() []domain.Unit {
	return idx.units
}

// Save writes the index to a JSON file.
func (idx *Index) Save(path string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("creating index directory: %w", err)
	}

	entries := make([]indexEntry, len(idx.units))
	for i, u := range idx.units {
		entries[i] = indexEntry{
			ID:       u.ID.String(),
			Language: u.ID.Language(),
			Path:     u.ID.Path(),
			Symbol:   u.ID.Symbol(),
			Type:     u.Type.String(),
		}
	}

	data, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		return fmt.Errorf("marshaling index: %w", err)
	}
	return os.WriteFile(path, data, 0o644)
}

// LoadIndex reads an index from a JSON file.
func LoadIndex(path string) (*Index, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading index: %w", err)
	}

	var entries []indexEntry
	if err := json.Unmarshal(data, &entries); err != nil {
		return nil, fmt.Errorf("parsing index: %w", err)
	}

	units := make([]domain.Unit, len(entries))
	for i, e := range entries {
		id := domain.NewUnitID(e.Language, e.Path, e.Symbol)
		ut, _ := domain.ParseUnitType(e.Type)
		units[i] = domain.NewUnit(id, ut)
	}

	return &Index{units: units}, nil
}

// DiffResult represents changes between two index snapshots.
type DiffResult struct {
	Added     []domain.Unit
	Removed   []domain.Unit
	Unchanged []domain.Unit
}

// Diff computes the difference between an old and new index.
func Diff(old, new_ *Index) DiffResult {
	oldSet := make(map[string]domain.Unit)
	for _, u := range old.units {
		oldSet[u.ID.String()] = u
	}

	newSet := make(map[string]domain.Unit)
	for _, u := range new_.units {
		newSet[u.ID.String()] = u
	}

	var result DiffResult

	// Added: in new but not old
	for k, u := range newSet {
		if _, found := oldSet[k]; !found {
			result.Added = append(result.Added, u)
		} else {
			result.Unchanged = append(result.Unchanged, u)
		}
	}

	// Removed: in old but not new
	for k, u := range oldSet {
		if _, found := newSet[k]; !found {
			result.Removed = append(result.Removed, u)
		}
	}

	return result
}
