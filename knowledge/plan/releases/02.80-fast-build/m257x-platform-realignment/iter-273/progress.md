# iter-273 — progress

**Type:** tik, under `TOK-08`. Route: gate clause 2 — the binding suite run at the shipping pin.

## Phase 1 — pre-registrations sealed

Four, sealed before the run. See `overview.md`.

## Phase 2 — the binding run

`./run-playthroughs.sh 2 --reset` — **unscoped**, so the harness's own gate is BINDING rather than
advisory. `SUITE_RC=1`. Wall-clock **20:46:18Z → 20:49:07Z = 169 s**, of which Playwright reports 2.6 m of
test time across **215** tests (214 passed / 1 failed). CONTENDED, host load ~4; **not a baseline**.

### Clause 2's true state at the shipping pin

From `report/last-report.json` — the machine verdict, not the console tail:

| field | value |
|---|---|
| `total` (manifest use cases) | **31** |
| `passing` | **29** |
| `failing` | **1** |
| `unimplemented` | **1** — the declared `will-not-build` verdict (`onboarding.enterprise-workforce-standard.UC1`) |
| `unimplementable-without-platform-edit` | **0** |

**Live = 29 passing + 1 failing = 30.** Clause 2 wants **30 live / 0 failing / 0 error** and reads
**30 live / 1 failing / 0 error**. It is **exactly one Playthrough** from met, and that Playthrough is the
one iter-272 decided.

The single failure is `workforce-intelligence.talent-pool.UC1` — `pt-workforce-succession` — failing on the
same assertion, with the same message, as in iter-272's two scoped runs. **Three runs, one identity.**

## Phase 3 — pre-registrations graded

| PR | verdict | evidence |
|---|---|---|
| **PR-1** — succession is the only failure | **HOLDS** | 1 `[FAIL]`, 29 `[PASS]`, 1 `[TODO]`; the failure is `talent-pool.UC1` |
| **PR-2** — the live denominator is 30 | **HOLDS** | 29 passing + 1 failing = 30 live, + 1 TODO = 31 total — matching `playthroughs.md`'s authoritative count exactly |
| **PR-3** — zero errors | **HOLDS** | no error state anywhere; `unimplementable-without-platform-edit` is **0** as well |
| **PR-4** — under 20 minutes | **HOLDS** | **169 s** |

**On PR-4, the honest form matters more than the win.** The milestone has been carrying *"the last full
suite took ~45 min"*. This run took **169 s** for reset + 215 tests. Those two numbers are **not a
contradiction I am entitled to resolve** — I did not measure what the 45-minute figure measured, and it may
have included a bring-up, a cold cache, or a period of much heavier contention. What is measured is this:
**at this pin, on this host, at this load, the binding suite is 169 s.** State the environment with the
number; do not retire someone else's measurement with your own.

## Close — 2026-08-10

**Outcome:** Clause 2 has a **binding** measurement at the shipping pin for the first time: **30 live / 1
failing / 0 error**, one Playthrough short, and that Playthrough is the one whose mechanism iter-272 closed.
The inherited one-failure prior is confirmed rather than assumed, and the suite is cheap enough (169 s) that
a fix can be graded in minutes rather than deferred for cost.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue

**Decisions:** `D-M257x-273-1` (a scoped suite run cannot grade clause 2 — the harness says so itself, and
the milestone had been grading from one).

**Side-deliverables:** none.

**Routes carried forward:**
- **`FIX-M257x-272-succession-hero-has-no-qualifying-surface`** — now the **only** thing between the suite
  and clause 2, with a measured denominator to grade against. iter-274's target.
- Gate **clause 5** — the documentation-accuracy reading, unmeasured since iter-131.
- The inherited queue (`FIX-M257x-269`, `ROUTE-M257x-270-directus-consumer-cms-key`, `FIX-M257x-266`,
  `FIX-M257x-265`, `ROUTE-M257x-h59`, `ROUTE-M257x-h65`) → open.

**Lessons:**
1. **A harness that labels its own output advisory is telling you the gate cannot read it.** The runner
   printed *"this run was SCOPED — its artifacts are advisory … Re-run unscoped for a binding verdict"* on
   every scoped run, and the milestone graded clause 2 from scoped runs anyway. **Read the provenance line
   the tool prints about itself** — it is the cheapest census in the building.
2. **Measure the cost of the feedback loop before deciding what you can afford.** The one-failure prior
   went un-rechecked partly because a full suite was believed to cost ~45 min. It costs 169 s here. A
   wrong cost estimate does not just waste time — it silently reshapes which measurements a plan is
   willing to take.
