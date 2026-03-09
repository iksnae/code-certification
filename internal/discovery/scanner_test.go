package discovery_test

import (
	"path/filepath"
	"testing"

	"github.com/iksnae/code-certification/internal/discovery"
	"github.com/iksnae/code-certification/internal/domain"
)

func repoPath(name string) string {
	return filepath.Join("..", "..", "testdata", "repos", name)
}

func TestGenericScanner_DiscoverAllFiles(t *testing.T) {
	s := discovery.NewGenericScanner(nil, nil)
	units, err := s.Scan(repoPath("go-simple"))
	if err != nil {
		t.Fatalf("Scan() error: %v", err)
	}

	// Should find main.go and internal/service/sync.go
	if len(units) < 2 {
		t.Errorf("len(units) = %d, want at least 2", len(units))
	}

	// All should be file-level units
	for _, u := range units {
		if u.Type != domain.UnitTypeFile {
			t.Errorf("unit %s has type %v, want file", u.ID, u.Type)
		}
	}
}

func TestGenericScanner_IncludePatterns(t *testing.T) {
	s := discovery.NewGenericScanner([]string{"*.go"}, nil)
	units, err := s.Scan(repoPath("go-simple"))
	if err != nil {
		t.Fatalf("Scan() error: %v", err)
	}

	for _, u := range units {
		if filepath.Ext(u.ID.Path()) != ".go" {
			t.Errorf("unit %s should not be included (only *.go)", u.ID)
		}
	}
}

func TestGenericScanner_ExcludePatterns(t *testing.T) {
	s := discovery.NewGenericScanner(nil, []string{"internal/**"})
	units, err := s.Scan(repoPath("go-simple"))
	if err != nil {
		t.Fatalf("Scan() error: %v", err)
	}

	for _, u := range units {
		path := u.ID.Path()
		if filepath.Dir(path) == "internal" || filepath.Dir(filepath.Dir(path)) == "internal" {
			t.Errorf("unit %s should be excluded by internal/** pattern", u.ID)
		}
	}
}

func TestGenericScanner_NestedDirs(t *testing.T) {
	s := discovery.NewGenericScanner([]string{"*.go"}, nil)
	units, err := s.Scan(repoPath("go-simple"))
	if err != nil {
		t.Fatalf("Scan() error: %v", err)
	}

	// Should find files in nested dirs
	found := make(map[string]bool)
	for _, u := range units {
		found[u.ID.Path()] = true
	}
	if !found["main.go"] {
		t.Error("missing main.go")
	}
	if !found["internal/service/sync.go"] {
		t.Error("missing internal/service/sync.go")
	}
}

func TestGenericScanner_StableIDs(t *testing.T) {
	s := discovery.NewGenericScanner([]string{"*.go"}, nil)
	units1, err := s.Scan(repoPath("go-simple"))
	if err != nil {
		t.Fatal(err)
	}
	units2, err := s.Scan(repoPath("go-simple"))
	if err != nil {
		t.Fatal(err)
	}

	if len(units1) != len(units2) {
		t.Fatal("scan should be deterministic")
	}
	for i := range units1 {
		if units1[i].ID.String() != units2[i].ID.String() {
			t.Errorf("ID mismatch: %s != %s", units1[i].ID, units2[i].ID)
		}
	}
}

func TestGenericScanner_TSFiles(t *testing.T) {
	s := discovery.NewGenericScanner([]string{"*.ts"}, nil)
	units, err := s.Scan(repoPath("ts-simple"))
	if err != nil {
		t.Fatal(err)
	}

	if len(units) < 2 {
		t.Errorf("len(units) = %d, want at least 2", len(units))
	}

	for _, u := range units {
		if u.ID.Language() != "ts" {
			t.Errorf("unit %s language = %q, want ts", u.ID, u.ID.Language())
		}
	}
}

func TestGenericScanner_EmptyDir(t *testing.T) {
	dir := t.TempDir()
	s := discovery.NewGenericScanner(nil, nil)
	units, err := s.Scan(dir)
	if err != nil {
		t.Fatalf("Scan(empty) error: %v", err)
	}
	if len(units) != 0 {
		t.Errorf("empty dir should have 0 units, got %d", len(units))
	}
}
