package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestVersionCmd(t *testing.T) {
	// Just verify it doesn't panic
	versionCmd.Run(versionCmd, nil)
}

func TestInitCmd(t *testing.T) {
	dir := t.TempDir()
	orig, _ := os.Getwd()
	defer os.Chdir(orig)
	os.Chdir(dir)

	if err := initCmd.RunE(initCmd, nil); err != nil {
		t.Fatalf("init error: %v", err)
	}

	// Verify directory structure
	for _, sub := range []string{
		".certification",
		".certification/policies",
		".certification/records",
		".certification/overrides",
	} {
		if _, err := os.Stat(filepath.Join(dir, sub)); err != nil {
			t.Errorf("missing directory: %s", sub)
		}
	}

	// Verify config file
	if _, err := os.Stat(filepath.Join(dir, ".certification", "config.yml")); err != nil {
		t.Error("missing config.yml")
	}

	// Verify policy file
	if _, err := os.Stat(filepath.Join(dir, ".certification", "policies", "global.yml")); err != nil {
		t.Error("missing global.yml policy")
	}
}

func TestInitCmd_Idempotent(t *testing.T) {
	dir := t.TempDir()
	orig, _ := os.Getwd()
	defer os.Chdir(orig)
	os.Chdir(dir)

	// Run init twice — second should not error
	_ = initCmd.RunE(initCmd, nil)
	if err := initCmd.RunE(initCmd, nil); err != nil {
		t.Fatalf("second init should not error: %v", err)
	}
}
