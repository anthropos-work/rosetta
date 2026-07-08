---
milestone: M211
slug: bringup-acceptance
version: v2.1 "quick change"
milestone_shape: iterative
status: planned
created: 2026-07-08
last_updated: 2026-07-08
complexity: large
depends_on: M209, M210
exit_gate: "/dev-up AND /demo-up both GREEN cold on the merged platform — 4-subgraph compose / no skiller container; snapshot recapture→replay loads public.* (taxonomy replay rc 0, ~42,763 public skills); seed resolves real public node-ids (closure green); verify passes with a merged-platform assertion (no skiller schema/subgraph/container); M42 coverage sweep + v2.0 Playthroughs suite GREEN; 0 residual skiller-schema references in any queried path"
iteration_protocol_ref: corpus/ops/verification.md (+ corpus/ops/demo/coverage-protocol.md + corpus/ops/demo/playthroughs.md gates)
delivers: proof the merged platform stands up end-to-end via the re-grounded tooling
---

# M211 — Bring-up acceptance: dev-up + demo-up green on the merged platform

## Goal
Prove the whole chain works end-to-end on the **merged platform** with the **re-grounded tooling** — a real user
can bring up a **dev** stack AND a **demo** stack, set-dressed + seeded + verified, with **zero skiller residue**.

## Exit gate (measurable)
From a re-synced state, **`/dev-up` AND `/demo-up` both go GREEN cold**:
- platform composes **4 subgraphs** / **no skiller container**;
- snapshot **recapture → replay** loads `public.*` (taxonomy replay exits **0**, ~**42,763** public skills);
- **seed** succeeds — the re-pointed taxonomy resolvers resolve real public node-ids, `datadna` closure GREEN;
- **verify** (`verification.md` net) passes with a **merged-platform assertion** (no skiller
  schema/subgraph/container expected; `readiness.sh` schema probe GREEN);
- the **M42 coverage sweep** + the **v2.0 Playthroughs suite** stay GREEN (10 live Playthroughs on cold
  reset-to-seed);
- **0 residual skiller-schema references** in any path the tooling queries.

## Why iterative (not section)
The merged 4-subgraph platform has **never been stood up locally** with the re-grounded tooling. Bring-up *will*
surface unknown fix-loops — migration ordering, the capture column-mapping caveat, vestigial container/clone
cleanup, cache-key behavior. A fixed `In:` checklist would be speculative; the exit gate is the commitment. Mirrors
v1.10b M53's cold-rebuild-acceptance role, but **iterative** because the merge is unverified at bring-up.

## Iteration protocol
The fit-up/dress-rehearsal **fix → re-measure → re-run bring-up** loop, driven by `corpus/ops/verification.md`
(the scoped non-fatal auto-verify net) + the `coverage-protocol.md` (presence) + `playthroughs.md` (function)
gates. Each tik: run the bring-up, triage a failure, route the fix to its surface (rext → M209-class change on the
authoring copy + re-tag; corpus → M210-class doc fix; stack → re-sync), re-measure. Close-on-gate.

## Three-fate note
Fixes surfaced mid-iter route per the three-fate rule. A surface that **cannot be driven without a platform-repo
edit ESCALATES** (the `unimplementable-without-platform-edit` state) — it never edits the platform. The platform
already did the merge; v2.1 stays tooling + docs + re-sync only.

## Running ledger
_(iter closeouts append below — see progress.md)_
