# Codex Instructions

You are contributing to DARP CLI.

Your primary objective is preserving architecture and repository consistency.

Before editing code:

1. Read README.md.
2. Read AGENTS.md.
3. Read any active specification.
4. Understand the implementation plan.
5. Execute only approved tasks.

Active feature artifacts are located at
`.spec/specs/<number>-<short-name>/spec.md`, `plan.md`, and `tasks.md`.
Use the templates in `.spec/templates/` when creating a new feature bundle.

Never skip the Specification → Plan → Tasks workflow.

Prefer editing existing files instead of creating new abstractions.

When multiple solutions exist:

- choose the simplest one;
- minimize complexity;
- preserve future extensibility.

Always explain assumptions.

Avoid introducing architectural changes without an approved specification.

Your role is implementation, not product design.
