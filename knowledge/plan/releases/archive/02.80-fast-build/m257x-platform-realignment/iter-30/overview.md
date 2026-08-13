---
milestone: M257x
iter: 30
iteration_type: tik
status: closed-fixed
---

**Type:** tik (under `TOK-01: instrument first, then follow`)

# iter-30 — the role text that two surfaces stopped rendering

## Active strategy reference

`TOK-01: instrument first, then follow` (milestone-root `decisions.md`). This iter is squarely inside it:
step 1 repairs the instrument that a diagnostic run corrupts, step 2 uses the repaired instrument.

## Step 0 — re-survey (mandatory), and what it already refuted

Clause 2's failing set is **not re-measured** this iter: iter-29 established it is deterministic across
three full `--reset` runs on an unchanged build, and the hand-off's own rule is that re-measuring to
"confirm" is waste. The re-survey instead **re-read the binding artifact per-id** — `e2e/report/last-run.json`
from run C — which is a thing no prior iter appears to have done for all five ids at once.

It refuted the inherited framing on **three** counts before any work started:

1. **`FIX-M257x-iter27-succession-hero-not-rendered` is misnamed.** The failing assertion in
   `workforce-succession.spec.ts:83` is `getByRole('heading', { name: /^DevOps Engineer$/ })` — a **key-role
   card heading**. It is not a hero-absence assert at all. The hero-absence assert is `:90`, and the run
   never reached it.
2. **The hand-off's taxonomy claim is false.** It stated *"the taxonomy has **no** exact `DevOps Engineer`
   role"*, listing near-misses. Measured on live `demo-1`: `public.job_roles` holds **`DevOps Engineer`
   exactly** (among 32 `%devops%` rows). The "the card may render a resolved taxonomy role that legitimately
   can't match" escape hatch the hand-off offered **does not exist**.
3. **The seeded data is intact on every axis the surfaces could read.** Pat Ellis
   (`23f24e3f-38fb-5027-9e07-2ef49a644af5`) has `memberships.job_role_id = J-DEVOPS-AECC` **and**
   `job_role_name = 'DevOps Engineer'`; 40/40 Meridian Labs members carry both.

## Cluster / target identified

TOK-01 names no stale target; the hand-off routed two "next tik" items. Both survive re-survey, and the
re-read shows they are **one candidate signature, sized honestly at two ids, not four**:

| id | failing assert | shape |
|---|---|---|
| `pt-workforce-funnel` | `…filter(hasText:/^Pat Ellis/).getByText('DevOps Engineer')` not found — **her card IS visible** (preceding assert passes) | the string is absent inside a rendered card |
| `pt-workforce-succession` | `heading /^DevOps Engineer$/` not found | the string is absent as a role-card heading |

**Deliberately NOT folded in** (the milestone has dissolved three clusters on measurement, most recently 2 of
4): `pt-activity-drilldown` fails on `heroRow.count() > 0` — the hero's **name** is absent from the
per-member breakdown, a different claim, and it never reaches its own role assert. `pt-orgadmin-role-create`
is a 60 s `waitForURL` timeout. `pt-onboarding-hiring-candidate` is a missing `/sim/…organizationId=` link.
Three singletons until measured otherwise.

## Hypothesis

**H1 — read-side.** The seeded role is present in the DB on both the id and the name axis, so the two
surfaces have stopped *rendering* a value they still hold. Cause is app-side render/DOM, not seed data.

**H2 — locator-side.** The value renders and the page objects no longer find it (the `^`-anchored
`hasText` in `memberSpotlight`, the `^…$`-exact heading match in `keyRoleCard`) because the DOM shape moved
— iter-22 found an uncaught Next 15→16 upgrade in this app.

H1 and H2 are distinguished by **one live DOM read** of the two surfaces. That read is the iter's core measurement.

## Prediction, recorded BEFORE the measurement

*(the iter-28/29 discipline — a number that beats its prediction deserves more suspicion than one that meets it)*

- I predict the DOM read finds Pat Ellis's spotlight card **rendering some role-ish string that is not the
  literal `DevOps Engineer`** (blank, truncated, or a different label), i.e. **H1**.
- **Declared acceptable in advance:** if the read finds `DevOps Engineer` present in the DOM and the locator
  simply fails to reach it (**H2**), that is an equally valuable result and **not a failure of the iter** —
  it converts two "platform drift" failures into a page-object repair, which is rext-side and landable.
- **Also declared acceptable:** finding that the two ids do **not** share a cause. The signature is a
  hypothesis, not a grouping.

## Phase plan (2 planned lines — declared, so the scope-creep tripwire counts against THIS shape)

1. **Instrument repair — `FIX-M257x-iter27-scoped-run-clobbers-binding-report`.** A scoped `--grep` run
   overwrites `e2e/report/last-run.json`, and nothing in the file says which invocation produced it. Every
   diagnostic run in step 2 is a scoped run, so this is a **prerequisite**, not a bundled extra. Make the
   artifact self-describing (record the invocation's scope) and fence it.
2. **Live DOM read + root cause.** Drive the two surfaces as the manager hero, dump the relevant subtree,
   decide H1 vs H2, and land the fix if it is rext-side.

Anything else that surfaces routes forward with a named handler.

## Expected lift

Clause 2 is `25 live / 5 failing / 1 unimplemented`. A landed fix for the shared cause would move it toward
`27`. **No re-measurement of the full suite is planned in this iter** — a binding run is ~35–40 min serial
and is its own iteration. This iter's deliverable is the instrument repair plus a *root cause named with
evidence*; the confirming full run belongs to iter-31.

## Escalation conditions

- The fix requires a platform-repo edit → **route forward, do not edit**. Zero platform edits is binding.
- Platform origin moves off `2adcf71` → re-scope trigger occurrence 2 → STOP and escalate.

## Acceptable close-no-lift outcomes

H2 confirmed with the locator repaired but unverified by a full run; or H1 confirmed and root-caused to a
platform-side render change that rext must not fix — either is a complete iter, because the deliverable is
the named cause plus the instrument repair.
