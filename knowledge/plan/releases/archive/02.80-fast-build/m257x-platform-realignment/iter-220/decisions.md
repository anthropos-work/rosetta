# iter-220 — decisions

## `D-M257x-220-1` — direction B is landed; direction A is not, and the reason is a denominator

A README is a registry with two directions. **B** (a cited file exists) is decidable from the filesystem
and is now enumerated on every run: **97** citations across **32** READMEs, **0** dangling. **A** (a file
that ought to be cited is cited) needs a rule for *ought*, and iter-179 already measured the obvious one
as wrong — *"the `10 of 63` you get from all test modules is the wrong ratio"*. So A is **printed**
(12 of 76) and its **shape** asserted (strict subset, non-empty gap), never its size. Landing a number
whose denominator nobody has justified is the defect this milestone keeps finding, not a partial fix.

## `D-M257x-220-2` — the probe's own false RED is KEPT as the fence's scope control

The first probe checked each citation against `stack-core/` alone and flagged `exposure_claim_guard.py`,
which is real and lives in `stack-injection/`. Deleting that mistake would have left a census whose only
output is a zero and no evidence it can produce anything else. Instead it is `test_02`, asserting **both**
halves — the narrow pool fires, and everything it names is real — so the arm goes RED either if the
matcher becomes inert or if a genuinely dangling citation appears and `test_01` misses it. `§9`: *a
census returning ZERO must prove its instrument*, and the honest proof was already in hand.

## `D-M257x-220-3` — this is iters 209 and 214's class, committed a fourth time by the same session

*A scoped reading is evidence about its scope alone* — written down at iter-209, again at iter-214,
again in iter-219's own close, and reproduced here inside ten minutes. The lesson recorded is not
*"remember the scope"*, which has now failed four times; it is **structural**: the scope must be an
argument the instrument's own control varies, so getting it wrong shows up as a test result rather than
as a reading nobody re-checks.
