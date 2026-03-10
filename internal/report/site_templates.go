package report

import (
	"fmt"
	"html/template"
	"strings"
)

// siteCSS is the embedded stylesheet for the static site.
const siteCSS = `
:root {
  --bg: #ffffff; --fg: #1a1a2e; --muted: #6b7280;
  --border: #e5e7eb; --card-bg: #f9fafb; --hover: #f3f4f6;
  --grade-a: #2E8B57; --grade-a-minus: #3DA06A; --grade-b-plus: #4A6B82;
  --grade-b: #4A6B82; --grade-c: #E0A100; --grade-d: #F59E0B; --grade-f: #DC2626;
  --expired: #9CA3AF;
}
@media (prefers-color-scheme: dark) {
  :root {
    --bg: #0f172a; --fg: #e2e8f0; --muted: #94a3b8;
    --border: #334155; --card-bg: #1e293b; --hover: #334155;
  }
}
*, *::before, *::after { box-sizing: border-box; margin: 0; padding: 0; }
body {
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif;
  background: var(--bg); color: var(--fg); line-height: 1.6;
  max-width: 1200px; margin: 0 auto; padding: 1rem;
}
a { color: var(--grade-b-plus); text-decoration: none; }
a:hover { text-decoration: underline; }
h1 { font-size: 1.8rem; margin-bottom: 0.5rem; }
h2 { font-size: 1.4rem; margin: 1.5rem 0 0.75rem; border-bottom: 1px solid var(--border); padding-bottom: 0.25rem; }
h3 { font-size: 1.1rem; margin: 1rem 0 0.5rem; }
.breadcrumb { font-size: 0.9rem; color: var(--muted); margin-bottom: 1rem; }
.breadcrumb a { color: var(--muted); }
.summary-grid { display: grid; grid-template-columns: repeat(auto-fit, minmax(200px, 1fr)); gap: 1rem; margin: 1rem 0; }
.stat-card {
  background: var(--card-bg); border: 1px solid var(--border); border-radius: 8px;
  padding: 1rem; text-align: center;
}
.stat-card .value { font-size: 2rem; font-weight: 700; }
.stat-card .label { font-size: 0.85rem; color: var(--muted); }
table { width: 100%; border-collapse: collapse; margin: 0.5rem 0 1.5rem; font-size: 0.9rem; }
th, td { padding: 0.5rem 0.75rem; text-align: left; border-bottom: 1px solid var(--border); }
th { font-weight: 600; background: var(--card-bg); position: sticky; top: 0; }
tr:hover { background: var(--hover); }
.grade { font-weight: 700; padding: 0.15rem 0.5rem; border-radius: 4px; display: inline-block; color: white; font-size: 0.85rem; }
.grade-a, .grade-a- { background: var(--grade-a); }
.grade-b\\+, .grade-b { background: var(--grade-b); }
.grade-c { background: var(--grade-c); color: var(--fg); }
.grade-d { background: var(--grade-d); color: var(--fg); }
.grade-f { background: var(--grade-f); }
.grade-na { background: var(--expired); }
.bar-container { background: var(--border); border-radius: 4px; height: 20px; width: 100%; overflow: hidden; }
.bar-fill { height: 100%; border-radius: 4px; transition: width 0.3s; }
.status-certified { color: var(--grade-a); }
.status-certified_with_observations { color: var(--grade-c); }
.status-probationary { color: var(--grade-d); }
.status-decertified { color: var(--grade-f); }
.status-expired { color: var(--expired); }
.status-exempt { color: var(--muted); }
.search-box { width: 100%; padding: 0.75rem; font-size: 1rem; border: 1px solid var(--border); border-radius: 8px; background: var(--card-bg); color: var(--fg); margin: 1rem 0; }
.search-box:focus { outline: 2px solid var(--grade-b-plus); }
.search-count { font-size: 0.85rem; color: var(--muted); margin-bottom: 0.5rem; }
.obs-list { list-style: none; padding: 0; }
.obs-list li { padding: 0.25rem 0; }
.obs-ai { color: var(--grade-b-plus); }
.obs-suggestion { color: var(--grade-a); }
.nav-links { display: flex; justify-content: space-between; margin: 1.5rem 0; font-size: 0.9rem; }
footer { margin-top: 2rem; padding-top: 1rem; border-top: 1px solid var(--border); font-size: 0.8rem; color: var(--muted); text-align: center; }
@media (max-width: 768px) {
  .summary-grid { grid-template-columns: repeat(2, 1fr); }
  table { font-size: 0.8rem; }
  th, td { padding: 0.35rem 0.5rem; }
}
`

