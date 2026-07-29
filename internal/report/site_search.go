package report

import (
	"encoding/json"
	"fmt"
	"strings"
)

// SearchEntry is a compact representation of a unit for client-side search.
//
// search-index.js is a published artifact: it carries a per-unit Grade and Score
// to every reader of the site. Score is therefore a pointer — an unassessed unit
// has no score to publish, and omitting the key says so, where a float64 could
// only ever say "zero". The field is minified like the rest, which is why no
// rendered-cell assertion elsewhere could observe the fabricated value.
type SearchEntry struct {
	Name        string   `json:"n"`
	Path        string   `json:"p"`
	UnitID      string   `json:"id"`
	Grade       string   `json:"g"`
	Status      string   `json:"s"`
	Language    string   `json:"l"`
	Score       *float64 `json:"sc,omitempty"`
	Unsupported bool     `json:"u,omitempty"`
	PackageURL  string   `json:"pu"`
	UnitURL     string   `json:"uu"`
}

// ScoreKnown reports whether this entry's Score is a measurement. See
// Card.ScoreKnown.
func (e SearchEntry) ScoreKnown() bool { return !e.Unsupported }

// BuildSearchIndex creates a compact search index from a FullReport.
func BuildSearchIndex(r FullReport) []SearchEntry {
	entries := make([]SearchEntry, 0, len(r.Units))
	for _, u := range r.Units {
		name := u.Symbol
		if name == "" {
			name = shortFile(u.Path)
		}
		anchor := unitAnchor(u)
		dir := dirOf(u.Path)

		// An unassessed unit publishes no score. Its stored value is the
		// pipeline's placeholder, and shipping it as "sc":0 states a measured
		// total failure to every consumer of the index.
		var score *float64
		if u.ScoreKnown() {
			s := u.Score
			score = &s
		}

		entries = append(entries, SearchEntry{
			Name:        name,
			Path:        u.Path,
			UnitID:      u.UnitID,
			Grade:       u.Grade,
			Status:      u.Status,
			Language:    u.Language,
			Score:       score,
			Unsupported: u.Unsupported,
			PackageURL:  "packages/" + dir + "/index.html",
			UnitURL:     "units/" + anchor + ".html",
		})
	}
	return entries
}

// FormatSearchIndexJS returns a JavaScript constant declaration containing
// the search index as a JSON array, suitable for embedding in an HTML page.
func FormatSearchIndexJS(entries []SearchEntry) string {
	data, err := json.Marshal(entries)
	if err != nil {
		return fmt.Sprintf("const SEARCH_INDEX = []; // error: %s\n", err)
	}
	var b strings.Builder
	b.WriteString("const SEARCH_INDEX = ")
	b.Write(data)
	b.WriteString(";\n")
	return b.String()
}
