---
iter: 261
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-10T14:56:02Z
---

# iter-261 — clause 2: the Playthrough suite against the ADVANCED platform, for the first time

**Type:** tik
**Active strategy:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07),
under the user's binding `D-M257x-256-1`.

## Step 0 — Re-survey before targeting

iter-260 met **clause 1** under its literal unit (3/3 consecutive `--purge` + `up`, all
`green:true / warnings:0`). `ROUTE-M257x-260-clause-2-never-run` is now the critical path on the demo half,
and it is the more informative of the two remaining non-blocked lines: a green bring-up proves the stack
**assembles**; only a Playthrough proves the platform **does** anything, because a Playthrough logs in as a
seeded hero and plays a real journey to its outcome.

**The denominator was in genuine doubt and is now measured.** The consumed tooling is pinned at
`fast-build-m257x-iter-101` — **157 iters stale** by deliberate choice (`D-M257x-258-1`, kept as an
experimental control so the *advance* is the single changed variable). The corpus asserts 30 live
Playthroughs, but that is the count at v2.8 M256 and there was no reason to assume the pin carries it.
Measured at the pin:

| | |
|---|---|
| `@pt:` tags across `e2e/tests/*.spec.ts` | **30** |
| manifest `outcome: success` | **30** |
| manifest `outcome: blocked` | **1** |

**The pin carries exactly the corpus's 30 live + 1 verdicted TODO**, which is what clause 2 names
(*"30 live / 0 failing / 0 error"*). M256 closed before M257x opened, so the pin post-dates it — the count
is consistent rather than lucky, but it is now measured rather than assumed.

**Runner preconditions verified before sealing** (each has cost a prior iter a run): `stackseed` exists at
the derived per-stack path `demo-stack/stacks/demo-2/bin/stackseed` (built by cycle 3, 16:51 local);
Playwright and its Chromium are installed. The stack is a **localhost** demo — the bring-up fell back from
`--public-host` at rung 1/6 (no `tailscale` on PATH) — so `PT_HOST=localhost` / `PT_APP_SCHEME=http`, the
runner's defaults, are correct and **no `PT_HOST` is set**.

## Cluster / target identified

**Gate clause 2**, on the stack clause 1 just produced. `run-playthroughs.sh 2 --reset`, full suite, serial
(`workers:1`, the config's own default — **not overridden**).

## Hypothesis

The advance (`app 3eaadae68` / `next-web-app 19423a1fb` / `ant-academy 249430c3`) is not merely assemblable
but **functional**: a 157-iter-stale harness, written against pre-advance platform behaviour, still drives
30 real user journeys to their outcomes. Stated as the thing that could fail — this is the first place an
**advance-induced functional regression** could surface, and a bring-up cannot see one.

## Expected lift

Clause 2 answered with a number for the first time in this milestone. **Either outcome is a complete iter:**
30/30 meets the clause; anything less is the functional regression the milestone exists to find, and is
worth more than a green.

## Pre-registrations — sealed in this iter's FIRST commit, before the run

| | claim | prediction |
|---|---|---|
| PR-1 | `--reset` is **required**: `pt-world` is a decoupled seed a bare `demo-up` does not lay down, so a no-reset run would measure preconditions, not function | **HOLDS** |
| PR-2 | the pinned iter-101 harness **runs to a verdict** against the advanced refs — i.e. it does not die on a harness/locator incompatibility before producing a report | **HOLDS**, and this is a real risk given the 157-iter gap |
| PR-3 | **30 live / 0 failing / 0 error** — clause 2 MET | **AT GENUINE RISK.** A stale harness against advanced product code is exactly where a regression hides, and no prior iter has run this |
| PR-4 | any failures cluster in surfaces the advance touched (`app` / `next-web-app` journeys), **not** uniformly across products — a uniform failure means harness/environment, a clustered one means product | **HOLDS if PR-3 misses** (vacuous if PR-3 holds, and will be reported as vacuous rather than as a pass) |
| PR-5 | the `--reset` is **per-stack scoped**: `demo-1` is untouched and its 11 containers stay byte-identical | **HOLDS** — `stackseed --reset` is per-stack and N=0-guarded |

## Phase plan

- **Phase A** — seal.
- **Phase B** — `run-playthroughs.sh 2 --reset`, full suite, no `--grep`, no `PW_WORKERS` override. Capture
  the runner's own report + `demo-1` state.
- **Phase C** — grade. If failures exist, triage **by product**, and route rather than fix in-iter: a
  platform-behaviour fix is out of scope (0 platform edits), and a harness fix at a deliberately-frozen pin
  would destroy the control `D-M257x-258-1` exists to preserve.
- **Phase D** — close.

## Escalation conditions

- A failing Playthrough is **the finding**, not a retry loop. No re-run for a nicer number; no `--grep` to
  narrow onto passes.
- **The reset TRUNCATEs demo-2's world** and replaces it with `pt-world`. That is expected and safe —
  clause 1's evidence is already captured and committed at `8bf242a`, so nothing is lost.
- Anything touching `demo-1` is an immediate stop.

## Acceptable close-no-lift outcomes

A measured *"N of 30 fail, and here is the product cluster"* is a complete iter and the more valuable
result. So is *"the harness cannot run against the advance"* (PR-2 refuted), which would itself be a
first-class finding about the cost of the frozen pin.
