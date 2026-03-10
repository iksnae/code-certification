# Chore: Unlock A Grades via Structural Score Boosts

## Chore Description

Currently, **no unit can achieve an A- or A grade** regardless of code quality. The maximum achievable score is 0.894 (B+) because 4 of 9 scoring dimensions lack positive boost mechanisms from structural evidence:

| Dimension | Current Max | Problem |
|-----------|------------|---------|
| `security` | 0.80 base (0.85 w/ agent) | Only penalized for global vars/ignored errors; never rewarded for clean code |
| `architectural_fitness` | 0.80 base (0.85 w/ agent) | Only penalized for context issues/god objects; never rewarded for clean API design |
| `performance_appropriateness` | 0.80 base (0.85 w/ agent) | Only penalized for defer-in-loop; never rewarded for lean functions |
| `operational_quality` | 0.85 max (>10 git commits) | Only boosted by git history; structural cleanliness not considered |

**Math:** Best possible = (0.95 + 0.95 + 0.95 + 0.90 + 0.80 + 0.80 + 0.85 + 0.80 + 0.90) / 9 = **0.878** without agent, **0.894** with agent. A- requires **0.90**.

**Fix:** Add positive boosts to `scoreFromStructural` for units demonstrating clean practices in these 4 dimensions. A perfect unit should score ~0.917 (A-). This is purely a scoring calibration — no new evidence collection needed.

**Boost criteria (all based on existing metrics already collected by structural analyzer):**

- **security** → 0.90 if `errors_ignored == 0 && global_mutable_count == 0 && panic_calls == 0`
- **architectural_fitness** → 0.90 if `param_count <= 3 && has_doc_comment == 1 && context_not_first == 0 && method_count <= 10`
- **performance_appropriateness** → 0.90 if `func_lines <= 30 && defer_in_loop == 0 && complexity <= 10`
- **operational_quality** → 0.90 if `errors_ignored == 0 && os_exit_calls == 0 && panic_calls == 0` (production-readiness indicators)

## Relevant Files

- `internal/engine/scorer.go` — Contains `scoreFromStructural()` which applies structural evidence to dimension scores; this is the **only file that needs logic changes**
- `internal/engine/scorer_test.go` — Tests for the scorer; needs new test cases for positive boosts validating A-grade scores are achievable
- `internal/domain/dimension.go` — Dimension constants and `GradeFromScore()` thresholds; read-only reference (no changes needed)
- `internal/evidence/structural.go` — Structural analyzer producing the metrics used by scorer; read-only reference (no changes needed, all metrics already collected)

## Step by Step Tasks

### Step 1: Add test for A- grade achievability (TDD)

- In `internal/engine/scorer_test.go`, add `TestScoreFromStructural_AMinusGrade`:
  - Create structural evidence with clean metrics: `func_lines: 20`, `errors_ignored: 0`, `global_mutable_count: 0`, `panic_calls: 0`, `os_exit_calls: 0`, `defer_in_loop: 0`, `param_count: 2`, `has_doc_comment: 1`, `exported_name: 1`, `context_not_first: 0`, `method_count: 3`, `max_nesting_depth: 2`, `complexity: 5`, `naked_returns: 0`
  - Also provide lint (pass), test (pass), and metrics evidence (low complexity, short code)
  - Call `Score(evidence, emptyEvalResult)`
  - Assert `GradeFromScore(weightedAvg) >= GradeAMinus` (score >= 0.90)
  - Assert `security >= 0.90`, `architectural_fitness >= 0.90`, `performance_appropriateness >= 0.90`
- Run test — expect **fail** (current max is 0.878 without agent)

### Step 2: Add positive boosts to `scoreFromStructural`

