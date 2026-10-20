# check-dev-tools

Workflow definition for **check-dev-tools**. This page documents the checked-in GitHub Actions configuration; it does not assert that a run succeeds.

## At a glance

| | |
| --- | --- |
| Platform | GitHub Actions |
| Source | [`.github/workflows/check-dev-tools.yml`](../check-dev-tools.yml) |
| Trigger | pull_request: paths=.github/workflows/check-dev-tools.yml, dev-tools/**, libbeat/scripts/Makefile |

## Purpose

This workflow provides the `check-dev-tools` CI task: it prepares or invokes `actions/checkout@v7`, `actions/setup-go@v7`, `actions/setup-python@v7` and then runs `go install github.com/magefile/mage`, `make -C ${{ env.BEAT_MODULE }} check`. The checked-in source does not state a separate motivation.



## When it runs

- pull_request: paths=.github/workflows/check-dev-tools.yml, dev-tools/**, libbeat/scripts/Makefile

Event and job filters are evaluated as shown in the workflow source.

## Execution flow

| Job | Runner | Condition | Depends on | Configured work |
| --- | --- | --- | --- | --- |
| `check` | `ubuntu-latest` | `Always when its event is eligible` | None declared | `actions/checkout@v7`; `actions/setup-go@v7`; `actions/setup-python@v7`; Run check/update — runs `go install github.com/magefile/mage`; `make -C ${{ env.BEAT_MODULE }} check` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | No workflow inputs declared; environment names: BEAT_MODULE |
| Outputs | No job outputs declared; artifact/cache actions: none declared |
| Permissions | contents: read |
| Secrets | No `secrets.*` references detected in this file |

## Operational notes

- No additional retry, timeout, or deployment behavior is declared in the checked-in workflow.

## Sources

- Workflow: [`.github/workflows/check-dev-tools.yml`](../check-dev-tools.yml)
- External action/workflow reference: `actions/checkout@v7` (implementation is outside this checkout).
- External action/workflow reference: `actions/setup-go@v7` (implementation is outside this checkout).
- External action/workflow reference: `actions/setup-python@v7` (implementation is outside this checkout).
