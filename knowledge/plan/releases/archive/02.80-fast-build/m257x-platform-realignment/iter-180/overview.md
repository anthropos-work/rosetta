---
iter: 180
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-09
---

# iter-180 — a prose claim about SETS is a claim a machine can grade

**Type:** tik. **Active strategy: `TOK-08`** — *census the mechanical classes; stop sampling them.*
The class here is small, exactly enumerable, and mechanically decidable once it is given a grammar:
**set-relation assertions living inside classification rationales.**

## Step 0 — re-survey before targeting (mandatory)

Re-surveyed at HEAD `4b60aa2` (iter-179's commit), rext at `dc171ca`, trees clean modulo the user's
`.claude/settings.json`.

| route | state at re-survey |
|---|---|
| `FIX-M257x-iter177-derivation-registry-decline-rationale-is-false` | **open, and re-derived TRUE at HEAD by this iter before targeting.** iter-177 measured it and deliberately did not land it — *"a second line of investigation, and an edit to that file after the suite ran would have invalidated the run this close reports."* A clean Fate-1 target now. |
| `FIX-M257x-iter174-…-one-registry-of-two` | **closed at iter-179.** |
| `SURVEY-M257x-iter175-census-vs-discover_fences-classified-differently` | open — and this iter's measurement sharpens it rather than settling it. |

**Target: the false rationale, and the class it belongs to.**

## Cluster / target identified

`derivation_registry.py:224-227` declines two derivations with **one shared sentence**:

> *"each returns the same population `census` does (modulo `CENSUS_EXCLUSIONS`)"*

Re-derived independently at HEAD, four sets over the same population:

| derivation | size | relation to `census` |
|---|---|---|
| `guard_family::union` | **27** | `census ∪ CENSUS_EXCLUSIONS` — **exactly**. The sentence is TRUE here. |
| `guard_family::census` | 26 | — |
| `guard_family::declaring_modules` | **26** | differs by **two** members (`guard_family` in, `repair_postcondition` out); adding the exclusions back gives `union` minus `repair_postcondition`. **The sentence is FALSE here, in both directions.** |
| `repair_postcondition::discover_fences` | **26** | **identical to `declaring_modules`** — and it is REGISTERED while `declaring_modules` is declined *for being the same as `census`*, which it is not. |

`CENSUS_EXCLUSIONS = {guard_family}`. The counts of `census` and `declaring_modules` are **equal**, which
is why nothing caught it: every count-based comparison of the two reads green (the hazard iter-177's own
mutation control already characterises).

## Hypothesis

The repair is not "fix the sentence". A rationale that asserts a set relation is **executable**, and the
population that asserts one is small enough to enumerate: measured at HEAD, **2 of 76** entries name a
sibling derivation in backticks — and they are exactly the two under review. So the class can be closed
by giving the claim a grammar and grading it live, rather than by rewriting prose that will rot again.

**Expected to be wrong about:** whether a generic operand resolver (`module::attr`, called if callable,
flattened to a set of strings) actually covers both sites without a lookup table — `discover_fences`
returns a *pair* of lists, so the resolver must flatten, and if that turns out to need a per-site rule
the fence becomes a registry and the design is worse than the prose.

## Expected lift

No `P`/`N` reading is taken; **no clause-5 movement is claimed** (`§9` — UNMEASURED is not unmoved). The
deliverable is a corrected rationale **plus** the fence that makes the correction checkable: every
rationale naming a sibling derivation must carry a machine-gradeable `RELATION:` clause, and every
`RELATION:` clause must hold live, both directions (a clause naming no site is RED too).

## Phase plan

1. Re-derive the four sets at HEAD (done above — recorded before any edit).
2. Census the population of rationales asserting a sibling relation; state the denominator.
3. Give the claim a grammar; write the resolver; grade both sites live.
4. Correct the false rationale to what is measured, per function rather than one sentence for two.
5. Controls: anti-vacuity (population ≥ 2, both operands non-empty) + mutation controls that fire.
6. Protocol doc §8 rule if the lesson generalises.

## Escalation conditions

- If the resolver needs a per-site lookup table, **stop and keep the prose repair only** — a fence that
  is itself a registry is the tax iter-178 declined to pay, and the class is 2 sites wide.
- If correcting the rationale turns the frozen-expectation census RED, re-derive the cause before
  touching the derivation (`D-M257x-174-3`).

## Acceptable close-no-lift outcomes

Landing only the corrected rationale — with the four-set measurement recorded and the fence shown to be
unaffordable — is a complete iter, provided the falsification is written down.
