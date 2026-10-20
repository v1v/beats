# updatecli-compose

Workflow definition for **updatecli-compose**. This page documents the checked-in GitHub Actions configuration; it does not assert that a run succeeds.

## At a glance

| | |
| --- | --- |
| Platform | GitHub Actions |
| Source | [`.github/workflows/updatecli-compose.yml`](../updatecli-compose.yml) |
| Trigger | workflow_dispatch; schedule (Runs on a schedule; see the workflow source for current timing) |

## Purpose

This workflow runs Updatecli compose diff and apply through the referenced external action. It sends failure notifications to Slack; the exact runner implementation is external.


## When it runs

- workflow_dispatch
- schedule (Runs on a schedule; see the workflow source for current timing)

Event and job filters are evaluated as shown in the workflow source.

## Execution flow

| Job | Runner | Condition | Depends on | Configured work |
| --- | --- | --- | --- | --- |
| `compose` | `ubuntu-latest` | `Always when its event is eligible` | None declared | `actions/checkout@v7`; `docker/login-action`; `elastic/oblt-actions/updatecli/run-and-notify@v1`; `elastic/oblt-actions/updatecli/run-and-notify@v1` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | No workflow inputs declared; environment names: GITHUB_TOKEN |
| Outputs | No job outputs declared; artifact/cache actions: none declared |
| Permissions | contents: read; compose.contents: write; compose.packages: read; compose.pull-requests: write |
| Secrets | GITHUB_TOKEN, SLACK_BOT_TOKEN |

## Operational notes

- Exact schedule timing is intentionally omitted; consult the workflow source.

## Sources

- Workflow: [`.github/workflows/updatecli-compose.yml`](../updatecli-compose.yml)
- Related Updatecli configuration overview: [`.github/workflows/updatecli.d/README.md`](../../../.github/workflows/updatecli.d/README.md)
- External action/workflow reference: `actions/checkout@v7` (implementation is outside this checkout).
- External action/workflow reference: `docker/login-action` (implementation is outside this checkout).
- External action/workflow reference: `elastic/oblt-actions/updatecli/run-and-notify@v1` (implementation is outside this checkout).
