package main

import (
	"fmt"
	"os"
	"os/exec"
)

// runSubcommand runs a certify subcommand as a child process.
// This avoids initialization cycles from workspace functions referencing cmd vars.
// Uses os.Args[0] to call the same binary.
func runSubcommand(args ...string) error {
	cmd := exec.Command(os.Args[0], args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("subcommand %v failed: %w", args, err)
	}
	return nil
}
