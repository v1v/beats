# Generate release notes

Workflow definition for **Generate release notes**. This page documents the checked-in GitHub Actions configuration; it does not assert that a run succeeds.

## At a glance

| | |
| --- | --- |
| Platform | GitHub Actions |
| Source | [`.github/workflows/release-notes.yml`](../release-notes.yml) |
| Trigger | workflow_dispatch: inputs: bc_commit_sha (Commit SHA for the build candidate), version (Version to be released (for example, `9.2.1`)) |

## Purpose

This workflow acts as an event wrapper: it delegates `build-release-notes` to `elastic/elastic-agent-changelog-tool/.github/workflows/generate-release-notes.yml@main`. The external workflow implementation is not checked into this repository, and the source does not state a separate rationale for the wrapper.



## When it runs

- workflow_dispatch: inputs: bc_commit_sha (Commit SHA for the build candidate), version (Version to be released (for example, `9.2.1`))

Event and job filters are evaluated as shown in the workflow source.

## Execution flow

| Job | Runner | Condition | Depends on | Configured work |
| --- | --- | --- | --- | --- |
| `build-release-notes` | `Reusable workflow or not specified` | `Always when its event is eligible` | None declared | Calls reusable workflow `elastic/elastic-agent-changelog-tool/.github/workflows/generate-release-notes.yml@main` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | workflow_dispatch input `bc_commit_sha`; workflow_dispatch input `version`; environment names: none declared |
| Outputs | No job outputs declared; artifact/cache actions: none declared |
| Permissions | contents: read; build-release-notes.contents: read; build-release-notes.id-token: write |
| Secrets | No `secrets.*` references detected in this file |

## Operational notes

- Reusable workflow internals are hosted externally and are not available in this checkout.

## Sources

- Workflow: [`.github/workflows/release-notes.yml`](../release-notes.yml)
- External action/workflow reference: `elastic/elastic-agent-changelog-tool/.github/workflows/generate-release-notes.yml@main` (implementation is outside this checkout).
