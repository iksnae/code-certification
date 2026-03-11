# certify — code certification CLI
# https://github.com/iksnae/code-certification

version  := "0.12.0"
build_dir := justfile_directory() / "build/bin"

# ── Build ────────────────────────────────────────────────────────────────────

# Build the certify CLI
build:
  #!/usr/bin/env bash
  set -euo pipefail
  mkdir -p "{{build_dir}}"
  go build -ldflags "\
    -X main.Version={{version}} \
    -X main.Commit=$(git rev-parse --short HEAD) \
    -X main.Date=$(date -u +%Y-%m-%dT%H:%M:%SZ)" \
    -o "{{build_dir}}/certify" ./cmd/certify
  echo "✓ built → {{build_dir}}/certify"

# Install certify to $GOPATH/bin
install:
  go install -ldflags "\
    -X main.Version={{version}} \
    -X main.Commit=$(git rev-parse --short HEAD) \
    -X main.Date=$(date -u +%Y-%m-%dT%H:%M:%SZ)" \
    ./cmd/certify

# Remove build artifacts
clean:
  rm -rf build/ coverage.out coverage.html

# ── Quality ──────────────────────────────────────────────────────────────────

# Run all tests
test:
  go test ./... -count=1

# Run all tests with verbose output
test-verbose:
  go test ./... -count=1 -v

# Generate test coverage report
cover:
  go test -coverprofile=coverage.out ./...
  go tool cover -html=coverage.out -o coverage.html
  @echo "✓ Coverage report → coverage.html"

# Run golangci-lint
lint:
  golangci-lint run ./...

# Check gofmt compliance
fmt:
  @test -z "$$(gofmt -l .)" || (gofmt -l . && exit 1)
  @echo "✓ All files formatted"

# Run go vet
vet:
  go vet ./...

# Run all quality checks: fmt → vet → lint → test
check: fmt vet lint test
  @echo "✓ All checks passed"

# ── Certify ──────────────────────────────────────────────────────────────────

# Self-certify this repository (scan → certify → report)
certify: build
  ./build/bin/certify scan
  ./build/bin/certify certify --skip-agent --reset-queue
  ./build/bin/certify report --format card

# Generate the full report card
report-card: build
  ./build/bin/certify report --format full
  @echo "✓ Report card → .certification/REPORT_CARD.md"

# Generate static certification site
site: build
  ./build/bin/certify report --format site
  @echo "✓ Open .certification/site/index.html in a browser"

# ── Doctor ───────────────────────────────────────────────────────────────────

# Check development environment
doctor:
  #!/usr/bin/env bash
  echo "certify — environment check"
  echo "════════════════════════════"
  echo ""
  echo "Required:"
  command -v go        >/dev/null 2>&1 && echo "  ✓ go $(go version | awk '{print $3}')" || echo "  ✗ go — https://go.dev"
  command -v git       >/dev/null 2>&1 && echo "  ✓ git $(git --version | awk '{print $3}')" || echo "  ✗ git — https://git-scm.com"
  echo ""
  echo "Optional:"
  command -v golangci-lint >/dev/null 2>&1 && echo "  ✓ golangci-lint" || echo "  · golangci-lint — https://golangci-lint.run (enhanced lint evidence)"
  command -v gh            >/dev/null 2>&1 && echo "  ✓ gh CLI"        || echo "  · gh — https://cli.github.com (PR/issue integration)"
  test -n "${OPENROUTER_API_KEY:-}" && echo "  ✓ OPENROUTER_API_KEY set" || echo "  · OPENROUTER_API_KEY not set (agent review disabled)"
