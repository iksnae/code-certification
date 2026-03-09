package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/iksnae/code-certification/internal/domain"
	"github.com/iksnae/code-certification/internal/record"
	"github.com/iksnae/code-certification/internal/report"
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
  text      Quick summary to terminal (default)
  card      Visual report card box for terminal
  json      Full report as machine-readable JSON
  full      Complete per-unit report card (markdown)

Every run generates:
  .certification/REPORT_CARD.md  — full per-unit report card
  .certification/badge.json      — shields.io endpoint badge

Badge for your README:
  certify report --badge    prints the markdown snippet`,
	RunE: func(cmd *cobra.Command, args []string) error {
		root := reportPath
		if root == "" {
			root, _ = os.Getwd()
		}
		certDir := filepath.Join(root, ".certification")

		// --badge flag just prints the snippet
		if showBadge, _ := cmd.Flags().GetBool("badge"); showBadge {
			repo := detectRepoName(root)
			if repo == "" {
				return fmt.Errorf("could not detect repository name (no git remote)")
			}
			fmt.Println(report.BadgeMarkdown(repo, "main"))
			return nil
		}

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

		case "full":
			fr := report.GenerateFullReport(records, repo, commit, now)
			output = report.FormatFullMarkdown(fr)

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

		// Always save the full report card + badge
		saveReportCard(certDir, records, repo, commit, now)
		saveBadge(certDir, records, repo, commit, now)

		return nil
	},
}

func init() {
	reportCmd.Flags().StringVarP(&reportFormat, "format", "f", "text", "Output format (text, json, card, full)")
	reportCmd.Flags().StringVar(&reportPath, "path", "", "Path to repository (default: current directory)")
	reportCmd.Flags().BoolVar(&reportDetailed, "detailed", false, "Include dimension breakdowns, risk analysis, expiring units")
	reportCmd.Flags().StringVarP(&reportOutput, "output", "o", "", "Write report to file instead of stdout")
	reportCmd.Flags().Bool("badge", false, "Print the shields.io badge markdown for your README")
}

func saveReportCard(certDir string, records []domain.CertificationRecord, repo, commit string, now time.Time) {
	fr := report.GenerateFullReport(records, repo, commit, now)
	md := report.FormatFullMarkdown(fr)
	os.WriteFile(filepath.Join(certDir, "REPORT_CARD.md"), []byte(md), 0o644)
}

func saveBadge(certDir string, records []domain.CertificationRecord, repo, commit string, now time.Time) {
	c := report.GenerateCard(records, repo, commit, now)
	badge := report.GenerateBadge(c)
	data, err := report.FormatBadgeJSON(badge)
	if err != nil {
		return
	}
	os.WriteFile(filepath.Join(certDir, "badge.json"), data, 0o644)
}

func detectRepoName(root string) string {
	cmd := exec.Command("git", "remote", "get-url", "origin")
	cmd.Dir = root
	out, err := cmd.Output()
	if err != nil {
		return ""
	}
	url := strings.TrimSpace(string(out))
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
