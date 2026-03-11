import * as vscode from 'vscode';
import * as path from 'path';
import type { CertifyDataLoader } from '../dataLoader.js';
import type { RecordJSON } from '../types.js';
import { GRADE_EMOJI, DIMENSION_NAMES, DEEP_METRIC_LABELS } from '../constants.js';

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
  } else if (langId === 'typescript' || langId === 'typescriptreact' || langId === 'javascript' || langId === 'javascriptreact') {
    patterns.push(
      new RegExp(`^(export\\s+)?function\\s+${escaped}\\s*[(<]`),
      new RegExp(`^(export\\s+)?(const|let|var)\\s+${escaped}\\s*=`),
      new RegExp(`^(export\\s+)?class\\s+${escaped}\\s*`),
      new RegExp(`^\\s+${escaped}\\s*\\(`),  // method
    );
  } else if (langId === 'python') {
    patterns.push(
      new RegExp(`^(\\s*)def\\s+${escaped}\\s*\\(`),        // def name(
      new RegExp(`^(\\s*)async\\s+def\\s+${escaped}\\s*\\(`), // async def name(
      new RegExp(`^class\\s+${escaped}\\s*[:(]`),            // class Name:
    );
  } else if (langId === 'rust') {
    patterns.push(
      new RegExp(`^\\s*(pub\\s+)?(async\\s+)?fn\\s+${escaped}\\s*[(<]`), // fn name(
      new RegExp(`^\\s*(pub\\s+)?struct\\s+${escaped}\\s*`),              // struct Name
      new RegExp(`^\\s*(pub\\s+)?enum\\s+${escaped}\\s*`),                // enum Name
      new RegExp(`^\\s*(pub\\s+)?trait\\s+${escaped}\\s*`),               // trait Name
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

  // Dimension scores (0-100%)
  const measuredDims = DIMENSION_NAMES.filter(dim => dim in record.dimensions!);
  const dimItems = measuredDims.map(dim => {
    const val = record.dimensions![dim];
    const bar = '█'.repeat(Math.round(val * 10)) + '░'.repeat(10 - Math.round(val * 10));
    return {
      label: `${bar} ${(val * 100).toFixed(0)}%`,
      description: dim.replace(/_/g, ' '),
    };
  });

  // Deep analysis metrics (if present)
  const deepItems: Array<{ label: string; description: string }> = [];
  for (const [key, label] of Object.entries(DEEP_METRIC_LABELS)) {
    if (key in record.dimensions!) {
      const val = record.dimensions![key];
      deepItems.push({
        label: `  ${label}: ${Number.isInteger(val) ? val : val.toFixed(2)}`,
        description: key,
      });
    }
  }

  const items = [
    ...dimItems,
    ...(deepItems.length > 0 ? [
      { label: '── Deep Analysis ──', description: '' },
      ...deepItems,
    ] : []),
  ];

  await vscode.window.showQuickPick(items, {
    title: `${record.unit_id} — ${record.grade} (${Math.round(record.score * 100)}%) — ${measuredDims.length} dimensions${deepItems.length > 0 ? ` + ${deepItems.length} deep metrics` : ''}`,
    placeHolder: 'Dimension Scores + Deep Analysis',
  });
}
