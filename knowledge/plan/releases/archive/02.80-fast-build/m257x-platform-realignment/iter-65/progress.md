**Type:** tik — under `TOK-05`. Two routed `CHECK-*` items, one class: **an anchor that resolves and
still does not name the claim.**

# iter-65 — a citation must name its subject

## Phase A — the G6 subject rule

`CHECK-M257x-iter60-g6-citation-subject`. G6 requires a mid-fold variable to be recorded on **both**
sides — the config side it derives, the consumer side the corpus must cite — and it tested the second
with `if site in all_text`. A whole-corpus substring match: any document mentioning `main.go:446` for
any unrelated reason closed the finding.

`anchor_construct_guard`'s docstring calls this class the line the fence family does not cross,
because it means deciding what a sentence claims. **For a known token it is decidable**, and the rule
needs no claim-parsing: the site and the variable must appear in the same **block** — the unit
`_pin_window` established at iter-63.

**The live corpus is GREEN under the strengthening**, which is the right outcome and is *not* evidence
the rule works. The evidence is the fixtures:

| fixture | expected |
|---|---|
| site + variable in one block | closes the finding |
| site cited in ANOTHER block | does **not** close it |
| site named, variable nowhere | does **not** close it |
| variable alone, no site | does **not** close it (one side is not a claim) |
| wrapped prose across two lines | one block — closes it |
| a table row carrying both | one block — closes it |

## Phase B — the reach hole a fixture found

`test_a_site_named_with_no_variable_anywhere_does_NOT_close_it` came back **green when RED was
expected**, and the cause was not the citation rule (`D-M257x-65-2`):

```
universe = set(compose.rpc_addrs) | named_anywhere | set(env_example_names)
```

A variable the platform configures **nowhere** and the corpus names **nowhere**, yet `app` **reads**,
had **no row** — so G6 could not see it. That is the most-undocumented case there is: silent at boot,
absent from every document, invisible to the assertion whose job is to catch it. Universe widened to
include `set(app_reads)`.

**Second time this milestone a fixture has surfaced a reach hole the live corpus happened not to
exhibit** (iter-61 was the first). A fence's blind spots are not discoverable by running it on a tree
that does not contain them.

## Phase C — `pms:87`, adjudicated against the platform

`CHECK-M257x-iter64-pms-87-subject`. `service_taxonomy.md`'s Directus retraction appealed to
`platform-migration-status.md:87` as *"the corpus's own fenced source of truth"*. **That map has no
Directus row at all** — it maps *repos*; Directus is an external service. The anchor resolved, carried
content, passed `anchor_construct_guard`, and named `anthropos-studio-room`.

Repaired by naming what actually settles it — `git show a2a3ee6^:docker-compose.yml`, already cited in
the same paragraph — and recording why the old appeal was wrong, so the next reader does not re-derive
it. **Every mechanical check passed on this citation, twice**: iters 63 and 64 both re-pointed it
faithfully as the target row moved. **A re-point preserves intent; it cannot audit it.**

## Phase D — gates

| gate | result |
|---|---|
| `platform_predicate_guard` | **OK** |
| `platform_alignment_guard` · `anchor_construct_guard` · `markdown_structure_guard` · `corpus_index_guard` | OK |
| `tests/test_platform_predicate_guard.py` | **60 tests** (was 54), all pass |
| mutation battery | revert-to-substring-test → **2 RED**; drop-consumer-side-from-universe → **1 RED** |
| `stack-core` suite | **675 tests, 1F** — the perishable iter-48 fixture, the single expected failure |
| §5 rule 34 re-point | 1 citation into the edited passage, re-pointed by hand after checking the target |

## Close — 2026-08-04

**Outcome:** two routed `CHECK-*` items closed, both instances of *an anchor that resolves and does
not name the claim*. G6's two-sided test now requires the site and its variable in the same block —
decidable because the subject is a known token, which is exactly the carve-out
`anchor_construct_guard` could not take. Writing the fixture for it found a **second** defect: G6's
universe excluded the consumer side, so a variable configured nowhere, documented nowhere and read by
`app` had no row at all. And `pms:87` turned out to appeal to a map that has no Directus row —
resolved, content-bearing, faithfully re-pointed twice, and wrong the whole time.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — 4 of 5, unchanged; clause 5 is still graded only by a reading that returns zero.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** `D-M257x-65-1` (a citation must name its subject; decidable for a known token),
`D-M257x-65-2` (G6's universe excluded the consumer side — found by a fixture, not by the tree),
`D-M257x-65-3` (the `pms:87` anchor resolved to a row about something else; a re-point preserves
intent but cannot audit it).
**Side-deliverables:** none.
**Routes carried forward:**
- `FIX-M257x-iter63-app-citation-residual` (the 68 non-mainline `app` citations, routed WHOLE) ·
  `CHECK-M257x-iter63-quoting-a-retired-token` · `FIX-M257x-iter53-union-set`
  (**PENDING USER DECISION**) · `FIX-M257x-iter56-assignment-flake` (**NOT DECIDED** — needs a
  failure *rate*) · `CHECK-M257x-iter38-ai-act-classification` (needs an owner outside this
  milestone) · `CHECK-M257x-iter57-anchor-guard-bare-class` · `FENCE-M257x-iter54-refs-block` ·
  `FIX-M257x-iter57-within-block-drift` · `CHECK-M257x-iter58-derive-preregistrations` ·
  `CHECK-M257x-iter52-second-ai-manager` · `-cold-daemon-registry` · `-grep-vs-failclosed` ·
  `-empty-stdout-class` · `-baseline-refs` · RF-2/3/7–13 · root `CLAUDE.md`.

**Lessons:**

1. **"Undecidable in general" is not "undecidable here."** The fence family declined this class
   because judging a claim is hard. For a citation whose subject is a *named token*, it collapses to
   two tokens co-occurring in one paragraph.
2. **A GREEN live corpus is not evidence a new rule works.** The fixtures are. Say which is which in
   the same breath, or the strengthening reads as verified when it is merely quiet.
3. **A fixture can find what the tree cannot show you.** Twice this milestone now: the fence's blind
   spot is not in the tree, so only a case you construct will reveal it.
4. **A re-point preserves intent; it cannot audit it.** `pms:87` was faithfully carried through two
   line-map re-points, each preserving a citation that had never named its subject. Mechanical
   correctness on an anchor says nothing about the anchor being right.
