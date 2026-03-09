package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/code-certification/certify/internal/domain"
	"github.com/code-certification/certify/internal/record"
	"github.com/spf13/cobra"
)

var expireCmd = &cobra.Command{
	Use:   "expire",
	Short: "Mark overdue certifications as expired",
	RunE: func(cmd *cobra.Command, args []string) error {
		root, _ := os.Getwd()
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

func init() {
	rootCmd.AddCommand(expireCmd)
}
