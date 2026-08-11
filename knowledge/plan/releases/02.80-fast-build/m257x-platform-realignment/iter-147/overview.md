---
iter: 147
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-08
---

# iter-147 — the emitter census, inverted: every path that CHOOSES a compose profile

**Active strategy:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07)
— census a mechanical class exhaustively, state the denominator, fence it to zero. `TOK-08`'s sealed
refutation branch bars an agent-authored successor; this is a tik under it, not a strategy revision.

## Cluster / target identified — and the substitution, stated

`TOK-08`'s standing direction was carried by iter-146's route
**`SURVEY-M257x-iter146-other-retired-services-unaudited`**: iter-146 censused **one** retired platform
fact end-to-end (the deleted GraphQL router's `:5050`) and asked the same emitter question for
`skiller`, `skillpath`, `chronos`, `intelligence`, and the `storage`/`messenger`/`customerio-sync`
containers deleted at `838d907`.

**Re-survey (Phase 1 Step 0) says that population is inert, and says so with numbers.** The four retired
services that still carry a port token — `cms` 8090, `jobsimulation` 8400, `storage` 8300, `roadrunner`
10400 — appear on **74 lines** in `rosetta-extensions` outside `.git`/`node_modules`/`tests/fixtures`,
and every one is a probe-registry row, a test assertion about that registry, port-offset arithmetic in
`GUIDE.md`/`rosetta-demo`, or guard prose quoting the corpus. **Zero emitters.** A registry row is a
*consumer*: it is scoped out by `STACK_SERVICES` and, unscoped, it fails loud as a `down` warning.

**Substitution under the same strategy: invert the search.** iter-146 hunted a **known-dead token across
all files**; iter-147 enumerates **every endpoint the known emitters announce** and grades each for
liveness. The inversion is not cosmetic — it is the only form that can reach this iter's finding, because
the defect here is an **absent** value, and an absent value has no token to grep for.

## Hypothesis

The residual sits where iter-146's `D-M257x-146-2` predicts: on paths nothing exercises. Widening from
"what does the tooling *say*" to "what does the tooling *choose*" should surface a defect the token
census structurally cannot see.

## Expected lift

No `N` reading is planned, so **no `N` movement will be claimed** (`§9` guard-rail 1). The deliverable is
a censused class with its denominator, any live defect it contains repaired, and a fence holding it at
zero with a mutation control and an anti-vacuity control that can fire.

## Phase plan

- **A** — census: enumerate every compose invocation in `rosetta-extensions` that selects a profile, and
  grade how each obtains it. State the denominator.
- **B** — repair anything the census grades as a live defect, with the consequence named in the refusal.
- **C** — fence: behavioural where possible (assert on the argv the shipped code produces), RED-proved
  against the real pre-fix text recovered from `HEAD`, never a reconstruction.
- **D** — gates: the new fence, the affected section suite against iter-145's baseline, the guard family.

## Escalation conditions

A finding that requires a **platform-repo edit** escalates rather than lands (v2.8 holds 0 platform
edits). A finding whose repair would need `stack-demo/**` touched escalates — `demo-1` is up and clauses
1–2 are proven on it.

## Acceptable close-no-lift outcomes

The census returning **zero live defects** is a complete iter, provided the denominator is stated and the
population is enumerated rather than sampled. `TOK-08`'s whole premise is that a census that returns zero
is a *measurement*, where a reading that returns zero is only a sample.
