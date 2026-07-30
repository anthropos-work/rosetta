---
milestone: M256
iter: 07
iteration_type: tik
status: archived
active_strategy: TOK-01
created: 2026-07-28
---

# M256 · iter-07 — the negative-control mechanism for the 16 Playthroughs that don't write

**Type:** `tik` · **Active strategy:** `TOK-01` move 4 (the honesty items). This is
`NEGCTL-M256-ablation-harness` + its designated first proof target `FIX-M256-studio-false-green` +
`DOC-M256-llm-lane-premise` — which the routing table already groups as *"the same tik"*.

## Step 0 — re-survey (mandatory), and it caught an iter-06 defect

`ptvalidate --manifest-dir manifest --e2e-dir e2e/tests --seed-worlds …` **FAILED** on the tree iter-06 had
just committed: one ORPHAN Playthrough id. **iter-06's own fence file explained the grammar collision by
quoting the rejected tag verbatim in its header prose, and `discover.go` harvested that comment as a
Playthrough id.** A COMMENT minted a phantom test.

It survived the whole of iter-06 because `run-playthroughs.sh` reconciles with `ptreport`, which does **not**
scan `@pt:` tags — only `ptvalidate --e2e-dir` does, and that runs separately. So the harness was green three
consecutive times while the validator was red. Fixed (the mention is now spelled apart) **and fenced**: a new
fence test applies the validator's own rule inside the harness, across **every** file in `tests/` including the
unit specs, because that is where the phantom lived. `ptvalidate` now reports `21 live Playthrough(s)`, VALID.

Otherwise the survey confirms the routed target is current: `ptreport` 21/23 passing, 0 failing, 2 TODO;
clause 1 at 0.6245×; clause 2 mutating **5/5 MET**, negative controls **5 of 21**, `blocked` **0**.

## The re-scope risk is RETIRED — and the audit's premise was wrong (recorded here because it is early)

The milestone's stated re-scope trigger is *"> 3 un-homed curated UCs prove unimplementable"*, and onboarding is
**5 of the 9** UCs clause 3 must land. The Phase-0b audit (F5) concluded there is **no pre-onboarding state and
none can be declared**, because `UsersSeeder` writes a membership for every user unconditionally. **That
conflated org membership with onboarding completion. They are different columns.**

Measured this iter:

- Onboarding completion lives in **`public.user_params.onboarding`** (a `jsonb` column — found via
  `app/internal/data/ent/userparam_update.go` §`SetOnboarding`, served by the `onboarding(userId:)` query in
  `queries.graphqls:46`). There is **no onboarding table at all**.
- That column is **NULL for all 191 seeded users**. So the pre-onboarding state is not merely seedable — **it is
  the DEFAULT, already present for every seeded hero.**
- And it **drives**: `/onboarding` was probed live for both `pt-employee` and `pt-manager` and renders the real
  first step with working controls (`Upload` / `Skip` / `Next`), no redirect, no `/login` bounce.

**Verdict: onboarding is UNBUILT, not impossible.** The re-scope trigger is **not** tripped, and clause 3's
scope is **not** reduced. The build is routed to iter-08 (`ONBOARD-M256-build`) rather than crammed in here —
this iter's line is the clause-2 mechanism, and a 5-UC coverage cluster is not a side errand.

## Cluster / target identified

Clause 2's remaining negative-control gap is **16 Playthroughs that do not write** (14 `READ-ONLY` + 2
`UNKNOWN`). iter-06's D22 pattern — the pre-state read as the control — **cannot reach them**: there is no
mutation whose absence to demonstrate. They need a different mechanism, and the one with real evidence behind it
is **outcome ablation**.

## Hypothesis

**H1.** If the app's own data query is intercepted so the surface renders with **no data**, then a Playthrough
whose final assertion is genuinely anchored on its outcome will **not** match — while one anchored on page
chrome **will still match**, and is thereby exposed as a false green. Ablation is therefore both the negative
control *and* the false-green detector, using one mechanism.

**H2.** `pt-studio-advanced-generate` will still match under ablation, because iter-02 established its
completion matcher fires on the route's own page header. That makes it the mechanism's first proof target: the
harness is validated by **catching a defect already known to exist**, which is the strongest validation
available.

**The honest-control requirement.** An ablated page must still be *alive*. If ablation blanks the app entirely,
"the locator does not match" degenerates to "a broken page shows nothing" and proves nothing. So every control
asserts **two** things: the surface's chrome still renders, AND the outcome locator does not match. Without the
first, the second is worthless.

## Expected lift

- A reusable ablation helper + a negative-control spec, with the studio false green **caught, then fixed, then
  its control proven**.
- Clause 2 negative controls **5 → 6+** with the mechanism proven for the rest (the remaining 15 are mechanical
  once the harness exists; routed to iter-08 deliberately, not hand-waved).
- `DOC-M256-llm-lane-premise` corrected once, against the fixed behaviour.

## Phase plan

Steps 3 → 4 → 5 → 6 of `corpus/ops/demo/playthroughs.md` § The iteration protocol.
**Phase A** probe the ablation mechanism live before building on it (the milestone's standing lesson, now
three-for-three). **Phase B** build the helper + the control; fix the studio matcher. **Phase C** full suite +
`ptvalidate` with `--e2e-dir` this time. **Phase D** close, commit, tag, push.

## Escalation conditions

- If ablation blanks the app rather than emptying the surface, the mechanism fails the honest-control
  requirement → record the falsification and route a different mechanism, do **not** ship a control that cannot
  discriminate.
- If fixing the studio matcher needs a platform edit → escalate / diagnosed draft. Never a platform edit.

## Acceptable close-no-lift outcomes

- Ablation measured and found unable to satisfy the honest-control requirement. That is a real finding about the
  suite's shape and it re-aims clause 2's third sub-target.
