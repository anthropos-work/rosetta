# M258 iter-04 — progress

**Type:** tik · **Active strategy:** `TOK-01` step 1 — *measure the composition before engineering it.*

## Phase A — the headroom gate REFUSED, and the refusal is recorded as the result

`assert-headroom --profile macmini` (host **named**, never bare) failed `peak_load1` on every check:

| time (UTC) | `load1` | vs floor (cores−2 = 10) |
|---|---|---|
| 05:27 | 20.31 | FAIL |
| 05:28 | 45.94 | FAIL |
| 05:34 (waiter cap) | 18.48 | FAIL |
| 05:36 | 14.77 | FAIL |
| 05:37 | 13.05 → 11.39 | FAIL, decaying |
| 05:38 | **16.59** | FAIL — **bounced back up** |

The load is **third-party and bursty, not decaying**: `mds_stores` (Spotlight) at 183 %, the user's own
`a8-cart-runner` (`workspace/hyperspace/anima8`), and two `node` processes at ~100 % each. The hot Python
was **3.12**; mine is **3.14**. The one contender I had created — a background `stack-core` census sweep
— I stopped and verified gone, after which load *rose*.

Two bounded waiters were armed and both hit their cap. **A third was not armed**: this is the
*permanently contended* host the release documents, and *"do not wait for quiet"* is the standing rule.

**Consequence, stated plainly:** `buildbench run` **cannot execute a cycle at all** under this load —
`D-M255-1` makes the pre-rep assert abort (rc 1) before rep 1. So **no gate-quality timing was
obtainable today, by the instrument's own design.** That is a result, not an omission.

## Phase B — the cycle was run anyway, as an OPERATOR, for the halves contention cannot corrupt

`up-injected.sh`'s own pre-flights are **advisory by design** (`buildbench.py:36` — *"never block a
genuinely good bring-up on a soft signal"*), and the gate instrument's refusal is about **quoting a
number**, not about doing work. So the cold cycle was driven directly:

```
rosetta-demo down 1 --purge      → DOWN_RC=0
up-injected.sh 1 --no-public-host
```

launched 05:39:06Z, `load1` 16 → 62 during the purge+build.

**Every one of P1–P4 is now confirmed on a real bring-up, not only in unit tests:**

| fix | live evidence (from `/tmp/m258_up.log`, this run) |
|---|---|
| **P3** | `hostlock: acquired demo-1 [up-injected demo-1 (**localhost**) …]` — and **no** `public-host AUTO-DISCOVERED` line. iter-02's same line read `(marcos-mac-mini.taildc510.ts.net)`. The flag reaches the bring-up and single-box mode engages |
| **P4** | `:117  clerk-frontend: minted pk_test_MTI3LjAuMC4xOjE1NDAwJA (host=127.0.0.1:15400, round-trip OK)` — **this line is in the log at all**, which is the whole of the fix; `2>/dev/null` discarded it before |
| **P1** | `.env.demo-1` injection blocks **24 → 1** on the live stack, at the moment `inject.py` ran |
| **P2** | with exactly one block, first-wins and last-wins **agree by construction**; the minted key is the loopback `pk_test_MTI3…`, which is what the images will bake |

### F-9 re-confirmed live, and it is why P1 was load-bearing

`--purge` completed (`DOWN_RC=0`, *"demo-1: data purged"*, images removed, network removed) and the
**24-block `.env.demo-1` survived it intact** — measured at 05:40:30Z, after the purge, before
`inject.py` ran. `ROUTE-M258-iter02-purge-did-not-clear-the-stack-dir` is not a hypothesis;
the file outlived a full purge on this very cycle. Without P1 the 25th block would have been appended
here.

### Isolation of the user's stacks — verified mid-cycle

`demo-2` **11 containers** and the dev stack **5 containers**, resident and untouched while `demo-1` was
at 0 (`down --purge` names its scope: *"project demo-1 ONLY — dev stack untouched"*).

## Phase C — THE BATCH HALF, MEASURED FOR THE FIRST TIME

`run-playthroughs.sh 1 --reset`, launched 05:52:40Z at `load1` **30.38**, on the single-box stack.

