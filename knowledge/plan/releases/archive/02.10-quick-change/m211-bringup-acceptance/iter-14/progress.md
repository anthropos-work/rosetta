**Type:** tik (sim-embeddings cache re-key → M42e employee coverage GREEN). Under TOK-01 move (4).

# iter-14 — tik progress

## Execution log
1. **Root-caused the sim-embeddings cache-miss** (bring-up log line 224). The 2026-06-29 cache is keyed at
   digest `10146f281304c26de2444529e36cee96` (OLD pre-M209 whole-schema scoping); demo-1 now probes the
   narrowed 4-table digest `032c99ea47678187631c59c31b4ef059` (M209 row-surface `VersionTables()` narrowing,
   so unrelated cms/app migrations don't thrash the cache). Confirmed live via `stacksnap capture --dry-run`:
   `DRY-RUN "sim-embeddings" @ schema 032c99ea… — firewall plan OK (public-only)`.
2. **D1 empirical column-match (PASSED):** demo-1's live columns for all 4 tables (`cms.similarities` 13 cols,
   `similarity_categories` 4, `similarity_features` 7, `similarity_skills` 3) match the cache manifest columns
   EXACTLY (names + order). Payloads are all `cms.*` (no skiller→public rewrite — the surface never touched
   the merged schema). → pure re-key, no fabrication, same D1-gate logic iter-02 used.
3. **Cache re-key:** new dir `…/sim-embeddings/032c99ea…/`; **hardlinked** the 4 `.copy` payloads (links=2 →
   zero extra disk, old cache preserved), transformed the manifest (`schema_version`→`032c99ea…`, migration
   note appended). Payload BYTES unchanged → the per-table sha256 checksums stay valid.
4. **Replay:** `stacksnap replay --surface sim-embeddings --stack demo-1` → **rc 0**, "4 table(s) cleared,
   4 table(s), **1490 row(s) loaded**, reindexed [cms.similarities.small_embedding3]".
5. **Verify DB:** `cms.similarities` = **274** (all `organization_id IS NULL AND entity_type='simulation'`);
   children categories 278 / features 274 / skills 664 — exactly the manifest row_counts.
6. **Verify query path (cheap pre-check):** `recommendedUserSimulations(categories:[…])` flipped from
   "jobRoleID, skillIds, or categories must be provided" → **"unauthorized"** (passed arg-validation, reached
   the data layer, only needs a logged-in user) — the vector-search path over the now-populated
   `cms.similarities` is live.
7. **Re-ran M42e employee coverage** (demo-1, Maya, reap-safe): **GATE: MET ✅** — reachable=62/150,
   **failingSections 1→0**, personaFailures=0, escapes=0, notReached=0, frontier=EXHAUSTED (14.0m, rc 0).

## Re-measurement (gate sub-conditions)
| Sub-condition | Pre-iter | Post-iter |
|---|---|---|
| (e) M42 coverage — **employee vantage** | 1 section short (`/library/ai-simulations` empty) | **GREEN** (failingSections 0, GATE MET) |
| (e) M42 coverage — manager vantage | NOT MET | NOT MET (next: iter-15) |
| (e) v2.0 Playthroughs | NOT MET | NOT MET (routed) |
**Metric:** M42e employee coverage residual sections **1 → 0** (delta −1, met expected lift; employee half
of sub-condition (e) closed). Overall composite gate still needs manager coverage + Playthroughs + cold /dev-up.

## Close — 2026-07-08

**Outcome:** Filled the sim-embeddings surface on demo-1 via a **pure cache re-key** (`10146f28…`→`032c99ea…`,
the M209 row-surface digest-narrowing class) → replay loaded 274 public sims (1490 rows) into
`cms.similarities` → the `/library/ai-simulations` grid populates → **M42e employee coverage GATE: MET**
(failingSections 1→0; escapes/persona/crossPort/notReached all already 0). Cold-durable: a future cold
`/dev-up`/`/demo-up` probes the same `032c99ea…` digest → HIT (upgrades the set-dress sim-embeddings step
from skipped(cache-miss) → replayed).
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (M42e employee GREEN, but the composite M211 gate still needs M42m manager coverage + v2.0 Playthroughs + cold /dev-up)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (tik with +progress) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (1 tik) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (the sim-embeddings digest drift is the M209 row-surface `VersionTables()` narrowing, NOT a data change — the 4 tables' columns match demo-1 exactly → pure re-key, no payload rewrite, safe by the same empirical-column-match D1-gate iter-02 established), D2 (the re-key is cold-durable — keyed under demo-1's live-probed narrowed digest so a future cold bring-up's set-dress replay HITs)
**Side-deliverables:** none (the cache re-key is the planned deliverable; gitignored `.agentspace/snapshots/`).
**Routes carried forward (Fate-3 → next iters this session):** M42m manager coverage + the drifted studio-url/public-website-url demopatch hash re-pin (iter-15); v2.0 Playthroughs (iter-16+); cold `/dev-up` (final).
**Lessons:** The sim-embeddings surface carries the SAME M209 re-key as taxonomy (iter-02) — any row surface captured before M209's `VersionTables()`-narrowing has a stale whole-schema cache key and needs a pure re-key to the narrowed digest. Unlike taxonomy, sim-embeddings needed NO payload rewrite (it's a `cms.*` surface untouched by the skiller→public move) — just the key migration. The re-key template generalizes: `hardlink payloads + bump manifest schema_version to the live-probed digest` re-keys ANY pre-M209 row-surface cache at $0.
