---
iter: 266
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-10
active_strategy: TOK-08
route: FIX-M257x-263-dev-bringup-must-run-the-check
---

# iter-266 — a skill claims a pre-flight the dev path does not carry

**Type:** tik, under [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07).

## Step 0 — re-survey (mandatory, before targeting)

iter-263 relocated iter-262's defect: `INVITATION_HMAC_SECRET` was not undeclared, it was **a secret-source
gap plus an UNRUN CHECK** — `stacksecrets check` declares it critical and would have said so, and the dev
bring-up never ran it. Route `FIX-M257x-263-dev-bringup-must-run-the-check` was left open.

Re-surveyed at open, corpus `79d63d1`:

- `.claude/skills/stack-secrets/SKILL.md:22-24` asserts *"The pre-flight `check` **already rides inside**
  `/dev-up` + `/demo-up`"* — present tense, both stacks, no qualifier.
- The demo half is real: `demo-stack/up-injected.sh:1520-1544` builds and runs `stacksecrets`.
- `grep stacksecrets dev-stack/*.sh dev-stack/dev-stack` → **no output.**

So the target is live, and the *shape* has changed since iter-263 named it: this is not only a missing
step, it is a **claimed** step. The class is iter-265's, one layer up — a capability that exists on the
path that gets exercised and is *asserted* for the path that does not.

## Cluster / target identified

`FIX-M257x-263-dev-bringup-must-run-the-check`, re-framed by the re-survey as **a false capability claim**
rather than a missing feature. Under the milestone's own rule (§8, iter-186: *a correct exclusion is still
a defect while it is silent*), a claim that is false is worse than a gap that is honest.

## Hypothesis

`dev-stack/` has no secret handling at all; `/dev-up`'s own skill never claims one; and the assertion lives
in exactly one place — `stack-secrets`'s SKILL.md — where it reads as documentation of a shipped rider.
The correct repair is to make the claim true on the dev path **or** make it honest, and to fence which.

## Expected lift

Limb 3. A new engineer following `/dev-up` gets no coverage check, then hits a **silent `app Exited (0)`**
(the class `secretdna/demo.go:45-47` names verbatim) with nothing having warned them — while the skill that
owns the check says it already ran.

## Phase plan

1. Seal pre-registrations (first commit).
2. Measure: `dev-stack/` occurrences; `/dev-up` SKILL claims; the corpus-wide sweep of the assertion.
3. Re-measure the dev-vs-demo `stacksecrets check` asymmetry on today's tree.
4. Repair: honest claim + the documented dev step, per what the measurement supports.
5. Fence the claim so it cannot silently become false again.
6. Re-run the fences; close.

## Escalation conditions

- If the repair needs a **rext tag + pin bump** to land in `dev-stack/`, that bump would spend
  `D-M257x-258-1`'s frozen-pin control. Do not spend it inside this iter; route the tooling half and land
  the doc/skill half, exactly as iter-264 split `FIX-M257x-262`.
- If the check **does** ride inside `/dev-up` by a route the grep missed, the iter closes `closed-no-lift`
  with the falsification, and `FIX-M257x-263` closes with it.

## Acceptable close-no-lift outcomes

A documented falsification of PR-1 (the check does run on the dev path) closes the route with evidence.
