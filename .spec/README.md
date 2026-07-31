# Specification Workspace

This directory stores the repository-level governance and reusable assets that support the DARP Development Lifecycle.

## Contents

- `constitution.md`: immutable project principles
- `templates/`: reusable planning templates
- `specs/`: active feature bundles containing `spec.md`, `plan.md`, and `tasks.md`
- `archive/`: historical specifications and plans kept for traceability

## Lifecycle Role

`.spec/` supports the planning stages that must happen before implementation begins.

## Standard Feature Layout

Every feature uses one directory named `<number>-<short-name>`:

```text
.spec/specs/002-doctor/
├── spec.md
├── plan.md
├── tasks.md
└── adr/                         # optional
```

The three core files are created from the templates in `.spec/templates/` and
must remain together. ADRs are colocated with the feature when the feature has
an architectural decision.
