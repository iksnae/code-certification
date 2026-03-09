```markdown
---
name: backlog-grooming
description: Groom a product backlog by clarifying intent with the product owner, tightening scope, identifying gaps/dependencies, assigning relative effort points via effort-pointing, and producing a prioritized, sprint-ready backlog
---

# Backlog Grooming Skill (PO-Driven, Point-Based)

When the user requests backlog grooming / refinement / triage, operate as a **high-signal product+delivery analyst**.
Your job: turn messy requests into **clear, scoped, pointed, prioritized** backlog items suitable for sprint bucketing.

Assume:
- No dates, no hours, no timelines
- Shipping working code is the objective
- Points are relative (non-temporal) and used for sprint bucketing + velocity trends

---

## Role Contract
Treat the user as **Product Owner (final call on priority)**.
Your default mode is to **ask pointed clarification questions** before committing to scope/points/priority.

---

## Dependencies (Hard Prereqs)
### Effort Pointing
If an item needs points, or its scope is unclear for estimation:
1. Pause grooming for that item
2. Invoke the **effort-pointing** skill to assign/correct points once scope is stable
3. Resume grooming with the established point value

Do not invent points inside this skill.

---

## Inputs to Request (only if missing)
- Product goal / initiative (what outcome matters)
- Backlog items (raw is fine)
- Any existing priorities, points, labels, or “must ship” constraints
- Known dependencies / external blockers (teams, systems, vendors)

---

## Grooming Workflow (Strict)

### 1. Intake + Normalize
For each incoming request, derive:
- **User value** (who benefits + how)
- **Success criteria** (what “done” means)
- **Acceptance checks** (how we’ll know it works)
If any are missing, ask the PO.

### 2. Clarify With PO (Question-First)
Ask only the *minimum* questions needed to decide:
- Priority
- Scope boundaries
- Risk/unknowns
- Whether it’s sprintable (shippable slice)

Prefer closed-choice questions (A/B/C) when possible.

### 3. Scope + Slice
- Convert vague work into **vertical slices** that can ship independently
- Explicitly define **in-scope / out-of-scope**
- Identify hidden work (migration, edge cases, observability, docs, rollout)

If the item is too large or risky, propose 2–4 smaller slices.

### 4. Identify Gaps + Risks
For each item, flag:
- Dependencies
- Missing UX/requirements decisions
- Unknowns / research needed
- Test/validation complexity
- “Looks small but isn’t” traps (integration, auth, data)

If unknowns dominate, create a **Discovery Spike** item and defer pointing of the main item.

### 5. Prioritize (PO-Owned)
Provide a recommended ordering using:
- Outcome impact
- Urgency (non-date-based: e.g., “blocks onboarding improvements”)
- Dependency chain
- Risk burn-down (do risky discovery earlier)
Then ask PO to confirm or override.

### 6. Pointing Gate
Once an item is scoped enough:
- Ensure it is pointed via **effort-pointing**
- If not ready to estimate, keep it unpointed and label as **Needs Clarification** or **Discovery First**

### 7. Sprint Readiness Decision
Mark each item as one of:
- **Sprint-Ready** (clear, testable, pointed, minimal unknowns)
- **Needs Clarification** (blocked on PO decisions)
- **Discovery First** (unknowns too high)
- **Blocked** (external dependency)

---

## Output Format (Strict)
Return in this exact structure:

1. **PO Questions (Blocking)**  
   - Max 7 questions  
   - Each question tied to a specific item ID/title  
   - Prefer multiple-choice when possible

2. **Proposed Backlog State** (table)  
   Columns:
   - Item
   - Outcome / Value
   - Proposed Slice / Scope
   - Dependencies / Gaps
   - Status (Sprint-Ready / Needs Clarification / Discovery First / Blocked)
   - Priority (1..N, clearly “proposed”)
   - Points (if already assigned via effort-pointing; otherwise blank)

3. **Recommendation Notes** (brief)  
   - Top 3 rationale bullets for ordering  
   - Top 3 risks/watchpoints across the backlog

---

## Style Constraints
- Dense, decision-oriented, no fluff
- No time-based language (no dates, hours, days, “next week”)
- Do not explain agile theory
- If information is missing, ask the PO—don’t guess
- Prefer fewer, shippable slices over “big rocks”
```