// searchJS is the minimal client-side search script embedded in the index page.
const searchJS = `
document.addEventListener('DOMContentLoaded', function() {
  var input = document.getElementById('search-input');
  var table = document.getElementById('search-results');
  var count = document.getElementById('search-count');
  var allRows = document.getElementById('all-units-table');
  if (!input || !table || typeof SEARCH_INDEX === 'undefined') return;

  input.addEventListener('input', function() {
    var q = this.value.toLowerCase().trim();
    if (q.length === 0) {
      table.style.display = 'none';
      if (allRows) allRows.style.display = '';
      count.textContent = '';
      return;
    }
    if (allRows) allRows.style.display = 'none';
    var matches = SEARCH_INDEX.filter(function(e) {
      return e.n.toLowerCase().indexOf(q) >= 0 ||
             e.p.toLowerCase().indexOf(q) >= 0 ||
             e.id.toLowerCase().indexOf(q) >= 0 ||
             e.g.toLowerCase().indexOf(q) >= 0 ||
             e.s.toLowerCase().indexOf(q) >= 0 ||
             e.l.toLowerCase().indexOf(q) >= 0;
    });
    count.textContent = matches.length + ' of ' + SEARCH_INDEX.length + ' units';
    var html = '<table><tr><th>Unit</th><th>Path</th><th>Grade</th><th>Status</th><th>Language</th></tr>';
    matches.forEach(function(e) {
      html += '<tr><td><a href="' + e.uu + '">' + esc(e.n) + '</a></td>';
      html += '<td>' + esc(e.p) + '</td>';
      html += '<td><span class="grade grade-' + e.g.toLowerCase().replace('+','\\+') + '">' + esc(e.g) + '</span></td>';
      html += '<td class="status-' + e.s + '">' + esc(e.s) + '</td>';
      html += '<td>' + esc(e.l) + '</td></tr>';
    });
    html += '</table>';
    table.innerHTML = html;
    table.style.display = '';
  });
  function esc(s) {
    var d = document.createElement('div');
    d.appendChild(document.createTextNode(s));
    return d.innerHTML;
  }
});
`

// Template helper functions
var siteFuncMap = template.FuncMap{
	"gradeColor": func(grade string) string {
		return badgeColor(grade)
	},
	"gradeClass": func(grade string) string {
		return "grade-" + strings.ToLower(strings.ReplaceAll(grade, "+", "\\+"))
	},
	"gradeCSSClass": gradeCSSClass,
	"gradeClassSafe": func(grade string) string {
		g := strings.ToLower(grade)
		g = strings.ReplaceAll(g, "+", "plus")
		return "grade-" + g
	},
	"gradeEmoji": func(grade string) string {
		return gradeEmoji(grade)
	},
	"pct": func(score float64) string {
		return fmt.Sprintf("%.1f%%", score*100)
	},
	"pctWidth": func(score float64) string {
		return fmt.Sprintf("%.1f%%", score*100)
	},
	"scoreBar": func(score float64) template.HTML {
		color := badgeColor("B")
		if score >= 0.9 {
			color = badgeColor("A")
		} else if score >= 0.7 {
			color = badgeColor("B")
		} else if score >= 0.5 {
			color = badgeColor("C")
		} else {
			color = badgeColor("F")
		}
		return template.HTML(fmt.Sprintf(
			`<div class="bar-container"><div class="bar-fill" style="width:%.1f%%;background:#%s"></div></div>`,
			score*100, color))
	},
	"shortPath": func(path string) string {
		return shortFile(path)
	},
	"relPath": func(from, to string) string {
		// Compute relative path from one page to another
		// from: "packages/internal/engine/index.html"
		// to:   "units/foo.html"
		depth := strings.Count(from, "/")
		prefix := strings.Repeat("../", depth)
		return prefix + to
	},
	"sortedDimKeys": func(m map[string]float64) []string {
		return sortedKeys(m)
	},
	"hasPrefix": strings.HasPrefix,
}

// indexTemplate is the dashboard page template.
const indexTemplateStr = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>{{.Title}} — Certification Report</title>
<style>` + siteCSS + `</style>
</head>
<body>
<header>
<h1>{{.GradeEmoji}} {{.Title}} — Certification Report</h1>
<p style="color:var(--muted)">Commit: <code>{{.CommitSHA}}</code> · Generated: {{.GeneratedAt}}</p>
</header>

