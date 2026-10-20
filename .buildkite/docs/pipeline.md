# pipeline

Buildkite pipeline definition for pipeline.

## At a glance

| | |
| --- | --- |
| Platform | Buildkite |
| Source | [`.buildkite/pipeline.yml`](../pipeline.yml) |
| Trigger | Not specified in this checked-in YAML; start conditions may be configured in Buildkite. |

## Purpose

This is the repository-level Buildkite fan-out: on pull requests it triggers component pipelines when their configured paths change, while the serverless-beats pipeline is triggered for non-pull-request builds. The source does not state a separate motivation.

## When it runs

- Not specified in this checked-in YAML; start conditions may be configured in Buildkite.
- The top-level component triggers require a pull request and use `if_changed` include patterns. A separate serverless-beats trigger is conditioned on not being a pull request.

## Execution flow

| Step | Condition | Depends on | Configured work |
| --- | --- | --- | --- |
| `Trigger Auditbeat` | `build.pull_request.id != null; path filter: auditbeat/**, .buildkite/auditbeat/**, .go-version, go.mod, pytest.ini, dev-tools/**…` | None declared | Triggers downstream pipeline `auditbeat` |
| `Trigger Heartbeat` | `build.pull_request.id != null; path filter: heartbeat/**, .buildkite/heartbeat/**, .go-version, go.mod, pytest.ini, dev-tools/**…` | None declared | Triggers downstream pipeline `heartbeat` |
| `Trigger Filebeat` | `build.pull_request.id != null; path filter: filebeat/**, .buildkite/filebeat/**, .go-version, go.mod, pytest.ini, dev-tools/**…` | None declared | Triggers downstream pipeline `filebeat` |
| `Trigger x-pack/filebeat` | `build.pull_request.id != null; path filter: filebeat/**, x-pack/filebeat/**, .buildkite/x-pack/pipeline.xpack.filebeat.yml, x-pack/libbeat/**, x-pack/otel/**, .go-version…` | None declared | Triggers downstream pipeline `beats-xpack-filebeat` |
| `Trigger x-pack/dockerlogbeat` | `build.pull_request.id != null; path filter: x-pack/dockerlogbeat/**, .buildkite/x-pack/pipeline.xpack.dockerlogbeat.yml, x-pack/libbeat/**, .go-version, go.mod, pytest.ini…` | None declared | Triggers downstream pipeline `beats-xpack-dockerlogbeat` |
| `Trigger Metricbeat` | `build.pull_request.id != null; path filter: metricbeat/**, .buildkite/metricbeat/**, .go-version, go.mod, pytest.ini, dev-tools/**…` | None declared | Triggers downstream pipeline `beats-metricbeat` |
| `Trigger x-pack/metricbeat` | `build.pull_request.id != null; path filter: metricbeat/**, x-pack/metricbeat/**, .buildkite/x-pack/pipeline.xpack.metricbeat.yml, x-pack/libbeat/**, x-pack/otel/**, .go-version…` | None declared | Triggers downstream pipeline `beats-xpack-metricbeat` |
| `Trigger x-pack/osquerybeat` | `build.pull_request.id != null; path filter: x-pack/osquerybeat/**, .buildkite/x-pack/pipeline.xpack.osquerybeat.yml, x-pack/libbeat/**, .go-version, go.mod, pytest.ini…` | None declared | Triggers downstream pipeline `beats-xpack-osquerybeat` |
| `Trigger x-pack/winlogbeat` | `build.pull_request.id != null; path filter: winlogbeat/**, x-pack/winlogbeat/**, .buildkite/x-pack/pipeline.xpack.winlogbeat.yml, x-pack/libbeat/**, .go-version, go.mod…` | None declared | Triggers downstream pipeline `beats-xpack-winlogbeat` |
| `Trigger Deploy/K8S` | `build.pull_request.id != null; path filter: deploy/kubernetes/**, metricbeat/module/kubernetes/**, .go-version, libbeat/docs/version.asciidoc, .buildkite/deploy/kubernetes/**, .buildkite/pipeline.yml…` | None declared | Triggers downstream pipeline `deploy-k8s` |
| `Trigger Libbeat` | `build.pull_request.id != null; path filter: libbeat/**, .buildkite/libbeat/**, .go-version, go.mod, pytest.ini, dev-tools/**…` | None declared | Triggers downstream pipeline `beats-libbeat` |
| `Trigger pkg/systemmetrics` | `build.pull_request.id != null; path filter: pkg/systemmetrics/**, .buildkite/pkg/systemmetrics/**, .go-version, go.mod, dev-tools/**, .buildkite/pipeline.yml…` | None declared | Triggers downstream pipeline `beats-pkg-systemmetrics` |
| `Trigger x-pack/libbeat` | `build.pull_request.id != null; path filter: x-pack/libbeat/**, .buildkite/x-pack/pipeline.xpack.libbeat.yml, .go-version, libbeat/**, go.mod, pytest.ini…` | None declared | Triggers downstream pipeline `beats-xpack-libbeat` |
| `Trigger x-pack/otel` | `build.pull_request.id != null; path filter: x-pack/otel/**, .buildkite/x-pack/pipeline.xpack.otel.yml, .go-version, go.mod, dev-tools/**, pkg/systemmetrics/**…` | None declared | Triggers downstream pipeline `beats-xpack-otel` |
| `Trigger x-pack/auditbeat` | `build.pull_request.id != null; path filter: auditbeat/**, x-pack/auditbeat/**, .buildkite/x-pack/pipeline.xpack.auditbeat.yml, x-pack/libbeat/**, .go-version, go.mod…` | None declared | Triggers downstream pipeline `beats-xpack-auditbeat` |
| `Trigger x-pack/heartbeat` | `build.pull_request.id != null; path filter: heartbeat/**, x-pack/heartbeat/**, .buildkite/x-pack/pipeline.xpack.heartbeat.yml, x-pack/libbeat/**, .go-version, go.mod…` | None declared | Triggers downstream pipeline `beats-xpack-heartbeat` |
| `Trigger x-pack/packetbeat` | `build.pull_request.id != null; path filter: packetbeat/**, x-pack/packetbeat/**, .buildkite/x-pack/pipeline.xpack.packetbeat.yml, x-pack/libbeat/**, .go-version, go.mod…` | None declared | Triggers downstream pipeline `beats-xpack-packetbeat` |
| `Trigger Winlogbeat` | `build.pull_request.id != null; path filter: winlogbeat/**, .buildkite/winlogbeat/**, .go-version, go.mod, pytest.ini, dev-tools/**…` | None declared | Triggers downstream pipeline `beats-winlogbeat` |
| `Trigger Packetbeat` | `build.pull_request.id != null; path filter: packetbeat/**, .buildkite/packetbeat/**, .go-version, go.mod, pytest.ini, dev-tools/**…` | None declared | Triggers downstream pipeline `beats-packetbeat` |
| `Serverless beats tests(elastic-agent)` | `build.pull_request.id == null` | None declared | Triggers downstream pipeline `beats-agent-serverless-tests` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | Environment names referenced: BUILDKITE_BRANCH, BUILDKITE_COMMIT, BUILDKITE_PULL_REQUEST, BUILDKITE_PULL_REQUEST_BASE_BRANCH, GITHUB_PR_LABELS |
| Outputs | Artifact paths: none declared in this file; downstream pipelines: `auditbeat`, `beats-agent-serverless-tests`, `beats-libbeat`, `beats-metricbeat`, `beats-packetbeat`, `beats-pkg-systemmetrics`, `beats-winlogbeat`, `beats-xpack-auditbeat`, `beats-xpack-dockerlogbeat`, `beats-xpack-filebeat`, `beats-xpack-heartbeat`, `beats-xpack-libbeat`, `beats-xpack-metricbeat`, `beats-xpack-osquerybeat`, `beats-xpack-otel`, `beats-xpack-packetbeat`, `beats-xpack-winlogbeat`, `deploy-k8s`, `filebeat`, `heartbeat` |
| Permissions | Buildkite agent/plugin access is not detailed beyond this YAML |
| Secrets | No secret values are documented; named secret access is configured in referenced plugins or external Buildkite settings |

## Operational notes

- The root pipeline conditionally uploads component builds based on changed paths; downstream test/package behavior is defined in the referenced pipeline YAML files.

## Sources

- Pipeline definition: [`.buildkite/pipeline.yml`](../pipeline.yml)
