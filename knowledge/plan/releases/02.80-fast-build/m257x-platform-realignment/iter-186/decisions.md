# iter-186 — decisions

## `D-M257x-186-1` — the class was ENUMERATED this time, not swept by hand

iters 184 and 185 each found a member of *a fence's population is a registry too* by judgement, and
iter-185 routed the honest consequence: **nobody knows how many exist.** This iter enumerated instead —
**70 module-level string collections** in `stack-core` — and picked its target from the enumeration.

The selector that worked: **two literals naming the same population with different cardinality**
(iter-177's shape). `claim_census_guard.REXT_SECTION_NAMES` says **11**; `suite_census.SECTIONS` says
**5**. Both are about "the sections of `rosetta-extensions`", and nothing had ever compared them.

## `D-M257x-186-2` — the omission is CORRECT; the silence is the defect

The six sections `suite_census` omits are not empty. Measured 2026-08-09:

| section | non-Python tests |
|---|---|
| `stack-seeding` | **119** `*_test.go` |
| `stack-snapshot` | 45 |
| `clerkenstein` | 37 |
| `playthroughs` | 22 `*_test.go` + **45** `*.spec.ts` |
| `alignment` | 21 |
| `stack-secrets` | 20 |

**264 Go test files and 45 TypeScript specs.** No Python runner can collect any of them, so excluding
them is right. What was wrong is that the tuple said none of it, while this milestone has quoted the
resulting total as **"the whole-population baseline"** — in run 17's brief, and in iter closes. It is
the whole **Python** population, over **5 of 11** sections.

**Decision:** derive `SECTIONS` from disk; exclude the six **by name, each with its reason**
(`§5` rule 8); print the scope with every total (`§5` rule 60 — a scoped green is evidence about its
scope alone, and this one had not been saying its scope).

## `D-M257x-186-3` — this supplies `D-M257x-145-3`'s missing denominator and does NOT rule on it

`D-M257x-145-3` is the standing assumption *"the suite" = all five `rosetta-extensions` sections*,
pending the user's ruling, with a second axis (*which interpreter*). It is now measurable that **five is
not all** — the repo has eleven — so the assumption's own wording contained an unmeasured claim.

**That is evidence for the ruling, not the ruling.** Nothing in this iter states what "the suite" means,
and the assumption stays open exactly as the brief requires. What changed is that the question can no
longer be discussed without its denominator.

## `D-M257x-186-4` — an exclusion must be checkable, or it is a place to hide a live section

Four arms, two mutants, both RED-proven with three failures each:

- **every section on disk is collected or excluded by name** — dropping `stack-seeding` from the
  exclusion map fires 3 arms naming it;
- **no accounted section is missing from disk** — a stale exclusion is how a live section gets dropped
  silently, the name and reason surviving the subject;
- **every excluded section is non-empty in its own language** — an exclusion whose subject has no tests
  is either stale or was never true, and both read as green. Adding `knowledge` with a plausible reason
  fires 3 arms;
- **the two halves partition the disk exactly** — an overlap or a gap makes the split unreadable, which
  is what *"the whole-population baseline"* meant for sixteen iters.

Plus the control that keeps the derivation honest: `derive_sections` must read the disk, and the
last-measured tuple must be reachable **only** through the cannot-locate-the-repo path. A silent
fallback is how a derivation becomes a literal again.
