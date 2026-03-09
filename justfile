# Code Certification System — Justfile
# Go CLI for repository governance with time-bound code certification

# ============================================================================
# Variables
# ============================================================================

build_dir := justfile_directory() / "build/bin"

# ============================================================================
# Build Commands
# ============================================================================

# Build the certify CLI into build/bin/
build:
  #!/usr/bin/env bash
  set -eu -o pipefail
  if [ ! -f cmd/certify/main.go ]; then
    echo "⚠ No source code yet. Run tests first (TDD)."
    exit 0
  fi
  mkdir -p "{{build_dir}}"
  go build -o "{{build_dir}}/certify" ./cmd/certify
  echo "built certify → {{build_dir}}/certify"

# Remove build artifacts
clean:
  rm -rf build/

# ============================================================================
# Test Commands
# ============================================================================

# Run all tests
test:
  #!/usr/bin/env bash
  set -eu -o pipefail
  if ! find . -name '*_test.go' -not -path './vendor/*' | grep -q .; then
    echo "⚠ No test files yet. Start with TDD."
    exit 0
  fi
  go test ./...

# Run all tests with verbose output
test-verbose:
  #!/usr/bin/env bash
  set -eu -o pipefail
  if ! find . -name '*_test.go' -not -path './vendor/*' | grep -q .; then
    echo "⚠ No test files yet. Start with TDD."
    exit 0
  fi
  go test -v ./...

# Generate test coverage report
cover:
  #!/usr/bin/env bash
  set -eu -o pipefail
  if ! find . -name '*_test.go' -not -path './vendor/*' | grep -q .; then
    echo "⚠ No test files yet. Start with TDD."
    exit 0
  fi
  go test -coverprofile=coverage.out ./...
  go tool cover -html=coverage.out -o coverage.html
  echo "Coverage report: coverage.html"

# ============================================================================
# Lint & Quality Commands
# ============================================================================

# Run golangci-lint
lint:
  #!/usr/bin/env bash
  set -eu -o pipefail
  if ! find . -name '*.go' -not -path './vendor/*' | grep -q .; then
    echo "⚠ No Go files yet."
    exit 0
  fi
  golangci-lint run ./...

# Check gofmt compliance
fmt:
  #!/usr/bin/env bash
  set -eu -o pipefail
  if ! find . -name '*.go' -not -path './vendor/*' | grep -q .; then
    echo "⚠ No Go files yet."
    exit 0
  fi
  unformatted=$(gofmt -l .)
  if [ -n "$unformatted" ]; then
    echo "Unformatted files:"
    echo "$unformatted"
    exit 1
  fi
  echo "✓ All files formatted"

# Run go vet
vet:
  #!/usr/bin/env bash
  set -eu -o pipefail
  if ! find . -name '*.go' -not -path './vendor/*' | grep -q .; then
    echo "⚠ No Go files yet."
    exit 0
  fi
  go vet ./...

# Run all quality checks (fmt + vet + lint + test)
check: fmt vet lint test
  @echo "✓ All checks passed"

# ============================================================================
# Run Commands
# ============================================================================

# Run the certify CLI with arguments (e.g., just run -- --help)
run *ARGS:
  #!/usr/bin/env bash
  set -eu -o pipefail
  if [ ! -f cmd/certify/main.go ]; then
    echo "⚠ No source code yet. Run tests first (TDD)."
    exit 0
  fi
  go run ./cmd/certify {{ARGS}}

# ============================================================================
# Workspace Health
# ============================================================================

# Check development environment requirements
doctor:
  #!/usr/bin/env bash
  echo "Code Certification — Environment Check"
  echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
  echo ""
  echo "Required Tools:"
  command -v go >/dev/null 2>&1 && echo "  ✓ go $(go version | awk '{print $3}')" || echo "  ✗ go not found"
  command -v golangci-lint >/dev/null 2>&1 && echo "  ✓ golangci-lint $(golangci-lint --version 2>&1 | head -1 | awk '{print $4}')" || echo "  ✗ golangci-lint not found"
  command -v git >/dev/null 2>&1 && echo "  ✓ git $(git --version | awk '{print $3}')" || echo "  ✗ git not found"
  command -v gh >/dev/null 2>&1 && echo "  ✓ gh $(gh --version 2>&1 | head -1 | awk '{print $3}')" || echo "  ✗ gh not found"
  command -v just >/dev/null 2>&1 && echo "  ✓ just $(just --version 2>&1 | awk '{print $2}')" || echo "  ✗ just not found"
  echo ""
  echo "Optional:"
  command -v shellcheck >/dev/null 2>&1 && echo "  ✓ shellcheck available" || echo "  ⚠ shellcheck not found"
  test -n "${OPENROUTER_API_KEY:-}" && echo "  ✓ OPENROUTER_API_KEY is set" || echo "  ⚠ OPENROUTER_API_KEY not set (agent review disabled)"
  echo ""
  echo "Project State:"
  if find . -name '*.go' -not -path './vendor/*' 2>/dev/null | grep -q .; then
    echo "  Go files: present"
  else
    echo "  Go files: none yet (pre-implementation)"
  fi
  if [ -f go.mod ]; then
    echo "  Go module: $(head -1 go.mod)"
  else
    echo "  Go module: not initialized yet"
  fi

# ============================================================================
# Pi Extensions
# ============================================================================

