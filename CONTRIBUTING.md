# Contributing

Thanks for helping improve this project.

## Development

Requirements:

- Go **1.25** or newer (see `go.mod`; `toolchain go1.26.2` is pinned for consistency). CI runs against both 1.25 and 1.26.

Commands:

```bash
go test -race ./...
go vet ./...
```

Lint (install [golangci-lint](https://golangci-lint.run/) v2 locally):

```bash
golangci-lint run ./...
```

## Git hooks

This repo uses [Lefthook](https://lefthook.dev/) for local git hooks. One-time setup after cloning:

```bash
go install github.com/evilmartians/lefthook/v2@latest
lefthook install
```

Hooks:

- **pre-commit** — `gofmt -l`, `go vet`, `go mod tidy` check, and `golangci-lint run`.
- **commit-msg** — enforces [Conventional Commits](https://www.conventionalcommits.org/) (required for release-please).
- **pre-push** — `go test -race ./...`.

To bypass hooks in an emergency use `git commit --no-verify` / `git push --no-verify`, but prefer fixing the underlying issue.

## Commits and releases

Use [Conventional Commits](https://www.conventionalcommits.org/) (for example `feat:`, `fix:`, `chore:`, `docs:`, `ci:`). Breaking API changes should use a `!` after the type, e.g. `feat!: change Claims shape`.

[release-please](https://github.com/googleapis/release-please) opens release pull requests on `main` from that history. Merging a release PR creates a version tag and triggers [GoReleaser](https://goreleaser.com/) for CLI binaries.

## Pull requests

- Keep changes focused and covered by tests where it makes sense.
- Ensure `go test ./...` and lint pass locally before opening a PR.
- Describe the motivation and any breaking changes in the PR body (there is a template in `.github/`).

## Security

Do not open public issues for security problems. See [SECURITY.md](SECURITY.md).