| | |
|---|---|
| **wall-clock** | **129 s** (`BATCH_SECONDS=129`; the runner's own `215 passed (2.0m)`) |
| specs | `Running 215 tests using 1 worker` → **215 passed**, 0 failed |
| ptreport four-state | **`passing=30  failing=0  unimplemented=1  unimplementable=0`** — 30/31 (96.8 %) |
| **red set** | **EMPTY** |
| `BATCH_RC` | **0** |
| reset-to-seed | ran — the real `stackseed --reset` FK-ordered TRUNCATE, not an additive re-seed |

**It was checked for the green-without-checking failure this release exists to refuse.** 129 s is fast,
so the first question asked was whether the suite *ran*: it reports its own worker count and spec count
(`215 … using 1 worker`, serial per `workers:1`), the reset printed its per-table truncations, and
ptreport reconciled **31 manifest use cases**. The single `unimplemented` is the **known declared TODO**
(`onboarding.enterprise-workforce-standard.UC1`, carrying the machine-checked `will-not-build` verdict),
which is exactly the shipped state `playthroughs.md` records: **30 live Playthroughs + 1 verdicted TODO**.

⚠️ **This is n=1 and MUST NOT be quoted as a p50.** `C2` — M256 measured a **2.04× spread with no trend**
over six full-suite runs and escalated that this suite's timing is *not decidable at n=3 on this host*.
One sample has no spread to publish. **And it is not comparable to M256's 56.6 s**, which priced **18**
specs; this run priced **215**.

### The composed arithmetic — and why it is NOT a gate reading

| half | this cycle | gate-quality? |
|---|---|---|
| bring-up (`down --purge` + `up --no-public-host`) | **781 s** | ✗ contended (`load1` 16 → 62), and the instrument refused to run at all |
| batch | **129 s** | ✗ contended (`load1` 30.4 at launch) |
| **composed** | **910 s** | ✗ — **against a 480 s ceiling** |

**910 s does not fire the `re_scope_trigger`, and must not be reported as if it did.** That trigger reads
*"the composed **p50** exceeds 600 s after 3 tiks"*; this is a single contended sample, taken in the one
condition `D-M255-1` says is not a measurement. Reporting it as a gate miss would be the precise error
the headroom clause exists to prevent.

**What it does support is a shape**, and it is the useful finding: **the batch half is SMALL.** At 129 s
contended it is ~14 % of this cycle, so the 480 s ceiling is dominated almost entirely by the **bring-up**
half. Against M257's proven gateable bring-up p50 of **286.99 s**, a composition of ~287 + ~129 ≈ **416 s**
would sit **inside** 480 s — which is the first evidence the ceiling is reachable at all. **Stated as an
inference from two non-comparable numbers, not as a result.**

### ISOLATION on the newly built images — green, and now green for the right reason

| | |
|---|---|
| images | 8 |
| `own_pk` | `pk_test_MTI3…61fbfaf4` → decodes to **`127.0.0.1:15400`** |
| verdict | **`ok`, 0 failures** |

That fingerprint is **exactly the `own_pk` iter-02 expected and could not match**. The stack now mints,
bakes and asserts one key, and `.env.demo-1` holds **one** block (37 lines) — the offline replay's
prediction, confirmed on the live file after a real bring-up.

## Close — 2026-08-12

**Outcome:** **The batch half exists.** 129 s, 215 specs, **30/30 Playthroughs passing, red set EMPTY**,
`BATCH_RC=0` — `TOK-01` step 1's outstanding deliverable, discharged after two iters blocked on it. All
four iter-03 fixes were additionally confirmed on a real bring-up (single-box mode engaged, the minted-host
line visible in the log, `.env.demo-1` 24 → 1, ISOLATION green on fresh images). Both halves were taken
under third-party load and **neither is a gate number**; the headroom gate refused a `buildbench` cycle
outright, so the cycle was driven as an operator for the parts contention cannot corrupt.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: **n** *(the 600 s trigger reads
a **p50 after 3 tiks**; this is one contended sample and firing on it would be a category error)* —
(4) user-blocker: n — (5) cap-reached: n *(2 tiks)* — (6) protocol-stop: n — (7) budget-exhausted: n —
Outcome: **continue**
**Decisions:** D11, D12 (this iter's `decisions.md`)
**Side-deliverables:** none.
**Routes carried forward:**

- **`MEASURE-M258-gateable-composition`** → **iter-05**, and it is now *cheap*: `load1` fell to **6.17**
  immediately after the batch (the user's workload ended), which is **inside** the cores−2 floor of 10.
  Run the real `buildbench run 1 --reps 1 --profile macmini --no-public-host` for a **gateable**
  bring-up half, then the batch again — giving a second batch sample toward the spread `C2` demands.
- **`RESTORE-M258-world-contract`** (`TOK-01` step 3) — **now owed in fact, not in principle.** The
  batch's `--reset` TRUNCATEd the demo world and re-seeded **pt-world**, so `demo-1` is currently a
  Playthrough world behind a cockpit projected from the stories preset. This is exactly the state
  `overview.md` § *The world contract* describes, and resolution **(b) restore after** is the decided fix.
- Unchanged and still open: `FIX-M258-iter03-guard-scans-its-own-scratch` ·
  `CHECK-M258-iter02-studio-desk-is-the-untouched-leg` ·
  `ROUTE-M258-iter02-isolation-names-two-causes-not-three` ·
  `ROUTE-M258-iter02-headroom-defaults-to-billion`.
- **`ROUTE-M258-iter02-purge-did-not-clear-the-stack-dir`** — **upgraded from routed to OBSERVED**: the
  24-block file survived a full `--purge` on this cycle (measured after `DOWN_RC=0`, before `inject.py`).

**Lessons:**

- **The instrument refusing to measure is not the same as the work being impossible.** `buildbench` could
  not run a cycle under load, but `up-injected.sh` is an *operator* whose pre-flights are advisory by
  design. Separating "may I quote a number" from "may I do the work" turned a blocked iter into the one
  that produced the milestone's missing measurement.
- **Ask whether a fast green actually ran.** 129 s for a suite whose only prior figure was 56.6 s/18 specs
  invites acceptance. The spec count, worker count, per-table truncations and a 31-case reconciliation are
  what make it a result rather than a hope.
- **A single contended sample is not a trigger.** 910 s over a 480 s ceiling looks like a `re_scope_trigger`
  and is not one — the trigger is defined on a p50 over 3 tiks. **Read the trigger's own definition before
  firing it**, or a noisy sample renegotiates a gate that was never missed.
- **Two blocked iters can still have been the cheap path.** iters 02 and 03 produced no batch number, but
  the blocker they removed was real and the measurement took **129 s** once it was gone.
