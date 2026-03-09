package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "certify",
	Short: "Code Certification System — time-bound trust for every code unit",
	Long: `certify discovers code units, evaluates them against versioned policies,
collects deterministic evidence, scores across 9 quality dimensions,
assigns time-bound certification status, and generates reports.

Use 'certify init' to bootstrap certification in your repository.`,
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(scanCmd)
	rootCmd.AddCommand(certifyCmd)
	rootCmd.AddCommand(reportCmd)
}
