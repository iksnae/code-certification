// Package main is the entry point for the certify CLI.
package main

import (
	"fmt"
	"os"
)

func main() {
	registerCommands()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
