import * as vscode from 'vscode';
import { CertifyDataLoader } from './dataLoader.js';
import { createStatusBarItem } from './statusBar.js';
import { CertificationTreeProvider } from './treeView/CertificationTreeProvider.js';
import { CertifyCodeLensProvider, showDimensionScores } from './codeLens/CertifyCodeLensProvider.js';
import { CertifyDiagnostics } from './diagnostics/CertifyDiagnostics.js';
import { DashboardPanel } from './dashboard/DashboardPanel.js';
import { ConfigPanel } from './config/ConfigPanel.js';
import { findCertifyBinary, runInTerminal, promptInstall } from './certifyBinary.js';
import type { RecordJSON } from './types.js';

export function activate(context: vscode.ExtensionContext): void {
  const workspaceFolders = vscode.workspace.workspaceFolders;
  if (!workspaceFolders) return;

  const workspaceRoot = workspaceFolders[0].uri.fsPath;
  const dataLoader = new CertifyDataLoader(workspaceRoot);

  // Set context for view visibility
  vscode.commands.executeCommand('setContext', 'certify:hasData', dataLoader.hasCertification);
  dataLoader.onDataChanged(() => {
    vscode.commands.executeCommand('setContext', 'certify:hasData', dataLoader.hasCertification);
  });

  // Status bar
  const statusBar = createStatusBarItem(dataLoader);
  context.subscriptions.push(statusBar);

  // Tree view
  const treeProvider = new CertificationTreeProvider(dataLoader, workspaceRoot);
  vscode.window.registerTreeDataProvider('certifyUnits', treeProvider);

  // CodeLens
  const codeLensProvider = new CertifyCodeLensProvider(dataLoader, workspaceRoot);
  context.subscriptions.push(
    vscode.languages.registerCodeLensProvider({ language: 'go' }, codeLensProvider),
    vscode.languages.registerCodeLensProvider({ language: 'typescript' }, codeLensProvider),
    vscode.languages.registerCodeLensProvider({ language: 'typescriptreact' }, codeLensProvider),
  );

  // Diagnostics
  const diagnostics = new CertifyDiagnostics(dataLoader, workspaceRoot);
  context.subscriptions.push({ dispose: () => diagnostics.dispose() });

  // Commands
  context.subscriptions.push(
    vscode.commands.registerCommand('certify.openDashboard', () => {
      DashboardPanel.createOrShow(dataLoader);
    }),

    vscode.commands.registerCommand('certify.configureProvider', () => {
      ConfigPanel.createOrShow(workspaceRoot, context.secrets);
    }),

    vscode.commands.registerCommand('certify.scan', async () => {
      if (!await ensureBinary()) return;
      runInTerminal(['scan'], workspaceRoot);
    }),

    vscode.commands.registerCommand('certify.certify', async () => {
      if (!await ensureBinary()) return;
      const choice = await vscode.window.showQuickPick([
        { label: 'Quick (deterministic only)', args: ['certify', '--skip-agent', '--reset-queue'] },
        { label: 'Conservative (free AI)', args: ['certify', '--reset-queue'] },
        { label: 'Standard (3-stage)', args: ['certify', '--reset-queue'] },
        { label: 'Full batch (all units)', args: ['certify', '--reset-queue'] },
      ], { title: 'Certify: Choose Mode' });
      if (choice) runInTerminal(choice.args, workspaceRoot);
    }),

    vscode.commands.registerCommand('certify.report', async () => {
      if (!await ensureBinary()) return;
      runInTerminal(['report'], workspaceRoot);
      DashboardPanel.createOrShow(dataLoader);
    }),

    vscode.commands.registerCommand('certify.listModels', () => {
      ConfigPanel.createOrShow(workspaceRoot, context.secrets);
    }),

    vscode.commands.registerCommand('certify.installCLI', () => {
      const terminal = vscode.window.createTerminal({ name: 'Install Certify' });
      terminal.show();
      terminal.sendText('go install github.com/iksnae/code-certification/cmd/certify@latest');
    }),

    vscode.commands.registerCommand('certify.openUnit', async (filePath: string) => {
      if (filePath) {
        const uri = vscode.Uri.file(filePath);
        await vscode.window.showTextDocument(uri);
      }
    }),

    vscode.commands.registerCommand('certify.showDimensions', (record: RecordJSON) => {
      showDimensionScores(record);
    }),
  );

  // Cleanup
  context.subscriptions.push({ dispose: () => dataLoader.dispose() });
}

async function ensureBinary(): Promise<boolean> {
  const binary = await findCertifyBinary();
  if (!binary) {
    await promptInstall();
    return false;
  }
  return true;
}

export function deactivate(): void {
  // Disposables handled by context.subscriptions
}
