# M258 iter-03 — progress

**Type:** tik · **Active strategy:** `TOK-01` step 1 — *measure the composition before engineering it.*

## Phase 0 — re-survey refuted the routed diagnosis (details in `overview.md` § Step 0)

Three inherited claims tested against code; **all three false**:

| inherited claim | verdict |
|---|---|
| the UI images carry a **foreign / real-Clerk** key | **REFUTED** — `pk_test_bWFy…` decodes to `marcos-mac-mini.taildc510.ts.net:15400`, a **Clerkenstein-minted** key for this box (`mint_pk` = `pk_test_ + base64(host+"$")`) |
| `demo-1` is `--no-public-host` / localhost-bound (iter-02 `D2`/`D7`) | **REFUTED** — `cycle.log:26` *"public-host AUTO-DISCOVERED — marcos-mac-mini.taildc510.ts.net"*; `docker port` shows `0.0.0.0` |
| the `2>/dev/null \|\| true` **swallow** is why 24 blocks accumulated | **REFUTED as the cause** — `inject.py` never warns about accumulation, so there was no message to swallow. `\|\| true` is *deliberate* (`:2033-2035`) and guarded by a fail-loud `[ -n "$PK_DEMO" ]` on the next line |

**The real chain:** `inject.py` appends → 24 blocks → `_stack_minted_pk` reads **first-wins** while every
consumer reads **last-wins** → ISOLATION compares this run's images against an **older bring-up's key** →
false RED. And separately, `buildbench` **cannot express `--no-public-host`**, so the campaign silently
ran the one mode in which the Playthrough batch **cannot be driven from this host at all**.

## Phase A — P1–P4 landed (rext `d0051e7`, tag `fast-build-m258-iter-03`, **on origin**)

| # | fix | file |
|---|---|---|
| P1 | strip-then-append, idempotent, atomic (`os.replace`) | `stack-injection/inject.py` (`write_injection_block`) |
| P2 | dotenv sources read **last-wins**, matched on key name; JSON keeps first-match | `stack-core/buildbench.py` (`_stack_minted_pk`) |
| P3 | `--no-public-host` passthrough + mutual exclusion refused **before** any teardown + `bringup_argv` in every rep ledger | `stack-core/buildbench.py` |
| P4 | stop discarding `inject.py`'s stderr (**keep** `\|\| true` — it is deliberate) | `demo-stack/up-injected.sh:2036` |

**Each proven RED with its precondition absent**, against a *faithful* revert of the original (append-only
body restored verbatim, not a stub):

| fix | RED evidence |
|---|---|
| P1 | 2 failures — block stacks to 2, file grows per call |
| P2 | 2 failures — returns the first key; **the end-to-end case reproduces the exact false RED** |
| P3 | 2 errors + 1 failure — flag unrecognised, `bringup_argv` absent |

## Phase B — test gate

| suite | result |
|---|---|
| `stack-injection` (7 modules) | **341 OK** (8 skipped) |
| `stack-core` — `test_buildbench` + `test_isolation_assert_m257` | **181 OK** |
| `stack-core` doc-facing guards (4 modules) | 273 run, **271 OK**, 2 failures — **pre-existing, proven** |

⚠️ **The 2 failures are NOT mine and NOT new.** `test_decommissioned_instruction_guard` reports
`demo-stack/stacks/demo-1/clones/app/...` — files inside a **live demo's ephemeral platform clone**, a
path `demo-stack/.gitignore:8` ignores. Proven pre-existing by running the same guard in the **pristine
consumption clone at `fast-build-m257-close`** (zero changes of mine): **identical 2 failures**. The guard
scans its own workspace's scratch — *an instrument that lives inside its own subject measures itself*.
Routed, not absorbed.

⚠️ **The full `stack-core` sweep did not complete** (>9 min, no output). That is the same condition M257's
close recorded in commit `d9608455` — *"the stack-core sweep did not complete — measured three times, in
the same place."* Not re-diagnosed here; the proportionate gate above covers this iter's diff.

## Phase C — LIVE proof without a rebuild, and the measurement that could not be taken

### The false RED is gone, proven on the very stack that produced it

Replicating buildbench's **exact live wiring** (`buildbench.py:1573-1576`: `_image_sizes` / `_image_env` /
`_image_bundle_pks`) against the still-running `demo-1`:

