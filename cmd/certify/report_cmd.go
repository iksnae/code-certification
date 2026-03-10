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
	"github.com/iksnae/code-certification/internal/workspace"
	"github.com/spf13/cobra"
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
		root, _ := cmd.Flags().GetString("path")
		if root == "" {
			var err error
			root, err = os.Getwd()
			if err != nil {
				return fmt.Errorf("getting working directory: %w", err)
			}
		}

		wsMode, _ := cmd.Flags().GetBool("workspace")
		if wsMode {
			return runWorkspaceReport(root)
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

		// Read format flags
		reportFormat, _ := cmd.Flags().GetString("format")
		reportSite, _ := cmd.Flags().GetBool("site")
		reportDetailed, _ := cmd.Flags().GetBool("detailed")
		reportOutput, _ := cmd.Flags().GetString("output")

		// --site flag overrides format
		if reportSite {
			reportFormat = "site"
		}

		// Generate FullReport once — used by all formats except "text"
		needsFullReport := reportFormat != "text"
		var fr report.FullReport
		if needsFullReport {
			fr = report.GenerateFullReport(records, repo, commit, now)
		}

		var output string

		switch reportFormat {
		case "site":
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

		case "json":
			data, err := report.FormatJSON(fr)
			if err != nil {
				return err
			}
			output = string(data)

		case "full":
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
			output = report.FormatCardText(fr.Card)
		}

		// Write to file or stdout
		if output != "" {
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
		}

		// Always save report artifacts (REPORT_CARD.md, badge.json, per-unit reports)
		if needsFullReport {
			engine.SaveReportArtifacts(certDir, fr)
		} else {
			engine.SaveReportArtifactsFromStore(certDir, store, repo, commit, now)
		}

		cardPath := filepath.Join(certDir, "REPORT_CARD.md")
		badgePath := filepath.Join(certDir, "badge.json")
		fmt.Printf("\n✓ %s updated\n✓ %s updated\n", cardPath, badgePath)

		return nil
	},
}

func bindReportFlags() {
	reportCmd.Flags().StringP("format", "f", "card", "Output format (card, full, json, text)")
	reportCmd.Flags().String("path", "", "Path to repository (default: current directory)")
	reportCmd.Flags().Bool("detailed", false, "Include dimension breakdowns, risk analysis, expiring units")
	reportCmd.Flags().StringP("output", "o", "", "Write report to file instead of stdout")
	reportCmd.Flags().Bool("badge", false, "Print the shields.io badge markdown for your README")
	reportCmd.Flags().Bool("site", false, "Generate a static HTML site (shorthand for --format site)")
}

func runWorkspaceReport(root string) error {
	subs, err := workspace.DiscoverSubmodules(root)
	if err != nil {
		return fmt.Errorf("discovering submodules: %w", err)
	}

	if len(subs) == 0 {
		return fmt.Errorf("no git submodules found in %s", root)
	}

	fmt.Printf("🔍 Workspace report: %d submodule(s)\n\n", len(subs))

	// Build submodule summaries
	var summaries []workspace.SubmoduleSummary
	for _, s := range subs {
		summary := workspace.SubmoduleSummary{
			Name:       s.Name,
			Path:       s.Path,
			HasCertify: s.HasConfig,
			Commit:     s.Commit,
		}

		if s.HasConfig {
			card, err := workspace.LoadSubmoduleCard(root, s)
			if err != nil {
				fmt.Fprintf(os.Stderr, "  warning: loading %s: %v\n", s.Path, err)
			} else if card != nil {
				summary.Grade = card.OverallGrade
				summary.Score = card.OverallScore
				summary.Units = card.TotalUnits
				summary.Passing = card.Passing
				summary.Failing = card.Failing
				summary.PassRate = card.PassRate
				summary.StateAt = card.GeneratedAt
			}
		}

		summaries = append(summaries, summary)
	}

	// Aggregate
	wc := workspace.AggregateCards(summaries)

	// Create workspace .certification directory
	certDir := filepath.Join(root, ".certification")
	if err := os.MkdirAll(certDir, 0o755); err != nil {
		return fmt.Errorf("creating workspace .certification: %w", err)
	}

	// Write workspace REPORT_CARD.md
	md := workspace.FormatWorkspaceCardMarkdown(wc)
	cardPath := filepath.Join(certDir, "REPORT_CARD.md")
	if err := os.WriteFile(cardPath, []byte(md), 0o644); err != nil {
		return fmt.Errorf("writing workspace REPORT_CARD.md: %w", err)
	}

	// Write workspace report tree
	reportsDir := filepath.Join(certDir, "reports")
	count, err := workspace.GenerateWorkspaceReportTree(wc, reportsDir)
	if err != nil {
		return fmt.Errorf("generating workspace report tree: %w", err)
	}

	// Print summary
	emoji := "🟢"
	switch {
	case wc.OverallGrade == "F":
		emoji = "🔴"
	case wc.OverallGrade == "D":
		emoji = "🟠"
	case wc.OverallGrade == "C":
		emoji = "🟡"
	}

	fmt.Printf("  %s Workspace: %s (%.1f%%)\n", emoji, wc.OverallGrade, wc.OverallScore*100)
	fmt.Printf("  Units: %d · Passing: %d · Failing: %d\n\n", wc.TotalUnits, wc.TotalPassing, wc.TotalFailing)

	for _, s := range summaries {
		if !s.HasCertify {
			fmt.Printf("  %-30s  — not configured\n", s.Name)
		} else if s.Units == 0 {
			fmt.Printf("  %-30s  — no data\n", s.Name)
		} else {
			fmt.Printf("  %-30s  %s %-3s %5.1f%%  %d units\n",
				s.Name, gradeEmojiShort(s.Grade), s.Grade, s.Score*100, s.Units)
		}
	}

	fmt.Printf("\n✓ %s updated\n", cardPath)
	fmt.Printf("✓ %d report files written to %s\n", count, reportsDir)

	return nil
}

func gradeEmojiShort(grade string) string {
	switch grade {
	case "A", "A-", "B+", "B":
		return "🟢"
	case "C":
		return "🟡"
	case "D":
		return "🟠"
	case "F":
		return "🔴"
	default:
		return "⚪"
	}
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
