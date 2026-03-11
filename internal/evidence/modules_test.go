package evidence

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDiscoverModuleRoots_GoInSubdir(t *testing.T) {
	root := t.TempDir()

	// go.mod in code/ not root
	codeDir := filepath.Join(root, "code")
	os.MkdirAll(codeDir, 0o755)
	os.WriteFile(filepath.Join(codeDir, "go.mod"), []byte("module example.com/code"), 0o644)
	os.WriteFile(filepath.Join(codeDir, "main.go"), []byte("package main"), 0o644)

	roots := DiscoverModuleRoots(root)

	var goRoots []ModuleRoot
	for _, r := range roots {
		if r.Language == "go" {
			goRoots = append(goRoots, r)
		}
	}

	if len(goRoots) != 1 {
		t.Fatalf("expected 1 Go root, got %d: %+v", len(goRoots), goRoots)
	}
	if goRoots[0].RelPath != "code" {
		t.Errorf("RelPath = %q, want 'code'", goRoots[0].RelPath)
	}
	if goRoots[0].Marker != "go.mod" {
		t.Errorf("Marker = %q, want 'go.mod'", goRoots[0].Marker)
	}
}

func TestDiscoverModuleRoots_MultipleModules(t *testing.T) {
	root := t.TempDir()

	// Root has go.mod
	os.WriteFile(filepath.Join(root, "go.mod"), []byte("module root"), 0o644)

	// services/api also has go.mod
	apiDir := filepath.Join(root, "services", "api")
	os.MkdirAll(apiDir, 0o755)
	os.WriteFile(filepath.Join(apiDir, "go.mod"), []byte("module api"), 0o644)

	// Also a package.json at root
	os.WriteFile(filepath.Join(root, "package.json"), []byte("{}"), 0o644)

	roots := DiscoverModuleRoots(root)

	goRoots := GoModuleRoots(roots)
	if len(goRoots) != 2 {
		t.Fatalf("expected 2 Go roots, got %d", len(goRoots))
	}

	// Check we found both
	relPaths := map[string]bool{}
	for _, r := range goRoots {
		relPaths[r.RelPath] = true
	}
	if !relPaths[""] {
		t.Error("missing root module")
	}
	if !relPaths["services/api"] {
		t.Error("missing services/api module")
	}
}

func TestDiscoverModuleRoots_SkipsVendor(t *testing.T) {
	root := t.TempDir()

	os.WriteFile(filepath.Join(root, "go.mod"), []byte("module root"), 0o644)

	vendorDir := filepath.Join(root, "vendor", "dep")
	os.MkdirAll(vendorDir, 0o755)
	os.WriteFile(filepath.Join(vendorDir, "go.mod"), []byte("module dep"), 0o644)

	roots := DiscoverModuleRoots(root)
	goRoots := GoModuleRoots(roots)

	if len(goRoots) != 1 {
		t.Fatalf("expected 1 Go root (vendor excluded), got %d", len(goRoots))
	}
}

func TestUnitBelongsToModule(t *testing.T) {
	rootMod := ModuleRoot{RelPath: ""}
	codeMod := ModuleRoot{RelPath: "code"}
	apiMod := ModuleRoot{RelPath: "services/api"}

	tests := []struct {
		unit   string
		module ModuleRoot
		want   bool
	}{
		{"main.go", rootMod, true},
		{"code/main.go", rootMod, true},
		{"code/main.go", codeMod, true},
		{"code/pkg/util.go", codeMod, true},
		{"other/main.go", codeMod, false},
		{"services/api/handler.go", apiMod, true},
		{"services/web/handler.go", apiMod, false},
	}

	for _, tt := range tests {
		got := UnitBelongsToModule(tt.unit, tt.module)
		if got != tt.want {
			t.Errorf("UnitBelongsToModule(%q, %q) = %v, want %v", tt.unit, tt.module.RelPath, got, tt.want)
		}
	}
}

func TestFindModuleForUnit_MostSpecific(t *testing.T) {
	modules := []ModuleRoot{
		{RelPath: "", Language: "go"},
		{RelPath: "code", Language: "go"},
		{RelPath: "code/services/api", Language: "go"},
	}

	// Unit in code/services/api/ should match the most specific module
	m := FindModuleForUnit("code/services/api/handler.go", modules)
	if m == nil {
		t.Fatal("expected match")
	}
	if m.RelPath != "code/services/api" {
		t.Errorf("matched %q, want code/services/api", m.RelPath)
	}

	// Unit in code/ but not api/ should match code
	m = FindModuleForUnit("code/pkg/util.go", modules)
	if m == nil {
		t.Fatal("expected match")
	}
	if m.RelPath != "code" {
		t.Errorf("matched %q, want code", m.RelPath)
	}

	// Unit at root should match root module
	m = FindModuleForUnit("main.go", modules)
	if m == nil {
		t.Fatal("expected match")
	}
	if m.RelPath != "" {
		t.Errorf("matched %q, want root", m.RelPath)
	}
}

func TestFindModuleForUnit_NoMatch(t *testing.T) {
	modules := []ModuleRoot{
		{RelPath: "code", Language: "go"},
	}

	m := FindModuleForUnit("other/main.go", modules)
	if m != nil {
		t.Errorf("expected nil, got %q", m.RelPath)
	}
}

func TestDiscoverModuleRoots_MixedLanguages(t *testing.T) {
	root := t.TempDir()

	os.WriteFile(filepath.Join(root, "go.mod"), []byte("module root"), 0o644)
	os.MkdirAll(filepath.Join(root, "frontend"), 0o755)
	os.WriteFile(filepath.Join(root, "frontend", "package.json"), []byte("{}"), 0o644)
	os.MkdirAll(filepath.Join(root, "ml"), 0o755)
	os.WriteFile(filepath.Join(root, "ml", "pyproject.toml"), []byte("[project]"), 0o644)

	roots := DiscoverModuleRoots(root)

	langs := map[string]int{}
	for _, r := range roots {
		langs[r.Language]++
	}

	if langs["go"] != 1 {
		t.Errorf("go modules = %d, want 1", langs["go"])
	}
	if langs["ts"] != 1 {
		t.Errorf("ts modules = %d, want 1", langs["ts"])
	}
	if langs["py"] != 1 {
		t.Errorf("py modules = %d, want 1", langs["py"])
	}
}
