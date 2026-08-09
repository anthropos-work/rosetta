# iter-212 — decisions

## `D-M257x-212-1` — the five fences are FOLDED, not compared

iter-211 left the five comparing their private scopes on every run and routed the fold. iter-212 lands
it: `markdown_structure_guard`, `anchor_construct_guard`, `claim_twin_guard` and `repair_leak_guard`
now `return fence_provenance.corpus_sources(repo_root)`; `value_change_guard` shares at one hop (it
delegates to `repair_leak_guard`). **The `SCAN_GLOBS` / `SCAN_ROOTS` / `SCAN_FILES` constants are
KEPT** — `repair_leak_guard` tests membership against `SCAN_FILES` and prints
`scope=(SCAN_GLOBS + SCAN_FILES)` in its refusal line, so they are still operator-facing declarations
and iter-211's three-spelling arms still fence them against the shared set. A constant that no longer
derives anything but is still quoted must still be true.

## `D-M257x-212-2` — the wider member keeps its extra, DECLARED and CALLABLE

`platform_predicate_guard` declares `SCAN_ROOTS = ("corpus", ".claude")` — **the whole harness
directory**, a strict superset of the other four's `.claude/skills/**/*.md`. Folding it to
`corpus_sources()` alone would have been a **NARROWING of a shipped fence** disguised as a
de-duplication. Two options were live; the third was taken:

1. widen `corpus_sources()` to `.claude/**` — **rejected.** iter-209's stated precondition for widening
   the shared set at all was *zero false REDs measured on the added documents*; the added set is empty
   today, so that precondition is **unmeasurable**, not satisfied.
2. leave it private — rejected, that is the defect.
3. **taken:** it calls `corpus_sources()` **plus** a new named derivation
   `fence_provenance.claude_docs_outside_skills()`, whose size is **printed on every run** (0 today)
   while the **shape** — *this fence's scope is exactly `corpus_sources() | its declared extra`* — is
   **asserted**, and survives the size changing. `§5`: *print the SIZE, assert the SHAPE.*

## `D-M257x-212-3` — the census enumerates by EFFECT; the literal arm is kept, and NEITHER is complete

The new class calls every discovered collector and compares the returned SET, which no spelling can
evade. It does **not** replace iter-210's literal-string arm, because the two have **different blind
spots** and the union is the population:

- name-based discovery cannot see `clone_drift_guard` (an inline walk, no collector function);
- string-based matching could not see the five (`glob`/`rglob("*")`, never `rglob("*.md")`).

Each blind spot is declared with its reason and **reconciled in both directions** — an undeclared
literal-only module fails, and a declared one the census later reaches fails too.

**This arm's own first draft asserted the census was a superset of the literal arm and went RED on its
author within a minute.** Left in place as the arm's docstring rather than quietly corrected: the same
class as iter-210's case-sensitivity miss and iter-205's, one detector over.
