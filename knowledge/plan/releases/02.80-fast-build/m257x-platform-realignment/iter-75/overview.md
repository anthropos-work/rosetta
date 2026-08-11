---
iteration_type: tik
status: closed-fixed
opened: 2026-08-04
closed: 2026-08-04
active_strategy: TOK-05
---

# iter-75 — the 92 "unrepaired" citations: adjudicate before repairing

**Type:** tik, under [`TOK-05`](../decisions.md#tok-05-stop-repairing-claims-fence-the-predicates-under-them--2026-08-04).

## Step 0 — re-survey

iter-74 closed with the class already re-derived and **one of its two hand-off numbers corrected**
(`D-M257x-74-5`): the routed *"92 bare citations still unresolvable"* reproduces as a **count**
(91 distinct citations / 103 sites) but its **head list came from a different instrument** — the
routed heads named `gen.py` x10, and `gen.py` is cited eleven times in the corpus and **never** in
`file:N` form. Also carried: **88 of the 239 sites the guard calls "unresolvable" are URLs**
(`http://backend:8083`), not citations at all.

Guards at open — all five **OK**: `anchor_construct_guard` 177 resolved / 239 unresolvable
(`default x82, block-pinned x45, no-clone x30, ambiguous x20`) · `platform_alignment_guard` F 74 /
0 unresolvable · `platform_predicate_guard` OK · `markdown_structure_guard` 112 files ·
`corpus_index_guard` 84 docs.

## Cluster / target identified

`FIX-M257x-iter73-unresolvable-92`. TOK-05's ordering puts the citation classes second
(*fence → citations → map state → read*), and this is the last open one.

**The governing instruction is the orchestrator's: every routed count is a hypothesis. Adjudicate
before repairing.** Three consecutive routed backlogs in this milestone collapsed on adjudication
— 64 → 5, 23 → 1, 21 → 0 — because each was routed by pattern-match on a count.

## Hypothesis

The class is **not a repair backlog at all**; it is the ninth reach limit. iter-73 taught a bare
`<name>.<ext>:N` to **reach** the resolver and did not teach the resolver to **find** it — every
route `resolve()` owns is positional (a `<repo>/…` prefix, a repo-relative path inside a service
doc, the platform repo, an rext-relative path), and an ops doc citing `` `up-injected.sh:1487` ``
supplies no position at all.

**Pre-registered three-way fate, adjudicated against `git ls-files` across every clone** — the
repository's own answer, not a directory walk:

| fate | meaning | response |
|---|---|---|
| **UNIQUE** | the basename is exactly one tracked file in the universe | a reach limit — resolvable by a unique-basename rule |
| **MULTI** | more than one | **must stay unresolved**; guessing a directory is the documented over-match |
| **ABSENT** | nowhere | a real corpus defect — a citation naming a file that does not exist |

## Expected lift

Reach grows by the UNIQUE count; MULTI stays counted-and-named; ABSENT is repaired. **Findings are
not predicted** — iter-73's comparable widening turned the corpus RED with 6, and the honest
position at open is that this one could do the same.

## Phase plan

- **A** — adjudicate the class three ways (done at open).
- **B** — dry-run the rule with the guard untouched (iter-73 lesson 1 / iter-74 lesson 6), with a
  **positive control on the dry run itself** so a 0-finding result is a measurement rather than an
  impression.
- **C** — land it; re-measure; repair whatever it turns RED.
- **D** — gates: five corpus guards · `CITE_REF=worktree` still discriminating ·
  `tests/test_iter45_mechanical_fences.py` · mutation battery (inverted mutant per clause + a no-op
  control that must SURVIVE) · `stack-core` against its stated baseline **775 / 1F** (the perishable
  iter-48 fixture, by IDENTITY).
- **E** — close.

## Escalation conditions

- More than ~15 findings → **measure and route**, do not land-and-repair.
- Any rule that has to be *tuned* to make a number come out (§4 Trap A) → refuse and route with the
  falsification.

## Acceptable close-no-lift outcomes

- **The class adjudicates to MULTI-dominant** — i.e. most of it is genuinely undecidable. Recording
  that with the derivation closes the route on its own terms: an unresolvable citation that *cannot*
  be resolved without guessing is coverage, not debt.
