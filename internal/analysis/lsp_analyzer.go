package analysis

import (
	"fmt"
	"net/url"
	"os"
	osexec "os/exec"
	"path/filepath"

	"github.com/iksnae/code-certification/internal/analysis/lsp"
)

// LSPAnalyzer wraps an LSP client to provide Tier 2 analysis for non-Go languages.
type LSPAnalyzer struct {
	client  *lsp.Client
	lang    string
	rootDir string
	rootURI string
}

// LSPServerConfig describes how to start a language server.
type LSPServerConfig struct {
	Command string
	Args    []string
}

// DefaultLSPServers maps language identifiers to their default server configs.
var DefaultLSPServers = map[string]LSPServerConfig{
	"ts": {Command: "typescript-language-server", Args: []string{"--stdio"}},
	"py": {Command: "pyright-langserver", Args: []string{"--stdio"}},
	"rs": {Command: "rust-analyzer", Args: nil},
}

// NewLSPAnalyzer starts the appropriate language server and returns an analyzer.
// Returns nil, nil if no LSP server is available for the language.
func NewLSPAnalyzer(lang, rootDir string) (*LSPAnalyzer, error) {
	cfg, ok := DefaultLSPServers[lang]
	if !ok {
		return nil, nil // no server configured
	}

	// Check if the command is available
	if !commandExists(cfg.Command) {
		return nil, nil // server not installed — graceful degradation
	}

	absRoot, err := filepath.Abs(rootDir)
	if err != nil {
		return nil, err
	}
	rootURI := fileURI(absRoot)

	client, err := lsp.Start(cfg.Command, cfg.Args, absRoot)
	if err != nil {
		return nil, fmt.Errorf("start %s: %w", cfg.Command, err)
	}

	_, err = client.Initialize(rootURI)
	if err != nil {
		client.Shutdown()
		return nil, fmt.Errorf("initialize %s: %w", cfg.Command, err)
	}

	return &LSPAnalyzer{
		client:  client,
		lang:    lang,
		rootDir: absRoot,
		rootURI: rootURI,
	}, nil
}

// FanIn returns incoming call count for a symbol at the given position.
func (a *LSPAnalyzer) FanIn(file string, line, col int) (int, error) {
	uri := fileURI(filepath.Join(a.rootDir, file))

	items, err := a.client.CallHierarchyPrepare(uri, line, col)
	if err != nil || len(items) == 0 {
		return 0, err
	}

	incoming, err := a.client.CallHierarchyIncoming(items[0])
	if err != nil {
		return 0, err
	}
	return len(incoming), nil
}

// FanOut returns outgoing call count from a symbol at the given position.
func (a *LSPAnalyzer) FanOut(file string, line, col int) (int, error) {
	uri := fileURI(filepath.Join(a.rootDir, file))

	items, err := a.client.CallHierarchyPrepare(uri, line, col)
	if err != nil || len(items) == 0 {
		return 0, err
	}

	outgoing, err := a.client.CallHierarchyOutgoing(items[0])
	if err != nil {
		return 0, err
	}
	return len(outgoing), nil
}

// IsDeadCode returns true if symbol has zero references outside its file.
func (a *LSPAnalyzer) IsDeadCode(file string, line, col int) (bool, error) {
	uri := fileURI(filepath.Join(a.rootDir, file))

	refs, err := a.client.References(uri, line, col)
	if err != nil {
		return false, err
	}

	// Filter: exclude references in the same file
	externalRefs := 0
	for _, ref := range refs {
		if ref.URI != uri {
			externalRefs++
		}
	}
	return externalRefs == 0, nil
}

// OpenFile sends a textDocument/didOpen notification for a file.
func (a *LSPAnalyzer) OpenFile(file string) error {
	absPath := filepath.Join(a.rootDir, file)
	data, err := os.ReadFile(absPath)
	if err != nil {
		return err
	}
	uri := fileURI(absPath)
	langID := languageID(a.lang)
	return a.client.DidOpen(uri, langID, string(data))
}

// Shutdown gracefully stops the language server.
func (a *LSPAnalyzer) Shutdown() error {
	if a.client == nil {
		return nil
	}
	return a.client.Shutdown()
}

// fileURI converts a file path to a file:// URI.
func fileURI(path string) string {
	abs, _ := filepath.Abs(path)
	return "file://" + url.PathEscape(abs)
}

// languageID maps our language codes to LSP language identifiers.
func languageID(lang string) string {
	switch lang {
	case "ts":
		return "typescript"
	case "py":
		return "python"
	case "rs":
		return "rust"
	default:
		return lang
	}
}

// commandExists checks if a command is available on PATH.
func commandExists(cmd string) bool {
	_, err := execLookPath(cmd)
	return err == nil
}

// execLookPath is a variable for testing. Defaults to exec.LookPath.
var execLookPath = defaultLookPath

func defaultLookPath(cmd string) (string, error) {
	return osexec.LookPath(cmd)
}
