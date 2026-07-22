# DARP CLI

> **Developer AI Resource Platform**

DARP CLI is a developer-first command-line interface for discovering, installing, managing, versioning and evaluating AI assets.

The project is inspired by package managers such as npm, pip and cargo, but instead of managing software libraries, DARP manages reusable AI assets.

Examples of supported assets include:

- Prompts
- Instructions
- Skills
- Personas
- MCP Servers
- Workflows
- Templates
- Context Packages

## Vision

Build an open ecosystem for AI engineering where reusable assets can be versioned, shared and installed as easily as software packages.

## Getting Started

Before contributing to this project, read:

1. AGENTS.md
2. docs/PROJECT_CONTEXT.md
3. .spec/constitution.md

These documents define the project's vision, development methodology and architectural principles.

## Current Command

The current implemented commands are:

```bash
darp init
darp --help
darp --version
```

This command initializes the current directory as a DARP project by creating `darp.yaml` and the base `.darp/` structure.

`darp --help` shows the CLI description and useful commands.

`darp --version` shows the CLI version embedded at build time.

## Local Development

Build the binary for the current platform:

```bash
make build
```

Install the CLI into `~/.local/bin`:

```bash
make install
```

Check the computed version string:

```bash
make version
```

The `Makefile` computes the version automatically from Git:

- if `HEAD` is tagged with a SemVer tag like `v0.1.0`, the CLI version becomes `0.1.0`
- otherwise it falls back to a development version like `0.1.0-dev+115ee40`
- if the worktree is dirty, the suffix becomes `0.1.0-dev+115ee40.dirty`

## DARP Development Lifecycle (DDL)

Every contribution to DARP follows the DARP Development Lifecycle (DDL), a Specification-Driven methodology designed for collaborative software engineering between humans and AI agents.

```mermaid
flowchart TD

    B["Bootstrap Repository"]

    V["Vision"]
    C["Constitution"]
    PC["Project Context"]

    S["Specification"]
    ADR["Architecture Decision Record"]
    P["Implementation Plan"]
    T["Tasks"]

    I["Implementation"]
    R["Review"]
    REL["Release"]

    B --> V
    V --> C
    C --> PC
    PC --> S
    S --> ADR
    ADR --> P
    P --> T
    T --> I
    I --> R
    R --> REL

    classDef bootstrap fill:#ede9fe,stroke:#7c3aed,color:#111827;
    classDef foundation fill:#dbeafe,stroke:#2563eb,color:#111827;
    classDef planning fill:#dcfce7,stroke:#16a34a,color:#111827;
    classDef execution fill:#fef3c7,stroke:#d97706,color:#111827;

    class B bootstrap;
    class V,C,PC foundation;
    class S,ADR,P,T planning;
    class I,R,REL execution;
```

### Lifecycle Phases

| Phase | Purpose |
| --------- | ---------- |
| **Foundation** | Defines the identity and long-term direction of the project. |
| **Planning** | Describes what will be built and how it will be implemented. |
| **Execution** | Implements, validates and prepares the feature for release. |

## Principles

- Specification-Driven Development (SDD)
- AI-first development workflow
- Provider agnostic
- Reproducible
- Deterministic
- Extensible
- Open standards whenever possible

## Project Status

🚧 Early development

Current milestone:

- Repository foundation
- Development methodology
- Initial architecture
- First CLI feature: `darp init`

## Development Workflow

Every feature follows the same lifecycle:

Vision
→ Constitution
→ Specification
→ Plan
→ Tasks
→ Implementation
→ Review

No feature should be implemented before an approved specification exists.

## Repository Layout

The repository now contains the first executable CLI baseline together with the existing specification-driven structure.

- `docs/`: project context, roadmap and lifecycle documentation
- `.spec/`: constitution, reusable templates and archival planning artifacts
- `cmd/`, `internal/`, `pkg/`, `test/`: implementation areas for CLI and supporting packages
- `assets/`: AI-oriented repository assets for tools such as Copilot and Codex
- `.github/`: prompts, agent guidance and workflow skeletons

## Supported AI Providers (planned)

- OpenAI
- Anthropic
- Google Gemini
- OpenRouter
- DeepSeek

## License

Apache 2.0 (planned)
