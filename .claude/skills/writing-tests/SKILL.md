---
name: writing-tests
description: Write high-quality tests that validate behavior, preserve design integrity, and keep dependencies at the edges. Prefer fakes over mocks, introduce interfaces at seams, and test through stable contracts.
---

## Core Principles

- Test behavior and outcomes, not implementation details.
- Keep tests deterministic, fast, and isolated.
- Prefer **fakes** (in-memory implementations) over mocks.
- Use abstraction to create seams where external dependencies exist.
- Push IO and infrastructure to the edges; keep core logic pure where possible.
- Design for testability: dependencies flow inward, implementations live outward.

---

## Test Design Rules

### 1) Choose the Right Test Level
Use a small number of broad categories:

- **Unit tests**: core logic with all external dependencies replaced by fakes.
- **Contract tests**: verify that an implementation satisfies an interface/contract used by the core.
- **Integration tests**: limited set that exercise real boundaries end-to-end; keep few, run slower.

Default to unit tests unless the goal explicitly requires boundary verification.

---

### 2) Prefer Fakes Over Mocks
**Fakes** are lightweight implementations that behave like the real dependency but run in-memory or locally.

Use fakes to:
- capture writes and reads
- simulate time, IDs, randomness
- simulate external responses
- record calls for assertions (call history as data)

Avoid mocks that:
- assert call order tightly unless behavior requires it
- encode implementation details
- require complex setup to satisfy the test

If verification is necessary, verify outcomes first; verify interactions only when they represent the contract.

---

### 3) Create Seams With Interfaces
Whenever code touches an external dependency, define a boundary interface owned by the core.

Common boundaries:
- File storage
- Network/API clients
- Databases
- Clock/time
- Randomness/UUIDs
- Environment/config
- Process execution
- Message queues/events

Rules:
- The core depends on interfaces (ports), not implementations.
- Implementations (adapters) live at the edge and depend on the core.
- Tests use fakes that implement the same interface.

---

### 4) File Handling Guidance
Do not test core behavior through the real filesystem.

Instead:
- Extract a file boundary interface (examples of responsibilities):
  - read text/binary by path/key
  - write text/binary by path/key
  - list keys
  - exists/delete
- Use an in-memory fake file store in tests.
- Add a small set of boundary/contract tests for the filesystem adapter only if needed.

Assertions should focus on:
- what was written
- what was read/parsed
- error propagation rules
- behavior when files are missing/corrupt

---

### 5) API / Network Calling Guidance
Do not test core behavior by calling real network services.

Instead:
- Extract a boundary interface for the external service.
- Use fakes for normal responses, errors, retries, and timeouts.
- Add a limited contract test suite for the real adapter, verifying:
  - request shape (method/path/headers/body)
  - response parsing
  - error mapping

Keep network adapter tests narrow and explicit.

---

### 6) Arrange–Act–Assert Structure
Every test should clearly separate:
- **Arrange**: build inputs and fakes
- **Act**: call the behavior under test
- **Assert**: verify outcomes

Additional rules:
- One primary reason to fail per test.
- Name tests by behavior: “does X when Y”.

---

### 7) Assertions and Observability
Prefer assertions that are:
- stable
- meaningful
- minimal but complete

Verify:
- returned values / outputs
- state changes in fakes
- emitted domain events (if applicable)
- error types/messages only as part of a defined contract

Avoid brittle assertions:
- exact formatting unless formatting is the behavior
- deep internal state if it’s not part of the contract

---

### 8) Error Handling Tests
For each boundary:
- test the happy path
- test expected failure modes
- test unexpected errors are surfaced or mapped consistently

Define and validate:
- error categories
- mapping rules at boundaries (edge)
- core logic never depends on vendor-specific errors

---

### 9) Data Builders and Fixtures
Use builders to keep tests readable:
- object builders for domain entities
- minimal fixtures
- avoid global shared mutable fixtures

Rules:
- tests should be understandable without scrolling
- fixtures should be local to the test file/module when possible

---

### 10) Keep the Core Pure Where It Pays Off
When feasible:
- keep domain logic free of IO
- make dependencies explicit via parameters or injected interfaces
- keep functions small and composable

If code is currently entangled:
- introduce seams incrementally
- write characterization tests first (capture current behavior)
- refactor behind tests

---

## Execution Workflow (What To Do When Asked to Add Tests)

1) Identify the behavior to validate (inputs, outputs, invariants).
2) Identify dependencies touched (file, api, db, time, etc.).
3) Extract or confirm boundary interfaces owned by the core.
4) Implement fakes for those interfaces in test code.
5) Write unit tests for core behavior using fakes.
6) Add contract tests for adapters only when the adapter is non-trivial or historically risky.
7) Ensure test suite remains fast and deterministic.

---

## Deliverables

When creating tests:
- Provide a short list of test cases (behavior-focused).
- Provide fakes needed and what they record.
- Ensure each test has clear Arrange–Act–Assert.
- Avoid interaction-heavy expectations unless part of the interface contract.

---

## Quality Checklist

- [ ] Tests fail for the right reason
- [ ] Tests do not require network or filesystem for core behavior
- [ ] Fakes capture effects and enable outcome assertions
- [ ] Core depends on interfaces; adapters depend on core
- [ ] Minimal, stable assertions
- [ ] Error paths covered
- [ ] Readable naming and structure
