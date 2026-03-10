package main

import (
	"github.com/spf13/cobra"
)

var workspaceMode bool

var rootCmd = &cobra.Command{
	Use:   "certify",
	Short: "Certify — code trust, with an expiration date",
	Long: `Certify continuously evaluates every code unit in your repository,
scores it against versioned policies, and assigns time-bound
certification you can actually trust.

Use 'certify init' to bootstrap certification in your repository.
Use --workspace for multi-repo operation across git submodules.`,
}

// registerCommands wires all subcommands and their flags into rootCmd.
// Called from main() instead of using scattered init() functions,
// so no file in this package has an init() function.
func registerCommands() {
	rootCmd.PersistentFlags().BoolVar(&workspaceMode, "workspace", false, "Workspace mode: operate across git submodules")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(scanCmd)
	rootCmd.AddCommand(certifyCmd)
	rootCmd.AddCommand(reportCmd)
	rootCmd.AddCommand(modelsCmd)
	rootCmd.AddCommand(architectCmd)
	rootCmd.AddCommand(expireCmd)
	rootCmd.AddCommand(reviewCmd)

	// Bind per-command flags
	bindVersionInfo()
	bindCertifyFlags()
	bindScanFlags()
	bindReportFlags()
	bindInitFlags()
	bindArchitectFlags()
	bindExpireFlags()
	bindModelsFlags()
	bindReviewFlags()
}
