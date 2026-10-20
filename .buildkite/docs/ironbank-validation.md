# ironbank-validation

Buildkite pipeline definition for ironbank-validation.

## At a glance

| | |
| --- | --- |
| Platform | Buildkite |
| Source | [`.buildkite/ironbank-validation.yml`](../ironbank-validation.yml) |
| Trigger | Not specified in this checked-in YAML; start conditions may be configured in Buildkite. |

## Purpose

This pipeline defines `:package: :test_tube: Package and Validate Iron Bank for x-pack/heartbeat` for `ironbank-validation`. The source does not state a separate motivation.

## When it runs

- Not specified in this checked-in YAML; start conditions may be configured in Buildkite.

## Execution flow

| Step | Condition | Depends on | Configured work |
| --- | --- | --- | --- |
| `:package: :test_tube: Package and Validate Iron Bank for x-pack/heartbeat` | `No step condition declared` | None declared | Runs `export BEAT_VERSION=$$(make get-version)`; `echo "~~~ Using BEAT_VERSION: $$BEAT_VERSION"`; agents: provider=gcp; image=${IMAGE_UBUNTU_X86_64}; machineType=n2-standard-8; artifacts: `x-pack/heartbeat/build/distributions/*`, `x-pack/heartbeat/build/heartbeat-ironbank-*-docker-build-context/*` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | Environment names referenced: ASDF_MAGE_VERSION, IMAGE_UBUNTU_X86_64 |
| Outputs | Artifact paths: `x-pack/heartbeat/build/distributions/*`, `x-pack/heartbeat/build/heartbeat-ironbank-*-docker-build-context/*`; downstream pipelines: none declared |
| Permissions | Agent requirements are configured per step (see table). |
| Secrets | No secret values are documented; named secret access is configured in referenced plugins or external Buildkite settings |

## Operational notes

- This page describes only the checked-in YAML; runtime Buildkite settings and externally defined plugins may add behavior not represented here.

## Sources

- Pipeline definition: [`.buildkite/ironbank-validation.yml`](../ironbank-validation.yml)
