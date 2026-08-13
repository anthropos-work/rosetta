**Type:** tik (under `TOK-01: instrument first, then follow`)

# iter-35 — the drill-down's "first row" was never determinate

## Clause 2: `27 / 3 / 1` → **`28 / 2 / 1`**

Binding full run, cold `--reset`, all 31 manifest ids: **28 passing / 2 failing / 1 unimplemented**,
`ptreport` reporting no regressions among the 27 that were already green. Both predictions held.

| prediction (registered before any measurement) | outcome |
|---|---|
| **P1** — the breakdown is *populated* but lacks the hero: a **wrong-target** problem, not a no-data one | ✅ **CONFIRMED** — every assertion up to `:113` passed; only `heroRow.count() > 0` returned 0 |
| **P2** — fixing target selection moves clause 2 to `28 / 2 / 1` with no other id disturbed | ✅ **CONFIRMED exactly** |

P1 was written so that a locator/formatting bug and a data gap would falsify it differently — the
iter-30 lesson that *a failing assertion cannot distinguish "the data is missing" from "the accessor is
wrong."* It was right about the class and **wrong about the cause**, which is the part worth keeping.

## The premise was false, and it was false in a way two runs could not detect

The hypothesis inherited from iter-27 was that *new hero sessions disturbed the grid's ordering*. The
measurement says the ordering never existed. Querying demo-1's Postgres for the hero's org
(Meridian Labs), grouping sessions per content:

```
        tied_timestamp         | contents_tied | tied_without_hero | sessions
-------------------------------+---------------+-------------------+----------
 2026-08-01 21:14:45.300746+00 |            11 |                 2 |       45
```

**Eleven distinct contents share the identical `max(started_at)` — to the microsecond** — because the
seeder writes every backdated session at one instant. "The grid sorts by most-recent activity" therefore
**does not order the top of the grid at all**; the first row is whatever the backend's secondary sort
returns, and **2 of those 11 contents have no hero session** — one of them the 19-session row, a likely
winner under any count-based tiebreak.

The spec defended its own determinism in a comment: *"the first row is a hero-session content — measured
twice, on separate runs."* **Two agreeing draws from an 11-way tie is not a measurement of determinism.**
This is the *"check whether an assertion's truth is a tiebreak"* rule, and it had been standing in this
spec, in writing, as the justification for the coupling.

**The tie reproduces cold.** Re-measured after the binding run's `--reset`, the timestamp moved
(`19:58:13.602842` → `21:14:45.300746`) and the shape did not: **11 tied, 2 without the hero.** It is
structural, not incidental — which is why two runs agreed and a third did not.

**Corroboration nobody had connected.** `negative-controls.spec.ts:566-572` already works around this
without naming it: it uses the *STARTED* hero rather than the thriving one because *"measured, the
thriving hero is NOT in this content's breakdown."* Someone hit the same tie one test over, patched
around the symptom, and recorded the observation as a property of the hero instead of the ordering.

## The fix — select by the property the assertion is about

`drillIntoContentContaining(memberName)` scans the grid's content rows and drills the first whose
per-member breakdown contains that member, returning **-1** when none of the scanned rows does — so a
genuine absence **fails loudly** instead of silently drilling something arbitrary. The spec asserts
`drilledIndex >= 0` as its tenancy final.

This does not weaken the claim; it sharpens it. The old assertion said *"the hero is in THIS content's
breakdown"* where "this content" was a coin-flip. The new one says *"some content of this manager's org
has this hero in its breakdown"* — which a manager of any other seeded tenant fails, and which the
existing negative control independently exercises. The role assertion is unchanged and still carries the
"identifies the member, not just a row count" half.

## A defect only a scanning test could find

The first fixed run still failed — `locator.click` timing out at 5 s with
`<div role="tooltip" class="ant-tooltip-container"> … intercepts pointer events`. After `goBack()`, the
antd tooltip opened by hovering row *i*'s link **survives the re-render and covers row *i+1***. A
single-click test never hovers twice, so this had been unreachable for the life of the suite.

