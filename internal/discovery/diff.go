package discovery

import (
	"os/exec"
	"strings"

	"github.com/iksnae/code-certification/internal/domain"
)

// ChangedFiles returns files changed between two git refs (or working tree if ref2 is empty).
func ChangedFiles(root, ref1, ref2 string) ([]string, error) {
	args := []string{"diff", "--name-only"}
	if ref2 != "" {
		args = append(args, ref1+"..."+ref2)
	} else {
		args = append(args, ref1)
	}

	cmd := exec.Command("git", args...)
	cmd.Dir = root
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	var files []string
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if l != "" {
			files = append(files, l)
		}
	}
	return files, nil
}

// FilterChanged returns only units whose paths appear in the changed files list.
func FilterChanged(units []domain.Unit, changedFiles []string) []domain.Unit {
	changed := make(map[string]bool, len(changedFiles))
	for _, f := range changedFiles {
		changed[f] = true
	}

	var result []domain.Unit
	for _, u := range units {
		if changed[u.ID.Path()] {
			result = append(result, u)
		}
	}
	return result
}

// DetectMoves identifies likely file renames between old and new unit lists.
// Uses git's rename detection when available, falls back to symbol matching.
func DetectMoves(root string) ([]MovedFile, error) {
	cmd := exec.Command("git", "diff", "--name-status", "-M", "HEAD~1", "HEAD")
	cmd.Dir = root
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var moves []MovedFile
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) >= 3 && strings.HasPrefix(parts[0], "R") {
			moves = append(moves, MovedFile{
				OldPath: parts[1],
				NewPath: parts[2],
			})
		}
	}
	return moves, nil
}

// MovedFile represents a file rename detected by git.
type MovedFile struct {
	OldPath string
	NewPath string
}

// FilterByPaths returns units whose paths match any of the given path prefixes.
func FilterByPaths(units []domain.Unit, paths []string) []domain.Unit {
	if len(paths) == 0 {
		return units
	}

	var result []domain.Unit
	for _, u := range units {
		for _, p := range paths {
			if strings.HasPrefix(u.ID.Path(), p) || u.ID.Path() == p {
				result = append(result, u)
				break
			}
		}
	}
	return result
}
