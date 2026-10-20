# heartbeat-pipeline

Buildkite pipeline definition for heartbeat-pipeline.

## At a glance

| | |
| --- | --- |
| Platform | Buildkite |
| Source | [`.buildkite/heartbeat/heartbeat-pipeline.yml`](../heartbeat/heartbeat-pipeline.yml) |
| Trigger | Not specified in this checked-in YAML; start conditions may be configured in Buildkite. |

## Purpose

This pipeline runs the configured `heartbeat-pipeline` work, organized into `Check/Update`, `Mandatory Testing`, `Extended Tests`, `Extended Windows Tests`, `Packaging`. The source does not state a separate motivation.

## When it runs

- Not specified in this checked-in YAML; start conditions may be configured in Buildkite.

## Execution flow

```mermaid
flowchart TD
  heartbeat_check_update --> Wait_barrier
  heartbeat_mandatory_tests --> Wait_barrier
```

| Step | Condition | Depends on | Configured work |
| --- | --- | --- | --- |
| `Check/Update / Heartbeat: Run check/update` | `No step condition declared` | None declared | Runs `make -C heartbeat check update`; `make check-no-changes`; agents: image=${IMAGE_BEATS_WITH_HOOKS_LATEST}; cpu=4000m; memory=8Gi; useCustomGlobalHooks=true |
| `Check/Update / Heartbeat: Run pre-commit` | `No step condition declared` | None declared | Runs `pre-commit run --all-files`; agents: image=${IMAGE_BEATS_WITH_HOOKS_LATEST}; memory=2Gi; useCustomGlobalHooks=true |
| `Wait barrier` | `build.env("BUILDKITE_PULL_REQUEST") != "false"` | heartbeat-check-update | Waits at a Buildkite barrier |
| `Mandatory Testing / :ubuntu: Heartbeat: Ubuntu x86_64 Unit Tests` | `No step condition declared` | None declared | Runs `cd heartbeat`; `mage build unitTest`; agents: provider=gcp; image=${IMAGE_UBUNTU_X86_64}; machineType=${GCP_DEFAULT_MACHINE_TYPE}; artifacts: `heartbeat/build/*.xml`, `heartbeat/build/*.json` |
| `Mandatory Testing / :ubuntu: Heartbeat: Ubuntu x86_64 Unit Tests with requirefips build tag` | `No step condition declared` | None declared | Runs `cd heartbeat`; `mage unitTest`; agents: provider=gcp; image=${IMAGE_UBUNTU_X86_64}; machineType=${GCP_DEFAULT_MACHINE_TYPE}; plugins: `elastic/vault-secrets#v0.1.0`, `test-collector#v1.11.0`; artifacts: `heartbeat/build/*.xml`, `heartbeat/build/*.json` |
| `Mandatory Testing / :ubuntu: Heartbeat: Ubuntu x86_64 fips140=only Unit Tests` | `No step condition declared` | None declared | Runs `cd heartbeat`; `mage goFIPSOnlyUnitTest`; agents: provider=gcp; image=${IMAGE_UBUNTU_X86_64}; machineType=${GCP_DEFAULT_MACHINE_TYPE}; plugins: `elastic/vault-secrets#v0.1.0`, `test-collector#v1.11.0`; artifacts: `heartbeat/build/*.xml`, `heartbeat/build/*.json` |
| `Mandatory Testing / :rhel: Heartbeat: RHEL9 Unit Tests` | `No step condition declared` | None declared | Runs `cd heartbeat`; `mage build unitTest`; agents: provider=gcp; image=${IMAGE_RHEL9}; machineType=${GCP_DEFAULT_MACHINE_TYPE}; artifacts: `heartbeat/build/*.xml`, `heartbeat/build/*.json` |
| `Mandatory Testing / :windows: Heartbeat: Win 2016 Unit Tests` | `No step condition declared` | None declared | Runs `Set-Location -Path heartbeat`; `mage build unitTest`; agents: provider=gcp; image=${IMAGE_WIN_2016}; machine_type=${GCP_WIN_MACHINE_TYPE}; disk_type=pd-ssd; artifacts: `heartbeat/build/*.xml`, `heartbeat/build/*.json` |
| `Mandatory Testing / :windows: Heartbeat: Win 2022 Unit Test` | `No step condition declared` | None declared | Runs `Set-Location -Path heartbeat`; `mage build unitTest`; agents: provider=gcp; image=${IMAGE_WIN_2022}; machine_type=${GCP_WIN_MACHINE_TYPE}; disk_type=pd-ssd; artifacts: `heartbeat/build/*.xml`, `heartbeat/build/*.json` |
| `Mandatory Testing / :ubuntu: Heartbeat: Go Integration Tests` | `No step condition declared` | None declared | Runs `cd heartbeat`; `mage goIntegTest`; agents: provider=gcp; image=${IMAGE_UBUNTU_X86_64}; machineType=${GCP_HI_PERF_MACHINE_TYPE}; artifacts: `heartbeat/build/*.xml`, `heartbeat/build/*.json`, `heartbeat/build/integration-tests/*`, `heartbeat/build/integration-tests/Test*/*`, `heartbeat/build/integration-tests/Test*/data/**/*` |
| `Mandatory Testing / :windows: Heartbeat: Go Integration Tests` | `build.env("BUILDKITE_PULL_REQUEST") != "false"` | None declared | Runs `cd heartbeat`; `mage goIntegTest`; agents: provider=gcp; image=${IMAGE_WIN_2022}; machine_type=${GCP_WIN_MACHINE_TYPE}; disk_size=200; disk_type=pd-ssd; artifacts: `heartbeat/build/*.xml`, `heartbeat/build/*.json`, `heartbeat/build/integration-tests/*`, `heartbeat/build/integration-tests/Test*/*`, `heartbeat/build/integration-tests/Test*/data/**/*` |
| `Mandatory Testing / :ubuntu: Heartbeat: Python Integration Tests` | `No step condition declared` | None declared | Runs `cd heartbeat`; `mage pythonIntegTest`; agents: provider=gcp; image=${IMAGE_UBUNTU_X86_64}; machineType=${GCP_HI_PERF_MACHINE_TYPE}; artifacts: `heartbeat/build/*.xml`, `heartbeat/build/*.json`, `heartbeat/build/integration-tests/*`, `heartbeat/build/integration-tests/Test*/*`, `heartbeat/build/integration-tests/Test*/data/**/*` |
| `Extended Tests / :ubuntu: Heartbeat: Ubuntu arm64 Unit Tests` | `build.env("BUILDKITE_PULL_REQUEST") == "false" \|\| build.env("GITHUB_PR_LABELS") =~ /.*arm.*/` | None declared | Runs `cd heartbeat`; `mage build unitTest`; agents: provider=aws; image=${AWS_IMAGE_UBUNTU_ARM_64}; instanceType=${AWS_ARM_INSTANCE_TYPE}; artifacts: `heartbeat/build/*.xml` |
| `Extended Windows Tests / :windows: Heartbeat: Win 2019 Unit Tests` | `build.env("BUILDKITE_PULL_REQUEST") == "false" \|\| build.env("GITHUB_PR_LABELS") =~ /.*[Ww]indows.*/` | None declared | Runs `Set-Location -Path heartbeat`; `mage build unitTest`; agents: provider=gcp; image=${IMAGE_WIN_2019}; machine_type=${GCP_WIN_MACHINE_TYPE}; disk_type=pd-ssd; artifacts: `heartbeat/build/*.xml`, `heartbeat/build/*.json` |
| `Extended Windows Tests / :windows: Heartbeat: Win 11 Unit Tests` | `build.env("BUILDKITE_PULL_REQUEST") == "false" \|\| build.env("GITHUB_PR_LABELS") =~ /.*[Ww]indows.*/` | None declared | Runs `Set-Location -Path heartbeat`; `mage build unitTest`; agents: provider=azure; image=${IMAGE_WIN_11}; imageVersion=${IMAGE_WIN_11_VERSION}; vmSize=Standard_D8s_v5; diskSizeGb=256; artifacts: `heartbeat/build/*.xml`, `heartbeat/build/*.json` |
| `Extended Windows Tests / :windows: Heartbeat: Win 10 Unit Tests` | `build.env("BUILDKITE_PULL_REQUEST") == "false" \|\| build.env("GITHUB_PR_LABELS") =~ /.*[Ww]indows.*/` | None declared | Runs `Set-Location -Path heartbeat`; `mage build unitTest`; agents: provider=azure; image=${IMAGE_WIN_10}; imageVersion=${IMAGE_WIN_10_VERSION}; vmSize=Standard_D8s_v5; diskSizeGb=256; artifacts: `heartbeat/build/*.xml`, `heartbeat/build/*.json` |
| `Wait barrier` | `build.env("BUILDKITE_PULL_REQUEST") != "false"` | heartbeat-mandatory-tests | Waits at a Buildkite barrier |
| `Packaging / :ubuntu: Heartbeat: Packaging Linux` | `No step condition declared` | None declared | Runs `.buildkite/scripts/packaging/packaging.sh`; agents: provider=gcp; image=${IMAGE_UBUNTU_X86_64}; machineType=${GCP_HI_PERF_MACHINE_TYPE}; plugins: `elastic/vault-docker-login#v0.8.0`, `elastic/vault-docker-login#v0.8.0` |
| `Packaging / :ubuntu: Heartbeat: Packaging Linux arm64` | `No step condition declared` | None declared | Runs `.buildkite/scripts/packaging/packaging.sh`; agents: provider=aws; image=${AWS_IMAGE_UBUNTU_ARM_64}; instanceType=${AWS_ARM_INSTANCE_TYPE}; plugins: `elastic/vault-docker-login#v0.8.0`, `elastic/vault-docker-login#v0.8.0` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | Environment names referenced: ASDF_MAGE_VERSION, ASDF_NODEJS_VERSION, ASDF_PYTHON_VERSION, AWS_ARM_INSTANCE_TYPE, AWS_IMAGE_UBUNTU_ARM_64, FIPS, GCP_DEFAULT_MACHINE_TYPE, GCP_HI_PERF_MACHINE_TYPE, GCP_WIN_MACHINE_TYPE, IMAGE_BEATS_WITH_HOOKS_LATEST, IMAGE_RHEL9, IMAGE_UBUNTU_X86_64, IMAGE_WIN_10, IMAGE_WIN_10_VERSION, IMAGE_WIN_11, IMAGE_WIN_11_VERSION, IMAGE_WIN_2016, IMAGE_WIN_2019, IMAGE_WIN_2022, PLATFORMS, RACE_DETECTOR, TEST_COVERAGE |
| Outputs | Artifact paths: `heartbeat/build/*.json`, `heartbeat/build/*.xml`, `heartbeat/build/integration-tests/*`, `heartbeat/build/integration-tests/Test*/*`, `heartbeat/build/integration-tests/Test*/data/**/*`; downstream pipelines: none declared |
| Permissions | Agent requirements are configured per step (see table).; Plugins are configured per step (see table). |
| Secrets | No secret values are documented; named secret access is configured in referenced plugins or external Buildkite settings |

## Operational notes

- One or more steps declare an explicit timeout.
- One or more steps declare automatic retry behavior.

## Sources

- Pipeline definition: [`.buildkite/heartbeat/heartbeat-pipeline.yml`](../heartbeat/heartbeat-pipeline.yml)
- Referenced script: [`.buildkite/scripts/packaging/packaging.sh`](../../.buildkite/scripts/packaging/packaging.sh)
- Referenced local dependency: [`.buildkite/scripts/docker.sh`](../../.buildkite/scripts/docker.sh)
- Referenced local dependency: [`.buildkite/scripts/qemu.sh`](../../.buildkite/scripts/qemu.sh)