<div class="summary-grid">
<div class="stat-card"><div class="value">{{.OverallGrade}}</div><div class="label">Overall Grade</div></div>
<div class="stat-card"><div class="value">{{pct .OverallScore}}</div><div class="label">Overall Score</div></div>
<div class="stat-card"><div class="value">{{.TotalUnits}}</div><div class="label">Total Units</div></div>
<div class="stat-card"><div class="value">{{pct .PassRate}}</div><div class="label">Pass Rate</div></div>
<div class="stat-card"><div class="value">{{.Passing}}</div><div class="label">Passing</div></div>
<div class="stat-card"><div class="value">{{.Failing}}</div><div class="label">Failing</div></div>
</div>

{{if .HasGrades}}
<h2>Grade Distribution</h2>
<table>
<tr><th>Grade</th><th>Count</th><th>%</th><th>Bar</th></tr>
{{range .Grades}}<tr>
<td><span class="grade grade-{{.CSSClass}}">{{.Name}}</span></td>
<td>{{.Count}}</td><td>{{pct .Pct}}</td>
<td>{{scoreBar .Pct}}</td>
</tr>{{end}}
</table>
{{end}}

{{if .HasDimensions}}
<h2>Dimension Averages</h2>
<table>
<tr><th>Dimension</th><th>Score</th><th>Bar</th></tr>
{{range .Dimensions}}<tr>
<td>{{.Name}}</td><td>{{pct .Score}}</td><td>{{scoreBar .Score}}</td>
</tr>{{end}}
</table>
{{end}}

{{if .HasLanguages}}
<h2>By Language</h2>
<table>
<tr><th>Language</th><th>Units</th><th>Grade</th><th>Score</th></tr>
{{range .Languages}}<tr>
<td>{{.Name}}</td><td>{{.Units}}</td>
<td><span class="grade grade-{{gradeCSSClass .Grade}}">{{.Grade}}</span></td>
<td>{{pct .AverageScore}}</td>
</tr>{{end}}
</table>
{{end}}

{{if .HasPackages}}
<h2>Packages</h2>
<table>
<tr><th>Package</th><th>Units</th><th>Grade</th><th>Avg Score</th></tr>
{{range .Packages}}<tr>
<td><a href="packages/{{.Path}}/index.html">{{.Path}}/</a></td>
<td>{{.Units}}</td>
<td><span class="grade grade-{{.CSSClass}}">{{.Grade}}</span></td>
<td>{{pct .AvgScore}}</td>
</tr>{{end}}
</table>
{{end}}

{{if .HasTopIssues}}
<h2>Top Issues</h2>
<table>
<tr><th>Unit</th><th>Grade</th><th>Score</th><th>Issue</th></tr>
{{range .TopIssues}}<tr>
<td><a href="units/{{.Anchor}}.html">{{.Name}}</a></td>
<td><span class="grade grade-{{.CSSClass}}">{{.Grade}}</span></td>
<td>{{pct .Score}}</td><td>{{.Reason}}</td>
</tr>{{end}}
</table>
{{end}}

{{if .IncludeSearch}}
<h2>Search</h2>
<input type="text" id="search-input" class="search-box" placeholder="Filter by name, path, grade, status, or language…">
<div id="search-count" class="search-count"></div>
<div id="search-results" style="display:none"></div>
<script src="search-index.js"></script>
<script>` + searchJS + `</script>
{{end}}

<footer>
<p>{{.TotalUnits}} units certified · Generated by <a href="https://github.com/iksnae/code-certification">Certify</a></p>
{{if .ReportCardLink}}<p><a href="../REPORT_CARD.md">Markdown Report</a> · <a href="../badge.json">Badge JSON</a></p>{{end}}
</footer>
</body>
</html>`

// packageTemplateStr is the per-package page template.
const packageTemplateStr = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>{{.PackagePath}}/ — {{.Title}}</title>
<style>` + siteCSS + `</style>
</head>
<body>
<div class="breadcrumb">
<a href="{{.IndexURL}}">{{.Title}}</a> / <strong>{{.PackagePath}}/</strong>
</div>

<h1>{{.GradeEmoji}} {{.PackagePath}}/</h1>

<div class="summary-grid">
<div class="stat-card"><div class="value">{{.Grade}}</div><div class="label">Grade</div></div>
<div class="stat-card"><div class="value">{{pct .AvgScore}}</div><div class="label">Avg Score</div></div>
<div class="stat-card"><div class="value">{{.UnitCount}}</div><div class="label">Units</div></div>
<div class="stat-card"><div class="value">{{pct .PassRate}}</div><div class="label">Pass Rate</div></div>
</div>

<h2>Units</h2>
<table>
<tr><th>Unit</th><th>Type</th><th>Grade</th><th>Score</th><th>Status</th><th>Expires</th></tr>
{{range .Units}}<tr>
<td><a href="{{.UnitURL}}">{{.Name}}</a></td>
<td>{{.UnitType}}</td>
<td><span class="grade grade-{{.CSSClass}}">{{.Grade}}</span></td>
<td>{{pct .Score}}</td>
<td class="status-{{.Status}}">{{.Status}}</td>
<td>{{.ExpiresAt}}</td>
</tr>{{end}}
</table>

<footer>
<p>{{.UnitCount}} units · <a href="{{.IndexURL}}">← Back to Dashboard</a></p>
<p>Generated by <a href="https://github.com/iksnae/code-certification">Certify</a></p>
</footer>
</body>
</html>`

