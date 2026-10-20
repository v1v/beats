# Code Simplifier

Workflow definition for **Code Simplifier**. This page documents the checked-in GitHub Actions configuration; it does not assert that a run succeeds.

## At a glance

| | |
| --- | --- |
| Platform | GitHub Actions |
| Source | [`.github/workflows/trigger-code-simplifier.yml`](../trigger-code-simplifier.yml) |
| Trigger | schedule (Runs on a schedule; see the workflow source for current timing); workflow_dispatch |

## Purpose

This workflow acts as an event wrapper: it delegates `run` to `elastic/ai-github-actions/.github/workflows/gh-aw-code-simplifier.lock.yml@v0`. The external workflow implementation is not checked into this repository, and the source does not state a separate rationale for the wrapper.



## When it runs

- schedule (Runs on a schedule; see the workflow source for current timing)
- workflow_dispatch

Event and job filters are evaluated as shown in the workflow source.

## Execution flow

| Job | Runner | Condition | Depends on | Configured work |
| --- | --- | --- | --- | --- |
| `run` | `Reusable workflow or not specified` | `Always when its event is eligible` | None declared | Calls reusable workflow `elastic/ai-github-actions/.github/workflows/gh-aw-code-simplifier.lock.yml@v0` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | No workflow inputs declared; environment names: none declared |
| Outputs | No job outputs declared; artifact/cache actions: none declared |
| Permissions | actions: read; contents: write; copilot-requests: write; issues: write; pull-requests: write |
| Secrets | No `secrets.*` references detected in this file |

## Operational notes

- Exact schedule timing is intentionally omitted; consult the workflow source.
- Reusable workflow internals are hosted externally and are not available in this checkout.

## Sources

- Workflow: [`.github/workflows/trigger-code-simplifier.yml`](../trigger-code-simplifier.yml)
- External action/workflow reference: `elastic/ai-github-actions/.github/workflows/gh-aw-code-simplifier.lock.yml@v0` (implementation is outside this checkout).
