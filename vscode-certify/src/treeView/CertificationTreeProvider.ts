import * as vscode from 'vscode';
import * as path from 'path';
import type { CertifyDataLoader, ArchitectMeta } from '../dataLoader.js';
import type { RecordJSON, ProjectState, SubmoduleInfo } from '../types.js';
import { GRADE_EMOJI } from '../constants.js';

export class CertifyTreeItem extends vscode.TreeItem {
  constructor(
    public readonly label: string,
    public readonly collapsibleState: vscode.TreeItemCollapsibleState,
    public readonly record?: RecordJSON,
    public readonly dirPath?: string,
    public readonly itemType?: 'action' | 'info' | 'submodule' | 'unit' | 'directory' | 'architect',
  ) {
    super(label, collapsibleState);
  }
}

export class CertificationTreeProvider implements vscode.TreeDataProvider<CertifyTreeItem> {
  private _onDidChangeTreeData = new vscode.EventEmitter<CertifyTreeItem | undefined>();
  readonly onDidChangeTreeData = this._onDidChangeTreeData.event;
  private records: RecordJSON[] = [];
  private _projectState: ProjectState = 'no-config';
  private architectMeta: ArchitectMeta | null = null;

  constructor(
    private dataLoader: CertifyDataLoader,
    private workspaceRoot: string,
  ) {
    dataLoader.onDataChanged(() => this.refresh());
    this.refresh();
  }

  get projectState(): ProjectState {
    return this._projectState;
  }

  async refresh(): Promise<void> {
    this._projectState = this.dataLoader.detectProjectState();
    if (this._projectState === 'ready') {
      this.records = await this.dataLoader.loadAllRecords();
      this.architectMeta = await this.dataLoader.loadArchitectMeta();
    } else {
      this.records = [];
      this.architectMeta = null;
    }
    this._onDidChangeTreeData.fire(undefined);
  }

  getTreeItem(element: CertifyTreeItem): vscode.TreeItem {
    return element;
  }

  getParent(element: CertifyTreeItem): CertifyTreeItem | undefined {
    // Top-level items have no parent
    if (element.itemType === 'action' || element.itemType === 'info' || element.itemType === 'submodule') {
      return undefined;
    }
    if (element.dirPath && element.itemType === 'directory') {
      return undefined;
    }
    // Unit items belong to a directory — but we don't track parent refs, so return undefined
    return undefined;
  }

  getChildren(element?: CertifyTreeItem): CertifyTreeItem[] {
    if (!element) {
      return this.getRootChildren();
    }

    // Expand directory in ready state
    if (element.dirPath && this._projectState === 'ready') {
      return this.getDirectoryChildren(element.dirPath);
    }

    return [];
  }

  private getRootChildren(): CertifyTreeItem[] {
    switch (this._projectState) {
      case 'no-config':
        return this.buildNoConfigItems();
      case 'config-no-data':
        return this.buildConfigNoDataItems();
      case 'ready':
        return this.buildReadyItems();
      case 'workspace':
        return this.buildWorkspaceItems();
    }
  }

  // --- no-config state ---
  private buildNoConfigItems(): CertifyTreeItem[] {
    const init = new CertifyTreeItem('Initialize Certify', vscode.TreeItemCollapsibleState.None, undefined, undefined, 'action');
    init.iconPath = new vscode.ThemeIcon('add');
    init.command = { command: 'certify.init', title: 'Initialize' };
    init.tooltip = 'Run certify init to bootstrap .certification/ in this repository';

    const configure = new CertifyTreeItem('Configure AI Provider', vscode.TreeItemCollapsibleState.None, undefined, undefined, 'action');
    configure.iconPath = new vscode.ThemeIcon('gear');
    configure.command = { command: 'certify.configureProvider', title: 'Configure' };

    const info = new CertifyTreeItem('No .certification/ found', vscode.TreeItemCollapsibleState.None, undefined, undefined, 'info');
    info.iconPath = new vscode.ThemeIcon('info');
    info.description = 'Initialize to get started';

    return [init, configure, info];
  }

