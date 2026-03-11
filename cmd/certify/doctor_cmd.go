package main

import (
	"fmt"
	"os"

	"github.com/iksnae/code-certification/internal/doctor"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check setup and diagnose issues",
	Long: `Run health checks on your Certify installation and project setup.

Checks:
  - Environment (Go, Git)
  - Project setup (.certification/ directory)
  - Configuration validity
  - Policy packs
  - Optional tools (golangci-lint, gh)
  - AI providers (auto-detection)

Examples:
  certify doctor                # check current directory
  certify doctor --path /repo   # check specific repo`,
	RunE: func(cmd *cobra.Command, args []string) error {
		root := flagString(cmd, "path")
		if root == "" {
			var err error
			root, err = os.Getwd()
			if err != nil {
				return fmt.Errorf("getting working directory: %w", err)
			}
		}

		report := doctor.RunAll(root)
		fmt.Print(doctor.FormatReport(report))

		if report.HasFailures() {
			os.Exit(1)
		}
		return nil
	},
}

func bindDoctorFlags() {
	doctorCmd.Flags().String("path", "", "Path to repository (default: current directory)")
}
