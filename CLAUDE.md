# Code Certification System

## Project Overview

The Code Certification System is a **repository governance platform** implemented as a standalone **Go CLI** (`certify`). It continuously evaluates software code against versioned policies and assigns time-bound certification status to meaningful code units.

**Current State**: Pre-implementation. Documentation and tooling only. No Go source code yet.

## Product Documentation

- **`PRD.md`** — Full product requirements (27 sections: executive summary, problem statement, vision, architecture, certification model, expiry model, GitHub integration, bootstrap, pipeline, reporting, language support, technology stack, MVP scope)
- **`FEATURES.md`** — Feature acceptance checklist (27 sections, 200+ measurable criteria covering: product foundation, bootstrap, configuration, policy-as-code, unit discovery, evidence collection, evaluation engine, certification dimensions, trust ledger, expiry, invalidation, PR workflow, scheduled workflows, issue sync, agent review, governance, reporting, CLI, adapters, security, storage, rollout, operational quality, architecture)
- **`STORIES.md`** — User stories organized by epic (15 epics: bootstrap, discovery, configuration, evidence, evaluation, trust decay, invalidation, PR workflow, scheduled workflows, remediation, reporting, agent assistance, governance, CLI, operational quality)

## Spec Output

Implementation plans and architecture specs go in `specs/`.

## Development Approach

**TDD — Test-Driven Development**. Every feature is implemented test-first:
1. Write a failing test
2. Write minimal code to pass the test
3. Refactor

No code is written without a test that demands it.

## Build Commands

```bash
just doctor     # Verify development environment
just build      # Build CLI → build/bin/certify
just test       # Run all tests
just lint       # Run golangci-lint
just fmt        # Check gofmt compliance
just vet        # Run go vet
just check      # All quality gates (fmt + vet + lint + test)
just run        # Run CLI with args: just run -- --help
just clean      # Remove build artifacts
just cover      # Generate test coverage report
```

## Key Architecture Decisions

- **Go CLI**: Single binary (`certify`), no daemon, no GUI
- **OpenRouter for agent-assisted review**: Free-tier LLM models via OpenAI-compatible API. Optional — system works fully without it.
- **Policy-as-code**: Versioned YAML policy packs in `.certification/policies/`
- **Language-agnostic**: Core engine is language-independent; language-specific adapters for discovery and evidence
- **Repository-native**: All certification state lives in the target repo's `.certification/` directory
- **GitHub integration**: PR checks, scheduled recertification, remediation issue creation

## Project Structure

```
code-certification/
├── PRD.md                    # Product requirements
├── FEATURES.md               # Feature acceptance checklist
├── STORIES.md                # User stories
├── CLAUDE.md                 # This file — agent orientation
├── README.md                 # Project overview
├── justfile                  # Build, test, lint recipes
├── specs/                    # Implementation plans
├── extensions/               # Pi framework extensions (project-agnostic)
├── .claude/                  # Claude Code commands and skills
├── .codex/                   # Codex commands (symlinked to .claude/)
└── .pi/                      # Pi agent config, themes, skills
```

When Go source code is added (via TDD), the structure will grow:

```
├── cmd/certify/              # CLI entry point
├── internal/                 # Core packages
│   ├── engine/               # Certification engine
│   ├── policy/               # Policy loading and evaluation
│   ├── evidence/             # Evidence collection and normalization
│   ├── agent/                # OpenRouter agent-assisted review
│   ├── discovery/            # Code unit discovery
│   ├── record/               # Certification record management
│   └── report/               # Report generation
├── .certification/           # Example/test certification config
│   ├── config.yml
│   ├── policies/
│   ├── prompts/
│   └── records/
└── go.mod                    # Go module (created with first test)
```

## GitHub Issue Tracking

Implementation is tracked via **GitHub Issues** as epics. Each epic maps to a phase in `specs/v1-implementation.md`.

| Epic | Phase | Issue |
|------|-------|-------|
| Domain Foundation | Phase 1 | [#1](https://github.com/iksnae/code-certification/issues/1) |
| Configuration & Policy | Phase 2 | [#2](https://github.com/iksnae/code-certification/issues/2) |
| Unit Discovery & Indexing | Phase 3 | [#3](https://github.com/iksnae/code-certification/issues/3) |
| Evidence Collection | Phase 4 | [#4](https://github.com/iksnae/code-certification/issues/4) |
| Certification Engine | Phase 5 | [#5](https://github.com/iksnae/code-certification/issues/5) |
| Agent-Assisted Review | Phase 6 | [#6](https://github.com/iksnae/code-certification/issues/6) |
| Records, Reports & CLI | Phase 7 | [#7](https://github.com/iksnae/code-certification/issues/7) |
| GitHub Integration | Phase 8 | [#8](https://github.com/iksnae/code-certification/issues/8) |

### Development Workflow

When implementing any step from the v1 plan:

1. **Reference the epic** — check the relevant GitHub issue for context and task checklist
2. **TDD cycle** — write failing test → implement → refactor
3. **Update FEATURES.md** — check off completed criteria as you implement them (change `- [ ]` to `- [x]`)
4. **Update the epic** — check off completed steps in the GitHub issue
5. **Commit with references** — include `Refs #N` in commit messages to link to the epic (e.g., `feat: unit discovery — Go adapter via go/ast. Refs #3`)
6. **Push** — keep GitHub issues and code in sync

### FEATURES.md as Living Checklist

`FEATURES.md` is the **source of truth** for what's done and what's not. As you implement features:
- Check off criteria: `- [ ]` → `- [x]`
- Each epic's GitHub issue lists which FEATURES sections it covers
- At any point, `grep -c '\- \[x\]' FEATURES.md` shows total completed criteria
- At any point, `grep -c '\- \[ \]' FEATURES.md` shows total remaining criteria

This keeps the acceptance checklist synchronized with actual implementation progress.

## Agent-Assisted Review

The certification engine optionally uses **OpenRouter** (free-tier models) for LLM-powered code review. See `specs/project-bootstrap-justfile-commands-initial-commit.md` for the complete architecture including:
- OpenRouter API contract and error handling
- Free model catalog with capabilities (7 models researched)
- Task-to-model routing (prescreen → review → scoring → decision → remediation)
- Structured output schemas for dimension scoring
- Provider abstraction interface
- GitHub repository secret setup for `OPENROUTER_API_KEY`

The system works fully without agent assistance. If `OPENROUTER_API_KEY` is not set, agent review is skipped and certification proceeds with deterministic evidence only.
