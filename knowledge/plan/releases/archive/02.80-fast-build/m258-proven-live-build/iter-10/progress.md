# M258 iter-10 — progress

**Type:** tik · **Active strategy:** `TOK-01`. Instrument integrity (M257x clusters 2 + 3, routed here).

## Phase A — the prune, and a falsifiable prediction that held

One name added to the existing checked-in registry (`_CENSUS_SKIP`), not new machinery — the repo had
already solved this class once (`census_pruned`, M257x harden pass 69) and enforced it in one walker.

**Prediction, stated before running:** the working-tree ratchets should fall to the pristine-HEAD
values **248 / 236 / 657**.

| ratchet | before | after | pristine HEAD | ceiling |
|---|---|---|---|---|
| DOCSTRING | 258 | **248** | 248 | 240 (still breached, **+8** not +18) |
| COMMENT | 237 | **236** | 236 | 236 (**exact**, no longer breached) |
| TEST_MODULE | 672 | **657** | 657 | 653 (still breached, **+4** not +19) |

All three landed exactly on prediction. **The debt is real but roughly half what the polluted reading
said**, and `RATCHET-M257-literal-ceilings-breached` can now be graded against this repo instead of
against whatever a demo last cloned.

⚠️ My own new comment added 1 COMMENT literal (`4,560-file`) and was **paid down, not waived** — the
figure was reworded, not the ceiling raised.

## Phase B — fenced in both directions, and mutation-checked

Three tests added beside the virtualenv precedent they mirror:
`test_a_per_stack_clone_is_pruned` · `test_pruning_stacks_does_not_prune_the_repos_own_sections`
(the negative control — a prune that also silenced `dev-stack/*.py` would fail **GREEN**, and a
substring rather than component-exact test would drop a file literally named `stacks.py`) ·
`test_a_planted_stack_clone_does_not_move_a_single_census_count` (the end-to-end form: identical
output across the whole root-taking census family, before and after a clone tree appears).

**Mutation:** removing `"stacks"` from the set fails 2 of the 3. The control holds.

## Phase C — the fourth consumer, found by doing what Phase C said

Phase C committed to *checking* the other consumers rather than claiming reach. Doing so found a
**fourth**: `population()` carried its own hand-rolled filter (`"/tests/" in str(path) or
"__pycache__"`), bypassing the shared helper entirely — and it was the only consumer **reporting the
problem out loud**, demanding a `DECISIONS` entry for
`demo-stack/stacks/demo-1/clones/app/.claude/skills/…` and `…/app/studio/benchmark/…`.

Its own contract is *"a non-test **rext** module"*, so routing it through `census_pruned` enforces the
stated contract rather than adding an exclusion. Result: `unclassified()` and `stale_decisions()` both
clean, `reach()` = (20, 109, 192), and **a pre-existing RED
(`test_every_executable_derivation_is_classified`) is now GREEN** with no new failures anywhere in the
four modules that import the registry.

## Phase C′ — a methodology error, caught and recorded

The first before/after used a pristine `git archive HEAD` extract as the baseline, which showed
`test_every_executable_derivation_is_classified` as a **new** failure of mine. It was not: `git
archive` omits gitignored paths, so the extract **has no `stacks/` tree at all** and cannot express
the defect. Re-tested by A/B-ing *in the working tree* — it fails there without the prune too.

**A pristine-checkout baseline cannot detect a defect whose cause is an untracked directory.** For
this class the control must be the working tree with the change reverted. Recorded because the
apples-to-oranges comparison very nearly booked a fix as a regression.

## Close — 2026-08-12

**Outcome:** The three literal ratchets now measure **this repo**: 258/237/672 → **248/236/657**,
exactly the pristine-HEAD prediction, so the recorded debt is roughly half what the polluted reading
claimed and COMMENT is not breached at all. A **fourth** consumer of the same root cause was found by
actually checking rather than asserting reach, and fixing it turned a **pre-existing RED green**.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n *(2 tiks)* — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**
*(superseded in-flight: the user ruled on the milestone during this iter — see `decisions.md` `D52`.)*

**Decisions:** D48–D52

**Routes carried forward:**

- **`ROUTE-M258-iter10-hand-rolled-path-filters`** — `census_pruned` now serves the literal censuses
  and `population()`, but other guards (`decommissioned_instruction_guard`'s `rglob("*")`,
  `claim_census_guard`'s own `SKIP_DIRS` walk, `corpus_index_guard`) still carry independent filters.
  `FIX-M258-iter03-guard-scans-its-own-scratch` and
  `test_fence_provenance::test_the_escape_accepts_and_records` are **not** claimed fixed — not run.
- **`RATCHET-M257-literal-ceilings-breached`** — still breached on honest numbers (**+8** DOCSTRING,
  **+4** TEST_MODULE). Pay down or attribute-and-raise; this iter fixed the *measurement*, not the debt.
- ⚠️ `demo-2` (11) / dev (5) / `demo-1` (11) all resident throughout; nothing was torn down.

**Lessons:**

- **When a fence reports something absurd, read it as a symptom, not as noise.** `population()` was
  demanding a classification decision for the platform's `.claude/skills/` — the most useful signal
  in the whole family, and the easiest to dismiss as junk.
- **A pristine-checkout baseline is blind to untracked causes.** A/B in the working tree instead.
- **Check the other consumers instead of asserting the fix reaches them.** Phase C was written that
  way deliberately and it is the only reason the fourth consumer was found.
