import * as vscode from 'vscode';
import * as path from 'path';
import { PROVIDER_PRESETS } from '../constants.js';
import { writeConfig, readConfig } from './configWriter.js';
import type { CertifyConfig } from '../types.js';

/**
 * Watches VS Code settings for certify.provider.* and certify.agent.* changes,
 * syncs them to .certification/config.yml, and auto-fills preset values.
 */
export function activateSettingsSync(
  context: vscode.ExtensionContext,
  workspaceRoot: string,
): void {
  // Sync preset → baseUrl/apiKeyEnvVar on change
  context.subscriptions.push(
    vscode.workspace.onDidChangeConfiguration(e => {
      if (e.affectsConfiguration('certify.provider.preset')) {
        applyPreset();
      }
      if (
        e.affectsConfiguration('certify.provider') ||
        e.affectsConfiguration('certify.agent')
      ) {
        syncSettingsToConfig(workspaceRoot);
      }
    }),
  );

  // Initial sync if settings exist but config.yml may be stale
  const cfg = vscode.workspace.getConfiguration('certify');
  if (cfg.get<string>('provider.baseUrl') || cfg.get<string>('provider.preset')) {
    syncSettingsToConfig(workspaceRoot);
  }
}

function applyPreset(): void {
  const cfg = vscode.workspace.getConfiguration('certify');
  const presetName = cfg.get<string>('provider.preset', '');
  if (!presetName) return;

  const preset = PROVIDER_PRESETS.find(p => p.name === presetName);
  if (!preset) return;

  // Only auto-fill if the user hasn't manually set a different URL
  const currentUrl = cfg.get<string>('provider.baseUrl', '');
  const isDefaultOrEmpty = !currentUrl || PROVIDER_PRESETS.some(p => p.baseURL === currentUrl);

  if (isDefaultOrEmpty && preset.baseURL) {
    cfg.update('provider.baseUrl', preset.baseURL, vscode.ConfigurationTarget.Workspace);
  }
  cfg.update('provider.apiKeyEnvVar', preset.apiKeyEnvVar, vscode.ConfigurationTarget.Workspace);
}

async function syncSettingsToConfig(workspaceRoot: string): Promise<void> {
  const cfg = vscode.workspace.getConfiguration('certify');
  const baseUrl = cfg.get<string>('provider.baseUrl', '');
  const apiKeyEnv = cfg.get<string>('provider.apiKeyEnvVar', '');
  const enabled = cfg.get<boolean>('agent.enabled', true);
  const model = cfg.get<string>('agent.model', '');

  // Don't write an empty config
  if (!baseUrl && !model) return;

  const certDir = path.join(workspaceRoot, '.certification');

  const config: CertifyConfig = {
    mode: 'advisory',
    agent: {
      enabled,
      provider: {
        type: 'openai-compatible',
        base_url: baseUrl || undefined,
        api_key_env: apiKeyEnv || undefined,
      },
      models: model
        ? { prescreen: model, review: model, scoring: model }
        : undefined,
    },
  };

  try {
    await writeConfig(certDir, config);
  } catch {
    // config.yml directory may not exist yet — that's ok
  }
}

/**
 * Load existing config.yml values into VS Code settings (one-time bootstrap).
 */
export async function bootstrapFromConfig(workspaceRoot: string): Promise<void> {
  const certDir = path.join(workspaceRoot, '.certification');
  const existing = await readConfig(certDir);
  if (!existing?.agent?.provider?.base_url) return;

  const cfg = vscode.workspace.getConfiguration('certify');

  // Only bootstrap if settings are empty (don't overwrite user choices)
  if (cfg.get<string>('provider.baseUrl')) return;

  const baseUrl = existing.agent.provider.base_url;
  const apiKeyEnv = existing.agent.provider.api_key_env || '';
  const model = existing.agent.models?.prescreen || '';

  // Find matching preset
  const preset = PROVIDER_PRESETS.find(p => p.baseURL === baseUrl);

  if (preset) {
    await cfg.update('provider.preset', preset.name, vscode.ConfigurationTarget.Workspace);
  }
  await cfg.update('provider.baseUrl', baseUrl, vscode.ConfigurationTarget.Workspace);
  if (apiKeyEnv) {
    await cfg.update('provider.apiKeyEnvVar', apiKeyEnv, vscode.ConfigurationTarget.Workspace);
  }
  if (model) {
    await cfg.update('agent.model', model, vscode.ConfigurationTarget.Workspace);
  }
  if (existing.agent.enabled !== undefined) {
    await cfg.update('agent.enabled', existing.agent.enabled, vscode.ConfigurationTarget.Workspace);
  }
}
