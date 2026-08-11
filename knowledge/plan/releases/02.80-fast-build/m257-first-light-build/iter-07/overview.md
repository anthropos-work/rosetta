---
iter: 7
milestone: M257
iteration_type: tik
status: in-progress
opened: 2026-08-11
---

# iter-07 — `BASELINE-M257-macmini-n3`: the number this milestone has never had

**Type:** tik · **Active strategy:** [`TOK-02`](../decisions.md) — step 3, *take the baseline on the
contended box, and label it*.

## Step 0 — re-survey

- `macmini.json` still has **no `gated_baseline`** — the field the gate is a percentage of is still empty.
- The gate now names this host (iter-05), so a number measured here is finally a number the gate can read.
- Clause 1 now grades the machine the sample came from (iter-06), so a refusal from it is now *readable*.

Both of `TOK-02`'s preconditions are discharged. The target is unchanged and is now the only thing between
this milestone and pricing levers.

## Cluster / target identified

**There is no cold-cycle p50 for any host this project may still measure on.** `billion.json` carries the
only `gated_baseline` in the repo (666.29 s, n=3, x86_64/containerd, demo-only since `D-v28-14`), and
`overview.md`'s own rule is that a wall-clock never transfers between hosts. Everything M257 has said about
its own distance to the gate — including iter-04's ~420–455 s — is a **scaling of billion's phase table**,
not a measurement.

## Hypothesis

A cold `demo-down --purge` + `demo-up` cycle, repeated 3×, produces a p50 on this host. Whatever it is, it
replaces an estimate with a measurement and fills `gated_baseline`.

## The environment, stated up front because it is part of the result

This box is a **permanently contended developer workstation** and cannot be freed — `TOK-02` step 3 exists
because waiting for quiet is waiting forever. Observed load1 while opening this iter ranged **1.96 → 23.48**
(the 23.48 driven by the user's own concurrent `pytest` run in another repo, top process at 792 % CPU).

**Nothing of the user's is disturbed**, and that is a measured decision, not a hope:

| | measured at open |
|---|---|
| slot used | **`demo-1`** — registered, container-less, free (`demo-2` and the 5-container dev stack untouched) |
| VM memory | `MemAvailable` **9.87 GiB** of 11.67 GiB; all 16 running containers total **~1.43 GiB** |
| projected peak | ~1.43 + one 3.1 GiB build lane + ~1.3 GiB for demo-1's own containers ≈ **5.9 GiB** |
| VM disk | **51.6 GiB** free vs `disk_floor` 7 + `projected_image` 15 = **22 GiB** |
| identity gate | `profile_describes_host` → **`match`** (host 12 cores / engine NCPU 8 — the exact split iter-06 fixed, now recorded in the campaign artefacts) |

So the dev stack did **not** need to come down, and it did not.

## Expected lift

**The gate metric gets its first value on this host.** Whether that value passes 360 s is not this iter's
to decide — the iter's deliverable is a measured, labelled number plus the reps' own headroom verdicts.

## Escalation conditions

- A **pre-rep** refusal (disk) aborts the campaign by contract (`D-M255-1`) — report it, do not override.
- A **post-rep** clause-1 refusal does **not** abort: it is recorded per rep. **Report it as a RESULT**,
  with what the run would have measured — never as a failure to measure.
- If a rep fails for a *platform* reason (a broken bring-up rather than a contended one), that is a
  different finding and gets its own line.

## Acceptable close-no-lift outcomes

If every rep is refused by its own headroom clause, the iter still closes on a real deliverable: the
refusal, the loads that caused it, and the raw cycle times the reps would have contributed. That is the
`laptop.json` situation made *legible* instead of left as an absence — that profile records a refusal at
load1 10.69 and **no cycle number at all**, which is the outcome this iter exists to avoid repeating.