  // --- config-no-data state ---
  private buildConfigNoDataItems(): CertifyTreeItem[] {
    const scan = new CertifyTreeItem('Scan Repository', vscode.TreeItemCollapsibleState.None, undefined, undefined, 'action');
    scan.iconPath = new vscode.ThemeIcon('search');
    scan.command = { command: 'certify.scan', title: 'Scan' };

    const certify = new CertifyTreeItem('Run Certification', vscode.TreeItemCollapsibleState.None, undefined, undefined, 'action');
    certify.iconPath = new vscode.ThemeIcon('shield');
    certify.command = { command: 'certify.certify', title: 'Certify' };

    const configure = new CertifyTreeItem('Configure AI Provider', vscode.TreeItemCollapsibleState.None, undefined, undefined, 'action');
    configure.iconPath = new vscode.ThemeIcon('gear');
    configure.command = { command: 'certify.configureProvider', title: 'Configure' };

    const info = new CertifyTreeItem('Config found — no data yet', vscode.TreeItemCollapsibleState.None, undefined, undefined, 'info');
    info.iconPath = new vscode.ThemeIcon('info');
    info.description = 'Run scan then certify';

    return [scan, certify, configure, info];
  }

  // --- ready state (existing behavior + action items) ---
  private buildReadyItems(): CertifyTreeItem[] {
    const items: CertifyTreeItem[] = [];

    // Architect Review section
    items.push(...this.buildArchitectItems());

    // Package directories
    const dirs = new Map<string, RecordJSON[]>();
    for (const r of this.records) {
      const dir = path.dirname(r.unit_path) || '.';
      if (!dirs.has(dir)) dirs.set(dir, []);
      dirs.get(dir)!.push(r);
    }

    const dirItems = Array.from(dirs.entries())
      .sort(([a], [b]) => a.localeCompare(b))
      .map(([dir, records]) => {
        const grades = records.map(r => r.grade);
        const worst = worstGrade(grades);
        const item = new CertifyTreeItem(
          dir,
          vscode.TreeItemCollapsibleState.Collapsed,
          undefined,
          dir,
          'directory',
        );
        const emoji = GRADE_EMOJI[worst] ?? '⚪';
        item.description = `${emoji} ${records.length} units`;
        item.iconPath = new vscode.ThemeIcon('folder');
        return item;
      });

    items.push(...dirItems);
    return items;
  }

  private buildArchitectItems(): CertifyTreeItem[] {
    const items: CertifyTreeItem[] = [];

    if (this.architectMeta) {
      // Review exists — show summary and open action
      const meta = this.architectMeta;
      const label = `🏗 Architect Review`;
      const item = new CertifyTreeItem(label, vscode.TreeItemCollapsibleState.None, undefined, undefined, 'architect');
      const parts: string[] = [];
      if (meta.phases) parts.push(meta.phases);
      if (meta.recommendations) parts.push(`${meta.recommendations} recs`);
      if (meta.model) parts.push(meta.model.split('/').pop() ?? meta.model);
      item.description = parts.join(' · ');
      item.iconPath = new vscode.ThemeIcon('book');
      item.tooltip = [
        'Architectural Review',
        meta.model ? `Model: ${meta.model}` : '',
        meta.tokens ? `Tokens: ${meta.tokens.toLocaleString()}` : '',
        meta.duration ? `Duration: ${meta.duration}` : '',
        meta.phases ? `Phases: ${meta.phases}` : '',
        meta.recommendations ? `Recommendations: ${meta.recommendations}` : '',
      ].filter(Boolean).join('\n');
      item.command = {
        command: 'certify.openArchitectReview',
        title: 'Open Architect Review',
      };
      items.push(item);
    } else {
      // No review — show action to run one
      const item = new CertifyTreeItem('🏗 Run Architect Review', vscode.TreeItemCollapsibleState.None, undefined, undefined, 'action');
      item.iconPath = new vscode.ThemeIcon('beaker');
      item.description = 'AI architectural analysis';
      item.tooltip = 'Run certify architect to generate a comprehensive architectural review with comparative recommendations';
      item.command = {
        command: 'certify.architect',
        title: 'Run Architect Review',
      };
      items.push(item);
    }

    return items;
  }

