# Plan 001 - `darp init`

## Related Specification

- `.darp/specs/001-init.md`

## Implementation Strategy

Implement the first executable CLI path for DARP with a small, testable Go architecture:

1. add a root executable entrypoint that embeds `.darp/lifecycle.md` into the binary;
2. keep argument parsing and user messaging in a CLI layer;
3. isolate initialization rules in an internal application service;
4. isolate filesystem access behind a small interface to support deterministic tests;
5. validate creation and idempotency with automated tests built around temporary directories.

The implementation should use only the Go standard library.

## Affected Areas

- `go.mod`
- `main.go`
- `cmd/` or internal CLI packages for command dispatch
- `internal/` packages for initialization rules and filesystem access
- `test/` or package-local `_test.go` files for automated coverage
- `README.md`
- `Makefile`
- `.github/workflows/ci.yml`

## Milestones

1. Create the implementation plan and tasks for spec 001.
2. Bootstrap the Go module and executable entrypoint.
3. Implement the `darp init` command flow and initialization service.
4. Add tests for first-run initialization and repeat-run idempotency.
5. Update repository documentation and validation commands for the new executable baseline.

## Validation Strategy

- run `go test ./...`
- verify the service creates:
  - `darp.yaml`
  - `.darp/lifecycle.md`
  - `.darp/specs`
  - `.darp/tasks`
  - `.darp/docs`
  - `.darp/templates`
  - `.darp/examples`
- verify a second execution does not overwrite existing files
- verify the CLI exits successfully for both first-run and already-initialized scenarios

## Rollback or Recovery Notes

- the feature is additive and localized to the new Go implementation
- if necessary, revert only the files introduced for the first CLI feature
- repeated execution of `darp init` must remain safe and non-destructive

## Assumptions

- the repository module path should follow the configured Git remote: `github.com/darpbr/darp-cli`
- the lifecycle document stored at `.darp/lifecycle.md` is the canonical content to embed in the binary
- an already initialized project can be detected when either `darp.yaml` or `.darp` already exists, because creating missing pieces in that state could risk unexpected partial mutation
