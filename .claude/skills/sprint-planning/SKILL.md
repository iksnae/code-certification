---
name: sprint-planning
description: Plan an ad-hoc, point-based software sprint by selecting backlog items that align with project goals and maximize the likelihood of shipping working code at sprint end
---

# Sprint Planning Skill (Point-Based, Ad-Hoc)

When the user requests sprint planning, operate as a **throughput- and outcome-focused planner**.  
Assume **no dates, no hours, no fixed sprint length**.  
Sprints exist to **ship working code**, not to satisfy a timeline.

## Core Assumptions
- Capacity is measured **only in points**
- Sprint boundaries are flexible and defined by completion
- Velocity is inferred *after the fact* by points shipped per sprint
- Planning optimizes for **finish rate**, not utilization

## Inputs to Request (only if missing)
- Target project or initiative goal
- Prior sprint point throughput (if known)
- Backlog items with point estimates
- Known dependencies or sequencing constraints

## Planning Rules

### 1. Capacity Heuristic (Points Only)
- Use historical shipped points as the **upper bound**
- If no history exists:
  - Start conservatively
  - Bias toward under-commitment
- Never translate points into time

### 2. Sprint Goal Definition
- Define **one concrete shipping outcome**
- The sprint is considered successful only if this outcome is met
- If the goal slips, the sprint is not “done”

### 3. Backlog Selection
- Pull items that **directly advance the sprint goal**
- Prefer:
  - Smaller, finishable slices
  - Vertical slices over horizontal work
- Avoid large, ambiguous items unless explicitly approved

### 4. Commitment Structure
- **Committed Work**: required to ship the goal
- **Optional Work**: valuable but non-goal-critical
- The plan must remain valid if all optional work is dropped

### 5. Scope Discipline
- Actively identify:
  - Over-scoped items
  - Hidden dependencies
  - Work likely to block shipping
- Recommend re-slicing when risk is high

## Output Format (strict)
- **Sprint Goal** (1–2 lines, ship-focused)
- **Assumed Capacity** (points, with rationale)
- **Committed Work** (ordered list with point totals)
- **Optional Work** (clearly marked)
- **Risks & Invalidators** (what would force a re-plan)

## Style Constraints
- No dates, timelines, or calendar language
- No Scrum or agile theory explanations
- Be concise, opinionated, and execution-focused
- Every line must support a planning decision
