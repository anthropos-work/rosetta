# M7b — Spec notes

Seeded from the 2026-06-04 redesign research; fill in during build.

## The reuse map (M0 → data)
Reuse from `test/alignment/alignctl`: the DNA manifest loader, the gene enumeration, the `dna diff` machinery, the
weighted score/coverage rollup, the record/replay (golden) idea. Add new: the **structural operators** + the
**schema-as-source** capture.

## The structural operators (new)
- `fk-valid` — every foreign key in the seeded row resolves to an existing parent row.
- `constraint-satisfied` — NOT NULL / UNIQUE / CHECK / enum domains hold.
- `type-match` — column types + nullability match the platform's current schema.
- `row-count` — the surface produced ≥ the scenario's expected cardinality (e.g. 1k users).
These test STRUCTURE, not VALUES — a seeded user need not be a *real* person, but it must be a *schema-valid* user.

## Schema-as-source (the gene denominator)
Two candidate capture paths (decide during build):
1. Parse the platform's generated ent schema (the `internal/data/ent/schema/*.go` + generated migrations).
2. `atlas inspect` the migrated per-stack DB → a canonical schema snapshot (golden), refreshed on platform bump.
The drift diff compares the captured snapshot to the DNA's recorded expectation → added/changed/removed surfaces+fields.

## The seedable-surface catalog (from research — the M7c checklist)
~8–10 core surfaces: Users · Orgs/Memberships/Casbin · Features/Tiers · Taxonomy (skills/jobroles) · Content/Library
· SkillPath sessions · JobSim sessions+results · Assignments · Activity/Academy-progress. Each carries its
isolation class (per M7a's guard) + its seeding primitive. ~30 entity types underneath.

## Drift as a flagged event (the M1b analogue, for data)
`data-dna diff` exit-code contract (mirror `drift-check.sh`): 0 = no drift · 1 = schema moved (catalog stale) ·
2 = a conformance gene regressed · 3 = usage. CI-gateable on platform bumps so "the data model drifted and our
seeders broke" becomes a flagged, scored event instead of a silent failure.

## To confirm during build
- Whether `alignctl`'s operator interface is extensible enough to host structural operators, or M7b needs a
  sibling runner (`datadnarun`) like the Clerk runners.
- The coverage threshold that becomes M7c's gate (start ~90% of catalogued surfaces conformance-passing).
