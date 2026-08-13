---
milestone: M258
iter: 2
iteration_type: tik
status: closed-fixed-partial
created: 2026-08-12
---

# M258 iter-02 — tik

**Active strategy reference:** `TOK-01` (milestone-root [`decisions.md`](../decisions.md)) —
*measure the composition before engineering it.* This is step 1 of its four ordered tiks.

## Step 0 — re-survey before targeting (mandatory)

`TOK-01`'s next-tik direction was authored 12 minutes ago, so staleness is unlikely — but the re-survey
is the discipline, and it returned **three facts the strategy did not have**:

| checked | result |
|---|---|
| is the pin still stale? | **yes** — `rosetta/.agentspace/rext.tag` = `fast-build-m257-iter-09`, and the consumption clone `stack-demo/rosetta-extensions` is checked out at that same tag (`8956e69`). `R0` stands. |
| is `demo-1`'s slot free? | **yes** — `docker ps -a` matches **0** `demo-1-*` containers. |
| is the batch half still unmeasured? | **yes** — no corpus doc or milestone record publishes a wall-clock for the shipped 30-Playthrough suite. |

**New — and it sharpens `R0` rather than replacing it: the pin exists in THREE copies with THREE
different values.**

| copy | value |
|---|---|
| `rosetta/.agentspace/rext.tag` (canonical, M49 #1) | `fast-build-m257-iter-09` |
| `stack-demo/rosetta-extensions` checkout | `fast-build-m257-iter-09` (`8956e69`) |
| `stack-demo/rosetta-extensions/.agentspace/rext.tag` (untracked stray) | **`fast-build-m257x-iter-279`** |

The third is an untracked file *inside the consumption clone* naming a tag from the **previous
milestone**. It is very likely inert — the canonical pin is rosetta's — but "a pin that exists in three
copies" is the same shape as the stale-verdict family this release keeps paying for, and it is
**cheaper to check than to be surprised by**. Verify which copy the bring-up actually reads before
trusting the re-pin.

**Also observed:** the **authoring copy's** `demo-stack/stacks/registry.json` still lists `demo-1` as
live with offset 10000, while **zero** `demo-1-*` containers exist — a registry record outliving its
subject. The consumption clone's registry is `{}`. Recorded, not acted on: teardown is the sanctioned
way to reclaim, and this tik runs one.

**No substitution.** `TOK-01`'s named target is intact and still the right next thing.

## Target

The composition's **second half**, which has never been measured, plus the first half re-measured at a
corrected pin.

## Hypothesis

One cold cycle on `demo-1` — `demo-down --purge` + `demo-up --no-public-host`, then the full
Playthrough batch — yields (i) a bring-up figure at the corrected pin, (ii) **the first wall-clock the
batch half has ever had**, (iii) the restore-leg cost, and (iv) a composed figure against the 480 s
ceiling. That converts `TOK-01`'s "1–2 iters" from a hope into a prediction.

## Expected lift

**Metric delta 0 by design** — this tik measures, it does not optimise. Its deliverable is a number,
not a second saved. Success criterion: a *complete* cycle whose bring-up half is `rc=0` with a green
`autoverify.json`, and a batch that ran to completion with its red set enumerated.

## Phase plan

1. **Re-pin** (`R0`): `rosetta/.agentspace/rext.tag` → `fast-build-m257-close`; re-point the
   consumption clone; prove all three copies agree or prove the strays inert.
2. **Pre-flight** (Phase 0d, light): the harness is invokable and the gate instrument is the
   **post-close** one before a ~10-minute campaign is spent on it.
3. **Cold cycle** on `demo-1`, foreground-polled — never background-and-yield (the documented stall
   trap, `overview.md` § Iteration protocol).
4. **Batch**, run to completion; enumerate the red set rather than halting on it (`D-v28-3`).
5. **Record** every figure with `load1` and the environment.

## Escalation conditions

- **Route forward, do not escalate:** individual Playthrough reds (that is the batch gate's designed
  output), the content-stories sweep's `exit 2` (`F1` — it gates a *different* sweep and must not read
  as a batch blocker), any inherited item found already-closed.
- **Escalate:** the bring-up cannot reach `rc=0` at the corrected pin (that would make the composed
  gate ungradeable — M257 iter-05's lesson: *check a gate for gradeability before checking it for
  satisfaction*).

## Acceptable close-no-lift outcomes

A cycle that fails to complete still closes this tik **if** it names how far it got and why —
per `TOK-01`, a measurement on a contended box, labelled, beats waiting for a quiet one, and a
HEADROOM refusal is a **result**, not a failure to measure.
