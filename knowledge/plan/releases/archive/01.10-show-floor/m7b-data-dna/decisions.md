# M7b — Decisions

## M7b-D1 — extend M0 to a data dimension, don't fork it (design, 2026-06-04)
The user asked whether the Clerk alignment pattern can measure data. Research verdict: **partial hold** — reuse
M0's manifest/enumeration/diff/score, but data needs **new structural operators** (FK-valid / constraint /
type-match / row-count) and a **schema-as-source** (introspection, one-sided output→schema) because seeding tests
structural conformance, not behavioral value equivalence. Decision: M7b is **additive** — a new dimension on the
existing framework, not a change to the Clerk DNAs (which stay behavioral + green). This earns its own section
milestone because the new machinery is real, not a footnote in M7a.

## M7b-D2 — the data-DNA is the catalog that drives M7c (design, 2026-06-04)
The enumerated genome doubles as (a) the authoritative list of seeders to implement (M7c's checklist) and (b)
each seeder's acceptance gene. The **coverage metric** (catalogued surfaces with a passing conformance gene)
becomes M7c's exit gate. This is what makes the fleet *trackable* and *drift-proof*: a new platform surface shows
up as an un-covered gene; a changed schema shows up as a diff.

## M7b-D3 — a separate `datadna` harness, not a new alignctl runner (build, 2026-06-04)
The data dimension reuses M0's manifest/score/diff *structure* but its operators are **structural** (type-match,
constraint-satisfied [NOT-NULL + UNIQUE], fk-valid, row-count) and the measurement is **one-sided** (seeder
output → live schema, not source vs mirror). That's different enough that it's a standalone harness in
`rosetta-extensions/stack-seeding/dna/` + the `datadna` CLI (reusing the M7a `pg` layer, connecting directly to
the stack's Postgres), NOT a new runner under `alignctl`. The conceptual link is documented in
`corpus/architecture/alignment_testing.md` § "The data dimension (M7b)".

## M7b-D4 — the live proof caught a real bug + drove a harden (proof, 2026-06-04)
Proven live against the M7a-seeded `demo-1`: `introspect` (seeded-only) → `measure` **100% / Critical 100%**
across the 4 seeded surfaces → `diff` flags an injected column (exit 1) and reads clean on revert. Findings:
1. **`introspect` was populating the PLANNED surfaces** (where the table exists, e.g. `skiller.skills`), which
   made the DNA self-contradictory (`Validate` requires planned surfaces to have an empty shape) → `measure`/
   `diff` aborted. Fixed: `introspect` now captures **seeded surfaces only**; planned stay empty until M7c
   promotes them.
2. **The UNIQUE leg of `constraint-satisfied` was unwired** (the `PgIntrospector` didn't implement `DupLister`)
   — `measure` silently checked "0 UNIQUE". Harden wired `pg.CountDuplicatesSQL` + `PgIntrospector.CountDuplicates`;
   the UNIQUE leg now activates (verified: "1 UNIQUE column(s) satisfied" on the seeded clerk_id).
Also corrected `TestShippedManifest` — it hardcoded "memberships has 2 FKs" but the live table has 4 (the seeder
writes 2); the exact count is precisely what `diff` tracks, so the test now asserts the stable org+user FKs.

## Open (resolve during build)
- Ent-introspection vs `atlas inspect` golden for schema-as-source.
- DNA home: `rosetta-extensions/stack-seeding/dna/` vs `test/alignment/` (likely the former).
- Reuse `alignctl`'s operator interface vs a sibling `datadnarun` runner.
- The coverage threshold for M7c's gate (start ~90%).