- In `internal/engine/scorer.go`, add positive boost block at the **end** of `scoreFromStructural` (after all penalties, so penalties still override):

  ```go
  // Positive boosts for clean structural practices.
  // These reward units that demonstrate quality in dimensions
  // that are otherwise stuck at the 0.80 base.

  // Security: reward units with no ignored errors, no global state, no panics
  if ignored == 0 && globalMut == 0 && panicCalls == 0 {
      scores[domain.DimSecurity] = max(scores[domain.DimSecurity], 0.90)
  }

  // Architectural fitness: reward clean API design
  if params <= 3 && hasDoc == 1.0 && e.Metrics["context_not_first"] != 1.0 && methodCount <= 10 {
      scores[domain.DimArchitecturalFitness] = max(scores[domain.DimArchitecturalFitness], 0.90)
  }

  // Performance appropriateness: reward lean, efficient functions
  complexity := metricInt(e, "complexity")
  if funcLines > 0 && funcLines <= 30 && deferInLoop == 0 && complexity <= 10 {
      scores[domain.DimPerformanceAppropriateness] = max(scores[domain.DimPerformanceAppropriateness], 0.90)
  }

  // Operational quality: reward production-ready code (no panics, no os.Exit, proper error handling)
  if ignored == 0 && osExitCalls == 0 && panicCalls == 0 {
      scores[domain.DimOperationalQuality] = max(scores[domain.DimOperationalQuality], 0.90)
  }
  ```

