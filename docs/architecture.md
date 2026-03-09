# Architecture

## Overview

The Code Certification System is a Go CLI that evaluates code units against versioned policies and assigns time-bound certification status.

```
┌─────────────────────────────────────────────────────────────┐
│                     CLI (cmd/certify/)                       │
│  init │ scan │ certify │ report │ expire │ review │ version │
└────┬──────┬──────┬────────┬────────┬────────┬──────────────┘
     │      │      │        │        │        │
┌────▼──────▼──────▼────────▼────────▼────────▼──────────────┐
│                    Internal Packages                        │
│                                                             │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌───────────┐  │
│  │ discovery │  │ evidence │  │  engine  │  │   agent   │  │
│  │ (scan)    │  │ (collect)│  │ (certify)│  │ (review)  │  │
│  └────┬─────┘  └────┬─────┘  └────┬─────┘  └─────┬─────┘  │
│       │              │             │              │         │
│  ┌────▼─────┐  ┌─────▼────┐  ┌────▼─────┐  ┌────▼──────┐  │
│  │  config  │  │  policy  │  │  record  │  │  report   │  │
│  │ (load)   │  │ (eval)   │  │  (store) │  │ (format)  │  │
│  └──────────┘  └──────────┘  └──────────┘  └───────────┘  │
│                                                             │
│  ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌───────────┐  │
│  │  domain  │  │  expiry  │  │ override │  │  github   │  │
│  │ (types)  │  │ (calc)   │  │ (apply)  │  │ (actions) │  │
│  └──────────┘  └──────────┘  └──────────┘  └───────────┘  │
│                                                             │
│  ┌──────────┐                                               │
│  │  queue   │                                               │
│  │ (persist)│                                               │
│  └──────────┘                                               │
└─────────────────────────────────────────────────────────────┘
```

## Package Responsibilities

### `internal/domain/`
Core types with zero external dependencies. Defines: Unit, UnitID, Evidence, PolicyPack, CertificationRecord, Status, Grade, Dimension, Override, Config.

### `internal/discovery/`
Code unit discovery. Language adapters (Go via `go/ast`, TypeScript via regex, generic file-level). Index management with JSON persistence. Diff computation for added/removed/changed units.

### `internal/evidence/`
Evidence collection from external tools. Runs `go vet`, `go test`, `golangci-lint`, `git log`. Parses results into normalized `Evidence` structs. AST-based cyclomatic complexity measurement.

### `internal/config/`
Loads YAML config and policy packs. Policy matching (language, path targeting). Config validation.

### `internal/policy/`
Evaluates policy rules against evidence. Extracts metrics (lint errors, test failures, complexity, TODOs). Records violations with severity and dimension.

### `internal/engine/`
Certification pipeline: evaluate → score → status → grade → expiry → record. Multi-dimension scoring across 9 quality dimensions. Weighted averages.

### `internal/agent/`
Optional LLM-assisted review via OpenRouter. Pipeline architecture: Prescreen → Review → Scoring stages. Circuit breaker, model chain with fallback, adaptive messages. Model attribution tracking.

### `internal/expiry/`
Computes certification expiry windows based on risk factors: churn, complexity, coverage, security sensitivity, prior pass/fail history.

### `internal/record/`
Flat JSON file store for certification records. SHA256-hashed filenames for deterministic storage.

### `internal/report/`
Health reports (summary + detailed). Dimension breakdowns, by-language analysis, expiring-soon detection, highest-risk identification.

### `internal/override/`
Manual governance: exempt, extend/shorten windows, force review. YAML-based override definitions with required rationale.

### `internal/github/`
GitHub integration: PR comment formatting, issue creation/close commands, workflow YAML generation.

### `internal/queue/`
Persistent JSON-backed work queue for incremental processing across runs. Crash-safe (saves after each item).

## Data Model

### Configuration
```
.certification/
├── config.yml          # Mode, scope, agent, expiry, issues
├── policies/           # Versioned YAML policy packs
│   ├── global.yml
│   └── go-standard.yml
├── overrides/          # Manual exemptions/adjustments
├── records/            # Certification record JSON files
├── index.json          # Discovered unit index
└── queue.json          # Processing queue state
```

### Certification Record
Each record contains: unit identity, policy version, status, grade, score, confidence, dimension scores, evidence references, observations, actions, timestamps (certified_at, expires_at), source attribution.

## Design Principles

1. **Language-agnostic core**: Domain types have no language-specific assumptions
2. **Adapter pattern**: Language support added via Scanner interface implementations
3. **Deterministic first**: Agent review is optional and additive — never overrides deterministic evidence
4. **Repository-native**: All state lives in `.certification/` within the target repo
5. **Incremental**: Queue-based processing across multiple runs
6. **Auditable**: Records are JSON, versioned in git, human-reviewable
