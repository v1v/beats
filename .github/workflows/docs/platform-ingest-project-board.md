# Add issue to Ingest project

Workflow definition for **Add issue to Ingest project**. This page documents the checked-in GitHub Actions configuration; it does not assert that a run succeeds.

## At a glance

| | |
| --- | --- |
| Platform | GitHub Actions |
| Source | [`.github/workflows/platform-ingest-project-board.yml`](../platform-ingest-project-board.yml) |
| Trigger | issues: types=labeled |

## Purpose

This workflow provides the `Add issue to Ingest project` task by invoking `octokit/graphql-action@v2.x`. The source does not state a separate motivation.



## When it runs

- issues: types=labeled

Event and job filters are evaluated as shown in the workflow source.

## Execution flow

| Job | Runner | Condition | Depends on | Configured work |
| --- | --- | --- | --- | --- |
| `add_to_ingest_project` | `ubuntu-latest` | `Always when its event is eligible` | None declared | `octokit/graphql-action@v2.x`; `octokit/graphql-action@v2.x` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | No workflow inputs declared; environment names: AREA_FIELD_ID, CONTROL_PLANE_LABEL, DATA_PLANE_LABEL, DATA_PLANE_OPTION_ID, ELASTIC_AGENT_LABEL, GITHUB_TOKEN, INGEST_PROJECT_ID |
| Outputs | No job outputs declared; artifact/cache actions: none declared |
| Permissions | contents: read |
| Secrets | PROJECT_ASSIGNER_TOKEN |

## Operational notes

- No additional retry, timeout, or deployment behavior is declared in the checked-in workflow.

## Sources

- Workflow: [`.github/workflows/platform-ingest-project-board.yml`](../platform-ingest-project-board.yml)
- External action/workflow reference: `octokit/graphql-action@v2.x` (implementation is outside this checkout).
