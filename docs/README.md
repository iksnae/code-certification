# Certify — Documentation

## Install

```bash
# From Go module registry
go install github.com/iksnae/code-certification/cmd/certify@latest

# Or from source
git clone https://github.com/iksnae/code-certification.git
cd code-certification
go build -o certify ./cmd/certify/
sudo mv certify /usr/local/bin/    # optional: install globally
```

**Requirements:** Go 1.22+, Git

**Optional tools (enhanced evidence):**
- [golangci-lint](https://golangci-lint.run/) — Go lint evidence
- [ESLint](https://eslint.org/) — TypeScript/JavaScript lint
- [ruff](https://docs.astral.sh/ruff/) — Python lint
- [cargo clippy](https://doc.rust-lang.org/clippy/) — Rust lint
- [gh CLI](https://cli.github.com/) — PR/issue integration

**Optional LSP servers (Tier 2 deep analysis for non-Go languages):**
- `typescript-language-server` — fan-in/fan-out, dead code for TS
- `pyright` — fan-in/fan-out, dead code for Python
- `rust-analyzer` — fan-in/fan-out, dead code for Rust

Run `certify doctor` to see what's available and what's missing.

## Quick Start

```bash
cd /path/to/your-repo

# 1. Bootstrap certification
certify init

# 2. Discover code units
certify scan

# 3. Run certification
certify certify

# 4. View your report card
certify report --format full
```

Your report card is saved to `.certification/REPORT_CARD.md`.

## CLI Reference

| Command | Description |
|---------|-------------|
| `certify init` | Bootstrap `.certification/` directory with config, policies, workflows |
| `certify scan` | Discover certifiable code units, save to index |
| `certify certify` | Evaluate units against policies, collect evidence, assign status |
| `certify report` | Generate certification reports |
| `certify expire` | Mark overdue certifications as expired |
| `certify review` | Generate PR review annotation |
| `certify architect` | AI-powered 6-phase architectural review |
| `certify doctor` | Check environment, tools, analysis tiers, AI providers |
| `certify onboard` | Interactive step-by-step onboarding guide |
| `certify models` | List available models from AI provider |
| `certify review` | Generate PR review annotation |
| `certify version` | Show version information |

### `certify certify` Flags

| Flag | Description |
|------|-------------|
| `--skip-agent` | Deterministic only, no LLM review |
| `--batch N` | Process N units per run (0=all) |
| `--reset-queue` | Rebuild work queue from scratch |
| `--target path` | Only certify units under given path(s) |
| `--diff-base ref` | Only certify files changed since git ref |
| `--path dir` | Repository root (default: current directory) |

### `certify report` Flags

| Flag | Description |
|------|-------------|
| `--format text` | Quick terminal summary (default) |
| `--format card` | Visual report card box in terminal |
| `--format full` | Complete per-unit report card (markdown) |
| `--format json` | Full report as machine-readable JSON |
| `--format site` | Interactive HTML site (alias: `--site`) |
| `--site` | Shorthand for `--format site` |
| `--badge` | Print shields.io badge markdown for your README |
| `--output file` | Write to file instead of stdout |
| `--detailed` | Add dimension breakdowns to text format |
| `--path dir` | Repository root |

## Report Formats

### Report Card (`--format full`)

The **primary output** — a complete markdown document with:
- Summary card (overall grade, pass rate, unit counts)
- Dimension averages across all 9 quality dimensions
- Per-language breakdown with score ranges
- Every unit listed with type, grade, score, status, expiry
- Expandable detail for units with observations

Saved automatically to `.certification/REPORT_CARD.md`.

### Interactive Site (`--site`)

A self-contained static HTML site for browsing large codebases:

```bash
certify report --site
open .certification/site/index.html
```

Generates to `.certification/site/` with:
- **Dashboard** — summary stats, grade distribution, dimension averages, language breakdown, top issues, package listing
- **Package pages** — per-directory roll-ups with sortable unit tables
- **Unit pages** — full detail with dimension scores, AI observations, suggestions, actions, prev/next navigation
- **Client-side search** — instant fuzzy search across all units via embedded JSON index

Features:
- Works offline via `file://` protocol — no server needed
- Zero external dependencies — all CSS/JS embedded
- Dark mode support and mobile responsive
- At 559 units, generates 584 pages in under 2 seconds

### Badge

```bash
certify report --badge
```

Outputs a shields.io badge for your README that links to the report card.

### JSON (`--format json`)

Machine-readable version of the full report — same data, JSON format.

## Incremental Processing

The certify command uses a persistent work queue:

```bash
# Process 20 units at a time (useful with rate-limited agent review)
certify certify --batch 20
# ... wait, then resume from where you left off
certify certify --batch 20
# ... repeat until "Queue complete!"

# Start over
certify certify --reset-queue
```

## Agent-Assisted Review

Optional LLM-powered code review using open-weight models via [OpenRouter](https://openrouter.ai/).

**Auto-detection (recommended):** Just set `OPENROUTER_API_KEY` in your environment or as a GitHub secret. Certify automatically enables conservative AI mode — prescreen-only, free-tier models, 10k token budget.

**Full pipeline:** For more control, explicitly configure in `.certification/config.yml`:

```yaml
agent:
  enabled: true
  provider:
    type: openrouter
    base_url: https://openrouter.ai/api/v1
    api_key_env: OPENROUTER_API_KEY
```

**Disable:** Set `agent: enabled: false` in config, or use `--skip-agent` for a single run.

Models (Apache 2.0 licensed, fine-tunable):
- **Primary**: `qwen/qwen3-coder:free` — code-specialized, 262k context
- **Fallback**: `mistralai/mistral-nemo` — 12B, clean JSON output

Agent review is always optional. The system works fully without it.

## GitHub Actions

`certify init` creates three workflow files:
- **`certification-pr.yml`** — Runs on PRs, certifies changed files
- **`certification-nightly.yml`** — Nightly sweep for expired certs
- **`certification-weekly.yml`** — Weekly full certification + report card

## Further Reading

- [Architecture](architecture.md) — System design and package structure
- [Policy Authoring](policy-authoring.md) — Writing custom policy packs
- [Troubleshooting](troubleshooting.md) — Common issues and solutions
