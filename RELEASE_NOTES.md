# Release v0.1.6

**Date:** 2026-03-09

## Highlights

Provider configuration restored and improved — the visual configurator is back alongside new VS Code Settings integration. Documentation updated across the entire project to reflect all 10 supported AI providers.

## What's Changed

### Bug Fixes

- **fix(vscode): restore ConfigPanel for provider configuration** — The visual configurator webview was accidentally removed in v0.1.5. It's back, and now syncs saved config into VS Code Settings bidirectionally. ([#23010ad](../../commit/23010ad))

### Improvements

- **docs: comprehensive provider documentation** — All 10 supported providers (OpenRouter, OpenAI, Google AI Studio, Groq, Together, Fireworks, Ollama, LM Studio, vLLM, Custom) are now consistently documented across README, extension README, website docs, and troubleshooting guide. ([#ca969fb](../../commit/ca969fb))

- **docs: multi-provider config examples** — Configuration reference now includes OpenAI, OpenRouter, and local Ollama examples side by side. Troubleshooting updated with API key setup for all major providers. ([#ca969fb](../../commit/ca969fb))

## VS Code Extension (v0.1.6)

- ConfigPanel restored for guided provider setup
- ConfigPanel saves now sync into VS Code Settings
- Both configuration paths work: visual configurator + native Settings
- `Certify: Test Provider Connection` command added
- All 6 new settings documented: `provider.preset`, `provider.baseUrl`, `provider.apiKeyEnvVar`, `agent.enabled`, `agent.model`, `agent.strategy`

## Full Changelog

```
ca969fb docs: update provider docs across repo + extension
23010ad fix(vscode): restore ConfigPanel for provider configuration
2c5c7e2 chore(vscode): bump extension version to 0.1.4
```
