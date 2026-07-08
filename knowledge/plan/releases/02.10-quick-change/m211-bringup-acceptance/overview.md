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
exit_gate: "/dev-up AND /demo-up both GREEN cold on the merged platform — 4-subgraph compose / no skiller container; snapshot recapture→replay loads public.* (taxonomy replay rc 0, ~42,790 public skills); seed resolves real public node-ids (closure green); verify passes with a merged-platform assertion (no skiller schema/subgraph/container); M42 coverage sweep + v2.0 Playthroughs suite GREEN; 0 residual skiller-schema references in any queried path"
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
- snapshot **recapture → replay** loads `public.*` (taxonomy replay exits **0**, ~**42,790** public skills);
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

## Pre-surfaced bring-up requirement (Fate-3 from M208's live de-risk)
M208's live containerized de-risk already surfaced one concrete cold-bring-up fix-loop (the **M25-D9 class**),
pinned here so M211's first tik doesn't re-discover it:
- **Bootstrap the `extensions` schema before `make migrate` on a clean DB** — `CREATE SCHEMA extensions;
  CREATE EXTENSION vector SCHEMA extensions; CREATE EXTENSION pg_trgm SCHEMA extensions;` — with the
  `extensions.gin_trgm_ops` opclass **resolvable** (search_path handling). Without it, a clean `make reset-db`
  + migrate fails: app `20260518125439` + cms `20250116133510` (`extensions.vector(1536)` columns) →
  `schema "extensions" does not exist`, and app `20260623090000` (GIN-trigram index) can't resolve
  `extensions.gin_trgm_ops`. Not a merge defect — a bring-up-ordering prerequisite the merged taxonomy needs.
- **Add a PG-readiness wait before migrate on a cold bring-up** — reset-db currently races Postgres
  (`connection reset by peer` on the first migrate pass; a re-run succeeds).

Both are bring-up-tooling requirements (M25-D9 class). See M208 `spec-notes.md`/`decisions.md` Finding 1;
the `extensions.`-qualified capture column list is also cross-referenced in M209's Risk-2 (`overview.md` §In).

## Pre-surfaced recapture prerequisite (Fate-3 from M209's code re-ground)
M209 re-grounded the snapshot tooling to capture/replay `public.*` (const flip + digest-narrow + verified
column list + the ~42,790 `skills` MinRows floor) and confirmed it builds/tests GREEN — but it **could not run
the actual recapture**: no valid COPY-byte capture source is provisioned locally (values-blind-checked — no
`marco_read`/prod-read Postgres DSN in `.agentspace/secrets` or `platform/.env`; the merged `stack-dev` Postgres
holds **0** taxonomy rows so it can't be a source; the wired `postgres` MCP returns JSON rows, not COPY bytes —
see `corpus/ops/snapshot-cold-start.md`). So the existing cache
(`.agentspace/snapshots/taxonomy/c75ce94…`) is the **stale pre-merge skiller-keyed capture** — structurally
wrong for a merged stack and cache-key stale (M209 narrowed the digest, so the key changed too).
**M211's first tik must recapture** `public.*` via the sanctioned `stacksnap capture --surface taxonomy --dsn
<source>` path (a restored prod `pg_dump` **table-scoped** to the taxonomy tables — `-t public.skills …` — since
schema-scoped `-n public` is no longer small; or the `marco_read` prod read endpoint when provisioned), under
`AssertPublicOnly`, then verify replay exits 0 + ~42,790 public skills (the MinRows floor auto-catches an
empty/under-capture). This is a **data** prerequisite; the tooling CODE is ready.

## Three-fate note
Fixes surfaced mid-iter route per the three-fate rule. A surface that **cannot be driven without a platform-repo
edit ESCALATES** (the `unimplementable-without-platform-edit` state) — it never edits the platform. The platform
already did the merge; v2.1 stays tooling + docs + re-sync only.

## Running ledger
_(iter closeouts append below — see progress.md)_
