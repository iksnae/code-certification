# Feature: Project Bootstrap — Justfile, Commands & Initial Commit

## Feature Description
The code-certification project is a brand-new repository that currently contains only product documentation (PRD.md, FEATURES.md, STORIES.md) and agent/infrastructure files copied verbatim from the Khaos Machine workspace. Every operational file — justfile, Claude commands, pi agent prompts, skills, settings, and cached agent sessions — references a completely different project's architecture, submodules, binaries, and domain concepts. None of it applies here.

This feature strips all Khaos Machine content, replaces it with tooling purpose-built for the Code Certification System (a standalone Go CLI), documents the agent-assisted review architecture with concrete OpenRouter integration details, removes the `plan_w_team.md` orchestration command, and creates the initial git commit so TDD-driven development can begin.

There will be **no Go source code** in this commit. Every feature will be Test Driven — tests first, implementation second, refactor third. This commit establishes the project skeleton, tooling, and design documentation only.

Key deliverables:
1. **New justfile** — lean recipes for a single Go CLI project (build, test, lint, run, doctor)
2. **Updated Claude commands** — all spec output paths changed to `specs/`, all Khaos references removed, `plan_w_team.md` removed
3. **Updated pi agent prompts** — `prime.md` rewritten for Code Certification System
4. **Updated settings** — `settings.local.json` permissions scoped to this project only
5. **New `.gitignore`** — standard Go project ignores
6. **New `CLAUDE.md`** — project context for agent orientation
7. **New `README.md`** — project overview and getting started
8. **Agent-Assisted Review architecture** — OpenRouter integration design with researched free model catalog, provider abstraction, structured output schemas, and configuration scaffolding
9. **Purged Khaos content** — all skills, sessions, settings, and references to the Khaos Machine workspace removed completely
10. **Initial commit** — all files committed to `main` branch

## User Story
As a developer starting work on the Code Certification System
I want the project's tooling, commands, and justfile to accurately reflect this project
So that I can begin TDD-driven development with correct build commands, agent context, and a clean git history

## Problem Statement
Every operational file in this repository belongs to a different project (Khaos Machine). The justfile builds 5 Go binaries across 6 submodules. The Claude `prime.md` describes "an AI-powered storytelling platform." The settings grant permissions to paths like `/Users/k/Spikes/khaos.machine/`. The skills teach agents about Vue/Nuxt components and macOS installer packaging. The `.pi/agent-sessions/` directory contains cached conversations about screenplay analysis. None of this is relevant.

Additionally, the agent-assisted review capability — a core differentiator of the Code Certification System (PRD §16, FEATURES §16, STORIES §11) — has no design documentation, no model selection rationale, and no integration architecture. The system intends to use OpenRouter with free-tier models, but which models, for what tasks, with what API contract, error handling, and fallback strategy is undocumented.

Development cannot begin until all tooling reflects the actual project and the agent architecture is designed with enough detail to be TDD-implementable.

## Solution Statement
1. **Complete purge** — Remove every file, reference, skill, session, and permission that mentions Khaos Machine. No trace remains.
2. **Justfile rewrite** — Single Go CLI project recipes: `build`, `test`, `lint`, `run`, `clean`, `doctor`, `check`, plus pi extension launchers.
3. **Command updates** — All Claude commands point to `specs/` for plan output. `plan_w_team.md` removed. `prime.md` fully rewritten. `install.md` and `maintenance.md` rewritten for Go project lifecycle.
4. **Agent architecture document** — Concrete design based on live OpenRouter API research: model catalog with capabilities, task-to-model routing, provider interface, structured output schemas, rate limiting, graceful degradation, GitHub repo secret wiring for `OPENROUTER_API_KEY`.
5. **Project scaffolding** — `.gitignore`, `CLAUDE.md`, `README.md`. No `go.mod` yet (created during first TDD cycle when writing first test).
6. **Initial commit** — Clean `main` branch with everything staged.

## Agent-Assisted Review Architecture

### Overview

The Code Certification System includes **optional agent-assisted review** (PRD §16, FEATURES §16, STORIES §11). When enabled, the certification engine sends code units to an LLM for contextual review commentary, remediation suggestions, and quality dimension scoring. Agent output supplements — but never overrides — deterministic evidence from linters, type checkers, test runners, and static analyzers.

The system uses **OpenRouter** as the LLM service provider, consuming **free-tier models** exclusively. OpenRouter provides an OpenAI-compatible API, meaning the Go HTTP client targets a single well-documented API shape that can be swapped to direct OpenAI, Anthropic, or local models by changing base URL and API key.

### OpenRouter API Contract

**Base URL**: `https://openrouter.ai/api/v1`
**Auth**: `Authorization: Bearer $OPENROUTER_API_KEY`
**Endpoint**: `POST /chat/completions` (OpenAI-compatible)

**Required request headers**:
```
Authorization: Bearer $OPENROUTER_API_KEY
Content-Type: application/json
HTTP-Referer: https://github.com/code-certification/code-certification
X-Title: Code Certification System
```

The `HTTP-Referer` and `X-Title` headers are recommended by OpenRouter for free-tier priority routing.

**Request body** (OpenAI-compatible):
```json
{
  "model": "qwen/qwen3-coder:free",
  "messages": [
    {"role": "system", "content": "..."},
    {"role": "user", "content": "..."}
  ],
  "tools": [...],
  "tool_choice": "auto",
  "temperature": 0.2,
  "max_tokens": 4096,
  "response_format": {"type": "json_schema", "json_schema": {...}}
}
```

**Response shape** (OpenAI-compatible):
```json
{
  "id": "gen-...",
  "model": "qwen/qwen3-coder:free",
  "choices": [{
    "index": 0,
    "message": {
      "role": "assistant",
      "content": "...",
      "tool_calls": [...]
    },
    "finish_reason": "stop"
  }],
  "usage": {"prompt_tokens": 1234, "completion_tokens": 567, "total_tokens": 1801}
}
```

**Error codes**:
| Code | Meaning | Engine behavior |
|------|---------|----------------|
| 401 | Invalid/missing API key | Log error, disable agent review for run |
| 402 | Insufficient credits | Log warning, skip agent review |
| 429 | Rate limited | Retry with exponential backoff (max 3 retries) |
| 502 | Model temporarily unavailable | Fall back to next model in chain |
| 503 | Service unavailable | Fall back to next model, then skip agent review |

