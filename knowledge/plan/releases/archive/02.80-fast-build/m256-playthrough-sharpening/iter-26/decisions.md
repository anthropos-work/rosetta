# iter-26 — decisions

## D104 — the stage-0 capability is a FIELD OF ITS OWN, not a third `trajectory` value

**Context.** iter-24 measured that a stage-0 (day-0) end-user readiness seat **cannot be declared**:
`aiReadinessStageFor`'s hero branch is manager → 0, struggling → 1, **everything else → 3**. The natural first
instinct is a third trajectory — `trajectory: day-zero` reads well in YAML and slots into an existing enum.

**Decision.** `blueprint.Persona.AIReadiness` (`""` | `"not_started"`), read in **exactly one place**
(`aiReadinessStageFor`), checked **before** the trajectory-derived default.

**Why.** Two reasons, both concrete:

1. **Truth.** Readiness stage is **orthogonal** to the life-arc. The org's diagnostic is a *time-boxed cycle*;
   a member who joined mid-cycle has verified skills, a career, self-ratings and months of activity while
   having done none of its three steps. `trajectory: day-zero` would have asserted the opposite — that her
   skills, level band, growth arc and succession signal are all day-0 too.
2. **Blast radius.** Five non-test files switch on `Trajectory` (`persona.go` ×2 bands, `succession.go`,
   `hiring_funnel.go`, `cockpit.go`), each with a `default:` branch. A new enum value lands in every one of
   those defaults **silently** — a day-0 hero would have picked up the mid level band (62–85), 5 verified
   skills by default, `EffectiveSelfRated() == true`, and a positive succession signal, each by accident
   rather than by decision. The new field touches one call site and cannot leak.

**Corollary — the enum is CLOSED, and that is load-bearing.** An unrecognised value falls back to the derived
stage, which for an end-user is **3 = COMPLETED** — i.e. a typo does not produce a broken seat, it produces
exactly the already-finished seat the capability exists to avoid, and the Playthrough would go green on the
wrong surface. `validate()` therefore rejects anything but the two legal values by name
(`TestValidate_AIReadinessEnum`, mutant M4 RED).

**Not chosen:** a raw `ai_readiness_stage: <int>`. It admits 2 and 3, neither of which any seeder path was
built or tested for, and it leaks the funnel's internal encoding into the seed language.

## D105 — the UC lands with a DECLARED P6 boundary, not with a re-cut goal

**Context.** The curated final reads *"the member's AI-readiness flow is completed."* All three steps means a
real ~30-minute AI simulation and a live AI interview.

**Decision.** Land the UC scoped to the flow being **entered from day-0** and its **first step completed and
persisted**, with the next step observably unlocked — and say so, in the manifest note and in the corpus, as a
**P6 boundary** rather than quietly rewording the goal.

**Why.** The suite already has this shape and it is sanctioned: `pt-aisim-chat-launch` stops at `/start`,
`pt-skillpath-legacy` stopped at *"the completion control exists"*. What is NOT sanctioned is a final that
*sounds* like full completion while asserting one step. The two remaining steps are named where a reader will
find them, so this is a bounded proof, not a silent gap.

## D106 — ⚠️ RETRACTED BY MEASUREMENT. Superseded by D107. Kept because the mistake is the lesson.

**What D106 concluded (below): "the flake is a host stall; harden the liveness assert to 45 s."**
**What the next gate run measured: the SAME assert failed again, after the full 45 s.** The diagnosis was
wrong, the fix was reverted, and the real cause — mine — is D107.

**Why the wrong conclusion was reachable, stated so it is not repeated.** The evidence assembled below is
real and it is not weak: a 6-run duration table in which every *passing* reading sits on the iter-25 baseline,
one 5–10× outlier per run, everything else in the same run normal, and an independently-cleared side-effect
hypothesis. It all pointed at the host. What it never did was **reproduce the failure on demand** — and a
diagnosis that has not been reproduced is a story, not a finding. The reproduction (drive the surface five
times and count) took under three minutes and gave the answer immediately. **The rule: before hardening a
flaky assertion, make it flake on command.** iter-13 set that standard for `pt-assignment-assign` and it held;
here it was skipped because the duration table felt like enough.

The second half of D106 — leaving `pt-skillpath-legacy` alone — stands. Its 35.7 s stall against an explicit
30 s bound remains unexplained and unreproduced, and is recorded as such rather than papered over.

---

### D106 (as originally written — RETRACTED, see above)

**Context.** The first 3× gate attempt returned `1 failed / 192 passed` twice, on **two different tests**:
the iter-14 cross-tenant control's succession **liveness** assert (19.1 s, default 15 s `expect` timeout) and
`pt-skillpath-legacy`'s read-back (35.7 s against its explicit 30 s bound). Run 3 was clean at 193 passed.

**Measurement, before any conclusion** — the same two tests' durations across 6 consecutive full-suite runs
(3 at iter-25, 3 at iter-26):

