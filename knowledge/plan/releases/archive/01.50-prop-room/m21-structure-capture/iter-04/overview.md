---
iter: 04
milestone: M21
iteration_type: tik
status: closed-fixed
created: 2026-06-11
---

# M21 iter-04 — tik (third tik under TOK-01): apply Option-A structure → converge digest → replay exit 0

Under TOK-01, with the operator's **M21-D7 (option A)** direction. Target: advance furthest-passing-stage **2 → 4**.

## Active strategy reference
**TOK-01** (staged-pipeline; structure applied before row replay), refined by **M21-D7**: converge the digest by making
the per-stack schema MATCH prod's — create ALL 26 user-collection tables + the pinned 11.6.1 system tables.

## Re-survey (Phase 1 Step 0)
furthest-passing-stage = 2 (live-confirmed). The iter-03 route `STRUCT-M21-iter04-apply` is untouched. **Option-A
feasibility CONFIRMED this iter:** prod's 27-system-table digest = `b4cb55bc…` = a bootstrapped 11.6.1 exactly (no
version skew). So bootstrap 11.6.1 + create all 26 collection tables ⇒ full digest should converge to `6cd35278…`.

## Cluster / target identified
The digest (`pg.SchemaVersionSQL`) is over **column structure** (information_schema.columns), NOT registry rows. So
stage 4 (replay cache-hit + COPY) needs only the **26 collection tables to exist with prod-faithful columns** — the
directus_collections/fields/relations registry rows are needed for SERVING (stages 5–6), not for the digest/replay.
This iter therefore attacks stages 3→4 via the table DDL alone.

## Hypothesis
On a fresh bootstrapped 11.6.1 harness: create all 26 collection tables from the real prod DDL → the directus-schema
digest becomes `6cd35278…` → `stacksnap replay --surface directus` finds the cache HIT and COPYs the 9 captured
content tables → **exits 0** (stage 4). The cached rows are public-only + parent-consistent, and the DDL carries no
DB-level FK constraints (Directus relations are app-level), so the bulk COPY should not violate constraints.

## Expected lift
furthest-passing-stage **2 → 4** (structure-apply works AND replay exits 0). If a COPY type/constraint mismatch
surfaces, expect **2 → 3** (tables created + digest converges) with the replay issue characterized + routed.

## Phase plan (staged-pipeline tik)
1. Capture all 26 collections' faithful DDL (pg_catalog) + any referenced sequences over the sanctioned MCP read.
2. Fresh bootstrapped harness (11.6.1) → create the 26 tables → verify the digest == `6cd35278…` (stage 3).
3. `stacksnap replay --surface directus --stack dev-5` → expect exit 0, 9 tables / real rows loaded (stage 4).
4. Spot-verify the replayed rows (e.g. simulations count == 304).

## Escalation conditions
- If the digest does NOT converge after creating the 26 tables, diff the harness vs prod `information_schema.columns`
  to find the residual (extra/missing table or a type mismatch) → fix or characterize → route.
- If replay COPY fails on a specific table (type/null), characterize the column + route the per-table fix forward.

## Acceptable close-no-lift outcomes
- The 26-table apply + a documented digest residual (the exact tables/columns that don't converge) is a first-class
  characterization if convergence proves harder than the system-digest match predicts.

## Test discipline
`go build/vet/test ./...` for any stack-snapshot code touched; the harness is throwaway (torn down at close).
