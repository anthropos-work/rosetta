---
iter: 172
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-08
---

# iter-172 — the census counted two different things and printed them in one column

## Step 0 — re-survey before targeting

`TOK-08` still controls: *census the mechanical classes; stop sampling them.* iter-171 closed the last
runner disagreement and opened exactly one route:

> `SURVEY-M257x-iter171-runner-test-count-gap` — the same 35 modules yield **1073** tests under unittest and
> **1062** under pytest. *"Pre-existing, unexplained, and by rule 75's own logic an unstated scope: 11 tests
> execute under one runner and not the other, and nobody has named which."*

Re-surveyed at open: still open, and **its stated reading is a hypothesis that has not been tested.**
Target confirmed.

**Active strategy reference:** `TOK-08`. The subject is the milestone's own **instrument** — `suite_census.py`,
shipped at iter-170 and already relied on twice — which is where `§9` says the sharpest defects live.

## Cluster / target identified

iter-171 published a number *and* an interpretation of it in the same breath. The number (11) is real. The
interpretation — *"11 tests execute under one runner and not the other"* — assumes both columns count the
same thing. **That assumption is exactly the class this milestone keeps finding**, one layer in: rule 75
made the milestone state its runner; it did not make the instrument state its **unit**.

## Hypothesis

The gap is **not** a population difference. It is the instrument: the two branches of `run_one` derive
`tests` from different regexes over different summary grammars, and those grammars do not denote the same
set.

## Expected lift

No `P`/`N` reading (`TOK-08` puts the reading after the sweep; `§9`'s **UNMEASURED-is-not-unmoved**
refinement applies and the close will say so in those words).

Deliverable: the gap **decomposed to zero residue** — every one of the 11 attributed to a named module and a
named cause — the counter repaired so both runners denote the same set, a fence that fails when they
diverge again, and **iter-171's published interpretation corrected in place** if it is refuted.

## Phase plan

(A) decompose the 11 per-module, per-runner. (B) read the two parsers and name the unit each denotes.
(C) repair so both denote the same set. (D) fence with a synthetic module carrying one pass, one failure and
one skip — the minimal input on which a "passed-only" counter and a "ran" counter must disagree, so the
control fires by construction. (E) re-measure; the per-module diff must be **0** with the runner named.
(F) correct the iter-171 route in place, and close.

## Escalation conditions

- If the decomposition leaves **any** residue, the population reading is not refuted and must not be
  reported as if it were — say what is unexplained and route it.
- A repair that makes the two columns agree by *dropping* information (e.g. counting neither skips nor
  failures) is a flattering fix; `§5` refuses those. The counter must denote the larger, honest set.

## Acceptable close-no-lift outcomes

If the 11 turn out to be a genuine population difference after all, that is iter-171's reading confirmed and
this iter closes having *tested* it rather than repeated it — a first-class outcome under the protocol.
