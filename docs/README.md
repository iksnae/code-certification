# Code Certification System — Documentation

## Installation

### From Source

```bash
git clone https://github.com/iksnae/code-certification.git
cd code-certification
go build -o certify ./cmd/certify/
sudo mv certify /usr/local/bin/  # optional: install globally
```

### Requirements

- **Go 1.22+** for building
- **Git** for evidence collection (history, churn metrics)
- **golangci-lint** (optional) for enhanced lint evidence
- **OPENROUTER_API_KEY** (optional) for agent-assisted review

## Quickstart

```bash
# 1. Initialize in your repository
cd /path/to/your-repo
certify init

# 2. Discover code units
certify scan

# 3. Run certification (deterministic only)
certify certify --skip-agent

# 4. View results
certify report
certify report --detailed
certify report --format json
```

## CLI Reference

| Command | Description |
|---------|-------------|
| `certify init` | Bootstrap `.certification/` directory with config, policies, workflows |
| `certify scan` | Discover certifiable code units, save to index |
| `certify certify` | Evaluate units against policies, collect evidence, assign status |
| `certify report` | Generate health report (text or JSON) |
| `certify expire` | Mark overdue certifications as expired |
| `certify review` | Generate PR review annotation |
| `certify version` | Show version information |

### `certify certify` Flags

| Flag | Description |
|------|-------------|
| `--skip-agent` | Skip agent-assisted review |
| `--batch N` | Process N units per run (0=all) |
| `--reset-queue` | Rebuild work queue from scratch |
| `--target path` | Only certify units under given path(s) |
| `--diff-base ref` | Only certify files changed since git ref |
| `--path dir` | Repository root (default: current directory) |

### `certify report` Flags

| Flag | Description |
|------|-------------|
| `--format text\|json` | Output format |
| `--detailed` | Include dimension breakdowns, risk analysis |
| `--path dir` | Repository root |

## Incremental Processing

The certify command uses a persistent work queue for incremental processing:

```bash
# Process 20 units at a time (useful with rate-limited agent review)
certify certify --batch 20
# ... wait, then resume
certify certify --batch 20
# ... repeat until "Queue complete!"

# Start over
certify certify --reset-queue
```

## Agent-Assisted Review

When enabled, the system uses open-weight LLM models via OpenRouter for code review:

```yaml
# .certification/config.yml
agent:
  enabled: true
  provider:
    type: openrouter
    base_url: https://openrouter.ai/api/v1
    api_key_env: OPENROUTER_API_KEY
```

Set the API key as a repository secret for CI, or in `.env` for local use.

Models are selected for open-weight licensing (Apache 2.0) and fine-tunability:
- **Primary**: `qwen/qwen3-coder:free` (code-specialized, 262k context)
- **Fallback**: `mistralai/mistral-nemo` (12B, clean JSON output)

Agent review is always optional. The system works fully without it.

## GitHub Actions

Initialize creates three workflow files:
- `certification-pr.yml` — Runs on PRs, posts review comment
- `certification-nightly.yml` — Nightly sweep, commits updated records
- `certification-weekly.yml` — Weekly report generation

## Further Reading

- [Architecture](architecture.md) — System design and package structure
- [Policy Authoring](policy-authoring.md) — Writing custom policy packs
- [Troubleshooting](troubleshooting.md) — Common issues and solutions
