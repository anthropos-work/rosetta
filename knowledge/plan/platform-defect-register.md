# Platform-defect register

**Defects in the `anthropos-work` PLATFORM repositories that Rosetta's tooling found, and that Rosetta
cannot fix.** Zero platform-repo edits is a standing constraint, so every entry here is a **report**, not a
work item for this repo.

## Why this file exists

It was created at the **M256 close (v2.8 "fast build", 2026-07-30)** because the deferral audit found a
structural gap: **there was no platform-defect register anywhere in this repo.** M256 alone routed four
defects "to the platform", and all four lived only in that milestone's `decisions.md` — a file that flips to
`archived` at close and is never read again. A defect recorded inside a closed milestone has been *filed
where it cannot be found*.

The failure mode is the one the milestone spent 32 iterations on, one level up: a routing that looks
discharged because it was written down somewhere. M255's close had already been bitten by the sibling version
(four items routed to *"M255 harden resume"*, which was not a milestone and could not hold them).

## How to use it

- **Append, never rewrite.** An entry is evidence with a date.
- **Every entry carries `file:line`** for the deciding code. A defect report a platform engineer has to
  re-derive is half a report — and re-deriving it may mean re-doing the measurement that found it.
- **Distinguish MEASURED from INFERRED, per entry.** Several of these were found by driving a live demo; some
  siblings are reasoned from a shared code path and were never driven. The distinction is load-bearing and
  the entries say which is which.
- **Mark an entry `FIXED` with the platform commit** when it lands upstream. Do not delete it.

---

## Open

### `PLATFORM-M256-onboarding-step-not-resumed` — the org-prepared onboarding flow cannot resume
**Found:** M256 iter-31 (2026-07-30) · **Repo:** `next-web-app` · **Status:** open · **Severity:** high (the
first thing a real member does)

A member whose org pre-filled her profile confirms her role. The confirmation **persists** server-side
(`public.user_params.onboarding` gains a `role` step — verified in the DB on every attempt). She reloads
`/onboarding` and is back on **step one**, progress `0`, with no trace of what she confirmed; the screen is
**byte-identical to the pre-state**. Six fresh navigations across three browser sessions, hours apart, all
agree. She can never advance past the first step across a page load.

**Mechanism — one array, two consumers, opposite ends:**

```
packages/graphql/src/hooks/onboarding/useGetOnboardingStatus.tsx:25-27
    result.onboarding.steps?.sort((a, b) => sorterFn({ first: b.updatedAt, second: a.updatedAt }))
    → the array handed to the component is sorted NEWEST-FIRST
packages/ui/src/Onboarding/OnboardingUser.tsx:130-132
    const lastStep = reimport ? Import : steps?.[steps.length - 1]?.step;
    → takes the LAST element of a newest-first array, i.e. the OLDEST step ever taken
```

So `lastStep` is the *first* step the user ever completed, `managerImport` is true again, and the initial step
resolves to `Import` forever. The **host page reads the same array from index 0**
(`apps/web/.../onboarding/page.tsx:141-143`) and is therefore correct — which is why completion redirects
properly and nothing else ever looked wrong.

**Why nobody had seen it:** invisible for a NULL or single-element `steps` array (`length-1 == 0`, so both
readings coincide) — which is **every one of the 191 seeded users** and every hero any earlier iter could
reach. Only a multi-step array exposes it, and the only multi-step user in existence is a seat M256 iter-28
minted to reach the surface at all.

**Second defect on the same journey.** The prepared flow **cannot be completed on a demo**: one `Next` past
the skills screen reaches *"Add more skills"*, which renders *"We're having trouble loading your skills at the
moment"*, and its `Next` is **inert** (clicked five times, identical screen, progress stuck at 100).
`useClusterizeSkills` is the surface behind it.

**Provenance:** source read + six live observations. The mechanism above is a **source read**, stated as such.

---

### `DEFECT-M256-silent-forbidden-mutation` — a refused mutation renders nothing at all
**Found:** M256 iter-20, measured iter-23 (2026-07-30) · **Repo:** `next-web-app` · **Status:** open ·
**Severity:** high (it hid a real authorization gap for fifteen iterations)

A mutation the backend **refuses** is, from the user's side, indistinguishable from one that was never sent.
Reproduced deliberately on `demo-2` only (the `p3 admin → org:feature:taxonomy:write` grant revoked, the
journey driven, the grant restored byte-identically, `--policy-check` rc 0 afterwards).

