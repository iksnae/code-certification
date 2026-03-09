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
    } catch (err) {
      console.error('[Certify] loadFullReport via CLI failed:', err);
      // Fallback: build report from local records
      return this.buildReportFromRecords();
    }
  }

  private async buildReportFromRecords(): Promise<FullReport | null> {
    const records = await this.loadAllRecords();
    if (records.length === 0) return null;

    const badge = await this.loadBadge();
    const totalUnits = records.length;
    const passing = records.filter(r => r.status === 'certified' || r.status === 'certified_with_observations').length;
    const failing = records.filter(r => r.status === 'decertified').length;
    const expired = records.filter(r => r.status === 'expired').length;

    // Grade distribution
    const gradeDist: Record<string, number> = {};
    for (const r of records) {
      gradeDist[r.grade] = (gradeDist[r.grade] ?? 0) + 1;
    }

    // Average score
    const avgScore = records.reduce((sum, r) => sum + r.score, 0) / totalUnits;

    // Determine overall grade from badge or compute
    const overallGrade = badge?.message?.split(' ')[0] ?? computeGrade(avgScore);

    // Dimension averages
    const dimSums: Record<string, number> = {};
    const dimCounts: Record<string, number> = {};
    for (const r of records) {
      if (!r.dimensions) continue;
      for (const [dim, val] of Object.entries(r.dimensions)) {
        dimSums[dim] = (dimSums[dim] ?? 0) + val;
        dimCounts[dim] = (dimCounts[dim] ?? 0) + 1;
      }
    }
    const dimensionAverages: Record<string, number> = {};
    for (const dim of Object.keys(dimSums)) {
      dimensionAverages[dim] = dimSums[dim] / dimCounts[dim];
    }

    // Language breakdown
    const langMap = new Map<string, { scores: number[]; grades: string[] }>();
    for (const r of records) {
      const lang = r.unit_path.endsWith('.go') ? 'Go' : r.unit_path.endsWith('.ts') ? 'TypeScript' : 'Other';
      if (!langMap.has(lang)) langMap.set(lang, { scores: [], grades: [] });
      langMap.get(lang)!.scores.push(r.score);
      langMap.get(lang)!.grades.push(r.grade);
    }

    const languages = Array.from(langMap.entries()).map(([name, data]) => {
      const avg = data.scores.reduce((s, v) => s + v, 0) / data.scores.length;
      return { name, units: data.scores.length, average_score: avg, grade: computeGrade(avg) };
    });

    // Build units
    const units = records.map(r => ({
      unit_id: r.unit_id,
      unit_type: r.unit_type,
      path: r.unit_path,
      language: r.unit_path.endsWith('.go') ? 'Go' : r.unit_path.endsWith('.ts') ? 'TypeScript' : 'Other',
      symbol: r.unit_id.includes('#') ? r.unit_id.split('#').pop() : undefined,
      status: r.status,
      grade: r.grade,
      score: r.score,
      confidence: r.confidence,
      dimensions: r.dimensions ?? {},
      observations: r.observations,
      actions: r.actions,
      certified_at: r.certified_at,
      expires_at: r.expires_at,
      source: r.source,
    }));

    const repo = badge?.label ?? 'unknown';

    return {
      repository: repo,
      generated_at: new Date().toISOString(),
      card: {
        repository: repo,
        generated_at: new Date().toISOString(),
        overall_grade: overallGrade,
        overall_score: avgScore,
        pass_rate: totalUnits > 0 ? passing / totalUnits : 0,
        total_units: totalUnits,
        passing,
        failing,
        expired,
        observations: records.filter(r => r.status === 'certified_with_observations').length,
        grade_distribution: gradeDist,
        languages,
      },
      units,
      dimension_averages: dimensionAverages,
      language_detail: languages.map(l => ({
        ...l,
        grade_distribution: {},
        top_score: 0,
        bottom_score: 0,
      })),
    };
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

function computeGrade(score: number): string {
  if (score >= 0.93) return 'A';
  if (score >= 0.90) return 'A-';
  if (score >= 0.87) return 'B+';
  if (score >= 0.80) return 'B';
  if (score >= 0.70) return 'C';
  if (score >= 0.60) return 'D';
  return 'F';
}
