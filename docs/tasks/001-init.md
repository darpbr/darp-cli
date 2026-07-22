# Tasks 001 - `darp init`

## Related Plan

- `docs/plans/001-init.md`

## Task List

1. Create the Go module and executable entrypoint.
2. Implement CLI dispatch for the `init` command with friendly output.
3. Implement an initialization service that:
   - derives the project name from the current directory;
   - creates `darp.yaml` when absent;
   - creates the `.darp` directory tree when absent;
   - writes embedded `lifecycle.md` when absent;
   - returns a stable result for already initialized projects.
4. Add an OS-backed filesystem adapter and a test double for deterministic tests.
5. Add automated tests for:
   - fresh initialization;
   - repeated execution without overwriting files;
   - expected file contents and directory layout.
6. Update repository docs and developer commands for the new executable baseline.

## Validation Checklist

- `go test ./...`
- manual smoke flow:
  - `darp init`
  - `darp init` again
- confirm no files are overwritten on repeat execution

## Notes and Assumptions

- friendly messages may differ slightly from the examples in the specification as long as they stay clear and aligned with the required steps
- no external dependencies should be introduced for this first feature
