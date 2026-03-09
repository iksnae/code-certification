package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/code-certification/certify/internal/config"
	"github.com/code-certification/certify/internal/domain"
	gh "github.com/code-certification/certify/internal/github"
	"github.com/code-certification/certify/internal/record"
	"github.com/spf13/cobra"
)

var reviewCmd = &cobra.Command{
	Use:   "review",
	Short: "Generate PR review annotation",
	Long:  "Formats certification results as a PR comment. Use in GitHub Actions.",
	RunE: func(cmd *cobra.Command, args []string) error {
		root, _ := os.Getwd()
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

func init() {
	rootCmd.AddCommand(reviewCmd)
}
