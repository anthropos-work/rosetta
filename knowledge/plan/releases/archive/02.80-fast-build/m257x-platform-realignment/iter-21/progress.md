**Type:** tik — clause 5 (reconcile `corpus/services/**` + `corpus/architecture/**` against the iter-20 map).

## What was measured

Blast radius, scoped to the gate's two trees: **15 files / ~78 router hits**, and **7 files** still asserting
a 3-subgraph supergraph. The inherited "35 files / ~128 hits" is a whole-corpus figure.

Endpoint truth, from the platform clone at origin HEAD: `docker-compose.yml:334,352` bake
`:8082/graphql/query` into studio-desk and next-web. **`:5050` is dead locally.** rext had already handled
this at iter-13 (`test_frontend_build.py:915` asserts the port cannot survive) — **only the corpus lagged.**

## What landed — 4 commits

1. `a5126bc` — the sweep. 40 enumerated edits across 17 files, each asserted to match **exactly once**.
2. `36f7fa5` — the merge of `main` (see below), 18 hunks resolved as unions.
3. `c33e207` — the residue main's sweep did not reach.
4. `1531da1` + `3ebc6b7` — audit runs 2 and 3's findings.

The correction shape is deliberate and is the milestone's own lesson: **not "delete every router mention".**
The router has two states. In prod it is declared (`terraform:20 = 1`); locally it does not exist. So: a
two-state banner at the top of the four affected hub docs, every **live local instruction** re-pointed off
`:5050`, and every historical row **struck through rather than deleted** — so a reader who greps the dead
port lands on the explanation instead of on nothing.

## The branch was three commits behind `main`, and nothing measured it

The first audit read RED/11, eight of them *"cms.md and jobsimulation.md never mention the merge"*. **They do
— on `main`.** This branch was cut at `bf3f9bc`, three commits before PR #17 landed those banners. The audit
was measuring a tree no reader would ever see. `git rev-list --count HEAD..main` would have caught it, and it
is now in the open-of-iteration checklist. (D-M257x-21-2.)

The merge conflicted on 8 files / 18 hunks and was **resolved as a union**, not escalated — every hunk was
two correct texts, one carrying `main`'s merge detail, one carrying the router-deletion fact `main` lacks.
The rule this draws is in D-M257x-21-1.

## Re-measurement — four audit runs, and the fourth is the one that counts

| run | method | verdict |
|---|---|---|
| 1 | drift-surface sweep | **RED, 11 blockers** |
| 2 | same, after the merge + residue fixes | **RED, 5** |
| 3 | same, after those fixes | **YELLOW, 2** |
| 4 | **all 40 files read in full** | **RED, 21** |

Runs 2 and 3 each stated their findings were *pre-existing, not regressions*. So the 11 → 5 → 2 curve was
never convergence — it was **a grep measuring its own vocabulary.** The dominant failure mode here is a
correct banner contradicted by prose that never uses the banner's words: `make init-studio`,
`docker compose up -d graphql`, `JOBSIMULATION_RPC_ADDR=http://jobsimulation:8401`, a mermaid arrow. A
term-scoped audit cannot see any of them. (D-M257x-21-4 — written into `platform-alignment.md` §5.)

**Two of the 21 refute sentences that arrived from `main` in this iter's own merge** — `cms.md:6` and
`jobsimulation.md:7` say the merged services are *"not in the local compose"*, and `docker-compose.yml:144`
and `:83` say they are. The iter-20 map caught them within hours of their arrival, using the vocabulary term
(`running_but_unfederated`) that map introduced. First time in this milestone a corpus claim has been refuted
by the corpus's own fenced reference. (D-M257x-21-5.)

## Close — 2026-08-01

**Outcome:** The clause-5 sweep landed substantially — ~50 claims corrected across 19 files, four audit runs,
the branch brought level with `main` — but **clause 5 is NOT met**: the first full read of all 40 in-scope
files returns **21 blockers**, and the honest finding of this iter is that the earlier 11 → 5 → 2 curve was a
sampling artefact rather than convergence. The residual is now enumerated to `file:line` with a quote and a
correction each, so the next iter is mechanical.
**Type:** tik
**Status:** closed-fixed-partial (the sweep and the branch-level fix landed clean; the clause did not close,
and the remaining 21 are routed with a named handler)
**Gate:** NOT MET (3 of 5 — clauses 1, 3, 4)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (1 no-prog tik since iter-20 moved the metric) —
(3) re-scope: n (platform origin HEAD `2adcf71` re-checked at open AND close, unchanged — occurrence stays
1 of 2) — (4) user-blocker: n (the merge conflict was a union of two correct texts, resolved; D-M257x-21-1) —
(5) cap-reached: n (2 tiks of 5) — (6) protocol-stop: n — Outcome: continue.
**Session note:** the session exits here on **budget**, not on an enum condition. The numeric 5-tik cap was
NOT reached (2 tiks). Closing cleanly with the residual enumerated beats starting a 21-item sweep that cannot
finish — runs 6 and 9 of this milestone died mid-iteration and left uncommitted trees.
**Decisions:** D-M257x-21-1 … D-M257x-21-5.
**Side-deliverables:** the branch is now level with `main` (3 commits it had been missing for the whole
milestone); `platform-alignment.md` §5 gains the term-scoped-audit rule.
**Routes carried forward:**
- `DOC-M257x-iter21-full-read-residual` — **the 21 blockers, enumerated in `HANDOFF-next.md` with file:line,
  quote and correction.** Next tik. Expected to close clause 5.
- `CHECK-M257x-iter21-branch-behind-main` — add `git rev-list --count HEAD..main` to the open-of-iteration
  checklist beside the platform-HEAD re-check. Recorded in the handoff.
- clause 2's three causes, unchanged.

**Lessons:**
- **An audit scoped by search terms measures the terms, not the corpus.** 11 → 5 → 2 → **21**. The first
  three runs grepped the drift vocabulary; the fourth read the files. If the deliverable is *"this tree is
  true"*, the audit reads the tree — and the cost difference was minutes.
- **Check your branch against its base before you audit it.** Eight of the first run's eleven blockers were
  fixed on `main` and had been for days. One command would have shown it.
- **A sweep needs its own reader.** The audit caught a clause I had duplicated in my own previous commit.
- **Fixing the summary is the easy half.** Every one of runs 2–4's findings sat *below* a banner that was
  already correct. The reader who is misled is the one who scrolled.
