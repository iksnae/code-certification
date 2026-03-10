package report

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/iksnae/code-certification/internal/agent"
)

// FormatArchitectReport renders the architectural review as a markdown document.
// Part I (snapshot) is deterministic from data. Part II uses LLM analysis.
// Part III formats comparative recommendations with delta tables.
func FormatArchitectReport(result *agent.ArchitectResult, pc *agent.ProjectContext) string {
	var b strings.Builder

	writeArchHeader(&b, result, pc)
	writeArchExecutiveSummary(&b, result)
	writeArchPartI(&b, result, pc)
	writeArchPartII(&b, result)
	writeArchPartIII(&b, result)
	writeArchRiskMatrix(&b, result)
	writeArchRoadmap(&b, result)
	writeArchErrors(&b, result)
	writeArchThinking(&b, result)
	writeArchAppendix(&b, result, pc)

	return b.String()
}

func writeArchHeader(b *strings.Builder, result *agent.ArchitectResult, pc *agent.ProjectContext) {
	repo := pc.RepoName
	if repo == "" {
		repo = "Unknown"
	}
	fmt.Fprintf(b, "# 🏗 Architectural Review — %s\n\n", repo)
	fmt.Fprintf(b, "**Generated:** %s", time.Now().Format("2006-01-02 15:04"))
	if pc.CommitSHA != "" {
		fmt.Fprintf(b, " · **Commit:** `%s`", pc.CommitSHA)
	}
	if result.Model != "" {
		fmt.Fprintf(b, " · **Model:** `%s`", result.Model)
	}
	if result.TotalTokens > 0 {
		fmt.Fprintf(b, " · **Tokens:** %d", result.TotalTokens)
	}
	if result.Duration > 0 {
		fmt.Fprintf(b, " · **Duration:** %s", result.Duration.Round(time.Second))
	}
	fmt.Fprintf(b, " · **Phases:** %d/6\n\n", result.PhasesComplete)
}

func writeArchExecutiveSummary(b *strings.Builder, result *agent.ArchitectResult) {
	b.WriteString("## Executive Summary\n\n")
	if result.Phase6 != nil && result.Phase6.ExecutiveSummary != "" {
		b.WriteString(result.Phase6.ExecutiveSummary)
		b.WriteString("\n\n")
	} else {
		b.WriteString("*Executive summary not available (synthesis phase did not complete).*\n\n")
	}
}

