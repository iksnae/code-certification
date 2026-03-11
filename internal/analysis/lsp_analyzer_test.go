package analysis

import (
	"fmt"
	"testing"
)

func TestNewLSPAnalyzer_UnsupportedLanguage(t *testing.T) {
	a, err := NewLSPAnalyzer("cobol", "/tmp")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if a != nil {
		t.Error("expected nil analyzer for unsupported language")
	}
}

func TestNewLSPAnalyzer_ServerNotInstalled(t *testing.T) {
	// Mock command lookup to always fail
	orig := execLookPath
	defer func() { execLookPath = orig }()
	execLookPath = func(cmd string) (string, error) {
		return "", fmt.Errorf("not found: %s", cmd)
	}

	a, err := NewLSPAnalyzer("ts", "/tmp")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if a != nil {
		t.Error("expected nil analyzer when server not installed (graceful degradation)")
	}
}

func TestDetectLSPServers(t *testing.T) {
	// Mock: all commands not found
	orig := execLookPath
	defer func() { execLookPath = orig }()
	execLookPath = func(cmd string) (string, error) {
		return "", fmt.Errorf("not found: %s", cmd)
	}

	avail := DetectLSPServers()
	if len(avail) != 3 {
		t.Fatalf("expected 3 servers, got %d", len(avail))
	}
	for _, a := range avail {
		if a.Available {
			t.Errorf("%s should not be available", a.Language)
		}
	}
}

func TestDetectLSPServers_WithServer(t *testing.T) {
	// Mock: typescript-language-server found
	orig := execLookPath
	defer func() { execLookPath = orig }()
	execLookPath = func(cmd string) (string, error) {
		if cmd == "typescript-language-server" {
			return "/usr/local/bin/typescript-language-server", nil
		}
		return "", fmt.Errorf("not found: %s", cmd)
	}

	avail := DetectLSPServers()
	for _, a := range avail {
		if a.Language == "ts" && !a.Available {
			t.Error("TypeScript server should be available")
		}
		if a.Language == "py" && a.Available {
			t.Error("Python server should not be available")
		}
	}
}

func TestFormatLSPStatus(t *testing.T) {
	avail := []LSPAvailability{
		{Language: "ts", Available: true, Command: "typescript-language-server"},
		{Language: "py", Available: false, Command: "pyright-langserver", InstallHint: "pip install pyright"},
	}

	lines := FormatLSPStatus(avail)
	if len(lines) != 3 { // 1 for ts (available), 2 for py (unavailable + hint)
		t.Errorf("expected 3 lines, got %d: %v", len(lines), lines)
	}
}

func TestFileURI(t *testing.T) {
	uri := fileURI("/home/user/project/src/main.ts")
	if uri == "" {
		t.Error("expected non-empty URI")
	}
	// Should start with file://
	if len(uri) < 7 || uri[:7] != "file://" {
		t.Errorf("URI should start with file://, got %q", uri)
	}
}

func TestLanguageID(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"ts", "typescript"},
		{"py", "python"},
		{"rs", "rust"},
		{"go", "go"},
	}
	for _, tt := range tests {
		got := languageID(tt.input)
		if got != tt.want {
			t.Errorf("languageID(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}
