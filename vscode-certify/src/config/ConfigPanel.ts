import * as vscode from 'vscode';
import * as path from 'path';
import { PROVIDER_PRESETS, BRAND_COLORS } from '../constants.js';
import { writeConfig, testConnection } from './configWriter.js';
import { listModels } from '../certifyBinary.js';
import type { CertifyConfig, ModelInfo, ProviderPreset } from '../types.js';

export class ConfigPanel {
  private static instance: ConfigPanel | undefined;
  private panel: vscode.WebviewPanel;

  private constructor(
    panel: vscode.WebviewPanel,
    private workspaceRoot: string,
    private secrets: vscode.SecretStorage,
  ) {
    this.panel = panel;
    this.panel.onDidDispose(() => { ConfigPanel.instance = undefined; });
    this.panel.webview.onDidReceiveMessage(msg => this.handleMessage(msg));
    this.panel.webview.html = this.renderHTML();
  }

  static createOrShow(workspaceRoot: string, secrets: vscode.SecretStorage): void {
    if (ConfigPanel.instance) {
      ConfigPanel.instance.panel.reveal();
      return;
    }
    const panel = vscode.window.createWebviewPanel(
      'certifyConfig',
      'Certify: Configure AI Provider',
      vscode.ViewColumn.One,
      { enableScripts: true },
    );
    ConfigPanel.instance = new ConfigPanel(panel, workspaceRoot, secrets);
  }

  private async handleMessage(msg: Record<string, unknown>): Promise<void> {
    switch (msg.command) {
      case 'testConnection': {
        const key = msg.apiKeyEnvVar ? process.env[msg.apiKeyEnvVar as string] ?? '' : '';
        const result = await testConnection(msg.baseURL as string, key || undefined);
        this.panel.webview.postMessage({ command: 'testResult', ...result });
        break;
      }
      case 'fetchModels': {
        try {
          const models = await listModels(
            msg.baseURL as string,
            (msg.apiKeyEnvVar as string) || undefined,
            this.workspaceRoot,
          );
          this.panel.webview.postMessage({ command: 'modelsResult', models });
        } catch (err) {
          this.panel.webview.postMessage({
            command: 'modelsResult',
            models: [],
            error: (err as Error).message,
          });
        }
        break;
      }
      case 'saveConfig': {
        const config = msg.config as CertifyConfig;
        const certDir = path.join(this.workspaceRoot, '.certification');
        await writeConfig(certDir, config);

        // Sync to VS Code settings
        const vsCfg = vscode.workspace.getConfiguration('certify');
        if (config.agent?.provider?.base_url) {
          await vsCfg.update('provider.baseUrl', config.agent.provider.base_url, vscode.ConfigurationTarget.Workspace);
        }
        if (config.agent?.provider?.api_key_env !== undefined) {
          await vsCfg.update('provider.apiKeyEnvVar', config.agent.provider.api_key_env || '', vscode.ConfigurationTarget.Workspace);
        }
        if (config.agent?.models?.prescreen) {
          await vsCfg.update('agent.model', config.agent.models.prescreen, vscode.ConfigurationTarget.Workspace);
        }
        if (config.agent?.enabled !== undefined) {
          await vsCfg.update('agent.enabled', config.agent.enabled, vscode.ConfigurationTarget.Workspace);
        }

        vscode.window.showInformationMessage('Certify: Configuration saved');
        break;
      }
      case 'saveAPIKey': {
        const provider = msg.provider as string;
        const key = msg.key as string;
        await this.secrets.store(`certify.apiKey.${provider}`, key);
        vscode.window.showInformationMessage(`Certify: API key saved for ${provider}`);
        break;
      }
    }
  }

