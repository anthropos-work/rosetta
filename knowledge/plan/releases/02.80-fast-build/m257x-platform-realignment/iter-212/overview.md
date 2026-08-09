---
iter: 212
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-09
---

# iter-212 — the arm that enumerates the family enumerates ONE SPELLING of it

**Type:** tik
**Active strategy:** [`TOK-08`](../decisions.md) — *census the mechanical classes; stop sampling them.*
Class under work: **the fence family's own source-set derivation** — the class iters 209/210/211 have
been walking one member at a time.

## Cluster / target identified

iter-211 closed routing `SURVEY-M257x-iter211-A-and-B-still-spell-their-own-scope`: *"Five fences still
expand their scope from private constants rather than calling `fence_provenance.corpus_sources`. They are
now **compared** on every run … but they are not yet **shared**. Routing them through the shared
derivation is a five-module change with five behaviour surfaces and wants its own iter."*

**This is that iter.** The re-survey (Phase 1 Step 0) found two things the route did not predict, and the
second is the reason the route existed at all:

1. `fence_provenance.py:267-270` still ships **iter-210's retracted rationale** as the design
   justification for keeping the five separate — *"Those answer a DIFFERENT question — they fence those
   two documents' own prose."* iter-211 retracted exactly that claim. The retraction reached the
   milestone ledger, the journal and `progress.md`; **it did not reach the comment that ACTS on it.**
2. The arm iter-210 shipped to prevent a third fork — `test_only_ONE_module_spells_the_corpus_source_construct`
   — enumerates by **one literal string**, `CONSTRUCT = 'rglob("*.md")'`. Measured: it sees
   **2 modules** (`fence_provenance`, `clone_drift_guard`). The five fences derive the same set through
   `glob("corpus/**/*.md")` and `rglob("*")`+suffix, and are **invisible to it**. The fork it was written
   to catch was already present, five times, in a different spelling.

## Hypothesis

The family's source-set class is censusable **by effect** (call each collector and compare the returned
set) and is not censusable **by spelling**. Replacing the literal enumerator with a behavioural one, then
routing all five through the shared derivation, closes the class rather than sampling it again.

## Pre-registered, sealed in this iter's FIRST commit — before any repair

**Census taken at corpus `1bf2dde` / rext `a55559c`, by re-expanding each module's OWN constants
(not by calling its collector), `/usr/bin/python3`, `stack-core`, Python:**

| spelling | constants | n | Δ vs `fence_provenance.corpus_sources()` |
|---|---|---:|---|
| `fence_provenance.corpus_sources` | `CORPUS_DIR` + `EXTRA_SOURCES` + `SKILL_SOURCE_GLOB` | **114** | — |
| `markdown_structure_guard` | `SCAN_GLOBS` + `SCAN_FILES` | **114** | 0 |
| `anchor_construct_guard` | `SCAN_GLOBS` + `SCAN_FILES` | **114** | 0 |
| `claim_twin_guard` | `SCAN_GLOBS` + `SCAN_FILES` | **114** | 0 |
| `repair_leak_guard` | `SCAN_GLOBS` + `SCAN_FILES` | **114** | 0 |
| `platform_predicate_guard` | `SCAN_ROOTS` + `SCAN_SUFFIXES` + `SCAN_FILES` | **114** | 0 |

All **15** pairwise symmetric differences among the five are **0**.

**And the agreement of the fifth is BY ABSENCE, not by construction.** `SCAN_ROOTS = ("corpus",
".claude")` walks **all of `.claude`**; the other four glob `.claude/skills/**/*.md`. They agree only
because `.claude` today holds `settings.json`, `settings.local.json` and `skills/` — **0 markdown files
outside `skills/`**. `.claude/agents/*.md` and `.claude/commands/*.md` are standard locations in this
harness; the day one appears, the fifth diverges from the other four **silently**.

**Six claims, registered before deriving the repair:**

- **R1** — the literal arm sees **2** modules; the true speller population is **7**.
- **R2** — the five fences and the shared derivation return the **same 114** documents today.
- **R3** — `platform_predicate_guard`'s scope is a **strict superset** by rule, equal only by absence,
  and the extra is **0 documents** at this tree.
- **R4** — `fence_provenance.py:267-270` asserts a proposition iter-211 retracted.
- **R5** — routing all five through the shared derivation changes **no live fence's finding count**.
- **R6** — the sharing is provable only by a **staged** tree; a live tree cannot separate the two rules
  (R3 is why).

**STOP CONDITION, sealed before the repair:** if routing the five through `corpus_sources()` changes ANY
live fence's finding count on the real tree, **do not land the fold** — report the change, keep the
comparison arms, and route the fold to its own iter. (Same shape as iters 210/211's pre-registered
findings-change stop, which did not fire either time.)

## Expected lift

The class closes: one derivation for the whole family, enumerated **by behaviour**, with the wider
member's extra **declared and sized** rather than implicit. `§5` — *print the SIZE, assert the SHAPE.*

## Phase plan

A: seal this record (probe commit). B: repair the retracted comment. C: replace the literal enumerator
with a behavioural census. D: route the five through the shared derivation + declare the sized extra.
E: arms + mutation controls + staged depth control. F: re-run live fences against R5's stop condition.

## Escalation conditions

Stop condition above fires → land the census + comment fix only, route the fold. A live fence's findings
move → same. Anything requiring a platform-repo edit → out of scope, route.

## Acceptable close-no-lift outcomes

R2 falsified (the five do NOT agree today) would itself be the deliverable — it would mean the family has
been measuring different denominators since before iter-209, and the fold would become a behaviour change
needing its own measurement rather than a de-duplication.
