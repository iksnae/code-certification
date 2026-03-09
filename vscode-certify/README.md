# Certify — VSCode Extension

> Code trust, with an expiration date.

View certification report cards, configure AI providers, and run certification directly from Visual Studio Code.

## Features

### 📊 Dashboard
Interactive report card showing overall grade, grade distribution, quality dimension scores, language breakdown, and per-unit details. Auto-refreshes when certification data changes.

### 🌳 Tree View
Explorer sidebar showing all certified units organized by directory. Click any unit to navigate to its source.

### 🔍 CodeLens
Inline grade annotations on Go and TypeScript functions — see `🟢 B+ (87%)` directly above each function. Click for dimension score breakdown.

### 📊 Status Bar
Persistent grade badge showing overall grade, unit count, and pass rate. Click to open the dashboard.

### ⚠️ Diagnostics
Warning markers on files with Grade D or F units. Information markers on soon-to-expire certifications.

### ⚙️ Provider Configuration
Visual configuration for **any** OpenAI-compatible AI provider:

| Provider | Type | Free Tier |
|----------|------|-----------|
| OpenRouter | Cloud | ✅ 200+ models |
| Groq | Cloud | ✅ 30 req/min |
| Together | Cloud | ✅ $1 credit |
| Fireworks | Cloud | ✅ $1 credit |
| OpenAI | Cloud | ❌ |
| Google AI Studio | Cloud | ✅ |
| Ollama | Local | ✅ Free |
| LM Studio | Local | ✅ Free |
| vLLM | Local | ✅ Free |
| Custom | Any | Any |

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
| `Certify: Configure AI Provider` | Set up cloud or local AI provider |
| `Certify: Browse Available Models` | Fetch and browse models from provider |
| `Certify: Install CLI` | Install the certify binary via Go |

## Configuration

| Setting | Default | Description |
|---------|---------|-------------|
| `certify.codeLens.enabled` | `true` | Show grade annotations on functions |
| `certify.binaryPath` | `""` | Path to certify binary (auto-detected) |

## License

MIT
