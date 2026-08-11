# iter-105 — closeout

**Shape:** tik · `iter_shape: fence` · **`TOK-06` step 0** — `FIX-M257x-iter103-guard-tree-provenance`.

## The one-line answer

**A guard verdict now states the tree its configuration lives in.** And the measured cost of it never having
done so: **52 recorded family verdicts across 26 milestone artifacts, 0 of which name that tree** — all
re-graded to *provenance-unstated*, not to *void*.

## What was broken, restated as a mechanism

`guard_family.py` printed the corpus sha (`:344`) and the platform sha (`:367`). It did not print its own.

That is not a cosmetic gap, because **the fence tree is the input that decides the verdict.** iter-103 ran
the family from the **pinned per-stack clone** (`09d06070`) instead of the **authoring copy** (`944fc4a2`),
read **2 RED** against its own sheet of `14 GREEN`, and drafted two conclusions on the spot — *"the sheet
asserted a verdict it did not have"* and *"a fence names 8 in-scope sites the double reading missed, so
`N ≥ 41`"*. **Both were false.** The entire difference was one file, `claim_twin_waivers.json` (+40 lines),
and the 8 RED sites were exactly the 8 waived sites.

Neither the corpus nor the platform had moved. **So nothing in the transcript could have told the reader
which of the two verdicts they were holding.** Run a fence from last release's pin and you measure last
release's fence; every waiver, baseline row and assertion added since reads as a fresh RED at sites nobody
touched.

## Landed

| | |
|---|---|
| `stack-core/fence_provenance.py` | net-new. `fence_tree()` → path · sha · dirty · describe, read from **where Python loaded the module** (`D-M257x-105-1`); `line()`, `unknown()`, `stamp()` |
| `stack-core/guard_family.py` | states the fence tree **first**, before corpus and platform; **`EXIT 2 — UNMEASURED`** when it cannot determine its own tree, with `--allow-unknown-provenance` recording the gap as `--allow-not-run` does; DIRTY disclosed on the **summary** line, not only the header |
| 17 members | every `*_guard.py` + `repair_postcondition.py` stamps on **direct execution**, so a standalone verdict carries it too |
| `stack-core/tests/test_fence_provenance.py` | net-new, **19 tests, all green** |
| `corpus/ops/platform-alignment.md` | **§5 rule 50** (the incident, the rule, the three design details, the retroactive re-grade) + **§8 rider 3** under *print the REFERENCE with every verdict* — *"the reference is THREE trees, not two"* |

**The fence told the truth about itself on its first run.** Every family run inside this iter printed
`fence tree 944fc4a21 is DIRTY — the verdict was taken with uncommitted configuration`, because the fence's
own edits were uncommitted at the time. It was not contrived; it is the disclosure working.

## Controls — TOK-06's binding clause, discharged

TOK-06: *every fence needs a mutation control AND an anti-vacuity control that can actually fire.* Six fences
in this milestone have been green over universes they never examined, and one compared a string to itself.

- **Derived, not listed.** The conformance check calls `guard_family.census()` — the same derived census the
  runner is built on — so a new `*_guard.py` that does not stamp turns the test RED **without anyone
  remembering to add it** (`D-M257x-105-4`).
- **Parsed, not grepped.** It walks the module's `if __name__ == "__main__"` AST for a call to
  `fence_provenance.stamp`, per §8's *parsed construct, never a whole-file substring*.
- **Mutation controls — 4 RED shapes, each one a shape a `grep` would have passed:** no stamp · *imports it
  and never calls it* · *mentions it in a comment* · *stamps at import time* (which does not travel with a
  standalone run). **Plus the control on the controls** — a correctly stamped module must PASS, or four
  reds prove nothing.