**Rate limit headers** (returned on responses):
```
x-ratelimit-limit-requests: 20
x-ratelimit-remaining-requests: 18
x-ratelimit-reset-requests: 2s
```

Free tier limits: **~20 requests/minute per model**, **~200 requests/day per model** (varies by model and load). The engine must implement rate limiting proactively rather than relying solely on 429 responses.

### Free Model Catalog (Researched March 2026)

All models below are confirmed free (`pricing.prompt: "0"`, `pricing.completion: "0"`) via live query of `GET https://openrouter.ai/api/v1/models`.

#### Tier 1: Code Analysis (Primary)

**`qwen/qwen3-coder:free`** — Qwen3 Coder 480B A35B (MoE, 35B active)
- Context: 262,000 tokens | Max completion: 262,000 tokens
- Capabilities: tool calling (`tools`, `tool_choice`)
- Does NOT support: `structured_outputs`, `response_format` — must use tool calling or prompt-based JSON for structured responses
- Optimized for: agentic coding tasks, function calling, long-context reasoning over repositories
- **Use for**: Code review commentary, readability/maintainability analysis, remediation suggestions
- Default params: none specified (use explicit temperature)

**`qwen/qwen3-next-80b-a3b-instruct:free`** — Qwen3 Next 80B A3B (MoE, 3B active)
- Context: 262,144 tokens | Max completion: not specified (use max_tokens)
- Capabilities: tool calling, **`structured_outputs`**, `response_format`
- Optimized for: deterministic instruction-following, RAG, tool use, agentic workflows
- **Use for**: Dimension scoring (structured JSON), certification record generation
- Default params: none specified

#### Tier 2: Reasoning & Decision

**`openai/gpt-oss-120b:free`** — GPT-OSS 120B (MoE, 5.1B active)
- Context: 131,072 tokens | Max completion: 131,072 tokens
- Capabilities: tool calling, **`reasoning`** (configurable chain-of-thought)
- Moderated: yes (content filtered)
- **Use for**: Borderline certification status decisions, complex multi-factor risk assessment
- Default params: none specified

**`openai/gpt-oss-20b:free`** — GPT-OSS 20B (MoE, 3.6B active)
- Context: 131,072 tokens | Max completion: 131,072 tokens
- Capabilities: tool calling, **`reasoning`**
- Moderated: yes
- **Use for**: Lighter reasoning tasks, faster inference fallback for gpt-oss-120b
- Default params: none specified

#### Tier 3: Fast & Structured

**`mistralai/mistral-small-3.1-24b-instruct:free`** — Mistral Small 3.1 24B (dense)
- Context: 128,000 tokens | Max completion: not specified
- Capabilities: tool calling, **`structured_outputs`**, `response_format`
- **Use for**: Pre-screening (does unit need deep review?), bulk observation generation, simple policy compliance
- Default params: temperature 0.3

**`arcee-ai/trinity-mini:free`** — Trinity Mini 26B (MoE, 3B active, 128 experts)
- Context: 131,072 tokens | Max completion: not specified
- Capabilities: tool calling, **`structured_outputs`**, `response_format`, **`reasoning`**
- Engineered for: efficient reasoning over long contexts, multi-step agent workflows
- **Use for**: Alternative to Mistral Small for structured scoring with reasoning
- Default params: temperature 0.15, top_p 0.75

#### Tier 4: Fallback

**`meta-llama/llama-3.3-70b-instruct:free`** — Llama 3.3 70B (dense)
- Context: 128,000 tokens | Max completion: 128,000 tokens
- Capabilities: tool calling
- Moderated: yes
- **Use for**: Universal fallback when primary models are rate-limited or unavailable
- Default params: none specified

**`openrouter/free`** — Free Models Router (auto-routing)
- Context: ~200,000 tokens | Capabilities: all (routes to model that supports request features)
- **Use for**: Last-resort fallback — OpenRouter auto-selects best available free model
- Note: Non-deterministic model selection — do not use for reproducible scoring

### Task-to-Model Routing

Each certification task maps to the optimal model based on required capabilities:

```
Task                     Primary Model               Fallback Model              Required Capability
─────────────────────────────────────────────────────────────────────────────────────────────────────
Pre-screen               mistral-small-3.1-24b       arcee-ai/trinity-mini       structured_outputs
Code Review              qwen3-coder-480b            llama-3.3-70b               tools (large context)
Dimension Scoring        qwen3-next-80b              mistral-small-3.1-24b       structured_outputs
Status Determination     gpt-oss-120b                gpt-oss-20b                 reasoning
Remediation Suggestions  qwen3-coder-480b            llama-3.3-70b               tools (large context)
Emergency Fallback       llama-3.3-70b               openrouter/free             tools
```

### Agent Review Pipeline

```
Certification Engine
        │
        ▼
┌──────────────────┐
│  1. Pre-screen   │ mistral-small-3.1-24b (structured_outputs)
│  "Worth deep     │ → returns { "needs_review": bool, "reason": string }
│   review?"       │ → if false: skip agent review, use deterministic evidence only
└────────┬─────────┘
         │ yes
         ▼
┌──────────────────┐
│  2. Code Review  │ qwen3-coder-480b (tool calling, 262k context)
│  Full code sent  │ → returns observations, concerns, suggestions as text
│  with policy ctx │ → attached to certification record as agent_commentary
└────────┬─────────┘
         ▼
┌──────────────────┐
│  3. Score        │ qwen3-next-80b (structured_outputs)
│  9 dimensions    │ → returns JSON: { dimensions: {correctness: 0.85, ...}, confidence: 0.7 }
│  as structured   │ → parsed and validated against schema
│  JSON            │ → stored with source: "agent" marker
└────────┬─────────┘
         ▼
┌──────────────────┐
│  4. Decide       │ gpt-oss-120b (reasoning)
│  Combine agent + │ → input: deterministic evidence + agent scores + policy thresholds
│  deterministic   │ → returns: recommended status + rationale
│  evidence        │ → deterministic failures ALWAYS override agent recommendation
└────────┬─────────┘
         │ if failing
         ▼
┌──────────────────┐
│  5. Remediate    │ qwen3-coder-480b (tool calling)
│  Generate fix    │ → returns: specific remediation steps + code examples
│  suggestions     │ → included in GitHub issue body and PR annotations
└──────────────────┘

Any step failure → fallback model → if fallback fails → skip step, continue pipeline
All steps are optional — engine produces valid certification with zero agent input
```

