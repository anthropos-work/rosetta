**Type:** tik, under `TOK-05`. Lands the fence-level fix for the class the user's correction named.

# iter-91 — grade the cannot-tell: the stale-clone class, and the guard-pair enumeration

## The question, answered with a measurement

*Should `guard_family.py` refuse to run against a clone whose remote-tracking refs are stale, rather than
answering from what it can see?* — **A qualified no, and the qualification is the deliverable.**

**"Stale" is not locally decidable.** A clone fetched a minute ago can already be behind; one fetched last
week can be current. Answering it needs the network, and a fence that cannot run offline stops being run.
So the network check exists and is **opt-in** (`--verify-remote`).

**"Cannot see the objects it needs" IS locally decidable** — and that is the condition that actually bit. It
is fixed **at the point of use**, because only the guard knows which refs it needs; encoding sixteen
heterogeneous guards' ref requirements into the runner would rebuild §2's hand-maintained tuple in a new
costume.

## What was actually wrong, and it was worse than a stale clone

`platform_alignment_guard` resolves a citation at `origin/main`, then `HEAD`, then **silently at the
worktree** (provenance `worktree(no-ref)`). The provenance string was always there. **Nothing read it.** Its
only positive control was `subject_checked == 0` — total resolver failure — so **partial** blindness was
folded into a verdict.

| reference | before | after |
|---|---|---|
| `auto`, refs present | GREEN (90 resolved, 0 unresolvable) | **GREEN — unchanged** |
| the worktree fallback | RED, 8 findings, **4 unresolvable ungraded** | **UNMEASURED (exit 2)** |

The second row is the sharp one: it returned a **RED**, which *looks* like diligence, while 4 citations had
not been checked at all. A partial skip is worse than a total one because it arrives with a verdict attached.

## What landed

1. **`platform_alignment_guard`** — counts the silent worktree fallback per clone and grades it, plus
   `unresolvable > 0`, as **UNMEASURED (exit 2)** — neither pass nor fail. Escape hatch
   `ALIGNMENT_ALLOW_UNMEASURED=1` **records** the gap rather than hiding it, the same contract as
   `--allow-not-run`.
2. **`guard_family.py`** — the reference line, on every run: corpus sha, platform sha, its `origin/main`,
   whether they agree, fetch age. Refusal (`exit 2`, **before any guard runs**) when a platform-facing run
   has no `origin/main`. Opt-in `--verify-remote` for the only honest staleness check.

   > The gap nobody had named: the runner printed `platform=<dir>` and **never a sha**. Every `13 GREEN`
   > transcript in this milestone names a *directory*, not a commit — which is exactly how a green reading
   > gets quoted forward into a run brief with no way to re-check it.

3. **The 7-guard conjunction-pair enumeration** (Decision 1 item 3) — see `decisions.md` for the full
   21-pair grid. **12 pairs can interact; 5 were uncovered; 4 now have tests.**

The two landed here are the safety-critical ones, and they were covered by **nothing**: **every** G1/G6
escape test in the file drives `apply`, while **`revert` is the verb that writes AND runs
`git checkout -- <path>`**. `cmd_revert` did call the firewall — but nothing said so, and nothing would have
noticed its removal. Mutation-proven: deleting that call turns 3 subtests RED.

## Verification

- `stack-core`: **873 tests, 1 failure** — `test_claim_twin_guard_iter48_answer_key::test_02`, and it is
  **pre-existing, not mine.** Proven rather than assumed: the answer key was re-run against a copy of the
  milestone dir with `iter-90/` and `iter-91/` removed and produced **exactly the same 2 green-twin hits and
  101 claims**. It is substantive, and it belongs to the next iter: the claim it fires on (`iter-49/raw/C.md`
  C-2) asserts the `cms`/`jobsimulation` husks still run and still import `colony/authn` — which the
  `838d907`/`0c91421` platform move changed. **The answer key is stale because the platform moved**, which is
  this milestone's whole subject. Routed into iter-92's M810 sweep.
- `demopatch` conjunction battery: **8 tests, all green**, mutation-proven in both directions.
- Guard family at HEAD, with the new reference line: **15 GREEN · 0 RED · 0 could-not-check · 1 not-run**,
  every clone fetched, `platform @ 0c91421df (origin/main 0c91421df, in sync)`.

## Close — 2026-08-05

**Outcome:** the stale-clone class is fenced at the level where it is decidable — a guard run that could not
check part of what it was asked to check now reports **UNMEASURED** instead of a verdict, and every family
reading now records the commits it was taken against. The 7-guard pair grid is enumerated, and the two
safety-critical uncovered pairs (the path firewall on `revert`) are closed.
**Type:** tik
**Status:** closed-fixed — all three declared steps landed
**Gate:** NOT MET — **4 of 5, unchanged.** No reading taken; the instrument was under repair this iter, and
repairing the instrument inside a measuring pass is forbidden.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (2 tiks this session) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D-M257x-91-1 … D-M257x-91-5 (iter-91/decisions.md)
**Side-deliverables:** none.
**Routes carried forward:**
- `CHECK-M257x-iter91-g5xg7-journal-on-postcondition-failure` → the one interacting pair still uncovered;
  needs fault injection to trigger a genuine short write.
- `CHECK-M257x-iter91-guard-repo-root-scoping` → `story_org_count_guard` and `union_apply_guard` returned
  **GREEN against a temp dir with no corpus** — they ignore `--repo-root`. Surfaced by this iter's own
  mutation run. One of the two is in the family whose green this milestone quotes.
- `CHECK-M257x-iter91-claim-twin-answer-key-stale` → iter-92, folded into the M810 sweep (same cause).
- `CHECK-M257x-iter90-realmanifest-baseline` → re-derived and deliberately left open (`D-M257x-91-5`): the
  fix is to re-scope the assertion to *the anchor resolves exactly once*, not to re-pin a sha that will go
  stale on the next `make pull`.
- `CHECK-M257x-iter90-revert-idempotency` → still open.
**Lessons:**
- **A guard has three verdicts, not two.** Reserve an exit code for UNMEASURED and route every
  partial-blindness signal into it; an accept-the-gap flag is fine because it *records*, silence does not.
  Generalised into the protocol doc §8.
- **Print the reference with every verdict.** A verdict without its refs is an anecdote — and this milestone
  has been quoting one forward.
- **Fix a cannot-tell at the point of use, not in the runner.** The runner cannot know what sixteen
  heterogeneous guards need, and guessing rebuilds the tuple §2 exists to refuse.
