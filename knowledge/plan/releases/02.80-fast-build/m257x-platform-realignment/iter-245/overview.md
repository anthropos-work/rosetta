---
iter: 245
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-10
---

# iter-245 — the same question on the PLATFORM slice: wrong, or merely uncheckable?

**Active strategy reference:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07).

## Step 0 — re-survey

iter-244 closed the rext slice of one class and opened
`ROUTE-M257x-244-unresolvable-and-wrong-share-one-bucket` naming the larger half: the corpus's citations
**into the platform repos**. Same mechanism, bigger population, and squarely the user's redirect — these
are the corpus's claims *about the platform*.

`anchor_construct_guard` reports **883 resolved / 599 unresolvable, exit 0**. Of the 599, **377 are bare
`:NN`** and are refused for stated, correct reasons. The rest are path-qualified heads — `main.go` ×33,
`infrastructure` ×18, `studioManager.go` ×9 — booked as *out of reach* and therefore indistinguishable
from citations that are simply **wrong**.

## Substrate, enumerated before sealing

732 path-qualified `` `path:NN` `` citations across the 114 live documents, classified:

| class | count |
|---|---|
| **repo-rooted, clone PRESENT** (the gradeable population) | **335** — 218 single-line + **117 range** |
| repo-rooted, clone ABSENT | **131** — `infrastructure` ×45 dominating |
| repo-relative (needs the citing doc's service to resolve) | 203 |
| rext or intra-corpus (iter-244 / `corpus_citation_guard` territory) | 63 |

Clone set: `stack-demo/` — 13 repos present (`app`, `cms`, `jobsimulation`, `messenger`, `roadrunner`,
`storage`, `sentinel`, `next-web-app`, `studio-desk`, `ant-academy`, `platform`, `graphql-wundergraph`,
`rosetta-extensions`).

## Hypothesis

**Existence and construct are separable, and only the second one is hard.** `anchor_construct_guard`
declines the 490 range citations because *which line of a range carries the claim* is undecided — but
**whether the file exists is not undecided for a range**. So a large slice of the standing
`ROUTE-M257x-h59-range-anchors-are-ungraded` population is gradeable *right now*, on a weaker question
that is still worth asking.

## Pre-registered numeric claims — SEALED IN THIS COMMIT

Graded at the clone's worktree **and** at `origin/main`; a finding requires the path to be missing at
**both**, so a citation that names a file living at some ref the corpus might mean is never a finding.

| id | claim | prediction |
|---|---|---|
| **P-245-1** | distinct file paths, among the 335 clone-present citations, missing at BOTH refs | **≤ 40** |
| **P-245-2** | ≥ 1 of the findings is a **RANGE** citation — i.e. this iter reaches part of the `h59` ungraded population on the existence question | **YES** |
| **P-245-3** | `infrastructure` is absent from the clone set, so its **45** citations must report could-not-check, never green | **YES** |
| **P-245-4** | `app` supplies the plurality of findings (it supplies 228 of 335 citations) | **YES** |
| **P-245-5** | ≥ 1 finding lands in a **frozen-legacy** repo (`cms`/`messenger`/`roadrunner`/`jobsimulation`/`storage`/`graphql-wundergraph`; 84 citations) — the population `ROUTE-M257x-241` flagged as graded by nothing | **YES** |

**Falsification:** if P-245-1 returns **0**, the platform slice is clean, the route closes with a measured
zero, and the deliverable is the separation itself (existence-gradeable vs construct-gradeable) rather
than a repair.

## Phase plan

A — build the classifier (worktree + `origin/main`, per clone). B — read it, with denominators.
C — repair or disposition every finding. D — fence, if the class is non-empty. E — re-derive last.

## Escalation

A citation missing at both refs but plausibly correct at an **older** ref is a **ref** question, not a
path question — disposition it, do not "repair" it into a different claim.
