# iter-145 — decisions

## `D-M257x-145-1` — a diff SCOPE cannot grade a failure; only a bisect can

Harden passes 30/31/32 routed 21 failures as *"provably not ours"*, evidenced by
`git diff --name-only 6ad8866..HEAD` returning 5 files, all `stack-core`. **The command is correct and
the inference is not.** It asks *"did the iters in THIS window touch the failing section?"* — a
question whose answer is `no` for every defect older than the window, including one this milestone
introduced itself.

Graded: **12 of the 21 were introduced by M257x iter-13** (`4414527`, 2026-08-01), 132 iters before
the window opened. They are not merely "ours" — they are the milestone's own, in the section the
milestone's own definition of "the suite" excluded.

**The rule:** a failure's authorship is settled by finding the commit that turned it RED, not by
checking whether a recent diff overlaps it. Where a bisect is too expensive, the honest wording is
*"not attributable within this window"* — which is a statement about the measurement, not about the
code. It is the same distinction `§9` guard-rail 1 already enforces for the primary metric
(**UNMEASURED is not UNMOVED**), applied to attribution.

## `D-M257x-145-2` — keep the MEMBERSHIP literal, derive the COUNT, fence the two

`test_verify.py` held **seven** copies of one platform fact: a `BASES` map naming every
`services.sh` row, and six independent literals restating how many rows there are (13 · 13 · 12 · 10 ·
13 · 14). iter-13 updated the table and none of the seven.

The tempting fix — derive everything from `services.sh` — is **wrong**, and `§8`'s anti-vacuity rule
says why: the port expectations are the control for the offset arithmetic, and a control read out of
the subject asserts nothing. The tempting opposite — hand-correct all seven — re-commits the defect.

**The split that holds:**

| what | how | why |
|---|---|---|
| the **membership + ports** (`REGISTRY_BASES`) | hand-written, one copy | it is the anti-vacuity control; it must be independent of the table |
| every **count** | derived from that map | a count is not a control, it is a restatement — and restatements rot |
| the **agreement** between map and table | a named fence | so drift fails ONCE, loudly, naming the row |

The fence's value is not that it detects the drift — the arithmetic already did, six times. It is that
its message **names the row and the file**. `12 != 13` was visible on every run for four months and
told no reader anything. Verified by control (`§8`): a re-drifted copy fires both, and only the fence
says what moved.

## `D-M257x-145-3` — the scope call, stated as an assumption pending the user's ruling

**Assumption taken by this iter:** *"the suite" for M257x means all five `rosetta-extensions` sections
(`demo-stack`, `dev-stack`, `stack-core`, `stack-injection`, `stack-verify`), not `stack-core` alone.*

The user has **not** ruled on widening it, and this decision does not pre-empt the ruling — it records
what one widened run measured, so the ruling has evidence rather than a routed count:

- The narrow definition covered **1,281 of 3,062** tests executed. Every "whole suite" claim in 144 iters
  and 32 harden passes was made over ~42 % of the suite.
  <br>⚠️ **These operands read `1,280 of ~3,040` until M257x iter-173**: the denominator was
  `2,978 passed + 11 skipped`, dropping the same table's **22 failures** — a unit switched from *executed*
  to *passed-and-skipped*. The **~42 % is unchanged**, which is precisely why it went unnoticed for 28
  iters. Now fenced at the operand level by `stack-core/derived_count_guard.py`.
- Widening it **once** surfaced a live defect of the milestone's own making, RED for 132 iters.
- The cost is ~11½ minutes of wall clock for the four extra sections — **less than half** what
  `stack-core` alone costs (20–35 min).

The cost argument that would justify the narrow scope does not survive measurement: the excluded 58 %
is the cheap half.

**What is NOT claimed:** that every iter must run all five. `§5` rule 60's whole-suite debt is about
iters that *touch* rext, and this iter touched one file in one section. What is claimed is that the
phrase *"the whole suite"* must name its denominator — the milestone's own standing rule, applied to
the milestone's own reporting.
