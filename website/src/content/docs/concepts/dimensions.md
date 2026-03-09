---
title: Quality Dimensions
description: The 9 quality dimensions Certify scores every code unit against.
---

Certify evaluates every code unit across 9 quality dimensions. Each dimension is scored independently, then combined using weighted averaging to produce a final score and grade.

## The Dimensions

### Correctness
**Does the code do what it claims?**

Evidence: lint errors, vet issues, test failures, type errors.

### Maintainability
**How easy is it to modify safely?**

Evidence: cyclomatic complexity, function length, nesting depth.

### Readability
**How clear and understandable is the code?**

Evidence: line length, documentation presence, TODO/FIXME count, naming conventions.

### Testability
**How well is the code tested?**

Evidence: test coverage, test file existence, test-to-code ratio.

### Security
**Are there security concerns?**

Evidence: security-sensitive patterns, dependency vulnerabilities, hardcoded credentials.

### Architectural Fitness
**Does the code fit the system's architecture?**

Evidence: package structure, dependency direction, import patterns.

### Operational Quality
**How stable is the code in practice?**

Evidence: git churn (change frequency), contributor count, file age.

### Performance Appropriateness
**Are there performance concerns?**

Evidence: algorithmic complexity indicators, resource usage patterns.

### Change Risk
**How risky is this code to modify?**

Evidence: recent change frequency, author concentration, coupling metrics.

## Scoring

Each dimension produces a score from 0.0 to 1.0. These are combined using configurable weights:

```
Final Score = Σ (dimension_score × dimension_weight) / Σ weights
```

Default weights give equal importance to all dimensions. Customize in your policy configuration.

## Grades

| Grade | Score | Meaning |
|-------|-------|---------|
| **A** | ≥ 93% | Excellent across all dimensions |
| **A-** | ≥ 90% | Very strong, minor areas for improvement |
| **B+** | ≥ 87% | Strong, above average quality |
| **B** | ≥ 80% | Good, meets expectations |
| **C** | ≥ 70% | Acceptable but has notable issues |
| **D** | ≥ 60% | Below expectations, needs improvement |
| **F** | < 60% | Fails to meet quality standards |
