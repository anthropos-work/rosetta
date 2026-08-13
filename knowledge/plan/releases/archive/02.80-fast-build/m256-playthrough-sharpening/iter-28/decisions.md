# iter-28 — decisions

## D111 — the trigger was found by READING the platform, after a probe sweep could not find it

**Context.** `standard.UC2`'s blocker read *"the trigger is not yet identified"*. iter-18 had driven heroes
across **four** orgs (A, C, D) and got the identical import form every time, which established two negatives —
it is not *"the member has a profile"* and not an org-level narrative flag in this seed — and no positive.

**Found in one `useState`.** `packages/ui/src/Onboarding/OnboardingUser.tsx:135`, in the shared component
**both** apps mount:

```ts
const lastStep = reimport ? Import : steps?.[steps.length - 1]?.step;
const [managerImport] = useState(
  Boolean(lastStep === OnboardingStep.Import && organizationName && userStats));
```

`organizationName` and `userStats` are always supplied by the host page (`userStats` is a memo over an object
literal, so never undefined), which leaves **`steps`** — i.e. `public.user_params.onboarding`, NULL for every
seeded user. `managerImport` swaps `<ImportStep>` for `<EnterpriseUser>` (name + location + last experience +
stat cards) and relabels the forward control to *"Start"*.

**Why probing could never have found it.** Every seat in the world has the same `user_params.onboarding` value
— NULL — so *no* choice of hero, org, narrative or entitlement varies the input. A four-org sweep was sampling
a constant. **The generalisable rule: when a probe sweep returns the same answer for every vantage, the input
is not one of the axes you are varying, and more vantages will not help.** iter-27's D108 is the same lesson
from the other side (a routing *cause* found in one `useEffect` after four orgs of probing).

## D112 — the insert alone silently did nothing, and the seeder now fails loudly instead

`public.user_params` is populated **row-per-user at user-insert time** — 191 rows appeared within 300 ms of the
users COPY, all with `onboarding` NULL, written by **nothing in this seeding fleet**. So the natural
`CopyRowsIdempotent(..., "id")` (ON CONFLICT DO NOTHING) found the row already present and **skipped it**, with
no error anywhere: the seat kept being served the plain import form and the only symptom was *"the product does
not show the prepared summary."* That is the misattribution shape this milestone hunts, arriving through an
idempotency clause.

**Decision.** Insert-then-**heal** (`UPDATE … SET onboarding = $2::jsonb`), the same shape `users.go` uses for
`memberships.picture_url` (M44 §D fix1) — **and fail the seed** when a declaring hero's row could not be
reached. A capability whose no-op looks like a product defect must not be allowed to no-op quietly.

*Also worth recording:* the missing `audit.Record` was caught by the isolation guard on the first live run
(*"surface reports 1 row written but recorded NO audit entry"*). Two guards fired in this iter and both were
right.

## D113 — the Playthrough is ROUTED, not attempted, and this is a deliberate stop

**What landed:** the capability (4 mutants RED), the seat, and the **live proof** that the variant renders —
same run, two seats, total discrimination:

| locator | `pt-onboard-prepared` | `pt-free` (day-0) |
|---|---|---|
| `linkedinUrlInput` | **0** | 1 |
| `importFromLinkedInLabel` | **0** | 1 |
| `uploadButton` | **0** | 1 |
| `skipButton` | 1 | 1 |

**What did not, and why.** The UC's flow continues *"confirm or adjust the pre-filled role, refine the suggested
skills"* — and on the prepared variant the forward control is relabelled **"Start"** (`nextLabel` →
`startLabel`), whose handler (`OnboardingUser.tsx:470`) is a branch nobody has driven. Whether `Start`
**completes** onboarding or **advances** to the Role step is therefore unmeasured, and the session budget ran
short of measuring it.

Writing the spec anyway would mean asserting a multi-step journey nobody drove — precisely the failure iter-22
paid for (a spec parked since iter-04 that *could not have passed*) and iter-27 paid for again (a "read-back"
that read the seed). **The iter's own escalation condition anticipated this and named the stop.** So the
Playthrough is routed with the measurements a successor needs, and the residual is one probe run wide rather
than one unknown wide.

**Status is `closed-fixed-partial`, deliberately:** planned scope was the capability *and* the UC; the
capability landed clean and proven, the UC is routed as Fate 3 with a named handler.

## D114 — the seat lives in ORG B, and a RED gate is what taught that

**What happened.** The seat was first appended to **Org A**, and `pt-workforce-funnel` went **RED in all three
gate runs** — deterministically, 16.7 s each, identical message: its iter-14-sharpened final asserts that
**Pat Ellis's member-spotlight CARD carries her seeded role**, and one extra Org A member displaced Pat from
the spotlight entirely.

**This is D107 one axis over.** There, a hero's *role occupancy* perturbed a succession key-role card
(probabilistically). Here, an org's *member set* perturbed a member spotlight (deterministically). The
underlying fact is the same one and it is now twice-paid-for: **iter-13/14 deliberately re-aimed the negative
controls at seeded facts by NAME, and the cost of that sharpness is that adding a hero to an anchored org
perturbs another Playthrough's anchor.**

**The fix, and the rule it yields.** Org B (`pt-halcyon-retail`) is the **only** pt org that
`e2e/lib/seed-facts.ts` does not name — `SEEDED_ORGS = [PT_ORG_A, PT_ORG_C, PT_ORG_D]` — so no sharpened final
anchors on its composition. It is also a workforce org, which is what this use case needs. Moved there,
renamed (`Elin Marchetti`), role kept distinct from `pt-free`'s per the D107 fence. Gate: **196 passed ×3, rc 0,
0 flake.**

> **THE RULE, stated for the next author: before appending a hero, check whether `seed-facts.ts` names her org.
> If it does, expect to perturb an anchor — and prefer Org B, which nothing anchors on.**

**Re-proven on the SHIPPED seat** (not just the Org A draft), which matters because `organizationName` is one of
`managerImport`'s three inputs and the org changed:

| locator | `pt-onboard-prepared` (Org B) |
|---|---|
| `linkedinUrlInput` / `importFromLinkedInLabel` / `uploadButton` | **0 / 0 / 0** |
| the relabelled **`Start`** control | **1** |
| her own name on the summary | **1** |

That last row is new evidence the Org A probe did not have: the prepared summary genuinely renders **her**
(`<EnterpriseUser fullName=…>`), and the forward control really is relabelled — both of which the routed
Playthrough will assert.
