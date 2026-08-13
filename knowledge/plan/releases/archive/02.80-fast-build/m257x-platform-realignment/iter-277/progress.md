# iter-277 — read gate clause 5, the last open clause

**Type:** tik

## Phase A — the mechanical census (the `TOK-08` instrument)

Ran the full `stack-core` suite — 51 guards, 93 test modules — with the **exit code captured to a file**
rather than inferred, because a detached run that loses its exit code is not a result (run 38's standard).

**First run: `rc=1` — 12 failed, 2211 passed, 1 skipped, 593 subtests passed, in 53 min 13 s (CONTENDED).**

That number is **wrong**, and finding out why is this iter's main deliverable.

### The census was contaminated by the act of running it

`stack-core` has no checked-in Python environment, so pytest was installed into a virtualenv — and the
obvious place, `stack-core/.venv-check/`, is **inside the tree the census scans**. Attributed by file, the
breach is entirely `site-packages`: pip's vendored `urllib3`, `pygments`, `rich`, `tomli`.

Measured three ways rather than argued (same counter function, three trees):

| tree | `COMMENT_LITERAL` (ceiling 236) | `DOCSTRING_LITERAL` (ceiling 240) |
|---|---:|---:|
| rext @ `2833a64` (pre-iter-276) | **236** | **240** |
| rext @ `0a8674e` (**iter-276, tracked files only**) | **236** | **240** |
| the working tree **with the venv present** | **279** | **282** |

**Both ratchets sat EXACTLY at their ceilings, and iter-276's code comments added exactly ZERO.** The
+43 / +42 came from the measurement apparatus. The hypothesis this iter opened with — that the verbose
new comments in `jobroleref.go` had breached a literal ceiling — was plausible, specific, and **false**;
it was tested by extracting each commit with `git archive` (tracked files only, no untracked artifacts)
and re-running the identical counter, not by reading the diff.

**Re-run after relocating the venv outside the tree: 6 of the 12 failures vanished.** All six were the
frozen-expectation ceiling/ratchet/registry assertions — the `.py` census counting an installed library
as though it were our source.

## Phase B — what the census actually found

**12 failures → 1 real one.** The disposition, in full:

| # | failure | verdict |
|---:|---|---|
| 6 | `test_frozen_expectation_census_m257x` (ceilings, ratchets, derivation registry, noun coverage) | **artifact** — the venv inside the scanned tree |
| 4 | the `SUBFAILED` ceiling subtests | **artifact** — same cause |
| 1 | `test_route_disposition_guard::test_the_live_registry_is_consistent` | **REAL, and it was iter-276's** — fixed this iter |
| 1 | `test_fence_provenance::test_the_escape_accepts_and_records` | **REAL** — surfaces `clone_drift_guard`, routed forward |

### The real one that was ours, and the repair that re-created it

iter-276's close wrote the closed handler in an **elided** form. `route_disposition_guard` parses the
ellipsis as a truncated stem and says exactly what the consequence is:

```
RED m257x-platform-realignment: 'FIX-M257x-275-' is not a route id — carried at iter-276
EXIT 1 — a carry-forward that names a SET must ENUMERATE it (§5 rule 73): a glob leaves a
truncated stem behind … and both read as live backlog in every brief that quotes this queue.
```

A landed fix would have appeared as open backlog in every brief quoting the queue.

**The first repair re-tripped the guard.** Writing the id in full but *quoting the elided form in the
footnote explaining the rule* put the same stem back on the same line — the defect reproduced inside its
own explanation. The note now describes the elision without spelling it, and says so, because the next
reader's instinct will be the same one. `route_disposition_guard: OK · 0 malformed · 0 contradictions`.

### The real one that remains, and it is this run's own doing

`clone_drift_guard` is RED for a reason iter-276 created:

```
[D1 advanced] rosetta-extensions is at 0a8674e74, which the corpus never cites —
              22 commit(s) past the nearest of 12 cited sha(s); 44 citing site(s)
```

The corpus cites `rosetta-extensions` at 12 shas across **44 sites**; shipping the occupancy bound moved
the tooling 22 commits past the nearest of them. **This is the guard working**, and it is squarely
clause-5 material — *"and the corpus reflects that"* is the user's own third limb. It is a 44-site sweep,
i.e. a full iter, and is routed rather than half-done at the end of a session.

(The same run also reports three `UNMEASURED` pins where the corpus cites `app@1e457fa70` / `app@5ba17044`
against a clone now at `3eaadae68` — consistent with `CLAUDE.md`'s own note that `ad9f3c49` is 28 commits
behind current origin/main. Advisory, not RED, and folded into the same routed sweep.)

## Phase C — clause 5, graded honestly

**Clause 5 is NOT met, and this iter does not produce a `P`.**

- The **mechanical** half is one finding from clean, and that finding is a citation sweep this run
  created — not a pre-existing semantic pool.
- The **semantic** half was **not measured**. A `P`/`N` reading is a run-scale activity (iter-131 used 14
  seats); a scoped reading is evidence about its scope alone, and this milestone has been burned by
  publishing one as though it generalised. **No number is offered rather than a weak one.**

The two named clause-5 defects iter-131 routed forward were censused (a complete check over a named
population, not a sample) and **both are repaired**: the `infrastructure` hedge now survives only as
retraction prose recording the correction, and `architecture_overview.md` §4 lists the correct five
modules with `ai` explicitly excluded.

**Gate position: clause 5 remains the only open clause.** Clause 1 met, clause 2 met at iter-276,
clauses 3 and 4 hold.

## Close — 2026-08-11

**Outcome:** The census ran clean and **the reading is 12 failures → 1 real**: 10 were artifacts of
installing pytest inside the tree the census scans, 1 was iter-276's elided route id (fixed here, twice
— the first repair re-created the defect inside its own footnote), and 1 is `clone_drift_guard` going RED
because iter-276 advanced `rosetta-extensions` 22 commits past the nearest of 12 shas the corpus cites
across 44 sites. **Clause 5 is NOT met and no `P` is claimed** — the semantic half was not measured, and a
scoped reading would not have generalised.
**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (2 tiks) — (6) protocol-stop: n — (7) budget-exhausted: **y** — Outcome: **exit-7**

