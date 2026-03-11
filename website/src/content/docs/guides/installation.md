---
title: Installation
description: Install Certify in under a minute.
---

## Requirements

- **Go 1.22+** — [go.dev](https://go.dev)
- **Git** — [git-scm.com](https://git-scm.com)

### Optional

- **golangci-lint** — Enhanced lint evidence ([golangci-lint.run](https://golangci-lint.run))
- **gh CLI** — GitHub PR/issue integration ([cli.github.com](https://cli.github.com))

## Install from Go Module Registry

```bash
go install github.com/iksnae/code-certification/cmd/certify@latest
```

This installs the `certify` binary to your `$GOPATH/bin`.

## Install from GitHub Release

Download pre-built binaries from [GitHub Releases](https://github.com/iksnae/code-certification/releases):

```bash
# macOS (Apple Silicon)
curl -L https://github.com/iksnae/code-certification/releases/latest/download/certify-darwin-arm64 -o certify

# macOS (Intel)
curl -L https://github.com/iksnae/code-certification/releases/latest/download/certify-darwin-amd64 -o certify

# Linux (x86_64)
curl -L https://github.com/iksnae/code-certification/releases/latest/download/certify-linux-amd64 -o certify

# Linux (ARM64)
curl -L https://github.com/iksnae/code-certification/releases/latest/download/certify-linux-arm64 -o certify

chmod +x certify
sudo mv certify /usr/local/bin/
```

## Build from Source

```bash
git clone https://github.com/iksnae/code-certification.git
cd code-certification
go build -o certify ./cmd/certify/
```

Move the binary somewhere in your `$PATH`:

```bash
sudo mv certify /usr/local/bin/
```

Or use the project's `Justfile`:

```bash
just build    # builds to build/bin/certify
```

## Verify Installation

```bash
certify version
```

You should see:

```
certify v0.9.0 (commit: a04d3f3, built: 2026-03-11T16:45:00Z)
```

## VSCode Extension

See grades inline, open the dashboard, and configure AI providers — all from the editor.

```bash
code --install-extension iksnae.certify-vscode
```

[Extension guide →](/code-certification/guides/vscode-extension/)

## Next Steps

Start the guided onboarding:

```bash
cd your-repo
certify onboard
```

Or jump straight into the [Quick Start →](/code-certification/guides/quickstart/)
