---
iter: 03
milestone: M21
iteration_type: tik
status: closed-fixed-partial
created: 2026-06-11
---

# M21 iter-03 — tik (second tik under TOK-01): resolve the structure source

Under TOK-01. Planned target: resolve the structure-source question (iter-02's M21-D4/D6 blocker) and produce the real
9-collection structure artifact to advance furthest-passing-stage 2 -> 3.

## Active strategy reference
**TOK-01** (staged-pipeline; structure artifact applied before row replay).

## Re-survey (Phase 1 Step 0)
furthest-passing-stage = 2 (live-confirmed). The iter-02 routes (STRUCT-M21-iter03-source/-artifact) are untouched
and meaningful. The structure-source blocker (M21-D6) gated the artifact.

## What happened (plan evolved on evidence)
1. **Investigated self-contained sources** — `stack-dev/cms/internal/directus/collections/*.go` are a lossy read-side
   app-view (field names + Go types + relational aliases, not Postgres types/registry). Self-contained options
   exhausted -> M21-D6 -> surfaced the operator-gated source decision.
2. **Operator sanctioned a prod structural read** (the wired `postgres` MCP; the directus schema is in the same prod
   Postgres). Blocker RESOLVED. Saved as a standing policy memory.
3. **Captured the real faithful structure** for all 9 collections via bounded read-only structural reads:
   exact `pg_catalog` DDL (uuid/json/text/varchar(N)/timestamptz; matches the manifest column counts), and the
   registry inventory: 9 `directus_collections` / 217 `directus_fields` / 43 `directus_relations` (20 dangling to the
   17 uncaptured collections -> M23 referential closure).
4. **Decisive digest finding (M21-D5 resolved toward option B):** prod digest `6cd35278…` = the FULL 53-table directus
   schema (27 system + **26** user collections). Our surface captures **9** of 26. A bootstrapped stack + a
   9-collection structure can NEVER converge the full-schema digest -> the bootstrap+partial-collection model can only
   ever cache-HIT if the cache is **re-keyed per-surface** over only the captured content tables (option B), not by
   the whole-schema digest. This is the stage-4 architectural fork surfaced for decision.

## Outcome vs target
furthest-passing-stage stays **2**: the structure SOURCE is now in hand (real DDL + registry inventory) and the
stage-4 keying question is resolved in principle, but the artifact is not yet APPLIED to a live harness (the apply +
the per-surface re-key implementation is iter-04). closed-fixed-partial.

## Routes carried forward (-> iter-04, Fate-3 under TOK-01)
- `STRUCT-M21-iter04-apply` — build the structure artifact (DDL + registry rows captured over the sanctioned source)
  and apply it to a fresh bootstrapped harness; confirm the 9 tables + registry -> stage 3.
- `STRUCT-M21-digest-keying` — **operator chose option A (M21-D7):** keep whole-schema keying; converge by applying
  ALL 26 collections + pinning the Directus version (NOT the per-surface re-key). Surfaced + decided this iter.
- `STRUCT-M21-iter03-artifact` (carried) + `directus_files` wiring + M23 referential-closure of the 20 dangling relations.
