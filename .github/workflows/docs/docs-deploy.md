# docs-deploy

Workflow definition for **docs-deploy**. This page documents the checked-in GitHub Actions configuration; it does not assert that a run succeeds.

## At a glance

| | |
| --- | --- |
| Platform | GitHub Actions |
| Source | [`.github/workflows/docs-deploy.yml`](../docs-deploy.yml) |
| Trigger | workflow_run: types=completed; workflows=docs-build |

## Purpose

This workflow acts as an event wrapper: it delegates `deploy` to `elastic/docs-actions/.github/workflows/docs-deploy.yml@v1`. The external workflow implementation is not checked into this repository, and the source does not state a separate rationale for the wrapper.



## When it runs

- workflow_run: types=completed; workflows=docs-build

Event and job filters are evaluated as shown in the workflow source.

## Execution flow

| Job | Runner | Condition | Depends on | Configured work |
| --- | --- | --- | --- | --- |
| `deploy` | `Reusable workflow or not specified` | `Always when its event is eligible` | None declared | Calls reusable workflow `elastic/docs-actions/.github/workflows/docs-deploy.yml@v1` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | No workflow inputs declared; environment names: none declared |
| Outputs | No job outputs declared; artifact/cache actions: none declared |
| Permissions | contents: read; deployments: write; id-token: write; pull-requests: write; actions: read |
| Secrets | No `secrets.*` references detected in this file |

## Operational notes

- Reusable workflow internals are hosted externally and are not available in this checkout.

## Sources

- Workflow: [`.github/workflows/docs-deploy.yml`](../docs-deploy.yml)
- External action/workflow reference: `elastic/docs-actions/.github/workflows/docs-deploy.yml@v1` (implementation is outside this checkout).
