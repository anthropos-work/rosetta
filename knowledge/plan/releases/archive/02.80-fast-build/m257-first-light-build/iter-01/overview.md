---
iter: 1
milestone: M257
iteration_type: tok
tok_flavor: bootstrap
status: closed-fixed
created: 2026-07-31
---

# iter-01 — bootstrap tok: author the opening strategy

**Type:** tok (bootstrap) · **Strategy class:** `new-direction` · **Records:** `TOK-01`

The milestone's first iter. Its job is to author the **first** strategy, not to move the metric —
there is no prior strategy to revise and no prior iter to read. Per the iter protocol a bootstrap
tok does **not** terminate the call; the loop continues into iter-02 as a tik under `TOK-01`.

## Inputs

- `overview.md` — the gate's five clauses, the L1–L10 lever table, the two inherited blocks
- `spec-notes.md` — the Phase 0b audit verdict + the odysseus recon (F1–F4)
- `decisions.md` — `D120` / `D121` / `D122`, and the `USER-BLOCKER` record the gate cleared
- `corpus/ops/demo/build-budget.md` — the declared `iteration_protocol_ref`
- `knowledge/plan/roadmap.md` § Active — v2.8, decisions `D-v28-1 … D-v28-14`

## Pre-flight (Phase 0b) — ran, and it mattered

`/developer-kit:audit-kb-fidelity --milestone=M257` returned **RED**, which **blocked this tok from
authoring strategy** until remediated. That is the gate working as designed: the doc this milestone
is required to follow asserted in six places that the gate is measured on `billion` against a
`666.29 s` baseline, both superseded hours earlier by `D-v28-14`. A strategy authored before the
audit would have been built on it.

Re-run after remediation: **YELLOW — prior RED CLEARED** (11 CLOSED / 1 PARTIAL / 2 tracked-YELLOW,
0 blockers). Phase 0b's YELLOW path is *proceed, with the gaps as this tok's known-context*, and
they are carried explicitly in `TOK-01` § Known-context.

## Output

`TOK-01: instrument before baseline, baseline before levers` — recorded in the **milestone-root**
[`decisions.md`](../decisions.md). See it for the strategy, its rationale, the distance-to-gate
context, and the iter-02 direction.

## Why this is not a tik

No lever was priced and no cycle was run, deliberately. The gate's distance is **unknown** —
odysseus's baseline does not exist — and the release's own standing rule (*state the environment
with every number*) forbids inheriting billion's. Picking a lever now would mean pricing it against
a number from a machine this milestone is not allowed to use.
