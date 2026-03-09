# Code Certification System

A repository governance platform that continuously evaluates software code against defined engineering policies and assigns **time-bound certification status** to meaningful code units.

Unlike CI pipelines that only validate whether code passes at a specific moment, Code Certification establishes **explicit trust with expiration**, ensuring that code quality, security posture, maintainability, and architectural compliance are periodically re-evaluated as standards evolve.

## Status

**Pre-implementation** — Documentation and tooling only. No Go source code yet.

Development follows **strict TDD**: every feature is implemented test-first (test → implement → refactor).

## Prerequisites

- [Go](https://go.dev/) 1.22+
- [golangci-lint](https://golangci-lint.run/)
- [just](https://github.com/casey/just) command runner
- [git](https://git-scm.com/)
- [gh](https://cli.github.com/) GitHub CLI

## Quick Start

```bash
# Verify your development environment
just doctor

# Run tests (when they exist)
just test

# Build the CLI (when source exists)
just build

# Run all quality checks
just check
```

## Documentation

- **[PRD.md](PRD.md)** — Full product requirements document
- **[FEATURES.md](FEATURES.md)** — Feature acceptance checklist (27 sections, 200+ criteria)
- **[STORIES.md](STORIES.md)** — User stories organized by epic
- **[specs/](specs/)** — Implementation plans and architecture specs

## Architecture

The system is built as a **standalone Go CLI** (`certify`) that:

- Discovers certifiable code units in a repository
- Evaluates units against versioned policy packs
- Collects evidence from linters, type checkers, test runners, and static analyzers
- Optionally performs agent-assisted review via OpenRouter (free-tier LLM models)
- Assigns time-bound certification status with expiration
- Stores certification records in a repository-local `.certification/` directory
- Integrates with GitHub workflows for PR checks and scheduled recertification
- Creates remediation issues for failing units

## Agent-Assisted Review

The certification engine includes optional LLM-powered code review using **OpenRouter** with free-tier models. Agent review supplements deterministic evidence — it never overrides tool results. The system operates fully without agent assistance when no API key is configured.

See [specs/project-bootstrap-justfile-commands-initial-commit.md](specs/project-bootstrap-justfile-commands-initial-commit.md) for the complete agent architecture, model catalog, and integration design.

## License

TBD
