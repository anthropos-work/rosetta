---
iter: 65
milestone: M257x
iteration_type: tik
status: closed-fixed
opened: 2026-08-04
---

# iter-65 — a citation must name its subject

**Active strategy reference:** `TOK-05`. Two routed `CHECK-*` items that are the same class:
**an anchor that resolves and still does not name the claim.**

## Cluster / target identified

- `CHECK-M257x-iter60-g6-citation-subject` — G6's two-sided record was closed by `if site in
  all_text`, a whole-corpus substring test. The one instance that closed it was hand-checked.
- `CHECK-M257x-iter64-pms-87-subject` (opened by iter-64) — `service_taxonomy.md`'s Directus
  retraction cites `platform-migration-status.md:87` as *"the corpus's own fenced source of truth"*.

`anchor_construct_guard`'s own docstring names this as the line the fence family does not cross:
deciding what a sentence *claims*. **For G6 it is decidable**, because the subject is a known token.

## Hypothesis

Requiring the site and the variable in the same **block** (the unit `_pin_window` established at
iter-63) converts a hand-check into an assertion, without needing to parse a claim.

## Expected lift

G6's two-sided test scoped to the block; the `pms:87` instance adjudicated against artifacts.

## Phase plan

A. Strengthen the G6 test; watch RED on fixtures and mutants.
B. Adjudicate `pms:87` against the platform, not against another doc.
C. Guards + suite + the §5 rule 34 re-point.

## Escalation conditions

- If the strengthening turns the live corpus RED at scale, route the residual WHOLE rather than
  weakening the rule.

## Acceptable close-no-lift outcomes

- The strengthened rule finds nothing live — recorded as such, with the fixtures as the evidence
  that it *can* fire.
