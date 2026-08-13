# iter-31 — decisions

## D118 — the routed persistence repair was IMPOSSIBLE, and the fourth mis-stated blocker in six iters

iter-29 routed `ONBOARD-M256-prepared-persistence` as **one measurement**: put a POSITIVE locator on the ROLE
screen a reload lands on, *"because the component opens on `lastStep || Import` and `lastStep` is now `role`"*.
That reasoning is a source read with no observation behind it, and it is **wrong**. A fresh navigation re-serves
the **pre-state** screen — byte-identical, progress `0`, `[Skip] [Start]` — because the flow reads its "last
step" off the wrong end of a newest-first array. Full mechanism + both `file:line`s in the milestone
`decisions.md` § `PLATFORM-M256-onboarding-step-not-resumed`.

**This is iter-30's D117 recurring one iter after it was written.** D117: *a routed blocker must carry the
measurement that produced it, or be marked `(estimate, unmeasured)`.* iter-26's *"needs a stage-0 seat"*,
iter-28's *"trigger not identified"*, iter-30's *"a seeder change"* — and now iter-29's *"the reload lands on the
ROLE step"*, which is subtly worse than the other three: it is not a *pricing* estimate but a **claim about
observable behaviour**, phrased with a `file:line` beside it, so it reads as though it had been measured. **A
source read is a hypothesis. Citing a line number does not make it an observation** — the extension D118 adds to
D117.

## D119 — what to ship when the honest assertion does not exist: label it, do not invent it

There is **no** locator on `/onboarding` that discriminates confirmed from not-confirmed (the two screens are
textually identical), and no route that carries the confirmation. Three options were on the table:

1. **Restore an absence assertion** — forbidden by D115 and by the new fence; it is what could not fail.
2. **Extend the journey to completion** and assert the `/onboarding` → `/home` redirect, which the platform
   *does* honour. **Measured and refused:** the flow dead-ends at *"Add more skills"* with an error banner and an
   inert `Next` (five clicks, identical screen). It is unreachable, so a spec written for it *could not have
   passed* — the iter-22 failure exactly.
3. **Assert what IS observable, and label it for precisely that.**

(3) ships: a positive, hydration-proof assertion that a fresh navigation still serves *her* flow — the
confirmation neither ejected her nor completed it — with the docstring, the failure message and the manifest all
stating that **it is not the write's read-back**. It discriminates two real failure modes (a `done` step would
redirect; a broken flow would not serve), watched RED three ways (S3 route-anchor, S4 identity-anchor, S5 a
planted `done` step). The write stays proven by S1 against the screen the click reaches.

**The rule:** when the honest assertion does not exist, the deliverable is the **measurement plus the label**,
never a weaker assertion wearing the strong one's name. A comment claiming more than the code checks is the same
defect as an assertion that cannot fail — iter-31 found the page object's own docstring asserting that
`changeRoleControl()` *"on a FRESH navigation is the server-side read-back"*, written in good faith at iter-29
and false. **It has been corrected in place**, because a docstring is read as evidence by the next iter.

## D120 — a fence for a rule three iters applied by hand, sized by measurement first

Absence-after-navigation has now failed three times in three costumes — **dead** (iter-07, `bodyLen 24`),
**empty** (iter-22, the placeholder row that *is* a row), **not there YET** (iter-29, pre-hydration). The third
is why the first two fixes did not prevent it: TIME was the confounder, and *an absence assertion needs a
companion that proves WHEN it was read, not only WHERE.* `negative-controls.spec.ts` carries the rule in prose,
and iter-29 broke it **in a brand-new spec, one hour after writing that prose**. A rule re-broken that reliably
belongs in a test.

