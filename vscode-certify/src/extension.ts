import * as vscode from 'vscode';
import { CertifyDataLoader } from './dataLoader.js';
import { createStatusBarItem } from './statusBar.js';
import { CertificationTreeProvider } from './treeView/CertificationTreeProvider.js';
import { CertifyCodeLensProvider, showDimensionScores } from './codeLens/CertifyCodeLensProvider.js';
import { CertifyDiagnostics } from './diagnostics/CertifyDiagnostics.js';
import { DashboardPanel } from './dashboard/DashboardPanel.js';
import { activateSettingsSync, bootstrapFromConfig } from './config/settingsSync.js';
import { testConnection } from './config/configWriter.js';
import { findCertifyBinary, runInTerminal, promptInstall, listModels } from './certifyBinary.js';
import type { RecordJSON, ModelInfo } from './types.js';

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
      vscode.commands.executeCommand('workbench.action.openSettings', '@ext:iksnae.certify-vscode certify.provider');
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

    vscode.commands.registerCommand('certify.listModels', async () => {
      const cfg = vscode.workspace.getConfiguration('certify');
      const baseUrl = cfg.get<string>('provider.baseUrl', '');
      const apiKeyEnv = cfg.get<string>('provider.apiKeyEnvVar', '');

      if (!baseUrl) {
        const action = await vscode.window.showWarningMessage(
          'No provider configured. Set up a provider first.',
          'Open Settings',
        );
        if (action === 'Open Settings') {
          vscode.commands.executeCommand('workbench.action.openSettings', '@ext:iksnae.certify-vscode certify.provider');
        }
        return;
      }

      try {
        const models = await vscode.window.withProgress(
          { location: vscode.ProgressLocation.Notification, title: 'Fetching models...' },
          () => listModels(baseUrl, apiKeyEnv || undefined, workspaceRoot),
        );

        if (!models.length) {
          vscode.window.showWarningMessage('No models found at the configured provider.');
          return;
        }

        const items = models.map((m: ModelInfo) => ({
          label: m.id,
          description: [
            m.owned_by ? `by ${m.owned_by}` : '',
            m.context_window ? `${Math.round(m.context_window / 1000)}k ctx` : '',
          ].filter(Boolean).join(' · '),
        }));

        const picked = await vscode.window.showQuickPick(items, {
          title: 'Certify: Select Model',
          placeHolder: 'Search models...',
          matchOnDescription: true,
        });

        if (picked) {
          await cfg.update('agent.model', picked.label, vscode.ConfigurationTarget.Workspace);
          vscode.window.showInformationMessage(`Certify: Model set to ${picked.label}`);
        }
      } catch (err) {
        vscode.window.showErrorMessage(`Failed to fetch models: ${(err as Error).message}`);
      }
    }),

    vscode.commands.registerCommand('certify.testConnection', async () => {
      const cfg = vscode.workspace.getConfiguration('certify');
      const baseUrl = cfg.get<string>('provider.baseUrl', '');
      const apiKeyEnv = cfg.get<string>('provider.apiKeyEnvVar', '');

      if (!baseUrl) {
        const action = await vscode.window.showWarningMessage(
          'No provider configured. Set up a provider first.',
          'Open Settings',
        );
        if (action === 'Open Settings') {
          vscode.commands.executeCommand('workbench.action.openSettings', '@ext:iksnae.certify-vscode certify.provider');
        }
        return;
      }

      const apiKey = apiKeyEnv ? process.env[apiKeyEnv] ?? '' : '';

      const result = await vscode.window.withProgress(
        { location: vscode.ProgressLocation.Notification, title: 'Testing connection...' },
        () => testConnection(baseUrl, apiKey || undefined),
      );

      if (result.ok) {
        const detail = result.modelCount ? ` (${result.modelCount} models, ${result.latencyMs}ms)` : ` (${result.latencyMs}ms)`;
        vscode.window.showInformationMessage(`✓ Connected to provider${detail}`);
      } else {
        vscode.window.showErrorMessage(`✗ Connection failed: ${result.error}`);
      }
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
  activateSettingsSync(context, workspaceRoot);
  bootstrapFromConfig(workspaceRoot);

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
