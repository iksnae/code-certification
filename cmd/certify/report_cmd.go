package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/code-certification/certify/internal/record"
	"github.com/code-certification/certify/internal/report"
	"github.com/spf13/cobra"
)

var (
	reportFormat string
	reportPath   string
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Generate certification reports",
	RunE: func(cmd *cobra.Command, args []string) error {
		root := reportPath
		if root == "" {
			root, _ = os.Getwd()
		}
		certDir := filepath.Join(root, ".certification")

		store := record.NewStore(filepath.Join(certDir, "records"))
		records, err := store.ListAll()
		if err != nil {
			return fmt.Errorf("loading records: %w", err)
		}

		if len(records) == 0 {
			fmt.Println("No certification records found. Run 'certify certify' first.")
			return nil
		}

		h := report.Health(records)

		switch reportFormat {
		case "json":
			data, err := report.FormatJSON(h)
			if err != nil {
				return err
			}
			fmt.Println(string(data))
		default:
			fmt.Print(report.FormatText(h))
		}

		return nil
	},
}

func init() {
	reportCmd.Flags().StringVarP(&reportFormat, "format", "f", "text", "Output format (text, json)")
	reportCmd.Flags().StringVar(&reportPath, "path", "", "Path to repository (default: current directory)")
}
