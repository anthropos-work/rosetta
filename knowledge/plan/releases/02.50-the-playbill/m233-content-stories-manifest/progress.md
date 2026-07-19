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
