---
iter: 186
milestone: M257x
iteration_type: tik
status: in-flight
opened: 2026-08-09
---

# iter-186 — "the whole-population baseline" was 5 of 11 sections, and nothing said so

**Active strategy reference:** `TOK-08`. Third consecutive member of the class iter-184 named
(*a fence's POPULATION is a registry too*), reached through iter-185's routed
`SURVEY-M257x-iter185-other-declared-populations-unaudited`.

## Step 0 — re-survey before targeting

Rather than sweeping by hand a third time, the class was **enumerated**: 70 module-level string
collections exist in `stack-core`. Sorted into population-defining vs predicate-defining, the sharpest
hit is a **pair naming the same population with different cardinality** — iter-177's shape:

| literal | n | ground truth |
|---|---|---|
| `claim_census_guard.REXT_SECTION_NAMES` | 11 | matches the 11 non-`knowledge` dirs exactly ✔ |
| **`suite_census.SECTIONS`** | **5** | **6 sections omitted, and they are not empty** |

## Cluster / target identified

`suite_census.py` is the instrument behind this milestone's **whole-population suite baseline** — the
`3,369 passed · 9 failed · 4 skipped` figure quoted in briefs and closes. Its section tuple names 5 of
11, and the six omitted carry **264 Go test files + 45 TypeScript specs**.

## Hypothesis

The omission is **correct** (no Python runner collects a Go test) and **unstated**, which is the defect.
The repair is a derivation plus a named, reasoned exclusion — not a widening.

## Expected lift

No `P`/`N` reading. Deliverable: `SECTIONS` derived from disk, the six excluded **by name with a reason**,
the scope printed with every total, and a fence asserting the two halves partition the filesystem
exactly. **`D-M257x-145-3` is NOT ruled on** — it is the user's — but it gets the denominator it lacked.

## Phase plan

**A — enumerate the class · B — derive + name the exclusions · C — fence, mutation-proven · D — close.**

## Phase 0d — pre-flight tooling check (RUN)

A new `tests/` module, no new `*_guard.py` — so no `INVOCATIONS` entry and no movement of the README
fence-index triple. iter-179 measured that a `tests/` module is in none of the four fence registries'
populations; re-checked at close by running both registry guards rather than assumed.

## Escalation conditions

- If any excluded section carries **Python** tests, the exclusion is wrong rather than merely unstated,
  and the census has been under-reporting its own language.
- Nothing in this iter may state what "the suite" means. That is `D-M257x-145-3`, pending the user.

## Acceptable close-no-lift outcomes

If the six omitted sections carried no tests at all, the iter still closes complete: the population is
derived, the partition asserted, and the null result published with its number.
