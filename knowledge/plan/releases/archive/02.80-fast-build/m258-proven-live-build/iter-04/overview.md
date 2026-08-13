---
iteration_type: tik
status: in-flight
milestone: M258
iter: 04
active_strategy: TOK-01
created: 2026-08-12
---

# M258 iter-04 — take the batch half

**Type:** tik · **Active strategy:** `TOK-01` step 1 — *measure the composition before engineering it.*

## Step 0 — re-survey

`TOK-01` step 1's deliverable — **the first wall-clock the batch half has ever had** — is still
outstanding after two tiks, and the target is **unchanged and still meaningful**. What changed is *why*
it is outstanding:

| iter | why the batch half was not measured |
|---|---|
| 02 | believed the UI tier was wired to real Clerk (**refuted**); the true blocker was that the campaign silently ran **public-host**, in which the batch cannot be driven from this host at all |
| 03 | that blocker **removed and live-proven removed** — ISOLATION green on the stack that reded, `--no-public-host` now expressible. Blocked instead by **third-party host load** (`load1` 39–46 vs 12 cores) |

Re-verified at iter-04 open: `assert-headroom --profile macmini` still **FAILs** on `peak_load1`
(14.77 vs the cores−2 floor of 10), trending down from 45.94 — the load is decaying, not permanent.
Pin confirmed `fast-build-m258-iter-03`, on origin, and the consumption clone is checked out at it.

## Cluster / target identified

**`MEASURE-M258-batch-half`** — the milestone's primary unknown, and the only routed item that gates the
gate. Everything else in the queue is instrument hygiene and stays routed.

## Hypothesis

With the mode blocker gone, one cold `--no-public-host` cycle followed by the full Playthrough batch
yields both halves of the composition against the 480 s ceiling.

## Expected lift

The batch half **exists as a number**, with its environment and `load1` stated, plus the **red set** —
which is the gate's *other* clause (zero standing red) and, unlike a wall-clock, **survives contention**.

## Phase plan

- **Phase A** — headroom gate. Bounded wait for `load1 < 9.8`, already armed.
- **Phase B** — `buildbench run 1 --reps 1 --profile macmini --no-public-host --label m258-iter04`,
  foreground-polled. Confirms P1/P3/P4 live: `.env.demo-1` collapses to **one** block, the minted host is
  `127.0.0.1`, `bringup_argv` records the mode, and `inject.py`'s report is visible in the log.
- **Phase C** — `run-playthroughs.sh 1 --reset`, timed. `D-v28-3` semantics: runs to completion, never
  halts at first red, never retries, **one consolidated red set at batch end**.
- **Phase D** — close: publish the batch half **with its spread caveat** (`C2`) and the composed figure
  against 480 s.

## Escalation conditions

- **Non-empty red set** → ONE consolidated escalation at batch end (`D-v28-3`), never a halt.
- **Composed p50 > 600 s** → the declared `re_scope_trigger`, surfaced **with measurements attached**.
- **Headroom still refusing at the cap** → the timing is **not taken as a gate number**. Fall back to
  harvesting what contention cannot corrupt (the red set, the P1/P3/P4 live booleans) and label every
  wall-clock **contended / not-gateable**. A refusal is a result (`D9`).

## Acceptable close-no-lift outcomes

- A batch half that **misses** the composed budget — provided its spread is published beside it.
- A headroom refusal that holds for the whole iter, recorded with the load trace.

Neither is a failure; both are the measurement discipline `TOK-01` was written to protect.

## Known-context carried

`TOK-01` § *Known-context* #1–#6, minus `R0` (discharged iter-02). `C2`'s **2.04× decidability caveat**
is the live one: publish the spread beside any p50, and never quote M256's **56.6 s / 18 specs** as the
batch half — the suite is **30 live Playthroughs / 209 passed**.
