# Certify — VSCode Extension

> Code trust, with an expiration date.

View certification report cards, configure AI providers, and run certification directly from Visual Studio Code.

## Features

### 📊 Dashboard
Interactive report card showing overall grade, grade distribution, quality dimension scores, language breakdown, and per-unit details. Auto-refreshes when certification data changes.

### 🌳 Tree View
Explorer sidebar showing all certified units organized by directory. Click any unit to navigate to its source.

### 🔍 CodeLens
Inline grade annotations on Go, TypeScript, Python, and Rust functions — see `🟢 B+ (87%)` directly above each function. Click for dimension score breakdown including deep analysis metrics (fan-in, fan-out, dep depth, etc.).

### 📊 Status Bar
Persistent grade badge showing overall grade, unit count, and pass rate. Click to open the dashboard.

### ⚠️ Diagnostics
Warning markers on files with Grade D or F units. Information markers on soon-to-expire certifications. Hint markers on dead exports (unused public API). Info markers on high fan-in functions (change risk hotspots).

### 🔬 Deep Analysis
When using CLI v0.12.0+, the dashboard shows type-aware cross-file analysis: fan-in/fan-out call graph metrics, dead export detection, dependency depth, package instability, and coupling scores. The fan-in hotspots table highlights the riskiest functions to change.

### ⚙️ Provider Configuration
Configure **any** OpenAI-compatible AI provider — via the visual configurator or VS Code Settings (`Ctrl+,` → search "certify"):

| Provider | Type | Free Tier |
|----------|------|-----------|
| OpenRouter | Cloud | ✅ 200+ models |
| OpenAI | Cloud | Paid (gpt-4o-mini, gpt-4o, o3-mini) |
| Google AI Studio | Cloud | ✅ (Gemini 2.0 Flash) |
| Groq | Cloud | ✅ 30 req/min |
| Together | Cloud | ✅ $1 credit |
| Fireworks | Cloud | ✅ $1 credit |
| Ollama | Local | ✅ Free |
| LM Studio | Local | ✅ Free |
| vLLM | Local | ✅ Free |
| Custom | Any | Any |

**Two ways to configure:**
- **Visual configurator** — `Certify: Configure AI Provider` for guided preset selection, connection testing, and model browsing
- **VS Code Settings** — `certify.provider.*` and `certify.agent.*` settings for quick edits; syncs bidirectionally with `.certification/config.yml`

Dynamic model browser — fetch and search available models from any provider. No hardcoded model lists.

## Installation

### Prerequisites
- [Certify CLI](https://github.com/iksnae/code-certification) installed:
  ```bash
  go install github.com/iksnae/code-certification/cmd/certify@latest
  ```
- A repository with `.certification/` directory (run `certify init`)

### From VSIX
```bash
code --install-extension certify-vscode-0.1.0.vsix
```

### From Source
```bash
cd vscode-certify
npm install
npm run build
# Press F5 in VSCode to launch Extension Development Host
```

## Commands

| Command | Description |
|---------|-------------|
| `Certify: Scan Repository` | Discover certifiable code units |
| `Certify: Run Certification` | Evaluate units against policies |
| `Certify: Generate Report` | Generate report and open dashboard |
| `Certify: Open Dashboard` | Show interactive report card |
| `Certify: Configure AI Provider` | Visual provider/model setup |
| `Certify: Browse Available Models` | Fetch and browse models from provider |
| `Certify: Test Provider Connection` | Verify provider connectivity |
| `Certify: Run Doctor` | Check environment, analysis tiers, LSP servers |
| `Certify: Install CLI` | Install the certify binary via Go |

## Settings

All settings are available in VS Code Settings (`Ctrl+,` → search "certify"):

| Setting | Default | Description |
|---------|---------|-------------|
| `certify.codeLens.enabled` | `true` | Show grade annotations on functions |
| `certify.binaryPath` | `""` | Path to certify binary (auto-detected) |
| `certify.provider.preset` | `""` | Quick-select provider (OpenRouter, OpenAI, Ollama, etc.) |
| `certify.provider.baseUrl` | `""` | API base URL (auto-filled from preset) |
| `certify.provider.apiKeyEnvVar` | `""` | Env var name for API key |
| `certify.agent.enabled` | `true` | Enable AI-assisted reviews |
| `certify.agent.model` | `""` | Model ID for reviews |
| `certify.agent.strategy` | `conservative` | Review depth: conservative, standard, or full |

Settings sync bidirectionally with `.certification/config.yml`.

## License

MIT
