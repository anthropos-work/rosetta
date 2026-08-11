---
iter: 116
milestone: M257x
iteration_type: tik
iter_shape: reading
status: archived
opened: 2026-08-07
---

# iter-116 — `TOK-07` step 3: THE READ

**Active strategy reference:** [`TOK-07`: enumerate the predicate, not the anchor](../decisions.md) —
milestone-root `decisions.md`, authored 2026-08-06 as a **deliberate, non-terminating** tok. Its declared
step order is **enumerate → repair whole predicates → read LAST**. Steps 0, 1, 1b and 2 landed at iters
111, 112–113, 114 and 115. **This iter is step 3, and it is the whole of it.**

## Step 0 — re-survey before targeting

Re-ran the protocol's primary measurement's *preconditions* rather than the measurement itself (the
measurement **is** this iter's work, and running it twice would be the reading):

- **The target is still untouched and still meaningful.** `P` has been **UNMEASURED since iter-109**, not
  unmoved — iters 110–115 each said so in §9's mandated words. Nothing has re-measured it.
- **The repair it was waiting on is complete.** iter-115 closed **24 of 24 predicates at every one of 71
  enumerated sites**, `denominator: corpus-derived-per-predicate`, exit 0.
- **The subject has not moved under it.** All 14 platform clones are at the sha they held at iter-103's
  and iter-109's reads. The corpus moved +177 lines, all of it iter-115's repair.

No substitution. `TOK-07`'s named next step is current.

## Cluster / target identified

There is only one target: **take reading #29/#30 of `corpus/services/**` + `corpus/architecture/**`
against the platform clones, under the frozen instrument, and report `P` and `N` whatever they are.**

## Hypothesis

`TOK-07`'s thesis is that iter-109's flat `P` was caused by the repair's **denominator** — a site list
derived from what a prior reading *detected*, so twins one file away survived. iter-115 replaced that
denominator with a corpus-wide per-predicate enumeration and closed every instance. **If the thesis is
right, `P` collapses. If it is wrong, the pool is simply larger than a reading samples, and `P` holds.**

## Expected lift

This iter's deliverable is a **number**, not a lift. The reading may return any value and each is a real
result. The pre-registered rule (sealed before any seat is dealt, in `pre-registration.md`) attaches a
verdict to each branch, including one that **refutes the active strategy**:

- `P ≤ 6` → **DENOMINATOR** — the pool drains.
- `7 ≤ P ≤ 14` → **PARTIAL DRAIN**.
- `P ≥ 15` → **VOLUME** — `TOK-07`'s own pre-registered refutation of repair-and-read fires. The next move
  is a **re-scope conversation with the user**, not an eighth strategy revision.
- `P = 0` → clause 5 is met.

## Phase plan

1. Re-derive ground truth from scratch — partition (with its reproduction control), clone refs + fetch
   times with **no fetch**, instrument sha re-checked **after** copying, guard family with its own fence
   tree printed. → `ground-truth.md`
2. Author and **seal the pre-registration in its own commit before any seat is dealt.** → this commit
3. Deal 14 blind seats (7 × reading #29, 7 × reading #30) over the identical partition, in two batches of
   7. **Commit each seat verbatim the moment it lands.**
4. Adjudicate with 4 independent adjudicators, each re-deriving from the clones. **Adjudicate before
   reporting `P` or `N`.**
5. Grade every band; report the upheld rate **twice** (raw, and `wrong-tree`-separated); close.

## Escalation conditions

- **`P ≥ 15`** → do **not** author `TOK-08`. Report the refutation first and loudest, close the iter
  honestly, and exit with `EXIT_REASON: re-scope-trigger` — `TOK-07` pre-registered this on 2026-08-06,
  before iter-111 ran.
- **`P = 0`** → clause 5 is met; `GATE: MET`.
- A seat or adjudicator dying mid-flight → the per-seat verbatim commit bounds the loss to that seat.

## Acceptable close-no-lift outcomes

A reading is a complete iter whatever it returns. `0 < P < 15` is progress short of the gate and closes
`closed-fixed` on the strength of the measurement, not the number. **What would NOT be acceptable is a
reading that quietly omitted a band, re-cut a band after seeing the number, or repaired anything inside
the measuring pass.**

## Known soft spots this reading must not paper over

- **`FIX-M257x-iter113-adjudication-is-judgement`** is open and carries **three measured misses** from
  iter-115 — three sites publishing an enumerated predicate that the enumeration did not enumerate.
  Fifteen of sixteen "small class" verdicts rest on one agent's readings. **Band #3 is the direct test.**
- **`FIX-M257x-iter116-predicate-guard-takes-first-pin-in-block`** — net-new, found at this open, and the
  guard family is **RED** on it. Measured and characterised in `ground-truth.md`; it is a **false** RED
  (the corpus sentence is true at the ref it names) and it is **routed, not repaired**.
- Anchor rot is live during any corpus edit — but this iter makes none.
