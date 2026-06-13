# M24 — Progress

**Status:** building. **Shape:** section.

## Section checklist
_One checkbox per concrete deliverable, ticked as it lands. Sections 1–3 land in `rosetta` (docs);
sections 4–7 land in the `rosetta-extensions` authoring copy (hygiene strand)._

### Rosetta docs (corpus-wide truth-up)
- [x] **§1 — Stale local-Directus corrections** (verified-against-compose): corrected the false local-Directus
  claims in `corpus/architecture/external_services.md` (image 10.10.1 + admin/password + the fictional compose
  snippet + local-container troubleshooting + local-uploads dir — all false; platform compose has no directus
  service, only `cms`'s `DIRECTUS_BASE_ADDR`→prod), `corpus/architecture/service_taxonomy.md` Directus table,
  `corpus/ops/quick_ops.md` ports table. Each now states the prod-read default + points at the v1.5 local tooling.
- [ ] **§2 — Known-state / safety / directus-local rewrites**: rewrite the `snapshot-spec.md` known-state block
  (per-stack Directus now real; exit-4 semantics redefined), the `safety.md` §2 deltas, finish
  `corpus/ops/directus-local.md`.
- [ ] **§3 — Corpus-wide language sweep** (via `/update-knowledge`): sweep the "print-only / exit-4 / reads-live-
  from-prod" language across the skills + `CLAUDE.md` so the whole corpus tells the new truth.

### Rosetta-extensions hygiene strand (each small + independently landable)
- [ ] **§4 — (a) Go toolchain pin bump** to go1.25.11+ (lazy rebuild — bump the pin only, no dedicated rebuild).
- [ ] **§5 — (b) README index-row guard**: a lint that fails when a new corpus doc lacks its directory-README
  index row.
- [ ] **§6 — (c) Zero-critical-genes guard**: `dna.Validate` / `compare.pct` must reject/flag a zero-critical DNA
  scoring 100% (verified still absent at `compare.go:247-252` / `dna.go:168-183`).
- [ ] **§7 — (d) `/project-stats` scope fix**: stop scanning the gitignored `stack-*/` platform clones that inflate
  the absolutes.

## Build log
_(append per build session)_
