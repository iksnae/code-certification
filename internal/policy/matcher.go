// Package policy handles policy matching and evaluation.
package policy

import (
	"path/filepath"
	"strings"

	"github.com/iksnae/code-certification/internal/domain"
)

// Matcher determines which policy packs apply to a given unit.
type Matcher struct {
	packs []domain.PolicyPack
}

// NewMatcher creates a new policy matcher.
func NewMatcher(packs []domain.PolicyPack) *Matcher {
	return &Matcher{packs: packs}
}

// Packs returns all loaded policy packs.
func (m *Matcher) Packs() []domain.PolicyPack {
	return m.packs
}

// Match returns the policy packs that apply to the given unit.
func (m *Matcher) Match(unit domain.Unit) []domain.PolicyPack {
	var matched []domain.PolicyPack

	for _, pack := range m.packs {
		if m.matchesPack(pack, unit) {
			matched = append(matched, pack)
		}
	}

	return matched
}

func (m *Matcher) matchesPack(pack domain.PolicyPack, unit domain.Unit) bool {
	// Language filter
	if pack.Language != "" && pack.Language != unit.ID.Language() {
		return false
	}

	// Path pattern filter (if patterns exist, at least one must match)
	if len(pack.PathPatterns) > 0 {
		path := unit.ID.Path()
		matched := false
		for _, pattern := range pack.PathPatterns {
			// Handle simple glob patterns
			if matchPath(pattern, path) {
				matched = true
				break
			}
		}
		if !matched {
			return false
		}
	}

	return true
}

// matchPath matches a path against a glob pattern.
func matchPath(pattern, path string) bool {
	// Strip ** prefix for simple directory matching
	p := strings.TrimPrefix(pattern, "**/")

	// Try matching the base directory
	parts := strings.SplitN(p, "/", 2)
	if len(parts) == 2 {
		dir := parts[0]
		rest := parts[1]
		pathParts := strings.Split(path, "/")
		for i, part := range pathParts {
			if globMatch(dir, part) {
				remaining := strings.Join(pathParts[i+1:], "/")
				if rest == "*" || rest == "**" {
					return true
				}
				if globMatch(rest, remaining) {
					return true
				}
			}
		}
	}

	return globMatch(pattern, path)
}

// globMatch wraps filepath.Match, treating pattern errors as non-match.
func globMatch(pattern, name string) bool {
	ok, err := filepath.Match(pattern, name)
	return err == nil && ok
}
