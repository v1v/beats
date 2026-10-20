# docs-build

Workflow definition for **docs-build**. This page documents the checked-in GitHub Actions configuration; it does not assert that a run succeeds.

## At a glance

| | |
| --- | --- |
| Platform | GitHub Actions |
| Source | [`.github/workflows/docs-build.yml`](../docs-build.yml) |
| Trigger | pull_request: types=opened, synchronize, reopened; push: branches=main, [0-9]+.[0-9]+; merge_group |

## Purpose

This workflow acts as an event wrapper: it delegates `build` to `elastic/docs-actions/.github/workflows/docs-build.yml@v1`. The external workflow implementation is not checked into this repository, and the source does not state a separate rationale for the wrapper.



## When it runs

- pull_request: types=opened, synchronize, reopened
- push: branches=main, [0-9]+.[0-9]+
- merge_group

Event and job filters are evaluated as shown in the workflow source.

## Execution flow

| Job | Runner | Condition | Depends on | Configured work |
| --- | --- | --- | --- | --- |
| `build` | `Reusable workflow or not specified` | `Always when its event is eligible` | None declared | Calls reusable workflow `elastic/docs-actions/.github/workflows/docs-build.yml@v1` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | No workflow inputs declared; environment names: none declared |
| Outputs | No job outputs declared; artifact/cache actions: none declared |
| Permissions | contents: read; pull-requests: read |
| Secrets | No `secrets.*` references detected in this file |

## Operational notes

- Reusable workflow internals are hosted externally and are not available in this checkout.

## Sources

- Workflow: [`.github/workflows/docs-build.yml`](../docs-build.yml)
- External action/workflow reference: `elastic/docs-actions/.github/workflows/docs-build.yml@v1` (implementation is outside this checkout).
