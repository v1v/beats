# Observability Agentic Workflow — Issue Comment

Workflow definition for **Observability Agentic Workflow — Issue Comment**. This page documents the checked-in GitHub Actions configuration; it does not assert that a run succeeds.

## At a glance

| | |
| --- | --- |
| Platform | GitHub Actions |
| Source | [`.github/workflows/trigger-obs-aw-issue-comment.yml`](../trigger-obs-aw-issue-comment.yml) |
| Trigger | issue_comment: types=created |

## Purpose

This workflow acts as an event wrapper: it delegates `run-obs-aw-issue-comment` to `elastic/oblt-aw/.github/workflows/obs-aw-event-issue-comment.yml@main`. The external workflow implementation is not checked into this repository, and the source does not state a separate rationale for the wrapper.



## When it runs

- issue_comment: types=created

Event and job filters are evaluated as shown in the workflow source.

## Execution flow

| Job | Runner | Condition | Depends on | Configured work |
| --- | --- | --- | --- | --- |
| `run-obs-aw-issue-comment` | `Reusable workflow or not specified` | `Always when its event is eligible` | None declared | Calls reusable workflow `elastic/oblt-aw/.github/workflows/obs-aw-event-issue-comment.yml@main` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | No workflow inputs declared; environment names: none declared |
| Outputs | No job outputs declared; artifact/cache actions: none declared |
| Permissions | contents: read; run-obs-aw-issue-comment.actions: read; run-obs-aw-issue-comment.contents: write; run-obs-aw-issue-comment.copilot-requests: write; run-obs-aw-issue-comment.discussions: write; run-obs-aw-issue-comment.id-token: write; run-obs-aw-issue-comment.issues: write; run-obs-aw-issue-comment.pull-requests: write |
| Secrets | No `secrets.*` references detected in this file |

## Operational notes

- Reusable workflow internals are hosted externally and are not available in this checkout.

## Sources

- Workflow: [`.github/workflows/trigger-obs-aw-issue-comment.yml`](../trigger-obs-aw-issue-comment.yml)
- External action/workflow reference: `elastic/oblt-aw/.github/workflows/obs-aw-event-issue-comment.yml@main` (implementation is outside this checkout).
