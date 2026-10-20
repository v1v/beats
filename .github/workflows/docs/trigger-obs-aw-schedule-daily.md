# Observability Agentic Workflow — Schedule (daily)

Workflow definition for **Observability Agentic Workflow — Schedule (daily)**. This page documents the checked-in GitHub Actions configuration; it does not assert that a run succeeds.

## At a glance

| | |
| --- | --- |
| Platform | GitHub Actions |
| Source | [`.github/workflows/trigger-obs-aw-schedule-daily.yml`](../trigger-obs-aw-schedule-daily.yml) |
| Trigger | schedule (Runs on a schedule; see the workflow source for current timing); workflow_dispatch |

## Purpose

This workflow acts as an event wrapper: it delegates `run-obs-aw-schedule-daily` to `elastic/oblt-aw/.github/workflows/obs-aw-event-schedule.yml@main`. The external workflow implementation is not checked into this repository, and the source does not state a separate rationale for the wrapper.



## When it runs

- schedule (Runs on a schedule; see the workflow source for current timing)
- workflow_dispatch

Event and job filters are evaluated as shown in the workflow source.

## Execution flow

| Job | Runner | Condition | Depends on | Configured work |
| --- | --- | --- | --- | --- |
| `run-obs-aw-schedule-daily` | `Reusable workflow or not specified` | `Always when its event is eligible` | None declared | Calls reusable workflow `elastic/oblt-aw/.github/workflows/obs-aw-event-schedule.yml@main` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | No workflow inputs declared; environment names: none declared |
| Outputs | No job outputs declared; artifact/cache actions: none declared |
| Permissions | contents: read; run-obs-aw-schedule-daily.actions: read; run-obs-aw-schedule-daily.contents: write; run-obs-aw-schedule-daily.copilot-requests: write; run-obs-aw-schedule-daily.id-token: write; run-obs-aw-schedule-daily.issues: write; run-obs-aw-schedule-daily.pull-requests: write |
| Secrets | No `secrets.*` references detected in this file |

## Operational notes

- Exact schedule timing is intentionally omitted; consult the workflow source.
- Reusable workflow internals are hosted externally and are not available in this checkout.

## Sources

- Workflow: [`.github/workflows/trigger-obs-aw-schedule-daily.yml`](../trigger-obs-aw-schedule-daily.yml)
- External action/workflow reference: `elastic/oblt-aw/.github/workflows/obs-aw-event-schedule.yml@main` (implementation is outside this checkout).
