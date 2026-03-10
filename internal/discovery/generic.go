package discovery

import (
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/iksnae/code-certification/internal/domain"
)

// certifiableExts maps file extensions to language identifiers.
// Only extensions in this map are considered certifiable code.
var certifiableExts = map[string]string{
	// Go
	".go": "go",
	// TypeScript / JavaScript
	".ts": "ts", ".tsx": "ts",
	".js": "js", ".jsx": "js",
	".mjs": "js", ".cjs": "js",
	// Python
	".py": "py",
	// Ruby
	".rb": "rb",
	// Rust
	".rs": "rs",
	// Java / Kotlin
	".java": "java", ".kt": "kt",
	// Shell
	".sh": "sh", ".bash": "sh",
	// C / C++
	".c": "c", ".h": "c",
	".cpp": "cpp", ".cc": "cpp", ".cxx": "cpp", ".hpp": "cpp",
	// C#
	".cs": "cs",
	// Swift
	".swift": "swift",
	// PHP
	".php": "php",
	// Elixir / Erlang
	".ex": "elixir", ".exs": "elixir", ".erl": "erlang",
	// Scala
	".scala": "scala",
	// Lua
	".lua": "lua",
	// R
	".r": "r", ".R": "r",
	// Zig
	".zig": "zig",
	// SQL (stored procedures, migrations)
	".sql": "sql",
	// Protobuf / gRPC
	".proto": "proto",
}

// nonCertifiableNames are filenames (exact match) to always skip.
var nonCertifiableNames = map[string]bool{
	"package-lock.json": true,
	"yarn.lock":         true,
	"pnpm-lock.yaml":    true,
	"go.sum":            true,
	"go.mod":            true,
	"Cargo.lock":        true,
	"Gemfile.lock":      true,
	"composer.lock":     true,
	"poetry.lock":       true,
}

// scanSkipDirs are directory names to always skip during scanning
// (in addition to hidden dirs, vendor, node_modules already in skipDirs).
var scanSkipDirs = map[string]bool{
	"out":         true,
	"coverage":    true,
	"__pycache__": true,
	"target":      true, // Rust/Java build output
	"bin":         true,
	"obj":         true, // .NET build output
	".next":       true,
	".nuxt":       true,
}

// Backward-compatible alias
var languageFromExt = certifiableExts

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
			// Skip hidden, vendor, node_modules, build output dirs
			if strings.HasPrefix(name, ".") || skipDirs[name] || scanSkipDirs[name] {
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

		name := d.Name()

		// Skip hidden files
		if strings.HasPrefix(name, ".") {
			return nil
		}

		// Skip known non-certifiable filenames (lock files, checksums)
		if nonCertifiableNames[name] {
			return nil
		}

		// Skip minified files
		if strings.HasSuffix(name, ".min.js") || strings.HasSuffix(name, ".min.css") {
			return nil
		}

		// Only include files with certifiable extensions
		ext := filepath.Ext(path)
		lang := certifiableExts[ext]
		if lang == "" {
			return nil // not a certifiable file type
		}

		// Apply include filter
		if len(s.include) > 0 && !matchAny(s.include, rel) {
			return nil
		}

		// Apply exclude filter
		if matchAny(s.exclude, rel) {
			return nil
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
		if matched, err := filepath.Match(p, filepath.Base(path)); err == nil && matched {
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
