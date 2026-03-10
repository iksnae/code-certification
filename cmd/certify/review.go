package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/iksnae/code-certification/internal/config"
	"github.com/iksnae/code-certification/internal/domain"
	gh "github.com/iksnae/code-certification/internal/github"
	"github.com/iksnae/code-certification/internal/record"
	"github.com/spf13/cobra"
)

var reviewCmd = &cobra.Command{
	Use:   "review",
	Short: "Generate PR review annotation",
	Long:  "Formats certification results as a PR comment. Use in GitHub Actions.",
	RunE: func(cmd *cobra.Command, args []string) error {
		root := flagString(cmd, "path")
		if root == "" {
			var err error
			root, err = os.Getwd()
			if err != nil {
				return fmt.Errorf("getting working directory: %w", err)
			}
		}
		certDir := filepath.Join(root, ".certification")

		// Load config for mode
		cfg, err := config.LoadFromDir(certDir)
		if err != nil {
			cfg = defaultConfigObj()
		}

		store := record.NewStore(filepath.Join(certDir, "records"))
		records, err := store.ListAll()
		if err != nil {
			return fmt.Errorf("loading records: %w", err)
		}

		enforcing := cfg.Mode == domain.ModeEnforcing
		comment := gh.FormatPRComment(records, enforcing)
		fmt.Print(comment)

		return nil
	},
}

func bindReviewFlags() {
	reviewCmd.Flags().String("path", "", "Path to repository (default: current directory)")
}
