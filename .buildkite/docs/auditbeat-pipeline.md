# beats-auditbeat

Buildkite pipeline definition for beats-auditbeat.

## At a glance

| | |
| --- | --- |
| Platform | Buildkite |
| Source | [`.buildkite/auditbeat/auditbeat-pipeline.yml`](../auditbeat/auditbeat-pipeline.yml) |
| Trigger | Not specified in this checked-in YAML; start conditions may be configured in Buildkite. |

## Purpose

This pipeline runs the configured `beats-auditbeat` work, organized into `Check/Update`, `Mandatory Testing`, `Extended Tests`, `Extended Windows Tests`, `Packaging`. The source does not state a separate motivation.

## When it runs

- Not specified in this checked-in YAML; start conditions may be configured in Buildkite.

## Execution flow

```mermaid
flowchart TD
  auditbeat_check_update --> Wait_barrier
  auditbeat_mandatory_tests --> Wait_barrier
```

| Step | Condition | Depends on | Configured work |
| --- | --- | --- | --- |
| `Check/Update / Auditbeat: Run check/update` | `No step condition declared` | None declared | Runs `make -C auditbeat check update`; `make check-no-changes`; agents: image=${IMAGE_BEATS_WITH_HOOKS_LATEST}; cpu=4000m; memory=8Gi; useCustomGlobalHooks=true |
| `Check/Update / Auditbeat: Run pre-commit` | `No step condition declared` | None declared | Runs `pre-commit run --all-files`; agents: image=${IMAGE_BEATS_WITH_HOOKS_LATEST}; memory=2Gi; useCustomGlobalHooks=true |
| `Wait barrier` | `build.env("BUILDKITE_PULL_REQUEST") != "false"` | auditbeat-check-update | Waits at a Buildkite barrier |
| `Mandatory Testing / :ubuntu: Auditbeat: Ubuntu x86_64 Unit Tests` | `No step condition declared` | None declared | Runs `cd auditbeat`; `mage build unitTest`; agents: provider=gcp; image=${IMAGE_UBUNTU_2204_X86_64}; machineType=${GCP_DEFAULT_MACHINE_TYPE}; plugins: `test-collector#v1.10.2`; artifacts: `auditbeat/build/*.xml`, `auditbeat/build/*.json` |
| `Mandatory Testing / :ubuntu: Auditbeat: Ubuntu x86_64 Unit Tests with requirefips build tag` | `No step condition declared` | None declared | Runs `cd auditbeat`; `mage unitTest`; agents: provider=gcp; image=${IMAGE_UBUNTU_2204_X86_64}; machineType=${GCP_DEFAULT_MACHINE_TYPE}; plugins: `test-collector#v1.10.2`; artifacts: `auditbeat/build/*.xml`, `auditbeat/build/*.json` |
| `Mandatory Testing / :ubuntu: Auditbeat: Ubuntu x86_64 fips140=only Unit Tests` | `No step condition declared` | None declared | Runs `cd auditbeat`; `mage goFIPSOnlyUnitTest`; agents: provider=gcp; image=${IMAGE_UBUNTU_2204_X86_64}; machineType=${GCP_DEFAULT_MACHINE_TYPE}; plugins: `test-collector#v1.10.2`; artifacts: `auditbeat/build/*.xml`, `auditbeat/build/*.json` |
| `Mandatory Testing / :rhel: Auditbeat: RHEL9 Unit Tests` | `No step condition declared` | None declared | Runs `cd auditbeat`; `mage build unitTest`; agents: provider=gcp; image=${IMAGE_RHEL9}; machineType=${GCP_DEFAULT_MACHINE_TYPE}; plugins: `test-collector#v1.10.2`; artifacts: `auditbeat/build/*.xml`, `auditbeat/build/*.json` |
| `Mandatory Testing / :windows: Auditbeat: Win 2016 Unit Tests` | `No step condition declared` | None declared | Runs `Set-Location -Path auditbeat`; `mage build unitTest`; agents: provider=gcp; image=${IMAGE_WIN_2016}; machine_type=${GCP_WIN_MACHINE_TYPE}; disk_size=200; disk_type=pd-ssd; plugins: `test-collector#v1.10.2`; artifacts: `auditbeat/build/*.xml`, `auditbeat/build/*.json` |
| `Mandatory Testing / :windows: Auditbeat: Win 2022 Unit Tests` | `No step condition declared` | None declared | Runs `Set-Location -Path auditbeat`; `mage build unitTest`; agents: provider=gcp; image=${IMAGE_WIN_2022}; machine_type=${GCP_WIN_MACHINE_TYPE}; disk_size=200; disk_type=pd-ssd; plugins: `test-collector#v1.10.2`; artifacts: `auditbeat/build/*.xml`, `auditbeat/build/*.json` |
| `Mandatory Testing / :windows: Auditbeat: Win 2025 Unit Tests` | `No step condition declared` | None declared | Runs `Set-Location -Path auditbeat`; `mage build unitTest`; agents: provider=gcp; image=${IMAGE_WIN_2025}; machine_type=${GCP_WIN_MACHINE_TYPE}; disk_size=200; disk_type=pd-ssd; plugins: `test-collector#v1.10.2`; artifacts: `auditbeat/build/*.xml`, `auditbeat/build/*.json` |
| `Extended Tests / :ubuntu: Auditbeat: Ubuntu x86_64 Integration Tests -- {{matrix.image}}` | `build.env("BUILDKITE_PULL_REQUEST") == "false" \|\| build.env("GITHUB_PR_LABELS") =~ /.*integrations.*/` | None declared | Runs `cd auditbeat`; `mage build integTest`; agents: provider=gcp; image={{matrix.image}}; machineType=${GCP_DEFAULT_MACHINE_TYPE}; plugins: `test-collector#v1.10.2`; artifacts: `auditbeat/build/*.xml`, `auditbeat/build/*.json` |
| `Extended Tests / :ubuntu: Auditbeat: Ubuntu arm64 Integration Tests -- {{matrix.image}}` | `build.env("BUILDKITE_PULL_REQUEST") == "false" \|\| build.env("GITHUB_PR_LABELS") =~ /.*integrations.*/` | None declared | Runs `cd auditbeat`; `mage build integTest`; agents: provider=aws; image={{matrix.image}}; instanceType=${AWS_ARM_INSTANCE_TYPE}; plugins: `test-collector#v1.10.2`; artifacts: `auditbeat/build/*.xml`, `auditbeat/build/*.json` |
| `Extended Tests / :ubuntu: Auditbeat: Ubuntu arm64 Unit Tests` | `build.env("BUILDKITE_PULL_REQUEST") == "false" \|\| build.env("GITHUB_PR_LABELS") =~ /.*arm.*/` | None declared | Runs `cd auditbeat`; `mage build unitTest`; agents: provider=aws; image=${IMAGE_UBUNTU_2204_ARM64}; instanceType=${AWS_ARM_INSTANCE_TYPE}; plugins: `test-collector#v1.10.2`; artifacts: `auditbeat/build/*.xml`, `auditbeat/build/*.json` |
| `Extended Windows Tests / :windows: Auditbeat: Win 2019 Unit Tests` | `build.env("BUILDKITE_PULL_REQUEST") == "false" \|\| build.env("GITHUB_PR_LABELS") =~ /.*[Ww]indows.*/` | None declared | Runs `Set-Location -Path auditbeat`; `mage build unitTest`; agents: provider=gcp; image=${IMAGE_WIN_2019}; machine_type=${GCP_WIN_MACHINE_TYPE}; disk_size=100; disk_type=pd-ssd; plugins: `test-collector#v1.10.2`; artifacts: `auditbeat/build/*.xml`, `auditbeat/build/*.json` |
| `Extended Windows Tests / :windows: Auditbeat: Win 10 Unit Tests` | `build.env("BUILDKITE_PULL_REQUEST") == "false" \|\| build.env("GITHUB_PR_LABELS") =~ /.*[Ww]indows.*/` | None declared | Runs `Set-Location -Path auditbeat`; `mage build unitTest`; agents: provider=azure; image=${IMAGE_WIN_10}; imageVersion=${IMAGE_WIN_10_VERSION}; vmSize=Standard_D8s_v5; diskSizeGb=256; plugins: `test-collector#v1.10.2`; artifacts: `auditbeat/build/*.xml`, `auditbeat/build/*.json` |
| `Extended Windows Tests / :windows: Auditbeat: Win 11 Unit Tests` | `build.env("BUILDKITE_PULL_REQUEST") == "false" \|\| build.env("GITHUB_PR_LABELS") =~ /.*[Ww]indows.*/` | None declared | Runs `Set-Location -Path auditbeat`; `mage build unitTest`; agents: provider=azure; image=${IMAGE_WIN_11}; imageVersion=${IMAGE_WIN_11_VERSION}; vmSize=Standard_D8s_v5; diskSizeGb=256; plugins: `test-collector#v1.10.2`; artifacts: `auditbeat/build/*.xml`, `auditbeat/build/*.json` |
| `Wait barrier` | `build.env("BUILDKITE_PULL_REQUEST") != "false"` | auditbeat-mandatory-tests | Waits at a Buildkite barrier |
| `Packaging / :ubuntu: Auditbeat: Packaging Linux` | `build.pull_request.id != null` | None declared | Runs `.buildkite/scripts/packaging/packaging.sh`; agents: provider=gcp; image=${IMAGE_UBUNTU_2204_X86_64}; machineType=${GCP_HI_PERF_MACHINE_TYPE}; plugins: `elastic/vault-docker-login#v0.8.0`, `elastic/vault-docker-login#v0.8.0` |
| `Packaging / :ubuntu: Auditbeat: Packaging Linux arm64` | `build.pull_request.id != null` | None declared | Runs `.buildkite/scripts/packaging/packaging.sh`; agents: provider=aws; image=${IMAGE_UBUNTU_2204_ARM64}; instanceType=${AWS_ARM_INSTANCE_TYPE}; plugins: `elastic/vault-docker-login#v0.8.0`, `elastic/vault-docker-login#v0.8.0` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | Environment names referenced: ASDF_MAGE_VERSION, AWS_ARM_INSTANCE_TYPE, FIPS, GCP_DEFAULT_MACHINE_TYPE, GCP_HI_PERF_MACHINE_TYPE, GCP_WIN_MACHINE_TYPE, IMAGE_BEATS_WITH_HOOKS_LATEST, IMAGE_RHEL9, IMAGE_UBUNTU_2004_ARM64, IMAGE_UBUNTU_2004_X86_64, IMAGE_UBUNTU_2204_ARM64, IMAGE_UBUNTU_2204_X86_64, IMAGE_UBUNTU_2404_ARM64, IMAGE_UBUNTU_2404_X86_64, IMAGE_WIN_10, IMAGE_WIN_10_VERSION, IMAGE_WIN_11, IMAGE_WIN_11_VERSION, IMAGE_WIN_2016, IMAGE_WIN_2019, IMAGE_WIN_2022, IMAGE_WIN_2025, PLATFORMS, RACE_DETECTOR, TEST_COVERAGE |
| Outputs | Artifact paths: `auditbeat/build/*.json`, `auditbeat/build/*.xml`; downstream pipelines: none declared |
| Permissions | Agent requirements are configured per step (see table).; Plugins are configured per step (see table). |
| Secrets | No secret values are documented; named secret access is configured in referenced plugins or external Buildkite settings |

## Operational notes

- One or more steps declare an explicit timeout.
- One or more steps declare automatic retry behavior.

## Sources

- Pipeline definition: [`.buildkite/auditbeat/auditbeat-pipeline.yml`](../auditbeat/auditbeat-pipeline.yml)
- Referenced script: [`.buildkite/scripts/packaging/packaging.sh`](../../.buildkite/scripts/packaging/packaging.sh)
- Referenced local dependency: [`.buildkite/scripts/docker.sh`](../../.buildkite/scripts/docker.sh)
- Referenced local dependency: [`.buildkite/scripts/qemu.sh`](../../.buildkite/scripts/qemu.sh)
