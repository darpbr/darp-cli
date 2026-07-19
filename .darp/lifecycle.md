# `.darp/lifecycle.md`

````md
# DARP Lifecycle

## Purpose

The DARP Lifecycle defines how software is planned, designed, implemented and maintained.

It is the single source of truth describing the engineering process adopted by this repository.

This lifecycle is designed for collaborative software development between humans and AI agents.

It is intentionally provider-agnostic and independent of any specific AI model or tool.

---

# Principles

Every change should be:

- intentional
- traceable
- incremental
- explainable
- reviewable

Implementation is never the first step.

Planning always precedes coding.

---

# Lifecycle

```text
Vision
    ↓
Constitution
    ↓
Project Context
    ↓
Specification
    ↓
Architecture Decision Record (ADR)
    ↓
Implementation Plan
    ↓
Tasks
    ↓
Implementation
    ↓
Review
    ↓
Release
```

---

# Phase Descriptions

## Vision

Defines why the project exists.

Changes rarely.

---

## Constitution

Defines the permanent engineering principles.

It has higher priority than every other document except the Vision.

---

## Project Context

Defines the shared knowledge required to understand the project.

Includes terminology, scope, constraints and long-term objectives.

---

## Specification

Describes **what** will be built.

A Specification must be approved before implementation begins.

---

## Architecture Decision Record (ADR)

Captures significant architectural decisions.

Every important technical decision should be documented.

---

## Implementation Plan

Describes **how** the Specification will be implemented.

Large features may require multiple implementation plans.

---

## Tasks

Break the implementation plan into small, verifiable work items.

Tasks should be independent whenever possible.

---

## Implementation

Only begins after the previous phases have been completed.

Implementation should never redefine the Specification.

---

## Review

Verifies correctness, consistency and alignment with the project principles.

Reviews include both human and AI feedback whenever appropriate.

---

## Release

Makes the completed work available for consumption.

Every release should be traceable to its Specification and ADRs.

---

# Document Hierarchy

When documents disagree, follow this order:

1. Vision
2. Constitution
3. Project Context
4. Approved Specification
5. Architecture Decision Record
6. Implementation Plan
7. Tasks

Lower-level documents must never contradict higher-level documents.

---

# AI Collaboration

AI agents participate throughout the lifecycle.

Agents should:

- read existing documentation before proposing changes;
- preserve architectural consistency;
- explain assumptions;
- avoid speculative implementations;
- request clarification when requirements conflict.

AI agents assist engineering decisions.

They do not replace project governance.

---

# Future Evolution

The lifecycle is intentionally small.

Additional phases or artifacts should only be introduced when they clearly improve traceability, maintainability or developer experience.

The goal is to keep the lifecycle understandable, scalable and reusable across different projects.
````
