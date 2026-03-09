package discovery_test

import (
	"os"
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

func TestGenericScanner_SkipsNonCertifiable(t *testing.T) {
	dir := t.TempDir()
	// Create a mix of certifiable and non-certifiable files
	certifiable := []string{"main.go", "app.ts", "lib.py", "run.sh", "server.rs", "Main.java"}
	nonCertifiable := []string{
		"logo.png", "icon.svg", "photo.jpg", "anim.gif", "pic.webp",  // images
		"data.json", "config.xml", "style.css", "page.html", "doc.md", // data/markup
		"app.wasm", "lib.so", "prog.exe",                               // binaries
		"bundle.min.js", "bundle.min.css",                               // minified
		"package-lock.json", "yarn.lock", "go.sum",                      // lock files
		"output.map", "bundle.vsix",                                     // artifacts
		"font.woff", "font.woff2", "font.ttf", "font.eot",             // fonts
		"video.mp4", "audio.mp3",                                        // media
		"archive.zip", "backup.tar.gz",                                  // archives
	}

	for _, f := range certifiable {
		writeFile(t, dir, f, "// code")
	}
	for _, f := range nonCertifiable {
		writeFile(t, dir, f, "content")
	}

	s := discovery.NewGenericScanner(nil, nil)
	units, err := s.Scan(dir)
	if err != nil {
		t.Fatalf("Scan() error: %v", err)
	}

	paths := make(map[string]bool)
	for _, u := range units {
		paths[u.ID.Path()] = true
	}

	// All certifiable files should be found
	for _, f := range certifiable {
		if !paths[f] {
			t.Errorf("certifiable file %q was not discovered", f)
		}
	}

	// No non-certifiable files should be found
	for _, f := range nonCertifiable {
		if paths[f] {
			t.Errorf("non-certifiable file %q should not be discovered", f)
		}
	}
}

func TestGenericScanner_SkipsBuildOutputDirs(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, dir, "src/main.go", "package main")
	writeFile(t, dir, "dist/bundle.js", "compiled")
	writeFile(t, dir, "out/index.js", "compiled")
	writeFile(t, dir, "coverage/lcov.info", "data")

	s := discovery.NewGenericScanner(nil, nil)
	units, err := s.Scan(dir)
	if err != nil {
		t.Fatalf("Scan() error: %v", err)
	}

	for _, u := range units {
		path := u.ID.Path()
		for _, bad := range []string{"dist/", "out/", "coverage/"} {
			if filepath.HasPrefix(path, bad) {
				t.Errorf("build output %q should be excluded", path)
			}
		}
	}
}

func writeFile(t *testing.T, dir, name, content string) {
	t.Helper()
	full := filepath.Join(dir, name)
	if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(full, []byte(content), 0o644); err != nil {
		t.Fatal(err)
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
