# beats-xpack-osquerybeat

Buildkite pipeline definition for beats-xpack-osquerybeat.

## At a glance

| | |
| --- | --- |
| Platform | Buildkite |
| Source | [`.buildkite/x-pack/pipeline.xpack.osquerybeat.yml`](../x-pack/pipeline.xpack.osquerybeat.yml) |
| Trigger | Not specified in this checked-in YAML; start conditions may be configured in Buildkite. |

## Purpose

This pipeline runs the configured `beats-xpack-osquerybeat` work, organized into `Check/Update`, `Mandatory Tests`, `Extended Windows Tests`, `Packaging`. The source does not state a separate motivation.

## When it runs

- Not specified in this checked-in YAML; start conditions may be configured in Buildkite.

## Execution flow

```mermaid
flowchart TD
  x_pack_osquerybeat_check_update --> Wait_barrier
  __step____x_pack_osquerybeat_mandatory_tests__ --> Wait_barrier
```

| Step | Condition | Depends on | Configured work |
| --- | --- | --- | --- |
| `Check/Update / x-pack/osquerybeat: Run check/update` | `No step condition declared` | None declared | Runs `make -C x-pack/osquerybeat check update`; `make check-no-changes`; agents: image=${IMAGE_BEATS_WITH_HOOKS_LATEST}; cpu=4000m; memory=8Gi; useCustomGlobalHooks=true |
| `Check/Update / x-pack/osquerybeat: Run pre-commit` | `No step condition declared` | None declared | Runs `pre-commit run --all-files`; agents: image=${IMAGE_BEATS_WITH_HOOKS_LATEST}; memory=2Gi; useCustomGlobalHooks=true |
| `Wait barrier` | `build.env("BUILDKITE_PULL_REQUEST") != "false"` | x-pack-osquerybeat-check-update | Waits at a Buildkite barrier |
| `Mandatory Tests / :ubuntu: x-pack/osquerybeat: Ubuntu x86_64 Unit Tests` | `No step condition declared` | None declared | Runs `cd x-pack/osquerybeat`; `mage build unitTest`; agents: provider=gcp; image=${IMAGE_UBUNTU_X86_64}; machineType=${GCP_DEFAULT_MACHINE_TYPE}; artifacts: `x-pack/osquerybeat/build/*.xml`, `x-pack/osquerybeat/build/*.json` |
| `Mandatory Tests / :ubuntu: x-pack/osquerybeat: Go Integration Tests` | `No step condition declared` | None declared | Runs `cd x-pack/osquerybeat`; `mage goIntegTest`; agents: provider=gcp; image=${IMAGE_UBUNTU_X86_64}; machineType=${GCP_HI_PERF_MACHINE_TYPE}; artifacts: `x-pack/osquerybeat/build/*.xml`, `x-pack/osquerybeat/build/*.json`, `x-pack/osquerybeat/build/integration-tests/*`, `x-pack/osquerybeat/build/integration-tests/Test*/*`, `x-pack/osquerybeat/build/integration-tests/Test*/data/**/*` |
| `Mandatory Tests / :windows: x-pack/osquerybeat: Go Integration Tests` | `No step condition declared` | None declared | Runs `cd x-pack/osquerybeat`; `mage goIntegTest`; agents: provider=gcp; image=${IMAGE_WIN_2022}; machineType=${GCP_HI_PERF_MACHINE_TYPE}; artifacts: `x-pack/osquerybeat/build/*.xml`, `x-pack/osquerybeat/build/*.json`, `x-pack/osquerybeat/build/integration-tests/*`, `x-pack/osquerybeat/build/integration-tests/Test*/*`, `x-pack/osquerybeat/build/integration-tests/Test*/data/**/*` |
| `Mandatory Tests / :windows: x-pack/osquerybeat: Win 2016 Unit Tests` | `No step condition declared` | None declared | Runs `Set-Location -Path x-pack/osquerybeat`; `mage build unitTest`; agents: provider=gcp; image=${IMAGE_WIN_2016}; machine_type=${GCP_WIN_MACHINE_TYPE}; disk_size=100; disk_type=pd-ssd; artifacts: `x-pack/osquerybeat/build/*.xml`, `x-pack/osquerybeat/build/*.json` |
| `Mandatory Tests / :windows: x-pack/osquerybeat: Win 2022 Unit Tests` | `No step condition declared` | None declared | Runs `Set-Location -Path x-pack/osquerybeat`; `mage build unitTest`; agents: provider=gcp; image=${IMAGE_WIN_2022}; machine_type=${GCP_WIN_MACHINE_TYPE}; disk_size=100; disk_type=pd-ssd; artifacts: `x-pack/osquerybeat/build/*.xml`, `x-pack/osquerybeat/build/*.json` |
| `Extended Windows Tests / :windows: x-pack/osquerybeat: Win 10 Unit Tests` | `build.env("BUILDKITE_PULL_REQUEST") == "false" \|\| build.env("GITHUB_PR_LABELS") =~ /.*[Ww]indows.*/` | None declared | Runs `Set-Location -Path x-pack/osquerybeat`; `mage build unitTest`; agents: provider=azure; image=${IMAGE_WIN_10}; imageVersion=${IMAGE_WIN_10_VERSION}; vmSize=Standard_D8s_v5; diskSizeGb=256; artifacts: `x-pack/osquerybeat/build/*.xml`, `x-pack/osquerybeat/build/*.json` |
| `Extended Windows Tests / :windows: x-pack/osquerybeat: Win 11 Unit Tests` | `build.env("BUILDKITE_PULL_REQUEST") == "false" \|\| build.env("GITHUB_PR_LABELS") =~ /.*[Ww]indows.*/` | None declared | Runs `Set-Location -Path x-pack/osquerybeat`; `mage build unitTest`; agents: provider=azure; image=${IMAGE_WIN_11}; imageVersion=${IMAGE_WIN_11_VERSION}; vmSize=Standard_D8s_v5; diskSizeGb=256; artifacts: `x-pack/osquerybeat/build/*.xml`, `x-pack/osquerybeat/build/*.json` |
| `Extended Windows Tests / :windows: x-pack/osquerybeat: Win 2019 Unit Tests` | `build.env("BUILDKITE_PULL_REQUEST") == "false" \|\| build.env("GITHUB_PR_LABELS") =~ /.*[Ww]indows.*/` | None declared | Runs `Set-Location -Path x-pack/osquerybeat`; `mage build unitTest`; agents: provider=gcp; image=${IMAGE_WIN_2019}; machineType=${GCP_WIN_MACHINE_TYPE}; disk_size=100; disk_type=pd-ssd; artifacts: `x-pack/osquerybeat/build/*.xml`, `x-pack/osquerybeat/build/*.json` |
| `Wait barrier` | `build.env("BUILDKITE_PULL_REQUEST") != "false"` | {'step': 'x-pack-osquerybeat-mandatory-tests'} | Waits at a Buildkite barrier |
| `Packaging / :ubuntu: Packaging Linux` | `build.pull_request.id != null` | None declared | Runs `.buildkite/scripts/packaging/packaging.sh`; agents: provider=gcp; image=${IMAGE_UBUNTU_X86_64}; machineType=${GCP_HI_PERF_MACHINE_TYPE}; disk_size=100; disk_type=pd-ssd; plugins: `elastic/vault-docker-login#v0.8.0`, `elastic/vault-docker-login#v0.8.0` |
| `Packaging / :ubuntu: x-pack/osquerybeat: Packaging arm64` | `build.pull_request.id != null` | None declared | Runs `.buildkite/scripts/packaging/packaging.sh`; agents: provider=aws; image=${AWS_IMAGE_UBUNTU_ARM_64}; instanceType=${AWS_ARM_INSTANCE_TYPE}; plugins: `elastic/vault-docker-login#v0.8.0`, `elastic/vault-docker-login#v0.8.0` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | Environment names referenced: ASDF_MAGE_VERSION, AWS_ARM_INSTANCE_TYPE, AWS_IMAGE_UBUNTU_ARM_64, GCP_DEFAULT_MACHINE_TYPE, GCP_HI_PERF_MACHINE_TYPE, GCP_WIN_MACHINE_TYPE, IMAGE_BEATS_WITH_HOOKS_LATEST, IMAGE_UBUNTU_X86_64, IMAGE_WIN_10, IMAGE_WIN_10_VERSION, IMAGE_WIN_11, IMAGE_WIN_11_VERSION, IMAGE_WIN_2016, IMAGE_WIN_2019, IMAGE_WIN_2022, PLATFORMS, RACE_DETECTOR, TEST_COVERAGE |
| Outputs | Artifact paths: `x-pack/osquerybeat/build/*.json`, `x-pack/osquerybeat/build/*.xml`, `x-pack/osquerybeat/build/integration-tests/*`, `x-pack/osquerybeat/build/integration-tests/Test*/*`, `x-pack/osquerybeat/build/integration-tests/Test*/data/**/*`; downstream pipelines: none declared |
| Permissions | Agent requirements are configured per step (see table).; Plugins are configured per step (see table). |
| Secrets | No secret values are documented; named secret access is configured in referenced plugins or external Buildkite settings |

## Operational notes

- One or more steps declare an explicit timeout.
- One or more steps declare automatic retry behavior.

## Sources

- Pipeline definition: [`.buildkite/x-pack/pipeline.xpack.osquerybeat.yml`](../x-pack/pipeline.xpack.osquerybeat.yml)
- Referenced script: [`.buildkite/scripts/packaging/packaging.sh`](../../.buildkite/scripts/packaging/packaging.sh)
- Referenced local dependency: [`.buildkite/scripts/docker.sh`](../../.buildkite/scripts/docker.sh)
- Referenced local dependency: [`.buildkite/scripts/qemu.sh`](../../.buildkite/scripts/qemu.sh)
