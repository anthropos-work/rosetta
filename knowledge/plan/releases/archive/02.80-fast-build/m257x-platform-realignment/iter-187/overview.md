---
iter: 187
milestone: M257x
iteration_type: tik
status: archived
opened: 2026-08-09
controlling_strategy: TOK-08
---

# iter-187 — the exclusion registry is at SECTION grain; the exclusion is at (section × language) grain

**Active strategy reference:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07)
— *census the mechanical classes; stop sampling them.* This iter censuses the (section × language)
matrix instead of re-reading the section list.

## Step 0 — re-survey before targeting

Run 17's brief named `SURVEY-M257x-iter185-other-declared-populations-unaudited` (70 enumerated
collections, population-vs-predicate classification) as the natural next target. **Re-survey substitutes**,
under the same strategy and for the reason iter-186 itself recorded: the residual of that route is
explicitly *judgement* (*"there is no syntactic marker, so the split is judgement"*), and TOK-08's whole
point is to stop paying for judgement where an enumeration is available.

The enumeration available here is one command: **the (section × language) test-file matrix**. iter-186
partitioned the repo's **11** sections into 5 collected / 6 excluded-by-language. That partition is at
**section** grain. The runner's actual reach is at **(section × language)** grain, and the two do not
coincide.

Measured 2026-08-09 over `.agentspace/rosetta-extensions` @ `4afbd6c`:

| section | `test_*.py` | `*_test.py` | `*_test.go` | `*.spec.ts` | bucket |
|---|---:|---:|---:|---:|---|
| alignment | 0 | 0 | 21 | 0 | excluded |
| clerkenstein | 0 | 0 | 37 | 0 | excluded |
| demo-stack | 35 | 0 | 0 | 0 | **collected** |
| dev-stack | 5 | 0 | 0 | 0 | **collected** |
| playthroughs | 0 | 0 | 22 | 45 | excluded |
| stack-core | 67 | 0 | 0 | 0 | **collected** |
| stack-injection | 7 | 0 | 0 | 0 | **collected** |
| stack-secrets | 0 | 0 | 20 | 0 | excluded |
| stack-seeding | 0 | 0 | 119 | 0 | excluded |
| stack-snapshot | 0 | 0 | 45 | 0 | excluded |
| **stack-verify** | 5 | 0 | 0 | **30** | **collected** |

**`stack-verify` is declared COLLECTED and carries 30 Playwright `*.spec.ts` that no instrument reads and
no registry names.** Every arm of `test_suite_census_population.py` is green over it, correctly: it is in
exactly one bucket, and it does carry Python tests. The arms check *membership*; the hole is *within* a
member.

## Cluster / target identified

Two mechanically-decidable defects, both in `stack-core/suite_census.py` + its iter-186 fence:

1. **D1 — a silent exclusion inside a section declared as read.** `LANGUAGE_EXCLUDED_SECTIONS` records
   only whole sections. `stack-verify`'s 30 specs are excluded in fact and named nowhere, so the printed
   scope line — `scope: 5 of 11 sections — Python only` — reads as *those five are fully read*. This is
   iter-186's own rule (*a CORRECT exclusion is still a defect while it is silent*) recurring one grain
   down, and it also corrects the repo-wide figure iter-186 published: the unread non-Python population is
   **264 Go + 75 TS**, not 264 + 45. iter-186's number was scoped to *the six excluded sections* and is
   true as written; what nothing states is the repo-wide total.
2. **D2 — the fence's Python-presence predicate is a SUPERSET of the collector's.** Harden pass 42 added
   `test_every_COLLECTED_section_actually_carries_PYTHON_tests` precisely so a silent zero cannot hide in
   a quoted total. Its predicate is `rglob("test_*.py") + rglob("*_test.py")`; the collector `modules()`
   globs **`test_*.py` only**. A section whose Python tests are all spelled `*_test.py` passes the arm and
   contributes a silent zero — the defect the arm exists to prevent, re-entering through the spelling
   (`§5` r70/71). **Hazard size, measured, not asserted: 0 files today** across all 11 sections — so this
   is latent, and it is one line to close by deriving the arm from the collector itself.

## Hypothesis

Making the registry (section × language)-grained, and deriving the presence arm from `modules()` rather
than from a hand-written superset glob, converts both a live silent exclusion and a latent spelling hole
into printed facts and RED-provable arms — with no widening of what the census actually runs.

## Expected lift

No `P`/`N` reading (`§9`: UNMEASURED, not unmoved). Instrument-quality lift: 1 live silent exclusion named
(30 specs), 1 latent superset predicate closed, 1 repo-wide figure stated for the first time, ≥2 new fence
arms each mutation-proven RED.

## Phase plan

- **A** — measure the (section × language) matrix (done above; re-derived in code as the fence's input).
- **B** — repair `suite_census.py`: derive the within-section unread population from disk, name it by
  reason, print it in the scope block ahead of any total.
- **C** — extend `test_suite_census_population.py`: both-direction arms over the within-section exclusion
  registry + re-base the presence arm on `S.modules()`.
- **D** — mutation-prove the new arms RED; run the touched modules under **both** runners (`§5` r75/76).
- **E** — publish: `platform-alignment.md` §8 rule; route residuals.

## Escalation conditions

- If the 30 specs turn out to be collected by some *other* instrument this milestone quotes, the finding
  is a duplicate-population problem, not a silent-exclusion one — re-scope to that instead and say so.
- If repairing the presence arm turns any existing section RED, that is a live defect, not a fence bug —
  land the finding, do not weaken the arm.

## Acceptable close-no-lift outcomes

- The within-section unread population is already named somewhere authoritative and the scope line simply
  does not print it → the finding shrinks to a printing defect; record the falsification and close.
