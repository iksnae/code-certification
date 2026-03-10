package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/record"
	"github.com/iksnae/code-certification/internal/workspace"
	"github.com/spf13/cobra"
)

var expirePath string

var expireCmd = &cobra.Command{
	Use:   "expire",
	Short: "Mark overdue certifications as expired",
	RunE: func(cmd *cobra.Command, args []string) error {
		root := expirePath
		if root == "" {
			root, _ = os.Getwd()
		}

		if workspaceMode {
			return runWorkspaceExpire(root)
		}

		certDir := filepath.Join(root, ".certification")
		store := record.NewStore(filepath.Join(certDir, "records"))

		records, err := store.ListAll()
		if err != nil {
			return fmt.Errorf("loading records: %w", err)
		}

		now := time.Now()
		var expired int

		for _, rec := range records {
			if rec.Status == domain.StatusCertified || rec.Status == domain.StatusCertifiedWithObservations {
				if now.After(rec.ExpiresAt) {
					rec.Status = domain.StatusExpired
					rec.Observations = append(rec.Observations, fmt.Sprintf("Expired at %s", now.Format(time.RFC3339)))
					if err := store.Save(rec); err != nil {
						fmt.Fprintf(os.Stderr, "warning: saving expired record for %s: %v\n", rec.UnitID, err)
					}
					expired++
				}
			}
		}

		fmt.Printf("✓ Checked %d records, %d newly expired\n", len(records), expired)
		return nil
	},
}

func bindExpireFlags() {
	expireCmd.Flags().StringVar(&expirePath, "path", "", "Path to repository (default: current directory)")
}

func runWorkspaceExpire(root string) error {
	subs, err := workspace.DiscoverSubmodules(root)
	if err != nil {
		return fmt.Errorf("discovering submodules: %w", err)
	}

	configured := workspace.ConfiguredSubmodules(subs)
	if len(configured) == 0 {
		return fmt.Errorf("no configured submodules found — run 'certify init --workspace' first")
	}

	fmt.Printf("🔍 Workspace expire: %d submodule(s)\n\n", len(configured))

	for _, s := range configured {
		fmt.Printf("  → Expiring %s...\n", s.Path)
		subPath := filepath.Join(root, s.Path)
		if err := runSubcommand("expire", "--path", subPath); err != nil {
			fmt.Fprintf(os.Stderr, "    warning: expire failed for %s: %v\n", s.Path, err)
		}
	}

	fmt.Println("\n✓ Workspace expiry complete.")
	return nil
}
