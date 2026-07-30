---
iter: 18
milestone: M256
iteration_type: tik
status: archived
opened: 2026-07-29
---

# iter-18 — price the onboarding cluster by driving it, not by reading its blockers

**Active strategy:** `TOK-01` move 3/4 — onboarding was ordered *after* org-admin because it was recorded
**seed-blocked**, and its cost was priced as *"a seeder + capability + roster seat, not just specs."* This
iter tests that price against the live surface.

## Cluster / target identified

Onboarding is **the long pole**: 1 landed (`onboarding.completion.UC1`, net-new at iter-08) and **4 curated
UCs TODO** — the largest remaining block of exit-gate clause 3. Each carries a written blocker in
`manifest/onboarding.yaml`, and **none of those blockers has ever been driven.** iter-08 wrote them from a
single Phase-A pass that only ever clicked **Skip**.

The milestone's own record says what to do with an unmeasured blocker: iter-07 refuted the pre-flight
audit's F5 ("no pre-onboarding state can exist") by *reading the schema*, and iter-17 refuted a
four-times-measured wall by *trying a different input modality*. **A blocker that has not been driven is a
hypothesis.** So this iter drives all four before writing a line of harness code.

## Hypothesis

At least one of the four recorded blockers is wrong, and the cheapest UC is cheaper than its note claims.
Specifically: `enterprise-workforce-standard.UC1`'s note pins it on *"a résumé fixture (`fixtures/` is
reserved and still EMPTY) plus a real async LLM import"* — but the UC's own flow says **"choose an import
source"**, and LinkedIn is a source. If the LinkedIn source takes a URL, the fixture half of that blocker is
false.

## Expected lift

Clause 3 landed-half **onboarding 1 → 2 of 5**, *if* a UC proves deterministically landable. If the only
route to a UC runs through a live third-party dependency, the honest outcome is a **sharpened written
verdict** instead — the milestone holds 28/28 written verdicts and 0 `unimplementable`, and a verdict built
on measurement is worth more than the unmeasured one it replaces.

## Phase plan

- **A — probe all four blockers live** on `demo-2`, before any harness code.
- **B — implement** whatever A proves landable; otherwise land the durable assets A produced and rewrite the
  verdicts from measurement.
- **C — mutation-verify** every assertion added: watch it go RED.
- **D — re-measure** 3 consecutive cold reset-to-seed runs; restore the drifted cockpit fixture + sha-verify.

## Escalation conditions

- A UC proves `unimplementable-without-platform-edit` → the re-scope trigger counts it (`> 3` escalates).
  The milestone holds **0**; do not spend one casually.
- A route exists but only through a **live third-party network dependency** → this is a **P6 determinism**
  decision, not a capability one. Record the measurement, refuse the landing, say so in the verdict.

## Acceptable close-no-lift outcomes

All four blockers driven and re-priced from measurement, with the assets that pricing required checked in —
even if no UC lands. What is **not** acceptable is re-deferring any of the four on the strength of a note
nobody drove.

## Scope

**One line of investigation: the onboarding cluster.** The last sharpenable negative control
(`pt-hiring-recruiter-compare`, controls 22 → 23) is the strongest remaining clause-2 target and is
deliberately **routed to iter-19** rather than opened here — the tripwire's rule is to land what is complete
and route the rest.
