# Observability Agentic Workflow — Pull Request

Workflow definition for **Observability Agentic Workflow — Pull Request**. This page documents the checked-in GitHub Actions configuration; it does not assert that a run succeeds.

## At a glance

| | |
| --- | --- |
| Platform | GitHub Actions |
| Source | [`.github/workflows/trigger-obs-aw-pull-request.yml`](../trigger-obs-aw-pull-request.yml) |
| Trigger | pull_request: types=opened, synchronize, reopened, labeled |

## Purpose

This workflow acts as an event wrapper: it delegates `run-obs-aw-pull-request` to `elastic/oblt-aw/.github/workflows/obs-aw-event-pull-request.yml@main`. The external workflow implementation is not checked into this repository, and the source does not state a separate rationale for the wrapper.



## When it runs

- pull_request: types=opened, synchronize, reopened, labeled

Event and job filters are evaluated as shown in the workflow source.

## Execution flow

| Job | Runner | Condition | Depends on | Configured work |
| --- | --- | --- | --- | --- |
| `run-obs-aw-pull-request` | `Reusable workflow or not specified` | `Always when its event is eligible` | None declared | Calls reusable workflow `elastic/oblt-aw/.github/workflows/obs-aw-event-pull-request.yml@main` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | No workflow inputs declared; environment names: none declared |
| Outputs | No job outputs declared; artifact/cache actions: none declared |
| Permissions | contents: read; run-obs-aw-pull-request.actions: read; run-obs-aw-pull-request.contents: write; run-obs-aw-pull-request.copilot-requests: write; run-obs-aw-pull-request.discussions: write; run-obs-aw-pull-request.id-token: write; run-obs-aw-pull-request.issues: write; run-obs-aw-pull-request.pull-requests: write |
| Secrets | No `secrets.*` references detected in this file |

## Operational notes

- Reusable workflow internals are hosted externally and are not available in this checkout.

## Sources

- Workflow: [`.github/workflows/trigger-obs-aw-pull-request.yml`](../trigger-obs-aw-pull-request.yml)
- External action/workflow reference: `elastic/oblt-aw/.github/workflows/obs-aw-event-pull-request.yml@main` (implementation is outside this checkout).
