---
iteration_type: tik
status: in-flight
active_strategy: TOK-08
---

# iter-211 — four spellings of one source set; three agree at 114 by coincidence and one is 92

**Active strategy reference:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07).

## Step 0 — re-survey, and it RETRACTS one of iter-210's own routes

iter-210 routed `SURVEY-M257x-iter210-five-fences-scan-only-the-two-root-docs`, on the strength of five
modules declaring `SCAN_FILES = ("CLAUDE.md", "README.md")`. Re-surveyed with three lines of context
instead of one, **that route is false and is retracted here**:

```python
SCAN_GLOBS = ("corpus/**/*.md", ".claude/skills/**/*.md")     # markdown_structure_guard,
SCAN_FILES = ("CLAUDE.md", "README.md")                       # anchor_construct_guard,
                                                              # claim_twin_guard, repair_leak_guard
SCAN_ROOTS = ("corpus", ".claude")                            # platform_predicate_guard
SCAN_FILES = ("CLAUDE.md", "README.md")
```

`SCAN_FILES` is the **second half** of a two-part scope, never the whole of it. Reading one constant and
inferring a fence's subject is the same error as iter-208's adjective — and it was committed one iter
after that finding.

**And the correction inverts the story of iters 209–210.** Four fences have scanned
`.claude/skills/**` all along. `corpus_citation_guard` was not the family widening its scope at
iter-209; it was **the outlier catching up**, and `repair_leak_guard` had already written down the
hazard beside its own declaration:

> *fences ask different questions about the same surface, and a scope that drifted between them would
> make "one fence is silent" ambiguous between "clean" and "not looking".*

The risk was declared. Nothing checked it.

## Cluster / target identified — four spellings, resolved

| shape | fences | documents |
|---|---|---|
| **A** `SCAN_GLOBS` + `SCAN_FILES` | `markdown_structure_guard`, `anchor_construct_guard`, `claim_twin_guard`, `repair_leak_guard` | **114** |
| **B** `SCAN_ROOTS` + `SCAN_FILES` | `platform_predicate_guard` | **114** |
| **C** `fence_provenance.corpus_sources` | `corpus_citation_guard`, `retracted_pin_guard` | **114** |
| **D** private `corpus/**` walk | `clone_drift_guard` | **92** |

A ≡ B ≡ C exactly today (all symmetric differences 0); D is corpus-only and short by the 22 documents
iter-210 measured and declared.

**Three agreeing derivations is not one derivation.** And their agreement is a coincidence of this tree:
A globs `.claude/skills/**/*.md` while C globs `.claude/skills/*/*.md`. They differ on any skill
document nested one level deeper, and no skill has one today.

## Hypothesis

The latent A-vs-C divergence is provable on a staged tree, and closing it costs nothing on the real one
(A − C = 0 today). The durable move is the same as iter-210's: **one derivation, and an arm over the
family's spellings rather than over their current agreement.**

## Expected lift

1. `SKILL_SOURCE_GLOB` becomes recursive so C matches A **by construction**, not by coincidence.
2. A family arm compares every spelling's resolved set, with D declared and sized.
3. iter-210's false route is retracted in place.

## Phase plan

1. Deepen `SKILL_SOURCE_GLOB`; re-measure both live fences (must not move findings).
2. Family-agreement arm + a staged nested-skill-doc mutation control that separates `*` from `**`.

## Escalation conditions

- If deepening the glob changes either fence's **findings** on the real tree, stop and route — the
  same stop condition iter-210 pre-registered.

## Acceptable close-no-lift outcomes

- Everything already agrees and the mutation control is the only new evidence → the family gains a
  comparison it never had, which is `§5` iter-175 applied to four derivations instead of two.
