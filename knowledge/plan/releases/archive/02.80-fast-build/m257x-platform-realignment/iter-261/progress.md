# iter-261 — progress

**Type:** tik
**Active strategy:** `TOK-08`, under the user binding `D-M257x-256-1`.

Pre-registrations sealed in this iter's FIRST commit, before the suite ran.

## Phase B — the run

`run-playthroughs.sh 2 --reset`, full suite, **no `--grep`**, **no `PW_WORKERS` override** (serial
`workers:1`, the config's own default). Window `14:57:08Z → 14:59:50Z` = **162 s CONTENDED** — not a
baseline. `SUITE_EXIT=1`.

The reset ran the real path: a full FK-ordered TRUNCATE (36 `public.*` tables) followed by a fresh
`pt-world.seed.yaml` seed — **not** an additive re-seed. Playwright then ran **215 tests / 1 worker →
1 failed, 214 passed (2.6 m)**. (215 ≫ 31 because the suite carries unit *fence* specs alongside the
Playthroughs; the Playthrough denominator is the gate's, below.)

### The gate's verdict — clause 2 is NOT MET

```
Playthroughs coverage: 29/31 passing (93.5%)
  passing=29  failing=1  unimplemented=1  unimplementable=0
ptreport: GATE no-regressions FAILED (a Playthrough is failing)
```

Clause 2 asks for **30 live / 0 failing / 0 error**. Measured: **29 passing, 1 failing**, plus the 1
`will-not-build` verdicted TODO that the corpus already declares. **29 ≠ 30, and one failing is not zero.**

## Phase C — triage

**The single failure is `workforce-intelligence.talent-pool.UC1` (`@pt:pt-workforce-succession`)** — the
manager's succession / at-risk view. `Expected: > 0, Received: 0` after a 15 s predicate timeout: the
projection never names the org's seeded hero.

**It is not an environment or harness failure, and the evidence is on the page itself.** From the captured
snapshot (`failing-page-snapshot.md`, text only — no media was opened):

- the page **renders**, with `heading "Succession Planning"`, the correct org (`Meridian Labs · Workforce`)
  and the correct logged-in manager (`Morgan`) — so login, tenant scoping, routing and the locator all work;
- the surface's own cards render, including **`Good data coverage`** and **`At-risk people`**;
- and **both projection tables render `img "No data"`** (refs `e599`, `e805`). The surface is **populated
  with chrome and empty of rows**.

**Its three siblings pass on the same login, same org, same seed, same run** — `workforce-roster`,
`workforce-funnel` and `workforce-org-feedback` are all `[PASS]`. So the org **is** populated and the
manager **can** read workforce surfaces; exactly one **computed projection** comes back empty. The spec's
own header says it: *"Succession/at-risk are COMPUTED PROJECTIONS."*

### Two candidate causes, and this iter does not choose between them

1. **Product** — the advance changed how the succession/at-risk projection is computed, so it no longer
   selects the seeded population.
2. **Seed-contract drift** — `pt-world` at the frozen iter-101 pin no longer supplies whatever the advanced
   projection requires.

Both are *product/contract*, not harness — which is what PR-4 pre-registered as the discriminator, and it
fired in that direction. **Discriminating 1 from 2 is not available inside this iter**: it needs either a
pin bump (which would destroy the single-changed-variable control `D-M257x-258-1` exists to preserve) or an
`app` source read, and the milestone forbids platform edits either way. It is routed, not guessed.

## Phase D — grading the pre-registrations

| | prediction | outcome | note |
|---|---|---|---|
| PR-1 | `--reset` is required | **HELD IN SUBSTANCE, counterfactual NOT RUN** | the reset demonstrably laid `pt-world` down (TRUNCATE-then-seed in the log). That a *bare* `demo-up` lacks it was **not** tested — I ran only with `--reset`, so I record this as untested rather than confirmed |
| PR-2 | the 157-iter-stale harness runs to a verdict against the advance | **HELD** | 215 tests executed, report JSON written, gate rendered a verdict. The frozen pin is not too stale to measure with |
| PR-3 | **30 live / 0 failing / 0 error** | **REFUTED** | 29 passing / 1 failing. This was pre-registered *at genuine risk* and the risk was real |
| PR-4 | failures cluster in product surfaces, not uniformly | **HELD** — with its power stated | 1 failure, 3 passing siblings, page renders with correct org/manager and empty tables. Clustered, not uniform. **n=1 limits how much this can carry**, and it is reported as a direction, not a proof |
| PR-5 | the reset is per-stack; `demo-1` untouched | **HELD** | `demo-1`'s 11 containers `diff`-identical on name + status + **container ID** against iter-260's baseline, after a full 36-table TRUNCATE on `demo-2` |

## Close — 2026-08-10

**Outcome:** **Clause 2 is measured for the first time in this milestone, and it is NOT MET: 29/31
passing, 1 failing.** The failure is a real, single-surface, product-side finding — the manager's
succession/at-risk **computed projection** renders its chrome and returns **no rows**, while its three
sibling workforce Playthroughs pass on the same login and seed. The 157-iter-stale harness proved
**capable** of measuring the advance (PR-2), which is itself a result about the frozen pin.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue

**Why `closed-fixed` and not `closed-no-lift`:** the iter's planned scope was *run the suite and grade it*,
and that landed completely, with the clause answered by a number for the first time. A RED gate is the
deliverable here, not a failure to deliver.

**Decisions:** none new. `D-M257x-258-1`'s frozen pin is **vindicated as a control** — because the harness
did run, the failure is attributable to the advance-or-seed contract rather than to tooling currency.

**Side-deliverables:** none.

**Routes carried forward:**
- `ROUTE-M257x-261-succession-projection-is-empty` → **new, and it is the milestone's first measured
  functional regression candidate.** Discriminate *product change* from *seed-contract drift* without
  bumping the frozen pin. Handler: `FIX-M257x-261-discriminate-succession-empty`.
- `ROUTE-M257x-260-clause-2-never-run` → **CLOSED.** It has now run; its successor is the route above.
- `ROUTE-M257x-258-no-dev-stack-on-this-box` → **UNBLOCKED by the user mid-iter** (prohibition lifted;
  relocated-path option withdrawn). The documented `/dev-up` at `stack-dev/` is the **next iter** and the
  milestone's critical path.
- All earlier routes unchanged and open.

**Lessons:**
1. **A green bring-up and a working platform are different claims, and only the second needs a user.**
   Clause 1 went 3/3 green on the same stack that fails a Playthrough. Assembly is not function, and the
   gate is right to ask for both.
2. **"The page is empty" is a much stronger finding than "the test failed" — and it costs one text read.**
   The snapshot separated *renders-with-no-rows* from *does-not-render*, which is the whole difference
   between a computed-projection defect and a routing/login/harness defect, and it ruled out three
   explanations without opening the video or the screenshot.
3. **Grade a pre-registration you did not actually test as untested.** PR-1 was easy to wave through on
   circumstantial evidence; the counterfactual run never happened and the record says so.
