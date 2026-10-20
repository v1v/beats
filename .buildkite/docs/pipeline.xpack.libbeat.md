# beats-xpack-libbeat

Buildkite pipeline definition for beats-xpack-libbeat.

## At a glance

| | |
| --- | --- |
| Platform | Buildkite |
| Source | [`.buildkite/x-pack/pipeline.xpack.libbeat.yml`](../x-pack/pipeline.xpack.libbeat.yml) |
| Trigger | Not specified in this checked-in YAML; start conditions may be configured in Buildkite. |

## Purpose

This pipeline runs the configured `beats-xpack-libbeat` work, organized into `Check/Update`, `Mandatory Tests`, `Extended Tests`. The source does not state a separate motivation.

## When it runs

- Not specified in this checked-in YAML; start conditions may be configured in Buildkite.

## Execution flow

```mermaid
flowchart TD
  x_pack_libbeat_check_update --> Wait_barrier
```

| Step | Condition | Depends on | Configured work |
| --- | --- | --- | --- |
| `Check/Update / x-pack/libbeat: Run check/update` | `No step condition declared` | None declared | Runs `make -C x-pack/libbeat check update`; `make check-no-changes`; agents: image=${IMAGE_BEATS_WITH_HOOKS_LATEST}; cpu=4000m; memory=8Gi; ephemeralStorage=40G; useCustomGlobalHooks=true |
| `Check/Update / x-pack/libbeat: Run pre-commit` | `No step condition declared` | None declared | Runs `pre-commit run --all-files`; agents: image=${IMAGE_BEATS_WITH_HOOKS_LATEST}; memory=2Gi; useCustomGlobalHooks=true |
| `Wait barrier` | `build.env("BUILDKITE_PULL_REQUEST") != "false"` | x-pack-libbeat-check-update | Waits at a Buildkite barrier |
| `Mandatory Tests / :ubuntu: x-pack/libbeat: Ubuntu x86_64 Unit Tests` | `No step condition declared` | None declared | Runs `cd x-pack/libbeat`; `mage build unitTest`; agents: provider=gcp; image=${IMAGE_UBUNTU_X86_64}; machineType=${GCP_DEFAULT_MACHINE_TYPE}; artifacts: `x-pack/libbeat/build/*.xml`, `x-pack/libbeat/build/*.json` |
| `Mandatory Tests / :ubuntu: x-pack/libbeat: Ubuntu x86_64 Unit Tests with requirefips build tag` | `No step condition declared` | None declared | Runs `cd x-pack/libbeat`; `mage unitTest`; agents: provider=gcp; image=${IMAGE_UBUNTU_X86_64}; machineType=${GCP_DEFAULT_MACHINE_TYPE}; artifacts: `x-pack/libbeat/build/*.xml`, `x-pack/libbeat/build/*.json` |
| `Mandatory Tests / :ubuntu: x-pack/libbeat: Ubuntu x86_64 fips140=only Unit Tests` | `No step condition declared` | None declared | Runs `cd x-pack/libbeat`; `mage goFIPSOnlyUnitTest`; agents: provider=gcp; image=${IMAGE_UBUNTU_X86_64}; machineType=${GCP_DEFAULT_MACHINE_TYPE}; artifacts: `x-pack/libbeat/build/*.xml`, `x-pack/libbeat/build/*.json` |
| `Mandatory Tests / :ubuntu: x-pack/libbeat: Go Integration Tests` | `No step condition declared` | None declared | Runs `cd x-pack/libbeat`; `mage goIntegTest`; agents: provider=gcp; image=${IMAGE_UBUNTU_X86_64}; machineType=${GCP_HI_PERF_MACHINE_TYPE}; artifacts: `x-pack/libbeat/build/*.xml`, `x-pack/libbeat/build/*.json`, `x-pack/libbeat/build/integration-tests/*`, `x-pack/libbeat/build/integration-tests/Test*/*`, `x-pack/libbeat/build/integration-tests/Test*/data/**/*` |
| `Mandatory Tests / :windows: x-pack/libbeat: Go Integration Tests` | `No step condition declared` | None declared | Runs `cd x-pack/libbeat`; `mage goIntegTest`; agents: provider=gcp; image=${IMAGE_WIN_2022}; machineType=${GCP_HI_PERF_MACHINE_TYPE}; artifacts: `x-pack/libbeat/build/*.xml`, `x-pack/libbeat/build/*.json`, `x-pack/libbeat/build/integration-tests/*`, `x-pack/libbeat/build/integration-tests/Test*/*`, `x-pack/libbeat/build/integration-tests/Test*/data/**/*`; skipped: Skipped due to https://github.com/elastic/beats/issues/44537 |
| `Mandatory Tests / :ubuntu: x-pack/libbeat: Python Integration Tests` | `No step condition declared` | None declared | Runs `cd x-pack/libbeat`; `mage pythonIntegTest`; agents: provider=gcp; image=${IMAGE_UBUNTU_X86_64}; machineType=${GCP_HI_PERF_MACHINE_TYPE}; artifacts: `x-pack/libbeat/build/*.xml`, `x-pack/libbeat/build/*.json`, `x-pack/libbeat/build/integration-tests/*`, `x-pack/libbeat/build/integration-tests/Test*/*`, `x-pack/libbeat/build/integration-tests/Test*/data/**/*` |
| `Extended Tests / :ubuntu: x-pack/libbeat: Ubuntu arm64 Unit Tests` | `build.env("BUILDKITE_PULL_REQUEST") == "false" \|\| build.env("GITHUB_PR_LABELS") =~ /.*arm.*/` | None declared | Runs `cd x-pack/libbeat`; `mage build unitTest`; agents: provider=aws; image=${AWS_IMAGE_UBUNTU_ARM_64}; instanceType=${AWS_ARM_INSTANCE_TYPE}; artifacts: `x-pack/libbeat/build/*.xml`, `x-pack/libbeat/build/*.json` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | Environment names referenced: ASDF_MAGE_VERSION, AWS_ARM_INSTANCE_TYPE, AWS_IMAGE_UBUNTU_ARM_64, FIPS, GCP_DEFAULT_MACHINE_TYPE, GCP_HI_PERF_MACHINE_TYPE, GCP_WIN_MACHINE_TYPE, IMAGE_BEATS_WITH_HOOKS_LATEST, IMAGE_UBUNTU_X86_64, IMAGE_WIN_2022, RACE_DETECTOR, TEST_COVERAGE |
| Outputs | Artifact paths: `x-pack/libbeat/build/*.json`, `x-pack/libbeat/build/*.xml`, `x-pack/libbeat/build/integration-tests/*`, `x-pack/libbeat/build/integration-tests/Test*/*`, `x-pack/libbeat/build/integration-tests/Test*/data/**/*`; downstream pipelines: none declared |
| Permissions | Agent requirements are configured per step (see table). |
| Secrets | No secret values are documented; named secret access is configured in referenced plugins or external Buildkite settings |

## Operational notes

- One or more steps declare automatic retry behavior.

## Sources

- Pipeline definition: [`.buildkite/x-pack/pipeline.xpack.libbeat.yml`](../x-pack/pipeline.xpack.libbeat.yml)
