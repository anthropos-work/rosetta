---
iteration_type: tik
status: in-flight
milestone: M258
iter: 05
active_strategy: TOK-01
created: 2026-08-12
---

# M258 iter-05 — a GATEABLE bring-up half, and the 395-vs-287 question

**Type:** tik · **Active strategy:** `TOK-01` step 1 — *measure the composition before engineering it.*

## Step 0 — re-survey

`TOK-01` step 1 is now **half discharged**: iter-04 produced the batch half (129 s, n=1, contended). What
step 1 still owes is a **gate-quality** number — every bring-up figure this milestone has taken so far was
refused by, or taken outside, the gate instrument:

| figure | mode | gateable? |
|---|---|---|
| M257 iter-09 **286.99 s** (n=3 p50) | `--public-host billion…` on `macmini` | ✅ the inherited baseline |
| iter-02 **395.31 s** (n=1) | public-host (**auto-discovered**, believed single-box) | ✗ n=1, and the mode was not what the record claimed |
| iter-04 **781 s** (n=1) | single-box, operator-driven | ✗ headroom refused; `load1` 16 → 62 |

**Re-survey at open:** `assert-headroom --profile macmini` → **OK** — `load1` **3.49**, free **65.9 GiB**,
`lanes=1 max_parallel_ui_lanes=2`. The third-party load that blocked iters 03–04 **ended** immediately
after the batch. The window is open, and on this host it is **bursty** — so the campaign was launched
first and this plan written while it runs.

## Cluster / target identified

**`MEASURE-M258-gateable-composition`.** One `buildbench run 1 --reps 3 --profile macmini
--no-public-host --label m258-iter05-gateable`. `--reps 3` rather than 1 deliberately: it is the only
thing that answers the **inherited priority question** — *explain the 395.31 vs 286.99 delta before
treating either as "the bring-up half"* — because that question is a comparison of a **n=1 contended
public-host sample** against a **n=3 p50**, and the missing term is a n=3 p50 in the mode the gate is
actually taken in.

## Hypothesis

A single-box n=3 p50 lands **near M257's 286.99 s**, and the 395.31 s excess is explained by some
combination of (a) n=1 sampling under `load1` 2.26, and (b) `ui_studio_desk` — the 115.35 s leg L1 never
touched (`CHECK-M258-iter02-studio-desk-is-the-untouched-leg`). The campaign's per-sub-phase p50 table
tests (b) directly.

## Expected lift

A **gateable** bring-up p50 with min/max, in the mode `TOK-01` declared the gate is taken against — plus
the per-sub-phase breakdown that either confirms or kills the studio-desk suspicion. Composed with
iter-04's 129 s batch, the **first honest arithmetic against the 480 s ceiling**.

## Phase plan

- **Phase A** — headroom pre-flight (**done: OK, `load1` 3.49**).
- **Phase B** — the 3-rep campaign, foreground-polled, `bringup_argv` confirming the mode per rep.
- **Phase C** — read the report: p50, min/max, `gateable`, per-sub-phase table, ISOLATION + identity per
  rep. Compose with the batch half; state the environment with every number.
- **Phase D** — close.

## Escalation conditions

- **Composed p50 > 600 s** → the declared `re_scope_trigger` — but only if the p50 is **gateable**, per
  `D12`. A contended or partial campaign does not fire it.
- **A rep aborting on headroom mid-campaign** → a partial result, reported as partial. Expected on this
  bursty host and **not** a failure.

## Acceptable close-no-lift outcomes

A campaign that aborts because the user's load returned mid-run. The reps that completed are still the
first gateable single-box samples this milestone has, and a refusal is a result (`D9`).

## Known-context carried

`C2` — publish the **spread** beside any p50; `buildbench` reports `min`/`max` natively, which is what
that caveat asks for. Never quote M256's 56.6 s as the batch half.
