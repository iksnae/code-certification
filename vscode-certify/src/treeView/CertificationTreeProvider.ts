import * as vscode from 'vscode';
import * as path from 'path';
import type { CertifyDataLoader } from '../dataLoader.js';
import type { RecordJSON } from '../types.js';
import { GRADE_EMOJI } from '../constants.js';

export class CertifyTreeItem extends vscode.TreeItem {
  constructor(
    public readonly label: string,
    public readonly collapsibleState: vscode.TreeItemCollapsibleState,
    public readonly record?: RecordJSON,
    public readonly dirPath?: string,
  ) {
    super(label, collapsibleState);

    if (record) {
      const emoji = GRADE_EMOJI[record.grade] ?? '⚪';
      this.description = `${emoji} ${record.grade} · ${Math.round(record.score * 100)}%`;
      this.tooltip = `${record.unit_id}\nStatus: ${record.status}\nGrade: ${record.grade} (${Math.round(record.score * 100)}%)`;
      this.contextValue = 'certifyUnit';

      // Click to open source file
      if (record.unit_path) {
        this.command = {
          command: 'vscode.open',
          title: 'Open Source',
          arguments: [vscode.Uri.file(record.unit_path)],
        };
      }
    }
  }
}

export class CertificationTreeProvider implements vscode.TreeDataProvider<CertifyTreeItem> {
  private _onDidChangeTreeData = new vscode.EventEmitter<CertifyTreeItem | undefined>();
  readonly onDidChangeTreeData = this._onDidChangeTreeData.event;
  private records: RecordJSON[] = [];

  constructor(
    private dataLoader: CertifyDataLoader,
    private workspaceRoot: string,
  ) {
    dataLoader.onDataChanged(() => this.refresh());
    this.refresh();
  }

  async refresh(): Promise<void> {
    this.records = await this.dataLoader.loadAllRecords();
    this._onDidChangeTreeData.fire(undefined);
  }

  getTreeItem(element: CertifyTreeItem): vscode.TreeItem {
    return element;
  }

  getChildren(element?: CertifyTreeItem): CertifyTreeItem[] {
    if (!element) {
      // Root level: directories
      const dirs = new Map<string, RecordJSON[]>();
      for (const r of this.records) {
        const dir = path.dirname(r.unit_path) || '.';
        if (!dirs.has(dir)) dirs.set(dir, []);
        dirs.get(dir)!.push(r);
      }

      return Array.from(dirs.entries())
        .sort(([a], [b]) => a.localeCompare(b))
        .map(([dir, records]) => {
          const grades = records.map(r => r.grade);
          const worst = worstGrade(grades);
          const item = new CertifyTreeItem(
            dir,
            vscode.TreeItemCollapsibleState.Collapsed,
            undefined,
            dir,
          );
          const emoji = GRADE_EMOJI[worst] ?? '⚪';
          item.description = `${emoji} ${records.length} units`;
          item.iconPath = new vscode.ThemeIcon('folder');
          return item;
        });
    }

    if (element.dirPath) {
      // Children: units in this directory
      return this.records
        .filter(r => path.dirname(r.unit_path) === element.dirPath)
        .sort((a, b) => a.unit_id.localeCompare(b.unit_id))
        .map(r => {
          const name = r.unit_id.includes('#')
            ? r.unit_id.split('#').pop()!
            : path.basename(r.unit_path);
          const item = new CertifyTreeItem(
            name,
            vscode.TreeItemCollapsibleState.None,
            r,
          );
          // Set the command to open with workspace-relative path resolved
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

    return [];
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
