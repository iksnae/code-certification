# Plan: Multi-Language Deep Analysis

## Objective

Bring every supported language to parity with Go's current AST analysis, then surpass it with type-aware cross-file analysis. Fill all 9 quality dimensions with deterministic, evidence-based signals — no LLM required for scoring.

## Current State (after Sprints 1–6)

### Evidence coverage by language

| Evidence | Go | TypeScript | Python | Rust |
|----------|:---:|:----------:|:------:|:----:|
| Symbol-level discovery | ✅ go/ast | ✅ tree-sitter (108 symbols) | ✅ tree-sitter | ✅ tree-sitter |
| Cyclomatic complexity | ✅ per-function | ✅ tree-sitter | ✅ tree-sitter | ✅ tree-sitter |
| Cognitive complexity | ✅ Sonar-style | ✅ tree-sitter | ✅ tree-sitter | ✅ tree-sitter |
| Structural metrics (27) | ✅ full go/ast | ✅ 22 metrics | ✅ 22 metrics | ✅ 22 metrics |
| Algo complexity | ✅ loop analysis | ✅ loop/recursion | ✅ loop/recursion | ✅ loop/recursion |
| Unsafe import detection | ✅ 8 packages | ✅ child_process/eval/vm | ✅ subprocess/pickle/ctypes | ✅ std::process/libc |
| Hardcoded secrets | ✅ pattern match | ✅ pattern match | ✅ pattern match | ✅ pattern match |
| Error handling patterns | ✅ ignored/unwrapped | ✅ empty catch | ✅ empty except:pass | ✅ panic!/unwrap() |
| Lint | ✅ go vet + golangci-lint | ✅ ESLint | ✅ ruff | ✅ cargo clippy |
| Tests | ✅ go test | ✅ Jest/Vitest | ✅ pytest | ✅ cargo test |
| Coverage | ✅ per-unit go cover | ⚠️ LCOV (if present) | ⚠️ Cobertura (if present) | ❌ (tarpaulin planned) |
| Git history | ✅ | ✅ | ✅ | ✅ |
| Code metrics (lines) | ✅ | ✅ | ✅ | ✅ |
| Nested module support | ✅ go.mod in subdirs | ✅ package.json in subdirs | ✅ pyproject.toml in subdirs | ✅ Cargo.toml in subdirs |
| **Fan-in / fan-out** | ❌ (Sprint 7) | ❌ (Sprint 11) | ❌ (Sprint 11) | ❌ |
| **Dead code** | ❌ (Sprint 7) | ❌ (Sprint 11) | ❌ (Sprint 11) | ❌ |
| **Interface compliance** | ❌ (Sprint 8) | ❌ | ❌ | ❌ |
| **Package dep graph** | ❌ (Sprint 8) | ❌ | ❌ | ❌ |

### Dimension signal strength (current)

| Dimension | Go | TS | Python | Rust |
|-----------|:---:|:---:|:------:|:----:|
| correctness | █████████░ 9 | ███████░░░ 7 | ███████░░░ 7 | ███████░░░ 7 |
| maintainability | ████████░░ 8 | ██████░░░░ 6 | ██████░░░░ 6 | ██████░░░░ 6 |
| readability | █████████░ 9 | ███████░░░ 7 | ███████░░░ 7 | ███████░░░ 7 |
| testability | ██████░░░░ 6 | ████░░░░░░ 4 | ████░░░░░░ 4 | ████░░░░░░ 4 |
| security | ████░░░░░░ 4 | ███░░░░░░░ 3 | ███░░░░░░░ 3 | ███░░░░░░░ 3 |
| arch_fitness | ██░░░░░░░░ 2 | █░░░░░░░░░ 1 | █░░░░░░░░░ 1 | █░░░░░░░░░ 1 |
| operational_quality | ███░░░░░░░ 3 | ██░░░░░░░░ 2 | ██░░░░░░░░ 2 | ██░░░░░░░░ 2 |
| performance | ████░░░░░░ 4 | ███░░░░░░░ 3 | ███░░░░░░░ 3 | ███░░░░░░░ 3 |
| change_risk | ███░░░░░░░ 3 | ██░░░░░░░░ 2 | ██░░░░░░░░ 2 | ██░░░░░░░░ 2 |

Remaining gaps: arch_fitness (needs dep graph/call graph), change_risk (needs fan-in), testability (needs interface compliance/dead code).

---

## Architecture

### Analysis tiers

```
┌─────────────────────────────────────────────────────────────────┐
│                        Evidence Pipeline                         │
├─────────────┬───────────────┬─────────────────┬─────────────────┤
│  Tier 0     │  Tier 1       │  Tier 2         │  Tier 3         │
│  Universal  │  Syntax AST   │  Type-Aware     │  Cross-Project  │
│  (today)    │  (per-file)   │  (per-project)  │  (workspace)    │
├─────────────┼───────────────┼─────────────────┼─────────────────┤
│ Git history │ Tree-sitter   │ go/types+pkgs   │ Workspace dep   │
│ Line counts │ or native AST │ LSP for non-Go  │   graph         │
│ TODO/FIXME  │               │                 │ Cross-module    │
│ Lint tools  │ Discovery     │ Call hierarchy   │   coupling      │
│ Test runner │ Nesting depth │ Fan-in/fan-out  │ API surface     │
│ Coverage    │ Params/returns│ Dead code       │   stability     │
│             │ Complexity    │ Interface impl  │                 │
│             │ Doc comments  │ Package cycles  │                 │
│             │ Import list   │ Type safety     │                 │
│             │ Error patterns│ Dep inversion   │                 │
│             │ Security sigs │                 │                 │
│             │ Cognitive cplx│                 │                 │
├─────────────┼───────────────┼─────────────────┼─────────────────┤
│ All langs   │ All langs     │ Go: native      │ All langs via   │
│ No deps     │ tree-sitter   │ TS: tsserver    │   submodule     │
│             │ (CGo) or      │ Py: pyright     │   config        │
│             │ Go: go/ast    │ Rs: rust-anlzr  │                 │
│             │ (no deps)     │ Optional—needs  │                 │
│             │               │ toolchain       │                 │
└─────────────┴───────────────┴─────────────────┴─────────────────┘
```

### Technology choices per tier

| Tier | Go | TypeScript | Python | Rust | Others |
|------|-----|-----------|--------|------|--------|
| **T0: Universal** | git, go test, go vet | git | git | git | git |
| **T1: Syntax** | `go/ast` (keep) | tree-sitter-typescript | tree-sitter-python | tree-sitter-rust | tree-sitter-* |
| **T2: Types** | `go/types` + `go/packages` | `tsserver` (subprocess) | `pyright` (subprocess) | `rust-analyzer` (subprocess) | LSP subprocess |
| **T3: Workspace** | `go/packages` multi-module | workspace config | workspace config | cargo workspace | workspace config |

### Why tree-sitter for Tier 1 (not more regex adapters)

- One parser framework → one code path for structural metrics across all languages
- Concrete syntax trees, not abstract — preserves comments, whitespace, exact positions
- Grammar files are maintained by large communities (GitHub, Neovim, Zed)
- Incremental parsing — fast enough for large repos
- We keep `go/ast` for Go because it's faster and richer (comments, type hints)

### Why LSP subprocess for Tier 2 (not tree-sitter)

