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

When Certify detects an available provider, it automatically enables **conservative mode**:
- **Prescreen only** — single cheap LLM call per file
- **Free-tier models** — zero cost by default (`qwen/qwen3-coder:free`)
- **10k token budget** — processes ~200 files before stopping
- **Circuit breaker** — 3 consecutive failures → falls back to deterministic
- **Scan suggestions** — `certify scan` prints AI-powered policy recommendations
- **Multi-provider fallback** — cloud providers tried first, then local

### Supported Providers

Any OpenAI-compatible endpoint works. These are auto-detected:

| Provider | Detection | Cost | Setup |
|----------|-----------|------|-------|
| **OpenRouter** | `OPENROUTER_API_KEY` env var | Free tier + paid (200+ models) | [openrouter.ai](https://openrouter.ai) |
| **OpenAI** | `OPENAI_API_KEY` env var | Paid (~$0.15/1M tokens for gpt-4o-mini) | [platform.openai.com](https://platform.openai.com) |
| **Google AI Studio** | `GEMINI_API_KEY` env var | Free tier (Gemini 2.0 Flash) | [aistudio.google.com](https://aistudio.google.com) |
| **Groq** | `GROQ_API_KEY` env var | Free tier (30 req/min) | [groq.com](https://groq.com) |
| **Together** | `TOGETHER_API_KEY` env var | Free $1 credit | [together.ai](https://together.ai) |
| **Fireworks** | `FIREWORKS_API_KEY` env var | Free $1 credit | [fireworks.ai](https://fireworks.ai) |
| **Ollama** | `OLLAMA_HOST` env or auto-probe `localhost:11434` | Free (local) | [ollama.com](https://ollama.com) |
| **LM Studio** | `LM_STUDIO_URL` env or auto-probe `localhost:1234` | Free (local) | [lmstudio.ai](https://lmstudio.ai) |
| **vLLM** | Auto-probe `localhost:8000` | Free (local) | [vllm.ai](https://vllm.ai) |

Certify checks in this order: `OPENROUTER_API_KEY` → `CERTIFY_API_KEY` → `OPENAI_API_KEY` → `GROQ_API_KEY` → Ollama → LM Studio. Cloud providers come first, local providers serve as fallback.

You can also configure any custom OpenAI-compatible endpoint via `.certification/config.yml` or the VS Code extension settings.

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

## Related

- [Architect Review](/code-certification/advanced/architect-review/) — 6-phase AI architectural analysis (separate from per-unit agent review)
- [Workspace Mode](/code-certification/advanced/workspace/) — multi-repo certification across git submodules
