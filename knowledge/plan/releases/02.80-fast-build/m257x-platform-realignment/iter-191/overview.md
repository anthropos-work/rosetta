---
iter: 191
milestone: M257x
iteration_type: tik
status: in-progress
opened: 2026-08-09
controlling_strategy: TOK-08
---

# iter-191 — the guard says "all 119 scanned doc(s) agree" and scanned 164 files

**Active strategy reference:** [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07).

## Step 0 — re-survey before targeting

`SURVEY-M257x-iter188-the-other-walks-are-unmeasured` named **three** name-based prune rules in
`stack-core`. Two are handled (`_SKIP_DIRS` at iter-188, the `platform_predicate_guard` pair at 189/190).
The third — `story_org_count_guard._EXCLUDED_DIRS` — is re-surveyed here, and the interesting finding is
**not** the prune list: that one is component-matched, reasoned, and carries the story of the bug it
memorialises. The finding is one line further down.

## Cluster / target identified — two planned lines

**L1 — the printed denominator is a different unit from the population it describes.**
`_SCANNED_SUFFIXES = (".md", ".sh", ".yaml", ".yml")` is what `find_violations` walks. The number the
guard *prints* — and the number its **refusal gate** keys on — is computed from `root.rglob("*.md")`
alone:

| | files |
|---|---:|
| printed: *"all N scanned doc(s) agree"* | **119** |
| actually scanned (`.md` 119 · `.yaml` 33 · `.sh` 9 · `.yml` 3) | **164** |

**The printed number covers 72.6 % of what the guard checked**, and the error runs both ways: a scope
with no markdown but 45 shell/YAML sites refuses as *"0 markdown file(s) in scope. Nothing was checked"*
while there were 45 files to check. iter-172's unit class, and iter-186's *print the scope with the
number*, in the same expression.

**L2 — `_excluded` is applied to ABSOLUTE paths**, so the exclusion is scoped by the whole filesystem
path rather than by the tree being scanned. This module's own comment memorialises the bug it came from
— *"an exclusion that matches the absolute path is an exclusion that can swallow the whole repo. Match
components; never substrings"* — and it fixed the **matching mode** while leaving the **scope**: both
halves were in the same sentence of the original bug report. Sized before repair: **0 difference** on
this box (no scan root's absolute path carries an excluded component), and the module's refusal gate
fails closed, so the residual is latent rather than live.

## Hypothesis

Deriving the denominator from `_SCANNED_SUFFIXES` (one derivation, not a restatement), printing the
per-suffix scope, keying the refusal on the true total, and scoping the exclusion to the scan root turns
a 72.6 %-of-itself number into a stated one and closes the last member of iter-188's routed set.

## Expected lift

No `P`/`N` reading (`§9`). 1 live printed-denominator defect (119 → 164) with its refusal gate; 1 latent
absolute-scope residual closed; ≥5 arms mutation-proven; `SURVEY-M257x-iter188-…` closed.

## Phase plan

- **A** — measure both (done).
- **B** — derive the denominator; print the per-suffix scope + the exclusion's reach; scope `_excluded`.
- **C** — fence both directions + an instrument control.
- **D** — mutation-prove; both runners.
- **E** — publish; route residuals.

## Escalation conditions

- If the true denominator changes the guard's **verdict** (not just its number), this is a live
  correctness defect rather than a reporting one — grade it that way.
- If scoping `_excluded` to the root changes what is excluded anywhere, the absolute form was
  load-bearing; keep it and fence the difference.

## Acceptable close-no-lift outcomes

- The markdown-only denominator turns out to be deliberate (a comment or arm says the claim class only
  ever appears in markdown) → the finding shrinks to a labelling defect; record it and close.
