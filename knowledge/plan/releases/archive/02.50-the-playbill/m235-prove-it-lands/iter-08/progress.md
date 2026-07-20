# iter-08 — progress

**Type:** tik (cleanup-shape — plan routing + doc sync, 0 production seeder code)

## What landed
1. **M236 Fate-3 handoff (user-authorized).** EDITED `m236-prove-on-billion/overview.md` `In:` list to OWN:
   (a) AUTHOR the new content-stories seat-login coverage/Playthrough plumbing (Playthroughs-style seat-login +
   exact dynamic `/sim/<slug>/result/<sessionId>` URL resolution + the shared result page-object) + CALIBRATE
   against a live render; (b) prove EVERY (session×action) lands — simulation AND the 3 M235 non-sim sections
   — with the per-section live-calibration checklists; (c) the M230 carry-forward live items (ANT_ACADEMY
   coverage descriptor + next-web clone re-anchor + getPublicCatalogView 2nd manifest). Updated the M236
   open-questions (resolved) + `last_updated`.
2. **RECORDED USER-BLOCKER-M235-02's RESOLUTION** + the M236 Fate-3 handoff + the milestone pragmatic-close
   disposition in the M235 milestone `decisions.md`.
3. **Corpus doc sync:** `content-stories-spec.md` §2 (the non-sim registry) + §6 (the 3 sections BUILT, live
   proof M236); `content-stories-routes.md` §7 (build-status note — non-sim sections built + unit-proven, live
   proof M236).

## Close — 2026-07-20

**Outcome:** the live-coupled work is routed to M236 (user-authorized Fate-3), USER-BLOCKER-M235-02's
resolution is recorded, and the corpus reflects "non-sim sections BUILT (offline); live proof M236." The
offline-buildable scope is EXHAUSTED (3/3 non-sim sections built + unit-proven across iters 05–07).
**Type:** tik (cleanup)
**Status:** closed-fixed
**Gate:** NOT MET (the milestone's live (session×action)-lands gate is M236; M235 delivered the built +
unit-proven substrate, the user's "build non-sim seeders, then close" mandate).
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n (4 tiks) — (6) protocol-stop: **y** — Outcome: exit-6 (protocol-stop: offline clusters
exhausted; live proof routed to M236).
**Decisions:** D1 (M236 Fate-3 handoff), D2 (cleanup-iter, rosetta-only), D3 (pragmatic-close) — see
iter-08/decisions.md.
**Side-deliverables:** none.
**Routes carried forward:** ALL remaining M235 gate items are M236 (the live proof + the seat-login coverage
plumbing + the per-section calibration + the M230 carry-forward live items) — recorded in M236's `In:` list.
**Lessons:** when a milestone's remaining scope is exclusively live-render-coupled and a peer milestone (M236)
already runs the live cold bring-up, the honest move is a clean Fate-3 handoff (edit the target's plan +
record it) + a pragmatic-close — NOT faking a live gate or building blind against an uncalibrated harness.
