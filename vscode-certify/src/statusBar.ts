import * as vscode from 'vscode';
import type { CertifyDataLoader } from './dataLoader.js';
import { GRADE_COLORS } from './constants.js';

export function createStatusBarItem(dataLoader: CertifyDataLoader): vscode.StatusBarItem {
  const item = vscode.window.createStatusBarItem(vscode.StatusBarAlignment.Left, 100);
  item.command = 'certify.openDashboard';
  item.tooltip = 'Open Certify Dashboard';

  async function update(): Promise<void> {
    const badge = await dataLoader.loadBadge();
    if (badge) {
      item.text = `$(shield) ${badge.message}`;
      const color = GRADE_COLORS[badge.message.split(' ')[0]] ?? '#9CA3AF';
      item.color = color;
      item.show();
    } else {
      item.hide();
    }
  }

  update();
  dataLoader.onDataChanged(() => update());

  return item;
}
