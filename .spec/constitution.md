# DARP Constitution

Version: 1.0

This document defines the immutable principles governing DARP CLI.

---

## Principle 1

Specification Before Implementation

No feature may be implemented before an approved specification exists.

---

## Principle 2

Architecture First

Architectural consistency has priority over implementation speed.

---

## Principle 3

Provider Agnostic

The project must never depend on a single AI provider.

All provider-specific implementations must be isolated.

---

## Principle 4

AI-First Development

The repository is optimized for AI-assisted software engineering.

Repository structure should remain understandable by both humans and AI agents.

---

## Principle 5

Deterministic Behavior

Whenever possible:

- deterministic outputs
- reproducible workflows
- explicit assumptions

Avoid hidden behavior.

---

## Principle 6

Incremental Evolution

Large changes should be divided into small, reviewable specifications.

---

## Principle 7

Documentation is part of the software architecture.

Documentation is part of the implementation.

Every architectural change must update documentation.

---

## Principle 8

Simplicity

Prefer:

- simple APIs
- small packages
- explicit code
- low coupling

Avoid unnecessary abstractions.

---

## Principle 9

Open Standards

Prefer open specifications over proprietary formats whenever possible.

---

## Principle 10

Extensibility

Every subsystem should allow future extensions without major redesign.

---

## AI Transparency

AI-generated changes should be explainable.

Agents must document assumptions whenever they influence architectural decisions.

---

## Development Workflow

Every feature follows:

Vision

↓

Specification

↓

Plan

↓

Tasks

↓

Implementation

↓

Review

↓

Merge

Skipping steps is not allowed.

---

## Governance

This constitution has precedence over implementation decisions.

When conflicts arise:

Constitution

↓

Architecture

↓

Specification

↓

Implementation
