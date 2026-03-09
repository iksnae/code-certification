# Certify — Brand Guide

## Product Name
**Certify**

## Product Category
**Continuous Code Certification**

## Short Description
Certify continuously evaluates every code unit in your repository, scores it against versioned policies, and assigns time-bound certification you can actually trust.

## Tagline
**Code trust, with an expiration date.**

---

## Brand Narrative

Most engineering teams assume code that once passed review can be trusted indefinitely.

That assumption is wrong.

Standards evolve. Dependencies change. Systems grow more complex. Code that once met expectations slowly drifts away from them.

Certify introduces a new engineering discipline: **code certification**.

Certify evaluates code against defined policies, assigns measurable quality scores, and issues certifications that **expire intentionally**. When certification expires, code must be reviewed again against the latest standards.

Instead of treating quality as a one-time event, Certify makes it a **continuous process of trust, verification, and renewal**.

CI tells you whether code passes right now.

**Certify tells you whether code should still be trusted.**

---

## Brand Pillars

### 1. Time-Bound Trust
Code certifications expire by design. Trust must be renewed as standards evolve.

### 2. Evidence-Driven Quality
Certification decisions are based on deterministic evidence and policy evaluation.

### 3. Continuous Governance
Quality governance becomes an automated, recurring process rather than an occasional audit.

### 4. Repository-Native Transparency
Certification state lives inside the repository and evolves alongside the codebase.

### 5. Developer-First Integration
Certify integrates into GitHub workflows and developer tooling without disrupting existing practices.

---

## Core Terminology

Standard terminology across docs, CLI output, and UI.

| Term | Definition |
|------|------------|
| Code Unit | A certifiable piece of code (file, function, module, workflow, etc.) |
| Certification | A time-bound approval of a code unit's quality |
| Certification Window | The time period for which certification is valid |
| Recertification | Re-evaluating a unit after expiration or change |
| Policy Pack | A set of rules used to evaluate code quality |
| Report Card | The summary output showing certification results |
| Certification Badge | Repository badge showing certification status |

---

## Certification Status Language

Statuses must remain consistent across all documentation.

| Status | Meaning |
|--------|---------|
| Certified | Meets all required policies |
| Certified with Observations | Acceptable but with minor issues |
| Probationary | Requires improvement soon |
| Expired | Certification window has elapsed |
| Decertified | Fails required policies |
| Exempt | Explicitly excluded from certification |

---

## Messaging

### One-Sentence Pitch
Certify continuously evaluates your repository and assigns time-bound certification so you always know which code is still trustworthy.

### Short Product Description
Certify is a code governance platform that evaluates every code unit in a repository, scores it against versioned policies, and assigns certification that expires intentionally to prevent quality drift.

### Elevator Pitch
Most codebases assume that code once reviewed stays trustworthy forever. Certify replaces that assumption with measurable, time-bound certification. It continuously evaluates code quality, assigns certification status, and ensures trust is renewed over time.

---

## CLI Branding

### Command
`certify`

### CLI Tone
Output should be:
- Precise
- Concise
- Evidence-based
- Calm and professional
- Free of unnecessary verbosity

---

## Visual Identity

### Design Personality

The brand should communicate:
- Precision
- Trust
- Durability
- Engineering rigor
- Minimalism

Avoid:
- Playful mascots
- Flashy startup colors
- Security-fear marketing
- Compliance bureaucracy aesthetics

### Logo Concept

**Primary Motif: Certification Seal**

A circular geometric mark representing verification, renewal cycles, and structured governance.

- **Outer ring** — the certification lifecycle
- **Inner mark** — validated trust
- Possible elements: segmented ring, minimal check mark, geometric seal

### Color Palette

**Primary:**

| Name | Hex |
|------|-----|
| Graphite | `#1C1C1C` |
| Slate | `#2C3A40` |
| Steel Blue | `#4A6B82` |

**Status:**

| Status | Color | Hex |
|--------|-------|-----|
| Certified | Green | `#2E8B57` |
| Observations | Amber | `#E0A100` |
| Probationary | Warning | `#F59E0B` |
| Expired | Gray | `#9CA3AF` |
| Decertified | Red | `#DC2626` |

**Backgrounds:**

| Use | Hex |
|-----|-----|
| Light | `#F8FAFC` |
| Muted | `#EEF2F5` |

### Typography

**Primary:** Inter, IBM Plex Sans, Source Sans, JetBrains Sans

**Monospace (CLI/docs):** JetBrains Mono, Source Code Pro

### Badge Emoji

| Status | Emoji |
|--------|-------|
| Certified | 🟢 |
| Observations | 🟡 |
| Probationary | 🟠 |
| Decertified | 🔴 |
| Expired | ⚪ |

---

## Voice and Tone

All messaging should be:
- Precise
- Technical
- Calm
- Evidence-oriented
- Trustworthy

**Avoid:** "revolutionary", "AI powered magic", "instant code perfection"

**Prefer:** certification, evidence, policy, trust, verification, governance, lifecycle

---

## Guiding Principle

**Trust in code should never be permanent.**

Certify exists to ensure that code trust is **measured, visible, and renewed over time.**
