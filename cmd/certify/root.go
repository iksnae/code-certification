package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "certify",
	Short: "Certify — code trust, with an expiration date",
	Long: `Certify continuously evaluates every code unit in your repository,
scores it against versioned policies, and assigns time-bound
certification you can actually trust.

Use 'certify init' to bootstrap certification in your repository.
Use --workspace for multi-repo operation across git submodules.`,
}

// flagString reads a string flag, returning "" on error (flag not registered = programming bug).
func flagString(cmd *cobra.Command, name string) string {
	v, err := cmd.Flags().GetString(name)
	if err != nil {
		return ""
	}
	return v
}

// flagBool reads a bool flag, returning false on error.
func flagBool(cmd *cobra.Command, name string) bool {
	v, err := cmd.Flags().GetBool(name)
	if err != nil {
		return false
	}
	return v
}

// flagInt reads an int flag, returning 0 on error.
func flagInt(cmd *cobra.Command, name string) int {
	v, err := cmd.Flags().GetInt(name)
	if err != nil {
		return 0
	}
	return v
}

// flagStringSlice reads a string slice flag, returning nil on error.
func flagStringSlice(cmd *cobra.Command, name string) []string {
	v, err := cmd.Flags().GetStringSlice(name)
	if err != nil {
		return nil
	}
	return v
}

// registerCommands wires all subcommands and their flags into rootCmd.
// Called from main() instead of using scattered init() functions,
// so no file in this package has an init() function.
func registerCommands() {
	rootCmd.PersistentFlags().Bool("workspace", false, "Workspace mode: operate across git submodules")

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
