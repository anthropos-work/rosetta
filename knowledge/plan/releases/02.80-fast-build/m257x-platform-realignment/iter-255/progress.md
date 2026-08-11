**Type:** tik · **Active strategy:** `TOK-08` — census the mechanical classes; stop sampling them.
**Route:** `ROUTE-M257x-249-fresh-checkout-hostile-tests`. iter-253 named the class, iter-254 took it
**22 → 12**; this is the rest of `TOK-08`'s *"run it to zero."*

## Open — 2026-08-10

Sealed PR-1…PR-5 (`666c3fb`). Three of five predict against a clean sweep, including `PR-3`, which says
this iter does **not** finish the class.

## What happened

### The class is at ZERO on a fresh checkout, and it was verified in both directions

| | frozen pair (rosetta `1f1e0be` / rext `d739952`) | live tree |
|---|---|---|
| the 22-member target | **0 failed · 20 skipped · 2 passed** | **22 passed · 0 skipped** (1,070 s) |

Every one of the 12 residual members got its **own** declaration — none fell out as a cascade — and every
one skips on a fresh clone naming its precondition while still running and passing here.

### Only two preconditions exist, and that is now a bounded negative

Read one failure at a time across the full population: **clone set 19, `node_modules` 3.** No third cause
and no non-environmental member surfaced, which refutes `PR-1` — the class is *closed*, not merely
unexhausted, because every member was examined rather than sampled.

The residual's symptom was a **collapsed denominator** rather than a named missing target, which is why
it looked like a different class and was not: *"resolution collapsed — the zero is vacuous"* (136 against
a floor of 200), *"the widened denominator collapsed"* (30 against 150), *"resolves in neither pool and is
undeclared"* (15 citations, every one a real file inside a clone —
`UPGRADE-IMPACT-next16.md` at `stack-demo/next-web-app/`, `internal/askengine/rules.md` in `app`).

And the **half-precondition** shape held to the end: every one of these arms already skipped on
`if not (root / "corpus").is_dir()` — a check a fresh clone **satisfies**, because the corpus is exactly
what a fresh clone has.

### The finding: a declaration weakens every mutation proof that depends on that test firing

`test_m257x_mechanical_fences_mutation_battery::test_01` was the last member, and its failure text
**changed mid-repair**. Before iters 254–255 it read *"the declared-GREEN control went RED"*. After the
detecting arms gained their preconditions it read:

> `THEATRE: mutant 'anchor-header-lookahead-dropped' left the suite GREEN.`

**That verdict is correct**, and the battery made it about my own repair: on a fresh checkout the mutant
really is undetected, because its detector now skips. The battery was not broken by the declarations — it
*measured* them.

So the rule (`D-M257x-255-1`) is not *"the battery needs clones too"*. It is: **a mutation battery is
evidence only where the arms that would detect the mutation actually run**, and every precondition added
anywhere in a suite narrows the trees on which that battery's GREEN means anything. `§5` rule 77's lesson
— *a battery's GREEN is only evidence if the run was real* — arriving from a completely different
direction.

### Some of these failures come with a work order attached

Worth separating from the class's general shape (`D-M257x-255-3`). Most members merely mislead;
two **instruct**:

- `test_no_exemption_outlives_its_site` names `CLAUDE.md:226:internal/jobsimulation/runner` and
  `frontend_architecture.md:39:NEXT_PUBLIC_BACKEND_API_URL` as exemptions *"matching nothing"* — fiction
  to be retired. **Both are real.** Acting on the report deletes two correct exemptions.
- `test_01_no_undeclared_markdown_citation_resolves_nowhere` ends its message *"**Report it as a corpus
  defect — do not widen a class to absorb it**"* — on an unprovisioned box, 15 defect reports against a
  corpus that is right.

## Pre-registration grading (sealed at `666c3fb`)

| # | claim | prediction | outcome |
|---|---|---|---|
| **PR-1** | the 12 need only the two established preconditions (no third appears) | **false** | **REFUTED** — no third appeared; the class closes at two |
| **PR-2** | ≥ 1 of the 12 resolves as a cascade, no edit of its own | **true** | **REFUTED** — all 12 needed their own declaration; the only cascade in this route was iter-254's |
| **PR-3** | the class reaches **0** this iter | **false** | **REFUTED** — it reached 0 |
| **PR-4** | no arm repaired this iter becomes a live SKIP | **true** | **HELD** — live 22 passed / 0 skipped |
| **PR-5** | ≥ 1 of the 12 is a REAL defect iter-253's control mis-classified | **false** | **HELD** — none; the control stands over the full population |

**2 of 5.** Run trend: **1/5 → 3/5 → 5/5 → 4/4 → 2/5 → 3/5 → 2/5.** All three refutations are
**pessimistic** misses — I bet against a clean sweep three ways and got one. Booked on the same terms the
milestone books optimistic ones (`D-M257x-255-2`): *an estimate wrong in the comfortable direction is
still wrong*, and the cause is legible — iter-254 paid for the predicates and the reading discipline, so
the marginal cost per member collapsed to a `--tb=short` read and three lines.

