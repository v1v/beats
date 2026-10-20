# pipeline-scheduler

Buildkite pipeline definition for pipeline-scheduler.

## At a glance

| | |
| --- | --- |
| Platform | Buildkite |
| Source | [`.buildkite/pipeline-scheduler.yml`](../pipeline-scheduler.yml) |
| Trigger | The checked-in definition is an uploadable step pipeline; its schedule is not defined here. |

## Purpose

This intermediate pipeline runs the release-branch scheduler script and uploads the generated trigger steps. The checked-in comments explain that it provides an agent image unavailable to the upstream catalog pipeline configuration; the exact scheduler registration is external to this YAML.

## When it runs

- The checked-in definition is an uploadable step pipeline; its schedule is not defined here.
- `PIPELINES_TO_TRIGGER` and optional `EXCLUDE_BRANCHES` are consumed by [`pipeline-scheduler.py`](../pipeline-scheduler.py); the script downloads active release branches and uploads generated trigger steps. The scheduler invocation itself is configured outside this YAML.

## Execution flow

| Step | Condition | Depends on | Configured work |
| --- | --- | --- | --- |
| `:pipeline: Generate trigger steps for $PIPELINES_TO_TRIGGER` | `No step condition declared` | None declared | Runs `.buildkite/pipeline-scheduler.py`; agents: image=docker.elastic.co/ci-agent-images/platform-ingest…; useCustomGlobalHooks=true |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | `PIPELINES_TO_TRIGGER` (required) and `EXCLUDE_BRANCHES` (optional), read by the referenced scheduler script |
| Outputs | Artifact paths: none declared in this file; downstream pipelines: none declared |
| Permissions | Agent requirements are configured per step (see table). |
| Secrets | No secret values are documented; named secret access is configured in referenced plugins or external Buildkite settings |

## Operational notes

- The command generates `steps.yml` and uploads it to Buildkite. Generated branches depend on the remote active-branches JSON source; runtime-generated steps are not fully visible in this static YAML.

## Sources

- Pipeline definition: [`.buildkite/pipeline-scheduler.yml`](../pipeline-scheduler.yml)
- Referenced script: [`.buildkite/pipeline-scheduler.py`](../../.buildkite/pipeline-scheduler.py)
