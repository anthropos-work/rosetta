# M258 — retro

**Closed 2026-08-12 · `closed-incomplete` — ACHIEVED BY USER RULING, NOT ON GATE.**

## Summary

M258 set out to make a stack come up **and prove itself**, in one cold command, fast enough that this
becomes the normal way to bring one up. It shipped that: the bring-up now ends in a **batch gate** that
drives every seeded hero's journey to completion, emits one consolidated red set, exits non-zero and loudly
on a non-empty one, and leaves the stack **UP regardless** — with a **world-contract restore leg** because
the suite's reset is destructive to the presenter demo.

What it did **not** ship is the number. **Clause 3 — composed p50 ≤ 480 s over 3 consecutive cold cycles —
was never measured clean, and must never be recorded as met.** The host is a permanently contended
workstation; the user concluded that contention is not removable and ruled the goal achieved on the other
four clauses plus a ~402 s projection. That is the shape of M257x's `TOK-09`, not of M257's fired gate.

The user then reshaped the milestone a second time, adding **space optimisation as a new goal** (`TOK-02`)
under the constraint *space must not be bought with time*, with a hard end state: exactly one stack up,
built with the new mechanism from the newest platform repos. Both delivered.

## Incidents this cycle

1. **The 15-red batch verdict (iter-15 → iter-17).** A cold batch returned `verdict: red, red_count: 15`
   and was **escalated, not swept**. Attributed to **our own tooling**: platform `766df6c` folded
   `sentinel` into `app` — the 8th merge — while three of our post-seed reload sites still drove the
   deleted container's RPC and logged the miss as *"non-fatal — a non-AI-sim run is unaffected"*, which was
   false. A stale casbin enforcer refuses **every** org-scoped read and write with `forbidden` at HTTP 200.
   Fixed, proven green cold. **`batch_seconds` 629 → 129: the suite was slow because it was broken.**
2. **The set-dress 3.49× scare (iter-08).** `set_dress` measured 283.53 s against iter-05's 81.23 s, which
   would have put the clean total near 600 s. **Refuted by two cheap checks** — all 3 reps did identical
   work (same digest), and a `git diff --name-only` filtered for the set-dress paths returned **zero
   files**. Environmental, not a regression.
3. **A false ISOLATION RED cost iter-02 its measurement.** `.env.demo-1` held 24 appended Clerkenstein
   blocks; `_stack_minted_pk` read **first-wins** while every consumer read **last-wins**. Both causes the
   assert named were **refuted**; the real chain was a third.
4. **P2 — one test failure introduced by this close, fixed inline.** A prose comment it added contained a
   phrase `test_frontend_build.py` counts, taking a fence 3 → 4. Reworded. Trivial, and recorded because
   the class is not: **a comment can break a fence that counts a phrase.**
5. **P2 — the milestone's own fence-set runner reported two false REDs**, both the runner using the wrong
   invocation contract. The harden ledger had written that trap down one pass earlier, and it still caught
   its own author.

## What went well

- **The escalation was closed, not carried.** `D-v28-3` says a non-empty red set escalates for
  renegotiation. It did, once, with a diagnosis and a fix — not a disposition.
- **The partition was the proof.** 15 failing journeys all org-scoped, 15 passing all user-scoped. That one
  observation named the mechanism before any code was read, and three competing stories (contention, a
  partial seed, "a table moved") were **refuted from artifacts already captured** — no new runs.
- **Refutation over adjective, repeatedly.** The set-dress scare died to a `git diff`. The ISOLATION RED's
  two named causes both fell. `FIX-M257-content-stories-pair-count`'s inherited description — carried
  through three milestones — was **refuted at this close** rather than re-shipped. **A code read is not a
  measurement.**
- **The order that protects the user was enforced in code, not by discipline.** `teardown-others.sh`
  refuses unless the survivor is up, and re-checks between every step. The user's stack was torn down last
  of all, after its replacement had proven itself.
- **The constraint was honoured rather than quoted.** 11.54 GB reclaimed at **zero build-time cost**, with
  the 21.03 GB reclaimable build cache **deliberately untouched**.