## Close — 2026-08-10

**Outcome:** `ROUTE-M257x-249-fresh-checkout-hostile-tests` **CLOSED at zero.** The 12 residual members
each declared their precondition, and the class is now **0 failed / 20 skipped** on a clean clone of both
repos and **22 passed / 0 skipped** live. Two preconditions cover the whole 22 — clone set 19,
`node_modules` 3 — with no third cause and no non-environmental member across a population read one
failure at a time, so it is a bounded negative rather than an exhausted search. The last member paid for
the iter on its own: the mutation battery reported **`THEATRE`** at my own repair, because declaring a
precondition on a test silently narrows the trees on which any mutation proof that depends on it means
anything.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — (7) budget-exhausted: y — Outcome: exit-7

**Suite state at close** — Python (`stack-core`, `/usr/bin/python3 -m pytest`, CPython 3.9.6):
the 22-member target **frozen 0 failed / 20 skipped / 2 passed**, **live 22 passed / 0 skipped (1,070 s)**;
`test_anchor_subject_census_m257x.py` + `test_m257x_corpus_file_citations.py` **live 31 passed**, **frozen
25 passed / 6 skipped**; `test_frozen_expectation_census_m257x.py` **99 passed** after the ceiling re-pin.
Three literal ceilings **exact +0** (234 / 220 / 634). *No whole-section re-run this iter — the scope is
7 files and both trees were measured on that scope.*

**Side-deliverables:**
- `TEST_MODULE_LITERAL_CEILING` 628 → 634 and `COMMENT_LITERAL_CEILING` 219 → 220, re-pinned with
  recorded reasons. **The comment arrow breached the ceiling it was written to raise for the second
  consecutive iter**, so it is folded into the arrow text as a property rather than an anecdote
  (`D-M257x-255-5`).

**Routes carried forward:**
- `ROUTE-M257x-249-fresh-checkout-hostile-tests` → **CLOSED.** 22 of 22, verified both directions. The
  standing instrument is `suite_census.py --fresh-checkout`; the acceptance test is that it stays green.
- `ROUTE-M257x-255-a-declaration-narrows-every-mutation-proof` → **new, and it generalises past this
  route.** `D-M257x-255-1`: the milestone has ~30 mutation proofs in `hardening-ledger.md`, and each is
  evidence only on trees where its detecting arms run. Nothing currently states each battery's reach.
  Handler: `FIX-M257x-255-battery-reach-is-unstated`.
- `ROUTE-M257x-255-the-class-can-regrow` → **new.** The class is at zero and **nothing watches it**: the
  census must be RUN to notice, and no close step runs it. Same shape as
  `ROUTE-M257x-253-the-iter-loop-runs-no-ratchet`, and the two should be answered together — a single
  cheap close-step reading both.
- `ROUTE-M257x-253-the-iter-loop-runs-no-ratchet` · `ROUTE-M257x-254-six-spellings-of-one-root` ·
  `ROUTE-M257x-253-suite-census-is-undocumented-in-rext` ·
  `ROUTE-M257x-251-two-trees-both-called-a-fresh-checkout` ·
  `ROUTE-M257x-250-workspace-tier-is-invisible-without-a-workspace` ·
  `ROUTE-M257x-249-anchor-offset-has-three-populations` → open.
- Still open, untouched: `ROUTE-M257x-246-two-of-four-censused-surfaces-still-have-no-fence` ·
  `ROUTE-M257x-244-two-fences-entered-the-family-unindexed` ·
  `ROUTE-M257x-244-unresolvable-and-wrong-share-one-bucket` · `ROUTE-M257x-h59-range-anchors-are-ungraded` ·
  `ROUTE-M257x-241-wider-citation-surface-is-ungraded` ·
  `ROUTE-M257x-240-prereq-floors-live-in-three-parallel-blocks` ·
  `ROUTE-M257x-238-container-vs-native-is-undrawn` · `ROUTE-M257x-237-hardcoded-vs-settable` ·
  `ROUTE-M257x-236-disclosure-scope-is-document-level` · `ROUTE-M257x-235-fence-scope-is-unread` ·
  `ROUTE-M257x-235-runnable-block-has-two-halves`.

**Lessons:**
1. **A repair can invalidate a proof that was not part of it.** The battery's `THEATRE` verdict was about
   the declarations, not about the fence. Before adding a precondition, ask which mutation proofs depend
   on that arm firing — nothing in this repo currently answers that.
2. **Grading a pre-registration is itself a derived figure.** `PR-1` was written up as HELD and is
   REFUTED; the seal states a CLAIM and separately my PREDICTION, and inverting them is a one-word error
   that flatters the author. Caught and corrected in place before the close.
3. **Three pessimistic misses in one iter is a signal about estimating, not about the work.** The
   expensive part of a mechanical class is the first two members; after the predicates and the reading
   discipline exist, the marginal member is three lines. Estimate the *tail* separately from the head.
4. **A false alarm with a work order is a distinct severity.** Two of these arms did not merely mislead —
   they instructed a reader to delete correct exemptions and to file 15 defect reports against a correct
   corpus.
