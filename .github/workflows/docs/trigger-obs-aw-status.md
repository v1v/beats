# Observability Agentic Workflow — Status

Workflow definition for **Observability Agentic Workflow — Status**. This page documents the checked-in GitHub Actions configuration; it does not assert that a run succeeds.

## At a glance

| | |
| --- | --- |
| Platform | GitHub Actions |
| Source | [`.github/workflows/trigger-obs-aw-status.yml`](../trigger-obs-aw-status.yml) |
| Trigger | status |

## Purpose

This workflow acts as an event wrapper: it delegates `run-obs-aw-status` to `elastic/oblt-aw/.github/workflows/obs-aw-event-status.yml@main`. The external workflow implementation is not checked into this repository, and the source does not state a separate rationale for the wrapper.



## When it runs

- status

Event and job filters are evaluated as shown in the workflow source.

## Execution flow

| Job | Runner | Condition | Depends on | Configured work |
| --- | --- | --- | --- | --- |
| `run-obs-aw-status` | `Reusable workflow or not specified` | `github.event.state == 'failure' && contains(github.event.context, 'buildkite')` | None declared | Calls reusable workflow `elastic/oblt-aw/.github/workflows/obs-aw-event-status.yml@main` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | No workflow inputs declared; environment names: none declared |
| Outputs | No job outputs declared; artifact/cache actions: none declared |
| Permissions | contents: read; run-obs-aw-status.actions: read; run-obs-aw-status.contents: read; run-obs-aw-status.copilot-requests: write; run-obs-aw-status.id-token: write; run-obs-aw-status.issues: write; run-obs-aw-status.pull-requests: write |
| Secrets | BUILDKITE_API_TOKEN, BUILDKITE_LOGS_API_TOKEN |

## Operational notes

- Reusable workflow internals are hosted externally and are not available in this checkout.

## Sources

- Workflow: [`.github/workflows/trigger-obs-aw-status.yml`](../trigger-obs-aw-status.yml)
- External action/workflow reference: `elastic/oblt-aw/.github/workflows/obs-aw-event-status.yml@main` (implementation is outside this checkout).
