# mergify backport labels copier

Workflow definition for **mergify backport labels copier**. This page documents the checked-in GitHub Actions configuration; it does not assert that a run succeeds.

## At a glance

| | |
| --- | --- |
| Platform | GitHub Actions |
| Source | [`.github/workflows/mergify-labels-copier.yml`](../mergify-labels-copier.yml) |
| Trigger | pull_request: types=opened |

## Purpose

This workflow provides the `mergify backport labels copier` task by invoking `elastic/oblt-actions/mergify/labels-copier@v1`. The source does not state a separate motivation.



## When it runs

- pull_request: types=opened

Event and job filters are evaluated as shown in the workflow source.

## Execution flow

| Job | Runner | Condition | Depends on | Configured work |
| --- | --- | --- | --- | --- |
| `mergify-backport-labels-copier` | `ubuntu-latest` | `startsWith(github.head_ref, 'mergify/bp/')` | None declared | `elastic/oblt-actions/mergify/labels-copier@v1` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | No workflow inputs declared; environment names: none declared |
| Outputs | No job outputs declared; artifact/cache actions: none declared |
| Permissions | contents: read; mergify-backport-labels-copier.pull-requests: write; mergify-backport-labels-copier.repository-projects: read |
| Secrets | No `secrets.*` references detected in this file |

## Operational notes

- No additional retry, timeout, or deployment behavior is declared in the checked-in workflow.

## Sources

- Workflow: [`.github/workflows/mergify-labels-copier.yml`](../mergify-labels-copier.yml)
- External action/workflow reference: `elastic/oblt-actions/mergify/labels-copier@v1` (implementation is outside this checkout).
