---
name: plan-from-latest-code-review-report
description: Locate the most recent code review report in `reports/`, extract findings and recommendations, then use the planning tool to produce an actionable remediation plan that aligns with repository intent and your known best practices. The output is a structured plan (and optionally a new plan artifact in `reports/`) with prioritized tasks, sequencing, and acceptance criteria.

---

## Constraints (Non-Negotiables)

- Assume NOTHING about the project language, framework, tooling, or architecture.
- Use ONLY the contents of the latest report + repo docs to infer intent and constraints.
- Do not invent issues not present in the report.
- Do not change code unless explicitly requested; this skill produces a plan only.
- Comply with “known best practices” by:
  - preserving established architecture conventions found in the repo
  - favoring incremental changes
  - prioritizing correctness and safety
  - keeping changes testable and reviewable

---

## Inputs

- Optional: focus areas (e.g., “architecture first”, “testing first”, “security first”)
- Optional: max scope (e.g., “only P0/P1”)
- Optional: desired output format (default: Markdown plan)

If not provided:
- Plan for ALL findings, prioritizing P0 → P3.

---

## Step-by-Step Execution

### Step 1 — Identify Latest Report
1. Confirm `reports/` exists.
2. List candidate report files:
   - Prefer `reports/code-review-YYYY-MM-DD.md` pattern if present
   - Otherwise consider all `*.md` in `reports/`
3. Determine “latest” using:
   - date in filename when available
   - otherwise git history (most recent commit touching that file)
4. Open and read the latest report.

---

### Step 2 — Extract & Normalize Findings
Parse the report into a normalized set of items:
- Priority: P0/P1/P2/P3
- Title
- Locations (paths)
- Symptom
- Impact
- Recommendation
- Effort (S/M/L)

If the report uses different labels, map them to this schema without inventing detail.

---

### Step 3 — Validate Against Project Intent + Best Practices
1. Reconcile recommendations with:
   - stated architecture intentions (README/ARCHITECTURE/ADRs)
   - existing conventions (folder layout, naming, boundaries)
2. If a recommendation conflicts with repo intent:
   - note the conflict explicitly
   - propose an alternative aligned with intent
3. Keep changes incremental and low-risk where possible.

---

### Step 4 — Use Planning Tool To Produce Remediation Plan
Create a plan that includes:

#### 4.1 Workstreams (Suggested)
- Correctness & Reliability
- Architecture & Boundaries
- SOLID & Design Improvements
- Testing & Quality Gates
- Security & Dependency Hygiene

#### 4.2 Milestones
- M0: Triage & guardrails
- M1: P0 fixes
- M2: P1 remediation
- M3: P2 refactors + pattern improvements
- M4: P3 polish + consistency

#### 4.3 Task Format (Required)
For each task:
- ID
- Priority (P0–P3)
- Scope (paths/modules)
- Description (what)
- Rationale (why)
- Approach (how at a high level, no code)
- Acceptance Criteria (testable)
- Risk/Notes
- Dependencies (if any)
- Effort (S/M/L)

#### 4.4 Sequencing Rules
- Fix P0/P1 correctness or security before refactors.
- Add/adjust tests before large refactors when feasible.
- Refactor in small slices, maintaining working state.
- Favor seams and interfaces to reduce coupling.
- Avoid broad rewrites.

---

### Step 5 — Output the Plan
Deliver the plan in one or both of:
- Chat output (primary)
- Optional file: `reports/remediation-plan-<YYYY-MM-DD>.md` (only if repo conventions support adding plans)

Do not commit unless explicitly instructed.

---

## Allowed Commands (Discovery Only)

- `ls reports/`
- `find reports -maxdepth 1 -type f`
- `git log --name-only -n 50 -- reports/`
- `git show <ref>:reports/<file>`
- read/open files directly

No destructive commands.

---

## Success Criteria

This skill succeeds if it:
- Reliably picks the latest report in `reports/`
- Extracts all findings without fabrication
- Produces a prioritized, sequenced remediation plan
- Aligns recommendations with project intent and known best practices
- Provides clear acceptance criteria for each task
