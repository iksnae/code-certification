import * as vscode from 'vscode';
import * as fs from 'fs';
import * as path from 'path';
import type { IndexEntry, RecordJSON, BadgeJSON, FullReport, CertifyConfig } from './types.js';
import { GRADE_COLORS } from './constants.js';
import { runCertifyJSON } from './certifyBinary.js';

export class CertifyDataLoader {
  private certDir: string;
  private _onDataChanged = new vscode.EventEmitter<void>();
  readonly onDataChanged = this._onDataChanged.event;
  private watcher: vscode.FileSystemWatcher | undefined;

  constructor(private workspaceRoot: string) {
    this.certDir = path.join(workspaceRoot, '.certification');
    this.startWatching();
  }

  private startWatching(): void {
    const pattern = new vscode.RelativePattern(this.certDir, '**/*');
    this.watcher = vscode.workspace.createFileSystemWatcher(pattern);
    this.watcher.onDidChange(() => this._onDataChanged.fire());
    this.watcher.onDidCreate(() => this._onDataChanged.fire());
    this.watcher.onDidDelete(() => this._onDataChanged.fire());
  }

  get hasCertification(): boolean {
    return fs.existsSync(path.join(this.certDir, 'config.yml'));
  }

  async loadIndex(): Promise<IndexEntry[]> {
    const indexPath = path.join(this.certDir, 'index.json');
    try {
      const data = await fs.promises.readFile(indexPath, 'utf-8');
      return JSON.parse(data) as IndexEntry[];
    } catch {
      return [];
    }
  }

  async loadAllRecords(): Promise<RecordJSON[]> {
    const recordsDir = path.join(this.certDir, 'records');
    try {
      const files = await fs.promises.readdir(recordsDir);
      const records: RecordJSON[] = [];
      for (const file of files) {
        if (!file.endsWith('.json') || file.endsWith('.history.jsonl')) continue;
        try {
          const data = await fs.promises.readFile(path.join(recordsDir, file), 'utf-8');
          records.push(JSON.parse(data) as RecordJSON);
        } catch {
          // skip corrupted records
        }
      }
      return records;
    } catch {
      return [];
    }
  }

  async loadConfig(): Promise<CertifyConfig | null> {
    const configPath = path.join(this.certDir, 'config.yml');
    try {
      const yaml = await import('yaml');
      const data = await fs.promises.readFile(configPath, 'utf-8');
      return yaml.parse(data) as CertifyConfig;
    } catch {
      return null;
    }
  }

  async loadBadge(): Promise<BadgeJSON | null> {
    const badgePath = path.join(this.certDir, 'badge.json');
    try {
      const data = await fs.promises.readFile(badgePath, 'utf-8');
      return JSON.parse(data) as BadgeJSON;
    } catch {
      return null;
    }
  }

  async loadFullReport(): Promise<FullReport | null> {
    try {
      return await runCertifyJSON<FullReport>(
        ['report', '--format', 'json'],
        this.workspaceRoot,
      );
    } catch {
      return null;
    }
  }

  getUnitsForFile(filePath: string, records: RecordJSON[]): RecordJSON[] {
    const relPath = path.relative(this.workspaceRoot, filePath);
    return records.filter(r => r.unit_path === relPath);
  }

  getGradeColor(grade: string): string {
    return GRADE_COLORS[grade] ?? '#9CA3AF';
  }

  dispose(): void {
    this.watcher?.dispose();
    this._onDataChanged.dispose();
  }
}
