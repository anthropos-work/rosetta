---
iter: 153
milestone: M257x
iteration_type: tik
status: archived
opened: 2026-08-08
---

# iter-153 — the probe scope is derived from the wrong artifact

**Active strategy reference:** `TOK-08` — census the mechanical classes; stop sampling them.

**Step 0 — re-survey.** The routed target is `FIX-M257x-h33-derive-includes-stack-override` (harden
pass 35): *union the stack's generated override compose into `generate.sh`'s derivation, so a demo's UI
tier is probed rather than disclosed away.* Pass 35 routed it as **Fate 3, "needs a live demo to
verify"**. The re-survey **refutes that precondition**: `gen_injected_override.py` is pure Python text
emission and `stack-demo/platform` is a real clone at `0c91421`, so the stack's own generated override
can be produced HERE, across the whole flag matrix, with no docker and no bring-up. The target stands and
is landable in-iter.

**Cluster / target identified.** The scope question — *"what services does this stack run?"* — is
answered in three places from two different artifacts:

| site | artifact read | answer (default demo) |
|---|---|---|
| `up-injected.sh:2682-2690` | platform compose **+ 3 hand-written literals** | 8 |
| `stack-verify/reports/generate.sh` | platform compose **only** | 5 |
| the stack's own generated override | *(it IS the artifact)* | **11** |

**Hypothesis.** `generate.sh`'s scope is derived from the platform's unmodified compose, which is not the
thing that decides what a stack runs. The stack's own generated override is. Union it in — and where the
probe registry cannot express a service the stack declares, **declare that gap** rather than let the
service vanish from the denominator (`D-M257x-151-1`: an absent-arm that reads a comment cannot fire;
`D-M257x-152-2`: a census cannot find a value absent from its denominator).

**Expected lift.** Not an `N` reading. A live-code defect closed on the `/test-platform` path + a fence
that enumerates the population, with the denominator stated in both directions.

**Phase plan.** A (census the scope-answering sites + measure the generator's flag matrix) → B (repair
`generate.sh`'s derivation + declare the registry gap) → C (fence, with mutation + anti-vacuity controls
run against a REAL generated override) → D (scoped suite run + close).

**Escalation conditions.** If the union requires `stack-verify` to import `demo-stack` *code* (rather
than read a generated *artifact*), stop and route — a verify section reaching into a bring-up section's
implementation is a coupling this milestone should not add.

**Acceptable close-no-lift outcomes.** If the measured override service set turns out to equal the
platform set on every flag combination, the harden finding is falsified and the iter closes on that.
