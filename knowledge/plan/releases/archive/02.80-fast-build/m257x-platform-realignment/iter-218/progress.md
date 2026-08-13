**Type:** tik · **Active strategy:** [`TOK-08`](../decisions.md) — *census the mechanical classes; stop
sampling them.*

# iter-218 — iter-217 taught ONE of the repo's number-matchers to read bold; six others never learned

## The finding

iter-217 closed the literal censuses' emphasis hole by widening `derivation_registry._MEASURED_RE`.
Re-surveyed immediately: **`_MEASURED_RE` is not this repo's only number-matcher.** A census **by
effect** — import every `stack-core` module, walk its module-level compiled patterns, probe each with a
plain figure and the same figure bolded — returns **seven** constructs, and the split is the finding:

> **Every emphasis-aware construct lives in ONE module** (`derived_count_guard`). The blind ones live in
> **three different modules**, which is why no reader had ever seen them as one class.

Iters 209–212's shape, one construct family over — and **iter-217's own repair, one iter late.**

## ⚠️ The census corrected its own author inside a minute

The iter opened with a **hand-enumeration sealed at `0f991bc` naming SIX constructs across four
modules.** The by-effect probe returned a **seventh** — `suite_census.RE_PYTEST_TERM` — within a minute
of existing, in the iter written to demonstrate that hand-enumeration is the wrong instrument.
Corrected in place, appended and not substituted. **That is `TOK-08` in one line: a reading SAMPLES, a
fence CENSUSES.**

## The target: G2 repo-count — the fence closest to this milestone's own subject

`platform_predicate_guard._REPO_COUNT` grades every corpus sentence claiming how many repos the clone
set has, and `repos.yml` membership moved **three times** inside M257x's own window. Measured over
`corpus_sources`' **114** documents inside G2's own `_CLONE_CONTEXT` window: **21 raw matches, 2
emphasis-blind** — and the two point in **opposite directions**, which is the whole reason the repair is
two rules rather than one:

| site | verdict |
|---|---|
| `corpus/ops/update_guide.md:96` — *"updates the **4** repos defined in `repos.yml`"* | **in scope**, and exactly what G2 exists to grade. TRUE today, so a **LATENT** false-GREEN, not a live one |
| `corpus/architecture/org-repos.md:71` — *"sources **9** service repos' `//terraform` modules"* | **out of scope**, and widening alone makes it a **FALSE RED** (9 ≠ 4). **The blind spot was hiding a false RED as well as a false GREEN** |

### The narrowing that does not work, refuted before it was written

Requiring the clone context on the **same line** instead of a ±4-line window separates nothing — line 71
itself says *"does not clone it"* — and costs **21 → 14** live matches.

### The rule that does, and its price

Excluding a **possessive** `repos'` — *the number counts the repos whose modules are sourced, not the
clone set* — costs **0** live matches (21 → 21), leaves the widened matcher adding **exactly 1** claim,
and that claim grades **GREEN**. **Zero false REDs**, which was iter-209's precondition on widening any
shipped guard and is not negotiable.

**Stop condition HELD, verified on both trees:** G2 reach **15 → 16** repo-count claims;
`platform_predicate_guard: OK` before and after (pre-repair tree unpacked with `git archive`, so the
matcher was the only variable).

## The other blind matchers, disposed of by name with their sizes

- `guard_family._STATED_COUNT` — reads a guard's own **stdout**. Live exposure **0** printed emphasised
  finding-counts; the 4 textual hits are prose *about* counts, which it never reads. **Declared.**
- `suite_census._DECLARED_GO_COUNT_RE` — reads this module's own exclusion-registry `reason` strings:
  **0 of 21** carry emphasis. **Declared.**
- `suite_census.RE_PYTEST_TERM` — parses pytest's own summary line. Machine output. **Declared.**
- `derived_count_guard._N_OF_M` — **not blind by this census's measure at all**, and the unit matters:
  it already reads `**28 of 29**`. It is blind to the **split** spelling `28** of **29`, whose live
  population across the guard's own roots is exactly **one** — `hardening-ledger.md:4766`, the harden's
  own `5 of 51`. **ROUTED whole, not smuggled in as a waiver**: declaring it blind here would have
  graded it against a probe set that never tests it, and widening it turns the guard RED until a
  `N_OF_M_DISPOSITIONS` entry is written for the claim it surfaces.

## Scope, stated rather than implied (`§5` r60)

`/usr/bin/python3 -m pytest` (**pytest 8.4.2 / CPython 3.9.6**), **Python**, `stack-core` only, changed-code
reach: **290 passed / 0 failed** across `test_frozen_expectation_census_m257x`, `test_m257x_emphasis_family`,
`test_m257x_emphasis_reach` and `test_platform_predicate_guard` (43 s). `derivation_registry --ceilings`
exits **0**, all three `exact +0` after re-pinning **175 → 178** (the `_REPO_COUNT` widening's own recorded
reason, quoting the figures it was measured from — the same mechanism iter-217 recorded, one iter later)
and **542 → 559** (this iter's own arms). **No whole-section run** — the tree was edited throughout. No
Go, no TypeScript, no non-`stack-core` Python section.

## Close — 2026-08-09

**Outcome:** the number-matcher family is enumerated **by effect** rather than by reading, and the split
that hid it — every emphasis-aware construct in one module, the blind ones scattered across three — is
now fenced so a matcher added tomorrow enrols itself. G2, the repo-count fence, reads this repo's bold
figures for the first time (**15 → 16** claims) and declines the possessive that would have made a
correct sentence a false RED. **Zero false REDs; verdict `OK` before and after.**
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (fiftieth consecutive `closed-fixed`;
**no `P`/`N` reading taken — UNMEASURED, not unmoved**, `§9`) — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n — **counted, not felt: iters 217, 218 = two tiks this run against a cap of five** —
(6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-218-1` … `D-M257x-218-4` (see [`decisions.md`](decisions.md))

**Side-deliverables:** none this iter.

**Routes carried forward:**
- `SURVEY-M257x-iter218-n-of-m-split-emphasis` — **NEW.** `28** of **29` is invisible to
  `derived_count_guard._N_OF_M`; live population **1** (`hardening-ledger.md:4766`, the harden's own
  `5 of 51`, undispositioned). Landing it needs a `N_OF_M_DISPOSITIONS` entry or the guard goes RED.
- `SURVEY-M257x-iter217-the-ratchets-have-no-pre-edit-whole-section-reading` — unchanged.
- All routes from iters 207–216 unchanged, plus the standing queue.

**Lessons:**
- **A census by EFFECT beats an enumeration by NAME, and this iter is its own proof** — the sealed
  hand-count said six and the probe said seven, on the same author, in the same hour.
- **Widen and scope in ONE edit when the blind spot hides errors in both directions.** Here the same
  two characters that would have surfaced a true claim would have flagged a correct one.
- **A construct blind to a spelling your probe set never tests is not "blind by design"** — it is
  ungraded, and calling it a declared exception would have parked it where nobody looks.
