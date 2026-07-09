# iter-14 — Decisions

**D1 — The sim-embeddings digest drift is a KEY-algorithm change (M209), not a data change → pure re-key.**
The 2026-06-29 sim-embeddings cache is keyed `10146f281304c26de2444529e36cee96`; demo-1 probes
`032c99ea47678187631c59c31b4ef059`. sim-embeddings is a ROW surface → `capture.Surface.VersionTables()`
returns its own 4 tables, so its staleness digest is over ONLY `cms.similarities` +
`similarity_{categories,features,skills}`. M209 introduced this narrowing (the comment: "critical post
skiller-in-app merge, where a whole-schema digest would change on ANY app migration and thrash the cache").
The 2026-06-29 capture predates the narrowing → its key is the OLD wide/whole-schema digest. Demo-1's live
columns for all 4 tables match the cache manifest columns EXACTLY (names + order) — the empirical D1-gate
iter-02 established. So the data is faithful and only the KEY is stale → a pure re-key (hardlink + manifest
`schema_version` bump), no payload rewrite (unlike taxonomy, this surface is `cms.*`, untouched by
skiller→public). Replay COPY is self-validating (fails loudly on any type mismatch) — it succeeded rc 0.

**D2 — Re-key under demo-1's LIVE-probed narrowed digest so the fix is cold-durable.** Keyed the migrated
cache at `032c99ea…` (the digest a cold `/dev-up`/`/demo-up` will probe for this surface, identical cms
schema across dev+demo). So a future cold bring-up's set-dress `stacksnap replay --surface sim-embeddings`
HITs → the sim-embeddings step upgrades from `skipped(cache-miss)` to `replayed`. Non-fatal either way (the
bring-up already tolerates a sim-embeddings cache miss), but this closes the AI-sims library grid cold.
