---
iter: 260
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-10T14:23:27Z
---

# iter-260 — clause 1 under its LITERAL unit: three consecutive `--purge` + `up` cycles

**Type:** tik
**Active strategy:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07),
under the user's binding `D-M257x-256-1`.

## Step 0 — Re-survey before targeting

iter-258 discharged `ROUTE-M257x-256-the-advance-is-unproven` for the **demo** half; iter-259 escalated the
**dev** half to the user, where it remains. Neither of the two lines that do not depend on that answer —
clause 1's remaining cycles, and clause 2's Playthrough suite — has moved.

The re-survey corrected one briefed fact before targeting, and it enlarges the work rather than shrinking it:

> **iter-258's cycle is not a clause-1 cycle.** Clause 1's unit is written as *"a cold `demo-down --purge` +
> `demo-up` … reaches `autoverify green:true / 0 warnings` across **3 consecutive cycles**."* iter-258 ran the
> **bare** `up-injected.sh 2` into a slot whose containers had never existed — a fresh-slot `up` with **no
> preceding purge** — and its own close says so (*"it was a fresh-slot bring-up rather than a `--purge` re-use
> cycle"*). The `up` half is proven; the **`down --purge` half has never run in this milestone's evidence at
> all**, and it is half of the clause's named unit.

So the count is **not** "1 of 3 with 2 to go." Under the clause's literal unit it is **0 of 3**, and the
briefing's "exactly 1 has happened" is the lenient reading. This iter takes the strict one, because the
alternative is to adjudicate a gate clause in the direction that happens to be cheaper — the exact move this
milestone has caught repeatedly.

**Measured preconditions** (2026-08-10T14:23:27Z): disk **196 GiB** free; Docker VM **11.67 GiB** / 8 CPU;
load averages **5.79 / 7.13 / 8.26** — the host is CONTENDED and third-party. `demo-1` is up 4 days with 11
containers and is **out of scope under every outcome**. `demo-2` is up and green from iter-258
(`ts 2026-08-10T14:04:49Z`). Host-native listeners observed before any teardown: demo-2 cockpit `:27700`
(pid 43878), demo-2 ant-academy `:23077` (pid 63229), demo-1 cockpit `:17700` (pid 75363).

## Cluster / target identified

**Gate clause 1**, under its literal unit. Three consecutive `rosetta-demo down 2 --purge` + `up-injected.sh 2`
cycles, each reaching `green:true / warnings:0`. Slot **2** is ours; slot 1 is not touched.

Clause 2's Playthrough suite is deliberately **not** in this iter's scope. Clause 2 reads *"the full Playthrough
suite passes on **that stack**"* — the stack clause 1 produces. Running the suite now, on a stack about to be
purged twice, would measure a stack that no longer exists by the time clause 1 closes, and its reset-to-seed
lifecycle would mutate state between cycles that are supposed to be consecutive. It is routed to iter-261.

## Hypothesis

The advance (`app 3eaadae68` / `next-web-app 19423a1fb` / `ant-academy 249430c3`) is not merely buildable once —
it is **reproducibly** buildable, and the teardown path that clause 1 names is intact. Expressed as the thing
that could fail: a repeat cycle exercises `down --purge` (data-dir removal, host-native listener reaping,
registry slot release) which the fresh-slot bring-up never touched, and the M217 leak class lives exactly there.

## Expected lift

Gate clause 1 moves from **0 of 3 (strict) / 1 of 3 (lenient)** to **3 of 3 under the strict reading**, which
satisfies it under both. If budget permits only two cycles, the honest result is 2 of 3 strict and the third
routes forward — a partial that is stated as a partial, never rounded up by adopting the lenient reading.

## Pre-registrations — sealed in this iter's FIRST commit, before any teardown

| | claim | prediction |
|---|---|---|
| PR-1 | iter-258's fresh-slot bring-up does **not** satisfy clause 1's literal unit; the strict count before this iter is **0 of 3** | **HOLDS** — from the clause text + iter-258's own close |
| PR-2 | `down 2 --purge` leaves **demo-1's 11 containers untouched**, and demo-1's cockpit `:17700` still bound afterwards | **HOLDS** — `down` is hard-scoped `-p demo-N` |
| PR-3 | all three cycles reach `green:true / warnings:0`, `EXIT_CODE=0` | **HOLDS** — cycle 1 was green under identical refs + identical tooling pin |
| PR-4 | per-cycle wall time lands in a broad band around iter-258's 717 s, and the spread across three cycles is **larger than any inference about build performance could survive** — i.e. these are timings, not a baseline | **HOLDS**, and the number is published CONTENDED or not at all |
| PR-5 | no host-native listener orphan accumulates across three cycles: after each `down --purge`, `:27700` and `:23077` are free; after each `up`, exactly one process holds each | **HOLDS** — but this is the M217 leak class and it is the one at genuine risk |

## Phase plan

- **Phase A** — seal these pre-registrations in the iter's first commit.
- **Phase B** — cycle 1: `down 2 --purge` → verify demo-1 intact + listeners free → `up-injected.sh 2` →
  capture `autoverify.json` + `EXIT_CODE` + paired clock reads.
- **Phase C** — cycles 2 and 3, identically. Budget-aware: each cycle's evidence is captured **as it closes**,
  so a stop between cycles yields a truthful partial rather than a lost measurement.
- **Phase D** — grade PR-1…PR-5, close.

## Escalation conditions

- A cycle that fails to reach `green:true` is **the finding**, not a retry loop — capture it, do not re-run to
  get a nicer number, and close the iter on the failure with the evidence.
- Any teardown that touches a `demo-1-*` container is an immediate stop; it is forbidden this run.
- A mid-cycle ENOSPC presents as `redis exited (1)` (M239-F1), not as a disk error — check `df` before booking
  any such crash as a defect.

## Acceptable close-no-lift outcomes

A measured *"the second cycle does not reproduce"* would be a complete iter and a more valuable one than three
greens: it would mean the advance builds once and not twice, which is precisely what one green cycle cannot
distinguish. Recording that with its log is Fate 1.
