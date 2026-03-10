package main

import (
	"fmt"
	"runtime/debug"

	"github.com/spf13/cobra"
)

// Version is set at build time via -ldflags.
var (
	Version = "dev"
	Commit  = "unknown"
	Date    = "unknown"
)

func init() {
	// When installed via `go install`, ldflags aren't set.
	// Fall back to embedded build info from the Go toolchain.
	if Version == "dev" {
		if info, ok := debug.ReadBuildInfo(); ok {
			if info.Main.Version != "" && info.Main.Version != "(devel)" {
				Version = info.Main.Version
			}
			for _, s := range info.Settings {
				if s.Key == "vcs.revision" && len(s.Value) >= 7 {
					Commit = s.Value[:7]
				}
				if s.Key == "vcs.time" {
					Date = s.Value
				}
			}
		}
	}
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("certify %s (commit: %s, built: %s)\n", Version, Commit, Date)
	},
}
