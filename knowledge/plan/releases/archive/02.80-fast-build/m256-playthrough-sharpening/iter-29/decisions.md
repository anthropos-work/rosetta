# iter-29 — decisions

## D115 — the fresh-navigation read-back was REMOVED, because a mutant proved it could not fail

**The assertion, and why it looked right.** The spec's first draft ended with `goto('/onboarding')` and asserted
the prepared summary was **gone**. The reasoning is sound and still is: `managerImport` is
`lastStep === Import && …`, so once a `role` step is persisted the route *cannot* re-open on `<EnterpriseUser>`.
That made it look like a textbook server-side read-back on a fresh navigation — the shape this milestone asks
for.

**The mutant (iter-27's standing Q1, run as S1c).** Delete the action *and* the intermediates, leaving only that
read-back. It **PASSED**. `toHaveCount(0)` immediately after a navigation is satisfied by a page that has **not
hydrated yet** — so the assertion was green whether or not the write happened.

**Decision: remove it, do not weaken it.** An absence assertion that a mid-hydration page satisfies is worse
than no assertion, because it *reads* as proof. The honest repair needs a POSITIVE locator on the screen a
reload actually lands on — which is the **Role** step (the component's initial step is `lastStep || Import`, and
`lastStep` is now `role`), a screen nothing has driven. Asserting an undriven screen is the failure iter-22 and
iter-27 each paid for, so the persistence half is **routed** with the measurement it needs. The DB write itself
(`role` appended to `public.user_params.onboarding`) was measured directly during authoring.

**Why this is the third variation on one theme, and worth naming as such.** iter-12 established
*liveness-before-absence* after an ablation produced a dead page that satisfied every absence assertion.
iter-22 found `rows > 0` satisfied by a table saying *"No roles match your filters"*. This is the same defect
with **time** as the confounder rather than an empty state or a dead page: the page is fine, it simply is not
there yet. **An absence assertion needs a companion that proves WHEN it was read, not only WHERE.**

**What the Playthrough proves instead, and it is mutation-proven both ways:**

| mutant | outcome |
|---|---|
| **S1** delete the click | **RED** at "Change Role" — the action is load-bearing |
| **S2** remove `onboarding: org_prepared` and RESEED | **RED** at the liveness assert — the summary never renders and the import form returns; the SEED is load-bearing |
| **S1c** only the fresh-nav read-back remains | **PASSED** → removed (this decision) |

So the UC lands on two proven halves — *the prepared summary is served **instead of** the import form* (a
cross-vantage control against every other seat in the world) and *confirming advances into the skills step with
her role's real taxonomy skills* — with the persistence half named rather than faked.
