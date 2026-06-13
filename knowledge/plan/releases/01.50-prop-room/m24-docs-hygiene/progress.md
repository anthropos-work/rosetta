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
- [x] **§4 — (a) Go toolchain pin bump** (ext `423bac7`): added an explicit `toolchain go1.25.11` directive to
  all 4 go.mod (alignment, clerkenstein, stack-snapshot, stack-seeding) + tightened the clerkenstein CI workflow
  `setup-go` from floating `"1.25"` to `"1.25.11"`. Lazy — pin only, no rebuild (parse-verified GOTOOLCHAIN=local).
- [x] **§5 — (b) README index-row guard** (ext `d6dd8fc`): `stack-core/corpus_index_guard.py` — for every
  README-bearing corpus dir, every other `*.md` must be referenced (by filename) in that README, else exit 1.
  8 unittest tests (full stack-core suite green, 77 tests). **Dog-food:** the guard surfaced 7 pre-existing gaps
  (5 architecture + 2 ops docs) — all backfilled rosetta-side so the guard passes clean on the live corpus
  (exit 0). Every M24-touched/created doc (directus-local.md etc.) is indexed.
- [x] **§6 — (c) Zero-critical-genes guard** (ext `04de89e`): `dna.Validate` now rejects a DNA with no critical
  gene (the vacuous-100% `pct(0,0)` hole); `compare.Report` gained a `CriticalGenes` count + `GateMet` refuses a
  non-zero critical gate when it's 0 (defence-in-depth). Tested; full alignment suite + vet + gofmt clean; the 5
  shipped DNAs all have a critical gene (live gate unaffected). Corpus: documented in `alignment_testing.md`.
- [x] **§7 — (d) `/project-stats` scope fix** (developer-kit `825cdce`): the `/project-stats` skill is the shared
  developer-kit `stats.sh` (no stats tooling exists in rosetta-extensions). Added `*/stack-*/*` to `PRUNE_PATHS`
  + a general `drop_gitignored` filter on the code-size scan, so the gitignored `stack-*/` platform clones (which
  inflated rosetta's count by ~2M lines / 9,235 files) are no longer scanned. Verified the collapse to the
  gitignore-respecting truth. Landed at the script's real source (the plugin — not a platform repo / not the
  corpus); rationale + a surfaced pre-existing doc-counter limitation in `decisions.md` M24-D3.

## Build log
- **2026-06-13** — full M24 build in one session. Phase 0b GREEN. §1–§3 (rosetta docs): corrected the false
  local-Directus claims (no directus service in the platform compose), converged the known-state + finished
  `directus-local.md` on the M23 cutover, swept the print-only/exit-4/live-from-prod framing across 5 skills +
  CLAUDE.md + 5 demo-corpus docs. §4–§7 (hygiene): Go toolchain pin → go1.25.11 (ext), the zero-critical-genes
  guard (ext), the README-index-row guard + 7-gap backfill (ext + corpus), the `/project-stats` gitignored-clone
  scope fix (developer-kit). Commits split: rosetta `m24/docs-hygiene`, ext `main` (4 commits), developer-kit
  `main` (1 commit). All 4 hygiene items Fate-1.
