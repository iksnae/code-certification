package discovery

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// LanguageInfo describes a detected language in a repository.
type LanguageInfo struct {
	Name      string // "go", "typescript", "python", etc.
	FileCount int
	HasConfig bool   // go.mod, package.json, pyproject.toml, etc.
	Adapter   string // "go", "ts", "generic"
}

// DetectLanguages walks a directory and detects programming languages present.
func DetectLanguages(root string) []LanguageInfo {
	counts := make(map[string]int)
	configs := make(map[string]bool)

	filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() {
			name := d.Name()
			if strings.HasPrefix(name, ".") || name == "vendor" || name == "node_modules" ||
				name == "dist" || name == "build" || name == "testdata" {
				return filepath.SkipDir
			}
			return nil
		}

		name := d.Name()
		ext := filepath.Ext(name)

		// Count files by extension
		switch ext {
		case ".go":
			counts["go"]++
		case ".ts", ".tsx":
			counts["typescript"]++
		case ".js", ".jsx":
			counts["javascript"]++
		case ".py":
			counts["python"]++
		case ".rs":
			counts["rust"]++
		case ".rb":
			counts["ruby"]++
		case ".java":
			counts["java"]++
		case ".kt":
			counts["kotlin"]++
		case ".sh", ".bash":
			counts["shell"]++
		}

		// Detect config files
		switch name {
		case "go.mod":
			configs["go"] = true
		case "package.json":
			configs["typescript"] = true
			configs["javascript"] = true
		case "tsconfig.json":
			configs["typescript"] = true
		case "pyproject.toml", "setup.py", "requirements.txt":
			configs["python"] = true
		case "Cargo.toml":
			configs["rust"] = true
		case "Gemfile":
			configs["ruby"] = true
		case "pom.xml", "build.gradle":
			configs["java"] = true
		}

		return nil
	})

	// Build results
	adapterMap := map[string]string{
		"go":         "go",
		"typescript": "ts",
		"javascript": "ts",
		"python":     "generic",
		"rust":       "generic",
		"ruby":       "generic",
		"java":       "generic",
		"kotlin":     "generic",
		"shell":      "generic",
	}

	var langs []LanguageInfo
	for lang, count := range counts {
		if count == 0 {
			continue
		}
		adapter := adapterMap[lang]
		if adapter == "" {
			adapter = "generic"
		}
		langs = append(langs, LanguageInfo{
			Name:      lang,
			FileCount: count,
			HasConfig: configs[lang],
			Adapter:   adapter,
		})
	}

	// Sort by file count descending
	sort.Slice(langs, func(i, j int) bool {
		return langs[i].FileCount > langs[j].FileCount
	})

	return langs
}

// DetectedAdapters returns the unique adapter names needed for detected languages.
func DetectedAdapters(langs []LanguageInfo) []string {
	seen := make(map[string]bool)
	var adapters []string
	for _, l := range langs {
		if !seen[l.Adapter] {
			seen[l.Adapter] = true
			adapters = append(adapters, l.Adapter)
		}
	}
	return adapters
}
