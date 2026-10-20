# github-commands-comment

Workflow definition for **github-commands-comment**. This page documents the checked-in GitHub Actions configuration; it does not assert that a run succeeds.

## At a glance

| | |
| --- | --- |
| Platform | GitHub Actions |
| Source | [`.github/workflows/github-commands-comment.yml`](../github-commands-comment.yml) |
| Trigger | pull_request_target: types=opened |

## Purpose

This workflow provides the `github-commands-comment` task by invoking `elastic/oblt-actions/elastic/github-commands@v1`. The source does not state a separate motivation.



## When it runs

- pull_request_target: types=opened

Event and job filters are evaluated as shown in the workflow source.

## Execution flow

| Job | Runner | Condition | Depends on | Configured work |
| --- | --- | --- | --- | --- |
| `comment` | `ubuntu-latest` | `Always when its event is eligible` | None declared | `elastic/oblt-actions/elastic/github-commands@v1` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | No workflow inputs declared; environment names: none declared |
| Outputs | No job outputs declared; artifact/cache actions: none declared |
| Permissions | contents: read; comment.pull-requests: write |
| Secrets | No `secrets.*` references detected in this file |

## Operational notes

- No additional retry, timeout, or deployment behavior is declared in the checked-in workflow.

## Sources

- Workflow: [`.github/workflows/github-commands-comment.yml`](../github-commands-comment.yml)
- External action/workflow reference: `elastic/oblt-actions/elastic/github-commands@v1` (implementation is outside this checkout).
