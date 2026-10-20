# check-default

Workflow definition for **check-default**. This page documents the checked-in GitHub Actions configuration; it does not assert that a run succeeds.

## At a glance

| | |
| --- | --- |
| Platform | GitHub Actions |
| Source | [`.github/workflows/check-default.yml`](../check-default.yml) |
| Trigger | pull_request; push: branches=main, 7.1*, 8.* |

## Purpose

This workflow runs the repository default checks: it checks out the source, sets up the configured Go and Python toolchains, and runs `make check-default`. The source does not state a separate motivation.



## When it runs

- pull_request
- push: branches=main, 7.1*, 8.*

Event and job filters are evaluated as shown in the workflow source.

## Execution flow

| Job | Runner | Condition | Depends on | Configured work |
| --- | --- | --- | --- | --- |
| `check` | `ubuntu-latest` | `Always when its event is eligible` | None declared | `actions/checkout@v7`; `actions/setup-go@v7`; Fix Code is not compatible with Python 3.12 (`actions/setup-python@v7`); Run check-default — runs `go install github.com/magefile/mage`; `make check-default` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | No workflow inputs declared; environment names: none declared |
| Outputs | No job outputs declared; artifact/cache actions: none declared |
| Permissions | contents: read; pull-requests: read |
| Secrets | No `secrets.*` references detected in this file |

## Operational notes

- No additional retry, timeout, or deployment behavior is declared in the checked-in workflow.

## Sources

- Workflow: [`.github/workflows/check-default.yml`](../check-default.yml)
- External action/workflow reference: `actions/checkout@v7` (implementation is outside this checkout).
- External action/workflow reference: `actions/setup-go@v7` (implementation is outside this checkout).
- External action/workflow reference: `actions/setup-python@v7` (implementation is outside this checkout).
