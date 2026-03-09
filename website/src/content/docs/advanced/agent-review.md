---
title: Agent-Assisted Review
description: Optional LLM-powered code review using open-weight models.
---

Certify can supplement deterministic evidence with LLM-powered code review. This is entirely optional — the system works fully without it.

## Auto-Detection (Recommended)

The simplest way to enable AI-assisted review: **just add your API key.**

```bash
export OPENROUTER_API_KEY=sk-or-v1-your-key-here
```

For CI, add `OPENROUTER_API_KEY` as a GitHub repository secret. That's it — no config changes needed.

When Certify detects an API key in the environment, it automatically enables **conservative mode**:
- **Prescreen only** — single cheap LLM call per file
- **Free-tier models** — zero cost by default (`qwen/qwen3-coder:free`)
- **10k token budget** — processes ~200 files before stopping
- **Circuit breaker** — 3 consecutive failures → falls back to deterministic
- **Scan suggestions** — `certify scan` prints AI-powered policy recommendations

Certify checks `OPENROUTER_API_KEY` first, then `CERTIFY_API_KEY`.

### Disable Auto-Detection

To explicitly disable AI even when a key is present:

```yaml
agent:
  enabled: false
```

The `--skip-agent` CLI flag also overrides auto-detection for any single run.

## Full Pipeline (Advanced)

For more control, explicitly configure the full multi-stage agent pipeline in `.certification/config.yml`:

```yaml
agent:
  enabled: true
  provider:
    type: openrouter
    base_url: https://openrouter.ai/api/v1
    api_key_env: OPENROUTER_API_KEY
  models:
    prescreen: qwen/qwen3-coder:free
    review: qwen/qwen3-coder:free
    scoring: qwen/qwen3-coder:free
    fallback: mistralai/mistral-nemo
  rate_limit:
    requests_per_minute: 20
    retry_max: 3
```

This enables the 3-stage pipeline (prescreen → review → scoring) with a 50k token budget.

## How It Works

When agent review is enabled, Certify runs a multi-stage pipeline:

1. **Prescreen** — Quick assessment of code quality (should we do a full review?)
2. **Review** — Detailed code review with observations
3. **Scoring** — Agent-assigned dimension scores

Agent scores are blended with deterministic evidence. Agent review **supplements** — it never overrides tool results.

## Models

All models are selected for open-weight licensing (Apache 2.0) and fine-tunability:

| Model | License | Context | Cost |
|-------|---------|---------|------|
| `qwen/qwen3-coder:free` | Apache 2.0 | 262k | Free |
| `mistralai/mistral-nemo` | Apache 2.0 | 128k | ~$0.002/run |

The model chain tries free models first, then falls through to paid models on rate limits or failures.

## Model Attribution

Every certification record that uses agent review records which models contributed:

```json
{
  "source": "deterministic+agent:qwen/qwen3-coder:free,mistralai/mistral-nemo"
}
```

## Circuit Breaker

If 5 consecutive agent calls fail, the circuit breaker opens and agent review is automatically disabled for the remainder of the run. Certification continues with deterministic evidence only.

## Rate Limits

Free-tier models have rate limits. Strategies:

- Use `--batch 20` to process fewer units per run
- The work queue saves progress — resume later
- Paid fallback models handle overflow automatically
- Use `--skip-agent` for deterministic-only runs

## Without Agent Review

When agent review is disabled or unavailable, Certify uses deterministic evidence only:

- Lint results (go vet, golangci-lint)
- Test status
- Git history and churn metrics
- AST-based complexity analysis
- Code metrics (line count, documentation presence)

This produces reliable, reproducible certification results.
