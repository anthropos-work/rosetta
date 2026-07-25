---
iter: 04
milestone: M254
iteration_type: tik
iter_shape: measurement-cluster
status: closed-fixed-partial
created: 2026-07-25
---

# M254 · iter-04 — read-only + latency measurement cluster (gates b/c/d/f/g/h)

**Type:** tik · **Active strategy:** TOK-01 (clusters 2+3 — read-only confirmation sweeps + latency).
Drove the read-only sweeps + the studio-FCP leg against the ONE live billion bring-up (from iter-03).

## Measured verdicts
- **(a) RE-CONFIRMED green LIVE** — regenerated a fresh autoverify (~9h post-bring-up) with the correct
  `--services`: `green:true, 0 warnings`, demo-patches ALL applied (the aireadiness fix holds), all probes pass.
- **(b) MET-with-disposition** — content-stories sweep LANDED 45/47; the 2 non-landing (voice MANAGER views
  `hire-voice-fail` + `asmt-voice-pass-en`, ~230 chars) are **coordinator-approved presence-only** (symmetric
  extension of `DEF-M240-01`: Bunny-keyless demo box → voice result can't render real media, either vantage).
  Gate (b) = **45/45 landable all landed + 4 voice cells presence-only**; surface renders, no fabricated CTA.
- **(d) MET (both vantages)** — Northwind AI-readiness coverage: manager `dana-manager` `gateMet:true` 70 pages
  0 escapes, `/ai-readiness` **8/8 sections incl all 3 drift-fixes** (by-tag ✓, interview-findings ✓ 3706 chars,
  handled-for-you ✓); employee `aria-completed` `gateMet:true` 63 pages 0 escapes 0 failingSections.
- **(c) prod-eject side PROVEN** — both coverage crawls show **0 escapes across 133 pages**. The explicit
  Back-to-Cockpit ITEM render/resolve check is the remaining piece (pending, routed forward).

## Residuals routed forward
- **(f) RESIDUAL — studio-desk session-carry on --public-host.** studio-desk `:19000` → 302 → `:13000/login`
  (its baked `VITE_CLERK_SIGN_IN_URL`); the cockpit session isn't carried to studio-desk, so the FCP flow lands
  on the stack web app (:13000 — NOT a prod-eject) and the studio `.page-skeleton` never paints. The
  roadmap-flagged env-sensitive FCP risk materialized. **Fix-iter needed** (studio-desk session-carry, or the
  FCP flow establishing studio's session). p95 not gradeable until studio actually loads.
- **(g) RESIDUAL — 9 host-sensitive test-health fails** (isolated-clean, listener killed): 2 pure nvm
  env-artifacts (`test_missing_node_documents` ×2 — billion's nvm node v22 defeats the node-absent stub) +
  7 needing disentangle/fix (launcher reap/stop intra-run listener leakage, 2 mutation meta-tests, a NEW
  `test_apply_revert_round_trip_on_the_real_next_config` = allowedDevOrigins patch vs the real next.config —
  possible consolidation drift). **Fix-iter needed** (rext, rung-zero).
- **(b) follow-up** (non-blocking) — reflect the presence-only disposition in tooling: `manager_presence_only`
  flag + `content-denominator.json` 47→45 + re-seed. Tracked, not gate-blocking.
- **Minor tooling drift** — `verify.sh` default service list still includes stale `skillpath`; standalone
  autoverify without `--services` probes a non-existent skillpath. Not gate-blocking (bring-up passes explicit
  `--services`). Follow-up: drop skillpath from the default.

## Still pending
- **(c)** explicit Back-to-Cockpit ×4 + studio logo/back/logout render check.
- **(e)** studio sim-builder-generate Playthrough (~10-min async).
- **(h)** p95 click→ACCESS < 5 s latency (SOLO) + Playthroughs green (fresh green autoverify in hand).
