---
iter: 5
milestone: M221
iteration_type: tik
iter_shape: tik
status: open
opened: 2026-07-15
strategy_ref: TOK-01 (../decisions.md)
---

# iter-05 (tik) — Phase C, cycle 2: FIX the dominant blocker + re-prove

**Type:** tik, under TOK-01 ("fix-then-prove, host-isolation FIRST"). iter-04 captured the baseline (3/8 MET)
and ironclad-diagnosed the dominant root cause. This tik LANDS the fixes, re-tags rext, re-pins `billion`, and
**re-runs the live cold battery cycle** (a DEFAULT `up-injected.sh 1`, NO FLAGS, cold reset-to-seed, driven from
a tailnet peer) to measure the new gate distance.

## What this tik landed (all RED-fenced, rext `cue-to-cue-m221-r3`)

- **F1 — the store-root SHADOWING bug (dominant, de-risked to certainty in iter-04).** The snapshot replay
  resolved the WRONG cache root: a consumption clone's own empty `.agentspace` shadowed the real workspace
  cache (`panorama/.agentspace/snapshots`, 330k rows), so replay reported "cache miss", the taxonomy load was
  skipped, `public.skills=0`, and gates 2/4/5/6 cascaded to empty. Fixed in three places, each RED-fenced:
  - **(A)** `workspaceRootFrom` now PREFERS an ancestor whose `.agentspace` holds `snapshots/` over a nearer
    empty marker (generic — every caller benefits).
  - **(B)** `dev-setdress.sh` resolves the workspace cache root and passes an explicit `--store` on every replay
    (deterministic, cwd-independent).
  - **(C)** `replay` is now LOUD about an EMPTY store (the wrong-root smell / D17), not a benign "cache miss".
- **F5 — the native-academy teardown reap (field-confirmed still-broken in iter-04).** The academy survived
  `down --purge` on :13077. Root cause re-characterised: a supervised dev server RESPAWNS its socket-holding
  worker, so reaping only the listener let `next dev` re-bind. `reap_port` now kills identity-matching
  ANCESTORS too (the supervisor), identity-scoped + settle-before-reprobe. RED-fenced with a real
  supervisor→worker respawn harness.
- **F5b — the gate-8 generated-file dirt.** `ant-academy.sh --stop` now path-scoped `git checkout` restores
  the two predev-hook-generated files (`code/public/catalog.json` + `content/index.md`) so the clone is left
  git-clean. RED-fenced.
- **F2 — directus re-capture: NOT pre-emptively done.** iter-04's F2 divergence was measured against the
  BROKEN store; the directus cache HAS a `_structure.sql` artifact, so with F1 fixed the M21 auto-provision
  should bridge the bootstrapped-gap schema to the captured one. Measured live rather than speculatively
  re-captured (D-M221-05e).

## The re-prove — the 8-condition gate + M219's five readiness gates

_(gate table + measurement filled at close — see `progress.md` and the milestone `findings` capture)_

## Out of scope (routed forward if they balloon — Fate-3)
F4 (academy client-side render defect) and F10 (freshness-abort + `assert_ports_free` field-exercise) fold in
only if the cycle stays clean; either ballooning past ~1 h routes to iter-06.
