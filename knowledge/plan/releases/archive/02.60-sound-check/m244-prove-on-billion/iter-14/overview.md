---
iter: 14
milestone: M244
iteration_type: tik
status: closed-no-lift
created: 2026-07-22
---

# iter-14 — gate (c): the stack-verify coverage sweep, live on billion

**Type:** tik · **Active strategy:** TOK-02 (sweep the remaining gates live on the green m244 seed)

## Active strategy reference
TOK-02. Run-6 priority 2: drive gate (c) — the 40 live-browser specs — live from the peer harness on the green m244 seed.

## Cluster / target identified
Gate (c) is 24 stack-verify + 16 Playthroughs. A seed constraint decides the order: the 16 Playthroughs need the dedicated **pt-world** seed and their `--reset` does a full TRUNCATE that would DESTROY the content/stories demo seed; the 24 stack-verify specs run against the CURRENT demo/stories seed. So all demo-seed gate work (24 stack-verify, f, h, d) must precede any pt-world reset, and the Playthroughs run LAST. iter-14 therefore targets the **stack-verify coverage sweep** — the core stack-verify live-browser proof (`coverage.spec.ts`, both hero vantages) — read-only on the current demo seed, no reset.

## Hypothesis
The coverage sweep (M42e employee + M42m manager) is green on the m244 seed from the peer (`COVERAGE_HOST` + https).

## Expected lift
If the whole of gate (c) landed → 5/8. Realistically iter-14 measures the coverage core; the discrete stack-verify specs + Playthroughs route forward.

## Phase plan
Run `run-coverage.sh 1 {employee,manager} --host billion` → read the semantic gate report → characterize any failure → route.

## Escalation conditions
A failure requiring a platform edit → STOP. (None found — the one failure is a demo-side academy render, gate-(d)-owned.)

## Acceptable close-no-lift outcomes
The coverage sweep runs to completion and characterizes its residual (e.g. a known gate-(d) coupling) even if gate (c) does not tick — a complete measurement cycle.
