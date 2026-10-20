# pre-commit

Workflow definition for **pre-commit**. This page documents the checked-in GitHub Actions configuration; it does not assert that a run succeeds.

## At a glance

| | |
| --- | --- |
| Platform | GitHub Actions |
| Source | [`.github/workflows/pre-commit.yml`](../pre-commit.yml) |
| Trigger | pull_request; push: branches=main, 7.1*, 8.*, 9.* |

## Purpose

This workflow provides the `pre-commit` task by invoking `elastic/oblt-actions/pre-commit@v1`. The source does not state a separate motivation.



## When it runs

- pull_request
- push: branches=main, 7.1*, 8.*, 9.*

Event and job filters are evaluated as shown in the workflow source.

## Execution flow

| Job | Runner | Condition | Depends on | Configured work |
| --- | --- | --- | --- | --- |
| `pre-commit` | `ubuntu-latest` | `Always when its event is eligible` | None declared | `elastic/oblt-actions/pre-commit@v1` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | No workflow inputs declared; environment names: none declared |
| Outputs | No job outputs declared; artifact/cache actions: none declared |
| Permissions | contents: read |
| Secrets | No `secrets.*` references detected in this file |

## Operational notes

- No additional retry, timeout, or deployment behavior is declared in the checked-in workflow.

## Sources

- Workflow: [`.github/workflows/pre-commit.yml`](../pre-commit.yml)
- External action/workflow reference: `elastic/oblt-actions/pre-commit@v1` (implementation is outside this checkout).
