# Observability Agentic Workflow — Issues

Workflow definition for **Observability Agentic Workflow — Issues**. This page documents the checked-in GitHub Actions configuration; it does not assert that a run succeeds.

## At a glance

| | |
| --- | --- |
| Platform | GitHub Actions |
| Source | [`.github/workflows/trigger-obs-aw-issues.yml`](../trigger-obs-aw-issues.yml) |
| Trigger | issues: types=opened, labeled, edited; workflow_dispatch |

## Purpose

This workflow acts as an event wrapper: it delegates `run-obs-aw-issues` to `elastic/oblt-aw/.github/workflows/obs-aw-event-issues.yml@main`. The external workflow implementation is not checked into this repository, and the source does not state a separate rationale for the wrapper.



## When it runs

- issues: types=opened, labeled, edited
- workflow_dispatch

Event and job filters are evaluated as shown in the workflow source.

## Execution flow

| Job | Runner | Condition | Depends on | Configured work |
| --- | --- | --- | --- | --- |
| `run-obs-aw-issues` | `Reusable workflow or not specified` | `Always when its event is eligible` | None declared | Calls reusable workflow `elastic/oblt-aw/.github/workflows/obs-aw-event-issues.yml@main` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | No workflow inputs declared; environment names: none declared |
| Outputs | No job outputs declared; artifact/cache actions: none declared |
| Permissions | contents: read; run-obs-aw-issues.actions: read; run-obs-aw-issues.contents: write; run-obs-aw-issues.copilot-requests: write; run-obs-aw-issues.discussions: write; run-obs-aw-issues.id-token: write; run-obs-aw-issues.issues: write; run-obs-aw-issues.pull-requests: write |
| Secrets | No `secrets.*` references detected in this file |

## Operational notes

- Reusable workflow internals are hosted externally and are not available in this checkout.

## Sources

- Workflow: [`.github/workflows/trigger-obs-aw-issues.yml`](../trigger-obs-aw-issues.yml)
- External action/workflow reference: `elastic/oblt-aw/.github/workflows/obs-aw-event-issues.yml@main` (implementation is outside this checkout).
