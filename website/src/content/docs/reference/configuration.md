---
title: Configuration
description: Configure Certify for your repository.
---

## Config File

`certify init` creates `.certification/config.yml`:

```yaml
# Certify — Configuration
mode: advisory

scope:
  include: []
  exclude:
    - "vendor/**"
    - "node_modules/**"
    - "dist/**"
    - "build/**"
    - "testdata/**"
    - "**/*_test.go"
    - "**/*.test.ts"
    - "**/*.spec.ts"

agent:
  enabled: false

expiry:
  default_window_days: 90
  min_window_days: 7
  max_window_days: 365

issues:
  enabled: false
```

## Sections

### `mode`

| Value | Behavior |
|-------|----------|
| `advisory` | Reports results but does not block. Default. |
| `enforcing` | Can block PRs and fail CI on certification failures. |

### `scope`

Controls which files are included in certification.

- **`include`** — Glob patterns to include. Empty means everything.
- **`exclude`** — Glob patterns to exclude. Matched files are skipped during discovery.

### `agent`

Optional LLM-assisted review. See [Agent-Assisted Review](/code-certification/advanced/agent-review/).

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
    fallback: mistralai/mistral-nemo
  rate_limit:
    requests_per_minute: 20
    retry_max: 3
```

### `expiry`

Controls certification window duration.

| Field | Default | Description |
|-------|---------|-------------|
| `default_window_days` | 90 | Standard certification window |
| `min_window_days` | 7 | Minimum window (high-risk code) |
| `max_window_days` | 365 | Maximum window (stable code) |

Risk factors automatically adjust the window — high-churn code gets shorter windows.

### `issues`

GitHub issue integration for remediation tracking.

```yaml
issues:
  enabled: true
```

When enabled, failing certifications create GitHub issues via the `gh` CLI.
