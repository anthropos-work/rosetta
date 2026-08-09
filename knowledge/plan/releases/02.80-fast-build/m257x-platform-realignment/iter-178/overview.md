---
iter: 178
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-09
controlling_strategy: TOK-08
---

# iter-178 — the `N of M` prose class, censused inside clause 5 and given a disposition each

**Active strategy reference:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07)
— *census the mechanical classes; stop sampling them.* This iter takes the class iter-177's new §8 rule
is **about** and does to it what that rule demands: enumerate every published `N of M` count on the
clause-5 surface, and give each one a disposition instead of a mood.

## Step 0 — re-survey before targeting (mandatory)

`derived_count_guard` prints, on **every** run including green ones:

> `NOT REACHED: the N of M prose shape. M names no source, so attributing it to a table is an inference.`

That NOT-REACHED clause is routed as `SURVEY-M257x-iter173-derived-count-guard-reach` and has been open
since iter-173. **It has never been measured** — nobody has said how large the unreached class is. Under
`TOK-08` that is the first thing to fix: *a reading SAMPLES; a fence CENSUSES*, and a class whose size is
unknown cannot be ranked against the others.

**Measured (read-only, at corpus `794b167`):**

| surface | `N of M` occurrences |
|---|---|
| `corpus/**` + `CLAUDE.md`, bolded `**N of M**` | 18 |
| `corpus/**` + `CLAUDE.md`, all forms | 61 |
| **`corpus/architecture/**` + `corpus/services/**` — the clause-5 surface** | **9 occurrences over 8 lines** |

Nine is small, and that is a finding in itself: the shape `derived_count_guard` declines to reach is
**not** a large hole on the surface the gate is graded over. It is nine sites, and nine sites can be
enumerated exactly.

## Cluster / target identified

The clause-5 population, enumerated in full — and **split by derivability first**, per §8 iter-173
(*a DERIVED number is censusable; an OBSERVED one is not — split the class before you scope the fence*):

| # | site | claim | class |
|---|---|---|---|
| 1 | `architecture/service_taxonomy.md:300` | `academy.graphqls`, **1 of 43** files | derivable |
| 2 | `services/ant-academy.md:77` | the same claim, restated | derivable (**twin of 1**) |
| 3 | `architecture/architecture_overview.md:398` | org auto-filter on **31 of 135** schemas | derivable |
| 4 | `architecture/shared_libraries.md:257` | `taxonomy` required by **6 of 7** on-disk Go repos (×2 forms) | derivable |
| 5 | `architecture/org-repos.md:24` | **6 of 13** clone trees behind their own origin | observed-at-a-past-instant |
| 6 | `services/customerio-sync.md:55` | identical for **7,392 of 7,397** rows | observed (prod DB) |
| 7 | `services/ai-readiness.md:446` | ≈**156 of 199** members completed | observed (a DB) |
| 8 | `services/ai-readiness.md:705` | **4 of 5** tiles rendered empty | historical (a fixed defect) |

## Hypothesis

Two parts, and the second is the deliverable:

1. **The derivable half is re-derived at the refs it cites.** If any is false it is a clause-5 defect
   and is repaired; if all are true, that is a census returning ZERO and its instrument must be proved
   (§9 iter-149).
2. **The class stops being "NOT REACHED" and becomes "enumerated and dispositioned."** Every `N of M`
   site on the clause-5 surface must carry a disposition — `DERIVABLE` (with the derivation named) or
   `OBSERVED:<why it cannot be re-derived here>` — and a site with none is RED. That is the same shape
   as the family runner's `reconcile()` and iter-176's registry fence: *a member omitted from a census
   reads exactly like a member that passed.*

The guard already exists and already owns the disclosure, so the arm goes **into
`derived_count_guard`** rather than into a new module. Adding a fence module would drag a README row, a
`derivation_registry` entry, a `guard_family` invocation and a battery seed entry behind it — four
registries, three of which this milestone has already caught rotting — for a check that belongs in the
guard whose own output declines it.

## Expected lift

**No `P`/`N` reading is taken this iter, so no clause-5 movement is claimed** (`§9` iter-type
refinement). Stated in counts: clause-5 `N of M` sites with a disposition **0 → 9**; derivable ones
re-derived at their cited ref **0 → 4**; NOT-REACHED clauses that name their own size **0 → 1**.

## Phase plan

* **A — census + split** (done in Step 0).
* **B — re-derive the derivable half** against the on-disk clone set, recording the **substrate**
  (`D-M257x-122-4`: before believing a defect, read the substrate line).
* **C — Arm D**: the disposition table + the both-directions reconcile in `derived_count_guard`, with a
  mutation control and an anti-vacuity control that can fire.
* **D — run**: the guard, its tests, and the `stack-core` section under the named runner.
* **E — route** whatever the census surfaces and this iter does not land.

## Escalation conditions

* A derivable site that re-derives **false** is a clause-5 defect: repair it in place, and say so in the
  close rather than folding it into the fence's headline.
* If Arm D's population turns out to need a clone to enumerate (rather than to *check*), stop — a fence
  that cannot run without a stack is not a fence for this surface.
* If adding the arm turns `derived_count_guard` RED anywhere else, that is a real finding; do not narrow
  the arm to make it quiet (iter-158).

## Acceptable close-no-lift outcomes

* All four derivable sites are true and the arm proves nothing new → a census returning ZERO, which is
  a first-class outcome **only if the instrument is proved to fire** (§9 iter-149). The mutation control
  is the proof; without it this iter closes `closed-no-lift`.

## Out of scope, routed not taken

* The other **52** `N of M` occurrences outside clause 5 (`corpus/ops/**`, `corpus/tools/**`,
  `CLAUDE.md`). `D-M257x-129-2` binds: work outside clause 5 is booked to the user's standing ask, never
  to the clause. The arm is scoped to the clause-5 surface and says so.