Tree-sitter is syntax-only — it cannot:
- Resolve `import { Foo } from "./bar"` to the actual definition
- Tell you that `MyStruct` implements `io.Reader`
- Build a call graph across files
- Detect unused exports
- Know that `x` is type `error` vs `string`

Only a type checker can do this. For Go we have `go/types` in-process. For other languages, the cheapest path to a type checker is the language's own LSP server — it's already built, tested, and maintained by the language team.

### Why subprocess (not embedded)

Embedding `tsserver` or `pyright` means bundling Node.js. Embedding `rust-analyzer` means bundling a Rust binary. Instead:
- Detect installed language servers (same pattern as `certify doctor`)
- Spawn as subprocess, communicate via stdin/stdout JSON-RPC
- Graceful degradation: if no LSP available, Tier 2 evidence is simply absent
- `certify doctor` already checks for optional tools — LSP servers become another optional tool

---

## Implementation Plan

### Phase 1: Unified Syntax Analysis (Tier 1)
**Goal:** Every language gets the same 22+ structural metrics Go currently has.

#### 1A: `internal/analysis/` package — abstract analyzer interface

```go
// Analyzer provides language-agnostic structural analysis.
type Analyzer interface {
    // Discover finds all symbols in a source file.
    Discover(path string, src []byte) ([]Symbol, error)

    // Analyze returns structural metrics for a specific symbol.
    Analyze(path string, src []byte, symbol string) (Metrics, error)

    // AnalyzeFile returns file-level metrics.
    AnalyzeFile(path string, src []byte) (FileMetrics, error)
}

// Symbol represents a discovered code unit.
type Symbol struct {
    Name      string
    Kind      SymbolKind // Function, Method, Class, Interface, Type, Constant
    StartLine int
    EndLine   int
    Parent    string     // enclosing type/class (empty for top-level)
}

// Metrics mirrors current StructuralMetrics but is language-agnostic.
type Metrics struct {
    // Shape
    ParamCount      int
    ReturnCount     int
    FuncLines       int
    MaxNestingDepth int

    // Documentation
    HasDocComment bool
    IsExported    bool

    // Complexity
    CyclomaticComplexity int
    CognitiveComplexity  int  // NEW — Sonar-style
    LoopNestingDepth     int
    RecursiveCalls       int

    // Error handling
    ErrorsIgnored    int
    ErrorsNotWrapped int  // NEW — fmt.Errorf without %w, catch without rethrow
    NakedReturns     int
    PanicCalls       int  // or throw without catch, unwrap() in Rust

    // Security
    UnsafeImports    []string // NEW — os/exec, unsafe, eval, subprocess
    HardcodedSecrets int      // NEW — string literals matching secret patterns

    // Design
    MethodCount     int
    IsConstructor   bool
    DeferInLoop     int
    ContextNotFirst bool  // Go-specific, but interface allows lang-specific flags

    // Performance
    AlgoComplexity    string
    NestedLoopPairs   int
    QuadraticPatterns int

    // Language-specific extras (opaque map for non-standard metrics)
    Extra map[string]float64
}
```

**Test strategy:** Each analyzer tested against known source snippets with expected metrics. Compare Go analyzer output against current `AnalyzeGoFunc` output for identical inputs — must match exactly.

#### 1B: Go analyzer — wrap existing `go/ast` code

Refactor `internal/evidence/structural.go` functions into the `Analyzer` interface. Zero behavior change — existing tests continue to pass.

```
internal/analysis/
    analyzer.go      — interface definitions
    go_analyzer.go   — wraps existing go/ast analysis
    go_analyzer_test.go
```

#### 1C: Tree-sitter analyzers for TS, Python, Rust

```
internal/analysis/
    treesitter.go          — shared tree-sitter helpers
    ts_analyzer.go         — TypeScript / JavaScript
    ts_analyzer_test.go
    python_analyzer.go     — Python
    python_analyzer_test.go
    rust_analyzer.go       — Rust
    rust_analyzer_test.go
```

Each implements the same `Analyzer` interface. What tree-sitter gives us per language:

**TypeScript:**
- `function_declaration`, `method_definition`, `class_declaration` → discovery
- `formal_parameters` → param count
- `return_statement` → return count
- `if_statement`, `for_statement`, `while_statement` → nesting depth
- `binary_expression` with `&&`/`||` → cyclomatic complexity
- `try_statement`, `catch_clause` → error handling patterns
- `import_statement` → security import analysis
- `comment` → doc comment detection (JSDoc `/** */`)
- `template_string` with `${...}` inside loops → quadratic pattern (string building)

**Python:**
- `function_definition`, `class_definition` → discovery
- `parameters` → param count
- `return_statement` → returns
- `if_statement`, `for_statement`, `while_statement`, `with_statement` → nesting
- `except_clause` with `pass` → errors ignored
- `import_statement`, `import_from_statement` → security (subprocess, eval, pickle, os.system)
- `decorator` → pattern detection (@staticmethod, @property)
- `raise` in `except` → error wrapping check
- Comments, docstrings (`expression_statement` > `string`) → documentation

**Rust:**
- `function_item`, `impl_item` → discovery
- `parameters` → param count
- `match_expression`, `if_expression`, `loop_expression` → nesting
- `unsafe_block` → security
- `.unwrap()`, `.expect()` → error handling (panic equivalent)
- `use_declaration` → imports
- `/// doc` comments → documentation
- `async fn`, `.await` → concurrency patterns

#### 1D: New structural metrics (all languages)

Added to the `Metrics` struct and computed by all analyzers:

| Metric | Dimension | What it detects | Method |
|--------|-----------|----------------|--------|
| `cognitive_complexity` | readability | Sonar-style: nesting increment + flow break | AST walk with depth counter |
| `errors_not_wrapped` | operational_quality | Go: `fmt.Errorf` without `%w`. TS: `catch` without `throw`. Py: bare `except: pass` | AST pattern match |
| `unsafe_imports` | security | Language-specific dangerous import list | Import node inspection |
| `hardcoded_secrets` | security | String literals matching `password=`, `secret=`, API key patterns, high-entropy base64 | String literal scan with regex |
| `assertion_density` | testability | In test files: assertions per test function | AST count of assert/expect/require calls |
| `type_annotation_pct` | readability | Python/TS: % of params with type annotations | AST param type presence check |
| `empty_catch_blocks` | correctness | catch/except/recover with empty body | AST block emptiness check |
| `magic_numbers` | readability | Numeric literals outside const/define (except 0, 1, -1) | AST literal inspection |
| `deeply_nested_callbacks` | readability | TS/JS: callback depth > 3 (callback hell) | AST nesting of function expressions |

#### 1E: Wire into evidence pipeline

Modify `Certifier.collectStructuralEvidence()`:

```go
func (c *Certifier) collectStructuralEvidence(unit domain.Unit, srcCode string, ev *[]domain.Evidence) {
    lang := unit.ID.Lang()
    analyzer := analysis.ForLanguage(lang) // returns Go, TS, Python, Rust, or nil
    if analyzer == nil {
        return
    }
    metrics, err := analyzer.Analyze(unit.ID.Path(), []byte(srcCode), unit.ID.Symbol())
    if err != nil {
        return
    }
    *ev = append(*ev, metrics.ToEvidence())
}
```

