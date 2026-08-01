---
milestone: M257x
iter: 28
---

**Type:** tik (under `TOK-01: instrument first, then follow`)

# iter-28 — the binding clause-2 number: `25 live / 5 failing / 1 unimplemented`

## What was done

One full `--reset` Playthrough run on `demo-1`, consuming rext at `fast-build-m257x-iter-27` (`b718149`)
from origin — the instrument's binding mode (the ptreport gate binds only on a full run). 209 specs.

    Playthroughs coverage: 25/31 passing (80.6%)

**`23 / 7 / 1`  →  `25 / 5 / 1`.** Gate clause 2 is NOT met (it wants `30 / 0 / 0`), but this is its
second movement in the milestone and the largest.

## The comparison is a sorted-id diff, never two summary lines

iter-19's rule: `23/7/1` and `25/5/1` tell you nothing about *which* ids moved.

    $ diff iter27-failing.txt iter28-failing.txt
    2d1
    < pt-assignment-assign
    6d4
    < pt-workforce-org-feedback

**Two removals, ZERO additions.** The five survivors are byte-identical to five of the seven.

| removed | attributable? |
|---|---|
| `pt-workforce-org-feedback` | **YES** — iter-27's fix, predicted in advance, and the prediction was made before the run |
| `pt-assignment-assign` | **NO. Recorded as an open question, not an attribution.** |

## The unpredicted removal, and why it is NOT being claimed

`pt-assignment-assign` asserts *"the assignable-affordance count drops by exactly one"*
(`expect(after).toBe(before - 1)`), and iter-26's run failed it with **Expected 15 / Received 14** — i.e.
`before` read **16** and `after` read **14**, a drop of **two**. `assignableCount()` counts *members with no
skill-path assignment*, read live from a hydrating grid.

Nothing in iter-27 touched assignments. A plausible mechanism exists — `before` sampled while the grid was
still hydrating, the exact class `activity-drilldown.spec.ts`'s own comment documents for this app
(*"while the grid hydrates it renders TWENTY `<tr>` with no cell content"*) — **but it was not measured, so
it is not the finding.** This milestone has already had one inference refuted an iter after it was made.

**What IS the finding: this is the SECOND un-attributed flip.** iter-26 recorded
`hiring.recruiter-comparison.UC1` flipping to passing with *"plausible mechanism, nothing measured"*, and it
is still open. Two of the seven failures have now resolved themselves between full reset runs with no
targeted change. That means **the clause-2 metric carries an unquantified flake component** — and the gate
demands `30 / 0 / 0`, a conjunction that a flaky suite cannot satisfy *reliably* even once it can satisfy it
*once*. Routed as `CHECK-M257x-iter28-clause2-flake-component`; the cheapest measurement is to re-run the
full suite twice more against an unchanged build and diff the three id sets.

## Side effect, deliberate: the clobbered artifact is repaired

`FIX-M257x-iter27-scoped-run-clobbers-binding-report` — iter-27's scoped diagnostic run had overwritten
`e2e/report/last-run.json`, reducing the binding 209-spec artifact to 1 spec. This run restores it
(`specs=209`). **The underlying defect stands**: nothing in the file distinguishes a binding full run from
an advisory scoped one, so the next scoped run destroys it again. The route stays open.

This run also exercised the reset path iter-27's manual `stackseed` invocation skipped — `--reset` also
refreshes the fake-FAPI roster, re-exports the cockpit manifest, and reloads Sentinel's casbin enforcer
(iter-27's hand-rolled Sentinel reload returned `000`, i.e. never connected, because it used the un-offset
port). So iter-27's fix is now proven through the full documented path, not only through a hand-driven one.

## Close — 2026-08-01

**Outcome:** gate clause 2 moves `23/7/1` → `25/5/1` on a binding full run; two removals, zero additions,
one attributable to iter-27 and one deliberately not attributed.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (3 of 5 — clause 2 needs `30/0/0`; clauses 1, 3, 4 hold)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n (platform origin `2adcf71`
re-checked at open and close; occurrence stays 1 of 2) — (4) user-blocker: n — (5) cap-reached: n —
(6) protocol-stop: n — Outcome: continue
**Decisions:** the substitution recorded in this iter's `overview.md` Step 0 (the succession dig is
open-ended; the routed full measurement was the higher-value target under the same TOK).
**Side-deliverables:** the 209-spec binding artifact restored.

**Routes carried forward:**

| item | why | target |
|---|---|---|
| `CHECK-M257x-iter28-clause2-flake-component` | **NEW, and it bears on whether clause 2 is reachable as written.** Two of seven failures have now flipped to passing between full reset runs with no targeted change (`hiring.recruiter-comparison.UC1` at iter-26, `pt-assignment-assign` here), neither attributed. Measure it: run the full suite twice more against an unchanged build and diff the three id sets. A gate of `30/0/0` over a suite with an uncharacterized flake component is a different problem from five failing Playthroughs. | next tik |
| `FIX-M257x-iter27-succession-hero-not-rendered` | Unchanged and still the best-evidenced remaining failure — her interview row EXISTS, FK'd to her real session. Read-side. NB the row the spec wants is a **computed projection** ("Rare skill held only by this person / In fragile role"), so the question is what the app derives, not what the seed wrote. | next tik |
| `FIX-M257x-iter27-funnel-card-role-missing` | Unchanged. Her card renders; only the role text inside it is missing while the DB carries it on three axes. DOM/locator-shaped. | next tik |
| `FIX-M257x-iter27-scoped-run-clobbers-binding-report` | Artifact repaired, **defect not fixed** — the next scoped run destroys it again. | next tik |
| `CHECK-M257x-iter27-drilldown-target-coupling` · `CHECK-M257x-iter27-assignment-affordance-count` | The latter is now **partly answered by the flip** — it passed once. Fold it into the flake measurement rather than treating it as a separate defect. | later tik |

**Lessons:**

- **Predict before you measure, and say so in writing beforehand.** iter-28's `overview.md` recorded the
  expectation `24/6/1` *before* the run. The result was `25/5/1`, and because the prediction was on record
  the extra removal was immediately visible as something requiring explanation rather than as a bonus to be
  folded into the headline. A number that beats its prediction deserves more suspicion than one that meets it.
- **A metric that improves without a cause is not obviously good news.** Two un-attributed flips in three
  iters is evidence about the *instrument*, not only about the build.
