package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestMain(m *testing.M) {
	registerCommands()
	os.Exit(m.Run())
}

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

func TestDetectRepoName(t *testing.T) {
	// Create a temp git repo with known remote
	dir := t.TempDir()
	run := func(args ...string) {
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Dir = dir
		if out, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("%v: %s", args, out)
		}
	}
	run("git", "init")
	run("git", "remote", "add", "origin", "https://github.com/someuser/their-repo.git")

	got := detectRepoName(dir)
	if got != "someuser/their-repo" {
		t.Errorf("HTTPS remote: got %q, want %q", got, "someuser/their-repo")
	}

	// Switch to SSH
	run("git", "remote", "set-url", "origin", "git@github.com:someuser/their-repo.git")
	got = detectRepoName(dir)
	if got != "someuser/their-repo" {
		t.Errorf("SSH remote: got %q, want %q", got, "someuser/their-repo")
	}
}
