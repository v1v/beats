# pipeline.xpack.elastic-agent

Buildkite pipeline definition for pipeline.xpack.elastic-agent.

## At a glance

| | |
| --- | --- |
| Platform | Buildkite |
| Source | [`.buildkite/x-pack/elastic-agent/pipeline.xpack.elastic-agent.yml`](../x-pack/elastic-agent/pipeline.xpack.elastic-agent.yml) |
| Trigger | Not specified in this checked-in YAML; start conditions may be configured in Buildkite. |

## Purpose

This pipeline defines `Example test` for `pipeline.xpack.elastic-agent`. The source does not state a separate motivation.

## When it runs

- Not specified in this checked-in YAML; start conditions may be configured in Buildkite.

## Execution flow

| Step | Condition | Depends on | Configured work |
| --- | --- | --- | --- |
| `Example test` | `No step condition declared` | None declared | Runs `echo "Hello!"` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | Environment names referenced: none detected |
| Outputs | Artifact paths: none declared in this file; downstream pipelines: none declared |
| Permissions | Buildkite agent/plugin access is not detailed beyond this YAML |
| Secrets | No secret values are documented; named secret access is configured in referenced plugins or external Buildkite settings |

## Operational notes

- This page describes only the checked-in YAML; runtime Buildkite settings and externally defined plugins may add behavior not represented here.

## Sources

- Pipeline definition: [`.buildkite/x-pack/elastic-agent/pipeline.xpack.elastic-agent.yml`](../x-pack/elastic-agent/pipeline.xpack.elastic-agent.yml)
