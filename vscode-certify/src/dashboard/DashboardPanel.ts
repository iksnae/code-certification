import * as vscode from 'vscode';
import type { CertifyDataLoader } from '../dataLoader.js';
import type { FullReport } from '../types.js';
import { GRADE_COLORS, GRADE_EMOJI, DIMENSION_NAMES, BRAND_COLORS } from '../constants.js';

export class DashboardPanel {
  private static instance: DashboardPanel | undefined;
  private panel: vscode.WebviewPanel;
  private dataLoader: CertifyDataLoader;
  private disposables: vscode.Disposable[] = [];

  private constructor(panel: vscode.WebviewPanel, dataLoader: CertifyDataLoader) {
    this.panel = panel;
    this.dataLoader = dataLoader;

    this.panel.onDidDispose(() => {
      DashboardPanel.instance = undefined;
      this.disposables.forEach(d => d.dispose());
    });

    this.panel.webview.onDidReceiveMessage(async msg => {
      if (msg.command === 'openFile' && msg.path) {
        const uri = vscode.Uri.file(msg.path);
        await vscode.window.showTextDocument(uri);
      }
    });

    this.dataLoader.onDataChanged(() => this.update(), undefined, this.disposables);
    this.update();
  }

  static createOrShow(dataLoader: CertifyDataLoader): void {
    if (DashboardPanel.instance) {
      DashboardPanel.instance.panel.reveal();
      return;
    }

    const panel = vscode.window.createWebviewPanel(
      'certifyDashboard',
      'Certify Dashboard',
      vscode.ViewColumn.One,
      { enableScripts: true },
    );

    DashboardPanel.instance = new DashboardPanel(panel, dataLoader);
  }

  private async update(): Promise<void> {
    const report = await this.dataLoader.loadFullReport();
    this.panel.webview.html = report
      ? this.renderHTML(report)
      : this.renderEmpty();
  }

  private renderEmpty(): string {
    return `<!DOCTYPE html>
<html><body style="font-family: var(--vscode-font-family); padding: 2rem; text-align: center;">
  <h2>No certification data</h2>
  <p>Run <code>Certify: Scan</code> then <code>Certify: Run Certification</code> to generate data.</p>
</body></html>`;
  }

