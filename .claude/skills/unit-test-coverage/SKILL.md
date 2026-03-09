---
name: unit-test-coverage
description: Increase unit test coverage by identifying untested code paths, adding isolated unit tests, and improving testability without assuming any specific language, framework, or tooling.
---

# Unit Test Coverage Improvement Skill

## Mission
Raise unit test coverage *meaningfully* (not just numbers) by:
- targeting high-risk / high-value code
- writing deterministic, isolated unit tests
- improving testability with minimal, safe refactors
- preventing regressions by adding coverage for reported/likely failures

This skill must **not assume** any specific project architecture, language, test framework, runner, coverage tool, CI system, or mocking library.

## Operating Principles (non-negotiable)
1. **Prefer isolated unit tests** over integration tests unless isolation is impossible.
2. **Test behavior and contracts**, not implementation details (but allow pragmatic seam-testing when needed).
3. **Deterministic tests**: no network, real time, random, global state, real filesystem, or external services unless explicitly controlled.
4. **Small, safe refactors only**: introduce seams (interfaces/wrappers/injection) to enable testing; do not redesign architecture.
5. **Coverage with intent**: prioritize branches, edge cases, and error handling—avoid “assert true” tests.

## Workflow

### Step 0 — Align on constraints (minimal questions)
If the user didn’t specify, ask only what’s necessary:
- What’s the target: overall %, per-module, or specific files/functions?
- Any exclusions (generated code, vendor, migrations, legacy)?
- Is there an existing failing bug or incident to cover?

If answers are unavailable, proceed by:
- picking the most changed/recent/high-risk areas
- prioritizing public APIs / domain logic over glue code

### Step 1 — Establish a baseline
Goal: identify *where* missing coverage is concentrated.

Do not assume tooling. Instead:
- Locate how the project runs tests (e.g., scripts, docs, CI config).
- Locate how coverage is produced (if present).
- If no coverage exists, propose adding it using the project’s existing ecosystem, but keep instructions generic:
  - “Use your language’s standard coverage tool (or the one already used in CI) and generate a per-file report.”

Deliverable:
- A short list of **top 5–10 files/modules** with the biggest coverage gaps and why they matter.

### Step 2 — Choose targets (maximize impact)
Rank candidates using this heuristic:
1. Code that is **critical** (auth, money, permissions, safety)
2. Code with **high churn** (recently changed)
3. Code with **complex branching** (many conditions/early returns)
4. Code with **bug history**
5. Public-facing APIs / boundary adapters (but test internals via unit seams)

Avoid:
- trivial getters/setters
- pure wiring/DI configuration unless it contains logic
- auto-generated code (unless it has real logic)

### Step 3 — Design tests before writing them
For each chosen unit:
- Identify observable behavior:
  - outputs / return values
  - thrown/returned errors
  - state transitions
  - interactions (calls) only when the interaction is the contract
- List key partitions:
  - happy path
  - boundary values
  - invalid inputs
  - empty/null cases
  - error/timeout paths (simulated)
  - branch edges (each condition true/false)
- Define fixtures:
  - keep them minimal
  - prefer explicit builders/factories over huge shared fixtures

Deliverable per unit:
- A mini test plan: “cases → expected result → required stubs/fakes”.

### Step 4 — Make code testable (minimal refactor patterns)
If the unit is hard to test, use the smallest applicable seam:

**Seam patterns (pick one):**
- **Dependency injection**: pass collaborators as parameters/constructor args.
- **Wrapper/adaptor** around time, random, UUID, environment, IO.
- **Pure function extraction**: move logic into pure function, test that, leave thin wrapper.
- **Strategy/Interface**: replace hard-coded static/global calls with an injected interface.

Rules:
- No behavior change without explicit agreement.
- Refactors must be covered by tests (either existing or newly added).
- Prefer local changes; avoid rippling changes through many modules.

### Step 5 — Implement tests (quality checklist)
Each test should be:
- **Arrange / Act / Assert** (or Given/When/Then)
- Single reason to fail
- Clear naming: `when_<condition>_then_<outcome>`
- No shared mutable state between tests
- Stable ordering (no dependence on other tests)

**Assertions**
- Assert on outcomes and invariants.
- For interactions, assert only:
  - that a collaborator was called with correct key data **when that call is the contract**, and
  - that errors are handled correctly.

**Mocking/Stubbing guidance**
- Prefer fakes/stubs over mocks when possible.
- Avoid mocking the unit under test.
- Mock at boundaries (network/db/fs/time), not in the middle of domain logic.

### Step 6 — Expand coverage intentionally
After initial tests:
- Re-run coverage and identify remaining uncovered lines/branches.
- Add tests only when they represent a meaningful scenario.
- Cover:
  - else branches
  - early returns
  - exception paths
  - defensive checks

### Step 7 — Guardrails and regression prevention
- Add tests for any bug being addressed (or recently observed failures).
- If a bug fix is included, follow:
  1) reproduce with failing test
  2) fix
  3) keep the test

### Step 8 — Report results in a useful format
Provide a concise summary:
- Coverage delta (overall and for targeted files) if available
- Tests added (count + what they cover)
- Refactors made (what seam + why safe)
- Remaining gaps (top 3) and next recommended actions

## Common Pitfalls (and what to do instead)
- **Chasing %**: Don’t add meaningless tests; target risk/branching.
- **Flaky tests**: Replace time/random/async nondeterminism with injected seams.
- **Over-mocking**: If you can’t refactor to seams, use a fake implementation.
- **Testing private internals**: Prefer public behavior; if unavoidable, extract pure logic.

## “Quick Start” Prompt Template (for the user to invoke this skill)
When you ask Roo for coverage help, include:
- the file(s) or module(s) to target
- desired coverage goal (or “maximize in these files”)
- how to run tests (command/script) if known
- how to run coverage (command/script) if known
