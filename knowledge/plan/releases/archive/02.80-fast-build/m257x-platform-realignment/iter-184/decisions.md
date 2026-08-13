# iter-184 — decisions

## `D-M257x-184-1` — the glob denotes the EMPTY SET, and it was in every brief

iter-182 carried forward `` `SURVEY-M257x-iter181-*` ``. The id grammar is a greedy character class, so
what the fence enumerated was the stem `SURVEY-M257x-iter181-` — and, measured, **iter-181 created no
route at all**: its close names seven ids, every one pre-existing, and closes one of them. The glob
denotes **∅**, while reading to a human as a live cluster of iter-181 surveys. Run 17's orchestrator
brief quotes it.

`§5` rule 73 already says **a glob is not a derivation**. What was missing was a population to check it
against, and iter-183 built one. **Decision:** refuse any carried member that is not a well-formed route
id, name its site, and repair iter-182 by **striking** the glob rather than deleting it.

## `D-M257x-184-2` — strikethrough is the RETRACTION grammar, so the guard reads it instead of demanding a deletion

The obvious repair — delete the glob — rewrites a closed iter's published record, which this milestone
does not do. The corpus already has the idiom: `CLAUDE.md` writes `~~ai~~`, `~~authn~~`,
`~~GraphQL/Cosmo Router~~`, and this registry used it once before, at iter-29. So the guard learns it: an
id inside a `~~struck~~` span is **withdrawn, not carried**, the record stays legible, and the
withdrawal is legible beside it.

**With two safeguards, because a stripper that hides a live route is this fence's own failure mode.**
The count of struck spans **prints on every run** (a silent stripper is a way to hide ids from a
census), and the pattern is deliberately **not** `re.S`: a non-greedy DOTALL span would let two
unrelated `~~` marks pages apart swallow every id between them. Audited live: exactly **2** struck spans
exist and they withdraw exactly the glob and one `CHECK-` id from iter-29 — no live route.

## `D-M257x-184-3` — the fence shipped one iter ago read 57.8 % of its own subject, and the cause was a hand-maintained tuple

iter-183 declared the id KIND as `(FIX|SURVEY|SWEEP|PROBE|TASK)` and called the result **"the route
registry"**. Measured inside the carry-forward blocks themselves:

| kind | distinct ids | in iter-183's population? |
|---|---|---|
| `FIX` | 151 | yes |
| **`CHECK`** | **76** | **no** |
| `SURVEY` | 37 | yes |
| **`DOC`** | 28 | **no** |
| **`FENCE`** | 15 | **no** |
| **`MEASURE`** | 10 | **no** |
| **`DEF`** | 3 | **no** |
| **`HOST`** / **`REPOINT`** / **`READ`** | 2 / 2 / 2 | **no** |
| `SWEEP` | 1 | yes |
| `PROBE`, `TASK` | **0** | declared, never occur |

**189 of 327 — 57.8 %.** Two of the five declared kinds do not exist. This is `§2`'s hand-maintained
tuple returning **inside the fence built to stop registries rotting**, one iter after it shipped, and
nothing would have said so: the guard was green, its controls were green, and its census line printed a
confident denominator.

**Decision:** derive the kind (`[A-Z]{3,10}-M<n>-…`) so there is no tuple to maintain. Re-derived, the
registry is **312 routes / 1,156 dispositions / 36 closures / 0 contradictions** — the consistency
property holds on the wider population, so widening cost nothing but honesty. **Both readings are
correct about their own population and neither may be quoted without it** (iter-177).

## `D-M257x-184-4` — well-formed means NOT TRUNCATED, and nothing stronger

The first predicate demanded `<kind>-<milestone>-<origin>-<slug>` with a multi-part slug. Run against the
widened population it booked **`HOST-M257x-toolchain`** as malformed — a legitimate id whose kind simply
is not iter-scoped. **A rule tuned to the majority shape manufacturing a finding about the minority
shape** is the instrument-side error `D-M257x-122-4` describes, and it was caught only because the
population had just been widened to contain a counter-example.

So the predicate asserts exactly the defect and nothing else: **the id may not end in `-`**. A glob stem
fails; a hard-wrapped head fails; a one-word slug passes.
