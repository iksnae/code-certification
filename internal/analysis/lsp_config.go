package analysis

import "fmt"

// LSPAvailability describes what LSP servers are available per language.
type LSPAvailability struct {
	Language  string
	Available bool
	Command   string
	InstallHint string
}

// DetectLSPServers checks which language servers are installed.
func DetectLSPServers() []LSPAvailability {
	results := []LSPAvailability{
		{
			Language:    "ts",
			Command:     "typescript-language-server",
			InstallHint: "npm i -g typescript-language-server typescript",
		},
		{
			Language:    "py",
			Command:     "pyright-langserver",
			InstallHint: "pip install pyright",
		},
		{
			Language:    "rs",
			Command:     "rust-analyzer",
			InstallHint: "rustup component add rust-analyzer",
		},
	}

	for i := range results {
		results[i].Available = commandExists(results[i].Command)
	}

	return results
}

// FormatLSPStatus returns a human-readable status for doctor output.
func FormatLSPStatus(avail []LSPAvailability) []string {
	var lines []string
	for _, a := range avail {
		if a.Available {
			lines = append(lines, fmt.Sprintf("  ✅ %s: %s found", langDisplayName(a.Language), a.Command))
		} else {
			lines = append(lines, fmt.Sprintf("  ⚠️  %s: %s not found", langDisplayName(a.Language), a.Command))
			lines = append(lines, fmt.Sprintf("     → Install: %s", a.InstallHint))
		}
	}
	return lines
}

// langDisplayName returns a display name for a language code.
func langDisplayName(lang string) string {
	switch lang {
	case "ts":
		return "TypeScript"
	case "py":
		return "Python"
	case "rs":
		return "Rust"
	case "go":
		return "Go"
	default:
		return lang
	}
}
