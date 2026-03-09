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

// extToLanguage maps file extensions to language names.
var extToLanguage = map[string]string{
	".go":   "go",
	".ts":   "typescript",
	".tsx":  "typescript",
	".js":   "javascript",
	".jsx":  "javascript",
	".py":   "python",
	".rs":   "rust",
	".rb":   "ruby",
	".java": "java",
	".kt":   "kotlin",
	".sh":   "shell",
	".bash": "shell",
}

// configToLanguages maps config filenames to the languages they indicate.
var configToLanguages = map[string][]string{
	"go.mod":           {"go"},
	"package.json":     {"typescript", "javascript"},
	"tsconfig.json":    {"typescript"},
	"pyproject.toml":   {"python"},
	"setup.py":         {"python"},
	"requirements.txt": {"python"},
	"Cargo.toml":       {"rust"},
	"Gemfile":          {"ruby"},
	"pom.xml":          {"java"},
	"build.gradle":     {"java"},
}

// languageToAdapter maps language names to adapter identifiers.
var languageToAdapter = map[string]string{
	"go":         "go",
	"typescript": "ts",
	"javascript": "ts",
}

// skipDirs are directories to skip during language detection.
var skipDirs = map[string]bool{
	"vendor": true, "node_modules": true, "dist": true,
	"build": true, "testdata": true,
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
			if strings.HasPrefix(name, ".") || skipDirs[name] {
				return filepath.SkipDir
			}
			return nil
		}

		name := d.Name()

		// Count files by extension
		if lang, ok := extToLanguage[filepath.Ext(name)]; ok {
			counts[lang]++
		}

		// Detect config files
		if langs, ok := configToLanguages[name]; ok {
			for _, l := range langs {
				configs[l] = true
			}
		}

		return nil
	})

	return buildLanguageList(counts, configs)
}

// buildLanguageList converts raw counts and config flags into sorted LanguageInfo.
func buildLanguageList(counts map[string]int, configs map[string]bool) []LanguageInfo {
	var langs []LanguageInfo
	for lang, count := range counts {
		if count == 0 {
			continue
		}
		adapter := languageToAdapter[lang]
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
