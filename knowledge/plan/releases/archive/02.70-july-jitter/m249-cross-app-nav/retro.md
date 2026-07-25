# M249 — Retro (cross-app navigation: Back-to-Cockpit + studio prod-eject fix)

## Summary
Closed the demo's navigation loop: a fail-closed "← Back to Cockpit" item in all **4** sub-app menus
(one shared `packages/ui` patch covers web+hiring; the **first-ever studio-desk SOURCE patch trio** image-baked
via a net-new `build_frontend_studio_desk` patch ladder + patch-set fingerprint; a native-run ant-academy
helper) reading a per-stack `COCKPIT_URL` @ 7700+OFFSET, **plus** the studio logo/back/logout **prod-eject fix**
(read this stack's app; original prod host kept as the `|| …` fallback → behaviour-identical off-demo). Patch
inventory 16→**21**. Introduced the **additive-UI injection** pattern (a NEW menu element vs a URL rewrite) as
`demopatch-spec.md` §8. Shipped: rext code-of-record `july-jitter-m249-harden @ 8ab5192` (138 M249-touched tests
GREEN, flake 5/5, +7 harden unit tests); 4 corpus docs; **LIVE-verified GREEN on demo-2** (4/4 menus @ `:27700`;
studio `:23000` baked, 0 effective ejects). Deferral audit GREEN (0 own deferrals). 0 platform-repo edits.

## Incidents This Cycle
- **P2 — a stranded `.env.production.local` overlay silently defeated the studio bake (D5, found at the live
  render-confirm).** The first demo-2 studio rebuild shipped a bundle with the "Back to Cockpit" item fail-closed
  to nothing: a Jul-17 stranded `.env.production.local` (a crashed prior build's leftover) tripped the M214
  `if [ -e "$desk_env" ]` "skip if it exists (repo-shipped?)" branch, so NEITHER `VITE_CLERK_SIGN_IN_URL` NOR
  `VITE_COCKPIT_URL` was baked. But studio-desk's real repo **gitignores `.env.*.local`** — it can never ship
  one, so a pre-existing overlay is ALWAYS a leftover. Fix: write the overlay unconditionally (`>` truncate, same
  contract as next-web's `apps/web/.env.local`); the RETURN trap removes it. Regression-fenced
  (`test_studio_desk_overwrites_a_stranded_overlay` + a failed-build trap test); the obsolete non-clobber test
  removed. rext `bcbb779`. Caught by the milestone's own live render-confirm — the "silently-unbaked" trap the
  patch-set fingerprint kills one level up.

## What Went Well
- The demopatch **anchor→replace ladder already supported insertion** (the R4 mitigation), so the additive-UI
  patches were a known-shape extension, not new machinery — exactly the `section` bet the shape rationale made.
- The **live render-confirm before finalizing** surfaced D5 (the stranded-overlay bake defeat) that a bundle-only
  or source-only check would have missed — the item looked present in source but was fail-closed to nothing in
  the shipped bundle.
- Harden was **pure test-depth on already-correct, LIVE-verified code** — 0 production bugs; the +7 unit tests
  (chain sha-linkage, a deep fail-closed guard fence, the native-helper REFUSE ladder) closed 3 genuine gaps
  without touching behaviour.
- Zero platform edits: every fix rides the demo's ephemeral gitignored clone; the canonical repos are untouched.

## What Didn't
- The first studio bundle shipped an inert item because the bake was silently skipped — the source patch was
  correct but the ENV WIRING (the overlay write) had a pre-existing stranding hazard. A reminder that for a
  baked-at-build feature, "the patch applied" is necessary but not sufficient — the bake must be verified in the
  shipped artifact (the fingerprint + live render-confirm now enforce this).
- A full native-browser render of the ant-academy item was out of reach on demo-2 (the pinned consumption
  clone's launcher + a test orphan squatting `:23077`); verified at source+env level instead. The live-browser
  academy proof rides M254.

## Carried Forward
- **The 2 pre-existing `test_ant_academy*` failures → M254 (Fate-2, confirmed-covered).** NOT M249 regressions
  (verified identical on committed rext HEAD) and NOT M249's patch domain — they are a subset of the **8**
  live/env/docker-gated demo-stack failures M251's close already fated Fate-2 → M254 (gate parts (g) live-box
  test-health + (h) live-browser re-prove). No sibling `overview.md` edit; recorded in `decisions.md` + the
  deferral-audit report.
- **M253 (studio-first-paint) is now unblocked** — it extends the `build_frontend_studio_desk` studio patch
  ladder this milestone created.

## Metrics Delta
- rext demo-stack tests: `test_back_to_cockpit_m249.py` 25 → **32** (harden +7); full M249-touched suite 131 →
  **138** GREEN at code-of-record. Flake gate 5/5 clean (close) + 3/3 (harden).
- Patch inventory: **16 → 21** (`11 next-web-app · 2 app · 5 ant-academy · 3 studio-desk`) — ground-truthed
  against the real rext `patches/` dir at close, fenced by `test_patch_inventory.py`.
- Live (demo-2): 4/4 app menus carry Back-to-Cockpit @ `:27700`; studio prod-eject fixed (`:23000` baked, 0
  effective ejects).
- Flake: 0. Platform-repo edits: 0. Code-of-record: `july-jitter-m249-harden @ 8ab5192` (on origin, rung-zero verified).
- Full metrics: `metrics.json`.