func writeArchPartI(b *strings.Builder, result *agent.ArchitectResult, pc *agent.ProjectContext) {
	b.WriteString("---\n\n## Part I: Architecture Snapshot (As-Is)\n\n")

	snap := result.Snapshot
	if snap == nil {
		b.WriteString("*No snapshot data available.*\n\n")
		return
	}

	// Package Map (deterministic from snapshot data)
	if len(snap.Packages) > 0 {
		b.WriteString("### Package Map\n\n")
		b.WriteString("| Package | Units | Avg Score | Grade | Observations | Top Issues |\n")
		b.WriteString("|---------|------:|----------:|:-----:|-------------:|------------|\n")
		for _, pkg := range snap.Packages {
			issues := "-"
			if len(pkg.TopIssues) > 0 {
				issues = strings.Join(pkg.TopIssues, ", ")
			}
			fmt.Fprintf(b, "| %s | %d | %.1f%% | %s | %d | %s |\n",
				pkg.Path, pkg.Units, pkg.AvgScore*100, pkg.Grade, pkg.Observations, issues)
		}
		b.WriteString("\n")
	}

	// Dependency Graph (deterministic)
	if len(snap.DependencyEdges) > 0 {
		b.WriteString("### Dependency Graph\n\n```\n")
		depMap := make(map[string][]string)
		for _, e := range snap.DependencyEdges {
			depMap[e.From] = append(depMap[e.From], e.To)
		}
		var fromPkgs []string
		for k := range depMap {
			fromPkgs = append(fromPkgs, k)
		}
		sort.Strings(fromPkgs)
		for _, from := range fromPkgs {
			fmt.Fprintf(b, "%s → [%s]\n", from, strings.Join(depMap[from], ", "))
		}
		b.WriteString("```\n\n")
	}

	// Layer Structure (phase 1 narrative + snapshot layers)
	b.WriteString("### Layer Structure\n\n")
	if len(snap.Layers) > 0 {
		layerPkgs := make(map[string][]string)
		for pkg, layer := range snap.Layers {
			layerPkgs[layer] = append(layerPkgs[layer], pkg)
		}
		for _, layer := range []string{"cmd", "domain", "internal", "pkg", "other"} {
			pkgs := layerPkgs[layer]
			if len(pkgs) == 0 {
				continue
			}
			sort.Strings(pkgs)
			fmt.Fprintf(b, "- **%s:** %s\n", layer, strings.Join(pkgs, ", "))
		}
		b.WriteString("\n")
	}
	if result.Phase1 != nil {
		if result.Phase1.DependencyAssessment != "" {
			b.WriteString(result.Phase1.DependencyAssessment)
			b.WriteString("\n\n")
		}
		for _, layer := range result.Phase1.Layers {
			fmt.Fprintf(b, "**%s** (%s): %s\n\n",
				layer.Name, strings.Join(layer.Packages, ", "), layer.Description)
		}
		for _, flow := range result.Phase1.DataFlows {
			fmt.Fprintf(b, "- `%s` → `%s`: %s\n", flow.From, flow.To, flow.Description)
		}
		if len(result.Phase1.DataFlows) > 0 {
			b.WriteString("\n")
		}
	}

	// Hotspots (deterministic)
	if len(snap.Hotspots) > 0 {
		b.WriteString("### Hotspots\n\n")
		b.WriteString("| Rank | Package | Units | Score | Risk Factor |\n")
		b.WriteString("|-----:|---------|------:|------:|------------:|\n")
		for i, h := range snap.Hotspots {
			risk := float64(h.Units) * (1.0 - h.AvgScore)
			fmt.Fprintf(b, "| %d | %s | %d | %.1f%% | %.2f |\n",
				i+1, h.Path, h.Units, h.AvgScore*100, risk)
		}
		b.WriteString("\n")
	}

	// Coupling Analysis (deterministic + phase 1 narrative)
	if len(snap.CouplingPairs) > 0 {
		b.WriteString("### Coupling Analysis\n\n")
		b.WriteString("| Package A | Package B | Edges |\n")
		b.WriteString("|-----------|-----------|------:|\n")
		for _, cp := range snap.CouplingPairs {
			fmt.Fprintf(b, "| %s | %s | %d |\n", cp.PkgA, cp.PkgB, cp.EdgeCount)
		}
		b.WriteString("\n")
	}
}

