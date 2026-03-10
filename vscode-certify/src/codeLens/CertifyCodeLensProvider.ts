import * as vscode from 'vscode';
import * as path from 'path';
import type { CertifyDataLoader } from '../dataLoader.js';
import type { RecordJSON } from '../types.js';
import { GRADE_EMOJI, DIMENSION_NAMES } from '../constants.js';

export class CertifyCodeLensProvider implements vscode.CodeLensProvider {
  private _onDidChangeCodeLenses = new vscode.EventEmitter<void>();
  readonly onDidChangeCodeLenses = this._onDidChangeCodeLenses.event;
  private records: RecordJSON[] = [];

  constructor(
    private dataLoader: CertifyDataLoader,
    private workspaceRoot: string,
  ) {
    dataLoader.onDataChanged(() => {
      this.loadRecords();
      this._onDidChangeCodeLenses.fire();
    });
    this.loadRecords();
  }

  private async loadRecords(): Promise<void> {
    this.records = await this.dataLoader.loadAllRecords();
  }

  provideCodeLenses(document: vscode.TextDocument): vscode.CodeLens[] {
    const enabled = vscode.workspace.getConfiguration('certify').get<boolean>('codeLens.enabled', true);
    if (!enabled) return [];

    const relPath = path.relative(this.workspaceRoot, document.uri.fsPath);
    const fileRecords = this.records.filter(r => r.unit_path === relPath);
    if (fileRecords.length === 0) return [];

    const lenses: vscode.CodeLens[] = [];
    const text = document.getText();

    for (const record of fileRecords) {
      const symbol = record.unit_id.includes('#')
        ? record.unit_id.split('#').pop()!
        : null;

      if (!symbol) continue;

      const line = findSymbolLine(text, symbol, document.languageId);
      if (line < 0) continue;

      const range = new vscode.Range(line, 0, line, 0);
      const emoji = GRADE_EMOJI[record.grade] ?? '⚪';
      const score = Math.round(record.score * 100);

      lenses.push(new vscode.CodeLens(range, {
        title: `${emoji} ${record.grade} (${score}%)`,
        command: 'certify.showDimensions',
        arguments: [record],
        tooltip: `Certify: ${record.status} — click for dimension scores`,
      }));
    }

    return lenses;
  }
}

function findSymbolLine(text: string, symbol: string, langId: string): number {
  const lines = text.split('\n');
  const escaped = symbol.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');

  // Language-specific patterns
  const patterns: RegExp[] = [];
  if (langId === 'go') {
    patterns.push(
      new RegExp(`^func\\s+${escaped}\\s*\\(`),           // func Name(
      new RegExp(`^func\\s+\\([^)]+\\)\\s+${escaped}\\s*\\(`), // func (r *T) Name(
      new RegExp(`^type\\s+${escaped}\\s+`),               // type Name struct
    );
  } else if (langId === 'typescript' || langId === 'typescriptreact' || langId === 'javascript') {
    patterns.push(
      new RegExp(`^(export\\s+)?function\\s+${escaped}\\s*[(<]`),
      new RegExp(`^(export\\s+)?(const|let|var)\\s+${escaped}\\s*=`),
      new RegExp(`^(export\\s+)?class\\s+${escaped}\\s*`),
      new RegExp(`^\\s+${escaped}\\s*\\(`),  // method
    );
  }

  for (let i = 0; i < lines.length; i++) {
    const line = lines[i];
    for (const pattern of patterns) {
      if (pattern.test(line)) return i;
    }
  }

  // Fallback: simple text match
  for (let i = 0; i < lines.length; i++) {
    if (lines[i].includes(symbol)) return i;
  }

  return -1;
}

export async function showDimensionScores(record: RecordJSON): Promise<void> {
  if (!record.dimensions || Object.keys(record.dimensions).length === 0) {
    vscode.window.showInformationMessage(`${record.unit_id}: Grade ${record.grade} (${Math.round(record.score * 100)}%)`);
    return;
  }

  // Only show dimensions that have actual measurements for this unit
  const measuredDims = DIMENSION_NAMES.filter(dim => dim in record.dimensions!);
  const items = measuredDims.map(dim => {
    const val = record.dimensions![dim];
    const bar = '█'.repeat(Math.round(val * 10)) + '░'.repeat(10 - Math.round(val * 10));
    return {
      label: `${bar} ${(val * 100).toFixed(0)}%`,
      description: dim.replace(/_/g, ' '),
    };
  });

  await vscode.window.showQuickPick(items, {
    title: `${record.unit_id} — ${record.grade} (${Math.round(record.score * 100)}%) — ${measuredDims.length} dimensions measured`,
    placeHolder: 'Dimension Scores (measured only)',
  });
}
