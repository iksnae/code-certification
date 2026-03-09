package discovery

import (
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/iksnae/code-certification/internal/domain"
)

// languageFromExt maps file extensions to language identifiers.
var languageFromExt = map[string]string{
	".go":   "go",
	".ts":   "ts",
	".tsx":  "ts",
	".js":   "js",
	".jsx":  "js",
	".py":   "py",
	".rb":   "rb",
	".rs":   "rs",
	".java": "java",
	".kt":   "kt",
	".sh":   "sh",
	".bash": "sh",
}

// GenericScanner discovers all files as file-level units.
type GenericScanner struct {
	include []string // glob patterns to include (nil = all)
	exclude []string // glob patterns to exclude
}

// NewGenericScanner creates a scanner with include/exclude patterns.
func NewGenericScanner(include, exclude []string) *GenericScanner {
	return &GenericScanner{include: include, exclude: exclude}
}

// Scan walks the directory tree and returns file-level units.
func (s *GenericScanner) Scan(root string) ([]domain.Unit, error) {
	var units []domain.Unit

	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			name := d.Name()
			// Skip hidden, vendor, node_modules
			if strings.HasPrefix(name, ".") || name == "vendor" || name == "node_modules" {
				return filepath.SkipDir
			}
			return nil
		}

		// Get relative path
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)

		// Apply include filter
		if len(s.include) > 0 && !matchAny(s.include, rel) {
			return nil
		}

		// Apply exclude filter
		if matchAny(s.exclude, rel) {
			return nil
		}

		// Skip binary/hidden files
		ext := filepath.Ext(path)
		if ext == "" || strings.HasPrefix(d.Name(), ".") {
			return nil
		}

		// Determine language
		lang := languageFromExt[ext]
		if lang == "" {
			lang = "file"
		}

		id := domain.NewUnitID(lang, rel, "")
		units = append(units, domain.NewUnit(id, domain.UnitTypeFile))
		return nil
	})
	if err != nil {
		return nil, err
	}

	// Sort for deterministic output
	sort.Slice(units, func(i, j int) bool {
		return units[i].ID.String() < units[j].ID.String()
	})

	return units, nil
}

// matchAny checks if the path matches any of the glob patterns.
func matchAny(patterns []string, path string) bool {
	for _, p := range patterns {
		if matched, _ := filepath.Match(p, filepath.Base(path)); matched {
			return true
		}
		// Also try matching against full path for ** patterns
		if strings.Contains(p, "/") || strings.Contains(p, "**") {
			// Simple prefix match for directory patterns like "internal/**"
			prefix := strings.TrimSuffix(p, "/**")
			prefix = strings.TrimSuffix(prefix, "/*")
			if strings.HasPrefix(path, prefix+"/") || strings.HasPrefix(path, prefix) {
				return true
			}
		}
	}
	return false
}
