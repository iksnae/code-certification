---
name: project-learnings
description: Build or refresh a LEARNINGS.md project map so an agent can quickly get bearings before starting work in this repo
---

# Project Learnings / Context Map Skill

## Goal
When the user asks to “get project context”, “map the repo”, “help me understand this codebase”, “onboard”, “where is X implemented?”, or similar: ensure a **LEARNINGS.md** exists and is up to date, then use it as the primary orientation doc for the session.

This skill is documentation-first: create or refresh **LEARNINGS.md** and then rely on it to guide subsequent work.

## File to manage
- Primary: `LEARNINGS.md` at the repository root.

## Workflow

### 1) Locate or create `LEARNINGS.md`
1. Check whether `LEARNINGS.md` exists at repo root.
2. If it exists:
   - Load it and treat it as authoritative “project bearings”.
   - Only update it if you discover missing/incorrect info during mapping.
3. If it does **not** exist:
   - Create it with the template below (keep it concise).
   - Then proceed to populate it based on repo inspection.

**Template (use as initial content if file is missing):**
```md
# LEARNINGS.md

> Project bearings for agents and humans. Update when you discover new context.

## 1. What this project is
- Purpose:
- Primary users / customers:
- Key outcomes / success metrics:

## 2. How to run it (developer quickstart)
- Prereqs:
- Install:
- Configure:
- Run:
- Test:
- Lint/format:
- Common commands:

## 3. Architecture map (high level)
- Components/services:
- Data flow:
- External dependencies / integrations:
- Storage:
- Auth:

## 4. Repo map (where things are)
- Key directories:
- Entry points:
- Configuration files:
- Build/deploy scripts:
- CI/CD:

## 5. Key workflows
- Most common developer workflows:
- Release process:
- Migration process (if any):

## 6. Gotchas / sharp edges
- Known pitfalls:
- Non-obvious conventions:
- Performance hotspots:

## 7. Glossary
- Domain terms:

## 8. Open questions
- Unknowns to clarify:
````

### 2) Build a “repo map” efficiently

Prioritize speed and signal over completeness.

1. Identify the project type(s):

   * Look for: `README*`, `package.json`, `pnpm-lock.yaml`, `yarn.lock`, `requirements.txt`, `pyproject.toml`, `Cargo.toml`, `go.mod`, `pom.xml`, `.csproj`, `Gemfile`, `composer.json`, etc.
2. Find entry points and important routes:

   * Web: server start file, router definitions, main application module.
   * CLI: `main` files, command registries.
   * Libraries: public API surface, `index.*`, exported modules.
3. Locate configuration & environment:

   * `.env*`, config directories, `config.*`, `settings.*`, Helm/Terraform, docker compose, k8s manifests.
4. Map tests:

   * Where tests live, how they run, any test helpers/fixtures.
5. Map build/deploy:

   * CI workflows (e.g., GitHub Actions), scripts, containers, infra.

### 3) Populate / refresh `LEARNINGS.md`

As you learn, update `LEARNINGS.md` with:

* The simplest “how to run” commands that likely work.
* A top-level directory map with 1–2 lines per important folder.
* The most important entry points and “where to change X”.
* Any conventions (linting, formatting, branching, commit style).
* A short list of “open questions” rather than guessing.

Rules for edits:

* Keep it skimmable. Prefer bullet points over paragraphs.
* If uncertain, label as **Assumption** or add to **Open questions**.
* Don’t delete useful history—append with dated notes if needed.

### 4) Use `LEARNINGS.md` during the session

After creating/updating it:

* Summarize the repo map and key commands from `LEARNINGS.md`.
* When the user asks “where is X?”, consult `LEARNINGS.md` first, then dive into code.
* If you find a new key path or rule, immediately add it back into `LEARNINGS.md`.

## Quality bar

At the end of this skill run, `LEARNINGS.md` should enable a new agent to:

* Understand what the project is
* Run it locally (or know exactly what’s missing)
* Know where main components live
* Know where to start making changes safely
