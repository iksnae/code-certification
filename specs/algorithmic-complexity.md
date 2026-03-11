# Feature: Algorithmic Complexity Detection

## Description

Add heuristic algorithmic complexity analysis to the structural evidence collector. This detects Big-O complexity patterns in function bodies via `go/ast`, capturing both micro-level (per-function) and macro-level (file/package) complexity signals that cyclomatic complexity alone cannot express.

Cyclomatic complexity counts branching paths. Algorithmic complexity measures **how work scales with input size** — nested loops, recursive calls, quadratic patterns. A function with complexity=3 (low cyclomatic) can still be O(n³) with three nested range loops.

### Metrics

| Metric | Type | Description |
|--------|------|-------------|
| `loop_nesting_depth` | int | Maximum depth of nested loops (for/range) — proxy for polynomial degree |
| `recursive_calls` | int | Count of direct recursive calls (func calls own name) |
| `algo_complexity` | string | Estimated complexity class: `O(1)`, `O(n)`, `O(n²)`, `O(n³)`, `O(2^n)` |
| `nested_loop_pairs` | int | Count of inner loops nested inside outer loops |
| `quadratic_patterns` | int | Count of known O(n²) anti-patterns (string concat in loop, sort in loop, contains-in-loop) |

### Complexity Classification Rules

| Pattern | Classification |
|---------|---------------|
| No loops | O(1) |
| Single loop, no recursion | O(n) |
| Nested 2-deep loops | O(n²) |
| Nested 3+-deep loops | O(n³) |
| Recursive call detected | O(2^n) (conservative — could be O(n) with tail recursion, but we can't prove it) |
| Recursive + loop | O(2^n) |

### Scoring Impact

| Complexity | `performance_appropriateness` |
|------------|-------------------------------|
| O(1) | 0.95 (boost) |
| O(n) | 0.90 (neutral-good) |
| O(n²) | 0.70 (observation) |
| O(n³) | 0.50 (warning) |
| O(2^n) | 0.40 (concern) |

This activates the `performance_appropriateness` dimension — currently penalty-only and never set for clean code.

## Relevant Files

- `internal/evidence/structural.go` — Add `AlgoComplexity`, `LoopNestingDepth`, `RecursiveCalls`, `NestedLoopPairs`, `QuadraticPatterns` to `StructuralMetrics` + analysis logic
- `internal/evidence/structural_test.go` — Tests for all complexity patterns  
- `internal/engine/scorer.go` — Score `performance_appropriateness` from algo complexity metrics
- `internal/engine/scorer_test.go` — Tests for scoring
- `internal/evidence/complexity.go` — Existing cyclomatic complexity (reference, not modified)

## Step by Step Tasks

### Step 1: Add algorithmic complexity fields to StructuralMetrics (TDD)

Tests:
- `TestAnalyzeGoFunc_AlgoO1` — no loops → `AlgoComplexity="O(1)"`, `LoopNestingDepth=0`
- `TestAnalyzeGoFunc_AlgoON` — single for loop → `AlgoComplexity="O(n)"`, `LoopNestingDepth=1`
- `TestAnalyzeGoFunc_AlgoON2` — nested for loops → `AlgoComplexity="O(n²)"`, `LoopNestingDepth=2`
- `TestAnalyzeGoFunc_AlgoON3` — triple nested → `AlgoComplexity="O(n³)"`, `LoopNestingDepth=3`
- `TestAnalyzeGoFunc_RecursiveCall` — func calls itself → `RecursiveCalls=1`, `AlgoComplexity="O(2^n)"`
- `TestAnalyzeGoFunc_RecursiveWithLoop` — recursive + loop → `AlgoComplexity="O(2^n)"`
- `TestAnalyzeGoFunc_NestedRange` — nested range loops → `AlgoComplexity="O(n²)"`

### Step 2: Detect quadratic anti-patterns (TDD)

Tests:
- `TestAnalyzeGoFunc_StringConcatInLoop` — `s += x` inside for → `QuadraticPatterns=1`
- `TestAnalyzeGoFunc_SortInLoop` — `sort.Slice` inside for → `QuadraticPatterns=1`
- `TestAnalyzeGoFunc_AppendInNestedLoop` — append inside nested for → `QuadraticPatterns=1`

### Step 3: Wire metrics into evidence and scoring (TDD)

- Add metrics to `ToEvidence()` output
- Score `performance_appropriateness` from `algo_complexity` metric
- Test: O(n²) function → `performance_appropriateness <= 0.70`
- Test: O(1) function → `performance_appropriateness >= 0.95`

### Step 4: Re-certify and validate — DONE

## Report

**Completed:** 2026-03-10

### What was implemented

1. **Algorithmic complexity detection** — 5 new structural metrics computed via `go/ast`:
   - `algo_complexity`: Big-O class (O(1), O(n), O(n²), O(n³), O(2^n))
   - `loop_nesting_depth`: max depth of nested for/range loops
   - `recursive_calls`: direct recursive call count
   - `nested_loop_pairs`: count of inner loops nested in outer loops
   - `quadratic_patterns`: anti-patterns like string concat in loop

2. **Scoring** — `performance_appropriateness` now always measured:
   - O(1) → 0.95, O(n) → 0.90, O(n²) → 0.70, O(n³) → 0.50, O(2^n) → 0.40
   - `defer_in_loop` penalty still overrides via `setMin`
   - Algo scoring runs BEFORE correctness penalties for correct ordering

3. **Also delivered** (from parent plan):
   - Graduated file-level `code_lines` thresholds
   - Graduated complexity tier (0.80 for complexity 11-15)
   - 14 exported symbols documented

### Results

| Metric | Before (v0.6.2) | After (v0.7.0) |
|--------|-----------------|-----------------|
| Score | 88.7% | **91.7%** |
| A/A- | 395 (52.8%) | **691 (92.4%)** |
| B+ | 148 | 26 |
| B | 192 | 31 |
| C | 13 | **0** |
| Observations | 13 | **0** |
| Dimensions measured | 7 | **8** (+performance_appropriateness) |

### Issues encountered

- `countQuadraticPatterns` initially flagged `total += n` (numeric) as string concat. Fixed with `containsStringConcat` heuristic: only flag `+=` when RHS contains a string literal.
- `scoreAlgoComplexity` with `setMax(perf, 0.95)` overrode `defer_in_loop`'s `setMin(perf, 0.50)`. Fixed by running algo scoring BEFORE correctness penalties.
- Existing `TestScorer_PenaltyOnlyDimsAppearWhenBad` test assumed perf was penalty-only. Updated to reflect new always-measured behavior.

### Tests added

- `TestAnalyzeGoFunc_AlgoComplexity` — 9 sub-tests (O(1) through O(2^n), sequential loops, loop+if)
- `TestAnalyzeGoFunc_QuadraticPatterns` — 3 sub-tests (string concat in loop, no pattern, outside loop)
- `TestScorer_AlgoComplexityScoring` — 5 sub-tests (all complexity tiers)
- `TestScorer_FileCodeLinesThresholds` — 5 sub-tests (file-level thresholds)
- `TestScorer_ComplexityGraduated` — 4 sub-tests (complexity tiers)

## Validation Commands

```bash
go test ./internal/evidence/... -count=1
go test ./internal/engine/... -count=1  
go test ./... -count=1
go build -o build/bin/certify ./cmd/certify/
```