**Measured across every channel a user or operator could learn from:** HTTP **200** with the error inside it ·
`[role=alert]` **present and EMPTY** · no `[role=status]` · antd `message`/`notification`/`form-item-explain`
all empty · the dialog **stays open with `Save` still ENABLED**, inviting an identical retry · URL unchanged ·
catalog total **49 → 49** · browser console says nothing about it · one **uncaught page error**.

**Two defects, one symptom:**

1. `packages/ui/src/JobRoles/Form/AddJobRole.tsx` `handleSubmit` handles exactly one error shape and
   `throw error`s the rest out of an async click handler — an unhandled rejection React renders nothing for.
   `onClose()` sits after the try/catch, which is why the dialog stays open. The empty `[role=alert]` is the
   **duplicate-warning slot**, never populated — *the app has one error surface here and it is reserved for a
   different error*. (`throw error;` from a catch appears **exactly once** in all of `packages/ui`.)
2. **The systemic half:** `apps/web/src/providers/Query.provider.tsx` sets
   `mutations: { onError: (e) => { captureException(e); PosthogClient.captureException(e) } }` — Sentry and
   PostHog, **no user surface**. Every mutation in the app is silent on failure unless it builds its own.

**And a dead contract that makes it look handled:** six mutations across four `hooks/organization/*` files
declare `meta: { error: '…' }` human-readable failure sentences. **No handler reads them** — there is **no
`MutationCache` anywhere** (0 occurrences); the only `meta.error` consumer is `QueryCache.onError`, which uses
it as a **Sentry tag**. So the strings are inert, on precisely the org-admin write set. *The authors wrote
failure messages and the framework never wired them up*, which is a more useful report than "the form is
silent" because it names a fix using a convention the codebase already believes it has.

**Suggested fix (not applied):** add a `MutationCache` whose `onError` reads `mutation.meta.error` and renders
it; replace `AddJobRole`'s `throw error` with the same path. Turns six dead strings live and gives every
future mutation a default surface.

**⚠️ Sweep residual — MEASURED vs INFERRED, stated because the claim is a negative.** Only `createJobRole` was
refused **LIVE**. The dead `meta.error` strings and the Sentry-only global handler are **definitive by
source**. That a refused tags-create / member-tag / settings-toggle would look **equally silent** is an
**inference** from the shared global handler — it was never driven, because each would have meant another
revoke/restore cycle on a stack later iters depended on. *Driving one sibling refusal closes that gap in one
revoke.*

---

### `PLATFORM-M256-keyrole-nondeterminism` — a succession key-role card appears nondeterministically
**Found:** M256 iter-26 (2026-07-30) · **Repo:** platform (succession ranking) · **Status:** open ·
**Severity:** low-medium (not a defect a presenter would see; it reddens automated batches)

A succession **key-role card**'s presence varies between page loads once its role has **2 occupants** —
measured **4 of 5 loads at occupancy 2** against **5 of 5 at occupancy 1**. Most plausibly a top-N ranking
with an unstable tiebreak.

**Cost, recorded because it is the reason this is filed rather than shrugged at:** two gate cycles. The
iter-14 cross-tenant negative control anchors its LIVENESS floor on that card, so it went RED reading
*"succession failed to compute for the contrast tenant"* — and a 45 s timeout did **not** fix it, because the
cause was the seed's role occupancy, not the clock.

**Mitigated our side, not fixed:** hero roles must be pairwise distinct within a story
(`playthroughs/e2e/tests/seed-facts-fence.unit.spec.ts`, mutant N1 RED). A batch can still redden on this.

---

### `PLATFORM-M256-cv-upload-never-parses` — a valid CV upload POSTs 200 and never advances
**Found:** M256 iter-18 (2026-07-29) · **Repo:** `next-web-app` / the import pipeline · **Status:** open ·
**Severity:** medium — **it is the reason a curated use case is `will-not-build`**

The profile-import CV route POSTs **200** for a valid PDF **and** for a docx alike, while the forward control
**never enables** (waited 100 s+). Measured with a purpose-built synthetic fixture
(`playthroughs/fixtures/synthetic-cv-sre.{pdf,docx}`, a wholly invented CV whose employers and school occur
nowhere in the seed, the taxonomy, or any real registry — so an assertion naming them can only be satisfied by
*that file having been imported*).

**Consequence for coverage, and it is the honest kind:** this is the deterministic alternative that would have
let `onboarding.enterprise-workforce-standard.UC1` be a Playthrough. With it blocked, the only advancing path
**scrapes a live public third-party profile** on a site that blocks automation — so the use case carries a
machine-checked `disposition: will-not-build` verdict instead (M256 `D104`/`D122`). **The two fixture files
ARE the evidence for that verdict**, which is why they ship despite having no consumer.
