# M264 — Progress

**Status: COMPLETE.** 2026-08-15.

- [x] the count-claims across the corpus + CLAUDE.md
- [x] the `taxonomy`-folded-into-`app` correction (five org modules → **zero**)
- [x] re-grade the four taxonomy repos' "dormant, DECIDE" verdicts
- [x] `repair_postcondition` green

## The realignment was NOT "replace 42,790 with 3,562"

That would have been the obvious edit and it would have been wrong. `≥42,790 / ≥22,470` describes
**production as this corpus last measured it**; `3,562 / 706` describes **the governed canon as an
artifact in the repo**. They are different subjects, and merging them is the easiest mistake
available here. The canonical section now carries a three-row table naming each figure with what it
describes, and says plainly that **we could not verify whether production has adopted the canon** —
the platform's plan says not, but a plan is a document, and neither `db-query` path was available.

## What actually changed

**"60K skills" → REFUTED.** It was graded UNVERIFIED on correct reasoning (a public-only capture
cannot see private rows, so 42,790 is a floor). The platform then published its own pre-consolidation
total — **43,584 / 22,511** — which contradicts 60K outright.

**And the floor language was vindicated.** Insisting on *"≥, never ="* across three milestones made a
prediction: there is a private remainder. Measured against the platform's totals it is **794 skills /
41 job roles** — exactly the shape predicted. That is the payoff for the pedantry, and it is recorded
where a reader meets the figures.

**`app` requires ZERO org-private modules.** Both `shared_libraries.md` and `CLAUDE.md` said **five**.
At `4bccda085`, `grep anthropos-work/ app/go.mod` returns one line — the `module` declaration. The
block was **7** at `b948604f`, **5** at `3eaadae6`, **0** at `4bccda085`: three answers inside one
release cycle, so both sites now tell the reader to re-measure before citing. Not a documentation
detail — this is what broke Clerkenstein's disarm and made every demo login 401 (M263).

**The four taxonomy repos keep their grading, now evidenced.** They were graded dormant on commit
dates; a taxonomy revamp looked like it should have woken them. It did not: the canon is checked into
`app`, `skills-and-job-roles` is a *different lineage* (12,201 / 1,893, ESCO, 2024-04-08) rather than
an older version of the 43,584-row catalogue, and ESCO survives as **provenance** on the canon.

## One thing the fence caught

The first attempt at this commit was **blocked**: the re-grading introduced a second `454 of 706`
claim and the ratchet requires every `N of M` on the published surface to carry a written
disposition. Registered as the org-register twin. The block was correct — that is the fence working,
not fighting.
