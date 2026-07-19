# Project Context

## Project

DARP CLI

Developer AI Resource Platform

---

## Vision

DARP aims to become the package manager for AI assets.

Instead of managing software libraries, DARP manages reusable assets used by AI assistants and AI-enabled development tools.

---

## Long-Term Goals

- Discover AI assets
- Install AI assets
- Version AI assets
- Share AI assets
- Evaluate AI assets
- Convert assets between formats
- Support multiple AI providers
- Integrate with MCP ecosystem

---

## Current Non Goals

- Build an AI assistant.
- Replace existing package managers.
- Execute AI models.
- Provide cloud services.
- Become an IDE.

---

## Core Concepts

### Asset

Reusable AI artifact.

Examples:

- Prompt
- Instruction
- Skill
- Persona
- MCP Server
- Workflow
- Template

---

### Persona

A collection of assets installed together.

---

### Registry

Central repository containing versioned assets.

---

### Package

A distributable unit containing one or more assets.

---

### Provider

AI provider used for reasoning.

Examples:

- OpenAI
- Anthropic
- Google
- DeepSeek
- OpenRouter

---

## Architecture

CLI

↓

Registry

↓

Providers

↓

Assets

---

## Technology

Language

- Go

Package manager

- Go Modules

Version Control

- Git

Hosting

- GitHub

Development

- VS Code

AI Assistants

- GitHub Copilot
- OpenAI Codex

---

## Current Scope

Current milestone:

- Repository structure
- Development methodology
- SDD workflow
- Initial CLI architecture

---

## Out of Scope

Current version does NOT include:

- Web UI
- Marketplace
- Asset execution
- Cloud synchronization
- Authentication server

These features may be implemented in future projects.

---

## Project Vocabulary

### Glossary

Asset

Reusable AI artifact.

Persona

Collection of assets.

Registry

Repository of versioned assets.

Provider

LLM provider.

Package

Installable unit.

Specification

Formal description of a feature.

Plan

Technical implementation strategy.

Task

Executable implementation step.