The pointer is now parked at `(0,0)` and the tooltip awaited `detached` before each drill.
**`force: true` would also have made it pass** — by refusing to check that the element is clickable,
which is the assertion the click exists to make.

## Evidence

`evidence/tie-measurement.txt` (the post-reset re-measurement) · `binding-run.log.txt` (the binding
`--reset` run) · rext `24ec6c9`, tag `fast-build-m257x-iter-35` **verified on origin**.

## Close — 2026-08-02

**Outcome:** clause 2 moves **`27 / 3 / 1` → `28 / 2 / 1`** on a binding cold-reset run — the predicted
value exactly. Root cause was not the inherited hypothesis: the drill target was chosen by grid position
under an **11-way timestamp tie**, so the assertion's truth was a tiebreak. Fixed by selecting the target
on hero participation; a tooltip-interception defect that only a multi-row scan could reach was fixed on
the way.
**Type:** tik
**Status:** closed-fixed (planned scope was measure → fix → confirm bindingly in one iter; all three
landed, and the 5-minute suite made the same-iteration confirmation possible)
**Gate:** NOT MET (**3 of 5**. Clause 2 needs `30 / 0 / 0`; two survivors remain)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n (platform origin `2adcf71`
re-fetched at open and close, unchanged; trigger stays at occurrence 1 of 2) — (4) user-blocker: n —
(5) cap-reached: n (2 tiks this session) — (6) protocol-stop: n — Outcome: continue
**Decisions:** `D-M257x-35-1` … `D-M257x-35-3` (`iter-35/decisions.md`).
**Side-deliverables:** none.
**Re-pin:** **YES** — rext runtime source changed. `.agentspace/rext.tag` and the `stack-demo`
consumption clone both moved to `fast-build-m257x-iter-35`, and the tag is confirmed on origin via
`git ls-remote`. No `stackseed` rebuild needed: the change is `playthroughs/e2e/` TypeScript, not
`stack-seeding` Go.
**Routes carried forward:**
- `FIX-M257x-iter32-hiring-candidate-sim-link` — `pt-onboarding-hiring-candidate`, the assigned position
  not rendering as a startable org-scoped `/sim/<slug>?organizationId=` link.
- `FIX-M257x-iter32-orgadmin-role-create-timeout` — `pt-orgadmin-role-create`, 60 s `waitForURL` timeout
  after Save. **These two are all that stand between clause 2 and MET.**
- **`CHECK-M257x-iter35-negative-control-rests-on-the-same-tie`** (new) —
  `negative-controls.spec.ts:564` still drills row 0 and asserts the STARTED hero is present. That is the
  same tiebreak, one test over, and it is currently green by luck. It did not fail this run, so fixing it
  here would have been unmeasured work; but it is a known latent flake with a named cause.
- **`CHECK-M257x-iter35-seeder-writes-one-instant`** (new) — the deeper cause is that the seeder stamps
  every backdated session with a single timestamp, which flattens *all* recency ordering in the product,
  not just this test's. A believability question as much as a test one.

**Lessons:**
- **"Measured twice, on separate runs" is not evidence of determinism when the population is a tie.**
  Two agreeing draws from an 11-way tie agree ~all the time and prove nothing. Before trusting a
  positional selector, ask what the *distribution* of the ordering key is — not whether two runs matched.
- **Select on the property the assertion is about.** The test cared about hero participation and selected
  on grid position; every failure since has been about the gap between those two things.
- **A workaround recorded as a fact hides its own cause.** The negative control already knew the thriving
  hero was missing from row 0 and wrote it down as a property of *the hero*. Naming it as a property of
  *the ordering* would have found this a milestone earlier.
- **Widening a test's traversal finds defects narrower tests cannot reach.** The tooltip interception was
  latent for the life of the suite because nothing had ever hovered two rows in one run.
