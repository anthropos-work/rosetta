---
iter: 04
milestone: M224
iteration_type: tik
status: archived
created: 2026-07-16
---

# iter-04 — THE FIRST GATE READING (baseline render + attribution)

**Type:** tik — under **TOK-01** (recruiter-render-first).

## Step 0 — Re-survey (mandatory)
TOK-01's `Next-tik direction` framed iter-02 as "enabling-scaffold + baseline-render," but the render half
was budget-split across sessions: iter-02 landed the Clerkenstein isHiring wiring, iter-03 landed the recruiter
seat, both metric-neutral by design. **The named target — the baseline render — is still untouched and is the
gate metric.** No substitution: this iter IS the deferred baseline-render half of TOK-01's opening move.
Current baseline = UNMEASURED (the surface has never rendered on a local stack).

## Active strategy reference
**TOK-01 (recruiter-render-first).** Every early tik targets rows-on-the-recruiter-scoreboard; the fix (if any)
is chosen by attribution against `hiring.md`'s traced read-path in the M219-trap order (Clerkenstein-identity →
render-gate-bypass → seed-gap). This iter is **reach + measure + attribute — NO fix before the attribution.**

## Cluster / target identified
The gate metric itself: `min(comparable candidate rows painted across the 5 seeded HIRING positions)` on
`rae-recruiter`'s comparison surface (`/enterprise/activity-dashboard → AI-Simulations → [simId]`), on a COLD
reset-to-seed LOCAL demo. Data is seeded (M223 funnel: ~40 of 45 candidates assessed on all 5 positions, score
band 30–95 ± jitter — non-degenerate). The render risk is concentrated in the render-GATE + the Clerkenstein
re-skin, exactly where the baseline reading is load-bearing.

## Hypothesis
The scaffold (isHiring wiring + recruiter seat + hiring funnel, all at `casting-call-m224-iter03`) is in place, so
SOME rows should paint. The open question the baseline answers: do ~40 rows paint per sim (near-gate), or does an
M219-class render-gate bypass the seed → 0 (or degenerate) rows? The number + its attribution is this iter's win.

## Expected lift
Qualitative — this is the **baseline**, so "lift" = producing the first measured `min(rows-per-sim)` + a
per-sim breakdown + score-distribution + closure/eject sub-readings + an attribution verdict. A gate-met reading
(≥40 × 5) is possible but not expected on the first look; an attributed gap is the acceptable-and-expected outcome.

## Phase plan (coverage-protocol A–E, render-gate variant)
- **Tooling (enabling):** author `tests/render-hiring-comparison.spec.ts` under `stack-verify/e2e/` + a
  `run-hiring-render.sh` driver — log in as `rae-recruiter` via the cockpit handshake, reach each of the 5
  `[simId]` comparison scoreboards (discover the sim links from the activity-dashboard DOM), count comparable
  candidate rows per sim, capture the score spread, and flag junk names / prod-eject. Harness-only (run from the
  authoring copy against the live demo — no rebuild needed to re-measure; the DATA is at the consumed iter03 tag).
- **Phase A — Sweep (baseline):** bring up LOCAL `demo-1` (`up-injected.sh 1 --no-public-host`, offset ports,
  auto-set-dress + stories seed), run the render-probe → the baseline `min(rows-per-sim)` + per-sim breakdown.
- **Phase B — Triage/attribute:** classify the gap per `hiring.md` read-path (Clerkenstein-identity vs
  render-gate-bypass vs seed-gap). NO fix before attribution (the M219 trap).
- **Phase C/D:** if a fix is clearly attributed AND lands within budget under the zero-platform-edit line
  (data-only seed change / Clerkenstein wiring / a sha-pinned demo-patch), land it + re-measure. Otherwise route
  the fix forward to iter-05 with the attribution as its evidence.
- **Phase E — Close:** grade on whether the baseline + attribution landed (the planned scope). Tag rext iter04
  (the probe tooling) + update `.agentspace/rext.tag`.

## DeepLinkCatalog (in-scope item 2)
Confirm the recruiter `jump_to: /enterprise/activity-dashboard` lands correctly; decide whether a per-`[simId]`
`NeedsID` DeepLinkCatalog entry is warranted once the seeded sim ids are known from the live surface.

## Escalation conditions
- **next-web OOM / UI-tier fails for memory** (Docker VM 9.7 GiB < 12 GB prereq) → STOP, `EXIT_REASON: user-blocker`,
  report the OOMed container + exit code so the user can raise Docker memory. Do NOT retry-thrash the bring-up.
- A render-gate that can ONLY be closed by a platform edit → re-scope-trigger (escalate; never edit the platform).

## Acceptable close-no-lift outcomes
Producing a measured baseline + a documented attribution IS the deliverable even if the gate is far off and no fix
lands this iter (the number + attribution is the win). A baseline that reveals a clean fix routes it to iter-05.
