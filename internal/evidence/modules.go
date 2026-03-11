package evidence

import (
	"os"
	"path/filepath"
	"strings"
)

// ModuleRoot represents a discovered language module root directory.
type ModuleRoot struct {
	Path     string // Absolute path to the module directory
	RelPath  string // Relative path from repo root (empty = repo root)
	Language string // "go", "ts", "py", "rs", etc.
	Marker   string // The file that identified this as a module root (go.mod, package.json, etc.)
}

// moduleMarkers maps marker filenames to the language they indicate.
var moduleMarkers = map[string]string{
	"go.mod":           "go",
	"go.work":          "go",
	"package.json":     "ts",
	"tsconfig.json":    "ts",
	"pyproject.toml":   "py",
	"setup.py":         "py",
	"setup.cfg":        "py",
	"requirements.txt": "py",
	"Cargo.toml":       "rs",
	"Gemfile":          "rb",
	"pom.xml":          "java",
	"build.gradle":     "java",
	"build.gradle.kts": "java",
	"mix.exs":          "elixir",
	"Package.swift":    "swift",
}

// DiscoverModuleRoots walks from the repository root and finds all directories
// that contain language module markers (go.mod, package.json, Cargo.toml, etc.).
//
// This handles monorepos and nested modules: a Go project might have go.mod in
// code/, services/api/, or multiple locations. Each is returned as a ModuleRoot.
func DiscoverModuleRoots(repoRoot string) []ModuleRoot {
	var roots []ModuleRoot
	seen := make(map[string]bool)

	filepath.WalkDir(repoRoot, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() {
			name := d.Name()
			if strings.HasPrefix(name, ".") || name == "vendor" || name == "node_modules" ||
				name == "dist" || name == "build" || name == "target" || name == "__pycache__" {
				return filepath.SkipDir
			}
			return nil
		}

		lang, isMarker := moduleMarkers[d.Name()]
		if !isMarker {
			return nil
		}

		dir := filepath.Dir(path)
		// Deduplicate: if we already found a marker for this language in this dir, skip
		key := dir + ":" + lang
		if seen[key] {
			return nil
		}
		seen[key] = true

		rel, err := filepath.Rel(repoRoot, dir)
		if err != nil {
			rel = dir
		}
		if rel == "." {
			rel = ""
		}
		rel = filepath.ToSlash(rel)

		roots = append(roots, ModuleRoot{
			Path:     dir,
			RelPath:  rel,
			Language: lang,
			Marker:   d.Name(),
		})

		return nil
	})

	return roots
}

// GoModuleRoots returns only Go module roots from the discovered modules.
func GoModuleRoots(roots []ModuleRoot) []ModuleRoot {
	var goRoots []ModuleRoot
	for _, r := range roots {
		if r.Language == "go" {
			goRoots = append(goRoots, r)
		}
	}
	return goRoots
}

// UnitBelongsToModule checks if a unit path (relative to repo root)
// falls within a module's directory.
func UnitBelongsToModule(unitPath string, module ModuleRoot) bool {
	if module.RelPath == "" {
		return true // root module owns everything
	}
	return strings.HasPrefix(unitPath, module.RelPath+"/") || unitPath == module.RelPath
}

// FindModuleForUnit returns the most specific module root that contains
// the given unit path. Returns nil if no module contains the unit.
func FindModuleForUnit(unitPath string, modules []ModuleRoot) *ModuleRoot {
	var best *ModuleRoot
	bestLen := -1

	for i, m := range modules {
		if UnitBelongsToModule(unitPath, m) {
			// Prefer the most specific (longest prefix) module
			if len(m.RelPath) > bestLen {
				best = &modules[i]
				bestLen = len(m.RelPath)
			}
		}
	}

	return best
}
