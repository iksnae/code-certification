import * as fs from 'fs';
import * as path from 'path';
import * as http from 'http';
import * as https from 'https';
import type { CertifyConfig } from '../types.js';

export async function readConfig(certDir: string): Promise<CertifyConfig | null> {
  const configPath = path.join(certDir, 'config.yml');
  try {
    const yaml = await import('yaml');
    const data = await fs.promises.readFile(configPath, 'utf-8');
    return yaml.parse(data) as CertifyConfig;
  } catch {
    return null;
  }
}

export async function writeConfig(certDir: string, config: CertifyConfig): Promise<void> {
  const yaml = await import('yaml');
  const configPath = path.join(certDir, 'config.yml');

  // Read existing to preserve comments/structure for non-agent sections
  let existing: Record<string, unknown> = {};
  try {
    const data = await fs.promises.readFile(configPath, 'utf-8');
    existing = yaml.parse(data) ?? {};
  } catch {
    // file doesn't exist
  }

  // Merge: config overrides existing
  const merged = { ...existing, ...config };
  const output = '# Certify — Configuration\n' + yaml.stringify(merged);
  await fs.promises.writeFile(configPath, output, 'utf-8');
}

export interface ConnectionTestResult {
  ok: boolean;
  latencyMs: number;
  error?: string;
  modelCount?: number;
}

export async function testConnection(baseURL: string, apiKey?: string): Promise<ConnectionTestResult> {
  const url = baseURL.replace(/\/+$/, '') + '/models';
  const start = Date.now();

  return new Promise(resolve => {
    const client = url.startsWith('https') ? https : http;
    const req = client.get(url, {
      timeout: 5000,
      headers: apiKey ? { Authorization: `Bearer ${apiKey}` } : {},
    }, res => {
      let body = '';
      res.on('data', chunk => body += chunk);
      res.on('end', () => {
        const latencyMs = Date.now() - start;
        if (res.statusCode === 200) {
          try {
            const data = JSON.parse(body);
            const count = Array.isArray(data.data) ? data.data.length : 0;
            resolve({ ok: true, latencyMs, modelCount: count });
          } catch {
            resolve({ ok: true, latencyMs });
          }
        } else {
          resolve({ ok: false, latencyMs, error: `HTTP ${res.statusCode}` });
        }
      });
    });

    req.on('error', err => {
      resolve({ ok: false, latencyMs: Date.now() - start, error: err.message });
    });

    req.on('timeout', () => {
      req.destroy();
      resolve({ ok: false, latencyMs: Date.now() - start, error: 'Connection timeout' });
    });
  });
}
