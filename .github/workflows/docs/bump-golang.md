# bump-golang

Workflow definition for **bump-golang**. This page documents the checked-in GitHub Actions configuration; it does not assert that a run succeeds.

## At a glance

| | |
| --- | --- |
| Platform | GitHub Actions |
| Source | [`.github/workflows/bump-golang.yml`](../bump-golang.yml) |
| Trigger | workflow_dispatch; schedule (Runs on a schedule; see the workflow source for current timing) |

## Purpose

This workflow runs the local Updatecli configuration that tracks the latest Go release for the current minor line and updates the module version, `.go-version`, docs, development container, and Go-based Dockerfiles.


## When it runs

- workflow_dispatch
- schedule (Runs on a schedule; see the workflow source for current timing)

Event and job filters are evaluated as shown in the workflow source.

## Execution flow

| Job | Runner | Condition | Depends on | Configured work |
| --- | --- | --- | --- | --- |
| `bump` | `ubuntu-latest` | `Always when its event is eligible` | None declared | `actions/checkout@v7`; `elastic/oblt-actions/updatecli/run-and-notify@v1` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | No workflow inputs declared; environment names: GITHUB_TOKEN |
| Outputs | No job outputs declared; artifact/cache actions: none declared |
| Permissions | contents: read; bump.contents: write; bump.pull-requests: write |
| Secrets | GITHUB_TOKEN, SLACK_BOT_TOKEN |

## Operational notes

- Exact schedule timing is intentionally omitted; consult the workflow source.

## Sources

- Workflow: [`.github/workflows/bump-golang.yml`](../bump-golang.yml)
- Updatecli configuration: [`.github/workflows/updatecli.d/values.d/scm.yml`](../../../.github/workflows/updatecli.d/values.d/scm.yml)
- Updatecli configuration: [`.github/workflows/updatecli.d/bump-golang.yml`](../../../.github/workflows/updatecli.d/bump-golang.yml)
- External action/workflow reference: `actions/checkout@v7` (implementation is outside this checkout).
- External action/workflow reference: `elastic/oblt-actions/updatecli/run-and-notify@v1` (implementation is outside this checkout).
