---
name: test-platform
description: Verify a running Anthropos platform with black-box probes (HTTP, GraphQL, Connect-RPC, DB, Playwright), drive each repo's own test suite, and inventory test coverage. Writes a timestamped markdown report to `.agentspace/test-platform/`. Use when you want to know what's actually working — and what isn't — across the whole stack without touching service internals.
argument-hint: [scope: 'live' | 'repos' | 'census' | 'full']
---

# Anthropos Platform Verification

This skill runs three orthogonal verifications against the Anthropos platform and produces one consolidated report:

1. **Live verification** — black-box probes against the running stack
2. **Repo tests** — invoke each platform repo's own test suite
3. **Census** — read-only inventory of test files per repo (development-health signal)

Pick the scope you want via the argument: `live`, `repos`, `census`, `full`. Default is `live` (fast — useful as a quick "is the platform up?" check).

## Your Mission

1. **Decide scope**. Default to `live`. If the user passes `repos`, `census`, or `full`, honor it.
2. **Confirm pre-conditions** (see below) before running anything that takes more than a few seconds.
3. **Run the underlying tooling** — `stack-verify/reports/generate.sh <scope>` from the `rosetta-extensions` clone (see *How to invoke*). Do not reinvent the probes — the bash scripts under `stack-verify/` are the source of truth.
4. **Surface the report** to the user: print the path, summarize pass/fail, point at any 0-test repos or service failures, and suggest follow-up actions.

## Scope semantics

| Scope | What runs | Typical duration | When to use |
|---|---|---|---|
| `live` (default) | Liveness + readiness probes (`stack-verify/live/verify.sh`) | seconds | Quick "is the platform up?" check, after `make up` |
| `repos` | Each platform repo's own test suite via `stack-verify/repos/run.sh` | minutes (10-30+) | Pre-commit / post-update verification of test-suite health |
| `census` | Test-file inventory via `stack-verify/census/inventory.sh` | seconds | "Which repos lack tests?" development-health audit |
| `full` | All three sequentially | 10-30+ min | Full health check, daily / pre-release |

## Pre-conditions per scope

| Scope | Requires | Verify with |
|---|---|---|
| `live` | Platform running (`make ps` shows containers up) | `cd stack-dev/platform && make ps` |
| `repos` | All non-studio repos cloned (`make init`); language toolchains installed (Go, pnpm 10.x, Node 24, npm) | `ls stack-dev/` and `node -v` |
| `census` | All repos cloned (read-only — no toolchain needed) | `ls stack-dev/` |
| `full` | All of the above | — |

If `make ps` shows the platform is down and the user asked for `live` or `full`, **ask** whether to run `/start-platform` first instead of probing a dead stack.

## Confirmation Policy

**Proceed WITHOUT confirmation**:
- Running `live` or `census` (both read-only, seconds)
- Reading existing reports under `.agentspace/test-platform/`