func writeArchPartII(b *strings.Builder, result *agent.ArchitectResult) {
	b.WriteString("---\n\n## Part II: Analysis\n\n")

	// Code Quality & Patterns
	b.WriteString("### Code Quality & Patterns\n\n")
	if result.Phase2 != nil && len(result.Phase2.Findings) > 0 {
		for _, f := range result.Phase2.Findings {
			severity := f.Severity
			if severity == "" {
				severity = "medium"
			}
			emoji := severityEmoji(severity)
			fmt.Fprintf(b, "%s **%s** — %s\n\n", emoji, f.Package, f.Issue)
		}
	} else {
		b.WriteString("*No code quality findings (phase did not complete).*\n\n")
	}

	// Test Strategy & Coverage
	b.WriteString("### Test Strategy & Coverage\n\n")
	if result.Phase3 != nil {
		if result.Phase3.StrategyAssessment != "" {
			b.WriteString(result.Phase3.StrategyAssessment)
			b.WriteString("\n\n")
		}
		if len(result.Phase3.CoverageGaps) > 0 {
			b.WriteString("**Coverage Gaps:**\n\n")
			for _, g := range result.Phase3.CoverageGaps {
				fmt.Fprintf(b, "- `%s` (score: %.1f%%): %s\n", g.Package, g.CurrentScore*100, g.Issue)
			}
			b.WriteString("\n")
		}
	} else {
		b.WriteString("*No test strategy assessment (phase did not complete).*\n\n")
	}

	// Security & Operations
	b.WriteString("### Security & Operations\n\n")
	if result.Phase4 != nil && len(result.Phase4.Concerns) > 0 {
		for _, c := range result.Phase4.Concerns {
			emoji := "🔒"
			switch c.Area {
			case "operations":
				emoji = "⚙️"
			case "config":
				emoji = "📋"
			case "dependencies":
				emoji = "📦"
			}
			fmt.Fprintf(b, "%s **%s** — %s\n", emoji, c.Area, c.Description)
			if len(c.AffectedPackages) > 0 {
				fmt.Fprintf(b, "  Affected: `%s`\n", strings.Join(c.AffectedPackages, "`, `"))
			}
			if len(c.Metrics) > 0 {
				var metricParts []string
				for k, v := range c.Metrics {
					metricParts = append(metricParts, fmt.Sprintf("%s: %v", k, v))
				}
				sort.Strings(metricParts)
				fmt.Fprintf(b, "  Metrics: %s\n", strings.Join(metricParts, ", "))
			}
			b.WriteString("\n")
		}
	} else {
		b.WriteString("*No security/operations concerns (phase did not complete).*\n\n")
	}
}

func writeArchPartIII(b *strings.Builder, result *agent.ArchitectResult) {
	b.WriteString("---\n\n## Part III: Recommendations (Current → Proposed)\n\n")

	if result.Phase5 == nil || len(result.Phase5.Recommendations) == 0 {
		b.WriteString("*No recommendations generated (phase did not complete).*\n\n")
		return
	}

	for _, rec := range result.Phase5.Recommendations {
		fmt.Fprintf(b, "### %s\n\n", rec.Title)

		// Delta table
		if len(rec.Deltas) > 0 {
			b.WriteString("| Metric | Current | Projected | Delta |\n")
			b.WriteString("|--------|--------:|----------:|------:|\n")
			for _, d := range rec.Deltas {
				delta := computeDeltaDisplay(d.Current, d.Projected)
				fmt.Fprintf(b, "| %s | %s | %s | %s |\n", d.Metric, d.Current, d.Projected, delta)
			}
			b.WriteString("\n")
		}

		if rec.CurrentState != "" {
			fmt.Fprintf(b, "**Current:** %s\n\n", rec.CurrentState)
		}
		if rec.ProposedState != "" {
			fmt.Fprintf(b, "**Proposed:** %s\n\n", rec.ProposedState)
		}
		if len(rec.AffectedUnits) > 0 {
			fmt.Fprintf(b, "**Affected:** `%s`\n\n", strings.Join(rec.AffectedUnits, "`, `"))
		}
		if rec.Effort != "" {
			fmt.Fprintf(b, "**Effort:** %s", rec.Effort)
		}
		if rec.Justification != "" {
			fmt.Fprintf(b, " · **Justification:** %s", rec.Justification)
		}
		b.WriteString("\n\n")
	}
}

func writeArchRiskMatrix(b *strings.Builder, result *agent.ArchitectResult) {
	b.WriteString("---\n\n## Risk Matrix\n\n")

	if result.Phase6 == nil || len(result.Phase6.RiskMatrix) == 0 {
		b.WriteString("*No risk matrix (synthesis phase did not complete).*\n\n")
		return
	}

	b.WriteString("| Risk | Severity | Likelihood | Related Recommendation |\n")
	b.WriteString("|------|----------|------------|------------------------|\n")
	for _, r := range result.Phase6.RiskMatrix {
		ref := r.RecommendationRef
		if ref == "" {
			ref = "-"
		}
		fmt.Fprintf(b, "| %s | %s | %s | %s |\n", r.Risk, r.Severity, r.Likelihood, ref)
	}
	b.WriteString("\n")
}