  private renderHTML(): string {
    const presetCards = PROVIDER_PRESETS.map(p => {
      const badge = p.local ? '<span class="badge local">Local</span>'
        : p.name === 'OpenRouter' || p.name === 'Groq' ? '<span class="badge free">Free Tier</span>' : '';
      return `<div class="preset-card" data-name="${p.name}" data-url="${p.baseURL}" data-env="${p.apiKeyEnvVar}" data-local="${p.local}">
        <div class="preset-name">${p.name} ${badge}</div>
        <div class="preset-desc">${p.description}</div>
      </div>`;
    }).join('');

    return `<!DOCTYPE html>
<html>
<head>
<style>
  body { font-family: var(--vscode-font-family); padding: 1.5rem; color: var(--vscode-foreground); background: var(--vscode-editor-background); margin: 0; }
  h1 { font-size: 1.3rem; margin: 0 0 1rem; }
  h2 { font-size: 1rem; margin: 1.5rem 0 0.5rem; border-bottom: 1px solid var(--vscode-widget-border); padding-bottom: 0.25rem; }
  .preset-grid { display: grid; grid-template-columns: repeat(auto-fill, minmax(220px, 1fr)); gap: 8px; }
  .preset-card { padding: 10px; border: 1px solid var(--vscode-widget-border); border-radius: 6px; cursor: pointer; transition: border-color 0.2s; }
  .preset-card:hover, .preset-card.selected { border-color: ${BRAND_COLORS.steelBlue}; background: var(--vscode-list-hoverBackground); }
  .preset-name { font-weight: 600; font-size: 0.9rem; }
  .preset-desc { font-size: 0.75rem; opacity: 0.7; margin-top: 4px; }
  .badge { display: inline-block; font-size: 0.65rem; padding: 1px 6px; border-radius: 3px; margin-left: 4px; }
  .badge.free { background: ${BRAND_COLORS.certified}; color: white; }
  .badge.local { background: ${BRAND_COLORS.steelBlue}; color: white; }
  .form-group { margin: 0.75rem 0; }
  label { display: block; font-size: 0.85rem; margin-bottom: 4px; font-weight: 500; }
  input[type=text], input[type=password], select { width: 100%; padding: 6px 10px; border: 1px solid var(--vscode-input-border); background: var(--vscode-input-background); color: var(--vscode-input-foreground); border-radius: 4px; box-sizing: border-box; }
  button { padding: 6px 16px; border: none; border-radius: 4px; cursor: pointer; font-size: 0.85rem; }
  .btn-primary { background: ${BRAND_COLORS.steelBlue}; color: white; }
  .btn-secondary { background: var(--vscode-button-secondaryBackground); color: var(--vscode-button-secondaryForeground); }
  .btn-row { display: flex; gap: 8px; margin-top: 1rem; }
  .status { font-size: 0.8rem; margin-top: 4px; }
  .status.ok { color: ${BRAND_COLORS.certified}; }
  .status.err { color: ${BRAND_COLORS.decertified}; }
  .model-list { max-height: 300px; overflow-y: auto; border: 1px solid var(--vscode-widget-border); border-radius: 4px; margin-top: 8px; }
  .model-item { padding: 6px 10px; cursor: pointer; font-size: 0.8rem; font-family: var(--vscode-editor-font-family); border-bottom: 1px solid var(--vscode-widget-border); }
  .model-item:hover { background: var(--vscode-list-hoverBackground); }
  .model-item.selected { background: var(--vscode-list-activeSelectionBackground); color: var(--vscode-list-activeSelectionForeground); }
  .radio-group { display: flex; gap: 1rem; }
  .radio-group label { display: inline-flex; align-items: center; gap: 4px; font-weight: 400; }
  #modelFilter { margin-top: 8px; }
  .hidden { display: none; }
</style>
</head>
<body>
  <h1>Certify — AI Provider Configuration</h1>
  <p style="font-size:0.85rem;opacity:0.7;">Any OpenAI-compatible endpoint works. Pick a preset or enter a custom URL.</p>

  <h2>1. Select Provider</h2>
  <div class="preset-grid">${presetCards}</div>

  <div id="configSection" class="hidden">
    <h2>2. Connection</h2>
    <div class="form-group">
      <label>API Base URL</label>
      <input type="text" id="baseURL" placeholder="https://api.example.com/v1" />
    </div>
    <div class="form-group" id="apiKeyGroup">
      <label>API Key</label>
      <input type="password" id="apiKey" placeholder="sk-..." />
      <button class="btn-secondary" onclick="saveKey()" style="margin-top:4px">Save to VSCode</button>
    </div>
    <div class="btn-row">
      <button class="btn-secondary" onclick="testConn()">Test Connection</button>
      <button class="btn-primary" onclick="fetchModels()">Fetch Models</button>
    </div>
    <div class="status" id="connStatus"></div>

    <h2>3. Select Model</h2>
    <input type="text" id="modelFilter" class="filter" placeholder="Search models..." />
    <div class="model-list" id="modelList"><div style="padding:10px;opacity:0.5">Click "Fetch Models" to load available models</div></div>
    <div class="form-group">
      <label>Selected Model</label>
      <input type="text" id="selectedModel" placeholder="(select from list or type model ID)" />
    </div>

    <h2>4. Strategy</h2>
    <div class="radio-group">
      <label><input type="radio" name="strategy" value="conservative" checked /> Conservative (prescreen only)</label>
      <label><input type="radio" name="strategy" value="standard" /> Standard (3-stage)</label>
      <label><input type="radio" name="strategy" value="full" /> Full (5-stage)</label>
    </div>

    <div class="btn-row" style="margin-top:1.5rem">
      <button class="btn-primary" onclick="saveConfig()">Save Configuration</button>
    </div>
  </div>

  <script>
    const vscode = acquireVsCodeApi();
    let selectedPreset = null;
    let models = [];

    document.querySelectorAll('.preset-card').forEach(card => {
      card.addEventListener('click', () => {
        document.querySelectorAll('.preset-card').forEach(c => c.classList.remove('selected'));
        card.classList.add('selected');
        selectedPreset = {
          name: card.dataset.name,
          url: card.dataset.url,
          env: card.dataset.env,
          local: card.dataset.local === 'true',
        };
        document.getElementById('baseURL').value = selectedPreset.url;
        document.getElementById('apiKeyGroup').classList.toggle('hidden', selectedPreset.local);
        document.getElementById('configSection').classList.remove('hidden');
      });
    });

    function testConn() {
      const baseURL = document.getElementById('baseURL').value;
      const env = selectedPreset?.env || '';
      document.getElementById('connStatus').textContent = 'Testing...';
      vscode.postMessage({ command: 'testConnection', baseURL, apiKeyEnvVar: env });
    }

    function fetchModels() {
      const baseURL = document.getElementById('baseURL').value;
      const env = selectedPreset?.env || '';
      document.getElementById('connStatus').textContent = 'Fetching models...';
      vscode.postMessage({ command: 'fetchModels', baseURL, apiKeyEnvVar: env });
    }

    function saveKey() {
      const key = document.getElementById('apiKey').value;
      if (!key || !selectedPreset) return;
      vscode.postMessage({ command: 'saveAPIKey', provider: selectedPreset.name, key });
    }

    function saveConfig() {
      const baseURL = document.getElementById('baseURL').value;
      const model = document.getElementById('selectedModel').value;
      const strategy = document.querySelector('input[name=strategy]:checked')?.value || 'conservative';
      const env = selectedPreset?.env || '';

      const config = {
        mode: 'advisory',
        agent: {
          enabled: true,
          provider: { type: 'openai-compatible', base_url: baseURL, api_key_env: env || undefined },
          models: { prescreen: model, review: model, scoring: model },
        },
      };
      vscode.postMessage({ command: 'saveConfig', config });
    }

    window.addEventListener('message', e => {
      const msg = e.data;
      if (msg.command === 'testResult') {
        const el = document.getElementById('connStatus');
        el.className = 'status ' + (msg.ok ? 'ok' : 'err');
        el.textContent = msg.ok
          ? '✓ Connected (' + msg.latencyMs + 'ms' + (msg.modelCount ? ', ' + msg.modelCount + ' models' : '') + ')'
          : '✗ ' + (msg.error || 'Failed');
      }
      if (msg.command === 'modelsResult') {
        models = msg.models || [];
        const el = document.getElementById('connStatus');
        if (msg.error) {
          el.className = 'status err';
          el.textContent = '✗ ' + msg.error;
        } else {
          el.className = 'status ok';
          el.textContent = '✓ ' + models.length + ' models loaded';
        }
        renderModels(models);
      }
    });

    function renderModels(list) {
      const container = document.getElementById('modelList');
      if (list.length === 0) {
        container.innerHTML = '<div style="padding:10px;opacity:0.5">No models found. Enter model ID manually below.</div>';
        return;
      }
      container.innerHTML = list.map(m =>
        '<div class="model-item" data-id="' + m.id + '">' +
        '<strong>' + m.id + '</strong>' +
        (m.owned_by ? ' <span style="opacity:0.5">(' + m.owned_by + ')</span>' : '') +
        (m.context_window ? ' <span style="opacity:0.5">' + Math.round(m.context_window/1000) + 'k ctx</span>' : '') +
        '</div>'
      ).join('');

      container.querySelectorAll('.model-item').forEach(item => {
        item.addEventListener('click', () => {
          container.querySelectorAll('.model-item').forEach(i => i.classList.remove('selected'));
          item.classList.add('selected');
          document.getElementById('selectedModel').value = item.dataset.id;
        });
      });
    }

    document.getElementById('modelFilter')?.addEventListener('input', e => {
      const q = e.target.value.toLowerCase();
      const filtered = models.filter(m => m.id.toLowerCase().includes(q) || (m.owned_by || '').toLowerCase().includes(q));
      renderModels(filtered);
    });
  </script>
</body>
</html>`;
  }
}
