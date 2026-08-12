---
milestone: M258
iter: 9
iteration_type: tik
status: closed-fixed
opened: 2026-08-12
---

# M258 iter-09 — split the `set_dress` anchor so `LEVER-M257-L5-setdress` can be aimed

**Type:** tik · **Active strategy:** `TOK-01` (bootstrap) — *measure the composition before engineering
it.* This iter is that strategy applied one level down: **the largest phase in the cycle is a single
un-attributed span, and it must be measured before any lever is spent on it.**

## Step 0 — re-survey before targeting (mandatory)

`TOK-01`'s next-tik direction is step 4, *the composed 3× cold campaign*. iter-08 **ran** that campaign
and it is **not** stale — but it returned `headroom=FAIL` on 3/3 reps, so its only remaining input is
**~45 minutes of host at `load1 < 10`**, which is not a work item. Measured at 09:22:06Z: **`load1
24.86`** (1/5/15 = 24.86 / 20.88 / 22.57). The window is not open and cannot be opened by working.

The re-survey therefore substitutes the target **without leaving the strategy**: iter-08's own close
routes `FIX-M258-iter08-set-dress-has-no-internal-attribution` as *"the natural iter-09 if the box stays
loud"*, and it **needs no host window**. Step 4 is not abandoned — it is **armed** (Phase C) rather than
waited on, per iter-08's `D29`.

## Cluster / target identified

`set_dress` is the largest single phase in the cycle and **89 % of it is one anchor-free span.** From
iter-08 Phase D: `370.98 → 623.71` = **252.73 s** with no intermediate log line, "running from the
Directus bootstrap straight to the taxonomy replay's completion. Two distinct operations, one anchor."

This is `D17`'s shape one level down — and `D17` is the defect where **a 166 s phase hid inside a 2 s one
and the phase table still summed.** *A table that adds up is not a table that attributes.*

## Hypothesis

The span is not "two operations" but **one dominant operation with a small head**, and the dominant one
is a single `stacksnap replay --surface taxonomy` that moves **~1.47 GB** of payload
(`public.skill_embeddings.copy` **825 MB** / `public.job_role_embeddings.copy` **364 MB**, measured in
the local snapshot store) and rebuilds **two** pgvector indexes (42,790 + 22,470 vectors). Splitting the
anchor will therefore not divide the cost in two — it will **concentrate it**, and the useful split is
the one *inside* the replay: **COPY versus REINDEX**, because `LEVER-M257-L5-setdress`'s remedy is
entirely different for each.

## Phase plan (two planned lines — a deliberate two-level shape, declared here)

- **Phase A — the retroactive level (buildbench).** The sub-phase boundaries **already exist as log
  lines** in iter-08's three captured reps. Derive nested `set_dress` sub-phases and prove them against
  **real** logs, not fixtures: each attributes separately **and** Σ sub-phases == parent exactly.
- **Phase B — the level below (stacksnap).** `replay.Run` currently emits nothing between start and
  finish. Instrument its five documented phases (verify / clear / copy / reindex / advance-sequences)
  with per-table detail and an **explicit unattributed residual**, and print them. This is the number
  L5 needs and no captured log can supply.
- **Phase C — arm, do not poll.** Re-pin + re-tag so the campaign consumes the new instrument, then arm
  `autoarm-campaign.sh`. `TOK-01` step 4 stays live without a human in the trigger.

## Expected lift

**No movement on the composed p50 is expected or claimed** — this is instrument work, and an instrument
that changed the number it measures would be the defect. The measurable deliverable is **attribution**:
the 252.73 s span resolved into named, separately-attributing sub-phases whose sum is exact, and the
copy-vs-reindex split made emittable. Success criterion: **L5 is aimable** — a reader can name which
operation to attack and with what expected ceiling.

## Escalation conditions

- Sub-phases that do **not** sum to their parent → the split is wrong; do not ship it (this is the
  whole point of the iter, and `D17` is the precedent for shipping a table that sums but misattributes).
- A `stacksnap` change that alters replay **behaviour** rather than only observing it → revert; the
  instrument must not touch the thing measured.
- Any need to tear down / re-seed `demo-2` or the dev stack → stop, escalate. Never.

## Acceptable close-no-lift outcomes

- The split lands, sums, and shows the cost is **irreducibly** one operation with no internal structure
  worth aiming at (i.e. L5 has no purchase). That is a falsification with a number behind it and closes
  the routed item honestly.
- The host never opens a window and clause 3 stays NOT MET. Expected; `budget-exhausted`, not a blocker.
