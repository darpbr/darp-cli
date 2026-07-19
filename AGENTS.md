# AGENTS.md

This document defines how AI agents must contribute to this repository.

## Mission

Help build DARP CLI while preserving architectural consistency.

Agents must prioritize quality over speed.

## Project Goal

DARP CLI is a package manager for AI assets.

The CLI manages reusable artifacts instead of source code libraries.

Examples:

- Prompt
- Instruction
- Skill
- Persona
- MCP Server
- Workflow

## Development Methodology

This repository follows Specification-Driven Development (SDD).

Implementation order is always:

1. Vision
2. Constitution
3. Specification
4. Plan
5. Tasks
6. Implementation
7. Review

Never skip steps.

## Architecture Principles

- Keep packages cohesive.
- Minimize coupling.
- Prefer composition over inheritance.
- Keep commands independent.
- Avoid global state.
- Favor deterministic behavior.
- Design for extensibility.

## Coding Principles

- Keep functions small.
- Prefer readability over cleverness.
- Document exported APIs.
- Write tests whenever functionality is implemented.
- Never introduce unnecessary abstractions.

## Decision Hierarchy

When conflicts arise, follow this order:

1. Constitution
2. Project Context
3. Approved Specification
4. ADR
5. Plan
6. Tasks

Never override a higher-level document.

## Repository Rules

Do not:

- implement features outside an approved specification;
- create hidden behavior;
- duplicate business logic;
- introduce unnecessary dependencies.

Always:

- update documentation;
- explain assumptions;
- preserve backward compatibility whenever possible.

## AI Behaviour

When uncertain:

- ask for clarification;
- do not invent requirements;
- document assumptions.

When implementing:

- prefer incremental changes;
- keep commits focused;
- avoid unrelated refactoring.
