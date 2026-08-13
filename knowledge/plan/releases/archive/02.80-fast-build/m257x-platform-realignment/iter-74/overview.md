---
iteration_type: tik
status: closed-fixed
opened: 2026-08-04
closed: 2026-08-04
active_strategy: TOK-05
---

# iter-74 — the ambiguous class: is 12 → 39 reach, or rot?

**Type:** tik, under [`TOK-05: stop repairing claims; fence the predicates under them`](../decisions.md#tok-05-stop-repairing-claims-fence-the-predicates-under-them--2026-08-04).

## Step 0 — re-survey (mandatory, run BEFORE targeting)

Five corpus guards run at open, all **OK**, and every number below is re-derived here rather than
inherited (§5 rule 32 — *re-derive the hand-off's numbers, including the orchestrator's*):

| guard | invocation | verdict at open |
|---|---|---|
| `platform_alignment_guard` | `… corpus/architecture/platform-migration-status.md stack-demo/platform/repos.yml` | **OK** — F 74 citations (20 subject-checked · 53 range-only · 1 outside a service block · **0 unresolvable**) |
| `anchor_construct_guard` | `--repo-root $PWD` | **OK** — 177 resolved / 112 files / 239 unresolvable; ref chosen by `default x63, block-pinned x45, ambiguous x39, no-clone x30` |
| `platform_predicate_guard` | `--repo-root $PWD --platform stack-demo/platform --app stack-demo/app` | **OK** — G1 99 · G2 3 · G3 3 · G4 13 · G5 24 (1 enumerated + 21 free prose + 2 ref-pinned) · G5b 4 · G6 7 · G7 21/22 · G8 8/8 |
| `markdown_structure_guard` | `--repo-root $PWD` | **OK** — 112 published files |
| `corpus_index_guard` | `corpus` | **OK** — 84 docs across 6 index-bearing dirs |

The orchestrator's hand-off named two numbers. Both were re-derived, and **one of them does not
reproduce as described**:

- **39 ambiguous — REPRODUCES exactly.**
- **"92 bare citations still unresolvable" — the COUNT reproduces (91 distinct bare-code citations
  / 103 sites), the HEAD LIST does not.** iter-73 handed forward *"`gen.py` x10, `intelligence.go`
  x8, `main.go` x7"*; measured now the heads are **`up-injected.sh` x32, `intelligence.go` x5,
  `20260722104506.sql` x5, `main.go` x2 — and `gen.py` x0.** `gen.py` is cited **eleven times** in
  the corpus and **not once** in the `file:N` form: every one is a RANGE (`` `gen.py:484-492` ``),
  which the guard's regex cannot match because it requires a closing backtick right after the
  digits. So iter-73's head list came from a **different, broader instrument** than its count. The
  count is sound; the heads were not measured with the regex they were attributed to. **This class
  is iter-75's, not this iter's** — recorded here so the number that gets repaired is the one that
  was measured.

## Cluster / target identified

`CHECK-M257x-iter73-ambiguous-grew` — the `ambiguous` ref-source bucket went **12 → 39** across
iter-73's reach widening. The route says explicitly that whether this is *"a corpus-writing habit
worth changing or a fence limitation"* is **not settled**, and the orchestrator's framing is the
governing one: **growth here is a sign the fence sees more, not that the corpus got worse —
establish which before repairing.**

TOK-05's `Next-tik direction` names the citation classes as the ordering's second rung
(*fence → citations → map state → read*), and this is the citation rung.

## Hypothesis — two parts, both mechanically decidable

**H1 (attribution).** The 12 → 39 growth is **entirely** the newly-reachable `bare-code`
alternative. Split the class by which regex alternative matched: if the `path`-qualified ambiguous
count is still **12** and the `bare-code` ambiguous count is **27**, the pre-existing class did not
move by one and the growth is reach, not rot.

**H2 (a real fence defect inside the class).** `anchor_construct_guard._block_of` computes a
**blank-line-delimited** window. A markdown **table has no blank lines between its rows**, so for
any citation inside a table the "block" is the *entire table* — and every sha named in ANY row
pollutes every citation in EVERY row. That contradicts **§5 rule 33**, which this milestone already
derived and already implements in the sibling guard: *"a pin's scope is the claim's own block — a
markdown **CELL** in a table, a wrapped sentence in prose."* Two guards, two definitions of
*block*, and only one of them matches the rule the corpus records.

The prediction that makes H2 refutable: **21 of the 39 ambiguous sites are in
`platform-migration-status.md`**, whose citation region is a 10-row repo table. If H2 holds,
narrowing the window to the table ROW reclassifies most of those; if it does not, the count barely
moves and H2 is wrong.

## Expected lift

- `ambiguous` falls from 39 to well under 20 (H2), with the residue being genuine two-ref contrast
  blocks — the case `block_ref` was written to fall back on.
- Some citations move `default` → `block-pinned`, i.e. get read at a **different file** than they
  are being read at today. **That is the payoff, not a side effect**: any finding it surfaces is a
  citation that has been graded against the wrong copy of the code.
- Net corpus repairs: unknown at open, and deliberately not predicted.

## Phase plan

- **A** — derive the attribution split (H1) and the table-row structure (H2). *(done at open; see
  `progress.md`)*
- **B** — implement the row-scoped window, watched RED before trusting it (§8 rule 5: an inverted
  mutant AND a no-op positive control that must SURVIVE — iter-73's control did not, and *if the
  control fails the battery has not run*).
- **C** — re-measure; adjudicate every citation whose ref changed; repair what is genuinely wrong.
- **D** — gates: five corpus guards, the `CITE_REF=worktree` escape hatch still discriminating,
  `tests/test_iter45_mechanical_fences.py`, the mutation battery, the `stack-core` suite against its
  **stated** baseline (769 / 1F, the perishable iter-48 fixture, matched by IDENTITY).
- **E** — close.

## Escalation conditions

- If the row-scoped window turns the corpus RED with more than ~15 findings, **measure and route**
  rather than land-and-repair (iter-73's lesson 1 — dry-run the widening before landing it).
- If narrowing the window requires a rule that has to be *tuned* to make the numbers come out (§4
  Trap A), the fix is refused and the class is routed with the falsification.

## Acceptable close-no-lift outcomes

- **H2 refuted** — the table-row window does not reduce `ambiguous`, or reduces it only by
  mis-scoping. Recording that, with the measurement, is a complete iter: it settles the open
  question the route left explicitly unsettled.
- **H1 alone** — even if H2 is refused, establishing that the growth is 100% reach closes
  `CHECK-M257x-iter73-ambiguous-grew` on its own terms.