**Migration:** The `isGo` check disappears. Every language goes through the same path.

#### 1F: Replace regex TypeScript discovery

`TSAdapter.Scan()` switches from regex to tree-sitter. Finds all symbols, not just exports. Includes line ranges for per-unit metrics attribution.

#### 1G: Add language-specific lint tool integration

Extend `ToolExecutor.CollectAll()` to detect and run:

| Language | Lint Tool | Detection |
|----------|-----------|-----------|
| Go | go vet + golangci-lint | ✅ already done |
| TypeScript | `eslint` or `biome` | `which eslint` or check `node_modules/.bin/eslint` |
| Python | `ruff` (fast, covers flake8+pylint+pyflakes) | `which ruff` |
| Rust | `cargo clippy` | `which cargo` |
| Java | `checkstyle` or `pmd` | `which checkstyle` |

Each produces `LintFinding` structs via output parsing. Per-unit attribution works the same way (filter by file + line range).

#### 1H: Add language-specific test runner integration

Extend `ToolExecutor` to detect and run:

| Language | Test Command | Coverage |
|----------|-------------|----------|
| Go | `go test -coverprofile` | ✅ already done |
| TypeScript | `npx jest --json` or `npx vitest run --reporter=json` | Parse JSON for pass/fail + coverage |
| Python | `python -m pytest --json-report` or parse `coverage.py` | `coverage json` |
| Rust | `cargo test -- --format=json` (nightly) or parse `cargo test` output | `cargo tarpaulin --out json` |

Graceful: if tool not installed, test evidence is absent. Doctor reports it.

---

### Phase 2: Deep Go Analysis (Tier 2 — Go)
**Goal:** Type-aware, cross-file analysis for Go using `go/types` + `go/packages`. No external deps.

#### 2A: `internal/analysis/go_deep.go` — type-aware analyzer

```go
// DeepGoAnalyzer performs type-aware analysis using go/packages.
type DeepGoAnalyzer struct {
    pkgs []*packages.Package
    fset *token.FileSet
}

// Load initializes the analyzer by loading all packages in the project.
func LoadGoProject(root string, patterns ...string) (*DeepGoAnalyzer, error) {
    cfg := &packages.Config{
        Mode: packages.NeedName | packages.NeedFiles | packages.NeedSyntax |
              packages.NeedTypes | packages.NeedTypesInfo | packages.NeedDeps |
              packages.NeedImports,
        Dir: root,
    }
    pkgs, err := packages.Load(cfg, patterns...)
    // ...
}
```

#### 2B: Call graph → fan-in / fan-out

Using `golang.org/x/tools/go/callgraph` with the VTA (Variable Type Analysis) algorithm:

```go
func (a *DeepGoAnalyzer) CallGraph() *callgraph.Graph {
    // VTA is the best balance of precision and speed
    prog, ssaPkgs := ssautil.AllPackages(a.pkgs, ssa.InstantiateGenerics)
    prog.Build()
    return vta.CallGraph(ssaPkgs, cha.CallGraph(prog))
}

// FanIn returns the number of call sites that invoke this function.
func (a *DeepGoAnalyzer) FanIn(funcName string) int { ... }

// FanOut returns the number of distinct functions called by this function.
func (a *DeepGoAnalyzer) FanOut(funcName string) int { ... }
```

New metrics and scoring:

| Metric | Dimension | Thresholds |
|--------|-----------|-----------|
| `fan_in` | change_risk | ≤5 → 0.95, ≤10 → 0.85, ≤20 → 0.70, >20 → 0.50 |
| `fan_out` | maintainability | ≤5 → 0.95, ≤10 → 0.85, ≤15 → 0.70, >15 → 0.55 |
| `coupling_score` | architectural_fitness | fan_in × fan_out normalized |

#### 2C: Dead code detection

```go
// UnusedExports returns exported symbols with zero external references.
func (a *DeepGoAnalyzer) UnusedExports() []UnusedSymbol {
    for each exported symbol in each package:
        refs := a.References(symbol)
        externalRefs := filter(refs, not in same package)
        if len(externalRefs) == 0 && not main package:
            unused = append(unused, symbol)
}
```

| Metric | Dimension | Scoring |
|--------|-----------|---------|
| `is_dead_code` | maintainability | 0 → neutral, 1 → setMin(0.60) |
| `dead_export_count` (package) | architectural_fitness | Package-level aggregate |

#### 2D: Interface compliance / dependency inversion

```go
// InterfaceSatisfaction checks whether function params use interfaces or concrete types.
func (a *DeepGoAnalyzer) ParamAbstraction(funcObj *types.Func) AbstractionScore {
    sig := funcObj.Type().(*types.Signature)
    for each param:
        if param type is interface → good
        if param type is concrete struct from external package → violation
}
```

| Metric | Dimension | What it measures |
|--------|-----------|-----------------|
| `concrete_deps` | testability, arch_fitness | Params that accept concrete types instead of interfaces |
| `interface_size` | arch_fitness | Methods in interfaces this type implements (ISP) |

#### 2E: Package dependency graph

```go
// PackageDeps builds the full import DAG.
func (a *DeepGoAnalyzer) PackageDeps() DepGraph {
    // For each package, record direct imports
    // Detect cycles (should be impossible in Go, but transitive dep depth matters)
}

// DepDepth returns the maximum transitive import depth for a package.
func (a *DeepGoAnalyzer) DepDepth(pkgPath string) int { ... }

// Instability computes Robert C. Martin's instability metric: Ce / (Ca + Ce)
// Ce = efferent coupling (imports), Ca = afferent coupling (imported by)
func (a *DeepGoAnalyzer) Instability(pkgPath string) float64 { ... }
```

| Metric | Dimension | Scoring |
|--------|-----------|---------|
| `dep_depth` | architectural_fitness | ≤3 → 0.95, ≤5 → 0.85, ≤8 → 0.70, >8 → 0.55 |
| `instability` | architectural_fitness | Package-level. Unstable concrete = bad, unstable abstract = ok |
| `has_dep_cycle` | architectural_fitness | (Not possible in Go but tracked for other languages) |

#### 2F: Error wrapping analysis

```go
// ErrorWrapping checks if errors are properly wrapped with context.
func (a *DeepGoAnalyzer) ErrorWrapping(fn *ast.FuncDecl) ErrorWrappingResult {
    // Find: fmt.Errorf("...", err) without %w
    // Find: functions that receive error and return error but never wrap
    // Find: errors.New() in functions that receive an error param
}
```

| Metric | Dimension | Scoring |
|--------|-----------|---------|
| `errors_not_wrapped` | operational_quality | 0 → 0.90, 1-2 → 0.75, 3+ → 0.55 |
| `error_context_ratio` | operational_quality | % of error returns that add context |

#### 2G: Cognitive complexity (Go)

