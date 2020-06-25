# Krank nmap scanner

[![GitHub Release](https://img.shields.io/github/v/release/tacheshun/krank)](https://github.com/tacheshun/krank/releases)
[![go.dev](https://img.shields.io/badge/go.dev-reference-blue.svg)](https://pkg.go.dev/github.com/tacheshun/krank)
[![go.mod](https://img.shields.io/github/go-mod/go-version/tacheshun/krank)](go.mod)
[![Build Status](https://img.shields.io/github/workflow/status/tacheshun/krank/build)](https://github.com/tacheshun/krank/actions?query=workflow%3Abuild+branch%3Amaster)
[![Go Report Card](https://goreportcard.com/badge/github.com/tacheshun/krank)](https://goreportcard.com/report/github.com/tacheshun/krank)

## Build

- Terminal: `make` to get help for make targets.
- Terminal: `make all` to execute a full build.
- Visual Studio Code: `Terminal` → `Run Build Task... (CTRL+ALT+B)` to execute a fast build.

## Release

The release workflow is triggered each time a tag with `v` prefix is pushed.

This repo uses [Github Tag Bump](https://github.com/marketplace/actions/github-tag-bump) for auto tagging on master branch. It automatically triggers the release workflow.

- Add `#minor` to your commit message to bump minor version.
- Add `#major` to your commit message to bump major version. DANGER! Use it with caution and make sure you understand the consequences. More info: [Go Wiki](https://github.com/golang/go/wiki/Modules#releasing-modules-v2-or-higher), [Go Blog][https://blog.golang.org/v2-go-modules].

## Maintainance

Remember to update Go version in [.github/workflows](.github/workflows), [Makefile](Makefile) and [devcontainer.json](.devcontainer/devcontainer.json).

Notable files:
- [devcontainer.json](.devcontainer/devcontainer.json) - Visual Studio Code Remote Container configuration
- [.github/workflows](.github/workflows) - GitHub Actions workflows
- [.vscode](.vscode) - Visual Studio Code configuration files
- [.golangci.yml](.golangci.yml) - golangci-lint configuration
- [.goreleaser.yml](.goreleaser.yml) - GoReleaser configuration
- [Makefile](Makefile) - Make targets used for development, [CI build](.github/workflows) and [.vscode/tasks.json](.vscode/tasks.json)
- [tools.go](tools.go) - build tools 
