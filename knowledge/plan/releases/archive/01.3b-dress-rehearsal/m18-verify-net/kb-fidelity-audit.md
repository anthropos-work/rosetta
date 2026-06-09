---
title: "KB Fidelity Audit — M18 (The verification safety net)"
date: 2026-06-09
scope: milestone:M18
invoked-by: build-milestone
---

## Verdict
GREEN

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Auto-verify contract (default-on + non-fatal at bring-up tail) | `corpus/ops/verification.md` (NET-NEW — overview `Delivers →` owns it) | `demo-stack/up-injected.sh`, `dev-stack/dev-stack`, `stack-verify/` | DOC-ONLY (pre-promoted deliverable) |
| Offset/scope-aware `stack-verify` | `corpus/ops/rosetta_demo.md` §registry; `stack-core/README.md`; `stack-verify/README.md` | `stack-verify/lib/services.sh`, `live/verify.sh`, `lib/readiness.sh`, `repos/run.sh`, `census/inventory.sh` | PAIRED |
| Unified registry (recorded ports = the offset source) | `corpus/ops/rosetta_demo.md` §"Unified stack registry + first-available-N" | `stack-core/stack_registry.py` | PAIRED |
| `/test-platform` skill semantics (the verify surface the new doc relates to) | `.claude/skills/test-platform/SKILL.md`; `stack-verify/README.md` | `stack-verify/reports/generate.sh` | PAIRED |
| The `$DEVDIR` bug + cheap-win asserts (ISSUE-7 catcher) | `.agentspace/demo-up-issue.md` ISSUE-12/14 (diagnosis of record) | `stack-verify/repos/run.sh:~108`, `census/inventory.sh:~75`; `demo-stack/up-injected.sh` tail | PAIRED (diagnosis) |

## Fidelity Findings

1. **Registry schema claim — ALIGNED.** `rosetta_demo.md` §registry states the per-stack record is
   `{type: dev|demo, n, ports, status, created}`, "pure runtime (gitignored), flock-guarded, atomic writes."
   Verified against `stack_registry.py` (`_write` is temp+os.replace; `_locked` is `fcntl.flock(LOCK_EX)`;
   `.stacks/` is gitignored). Matches.

2. **"verify is operator-driven only" — ALIGNED (no stale claim to contradict).** A grep across
   `corpus/ops/{rosetta_demo,idempotency,safety,README}.md` + root `CLAUDE.md` finds NO claim that any
   bring-up auto-runs verify or a health probe today. The docs correctly reflect the current state
   (`/test-platform` is the only verify surface, manual). M18 ADDS the auto-run — it does not correct a
   stale promise. (This is exactly the gap ISSUE-12a/ISSUE-14 record.)

3. **`stack-verify` env contract — ALIGNED.** `stack-verify/README.md` + the `/test-platform` SKILL both
   document `STACK_ROOT` (defaults to the stack owning the clone) and `REPORT_DIR`. Matches the scripts'
   `${STACK_ROOT:-...}` resolution. (Note: the scripts' INTERNAL `$DEVDIR` reference in `repos/run.sh` +
   `census/inventory.sh` is the latent bug M18 fixes — it is not a doc claim; the doc correctly says `STACK_ROOT`.)

## Completeness Gaps

- **`corpus/ops/verification.md` is a BLIND-AREA that is PRE-PROMOTED to a milestone deliverable.** The
  M18 `overview.md` `Delivers → rosetta` line explicitly names `corpus/ops/verification.md` (net-new). Per
  Phase 5, a blind area whose doc is already a declared `Delivers →` deliverable does NOT block — the
  knowledge production IS part of the milestone. No action needed beyond authoring it during the build
  (the §"verification.md authored" section in `progress.md`).

## Applied Fixes
None required — all PAIRED topics are ALIGNED; the one blind area is a pre-promoted deliverable.

## Open Items (require user decision)
None.

## Gate Result
GREEN — proceed to Phase 1. Every dependency topic is PAIRED + ALIGNED; the net-new `verification.md` is a
declared milestone deliverable (not an unfilled gap). SEVERITY=clear.