**Decisions:** `D-M257x-277-1` (the census was contaminated by the act of running it — measured, not
argued), `D-M257x-277-2` (the elided route id, and the repair that re-created it), `D-M257x-277-3` (no
`P` is claimed rather than a scoped one).

**Side-deliverables:** none.

**Routes carried forward:**
- **`FIX-M257x-277-corpus-cites-a-rext-sha-that-no-longer-exists`** — **the next iter's highest-value
  target.** 44 citing sites across 12 shas; `rosetta-extensions` is 22 commits past the nearest. Created
  by iter-276 and owed by this milestone under the user's third limb.
- **`ROUTE-M257x-277-the-census-cannot-be-run-from-inside-its-own-tree`** — `stack-core` ships no
  environment, and the natural place to make one breaches two ratchets that sit exactly at their
  ceilings. Either the counters exclude virtualenvs or the runbook states the venv must live outside the
  tree. Until then **every future census run reproduces this**, and it costs 53 min to discover.
- **Clause 5's semantic reading is still unmeasured** (last: iter-131, `P = 29 / N = 47`, a floor).
- Unchanged and not absorbed: `ROUTE-M257x-274-successor-half-is-uncovered`,
  `ROUTE-M257x-274-tie-order-is-unstable`, `FIX-M257x-269`,
  `ROUTE-M257x-270-directus-consumer-cms-key`, `FIX-M257x-266`, `FIX-M257x-265`, `ROUTE-M257x-h59`,
  `ROUTE-M257x-h65`.

**Lessons:**
1. **An instrument installed inside its own subject measures itself.** Ten of twelve findings were the
   measurement apparatus. The tell was that both breached ratchets sat *exactly* at their ceilings
   beforehand — a suspiciously round coincidence that turned out to mean the ratchets were fine and
   something new had been added underneath them.
2. **Quoting a defect can commit it.** The repair for an elided route id re-tripped the same guard by
   spelling the elided form inside the footnote explaining the rule. Guards read prose, including prose
   *about* guards.
3. **Shipping tooling puts the corpus out of date by construction.** Advancing `rosetta-extensions` to
   fix a Playthrough moved it past every sha the corpus cites. The tooling half and the documentation
   half of this milestone are coupled: a green clause 2 mechanically costs clause 5 something, and that
   debt should be paid in the same iter that creates it, not discovered by a fence 50 minutes later.
