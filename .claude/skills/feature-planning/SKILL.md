---
name: feature-planning
description: Plan a new feature for an existing codebase using structured planning and an explicit task todo list, keeping risks, scope, and steps clear before implementation
---

# Feature Planning (Plan + TODO Driven)

When the user asks to plan/design/scope a new feature for an existing codebase, follow this workflow.
Do NOT jump straight to coding.

## 0) Mandatory Operating Mode: Planning + TODO
- Start by producing a clear written plan (sections below).
- Then immediately create a Task Todo List representing the plan’s work items.
- Keep the Task Todo List continuously updated:
  - Mark items complete as they’re done
  - Add new items when discoveries require more work
  - Reorder/prioritize when new constraints appear
- Use the todo list as the single source of truth for “what’s next”.

## 1) Context & Goals (Planning Section)
Collect or infer (and clearly label assumptions):
- Feature goal / user outcome
- Non-goals (what is explicitly out of scope)
- Target users / personas (if relevant)
- Constraints: performance, compatibility, security, rollout, timeline
- Success criteria: measurable acceptance criteria

If key context is missing, ask targeted questions BEFORE designing.

## 2) Codebase Recon (Planning Section)
Identify where the feature should live by mapping:
- Existing modules/components/services related to the feature
- Extension points vs high-risk areas
- Public APIs vs internal interfaces
- Data models involved
- Cross-cutting concerns: auth, logging, caching, background jobs, UI routing, etc.

Output:
- “Likely touched areas” list (paths/components)
- “Dependencies & integrations” list
- “Unknowns to verify” list

Add each unknown/verification step to the Task Todo List.

## 3) Risks & Trade-offs (Planning Section)
Explicitly call out:
- Breaking-change risks
- Backwards compatibility requirements
- Migration needs (data/schema/config)
- Performance/scalability risks
- Testing complexity
- Security/privacy implications

Add mitigations as todo items.

## 4) Proposed Design (Planning Section)
Provide:
- High-level approach (extend/compose/refactor/isolate)
- New components/modules to introduce (if any)
- Changes to existing components (minimise blast radius)
- Data flow + control flow (text diagram is fine)
- Error handling + observability expectations
- Rollout strategy (feature flag, gradual rollout, canary) when relevant

Convert design into actionable todo items:
- Each todo should be small, reviewable, and testable.

## 5) Incremental Implementation Plan (Planning Section)
Break into steps that can ship safely:
1. Prep / refactors (only if needed)
2. Feature scaffolding
3. Core logic
4. Integration points
5. Tests
6. Docs + metrics + rollout

Add each step to the Task Todo List (with clear acceptance criteria).

## 6) Testing Plan (Planning Section)
Define:
- Unit tests to add/update
- Integration/contract tests (if applicable)
- Regression risks
- Edge cases + failure modes

Create explicit todo items for each test batch.

## 7) Output Format Requirements
When responding, always include:
- A structured plan (sections 1–6)
- A short “Next up” list that matches the top incomplete todo items
- A list of assumptions (if any)
- A list of questions/blockers (if any)

## 8) Guardrails
- Do not write production code unless the user explicitly asks to move from planning into implementation.
- If the user requests implementation, first confirm the plan/todo list is up to date, then proceed step-by-step following the todo order.