func writeArchRoadmap(b *strings.Builder, result *agent.ArchitectResult) {
	b.WriteString("## Prioritized Roadmap\n\n")

	if result.Phase6 == nil || len(result.Phase6.Roadmap) == 0 {
		b.WriteString("*No roadmap (synthesis phase did not complete).*\n\n")
		return
	}

	b.WriteString("| # | Item | Effort | Impact | Current → Projected |\n")
	b.WriteString("|--:|------|--------|--------|---------------------|\n")
	for _, item := range result.Phase6.Roadmap {
		delta := item.DeltaSummary
		if delta == "" {
			delta = "-"
		}
		fmt.Fprintf(b, "| %d | %s | %s | %s | %s |\n",
			item.Priority, item.Title, item.Effort, item.Impact, delta)
	}
	b.WriteString("\n")
}

func writeArchErrors(b *strings.Builder, result *agent.ArchitectResult) {
	if len(result.Errors) == 0 {
		return
	}

	b.WriteString("## ⚠️ Incomplete Phases\n\n")
	for _, e := range result.Errors {
		fmt.Fprintf(b, "- %s\n", e)
	}
	b.WriteString("\n")
}

func writeArchThinking(b *strings.Builder, result *agent.ArchitectResult) {
	// Collect all non-empty thinking blocks
	var hasThinking bool
	for _, t := range result.Thinking {
		if t != "" {
			hasThinking = true
			break
		}
	}
	if !hasThinking {
		return
	}

	phaseNames := agent.ArchitectPhaseNames()

	b.WriteString("## 🧠 Agent Reasoning\n\n")
	for i, t := range result.Thinking {
		if t == "" {
			continue
		}
		name := "Unknown"
		if i < len(phaseNames) {
			name = phaseNames[i]
		}
		fmt.Fprintf(b, "<details>\n<summary>Phase %d: %s</summary>\n\n", i+1, name)
		b.WriteString(t)
		b.WriteString("\n\n</details>\n\n")
	}
}

func writeArchAppendix(b *strings.Builder, result *agent.ArchitectResult, pc *agent.ProjectContext) {
	b.WriteString("---\n\n## Appendix: Data Sources\n\n")

	snap := result.Snapshot
	if snap != nil {
		fmt.Fprintf(b, "- **%d** units across **%d** packages", snap.Metrics.TotalUnits, snap.Metrics.TotalPackages)
		if snap.Metrics.AvgScore > 0 {
			grade := "N/A"
			for _, g := range []string{"A", "A-", "B+", "B", "C", "D", "F"} {
				if snap.Metrics.GradeDistribution[g] > 0 {
					grade = g // just use first non-zero as representative
					break
				}
			}
			_ = grade
			fmt.Fprintf(b, " · Score: %.1f%%", snap.Metrics.AvgScore*100)
		}
		b.WriteString("\n")
	}
	b.WriteString("- Evidence: lint, test, coverage, structural, git history\n")
	if pc.CommitSHA != "" {
		fmt.Fprintf(b, "- Snapshot computed from certification records at `%s`\n", pc.CommitSHA)
	}
	b.WriteString("\n---\n\n*Generated by [Certify](https://github.com/iksnae/code-certification) `architect` command.*\n")
}

// computeDeltaDisplay computes a human-readable delta string.
func computeDeltaDisplay(current, projected string) string {
	if current == "unknown" || projected == "unknown" {
		return "—"
	}
	// Try to show as difference if both are simple values
	return fmt.Sprintf("%s → %s", current, projected)
}

func severityEmoji(severity string) string {
	switch severity {
	case "critical":
		return "🔴"
	case "high":
		return "🟠"
	case "medium":
		return "🟡"
	case "low":
		return "🟢"
	default:
		return "⚪"
	}
}
