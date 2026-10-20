# fragment-in-pr

Workflow definition for **fragment-in-pr**. This page documents the checked-in GitHub Actions configuration; it does not assert that a run succeeds.

## At a glance

| | |
| --- | --- |
| Platform | GitHub Actions |
| Source | [`.github/workflows/fragment-in-pr.yml`](../fragment-in-pr.yml) |
| Trigger | pull_request: types=labeled, unlabeled, opened, reopened, synchronize |

## Purpose

This workflow provides the `fragment-in-pr` CI task: it prepares or invokes `actions/checkout@v7`, `actions/setup-go@v7` and then runs `GOBIN=$PWD/bin go install github.com/elastic/elastic-agent-changelog-tool@latest`, `./bin/elastic-agent-changelog-tool pr-has-fragment --repo ${{ github.event.repository.name }} ${{github.event.numbe`, `gv=$(awk '$0 ~ /[0-9]+\.[0-9]+\.[0-9]+/ {print $4}' libbeat/version/version.go)`, `echo "beats-version=$gv" >> $GITHUB_OUTPUT`. The checked-in source does not state a separate motivation.



## When it runs

- pull_request: types=labeled, unlabeled, opened, reopened, synchronize

Event and job filters are evaluated as shown in the workflow source.

## Execution flow

| Job | Runner | Condition | Depends on | Configured work |
| --- | --- | --- | --- | --- |
| `fragments` | `ubuntu-latest` | `!contains(github.event.pull_request.labels.*.name, 'skip-changelog') && !contains(github.event.pull_request.labels.*.name, 'backport')` | None declared | `actions/checkout@v7`; `actions/setup-go@v7`; check pr-has-fragment — runs `GOBIN=$PWD/bin go install github.com/elastic/elastic-agent-changelog-tool@latest`; `./bin/elastic-agent-changelog-tool pr-has-fragment --repo ${{ github.event.repository.name }} ${{github.event.numbe`; get beats-agent version — runs `gv=$(awk '$0 ~ /[0-9]+\.[0-9]+\.[0-9]+/ {print $4}' libbeat/version/version.go)`; `echo "beats-version=$gv" >> $GITHUB_OUTPUT`; check fragment-syntax — runs `./bin/elastic-agent-changelog-tool build --version ${{steps.beats-version.outputs.beats-version}} --owner elastic -` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | No workflow inputs declared; environment names: none declared |
| Outputs | No job outputs declared; artifact/cache actions: none declared |
| Permissions | contents: read |
| Secrets | No `secrets.*` references detected in this file |

## Operational notes

- No additional retry, timeout, or deployment behavior is declared in the checked-in workflow.

## Sources

- Workflow: [`.github/workflows/fragment-in-pr.yml`](../fragment-in-pr.yml)
- External action/workflow reference: `actions/checkout@v7` (implementation is outside this checkout).
- External action/workflow reference: `actions/setup-go@v7` (implementation is outside this checkout).
