---
iter: 227
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-09
---

# iter-227 — does each archived repo agree with what this corpus says about it?

**Type:** tik, under `TOK-08` (*a reading SAMPLES; a fence CENSUSES*), on the **corpus half** of the user
redirect of 2026-08-09.

## Step 0 — re-survey before targeting (mandatory)

Re-surveyed against the open route queue. The **working-stack** half is currently blocked: gate clause 1
needs `ROUTE-M257x-225-no-profile-for-sanctioned-host`, whose prerequisite is a `buildbench` measurement
on a **quiet** box — and this host is running an agent, which is precisely the condition that made
`laptop.json`'s own cycle attempt get refused at `peak load1 10.69`. Sizing it was iter-225's deliverable;
running it is not this iter's, and pretending otherwise would be the fabrication the brief forbids.

So this iter takes the **corpus half**, and it generalises iter-224's sharpest finding, which was found
**once** and never asked as a question:

> **A resolved anchor quoting a verbatim line is not the source's position.** `messenger` retracted its
> "rollback path" claim in `CLAUDE.md` and left the same claim standing in `terraform/main.tf`, at one
> ref. The corpus quoted the stale half.

That was one repo, found by accident while checking something else. **Nothing has asked whether it is
true of the others** — and the six archived repos are exactly where this corpus makes its most
load-bearing status claims.

## Cluster / target identified

The **six archived / merged repos**, all now at their origin tips as of iter-224: `cms`,
`graphql-wundergraph`, `jobsimulation`, `messenger`, `roadrunner`, `storage`.

Each ships its **own** `CLAUDE.md` / `README.md` — the repo's *stated position on its own status*. The
corpus carries a row for each in `platform-migration-status.md` plus a service doc. **The two have never
been compared.** Where they disagree, the corpus is either ahead of the repo (fine, and it has been
before — iter-224 showed the corpus knew refs its clone set did not) or carrying something the repo has
retracted (the messenger defect).

## Hypothesis

At least one more repo states a status in its own `CLAUDE.md` that the corpus's row for it does not
reflect — most likely among the four that advanced in the last four days, since two of those advances
were *explicitly* doc-correction commits.

## Expected lift

A **6-row census** with an explicit verdict per repo — `agrees` / `corpus-ahead` / **`corpus-carries-a-
retracted-claim`** — every disagreement repaired, and the comparison written down so it can be re-run.

## Phase plan

1. **Seal predictions** (this commit — `probe(M257x/227)`).
2. Read each repo's own status statement at its origin tip.
3. Compare to the corpus's row + service doc; classify all six.
4. Repair disagreements; re-run fences.

## Escalation conditions

- If a repo's `CLAUDE.md` contradicts **the platform's `repos.yml`** rather than the corpus, that is a
  platform-internal inconsistency and gets recorded, not repaired — this corpus does not edit the platform.

## Acceptable close-no-lift outcomes

**Finding all six in agreement is a first-class result** — it would bound the messenger defect to a single
site instead of leaving it an open worry, and that is worth knowing precisely.

## Pre-registered predictions — SEALED IN THIS COMMIT

| id | prediction | rationale |
|---|---|---|
| **P-227-1** | **≥ 4 of 6** ship a `CLAUDE.md` carrying an explicit status / freeze statement | two shipped *"freeze the repo"* commits on 2026-08-05 |
| **P-227-2** | **≥ 1** repo beyond `messenger` states a status the corpus's row does not reflect | the messenger defect had no reason to be unique |
| **P-227-3** | `roadrunner` and `graphql-wundergraph` — the two that did **not** advance — carry **no** freeze commit, so their archived status is asserted by the corpus and **not** by the repo | they were already at their tips, and neither appears in the 2026-08-05 freeze wave |
| **P-227-4** | **0** disagreements will be cases of the corpus being *wrong about the platform's direction* — where they differ, the corpus will be **stale on detail**, not **wrong on direction** | every prior instance this milestone found was a detail rot |

**If P-227-2 is refuted (0 further disagreements), the messenger defect is bounded to one site** and the
iter reports that as the finding, not as an absence of work.
