# xpack-filebeat-macos-unit-tests

Workflow definition for **xpack-filebeat-macos-unit-tests**. This page documents the checked-in GitHub Actions configuration; it does not assert that a run succeeds.

## At a glance

| | |
| --- | --- |
| Platform | GitHub Actions |
| Source | [`.github/workflows/xpack-filebeat-macos-unit-tests.yml`](../xpack-filebeat-macos-unit-tests.yml) |
| Trigger | workflow_dispatch; pull_request: paths=filebeat/**, x-pack/filebeat/**, x-pack/libbeat/**, go.mod, pytest.ini, dev-tools/**, libbeat/**, .github/workflows/xpack-filebeat-macos-unit-tests.yml, .github/actions/unit-test/action.yml; push: branches=main; paths=filebeat/**, x-pack/filebeat/**, x-pack/libbeat/**, go.mod, pytest.ini, dev-tools/**, libbeat/**, .github/workflows/xpack-filebeat-macos-unit-tests.yml, .github/actions/unit-test/action.yml |

## Purpose

This workflow runs the `x-pack/filebeat` macOS build and unit-test job for the configured pull-request and push path filters, and also supports manual dispatch. It delegates build and unit-test commands to the shared local composite action. The source does not state a separate motivation.

## When it runs

- workflow_dispatch
- pull_request: paths=filebeat/**, x-pack/filebeat/**, x-pack/libbeat/**, go.mod, pytest.ini, dev-tools/**, libbeat/**, .github/workflows/xpack-filebeat-macos-unit-tests.yml, .github/actions/unit-test/action.yml
- push: branches=main; paths=filebeat/**, x-pack/filebeat/**, x-pack/libbeat/**, go.mod, pytest.ini, dev-tools/**, libbeat/**, .github/workflows/xpack-filebeat-macos-unit-tests.yml, .github/actions/unit-test/action.yml

Event and job filters are evaluated as shown in the workflow source.

## Execution flow

| Job | Runner | Condition | Depends on | Configured work |
| --- | --- | --- | --- | --- |
| `unit-tests` | `${{ matrix.os }}` | `Always when its event is eligible` | None declared | `actions/checkout@v7`; `./.github/actions/unit-test` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | local action input `beat` = `${{ env.BEAT }}`; environment names: BEAT |
| Outputs | No job outputs declared; artifact/cache actions: none declared |
| Permissions | contents: read |
| Secrets | No `secrets.*` references detected in this file |

## Operational notes

- These jobs use the shared local composite action in `.github/actions/unit-test/action.yml`; it installs the Go and Python toolchains and runs the beat build and unit tests for the configured `beat` input.
- unit-tests timeout: 60 minutes
- unit-tests uses a matrix/strategy declared in the source

## Sources

- Workflow: [`.github/workflows/xpack-filebeat-macos-unit-tests.yml`](../xpack-filebeat-macos-unit-tests.yml)
- Referenced local source: [`.github/actions/unit-test/action.yml`](../../../.github/actions/unit-test/action.yml)
- External action/workflow reference: `actions/checkout@v7` (implementation is outside this checkout).
