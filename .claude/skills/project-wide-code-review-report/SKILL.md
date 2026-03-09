---
name: project-wide-code-review-report
description: Perform a repository-wide code review to align the system with project intentions and industry best practices. Detect code smells, correctness risks, architectural issues (including Clean Architecture drift), SOLID alignment gaps, and opportunities to apply design patterns. Produce a written report in the repo under `reports/` with prioritized issues and actionable recommendations.
---



## Name
project-wide-code-review-report

## Description
Perform a repository-wide code review to align the system with project intentions and industry best practices. Detect code smells, correctness risks, architectural issues (including Clean Architecture drift), SOLID alignment gaps, and opportunities to apply design patterns. Produce a written report in the repo under `reports/` with prioritized issues and actionable recommendations.

---

## Non-Negotiables (Constraints)

- Assume NOTHING about the project’s language, frameworks, architecture, domain, or build system.
- Use existing repo materials to infer intent: READMEs, docs, ADRs, issues, comments, commit history, folder naming, and architecture notes.
- If project intent is unclear, perform lightweight research (web) only to clarify:
  - external dependencies / frameworks referenced
  - architectural patterns explicitly named
  - industry best practices relevant to the repo’s domain (e.g., serverless, mobile, CLI, library, monolith)
- Do not introduce new tooling requirements or mandate rewrites.
- Do not modify production code as part of this skill unless explicitly asked.
- Output must be a report file placed in `reports/` (create directory if missing).

---

## Inputs the Skill Should Accept

- Optional: requested review focus (e.g., “security”, “performance”, “maintainability”, “architecture”)
- Optional: target scope (entire repo default; allow path filters)
- Optional: preferred severity model (default included below)

If no inputs provided, review the entire repository and use default severity/prioritization.

---

## Output Artifact

Create one new report file:

- Path: `reports/code-review-<YYYY-MM-DD>.md`
- Content: structured, actionable, and repo-intent-aware
- Include:
  - executive summary
  - top risks & priorities
  - findings grouped by category
  - concrete recommendations (small/medium/large)
  - suggested follow-up tasks (issue-ready bullets)

Do not include any private keys, secrets, or sensitive values in the report.

---

## Review Dimensions (What to Evaluate)

### 1) Correctness & Reliability
- error handling consistency
- edge case handling
- data validation at boundaries
- concurrency hazards (if applicable)
- idempotency and retry safety (if relevant)
- logging/telemetry gaps

### 2) Maintainability & Readability
- naming consistency
- cyclomatic complexity hotspots
- long functions / god objects
- dead code, duplication, tight coupling
- inconsistent formatting or style drift (only note; don’t enforce a specific formatter)

### 3) Architecture & Boundaries
- separation of concerns
- module/package layering
- dependency direction (inward/outward)
- Clean Architecture alignment (if the repo claims it; otherwise describe boundary issues neutrally)
- domain vs infrastructure mixing
- boundary adapters (IO, persistence, network) leaking into core logic

### 4) SOLID Alignment
- SRP: overly broad modules/services
- OCP: switch/if ladders; fragile extension points
- LSP: subtype misuse; contract violations
- ISP: bloated interfaces; unnecessary dependencies
- DIP: high-level modules depending on low-level details

### 5) Design Patterns Opportunities (Pragmatic)
- where patterns could reduce coupling or clarify intent:
  - Strategy, Factory, Adapter, Facade, Decorator, Observer, Command, Repository, Unit of Work, State
- note tradeoffs and keep recommendations incremental

### 6) Testing & Quality Gates
- test organization and coverage gaps (qualitative)
- flaky tests indicators
- missing test seams / hard-to-test design
- missing static checks (lint/typecheck) only if repo indicates they’re expected

### 7) Security & Dependency Hygiene (Basic)
- obvious secret leakage (keys, tokens)
- unsafe deserialization / injection vectors (if recognizable)
- dependency risk signals (outdated critical deps) ONLY if the repo lists them

---

## Severity & Priority Model

Use this scale in the report:

- **P0 (Critical):** correctness/security risks, data loss, severe architectural blockers
- **P1 (High):** likely bugs, major maintainability problems, significant coupling
- **P2 (Medium):** code smells, moderate refactors, design improvements
- **P3 (Low):** polish, style consistency, minor cleanups

Each finding must include:
- priority (P0–P3)
- location(s) (file paths)
- symptom (what’s wrong)
- impact (why it matters)
- recommendation (what to do)
- effort estimate (S/M/L)
- optional: pattern suggestion (if applicable)

