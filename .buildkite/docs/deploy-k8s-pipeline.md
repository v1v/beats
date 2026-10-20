# deploy-k8s-pipeline

Buildkite pipeline definition for deploy-k8s-pipeline.

## At a glance

| | |
| --- | --- |
| Platform | Buildkite |
| Source | [`.buildkite/deploy/kubernetes/deploy-k8s-pipeline.yml`](../deploy/kubernetes/deploy-k8s-pipeline.yml) |
| Trigger | Not specified in this checked-in YAML; start conditions may be configured in Buildkite. |

## Purpose

This pipeline runs the configured `deploy-k8s-pipeline` work, organized into `Deploy/K8S`. The source does not state a separate motivation.

## When it runs

- Not specified in this checked-in YAML; start conditions may be configured in Buildkite.

## Execution flow

| Step | Condition | Depends on | Configured work |
| --- | --- | --- | --- |
| `Deploy/K8S / Deploy/K8S: Run pre-commit` | `No step condition declared` | None declared | Runs `pre-commit run --all-files`; agents: image=${IMAGE_BEATS_WITH_HOOKS_LATEST}; cpu=2000m; memory=4Gi; useCustomGlobalHooks=true |
| `Deploy/K8S / Checks` | `No step condition declared` | None declared | Runs `make -C deploy/kubernetes all`; `make check-no-changes`; agents: provider=gcp; image=${IMAGE_UBUNTU_X86_64}; machineType=${GCP_HI_PERF_MACHINE_TYPE} |
| `Deploy/K8S / K8S Test/K8S version: {{matrix}}` | `No step condition declared` | None declared | Configured step; agents: provider=gcp; image=${IMAGE_UBUNTU_X86_64}; machineType=${GCP_HI_PERF_MACHINE_TYPE} |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | Environment names referenced: ASDF_KIND_VERSION, ASDF_MAGE_VERSION, GCP_HI_PERF_MACHINE_TYPE, IMAGE_BEATS_WITH_HOOKS_LATEST, IMAGE_UBUNTU_X86_64, K8S_VERSION, MODULE |
| Outputs | Artifact paths: none declared in this file; downstream pipelines: none declared |
| Permissions | Agent requirements are configured per step (see table). |
| Secrets | No secret values are documented; named secret access is configured in referenced plugins or external Buildkite settings |

## Operational notes

- One or more steps declare automatic retry behavior.

## Sources

- Pipeline definition: [`.buildkite/deploy/kubernetes/deploy-k8s-pipeline.yml`](../deploy/kubernetes/deploy-k8s-pipeline.yml)
