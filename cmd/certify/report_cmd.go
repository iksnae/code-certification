package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/iksnae/code-certification/internal/engine"
	"github.com/iksnae/code-certification/internal/record"
	"github.com/iksnae/code-certification/internal/report"
	"github.com/spf13/cobra"
)

var (
	reportFormat   string
	reportPath     string
	reportDetailed bool
	reportOutput   string
	reportSite     bool
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Generate certification reports",
	Long: `Generate certification reports.

Every run updates:
  .certification/REPORT_CARD.md  — full per-unit report card
  .certification/badge.json      — shields.io endpoint badge

Formats:
  card      Visual report card in terminal (default)
  full      Complete per-unit report card (markdown)
  json      Machine-readable full report
  text      Brief health summary
  site      Static HTML site (browsable, searchable, works offline)

Badge for your README:
  certify report --badge

Static site (for large repos):
  certify report --format site
  certify report --site`,
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

		snapshotPath := filepath.Join(certDir, "state.json")
		store := record.NewStoreWithSnapshot(filepath.Join(certDir, "records"), snapshotPath)
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

		// --site flag overrides format
		if reportSite {
			reportFormat = "site"
		}

		var output string

		switch reportFormat {
		case "site":
			fr := report.GenerateFullReport(records, repo, commit, now)
			siteDir := filepath.Join(certDir, "site")
			cfg := report.SiteConfig{
				OutputDir:     siteDir,
				Title:         repo,
				IncludeSearch: true,
			}
			if cfg.Title == "" {
				cfg.Title = "Certification Report"
			}
			if err := report.GenerateSite(fr, cfg); err != nil {
				return fmt.Errorf("generating site: %w", err)
			}
			// Count generated pages
			pageCount := len(fr.Units) + 1 // units + index
			dirs := make(map[string]bool)
			for _, u := range fr.Units {
				dirs[reportDirOf(u.Path)] = true
			}
			pageCount += len(dirs) // package pages
			fmt.Printf("✓ Static site generated → %s (%d pages)\n", siteDir, pageCount)
			fmt.Println("  Open .certification/site/index.html in a browser")

			// Still save markdown report + badge via consolidated function
			engine.SaveReportArtifacts(certDir, store, repo, commit, now)

			cardPath := filepath.Join(certDir, "REPORT_CARD.md")
			badgePath := filepath.Join(certDir, "badge.json")
			fmt.Printf("\n✓ %s updated\n✓ %s updated\n", cardPath, badgePath)
			return nil

		case "json":
			fr := report.GenerateFullReport(records, repo, commit, now)
			data, err := report.FormatJSON(fr)
			if err != nil {
				return err
			}
			output = string(data)

		case "full":
			fr := report.GenerateFullReport(records, repo, commit, now)
			output = report.FormatFullMarkdown(fr)

		case "text":
			if reportDetailed {
				d := report.Detailed(records, now)
				output = report.FormatDetailedText(d)
			} else {
				h := report.Health(records)
				output = report.FormatText(h)
			}

		default: // "card" or unrecognized
			c := report.GenerateCard(records, repo, commit, now)
			output = report.FormatCardText(c)
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

		// Always save the full report card + badge via consolidated function
		engine.SaveReportArtifacts(certDir, store, repo, commit, now)

		cardPath := filepath.Join(certDir, "REPORT_CARD.md")
		badgePath := filepath.Join(certDir, "badge.json")
		fmt.Printf("\n✓ %s updated\n✓ %s updated\n", cardPath, badgePath)

		return nil
	},
}

func init() {
	reportCmd.Flags().StringVarP(&reportFormat, "format", "f", "card", "Output format (card, full, json, text)")
	reportCmd.Flags().StringVar(&reportPath, "path", "", "Path to repository (default: current directory)")
	reportCmd.Flags().BoolVar(&reportDetailed, "detailed", false, "Include dimension breakdowns, risk analysis, expiring units")
	reportCmd.Flags().StringVarP(&reportOutput, "output", "o", "", "Write report to file instead of stdout")
	reportCmd.Flags().Bool("badge", false, "Print the shields.io badge markdown for your README")
	reportCmd.Flags().BoolVar(&reportSite, "site", false, "Generate a static HTML site (shorthand for --format site)")
}

func reportDirOf(path string) string {
	idx := strings.LastIndex(path, "/")
	if idx < 0 {
		return "."
	}
	return path[:idx]
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

	if idx := strings.Index(url, "github.com/"); idx >= 0 {
		return url[idx+len("github.com/"):]
	}
	if idx := strings.LastIndex(url, ":"); idx >= 0 && !strings.Contains(url[:idx], "//") {
		return url[idx+1:]
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
