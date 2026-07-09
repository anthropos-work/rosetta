---
iter: 14
milestone: M211
iteration_type: tik
status: planned
created: 2026-07-08
---

# iter-14 — tik: fill sim-embeddings (cache re-key) → M42e employee coverage GREEN

## Active strategy reference
**TOK-01** ("Warm-first cache-migrate, then cold-prove both stacks") — move (4) "iterate warm-first,
prove cold". This tik closes the last M42e-employee residual on the live demo-1 substrate.

## Step 0 — Re-survey (done)
M42e employee coverage vs demo-1 is **ONE section short**: escapes=0, personaFailures=0,
crossPortFollowFails=0, notReached=0, reachable=62. The sole failing section is `/library/ai-simulations`
(the `sim-card-grid` is empty). Live probe confirms **`cms.similarities` = 0 rows** on demo-1 (the
pgvector simulation-similarity index behind `searchSimulations`), so the AI-sims grid renders empty.

## Cluster / target identified
The **sim-embeddings** snapshot surface (`cms.similarities` + `cms.similarity_{categories,features,skills}`)
was **skipped during demo-1 set-dress** (bring-up log line 224: `cache miss: no snapshot for
sim-embeddings/032c99ea47678187631c59c31b4ef059`). Root cause (analogous to iter-02's taxonomy re-key):
- The captured cache (2026-06-29, 274 public sims + 278/274/664 child rows) is keyed at digest
  **`10146f281304c26de2444529e36cee96`** — computed under the OLD (pre-M209) whole-schema digest scoping.
- M209 narrowed row-surface digests to their OWN `VersionTables()` (4 sim tables) so unrelated cms/app
  migrations don't thrash the cache. The merged demo-1 stack now probes the narrowed digest
  **`032c99ea47678187631c59c31b4ef059`** → cache miss → surface skipped.
- **Empirical D1-gate PASSED:** demo-1's live-introspected columns for all 4 tables match the cache
  manifest columns EXACTLY (names + order). Payloads are all `cms.*` (no skiller→public rewrite needed —
  the surface never touched the merged schema). Embeddings + IDs are stable across the schema-prefix merge.

## Hypothesis
A **pure cache re-key** (hardlink the 4 payloads under the new digest dir + transform the manifest
`schema_version` → `032c99ea…`, payload bytes unchanged → sha256 checksums hold) makes `stacksnap replay
--surface sim-embeddings --stack demo-1` HIT → COPY loads 274 public sims into `cms.similarities` →
`searchSimulations` vector-scans real data → the `/library/ai-simulations` grid populates → M42e GREEN.
The replay COPY is self-validating (fails loudly on any type mismatch) — the same safety guarantee iter-02
relied on. Cold-durable: a future cold `/dev-up`/`/demo-up` probes the same `032c99ea…` digest → HIT.

## Expected lift
M42e employee coverage: failingSections **1 → 0** → **GATE: MET** for the employee vantage (all of
escapes/persona/crossPort/notReached already 0). Sub-condition (e) partial (employee half of coverage).

## Phase plan
1. Re-key the sim-embeddings cache `10146f28…` → `032c99ea…` (hardlink payloads + transform manifest).
2. `stacksnap replay --surface sim-embeddings --stack demo-1` → expect rc 0, ~274 sims + child rows loaded,
   reindexed `cms.similarities.small_embedding3`.
3. Verify `cms.similarities` count = 274 (public sims) on demo-1.
4. Re-run the M42e employee coverage sweep (reap-safe, in-turn) → expect failingSections 1→0, GATE MET.
5. Route the fix durably (the cache re-key is the fix; document the pattern for the demo-up log's
   sim-embeddings step so a future cold bring-up hits).

## Escalation conditions
- If the D1 column-match had FAILED (it did not) → user-blocker (a genuine schema divergence, not a re-key).
- If replay COPY fails on a type mismatch → investigate (would indicate a real merge shape change, not a
  stale key) → user-blocker.
- If the grid stays empty AFTER `cms.similarities` populates → a frontend/federation gap, not data →
  route to the documented empty-query→`publicJobSimulations` demopatch fallback (per orchestrator brief).

## Acceptable close-no-lift outcomes
If replay populates `cms.similarities` but the grid still needs a frontend demopatch (federation can't
serve the vector search), close-fixed-partial with the data landed + the demopatch routed — the data gap
is closed either way.
