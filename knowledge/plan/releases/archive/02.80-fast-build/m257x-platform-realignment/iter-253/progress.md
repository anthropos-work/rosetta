**Type:** tik · **Active strategy:** `TOK-08` — census the mechanical classes; stop sampling them.
**Route:** `ROUTE-M257x-249-fresh-checkout-hostile-tests`, the largest thing iter-249 found.

## Open — 2026-08-10

Sealed PR-1…PR-5 before running anything (`1f1e0be`). Two of the five predict against numbers this
milestone has already published.

## What happened

### The instrument, and why it is not a text scan

The tempting fence was built and measured **first**, so it could be refuted before it was designed in:
*a test file that names operator-local state (`stack-dev` / `stack-demo` / `.agentspace`) and carries no
skip idiom.* It selects **8 files, 6 in `stack-core`**, against a dynamic class of **13**. Worse, and this
is the number that closes the question:

> **7 of the 13 failing files are not in the 50-file grep candidate set at all.**

They reach operator-local state without ever spelling it — they derive it. `test_toolchain_floor_guard`
had already shown the same thing from the other side at iter-249: it **carried** a `skipUnless` and failed
anyway, having declared **half** a precondition. The defect is at the **test-function** grain and is
conditional on state, so it is decidable only by **running the tests with that state absent**.

The demonstration is in this iter's own new test file, which the grep flags — on a comment that reads
*"Not asserting the literal `.agentspace/rosetta-extensions`"*. A text predicate cannot tell an assertion
from a sentence disclaiming one.

### The reading — frozen pair, then the live control

Frozen with iter-249's recipe (`git clone --local --shared`, both repos, **0.5 s / 54 MB**, 4,222 files /
2,879 markdown) at rosetta `1f1e0be` + rext `d739952`:

| tree | result |
|---|---|
| **frozen pair**, whole `stack-core` section | **27 failed · 2,060 passed · 27 skipped** (758.8 s) |
| **live control**, the same 27 node-ids | **22 passed · 5 failed** (838.8 s) |

*Both timings were taken with a second suite running concurrently on the same box and are **not**
comparable to iter-249's 726 s / 806 s. The verdicts are unaffected; the clocks are not.*

**The partition: 22 BOX · 5 REAL · 0 DECLARED**, and the 5 REAL are the finding.

### The 5 REAL — a ratchet that was breached before this iter touched anything

All five are in `test_frozen_expectation_census_m257x.py`, and they fail on **both** trees, so they are not
about the box:

```
TEST_MODULE_LITERAL_CEILING   live 624  ceiling 622  BREACHED +2   ← on the FROZEN tree, at HEAD
```

`TEST_MODULE_LITERAL_CEILING` was pinned at **622** by harden pass 61. Iters 249–252's own new test files
pushed the live population to **624** and **nothing noticed for four iters**, because the module that
asserts it is not run by the iter loop — the precise shape of
`ROUTE-M257x-h59-rext-edits-fire-no-fence-anywhere`, whose measured cost last session was 11 silent REDs.
Here are 5 more, and this time the instrument that found them was pointed somewhere else entirely.

**Without the control these five would have been published as fresh-checkout artifacts.** That is
`D-M257x-253-2`, and it is the reason the report prints REAL findings *first*.

### The 22 BOX — and the class is NOT growing

The arithmetic closes exactly, in both directions:

| | |
|---|---|
| iter-249's frozen class | **29** |
| repaired since (5 `LiveTree` + `toolchain_floor_guard` @ iter-249, `rext_path_guard` @ iter-250) | **− 7** |
| **newly manufactured BOX members** | **+ 0** |
| = this iter's BOX | **22**, across **12** files |

Every one of the 22 was already in iter-249's 29. So iter-249's *"the class is still being
manufactured"* — true of the files' authorship — is **not** true of the population since: three repairs
landed, nothing new arrived, and `PR-3` is refuted on the measurement rather than on the story.
`rext_path_guard` leaving the list is iter-250's repair confirmed from the outside.

### The names are now durable, which is what `ROUTE-M257x-249` actually needed

iter-249 **did** write its 29 node-ids down. It wrote them to
`.agentspace/scratch/work-m257x/iter249-failing-ids.txt`, and `git check-ignore` answers
`.gitignore:138:.agentspace/`. The names existed and could not survive the session — the same
git-ignored-substrate mechanism as `ROUTE-M257x-h59`. They are rescued into this iter's `evidence/`, which
is the only reason the diff above could be computed at all. Four files ship: iter-249's 29, this iter's
27, the 5 REAL and the 22 BOX.