**Sized before adopting** (iter-15 D74's discipline): **29 files · 62 navigation sites · 184 liveness witnesses ·
37 absence assertions · 0 violations.** So the invariant is already true everywhere and the fence costs **zero
edits** — it buys the *next* spec. **Fail-closed on three floors**, because a scan that matches nothing passes
every assertion, and a fence is the worst place in the tree to commit this milestone's signature defect. Proven
discriminating by re-introducing iter-29's exact S1c shape and watching it name the `file:line`.

Deliberately scoped to `tests/`, not `lib/`: page objects own locators, not assertions, and the
navigation-then-assertion adjacency this rule is about only exists in a spec.

## D121 — a `TODO` must carry a WRITTEN VERDICT, fenced in both directions and against vacuity

The exit gate asks for *"a written verdict for every remaining uncovered use case — zero silent gaps"*, and that
clause lived entirely in prose. The machine-readable half of an uncovered use case was `playthrough: TODO`, so
`ptreport` rendered **every** one with the same sentence: *"declared use case, no Playthrough yet
(build-reference gap)"*. For `onboarding.enterprise-workforce-standard.UC1` that sentence is **false** — it is a
measured refusal, not an effort gap — and the four-state map is what tooling and reviewers actually read. *The
gate was about to be closed around a use case the artifact described wrongly.*

**Decision.** A use case with no Playthrough carries a `verdict` block: a **closed** disposition
(`will-not-build` | `not-yet-built`), a `measured_by`, a `rationale`, and a `handler`. `ptreport` renders that
instead of the generic sentence. The four states and their glyphs are **unchanged** — `will-not-build` is not a
fifth state; a new glyph would imply a state the reconciler does not have.

**Three design choices, each with a reason from this milestone's own history:**

- **Bidirectional.** A TODO without a verdict fails; a **live** use case that still carries one fails too. A
  stale verdict on a proven use case is a claim with **no expiry** — nothing tells a reader it is out of date.
  Landing a use case must *force* the verdict's removal rather than leave it to diligence. (M255's knob guard is
  the precedent: a doc-promised flag with no parser entry is a false promise; a parser flag with no doc row is
  undiscoverable.)
- **Anti-vacuity.** A presence check is satisfied by `rationale: TODO` — the same silence in a new schema. Hence
  an 80-character floor, nine placeholder spellings blacklisted, and `measured_by` required (the mechanical form
  of **D117**).
- **The enum is closed with NO fallback member.** This milestone has twice shipped a seed enum whose
  unrecognised value fell back to a permissive default (iter-26 `ai_readiness` → stage 3 COMPLETED; iter-28
  `onboarding` → the day-0 form), and in both cases the fallback would have produced a Playthrough that looked
  like a **product regression**. A verdict that falls back is a verdict nobody wrote.
- **The handler asymmetry carries the meaning.** `not-yet-built` MUST name a handler (a gap with nobody assigned
  is the *"later"* the three-fate rule forbids); `will-not-build` must NOT (a refusal with an assignee is a
  contradiction). Without the asymmetry the two dispositions blur into *"TODO with a paragraph attached"*.

**Two existing tests failed, and that was the point.** `TestValidate_TODOIsLegitimate` and the `ptvalidate`
single-file fixture both encoded *"a bare TODO is legitimate"* — the precise position this decision revokes.
Renamed and updated, never relaxed.

**Proven on the SHIPPED manifest, not only on fixtures** (iter-16's lesson: five green unit tests once drove a
mock path the real client never used). Six live mutants RED: the block deleted · an unrecognised disposition · a
blank `measured_by` · a gap with no handler · a refusal *with* a handler · a stale verdict on the use case
iter-29 landed.

## D122 — `standard.UC1` is NOT landed, and no route to landing it was sought (D104 upheld)

Recorded explicitly because the temptation is real and iter-18's refusal is now in its **thirteenth** iter. The
iter's plan forbade looking for a way to land it; nothing was probed, nothing was drafted, and the deliverable
was the **verdict**. `disposition: will-not-build`, `measured_by: iter-18 + D104`, no handler. The reasons are
unchanged and are now in the artifact the tooling reads rather than only in a story `note`: the only advancing
path scrapes a live public third-party profile from a site that blocks automation (a real person's profile would
become a permanent fixture, and **its RED would read as a product regression**), and the deterministic route is
blocked by a measured product defect we do not own.

## Safety

Every write was to **demo-2's own Postgres**. Production was neither written nor read. The one hero row touched
(`public.user_params` for `elin.marchetti2@…`, to reach the pre-confirm state and later to plant the mutant
`done` step) was **backed up first** and **restored byte-identically** (`diff` clean, verified twice). The
DRIFTED cockpit-manifest fixture was backed up before the gate and **restored + sha-verified `99e2f315` after
each of the three `--reset` runs**. `stackseed --policy-check --stack demo-2` rc 0, `live=18 expected=18`.
`docker ps -a`: 16 Up / 0 exited. **Zero platform-repo edits** — the platform tree was read only
(`OnboardingUser.tsx`, `useGetOnboardingStatus.tsx`, `onboarding/page.tsx`).
