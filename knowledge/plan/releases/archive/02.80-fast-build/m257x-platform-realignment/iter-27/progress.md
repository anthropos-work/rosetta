---
milestone: M257x
iter: 27
---

**Type:** tik (under `TOK-01: instrument first, then follow`)

# iter-27 — the cluster split, and the one seeder that never asked

## What was done

**Phase 1 — the per-id split.** Read the failing set from iter-26's own full-run artifact rather than from
the hand-off, and corrected the hand-off's enumeration on two points before any code was written (the
drilldown's failure text had been attributed to `assign-and-track.UC2`; `pt-assignment-assign` is an
affordance **count**, 15 vs 14, and is a singleton, not a member of the class).

Then measured **per id**, which refutes the hand-off's "four share ONE coherent signature":

| id | mechanism | evidence |
|---|---|---|
| `pt-workforce-org-feedback` | **seed-side absence** | her `feedback` draw is 0.8305 vs a 0.45 share |
| `pt-workforce-succession` | **read-side** | her interview row **exists** (`957d5253-…`) |
| `pt-workforce-funnel` | **not a hero absence** | her card's visible assert **passed**; only the role text inside it is missing, and the role is in the DB on 3 axes |
| `pt-activity-drilldown` | conditional on the drill target | needs a session on the sim the drill picks |

**Phase 2 — the fix.** The hero is population slot 1: `deterministicUUID("demo-1:story:pt-org-a:user:1")`
equals her live user id **exactly**. Her feedback share draw is fixed (the key prefix is pinned by
`stack: demo-1` *in the seed YAML*), so she was excluded permanently. Independently reproduced in Python: the
predicted in-share slots `[7,12,24,27,28,34,36,38,40]` are **byte-for-byte** the 9 rows the org held — the
match is what confirms the model rather than the guess.

`feedback.go` was the **only** one of six share-gated seeders that never resolved `personaIndexMapForStory`
— five siblings each make a deliberate, documented hero decision. Fixed with the family's own
`!isHero && !memberInShare(…)` guard plus the org-less guard the siblings already carry.

**Phase 3 — the fence.** A *classification*, not a rule (`D-M257x-27-1`): "heroes are always included" would
be **false** — `population_evidence.go` excludes them on purpose. Declare the policy per seeder, derive the
scope from the AST, check both directions. Reports what it checked (*53 sources scanned, 6 share-gated,
hero-always enforced in 2*) and **fails closed** if it scans nothing.

**Phase 4 — live proof** on `demo-1`, rext consumed at `fast-build-m257x-iter-27` (`b718149`) from origin:

| | before | after |
|---|---|---|
| hero feedback rows | **0** | **1** |
| org feedback rows | 9 | **11** (both heroes) |
| total feedback rows | 62 | 66 |
| audited rows | 55 729 (iter-26) | 55 733 |

…and at the **surface**, not only in the table (§5 rule 14): a scoped run reports
`[PASS] workforce-intelligence.organization-feedback.UC1`. Recorded as **advisory** — the harness itself
prints that every un-selected id correctly reads "did not run" and the ptreport gate binds only on a full
run. **No clause-2 number is claimed from it.**

## Close — 2026-08-01

**Outcome:** the `manager-reads-empty` cluster is split into ≥3 distinct mechanisms (2 of the 4 refuted as
"hero absent" outright); the one genuine seed-side absence is fixed at the family invariant it violated and
proven live at both the data layer and the surface. No clause-2 number claimed — a full run is its own iter.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (3 of 5 — clause 2 still wants 30/0/0; last binding number remains iter-26's 23/7/1)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n (platform origin `2adcf71`
re-checked at open and close, unchanged; occurrence stays 1 of 2) — (4) user-blocker: n — (5) cap-reached: n
— (6) protocol-stop: n — Outcome: continue
**Decisions:** D-M257x-27-1 (declare the hero policy, derive the scope) · D-M257x-27-2 (the cluster was not a
cluster) · D-M257x-27-3 (the inverted mutant) · D-M257x-27-4 (the suite was silent, not arguing)
**Side-deliverables:** none — the org-less guard shipped inside the planned fix because unconditional hero
inclusion requires it for correctness, not as an unrelated find.

**Routes carried forward:**

| item | why | target |
|---|---|---|
| `FIX-M257x-iter27-succession-hero-not-rendered` | Her interview row EXISTS (`957d5253-…`, FK'd to her real session). Read-side: name the query the Succession view runs and find what filters her out. Do **not** re-derive the seed side — it is measured. | next tik |
| `FIX-M257x-iter27-funnel-card-role-missing` | **Her card renders** (the visible assert passed); only the role text inside it is missing, while `user_basic_info.job_title`, the current `user_experiences` row and `job_role_id` all carry "DevOps Engineer" and 40/40 members have a title. DOM/locator-shaped — needs a live browser read of the card subtree, not a DB query. | next tik |
| `CHECK-M257x-iter27-drilldown-target-coupling` | The assert is conditional on which content `drillIntoActiveContent()` selects; the hero must hold a session on *that* sim. Establish whether the coupling is intended before treating it as a defect. | later tik |
| `CHECK-M257x-iter27-assignment-affordance-count` | The mis-filed singleton: `expect(received).toBe(expected)` **15 vs 14** — the count dropped by *two*, or the baseline was read after a mutation. Not a hero absence; do not batch it with the class. | later tik |
| `FIX-M257x-iter27-scoped-run-clobbers-binding-report` | A **scoped** diagnostic run overwrites `e2e/report/last-run.json` — iter-26's binding 209-spec artifact became a 1-spec one, and nothing in the file distinguishes a binding full run from an advisory scoped one. The measurement that grades the gate is destroyed by a run the harness itself calls non-binding. Same family as §5 rule 12 (*say which invocation produced the number*), one layer down: **say which invocation produced the FILE**. | next tik |
| `MEASURE-M257x-iter28-clause2` | The full `--reset` run. Budget it as an **entire iteration** (serial, `workers:1`; iter-25's mistake). Expect `pt-workforce-org-feedback` to have flipped — **expect, not claim**. | own iter |

**Lessons:**

- **A shared symptom is not a shared cause, and the cheapest way to prove it is to read the assert that
  came BEFORE the one that failed.** `pt-workforce-funnel` was in the cluster for three iters on the
  strength of its failure text; its *preceding* assertion says the hero's card is on the page. One line of
  the report, never read, would have removed it from the class at iter-15.
- **A deterministic hash is a decision-maker nobody reviews.** Five seeders reviewed it; one didn't, and the
  one that didn't produced a permanently-absent protagonist that looked like a platform-drift bug for
  twelve iterations. Promoted into the fence rather than the prose.
- **Include the inverted mutant.** Promoted to `platform-alignment.md` §8 rule 5 (`D-M257x-27-3`), in the
  same commit as this iter, alongside the no-op control it depends on.
