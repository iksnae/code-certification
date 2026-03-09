import * as vscode from 'vscode';
import { execFile, spawn } from 'child_process';
import { promisify } from 'util';
import * as path from 'path';
import * as fs from 'fs';
import type { ModelInfo } from './types.js';

const execFileAsync = promisify(execFile);

let cachedBinaryPath: string | null = null;

export async function findCertifyBinary(): Promise<string | null> {
  if (cachedBinaryPath && fs.existsSync(cachedBinaryPath)) {
    return cachedBinaryPath;
  }

  // Check user setting
  const configPath = vscode.workspace.getConfiguration('certify').get<string>('binaryPath');
  if (configPath && fs.existsSync(configPath)) {
    cachedBinaryPath = configPath;
    return configPath;
  }

  // Check workspace build directory
  const workspaceFolders = vscode.workspace.workspaceFolders;
  if (workspaceFolders) {
    for (const folder of workspaceFolders) {
      const localBin = path.join(folder.uri.fsPath, 'build', 'bin', 'certify');
      if (fs.existsSync(localBin)) {
        cachedBinaryPath = localBin;
        return localBin;
      }
    }
  }

  // Check PATH
  try {
    const { stdout } = await execFileAsync('which', ['certify']);
    const p = stdout.trim();
    if (p) {
      cachedBinaryPath = p;
      return p;
    }
  } catch {
    // not in PATH
  }

  // Check GOPATH
  try {
    const { stdout } = await execFileAsync('go', ['env', 'GOPATH']);
    const gopath = stdout.trim();
    const goBin = path.join(gopath, 'bin', 'certify');
    if (fs.existsSync(goBin)) {
      cachedBinaryPath = goBin;
      return goBin;
    }
  } catch {
    // go not installed
  }

  return null;
}

export interface RunResult {
  stdout: string;
  stderr: string;
  exitCode: number;
}

export async function runCertify(args: string[], cwd: string): Promise<RunResult> {
  const binary = await findCertifyBinary();
  if (!binary) {
    throw new Error('certify binary not found. Run "Certify: Install CLI" to install.');
  }

  try {
    const { stdout, stderr } = await execFileAsync(binary, args, {
      cwd,
      timeout: 120_000,
      maxBuffer: 10 * 1024 * 1024,
    });
    return { stdout, stderr, exitCode: 0 };
  } catch (err: unknown) {
    const e = err as { stdout?: string; stderr?: string; code?: number };
    return {
      stdout: e.stdout ?? '',
      stderr: e.stderr ?? '',
      exitCode: e.code ?? 1,
    };
  }
}

export async function runCertifyJSON<T>(args: string[], cwd: string): Promise<T> {
  const result = await runCertify(args, cwd);
  if (result.exitCode !== 0 && !result.stdout.trim()) {
    throw new Error(`certify ${args.join(' ')} failed: ${result.stderr}`);
  }
  return JSON.parse(result.stdout) as T;
}

export async function listModels(providerURL: string, apiKeyEnv: string | undefined, cwd: string): Promise<ModelInfo[]> {
  const args = ['models', '--provider-url', providerURL];
  if (apiKeyEnv) {
    args.push('--api-key-env', apiKeyEnv);
  }
  return runCertifyJSON<ModelInfo[]>(args, cwd);
}

export function runInTerminal(args: string[], cwd: string): void {
  const terminal = vscode.window.createTerminal({
    name: 'Certify',
    cwd,
  });
  terminal.show();
  terminal.sendText(`certify ${args.join(' ')}`);
}

export async function promptInstall(): Promise<void> {
  const action = await vscode.window.showWarningMessage(
    'Certify CLI not found. Install it?',
    'Install via Go',
    'Cancel',
  );
  if (action === 'Install via Go') {
    const terminal = vscode.window.createTerminal({ name: 'Install Certify' });
    terminal.show();
    terminal.sendText('go install github.com/iksnae/code-certification/cmd/certify@latest');
  }
}