### Structured Output Schema (Dimension Scoring)

Used by Step 3 (qwen3-next-80b via `structured_outputs`):

```json
{
  "type": "json_schema",
  "json_schema": {
    "name": "dimension_scores",
    "strict": true,
    "schema": {
      "type": "object",
      "properties": {
        "dimensions": {
          "type": "object",
          "properties": {
            "correctness":                { "type": "number", "minimum": 0, "maximum": 1 },
            "maintainability":            { "type": "number", "minimum": 0, "maximum": 1 },
            "readability":                { "type": "number", "minimum": 0, "maximum": 1 },
            "testability":                { "type": "number", "minimum": 0, "maximum": 1 },
            "security":                   { "type": "number", "minimum": 0, "maximum": 1 },
            "architectural_fitness":      { "type": "number", "minimum": 0, "maximum": 1 },
            "operational_quality":        { "type": "number", "minimum": 0, "maximum": 1 },
            "performance_appropriateness": { "type": "number", "minimum": 0, "maximum": 1 },
            "change_risk":                { "type": "number", "minimum": 0, "maximum": 1 }
          },
          "required": ["correctness", "maintainability", "readability", "testability", "security", "architectural_fitness", "operational_quality", "performance_appropriateness", "change_risk"],
          "additionalProperties": false
        },
        "observations": {
          "type": "array",
          "items": { "type": "string" }
        },
        "confidence": { "type": "number", "minimum": 0, "maximum": 1 },
        "summary": { "type": "string" }
      },
      "required": ["dimensions", "observations", "confidence", "summary"],
      "additionalProperties": false
    }
  }
}
```

For models without `structured_outputs` support (e.g., qwen3-coder), the same schema is enforced via tool calling:

```json
{
  "type": "function",
  "function": {
    "name": "submit_dimension_scores",
    "description": "Submit the quality dimension scores for a code unit",
    "parameters": { ... same schema as above ... }
  }
}
```

### Provider Abstraction (Go Interface)

```go
// internal/agent/provider.go

// Provider abstracts LLM communication behind an OpenAI-compatible interface.
type Provider interface {
    ChatCompletion(ctx context.Context, req ChatRequest) (ChatResponse, error)
    Name() string
    Models() map[TaskType]ModelConfig
}

// ModelConfig holds per-model capability metadata from the OpenRouter catalog.
type ModelConfig struct {
    ID                     string  // e.g. "qwen/qwen3-coder:free"
    ContextLength          int     // e.g. 262000
    MaxCompletionTokens    int     // e.g. 262000, 0 if unspecified
    SupportsTools          bool
    SupportsStructuredOutput bool
    SupportsReasoning      bool
    IsModerated            bool
    DefaultTemperature     float64 // 0 means not specified
}

// TaskType identifies the certification pipeline step.
type TaskType string

const (
    TaskPrescreen    TaskType = "prescreen"
    TaskReview       TaskType = "review"
    TaskScoring      TaskType = "scoring"
    TaskDecision     TaskType = "decision"
    TaskRemediation  TaskType = "remediation"
)
```

### Configuration

```yaml
# .certification/config.yml (agent section)
agent:
  enabled: true
  provider:
    type: openrouter                    # openrouter | openai | local
    base_url: https://openrouter.ai/api/v1
    api_key_env: OPENROUTER_API_KEY     # env var name — actual key in GitHub repo secret
                                        # CI: ${{ secrets.OPENROUTER_API_KEY }}
                                        # Local: export OPENROUTER_API_KEY=sk-or-v1-...
    http_referer: https://github.com/code-certification/code-certification
    x_title: Code Certification System

  models:
    prescreen:   mistralai/mistral-small-3.1-24b-instruct:free
    review:      qwen/qwen3-coder:free
    scoring:     qwen/qwen3-next-80b-a3b-instruct:free
    decision:    openai/gpt-oss-120b:free
    remediation: qwen/qwen3-coder:free
    fallback:    meta-llama/llama-3.3-70b-instruct:free

  rate_limit:
    requests_per_minute: 20             # per-model free tier limit
    retry_max: 3
    retry_backoff_base_ms: 1000         # exponential: 1s, 2s, 4s

  review_scope:
    skip_if_deterministic_score_above: 0.9
    always_review_patterns:
      - "**/security/**"
      - "**/auth/**"

  prompts_dir: .certification/prompts   # versioned prompt templates
```

### GitHub Repository Secret Setup

The `OPENROUTER_API_KEY` is stored as a **GitHub repository secret** — never in code:

