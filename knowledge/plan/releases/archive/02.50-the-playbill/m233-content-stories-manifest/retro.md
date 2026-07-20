# M233 "content-stories manifest" — Retro

## Summary
Built the **manifest half** of Content stories: `BuildContentProducts` projects a `content_products[]` menu the
cockpit's 2nd tab reads — per content product, the played sessions each with player+manager seat keys + result paths +
`has_manager_view` + app_base + per-type icon — SINGLE-SOURCED from the SAME content-session fixture the M232 seeder
seeds from (the player seat OWNS the seeded session by construction). Honesty-gated (`CanonicalFileMatchesProjection` +
a teeth test) so `content-manifest.json` can't drift; fail-closed (drop-with-reason + `ValidateContentManifest` fails
loud — no fabricated links); emitted by `stackseed --content-export`. Open question resolved (`#D-M233-1`: a separate
JSON, the cockpit reads JSON not YAML). Deliverable: `content-stories-spec.md`. 0 platform edits.

## Incidents this cycle
- **None.** No P1/P2. No regressions, no flakes (5/5). The build's per-section review + the 2-pass harden (to 100%
  function coverage on the projector) meant the close review surfaced only one cosmetic doc gap.

## What went well
- The **single-source discipline** (D9) held cleanly: the projector re-uses the seeder's exact owner derivation
  (`eligiblePlayerOwnerSlots`) and session-id derivation (`contentStorySessionID`), so the honesty gate is a byte
  comparison against a checked-in canonical rather than a hand-maintained fixture. Cheap to keep honest.
- The **flat-index-survives-drops** invariant was recognized as load-bearing during harden and pinned by a dedicated
  test — the subtle failure mode (a per-product index reset silently re-owning every session after the first drop →
  byte-clean manifest, universally dead CTAs) is exactly the kind an adversarial review would hunt for, and it was
  already closed. The close's Phase 2c recorded the scenario as a landmine note for future editors.
- **Fail-closed by default**: a non-simulation session with no builder yet (M234's work) drops with a clear reason and
  `ValidateContentManifest` refuses the export — so a premature fixture addition fails loud, never ships a dead button.

## What didn't
- One decision (`D-M233-3`, slug-resolved-at-authoring) was blended into the spec at build time without its
  `(#D-M233-x)` back-ref tag, while D-M233-1/2/4 had theirs. Trivial, fixed at close — but a reminder that
  decision-triage tagging is easy to do inconsistently during a build.

## Carried forward (all Fate-2, confirmed in M234's `In:` list — not new deferrals)
- Bring-up export wiring (`up-injected.sh --content-export`) + cockpit tab render + `content-player-<idx>` seat
  registration → **M234**.
- Non-simulation product player-path builders (skill-path / academy) → **M234/M235** (need fixture route fields).
- Inherited standing carry (release-scoped, not M233): 14 pre-existing demo-stack test failures (REPEAT v2.4→v2.5) →
  v2.5 release-close re-anchor. See `audit-deferrals/deferral-audit-2026-07-19.md`.

## Metrics delta
- Whole-rext go test funcs 1902 → **1954** (M233 ~+23 in stack-seeding); `content_manifest.go` at **100% function
  coverage**; full stack-seeding suite GREEN (16 pkgs); flake **5/5**; **0 platform edits**. rext tags
  `playbill-m233-content-manifest` @ 9f0ab1c + `-hardened` @ c30fee3. See `metrics.json`.