**ASK for confirmation before**:
- Running `repos` or `full` (long-running; pulls compute on the user's machine)
- Re-running a scope that just ran (the previous report may already answer the question)

## How to invoke

The probes/runners live in the **`rosetta-extensions`** repo, section
`stack-verify/` (rosetta keeps only this skill). Locate the toolkit, then run its
driver with two env vars:

- `STACK_ROOT` — the stack being tested (the dir that holds `platform/`), e.g. `stack-dev`.
- `REPORT_DIR` — where to write the report. Use rosetta's `.agentspace/test-platform/`
  so reports land where this skill has always written them.

```bash
ROSETTA=/Users/kirality/Dropbox/Workspaces/swarm/rosetta
# Verification toolkit: prefer the target stack's own clone; else the authoring copy.
VERIFY="$ROSETTA/.agentspace/rosetta-extensions/stack-verify"
# (or "$ROSETTA/stack-dev/rosetta-extensions/stack-verify" once the dev stack has its own clone)
# If neither exists: git clone https://github.com/anthropos-work/rosetta-extensions.git \
#   --branch v1.2.0 "$ROSETTA/.agentspace/rosetta-extensions"

STACK_ROOT="$ROSETTA/stack-dev" \
REPORT_DIR="$ROSETTA/.agentspace/test-platform" \
  bash "$VERIFY/reports/generate.sh" <scope>
```

The script:
- Runs the underlying probes in order
- Writes `$REPORT_DIR/op_YYYYMMDD_HHMMSS_<scope>.md` (the human report)
- Also writes `op_YYYYMMDD_HHMMSS_<scope>.raw.txt` (raw stderr/stdout for failure forensics)
- Returns exit code 0 on full pass, 1 on any failure, 2 if anything was skipped due to missing tools / missing checkout

You report the **markdown path** to the user and quote the headline metrics (pass/fail counts, any flagged repos).

## Report Structure

The generated report has these sections (only those relevant to the chosen scope):

1. **Header** — date, scope, overall status, git branch / SHA, host
2. **Live verification** — liveness table + readiness table per service
3. **Repo test suites** — pass / fail / skip table per repo, with log paths
4. **Test census** — per-repo unit / integ / e2e / CI counts with health flag (`ok` / `no-tests` / `no-ci` / `not-cloned`)
5. **Notes** — summary of what to do next

## Critical Rules

- **Scope boundary**: the probes speak each service's **external interface only** — HTTP, GraphQL, Connect-RPC, psql, redis-cli, Playwright. Never import service internals into `stack-verify/`. If a check would require touching internals, it belongs in that service's own test suite, invoked by `scope=repos`.
- **Tooling home**: the probes live in `rosetta-extensions/stack-verify/`, not in rosetta. New/changed probes are built and tested in the `.agentspace/rosetta-extensions/` authoring copy and tagged; a stack runs them from its pinned clone. Never hand-write probe scripts into the rosetta corpus.
- **No duplication**: do not re-implement what a service already tests. The `repos` scope exists precisely to delegate to each repo's runner.
- **No mutations**: probes are read-only. The census never executes code. The repo runner invokes each repo's own runner (which may write to a local DB — that's expected for integration tests, but `repos` scope should NOT be run against shared infra).
- **Report only**: the skill produces a report. It does not commit anything, push anything, or fix anything. Fixes are a separate conversation with the user.

## Error Handling

1. **`make ps` shows nothing**: tell the user the platform isn't running and offer `/start-platform`.
2. **A specific service is down**: do not retry or restart. Report it and let the user decide.
3. **A repo test suite fails**: capture the log path from the report and quote the last 10 lines to the user. Do not attempt to fix the test — that's a per-repo PR.
4. **Missing toolchain (no Go, no pnpm, etc.)**: the runner marks the repo as `skipped` automatically. Surface this and recommend installation.

## Adding new probes or new repos

These edits happen in the `.agentspace/rosetta-extensions/` authoring copy (then commit + tag), never in rosetta:

* **New service**: edit `stack-verify/lib/services.sh` (registry row) + optionally `stack-verify/lib/readiness.sh` (deeper probe) + call the new readiness function from `stack-verify/live/verify.sh`.
* **New repo**: edit the `TEST_CMD` map in `stack-verify/repos/run.sh` and the `should_skip` logic if it needs a new toolchain.
* **New e2e flow**: add a `.spec.ts` under `stack-verify/e2e/tests/`. Keep it unauthenticated — authenticated flows belong to next-web-app's own E2E suite.

## Anti-patterns to refuse

- "Run /test-platform repos in production" — refuse; this script invokes per-repo runners which may exercise local DBs.
- "Add a probe that mutates state" — refuse; probes are read-only by design.
- "Have /test-platform fix failures" — refuse; this skill is read-only and reporting-only. Fix in a separate PR.

## Additional Resources

All in the `rosetta-extensions` clone, section `stack-verify/`:
- `stack-verify/README.md` — layout overview + the `STACK_ROOT`/`REPORT_DIR` contract
- `stack-verify/lib/services.sh` — current service registry
- `stack-verify/lib/readiness.sh` — readiness probe functions
- `stack-verify/reports/generate.sh` — top-level driver
