---
milestone: M257x
iter: 21
iteration_type: tik
status: closed-fixed-partial
opened: 2026-08-01
---

# iter-21 — clause 5: reconcile the corpus against the map

**Active strategy reference:** `TOK-01: instrument first, then follow` — step 4's second half, *"and the
reconciliation sweep"*. iter-20 landed the map (step 4's first half); this is the sweep that consumes it.

**Step 0 re-survey (done at open):** platform origin HEAD `2adcf71`, unchanged. Blast radius re-measured
in the clause-5 trees only (`corpus/services/**` + `corpus/architecture/**`): **15 files / ~78 router hits**,
plus 7 files still asserting a 3-subgraph supergraph. `DOC-M257x-iter14-corpus-router-drop`'s inherited
"35 files / ~128 hits" is a **whole-corpus** number; scoped to the gate's two trees it is smaller.

**Cluster / target identified:** gate clause 5 — *KB-fidelity audit GREEN, or YELLOW with 0 blockers, over
`corpus/services/**` + `corpus/architecture/**`.*

**Hypothesis:** the correction is **not** "delete every router mention". The router has two states — deleted
locally at `2adcf71`, still declared in prod — and the map landed at iter-20 is the reference that makes the
sweep mechanical rather than a judgement call per line.

**Expected lift:** gate 3 of 5 → 4 of 5.

**Phase plan:** measure the blast radius → sweep with a scripted, enumerated edit list (each edit asserted to
match **exactly once**, so a silent miss is impossible) → re-measure with a real KB-fidelity audit → fix →
re-audit until the verdict is claimable.

**Escalation conditions:** a second platform commit fires the re-scope trigger. A merge conflict against the
milestone branch is a user-blocker **unless** it is resolvable as a pure union of two correct texts.

**Acceptable close-no-lift outcomes:** an audit verdict that stays RED is a complete outcome if the residual
is enumerated well enough that the next iter is mechanical.
