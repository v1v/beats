# Backport to active branches

Workflow definition for **Backport to active branches**. This page documents the checked-in GitHub Actions configuration; it does not assert that a run succeeds.

## At a glance

| | |
| --- | --- |
| Platform | GitHub Actions |
| Source | [`.github/workflows/backport-active.yml`](../backport-active.yml) |
| Trigger | pull_request_target: types=closed, labeled; branches=main |

## Purpose

This workflow provides the `Backport to active branches` task by invoking `elastic/oblt-actions/github/backport-active@v1`. The source does not state a separate motivation.



## When it runs

- pull_request_target: types=closed, labeled; branches=main

Event and job filters are evaluated as shown in the workflow source.

## Execution flow

| Job | Runner | Condition | Depends on | Configured work |
| --- | --- | --- | --- | --- |
| `backport` | `ubuntu-latest` | `github.event.pull_request.merged == true && (   (github.event.action == 'closed' && contains(toJSON(github.event.pull_request.labels.*.name), 'backport-active-')) \|\|   (github.event.action == 'labeled' && startsWith(github.event.label.name, 'backport-active-')) )` | None declared | `elastic/oblt-actions/github/backport-active@v1` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | No workflow inputs declared; environment names: none declared |
| Outputs | No job outputs declared; artifact/cache actions: none declared |
| Permissions | pull-requests: write; contents: read |
| Secrets | No `secrets.*` references detected in this file |

## Operational notes

- No additional retry, timeout, or deployment behavior is declared in the checked-in workflow.

## Sources

- Workflow: [`.github/workflows/backport-active.yml`](../backport-active.yml)
- External action/workflow reference: `elastic/oblt-actions/github/backport-active@v1` (implementation is outside this checkout).