- **Note:** The variables `ignored`, `globalMut`, `panicCalls`, `funcLines`, `osExitCalls`, `deferInLoop`, `params`, `hasDoc`, `methodCount` are already computed earlier in `scoreFromStructural`. The complexity metric needs to be read from `e.Metrics["complexity"]` (it's available from the structural evidence, but currently only read in `scoreFromMetrics` — need to also read it here OR use a helper). Check if structural evidence includes complexity metric; if not, use a safe default.
- Run test from Step 1 — expect **pass**

### Step 3: Add test for penalty-overrides-boost behavior

- In `internal/engine/scorer_test.go`, add `TestScoreFromStructural_PenaltyOverridesBoost`:
  - Create structural evidence with `errors_ignored: 1` but otherwise clean
  - Assert `security` is penalized (≤ 0.60, from existing `min` clamp) despite other clean metrics
  - Assert `operational_quality` is NOT boosted (since `errors_ignored > 0`)
  - This validates that penalties still take precedence over boosts
- Run test — expect pass

### Step 4: Add test for partial boosts

- In `internal/engine/scorer_test.go`, add `TestScoreFromStructural_PartialBoosts`:
  - Create structural evidence with `func_lines: 50` (too long for perf boost), `param_count: 1`, `has_doc_comment: 1` (qualifies for arch boost)
  - Assert `architectural_fitness >= 0.90` (boosted)
  - Assert `performance_appropriateness == 0.80` (not boosted — func too long)
  - This validates each boost is independent
- Run test — expect pass

### Step 5: Verify existing tests still pass

- Run `go test ./internal/engine/... -count=1 -v`
- Run `go test ./... -count=1` — full suite
- Verify no regressions

### Step 6: Build and re-certify

- `go build -o build/bin/certify ./cmd/certify/`
- `./build/bin/certify certify --root . --skip-agent --batch --reset-queue`
- `./build/bin/certify report`
- Check report card for A-/A grade units

### Step 7: Validate A grades exist

- Parse records to confirm units with A- or A grades now exist
- Verify grade distribution includes A- entries
- Confirm no units were downgraded (score changes should only be upward)

## Validation Commands

```bash
# All tests pass
go test ./... -count=1

# Build succeeds
go build -o build/bin/certify ./cmd/certify/

# Re-certify
./build/bin/certify scan --root .
./build/bin/certify certify --root . --skip-agent --batch --reset-queue
./build/bin/certify report

# Verify A grades exist
python3 -c "
import json, glob
from collections import Counter

with open('.certification/index.json') as f:
    index = json.load(f)
index_ids = set(u['id'] for u in index)

grades = Counter()
a_units = []
for f in glob.glob('.certification/records/*.json'):
    with open(f) as fh:
        r = json.load(fh)
    if r['unit_id'] not in index_ids:
        continue
    grades[r.get('grade','')] += 1
    if r.get('grade','') in ('A', 'A-'):
        a_units.append((r['unit_id'], r.get('score',0)))

print('Grade distribution:', dict(grades))
print(f'A/A- units: {len(a_units)}')
for uid, score in sorted(a_units, key=lambda x: -x[1])[:20]:
    print(f'  {score:.3f} {uid}')
assert len(a_units) > 0, 'No A-grade units found!'
print('A GRADES UNLOCKED')
"
```

## Report

**Date:** 2026-03-10

### What Was Implemented

Added 4 positive structural score boosts in `scoreFromStructural()` that reward clean code practices in dimensions previously stuck at the 0.80 base:

1. **security** → 0.90 when `errors_ignored == 0 && global_mutable_count == 0 && panic_calls == 0`
2. **architectural_fitness** → 0.90 when `param_count <= 3 && has_doc_comment == 1 && context_not_first == 0 && method_count <= 10`
3. **performance_appropriateness** → 0.90 when `(func_lines == 0 || func_lines <= 30) && defer_in_loop == 0 && nesting <= 3`
4. **operational_quality** → 0.90 when `errors_ignored == 0 && os_exit_calls == 0 && panic_calls == 0`

### Results

| Metric | Before | After | Delta |
|--------|--------|-------|-------|
| Overall Score | 85.8% | **88.4%** | **+2.6%** |
| Overall Grade | B | **B+** | ⬆ |
| A- units | 0 | **340** | **+340** |
| B+ units | 326 | 186 | -140 (promoted to A-) |
| B units | 400 | 229 | -171 |
| C units | 22 | 11 | -11 |
| Observations | 29 | 18 | -11 |
| Go score | 85.7% | **88.6%** | **+2.9%** |

### Tests Added (4 new, 19 total)

1. `TestScoreFromStructural_AMinusGrade` — Clean unit achieves A- (≥0.90) with all 4 dimension boosts
2. `TestScoreFromStructural_PenaltyOverridesBoost` — `errors_ignored: 1` prevents security/operational boosts but not architectural
3. `TestScoreFromStructural_PartialBoosts` — Long function blocks performance boost but other boosts still apply
4. `TestScoreFromStructural_FileLevelUnit` — `func_lines: 0` (type/var units) treated as clean for performance boost

### Issues Encountered

1. **Complexity metric not in structural evidence**: Plan assumed `complexity` was available in structural evidence; it's only in metrics evidence. Replaced with `nesting <= 3` (available in structural) as the performance complexity proxy.

2. **File-level units have `func_lines: 0`**: Types and variable declarations don't have function bodies. The original plan condition `funcLines > 0 && funcLines <= 30` would exclude all file-level units. Fixed with `funcLines == 0 || funcLines <= 30`.

### Refactoring

- No refactoring needed — the change was a clean 16-line addition to the existing `scoreFromStructural` function, using variables already computed earlier in the same function.

### FEATURES.md

All criteria were already checked off. This change improves scoring calibration without adding features.

### Commits

```
d05496f feat: unlock A grades — add positive structural boosts for 4 stuck dimensions
```

## Notes

- **Boosts use `max()` not assignment:** This ensures they don't accidentally lower a score already above 0.90 from other evidence (e.g., agent review giving higher).
- **Penalties still win:** Since penalties use `min()` and come before boosts in the code flow, a unit with `panic_calls: 1` gets correctness clamped to 0.50 regardless of other clean metrics. However, the boost block runs after penalties, so we use `max()` which won't override a penalty-clamped value since `max(0.50, 0.90) = 0.90` would undo the penalty. **Fix:** Order matters — boosts should only fire when the prerequisite metrics are clean. Each boost condition already checks for clean metrics (e.g., security boost requires `panic_calls == 0`), so the penalty case won't trigger the boost.
- **Complexity metric in structural evidence:** The structural analyzer collects `complexity` via `go/ast` (cyclomatic). This is available in `e.Metrics["complexity"]` for structural evidence. However, `scoreFromMetrics` also reads complexity from the metrics evidence kind. The structural boost uses the structural evidence's own complexity metric.
- **File-level units (types, vars):** File-level units like `go://internal/domain/policy.go#PolicyPack` have `func_lines: 0`. The `funcLines > 0 && funcLines <= 30` condition handles this — units without function bodies won't get the performance boost from func_lines, but may still qualify if complexity is available. Consider treating `func_lines == 0` as clean for type definitions.
- **Expected impact:** ~300+ units currently at 0.878 (B+) with clean metrics should jump to 0.90+ (A-). This would shift the grade distribution dramatically: B+ 326 → ~50, A- ~300+, overall score 85.8% → ~89%.
