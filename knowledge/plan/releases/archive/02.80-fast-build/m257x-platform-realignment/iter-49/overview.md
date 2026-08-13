---
iter: 49
milestone: M257x
iteration_type: tik
status: archived
opened: 2026-08-03
---

# iter-49 — the two named fences, then the twelve, then the ninth reading

**Active strategy:** [`TOK-02: fence the prose the way the anchors are fenced`](../decisions.md#tok-02-fence-the-prose-the-way-the-anchors-are-fenced--2026-08-02)
— steps 1–3 (build the mechanical fences), step 4 (repair by CLAIM, fence-assisted), step 5 (ONE full read
at the frozen instrument). This iteration runs one more turn of 3 → 4 → 5, with the two fences iter-48
named as step 3's remainder.

**Multi-step planned shape, declared here** (Phase-2 scope-creep carve-out): three planned lines —
fences, repair, reading. A fourth unplanned line fires the tripwire.

## Step 0 — re-survey (mandatory, done before this plan settled)

- Platform origin re-fetched at open: **`2adcf71`, unchanged**. The re-scope trigger stays at occurrence
  1 of 2.
- `repair_postcondition --json` at open: **18 new sites**, all `claim_twin_guard`, every one keyed to a
  claim published by `iter-48/blocker-ledger.md` or an iter-48 raw seat report. This is the durable
  half of `D-M257x-48-12` — the audit commit's cost, still standing, exactly as that decision said it
  would be.
- The 12 blockers are still unrepaired and still at the anchors iter-48 recorded. Target is current.

## The user has ruled, twice

Clause 5 is **not re-cut**. It is met by a reading that returns zero blockers and by nothing else. The
user was offered close-at-4-of-5, a re-cut, or continue, and chose **continue**, knowing iter-48 measured
a run-to-run variance of about ±5 on a frozen instrument. This iteration does not re-open that question.
Its job is to make the cycle cheaper and more effective.

## Cluster / target identified

iter-48 traced all 12 blockers individually against the shipped fences and found **none of them
reachable**: every one sits at a **valid** anchor (so `anchor_construct_guard` is correctly silent),
`claim_twin_guard` only fires on claims already in a ledger, and `repair_leak_guard` is **verbatim-only**.
Two named handlers came out of that trace, and they are this iteration's steps 1 and 2.

## Hypothesis

1. **`FENCE-M257x-iter49-numeric-leak`** — the leak fence's blind spot is *structural, not a knob*
   (`D-M257x-48-9`): when a repair changes only a NUMBER, the rest of the sentence is present in both the
   removed and the added text, so the shingle that would locate the twin is suppressed as "the commit
   rewrote this", and the one token that did change is far below any distinctiveness floor. Lowering K
   was measured to buy two false positives and still miss it. The fix is to shingle over a **word-level
   diff**: a change becomes `(old tokens, surviving context)`, and the twin is found by the context
   carrying the OLD value. Expected to reach blocker **#3** (`architecture_overview.md:298`, 16-vs-23),
   which is a proven leak of `301d61a`.

2. **`FENCE-M257x-iter49-audit-commit-mode`** — the ratchet cannot distinguish *"a repair added a false
   claim"* from *"an audit refuted an existing one"* (`D-M257x-48-12`), so it hard-blocked iter-48's audit
   commit and the commit went through with `--no-verify`. The mode must let an audit commit land **without
   weakening anything** — so it must not touch the baseline at all (the monotonicity is the ratchet), and
   it must be **unable to launder a repair**.

3. **Repairing the 12 by CLAIM** (§5 rule 19), tree-wide, with the post-condition active, clears the 18
   ratchet sites and therefore the 8 traced `stack-core` failures.

## Expected lift — PRE-REGISTERED, before any report is read

**Predicted reading at step 3: 6 blockers**, split **2 induced / 4 pre-existing**.

Reasoning, stated so it can be refuted rather than reconstructed afterwards:

- The two-term model survives iter-48 with one input corrected. The **corpus term** is not zero — iter-48
  refuted that — and the eight readings `25 → 13 → 11 → 17 → 37 → 18 → 7 → 12` put the frozen
  instrument's run-to-run variance at about **±5**.
- The **repair term** has been 9, then 7, then 2. The two fences attack it directly, so it should not
  grow; predicted **2**.
- The **corpus term** measured 0 then 10 on the same instrument. Its expectation over the two readings is
  ~5; this repair removes 12 known members of it, but iter-48 showed a reading surfaces pre-existing text
  it had previously walked past. Predicted **4**.

**This prediction is registered before reading any seat report.** Four consecutive passes refuted their
own predictions; iter-41's held only once the instrument was frozen. **A zero reading sits inside the
instrument's own noise, so a zero would be weak evidence and a non-zero would be no surprise** — the
prediction is a falsifiable statement about the method, not a target to steer toward.

## Phase plan

| step | work | done when |
|---|---|---|
| 1 | `FENCE-M257x-iter49-numeric-leak` — built, **watched RED** on a fixture with a known answer key BEFORE any repair, mutation battery incl. inversions + a no-op control, reporting paths deleted to prove a test fails, false-positive count **pinned** | the fence names blocker #3 at `architecture_overview.md:298` from the `301d61a` diff, and its limits are pinned rather than believed |
| 2 | `FENCE-M257x-iter49-audit-commit-mode` — an honest audit mode for the ratchet, **provably unable to launder a repair** (that is the load-bearing test, and it is an INVERSION test, not a removal one) | an audit-shaped commit passes; a repair-shaped commit wearing the same flag is REFUSED |
| 3 | fixture capture for the 12 (perishable — §5 rule 21), then repair by CLAIM not by file, then ONE full 7-seat read at iter-41's frozen instrument | the reading is adjudicated and its induced/pre-existing split reported against the pre-registration above |

## Escalation conditions

- A fence that cannot be watched going RED before the repair spends its fixture → close the step
  `closed-no-lift` with the falsification, do not ship a fence that can only demonstrate GREEN.
- The audit mode passing a **repair-shaped** commit in test → hard stop; that mode is worse than
  `--no-verify`, because `--no-verify` is visible in a reflog and a laundering fence is not.
- A second platform commit invalidating an alignment attempt → `EXIT_REASON: re-scope-trigger`.

## Acceptable close-no-lift outcomes

A reading that returns non-zero is **not** a no-lift — it is the measurement clause 5 asks for, and
`overview.md` has pre-declared this since iter-47. The falsifications that would still satisfy the
protocol:

- The numeric-leak fence is built, watched RED, and **still** does not reach blocker #3 → then the class
  is pinned OUT OF REACH by a test (the `D-M257x-48-4` discipline), and the honest limit is the
  deliverable.
- The audit mode is shown to be unbuildable without a laundering hole → then the limit is written into
  the fence docstring and an audit commit continues to cost a recorded bypass.
