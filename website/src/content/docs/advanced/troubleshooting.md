---
title: Troubleshooting
description: Common issues and solutions.
---

## "loading index (run 'certify scan' first)"

The unit index doesn't exist yet:

```bash
certify scan
```

## "No certification records found"

Run certification before generating reports:

```bash
certify certify --skip-agent
```

## Agent review: 429 Too Many Requests

Free-tier models on OpenRouter have rate limits.

- Use `--batch 20` to process fewer units per run
- Wait a few minutes between runs (the queue saves progress)
- Set up a paid fallback model in config
- Use `--skip-agent` for deterministic-only certification

## Agent review: 401 Unauthorized

Your API key is invalid or not set:

```bash
export OPENROUTER_API_KEY=sk-or-v1-your-key-here
```

For CI, add `OPENROUTER_API_KEY` as a GitHub repository secret.

## Agent review: 402 Payment Required

Account budget exhausted on OpenRouter. The system will:

1. Abort the current model
2. Try free-tier fallback models
3. Skip agent review if all models fail

Certification continues with deterministic evidence only.

## All units show the same score

This usually means no language-specific evidence was collected. Check:

1. Is `go.mod` present? (needed for Go tools)
2. Is `go` in PATH?
3. Run `certify scan` to verify language detection
4. Check output for "Collected N repo-level evidence items"

## Queue stuck / old state

Reset the queue to start fresh:

```bash
certify certify --reset-queue
```

## "go vet" or "golangci-lint" not found

These tools are optional. Without them, the system uses available evidence. Install for better results:

```bash
# golangci-lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh
```

## Records directory too large

Each unit gets a small JSON file (~500 bytes). For very large repos (10k+ units):

1. Add `.certification/records/` to `.gitignore`
2. Use CI artifacts instead of committed records
3. Scope certification to critical paths via `scope.include`

## Getting Help

- [GitHub Issues](https://github.com/iksnae/code-certification/issues)
- `certify --help` for command reference
- `certify <command> --help` for flag details
