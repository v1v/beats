# post-dependabot

Workflow definition for **post-dependabot**. This page documents the checked-in GitHub Actions configuration; it does not assert that a run succeeds.

## At a glance

| | |
| --- | --- |
| Platform | GitHub Actions |
| Source | [`.github/workflows/post-dependabot.yml`](../post-dependabot.yml) |
| Trigger | push: branches=dependabot/go_modules/** |

## Purpose

This workflow provides the `post-dependabot` CI task: it prepares or invokes `actions/checkout@v7`, `actions/setup-go@v7` and then runs `make notice`, `echo "modified=$(if git status --porcelain --untracked-files=no | grep -q -E ' NOTICE.txt$'; then echo "true"; else`, `git config --global user.name 'dependabot[bot]'`, `git config --global user.email 'dependabot[bot]@users.noreply.github.com'`. The checked-in source does not state a separate motivation.



## When it runs

- push: branches=dependabot/go_modules/**

Event and job filters are evaluated as shown in the workflow source.

## Execution flow

| Job | Runner | Condition | Depends on | Configured work |
| --- | --- | --- | --- | --- |
| `update-notice` | `ubuntu-latest` | `Always when its event is eligible` | None declared | `actions/checkout@v7`; `actions/setup-go@v7`; update NOTICE.txt — runs `make notice`; check for modified NOTICE.txt — runs `echo "modified=$(if git status --porcelain --untracked-files=no \| grep -q -E ' NOTICE.txt$'; then echo "true"; else`; commit NOTICE.txt — runs `git config --global user.name 'dependabot[bot]'`; `git config --global user.email 'dependabot[bot]@users.noreply.github.com'` |
| `manage-pr-reviewers` | `ubuntu-latest` | `Always when its event is eligible` | None declared | Add reviewers (obs-infraobs-integrations) if PR has Team:Obs-InfraObs label — runs `PR_NUMBER=$(gh pr list --repo ${{ github.repository }} --head ${{ github.ref_name }} --json number --jq '.[0].numbe`; `if [ -n "$PR_NUMBER" ] && [ "$PR_NUMBER" != "null" ]; then` |

## Inputs, outputs, and access

| Item | Details |
| --- | --- |
| Inputs | No workflow inputs declared; environment names: GH_TOKEN |
| Outputs | No job outputs declared; artifact/cache actions: none declared |
| Permissions | contents: read; update-notice.contents: write; manage-pr-reviewers.pull-requests: write; manage-pr-reviewers.repository-projects: read |
| Secrets | GITHUB_TOKEN |

## Operational notes

- No additional retry, timeout, or deployment behavior is declared in the checked-in workflow.

## Sources

- Workflow: [`.github/workflows/post-dependabot.yml`](../post-dependabot.yml)
- External action/workflow reference: `actions/checkout@v7` (implementation is outside this checkout).
- External action/workflow reference: `actions/setup-go@v7` (implementation is outside this checkout).
