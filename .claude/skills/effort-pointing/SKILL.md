---
name: effort-pointing
description: Assign relative effort points to backlog items by reasoning about work across agent modes, focusing on complexity, uncertainty, and integration cost rather than time
---

# Effort Pointing Skill (Mode-Based)

When the user requests pointing or estimation, operate as a **calibrated estimator**.  
Points represent **relative effort and risk**, distributed across agent modes — never time.

## Core Assumptions
- Points are abstract, relative, and non-temporal
- Estimation reflects **cognitive load, uncertainty, and integration friction**
- Different agent modes contribute differently to total effort
- Points exist to support sprint bucketing and velocity projection

## Pointing Model

### What Points Capture
- Problem complexity
- Degree of unknowns
- Cross-cutting coordination
- Likelihood of rework
- Integration and validation effort

### What Points Explicitly Exclude
- Hours, days, or calendar time
- Individual developer speed
- Optimism or stretch goals

## Mode Contribution Heuristic

Estimate how effort is distributed across modes.  
Not all modes must be used, but all **material contributors must be considered**.

- **Architect mode**
  - System design
  - Boundary definition
  - Tradeoff analysis
  - Risk: incorrect abstractions

- **Code mode**
  - Implementation complexity
  - Refactoring depth
  - Algorithmic difficulty
  - Risk: hidden coupling

- **Debug mode**
  - Failure modes
  - Observability gaps
  - Non-obvious edge cases
  - Risk: prolonged uncertainty

- **Test / Validation**
  - Verification difficulty
  - Environment setup
  - Confidence to ship
  - Risk: false positives

## Pointing Rules

1. **Start with baseline complexity**
   - Simple, known patterns → low points
   - Novel or cross-cutting work → higher points

2. **Adjust for uncertainty**
   - Unknowns inflate points more than raw work
   - If success criteria are fuzzy, points go up

3. **Account for integration**
   - Touching multiple systems increases points
   - Backwards compatibility adds weight

4. **Bias toward finishability**
   - If work cannot realistically ship as scoped, recommend re-slicing
   - Smaller, shippable slices are preferred even if total points increase slightly

## Output Format (strict)
- **Item Summary**
- **Mode Effort Distribution** (bullets or percentages)
- **Point Assignment** (single number)
- **Estimation Rationale** (1–3 concise lines)
- **Risk Notes** (if any)

## Style Constraints
- No time-based language
- No justifications longer than necessary
- Be decisive; avoid hedging
- If pointing confidence is low, say so explicitly
