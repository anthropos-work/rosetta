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
- [x] **§2 — Known-state / safety / directus-local rewrites**: rewrote the `snapshot-spec.md` known-state block
  (the `--local-content` self-contained path now leads as the converged end-state; the prod-read/exit-4 path is
  the documented fallback; M23 retired from future-tense), finished `corpus/ops/directus-local.md` (status note
  M22→M23, the promised "data-plane cutover (M23)" + referential-closure sections added, M23 moved out of
  future-work, `cms`-only over-claims fixed). `safety.md` §2 verified **already M23-accurate** (landed in M23's
  own close) — investigated, nothing to change (Fate-1: work genuinely complete, not deferred).
- [x] **§3 — Corpus-wide language sweep** (via `/update-knowledge`): swept the "print-only / exit-4 / reads-live-
  from-prod / not-yet-automated" framing across the skills (dev-up SKILL+reference, demo-up, stack-snapshot,
  db-query) + `CLAUDE.md` AND the demo-facing corpus docs the sweep surfaced (demo/README.md, the 3 demo recipes,
  seeding-spec.md) — each now presents `--local-content`-executed (demo default-on / dev opt-in) as the converged
  self-contained state and prod-read/exit-4 as the documented fallback. Surfaced + threaded the real
  `--local-content` flag / `DEMO_NO_LOCAL_CONTENT` env into the skill descriptions (flag↔docs consistency).

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