# g1

# 1. default pi
pi:
    pi

# 2. Pure focus pi: strip footer and status line entirely
ext-pure-focus:
    pi -e extensions/pure-focus.ts

# 3. Minimal pi: model name + 10-block context meter
ext-minimal:
    pi -e extensions/minimal.ts -e extensions/theme-cycler.ts

# 4. Cross-agent pi: load commands from .claude/, .gemini/, .codex/ dirs
ext-cross-agent:
    pi -e extensions/cross-agent.ts -e extensions/minimal.ts

# 5. Purpose gate pi: declare intent before working, persistent widget, focus the system prompt on the ONE PURPOSE for this agent
ext-purpose-gate:
    pi -e extensions/purpose-gate.ts -e extensions/minimal.ts

# 6. Customized footer pi: Tool counter, model, branch, cwd, cost, etc.
ext-tool-counter:
    pi -e extensions/tool-counter.ts

# 7. Tool counter widget: tool call counts in a below-editor widget
ext-tool-counter-widget:
    pi -e extensions/tool-counter-widget.ts -e extensions/minimal.ts

# 8. Subagent widget: /sub <task> with live streaming progress
ext-subagent-widget:
    pi -e extensions/subagent-widget.ts -e extensions/pure-focus.ts -e extensions/theme-cycler.ts

# 9. TillDone: task-driven discipline — define tasks before working
ext-tilldone:
    pi -e extensions/tilldone.ts -e extensions/theme-cycler.ts

#g2

# 10. Agent team: dispatcher orchestrator with team select and grid dashboard
ext-agent-team:
    pi -e extensions/agent-team.ts -e extensions/theme-cycler.ts

# 11. System select: /system to pick an agent persona as system prompt
ext-system-select:
    pi -e extensions/system-select.ts -e extensions/minimal.ts -e extensions/theme-cycler.ts

# 12. Launch with Damage-Control safety auditing
ext-damage-control:
    pi -e extensions/damage-control.ts -e extensions/minimal.ts -e extensions/theme-cycler.ts

# 13. Agent chain: sequential pipeline orchestrator
ext-agent-chain:
    pi -e extensions/agent-chain.ts -e extensions/theme-cycler.ts

#g3

# 14. Pi Pi: meta-agent that builds Pi agents with parallel expert research
ext-pi-pi:
    pi -e extensions/pi-pi.ts -e extensions/theme-cycler.ts

#ext

# 15. Session Replay: scrollable timeline overlay of session history (legit)
ext-session-replay:
    pi -e extensions/session-replay.ts -e extensions/minimal.ts

# 16. Theme cycler: Ctrl+X forward, Ctrl+Q backward, /theme picker
ext-theme-cycler:
    pi -e extensions/theme-cycler.ts -e extensions/minimal.ts

# utils

# Open pi with one or more stacked extensions in a new terminal: just open minimal tool-counter
open +exts:
    #!/usr/bin/env bash
    args=""
    for ext in {{exts}}; do
        args="$args -e extensions/$ext.ts"
    done
    cmd="cd '{{justfile_directory()}}' && pi$args"
    escaped="${cmd//\\/\\\\}"
    escaped="${escaped//\"/\\\"}"
    osascript -e "tell application \"Terminal\" to do script \"$escaped\""

# Open every extension in its own terminal window
all:
    just open pi
    just open pure-focus
    just open minimal theme-cycler
    just open cross-agent minimal
    just open purpose-gate minimal
    just open tool-counter
    just open tool-counter-widget minimal
    just open subagent-widget pure-focus theme-cycler
    just open tilldone theme-cycler
    just open agent-team theme-cycler
    just open system-select minimal theme-cycler
    just open damage-control minimal theme-cycler
    just open agent-chain theme-cycler
    just open pi-pi theme-cycler

# ============================================================================
# Claude Code Commands
# ============================================================================

# Claude setup with Opus model (full capabilities)
cld-super:
  claude --model opus --dangerously-skip-permissions --init

# Deterministic workspace setup (fast, CI-friendly)
cld-init:
  claude --model haiku --dangerously-skip-permissions --init

# Deterministic workspace maintenance
cld-maintain:
  claude --model haiku --dangerously-skip-permissions --maintenance

# Agentic workspace setup (with agent analysis)
cld-init-agent:
  claude --model haiku --dangerously-skip-permissions --init "/install"

# Agentic workspace setup (interactive)
cld-init-interactive:
  claude --model haiku --dangerously-skip-permissions --init "/install true"

# Agentic workspace maintenance
cld-maintain-agent:
  claude --model haiku --dangerously-skip-permissions --maintenance "/maintenance"

# ============================================================================
# OpenAI Codex Commands
# ============================================================================

# Codex setup with full autonomy (maximum capability)
cdx-super:
  codex --dangerously-bypass-approvals-and-sandbox

# Deterministic workspace setup (fast, CI-friendly)
cdx-init:
  codex --full-auto "/install"

# Deterministic workspace maintenance
cdx-maintain:
  codex --full-auto "/maintenance"

# Agentic workspace setup (with agent analysis)
cdx-init-agent:
  codex --full-auto "/install"

# Agentic workspace setup (interactive)
cdx-init-interactive:
  codex "/install true"

# Agentic workspace maintenance
cdx-maintain-agent:
  codex --full-auto "/maintenance"
