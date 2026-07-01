---
iteration_type: tik
iter_shape: tooling
status: planned
---

# iter-08 — deep-link the demo entry (the user-chosen zero-edit strategy) + verify the frozen fast path

**Type:** tik — under TOK-01 (coverage-drive strand), applying the user-chosen **DEEP-LINK THE DEMO ENTRY**
strategy (run-5 brief): point Dana's cockpit `jump_to` + the coverage manifest at the AI-readiness dashboard
**with `?cycle=<latest-closed-cycle-id>`**, so the FE fires the cycle-scoped GET that the platform serves from
the FAST frozen path (`app buildResponseFromSnapshots`), reaching `(0,0)` without a platform edit.

## Active strategy reference
TOK-01 (active-cycle→closed-cycle showcase, coverage-drive strand). This iter is the **deep-link demo-entry**
refinement chosen by the user after iter-07 falsified the *default FE route* to the frozen path.

## Cluster / target identified
The residual `failingSections=5` (2 AI-readiness sections + 3 workforce-aggregate sections) — all skeleton
false-fails whose data is proven-correct in the DB (199 frozen snapshots, 78.4% stage-3). iter-07 root-caused
that the platform FE's DEFAULT dashboard GET omits `?cycle=` → the live-recompute wall. The user's chosen fix:
make the cockpit-sanctioned demo entry (Dana's `jump_to`) carry `?cycle=<closed-id>` so the FE fires the
cycle-scoped GET → the frozen `buildResponseFromSnapshots` branch → fast → the sections clear.

## Hypothesis
The deep-linked `?cycle=<closed-id>` entry hits `buildResponseFromSnapshots` (frozen read) which is FAST →
the 2 AI-readiness sections clear; the 3 workforce-aggregate sections either share the frozen deep-link or are
disclosed-allowed (per the run-5 brief).

## Expected lift
`failingSections 5 → 0` (or ≥ the 2 AI-readiness sections cleared via deep-link, the 3 aggregates disclosed).

## Phase plan (tooling-iter shape)
1. **Verify the deep-link hits the frozen fast path FIRST (cheap probe, mandatory)** — an authenticated
   dual-endpoint direct probe (`probe-aireadiness-deeplink.spec.ts`): lift Dana's token, hit `/cycles`
   (the FE gate) + the frozen data GET `?cycle=<closed>` DIRECTLY, measure both. If fast+correct → proceed to
   the cockpit + manifest deep-link edits + a GATED sweep. If FALSIFIED → the run-5 fallback
   (disclosed-presenter-note, needs sign-off).
2. (contingent on 1) cockpit `jump_to` + manifest `?cycle=` deep-link + GATED sweep, OR the falsification close.

## Escalation conditions
- Deep-link falsified (frozen path not fast) AND the only remaining zero-edit path (disclosed-allow) needs the
  user's explicit sign-off → **user-blocker** (surface, don't self-grant the disclosure).
- Any path requiring a platform edit → re-scope-trigger (never edit the platform).

## Acceptable close-no-lift outcomes
The deep-link hypothesis is falsified with full evidence (the frozen read is itself slow, so the deep-link
cannot clear the wall even in principle) — a complete investigation ending in characterization is a valid
closed-no-lift under the protocol.
