package main

import (
	"fmt"
	"os"

	"github.com/iksnae/code-certification/internal/doctor"
	"github.com/spf13/cobra"
)

var onboardCmd = &cobra.Command{
	Use:   "onboard",
	Short: "Interactive onboarding guide",
	Long: `Show a step-by-step onboarding checklist for your project.

Checks which steps are already complete and tells you what to do next.
Re-run at any time to see your progress.

Steps:
  1. Initialize (.certification/ setup)
  2. Discover code units (scan)
  3. Run certification (evaluate + score)
  4. Generate report card
  5. Architect review (optional, requires AI)
  6. Add badge to README

Examples:
  certify onboard                # check current directory
  certify onboard --path /repo   # check specific repo`,
	RunE: func(cmd *cobra.Command, args []string) error {
		root := flagString(cmd, "path")
		if root == "" {
			var err error
			root, err = os.Getwd()
			if err != nil {
				return fmt.Errorf("getting working directory: %w", err)
			}
		}

		plan := doctor.BuildOnboardPlan(root)
		fmt.Print(doctor.FormatOnboardPlan(plan))
		return nil
	},
}

func bindOnboardFlags() {
	onboardCmd.Flags().String("path", "", "Path to repository (default: current directory)")
}
