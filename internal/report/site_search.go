package report

import (
	"encoding/json"
	"fmt"
	"strings"
)

// SearchEntry is a compact representation of a unit for client-side search.
type SearchEntry struct {
	Name       string  `json:"n"`
	Path       string  `json:"p"`
	UnitID     string  `json:"id"`
	Grade      string  `json:"g"`
	Status     string  `json:"s"`
	Language   string  `json:"l"`
	Score      float64 `json:"sc"`
	PackageURL string  `json:"pu"`
	UnitURL    string  `json:"uu"`
}

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

		entries = append(entries, SearchEntry{
			Name:       name,
			Path:       u.Path,
			UnitID:     u.UnitID,
			Grade:      u.Grade,
			Status:     u.Status,
			Language:   u.Language,
			Score:      u.Score,
			PackageURL: "packages/" + dir + "/index.html",
			UnitURL:    "units/" + anchor + ".html",
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