| | iter-25 ×3 | iter-26 ×3 |
|---|---|---|
| cross-tenant control | 6.7 / 6.9 / 6.2 s | **19.1 (FAIL)** / 7.1 / 6.8 s |
| `pt-skillpath-legacy` | 4.3 / 3.5 / 3.5 s | 3.6 / **35.7 (FAIL)** / 4.7 s |
| suite wall-clock | 1.7 / 1.4 / 1.3 m | 1.8 / 2.3 / 1.7 m |
| `pt-onboarding-aireadiness-guided` | — | 8.4 / 7.0 / 6.7 s |

**Every passing reading sits on the iter-25 baseline.** One test in one run blew up 5–10× while everything
else in that same run was normal. That is the signature of a **host stall**, not of a slowdown introduced by
this iter — and the new Playthrough is a well-behaved 6.7–8.4 s. Independently checked and cleared: the
Playthrough's terminal *"Go to the AI Simulation"* click creates **no** `jobsimulation.sessions` row (Ola's 11
sessions are all seeded and backdated), so it is not loading the stack for the tests after it.

**Decision.** Harden **one** assertion, and leave the other alone.

- **Hardened:** the cross-tenant control's succession liveness assert, from the implicit 15 s default to an
  explicit **45 s** (~6.5× the measured p50). It is the **first touch** of the slowest of four dashboards that
  one test walks, an O(members) live projection with nothing before it to warm it, and — the reason it matters
  — when it times out the RED reads *"succession failed to compute for the contrast tenant"*: a liveness FLOOR
  failing for a reason unrelated to what it guards, which is misattribution of the kind this very file exists
  to prevent. Verified still discriminating: mutant **L1** points it at a role Org C does not have and it goes
  RED after the full 45 s, then green again at 6.5 s.
- **NOT hardened:** `pt-skillpath-legacy`. Its bound is already **explicit and 30 s** — ~9× its baseline, a
  considered iter-06 choice. Pushing it to 60 s would tolerate a 17× stall, and at that point the number is
  hiding a stall rather than absorbing one. Recorded as an observed host transient with its numbers instead.

**The general rule, worth carrying:** raising a timeout is legitimate when the *passing* distribution is
measured and tight and the tail is an outlier; it is hiding when the distribution has moved. The duration table
above is what makes the difference decidable, and it is why the table is in this record rather than a sentence
saying "it was probably slow."

## D107 — a hero's ROLE is an assertion anchor, and the anchor must be single-occupant

**The finding.** The day-0 seat was first declared `role: Data Analyst`, chosen for one reason only — that role
was **proven to resolve** in the public taxonomy, because Org C's COMPLETED hero already held it. That made
`Data Analyst` a **two-member** role in Org C. Measured, on one seed, one session, one page object:

| Org C `Data Analyst` occupancy | succession key-role card present |
|---|---|
| **2** (Robin + the day-0 seat) | **4 of 5** page loads |
| **1** (Robin only) | **5 of 5** page loads |

And at the suite level: **2 of 6** full-suite runs went RED — both times on `negative-controls.spec.ts`'s
cross-tenant control, and both times on its **LIVENESS floor**, i.e. reading as *"succession failed to compute
for the contrast tenant."* A control failing for a reason that has nothing to do with what it controls for.

**The fix.** Re-role the seat to `Supply Chain Analyst` — 10 `job_role_skills` (so `assertHeroRolesResolve`
stays green), held by **no other Org C member**, and logistics-coherent for Vertex Logistics.

**The fence.** `seed-facts-fence.unit.spec.ts` now requires **hero roles to be pairwise distinct within a
story**, with a self-test that injects a collision so the check is discriminating rather than trivially true.
Mutant **N1** puts `Data Analyst` back on the seat in the real seed and the fence goes RED naming
`Vertex Logistics: "Data Analyst" held by pt-ai-completed + pt-ai-onboard`. Two call-site comments carry the
same warning — one at the seat in the seed, one at the anchor in the control — so a future author meets it from
either direction.

**Two lessons worth keeping, both of which generalise past this iter:**

1. **A seeded hero's attributes are part of the test SUITE's contract, not just the demo's believability.**
   Four negative controls anchor on this world's seeded facts by NAME (org domain, org size, hero name, hero
   ROLE) — that is what iter-13/14 built, deliberately, to replace structural finals. The cost of that
   sharpness is that adding a hero can now perturb another Playthrough's anchor, silently and
   probabilistically. Picking a role because it "is known to resolve" copies the nearest hero's role, which is
   precisely the collision.
2. **The platform surface itself is non-deterministic here, and that is a real if minor product observation:**
   a succession key-role card's presence varies between page loads once the role has two occupants — most
   plausibly a top-N ranking with an unstable tiebreak. Not this milestone's to fix (zero platform edits), and
   it is not a demo defect the presenter would notice. Recorded so the next author who sees a key-role
   assertion flake looks at role occupancy before the clock.
