---
title: "KB Fidelity Audit — M17 (rerun-safety)"
date: 2026-06-08
scope: milestone:M17
invoked-by: build-milestone
---

## Verdict
GREEN

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| snapshot replay write path + per-stack-isolation class | `corpus/ops/snapshot-spec.md` (§replay, §`stacksnap` CLI) | `stack-snapshot/replay/replay.go`, `stack-snapshot/pg/pg.go` (`CopyIn`), `stack-snapshot/cmd/stacksnap/main.go` (`replayCmd`) | PAIRED |
| seeder + `--reset` model + n=0 guard + casbin gotcha | `corpus/ops/seeding-spec.md` (§CLI, §n=0-dev guard, §casbin gotcha) | `stack-seeding/cmd/stackseed/main.go` (`resetTables`, `doReset`), `stack-seeding/seeders/identity.go`, `stack-seeding/seeders/activity.go`, `stack-seeding/pg/pg.go` | PAIRED |
| tooling write-side / read-side safety contract (TRUNCATE target class) | `corpus/ops/safety.md` (§2.1 store registry, §2.2 3-layer guard, §2.5 n=0 guards) | `stack-seeding/isolation/isolation.go`, `stack-seeding/isolation/audit.go` | PAIRED |
| migrate re-run safety (ISSUE-7 residual + race) | `corpus/ops/seeding-spec.md` (§global policy prerequisite); demo-stack `migrate-demo.sh` comments | `demo-stack/migrate-demo.sh`, `demo-stack/tests/test_tooling.py::TestMigrateRaceGuard` | PAIRED |
| **per-component re-run / idempotency contract** | **`corpus/ops/idempotency.md` (net-new — milestone `Delivers →`)** | replay + seed write paths above + the 4 bring-up scripts | **BLIND-AREA (delivered by this milestone)** |

## Fidelity Findings

1. **snapshot-spec.md — replay re-run behavior.** Source: `corpus/ops/snapshot-spec.md` §replay/§`stacksnap` CLI. Expected: doc describes replay as "verify checksums → bulk COPY → rebuild pgvector index"; makes **no idempotency claim**. Actual: `replay.Run` (`replay/replay.go:48-93`) does exactly this — bare `CopyIn` per table, no TRUNCATE/skip; not idempotent. Verdict: **ALIGNED** (the doc does not over-claim idempotency, so no stale claim — the milestone ADDS the guard + doc coverage). Fix owner: doc gains a re-run subsection as part of M17's delivery (`idempotency.md` + a snapshot-spec.md cross-link).

2. **seeding-spec.md — `--reset` truncate list.** Source: §CLI / §n=0 guard. Expected: doc says `--reset` does "per-stack reset (refuses n=0 dev unless --force)"; doesn't enumerate the truncate set. Actual: `resetTables` (`cmd/stackseed/main.go:28-32`) = {memberships, users, organizations} — STALE vs the actual seeded surface set (skips sessions/activity/casbin). Verdict: **ALIGNED at the doc level** (the doc doesn't enumerate the list, so it isn't a false claim) but the CODE is incomplete — fixed in-scope by M17 (extend the truncate list). The seeding-spec.md gains an explicit truncate-set note pointing at idempotency.md.

3. **safety.md §2.5 — n=0 guard scope + replay-has-no-N=0-guard.** Source: §2.5. Expected: "Snapshot replay has NO N=0 guard ... replaying the real catalog into the main dev stack is harmless (public data, own isolated Postgres)." Actual: `replayCmd` has no N=0 guard (confirmed). Verdict: **ALIGNED**. M17's replay re-run guard is a per-stack-isolated TRUNCATE-then-reload — it stays within the "own isolated Postgres / public data" envelope §2.5 already describes, so the §2.5 claim remains true. The new TRUNCATE will be pinned by a target-class test to per-stack-isolated only.

4. **safety.md §2.1/§2.2 drift guards (load-bearing).** The M15 drift guards (`isolation/safety_doc_drift_test.go` + `stack-snapshot/cmd/stacksnap/main_drift_test.go`) pin the exact predicate literals + guard symbol names + bounded-read SQL the doc quotes. Verdict: **ALIGNED** (baseline test pass green). M17 constraint: any safety.md edit must keep these green byte-for-byte.

## Completeness Gaps

1. **(incidental)** `CopyIn` / `CopyRows` are COPY primitives that cannot express `ON CONFLICT` (COPY has no upsert form). The seeder's `ON CONFLICT` guard for det-UUID rows + the casbin g2 grant must therefore route through `Exec` (INSERT … ON CONFLICT) or a staging+merge — a design note for M17's `decisions.md`, not a doc gap.

## Applied Fixes
None needed pre-implementation — all PAIRED topics ALIGNED; the one BLIND-AREA is the milestone's own delivered doc.

## Open Items (require user decision)
None.

## Gate Result
**GREEN: proceed.** Every dependency topic is PAIRED + ALIGNED; no stale load-bearing claim the milestone would read as truth; the single BLIND-AREA (`corpus/ops/idempotency.md`) is the milestone's explicit `Delivers →` deliverable, authored as M17 work.
