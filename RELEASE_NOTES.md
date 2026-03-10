# Release v0.1.5

**Date:** 2026-03-09

## Highlights

The VS Code extension's AI provider configuration has moved from a custom webview panel to native VS Code Settings — searchable, familiar, and synced with `.certification/config.yml`.

## What's Changed

### Features

- **feat(vscode): native Settings UI for AI provider config** — All provider configuration now lives in VS Code Settings (`Ctrl+,` → search "certify"). Select a preset, set your model, and choose a strategy — all from the standard settings editor. ([#41ced73](../../commit/41ced73))

  - **Provider presets** — Dropdown with 10 presets (OpenRouter, Groq, Ollama, LM Studio, etc.) that auto-fill base URL and API key env var
  - **Browse Models** — `Certify: Browse Available Models` now uses a native QuickPick with search instead of a webview list
  - **Test Connection** — New `Certify: Test Provider Connection` command with notification feedback
  - **Settings sync** — Changes in VS Code Settings automatically sync to `.certification/config.yml`; existing config.yml values bootstrap into settings on activation

### Removed

- **ConfigPanel webview** — The custom HTML/CSS/JS webview panel has been removed in favor of native settings

## Upgrade Notes

- Your existing `.certification/config.yml` values will auto-populate into VS Code Settings on first activation
- Use `Ctrl+,` (or `Cmd+,`) → search `certify.provider` to configure your AI provider
- The `Certify: Configure AI Provider` command now opens Settings instead of a webview

## Full Changelog

```
41ced73 feat(vscode): move AI provider config to native VS Code settings
```
