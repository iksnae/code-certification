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

## Verify Installation

```bash
certify version
```

You should see:

```
certify v0.1.0 (abc1234) built 2026-03-09T12:00:00Z
```

## VSCode Extension

See grades inline, open the dashboard, and configure AI providers — all from the editor.

```bash
code --install-extension iksnae.certify-vscode
```

[Extension guide →](/code-certification/guides/vscode-extension/)

## Next Steps

[Quick Start — certify your first repository →](/code-certification/guides/quickstart/)
