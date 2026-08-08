---
iter: 149
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-08
---

# iter-149 — the emitter census, for every retired service rather than one

**Active strategy reference:** `TOK-08` — *census the mechanical classes; stop sampling them.* A reading
SAMPLES; a fence CENSUSES. Work the classes in descending measured size and report, per class, the
enumerated population and how many of it were defects.

## Cluster / target identified

`SURVEY-M257x-iter146-other-retired-services-unaudited`, opened by iter-146's own close and still the
largest named open route. iter-146 censused **one** retired platform fact end-to-end — the GraphQL router
platform `2adcf71` deleted — found 84 references, 82 correct, and **both misses were EMITTERS on
un-exercised paths**. It then wrote, in its own routes-carried-forward: *"the same emitter question is
unasked for every other one: `skiller`, `skillpath`, `chronos`, `intelligence`, the
`storage`/`messenger`/`customerio-sync` containers deleted at `838d907`."* It also wrote the constraint —
*"the emitter fence is one allowlist and one regex away from covering them; **do not widen it un-audited**
— grade a population first."*

**Step-0 re-survey (mandatory, and it changed nothing):** the target is untouched. `git log` since
iter-146 shows two commits (iter-147 profile-token census, iter-148 verify-scope census) and neither
touches the emitter fence or the retired-service token set. The route is still meaningful and still the
biggest.

## Hypothesis

The emitter class generalises, and so does its fence. The population is **not** one token — it is a
cross-product: **{retired service} × {arm}**, where the arms are the ways a dead service can be named in
executable content that an operator or a build consumes — its **base port**, its **container name**, its
**`host:port` address**, and its **RPC address variable**. iter-146 measured one cell of that grid
(router × port) and found 2 defects in 84 references. This iter enumerates the whole grid.

The honest prior, from `D-M257x-146-2` (*a repair's completeness tracks EXERCISE, not care*): the router
was the fact with the **most** exercise, so the residual for the less-exercised names could be higher —
or the class could be clean because the M257/iter-13 re-point swept them first. **Either answer is the
deliverable.** A census that returns zero is a result; a census that is never run is the defect.

## Expected lift

No `N` reading is planned, so **no `N` movement will be claimed** (`§9` guard-rail 1). The deliverable is
the class-level number `TOK-08` asks for: **population enumerated / defects found / fence that keeps it
at that number**, per retired service and per arm.

## Phase plan

1. **Enumerate the population.** Bind the retired set to the one that already exists and is already
   fenced — `claim_census_guard.ARCHIVED_SERVICE_NAMES` (12 names, each an archived row of
   `platform-migration-status.md`, itself machine-fenced against `repos.yml` in both directions by
   `platform_alignment_guard`). Do **not** author a thirteenth hand list; `-iter134` is open about the
   fence family having no shared predicate layer and this is one seam where it can have one.
2. **Census all four arms** across the tooling monorepo, classified per hit: emitter · fence-assertion ·
   comment/documentation · test fixture · derived-and-therefore-inert.
3. **Repair** whatever the census grades a live or latent emitter.
4. **Generalise the fence** — `test_deleted_router_endpoints.py` covers one hard-coded port; make it cover
   the enumerated set, keeping iter-146's comment carve-out and its both-directions pair control, and add
   an anti-vacuity control that binds the retired set to the shared source rather than re-declaring it.
5. **RED-proof it** on synthetic content in both directions before believing it.

## Escalation conditions

- If the census returns a defect whose repair needs a platform-repo edit → route forward; v2.8's zero-edit
  constraint holds.
- If binding to `ARCHIVED_SERVICE_NAMES` would make the fence RED on correct content (the rule-67 shape —
  a token carrying opposite obligations), do **not** widen: narrow the arm and say which arm was dropped.

## Acceptable close-no-lift outcomes

**A census returning zero emitter defects is a complete iter**, provided the population is enumerated,
the classification is stated, and the fence that holds it there is RED-proofed. That is the falsification
`D-M257x-146-2` predicts against, and recording it is worth more than another repair.
