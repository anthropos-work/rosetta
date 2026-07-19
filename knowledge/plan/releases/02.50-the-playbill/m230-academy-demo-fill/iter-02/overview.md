---
iter: iter-02
milestone: M230
iteration_type: tik
status: closed-fixed
created: 2026-07-19
---

# iter-02 — build the Option C fill mechanism (+ runtime-prove it)

**Type:** tik · **Active strategy:** TOK-01 (Option C — FS-as-published fallback demo-patch).

## Step 0 — re-survey
TOK-01's Next-tik direction named "build the demo-patch". Re-survey confirms it's untouched + still the right
target: the gate metric is still 0 real cards (F4), and no sibling milestone has filled it. Proceeding.

## Cluster / target
Build the `academy-fs-published-fallback` demo-patch in the rext authoring copy + wire it into `ant-academy.sh`,
per TOK-01. The fix mechanism is authorable + unit-verifiable WITHOUT a live demo (the heavy cold-`/demo-up`
proof is the gate's formal step, separate).

## Hypothesis
An env-gated (`ACADEMY_DEMO_FS_PUBLISHED`) fallback that reuses the tested `mergeDrafts(emptyCatalogView(), eids)`
and strips `_draft`/`_origin` makes `getServerCatalogView` render the full public FS catalog as PUBLISHED cards
(no chip) — behavior-identical to pristine when the env is unset.

## Phase plan (coverage-protocol A–E)
- A/B (measure/triage): baseline established (0 cards, F4) — code+launcher-confirmed in iter-01.
- C (fix): author manifest + native apply/revert helper + `ant-academy.sh` wiring + unit tests; tag rext.
- D (re-measure): a bounded standalone runtime proof (patched `next dev`, signed-in persona, SSR-curl the home,
  count rendered cards + assert 0 draft chips) — a strong proxy for the formal coverage-sweep re-measure.
- E (close): grade + route the formal cold-`/demo-up` gate forward.

## Escalation conditions
- The formal gate (coverage sweep on a cold `/demo-up`) is heavy infra + faces the local next-web clone drift →
  surface as a user-blocker for a go-ahead (do not fake the proof), per the orchestrator's infra guidance.

## Close
See `progress.md`.
