# beats-tests-pipeline

Buildkite pipeline definition for beats-tests-pipeline.

## At a glance

| | |
| --- | --- |
| Platform | Buildkite |
| Source | [`.buildkite/beats-tests-pipeline.yml`](../beats-tests-pipeline.yml) |
| Trigger | Not specified in this checked-in YAML; start conditions may be configured in Buildkite. |

## Purpose

This pipeline defines `Triggering Build for Auditbeat`, `Triggering Build for Heartbeat`, `Triggering Build for Filebeat`, `Triggering Build for x-pack/filebeat` and 15 additional steps for `beats-tests-pipeline`. The source does not state a separate motivation.

## When it runs

- Not specified in this checked-in YAML; start conditions may be configured in Buildkite.

## Execution flow

| Step | Condition | Depends on | Configured work |
| --- | --- | --- | --- |
| `Triggering Build for Auditbeat` | `No step condition declared` | None declared | Triggers downstream pipeline `auditbeat` |
| `Triggering Build for Heartbeat` | `No step condition declared` | None declared | Triggers downstream pipeline `heartbeat` |
| `Triggering Build for Filebeat` | `No step condition declared` | None declared | Triggers downstream pipeline `filebeat` |
| `Triggering Build for x-pack/filebeat` | `No step condition declared` | None declared | Triggers downstream pipeline `beats-xpack-filebeat` |
| `Triggering Build for x-pack/dockerlogbeat` | `No step condition declared` | None declared | Triggers downstream pipeline `beats-xpack-dockerlogbeat` |
| `Triggering Build for Metricbeat` | `No step condition declared` | None declared | Triggers downstream pipeline `beats-metricbeat` |
| `Triggering Build for x-pack/metricbeat` | `No step condition declared` | None declared | Triggers downstream pipeline `beats-xpack-metricbeat` |
| `Triggering Build for x-pack/osquerybeat` | `No step condition declared` | None declared | Triggers downstream pipeline `beats-xpack-osquerybeat` |
| `Triggering Build for x-pack/winlogbeat` | `No step condition declared` | None declared | Triggers downstream pipeline `beats-xpack-winlogbeat` |
| `Triggering Build for Deploy/K8S` | `No step condition declared` | None declared | Triggers downstream pipeline `deploy-k8s` |
| `Triggering Build for Libbeat` | `No step condition declared` | None declared | Triggers downstream pipeline `beats-libbeat` |
| `Triggering Build for pkg/systemmetrics` | `No step condition declared` | None declared | Triggers downstream pipeline `beats-pkg-systemmetrics` |
| `Triggering Build for x-pack/libbeat` | `No step condition declared` | None declared | Triggers downstream pipeline `beats-xpack-libbeat` |
| `Triggering Build for x-pack/otel` | `No step condition declared` | None declared | Triggers downstream pipeline `beats-xpack-otel` |
| `Triggering Build for x-pack/auditbeat` | `No step condition declared` | None declared | Triggers downstream pipeline `beats-xpack-auditbeat` |
| `Triggering Build for x-pack/heartbeat` | `No step condition declared` | None declared | Triggers downstream pipeline `beats-xpack-heartbeat` |
| `Triggering Build for x-pack/packetbeat` | `No step condition declared` | None declared | Triggers downstream pipeline `beats-xpack-packetbeat` |
| `Triggering Build for Winlogbeat` | `No step condition declared` | None declared | Triggers downstream pipeline `beats-winlogbeat` |
| `Triggering Build for Packetbeat` | `No step condition declared` | None declared | Triggers downstream pipeline `beats-packetbeat` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | Environment names referenced: BUILDKITE_BRANCH, BUILDKITE_COMMIT |
| Outputs | Artifact paths: none declared in this file; downstream pipelines: `auditbeat`, `beats-libbeat`, `beats-metricbeat`, `beats-packetbeat`, `beats-pkg-systemmetrics`, `beats-winlogbeat`, `beats-xpack-auditbeat`, `beats-xpack-dockerlogbeat`, `beats-xpack-filebeat`, `beats-xpack-heartbeat`, `beats-xpack-libbeat`, `beats-xpack-metricbeat`, `beats-xpack-osquerybeat`, `beats-xpack-otel`, `beats-xpack-packetbeat`, `beats-xpack-winlogbeat`, `deploy-k8s`, `filebeat`, `heartbeat` |
| Permissions | Buildkite agent/plugin access is not detailed beyond this YAML |
| Secrets | No secret values are documented; named secret access is configured in referenced plugins or external Buildkite settings |

## Operational notes

- This page describes only the checked-in YAML; runtime Buildkite settings and externally defined plugins may add behavior not represented here.

## Sources

- Pipeline definition: [`.buildkite/beats-tests-pipeline.yml`](../beats-tests-pipeline.yml)