  private renderHTML(report: FullReport): string {
    const c = report.card;
    const gradeColor = GRADE_COLORS[c.overall_grade] ?? BRAND_COLORS.expired;
    const emoji = GRADE_EMOJI[c.overall_grade] ?? '⚪';

    const gradeDist = Object.entries(c.grade_distribution)
      .sort(([a], [b]) => a.localeCompare(b))
      .map(([grade, count]) => {
        const color = GRADE_COLORS[grade] ?? '#9CA3AF';
        const pct = c.total_units > 0 ? (count / c.total_units) * 100 : 0;
        return `<div class="bar-segment" style="width:${pct}%;background:${color}" title="${grade}: ${count} (${pct.toFixed(1)}%)"></div>`;
      })
      .join('');

    const dimensions = DIMENSION_NAMES.map(dim => {
      const val = (report.dimension_averages[dim] ?? 0) * 100;
      const label = dim.replace(/_/g, ' ');
      return `<div class="dim-row">
        <span class="dim-label">${label}</span>
        <div class="dim-bar"><div class="dim-fill" style="width:${val}%"></div></div>
        <span class="dim-val">${val.toFixed(0)}%</span>
      </div>`;
    }).join('');

    const langs = c.languages.map(l =>
      `<tr><td>${l.name}</td><td>${l.units}</td><td>${l.passing}/${l.units}</td><td>${l.grade}</td><td>${(l.average_score * 100).toFixed(1)}%</td></tr>`
    ).join('');

    const unitRows = report.units.slice(0, 200).map(u => {
      const e = GRADE_EMOJI[u.grade] ?? '⚪';
      return `<tr class="unit-row" data-path="${u.path}">
        <td>${e} ${u.grade}</td><td>${(u.score * 100).toFixed(1)}%</td>
        <td>${u.language}</td><td class="unit-id">${u.symbol ?? u.path}</td>
      </tr>`;
    }).join('');

    return `<!DOCTYPE html>
<html>
<head>
<style>
  :root { --brand: ${BRAND_COLORS.steelBlue}; --grade: ${gradeColor}; }
  body { font-family: var(--vscode-font-family); padding: 1.5rem; color: var(--vscode-foreground); background: var(--vscode-editor-background); margin: 0; }
  h1 { font-size: 1.5rem; margin: 0 0 0.5rem; }
  .hero { text-align: center; padding: 1.5rem; border-radius: 8px; background: var(--vscode-editor-inactiveSelectionBackground); margin-bottom: 1.5rem; }
  .grade-big { font-size: 3rem; font-weight: 700; color: var(--grade); }
  .summary { display: flex; gap: 1rem; flex-wrap: wrap; margin-bottom: 1.5rem; }
  .card { flex: 1; min-width: 100px; padding: 0.75rem; border-radius: 6px; background: var(--vscode-editor-inactiveSelectionBackground); text-align: center; }
  .card-val { font-size: 1.5rem; font-weight: 700; }
  .card-label { font-size: 0.75rem; opacity: 0.7; }
  .section { margin-bottom: 1.5rem; }
  .section h2 { font-size: 1rem; border-bottom: 1px solid var(--vscode-widget-border); padding-bottom: 0.25rem; }
  .grade-bar { display: flex; height: 24px; border-radius: 4px; overflow: hidden; }
  .bar-segment { min-width: 2px; }
  .dim-row { display: flex; align-items: center; gap: 0.5rem; margin: 4px 0; }
  .dim-label { width: 180px; font-size: 0.8rem; text-transform: capitalize; }
  .dim-bar { flex: 1; height: 14px; background: var(--vscode-editor-inactiveSelectionBackground); border-radius: 3px; overflow: hidden; }
  .dim-fill { height: 100%; background: var(--brand); border-radius: 3px; transition: width 0.3s; }
  .dim-val { width: 40px; text-align: right; font-size: 0.8rem; }
  table { width: 100%; border-collapse: collapse; font-size: 0.85rem; }
  th, td { text-align: left; padding: 4px 8px; border-bottom: 1px solid var(--vscode-widget-border); }
  th { opacity: 0.7; font-weight: 600; }
  .unit-row { cursor: pointer; }
  .unit-row:hover { background: var(--vscode-list-hoverBackground); }
  .unit-id { font-family: var(--vscode-editor-font-family); font-size: 0.8rem; }
  .filter { width: 100%; padding: 6px 10px; margin-bottom: 8px; border: 1px solid var(--vscode-input-border); background: var(--vscode-input-background); color: var(--vscode-input-foreground); border-radius: 4px; }
</style>
</head>
<body>
  <div class="hero">
    <h1>Certify — Report Card</h1>
    <div class="grade-big">${emoji} ${c.overall_grade}</div>
    <div>${c.repository} · ${c.commit_sha ?? ''}</div>
  </div>

  <div class="summary">
    <div class="card"><div class="card-val">${c.total_units}</div><div class="card-label">Total Units</div></div>
    <div class="card"><div class="card-val" style="color:${BRAND_COLORS.certified}">${c.passing}</div><div class="card-label">Passing</div></div>
    <div class="card"><div class="card-val" style="color:${BRAND_COLORS.decertified}">${c.failing}</div><div class="card-label">Failing</div></div>
    <div class="card"><div class="card-val">${(c.pass_rate * 100).toFixed(1)}%</div><div class="card-label">Pass Rate</div></div>
  </div>

  <div class="section">
    <h2>Grade Distribution</h2>
    <div class="grade-bar">${gradeDist}</div>
  </div>

  <div class="section">
    <h2>Quality Dimensions</h2>
    ${dimensions}
  </div>

  <div class="section">
    <h2>Languages</h2>
    <table><tr><th>Language</th><th>Units</th><th>Passing</th><th>Grade</th><th>Avg Score</th></tr>${langs}</table>
  </div>

  <div class="section">
    <h2>Units (${report.units.length})</h2>
    <input class="filter" type="text" placeholder="Filter units..." id="unitFilter" />
    <table id="unitTable">
      <tr><th>Grade</th><th>Score</th><th>Language</th><th>Unit</th></tr>
      ${unitRows}
    </table>
  </div>

  <script>
    const vscode = acquireVsCodeApi();
    document.querySelectorAll('.unit-row').forEach(row => {
      row.addEventListener('click', () => {
        vscode.postMessage({ command: 'openFile', path: row.dataset.path });
      });
    });
    document.getElementById('unitFilter')?.addEventListener('input', e => {
      const q = e.target.value.toLowerCase();
      document.querySelectorAll('.unit-row').forEach(row => {
        row.style.display = row.textContent.toLowerCase().includes(q) ? '' : 'none';
      });
    });
  </script>
</body>
</html>`;
  }
}
