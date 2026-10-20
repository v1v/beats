# catalog-info

Workflow definition for **catalog-info**. This page documents the checked-in GitHub Actions configuration; it does not assert that a run succeeds.

## At a glance

| | |
| --- | --- |
| Platform | GitHub Actions |
| Source | [`.github/workflows/catalog-info.yml`](../catalog-info.yml) |
| Trigger | pull_request: branches=main; paths=catalog-info.yaml |

## Purpose

This workflow provides the `catalog-info` task by invoking `actions/checkout@v7`, `elastic/oblt-actions/elastic/validate-catalog@v1`. The source does not state a separate motivation.



## When it runs

- pull_request: branches=main; paths=catalog-info.yaml

Event and job filters are evaluated as shown in the workflow source.

## Execution flow

| Job | Runner | Condition | Depends on | Configured work |
| --- | --- | --- | --- | --- |
| `validate` | `ubuntu-latest` | `Always when its event is eligible` | None declared | `actions/checkout@v7`; `elastic/oblt-actions/elastic/validate-catalog@v1` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | No workflow inputs declared; environment names: none declared |
| Outputs | No job outputs declared; artifact/cache actions: none declared |
| Permissions | contents: read; validate.contents: read; validate.packages: read |
| Secrets | No `secrets.*` references detected in this file |

## Operational notes

- No additional retry, timeout, or deployment behavior is declared in the checked-in workflow.

## Sources

- Workflow: [`.github/workflows/catalog-info.yml`](../catalog-info.yml)
- External action/workflow reference: `actions/checkout@v7` (implementation is outside this checkout).
- External action/workflow reference: `elastic/oblt-actions/elastic/validate-catalog@v1` (implementation is outside this checkout).
