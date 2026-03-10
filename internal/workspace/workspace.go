// Package workspace provides multi-repo workspace support for certify.
// It discovers git submodules, checks their certification status, and
// aggregates reports across submodules.
package workspace

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/iksnae/code-certification/internal/record"
	"github.com/iksnae/code-certification/internal/report"
)

// Submodule represents a git submodule within a workspace.
type Submodule struct {
	Name      string // submodule name (typically same as path)
	Path      string // relative path within workspace
	URL       string // remote URL (if available)
	Commit    string // current checked-out commit SHA
	HasConfig bool   // true if .certification/config.yml exists
}

// DiscoverSubmodules finds all git submodules in the workspace root,
// checks which ones have certify configured, and returns them sorted by path.
func DiscoverSubmodules(root string) ([]Submodule, error) {
	cmd := exec.Command("git", "submodule", "status")
	cmd.Dir = root
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("git submodule status: %w", err)
	}

	subs := ParseSubmoduleStatus(string(out))
	for i := range subs {
		subs[i].HasConfig = CheckHasConfig(root, subs[i])
	}

	sort.Slice(subs, func(i, j int) bool {
		return subs[i].Path < subs[j].Path
	})

	return subs, nil
}

// ParseSubmoduleStatus parses the output of `git submodule status`.
// Each line format: `[+-U ]<commit> <path> (<branch>)`
// The leading character indicates: ' ' = normal, '+' = different commit,
// '-' = not initialized, 'U' = merge conflict.
func ParseSubmoduleStatus(output string) []Submodule {
	var subs []Submodule
	for _, line := range strings.Split(output, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Strip leading status character (+, -, U, or space)
		if len(line) > 0 && (line[0] == '+' || line[0] == '-' || line[0] == 'U') {
			line = line[1:]
		}

		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}

		commit := parts[0]
		path := parts[1]

		subs = append(subs, Submodule{
			Name:   path,
			Path:   path,
			Commit: commit,
		})
	}
	return subs
}

// CheckHasConfig checks whether a submodule has a .certification/config.yml file.
func CheckHasConfig(root string, sub Submodule) bool {
	configPath := filepath.Join(root, sub.Path, ".certification", "config.yml")
	_, err := os.Stat(configPath)
	return err == nil
}

// LoadSubmoduleCard loads certification state from a submodule and generates
// a report Card. Returns (nil, nil) if the submodule has no state.json.
func LoadSubmoduleCard(root string, sub Submodule) (*report.Card, error) {
	certDir := filepath.Join(root, sub.Path, ".certification")
	statePath := filepath.Join(certDir, "state.json")

	if _, err := os.Stat(statePath); os.IsNotExist(err) {
		return nil, nil
	}

	store := record.NewStoreWithSnapshot(filepath.Join(certDir, "records"), statePath)
	records, err := store.ListAll()
	if err != nil {
		return nil, fmt.Errorf("loading records for %s: %w", sub.Name, err)
	}

	if len(records) == 0 {
		return nil, nil
	}

	now := time.Now()
	card := report.GenerateCard(records, sub.Name, sub.Commit, now)
	return &card, nil
}

// ConfiguredSubmodules returns only submodules that have certify configured.
func ConfiguredSubmodules(subs []Submodule) []Submodule {
	var configured []Submodule
	for _, s := range subs {
		if s.HasConfig {
			configured = append(configured, s)
		}
	}
	return configured
}
