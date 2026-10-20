# beats-metricbeat

Buildkite pipeline definition for beats-metricbeat.

## At a glance

| | |
| --- | --- |
| Platform | Buildkite |
| Source | [`.buildkite/metricbeat/pipeline.yml`](../metricbeat/pipeline.yml) |
| Trigger | Not specified in this checked-in YAML; start conditions may be configured in Buildkite. |

## Purpose

This pipeline runs the configured `beats-metricbeat` work, organized into `Check/Update`, `Mandatory Tests`, `Extended Windows Tests`, `Packaging`. The source does not state a separate motivation.

## When it runs

- Not specified in this checked-in YAML; start conditions may be configured in Buildkite.
- The top-level component triggers require a pull request and use `if_changed` include patterns. A separate serverless-beats trigger is conditioned on not being a pull request.

## Execution flow

```mermaid
flowchart TD
  metricbeat_check_update --> Wait_barrier
  metricbeat_mandatory_tests --> Wait_barrier
```

| Step | Condition | Depends on | Configured work |
| --- | --- | --- | --- |
| `Check/Update / Metricbeat: Run check/update` | `No step condition declared` | None declared | Runs `make -C metricbeat check update`; `make check-no-changes`; agents: image=${IMAGE_BEATS_WITH_HOOKS_LATEST}; cpu=4000m; memory=8Gi; useCustomGlobalHooks=true |
| `Check/Update / Metricbeat: Run pre-commit` | `No step condition declared` | None declared | Runs `pre-commit run --all-files`; agents: image=${IMAGE_BEATS_WITH_HOOKS_LATEST}; memory=2Gi; useCustomGlobalHooks=true |
| `Wait barrier` | `build.env("BUILDKITE_PULL_REQUEST") != "false"` | metricbeat-check-update | Waits at a Buildkite barrier |
| `Mandatory Tests / :ubuntu: Metricbeat: Ubuntu x86_64 Unit Tests` | `No step condition declared` | None declared | Runs `cd metricbeat`; `mage build unitTest`; agents: provider=gcp; image=${IMAGE_UBUNTU_X86_64}; machineType=${GCP_DEFAULT_MACHINE_TYPE}; plugins: `test-collector#v1.10.2`; artifacts: `metricbeat/build/*.xml`, `metricbeat/build/*.json` |
| `Mandatory Tests / :ubuntu: Metricbeat: Ubuntu x86_64 Unit Tests with requirefips build tag` | `No step condition declared` | None declared | Runs `cd metricbeat`; `mage unitTest`; agents: provider=gcp; image=${IMAGE_UBUNTU_X86_64}; machineType=${GCP_DEFAULT_MACHINE_TYPE}; plugins: `test-collector#v1.10.2`; artifacts: `metricbeat/build/*.xml`, `metricbeat/build/*.json` |
| `Mandatory Tests / :ubuntu: Metricbeat: Ubuntu x86_64 fips140=only Unit Tests` | `No step condition declared` | None declared | Runs `cd metricbeat`; `mage goFIPSOnlyUnitTest`; agents: provider=gcp; image=${IMAGE_UBUNTU_X86_64}; machineType=${GCP_DEFAULT_MACHINE_TYPE}; plugins: `test-collector#v1.10.2`; artifacts: `metricbeat/build/*.xml`, `metricbeat/build/*.json` |
| `Mandatory Tests / :ubuntu: Metricbeat: Go Integration Tests (Module)` | `No step condition declared` | None declared | Runs `.buildkite/deploy/kubernetes/scripts/setup-k8s-env.sh`; agents: provider=gcp; image=${IMAGE_UBUNTU_X86_64}; machineType=${GCP_HI_PERF_MACHINE_TYPE}; plugins: `test-collector#v1.10.2`; artifacts: `metricbeat/build/*.xml`, `metricbeat/build/*.json`, `metricbeat/build/integration-tests/*`, `metricbeat/build/integration-tests/Test*/*`, `metricbeat/build/integration-tests/Test*/data/**/*` |
| `Mandatory Tests / :windows: Metricbeat: Go Integration Tests (Module)` | `build.env("BUILDKITE_PULL_REQUEST") != "false"` | None declared | Runs `cd metricbeat`; `mage goIntegTest`; agents: provider=gcp; image=${IMAGE_WIN_2022}; machine_type=${GCP_WIN_MACHINE_TYPE}; disk_size=200; disk_type=pd-ssd; plugins: `test-collector#v1.10.2`; artifacts: `metricbeat/build/*.xml`, `metricbeat/build/*.json`, `metricbeat/build/integration-tests/*`, `metricbeat/build/integration-tests/Test*/*`, `metricbeat/build/integration-tests/Test*/data/**/*`; skipped: Skipped due to https://github.com/elastic/beats/issues/44537 |
| `Mandatory Tests / :ubuntu: Metricbeat: Go fips140=only Integration Tests (Module)` | `No step condition declared` | None declared | Runs `.buildkite/deploy/kubernetes/scripts/setup-k8s-env.sh`; agents: provider=gcp; image=${IMAGE_UBUNTU_X86_64}; machineType=${GCP_HI_PERF_MACHINE_TYPE}; plugins: `test-collector#v1.10.2`; artifacts: `metricbeat/build/*.xml`, `metricbeat/build/*.json`, `metricbeat/build/integration-tests/*`, `metricbeat/build/integration-tests/Test*/*`, `metricbeat/build/integration-tests/Test*/data/**/*` |
| `Mandatory Tests / :ubuntu: Metricbeat: Python Integration Tests (Module)` | `No step condition declared` | None declared | Runs `.buildkite/deploy/kubernetes/scripts/setup-k8s-env.sh`; agents: provider=gcp; image=${IMAGE_UBUNTU_X86_64}; machineType=${GCP_HI_PERF_MACHINE_TYPE}; plugins: `test-collector#v1.10.2`; artifacts: `metricbeat/build/*.xml`, `metricbeat/build/*.json`, `metricbeat/build/integration-tests/*`, `metricbeat/build/integration-tests/Test*/*`, `metricbeat/build/integration-tests/Test*/data/**/*` |
| `Mandatory Tests / :windows: Metricbeat: Win 2016 Unit Tests` | `No step condition declared` | None declared | Runs `Set-Location -Path metricbeat`; `mage build unitTest`; agents: provider=gcp; image=${IMAGE_WIN_2016}; machine_type=${GCP_WIN_MACHINE_TYPE}; disk_size=100; disk_type=pd-ssd; plugins: `test-collector#v1.10.2`; artifacts: `metricbeat/build/*.xml`, `metricbeat/build/*.json` |
| `Mandatory Tests / :windows: Metricbeat: Win 2022 Unit Tests` | `No step condition declared` | None declared | Runs `Set-Location -Path metricbeat`; `mage build unitTest`; agents: provider=gcp; image=${IMAGE_WIN_2022}; machine_type=${GCP_WIN_MACHINE_TYPE}; disk_size=100; disk_type=pd-ssd; plugins: `test-collector#v1.10.2`; artifacts: `metricbeat/build/*.xml`, `metricbeat/build/*.json` |
| `Extended Windows Tests / :windows: Metricbeat: Win 10 Unit Tests` | `build.env("BUILDKITE_PULL_REQUEST") == "false" \|\| build.env("GITHUB_PR_LABELS") =~ /.*[Ww]indows.*/` | None declared | Runs `Set-Location -Path metricbeat`; `mage build unitTest`; agents: provider=azure; image=${IMAGE_WIN_10}; imageVersion=${IMAGE_WIN_10_VERSION}; vmSize=Standard_D8s_v5; diskSizeGb=256; plugins: `test-collector#v1.10.2`; artifacts: `metricbeat/build/*.xml`, `metricbeat/build/*.json` |
| `Extended Windows Tests / :windows: Metricbeat: Win 11 Unit Tests` | `build.env("BUILDKITE_PULL_REQUEST") == "false" \|\| build.env("GITHUB_PR_LABELS") =~ /.*[Ww]indows.*/` | None declared | Runs `Set-Location -Path metricbeat`; `mage build unitTest`; agents: provider=azure; image=${IMAGE_WIN_11}; imageVersion=${IMAGE_WIN_11_VERSION}; vmSize=Standard_D8s_v5; diskSizeGb=256; plugins: `test-collector#v1.10.2`; artifacts: `metricbeat/build/*.xml`, `metricbeat/build/*.json` |
| `Extended Windows Tests / :windows: Metricbeat: Win 2019 Unit Tests` | `build.env("BUILDKITE_PULL_REQUEST") == "false" \|\| build.env("GITHUB_PR_LABELS") =~ /.*[Ww]indows.*/` | None declared | Runs `Set-Location -Path metricbeat`; `mage build unitTest`; agents: provider=gcp; image=${IMAGE_WIN_2019}; machineType=${GCP_WIN_MACHINE_TYPE}; disk_size=100; disk_type=pd-ssd; plugins: `test-collector#v1.10.2`; artifacts: `metricbeat/build/*.xml`, `metricbeat/build/*.json` |
| `Extended Windows Tests / :windows: Metricbeat: Win 2025 Unit Tests` | `build.env("BUILDKITE_PULL_REQUEST") == "false" \|\| build.env("GITHUB_PR_LABELS") =~ /.*[Ww]indows.*/` | None declared | Runs `Set-Location -Path metricbeat`; `mage build unitTest`; agents: provider=gcp; image=${IMAGE_WIN_2025}; machineType=${GCP_WIN_MACHINE_TYPE}; disk_size=100; disk_type=pd-ssd; plugins: `test-collector#v1.10.2`; artifacts: `metricbeat/build/*.xml`, `metricbeat/build/*.json` |
| `Wait barrier` | `build.env("BUILDKITE_PULL_REQUEST") != "false"` | metricbeat-mandatory-tests | Waits at a Buildkite barrier |
| `Packaging / :ubuntu: Metricbeat: Packaging Linux` | `build.pull_request.id != null` | None declared | Runs `.buildkite/scripts/packaging/packaging.sh`; agents: provider=gcp; image=${IMAGE_UBUNTU_X86_64}; machineType=${GCP_HI_PERF_MACHINE_TYPE}; disk_size=100; disk_type=pd-ssd; plugins: `elastic/vault-docker-login#v0.8.0`, `elastic/vault-docker-login#v0.8.0` |
| `Packaging / :ubuntu: Metricbeat: Packaging Linux arm64` | `build.pull_request.id != null` | None declared | Runs `.buildkite/scripts/packaging/packaging.sh`; agents: provider=aws; image=${AWS_IMAGE_UBUNTU_ARM_64}; instanceType=${AWS_ARM_INSTANCE_TYPE}; plugins: `elastic/vault-docker-login#v0.8.0`, `elastic/vault-docker-login#v0.8.0` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | Environment names referenced: ASDF_KIND_VERSION, ASDF_MAGE_VERSION, AWS_ARM_INSTANCE_TYPE, AWS_IMAGE_UBUNTU_ARM_64, BEAT_PATH, FIPS, GCP_DEFAULT_MACHINE_TYPE, GCP_HI_PERF_MACHINE_TYPE, GCP_WIN_MACHINE_TYPE, IMAGE_BEATS_WITH_HOOKS_LATEST, IMAGE_UBUNTU_X86_64, IMAGE_WIN_10, IMAGE_WIN_10_VERSION, IMAGE_WIN_11, IMAGE_WIN_11_VERSION, IMAGE_WIN_2016, IMAGE_WIN_2019, IMAGE_WIN_2022, IMAGE_WIN_2025, K8S_VERSION, PLATFORMS, RACE_DETECTOR, TEST_COVERAGE |
| Outputs | Artifact paths: `metricbeat/build/*.json`, `metricbeat/build/*.xml`, `metricbeat/build/integration-tests/*`, `metricbeat/build/integration-tests/Test*/*`, `metricbeat/build/integration-tests/Test*/data/**/*`; downstream pipelines: none declared |
| Permissions | Agent requirements are configured per step (see table).; Plugins are configured per step (see table). |
| Secrets | No secret values are documented; named secret access is configured in referenced plugins or external Buildkite settings |

## Operational notes

- One or more steps declare an explicit timeout.
- One or more steps declare automatic retry behavior.
- The root pipeline conditionally uploads component builds based on changed paths; downstream test/package behavior is defined in the referenced pipeline YAML files.

## Sources

- Pipeline definition: [`.buildkite/metricbeat/pipeline.yml`](../metricbeat/pipeline.yml)
- Referenced script: [`.buildkite/deploy/kubernetes/scripts/setup-k8s-env.sh`](../../.buildkite/deploy/kubernetes/scripts/setup-k8s-env.sh)
- Referenced script: [`.buildkite/scripts/packaging/packaging.sh`](../../.buildkite/scripts/packaging/packaging.sh)
- Referenced local dependency: [`.buildkite/deploy/kubernetes/scripts/kind-setup.sh`](../../.buildkite/deploy/kubernetes/scripts/kind-setup.sh)
- Referenced local dependency: [`.buildkite/scripts/docker.sh`](../../.buildkite/scripts/docker.sh)
- Referenced local dependency: [`.buildkite/scripts/qemu.sh`](../../.buildkite/scripts/qemu.sh)
