# golangci-lint

Workflow definition for **golangci-lint**. This page documents the checked-in GitHub Actions configuration; it does not assert that a run succeeds.

## At a glance

| | |
| --- | --- |
| Platform | GitHub Actions |
| Source | [`.github/workflows/golangci-lint.yml`](../golangci-lint.yml) |
| Trigger | workflow_dispatch; pull_request: paths=**/*.go, **/go.mod, **/go.sum, .go-version, .golangci.yml, .github/workflows/golangci-lint.yml |

## Purpose

This workflow provides the `golangci-lint` CI task: it prepares or invokes `actions/checkout@v7`, `actions/setup-go@v7`, `golangci/golangci-lint-action` and then runs `sudo apt-get update && sudo apt-get install -y libpcap-dev librpm-dev`. The checked-in source does not state a separate motivation.



## When it runs

- workflow_dispatch
- pull_request: paths=**/*.go, **/go.mod, **/go.sum, .go-version, .golangci.yml, .github/workflows/golangci-lint.yml

Event and job filters are evaluated as shown in the workflow source.

## Execution flow

| Job | Runner | Condition | Depends on | Configured work |
| --- | --- | --- | --- | --- |
| `golangci` | `${{ matrix.os }}` | `Always when its event is eligible` | None declared | `actions/checkout@v7`; `actions/setup-go@v7`; Install Apt Package — runs `sudo apt-get update && sudo apt-get install -y libpcap-dev librpm-dev`; golangci-lint (`golangci/golangci-lint-action`) |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | No workflow inputs declared; environment names: CGO_ENABLED |
| Outputs | No job outputs declared; artifact/cache actions: none declared |
| Permissions | contents: read; pull-requests: read |
| Secrets | No `secrets.*` references detected in this file |

## Operational notes

- golangci uses a matrix/strategy declared in the source

## Sources

- Workflow: [`.github/workflows/golangci-lint.yml`](../golangci-lint.yml)
- External action/workflow reference: `actions/checkout@v7` (implementation is outside this checkout).
- External action/workflow reference: `actions/setup-go@v7` (implementation is outside this checkout).
- External action/workflow reference: `golangci/golangci-lint-action` (implementation is outside this checkout).
