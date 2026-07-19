# M233 — Decisions

## KB findings (from Phase 0b)

- **KB-1 (YELLOW) — ADDRESSED (Phase 5).** `seed-manifest-spec.md`'s block inventory omitted the M232
  `content_sessions` block. Backfilled: the §1 block table now lists `content_sessions`, and a new §8 documents
  both the M232 `content_sessions` block + the M233 `content-manifest.json` peer.

## Design decisions

- **D-M233-1 — the open question is resolved as "separate `content-manifest.json`" (the preferred option).**
  The presenter cockpit is **stdlib-only Python — it reads JSON, never YAML (no PyYAML, the supply-chain-GREEN
  posture)**. So the RENDER surface the content tab reads MUST be JSON — a `content-manifest.json` projected by
  `stackseed --content-export`, exactly analogous to the existing `cockpit-manifest.json` (the stories menu, also a
  separate JSON, also NOT the seed-generation-manifest.yaml). The source-PINS stay folded in
  `seed-generation-manifest.yaml`'s `content_sessions` block (M232). So `content_products[]` is a **peer manifest
  artifact** (the content analog of the population/stories manifest), honesty-gated by its own checked-in canonical
  `presets/content-manifest.json` + a `CanonicalFileMatchesProjection`-style test. This preserves the non-fatal
  bring-up (a broken content projection drops the tab, never blocks the cockpit) and keeps `manifest`/`seeders`
  decoupled from each other.

- **D-M233-2 — content-story sessions render in `apps/web` (app_base = web), not apps/hiring.** The seeder
  re-tenants every clone into a WORKFORCE org (`firstNonHiringStory`). M231's "HIRING → apps/hiring" is the
  org-ejection rule for genuinely-hiring ORGS (M224) — a different feature. app_base keys on the HOST-ORG type, so a
  HIRING-*sim_type* clone in a Workforce org renders in apps/web. sim_type still drives the icon + the interview
  flag-gated render. (The per-product app_base map still carries `academy` = the offset academy origin for the
  future academy product.)

- **D-M233-3 — slug resolved at authoring time into the fixture (`sim_slug`).** The player result path
  `/sim/<slug>/result/<sessionId>` resolves by `jobSimulationBySlug(slug)` (a TEXT slug, not the sim uuid). The 8
  pinned public sim slugs were resolved read-only (public-published, tenant-less — the db-access boundary) and baked
  into the fixture, so `BuildContentProducts` is fully OFFLINE + the projection is concrete + honesty-gated. The
  fail-closed resolver is then a pure function.

- **D-M233-4 — fail-closed = drop-with-reason (never fabricate) + a fail-LOUD guard on curated exhibits.**
  `BuildContentProducts` DROPS (never fabricates) any session that can't form a real link (missing slug/seat/route),
  returning the drops with reasons. `ValidateContentManifest` (the `--content-export` guard, the analog of
  `ValidateCockpitManifest`) FAILS LOUD if a curated SIMULATION exhibit was dropped — "a refusal nobody sees never
  happened" (the D17 / cockpit-guard philosophy). AI-labs presence-only (no player link) is a legitimate disposition,
  not a drop.

## Cross-milestone handoffs (three-fate rule)

- **M234 owns the bring-up export wiring + the cockpit read (Fate-2, already planned).** M233 delivers the
  `stackseed --content-export` verb + the honesty-gated `content-manifest.json`. Wiring `up-injected.sh` to export it
  at bring-up (as it already does for `cockpit-manifest.json`) + the cockpit tab reading it is squarely in M234's
  scope ("per-product sections rendering the M233 manifest" + "mint/resolve per-session player seats via roster.go +
  Clerkenstein" — the player-seat REGISTRATION that makes `content-player-<idx>` a live login). Confirmed covered by
  M234's `In:` list — no new deferral.

- **M234 owns the non-simulation product player-path builders (Fate-2, already planned).** The M233 product
  registry is SCHEMA-COMPLETE for all four products (simulation / skill-path-legacy / skill-path-new / ai-labs:
  app-base + icon + disposition). The fixture today carries only `simulation` sessions, so only the simulation
  player-path builder (`/sim/<slug>/result/<sessionId>`) is exercised. The skill-path / academy player-path builders
  (which need route fields — skillPathId, chapter slug — not yet in the fixture) land with M234/M235's fixture
  additions. `playerResultPath` returns a clear fail-closed reason for a non-simulation link-bearing session until
  then (no fabrication, no stub). Covered by M234's `In:` list.

## Adversarial review (close Phase 2c)

- **Scenario: the projector's flat session index silently de-syncs from the seeder's, re-owning every session
  after a drop.** The projected `player_seat` must OWN the seeded session, which holds only if the projector's
  flat index (`buildContentProductsFromSet`) advances in lock-step with the seeder's `for idx, cs := range
  Embedded().Sessions()`. The subtle failure: if the projector reset its index per-product, or skipped the
  increment on a dropped/unknown-product session, then the first drop would shift every subsequent owner by one
  — the manifest would compile, the honesty gate (which re-projects the SAME code) would still pass, and yet
  every CTA after the first drop would log in as the wrong member and land on a session that member never took.
  A byte-clean manifest with universally dead CTAs — invisible to any self-consistent test.
  - **Verification:** the projector increments `flat` BEFORE the known/drop checks
    (`content_manifest.go:222-231`), so drops consume their slot; `Set.Sessions()` flattens Products→Sessions in
    the exact declaration order the projector's nested loop walks; both `owners` (seeder) and `slots` (projector)
    derive from `eligiblePlayerOwnerSlots`. The invariant is pinned by `TestContentProducts_FlatIndexSurvivesDrops`
    (M233 harden), which feeds a drop-bearing set and asserts a later product's owner is unchanged. Handled; no
    new fix. (Recorded so a future editor who "tidies" the loop into a per-product index reset sees the landmine.)