// unitTemplateStr is the per-unit detail page template.
const unitTemplateStr = `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>{{.Name}} — {{.Title}}</title>
<style>` + siteCSS + `</style>
</head>
<body>
<div class="breadcrumb">
<a href="{{.IndexURL}}">{{.Title}}</a> / <a href="{{.PackageURL}}">{{.PackagePath}}/</a> / <strong>{{.Name}}</strong>
</div>

<h1>{{.GradeEmoji}} {{.Name}}</h1>

<h2>Identity</h2>
<table>
<tr><th>Field</th><th>Value</th></tr>
<tr><td>Unit ID</td><td><code>{{.UnitID}}</code></td></tr>
<tr><td>Type</td><td>{{.UnitType}}</td></tr>
<tr><td>Path</td><td><code>{{.Path}}</code></td></tr>
<tr><td>Language</td><td>{{.Language}}</td></tr>
{{if .Symbol}}<tr><td>Symbol</td><td><code>{{.Symbol}}</code></td></tr>{{end}}
</table>

<h2>Certification</h2>
<table>
<tr><th>Field</th><th>Value</th></tr>
<tr><td>Grade</td><td><span class="grade grade-{{.CSSClass}}">{{.Grade}}</span></td></tr>
<tr><td>Score</td><td>{{pct .Score}}</td></tr>
<tr><td>Status</td><td class="status-{{.Status}}">{{.Status}}</td></tr>
<tr><td>Confidence</td><td>{{pct .Confidence}}</td></tr>
<tr><td>Certified</td><td>{{.CertifiedAt}}</td></tr>
<tr><td>Expires</td><td>{{.ExpiresAt}}</td></tr>
<tr><td>Source</td><td><code>{{.Source}}</code></td></tr>
</table>

{{if .HasDimensions}}
<h2>Dimension Scores</h2>
<table>
<tr><th>Dimension</th><th>Score</th><th>Bar</th></tr>
{{range .Dimensions}}<tr>
<td>{{.Name}}</td><td>{{pct .Score}}</td><td>{{scoreBar .Score}}</td>
</tr>{{end}}
</table>
{{end}}

{{if .HasAIObservations}}
<h2>🤖 AI Assessment</h2>
{{range .AIObservations}}<p class="obs-ai">{{.}}</p>{{end}}
{{if .HasSuggestions}}
<h3>Suggestions</h3>
<ul class="obs-list">
{{range .Suggestions}}<li class="obs-suggestion">{{.}}</li>{{end}}
</ul>
{{end}}
{{end}}

{{if .HasOtherObservations}}
<h2>Observations</h2>
<ul class="obs-list">
{{range .OtherObservations}}<li>{{.}}</li>{{end}}
</ul>
{{end}}

{{if .HasActions}}
<h2>Required Actions</h2>
<ul>
{{range .Actions}}<li>{{.}}</li>{{end}}
</ul>
{{end}}

<div class="nav-links">
{{if .PrevURL}}<a href="{{.PrevURL}}">← {{.PrevName}}</a>{{else}}<span></span>{{end}}
{{if .NextURL}}<a href="{{.NextURL}}">{{.NextName}} →</a>{{else}}<span></span>{{end}}
</div>

<footer>
<p><a href="{{.PackageURL}}">← Back to {{.PackagePath}}/</a> · <a href="{{.IndexURL}}">Dashboard</a></p>
<p>Generated by <a href="https://github.com/iksnae/code-certification">Certify</a></p>
</footer>
</body>
</html>`
