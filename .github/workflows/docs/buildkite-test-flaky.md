# Buildkite Test Flaky

Workflow definition for **Buildkite Test Flaky**. This page documents the checked-in GitHub Actions configuration; it does not assert that a run succeeds.

## At a glance

| | |
| --- | --- |
| Platform | GitHub Actions |
| Source | [`.github/workflows/buildkite-test-flaky.yml`](../buildkite-test-flaky.yml) |
| Trigger | schedule (Runs on a schedule; see the workflow source for current timing); workflow_dispatch |

## Purpose

This workflow provides the `Buildkite Test Flaky` task by invoking `elastic/oblt-actions/github/create-token@v1`, `elastic/oblt-actions/buildkite/flaky-report@v1`. The source does not state a separate motivation.



## When it runs

- schedule (Runs on a schedule; see the workflow source for current timing)
- workflow_dispatch

Event and job filters are evaluated as shown in the workflow source.

## Execution flow

| Job | Runner | Condition | Depends on | Configured work |
| --- | --- | --- | --- | --- |
| `run` | `ubuntu-latest` | `Always when its event is eligible` | None declared | `elastic/oblt-actions/github/create-token@v1`; `elastic/oblt-actions/buildkite/flaky-report@v1` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | No workflow inputs declared; environment names: none declared |
| Outputs | No job outputs declared; artifact/cache actions: none declared |
| Permissions | contents: read; run.id-token: write |
| Secrets | BUILDKITE_TEST_ENGINE_TOKEN |

## Operational notes

- Exact schedule timing is intentionally omitted; consult the workflow source.

## Sources

- Workflow: [`.github/workflows/buildkite-test-flaky.yml`](../buildkite-test-flaky.yml)
- External action/workflow reference: `elastic/oblt-actions/buildkite/flaky-report@v1` (implementation is outside this checkout).
- External action/workflow reference: `elastic/oblt-actions/github/create-token@v1` (implementation is outside this checkout).
