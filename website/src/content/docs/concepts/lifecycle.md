---
title: Certification Lifecycle
description: How certifications are created, maintained, and renewed.
---

## The Core Principle

**Trust in code should never be permanent.**

Every certification in Certify has an expiration date. When that date passes, the certification lapses and the code must be re-evaluated against current standards.

## Status Flow

```
                    ┌──────────┐
        ┌──────────▶│ Certified │◀─── recertification
        │           └─────┬────┘
        │                 │ time passes
        │           ┌─────▼────┐
        │           │ Expired  │
        │           └─────┬────┘
        │                 │ re-evaluate
  passes│           ┌─────▼────────┐
  policy│     ┌─────│ Evaluating   │─────┐
        │     │     └──────────────┘     │ fails
        │     │                          │ policy
        │     ▼                          ▼
        └─────┘                   ┌──────────────┐
                                  │ Decertified  │
                                  └──────────────┘
```

## Statuses

| Status | What it means | What happens next |
|--------|--------------|-------------------|
| **Certified** | Unit meets all policies | Valid until expiration |
| **Certified with Observations** | Passes but has warnings | Valid, but watch items noted |
| **Probationary** | Below threshold, grace period | Must improve before next evaluation |
| **Expired** | Certification window elapsed | Must be re-evaluated |
| **Decertified** | Fails required policies | Needs remediation |
| **Exempt** | Excluded by human override | Not evaluated |

## Certification Window

Each certification has a time window — by default, **90 days**.

The window adjusts based on risk factors:

| Factor | Effect |
|--------|--------|
| High git churn | Shorter window |
| Many recent authors | Shorter window |
| Low complexity, stable code | Longer window |
| High test coverage | Longer window |

Configurable bounds:
- Minimum: 7 days
- Maximum: 365 days

## Recertification

When certification expires, the unit enters the queue for re-evaluation on the next `certify certify` run. The process is automatic — expired units are rediscovered, re-evaluated with fresh evidence, and assigned new certification status.

## Continuous Process

The recommended workflow:

1. **PR** — Certify changed files before merge
2. **Nightly** — Sweep for expired certifications
3. **Weekly** — Full recertification run + report card update

This ensures certification state stays current without manual effort.
