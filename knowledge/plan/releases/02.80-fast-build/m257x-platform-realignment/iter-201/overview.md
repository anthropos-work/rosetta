---
iter: 201
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-09
---

# iter-201 — the milestone's last open FIX route: the module one runner could not see

**Active strategy reference:** `TOK-08` — *census the mechanical classes; stop sampling them.* This iter
is the other half of the pattern: **once a class is censused to a single named member, close it.**

## Step 0 — re-survey (mandatory)

Re-surveyed at HEAD `aa1e367`. Every harden-filed route from passes 45–47 is now closed —
`h42` at iter-200, `h45` at iter-199, `h46` at iter-198 — leaving the milestone's **only open `FIX-`
route**, and the one iter-197 sized:

> `FIX-M257x-h44-claim-census-guard-is-single-runner` — *"Convert `test_claim_census_guard.py`'s 25
> pytest-style functions to `TestCase` so both runners collect them, **carefully**: iter-182 measured that
> the obvious translation loses tests silently."*

The last **measured** reading of the class is iter-197's, at corpus `592d583`: **123 modules · pytest
3,551 · unittest 3,526 · gap 25 · 1 module DISAGREES**, style `bare 1 · testcase 122`. Iters 198–200 each
added a test module, so the module count and totals have grown since; **no pre-conversion re-reading was
taken at iter-201's open, and none is reconstructed here** — a subtraction back from the post-conversion
figure would be exactly the agreeing reconstruction iter-197 was written to stop. What is asserted about
the starting state is only what iter-197 measured.

## Cluster / target identified

One module, 25 cases, 3 fixture APIs (`monkeypatch.setattr` ×6, `capsys.readouterr` ×6, `pytest.skip`
×2). Measured before planning, because the fixture surface is what decides whether this is an hour or a
day.

## Hypothesis

The cases do not need rewriting — only re-binding. Renaming them `_case_*` (so neither runner
double-collects) and **generating** the `TestCase` methods from the module namespace makes iter-182's
silent-loss failure mode unrepresentable rather than merely avoided.

## Expected lift

`gap 0 · 0 modules disagree`, with all 25 cases still running and the count asserted from two
independent derivations.

## Phase plan

1. Rename the 25 cases; shim the three fixture APIs; drop the `pytest` import entirely.
2. Generate the `TestCase` methods from `CASES`.
3. Lost-test controls: namespace-vs-class both ways, and an independent count from the file source.
4. Re-run the census; empty `RUNNER_COLLECTION_SPLIT` and repair iter-197's now-vacuous arms.

## Escalation conditions

- Any case behaving differently under the shims → stop and report; do not adjust a case body to fit.
- Fewer than 25 bound → the exact failure iter-182 named; stop.

## Explicitly NOT in scope

Changing any of the 25 case bodies. The conversion is a binding change, and if it needed a body change
it would not be one.