1. **Get free API key**: Sign up at [openrouter.ai](https://openrouter.ai) (no credit card required for free tier)
2. **Add to repo**: GitHub → **Settings → Secrets and variables → Actions → New repository secret**
   - Name: `OPENROUTER_API_KEY`
   - Value: `sk-or-v1-...`
3. **GitHub Actions wiring**:
```yaml
# .github/workflows/certification-pr.yml (excerpt)
jobs:
  certify:
    runs-on: ubuntu-latest
    env:
      OPENROUTER_API_KEY: ${{ secrets.OPENROUTER_API_KEY }}
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-go@v5
        with:
          go-version-file: go.mod
      - run: go build -o certify ./cmd/certify
      - run: ./certify review --changed-only
```
4. **Graceful degradation**: If `OPENROUTER_API_KEY` is empty or unset at runtime, the engine logs `info: agent review disabled (OPENROUTER_API_KEY not set)` and proceeds with deterministic-only certification. This is not an error.
5. **Local development**:
```bash
export OPENROUTER_API_KEY=sk-or-v1-...
# or add to ~/.zshrc / .envrc (gitignored)
```

### Agent Design Principles

1. **Agent output is advisory, not authoritative** — Deterministic evidence (lint, tests, type checks) always takes precedence. Agent scores are stored with `source: "agent"` and weighted separately in the certification record.
2. **The system works without agents** — If `agent.enabled: false` or API key is absent, certification runs with deterministic evidence only. Zero degradation of core functionality.
3. **Agent failures are graceful** — Rate limits, model outages, and invalid responses trigger fallback chain, then skip. Agent failures are warnings, never errors.
4. **Evidence provenance is explicit** — Every agent output includes: model ID, provider, timestamp, confidence score, prompt version, and `source: "agent"` marker.
5. **Free tier rate limits are respected** — Token-bucket rate limiter with proactive `x-ratelimit-remaining` header tracking. Bulk runs batch requests across models to distribute load.
6. **Prompt engineering is versioned** — System prompts stored in `.certification/prompts/*.v{N}.md`. Certification records reference prompt version for reproducibility.
7. **No secrets in code** — API key lives exclusively in GitHub repository secrets. Config references env var name only. Engine reads `os.Getenv()` at runtime.

### Planned File Structure (Agent System)

```
internal/
  agent/
    provider.go            # Provider interface
    openrouter.go          # OpenRouterProvider implementation
    openrouter_test.go     # HTTP mock tests for OpenRouter
    router.go              # TaskType → model routing + fallback chain
    router_test.go
    reviewer.go            # AgentReviewer pipeline orchestrator
    reviewer_test.go
    prompts.go             # Prompt template loading + versioning
    prompts_test.go
    schemas.go             # Structured output JSON schemas
    schemas_test.go
    types.go               # ChatRequest, ChatResponse, Message, Tool, ModelConfig, etc.
    ratelimit.go           # Token-bucket rate limiter with header tracking
    ratelimit_test.go

.certification/
  config.yml               # Full certification config including agent section
  prompts/
    prescreen.v1.md        # "Should this unit get deep review?"
    review.v1.md           # "Analyze this code for quality dimensions"
    scoring.v1.md          # "Score across 9 dimensions" + rubric
    decision.v1.md         # "Determine certification status from evidence"
    remediation.v1.md      # "Suggest specific fixes for failing dimensions"
```

## Relevant Files

### Existing Files to Modify
- **`justfile`** — 600+ lines of Khaos Machine commands. Complete rewrite to ~150 lines for single Go CLI.
- **`.claude/commands/feature.md`** — References `khaos-foundation/specs/` and "Khaos Machine platform". Change to `specs/` and "Code Certification System".
- **`.claude/commands/bug.md`** — Same `khaos-foundation/specs/` path and submodule references.
- **`.claude/commands/chore.md`** — Same `khaos-foundation/specs/` path and submodule references.
- **`.claude/commands/plan.md`** — References khaos-foundation build scripts (build.sh, sign.sh, release.sh, common.sh, KhaosFoundation.pkgproj, nfpm.yaml, packaging/, scripts/, docs/releases/) in plan format template. Replace with Go CLI project references.
- **`.claude/commands/prime.md`** — Entire content describes Khaos Machine (6 submodules, storytelling platform, screenplay analysis). Complete rewrite for Code Certification System.
- **`.claude/commands/install.md`** — References macOS packaging tools (pkgbuild, productsign, nfpm). Rewrite for Go development environment.
- **`.claude/commands/maintenance.md`** — References submodule health, upstream releases, khaos-foundation. Rewrite for single Go project.
- **`.claude/settings.local.json`** — Permissions list contains paths to `/Users/k/Spikes/khaos.machine/`, khaos-specific binaries, and ~60 overly-specific Bash permissions. Replace with clean, scoped permissions.

### Existing Files to Remove
- **`.claude/commands/plan_w_team.md`** — Team orchestration command. Remove from project (reference material only, not a project command). Also remove its hook validators since they validate plan_w_team-specific headers.
- **`.claude/skills/data-visualization-ux/`** — Khaos GUI React/Storybook skill (entire directory tree).
- **`.claude/skills/create-release/`** — Khaos-foundation release process.
- **`.claude/skills/platform-state/`** — Khaos Machine platform assessment.
- **`.claude/skills/frontend-design/`** — Web app UI design skill. Not applicable to Go CLI.
- **`.claude/skills/docusaurus-docs-management/`** — Docusaurus site management. Not applicable.
- **`.pi/skills/platform-state/`** — Duplicate of claude skill.
- **`.pi/skills/senior-ts/`** — TypeScript/Vue skill. Not applicable to Go-only project.
- **`.pi/skills/senior-go/SKILL.md`** — References Khaos Machine Go modules, Go Development Constitution. Remove (will create project-specific Go skill later during TDD).
- **`.pi/agent-sessions/`** — Entire directory. Cached Khaos Machine conversations (builder.json, documenter.json, planner.json, reviewer.json, scout.json). Stale data from wrong project.

### Existing Files to Keep (no changes needed)
- **`PRD.md`**, **`FEATURES.md`**, **`STORIES.md`** — Product documentation. Core project context.
- **`extensions/`** — Pi extensions (project-agnostic TypeScript).
- **`.pi/agents/builder.md`**, **`scout.md`**, **`planner.md`**, **`reviewer.md`**, **`documenter.md`**, **`plan-reviewer.md`**, **`red-team.md`**, **`bowser.md`** — Generic agent personas. No Khaos references.
- **`.pi/agents/teams.yaml`**, **`agent-chain.yaml`** — Generic orchestration config.
- **`.pi/agents/pi-pi/`** — Pi extension expert agents. Project-agnostic.
- **`.pi/themes/`** — Pi UI themes. Project-agnostic.
- **`.pi/settings.json`** — Pi settings (theme + prompts path). Keep as-is.
- **`.pi/skills/bowser/`** — Playwright automation skill. Project-agnostic.
- **`.codex/`** — Codex symlinks to .claude/. Keep as-is.
- **`.claude/commands/build.md`** — Generic "implement the plan" command. No Khaos references.
- **`.claude/commands/tdd.md`** — Generic TDD loop command. No Khaos references.
- **`.claude/commands/tools.md`** — Lists built-in tools. No Khaos references.
- **`.claude/commands/debug-gh-workflows.md`** — Generic GH workflow debugging. No Khaos references.
- **`.claude/skills/backlog-grooming/`**, **`development-loop/`**, **`effort-pointing/`**, **`feature-planning/`**, **`incremental-commit-all/`**, **`plan-from-latest-code-review-report/`**, **`project-learnings/`**, **`project-wide-code-review-report/`**, **`repo-root-cleanup/`**, **`sprint-planning/`**, **`unit-test-coverage/`**, **`web-search/`**, **`writing-tests/`**, **`fixture-audit/`** — Generic skills with no Khaos references.
- **`.claude/hooks/validators/`** — Keep `validate_file_contains.py` and `validate_new_file.py` (generic validators). The `plan_w_team.md` Stop hooks that reference them will be removed with the command.

### New Files
- **`CLAUDE.md`** — Project context for Claude Code / agent orientation.
- **`README.md`** — Project overview, getting started, architecture summary.
- **`.gitignore`** — Go project ignores.
- **`specs/`** — Directory for plan output (already created by this spec).

## Implementation Plan

### Phase 1: Foundation
Create essential project scaffolding and purge stale data:
- `.gitignore`, `CLAUDE.md`, `README.md`
- Delete all Khaos-specific skills, sessions, and the `plan_w_team.md` command
- `specs/` directory confirmed present

### Phase 2: Core Implementation
Rewrite all operational files:
- Complete justfile rewrite (~150 lines, single Go CLI)
- All Claude commands updated (6 files modified, 1 removed)
- `prime.md` fully rewritten
- `install.md` and `maintenance.md` fully rewritten
- `settings.local.json` rebuilt from scratch
- `plan.md` template references updated

### Phase 3: Integration
- Verify zero Khaos references remain in any operational file
- Verify justfile parses and doctor recipe runs
- Create initial git commit
- Verify clean working tree

## Step by Step Tasks
IMPORTANT: Execute every step in order, top to bottom.

### Step 1: Create `.gitignore`
- Create standard Go project `.gitignore` at project root:
  ```
  # Build output
  build/
  dist/
  certify

  # Go
  *.exe
  *.exe~
  *.dll
  *.so
  *.dylib
  *.test
  *.out
  *.prof
  coverage.out
  coverage.html

  # IDE
  .idea/
  .vscode/
  *.swp
  *.swo
  *~

  # OS
  .DS_Store
  Thumbs.db

  # Environment
  .env
  .envrc
  ```
- Verify: `cat .gitignore`

### Step 2: Create `README.md`
- Create project README with:
  - Project name: Code Certification System
  - One-paragraph description from PRD executive summary
  - Status: Pre-implementation (documentation and tooling only)
  - Prerequisites: Go 1.22+, golangci-lint, just, git, gh CLI
  - Quick start: `just doctor` to verify environment
  - Link to PRD.md, FEATURES.md, STORIES.md for full product context
  - Note that development follows TDD: tests first, implementation second
  - Brief mention of agent-assisted review via OpenRouter (see this spec for architecture)
- Verify: `cat README.md`

### Step 3: Create `CLAUDE.md`
- Create project context file for agent orientation:
  - Project: Code Certification System — a repository governance platform (Go CLI)
  - Current state: Pre-implementation. Documentation only. No Go source code yet.
  - Product docs: PRD.md (full requirements), FEATURES.md (acceptance checklist), STORIES.md (user stories)
  - Spec output: `specs/` directory
  - Development approach: TDD — every feature is test-driven (test → implement → refactor)
  - Build commands: `just build`, `just test`, `just lint`, `just check`, `just doctor`
  - Key architecture decisions: Go CLI, OpenRouter for agent-assisted review, policy-as-code, language-agnostic
  - Agent system: OpenRouter free-tier models, see `specs/project-bootstrap-justfile-commands-initial-commit.md` for full architecture
  - No submodules. Single Go module. Single binary (`certify`).
- Verify: `cat CLAUDE.md`

### Step 4: Delete Khaos-specific skills and sessions
- Remove directories and files:
  ```bash
  rm -rf .claude/skills/data-visualization-ux/
  rm -rf .claude/skills/create-release/
  rm -rf .claude/skills/platform-state/
  rm -rf .claude/skills/frontend-design/
  rm -rf .claude/skills/docusaurus-docs-management/
  rm -rf .pi/skills/platform-state/
  rm -rf .pi/skills/senior-ts/
  rm -rf .pi/skills/senior-go/
  rm -rf .pi/agent-sessions/
  ```
- Verify: `ls .claude/skills/ .pi/skills/` — confirm only generic skills remain
- Verify: `ls .pi/agent-sessions/ 2>/dev/null` — should not exist

### Step 5: Remove `plan_w_team.md` command
- Remove the file:
  ```bash
  rm .claude/commands/plan_w_team.md
  ```
- Verify: `ls .claude/commands/plan_w_team.md 2>/dev/null` — should not exist
- Note: The hook validators (`validate_file_contains.py`, `validate_new_file.py`) are generic utilities used by other commands too — keep them.

### Step 6: Rewrite `justfile`
- Replace entire justfile content with lean Go CLI recipes:
  - **Variables**: `build_dir := justfile_directory() / "build/bin"`
  - **Build**: `build` (go build → `build/bin/certify`), `clean` (rm -rf build/)
  - **Test**: `test` (go test ./...), `test-verbose` (go test -v ./...), `cover` (go test -coverprofile + go tool cover)
  - **Lint**: `lint` (golangci-lint run), `fmt` (gofmt -l check), `vet` (go vet ./...)
  - **Run**: `run *ARGS` (go run ./cmd/certify with args passthrough)
  - **Quality**: `check` (fmt + vet + lint + test combined gate)
  - **Doctor**: `doctor` (verify go, golangci-lint, git, gh, just are available with versions)
  - **Pi extensions**: Keep ALL `pi`, `ext-*`, and `open` recipes exactly as they are (project-agnostic)
  - **Claude/Codex agent launchers**: Keep `cld-*` and `cdx-*` recipes exactly as they are
  - **Remove everything else**: All submodule, build-tools, build-wfl, build-tui, build-gui, build-manager, wfl-*, gui*, tui, storybook*, embed-gui, release*, set-version, bench*, test-tools, test-wfl, test-tui, test-manager, doctor-versions, upstream, md-lint*, install-hooks, docs*, env, status, init, update recipes
- The `build` recipe should gracefully handle missing source:
  ```
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
  ```
- The `test` recipe should gracefully handle no Go files:
  ```
  test:
    #!/usr/bin/env bash
    set -eu -o pipefail
    if ! find . -name '*_test.go' -not -path './vendor/*' | grep -q .; then
      echo "⚠ No test files yet. Start with: just tdd"
      exit 0
    fi
    go test ./...
  ```
- Verify: `just --list` — should show only relevant recipes, no Khaos commands
- Verify: `just doctor` — should run and report tool availability
- Verify: `just build` — should print "No source code yet" gracefully
- Verify: `just test` — should print "No test files yet" gracefully

### Step 7: Rewrite `.claude/commands/prime.md`
- Complete rewrite. New content:
  - Description: "Quick-start agent understanding of the Code Certification System"
  - Purpose: Orient agent to the Code Certification System — a Go CLI that evaluates code against versioned policies and assigns time-bound certification status
  - Project overview: single Go CLI project, no submodules, TDD development approach
  - Product documentation map: PRD.md → FEATURES.md → STORIES.md
  - Current state: pre-implementation, documentation and tooling only
  - Target architecture: Go CLI (`certify`), `.certification/` directory in target repos, GitHub workflow integration, OpenRouter for agent-assisted review
  - Key concepts: certification units, certification records, policy packs, evidence, expiry model, certification status model
  - Build/test reference: `just build`, `just test`, `just lint`, `just check`, `just doctor`
  - Development approach: TDD — every feature implemented test-first
  - Report section: current state, architecture summary, recommended next steps
- Verify: `grep -c 'Khaos\|khaos\|storytelling\|screenplay' .claude/commands/prime.md` — must return 0

### Step 8: Update `.claude/commands/feature.md`
- Replace `khaos-foundation/specs/*.md` → `specs/*.md` (2 occurrences: instruction text and plan format)
- Replace "Khaos Machine platform" → "Code Certification System"
- Replace "This is a multi-repo workspace — the feature may span multiple submodules (khaos-manager, khaos-gui, khaos-tools, khaos-wfl, khaos-tui, khaos-foundation)." → "This is a single Go CLI project. All code lives in one repository."
- Remove "Start your research by reading `CLAUDE.md` and `AGENTS.md`" → "Start your research by reading `CLAUDE.md` and `PRD.md`"
- Verify: `grep -ci 'khaos' .claude/commands/feature.md` — must return 0

### Step 9: Update `.claude/commands/bug.md`
- Replace `khaos-foundation/specs/*.md` → `specs/*.md` (2 occurrences)
- Replace "This is a multi-repo workspace — the bug may span multiple submodules (khaos-manager, khaos-gui, khaos-tools, khaos-wfl, khaos-tui, khaos-foundation)." → "This is a single Go CLI project. All code lives in one repository."
- Replace "Start your research by reading `CLAUDE.md` and `AGENTS.md`" → "Start your research by reading `CLAUDE.md` and `PRD.md`"
- Verify: `grep -ci 'khaos' .claude/commands/bug.md` — must return 0

### Step 10: Update `.claude/commands/chore.md`
- Replace `khaos-foundation/specs/*.md` → `specs/*.md` (2 occurrences)
- Replace "This is a multi-repo workspace — the chore may span multiple submodules (khaos-manager, khaos-gui, khaos-tools, khaos-wfl, khaos-tui, khaos-foundation)." → "This is a single Go CLI project. All code lives in one repository."
- Replace "Start your research by reading `CLAUDE.md` and `AGENTS.md`" → "Start your research by reading `CLAUDE.md` and `PRD.md`"
- Verify: `grep -ci 'khaos' .claude/commands/chore.md` — must return 0

### Step 11: Update `.claude/commands/plan.md`
- In the Plan Format template, replace the "Key files" section. Remove:
  ```
  - `build.sh` - Main macOS build script
  - `build-linux.sh` - Linux packaging script
  - `build-package.sh` - Full pipeline orchestrator
  - `sign.sh` - Code signing and notarization
  - `release.sh` - GitHub release publishing
  - `common.sh` - Shared build utilities
  - `KhaosFoundation.pkgproj` - macOS installer project config
  - `nfpm.yaml` - Linux package spec
  - `packaging/` - Platform-specific installer assets
  - `scripts/` - WFL daemon helper scripts
  - `docs/releases/` - Versioned release notes
  ```
  Replace with:
  ```
  - `cmd/certify/` - CLI entry point
  - `internal/` - Core packages (engine, policy, evidence, agent, discovery)
  - `.certification/` - Configuration, policies, prompts, records
  - `justfile` - Build, test, lint recipes
  - `PRD.md` - Product requirements
  - `FEATURES.md` - Feature acceptance checklist
  - `STORIES.md` - User stories
  ```
- In the Validation Commands section, replace:
  ```
  - `shellcheck *.sh` - Lint all shell scripts
  - `./build.sh` - Verify macOS build completes successfully
  - `./build-linux.sh --arch amd64` - Verify Linux build completes successfully
  ```
  Replace with:
  ```
  - `just check` - Run fmt + vet + lint + test
  - `just build` - Verify build completes successfully
  ```
- Remove "Consider edge cases, error handling, and cross-platform (macOS/Linux) concerns" → "Consider edge cases, error handling, and testability"
- Verify: `grep -ci 'khaos\|pkgproj\|nfpm\|sign\.sh\|macOS build\|Linux packaging' .claude/commands/plan.md` — must return 0

### Step 12: Rewrite `.claude/commands/install.md`
- Complete rewrite. New content:
  - Description: "Verify development environment has required tools"
  - Workflow:
    1. Read `CLAUDE.md` for project context
    2. Verify key tools: `go` (Go 1.22+), `golangci-lint`, `git`, `gh` (GitHub CLI, authenticated), `just`
    3. Check `OPENROUTER_API_KEY` environment variable (info if missing, not error)
    4. Run `just doctor` for comprehensive check
    5. Report to user
  - Report format: tool versions, auth status, what passed, what failed, next steps
- Verify: `grep -ci 'pkgbuild\|productsign\|nfpm\|codesign\|khaos' .claude/commands/install.md` — must return 0

### Step 13: Rewrite `.claude/commands/maintenance.md`
- Complete rewrite. New content:
  - Description: "Run project maintenance checks and report health status"
  - Workflow:
    1. Read `CLAUDE.md` for project context
    2. Check git health: `git status`, clean working tree
    3. Check Go module: `go mod tidy` (verify no changes needed)
    4. Run lint: `just lint`
    5. Run tests: `just test`
    6. Check build: `just build`
    7. Report to user
  - Report format: health checks (git, module, lint, test, build), what passed, what failed, next steps
- Verify: `grep -ci 'khaos\|submodule\|upstream\|khaos-foundation' .claude/commands/maintenance.md` — must return 0

### Step 14: Rebuild `.claude/settings.local.json`
- Replace entire content with clean, scoped permissions for this project only:
  ```json
  {
    "permissions": {
      "allow": [
        "Bash(git *)",
        "Bash(gh *)",
        "Bash(go *)",
        "Bash(golangci-lint *)",
        "Bash(just *)",
        "Bash(ls *)",
        "Bash(find *)",
        "Bash(grep *)",
        "Bash(cat *)",
        "Bash(echo *)",
        "Bash(mkdir *)",
        "Bash(rm *)",
        "Bash(cp *)",
        "Bash(mv *)",
        "Bash(chmod *)",
        "Bash(pwd)",
        "Bash(command -v *)",
        "Bash(curl *)",
        "Bash(python3 *)",
        "Bash(open *)",
        "Bash(shellcheck *)",
        "WebSearch"
      ]
    }
  }
  ```
- Verify: `grep -ci 'khaos\|Spikes\|khaos.machine' .claude/settings.local.json` — must return 0
- Verify: `python3 -c "import json; json.load(open('.claude/settings.local.json'))"` — must parse without error

### Step 15: Verify zero Khaos references in operational files
- Run comprehensive scan:
  ```bash
  grep -rli 'khaos\|Khaos\|KHAOS' \
    justfile \
    .claude/commands/ \
    .claude/settings.local.json \
    .claude/skills/ \
    .pi/skills/ \
    .pi/agents/ \
    .pi/agent-sessions/ \
    CLAUDE.md \
    README.md \
    2>/dev/null
  ```
- Must return **zero results**. If any file matches, fix it before proceeding.
- Note: `extensions/` and `.pi/themes/` are not scanned — they are project-agnostic assets.

### Step 16: Verify justfile recipes
- Run: `just --list` — all recipes should list without parse errors
- Run: `just doctor` — should report tool availability
- Run: `just build` — should print "No source code yet" message
- Run: `just test` — should print "No test files yet" message
- Confirm NO Khaos-related recipes appear in `just --list` output

### Step 17: Initial git commit
- Ensure git is initialized (should already be from earlier):
  ```bash
  git init --initial-branch=main 2>/dev/null || true
  ```
- Stage all files:
  ```bash
  git add -A
  ```
- Review staged files:
  ```bash
  git status
  ```
  - Confirm expected files are staged (new: .gitignore, CLAUDE.md, README.md; modified: justfile, commands, settings; deleted: removed skills/sessions)
- Commit:
  ```bash
  git commit -m "chore: bootstrap code-certification project

  - Rewrite justfile for single Go CLI project (certify)
  - Update all Claude commands for Code Certification System
  - Rewrite prime.md, install.md, maintenance.md
  - Remove plan_w_team.md orchestration command
  - Rebuild settings.local.json with scoped permissions
  - Remove Khaos Machine skills, sessions, and references
  - Add .gitignore, CLAUDE.md, README.md
  - Add specs/ with agent-assisted review architecture
  - Document OpenRouter integration with free model catalog
  - Project ready for TDD-driven development"
  ```
- Verify clean state:
  ```bash
  git status
  ```
  — Should show "nothing to commit, working tree clean"

### Step 18: Run Validation Commands
- Execute every command in the Validation Commands section below
- Every command must pass. Zero failures.

## Testing Strategy

### Unit Tests
No unit tests for this feature — it is infrastructure/tooling only. All future features will be TDD (test → implement → refactor using `just tdd`).

### Manual Verification
1. `just --list` — shows only Code Certification recipes + pi extensions + agent launchers
2. `just doctor` — reports Go, golangci-lint, git, gh, just availability
3. `just build` — gracefully reports "No source code yet"
4. `just test` — gracefully reports "No test files yet"
5. `cat CLAUDE.md` — describes Code Certification System, not Khaos Machine
6. `cat README.md` — describes Code Certification System with getting started
7. `cat .claude/commands/prime.md` — orients to Code Certification, not storytelling
8. `git log --oneline -1` — shows bootstrap commit

### Edge Cases
- **Missing Go toolchain**: `just doctor` reports missing tools without crashing
- **No Go source code**: `just build` and `just test` handle gracefully with informative messages
- **Missing OPENROUTER_API_KEY**: Documented as expected — agent review disabled, not an error
- **Pi extensions**: All `ext-*` recipes still work since they reference `extensions/*.ts` files that exist unchanged

## Acceptance Criteria
1. `just --list` succeeds — no parse errors, no Khaos recipes
2. `just doctor` runs and reports tool availability
3. `just build` prints graceful "no source" message (exit 0)
4. `just test` prints graceful "no test files" message (exit 0)
5. Zero files in `justfile`, `.claude/commands/`, `.claude/settings.local.json`, `.claude/skills/`, `.pi/skills/`, `.pi/agents/`, `CLAUDE.md`, or `README.md` contain "khaos" (case-insensitive)
6. `.gitignore` exists with Go project patterns
7. `CLAUDE.md` exists and describes the Code Certification System
8. `README.md` exists with project overview and getting started
9. `specs/` directory exists with this plan
10. `.claude/commands/plan_w_team.md` does not exist
11. `.pi/agent-sessions/` directory does not exist
12. `.claude/skills/data-visualization-ux/` does not exist
13. `.claude/skills/frontend-design/` does not exist
14. `.pi/skills/senior-ts/` does not exist
15. `git log --oneline -1` shows the bootstrap commit
16. `git status --porcelain` returns empty (clean working tree)
17. All Claude commands reference `specs/` (not `khaos-foundation/specs/`) for plan output
18. `prime.md` describes Code Certification System with TDD development approach
19. `settings.local.json` is valid JSON with no Khaos-specific permissions
20. Agent architecture documented in this spec with: OpenRouter API contract, 7 researched free models with capabilities, task-to-model routing table, structured output schema, provider interface, configuration format, GitHub secret setup

## Validation Commands
Execute every command to validate the feature works correctly with zero regressions.

```bash
# 1. Justfile parses without errors
just --list

# 2. Doctor recipe runs
just doctor

# 3. Build handles no source gracefully
just build

# 4. Test handles no tests gracefully
just test

# 5. No Khaos references in any operational file
KHAOS_HITS=$(grep -rli 'khaos\|Khaos\|KHAOS' \
  justfile .claude/commands/ .claude/settings.local.json \
  .claude/skills/ .pi/skills/ .pi/agents/ \
  CLAUDE.md README.md 2>/dev/null | wc -l | tr -d ' ')
test "$KHAOS_HITS" -eq 0 && echo "PASS: zero Khaos references" || echo "FAIL: $KHAOS_HITS files still reference Khaos"

# 6. No old spec path in commands
OLD_PATH_HITS=$(grep -rl 'khaos-foundation/specs' .claude/commands/ 2>/dev/null | wc -l | tr -d ' ')
test "$OLD_PATH_HITS" -eq 0 && echo "PASS: spec paths updated" || echo "FAIL: old spec paths found"

# 7. Required new files exist
test -f .gitignore && echo "PASS: .gitignore" || echo "FAIL: .gitignore missing"
test -f CLAUDE.md && echo "PASS: CLAUDE.md" || echo "FAIL: CLAUDE.md missing"
test -f README.md && echo "PASS: README.md" || echo "FAIL: README.md missing"
test -d specs && echo "PASS: specs/" || echo "FAIL: specs/ missing"

# 8. Removed files/dirs don't exist
test ! -f .claude/commands/plan_w_team.md && echo "PASS: plan_w_team.md removed" || echo "FAIL: plan_w_team.md exists"
test ! -d .pi/agent-sessions && echo "PASS: agent-sessions removed" || echo "FAIL: agent-sessions exists"
test ! -d .claude/skills/data-visualization-ux && echo "PASS: data-viz removed" || echo "FAIL: data-viz exists"
test ! -d .claude/skills/frontend-design && echo "PASS: frontend-design removed" || echo "FAIL: frontend-design exists"
test ! -d .claude/skills/create-release && echo "PASS: create-release removed" || echo "FAIL: create-release exists"
test ! -d .pi/skills/senior-ts && echo "PASS: senior-ts removed" || echo "FAIL: senior-ts exists"
test ! -d .pi/skills/senior-go && echo "PASS: senior-go removed" || echo "FAIL: senior-go exists"
test ! -d .pi/skills/platform-state && echo "PASS: pi/platform-state removed" || echo "FAIL: pi/platform-state exists"
test ! -d .claude/skills/platform-state && echo "PASS: claude/platform-state removed" || echo "FAIL: claude/platform-state exists"
test ! -d .claude/skills/docusaurus-docs-management && echo "PASS: docusaurus removed" || echo "FAIL: docusaurus exists"

# 9. Settings is valid JSON
python3 -c "import json; json.load(open('.claude/settings.local.json'))" && echo "PASS: settings.local.json valid JSON" || echo "FAIL: invalid JSON"

# 10. Git state is clean with initial commit
git log --oneline -1
DIRTY=$(git status --porcelain | wc -l | tr -d ' ')
test "$DIRTY" -eq 0 && echo "PASS: clean working tree" || echo "FAIL: $DIRTY uncommitted changes"
```

## Notes
- **No `go.mod` in this commit** — The Go module will be created during the first TDD cycle when the first test file is written. `go mod init` requires a decision on module path that should align with the actual GitHub repository URL, which may not be finalized yet.
- **No Go source code** — This is deliberate. The project follows strict TDD. The first code will be a failing test, not a `main.go`.
- **Pi extensions kept as-is** — All `extensions/*.ts` files and their justfile recipes are project-agnostic pi framework tools. They work identically regardless of project.
- **Model availability may change** — OpenRouter's free model catalog is dynamic. Models may be added, removed, or have rate limits changed. The configuration is designed for easy model swapping via `.certification/config.yml`. The `openrouter/free` auto-router serves as the ultimate fallback.
- **`plan_w_team.md` retained as reference** — The team orchestration patterns documented in that file are useful reference material for understanding how to employ agents for certification tasks. The command itself is removed because the orchestration tooling (TaskCreate, TaskUpdate, etc.) is not applicable to this project. The concepts inform the agent pipeline design in this spec.
- **Skills cleanup is conservative** — Generic skills (backlog-grooming, writing-tests, development-loop, etc.) are kept even though some reference patterns from other projects in their examples. They provide useful methodology guidance that applies to any software project.

---

## Report

**Date of Completion**: March 9, 2026

**What was implemented**:
All 18 steps of the plan were executed in order. The code-certification project is now bootstrapped with project-specific tooling, zero references to any prior project, comprehensive agent-assisted review architecture documentation, and a clean initial git commit on `main`.

**Files created**: `.gitignore`, `CLAUDE.md`, `README.md`, `specs/project-bootstrap-justfile-commands-initial-commit.md`

**Files rewritten**: `justfile` (600+ lines → 250 lines), `.claude/commands/prime.md`, `.claude/commands/feature.md`, `.claude/commands/bug.md`, `.claude/commands/chore.md`, `.claude/commands/plan.md`, `.claude/commands/install.md`, `.claude/commands/maintenance.md`, `.claude/settings.local.json`

**Files removed**: `.claude/commands/plan_w_team.md`, `.claude/skills/data-visualization-ux/` (entire tree), `.claude/skills/create-release/`, `.claude/skills/platform-state/`, `.claude/skills/frontend-design/`, `.claude/skills/docusaurus-docs-management/`, `.pi/skills/platform-state/`, `.pi/skills/senior-ts/`, `.pi/skills/senior-go/`, `.pi/agent-sessions/` (5 cached sessions)

**Issues encountered**: None. All steps executed cleanly. Git was not yet initialized at project start — `git init --initial-branch=main` was run as part of Step 17.

**Refactoring**: No refactoring needed — this is a greenfield infrastructure commit with no prior code to refactor.

**Validation results**: All 10 validation command groups passed (20 individual checks). Zero failures. Clean working tree confirmed.
