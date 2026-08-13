**Type:** tik (under `TOK-01: instrument first, then follow`)

# iter-32 — the binding run

## What ran

One full, unscoped, `--reset` Playthrough run on `demo-1`, from the pinned consumption clone at
`fast-build-m257x-iter-31b`, against platform origin `2adcf71` (re-fetched at open **and** at close —
unchanged both times, so the re-scope trigger stays at occurrence 1 of 2).

    cd stack-demo/rosetta-extensions/playthroughs/e2e && ./run-playthroughs.sh 1 --reset

Preconditions checked at open rather than inherited from the hand-off — all four of the standing
checklist, plus the one run 17 discovered: **`stackseed` is a per-stack built binary and re-pinning the
clone does not rebuild it.** Its mtime `23:49:35` post-dates the pinned tag's commit `23:49:26`, so the
reset ran iter-31b's seeder and not stale code. The reset itself: 66 audited write attempts, **55 708
rows**, `isolation: clean`, policy advisory OK (18 grants), and all three of the steps that used to
degrade silently — sentinel enforcer reload, fake-FAPI roster export, cockpit-manifest export — ran and
reported success. (iter-25 fixed those; this is the second run to prove they hold.)

## The number

    Playthroughs coverage: 27/31 passing (87.1%)
      passing=27  failing=3  unimplemented=1  unimplementable=0

**`27 / 3 / 1` — exactly the figure pre-registered at iter-31's close, before any confirming run existed,
and restated in this iter's `overview.md` before the run started.** Prediction met, not beaten.

## The measurement is the diff, not the headline

Per the iter-19 rule, `25 → 27` is compared as a **sorted-id set difference**, never as two summary
lines: three matching totals could be three different failures. The two vocabularies in play had to be
reconciled first — iter-29 recorded **Playthrough handler ids** (`pt-*`) while ptreport's table prints
**manifest ids** (`org-admin.roles.UC1` …) — so the diff was taken in the `pt-*` vocabulary, extracted
mechanically from the run's own `@pt:` spec tags rather than mapped by eye:

    < pt-workforce-funnel        ← fixed at iter-30
    < pt-workforce-succession    ← fixed at iter-31
    (additions: 0)

**Two removals, ZERO additions.** Both removals were predicted in writing by the iters that caused them.
The three survivors are byte-identical to three of iter-29's five:

| id | manifest id | failure |
|---|---|---|
| `pt-activity-drilldown` | `assignment-monitoring.assign-and-track.UC2` | the org's seeded hero absent from the per-member results |
| `pt-onboarding-hiring-candidate` | `onboarding.enterprise-hiring.UC1` | assigned position does not render as a startable org-scoped `/sim/<slug>` link |
| `pt-orgadmin-role-create` | `org-admin.roles.UC1` | `page.waitForURL` timeout, 60 000 ms, after Save |

The 1 TODO is unchanged and declared: `onboarding.enterprise-workforce-standard.UC1`, *"will-not-build:
MEASURED, then deliberately refused."*

**Zero additions is the load-bearing half of this result.** iter-31 changed the role distribution of
**every** seeded org — the widest blast radius of the milestone — and its five negative controls were
explicitly recorded as "not the whole suite." They are now backed by 209 specs: the change regressed
nothing.

## The tok question, decided by the number rather than by argument

This iter's `overview.md` graded Phase 0 honestly: on the **headline** metric, iters 29, 30 and 31 form a
3-consecutive-no-progress-tik streak and rule 2 would make iter-32 a **triggered tok**. It was recorded as
a tik with the question left open, because the skill's own stale-trigger clause requires re-running the
primary measurement *before* revising a strategy, and that measurement was this iter's whole deliverable —
so tok and tik prescribed an identical next action and only the close would differ.

**The trigger was STALE.** The metric had moved 25 → 27; it had simply not been *read*, because this
milestone forbids quoting a scoped run as a clause-2 number (iter-25/iter-26) and the full read was
budgeted as its own iteration. Iters 30 and 31 each made real, attributable progress.

The generalisable lesson, and it is a defect in how the streak is counted rather than in the iters:
**"the metric did not move" and "nobody read the metric" are indistinguishable from the ledger.** A
protocol that makes its primary measurement expensive will manufacture phantom no-progress streaks, and
the streak rule will then fire a strategy revision against iters that were working. Routed as
`DOC-M257x-iter32-noprog-streak-counts-unread-metrics`.

## The measurement is ~8× cheaper than the hand-off believed — which changes the iteration economics

The hand-off budgeted *"~35–40 min serial, 209 specs"* and instructed that the run be given an entire
iteration. **Measured: 4 min 50 s wall, reset included** (launch → last log write). The run is legitimate
on every check that could distinguish a real run from a truncated one: 209 tests announced and 209
numbered, the reset's own audit present, real Playthroughs taking real time (the three failures at 16.6 s
/ 17.3 s / 60 s), ptreport's gate evaluated and correctly `FAILED`, and the artifact self-describing as
`binding: true, scoped: false`.

Where the old estimate came from: **"209 specs" was never a time proxy.** Only ~31 are Playthroughs; the
majority are unit specs (`mutation-class-fence`, `url-shapes`, the locator suites) running at 0–1 ms.

