**Type:** tik — under [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07).

# iter-211 — four spellings of one source set; three agree at 114 by coincidence, and iter-210's route was wrong

## RETRACTION — `SURVEY-M257x-iter210-five-fences-scan-only-the-two-root-docs` is false

iter-210 read `SCAN_FILES = ("CLAUDE.md", "README.md")` in five modules and routed forward that those
fences *"scan only the two root docs"*. Re-surveyed with three lines of context instead of one:

```python
SCAN_GLOBS = ("corpus/**/*.md", ".claude/skills/**/*.md")   # markdown_structure_guard,
SCAN_FILES = ("CLAUDE.md", "README.md")                     # anchor_construct_guard,
                                                            # claim_twin_guard, repair_leak_guard
SCAN_ROOTS = ("corpus", ".claude")                          # platform_predicate_guard
SCAN_FILES = ("CLAUDE.md", "README.md")
```

**`SCAN_FILES` is the second half of a two-part scope, never the whole of it.** Inferring a fence's
subject from one constant is the same error as iter-208's inserted adjective, committed one iter after
that finding and by the same session.

**And the correction inverts the story of iters 209–210.** Four fences have globbed
`.claude/skills/**` all along. `corpus_citation_guard` was never the family widening its scope at
iter-209 — it was **the outlier catching up.** `repair_leak_guard` had already written the hazard down,
in the comment directly above its own declaration:

> *fences ask different questions about the same surface, and a scope that drifted between them would
> make "one fence is silent" ambiguous between "clean" and "not looking".*

The risk was **declared and unchecked**, which is `§5` iter-189's rule (*a stated-but-unfenced rule is a
comment*) landing on the family that keeps recording it.

## The four spellings, resolved on the real tree

| shape | fences | documents |
|---|---|---|
| **A** `SCAN_GLOBS` + `SCAN_FILES` | `markdown_structure_guard`, `anchor_construct_guard`, `claim_twin_guard`, `repair_leak_guard` | **114** |
| **B** `SCAN_ROOTS` + `SCAN_FILES` | `platform_predicate_guard` | **114** |
| **C** `fence_provenance.corpus_sources` | `corpus_citation_guard`, `retracted_pin_guard` | **114** |
| **D** private `corpus/**` walk | `clone_drift_guard` | **92** (declared + sized at iter-210) |

A ≡ B ≡ C exactly — every symmetric difference **0**. **Three agreeing derivations are not one
derivation**, and here the agreement is a property of today's tree rather than of the code: **A globs
`.claude/skills/**/*.md` and C globbed `.claude/skills/*/*.md`.** They differ on any skill document
nested one directory deeper, and no skill has one today.

## What shipped

- `SKILL_SOURCE_GLOB` is **recursive** (`.claude/skills/**/*.md`), so C matches A **by construction**.
  Re-measured on the real tree: sources **114 → 114**, citation findings **0 → 0**, retracted-pin
  findings **3 → 3**, pins **2,201 → 2,201**. The pre-registered stop condition — a *findings* change —
  did not fire.
- Two arms in `tests/test_corpus_citation_guard.py` (`TheFAMILYSpellsOneScopeFourWays`): all three live
  spellings resolve to one set, expanded independently from the constants each family member declares;
  and a **staged mutation control that separates `*` from `**`** — a skill document one level deeper,
  which the one-level glob provably misses and the shipped glob must see.
- The retraction is written into the arm's own docstring, so the false route cannot be re-derived from
  the code that occasioned it.

## Close — 2026-08-09

**Outcome:** the family spells its source set four ways; three resolve to the same 114 documents **by
coincidence of this tree** and are now equal by construction and fenced against each other; the fourth
stays declared at 92. iter-210's route — that five fences read only the two root documents — is
**retracted**: they pair `SCAN_FILES` with `SCAN_GLOBS`/`SCAN_ROOTS`, four of them have read
`.claude/skills/**` all along, and iter-209's widening was the outlier rejoining the family rather than
the family being widened.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (forty-third consecutive `closed-fixed`;
**no `P`/`N` reading taken — UNMEASURED, not unmoved**, `§9`) — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: **y** — **counted, not felt: iters 207, 208, 209, 210, 211 = five tiks this run
against a cap of five** — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **exit-5**
**Decisions:** `D-M257x-211-1` … `D-M257x-211-2` (see [`decisions.md`](decisions.md))

**Audit:** `/usr/bin/python3 -m pytest` (**8.4.2** / CPython **3.9.6**), **Python**, `stack-core` only —
**156 passed** across `test_corpus_citation_guard` + `test_retracted_pin_guard` + `test_guard_family` +
`test_clone_drift_guard`; **78 passed** on the two changed-fence modules alone. Both live guards on the
real rosetta tree: citation **0 findings / 114 sources**; retracted-pin **3 findings / 114 sources /
2,201 pins**.
**RED-proof battery, mtime-mitigated (`§5` r77):** `SKILL_SOURCE_GLOB` reverted to one level → **the
depth arm RED**, the three-spelling equality arm green (correct — they still agree on this tree, which
is exactly why the depth control is needed). Restore sha-verified.
*Scope, stated rather than implied (`§5` r60): `stack-core` only, Python only, changed-code reach. No
whole-section run this run — the tree was edited throughout all five iters. No Go, no TypeScript; the
four non-`stack-core` Python sections were read at iter-208 and not re-read since.*

**Side-deliverables:** none this iter.

**Routes carried forward:**
- `SURVEY-M257x-iter210-five-fences-scan-only-the-two-root-docs` — **RETRACTED, false as written.**
  Superseded by the four-spelling table above.
- `SURVEY-M257x-iter211-A-and-B-still-spell-their-own-scope` — **NEW.** Five fences still expand their
  scope from private constants rather than calling `fence_provenance.corpus_sources`. They are now
  **compared** on every run, which is the property iter-210 argued for, but they are not yet **shared**.
  Routing them through the shared derivation is a five-module change with five behaviour surfaces and
  wants its own iter.
- `SURVEY-M257x-iter210-clone-drift-reads-a-third-corpus` — unchanged, declared and sized (+22 docs /
  +73 sha tokens / +5.0 %).
- All routes from iters 207–209, unchanged, plus the standing queue.

**Lessons:**
- **Read the whole declaration before routing a claim about it.** One constant of a two-constant scope
  produced a route that was false in both directions — it understated four fences' reach and inverted
  which member of the family was the outlier.
- **Three agreeing derivations are not one derivation.** Their agreement here was a fact about the
  absence of a nested file, not about the code; only a staged tree could tell them apart.
- **A declared hazard with no check is a comment.** `repair_leak_guard` wrote down the exact failure
  mode iters 209–211 then lived through.
