# go-utils

Small Go utilities with a versioned module and conventional releases.

[![CI](https://github.com/dendavidov/go-utils/actions/workflows/ci.yml/badge.svg)](https://github.com/dendavidov/go-utils/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/dendavidov/go-utils.svg)](https://pkg.go.dev/github.com/dendavidov/go-utils)
[![Go Report Card](https://goreportcard.com/badge/github.com/dendavidov/go-utils)](https://goreportcard.com/report/github.com/dendavidov/go-utils)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

## Install

Library:

```bash
go get github.com/dendavidov/go-utils@latest
```

CLI (requires a recent Go toolchain):

```bash
go install github.com/dendavidov/go-utils/cmd/jwt-token@latest
```

Prebuilt binaries appear on [GitHub Releases](https://github.com/dendavidov/go-utils/releases) when a tag is published.

## Package `pkg/jwt`

Mint and verify RS256 JWTs with custom claims (`id`, `email`, `isAccessToken`) and standard registered claims (`exp`, `iat`, `nbf`).

```go
import "github.com/dendavidov/go-utils/pkg/jwt"

priv, err := jwt.ParsePrivateKeyFromBase64(os.Getenv("JWT_PRIVATE_KEY_B64"))
if err != nil { /* ... */ }

token, err := jwt.GenerateToken(priv, jwt.Claims{
    ID: "user-1", Email: "a@example.com", IsAccess: true,
}, 24*time.Hour)

claims, err := jwt.ParseToken(token, &priv.PublicKey)
```

## CLI `jwt-token`

Prints a signed JWT to stdout (no trailing newline beyond what `fmt.Print` does for the token string).

```text
jwt-token -id USER -email user@example.com -key "$JWT_PRIVATE_KEY_B64"
jwt-token -id USER -email user@example.com -key-file ./private.pem -ttl 720h
```

Environment variables: `JWT_SUBJECT_ID`, `JWT_EMAIL`, `JWT_PRIVATE_KEY_B64`.

Flags override env when set. Use `-version` to print the release version (set at link time by GoReleaser).

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md). This project uses [Conventional Commits](https://www.conventionalcommits.org/) and [release-please](https://github.com/googleapis/release-please) for changelog and tags.

## License

MIT — see [LICENSE](LICENSE).
