# M11 — Spec notes

## Pre-flight audits — section 1 (presets)
- KB-fidelity audit (Phase 0b): **GREEN** — `knowledge/plan/releases/01.20-set-dressing/m11-richer-recipes/kb-fidelity-audit.md`.
- Reused across all M11 sections: same subsystem (corpus/ops/demo + extensions presets/snapshot), no load-bearing
  knowledge-doc moved between sections.

## Topic → doc → code triples (for fast future audits)
- presets → `corpus/ops/demo/README.md` → `rosetta-extensions/stack-seeding/presets/*.seed.yaml` + `blueprint/blueprint.go`
- recipes → `corpus/ops/demo/recipe-*.md` → operator flow over `stackseed` + `stacksnap`
- snapshot CLI → `corpus/ops/snapshot-spec.md` § "The `stacksnap` CLI" → `stack-snapshot/cmd/stacksnap/{main,surfaces}.go`
- data-DNA 100% → `corpus/ops/seeding-spec.md` data-DNA § → `stack-seeding/{dna,seeders/{taxonomy,content}_snapshot.go}`

## The snapshot/seed flow (the architecture the recipes + skill must reflect)
`stacksnap replay --surface taxonomy|directus` runs **out-of-band, BEFORE `stackseed`** (per-stack reference
data, not org-scoped). The seeder's `TaxonomySnapshotSeeder` / `ContentSnapshotSeeder` DAG nodes only **verify**
the replay happened (count public rows) and are the ordering/linkage anchor: `… → taxonomy/content (snapshot) →
sessions/assignments (which LINK content refs) → activity`. The preset `stack.seed.yaml` does NOT carry snapshot
config — snapshots are stack-global, replayed separately. So a "full-fidelity" curator flow is:
`/demo-up N → /demo-snapshot replay N (taxonomy + directus) → /demo-seed N --preset … → login`.

## Refreshed presets (section 1)
The preset YAMLs gain a documented **snapshot prerequisite** (a header comment block) so a curator knows the
real-catalog / real-content world needs `stacksnap replay` first; without it the seeder degrades gracefully to a
structural-only world (free content refs, empty catalog). The structural params (size/role_mix/tier_mix/activity)
are unchanged. A 4th `full-fidelity` note is added to the README family flow.

## Recipe family refresh (section 2)
- Both org/onboarding + skill-progression recipes gain a **snapshot-replay step** before seeding, and call out the
  set-dressed result (real catalog in the skills view, real templates behind seeded sessions).
- The `recipe-skill-progression.md` "waived / future-v1.2" Notes paragraph (now FALSE) is rewritten to the shipped
  reality (taxonomy + content are `snapshot-seeded`, coverage 100%).
- A new **`recipe-snapshot-world.md`** documents the capture→replay→set-dressed-world flow end-to-end (the curator
  recipe for the snapshot layer), cross-linked from the README index.

## /demo-snapshot skill (section 3 — M11-Q2 RESOLVED: a NEW skill, not a /demo-seed extension)
A dedicated `/demo-snapshot` skill driving the `stacksnap` CLI (`capture` / `replay` / `status`). Rationale: capture
is a privileged prod READ with its own safety contract (the read-side firewall), distinct from `/demo-seed`'s
per-stack WRITES; the spec already treats them as sibling extensions. Frontmatter + body follow the `/demo-seed`
convention. Source-of-truth pointer: `corpus/ops/snapshot-spec.md`.

## Corpus updates (section 4)
- CLAUDE.md skills table gains the `/demo-snapshot` row (guide → `corpus/ops/snapshot-spec.md`).
- snapshot-spec / seeding-spec cross-links to the demo family + the new recipe; the demo README indexes the
  snapshot-spec + the new recipe + lists `/demo-snapshot`.

## Release-close hygiene carry (section 5)
- **`stacksnap` CLI help text stale** (`cmd/stacksnap/main.go`): says "framework (M9a/M9b)" + omits the `directus`
  surface from the printed surface list though it's registered. Fix → list `directus` + bump the M-tag. (Extensions
  change → `stack-snapshot-m11` tag at close.)
