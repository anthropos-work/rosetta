# iter-27 — decisions

## D108 — the cross-app ROUTING half of the hiring UC is deliberately NOT asserted

**Context.** `onboarding.enterprise-hiring.UC1`'s curated intermediate reads *"the member is routed into the
hiring app, not the workforce app."* The manifest had flagged, for two releases, that a final observing this
might be proving the **cockpit's** routing rather than onboarding's, and that the discriminator must be found
before the use case could be honestly asserted.

**Found, by reading the source rather than by driving the surface.**
`apps/web/src/context/UserStatusContext.tsx:142-173` computes `userHasAllHiringOrgs` and, if true, hands off to
the hiring app. **There is no onboarding condition anywhere in that effect.** So the routing is owned by
*neither* onboarding *nor* the cockpit — it is a membership-shape redirect that fires on any `apps/web` page
load, for a member who has onboarded or not.

**Decision.** Do not assert it. Drive the **hiring** app — which is where the platform puts an all-hiring-org
member anyway — and assert the half onboarding actually performs: the hiring app's *own* onboarding route
(`apps/hiring/.../onboarding/page.tsx`) closes with `router.replace('/home')`, so completing the flow is what
moves her. The un-asserted half is **named** in the spec header, the manifest note and the corpus, with the
`file:line`, rather than dropped silently.

**Why this is landing the use case rather than dodging it.** The curated final is *"the member lands in the
hiring app with a hiring simulation assigned and ready to start"* — and that is what is proven, from the hiring
app's own flow. What is refused is a *green over `UserStatusContext`*, which would be the M219 lesson ("a
surface that renders is not the same as the RIGHT surface") committed one level up: the right surface, proven by
the wrong mechanism.

## D109 — TWO mutants passed, and both changed the spec. On a seeded world, "present after" is not evidence.

Both of these are the same defect wearing two hats, and neither was visible without the mutant.

**Q1 — skip the write entirely, navigate to `/home` by hand: PASSED.** Every assertion still held. Her home —
the greeting, the tenant chrome, the assigned position — is **seed state that exists before she onboards**. The
spec called its block a "read-back on a fresh navigation" and it was reading something the write never touched:
a Playthrough that would have gone green **without anyone onboarding**, which is the precise failure iter-24
caught in the *seed* and iter-22 caught in an *assertion*, arriving this time through the *page object*.

*Fix:* delete the manual navigation. `apps/hiring`'s onboarding closes with `router.replace('/home')`, so
waiting for the app itself to move her IS the write's observable consequence. **Q1b** (the same mutant against
the fixed spec) is **RED** — `waitForURL` times out on `/onboarding`. That one line is now the only assertion
here that the write can satisfy, and it is labelled as such in the source.

**Q5 — seed mutant, `trajectory: thriving` (→ ASSESSED, i.e. she has already taken the positions): PASSED.** So
the `assignedOnly` state is **not discriminated by this surface**: the candidate's home renders the org's
position as a startable, org-scoped link whether or not she has taken it.

*Response — state it, don't paper over it.* The final asserts the **affordance** (her position, in her tenant,
with a real title, startable). The "not yet taken" half is a **seed guarantee** (`heroHiringStage` →
`assignedOnly`), true by construction and untested by the UI because the UI does not surface it here.
`trajectory: struggling` is kept because it is the truer state for a day-0 candidate, **not** because an
assertion depends on it — and the spec says exactly that, so the next author does not assume it is load-bearing.

**The general rule, and it is the most transferable thing in this iter:** on a **seeded** world, *"the outcome
is present after the action"* is not evidence — it is only evidence if the outcome was **absent before**. Every
mutating Playthrough here should be able to answer "which single assertion fails if I delete the action?", and
if the answer is "none", the read-back is reading the seed. **Q1 is now a standing mutant worth running against
any new mutating Playthrough on this world.**

## D110 — the "assigned and ready to start" half is a declared seed guarantee, not an asserted claim

Recorded separately from D109 because it is the *scope* consequence rather than the *method* one: this use case
lands with its final scoped to what the surface can support. Same shape as iter-26's D105 P6 boundary — the
unproven half is named where a reader will find it (spec header, manifest note, corpus section), so the
Playthrough's green cannot be read as more than it is.
