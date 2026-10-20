# Issue Fixer

Workflow definition for **Issue Fixer**. This page documents the checked-in GitHub Actions configuration; it does not assert that a run succeeds.

## At a glance

| | |
| --- | --- |
| Platform | GitHub Actions |
| Source | [`.github/workflows/trigger-issue-fixer.yml`](../trigger-issue-fixer.yml) |
| Trigger | issues: types=labeled |

## Purpose

This workflow acts as an event wrapper: it delegates `flaky-test-fixer` to `elastic/ai-github-actions/.github/workflows/gh-aw-issue-fixer.lock.yml@v0`. The external workflow implementation is not checked into this repository, and the source does not state a separate rationale for the wrapper.



## When it runs

- issues: types=labeled

Event and job filters are evaluated as shown in the workflow source.

## Execution flow

| Job | Runner | Condition | Depends on | Configured work |
| --- | --- | --- | --- | --- |
| `flaky-test-fixer` | `Reusable workflow or not specified` | `contains(github.event.issue.labels.*.name, 'flaky-test') && contains(github.event.issue.labels.*.name, 'triaged')` | None declared | Calls reusable workflow `elastic/ai-github-actions/.github/workflows/gh-aw-issue-fixer.lock.yml@v0` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | No workflow inputs declared; environment names: none declared |
| Outputs | No job outputs declared; artifact/cache actions: none declared |
| Permissions | actions: read; contents: write; copilot-requests: write; discussions: write; issues: write; pull-requests: write |
| Secrets | No `secrets.*` references detected in this file |

## Operational notes

- Reusable workflow internals are hosted externally and are not available in this checkout.

## Sources

- Workflow: [`.github/workflows/trigger-issue-fixer.yml`](../trigger-issue-fixer.yml)
- External action/workflow reference: `elastic/ai-github-actions/.github/workflows/gh-aw-issue-fixer.lock.yml@v0` (implementation is outside this checkout).
