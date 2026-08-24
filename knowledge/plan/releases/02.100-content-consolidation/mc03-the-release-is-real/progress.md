# MC03 — Progress

**Status: `planned`.** Not started. No iters run.

> **This is an ITERATIVE milestone — there is no checklist.** Work is recorded as iter closeouts in
> § Running ledger, and the seven-clause ledger below is re-graded after each one.
>
> **A RED clause closes an iter successfully.** MC03 exists to find failures the delivery milestones' own
> green gates could not see; a checkpoint that never fails is not detecting anything. What a RED clause
> does **not** license is repairing the cluster in place — findings **route to the milestone that owns
> them** (`overview.md` § Routing, not absorbing), and a finding whose owner is already closed fires the
> **re-scope trigger** and goes to the user with the three honest options.

## Exit-gate ledger

Re-graded at every iter closeout. A clause is **GREEN with evidence cited** or **RED with the finding
named** — "looked fine" is not a grade, and a clause graded from stdout rather than a machine-readable
artifact is **not graded**.

| # | clause | status | evidence |
|---|---|---|---|
| 1 | Cold `/demo-down --purge` + `/demo-up` on the official demo host → READY, autoverify GREEN, **0 patches refused/skipped** (denominator pinned) | _not started_ | — |
| 2 | `/dev-up` on a local stack → READY with the same properties | _not started_ | — |
| 3 | Playthrough batch gate **GREEN on both**, verdict contains no `skipped` | _not started_ | — |
| 4 | **Each** annotation request demonstrated **individually** against the running stack, evidence per request | _not started_ | — |
| 5 | Every v2.10 `Delivers →` doc **exists** and describes **shipped behaviour**, incl. NET-NEW `voice-feasibility.md` | _not started_ | — |
| 6 | The M271 voice **GO/NO-GO verdict recorded in the corpus with its reasoning** (+ `safety.md` amended if it touches the shared-AWS exposure) | _not started_ | — |
| 7 | **No milestone closed on a claim this checkpoint cannot reproduce** — any that did is named + re-proven or carried honestly | _not started_ | — |

**Gate = all seven.** A single RED clause blocks the release close until it is either turned GREEN or
routed under the re-scope trigger with a user decision recorded.

## Running ledger

_Iter closeouts append here, newest last. One entry per iter:_

> **iter-NN — `<clause(s) driven>` — `<host>` / `<stack N>` / rext `<tag>` / `<date>`**
> **Did:** …  **Measured:** …  **Verdict:** GREEN | RED  **Routed to:** …  **Next:** …

_No iters run._

## Findings routed out of this milestone

_Every finding whose fix belongs elsewhere is listed here with its destination — this is the table
`/developer-kit:close-release` reads. A destination is a **milestone**, never "the next release" or "a
future pass"._

| # | finding | owning milestone | owner status | routing (re-open / carry / drop) | user decision |
|---|---|---|---|---|---|
| _(none yet)_ | | | | | |

## Iters

_None._
