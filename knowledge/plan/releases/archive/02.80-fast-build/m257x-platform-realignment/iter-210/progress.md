**Type:** tik — under [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07).

# iter-210 — iter-209 widened one of two identical source-set derivations, and the family had no way to notice

## The defect this iter's own predecessor created

`grep -l "def collect_sources"` returns two modules — `corpus_citation_guard.py` and
`retracted_pin_guard.py` — and they held the **same four lines**. The second carried this comment
above its copy:

> *Same source set as `corpus_citation_guard` — the published corpus plus the two root docs every agent
> opens. **Kept identical on purpose**: two fences over the corpus that disagree about what the corpus
> IS would each be measuring a different denominator.*

**The intent was right and the mechanism was a duplicate.** iter-209 widened one copy, and four minutes
later the family answered *"which documents are the corpus?"* with **114** from one fence and **94**
from the other — symmetric difference exactly the 20 skill documents.

`§5` iter-190 named this and named the repair: **two readers of one construct must SHARE the
derivation — agreement today is not the property.** A comment asserting the copies are kept identical
is the strongest possible evidence that nothing was checking.

## What shipped

- **`fence_provenance.corpus_sources(repo_root)`** — one derivation, in the module every fence in the
  family already imports, with `CORPUS_DIR` / `EXTRA_SOURCES` / `SKILL_SOURCE_GLOB` as its single
  definitions. Both guards now call it; both re-export `EXTRA_SOURCES` so existing callers and tests
  keep working.
- Measured when `retracted_pin_guard`'s copy was replaced: sources **94 → 114**, pins enumerated
  **2,193 → 2,201**, findings **3 → 3, zero new.** The escalation condition in `overview.md` was that a
  *findings* change would stop the iter; it did not fire.
- **Deliberately NOT folded in:** the five fences declaring `SCAN_FILES = ("CLAUDE.md", "README.md")`.
  Those answer a different question — they fence those two documents' own prose — and merging them
  would be the same conflation one grain up.
- Three arms in `tests/test_corpus_citation_guard.py` (`TheFamilyHasONESourceSetDerivation`): the two
  fences return the same set; **only one module may spell the construct**; and a staged re-fork must be
  detected, so the live arm is provably able to fire.

## The third copy, found by the arm within a minute of it existing

The sharing arm immediately reported a module the two-fence comparison could never have seen:
**`clone_drift_guard.py` walks `corpus/**` twice with its own private glob** (`:158`, `:217`), for
backticked sha tokens.

It is **not** folded in, and the reason is measured rather than asserted: the shared set adds **22
documents carrying 73 sha-shaped backticked tokens** — occurrences **1,454 → 1,527, +5.0 %** — and
`CLAUDE.md` alone is dense with shas. That is a **behaviour change, not a de-duplication**, so it is
declared, sized and routed.

The exception lives in a `DECLARED_PRIVATE` map carrying the reason and the measured cost, and it is
**reconciled in both directions**: an undeclared speller fails, and a declared entry for a module that
no longer keeps a private walk fails just as loudly — *a waiver outliving its subject reads as
coverage.*

## An instrument note

The sharing arm's first draft matched the lowercase word `corpus` in a line, and `fence_provenance`
spells it **`CORPUS_DIR`**. The arm's own anti-vacuity half — *"the OWNER no longer spells the construct
at all, so this arm is asserting an empty condition"* — fired on its author within a minute. That is
iter-205's case-sensitivity class, one detector over, and it is the reason the anti-vacuity half exists.

## Close — 2026-08-09

**Outcome:** the family's two corpus-source derivations were a copy-paste pair under a comment claiming
they were kept identical; iter-209 falsified that by moving one. They now share **one** derivation in
`fence_provenance`, the second fence's population rose **94 → 114** with **zero** new findings, and the
sharing itself is fenced — which immediately surfaced a **third** private copy in `clone_drift_guard`,
declared with its measured folding cost (+73 sha tokens, +5.0 %) rather than silently skipped.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (forty-second consecutive `closed-fixed`;
**no `P`/`N` reading taken — UNMEASURED, not unmoved**, `§9`) — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n — **counted: iters 207, 208, 209, 210 = four tiks this run against a cap of five** —
(6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-210-1` … `D-M257x-210-3` (see [`decisions.md`](decisions.md))

**Audit:** `/usr/bin/python3 -m pytest` (**8.4.2** / CPython **3.9.6**), **Python**, `stack-core` only —
**154 passed** across `test_corpus_citation_guard` + `test_retracted_pin_guard` + `test_guard_family` +
`test_clone_drift_guard`. Both live guards on the real rosetta tree: citation **0 findings over 114
sources**; retracted-pin **3 findings over 114 sources, 2,201 pins** (the same 3 as at 94 sources).
**RED-proof battery, mtime-mitigated (`§5` r77):** two mutations. (a) the skills line deleted from the
shared derivation → **iter-209's two contract arms RED**, iter-210's sharing arms green (correct — both
fences still agree); (b) a private copy re-forked into `retracted_pin_guard` → **iter-210's two arms
RED**. Restores sha-verified against `c30ccd58…` and the retracted-pin baseline.
*Scope, stated rather than implied (`§5` r60): `stack-core` only, Python only, changed-code reach. No
whole-section run; no Go, no TypeScript; the four non-`stack-core` Python sections were read at iter-208
and not re-read here.*

**Side-deliverables:** none this iter.

**Routes carried forward:**
- `SURVEY-M257x-iter210-clone-drift-reads-a-third-corpus` — **NEW.** `clone_drift_guard` walks
  `corpus/**` privately at two sites for sha tokens. Folding it into the shared set is a behaviour
  change worth **+22 documents / +73 sha tokens / +5.0 % occurrences**, `CLAUDE.md` being the dense one.
  Declared in `DECLARED_PRIVATE` with that measurement; needs its own iter to decide, not a drive-by.
- `SURVEY-M257x-iter210-five-fences-scan-only-the-two-root-docs` — **NEW.**
  `markdown_structure_guard`, `anchor_construct_guard`, `claim_twin_guard`, `repair_leak_guard` and
  `platform_predicate_guard` all declare `SCAN_FILES = ("CLAUDE.md", "README.md")`. Correct for their
  subject and **not** compared against anything; whether each still wants exactly those two is a claim
  nobody grades.
- All routes from iters 207–209, unchanged, plus the standing queue.

**Lessons:**
- **A comment claiming two copies are kept identical is evidence that nothing checks.** This one was
  right for as long as nobody moved either copy, and wrong four minutes after somebody did.
- **Fence the SHARING, not the agreement.** Agreement was true on the day the copies diverged; only a
  test that the construct has one home would have caught it — and that test found a third copy at once.
- **A repair that widens one member of a family creates a family defect.** iter-209 was correct and
  incomplete in the same commit.
