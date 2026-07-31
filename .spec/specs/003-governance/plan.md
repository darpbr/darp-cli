# Plan 003 - Governance

## Status

Draft

## Related Specification

- `spec.md`

## Implementation Strategy

Create the governance foundation described by the specification without
changing existing CLI behavior. The work is limited to lifecycle, quality-gate,
documentation, and agent-skill artifacts.

## Affected Areas

- `.darp/governance/`
- `.darp/lifecycle.md`
- `.agents/skills/`
- `README.md`

## Validation Strategy

- Confirm every required governance file exists.
- Confirm each agent skill contains a valid `SKILL.md`.
- Confirm the lifecycle references the governance skills and quality gates.
- Confirm existing CLI tests and commands remain unchanged.

## Assumptions

- `.darp/` stores DARP runtime contracts.
- `.agents/skills/` stores skills discovered by coding agents.
- Optional skill resources do not need to be created unless required by a skill.
