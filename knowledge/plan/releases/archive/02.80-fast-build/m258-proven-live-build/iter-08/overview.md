---
iteration_type: tik
status: closed-fixed-partial
iter: 08
milestone: M258
created: 2026-08-12
---

# M258 iter-08 — the composed 3× cold campaign, fired by an auto-arming waiter

**Active strategy reference:** `TOK-01` step **4** — *the composed 3× cold campaign against the gate,
with the spread published beside the p50.* Steps 1–3 discharged (bring-up 247.79 s, batch 129/160 s,
gate wired at `up-injected.sh:2844`, restore leg landed).

## Cluster / target identified

`TOK-01`'s next-tik direction and iter-07's carry-forward name the same single target: **gate clause 3**,
the only unproven clause. Clauses 1, 2, 4 and 5 are booleans and all five were graded at iter-07 —
4 ✅, 1 ⬜. Every precondition is discharged: the tag is on origin, the consumption clone is re-pinned
and **verified this iter to carry the feature under test** (`batch-gate.sh` present, hook at `:2844` —
the `D21` check, re-run rather than inherited), `postgres-schemas` is proven satisfiable by `D27`
arithmetic, and the campaign is scripted with a headroom-before-teardown guard.

**Re-survey (Phase 1 Step 0).** The target is unchanged and still meaningful — no iter has absorbed it.
What HAS changed since iter-07 closed is the **contention source**, and the change is material enough to
record: iter-07 named Spotlight (`mdworker_shared`/`mds`) + two `a8-cart-runner` processes. At iter-08
open the mix is **`anima8` `m270a2-iter132/wt/shifttrap/lab/st_pin3.py` at ~89 % CPU** plus **31 node
processes across two `hyper-studio-worktrees/` checkouts** (`room-parity`, `dev--open-sess`) — a
*different, larger* third-party load. The standing rule holds: none of it is mine to stop.

## Hypothesis

**The host's calm windows are real but short, and iter-07 could only catch one by coinciding with a
manual poll.** iter-07 polled by hand for ~30 minutes at ~2-minute granularity and recorded a minimum of
11.93 against a limit of 10 — but a window *was* independently observed at 08:11:20Z (three consecutive
`load1` samples under 5.0, decaying from 13.29). A window that opens and closes between two manual polls
is indistinguishable from no window at all.

So the fix is not more patience at the same granularity — it is to **remove the human from the trigger**:
arm a watcher that samples continuously and fires the campaign on the first sustained dip, whenever it
comes. A false start is already cheap by construction (`launch-iter07-campaign.sh` asserts headroom
*before* the teardown and exits 8 leaving `demo-1` intact), so the arming threshold can be tuned for
*catching* windows rather than for avoiding waste.

## Expected lift

**Clause 3 measured** — a composed p50 over 3 consecutive cold reps, with its **spread published beside
it** (`TOK-01`'s standing requirement, from M256's 2.04×-spread escalation), `load1` and the environment
stated with every figure. Projection from the separately-measured halves is **414.15 s vs the 480 s
ceiling**; the number is to be *reported*, not forced.

**Acceptable close-no-lift outcome:** if the host never offers a window, the iter closes with the
waiter's **measured** load time-series as the falsification — the same honest refusal iter-07 produced,
but with the trigger automated so the next window cannot be missed by absence.

## Phase plan

- **Phase A — pre-flight.** Verify the pin carries the feature (`D21`); verify the user's stacks resident
  (`demo-2`=11, dev=5); verify `demo-1` still presenter-usable. *(A partially complete at iter open.)*
- **Phase B — arm.** Build and launch the auto-arming waiter: continuous `load1` sampling, fire on a
  sustained dip, hand off to the already-validated campaign script.
- **Phase C — poll.** One waiter, ≥5 min interval, terminating condition + process-died branch,
  heartbeat every rep.
- **Phase D — measure.** If the campaign runs: compose the p50 **and the spread**, verify `batch_gate`
  attributes separately from `autoverify` in every rep (`D17` — *a table that adds up is not a table that
  attributes*), and check the outputs agree **with each other** (`D19`/`D20`).
- **Phase E — restore + close.** Verify the user's stacks resident after; verify `demo-1` left
  presenter-usable (12/12 cockpit seats in the 35-identity roster).

## Escalation conditions

- **Composed p50 > 600 s after 3 tiks** → the milestone's declared `re_scope_trigger` (split into a fast
  smoke lane + a full lane after) — surfaced **with measurements**, never as a question.
- **Non-empty red set** → escalates to the user as a **product verdict** per `D-v28-3` clause 4. Do
  **not** retry it (`retries: 0` is contract) and do **not** attribute it to load without evidence —
  record `load1` at the moment it happened.
- **No window within the session's budget** → close honestly with the time-series; the constraint is the
  host, and it is external.

## Acceptable close-no-lift outcomes

- The instrument refuses across the whole session and the campaign never fires. **A HEADROOM refusal is a
  RESULT** — provided the refusal is *measured* (a time-series, not a recollection) and the trigger was
  armed the whole time rather than sampled by hand.
