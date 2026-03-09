package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/code-certification/certify/internal/record"
	"github.com/code-certification/certify/internal/report"
	"github.com/spf13/cobra"
)

var (
	reportFormat   string
	reportPath     string
	reportDetailed bool
	reportOutput   string
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Generate certification reports",
	Long: `Generate certification reports in various formats.

Formats:
  text      Human-readable summary (default)
  json      Full report as machine-readable JSON (every unit + dimensions)
  card      Report card (text box with grade)
  markdown  Report card as GitHub-friendly markdown
  full      Complete per-unit report with all dimensions and scores

Output:
  By default, reports print to stdout. Use --output to write to a file.
  The report card is also saved to .certification/REPORT_CARD.md.
  The full report is saved to .certification/FULL_REPORT.md.`,
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

		now := time.Now()
		repo := detectRepoName(root)
		commit := detectCommit(root)

		var output string

		switch reportFormat {
		case "json":
			fr := report.GenerateFullReport(records, repo, commit, now)
			data, err := report.FormatJSON(fr)
			if err != nil {
				return err
			}
			output = string(data)

		case "card":
			c := report.GenerateCard(records, repo, commit, now)
			output = report.FormatCardText(c)

		case "markdown", "md":
			c := report.GenerateCard(records, repo, commit, now)
			output = report.FormatCardMarkdown(c)
			cardPath := filepath.Join(certDir, "REPORT_CARD.md")
			os.WriteFile(cardPath, []byte(output), 0o644)

		case "full":
			fr := report.GenerateFullReport(records, repo, commit, now)
			output = report.FormatFullMarkdown(fr)
			fullPath := filepath.Join(certDir, "FULL_REPORT.md")
			os.WriteFile(fullPath, []byte(output), 0o644)

		default:
			if reportDetailed {
				d := report.Detailed(records, now)
				output = report.FormatDetailedText(d)
			} else {
				h := report.Health(records)
				output = report.FormatText(h)
			}
		}

		// Write to file or stdout
		if reportOutput != "" {
			if err := os.MkdirAll(filepath.Dir(reportOutput), 0o755); err != nil {
				return err
			}
			if err := os.WriteFile(reportOutput, []byte(output), 0o644); err != nil {
				return err
			}
			fmt.Printf("✓ Report written to %s\n", reportOutput)
		} else {
			fmt.Print(output)
		}

		// Always save both report card and full report alongside records
		if reportFormat != "markdown" && reportFormat != "md" {
			c := report.GenerateCard(records, repo, commit, now)
			md := report.FormatCardMarkdown(c)
			os.WriteFile(filepath.Join(certDir, "REPORT_CARD.md"), []byte(md), 0o644)
		}
		if reportFormat != "full" {
			fr := report.GenerateFullReport(records, repo, commit, now)
			fullMD := report.FormatFullMarkdown(fr)
			os.WriteFile(filepath.Join(certDir, "FULL_REPORT.md"), []byte(fullMD), 0o644)
		}

		return nil
	},
}

func init() {
	reportCmd.Flags().StringVarP(&reportFormat, "format", "f", "text", "Output format (text, json, card, markdown, full)")
	reportCmd.Flags().StringVar(&reportPath, "path", "", "Path to repository (default: current directory)")
	reportCmd.Flags().BoolVar(&reportDetailed, "detailed", false, "Include dimension breakdowns, risk analysis, expiring units")
	reportCmd.Flags().StringVarP(&reportOutput, "output", "o", "", "Write report to file instead of stdout")
}

func detectRepoName(root string) string {
	cmd := exec.Command("git", "remote", "get-url", "origin")
	cmd.Dir = root
	out, err := cmd.Output()
	if err != nil {
		return ""
	}
	url := strings.TrimSpace(string(out))
	// Extract owner/repo from git URL
	url = strings.TrimSuffix(url, ".git")
	if idx := strings.LastIndex(url, ":"); idx >= 0 {
		return url[idx+1:]
	}
	if idx := strings.Index(url, "github.com/"); idx >= 0 {
		return url[idx+len("github.com/"):]
	}
	return url
}

func detectCommit(root string) string {
	cmd := exec.Command("git", "rev-parse", "--short", "HEAD")
	cmd.Dir = root
	out, err := cmd.Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}
