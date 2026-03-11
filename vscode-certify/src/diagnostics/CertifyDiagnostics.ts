import * as vscode from 'vscode';
import * as path from 'path';
import type { CertifyDataLoader } from '../dataLoader.js';
import type { RecordJSON } from '../types.js';

const LOW_GRADES = new Set(['D', 'F']);
const EXPIRY_WARNING_DAYS = 14;

export class CertifyDiagnostics {
  private collection: vscode.DiagnosticCollection;
  private records: RecordJSON[] = [];
  private disposables: vscode.Disposable[] = [];

  constructor(
    private dataLoader: CertifyDataLoader,
    private workspaceRoot: string,
  ) {
    this.collection = vscode.languages.createDiagnosticCollection('certify');

    // Update on data change
    dataLoader.onDataChanged(() => this.refresh(), undefined, this.disposables);

    // Update on document open
    vscode.workspace.onDidOpenTextDocument(doc => this.updateDocument(doc), undefined, this.disposables);

    this.refresh();
  }

  private async refresh(): Promise<void> {
    this.records = await this.dataLoader.loadAllRecords();
    this.collection.clear();

    // Update all open documents
    for (const doc of vscode.workspace.textDocuments) {
      this.updateDocument(doc);
    }
  }

  private updateDocument(document: vscode.TextDocument): void {
    const relPath = path.relative(this.workspaceRoot, document.uri.fsPath);
    const fileRecords = this.records.filter(r => r.unit_path === relPath);
    if (fileRecords.length === 0) return;

    const diagnostics: vscode.Diagnostic[] = [];
    const now = Date.now();

    for (const record of fileRecords) {
      // Low grade warning
      if (LOW_GRADES.has(record.grade)) {
        const diag = new vscode.Diagnostic(
          new vscode.Range(0, 0, 0, 0),
          `Certify: Grade ${record.grade} (${Math.round(record.score * 100)}%) — needs improvement`,
          vscode.DiagnosticSeverity.Warning,
        );
        diag.source = 'Certify';
        diagnostics.push(diag);
      }

      // Dead code hint (from deep analysis)
      if (record.dimensions?.['is_dead_code'] && record.dimensions['is_dead_code'] > 0) {
        const diag = new vscode.Diagnostic(
          new vscode.Range(0, 0, 0, 0),
          `Certify: Exported symbol has no external references (dead export)`,
          vscode.DiagnosticSeverity.Hint,
        );
        diag.source = 'Certify';
        diagnostics.push(diag);
      }

      // High fan-in warning (change risk)
      if (record.dimensions?.['fan_in'] && record.dimensions['fan_in'] > 20) {
        const fanIn = Math.round(record.dimensions['fan_in']);
        const diag = new vscode.Diagnostic(
          new vscode.Range(0, 0, 0, 0),
          `Certify: High fan-in (${fanIn} callers) — changes here affect many dependents`,
          vscode.DiagnosticSeverity.Information,
        );
        diag.source = 'Certify';
        diagnostics.push(diag);
      }

      // Expiring soon
      if (record.expires_at) {
        const expires = new Date(record.expires_at).getTime();
        const daysLeft = Math.ceil((expires - now) / (1000 * 60 * 60 * 24));
        if (daysLeft > 0 && daysLeft <= EXPIRY_WARNING_DAYS) {
          const diag = new vscode.Diagnostic(
            new vscode.Range(0, 0, 0, 0),
            `Certify: Certification expires in ${daysLeft} days`,
            vscode.DiagnosticSeverity.Information,
          );
          diag.source = 'Certify';
          diagnostics.push(diag);
        }
      }
    }

    if (diagnostics.length > 0) {
      this.collection.set(document.uri, diagnostics);
    }
  }

  dispose(): void {
    this.collection.dispose();
    this.disposables.forEach(d => d.dispose());
  }
}