---

## Execution Plan (Step-by-Step)

### Step A — Establish Project Intent (No Assumptions)
1. Identify top-level intent signals:
   - README(s), docs, ADRs, CONTRIBUTING, ARCHITECTURE, RFCs
2. Summarize intent in 5–10 bullets for the report (“What this project seems to be trying to do”).
3. If intent is unclear or contradictory:
   - search within the repo for keywords (e.g., “architecture”, “clean architecture”, “DDD”, “hexagonal”, “service”, “cli”)
   - perform minimal web research ONLY on explicitly named technologies used by the repo
   - do not guess business domain beyond what the repo says

### Step B — Inventory & Hotspot Mapping
1. Map directories/modules and their responsibilities.
2. Identify hotspots:
   - large files
   - high-change areas (git history)
   - core domain vs IO boundaries
3. Note any explicit architecture claims and assess alignment.

### Step C — Deep Review Pass
1. Review core flows (entrypoints → domain → persistence/IO).
2. Identify coupling and boundary violations.
3. Detect code smells and anti-patterns.
4. Evaluate SOLID and propose incremental improvements.
5. Identify design patterns that would simplify or clarify (only when justified).

### Step D — Report Drafting
1. Write a prioritized, actionable report with:
   - Executive Summary
   - Intent Summary
   - Priority Findings (P0/P1 first)
   - Category Findings (architecture, SOLID, smells, tests, security)
   - Recommendations Roadmap (Quick wins / Medium / Larger)
   - Suggested Follow-up Tasks (issue-ready list)

### Step E — Save Report in Repo
1. Ensure `reports/` exists; create it if missing.
2. Write `reports/code-review-<YYYY-MM-DD>.md`.
3. (Optional) Add a short entry to an existing docs index ONLY if the repo already maintains one.
4. Do not commit unless explicitly instructed.

---

## Allowed Git + Shell Commands (Examples)

Use simple commands for discovery; do not run destructive operations.

- `git status`
- `git ls-files`
- `git log --oneline --decorate -n 50`
- `git log --name-only -n 50`
- `git tag --list --sort=-creatordate | head`
- `find . -maxdepth 3 -type f`
- `rg -n "architecture|clean architecture|DDD|hexagonal|layer|domain|usecase|adapter|repository" .`
- `rg -n "TODO|FIXME|HACK|XXX" .`
- `rg -n "panic|throw|fatal|System.exit|os.Exit" .` (language-agnostic hints; interpret carefully)

Prefer reading files directly over running build steps.

---

## Report Template (Use This Structure)

# Code Review Report — <YYYY-MM-DD>

## 1. Executive Summary
- Overall assessment (2–5 bullets)
- Top priorities (P0/P1 list)
- Biggest architectural theme(s)

## 2. Project Intent (Inferred)
- Evidence-backed summary from repo docs/code organization
- Known constraints and goals (if stated in repo)

## 3. Repository Overview
- Major modules/directories and perceived responsibilities
- Key entrypoints (apps, services, CLIs, libraries)

## 4. Priority Findings (P0 / P1)
### Finding: <Title>
- Priority: P0/P1
- Locations: <paths>
- Symptom:
- Impact:
- Recommendation:
- Effort: S/M/L
- Notes / References:

(repeat)

## 5. Architectural & Boundary Findings
- Clean Architecture / layering observations (only if relevant)
- Dependency direction issues
- Domain/IO leakage
- Suggested refactors (incremental)

## 6. SOLID Alignment Findings
- SRP:
- OCP:
- LSP:
- ISP:
- DIP:

## 7. Code Smells & Maintainability Findings
- Duplication, complexity hotspots, naming consistency, dead code

## 8. Testing & Quality Recommendations
- Gaps, missing seams, suggested test strategies (repo-appropriate)

## 9. Security & Dependency Hygiene (Basic)
- Obvious issues only; avoid speculation

## 10. Design Pattern Opportunities
- Candidate patterns + rationale + tradeoffs
- Keep recommendations minimal and high-leverage

## 11. Suggested Roadmap
### Quick Wins (S)
### Medium (M)
### Larger Refactors (L)

## 12. Issue-Ready Task List
- [ ] <task> (priority, effort, owner suggestion optional)

---

## Success Criteria

This skill succeeds if:
- A report is created in `reports/` with clear prioritization and repo-specific, evidence-backed recommendations.
- Findings are concrete and location-based (paths), not generic advice.
- Recommendations reflect project intentions and constraints inferred from the repository and minimal research.