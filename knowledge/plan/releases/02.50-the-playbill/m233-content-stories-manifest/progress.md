# M233 — Progress

## Sections

### §1 — content_products projection model + fail-closed resolver ✅
- [x] Resolve the 8 pinned public sim uuids → their public slugs; add `sim_slug` to the content-session fixture + struct + validation (public, non-PII, needed for the player result path `/sim/<slug>/result/<sessionId>`)
- [x] New content-products projection types (product: id/name/app_base/icon + sessions) + `BuildContentProducts(bp)` resolver, single-sourced from `contentsession.Embedded()` + the blueprint
- [x] Resolve per session: player_seat (owner member, single-sourced with the seeder's owner derivation) + manager_seat (host org manager hero) + player/manager result paths + has_manager_view + app_base + per-type icon
- [x] Fail-closed / no-fabrication: drop (with a recorded reason) or downgrade a session that can't form a real link (missing slug/seat/route); never invent a link
- [x] Complete per-product app_base + icon maps (simulation / skill-path-legacy / skill-path-new-academy / ai-labs); per-sim_type icon map
- [x] Unit tests: projection shape, seat single-source, fail-closed drop, no-fabrication teeth, icon/app_base coverage (9 tests)

### §2 — honesty-gated content-manifest.json + --content-export (D-M233-1: separate JSON, cockpit reads JSON not YAML) ✅
- [x] `stackseed --content-export` — the standalone `content-manifest.json` render surface (mirrors `--cockpit-export`, non-fatal) M234 reads
- [x] Checked-in canonical `presets/content-manifest.json` (the honesty-gated artifact, 9 sessions)
- [x] `CanonicalFileMatchesProjection`-style honesty gate on `content-manifest.json` + a teeth test on the projection
- [x] CLI-flag ↔ docs consistency (both directions); the M232 `content_sessions` source-pins stay folded in `seed-generation-manifest.yaml` (no regen — buildContentSessions doesn't carry sim_slug)

### §3 — corpus deliverable content-stories-spec.md (manifest-schema half) ✅
- [x] Author `corpus/ops/demo/content-stories-spec.md` — the manifest schema, the projection, the honesty gate, the fail-closed resolver, the seat model, the app_base + icon maps, the open-question resolution
- [x] Update cross-refs (seed-manifest-spec.md KB-1 backfill + §8, session-clone-spec.md already refs it, demo/README.md, CLAUDE.md index)
- [x] Tag the rext authoring clone (`playbill-m233-content-manifest` @ 9f0ab1c)

## Status: all sections complete. rext tag `playbill-m233-content-manifest` (9f0ab1c). Ready for `/developer-kit:close-milestone`.

## M233: Hardening

### Pass 1 — 2026-07-19
**Scope manifest (milestone-touched Go, rext `stack-seeding`):** `seeders/content_manifest.go` (primary,
355 L — projector `BuildContentProducts`/`buildContentProductsFromSet`/`resolveSession`/`playerResultPath`/
`ValidateContentManifest`/`WriteContentManifest` + `simTypeIcon`/`managerHeroKey`/registry) · its
`content_manifest_test.go` · `contentsession/contentsession.go` (M233 `sim_slug` + `slugRE` + `Validate`) ·
`cmd/stackseed/main.go` `doContentExport` + `main_test.go` honesty gate · `seeders/content_stories{,_write}.go`
M233 single-source additions (`ownerSlot`/`eligiblePlayerOwnerSlots`/`contentStorySessionID` — already 100%).
The M232 validators in `content_stories_write.go` (`validTerminalEval` etc.) are adjacent (M232 code) — out of
M233 harden scope.

**Coverage delta (milestone-touched packages, statements):**
- seeders: 96.0% → 96.2% (+0.2) — but the signal is per-function: `content_manifest.go` went from six
  sub-100% functions to **100% across every function** (simTypeIcon 83.3→100, managerHeroKey 83.3→100,
  playerResultPath 80→100, isSimulationType 66.7→100, buildContentProductsFromSet 94.6→100, WriteContentManifest 83.3→100).
- contentsession: 91.6% → 93.7% (+2.1) — `Validate` 96.6→100, `ProductFor` 80→100 (`Embedded` stays 75% = the
  defensive panic on an invalid go:embed, a compile-time invariant with no runtime input path — not chased).
- cmd/stackseed: 64.3% → 64.5% — `doContentExport` 70.8→75 (the `--seed` guard).

**Tests added (Pass 1, 10):**
- `content_manifest_test.go`: +7 — presence-only AI-labs disposition (row without player path, NOT a drop);
  non-simulation link-bearing drop (the M234-deferred `playerResultPath` reason path); **flat-index-survives-drops**
  (the load-bearing seat single-source: the projection's flat session index advances through dropped/unknown-product
  sessions, staying aligned with the seeder's `Sessions()` index — a regression = dead CTAs); simTypeIcon fallback;
  managerHeroKey blank-hero skip; WriteContentManifest fail-closed on no-host-org (writes nothing); the M217
  wire-format lesson applied — empty projection marshals `"products":[]` never `null`.
- `contentsession_test.go`: +1 — `sim_slug` validation (malformed uppercase/underscore/edge-hyphen rejected,
  omitted allowed, valid round-trips) + `ProductFor` not-found branch.
- `cmd/stackseed/main_test.go`: +1 — `--content-export` requires `--seed` guard.

### Pass 2 — 2026-07-19
Closed the two remaining branches on the projector's export entry.
**Coverage delta:** seeders 96.1% → 96.2%; `content_manifest.go` now **100% across all functions**.
**Tests added (Pass 2, 2):**
- `TestContentProducts_NoEligiblePlayerOwner` — the third fail-closed drop-path: a non-hiring host org exists but
  is degenerate (Size==heroCount → no member slot), every session drops with "no eligible player owner", guard
  fails loud. (Distinct from no-host-org.)
- `TestWriteContentManifest_EncodeError` — a failing writer surfaces the wrapped encode error with 0 sessions
  reported (the I/O error path, previously unexercised).

**Bugs fixed inline:** none — the projector held under every drop-path, disposition, and wire-format probe.

**Flakes stabilized:** none observed (all new tests deterministic; 3 consecutive clean sequential runs).

**Knowledge backfill:** see the `**Knowledge backfill:**` note below.

### Stop condition
Stopped after Pass 2 (2 passes): the primary file `content_manifest.go` reached 100% function coverage, the
Pass-2 statement delta was < 2%, and the Step-2b scan found nothing new worth adding (the sole remaining gap is
the defensive `Embedded()` panic on a compile-time invariant — a contrived test there is a disguised deferral).
12 tests added, 0 bugs, 0 flakes. rext test commit + re-tag; no source change.

## M233: Final Review

Cross-cutting close review (Phases 1–5). The build+harden already reviewed per-section and drove
`content_manifest.go` to 100% function coverage, so the close review surfaced almost nothing new. Deferral
re-audit (Phase 1b) YELLOW, 0 blockers — see `audit-deferrals/deferral-audit-2026-07-19.md`.

### Scope
- [x] All 3 sections checked; no TODO/FIXME/HACK in the M233 code. The 2 Fate-2 handoffs (bring-up export +
      cockpit render + `content-player-<idx>` seat registration; non-simulation player-path builders) RE-VERIFIED
      present in M234's `overview.md` `In:` list — no Fate-3 roadmap edit needed.

### Code Quality
- [x] No findings. `content_manifest.go` is consistent with `cockpit.go` (dual json+yaml tags, fail-closed
      philosophy, `Validate*Manifest` guard). `go vet ./seeders/... ./contentsession/... ./cmd/stackseed/...` exit 0.

### Adversarial (Phase 2c)
- [x] Failure mode considered — the flat session index must stay aligned with the seeder's `Sessions()` index
      *through drops*, or every CTA after the first drop re-owns the wrong seat (dead links). VERIFIED single-sourced
      at both ends (seeder `content_stories.go:80` `owners[idx%len]` over `Embedded().Sessions()`; projector flat
      `idx++` before the drop check; `Set.Sessions()` flattens Products→Sessions in the same declaration order).
      Already pinned by `TestContentProducts_FlatIndexSurvivesDrops`. No new test needed.

### Documentation
- [x] `content-stories-spec.md` accurate; all 5 cross-refs resolve; indexed in `demo/README.md` + `CLAUDE.md`.
      No new top-level unit (a new file in the existing `seeders` package) → no per-unit handbook owed.

### Decision Triage
- [x] D-M233-3 (slug resolved at authoring time) → blended in §3 but was MISSING its `(#D-M233-3)` back-ref tag
      (D-M233-1/2/4 all had theirs). Added the tag. FIXED.
- [x] D-M233-1 / D-M233-2 / D-M233-4 → already blended + tagged in `content-stories-spec.md` (§1/§3/§4). Verified.
- [x] KB-1 → already addressed at build (seed-manifest-spec.md §8 backfill). Archive.
- [x] Cross-milestone handoff decisions (M234-owns-*) → archive (maintainer-only; the scope boundary is already
      documented in spec §6).

### Tests & Benchmarks
- [x] No gaps. Full rext `stack-seeding go test ./...` GREEN (16 pkgs); harden already reached 100% function
      coverage on the primary file. Flake gate run at Phase 8.
