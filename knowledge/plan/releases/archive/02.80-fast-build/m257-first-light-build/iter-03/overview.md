---
iter: 3
milestone: M257
iteration_type: tik
status: in-progress
created: 2026-07-31
---

# iter-03 — make the gate reachable

**Type:** tik · **Active strategy:** [`TOK-01`](../decisions.md) — still step (2)→(3); the campaign
cannot start while READY is unsatisfiable.

## Phase 1 Step 0 — re-survey

iter-02 closed hours ago and its routes are current. Confirmed: `odysseus.json` still absent (so
`buildbench --profile odysseus` refuses), B1 and B2 unfixed in rext, the odysseus stack still up on the
hand-applied `app/studio` hack.

## Active strategy reference

`TOK-01`: *"do not touch a lever until … odysseus can run a cycle"*. iter-02 proved it can run **a**
cycle — but only with a manual hack, and only to `green:false`. **This iter makes that sentence true
without an asterisk.** No lever here either.

## Cluster / target identified

iter-02's four blocking routes. **The gate is currently unreachable on any host**, so this is not
lever-adjacent work that could be deferred — it is the precondition for every number this milestone
will ever quote.

## Hypothesis

If B1 and B2 land in rext and the `load1` question is answered with the real instrument, then a cold
`--purge` + `demo-up` on odysseus can reach `green:true / 0 warnings` **reproducibly and without manual
intervention** — and iter-04 can run the baseline campaign against a definition that is actually
satisfiable.

## Expected lift

**Zero on the primary metric, again by design** — and this is the last iter for which that is true.
Grades on planned deliverables per Phase 4 Step 0.

## Phase plan

- **B1 `FIX-M257-seeders-local-mirror-drop`** — re-point the 6 failing `stack-seeding` seeders at
  canonical `public.job_simulation_sessions` / `public.skill_path_sessions`; correct the
  under-set-dress warning's misattributed cause; correct `content-stories-routes.md`'s MIRROR guidance.
- **B2 `FIX-M257-app-studio-acquisition`** — give `app/studio` a real acquisition path in
  `ensure-clones.sh` + `up-injected.sh`; correct `ensure-clones.sh:144-145`'s now-false premise. The
  hand-applied `cp` must stop being load-bearing.
- **`INVESTIGATE-M257-load1-48`** — is peak `load1` 48.7 real, reproducible and attributable? **If
  HEADROOM clause 1 cannot pass on this host, that is a re-scope signal and the milestone's shape
  changes.** Answer it with `buildbench`'s own sampler, not an ad-hoc probe.
- **Author `stack-core/hostprofiles/odysseus.json`** with the host facts already measured, deliberately
  **without** a `gated_baseline` (none is measured yet — the fence handles a baseline-less profile by
  design, and inventing one would be the exact defect `D120` exists to catch). This unblocks
  `buildbench --profile odysseus`, which currently refuses.

## Escalation conditions

- **`load1` confirms clause 1 is unsatisfiable on odysseus → surface immediately.** That is a gate-shape
  question, not an implementation one.
- A seeder re-point that needs a schema decision beyond re-pointing → route forward rather than guess.

## Acceptable close-no-lift outcomes

- B1/B2 land but a full green cycle is not achieved for an unrelated reason → `closed-fixed-partial`,
  with the green cycle routed to iter-04.
- The `load1` investigation returns "unreproducible / sampling artefact" → that is a **complete**
  answer and closes the route; a falsification is a first-class outcome.
