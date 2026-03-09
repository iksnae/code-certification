package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version is set at build time via -ldflags.
// Release: v0.1.3
var (
	Version = "dev"
	Commit  = "unknown"
	Date    = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("certify %s (commit: %s, built: %s)\n", Version, Commit, Date)
	},
}
