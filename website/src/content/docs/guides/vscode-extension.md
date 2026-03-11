---
title: VSCode Extension
description: View report cards, configure AI providers, and run certification from Visual Studio Code.
---

import { Tabs, TabItem, Card, CardGrid } from '@astrojs/starlight/components';

The Certify VSCode extension brings your certification data directly into the editor — grade annotations on functions, an interactive dashboard, diagnostics for failing units, and a visual AI provider configurator.

## Install

<Tabs>
  <TabItem label="Marketplace">
    Search **"Certify"** in the VSCode Extensions sidebar, or:

    ```bash
    code --install-extension iksnae.certify-vscode
    ```
  </TabItem>
  <TabItem label="From VSIX">
    Download the latest `.vsix` from [GitHub Releases](https://github.com/iksnae/code-certification/releases), then:

    ```bash
    code --install-extension certify-vscode-*.vsix
    ```
  </TabItem>
  <TabItem label="From Source">
    ```bash
    git clone https://github.com/iksnae/code-certification.git
    cd code-certification/vscode-certify
    npm install && npm run build
    ```
    Press **F5** in VSCode to launch the Extension Development Host.
  </TabItem>
</Tabs>

### Prerequisites

The extension requires the Certify CLI:

```bash
go install github.com/iksnae/code-certification/cmd/certify@latest
```

Or run **Certify: Install CLI** from the command palette.

---

## Features

### 📊 Dashboard

Open with **Certify: Open Dashboard** or click the status bar badge.

The dashboard shows:
- **Overall grade** with hero display
- **Summary cards** — total units, passing, failing, pass rate
- **Grade distribution** — horizontal stacked bar
- **Quality dimensions** — bar chart for all 9 dimensions
- **Language breakdown** — table with per-language grades
- **Unit table** — searchable list of every unit, click to navigate to source

The dashboard auto-refreshes when certification data changes.

### 🌳 Tree View

The **Certify** panel appears in the Explorer sidebar when `.certification/` exists. Units are organized by directory with grade indicators:

- 🟢 All units Grade B or better
- 🟡 Some Grade C units
- 🔴 Grade D or F units present

Click any unit to open its source file.

### 🔍 CodeLens

Inline annotations appear above Go, TypeScript, Python, and Rust functions:

```
🟢 B+ (87%)
func ProcessOrder(ctx context.Context, order *Order) error {
```

Click the annotation to see dimension scores **and** deep analysis metrics (fan-in, fan-out, dep depth, coupling, etc.) in a quick pick.

Supported symbol patterns:
- **Go**: `func`, methods, `type`
- **TypeScript/JavaScript**: `function`, `class`, `const`/`let`/`var`, methods
- **Python**: `def`, `async def`, `class`
- **Rust**: `fn`, `struct`, `enum`, `trait`

Disable with `certify.codeLens.enabled: false` in settings.

### 📊 Status Bar

A persistent badge in the bottom status bar shows the overall grade:

> `$(shield) B+ · 100% pass`

Click to open the dashboard.

### 🔬 Deep Analysis

When using CLI v0.12.0+, the dashboard includes a **Deep Analysis** section showing:
- Units with type-aware analysis
- Max fan-in / fan-out across the codebase
- Dead export count (unused public API)
- **Fan-in hotspots table** — the riskiest functions to change, sorted by caller count

### ⚠️ Diagnostics

The Problems panel shows:
- **⚠️ Warnings** for Grade D and F units
- **ℹ️ Information** for certifications expiring within 14 days
- **ℹ️ Information** for high fan-in functions (>20 callers — change risk)
- **💡 Hints** for dead exports (exported symbols with no external references)

### ⚙️ AI Provider Configuration

Configure providers two ways:

**Visual Configurator** — Open with **Certify: Configure AI Provider**:
1. Select a provider preset (or enter a custom URL)
2. Enter API key (cloud providers) — saved to VSCode SecretStorage
3. Click **Test Connection** to verify
4. Click **Fetch Models** to browse available models
5. Select a model and strategy
6. **Save Configuration** writes to `.certification/config.yml`

**VS Code Settings** — Open with `Ctrl+,` → search "certify":
- `certify.provider.preset` — Quick-select from 10 presets
- `certify.provider.baseUrl` — API endpoint (auto-filled from preset)
- `certify.provider.apiKeyEnvVar` — Env var name for API key
- `certify.agent.model` — Model ID (use **Browse Models** to discover)
- `certify.agent.strategy` — conservative / standard / full

Both methods sync bidirectionally with `.certification/config.yml`.

Supports **any OpenAI-compatible API**:

| Provider | Type | Free Tier |
|----------|------|-----------|
| OpenRouter | Cloud | ✅ 200+ models |
| OpenAI | Cloud | Paid (gpt-4o-mini, gpt-4o, o3-mini) |
| Google AI Studio | Cloud | ✅ (Gemini 2.0 Flash) |
| Groq | Cloud | ✅ 30 req/min |
| Together | Cloud | ✅ $1 credit |
| Fireworks | Cloud | ✅ $1 credit |
| Ollama | Local | ✅ Unlimited |
| LM Studio | Local | ✅ Unlimited |
| vLLM | Local | ✅ Unlimited |
| Custom | Any | — |

---

## Commands

| Command | Description |
|---------|-------------|
| **Certify: Scan Repository** | Discover all certifiable code units |
| **Certify: Run Certification** | Evaluate units against policies |
| **Certify: Generate Report** | Generate report card and open dashboard |
| **Certify: Open Dashboard** | Show interactive report card |
| **Certify: Configure AI Provider** | Visual provider/model setup |
| **Certify: Browse Available Models** | Fetch models from provider |
| **Certify: Test Provider Connection** | Verify provider connectivity |
| **Certify: Run Doctor** | Check environment, analysis tiers, tools, AI providers |
| **Certify: Install CLI** | Install `certify` via `go install` |
| **Certify: Go to Unit Source** | Navigate to a unit's source file |

---

## Settings

All settings are available in VS Code Settings (`Ctrl+,` → search "certify"):

| Setting | Default | Description |
|---------|---------|-------------|
| `certify.codeLens.enabled` | `true` | Show grade annotations on functions |
| `certify.binaryPath` | `""` | Path to `certify` binary (auto-detected if empty) |
| `certify.provider.preset` | `""` | Quick-select provider preset |
| `certify.provider.baseUrl` | `""` | API base URL (auto-filled from preset) |
| `certify.provider.apiKeyEnvVar` | `""` | Env var name for API key |
| `certify.agent.enabled` | `true` | Enable AI-assisted reviews |
| `certify.agent.model` | `""` | Model ID for reviews |
| `certify.agent.strategy` | `conservative` | Review depth (conservative/standard/full) |

Settings sync bidirectionally with `.certification/config.yml`.

Auto-detection checks:
1. `certify.binaryPath` setting
2. `build/bin/certify` in workspace
3. `certify` in `$PATH`
4. `$GOPATH/bin/certify`

---

## How It Works

The extension **does not reimplement certification logic**. It uses the `certify` CLI as the source of truth:

- **Data reading**: Loads `.certification/index.json`, `records/*.json`, `badge.json`, `config.yml` directly from disk
- **Report generation**: Shells out to `certify report --format json`
- **Model discovery**: Shells out to `certify models --provider-url <url>`
- **Running commands**: Opens a VSCode terminal with the appropriate `certify` command
- **File watching**: Monitors `.certification/` for changes and auto-refreshes all UI

This means the extension always shows exactly what the CLI produces — no drift, no inconsistency.
