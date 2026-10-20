# docs-preview-cleanup

Workflow definition for **docs-preview-cleanup**. This page documents the checked-in GitHub Actions configuration; it does not assert that a run succeeds.

## At a glance

| | |
| --- | --- |
| Platform | GitHub Actions |
| Source | [`.github/workflows/docs-preview-cleanup.yml`](../docs-preview-cleanup.yml) |
| Trigger | pull_request_target: types=closed |

## Purpose

This workflow acts as an event wrapper: it delegates `cleanup` to `elastic/docs-actions/.github/workflows/docs-preview-cleanup.yml@v1`. The external workflow implementation is not checked into this repository, and the source does not state a separate rationale for the wrapper.



## When it runs

- pull_request_target: types=closed

Event and job filters are evaluated as shown in the workflow source.

## Execution flow

| Job | Runner | Condition | Depends on | Configured work |
| --- | --- | --- | --- | --- |
| `cleanup` | `Reusable workflow or not specified` | `Always when its event is eligible` | None declared | Calls reusable workflow `elastic/docs-actions/.github/workflows/docs-preview-cleanup.yml@v1` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | No workflow inputs declared; environment names: none declared |
| Outputs | No job outputs declared; artifact/cache actions: none declared |
| Permissions | contents: none; deployments: write; id-token: write |
| Secrets | No `secrets.*` references detected in this file |

## Operational notes

- Reusable workflow internals are hosted externally and are not available in this checkout.

## Sources

- Workflow: [`.github/workflows/docs-preview-cleanup.yml`](../docs-preview-cleanup.yml)
- External action/workflow reference: `elastic/docs-actions/.github/workflows/docs-preview-cleanup.yml@v1` (implementation is outside this checkout).
