---
title: "KB Fidelity Audit — M20 Lifecycle convergence"
date: 2026-06-09
scope: milestone:M20
invoked-by: build-milestone
---

## Verdict
GREEN

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| Set-dress chaining (ISSUE-10) — demo-up runs replay→seed | `corpus/ops/snapshot-spec.md` §"Dev as a full-fidelity peer (M13)", `corpus/ops/safety.md` §2.5 | `dev-stack/dev-setdress.sh` (the proven M13 pass to reuse), `demo-stack/up-injected.sh` (chain site) | PAIRED |
| Atomicity / re-run safety (M17 guards make retry safe) | `corpus/ops/idempotency.md` | `stack-snapshot/replay/`, `stack-seeding` (TRUNCATE-then-reload, idempotent COPY) | PAIRED |
| Cold-start capture-source policy (DSN-export / dump-restore) | `corpus/ops/snapshot-spec.md` §"The capture-source policy", `corpus/ops/safety.md` §1.4 | `stack-snapshot/source/source.go`, `stack-snapshot/cmd/stacksnap` | PAIRED |
| Cold-start fresh-box WORKFLOW + MCP limitation (ISSUE-13) | — (net-new `corpus/ops/snapshot-cold-start.md`) | (the workflow doc + optional MCP spike) | BLIND-AREA → milestone deliverable |
| demo auto-set-dress preset (`small-200` over `dev-min`) | `corpus/ops/seeding-spec.md` §"The shipped presets" | `stack-seeding/presets/small-200.seed.yaml` | PAIRED |
| demo-up / demo-down skills + recipe family | `.claude/skills/demo-up/SKILL.md`, `corpus/ops/demo/README.md` + recipes | `demo-stack/up-injected.sh`, `rosetta-demo` | PAIRED |

## Fidelity Findings

1. **Capture-source policy (ALIGNED).** `snapshot-spec.md`'s capture-source precedence table (cache-hit → dump-ingest [default] → primary-read [fallback] → restore-from-snapshot / read-replica [upgrades]) matches `source.go` byte-for-byte: `DefaultPrecedence` lists the same four `Kind`s in the same order; `Available()` returns true only for `dump-ingest`/`primary-read`; the bounded session SQL (`SET TRANSACTION READ ONLY` + the three timeouts) matches the doc's code block. The "no offline pg_dump-FILE reader" claim (M9b-D9) is enforced — `cmd/stacksnap` requires `--dsn` and offers no `--dump` file path. **Verdict: ALIGNED.** This is the load-bearing contract the cold-start doc will describe; it is accurate.

2. **M13 set-dress pass (ALIGNED).** `snapshot-spec.md` §"Dev as a full-fidelity peer" describes the `dev-setdress.sh` three-step pass (per-stack Directus recipe + firewall-check → cache-first replay [non-fatal on stale/miss] → dev-min light seed) and the escapes (`--no-snapshot`, `--no-setdress`, n=0 guard). The code matches: `dev-setdress.sh` implements exactly that sequence, the n=0 hard-refuse-without-`--force` guard, and the non-fatal cache-miss warning. **Verdict: ALIGNED.** This is the pass M20 reuses verbatim for demo-up.

3. **Atomicity / re-run safety (ALIGNED).** `idempotency.md` (M17) + `snapshot-spec.md`/`seeding-spec.md` document replay's per-stack-isolated TRUNCATE-then-reload and the idempotent seed COPY + casbin `WHERE NOT EXISTS`. These make a re-run of the auto-chained set-dress safe. **Verdict: ALIGNED.**

4. **Safety contract n=0 + firewall guards (ALIGNED).** `safety.md` §2.5 precisely scopes the n=0 guard to auto-set-dress + `--reset` (replay has none, correctly). §1.4 documents the read-only bounded capture. M20 must preserve these byte-for-byte; the docs accurately state the boundary M20 cannot cross. **Verdict: ALIGNED.**

## Completeness Gaps

1. **The cold-start fresh-box workflow is undocumented (BLIND-AREA, by design).** ISSUE-13 describes the situation — a fresh box has no `~/.pgpass`, no staging dump, and the wired `postgres` MCP is a query tool not a `--dsn` `stacksnap` can `COPY` through — but there is no knowledge doc walking an operator from "fresh box" to "real catalog replayed". This is precisely the milestone's net-new deliverable (`overview.md` `Delivers → corpus/ops/snapshot-cold-start.md`), so it is **NOT a blocker** — it is the first work of the milestone, authored in full. The capture-source policy it builds on (Finding 1) is already accurate.

## Applied Fixes
None needed inline — all PAIRED topics are ALIGNED. No stale claims, no broken cross-references found in the M20-scope docs.

## Open Items (require user decision)
None.

## Gate Result
**GREEN: proceed.** The one BLIND-AREA (the cold-start fresh-box workflow doc) is already an explicit milestone deliverable per `overview.md`'s `Delivers →` line, authored as the milestone's first section — not a stale or unanchored gap. Every load-bearing contract M20 builds on (capture-source policy, the M13 set-dress pass, the M17 re-run guards, the safety n=0/firewall boundaries) is verified accurate against code. Implementation may enter Phase 1.
