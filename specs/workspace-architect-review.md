# Workspace-Scoped Architect Review

**Status:** Draft  
**Date:** 2026-03-11  
**Author:** Claude  
**Relates to:** `internal/workspace/`, `internal/agent/architect*.go`, `cmd/certify/architect_cmd.go`

## Problem

The architect review command (`certify architect`) has no `--workspace` support. In a multi-repo workspace with git submodules, each submodule can be individually certified but there is no holistic architectural review that:

1. Sees all submodules as components of a single system
2. Analyzes cross-submodule dependencies and integration boundaries
3. Identifies architectural concerns that only emerge at the workspace level (e.g., circular dependency chains between submodules, shared library version drift, inconsistent API contracts)
4. Treats workspace-level code (root-level scripts, CI, shared config) as infrastructure

Currently, `--workspace` is supported by: `init`, `scan`, `certify`, `report`, `expire`.  
Missing: `architect`.

## Design Principles

1. **Workspace = system, submodules = components** — The architect review should reason about the whole system, not just iterate per-submodule.
2. **Workspace-level source is infrastructural** — Root-level code (build scripts, CI, shared config, Makefiles, Justfiles) should be categorized as infrastructure, not application logic.
3. **Submodule snapshots are composed, not concatenated** — Build a single `WorkspaceArchSnapshot` that merges submodule data into a coherent cross-repo view.
4. **Each submodule's data stays grounded** — Per-submodule metrics come from their existing certification records. No re-scanning.

## Architecture

### New Types

```go
// WorkspaceArchSnapshot extends ArchSnapshot with cross-submodule data.
type WorkspaceArchSnapshot struct {
    SchemaVersion      int                      `json:"schema_version"`
    SubmoduleSnapshots []SubmoduleSnapshotEntry  `json:"submodule_snapshots"`
    CrossDependencies  []CrossDepEdge           `json:"cross_dependencies"`
    InfraFiles         []string                 `json:"infra_files"`
    AggregateMetrics   WorkspaceMetrics         `json:"aggregate_metrics"`
}

type SubmoduleSnapshotEntry struct {
    Name     string        `json:"name"`
    Path     string        `json:"path"`
    Commit   string        `json:"commit"`
    Snapshot *ArchSnapshot `json:"snapshot"` // nil if no cert data
    Role     string        `json:"role"`     // "service", "library", "tool", "infrastructure"
}

type CrossDepEdge struct {
    FromSubmodule string `json:"from_submodule"`
    ToSubmodule   string `json:"to_submodule"`
    Evidence      string `json:"evidence"` // how detected (go.mod replace, import path, config reference)
    Weight        int    `json:"weight"`
}

type WorkspaceMetrics struct {
    TotalSubmodules      int     `json:"total_submodules"`
    ConfiguredSubmodules int     `json:"configured_submodules"`
    TotalUnitsAcrossAll  int     `json:"total_units_across_all"`
    WeightedAvgScore     float64 `json:"weighted_avg_score"`
    WorstSubmodule       string  `json:"worst_submodule"`
    BestSubmodule        string  `json:"best_submodule"`
}
```

### Snapshot Building

`BuildWorkspaceSnapshot(root string, subs []workspace.Submodule) *WorkspaceArchSnapshot`:
1. For each configured submodule, load its certification records and call `BuildSnapshot(records, subRoot)`
2. Classify each submodule's role based on path heuristics and content (presence of `main.go` in `cmd/` → service/tool, exported packages → library)
3. Detect cross-submodule dependencies by scanning `go.mod` replace directives and import paths
4. Identify workspace-level infrastructure files (root `Justfile`, `Makefile`, `.github/`, `docker-compose.yml`, etc.)
5. Compute aggregate metrics

### LLM Context Formatting

`FormatWorkspaceForLLM(snap *WorkspaceArchSnapshot, maxTokensHint int) string`:

Produces a structured text block:
```
# Workspace Architecture Snapshot

**Schema:** v2 (workspace)
**Submodules:** 4 (3 configured)
**Total Units:** 1,247

## Submodule Overview
| Submodule | Role | Units | Avg Score | Grade | Packages |
|-----------|------|------:|----------:|:-----:|--------:|

## Cross-Submodule Dependencies
| From | To | Evidence |
|------|----|---------| 

## Infrastructure Files
- Justfile (build orchestration)
- .github/workflows/ (CI/CD)
- docker-compose.yml (local dev environment)

## Per-Submodule Architecture Summaries
### api-service (services/api)
[condensed single-submodule snapshot]

### shared-lib (lib/shared)
[condensed single-submodule snapshot]

## Workspace-Level Structural Metrics
[aggregated across all submodules]
```