### The deliverable

`suite_census.py --fresh-checkout` — freeze → census → **live control** → partition, extending the module
whose founding rule is already *"the third bucket is DECLARED, not sniffed"*. It prints its shape, its
denominator and its reach on every run; it refuses the shape it has not implemented (`rosetta-only`,
`ROUTE-M257x-251`); and its BOX message names the repair that is **not** allowed — an `ENV_GATED` entry,
which would green the instrument and leave every one of those tests failing for the reader.
**19 tests**, including a doctrine pin on that last sentence and a partition-exactness control.

Corpus side: **`§5` rule 78** in `corpus/ops/platform-alignment.md` carries the whole iters-249→253
thread, which had reached the protocol doc **nowhere** — clone-not-export, the two fresh checkouts, the
dynamic predicate, and the control.

## Pre-registration grading (sealed at `1f1e0be`)

| # | claim | prediction | outcome |
|---|---|---|---|
| **PR-1** | the static file-level predicate reproduces the dynamic class | false | **HELD** — 6 vs 13, and 7 of 13 files are outside the candidate set entirely |
| **PR-2** | the frozen run reproduces iter-249's **23 / 13** exactly | false | **HELD** — 27 node-ids; the *file* count is 13 by coincidence, the membership differs by 12 |
| **PR-3** | ≥ 1 hostile test manufactured by iters 250–252 | **true** | **REFUTED — 0.** The 5 new frozen REDs are older classes (iter-207, passes 52/55/61) that went RED for an unrelated reason, and no BOX member is new |
| **PR-4** | the grep candidate set is a strict superset of the failing files | **true** | **REFUTED** — 7 of 13 absent. Also: the set is **50** at HEAD, not the **51** sealed (51 counts this iter's own new file) |
| **PR-5** | every frozen failure passes live (0 real defects) | **true** | **REFUTED** — 5 fail on both trees |

**2 of 5.** Run trend: **1/5 → 3/5 → 5/5 → 4/4 → 2/5.** The three refuted are the three where I assumed
the class was what iter-249 described; the two that held are the two I had already measured a piece of.
`PR-4` also caught a figure of mine that was wrong in the seal itself — **50, not 51** — which is the
1-in-3 rate this milestone keeps re-learning about derived numbers.

## Close — 2026-08-10

**Outcome:** `ROUTE-M257x-249-fresh-checkout-hostile-tests` turned from **a count into a named, controlled
census**, and the control immediately earned itself. Frozen pair at rosetta `1f1e0be` / rext `d739952`:
**27 failed**, partitioned **22 BOX · 5 REAL · 0 DECLARED**. The 22 are exactly iter-249's residual after
three repairs, with **zero new members** — the class is not growing. The 5 are a **literal ratchet
breached at 624 against a pinned 622 on the frozen tree, i.e. before this iter changed anything**, unseen
for four iters because nothing in the iter loop runs the module that asserts it. Shipped
`suite_census --fresh-checkout` (19 tests) and `§5` rule 78, and made iter-249's names durable — they had
been written to a git-ignored path, which is what `ROUTE-M257x-249-a-reading-must-name-its-failures`
was really about.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: continue
**Decisions:** `D-M257x-253-1` (the class is decidable dynamically, never statically at the file grain) ·
`D-M257x-253-2` (the live control is the load-bearing half — it separated 5 real REDs from 22
environmental ones) · `D-M257x-253-3` (the repair is a `skipUnless` in the test, never an `ENV_GATED`
entry) · `D-M257x-253-4` (a reading's names must be durable; iter-249's were git-ignored) ·
`D-M257x-253-5` (a class-level `skipUnless` protects the class, not the file) · `D-M257x-253-6` (`--shape`
ships only the tree it implements).

**Suite state at close** — Python (`stack-core`, `/usr/bin/python3 -m pytest`, CPython 3.9.6):
`test_fresh_checkout_census_m257x.py` **19 passed**; `test_frozen_expectation_census_m257x.py` **RED→GREEN
after the ratchet re-pin** (see side-deliverable); population/collection meta-tests
(`test_suite_census_population.py` + `test_test_collection_fence.py`) **71 passed / 1 skipped**;
`test_guard_family.py` + `test_guard_family_verdict_line_m257x.py` + `test_fence_provenance.py` **green**.
Frozen-pair whole section: **27 failed / 2,060 passed / 27 skipped**.

**Side-deliverables:**
- **The two breached literal ratchets re-pinned with recorded reasons** (`derivation_registry.py`) —
  `TEST_MODULE_LITERAL_CEILING` and `COMMENT_LITERAL_CEILING`. The breach is split and attributed rather
  than absorbed: the TEST_MODULE excess was **+2 before this iter existed** (iters 249–252) and **+2**
  from this iter's own test module; the COMMENT excess is **entirely this iter's** (the frozen tree reads
  `exact +0`). Separate commit, separate decision; it does not upgrade the iter's close status.

**Routes carried forward:**
- `ROUTE-M257x-249-fresh-checkout-hostile-tests` → **still open, and now actionable**: 22 named node-ids
  across 12 files, each needing a `skipUnless` that states its precondition. Handler:
  `FIX-M257x-249-declare-the-clone-precondition`. The census is the acceptance test — it goes green when
  they skip.
- `ROUTE-M257x-253-the-iter-loop-runs-no-ratchet` → **new, and it is `ROUTE-M257x-h59` with a second
  witness.** Three literal ceilings are asserted only by a module nothing in the loop runs; two were
  breached at HEAD and stayed breached for four iters. The one-command form
  (`derivation_registry.py --ceilings`, ~1 s) exists and is not wired into any close step.
- `ROUTE-M257x-253-suite-census-is-undocumented-in-rext` → **new.** `suite_census.py` — the instrument
  behind this milestone's headline suite figures — appears in **no markdown in the rext repo**. It is
  documented only in the rosetta corpus. Sibling of `ROUTE-M257x-244-two-fences-entered-the-family-unindexed`.
- `ROUTE-M257x-249-a-reading-must-name-its-failures` → **re-aimed, not closed**: the discipline is
  *durability*, not *naming*. See `D-M257x-253-4`.
- `ROUTE-M257x-251-two-trees-both-called-a-fresh-checkout` → open; `--shape rosetta-only` is declared and
  refused rather than silently wrong.
- Still open, untouched: `ROUTE-M257x-250-workspace-tier-is-invisible-without-a-workspace` ·
  `ROUTE-M257x-249-anchor-offset-has-three-populations` ·
  `ROUTE-M257x-246-two-of-four-censused-surfaces-still-have-no-fence` ·
  `ROUTE-M257x-244-two-fences-entered-the-family-unindexed` ·
  `ROUTE-M257x-244-unresolvable-and-wrong-share-one-bucket` · `ROUTE-M257x-h59-range-anchors-are-ungraded` ·
  `ROUTE-M257x-241-wider-citation-surface-is-ungraded` ·
  `ROUTE-M257x-240-prereq-floors-live-in-three-parallel-blocks` ·
  `ROUTE-M257x-238-container-vs-native-is-undrawn` · `ROUTE-M257x-237-hardcoded-vs-settable` ·
  `ROUTE-M257x-236-disclosure-scope-is-document-level` · `ROUTE-M257x-235-fence-scope-is-unread` ·
  `ROUTE-M257x-235-runnable-block-has-two-halves`.

**Lessons:**
1. **A frozen list without a live control is not a classification.** 5 of 27 were real REDs at HEAD; the
   bucket built to excuse the environment would have swallowed every one of them. The control cost one
   scoped re-run and changed the meaning of the whole reading.
2. **Build the tempting instrument first, so it can be refuted before it is designed in.** The static
   predicate took four minutes and its two numbers — 6 of 13, and 7 of 13 outside the candidate set —
   are what justify the expensive dynamic one. A rejected instrument that was measured is evidence; a
   rejected instrument that was reasoned about is an opinion.
3. **"The names were never written" and "the names could not survive" want different fixes.** iter-249
   wrote all 29 correctly, into a git-ignored directory. The route said *name your failures*; what it
   needed was *put them where the next reading can diff against them*.
4. **A class-level `skipUnless` is not a file-level guarantee.** The repaired file has the same hostile
   count it started with, from four sibling classes that do not inherit the declaration.
5. **A sealed number can be wrong in the seal.** `PR-4` said 51; the candidate set is 50 at HEAD, and 51
   only counts this iter's own file. The seal makes the error legible instead of load-bearing.
