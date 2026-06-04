---
milestone: M7b
slug: data-dna
version: v1.1 "show floor"
milestone_shape: section
status: planned
created: 2026-06-04
last_updated: 2026-06-04
delivers: rosetta-extensions/stack-seeding/dna/ (the data-DNA manifest + catalog + drift check) + corpus/architecture/alignment_testing.md (the data dimension)
---

# M7b — The data-alignment dimension ("data DNA")

## Goal
Make seeding **drift-proof, trackable, and self-describing** by extending the v1.0 **M0 alignment framework** to
a new dimension — **data**. A machine-readable **data-DNA** (a) *enumerates the seedable surfaces* — becoming the
authoritative **catalog of seeders to build** (it drives M7c) — and (b) *measures each seeder's output conforms
to the platform's current schema*, so a platform schema change **surfaces as a DNA diff** instead of a silently
broken seeder. This is the user's core ask: "use the alignment pattern, this time the dimension is data; it will
help us list what seeder to implement and stop the data model from drifting out from under our tools."

## Why a section milestone (buildable)
The harness only needs the platform's ent schema (which exists) + M0's `alignctl` scaffolding (which exists). The
surface is enumerable — research counted ~30 entity types → ~8–10 core seeders. The `In:` list is writable now.
The novelty is bounded (new operators + introspection-as-source), not open-ended.

## How it extends M0 (honest about where the analogy holds / breaks — research 2026-06-04)
| M0 (Clerk, behavioral) | M7b (data, structural) |
|---|---|
| Capability = an endpoint/function | Capability = a **seedable surface** (e.g. `SeedUsers`) |
| Variant = an input/scenario class | Variant = a **scenario** (e.g. `org-1k-mixed-roles`) |
| Source = the live engine; input→output | **Source = the platform's ent schema, introspected** (one-sided: output→schema, no input) |
| Operators = exact / shape / normalized / error_class (**value** semantics) | **New operators** = FK-valid / constraint-satisfied / type-match / row-count (**structural** semantics) |
| Drift = re-score after a version bump (M1b) | Drift = **schema diff** — a new/changed/removed field or surface flags the catalog as stale |
| Score = aligned ÷ total genes | **Coverage** = catalogued surfaces with a passing conformance gene (the gate M7c is measured against) |

**Holds:** the enumerated genome *forces completeness + drift detection*; the manifest/diff/score machinery; the
record/replay idea (capture the schema via introspection, replay seeder output against it). **Breaks:** M0's
operators test value equivalence (two-sided); seeding tests structural conformance (one-sided). So M7b reuses
M0's manifest + enumeration + diff + score, and **adds the structural operators + schema-as-source** — genuinely
new machinery, which is why it earns its own milestone rather than a footnote in M7a.

## Scope
### In
- **The data-DNA manifest format**: each gene = `{surface (capability), scenario (variant), expected-shape
  descriptor, conformance operator, criticality}` — the `clerk-*.json` analogue, for data.
- **The seedable-surface catalog**: the ~8–10 core surfaces enumerated as capabilities (users · orgs/
  memberships/casbin · features/tiers · taxonomy · content/library · skillpath sessions · jobsim sessions+results
  · assignments · activity), with isolation-class cross-checked against M7a's guard. **This catalog is the
  checklist M7c builds against.**
- **The structural conformance operators**: FK-valid, constraint-satisfied, type-match, row-count — the place
  M0's value-operators don't fit.
- **Schema-as-source via introspection**: capture the platform's current expected shape per surface (ent schema
  introspection / golden schema snapshot) as the gene's denominator — refreshable when the platform moves.
- **Drift detection**: a `data-dna diff` (the M1b analogue) that flags when the platform adds/changes/removes a
  surface or field → the catalog is now incomplete/stale; **CI-gateable** (exit-code contract like
  `drift-check.sh`).
- **Reuse `alignctl` where it fits** (manifest load, enumeration, diff, score scaffolding); add the data
  operators + the introspection source where it doesn't.
- **Delivers** the data dimension into `corpus/architecture/alignment_testing.md` + the manifest/catalog/drift
  check under `rosetta-extensions/stack-seeding/dna/`.

### Out
- The seeder modules themselves (M7c implements against this catalog).
- The framework/orchestrator/safety guard (M7a).
- Re-architecting M0's Clerk DNAs (they stay behavioral; M7b is additive — a new dimension, not a change).
- AI-content fidelity (out of v1.1 entirely).

## Depends on
**M7a** (the seeder contract to enumerate + measure against) + **M0** (the alignment framework, in rosetta —
`test/alignment/alignctl` + `corpus/architecture/alignment_testing.md`). **Parallel with:** none (gates M7c by
producing its catalog + gate).

## Estimated complexity
**medium** — bounded new machinery (operators + introspection + diff) layered on existing scaffolding; the
risk is getting the introspection-as-source right so the drift diff is trustworthy.

## Open questions (resolve during build)
- Ent-schema introspection (parse the generated ent code / `atlas inspect`) vs a hand-curated golden schema snapshot.
- Where the data-DNA lives: `rosetta-extensions/stack-seeding/dna/` (alongside the seeders) vs `test/alignment/`
  (alongside the Clerk DNAs) — likely the former (data is an extensions concern; Clerk DNAs are rosetta's).
- Whether the conformance operators reuse `alignctl`'s operator interface or need a sibling runner.

## KB dependencies (read as contract)
- `corpus/architecture/alignment_testing.md` (the M0 model M7b extends — capability×variant, DNA, operators, score, diff)
- M0's `test/alignment/alignctl` source (the manifest/diff/score scaffolding to reuse)
- `corpus/services/{backend,skiller,jobsimulation,skillpath,cms}.md` (the ent schemas the catalog enumerates)
- M7a's `corpus/ops/seeding-spec.md` (the seeder contract + isolation classes the catalog cross-checks)

## Delivers → `corpus/architecture/alignment_testing.md` (the data dimension) + `rosetta-extensions/stack-seeding/dna/`
The "alignment, applied to data" section (capability/variant/operator/score reinterpreted structurally) + the
data-DNA manifest, the seedable-surface catalog, and the drift-check contract.

## Exit (section)
The data-DNA manifest + the full seedable-surface catalog exist and validate; the structural operators run; a
`data-dna diff` provably flags an injected schema change (a renamed/added column on a seeded surface); the
**coverage metric** is defined as M7c's gate — documented in `alignment_testing.md` + the extension's `dna/`.
