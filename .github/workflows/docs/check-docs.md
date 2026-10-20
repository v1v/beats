# check-docs

Workflow definition for **check-docs**. This page documents the checked-in GitHub Actions configuration; it does not assert that a run succeeds.

## At a glance

| | |
| --- | --- |
| Platform | GitHub Actions |
| Source | [`.github/workflows/check-docs.yml`](../check-docs.yml) |
| Trigger | pull_request: paths=.github/workflows/check-docs.yml, **/*.asciidoc, **/*.md, deploy/kubernetes/*-kubernetes.yaml |

## Purpose

This workflow provides the `check-docs` CI task: it prepares or invokes `actions/checkout@v7`, `actions/setup-go@v7`, `actions/setup-python@v7` and then runs `sudo apt-get update`, `sudo apt-get install -y libpcap-dev`, `sudo apt-get install -y libsystemd-dev`, `sudo apt-get install -y librpm-dev`. The checked-in source does not state a separate motivation.



## When it runs

- pull_request: paths=.github/workflows/check-docs.yml, **/*.asciidoc, **/*.md, deploy/kubernetes/*-kubernetes.yaml

Event and job filters are evaluated as shown in the workflow source.

## Execution flow

| Job | Runner | Condition | Depends on | Configured work |
| --- | --- | --- | --- | --- |
| `check` | `ubuntu-latest` | `Always when its event is eligible` | None declared | `actions/checkout@v7`; `actions/setup-go@v7`; Update package lists — runs `sudo apt-get update`; Install libpcap-dev — runs `sudo apt-get install -y libpcap-dev`; Install libsystemd-dev — runs `sudo apt-get install -y libsystemd-dev`; Install librpm-dev — runs `sudo apt-get install -y librpm-dev`; Fix Code is not compatible with Python 3.12 (`actions/setup-python@v7`); Run check — runs `make check` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | No workflow inputs declared; environment names: none declared |
| Outputs | No job outputs declared; artifact/cache actions: none declared |
| Permissions | contents: read; pull-requests: read |
| Secrets | No `secrets.*` references detected in this file |

## Operational notes

- No additional retry, timeout, or deployment behavior is declared in the checked-in workflow.

## Sources

- Workflow: [`.github/workflows/check-docs.yml`](../check-docs.yml)
- External action/workflow reference: `actions/checkout@v7` (implementation is outside this checkout).
- External action/workflow reference: `actions/setup-go@v7` (implementation is outside this checkout).
- External action/workflow reference: `actions/setup-python@v7` (implementation is outside this checkout).
