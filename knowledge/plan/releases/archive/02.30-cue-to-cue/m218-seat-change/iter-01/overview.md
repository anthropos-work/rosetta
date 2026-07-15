---
iter: 1
milestone: M218
iteration_type: tok
tok_flavor: bootstrap
status: closed-fixed
opened: 2026-07-13
closed: 2026-07-13
---

# iter-01 (bootstrap tok) — the 4-leg experiment

## Why a tok
Skill rule: iter-01 of an iterative milestone is **unconditionally** the bootstrap tok — it authors the first
strategy (TOK-01). The milestone's `overview.md` independently mandates the same opening move:

> **iter-01 is the HARNESS + the 4-leg experiment. Write NO fix before it runs.**

Both agree: **measure before fixing.** The confirmed cost budget (~18 s) did not sum to the reported symptom
(60–120 s), so writing a fix first would have been guessing which fix to build. This iter writes **zero
production code** and instead discharges the 4-leg experiment against the live green demo on `billion`.

## Inputs
- `overview.md` — the ranked suspect list C-1…C-6, the exit gate, the re-scope trigger.
- The Phase-0b KB-fidelity audit (**YELLOW**) — recorded in `spec-notes.md`.
- The **live demo on `billion`** (`demo-1`, rext `cue-to-cue-m217`, 16 containers, up 4 h).
- M217's hand-off: `autoverify.json` — the gate that says a stack is safe to measure.

## Measurement gate (M217's hard barrier, honoured)
```json
{"project":"demo-1","offset":10000,"warnings":0,"green":true}
```
**GREEN, 0 warnings** → this stack is legitimate to measure. (Path drift found — see D3.)

## Scope
- Run the 4-leg experiment; discriminate every suspect.
- Author **TOK-01** — the strategy the tiks will follow.
- **NO fix.** No production code. Not one line.

## Out
Any fix. The whole point is to not guess.
