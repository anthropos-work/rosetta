---
iter: 4
milestone: M221
iteration_type: tik
iter_shape: tik
status: closed-no-lift
opened: 2026-07-15
closed: 2026-07-15
strategy_ref: TOK-01 (../decisions.md)
---

# iter-04 (tik) — Phase C: the FIRST live battery cycle (BASELINE + diagnose)

**Type:** tik, under TOK-01 ("fix-then-prove, host-isolation FIRST"). Phase A (iter-02) landed the host lock;
Phase B (iter-03) landed the off-box demo-hygiene fixes. This tik is **Phase C, cycle 1** — the first live
cross-machine run on `billion`, following the **M215 direct-drive pattern**: bring up a **DEFAULT** demo
(**NO FLAGS**), drive it from a tailnet peer, and **capture EVERYTHING** into `findings.md`. Establishes the
baseline distance to the 8-condition gate. A clean pass on cycle 1 was **not** expected — the job is to surface
the live breakages.

## What this tik did
- Rolled the host rext pin to `cue-to-cue-m221-r2` (clone **+** `.agentspace/rext.tag` — F0).
- Cold reset-to-seed `up-injected.sh 1` **NO FLAGS** on `billion`, driven synchronously in bounded foreground
  polls (no detached/nohup scripts on the host; never graded a still-building stack).
- Measured all 8 gate conditions from this Mac (tailnet peer), plus the folded-in inherited carries.
- IRONCLAD-confirmed the **dominant** root cause (the snapshot store-root bug) via a controlled post-baseline
  diagnostic (`public.skills 0 -> 42,790`).
- Tore down, verified the host clean, released the lock, killed a survivor the teardown missed (F5).

## Outcome — `closed-no-lift` (baseline-with-full-findings)
No code fix landed this tik — **by design and honestly**: neither on-the-spot candidate was clean-to-land.
The academy carry's root cause was **disproven + re-characterized** (F4 — a client-side render defect, not the
empty local-content the carry hypothesized), so it is not a clean fix. The **dominant** breakage (the snapshot
store-root bug, F1) is IRONCLAD-diagnosed with an exact seam but needs a rext edit + RED-fence + re-tag +
re-cycle to ship+prove — **routed to iter-05 as the #1 fix, de-risked to certainty** rather than ballooned into
this baseline tik (the tik tripwire + anti-ballooning discipline). See `findings.md` for the full capture.

## Distance to gate (the headline)
**3 MET (1 latency, 3 orgs, 7 remote-by-default) / 4 NOT MET (2 catalog, 4 Dana, 5 Ben, 6 Aria) / 1 at-risk
(8 zero-edits).** The 4 NOT-MET conditions **collapse to ONE root cause** — the snapshot store-root resolution
bug (F1) skips the taxonomy replay, cascading through every taxonomy-dependent seeder. Fix F1 (+ the directus
re-capture F2) plausibly moves the baseline from 3-MET toward 6-7 MET in a single cycle.

## Out of scope (routed forward — Fate-3, iter-05+)
F1 store-root fix (#1) / F2 directus re-capture / F4 academy render-defect (re-characterized) / F5 native-academy
teardown reap + F5b gate-8 generated-file dirt / F7 c-3 router 403 re-check (with F2) / F10 freshness-abort +
assert_ports_free field-exercise. `BURNIN-M221-dev-public-host` + `F-M220-4` not reached this cycle.
