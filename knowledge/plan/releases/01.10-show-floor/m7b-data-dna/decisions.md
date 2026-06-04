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

## Open (resolve during build)
- Ent-introspection vs `atlas inspect` golden for schema-as-source.
- DNA home: `rosetta-extensions/stack-seeding/dna/` vs `test/alignment/` (likely the former).
- Reuse `alignctl`'s operator interface vs a sibling `datadnarun` runner.
- The coverage threshold for M7c's gate (start ~90%).