- **Nothing flattering was banked.** The last iter refused a ~290 s cycle because it was warm-cache on the
  quietest box of the milestone. That refusal is preserved verbatim in three artifacts.

## What didn't

- **The gate's own number was never taken.** Twenty iters, two campaigns, an auto-arming trigger — and
  still no clean p50. The milestone's one genuine unknown was environmental, and no amount of engineering
  addressed it. `TOK-01`'s instinct (*measure the composition before engineering it*) was right; what it
  could not anticipate was that the measurement itself was the thing the host would not give.
- **The milestone ran 20 iters and was never hardened until the end.** The final pass then found **four
  defects in its own youngest code** — including a fence **satisfied by its own comment** and two of three
  reload sites unfenced, one of them the restore leg. Hardening at ~10 tiks would have caught them sooner.
- **The close found eight more, two of them fail-opens to GREEN.** A crash in the red-set reader printed
  nothing, and "nothing" was graded as "red set empty"; a hook that re-decided the opt-out wrote no verdict
  at all, letting the previous run's file be read as this one's. **Both inverted the milestone's central
  claim, in the one script whose header enumerates the other fail-opens it closed.** The pattern is that
  every fixture hand-wrote the shape the producer never emits.
- **The corpus sweep needed two passes.** iter-18 corrected 26 files; the close found 16 more still
  asserting a live `sentinel`. §5 rule 54 — *a correction that reaches one cell is not a correction* — was
  written for exactly this, and this milestone re-proved it on itself.
- **The full `stack-core` sweep still does not complete on this host**, across two milestones now. Scoped
  runs plus pristine-extract attribution were used instead. **That is not "swept clean" and must not be
  upgraded to it.**

## Carried forward

Everything routes to **`/developer-kit:close-release` for v2.8** as **one conscious block fate, named item
by item** — M258 is the release's final milestone, so there is no later in-release destination.
**0 escape-hatch deferrals.** Full inventory: [`carry-forward.md`](carry-forward.md) +
[`deferrals-audit.md`](deferrals-audit.md).

Three items are owed the user's **explicit fate** at that close, because crossing a release boundary
revokes a deferral's authority: **`F2`** (`ptvalidate` unwired, M256→M257→M258) ·
**`PROFILE-M257-provisional-fields`** (M255→M257→M258) · **`RATCHET-M257-literal-ceilings-breached`**
(a pre-existing breach of 8 — 249 against a ceiling of 240, **never raised by anyone**).

And the standing safety item, still open: **a demo reached the production S3 bucket**, and containment is
proven by a unit test on the emitter and **on no running stack**.

## Metrics delta

| | |
|---|---|
| iters | 20 (18 tiks + 2 toks, one **user-directed**) · 0 orphan iters, 0 orphan commits |
| harden | 5 passes, **STABILIZED** · 18 mutants proven rejected · 4 defects fixed inline |
| close | **8** further defects fixed · **14** new fences, every one RED-proven against a *faithful* pre-fix mutant |
| tests | demo-stack 1140 pass / 9 pre-existing · dev-stack 180 (+14) · stack-verify 281 (+6) · stack-injection 335 (+1) · all 6 Go sections green |
| flakes | **0** — 5 consecutive sequential passes, identical counts |
| fences | **18/18 `rc=0`** against platform `766df6c` |
| ratchet | 249 vs a ceiling of 240 — **pre-existing by 8**; this close added **0**, and no ceiling was raised |
| space | **11.54 GB** of real SSD at **zero build-time cost**; build cache untouched |
| platform edits | **0** |

Full numbers, each with its status attached: [`metrics.json`](metrics.json).

## The lessons this milestone paid for

- **The partition is the proof.**
- **A fence can be satisfied by its own comment — and a comment can break a fence that counts a phrase.**
- **A "non-fatal" log line can be false.**
- **Never read an exit code through a pipe.**
- **A column can answer a different question than the one you asked it.**
- **A code read is not a measurement.**
- **A fence proven by an insufficient mutant is not proven.**
- **Never quote a fence count without its invocation and its ref.**
- **Refute with mechanisms, not adjectives.**
- **Never edit correct prose to make a fence green, never `--no-verify`, never raise a ratchet ceiling.**
