# Changelog

## [1.0.1](https://github.com/dendavidov/go-utils/compare/v1.0.0...v1.0.1) (2026-04-21)


### Bug Fixes

* **release:** GoReleaser v2 checksum field and goreleaser token ([8e0f9d7](https://github.com/dendavidov/go-utils/commit/8e0f9d7a23c777999d71cb6341bccd60b620e3d8))

## 1.0.0 (2026-04-21)


### ⚠ BREAKING CHANGES

* **pkg:** add importable jwt helpers and bump module to go-utils

### Features

* add jwt token ([baa0e55](https://github.com/dendavidov/go-utils/commit/baa0e550b56f0d228c9dfaec94c8124fc0a3e069))
* **cmd:** add jwt-token CLI for minting RS256 tokens ([bb571df](https://github.com/dendavidov/go-utils/commit/bb571df012ef61c1223356c12050ac5e08415ae3))
* **pkg:** add importable jwt helpers and bump module to go-utils ([84b99f6](https://github.com/dendavidov/go-utils/commit/84b99f607814bd3b6dde47c27c127d1da58e19b9))


### Bug Fixes

* **ci:** add actions:read for CodeQL analyze ([6ed0cf2](https://github.com/dendavidov/go-utils/commit/6ed0cf2f78e0ec842b74b36b002c764e8ecc533c))
* **ci:** correct release-please token and document PR permission ([b793353](https://github.com/dendavidov/go-utils/commit/b7933539023cacbce0535dd0c88b862447985692))
* **ci:** make Go 1.25 matrix actually test on Go 1.25 ([261fdeb](https://github.com/dendavidov/go-utils/commit/261fdeb54c94a6069ecb02f969de743d156abb79))
* **ci:** repair golangci-lint v2 config and Go 1.25 Actions matrix ([f417418](https://github.com/dendavidov/go-utils/commit/f417418134b7212e0a1d0b5c1c5b2faaf3e93da5))
* **ci:** run go test under bash on Windows ([27502a5](https://github.com/dendavidov/go-utils/commit/27502a5f4af883bf8d42ec16acca84655e989665))
* **ci:** run GoReleaser directly from release-please job ([9a42ae5](https://github.com/dendavidov/go-utils/commit/9a42ae5dd25c8e08c4bf2dad55540ca4b1835d96))
* **pkg:** preserve caller-set RegisteredClaims in GenerateToken ([3584098](https://github.com/dendavidov/go-utils/commit/35840981d44821159c18624dd212aa13eecab760))

## Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

Release notes after the initial modernization are produced by
[release-please](https://github.com/googleapis/release-please) from
[Conventional Commits](https://www.conventionalcommits.org/).

<!-- release-please will populate this file starting from the first release. -->
