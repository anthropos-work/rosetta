---
iter: 02
milestone: M256
iteration_type: tik
status: closed-fixed
opened: 2026-07-28
---

# M256 · iter-02 — the baseline (the denominator every later claim divides by)

**Type:** tik · **Active strategy:** `TOK-01` move 1 ("measure the denominator first, on this box, and
change nothing until it exists").

## Step 0 — re-survey

TOK-01's next-tik direction names the baseline measurement. Re-surveyed: `../progress.md` § Baseline reads
**"Not yet measured"**; no harness code has changed (iter-01 touched only `knowledge/` + `corpus/`). Target
stands, unmodified.

## Cluster / target identified

Gate clause 1 is **relative** (D-v28-12): median per-Playthrough ≤ **0.79×** a same-stack pre-work
baseline. With no measured starting point the clause is unfalsifiable — so this is the one iter that must
run before any other, and the one iter that must change **nothing**.

## Hypothesis

None to test — this is a measurement iter. The *deliverable* is a defensible denominator: `n=3` consecutive
suite runs on the live local `demo-2`, with the environment stated alongside the number
(`latency-budget.md`'s rule), plus the three figures D-v28-9 requires to be kept apart:

1. **median per non-LLM Playthrough** — the **gated** metric;
2. **suite wall-clock** — **reported, not gated** (the denominator grows 18 → ~27 inside this milestone);
3. **the studio lane** — `pt-studio-advanced-generate` (+ `-guided-`), an irreducible live-LLM
   round-trip, excluded from the median and budgeted separately.

## Expected lift

**Zero, by design.** A tik that changed the metric here would invalidate the thing it is producing.
Success = a recorded, reproducible baseline + a per-test breakdown good enough to *target* iter-03.

## Phase plan (protocol: `corpus/ops/demo/playthroughs.md` § The iteration protocol, steps 4 + 6)

1. `run-playthroughs.sh 2 --reset` ×3, serial, `retries: 0`, with the demo's `bin/` on `PATH` (the M204
   iter-05 gate-run prereq) — reset-to-seed each run so P6 holds across the three.
2. Harvest per-test durations from `report/last-run.json` after each run.
3. Compute the three figures; record with the environment; note any red test as a **pre-existing** red
   (D-v28-3's batch-gate rule: full batch to completion, one consolidated red set at the end).
4. Write the per-test table so iter-03 can target the slowest non-LLM tests directly.

## Escalation conditions

- The suite cannot run at all (login broken, seed broken) → `user-blocker`; a milestone with no measurement
  surface cannot proceed and must not fake one.
- A **non-empty red set** at batch end → per **D-v28-3** this escalates to the user for renegotiation
  (fix, or an explicit written disposition). A pre-existing red is not this iter's to silently absorb.
- Run-to-run variance so wide that a 21 % target sits inside the noise → record it and say so; the gate
  would need a decision, not a guess.

## Acceptable close-no-lift outcomes

Any honest baseline closes this iter `closed-fixed` — including one that shows the suite is *already* fast,
or that variance is high. What would be a failure is a single-run number reported as a baseline.
