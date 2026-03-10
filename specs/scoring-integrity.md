# Chore: Scoring Integrity — Only Score What We Measure

## Chore Description

The current scoring system includes 9 dimensions in the weighted average, but only 5-7 are actually measured by evidence. The unmeasured dimensions sit at a hardcoded 0.80 base, acting as constant diluters that make all scores converge toward 0.80 regardless of actual code quality. This creates several integrity problems:

### Problem 1: Fictional Dimensions

Two dimensions are **never measured** for any unit — they always report 0.80:

| Dimension | Evidence That Moves It | Units Measured | Reality |
|-----------|----------------------|----------------|---------|
| `architectural_fitness` | Only penalized: `context_not_first`, `method_count > 15` | 0/748 (0%) | Never triggered in practice |
| `performance_appropriateness` | Only penalized: `defer_in_loop` | 0/748 (0%) | Never triggered in practice |

These contribute ~22% of the weighted average (2/9) but carry zero signal. They make every score 2-3% worse than the measured reality.

### Problem 2: One-Way Dimensions

`security` only moves **downward** from 0.80 via `global_mutable_count` penalty. It has no positive measurement — 78% of units sit at 0.80 base, and the rest are penalized. A unit with no global vars gets 0.80 security, but there's no way to earn higher. Security is only "absence of one bad pattern" not "presence of good practices."

### Problem 3: Agent Review Blanket Boost

When agent review passes, ALL 9 dimensions get `max(score, 0.85)` — including the fictional ones. This means agent review inflates unmeasured dimensions from 0.80 → 0.85 with zero actual quality assessment.

### Problem 4: Base Score Is Arbitrary

The 0.80 base was chosen arbitrarily. It means a unit with NO evidence at all gets a B grade (80%). This is generous — a unit with zero evidence should arguably score 0 or be marked as unmeasurable.

## Audit: What's Actually Measured

| Dimension | Positive Evidence | Negative Evidence | Coverage |
|-----------|------------------|-------------------|----------|
| **correctness** | lint pass → 0.95 | lint fail, errors_ignored, panic, os_exit, defer_in_loop, context_not_first | 100% |
| **testability** | test pass → 0.90, coverage ≥80% → 0.95 | test fail, os_exit, init(), global_mutable_count | 100% |
| **maintainability** | complexity ≤5 → 0.95 | high complexity, params >5, func_lines >100, method_count >10, init() | 95% |
| **readability** | code_lines ≤50 → 0.95, has_doc → 0.90, func_lines ≤30 → 0.90 | no doc exported, nesting >3, naked_returns, func_lines >100 | 97% |
| **change_risk** | multi-author → 0.90 | *(none)* | 100% |
| **operational_quality** | commits >10 → 0.85 | *(none)* | 100% |
| **security** | *(none)* | global_mutable_count penalty | 22% (down only) |
| **architectural_fitness** | *(none)* | context_not_first, method_count >15 | 0% |
| **performance_appropriateness** | *(none)* | defer_in_loop | 0% |

## Relevant Files

- `internal/engine/scorer.go` — Score function and all dimension scoring logic; the core file that needs changes
- `internal/engine/scorer_test.go` — Tests for scorer; needs updated assertions for new behavior
- `internal/domain/dimension.go` — Dimension definitions, `WeightedAverage()`, `AllDimensions()`; needs `MeasuredDimensions` concept
- `internal/domain/dimension_test.go` — May need new tests for measured dimension tracking
- `internal/report/card.go` — Report card generation; should flag unmeasured dimensions
- `internal/report/report_tree.go` — Unit reports; should distinguish measured vs unmeasured dimensions
- `.certification/policies/go-standard.yml` — Policy rules that reference dimensions

## Step by Step Tasks

### Step 1: Track which dimensions have evidence (TDD)

- Add `measured` tracking to the `Score` function
- Return both `DimensionScores` and a `map[Dimension]bool` indicating which dimensions received actual evidence
- Write test: a unit with only lint evidence should mark `correctness` as measured, `architectural_fitness` as unmeasured
- Run test → fail → implement → pass

### Step 2: Exclude unmeasured dimensions from weighted average

- Modify `WeightedAverage` to accept an optional `measured` set
- Only include dimensions that have actual evidence in the average
- A unit with 7 measured dimensions should compute average over 7, not 9
- Write test: average of measured dims only should differ from all dims
- Run test → fail → implement → pass

### Step 3: Remove the 0.80 universal base

- Unmeasured dimensions should report as "unmeasured" (e.g., -1 or `NaN` or absent from the map) rather than 0.80
- Measured dimensions keep their evidence-based scores
- Write test: unmeasured dimension should not appear in scores map
- Run test → fail → implement → pass

### Step 4: Fix agent review blanket boost

- Agent review should only boost dimensions it actually assesses
- Replace `for all dims: max(score, 0.85)` with targeted boosts based on what the agent actually reviewed
- At minimum, only boost correctness, maintainability, readability, security (things an LLM can actually assess from code)
- Write test: agent review should not boost architectural_fitness or performance_appropriateness
- Run test → fail → implement → pass

### Step 5: Make reports transparent

- Unit reports should show which dimensions are measured vs unmeasured
- Report card should note the number of measured dimensions per unit
- Grade should be computed from measured dimensions only

### Step 6: Re-certify and validate

- Full re-certification with the new scoring
- Verify: scores now reflect actual measured quality
- Verify: units with more evidence types have more measured dimensions
- Verify: grade spread increases (bad code scores lower, good code scores higher)
- Verify: no unit gets credit for unmeasured dimensions

## Validation Commands

```bash
go test ./... -count=1
go build -o build/bin/certify ./cmd/certify/
./build/bin/certify certify --skip-agent --batch 0 --reset-queue
./build/bin/certify report

# Verify no dimension is universally at 0.80 base
python3 -c "
import json, glob, collections
with open('.certification/index.json') as f:
    index = json.load(f)
index_ids = set(u['id'] for u in index)
dim_at_base = collections.Counter()
total = 0
for f in glob.glob('.certification/records/*.json'):
    with open(f) as fh:
        r = json.load(fh)
    if r['unit_id'] not in index_ids: continue
    total += 1
    dims = r.get('dimensions', {})
    for d, v in dims.items():
        if abs(v - 0.80) < 0.001:
            dim_at_base[d] += 1
for d in sorted(dim_at_base.keys()):
    pct = dim_at_base[d] / total * 100
    if pct > 50:
        print(f'WARNING: {d} still at base for {pct:.0f}% of units')
        assert False, f'{d} is mostly unmeasured'
print('ALL dimensions have real measurements')
"
```

## Notes

- **This will change scores significantly.** Removing fictional dimensions from the average means scores will shift. That's the point — the current scores are partially fictional.
- **Consider removing fictional dimensions entirely** until real evidence exists. It's better to have 7 honest dimensions than 9 partially fictional ones.
- **`security` needs real measurement** — static analysis for input validation, SQL injection, path traversal, etc. The current "count global vars" proxy is weak.
- **`architectural_fitness` needs real measurement** — dependency depth, circular imports, layer violations, interface segregation. Not currently measurable from structural evidence alone.
- **`performance_appropriateness` needs real measurement** — benchmark results, allocation analysis, complexity class. Not currently measurable.
- **Don't add fake boosts to fill the gap.** If we can't measure it, we should say "unmeasured" not "0.80."
- **The PRD lists 9 dimensions** — we should keep tracking all 9 but clearly distinguish measured from aspirational in reports.
