---
iter: 01
milestone: M211
iteration_type: tok
tok_flavor: bootstrap
status: closed-fixed
created: 2026-07-08
---

# iter-01 — Bootstrap tok: author the M211 bring-up-acceptance strategy

**Type:** tok (bootstrap) — iter-01 of M211, authors the FIRST strategy. Does NOT terminate the call;
the loop continues into iter-02 (a tik) under TOK-01.

## Job
Author the opening strategy the first batch of tiks will follow toward the exit gate:
`/dev-up` AND `/demo-up` both GREEN **cold** on the merged 4-subgraph platform, via the re-grounded
tooling — recapture→replay loads `public.*` (~42,790 skills), seed closure green, verify (merged
assertion), M42 coverage + v2.0 Playthroughs GREEN, 0 residual skiller-schema refs.

## Inputs consulted
- `overview.md` (exit gate + the two pre-surfaced Fate-3 sections: M25-D9 bring-up requirement; the
  recapture prerequisite — the latter OVERRIDDEN by the user's cache-migration decision, see decisions.md).
- Protocol docs: `corpus/ops/verification.md` (the non-fatal auto-verify net), `coverage-protocol.md`
  (presence gate), `playthroughs.md` (function gate).
- Orchestrator context block (M208/M209/M210 verified facts; the cache-migration user decision; the
  docker-reap warning; the two-repo boundary).
- Live recon (see decisions.md D1–D5 + spec-notes baseline).

## Baseline (distance-to-gate)
Gate is a COMPOSITE across ~6 sub-conditions, each must hold on BOTH dev + demo, COLD. Current read:
| Sub-condition | State at iter-01 |
|---|---|
| (a) 4-subgraph / no-skiller compose | **MET** — warm `docker ps`: 11 anthropos containers, no skiller container |
| (b) recapture→replay loads `public.*` (~42,790) | **NOT MET** — cache is stale skiller-keyed (`c75ce94…`) |
| (c) seed closure green (real public node-ids) | **NOT MET** — untested on merged |
| (d) verify passes w/ merged-platform assertion | **NOT MET** — rext code ready (M209), unrun |
| (e) M42 coverage + v2.0 Playthroughs GREEN | **NOT MET** — untested on merged |
| (f) 0 residual skiller-schema refs in queried paths | code+corpus clean (M209/M210); runtime-unconfirmed |

Starting value: **1/6 sub-conditions demonstrably met, warm-only.** Gate wants all-6 on BOTH stacks, cold.

## TOK-01 strategy (recorded in milestone-root decisions.md)
**"Warm-first cache-migrate, then cold-prove both stacks."** Consume the re-grounded tooling; land the
dev-side M25-D9 pre-migrate hook (demo path already solves it); execute the user's cache-migration
(re-key `skiller.*→public.*` after empirical column-match); iterate the WARM merged stack to green
(reset-db → extensions-bootstrap → migrate → replay → seed → verify) for speed; then prove a full COLD
`/dev-up` + `/demo-up`; then M42 coverage + v2.0 Playthroughs.

## Next-tik direction (iter-02, first tik)
Target sub-condition **(b)**: taxonomy replay loads `public.*` (~42,790) into the WARM merged stack.
Prep steps folded in (all in service of the one measurable outcome): re-pin consumption to
`quick-change-m209`; empirically verify cached `skills` column-set == merged `public.skills`; land the
dev-side extensions-bootstrap + PG-wait pre-migrate hook if the dev path needs it (model:
`migrate-demo.sh`); execute the cache-migration (re-key manifest schema/payload/filter/public_via +
rename 10 payload files + resolve the new narrowed-digest cache key); run reset-db → migrate → stacksnap
taxonomy replay; measure replay rc (target 0) + `public.skills` count (target ~42,790).

## Phase plan (protocol: overview §Iteration protocol + verification.md)
Each tik: run a bring-up phase → triage the first failure → route the fix to its surface (rext
authoring-copy = M209-class + re-tag; corpus = M210-class doc fix; stack = re-sync) → re-measure.
Close-on-gate.

## Escalation conditions
- A surface that cannot be driven without a platform-repo edit → `unimplementable-without-platform-edit`
  ESCALATE (never edit the platform).
- Genuine column drift making the cache-migration unreconcilable → `user-blocker` (user decides: prod
  DSN or gate-partial). Do NOT fabricate rows.
