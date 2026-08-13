---
iter: 132
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-07
---

# iter-132 — the hedge the read retired, swept where it was published

**Resumed, not opened.** A prior session committed `fix(M257x/132)` (`1012679`, one line in
`org-repos.md`) and stopped before creating this dir or closing. That commit is this iter's first
deliverable; the dir and the rest of the sweep are the resumption. No work is re-done.

## Active strategy reference

**No successor strategy is authorable.** `TOK-08`'s sealed refutation branch fired at iter-119 and
instructs *"STOP. Do NOT author a successor strategy… hand the milestone back for a scope decision
from the user."* The milestone is at the user for that decision; iters 121–131 have run under the
user's direct brief, and so does this one. There is no `TOK-09` and this iter does not create one.

## Step 0 — re-survey

`unreadable_repo_claim_guard` at open: **OK**, with its standing NOTE — *"9 site(s) still hedge about
`infrastructure` while 13 report having read it."* The routed target
`FIX-M257x-iter131-infrastructure-hedge-stale` is live and unabsorbed. **One correction to the route
as written:** it says *"11 sites **+ `CLAUDE.md`**"*. `CLAUDE.md` was corrected at iter-124 and
carries the settled reading at `:194-203` and `:259`; its two remaining `clone set` hits are
`stack-demo`'s clone set (`:123`, `:132`) and `customerio-sync`'s `go.mod` (`:293`), neither of which
is this predicate. **The route over-stated by one file, and the file it named is the one every agent
loads.** Recorded rather than silently dropped.

## Cluster / target identified

iter-131's **P1** — the milestone's largest measured cluster, 19 of 80 blockers over six seats, and
the one `adj-1` (the single independent adjudicator) re-formulated. The predicate is **the
conflation**, not either conjunct:

> `infrastructure` is in no clone set / has never been read, **therefore** the folded services'
> production disposition is UNMEASURABLE.

Both conjuncts about the clone set are TRUE. The **inference** is false, because iter-123 read the
repo transiently at `13c248e6` and the corpus cites that read 28 times.

## Width, measured before repairing (§5 rule 57)

Four independent searches — `not measurable|unmeasurable`, `clone set`+`infrastructure`, `assert
neither|report both|do not assert either way`, and the union of the first two — over `corpus/**`,
`CLAUDE.md`, `README.md`. Union: **22 candidate lines in 11 files**, triaged to:

| group | sites | verdict |
|---|---|---|
| **A** — cms's M810 prod state inferred UNMEASURABLE from the clone-set premise | **8** | **FALSE — repair** |
| **B** — the production RPC address hedged on the same premise | **7** | **stale premise — repair** |
| C — jobsimulation's **GitHub archive state** (org API, never a clone) | 3 | TRUE — leave |
| D — the router banners' *"Vercel runtime configuration, in no clone set"* | 3 | TRUE, different subject — leave |
| E — already-corrected sites, and quotations of the retired hedge | 5 | leave (`org-repos.md:104`, `cms.md:81`, `platform-migration-status.md:88`,`:158`, `architecture_overview.md:227`) |
| F — `backend.md:51` (`customerio-sync`), `staging-bringup.md:461` (`colony`) | 2 | different repos — leave |

## Hypothesis

Group A is repaired by **substitution of the settled verdict**. Group B's premise can be *measured*
rather than re-hedged: clone `infrastructure` at the cited sha and read the production terraform for
the address the corpus says it cannot see.

## Expected lift

No `N` reading is taken this iter, so **no `N` movement is claimed** (§9's UNMEASURED rule, guard-rail
1). The deliverable is the sweep. Success = every group-A and group-B site repaired against a ref,
guard family no worse than at open, and the `unreadable_repo_claim_guard` NOTE resolved or explained.

## Phase plan

1. Enumerate + triage the population (done above).
2. **Settle group B's premise at source** rather than re-wording it — clone at the cited sha, two
   independent searches.
3. Repair groups A and B, each site against a ref.
4. Re-run the guard family; re-run `unreadable_repo_claim_guard`.
5. Close.

## Escalation conditions

- If the clone contradicts the corpus's settled cms verdict → **stop and surface**: that would put
  iter-123's read and iter-124's correction in question, not just this sweep.
- If a repair needs a platform-repo edit → route forward; v2.8 permits none.

## Acceptable close-no-lift outcomes

- The clone is unobtainable → group B is repaired with the honest *"not in the standing clone set,
  read transiently for a different question"* form, and the measurement is routed forward. That is a
  complete iter, not a partial one.