- **Anti-vacuity written against the SUBJECT, not the inputs** (§8's iter-94 rule): the discovered set must
  be **exactly** `guard_family.census()`, not merely non-empty — one stray file satisfies non-empty.
- **Live half:** a standalone guard run's **first** line is the stamp; suppression suppresses it; and the
  family states the tree **exactly once**, not once per member.

## Measurement — what this re-grades

```
grep -rnoE "[0-9]+ GREEN · [0-9]+ RED"  →  52 recorded verdicts across 26 milestone artifacts
                          naming the fence tree  →  0
```

The single line mentioning `rext` names a module path (`rext stack-core/guard_family`), not a sha.

**None of the 52 is thereby wrong**, and most were almost certainly taken from the authoring copy, which is
where the work happens. But *"almost certainly"* is the word this milestone exists to remove. They are
**unre-checkable** — strictly weaker than a green — so the honest grade is **provenance-unstated until
re-run**, stated once in §5 rule 50 as a reading instruction rather than stamped retroactively onto 26
artifacts (`D-M257x-105-5`: inventing evidence to fix a lack of evidence is the class, not the cure).

## Tests

**`stack-core`: 937 passed · 1 failed (16 m 52 s).** The single failure is
`test_claim_twin_guard_iter48_answer_key.py::test_02_the_green_twin_of_every_site_stays_SILENT`.

**It is PRE-EXISTING and this iter did not cause it** — proven, not assumed. Reproduced three ways via
read-only `git archive` (no worktree, no stash): the **pre-change rext tree** `944fc4a` against the live
corpus; `944fc4a` against the **run-open corpus** `22eaac4`; and the changed rext against `22eaac4`. Same
single failure, identical output, all three.

**What it is:** two *green-twin* fixtures (`corpus/04.md`, `corpus/05.md` — the rewritten prose that no
longer carries the refuted claim) now fire, both sourced from `iter-49/raw/C.md:57`, the C-2 blocker row.
`claim_twin_guard` itself is **GREEN on the live tree**; what is RED is its **discrimination control** —
§8 rule 5's *a battery of REDs cannot tell a discriminating fence from a brittle one*.

**The likely cause is worth naming even though the fix is routed:** iter-102 grew the adjudicated-claim set
**134 → 264** by publishing ledgers in the shape `claim_ledger.py` derives from. A ledger that grows derives
more claims, and one of them now matches prose written specifically to be clean. That is **repair-induction
in the fence layer** — the second inflow TOK-06 named, one layer below where TOK-06 was looking.

**Not fixed here, deliberately.** It is the iter's third line of investigation and the scope-creep tripwire
fires: deciding whether the *fixture* or the *ledger derivation* is wrong is a real investigation, not a
one-liner. Routed with a named handler, characterised rather than merely deferred.

## Guard family at close

`14 GREEN · 0 RED · 0 could-not-check · 3 not-run` over 17 members — **and for the first time the transcript
says which fence tree said so.**

## Gate

**Unchanged at 4 of 5, and no movement on `N` is claimed.** This is a step-0 instrument iter. It touches
clause 3's instrument (the guard family), never clause 5's (the graded read); the milestone's standing rule
that those are two instruments applies, and neither speaks for the other. Clause 5 was not re-cut, narrowed,
reinterpreted or argued.

## Housekeeping

Zero platform-repo edits. `stack-demo/**` untouched, no stack brought up or reconfigured, **no clone
fetched** (the `git archive` reproductions are read-only extractions of commits already local). rext stays on
`main`; **no tag cut** — nothing here must be consumed by a stack, so `944fc4a` + this commit fold into the
next cut alongside `4cb920a`.

## Close — 2026-08-06

**Outcome:** guard-tree provenance shipped with derived + AST-based + mutation-proven controls; 52 prior
verdicts across 26 artifacts re-graded provenance-unstated; `stack-core` 937/938 with the 1 failure proven
pre-existing at the run-open tree.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: **n — the
one RED is PRE-EXISTING, reproduced at the run-open tree with unmodified rext, does not block this iter's
deliverable, and is routed with a named handler; Phase 5 §4's NOT-list case "new findings discovered
mid-iter"** — (5) cap-reached: n (1 tik) — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue
**Decisions:** `D-M257x-105-1` (the tree is where Python loaded the module — flag/cwd/env each reintroduce
the defect) · `-2` (DIRTY disclosed, UNDETERMINABLE refused, and why the caveat rides the summary line) ·
`-3` (print FIRST — the iter-87 `headline()` shape) · `-4` (derived from the census, asserted over the AST,
anti-vacuity against the subject) · `-5` (re-grade to provenance-unstated, do not stamp 26 artifacts)
**Side-deliverables:** none — the pre-existing RED is a finding, not a fix.
**Routes carried forward:**
- **`FIX-M257x-iter105-claimtwin-green-twin-refire`** *(net-new)* — the green-twin discrimination control is
  RED; two fixtures fire from `iter-49/raw/C.md:57`. Characterised above; likely iter-102's 134 → 264 ledger
  growth. **Target: TOK-06 step 2 (the induction checks), where it belongs by class.** Decide whether the
  fixture or the derivation is wrong — do not silence the control.
- TOK-06 steps 1–4 unchanged: the drift fence next, then the induction checks, then the 33, then the read.
- Still open and unmoved: `FIX-M257x-iter56-assignment-flake`, `FIX-M257x-iter103-assignment-context-bleed`,
  `FIX-M257x-iter103-read-union`, `DEF-M257x-iter103-aws-bind-provenance`,
  `DEF-M257x-iter101-briefing-rext-tree`, `RF-2/3/7–14`, the five pass-22 items.
**Lessons:** **a fence that discloses its own weakness on its first run is doing its job, and the disclosure
should be the thing you cannot lose.** The DIRTY caveat was added to the *summary* line rather than only the
header for one reason — harden pass-20 measured that the summary line is what gets quoted forward — and it
fired immediately, on this iter's own uncommitted state. The general form: **put a caveat where the
quotable sentence is, not where the careful reader is.**
