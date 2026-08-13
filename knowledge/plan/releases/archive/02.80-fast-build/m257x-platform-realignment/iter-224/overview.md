---
iter: 224
milestone: M257x
iteration_type: tik
status: archived
opened: 2026-08-09
---

# iter-224 — the six clones nobody has fetched

**Type:** tik, under `TOK-08` (census the mechanical classes; stop sampling them), steered by the
**user's 2026-08-09 redirect**: the target is *(a) the corpus's claims about the platform* and *(b) an
actual working stack* — not the instruments that grade them.

## Step 0 — re-survey before targeting (mandatory)

`TOK-08`'s standing next-target language is "work the mechanical classes in descending measured size."
The user redirect narrows *which* mechanical class. iter-222 discovered the live one and iter-223
narrowed its consequence:

- **iter-222** — *nobody had fetched.* A plain `git fetch` against `platform` showed **15 corpus anchors
  had rotted**; the corpus was being graded against a stale local copy. Repaired **for `platform`**.
- **iter-223** — proved the 23 source patches still apply, so advancing the pin is safe. Left
  `ROUTE-M257x-222-other-clones-never-fetched` **open and unchanged**.

That route is this iter's target, and the re-survey confirms it is still untouched and still meaningful.
Measured this iter, before any fetch (`stat` on each clone's `.git/FETCH_HEAD`):

| clone | last fetch |
|---|---|
| ant-academy, app, next-web-app, platform, sentinel, studio-desk | **2026-08-09 19:41** (iter-222/223) |
| rosetta-extensions | 2026-08-06 11:19 |
| **cms, graphql-wundergraph, jobsimulation, messenger, roadrunner, storage** | **2026-08-05 23:24 — four days stale** |

The six stale ones are exactly the **archived / merged** repos — and they are the repos the corpus's
whole M810-teardown story is cited into (`cms/terraform/main.tf:39`,
`jobsimulation/terraform/main.tf:15-22`, `messenger/terraform/main.tf:29`, `storage/terraform/main.tf`).
**`clone_drift_guard` grades a corpus citation against the clone's LOCAL HEAD.** A clone that has not
been fetched cannot report drift — the guard returns green against a copy of the world from four days
ago. This is the iter-222 defect, in the half iter-222 did not reach.

## Cluster / target identified

The **150 citation occurrences** the corpus makes into those six repos — **123 of them carrying an
explicit `:NN` / `:NN-MM` line pin** — across **26 corpus files**. Population, measured pre-fetch:

| repo | occurrences | with `:NN` |
|---|---|---|
| cms | 52 | 46 |
| jobsimulation | 44 | 33 |
| messenger | 17 | 15 |
| storage | 17 | 11 |
| roadrunner | 11 | 10 |
| graphql-wundergraph | 9 | 8 |
| **total** | **150** | **123** |

## Hypothesis

Fetching the six will show at least one has moved past what the corpus knows, and the citations into the
moved repos will have rotted the same way `platform`'s did at iter-222 — silently, with every guard green.

## Expected lift

A **census, not a sample**: every one of the 123 pinned citations resolved at the repo's *origin* tip,
each rot named and repaired, and the blind spot (guard grades local HEAD, nobody fetches) fenced or
recorded. No `N`/`P` reading is claimed — this iter takes no graded reading.

## Phase plan

1. **Seal predictions** (this commit — `probe(M257x/224)`), before a single `git fetch`.
2. Fetch all six; measure `HEAD..origin/<default>` per repo.
3. Resolve all 123 pinned citations at each repo's origin tip; classify rot.
4. Repair what rotted; re-run the guard family.

## Escalation conditions

- If a fetch fails for auth/network reasons → record as UNMEASURED per `§8`'s three-verdict rule, do not
  report the repo as green.
- If a repo has advanced so far that the corpus's *claims* (not just its line numbers) are refuted →
  that is a finding for the migration-status map, and it lands here rather than being routed.

## Acceptable close-no-lift outcomes

**All six frozen and all 123 citations resolving is a first-class result** — it converts "nobody has
looked" into a measured statement, and it is the outcome the sealed predictions expect.

## Pre-registered predictions — SEALED IN THIS COMMIT, BEFORE ANY FETCH

Stated as falsifiable numbers before the first `git fetch`, per the milestone's standing method
(*derived figures are wrong at ~1 in 3 when written; derive or omit, never carry*).

| id | prediction | rationale |
|---|---|---|
| **P-224-1** | **≤ 2 of the 6** clones have advanced on origin past our local HEAD | the corpus asserts all six are frozen legacy since `838d907` / `2adcf71` |
| **P-224-2** | **≤ 8 of the 123** pinned citations fail to resolve at their repo's origin tip | follows from P-224-1 |
| **P-224-3** | **0 of 6** fetches fail for auth/network reasons | the same credentials fetched six other clones at 19:41 today |
| **P-224-4** | **`clone_drift_guard` reports GREEN both before and after the fetch** for at least one repo that HAS advanced — i.e. the guard is structurally blind to an unfetched clone, not merely quiet | it reads local `HEAD`; nothing in it fetches |

**If P-224-1 is exceeded (≥ 3 advanced), the corpus's "frozen legacy" framing for these repos is itself
at risk** and that, not the line numbers, becomes the finding.
