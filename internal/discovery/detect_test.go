package discovery_test

import (
	"testing"

	"github.com/code-certification/certify/internal/discovery"
)

func TestDetectLanguages_GoRepo(t *testing.T) {
	langs := discovery.DetectLanguages(repoPath("go-simple"))
	if len(langs) == 0 {
		t.Fatal("should detect at least one language")
	}
	if langs[0].Name != "go" {
		t.Errorf("primary language = %q, want go", langs[0].Name)
	}
	if langs[0].FileCount < 2 {
		t.Errorf("go file count = %d, want >= 2", langs[0].FileCount)
	}
}

func TestDetectLanguages_TSRepo(t *testing.T) {
	langs := discovery.DetectLanguages(repoPath("ts-simple"))
	if len(langs) == 0 {
		t.Fatal("should detect at least one language")
	}
	found := false
	for _, l := range langs {
		if l.Name == "typescript" {
			found = true
			if l.FileCount < 2 {
				t.Errorf("ts file count = %d, want >= 2", l.FileCount)
			}
		}
	}
	if !found {
		t.Error("should detect typescript")
	}
}

func TestDetectLanguages_EmptyDir(t *testing.T) {
	dir := t.TempDir()
	langs := discovery.DetectLanguages(dir)
	if len(langs) != 0 {
		t.Errorf("empty dir should detect 0 languages, got %d", len(langs))
	}
}

func TestDetectedAdapters(t *testing.T) {
	langs := []discovery.LanguageInfo{
		{Name: "go", Adapter: "go"},
		{Name: "typescript", Adapter: "ts"},
		{Name: "python", Adapter: "generic"},
	}
	adapters := discovery.DetectedAdapters(langs)
	if len(adapters) != 3 {
		t.Errorf("adapters = %d, want 3", len(adapters))
	}
}