Sonar's cognitive complexity algorithm, distinct from cyclomatic:
- +1 for each `if`, `else if`, `else`, `for`, `switch`, `select`
- +1 extra for each nesting level (cognitive penalty for nested logic)
- +1 for `break`/`continue` to a label, `goto`
- +1 for each boolean sequence (`&&`/`||` alternation)
- NOT +1 for `case` clauses (they're linear, not branching)

```go
func CognitiveComplexity(fn *ast.FuncDecl) int { ... }
```

| Metric | Dimension | Thresholds |
|--------|-----------|-----------|
| `cognitive_complexity` | readability | ≤8 → 0.95, ≤15 → 0.85, ≤25 → 0.70, >25 → 0.50 |

This is a stronger readability signal than cyclomatic complexity because it penalizes the patterns humans actually find hard to read.

---

### Phase 3: LSP Integration for Non-Go (Tier 2 — Multi-Language)
**Goal:** Type-aware analysis for TS, Python, Rust via their language servers.

#### 3A: `internal/analysis/lsp/` — generic LSP client

```go
// Client manages a language server subprocess lifecycle.
type Client struct {
    cmd     *exec.Cmd
    stdin   io.WriteCloser
    stdout  *bufio.Reader
    nextID  int64
}

// Start spawns the language server process.
func Start(command string, args []string, rootDir string) (*Client, error) { ... }

// Initialize sends the LSP initialize request.
func (c *Client) Initialize(rootURI string) error { ... }

// DocumentSymbols returns all symbols in a file.
func (c *Client) DocumentSymbols(uri string) ([]Symbol, error) { ... }

// References finds all references to a symbol at a given position.
func (c *Client) References(uri string, line, col int) ([]Location, error) { ... }

// CallHierarchyIncoming returns incoming calls to a symbol.
func (c *Client) CallHierarchyIncoming(uri string, line, col int) ([]CallHierarchyItem, error) { ... }

// CallHierarchyOutgoing returns outgoing calls from a symbol.
func (c *Client) CallHierarchyOutgoing(uri string, line, col int) ([]CallHierarchyItem, error) { ... }

// Diagnostics returns all diagnostics for a file.
func (c *Client) Diagnostics(uri string) ([]Diagnostic, error) { ... }

// Shutdown gracefully stops the server.
func (c *Client) Shutdown() error { ... }
```

This is ~300-400 lines. The LSP JSON-RPC protocol is well-specified and the same for all languages.

#### 3B: Language server detection and configuration

```yaml
# .certification/config.yml
analysis:
  lsp:
    typescript:
      command: "npx"
      args: ["typescript-language-server", "--stdio"]
      # or auto-detect from node_modules
    python:
      command: "pyright-langserver"
      args: ["--stdio"]
      # or "pylsp", or "ruff server"
    rust:
      command: "rust-analyzer"
      args: []
```

Auto-detection in `certify doctor`:
```
── Language Servers ──
  ✅ TypeScript: typescript-language-server found (via npx)
  ✅ Python: pyright 1.2.3
  ⚠️ Rust: rust-analyzer not found
     → Install: rustup component add rust-analyzer
```

#### 3C: LSP-powered evidence collection

For each language with an available LSP:

1. **Start server** (once per `certify certify` run)
2. **Open all files** in the project
3. **Collect diagnostics** → correctness evidence (type errors, unused imports)
4. **For each unit:**
   - `callHierarchy/incomingCalls` → fan-in
   - `callHierarchy/outgoingCalls` → fan-out
   - `textDocument/references` → reference count (dead code = 0)
   - `textDocument/implementation` → interface compliance
5. **Shutdown server**

Cost: ~5-30 seconds startup, then fast per-query. Acceptable for a certification run.

#### 3D: Graceful degradation

```
Evidence collection priority:
  1. Always: Tier 0 (git, line counts)
  2. Always: Tier 1 (tree-sitter structural — bundled, no external deps)
  3. If available: Tier 1 lint tools (eslint, ruff, clippy)
  4. If available: Tier 1 test runners (jest, pytest, cargo test)
  5. If available: Tier 2 LSP (fan-in/out, dead code, interface compliance)
```

Each missing tier just means fewer evidence items — scoring adapts because dimensions without evidence don't dilute the average (existing behavior).

`certify doctor` reports what's available at each tier:
```
── Analysis Capabilities ──
  ✅ Go: Tier 2 (go/types — full type analysis)
  ✅ TypeScript: Tier 2 (tsserver found)
  ⚠️ Python: Tier 1 only (pyright not found)
     → Install: pip install pyright
  ⚠️ Rust: Tier 1 only (rust-analyzer not found)
     → Install: rustup component add rust-analyzer
```

---

### Phase 4: Scorer Enhancements
**Goal:** New metrics flow into dimension scores with well-calibrated thresholds.

#### 4A: New scoring rules

```go
func scoreFromDeepAnalysis(e domain.Evidence, scores domain.DimensionScores) {
    m := e.Metrics

    // Fan-in → change_risk
    if fanIn, ok := m["fan_in"]; ok {
        switch {
        case fanIn <= 5:  setMax(scores, DimChangeRisk, 0.95)
        case fanIn <= 10: setMax(scores, DimChangeRisk, 0.85)
        case fanIn <= 20: setMax(scores, DimChangeRisk, 0.70)
        default:          setMin(scores, DimChangeRisk, 0.50)
        }
    }

    // Fan-out → maintainability
    if fanOut, ok := m["fan_out"]; ok {
        switch {
        case fanOut <= 5:  setMax(scores, DimMaintainability, 0.95)
        case fanOut <= 10: setMax(scores, DimMaintainability, 0.85)
        case fanOut <= 15: setMax(scores, DimMaintainability, 0.70)
        default:           setMin(scores, DimMaintainability, 0.55)
        }
    }

    // Dead code → maintainability
    if m["is_dead_code"] > 0 {
        setMin(scores, DimMaintainability, 0.60)
    }

    // Concrete deps → testability + arch_fitness
    if concreteDeps, ok := m["concrete_deps"]; ok && concreteDeps > 0 {
        setMin(scores, DimTestability, 0.65)
        setMin(scores, DimArchitecturalFitness, 0.65)
    }

    // Cognitive complexity → readability
    if cogCplx, ok := m["cognitive_complexity"]; ok {
        switch {
        case cogCplx <= 8:  setMax(scores, DimReadability, 0.95)
        case cogCplx <= 15: setMax(scores, DimReadability, 0.85)
        case cogCplx <= 25: setMax(scores, DimReadability, 0.70)
        default:            setMin(scores, DimReadability, 0.50)
        }
    }

    // Error wrapping → operational_quality
    if unwrapped, ok := m["errors_not_wrapped"]; ok {
        switch {
        case unwrapped == 0: setMax(scores, DimOperationalQuality, 0.90)
        case unwrapped <= 2: setMax(scores, DimOperationalQuality, 0.75)
        default:             setMin(scores, DimOperationalQuality, 0.55)
        }
    }

    // Unsafe imports → security
    if unsafeCount, ok := m["unsafe_import_count"]; ok && unsafeCount > 0 {
        setMin(scores, DimSecurity, 0.60)
    }

    // Hardcoded secrets → security
    if secrets, ok := m["hardcoded_secrets"]; ok && secrets > 0 {
        setMin(scores, DimSecurity, 0.30) // Critical
    }

    // Dep depth → architectural_fitness
    if depth, ok := m["dep_depth"]; ok {
        switch {
        case depth <= 3: setMax(scores, DimArchitecturalFitness, 0.95)
        case depth <= 5: setMax(scores, DimArchitecturalFitness, 0.85)
        case depth <= 8: setMax(scores, DimArchitecturalFitness, 0.70)
        default:         setMin(scores, DimArchitecturalFitness, 0.55)
        }
    }

    // Instability for concrete packages → architectural_fitness
    if instability, ok := m["instability"]; ok && instability > 0.8 {
        // High instability is fine for abstractions, bad for concrete impls
        if m["is_abstract"] == 0 {
            setMin(scores, DimArchitecturalFitness, 0.65)
        }
    }
}
```

#### 4B: New evidence kind

```go
const (
    EvidenceKindDeepAnalysis EvidenceKind = iota + 100 // Type-aware cross-file analysis
)
```

#### 4C: New policy rules (added to go-standard.yml)

```yaml
  - id: max-fan-out
    dimension: maintainability
    description: "Functions should not call more than 15 distinct functions"
    severity: warning
    metric: fan_out
    threshold: 15

  - id: max-fan-in
    dimension: change_risk
    description: "Functions called by more than 20 callers are high-risk change points"
    severity: warning
    metric: fan_in
    threshold: 20

  - id: no-dead-exports
    dimension: maintainability
    description: "Exported symbols should have at least one external reference"
    severity: info
    metric: is_dead_code
    threshold: 0

  - id: max-cognitive-complexity
    dimension: readability
    description: "Cognitive complexity should not exceed 25"
    severity: warning
    metric: cognitive_complexity
    threshold: 25

  - id: no-unsafe-imports
    dimension: security
    description: "Avoid unsafe, os/exec, eval, and similar dangerous imports"
    severity: warning
    metric: unsafe_import_count
    threshold: 0

  - id: no-hardcoded-secrets
    dimension: security
    description: "No hardcoded passwords, API keys, or secrets"
    severity: critical
    metric: hardcoded_secrets
    threshold: 0

  - id: wrap-errors
    dimension: operational_quality
    description: "Errors should be wrapped with context using %w or equivalent"
    severity: warning
    metric: errors_not_wrapped
    threshold: 0

  - id: max-dep-depth
    dimension: architectural_fitness
    description: "Package import depth should not exceed 8 levels"
    severity: warning
    metric: dep_depth
    threshold: 8
```

---

### Phase 5: Architect Review Integration
**Goal:** Feed all new metrics into the architect snapshot for grounded LLM analysis.

#### 5A: Expand `ArchSnapshot` aggregates

Add to the existing `StructuralAggregates`:

```go
type DeepAnalysisAggregates struct {
    AvgFanIn            float64
    MaxFanIn            int
    AvgFanOut           float64
    MaxFanOut           int
    DeadExportCount     int
    ConcreteDepsCount   int
    AvgCogComplexity    float64
    MaxCogComplexity    int
    ErrorsNotWrapped    int
    UnsafeImportCount   int
    HardcodedSecrets    int
    MaxDepDepth         int
    AvgInstability      float64
    PackagesWithCycles  int
}
```

#### 5B: Update phase prompts

Phase 2 (package analysis) prompt now includes:
```
Package Coupling:
| Package | Fan-In | Fan-Out | Instability | Dep Depth | Dead Exports |
...
```

Phase 4 (structural) prompt now includes:
```
Cognitive Complexity Hotspots:
| Function | Cyclomatic | Cognitive | Fan-In | Fan-Out |
...
```

Phase 5 (security) prompt now includes:
```
Security Findings:
| Unit | Finding | Risk |
| ... | unsafe import: os/exec | high |
| ... | hardcoded secret pattern | critical |
...
```

All grounded — every number in the prompt has an exact source in the snapshot.

---

### Phase 6: Documentation and Doctor Updates
**Goal:** Users understand what analysis is available and how to unlock more.

#### 6A: `certify doctor` — analysis tier reporting

```
── Analysis Tiers ──
  ✅ Go: Tier 2 (go/types — call graph, dead code, interface compliance)
     → 38 metrics per unit across all 9 dimensions
  ⚠️ TypeScript: Tier 1 (tree-sitter syntax analysis)
     → 22 metrics per unit. Install typescript-language-server for Tier 2.
  ❌ Python: Tier 0 (file-level only)
     → Install tree-sitter for Tier 1, pyright for Tier 2.

── Lint Tools ──
  ✅ golangci-lint 2.7.2
  ✅ eslint 9.x (via npx)
  ⚠️ ruff: not found → pip install ruff
  ⚠️ cargo clippy: not found → rustup component add clippy

── Test Runners ──
  ✅ go test (with coverage)
  ✅ jest (detected in package.json)
  ⚠️ pytest: not found → pip install pytest
```

#### 6B: Website docs

- `advanced/analysis-tiers.md` — explains the tier system, what each provides, how to upgrade
- Update `concepts/dimensions.md` — show all metrics feeding each dimension
- Update `reference/cli.md` — document `--deep` / `--tier` flags if added

---

## Execution Sequence

| Sprint | Phase | Deliverable | Status | Lines (actual/est) |
|--------|-------|-------------|--------|-------------------|
| **1** | 1A-1B+1D+4A | Analyzer interface + Go adapter + new metrics + scorer + policy v1.2 | ✅ Done | 1,130 actual |
| **2** | 1C | Tree-sitter TypeScript analyzer | ✅ Done | 963 actual |
| **3** | 1C | Tree-sitter Python + Rust analyzers | ✅ Done | 1,437 actual |
| **4** | 1D+4A-C | Scorer enhancements + all language policy packs v1.2-1.3 | ✅ Done | 592 actual |
| **5** | 1E-1F | Wire into pipeline, replace TS regex discovery, add Py/Rs detection | ✅ Done | 152 actual |
| **—** | (bug fix) | Nested Go module support (`DiscoverModuleRoots`) | ✅ Done | 457 actual |
| **6** | 1G-1H | Lint/test tool integration (ESLint, ruff, pytest, cargo clippy/test) | ✅ Done | 1,040 actual |
| **7** | 2A-2C | `go/packages` loader + call graph + dead code | ✅ Done | 540 actual |
| **8** | 2D-2E | Interface compliance + package dep graph | 🔲 Planned | ~600 est |
| **9** | 2F-2G | Error wrapping + cognitive complexity (type-aware refinement) | 🔲 Planned | ~400 est |
| **10** | 3A | LSP client infrastructure | 🔲 Planned | ~400 est |
| **11** | 3B-3C | LSP-powered TS/Py/Rs analysis | 🔲 Planned | ~500 est |
| **12** | 5A-5B | Architect snapshot + prompt updates for new metrics | 🔲 Planned | ~300 est |
| **13** | 6A-6B | Doctor analysis tier reporting + website docs | 🔲 Planned | ~500 est |

**Done:** ~5,771 lines (Sprints 1–6 + nested module fix)
**Remaining:** ~3,500 lines (Sprints 7–13)
**Total:** ~9,271 lines of production code + tests

### Notes on execution vs plan
- Sprint 4 absorbed the spec's Sprint 12 (scorer + policy rules) — they were natural to do together
- Sprint numbering shifted: original Sprint 12 is eliminated, Sprints 13/14 became 12/13
- Nested module fix was unplanned but critical — it was the #1 user-reported issue
- Sprint 1 was larger than estimated: included scorer integration + new metrics (not just parity)
- Tree-sitter LOC ran ~50% higher than estimated due to comprehensive test coverage

---

## Dimension Coverage Progression

### Before (v0.9.0) → After Sprints 1-6 → After Sprints 7-13 (projected)

| Dimension | Go: before → now → final | TS: before → now → final | Py/Rs: before → now → final |
|-----------|--------------------------|--------------------------|---------------------------|
| correctness | 6 → 9 → 11 | 1 → 7 → 9 | 1 → 7 → 8 |
| maintainability | 5 → 8 → 11 | 1 → 6 → 9 | 1 → 6 → 7 |
| readability | 6 → 9 → 9 | 1 → 7 → 7 | 1 → 7 → 7 |
| testability | 4 → 6 → 8 | 0 → 4 → 6 | 0 → 4 → 5 |
| security | 1 → 4 → 4 | 0 → 3 → 3 | 0 → 3 → 3 |
| arch_fitness | 2 → 2 → 7 | 0 → 1 → 5 | 0 → 1 → 3 |
| operational_quality | 1 → 3 → 3 | 1 → 2 → 3 | 1 → 2 → 2 |
| performance | 3 → 4 → 4 | 0 → 3 → 3 | 0 → 3 → 3 |
| change_risk | 2 → 3 → 5 | 1 → 2 → 4 | 1 → 2 → 3 |
| **Total signals** | **30 → 48 → 62** | **4 → 35 → 49** | **4 → 35 → 41** |

### Key transitions achieved (Sprints 1-6)
- **Go:** 30 → 48 signals (+5 new metrics, +lint/test nesting, +cognitive complexity)
- **TypeScript:** 4 → 35 signals (from regex exports-only to full tree-sitter AST + ESLint + Jest)
- **Python:** 4 → 35 signals (from file-level to symbol-level + ruff + pytest)
- **Rust:** 4 → 35 signals (from file-level to symbol-level + cargo clippy + cargo test)

### Key transitions remaining (Sprints 7-13)
- **Go:** 48 → 62 signals (fan-in/fan-out, dead code, dep graph, interface compliance)
- **TypeScript:** 35 → 49 signals (LSP-powered call hierarchy, dead code, diagnostics)
- **Python/Rust:** 35 → 41 signals (LSP-powered analysis where servers available)

---

## Risk Mitigation

| Risk | Mitigation |
|------|-----------|
| Tree-sitter CGo complicates cross-compilation | Provide `CGO_ENABLED=0` build without tree-sitter (falls back to regex adapters). Binary downloads for common platforms. |
| LSP servers are slow to start | Start once per certification run, reuse for all files. Cache results in `.certification/cache/`. |
| LSP servers may not be installed | Graceful degradation — Tier 2 evidence simply absent. Doctor reports what's missing. |
| `go/packages` is slow on large repos | Cache package load. Only re-analyze changed packages (use git diff). |
| Scoring thresholds may need calibration | All thresholds configurable in policy YAML. Ship conservative defaults, tune with real-world data. |
| New metrics change existing grades | Version the scoring algorithm. Document that v0.11+ includes deeper analysis. First run may lower scores — this is accurate, not a regression. |

---

## Dependencies

| Package | Purpose | Type |
|---------|---------|------|
| `golang.org/x/tools/go/packages` | Load Go packages with types | Go module (no CGo) |
| `golang.org/x/tools/go/callgraph` | Call graph algorithms | Go module (no CGo) |
| `golang.org/x/tools/go/ssa` | SSA form for call graph | Go module (no CGo) |
| `github.com/smacker/go-tree-sitter` | Tree-sitter Go bindings | CGo |
| `github.com/smacker/go-tree-sitter/typescript` | TS grammar | CGo |
| `github.com/smacker/go-tree-sitter/python` | Python grammar | CGo |
| `github.com/smacker/go-tree-sitter/rust` | Rust grammar | CGo |

Phase 2 (deep Go) requires only `golang.org/x/tools` — **no CGo**.
Phase 1 tree-sitter requires CGo but is optional via build tags.
Phase 3 LSP requires no new Go deps — just `os/exec` + `encoding/json`.

---

## Success Criteria

### Achieved ✅
1. ✅ `go test ./...` passes at every phase boundary (18 packages, 0 failures throughout)
2. ✅ Go analysis produces identical scores to pre-Sprint 1 (regression_test.go verified)
3. ✅ TypeScript units receive scores on ≥7 dimensions (was ≤3, now 7-8)
4. ✅ Python/Rust units receive symbol-level analysis (was file-level only)
5. ✅ No CGo required for core functionality — tree-sitter is opt-in
6. ✅ Nested module support — go.mod/package.json/Cargo.toml in subdirectories
7. ✅ `certify doctor` reports lint/test tool availability for all languages
8. ✅ Policy packs for all 4 languages with 11-19 rules each

### Remaining (Sprints 7-13)
9. `certify doctor` reports analysis tier per language (T0/T1/T2)
10. Go call graph provides fan-in/fan-out for every function
11. Dead export detection flags unused public API
12. LSP client provides type-aware analysis for non-Go languages
13. Architect review prompts cite all new metrics from snapshot (zero hallucination)
14. Website documents analysis tier system and dimension-to-metric mapping

---

## Remaining Sprint Plans (Detailed)

### Sprint 7 — Deep Go: Call Graph + Dead Code

**Goal:** Type-aware, cross-file Go analysis using `go/types` + `go/packages`. Fan-in, fan-out, and dead export detection.

**Dependency:** `golang.org/x/tools/go/packages`, `golang.org/x/tools/go/callgraph/vta`, `golang.org/x/tools/go/ssa`

**New file:** `internal/analysis/go_deep.go` (~500 lines)

**Implementation:**

```go
// DeepGoAnalyzer provides type-aware cross-file analysis for Go.
type DeepGoAnalyzer struct {
    pkgs  []*packages.Package
    fset  *token.FileSet
    graph *callgraph.Graph
}

// LoadGoProject initializes the analyzer by loading all packages.
func LoadGoProject(root string, patterns ...string) (*DeepGoAnalyzer, error)

// FanIn returns the number of call sites that invoke this function.
func (a *DeepGoAnalyzer) FanIn(pkg, funcName string) int

// FanOut returns the number of distinct functions called by this function.
func (a *DeepGoAnalyzer) FanOut(pkg, funcName string) int

// UnusedExports returns exported symbols with zero external references.
func (a *DeepGoAnalyzer) UnusedExports() []UnusedSymbol
```

**Integration point:** Called once per `certify certify` run (project-wide). Results cached in a map keyed by `UnitID`. The certifier looks up fan-in/fan-out per unit during `collectStructuralEvidence()` and appends metrics to the evidence.

**New metrics:**

| Metric | Key | Dimension | Scoring |
|--------|-----|-----------|---------|
| Fan-in | `fan_in` | change_risk | ≤5 → 0.95, ≤10 → 0.85, ≤20 → 0.70, >20 → 0.50 |
| Fan-out | `fan_out` | maintainability | ≤5 → 0.95, ≤10 → 0.85, ≤15 → 0.70, >15 → 0.55 |
| Dead code | `is_dead_code` | maintainability | 0 → neutral, 1 → cap at 0.60 |

**New policy rules (go-standard v1.4.0):**
```yaml
  - id: max-fan-out
    dimension: maintainability
    metric: fan_out
    threshold: 15
    severity: warning

  - id: max-fan-in
    dimension: change_risk
    metric: fan_in
    threshold: 20
    severity: warning

  - id: no-dead-exports
    dimension: maintainability
    metric: is_dead_code
    threshold: 0
    severity: info
```

**Tests (~300 lines):** `internal/analysis/go_deep_test.go`
- Test with multi-file Go project in testdata/
- Verify fan-in/fan-out counts against known call graph
- Verify dead export detection
- Verify graceful handling of build errors
- Verify caching (second call returns same results)

**Acceptance criteria:**
1. `go test ./internal/analysis/` passes with call graph tests
2. Running `certify certify` on this repo produces `fan_in`/`fan_out` metrics for Go units
3. Dead exports detected and flagged in evidence
4. Full test suite passes (18 packages)
5. Performance: project load + call graph < 30s for this repo (~900 units)

---

### Sprint 8 — Deep Go: Interface Compliance + Package Dep Graph

**Goal:** Dependency inversion detection and package coupling metrics.

**New file:** `internal/analysis/go_deps.go` (~400 lines)

**Implementation:**

```go
// ParamAbstraction checks whether function params use interfaces or concrete types.
func (a *DeepGoAnalyzer) ParamAbstraction(pkg, funcName string) AbstractionResult

// DepDepth returns the maximum transitive import depth for a package.
func (a *DeepGoAnalyzer) DepDepth(pkgPath string) int

// Instability computes Robert C. Martin's instability: Ce / (Ca + Ce)
func (a *DeepGoAnalyzer) Instability(pkgPath string) float64

// DetectCycles checks for import cycles (should be impossible in Go).
func (a *DeepGoAnalyzer) DetectCycles() [][]string
```

**New metrics:**

| Metric | Key | Dimension | Scoring |
|--------|-----|-----------|---------|
| Concrete deps | `concrete_deps` | testability, arch_fitness | >0 → cap at 0.65 |
| Dep depth | `dep_depth` | arch_fitness | ≤3 → 0.95, ≤5 → 0.85, ≤8 → 0.70, >8 → 0.55 |
| Instability | `instability` | arch_fitness | >0.8 for concrete → cap at 0.65 |
| Coupling score | `coupling_score` | arch_fitness | fan_in × fan_out normalized |

**Tests (~200 lines):** `internal/analysis/go_deps_test.go`
- Verify dep depth calculation with known import chains
- Verify instability formula against hand-calculated values
- Verify concrete dep detection for function params

**Acceptance criteria:**
1. Package dep graph built from `go/packages` import data
2. `dep_depth` and `instability` computed per package, attributed to units
3. Interface compliance check identifies concrete struct params
4. New policy rules added for `max-dep-depth`, `no-dead-exports`
5. Full test suite passes

---

### Sprint 9 — Type-Aware Refinement

**Goal:** Improve accuracy of existing metrics using type information.

**Changes to:** `internal/analysis/go_deep.go` (~200 lines additions)

**Improvements:**
1. **Error wrapping (type-aware):** Currently AST-only (checks for `fmt.Errorf` without `%w`). With `go/types`, can verify the value is actually of type `error` — reduces false positives when functions have non-error string formatting.
2. **Cognitive complexity (type-aware):** Type switches on `interface{}` are more complex than on concrete types. Add nesting increment for type assertion cascades.
3. **Method set analysis:** For types implementing large interfaces (>10 methods), flag potential ISP violations.
4. **Unused parameter detection:** With type info, can detect params that are never used in the function body (beyond `_` convention).

**New metrics:**

| Metric | Key | Dimension | Notes |
|--------|-----|-----------|-------|
| Unused params | `unused_params` | maintainability | Params received but never referenced |
| Interface size | `interface_size` | arch_fitness | Methods in interfaces this type implements |

**Tests (~100 lines):**
- Error wrapping: verify only actual `error` values are flagged
- Unused params: verify detection in multi-file context
- Interface size: verify against known interface implementations

**Acceptance criteria:**
1. Error wrapping false positive rate reduced (test with known edge cases)
2. Unused param detection works across files
3. Existing test suite still passes (no regressions)
4. Scoring changes documented as accuracy improvements

---

### Sprint 10 — LSP Client Infrastructure

**Goal:** Generic JSON-RPC 2.0 client for communicating with language servers via stdin/stdout.

**New package:** `internal/analysis/lsp/` (~400 lines)

**Implementation:**

```go
// internal/analysis/lsp/client.go

// Client manages a language server subprocess lifecycle.
type Client struct {
    cmd     *exec.Cmd
    stdin   io.WriteCloser
    stdout  *bufio.Reader
    nextID  atomic.Int64
    pending map[int64]chan json.RawMessage
    mu      sync.Mutex
}

// Start spawns a language server subprocess.
func Start(command string, args []string, rootDir string) (*Client, error)

// Initialize sends the LSP initialize request with capabilities.
func (c *Client) Initialize(rootURI string, capabilities ClientCapabilities) (*InitializeResult, error)

// Request sends a JSON-RPC request and waits for the response.
func (c *Client) Request(method string, params interface{}) (json.RawMessage, error)

// Notify sends a JSON-RPC notification (no response expected).
func (c *Client) Notify(method string, params interface{}) error

// Shutdown gracefully stops the server.
func (c *Client) Shutdown() error

// --- High-level LSP methods ---

// DocumentSymbols returns all symbols in a file.
func (c *Client) DocumentSymbols(uri string) ([]DocumentSymbol, error)

// References finds all references to a symbol at position.
func (c *Client) References(uri string, line, col int) ([]Location, error)

// CallHierarchyIncoming returns incoming calls to symbol at position.
func (c *Client) CallHierarchyIncoming(uri string, line, col int) ([]CallHierarchyIncomingCall, error)

// CallHierarchyOutgoing returns outgoing calls from symbol at position.
func (c *Client) CallHierarchyOutgoing(uri string, line, col int) ([]CallHierarchyOutgoingCall, error)

// Diagnostics returns all diagnostics for a file.
func (c *Client) Diagnostics(uri string) ([]Diagnostic, error)
```

**Files:**
- `internal/analysis/lsp/client.go` — JSON-RPC transport + lifecycle
- `internal/analysis/lsp/types.go` — LSP type definitions (subset: what we need)
- `internal/analysis/lsp/client_test.go` — unit tests with mock server

**Key design decisions:**
- One client per language per `certify certify` run (start once, query many)
- Content-Length framing per LSP spec (`Content-Length: N\r\n\r\n{...}`)
- Async read loop dispatches responses to pending requests by ID
- 30-second timeout per request (configurable)
- Server detection: `certify doctor` checks for `typescript-language-server`, `pyright`, `rust-analyzer`

**Tests (~150 lines):**
- JSON-RPC message framing (encode/decode)
- Request/response matching by ID
- Timeout handling
- Graceful shutdown sequence (shutdown → exit)
- Mock server for integration test (Go subprocess that speaks LSP)

**Acceptance criteria:**
1. Client can start/initialize/shutdown a language server
2. Request/response cycle works with proper content-length framing
3. Concurrent requests are correctly dispatched
4. Timeout fires for unresponsive servers
5. No goroutine leaks on shutdown

---

### Sprint 11 — LSP-Powered Analysis for TS/Py/Rs

**Goal:** Use language servers to get fan-in/fan-out and dead code for TypeScript, Python, and Rust.

**New files:**
- `internal/analysis/lsp_analyzer.go` (~300 lines) — generic LSP-backed analyzer
- `internal/analysis/lsp_config.go` (~100 lines) — server detection + config

**Implementation:**

```go
// LSPAnalyzer wraps an LSP client to provide Tier 2 analysis.
type LSPAnalyzer struct {
    client   *lsp.Client
    language string
    rootURI  string
}

// NewLSPAnalyzer starts the appropriate language server and returns an analyzer.
func NewLSPAnalyzer(lang, rootDir string) (*LSPAnalyzer, error)

// FanIn returns incoming call count for a symbol.
func (a *LSPAnalyzer) FanIn(file string, line, col int) (int, error)

// FanOut returns outgoing call count from a symbol.
func (a *LSPAnalyzer) FanOut(file string, line, col int) (int, error)

// IsDeadCode returns true if symbol has zero references outside its file.
func (a *LSPAnalyzer) IsDeadCode(file string, line, col int) (bool, error)

// Diagnostics returns type-checker diagnostics for a file.
func (a *LSPAnalyzer) Diagnostics(file string) ([]lsp.Diagnostic, error)
```

**Language server detection:**

| Language | Server | Detection | Install hint |
|----------|--------|-----------|-------------|
| TypeScript | `typescript-language-server` | `which typescript-language-server` or `npx` | `npm i -g typescript-language-server` |
| Python | `pyright` or `pylsp` | `which pyright-langserver` or `which pylsp` | `pip install pyright` |
| Rust | `rust-analyzer` | `which rust-analyzer` | `rustup component add rust-analyzer` |

**Configuration (optional in .certification/config.yml):**
```yaml
analysis:
  lsp:
    typescript:
      command: "npx"
      args: ["typescript-language-server", "--stdio"]
    python:
      command: "pyright-langserver"
      args: ["--stdio"]
    rust:
      command: "rust-analyzer"
```

**Integration point:** During `certify certify`, if an LSP server is available:
1. Start server once (first unit of that language)
2. For each unit: query callHierarchy + references
3. Append `fan_in`, `fan_out`, `is_dead_code` to evidence
4. Shutdown server after all units processed

**Graceful degradation:** If no LSP server found, Tier 2 metrics are absent. Scoring adapts — dimensions without evidence don't dilute the average (existing behavior).

**New metrics (for TS/Py/Rs):**

| Metric | Key | Dimension | Same thresholds as Go Sprint 7 |
|--------|-----|-----------|-------------------------------|
| Fan-in | `fan_in` | change_risk | ≤5/≤10/≤20/>20 |
| Fan-out | `fan_out` | maintainability | ≤5/≤10/≤15/>15 |
| Dead code | `is_dead_code` | maintainability | cap at 0.60 |
| Type diagnostics | `type_errors` | correctness | >0 → cap at 0.50 |

**Tests (~150 lines):**
- LSPAnalyzer with mock server returning canned callHierarchy responses
- FanIn/FanOut counting from callHierarchy items
- Dead code from empty references response
- Graceful handling when server unavailable

**Acceptance criteria:**
1. TypeScript fan-in/fan-out works with `typescript-language-server`
2. Python analysis works with `pyright`
3. Graceful degradation when server not installed
4. `certify doctor` reports LSP availability per language
5. Performance: < 60s for a 100-file TypeScript project

---

### Sprint 12 — Architect Snapshot + Prompt Updates

**Goal:** Feed all new Tier 2 metrics into the architect snapshot so LLM analysis is fully grounded.

**Changes to:** `internal/agent/architect_snapshot.go` (~200 lines additions)

**New aggregates:**

```go
type DeepAnalysisAggregates struct {
    AvgFanIn            float64
    MaxFanIn            int
    AvgFanOut           float64
    MaxFanOut           int
    DeadExportCount     int
    ConcreteDepsCount   int
    AvgCogComplexity    float64
    MaxCogComplexity    int
    ErrorsNotWrapped    int
    UnsafeImportCount   int
    HardcodedSecrets    int
    MaxDepDepth         int
    AvgInstability      float64
}
```

**Prompt updates:**

Phase 2 (package analysis):
```
Package Coupling:
| Package | Fan-In | Fan-Out | Instability | Dep Depth | Dead Exports |
...
```

Phase 4 (structural hotspots):
```
Complexity Hotspots:
| Function | Cyclomatic | Cognitive | Fan-In | Fan-Out |
...
Top units with fan-in > 10 are high-risk change points.
```

Phase 5 (security):
```
Security Findings:
| Unit | Finding | Severity |
| pkg/exec.go#Run | unsafe import: os/exec | warning |
| config/load.go#Parse | hardcoded secret pattern | critical |
...
```

**Schema version:** Bump to v3 (v2 was snapshot alignment, v3 adds deep analysis).

**Tests (~100 lines):**
- Verify DeepAnalysisAggregates populated from mock records
- Verify prompt formatting includes new tables
- Verify schema version in header and appendix

**Acceptance criteria:**
1. `DeepAnalysisAggregates` computed during `BuildArchSnapshot()`
2. All 6 phase prompts reference deep analysis metrics where relevant
3. Schema version bumped and documented
4. All prompt data has exact source in snapshot (zero hallucination)

---

### Sprint 13 — Doctor Analysis Tiers + Website Documentation

**Goal:** Users understand what analysis is available, how to unlock more, and all features are documented.

**Changes to:** `internal/doctor/doctor.go` (~150 lines additions)

**Doctor output — analysis tier reporting:**
```
── Analysis Tiers ──
  ✅ Go: Tier 2 (go/types — call graph, dead code, interface compliance)
     → 62 metrics per unit across 9 dimensions
  ✅ TypeScript: Tier 2 (typescript-language-server found)
     → 49 metrics per unit across 9 dimensions
  ⚠️ Python: Tier 1 (tree-sitter syntax analysis only)
     → 35 metrics per unit. Install pyright for Tier 2:
     → pip install pyright
  ⚠️ Rust: Tier 1 (tree-sitter syntax analysis only)
     → 35 metrics per unit. Install rust-analyzer for Tier 2:
     → rustup component add rust-analyzer
```

**Website pages (~350 lines total):**

1. **`advanced/analysis-tiers.md`** (new, ~200 lines)
   - Explains Tier 0/1/2/3 system
   - What each tier provides per language
   - How to unlock higher tiers (install tools/servers)
   - Metric-to-dimension mapping table
   - Performance expectations per tier

2. **Update `concepts/dimensions.md`** (~50 lines)
   - Add "Metrics" column showing all signals feeding each dimension
   - Note which metrics require which tier

3. **Update `reference/cli.md`** (~50 lines)
   - Document any new flags (e.g., `--tier`, `--deep`)
   - Update `certify doctor` output examples

4. **Update `guides/quickstart.md`** (~50 lines)
   - Note that `certify doctor` shows analysis capabilities
   - Recommend installing lint/test tools for best results

**Tests:**
- Doctor tier detection test (mock with/without LSP servers)
- Website build verification (`npm run build` succeeds)

**Acceptance criteria:**
1. `certify doctor` shows analysis tier per detected language
2. Doctor suggests install commands for missing Tier 2 servers
3. Website builds with new/updated pages
4. All metric-to-dimension mappings documented
5. Release notes drafted for v0.12.0