| | iter-02 | iter-03 (same stack, same images, no rebuild) |
|---|---|---|
| images checked | 8 | 8 |
| `own_pk_fingerprint` | `pk_test_MTI3…61fbfaf4` (an **older** bring-up's) | **`pk_test_bWFy…52038877`** |
| ISOLATION | **FAIL**, 3 × `foreign_pk` | **`ok: True`, 0 failures**, `foreign_pks {}`, `foreign_origins {}` |

The reader now returns **exactly the key iter-02 recorded the images as carrying**. So the images were
never foreign, the campaign was never dirty, and **no demo ever reached production auth.**

### P1 on the real 128-line artifact (replayed on a copy — nothing live mutated)

`128 lines / 24 blocks` → **`37 lines / 1 block`**; re-run **byte-identical** (idempotent); 4 `DESK_CLERK_`
lines preserved; base lines preserved; final pk re-minted for `127.0.0.1:15400`.

> Honest residue: the 32 retained lines include ~23 blank separators the old blocks left. Cosmetic, on a
> first migration only — interior blanks are not stripped because they may be another owner's spacing, and
> the file **no longer grows**, which is the actual requirement (proven byte-identical).

### The batch half — STILL NOT MEASURED, and this time the blocker is the host

`assert-headroom --profile macmini` (host **named**, per the routed item) → **FAIL**:

```
peak_load1: 20.31 exceeded cores-2 (10) on a 12-core host — the run was CPU-bound,
so its phase timings are contended and not comparable
```

At the decision point `load1` was **39.05 / 45.94** against 12 cores, and the top consumers were **not
mine** — `mds_stores` (Spotlight) at 183 %, `a8-cart-runner` from the user's own `anima8` project, and
node processes; my interpreter is 3.14, the hot Python was 3.12. I stopped the one contender I *had*
created (the background census sweep) and verified it gone; the load kept climbing.

**No timing was taken, deliberately.** Three reasons, and the first is sufficient:

1. **The instrument refuses.** `D-M255-1`: *a gate number measured on a host without headroom is not a
   number.* A rep that fails headroom is `not gateable` — running it produces no usable figure anyway.
2. **A bring-up is three parallel Docker build lanes.** Launching that onto a box the user is actively
   working on degrades their session to buy a number the harness would then discard.
3. `TOK-01`'s rule, and iter-02's `D5`: **a refusal to measure is a result; a meaningless measurement is a
   liability.**

**Booleans survive contention and were taken** (above). **Timings do not, and were not.**

## Metric

| | |
|---|---|
| batch half | **still unmeasured** — host contention, not a defect |
| its blocker | **REMOVED and live-proven**: ISOLATION green on the stack that reded, `--no-public-host` now expressible |
| bring-up half | not re-taken (no headroom) — the n=1 **395.31 s** remains, and is now known to be a **public-host** figure |

**Metric delta 0 on the primary unknown; the precondition for it moved from "blocked by a defect" to
"blocked by someone else's CPU."** Those are different states and are recorded as such.

> **A correction the 395.31 s number needs.** It was taken in **public-host** mode, which pays a
> `tailscale serve` + cert-mint leg (`cycle.log:143-144`, a real cert re-mint) that M257's **286.99 s**
> `--public-host billion…` p50 also paid — so mode alone does *not* explain the +108.32 s.
> `CHECK-M258-iter02-studio-desk-is-the-untouched-leg` stays the live suspect. **Stated as an open
> attribution, not resolved** — the honest answer needs n≥3 in a named mode.

## Scope-creep tripwire

Lines opened: (1) the refuted diagnosis, (2) the first/last-wins reader, (3) the host-mode passthrough,
(4) the stderr swallow. All four are the **planned multi-step shape** declared in `overview.md` — the
carve-out case. The **guard failure** found in Phase B was the 3rd *unplanned* line and was **routed, not
absorbed** (proven pre-existing in ~2 min, then dropped).

## Close — 2026-08-12

**Outcome:** The routed diagnosis was refuted in all three of its claims, the real four-part chain was
found, fixed, published to origin and **live-proven on the stack that produced the failure** — ISOLATION
went `FAIL (3× foreign_pk)` → **`ok: True, 0 failures`** with no rebuild, and the 24-block env file
collapses to 1 idempotently. **No demo was ever wired to production auth.** The batch half remains
unmeasured: its defect-blocker is gone, but the host was under sustained third-party load (`load1` 39–46
vs 12 cores) and the headroom gate correctly refuses to produce a number there.
**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n *(this iter is a tik; 2 no-prog tiks, not 3)* —
(3) re-scope: n *(no composed p50 exists to fire the 600 s valve)* — (4) user-blocker: **n** *(no decision
is owed: the fixes were unambiguous, and host contention is an environmental condition, not a fork —
Phase 5 §4's NOT-list)* — (5) cap-reached: n *(1 tik)* — (6) protocol-stop: n — (7) budget-exhausted: n —
Outcome: **continue**
**Decisions:** D8, D9, D10 (this iter's `decisions.md`)
**Side-deliverables:** `bringup_argv` in every rep ledger (P3's companion — the mode a number was taken in
is now readable from the record, not recoverable by grepping `cycle.log`); `build-budget.md` documents the
host-mode flag.
**Routes carried forward:**

- **`MEASURE-M258-batch-half`** → **iter-04**, unchanged as `TOK-01` step 1's outstanding deliverable. The
  defect blocker is discharged; what it now needs is headroom. Run
  `buildbench run 1 --reps 1 --profile macmini --no-public-host` then the batch, and **state `load1` with
  both numbers**.
- **`FIX-M258-iter03-guard-scans-its-own-scratch`** → iter-04 or later.
  `test_decommissioned_instruction_guard` walks `demo-stack/stacks/**` — gitignored ephemeral demo clones
  — and reports the *platform's* source as a rext named-consumer list. **Proven pre-existing** at
  `fast-build-m257-close`. Fires on any box with a demo stack dir present, which is every box that has
  ever run this milestone. Likely the same root as M257's *"the stack-core sweep did not complete."*
- **`CHECK-M258-iter02-studio-desk-is-the-untouched-leg`** → carried, **unresolved and now sharper**: mode
  does not explain the +108.32 s, so `ui_studio_desk` (115.35 s, the leg L1 never touched) stays the
  suspect. n≥3 in a named mode before any claim.
- **`ROUTE-M258-iter02-isolation-names-two-causes-not-three`** → **still open, and now three-for-three**:
  the message names two causes, both refuted at iter-02, and the true cause (a stale own_pk from a
  first-wins read) is a **third**. The text should name the reader as a candidate.
- **`ROUTE-M258-iter02-headroom-defaults-to-billion`** → still open (bare `assert-headroom` grades against
  `billion.json`; every invocation here named `macmini` explicitly).
- **`ROUTE-M258-iter02-purge-did-not-clear-the-stack-dir`** → still open. P1 makes the *symptom* benign
  (the file no longer grows), but `--purge` still leaves the stack dir — `verification.md` **F-9**.

**Lessons:**

- **Re-verify an inherited diagnosis before implementing it — including one written by the milestone's own
  previous iter.** All three of its claims were false, and one ("wired to a real Clerk app") would have
  sent a fixer at the `dockerignore`/build path, which is real, documented, and *not this defect*. The
  decisive check cost one line: `base64 -d` on a key that is public by design.
- **A false RED is not the safe direction.** The clause's own docstring already says *wrong-and-loud is
  not fail-closed — it is a fence that cannot be believed*, about fingerprints, one line from the
  ordering bug that produced exactly that outcome. **A principle written next to the code does not apply
  itself.**
- **"No flag" is not "the default you assume."** `up-injected.sh` auto-discovers by default, so passing
  neither host flag selected the *opposite* of the mode `TOK-01` declared it was gating — and nothing in
  the campaign record said so. **Record the argv with the number.**
- **When the reader and the writer disagree about ordering, the file is not the bug — the disagreement
  is.** Two writers of one dotenv, one idempotent and one appending, and a reader taking neither
  convention.
- **A boolean survives contention; a timing does not.** Under `load1` 39–46 the live ISOLATION verdict was
  still worth taking, and the wall-clock was not. Knowing which of your deliverables is which is what lets
  a contended host cost you one of them instead of both.
