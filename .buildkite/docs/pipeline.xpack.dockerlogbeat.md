# beats-xpack-dockerlogbeat

Buildkite pipeline definition for beats-xpack-dockerlogbeat.

## At a glance

| | |
| --- | --- |
| Platform | Buildkite |
| Source | [`.buildkite/x-pack/pipeline.xpack.dockerlogbeat.yml`](../x-pack/pipeline.xpack.dockerlogbeat.yml) |
| Trigger | Not specified in this checked-in YAML; start conditions may be configured in Buildkite. |

## Purpose

This pipeline runs the configured `beats-xpack-dockerlogbeat` work, organized into `Check/Update`, `Mandatory Tests`, `Packaging`. The source does not state a separate motivation.

## When it runs

- Not specified in this checked-in YAML; start conditions may be configured in Buildkite.

## Execution flow

```mermaid
flowchart TD
  x_pack_dockerlogbeat_check_update --> Wait_barrier
  xpack_dockerlogbeat_mandatory_tests --> Wait_barrier
```

| Step | Condition | Depends on | Configured work |
| --- | --- | --- | --- |
| `Check/Update / x-pack/dockerlogbeat: Run check/update` | `No step condition declared` | None declared | Runs `make -C x-pack/dockerlogbeat check update`; `make check-no-changes`; agents: image=${IMAGE_BEATS_WITH_HOOKS_LATEST}; cpu=4000m; memory=8Gi; useCustomGlobalHooks=true |
| `Check/Update / x-pack/dockerlogbeat: Run pre-commit` | `No step condition declared` | None declared | Runs `pre-commit run --all-files`; agents: image=${IMAGE_BEATS_WITH_HOOKS_LATEST}; memory=2Gi; useCustomGlobalHooks=true |
| `Wait barrier` | `build.env("BUILDKITE_PULL_REQUEST") != "false"` | x-pack-dockerlogbeat-check-update | Waits at a Buildkite barrier |
| `Mandatory Tests / :ubuntu: x-pack/dockerlogbeat: Ubuntu x86_64 Unit Tests` | `No step condition declared` | None declared | Runs `.buildkite/scripts/qemu.sh`; agents: provider=gcp; image=${IMAGE_UBUNTU_X86_64}; machineType=${GCP_DEFAULT_MACHINE_TYPE}; artifacts: `x-pack/dockerlogbeat/build/*.xml`, `x-pack/dockerlogbeat/build/*.json` |
| `Mandatory Tests / :ubuntu: x-pack/dockerlogbeat: Go Integration Tests` | `No step condition declared` | None declared | Runs `cd x-pack/dockerlogbeat`; `mage goIntegTest`; agents: provider=gcp; image=${IMAGE_UBUNTU_X86_64}; machineType=${GCP_DEFAULT_MACHINE_TYPE}; artifacts: `x-pack/dockerlogbeat/build/*.xml`, `x-pack/dockerlogbeat/build/*.json`, `x-pack/dockerlogbeat/build/integration-tests/*`, `x-pack/dockerlogbeat/build/integration-tests/Test*/*`, `x-pack/dockerlogbeat/build/integration-tests/Test*/data/**/*` |
| `Mandatory Tests / :windows: x-pack/dockerlogbeat: Go Integration Tests` | `No step condition declared` | None declared | Runs `cd x-pack/dockerlogbeat`; `mage goIntegTest`; agents: provider=gcp; image=${IMAGE_WIN_2022}; machineType=${GCP_DEFAULT_MACHINE_TYPE}; artifacts: `x-pack/dockerlogbeat/build/*.xml`, `x-pack/dockerlogbeat/build/*.json`, `x-pack/dockerlogbeat/build/integration-tests/*`, `x-pack/dockerlogbeat/build/integration-tests/Test*/*`, `x-pack/dockerlogbeat/build/integration-tests/Test*/data/**/*` |
| `Wait barrier` | `build.env("BUILDKITE_PULL_REQUEST") != "false"` | xpack-dockerlogbeat-mandatory-tests | Waits at a Buildkite barrier |
| `Packaging / :ubuntu: x-pack/dockerlogbeat: Packaging Linux` | `build.pull_request.id != null` | None declared | Runs `.buildkite/scripts/packaging/packaging.sh`; agents: provider=gcp; image=${IMAGE_UBUNTU_X86_64}; machineType=${GCP_HI_PERF_MACHINE_TYPE}; plugins: `elastic/vault-docker-login#v0.8.0`, `elastic/vault-docker-login#v0.8.0` |
| `Packaging / :ubuntu: x-pack/dockerlogbeat: Packaging arm64` | `build.pull_request.id != null` | None declared | Runs `.buildkite/scripts/packaging/packaging.sh`; agents: provider=aws; image=${AWS_IMAGE_UBUNTU_ARM_64}; instanceType=${AWS_ARM_INSTANCE_TYPE}; plugins: `elastic/vault-docker-login#v0.8.0`, `elastic/vault-docker-login#v0.8.0` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | Environment names referenced: ASDF_MAGE_VERSION, AWS_ARM_INSTANCE_TYPE, AWS_IMAGE_UBUNTU_ARM_64, GCP_DEFAULT_MACHINE_TYPE, GCP_HI_PERF_MACHINE_TYPE, IMAGE_BEATS_WITH_HOOKS_LATEST, IMAGE_UBUNTU_X86_64, IMAGE_WIN_2022, PLATFORMS, RACE_DETECTOR, TEST_COVERAGE |
| Outputs | Artifact paths: `x-pack/dockerlogbeat/build/*.json`, `x-pack/dockerlogbeat/build/*.xml`, `x-pack/dockerlogbeat/build/integration-tests/*`, `x-pack/dockerlogbeat/build/integration-tests/Test*/*`, `x-pack/dockerlogbeat/build/integration-tests/Test*/data/**/*`; downstream pipelines: none declared |
| Permissions | Agent requirements are configured per step (see table).; Plugins are configured per step (see table). |
| Secrets | No secret values are documented; named secret access is configured in referenced plugins or external Buildkite settings |

## Operational notes

- One or more steps declare an explicit timeout.
- One or more steps declare automatic retry behavior.

## Sources

- Pipeline definition: [`.buildkite/x-pack/pipeline.xpack.dockerlogbeat.yml`](../x-pack/pipeline.xpack.dockerlogbeat.yml)
- Referenced script: [`.buildkite/scripts/packaging/packaging.sh`](../../.buildkite/scripts/packaging/packaging.sh)
- Referenced script: [`.buildkite/scripts/qemu.sh`](../../.buildkite/scripts/qemu.sh)
- Referenced local dependency: [`.buildkite/scripts/docker.sh`](../../.buildkite/scripts/docker.sh)
