# Agent Review — Redesign

## Design Patterns

### Pipeline Pattern (Chain of Responsibility)
Each review step is a `Stage` with a clear contract. Stages are composed into
a Pipeline. Each stage can short-circuit (skip remaining stages).

### Strategy Pattern
Different review strategies for different contexts:
- **Quick**: Prescreen only — cheap gate, no deep review
- **Standard**: Prescreen → Review → Score (3 calls max)
- **Full**: All 5 stages (only for failing/borderline units)

### Circuit Breaker
Wraps the Provider. After N consecutive failures, opens the circuit and
returns cached/fallback results without calling the API.

### Batch Coordinator
Owns the file-level dedup and token budget. The certify command delegates
to the coordinator instead of managing review state itself.

## Architecture

```
certify_cmd.go
  └─ ReviewCoordinator          (batch orchestration)
       ├─ CircuitBreaker         (wraps Provider)
       │    └─ RateLimitedProvider (wraps Provider with RateLimiter)
       │         └─ OpenRouterProvider
       ├─ Pipeline               (chain of stages)
       │    ├─ PrescreenStage
       │    ├─ ReviewStage
       │    ├─ ScoringStage
       │    ├─ DecisionStage
       │    └─ RemediationStage
       ├─ PromptRegistry         (loads + caches templates)
       └─ Router                 (task → model mapping)
```

## Key Decisions
- Prescreen parses JSON loosely (regex fallback for `needs_review`)
- Token budget: coordinator tracks total tokens, stops when budget exhausted
- File-level dedup lives in coordinator, not in certify_cmd
- Each stage returns `(StageResult, bool)` where bool = continue pipeline
- Pipeline selects strategy based on deterministic evidence quality
