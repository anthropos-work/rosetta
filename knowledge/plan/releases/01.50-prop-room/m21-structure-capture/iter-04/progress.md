**Type:** tik (third tik under TOK-01, with the M21-D7 option-A direction).

# M21 iter-04 — progress

## Work done (live, against Docker + sanctioned prod MCP read)
1. **Option-A feasibility confirmed:** prod's 27-system-table digest = `b4cb55bcee08c76f2c37980da460a683` = a freshly
   bootstrapped `directus/directus:11.6.1` exactly → **no version skew** (the pinned image's system tables match prod).
2. **Captured all 26 collections' faithful DDL** (the 9 from iter-03 + the 17 others) + the 8 junction-table sequences,
   via bounded read-only `pg_catalog` reads. Assembled as `iter-04/structure.sql` (the validated structure artifact).
3. **Applied the structure to a fresh bootstrapped harness** (11.6.1, schema `directus`): 8 `CREATE SEQUENCE` + 26
   `CREATE TABLE` → 53 tables total → **schema digest = `6cd35278edbc8a7962053a9d7ebfc480` = the prod cache key
   EXACTLY**. Stage 3 (structure-apply + digest convergence) **PASSES**.
4. **`stacksnap replay --surface directus --stack dev-5` → EXIT 0**: `9 table(s) cleared, 9 table(s), 10128 row(s)
   loaded`. Verified the real public content landed: simulations=**304** (= the manifest's public-sim count),
   skill_paths=22, roles=953, sequences=304. Stage 4 (replay-exit-0) **PASSES**.

## Re-measure
- Pre-iter furthest-passing-stage: **2**. Post-iter: **4** (stage 3 digest-converge + stage 4 replay exit 0,
  demonstrated end-to-end with the real prod structure). Delta **+2** — met the planned 2→4 lift.
- **Honest caveat:** the structure-apply was done by hand-applying the captured real-DDL artifact (`structure.sql`)
  via psql. The exit GATE requires `stacksnap` ITSELF to apply the captured structure — so the CODE-ification (the
  capture-side structure extension + provision automation) is the remaining stage-3 tooling work (iter-05+). What is
  proven: the Option-A path (bootstrap → apply 26-collection structure → digest converges → replay exits 0 with real
  rows) works end-to-end. → M21-D8.

## Close — 2026-06-11

**Outcome:** Option A validated end-to-end: 26-collection structure apply converges the digest to `6cd35278…` and
`stacksnap replay` exits 0 loading the real public content (10128 rows, simulations=304). furthest-passing-stage
**2 → 4**.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (gate = stage 6, serve anonymously; also requires the structure-apply to be stacksnap-automated)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (tik; and this iter broke the no-prog streak) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (1 tik this session) — (6) protocol-stop: n — Outcome: continue
**Decisions:** M21-D8 (Option A validated end-to-end through stage 4; the manual-apply→code-ification caveat).
**Side-deliverables:** the `iter-04/structure.sql` artifact (real 26-collection public content-model DDL + sequences).
**Routes carried forward (→ iter-05+ under TOK-01):**
  - `STRUCT-M21-iter05-serve` — stages 5–6: load the registry rows (directus_collections/fields/relations) + an
    anonymous public read permission + boot Directus + `GET /items/simulations?limit=1` → 200 (the gate). The
    Directus permission model is the milestone's flagged live-only risk.
  - `STRUCT-M21-codeify` — build the stacksnap capture-side structure extension (capture the 26-collection DDL +
    registry over `--dsn` into the snapshot; apply before row replay in provision) so the GATE's "stacksnap applies
    the captured structure" is met by the tooling, not a hand-applied artifact.
  - Carried: `directus_files` wiring; M23 referential closure.
**Lessons:** for a whole-schema-digest cache, convergence is a pure column-structure match — proven by creating all 26
collection tables and watching the digest land on the prod key. The registry ROWS don't affect the digest (they're
system-table rows, not columns), so stage 4 (cache-hit + replay) decoupled cleanly from stages 5–6 (serve, which need
the registry + permissions). Splitting the gate this way let one tik bank a +2 lift.