### Phase Prompts

The workspace architect review uses the same 6-phase structure but with workspace-aware system prompts:

- **Phase 1 (Architecture Narration):** Describe the system as a whole. What role does each submodule play? How do they compose? What are the integration boundaries?
- **Phase 2 (Code Quality):** Cross-submodule quality patterns. Shared library quality affecting consumers. Inconsistent patterns between submodules.
- **Phase 3 (Test Strategy):** Integration testing across submodule boundaries. Contract testing. Are shared libraries adequately tested for their consumers?
- **Phase 4 (Security & Ops):** Cross-submodule security surface. Shared secrets handling. Deployment coupling. Infrastructure-as-code quality.
- **Phase 5 (Recommendations):** System-level improvements. Submodule boundary changes. Shared library extraction opportunities.
- **Phase 6 (Synthesis):** System-level executive summary referencing submodule interactions.

### CLI Integration

```
certify architect --workspace    # workspace-level review
certify architect                # single-repo review (unchanged)
```

The `--workspace` flag:
1. Discovers submodules via `workspace.DiscoverSubmodules(root)`
2. Builds `WorkspaceArchSnapshot` instead of `ArchSnapshot`
3. Uses workspace-specific phase prompts
4. Writes output to workspace-level `.certification/ARCHITECT_REVIEW.md`

## Work Items

### WI-1: WorkspaceArchSnapshot types and builder
- Add `WorkspaceArchSnapshot`, `SubmoduleSnapshotEntry`, `CrossDepEdge`, `WorkspaceMetrics` types
- Implement `BuildWorkspaceSnapshot()` — loads per-submodule records, builds per-submodule snapshots, detects cross-deps
- Implement submodule role classification heuristic
- Implement cross-dependency detection (go.mod replace, shared import paths)
- Implement infra file detection

### WI-2: Workspace LLM context formatting
- Implement `FormatWorkspaceForLLM()` — renders workspace snapshot for LLM consumption
- Condense per-submodule snapshots to fit token budget (top-level metrics + hotspots only)
- Include cross-dependency graph and infrastructure listing

### WI-3: Workspace phase prompts
- Write 6 workspace-aware system prompts
- All prompts include anti-hallucination grounding language
- Prompts reference workspace-specific data sections (cross-deps, submodule roles, infra files)

### WI-4: Workspace ProjectContext and review pipeline
- Add `WorkspaceProjectContext` (or extend `ProjectContext` with workspace fields)
- Wire workspace snapshot building into `GatherWorkspaceContext()`
- Ensure `ArchitectReviewer.Review()` works with workspace context (may need workspace-specific prompt selection)

### WI-5: CLI integration
- Add `--workspace` handling to `architect_cmd.go`
- Implement `runWorkspaceArchitect()` 
- Write output to workspace-level `.certification/ARCHITECT_REVIEW.md`

### WI-6: Workspace architect report formatting
- Extend `report.FormatArchitectReport()` to handle workspace results
- Include per-submodule summaries in report
- Include cross-dependency visualization

## Test Plan

### Unit Tests (TDD — write first)
1. `TestBuildWorkspaceSnapshot_MultipleSubmodules` — builds snapshot from 2+ submodule record sets
2. `TestBuildWorkspaceSnapshot_ClassifiesRoles` — service vs library vs tool detection  
3. `TestBuildWorkspaceSnapshot_DetectsCrossDeps` — go.mod replace directive parsing
4. `TestBuildWorkspaceSnapshot_InfraFiles` — detects Justfile, Makefile, CI configs
5. `TestBuildWorkspaceSnapshot_AggregateMetrics` — weighted avg, best/worst submodule
6. `TestFormatWorkspaceForLLM_Structure` — output contains expected sections
7. `TestFormatWorkspaceForLLM_TokenBudget` — respects size limits
8. `TestWorkspacePhasePrompts_AllContainGrounding` — anti-hallucination in all 6
9. `TestGatherWorkspaceContext` — end-to-end context building from temp dirs

### Integration Tests
10. `TestWorkspaceArchitect_EndToEnd` — full pipeline with mock provider

## Commit Plan

Single commit: `feat(architect): add workspace-scoped architectural review`
