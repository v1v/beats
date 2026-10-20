# aws-tests-pipeline

Buildkite pipeline definition for aws-tests-pipeline.

## At a glance

| | |
| --- | --- |
| Platform | Buildkite |
| Source | [`.buildkite/aws-tests-pipeline.yml`](../aws-tests-pipeline.yml) |
| Trigger | Not specified in this checked-in YAML; start conditions may be configured in Buildkite. |

## Purpose

This pipeline defines `:ubuntu: x-pack/filebeat: AWS Tests`, `:ubuntu: x-pack/metricbeat: AWS Tests` for `aws-tests-pipeline`. The source does not state a separate motivation.

## When it runs

- Not specified in this checked-in YAML; start conditions may be configured in Buildkite.

## Execution flow

| Step | Condition | Depends on | Configured work |
| --- | --- | --- | --- |
| `:ubuntu: x-pack/filebeat: AWS Tests` | `No step condition declared` | None declared | Runs `.buildkite/scripts/initCloudEnv.sh`; agents: provider=aws; image=${AWS_IMAGE_UBUNTU_ARM_64}; instanceType=${AWS_ARM_INSTANCE_TYPE}; plugins: `elastic/oblt-aws-auth#v0.2.0`; artifacts: `x-pack/filebeat/build/*.xml`, `x-pack/filebeat/build/*.json`; skipped: skipping:  elastic/ingest-dev#3467 |
| `:ubuntu: x-pack/metricbeat: AWS Tests` | `No step condition declared` | None declared | Runs `.buildkite/scripts/initCloudEnv.sh`; agents: provider=aws; image=${AWS_IMAGE_UBUNTU_ARM_64}; instanceType=${AWS_ARM_INSTANCE_TYPE}; plugins: `elastic/oblt-aws-auth#v0.2.0`; artifacts: `x-pack/metricbeat/build/*.xml`, `x-pack/metricbeat/build/*.json` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | Environment names referenced: ASDF_MAGE_VERSION, ASDF_TERRAFORM_VERSION, AWS_ARM_INSTANCE_TYPE, AWS_IMAGE_UBUNTU_ARM_64, MODULE, MODULE_DIR, RACE_DETECTOR, TEST_COVERAGE, TEST_TAGS |
| Outputs | Artifact paths: `x-pack/filebeat/build/*.json`, `x-pack/filebeat/build/*.xml`, `x-pack/metricbeat/build/*.json`, `x-pack/metricbeat/build/*.xml`; downstream pipelines: none declared |
| Permissions | Agent requirements are configured per step (see table).; Plugins are configured per step (see table). |
| Secrets | No secret values are documented; named secret access is configured in referenced plugins or external Buildkite settings |

## Operational notes

- This page describes only the checked-in YAML; runtime Buildkite settings and externally defined plugins may add behavior not represented here.

## Sources

- Pipeline definition: [`.buildkite/aws-tests-pipeline.yml`](../aws-tests-pipeline.yml)
- Referenced script: [`.buildkite/scripts/initCloudEnv.sh`](../../.buildkite/scripts/initCloudEnv.sh)
- Referenced local dependency: [`.buildkite/deploy/docker/docker-compose.yml`](../../.buildkite/deploy/docker/docker-compose.yml)
