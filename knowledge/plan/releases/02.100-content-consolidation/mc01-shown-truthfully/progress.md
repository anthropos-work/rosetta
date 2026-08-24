# MC01 — progress

**Status: `planned`.** Not started. No iters run. **Blocked until M266 and M268 have both closed** — a clause
graded against work still in flight measures the flight, not the cluster.

> **This is an ITERATIVE milestone — there is no section checklist.** There is a clause ledger (what the gate
> says) and a running ledger (what each iter found). **A red clause is this milestone WORKING**: its output is
> a **classification and a routing**, not a fix applied here.

## Exit-gate clause ledger

Each clause is green only when its **recorded evidence** exists in [`spec-notes.md`](spec-notes.md).
"It looks right" is not evidence.

| # | clause | required evidence | status |
|---|---|---|---|
| 1 | cockpit org-stories: non-manager heroes **two-per-row**, manager card **full-row**, cold stack | stated viewport width + observed cards-per-row, in a browser | _not started_ |
| 2 | "Candidate Hiring & Comparison": non-manager heroes read **CANDIDATE**, pills read **PERFORMING / UNDER-PERFORMING** | badge text read on every non-manager card; zero EMPLOYEE / THRIVING / STRUGGLING | _not started_ |
| 3 | content-story cards: **no nested `.ctcol`**, language pills carry an **inline-SVG flag**, **Academy has no pass/fail chip** while verdict-carrying products still do | three observations on the served page, graded against M266's `has_verdict` decision | _not started_ |
| 4 | **whole** of `public.job_simulation_sessions`: every row satisfies `(score >= 60) == (completion == passed)` | query or fence test + **violating rows AND total row count** | _not started_ |
| 5 | `/enterprise/assignments-list` shows **3–5 programs** across **>= 3 distinct stages** | browser observation + the written reading rule for "stage" | _not started_ |
| 6 | `cockpit-spec.md`, `content-stories-spec.md`, `seeding-spec.md`, `stories-spec.md` each describe the **shipped** behaviour, read **AGAINST THE RUNNING STACK** | per-doc read with sections named; **drift is a gate failure, not a follow-up** | _not started_ |

## Stack under test

_None. Record stack id, flags (incl. whether `--public-host`), consumed `rosetta-extensions` tag (verified **on
origin**), platform refs, cold/warm, and timestamp in [`spec-notes.md`](spec-notes.md) § *The stack under test*
before grading anything._

## Routings issued

Findings sent **back** to the milestone that owns them, per the re-scope trigger.

| # | finding | clause | routed to | date |
|---|---|---|---|---|
| _(none yet)_ | | | | |

## Running ledger

_Iter closeouts append here — newest last. One entry per iter: what was measured, on which stack, what came
back, how each finding was classified, and where it was routed. **Record red clauses in full** — a checkpoint's
value is in what it caught, and an iter that found nothing should say so plainly rather than pad._

_No iters run._
