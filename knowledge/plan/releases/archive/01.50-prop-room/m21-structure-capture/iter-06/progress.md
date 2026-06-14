**Type:** tik (fifth tik under TOK-01) — code-ification, capture-side core.

# M21 iter-06 — progress

## Work done (Go, in the rosetta-extensions authoring copy)
1. **manifest additive `Structure` field** (`StructureArtifact{Payload, SHA256, Statements}`) — optional +
   back-compatible (older manifests omit it), no FormatVersion bump (the Predicate precedent). Validation added.
2. **`directus/structure.go`** — `CaptureStructure(ctx, runner)` reads the directus content-model schema **dynamically**
   (every `directus.*` base table NOT `directus_*`) and assembles the ordered structure SQL: CREATE SEQUENCE →
   CREATE TABLE (prod-faithful `pg_catalog` types + defaults) → ALTER TABLE ADD PRIMARY KEY. PKs captured per M21-D9.
3. **`pg.QueryRowString`** — the text-blob runner the catalog queries use.
4. **Privilege-visibility alignment (M21-D10):** the queries intersect `pg_catalog` with `information_schema.columns`
   (the digest's privilege-filtered view). Verified on prod: `sim_tasks_criterion` + `sim_tasks_criterion_check` are
   `pg_class`-visible to the read role but absent from its `information_schema`, so a naive `pg_class` capture would
   over-capture by 2 (28 vs 26) and never converge the digest. Fixed → 26.
5. **Tests** (all green, 12 packages): assembly/ordering, no-sequences, empty-DDL guard, error propagation, query
   scoping (system-table exclusion + information_schema alignment), manifest additive-field round-trip + back-compat.
   Catalog SQL **live-validated on prod**: 26 tables / 8 sequences / 26 PKs.

## Re-measure
furthest-passing-stage stays **6 (demonstrated)** — this is a code-ification BUILD iter (the tooling toward
gate-met-by-automation), graded on deliverables landing, not an ordinal move. The metric flips to "met (automated)"
once iter-07 (apply) + iter-08 (serve rows) wire it through stacksnap. Committed: rosetta-extensions `2c42ed5`.

## Close — 2026-06-11

**Outcome:** Shipped the M21 capture-side structure extension core — the additive manifest artifact + the dynamic,
digest-aligned directus structure-SQL generator (DDL + PKs) + `pg.QueryRowString`, fully unit-tested + live-validated.
Found + fixed the privilege-visibility alignment subtlety (M21-D10).
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (capture core built; apply + serve-rows + wiring remain — iter-07/08)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (tik) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (3 tiks this session) — (6) protocol-stop: n — Outcome: continue
**Decisions:** M21-D10 (privilege-visibility alignment: capture must scope to the digest's information_schema view).
**Side-deliverables:** none separable (all in-scope code-ification core).
**Routes carried forward (→ iter-07/08 under TOK-01):**
  - `STRUCT-M21-iter07-apply` — wire `CaptureStructure` into the directus capture (store the artifact as a payload +
    set `manifest.Structure`) AND apply it in provision/replay BEFORE the row replay; redefine the exit-4/5 semantics
    so a structure-bearing replay provisions then loads. Live integration test (capture → fresh bootstrap → apply →
    digest converges → replay exit 0).
  - `STRUCT-M21-iter08-serve` — capture + apply the directus_collections registration rows + the public read
    permissions (the serve half, M21-D9) so a `stacksnap`-provisioned stack serves anonymously without hand SQL.
  - Carried: the firewall structural-metadata admissibility class (overview In-scope); `directus_files` ref capture; M23.
**Lessons:** a capture that must reproduce a privilege-filtered digest has to scope to the SAME privilege view
(information_schema), not the raw catalog (pg_class) — else it silently over-captures and never converges. Generalize
to any digest-keyed capture against a least-privilege read role.
