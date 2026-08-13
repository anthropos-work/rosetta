---
iteration_type: tik
status: closed-fixed
opened: 2026-08-02
---

# iter-35 — clause 2: the drill-down's positional coupling

**Active strategy reference:** `TOK-01: instrument first, then follow`. Clause 2 is a
rosetta-extensions clause; TOK-01's step 5 ("prove it cold") is the phase, and the instrument
(the Playthrough suite) is proven deterministic (iter-29) and cheap — **4 min 50 s**, reset included
(iter-32, correcting a 35–40 min estimate carried unchecked across seven hand-offs).

## Step 0 — re-survey before targeting

| check | result |
|---|---|
| platform origin HEAD re-fetched at open | **`2adcf71` — unchanged.** Trigger stays at occurrence 1 of 2 |
| rosetta tree clean at `2c9befe` (iter-34 committed) | ✅ |
| rext `main` @ `b2b46cb`, clean; pin `fast-build-m257x-iter-31b` on origin | ✅ no re-pin (iter-34 touched no rext source) |
| `demo-1` up? | ✅ 8+ containers running, carrying iter-31b's seed |
| target still meaningful? | **YES.** Clause 2 is `27 / 3 / 1`; `pt-activity-drilldown` is the best-evidenced of the three survivors and nothing has absorbed it |

Not stale. No substitution.

## Cluster / target identified

`CHECK-M257x-iter27-drilldown-target-coupling`. `activity-drilldown.spec.ts:113` fails on
`heroRow.count() > 0` — the hero's name absent from the per-member breakdown of the content row the
test drills. It never reaches its own role assertion, so it is **not** in the role-text family iters
30/31 fixed.

The coupling is **positional**: `activity-dashboard-page.ts:77-79` drills `contentRows().first()`,
and the spec's own comment (`:103-105`) rests its determinism on *"the grid sorts by most-recent
activity and the seeded heroes' sessions are dated today, so the first row is a hero-session
content — measured twice."* iter-27 added hero feedback rows and plausibly disturbed that ordering.

## Hypothesis

The first content row is no longer a content the hero participated in. The assertion's **stated
purpose** (`:97-98`) is *"this is the manager's OWN tenant's breakdown rather than any populated
org's"* — a claim about tenancy, not about row order. **Selecting the drill target by hero
participation rather than by grid position preserves the purpose exactly and removes a coupling the
test never intended to have.**

## PRE-REGISTERED PREDICTION (before any measurement)

**P1.** The measurement will show the first content row IS populated (the earlier assertions at
`:70-95` pass) but its per-member breakdown does **not** contain the hero — i.e. the failure is a
*wrong-target* problem, not a *no-data* problem.
*Falsified by:* an empty breakdown, or the hero being present and the match failing on name
formatting (which would make it a locator bug, not a coupling bug — the iter-30 lesson that a failing
assertion cannot distinguish "data missing" from "accessor wrong" applies, and this prediction is
written so the two are distinguishable).

**P2.** Fixing the target selection moves clause 2 to **28 / 2 / 1** and does **not** disturb the
other 27. *Falsified by:* any previously-passing id regressing.

## Expected lift

Clause 2: `27 / 3 / 1` → `28 / 2 / 1`, confirmed by a **binding full run in this same iteration**
(the suite is 5 minutes; iter-32's correction means a fix iter no longer defers its own measurement).

## Phase plan

1. **Phase A — measure, don't reason.** Run the single spec scoped and capture *which* content the
   first row is and whether the hero appears anywhere in that breakdown. Do not touch code first.
2. **Phase B — fix by evidence**, in `rosetta-extensions` (page object + spec), with the negative
   control iter-31 established: the fix must not make the assertion true for a *different* org.
3. **Phase C — binding full run** (`--reset`), all 30.
4. **Phase D — close**, re-pin rext if runtime source changed, verify the tag on origin.

## Escalation conditions

- A platform commit at close → re-scope trigger fires, exit.
- The fix would need a platform-repo edit → escalate; 0 platform edits is binding.
- A previously-passing id regresses → treat as a real finding, not noise; iter-29 proved this
  instrument deterministic.

## Acceptable close-no-lift outcomes

A measurement that falsifies P1 — showing the failure is a locator/formatting problem or a genuine
data gap rather than target coupling — is a complete iter even with no metric move, provided the
falsification is recorded and the next target named.
