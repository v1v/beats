# beats-filebeat

Buildkite pipeline definition for beats-filebeat.

## At a glance

| | |
| --- | --- |
| Platform | Buildkite |
| Source | [`.buildkite/filebeat/filebeat-pipeline.yml`](../filebeat/filebeat-pipeline.yml) |
| Trigger | Not specified in this checked-in YAML; start conditions may be configured in Buildkite. |

## Purpose

This pipeline runs the configured `beats-filebeat` work, organized into `Check/Update`, `Mandatory Tests`, `Extended Tests`, `Extended Windows Tests`, `Packaging`. The source does not state a separate motivation.

## When it runs

- Not specified in this checked-in YAML; start conditions may be configured in Buildkite.

## Execution flow

```mermaid
flowchart TD
  filebeat_check_update --> Wait_barrier
  filebeat_mandatory_tests --> Wait_barrier
```

| Step | Condition | Depends on | Configured work |
| --- | --- | --- | --- |
| `Check/Update / Filebeat: Run check/update` | `No step condition declared` | None declared | Runs `make -C filebeat check update`; `make check-no-changes`; agents: image=${IMAGE_BEATS_WITH_HOOKS_LATEST}; cpu=4000m; memory=8Gi; useCustomGlobalHooks=true |
| `Check/Update / Filebeat: Run pre-commit` | `No step condition declared` | None declared | Runs `pre-commit run --all-files`; agents: image=${IMAGE_BEATS_WITH_HOOKS_LATEST}; memory=2Gi; useCustomGlobalHooks=true |
| `Wait barrier` | `build.env("BUILDKITE_PULL_REQUEST") != "false"` | filebeat-check-update | Waits at a Buildkite barrier |
| `Mandatory Tests / :ubuntu: Filebeat: Ubuntu x86_64 Unit Tests` | `No step condition declared` | None declared | Runs `cd filebeat`; `mage build unitTest`; agents: provider=gcp; image=${IMAGE_UBUNTU_X86_64}; machineType=${GCP_DEFAULT_MACHINE_TYPE}; plugins: `test-collector#v1.10.2`; artifacts: `filebeat/build/*.xml`, `filebeat/build/*.json`, `filebeat/build/input-test/**`, `filebeat/build/integration-tests/**` |
| `Mandatory Tests / :ubuntu: Filebeat: Ubuntu x86_64 Unit Tests with requirefips build tag` | `No step condition declared` | None declared | Runs `cd filebeat`; `mage unitTest`; agents: provider=gcp; image=${IMAGE_UBUNTU_X86_64}; machineType=${GCP_DEFAULT_MACHINE_TYPE}; plugins: `test-collector#v1.10.2`; artifacts: `filebeat/build/*.xml`, `filebeat/build/*.json`, `filebeat/build/input-test/**`, `filebeat/build/integration-tests/**` |
| `Mandatory Tests / :ubuntu: Filebeat: Ubuntu x86_64 fips140=only Unit Tests` | `No step condition declared` | None declared | Runs `cd filebeat`; `mage goFIPSOnlyUnitTest`; agents: provider=gcp; image=${IMAGE_UBUNTU_X86_64}; machineType=${GCP_DEFAULT_MACHINE_TYPE}; plugins: `test-collector#v1.10.2`; artifacts: `filebeat/build/*.xml`, `filebeat/build/*.json`, `filebeat/build/input-test/**`, `filebeat/build/integration-tests/**` |
| `Mandatory Tests / :ubuntu: Filebeat: Go Integration Tests` | `No step condition declared` | None declared | Runs `cd filebeat`; `mage goIntegTest`; agents: provider=gcp; image=${IMAGE_UBUNTU_X86_64}; machineType=${GCP_HI_PERF_MACHINE_TYPE}; plugins: `test-collector#v1.10.2`; artifacts: `filebeat/build/*.xml`, `filebeat/build/*.json`, `filebeat/build/input-test/**`, `filebeat/build/integration-tests/**` |
| `Mandatory Tests / :windows: Filebeat: Go Integration Tests` | `build.env("BUILDKITE_PULL_REQUEST") != "false"` | None declared | Runs `cd filebeat`; `mage goIntegTest`; agents: provider=gcp; image=${IMAGE_WIN_2022}; machine_type=${GCP_WIN_MACHINE_TYPE}; disk_size=200; disk_type=pd-ssd; plugins: `test-collector#v1.10.2`; artifacts: `filebeat/build/*.xml`, `filebeat/build/*.json`, `filebeat/build/input-test/**`, `filebeat/build/integration-tests/**`; skipped: Skipped due to https://github.com/elastic/beats/issues/44537 |
| `Mandatory Tests / :ubuntu: Filebeat: Go fips140=only Integration Tests` | `No step condition declared` | None declared | Runs `cd filebeat`; `mage goFIPSOnlyIntegTest`; agents: provider=gcp; image=${IMAGE_UBUNTU_X86_64}; machineType=${GCP_HI_PERF_MACHINE_TYPE}; plugins: `test-collector#v1.10.2`; artifacts: `filebeat/build/*.xml`, `filebeat/build/*.json`, `filebeat/build/input-test/**`, `filebeat/build/integration-tests/**` |
| `Mandatory Tests / :ubuntu: Filebeat: Python Integration Tests` | `No step condition declared` | None declared | Runs `cd filebeat`; `mage pythonIntegTest`; agents: provider=gcp; image=${IMAGE_UBUNTU_X86_64}; machineType=${GCP_HI_PERF_MACHINE_TYPE}; plugins: `test-collector#v1.10.2`; artifacts: `filebeat/build/*.xml`, `filebeat/build/*.json`, `filebeat/build/integration-tests/*`, `filebeat/build/integration-tests/Test*/*`, `filebeat/build/integration-tests/Test*/data/**/*` |
| `Mandatory Tests / :ubuntu: Filebeat: Log input as Filestream Python Integration Tests` | `No step condition declared` | None declared | Runs `cd filebeat`; `mage PythonIntegTestLogAsFilestream`; agents: provider=gcp; image=${IMAGE_UBUNTU_X86_64}; machineType=${GCP_HI_PERF_MACHINE_TYPE}; plugins: `test-collector#v1.10.2`; artifacts: `filebeat/build/*.xml`, `filebeat/build/*.json`, `filebeat/build/integration-tests/*`, `filebeat/build/integration-tests/Test*/*`, `filebeat/build/integration-tests/Test*/data/**/*` |
| `Mandatory Tests / :windows: Filebeat: Win 2016 Unit Tests` | `No step condition declared` | None declared | Runs `Set-Location -Path filebeat`; `mage build unitTest`; agents: provider=gcp; image=${IMAGE_WIN_2016}; machine_type=${GCP_WIN_MACHINE_TYPE}; disk_size=200; disk_type=pd-ssd; plugins: `test-collector#v1.10.2`; artifacts: `filebeat/build/*.xml`, `filebeat/build/*.json`, `filebeat/build/integration-tests/*`, `filebeat/build/integration-tests/Test*/*`, `filebeat/build/integration-tests/Test*/data/**/*` |
| `Mandatory Tests / :windows: Filebeat: Win 2022 Unit Tests` | `No step condition declared` | None declared | Runs `Set-Location -Path filebeat`; `mage build unitTest`; agents: provider=gcp; image=${IMAGE_WIN_2022}; machine_type=${GCP_WIN_MACHINE_TYPE}; disk_size=200; disk_type=pd-ssd; plugins: `test-collector#v1.10.2`; artifacts: `filebeat/build/*.xml`, `filebeat/build/*.json` |
| `Extended Tests / :ubuntu: Filebeat: Ubuntu arm64 Unit Tests` | `build.env("BUILDKITE_PULL_REQUEST") == "false" \|\| build.env("GITHUB_PR_LABELS") =~ /.*arm.*/` | None declared | Runs `cd filebeat`; `mage build unitTest`; agents: provider=aws; image=${AWS_IMAGE_UBUNTU_ARM_64}; instanceType=${AWS_ARM_INSTANCE_TYPE}; plugins: `test-collector#v1.10.2`; artifacts: `filebeat/build/*.xml`, `filebeat/build/*.json`, `filebeat/build/input-test/**`, `filebeat/build/integration-tests/**` |
| `Extended Windows Tests / :windows: Filebeat: Win 2025 Unit Tests` | `build.env("BUILDKITE_PULL_REQUEST") == "false" \|\| build.env("GITHUB_PR_LABELS") =~ /.*[Ww]indows.*/` | None declared | Runs `Set-Location -Path filebeat`; `mage build unitTest`; agents: provider=gcp; image=${IMAGE_WIN_2025}; machine_type=${GCP_WIN_MACHINE_TYPE}; disk_size=200; disk_type=pd-ssd; plugins: `test-collector#v1.10.2`; artifacts: `filebeat/build/*.xml`, `filebeat/build/*.json` |
| `Extended Windows Tests / :windows: Filebeat: Win 2019 Unit Tests` | `build.env("BUILDKITE_PULL_REQUEST") == "false" \|\| build.env("GITHUB_PR_LABELS") =~ /.*[Ww]indows.*/` | None declared | Runs `Set-Location -Path filebeat`; `mage build unitTest`; agents: provider=gcp; image=${IMAGE_WIN_2019}; machine_type=${GCP_WIN_MACHINE_TYPE}; disk_size=200; disk_type=pd-ssd; plugins: `test-collector#v1.10.2`; artifacts: `filebeat/build/*.xml`, `filebeat/build/*.json` |
| `Extended Windows Tests / :windows: Filebeat: Win 11 Unit Tests` | `build.env("BUILDKITE_PULL_REQUEST") == "false" \|\| build.env("GITHUB_PR_LABELS") =~ /.*[Ww]indows.*/` | None declared | Runs `Set-Location -Path filebeat`; `mage build unitTest`; agents: provider=azure; image=${IMAGE_WIN_11}; imageVersion=${IMAGE_WIN_11_VERSION}; vmSize=Standard_D8s_v5; diskSizeGb=256; plugins: `test-collector#v1.10.2`; artifacts: `filebeat/build/*.xml`, `filebeat/build/*.json` |
| `Extended Windows Tests / :windows: Filebeat: Win 10 Unit Tests` | `build.env("BUILDKITE_PULL_REQUEST") == "false" \|\| build.env("GITHUB_PR_LABELS") =~ /.*[Ww]indows.*/` | None declared | Runs `Set-Location -Path filebeat`; `mage build unitTest`; agents: provider=azure; image=${IMAGE_WIN_10}; imageVersion=${IMAGE_WIN_10_VERSION}; vmSize=Standard_D8s_v5; diskSizeGb=256; plugins: `test-collector#v1.10.2`; artifacts: `filebeat/build/*.xml`, `filebeat/build/*.json` |
| `Wait barrier` | `build.env("BUILDKITE_PULL_REQUEST") != "false"` | filebeat-mandatory-tests | Waits at a Buildkite barrier |
| `Packaging / :ubuntu: Filebeat: Packaging Linux` | `No step condition declared` | None declared | Runs `.buildkite/scripts/packaging/packaging.sh`; agents: provider=gcp; image=${IMAGE_UBUNTU_X86_64}; machineType=${GCP_HI_PERF_MACHINE_TYPE}; disk_size=100; disk_type=pd-ssd; plugins: `elastic/vault-docker-login#v0.8.0`, `elastic/vault-docker-login#v0.8.0` |
| `Packaging / :ubuntu: Filebeat: Packaging arm64` | `No step condition declared` | None declared | Runs `.buildkite/scripts/packaging/packaging.sh`; agents: provider=aws; image=${AWS_IMAGE_UBUNTU_ARM_64}; instanceType=${AWS_ARM_INSTANCE_TYPE}; plugins: `elastic/vault-docker-login#v0.8.0`, `elastic/vault-docker-login#v0.8.0` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | Environment names referenced: ASDF_KIND_VERSION, ASDF_MAGE_VERSION, AWS_ARM_INSTANCE_TYPE, AWS_IMAGE_UBUNTU_ARM_64, FIPS, GCP_DEFAULT_MACHINE_TYPE, GCP_HI_PERF_MACHINE_TYPE, GCP_WIN_MACHINE_TYPE, IMAGE_BEATS_WITH_HOOKS_LATEST, IMAGE_UBUNTU_X86_64, IMAGE_WIN_10, IMAGE_WIN_10_VERSION, IMAGE_WIN_11, IMAGE_WIN_11_VERSION, IMAGE_WIN_2016, IMAGE_WIN_2019, IMAGE_WIN_2022, IMAGE_WIN_2025, K8S_VERSION, PLATFORMS, PYTEST_ADDOPTS, RACE_DETECTOR, TEST_COVERAGE |
| Outputs | Artifact paths: `filebeat/build/*.json`, `filebeat/build/*.xml`, `filebeat/build/input-test/**`, `filebeat/build/integration-tests/*`, `filebeat/build/integration-tests/**`, `filebeat/build/integration-tests/Test*/*`, `filebeat/build/integration-tests/Test*/data/**/*`; downstream pipelines: none declared |
| Permissions | Agent requirements are configured per step (see table).; Plugins are configured per step (see table). |
| Secrets | No secret values are documented; named secret access is configured in referenced plugins or external Buildkite settings |

## Operational notes

- One or more steps declare an explicit timeout.
- One or more steps declare automatic retry behavior.

## Sources

- Pipeline definition: [`.buildkite/filebeat/filebeat-pipeline.yml`](../filebeat/filebeat-pipeline.yml)
- Referenced script: [`.buildkite/scripts/packaging/packaging.sh`](../../.buildkite/scripts/packaging/packaging.sh)
- Referenced local dependency: [`.buildkite/scripts/docker.sh`](../../.buildkite/scripts/docker.sh)
- Referenced local dependency: [`.buildkite/scripts/qemu.sh`](../../.buildkite/scripts/qemu.sh)
