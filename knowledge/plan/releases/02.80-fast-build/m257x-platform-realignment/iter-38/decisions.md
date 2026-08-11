# iter-38 — decisions

## D-M257x-38-1: run the pass WIDE, against the routed instruction — rule 18 weights, it does not narrow

**Choice.** All 40 in-scope files read top-to-bottom by six auditors, with the 8 previously-repaired files
weighted (a dedicated auditor per one or two of them). The routed instruction said *"scope the next pass to
the 9 changed files."*

**Why.** Two reasons, and the second is the load-bearing one.

1. Clause 5's own wording asks for an audit *"over `corpus/services/**` + `corpus/architecture/**`."* A
   pass reading 8 of 40 files cannot return a verdict about 40, however well-aimed. iter-21 is this
   milestone's precedent: a scoped audit converged on a small number that a full read multiplied by five.
2. §5 rule 18's measurement is a **density** claim (0.69 vs 0.074 per file). Density licenses **where to
   spend effort**; it does not license **what to leave unread**. Those are different operations and the
   distinction is exactly what a "9× " number invites you to blur.

**Vindicated by the outcome, which is why it is recorded rather than assumed:** the density ratio
reproduced (~7.3×) *and* the untouched files still yielded **6 of 17** blockers — two of them in a file
already read cover-to-cover twice, one of them the most consequential claim in the corpus. Following the
route would have found 11 and declared the rest clean.

## D-M257x-38-2: retract the premise; do NOT re-classify

**Choice.** `security_compliance.md` and `ai_architecture.md` now state that the per-check verdicts are
LLM-produced and that the stated basis for the **EU AI Act Limited-Risk** classification does not hold at
platform HEAD. Neither doc asserts a different classification. Both say the re-derivation is a question for
counsel, and `CHECK-M257x-iter38-ai-act-classification` is routed with no owner inside this milestone.

**Why.** The corpus's defect was asserting a **legal conclusion** from a **technical premise** it had not
measured. Replacing it with a different legal conclusion — "therefore High Risk" — would repeat the defect
with the sign flipped, and this milestone has already had one correction that was wrong in the opposite
direction (the tenancy fence, twice). What the corpus can support is the measurement; it cannot support the
classification either way.

**And the retraction was itself over-stated on first cut.** The adversarial pass showed `EngineTextDiff`
checks run deterministically alongside the LLM ones, so *"the verdicts come from an LLM"* is false as a
universal. Both files now say **most**, not all — which is weaker, and true.

## D-M257x-38-3: cite the DISPATCH, never the registration map

**Choice.** The retraction's mechanism citation is `basevalidator/criterion.go:127 → :428` (the hardcoded
switch and the LLM checker it constructs), not `v3/validator/validator.go:60-61` (`checkerEngines`).

**Why.** `checkerEngines` is declared, populated, passed down and **never read** — verified by grepping
every occurrence. It is a perfectly convincing citation for a mechanism that does not run through it. On a
page whose entire purpose is to tell a reader *"re-derive this compliance claim"*, handing them a dead
field would have reproduced the original defect one level down: a claim that looks sourced and is not.

## D-M257x-38-4: clause 5 stays open

**Choice.** `Gate: NOT MET`, despite 23 blockers found and fixed.

**Why.** The clause is met by a **reading that returns zero**, not by a repair that clears its own findings.
A pass that fixes everything it finds tells you the repair worked; it tells you nothing about what the pass
missed — and this pass's own adversarial half found six defects the pass itself had just created. iters 33
and 34 both refused the clause on this ground. Holding the line costs one more iteration and is the only
thing that makes the eventual claim mean anything.