A candidate mechanism for the change since iter-25's *"65 of 209 in ~35 min"* — **recorded as a candidate,
explicitly NOT attributed** — is iter-24's Directus re-point, which removed a 96-line 403 storm that was
on every page load. This milestone has had an inference of exactly this shape refuted one iter after it
was made; the honest statement is that the runtime collapsed at some point between iter-25 and now and
nothing here measured why. Routed as `CHECK-M257x-iter32-suite-runtime-collapse`.

The consequence is immediate and practical: **a binding clause-2 read no longer needs to be an iteration.**
It can close a fix iter that lands one.

## Side observation — iter-30's artifact guard fired for the first time

`e2e/report/last-binding-run.provenance.json` and `report/last-binding-report.json` were **written by this
run** and did not exist before it: iter-30 shipped the guard, but every run since had been scoped, so this
is the first snapshot it has ever taken. `{"binding": true, "scoped": false, "grep_pattern": "", "stack":
"demo-1"}`. The binding verdict is now self-describing and cannot be clobbered by a scoped diagnostic —
which is what `FIX-M257x-iter27-scoped-run-clobbers-binding-report` promised and this is its first proof.

## Evidence

`iter-32/evidence/` — `binding-failing-set.txt` (the three `pt-*` ids, extracted from the run's own spec
tags), `ptreport-verdict.txt`, `last-binding-report.json`, `last-binding-run.provenance.json`.

## Close — 2026-08-02

**Outcome:** clause 2 moves `25 / 5 / 1` → **`27 / 3 / 1`** on a binding full `--reset` run — exactly the
pre-registered prediction, by a sorted-id diff of two removals and **zero additions**, confirming iter-30's
and iter-31's fixes through the documented path and clearing iter-31's org-wide seed change of any
regression across 209 specs.
**Type:** tik
**Status:** closed-fixed (the planned deliverable was the binding measurement; it landed, met its
pre-registered prediction, and produced an id-level attribution rather than a headline)
**Gate:** NOT MET (3 of 5 — clause 2 now `27 / 3 / 1`, needs `30 / 0 / 0`; three failures remain)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (this iter was a tik; the streak was measured
STALE, so no tok fired) — (3) re-scope: n (platform origin `2adcf71` re-fetched at open and close,
unchanged) — (4) user-blocker: n — (5) cap-reached: n (1 tik this session) — (6) protocol-stop: n —
Outcome: continue
**Decisions:** the Phase 0 tok question was deferred to the measurement rather than argued (recorded in
`overview.md` before the run); the runtime finding is recorded as a candidate mechanism, not an
attribution.
**Side-deliverables:** none — no source of any kind was modified this iter. **No rext re-pin**, because no
rext runtime source changed; the pin stays `fast-build-m257x-iter-31b`.
**Routes carried forward:**
- `CHECK-M257x-iter27-drilldown-target-coupling` — `pt-activity-drilldown`, the best-evidenced of the
  three survivors. Fails at `activity-drilldown.spec.ts:113` on `heroRow.count() > 0`; an ordering
  coupling (the grid sorts by most-recent activity), not the role-text family the other two removals were.
  **Next target.**
- `CHECK-M257x-iter32-suite-runtime-collapse` — the suite got ~8× faster between iter-25 and now; candidate
  mechanism recorded, nothing measured.
- `DOC-M257x-iter32-noprog-streak-counts-unread-metrics` — an unread metric is indistinguishable from an
  unmoved one; worth a line in `platform-alignment.md`.
- `FIX-M257x-iter32-hiring-candidate-sim-link` and `FIX-M257x-iter32-orgadmin-role-create-timeout` — the
  other two survivors, now reproducible on demand.
- Unchanged: clause 5's full 40-file re-read (untouched for three runs, now the longest-standing clause).

**Lessons:**
- **Pre-register the number, then honour it.** `27 / 3 / 1` was written down before a confirming run
  existed, and the run returned it exactly. The value is not that the guess was right — it is that no
  post-hoc framing was available in either direction.
- **Diff in ONE vocabulary, and extract it mechanically.** The prior binding set was recorded in `pt-*`
  handler ids and the new verdict prints manifest ids. Eyeballing that correspondence is how a "zero
  additions" claim becomes false without anyone noticing; the ids came from the run's own `@pt:` tags.
- **"The metric did not move" ≠ "nobody read the metric."** The tok trigger cannot tell those apart, and
  a protocol with an expensive primary measurement will manufacture phantom no-progress streaks. Decide
  the trigger by measuring, never by counting ledger rows.
- **Re-measure your own cost estimates.** A budget inherited across seven hand-offs was wrong by ~8×, and
  it had been shaping how work was scheduled — one whole iteration reserved for a five-minute command.

> **Note on the evidence dir:** `.gitignore:147` excludes `knowledge/plan/**/*-report.json`, so the
> ptreport artifact is checked in as `last-binding-report.json.txt`. Same class as the `*.log` exclusion
> at `.gitignore:89` — a `.json` named `*-report.json` under `knowledge/plan/` is silently local-only.
