# Duplicate Issue Detector

Workflow definition for **Duplicate Issue Detector**. This page documents the checked-in GitHub Actions configuration; it does not assert that a run succeeds.

## At a glance

| | |
| --- | --- |
| Platform | GitHub Actions |
| Source | [`.github/workflows/trigger-duplicate-issue-detector.yml`](../trigger-duplicate-issue-detector.yml) |
| Trigger | issues: types=opened |

## Purpose

This workflow acts as an event wrapper: it delegates `run` to `elastic/ai-github-actions/.github/workflows/gh-aw-duplicate-issue-detector.lock.yml@v0`. The external workflow implementation is not checked into this repository, and the source does not state a separate rationale for the wrapper.



## When it runs

- issues: types=opened

Event and job filters are evaluated as shown in the workflow source.

## Execution flow

| Job | Runner | Condition | Depends on | Configured work |
| --- | --- | --- | --- | --- |
| `run` | `Reusable workflow or not specified` | `contains(fromJSON('["OWNER","MEMBER","COLLABORATOR"]'), github.event.issue.author_association)` | None declared | Calls reusable workflow `elastic/ai-github-actions/.github/workflows/gh-aw-duplicate-issue-detector.lock.yml@v0` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | No workflow inputs declared; environment names: none declared |
| Outputs | No job outputs declared; artifact/cache actions: none declared |
| Permissions | actions: read; contents: read; copilot-requests: write; discussions: write; issues: write; pull-requests: read |
| Secrets | No `secrets.*` references detected in this file |

## Operational notes

- Reusable workflow internals are hosted externally and are not available in this checkout.

## Sources

- Workflow: [`.github/workflows/trigger-duplicate-issue-detector.yml`](../trigger-duplicate-issue-detector.yml)
- External action/workflow reference: `elastic/ai-github-actions/.github/workflows/gh-aw-duplicate-issue-detector.lock.yml@v0` (implementation is outside this checkout).
