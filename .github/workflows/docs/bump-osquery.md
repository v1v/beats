# bump-osquery

Workflow definition for **bump-osquery**. This page documents the checked-in GitHub Actions configuration; it does not assert that a run succeeds.

## At a glance

| | |
| --- | --- |
| Platform | GitHub Actions |
| Source | [`.github/workflows/bump-osquery.yml`](../bump-osquery.yml) |
| Trigger | workflow_dispatch; schedule (Runs on a schedule; see the workflow source for current timing) |

## Purpose

This workflow runs the local Updatecli configuration that tracks the latest osquery release and updates the bundled osquery version and checksums through the repository updater.


## When it runs

- workflow_dispatch
- schedule (Runs on a schedule; see the workflow source for current timing)

Event and job filters are evaluated as shown in the workflow source.

## Execution flow

| Job | Runner | Condition | Depends on | Configured work |
| --- | --- | --- | --- | --- |
| `bump` | `ubuntu-latest` | `Always when its event is eligible` | None declared | `actions/checkout@v7`; `actions/setup-go@v7`; `elastic/oblt-actions/updatecli/run@v1`; `elastic/oblt-actions/slack/send@v1` |

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

- Workflow: [`.github/workflows/bump-osquery.yml`](../bump-osquery.yml)
- Update helper referenced by the Updatecli configuration: [`x-pack/osquerybeat/scripts/update_osquery/main.go`](../../../x-pack/osquerybeat/scripts/update_osquery/main.go)
- Updatecli configuration: [`.github/workflows/updatecli.d/values.d/scm.yml`](../../../.github/workflows/updatecli.d/values.d/scm.yml)
- Updatecli configuration: [`.github/workflows/updatecli.d/bump-osquery.yml`](../../../.github/workflows/updatecli.d/bump-osquery.yml)
- External action/workflow reference: `actions/checkout@v7` (implementation is outside this checkout).
- External action/workflow reference: `actions/setup-go@v7` (implementation is outside this checkout).
- External action/workflow reference: `elastic/oblt-actions/slack/send@v1` (implementation is outside this checkout).
- External action/workflow reference: `elastic/oblt-actions/updatecli/run@v1` (implementation is outside this checkout).