  private getDirectoryChildren(dirPath: string): CertifyTreeItem[] {
    return this.records
      .filter(r => path.dirname(r.unit_path) === dirPath)
      .sort((a, b) => a.unit_id.localeCompare(b.unit_id))
      .map(r => {
        const name = r.unit_id.includes('#')
          ? r.unit_id.split('#').pop()!
          : path.basename(r.unit_path);
        const emoji = GRADE_EMOJI[r.grade] ?? '⚪';
        const item = new CertifyTreeItem(
          name,
          vscode.TreeItemCollapsibleState.None,
          r,
          undefined,
          'unit',
        );
        item.description = `${emoji} ${r.grade} · ${Math.round(r.score * 100)}%`;
        item.tooltip = `${r.unit_id}\nStatus: ${r.status}\nGrade: ${r.grade} (${Math.round(r.score * 100)}%)`;
        item.contextValue = 'certifyUnit';
        if (r.unit_path) {
          item.command = {
            command: 'vscode.open',
            title: 'Open Source',
            arguments: [vscode.Uri.file(path.join(this.workspaceRoot, r.unit_path))],
          };
        }
        item.iconPath = new vscode.ThemeIcon('symbol-function');
        return item;
      });
  }

  // --- workspace state ---
  private buildWorkspaceItems(): CertifyTreeItem[] {
    const submodules = this.dataLoader.discoverSubmodules();
    const items: CertifyTreeItem[] = [];

    // Summary
    const configured = submodules.filter(s => s.hasConfig);
    const totalUnits = configured.reduce((sum, s) => sum + (s.units ?? 0), 0);
    const summary = new CertifyTreeItem(
      `Workspace: ${configured.length}/${submodules.length} modules`,
      vscode.TreeItemCollapsibleState.None,
      undefined,
      undefined,
      'info',
    );
    summary.iconPath = new vscode.ThemeIcon('repo');
    summary.description = `${totalUnits} units`;
    items.push(summary);

    // Submodules
    for (const sub of submodules) {
      const item = new CertifyTreeItem(
        sub.name,
        vscode.TreeItemCollapsibleState.None,
        undefined,
        undefined,
        'submodule',
      );

      if (sub.hasConfig) {
        const emoji = sub.grade ? (GRADE_EMOJI[sub.grade] ?? '⚪') : '⚪';
        item.description = sub.grade
          ? `${emoji} ${sub.grade} · ${sub.units ?? 0} units`
          : `${sub.units ?? 0} units`;
        item.iconPath = new vscode.ThemeIcon('folder-library');
        item.tooltip = `${sub.path}\nCertified: ${sub.grade ?? 'unknown'}`;
        // Open the submodule's report card
        const reportCard = path.join(this.workspaceRoot, sub.path, '.certification', 'REPORT_CARD.md');
        item.command = {
          command: 'vscode.open',
          title: 'Open Report Card',
          arguments: [vscode.Uri.file(reportCard)],
        };
      } else {
        item.description = 'not configured';
        item.iconPath = new vscode.ThemeIcon('folder');
        item.tooltip = `${sub.path}\nNo .certification/ — run init`;
        item.command = {
          command: 'certify.init',
          title: 'Initialize',
        };
      }

      items.push(item);
    }

    return items;
  }
}

const GRADE_ORDER = ['A', 'A-', 'B+', 'B', 'C', 'D', 'F'];

function worstGrade(grades: string[]): string {
  let worstIdx = 0;
  for (const g of grades) {
    const idx = GRADE_ORDER.indexOf(g);
    if (idx > worstIdx) worstIdx = idx;
  }
  return GRADE_ORDER[worstIdx] ?? 'F';
}
