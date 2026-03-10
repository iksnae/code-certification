import * as vscode from 'vscode';
import { CertifyDataLoader } from './dataLoader.js';
import { createStatusBarItem } from './statusBar.js';
import { CertificationTreeProvider } from './treeView/CertificationTreeProvider.js';
import { CertifyCodeLensProvider, showDimensionScores } from './codeLens/CertifyCodeLensProvider.js';
import { CertifyDiagnostics } from './diagnostics/CertifyDiagnostics.js';
import { DashboardPanel } from './dashboard/DashboardPanel.js';
import { ConfigPanel } from './config/ConfigPanel.js';
import { activateSettingsSync, bootstrapFromConfig } from './config/settingsSync.js';
import { findCertifyBinary, runInTerminal, promptInstall } from './certifyBinary.js';
import type { RecordJSON } from './types.js';

export function activate(context: vscode.ExtensionContext): void {
  const workspaceFolders = vscode.workspace.workspaceFolders;
  if (!workspaceFolders) return;

  const workspaceRoot = workspaceFolders[0].uri.fsPath;
  const dataLoader = new CertifyDataLoader(workspaceRoot);

  // Detect project state and set contexts
  function updateContexts(): void {
    const state = dataLoader.detectProjectState();
    vscode.commands.executeCommand('setContext', 'certify:hasData', state === 'ready' || state === 'workspace');
    vscode.commands.executeCommand('setContext', 'certify:isWorkspace', state === 'workspace');
    vscode.commands.executeCommand('setContext', 'certify:hasConfig', state !== 'no-config');
  }
  updateContexts();
  dataLoader.onDataChanged(() => updateContexts());

  // Status bar
  const statusBar = createStatusBarItem(dataLoader);
  context.subscriptions.push(statusBar);

  // Tree view — always registered, shows state-appropriate content
  const treeProvider = new CertificationTreeProvider(dataLoader, workspaceRoot);
  vscode.window.registerTreeDataProvider('certifyUnits', treeProvider);

  // CodeLens — only useful when data exists
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
    // Init
    vscode.commands.registerCommand('certify.init', async () => {
      if (!await ensureBinary()) return;
      runInTerminal(['init'], workspaceRoot);
    }),

    // Dashboard
    vscode.commands.registerCommand('certify.openDashboard', () => {
      DashboardPanel.createOrShow(dataLoader);
    }),

    // Configure
    vscode.commands.registerCommand('certify.configureProvider', () => {
      ConfigPanel.createOrShow(workspaceRoot, context.secrets);
    }),

    // Scan
    vscode.commands.registerCommand('certify.scan', async () => {
      if (!await ensureBinary()) return;
      runInTerminal(['scan'], workspaceRoot);
    }),

    // Certify
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

    // Report
    vscode.commands.registerCommand('certify.report', async () => {
      if (!await ensureBinary()) return;
      runInTerminal(['report'], workspaceRoot);
      DashboardPanel.createOrShow(dataLoader);
    }),

    // Workspace commands
    vscode.commands.registerCommand('certify.workspaceScan', async () => {
      if (!await ensureBinary()) return;
      runInTerminal(['scan', '--workspace'], workspaceRoot);
    }),

    vscode.commands.registerCommand('certify.workspaceCertify', async () => {
      if (!await ensureBinary()) return;
      const choice = await vscode.window.showQuickPick([
        { label: 'Quick (deterministic only)', args: ['certify', '--workspace', '--skip-agent', '--reset-queue'] },
        { label: 'Conservative (free AI)', args: ['certify', '--workspace', '--reset-queue'] },
        { label: 'Full batch (all units)', args: ['certify', '--workspace', '--reset-queue'] },
      ], { title: 'Certify Workspace: Choose Mode' });
      if (choice) runInTerminal(choice.args, workspaceRoot);
    }),

    vscode.commands.registerCommand('certify.workspaceReport', async () => {
      if (!await ensureBinary()) return;
      runInTerminal(['report', '--workspace'], workspaceRoot);
    }),

    // Existing utility commands
    vscode.commands.registerCommand('certify.listModels', () => {
      ConfigPanel.createOrShow(workspaceRoot, context.secrets);
    }),

    vscode.commands.registerCommand('certify.testConnection', () => {
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

  // Settings sync: VS Code settings ↔ .certification/config.yml
  if (dataLoader.hasCertification) {
    activateSettingsSync(context, workspaceRoot);
    bootstrapFromConfig(workspaceRoot);
  }

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
