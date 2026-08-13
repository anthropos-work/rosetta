---
milestone: M257x
iter: 30
---

**Type:** tik (under `TOK-01: instrument first, then follow`)

# iter-30 — one failure fixed at the DOM, one root-caused to the seed's own role diversity

## Line 1 (planned) — the instrument repair

`FIX-M257x-iter27-scoped-run-clobbers-binding-report` is closed. `run-playthroughs.sh` now writes a
provenance sidecar naming the invocation that produced the artifacts, and — the half that actually saves
a measurement — **only an UNSCOPED run may lay down `last-binding-run.json` / `last-binding-report.json`**,
so no number of scoped diagnostics can destroy the last binding verdict. The runner always *knew* which it
was (`SCOPED`, computed over the whole argv, is what the gate branches on); the knowledge simply never
reached the artifact.

Guarded by `TestRunnerSafety_RunProvenance`, which **executes** the extracted shell function against a temp
dir under both classifications rather than grepping for it. Mutant battery, with the declared-GREEN control
that makes the RED verdicts interpretable:

| mutant | expected | actual |
|---|---|---|
| M0 no-op (comment-only) | GREEN | GREEN ✓ |
| M1 removal (the call deleted) | RED | RED ✓ |
| M2 **inversion** (`-ne 1` → `-eq 1`: scoped runs preserve, full runs don't) | RED | RED ✓ |
| M3 **label-only "fix"** (provenance kept, PRESERVE block dropped) | RED | RED ✓ |

M3 is the one worth having: a fix that only *labelled* the artifact would satisfy every plausible
"does it say it's scoped" assertion while still destroying the file, which is the actual loss iter-29 was
hand-copying around.

## Line 2 (planned) — the live DOM read, and what it settled

Both hypotheses in the overview were partly wrong, and the two ids do **not** share a cause. Stated first
because it is the third cluster this milestone has dissolved on measurement.

### `pt-workforce-funnel` — FIXED, verified live

The role was never missing. `DEVOPS_IN_FUNNEL_PAGE=true`: the string is on the page. The accessor was
addressing the wrong node. Measured — **eight** divs match `^Pat Ellis` on the manager's dashboard:

    [0] "Pat Ellis / 25 sims / 3.9-5 / 4.3-5 / +0.5 / Product"      ← a FEEDBACK-rating card, no role
    [1] "Pat Ellis / 25 sims / 3.9-5 / 4.3-5 / +0.5"
    [2] "Pat Ellis / 25 sims"
    [3] "Pat Ellis / DevOps Engineer / 25-25 passed / 100% / Product / Morgan Reyes / …"   ← the cluster
    [4]  (idem)
    [5] "Pat Ellis / DevOps Engineer / 25-25 passed / 100% / Product"     ← HER learner card
    [6] "Pat Ellis / DevOps Engineer / 25-25 passed / 100%"
    [7] "Pat Ellis / DevOps Engineer"

`.first()` returned `[0]`. The dashboard was rendering the role correctly, seven nodes away, for however
many iterations this has been attributed to platform drift.

The repair discriminates on the learner cluster's own marker (`passed`) and takes `.last()`. The
discriminator is deliberately a **different string from the claim** the Playthrough then makes inside the
card (§5 rule 7 — a probe must not be able to satisfy itself; selecting the card *by* the role would have
made the role assertion vacuous).

**Verified live on `demo-1`:** `pt-workforce-funnel` **passes, 1.4 s**. The five negative controls all
pass, and one of them is the anti-vacuity control that matters — the contrast tenant's own hero is still
found through the same accessor (so it is not over-narrow) while Pat Ellis is still absent from that
tenant's dashboard (so it is not over-broad).

Mutation is written up in `decisions.md` D2 and generalised into the protocol doc §8 rule 5: the
single-clause mutants **survived**, and the discriminating control is the full revert.

### `pt-workforce-succession` — root-caused, NOT fixed, routed with the cause named

The inherited handler name (`…-succession-hero-not-rendered`) is wrong, and so was the hand-off's stated
mechanism. The failing assertion is `workforce-succession.spec.ts:83` —
`getByRole('heading', { name: /^DevOps Engineer$/ })`, a **key-role card**, not a hero-absence assert. The
run never reaches the hero assert at `:90`.

Measured on the live surface: the page renders **"Roles by risk"**, and it is a *capped, tie-broken* list.

- `Critical roles: 28` (risk ≥ 50); the page emits **25** role headings.
- Every rendered card reads the identical `MEDIUM / risk 68`, so all 28 are **tied**.
- The seeded org is **40 members across 39 distinct job roles** — one incumbent each, bar a single role
  with two. `DevOps Engineer` has exactly **1**.

So whether the hero's role appears is decided by a tiebreak among 28 equal-risk roles, of which the view
shows 25 — and the rendered set is dominated by `3D…` / `A…` names. `DEVOPS_IN_SUCCESSION_PAGE=false`.

**This is not platform drift and it is not a locator bug.** It is an assertion whose truth depends on an
uncontrolled tiebreak, which happened to be heads when iter-14 measured it. The org's role atomisation
(one role per member, from the generated batch) is what pushed her out. The fix is a **seed-shape**
question — give the proof role real incumbency so it ranks on its own merits — which is a third line of
work and lands squarely on the scope-creep tripwire. Routed with the cause measured, not with a hunch.

### Three ids deliberately NOT folded in

`pt-activity-drilldown` fails on `heroRow.count() > 0` — the hero's **name** absent from the per-member
breakdown, and it never reaches its own role assert. `pt-orgadmin-role-create` is a 60 s `waitForURL`
timeout. `pt-onboarding-hiring-candidate` is a missing `/sim/…organizationId=` link. Three singletons
until measured otherwise.

## Refutations of the inherited hand-off (three, all before any work started)

1. **`FIX-M257x-iter27-succession-hero-not-rendered` is misnamed** — see above.
2. **"The taxonomy has no exact `DevOps Engineer` role" is FALSE.** `public.job_roles` holds it exactly
   (among 32 `%devops%` rows). The escape hatch the hand-off offered — *the card may render a resolved
   taxonomy role that legitimately cannot match* — does not exist.
3. **The seed data is intact on every axis.** `memberships.job_role_id = J-DEVOPS-AECC` **and**
   `job_role_name = 'DevOps Engineer'` for the hero, 40/40 for the org. (Incidental: `job_role_title` is
   empty for **all 190** memberships platform-wide — nothing fills it. Noted, not chased.)

## Close — 2026-08-01

**Outcome:** `pt-workforce-funnel` fixed and verified green live (an accessor addressing a role-less
feedback card instead of the learner card); `pt-workforce-succession` root-caused to a capped tie-broken
role list over a 39-roles-for-40-members seed, and routed; the scoped-run artifact clobber closed with an
executed guard. The inherited two-id "shared signature" is refuted — the two failures share nothing.
**Type:** tik
**Status:** closed-fixed (both planned lines landed: the instrument repair, and the DOM read that named
both root causes — one of which is fixed)
**Gate:** NOT MET (3 of 5; clause 2 measured `25 / 5 / 1` at iter-29, one id now fixed but **not**
re-measured — see below)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n (platform origin `2adcf71`
re-checked at open and close; occurrence stays 1 of 2) — (4) user-blocker: n — (5) cap-reached: n —
(6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (provenance + preservation), D2 (the individually-sufficient-clauses mutation rule),
D3 (succession routed rather than fixed, under the tripwire).

**Honesty on the metric.** Clause 2 is **NOT claimed to be `26 / 4 / 1`.** The funnel green was measured by
a *scoped* run, which is advisory by the runner's own gate — and the whole point of line 1 is that a scoped
artifact must not be read as a binding verdict. A binding run is ~35–40 min serial and is its own
iteration. What is claimed: one previously-failing id passes on demand, with a full-revert mutant proving
attribution.

**Side-deliverables:** none.

**Routes carried forward:**

| item | why | target |
|---|---|---|
| `FIX-M257x-iter30-succession-role-tiebreak` (**supersedes** `FIX-M257x-iter27-succession-hero-not-rendered`, which named the wrong assert) | The hero's role is 1 of 28 tied `risk 68` roles in a list that renders 25. Fix is seed-shape: real incumbency for the proof role. Verify it does not perturb the 39-role aggregates other specs read. | next tik |
| `MEASURE-M257x-iter30-clause2-binding-run` | A full binding run to convert the funnel fix into a clause-2 number. Budget it as a whole iteration. | after the succession fix, so one run buys both |
| `CHECK-M257x-iter30-scoped-classifier-misses-filenames` | `note_scoping_flag` classifies a bare spec **filename** as NOT scoped (deliberately, per its own test), yet a filename narrows the suite exactly as `--grep` does — so `./run-playthroughs.sh 1 tests/x.spec.ts` would now write a *binding* snapshot from a one-spec run. Found while wiring line 1; not touched, because flipping it changes an existing tested contract. | later tik |
| `CHECK-M257x-iter28-assignment-flip-is-stateful` | Unchanged. Note the new candidate mechanism: iter-27 gave the hero feedback rows, which render a **new card** (`[0]` above) — a plausible non-coincidental source of an affordance-count change. Still to be measured, not inferred. | later tik |
| `CHECK-M257x-iter27-drilldown-target-coupling` | Unchanged; `pt-activity-drilldown` is a hero-name absence, a different claim from the role-text ids. | later tik |
| `DOC-M257x-iter30-job-role-title-unfilled` | `memberships.job_role_title` is empty for all 190 rows platform-wide. Harmless today; worth one line in the map if any surface ever reads it. | later tik |

**Lessons:**

- **Read the report per-id before planning anything.** Three inherited claims died in the first ten
  minutes — including a factual claim about the taxonomy that a single `SELECT` refuted, and a handler
  name that pointed at the wrong assertion. The binding artifact was on disk the whole time.
- **A surviving mutant is not automatically a refutation of the fix.** Check first whether the fix is a
  conjunction of individually-sufficient clauses; then the discriminating control is the **full revert**.
  Now protocol doc §8 rule 5.
- **"The data is missing" and "the accessor is wrong" look identical from a failing assertion.** Eight
  nodes matched; the seventh through eighth carried exactly what the test wanted. The DB query that
  "proved the data was there" had already been run by an earlier iter and was used to conclude the
  *platform* had drifted.
- **Fix the instrument before using it, in the same iter, and the ordering pays for itself.** Every
  diagnostic in line 2 was a scoped run — precisely the thing that used to overwrite the binding artifact.
