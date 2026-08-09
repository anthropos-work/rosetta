# iter-212 — the arm that enumerates the family enumerates ONE SPELLING of it

**Type:** tik — under [`TOK-08`](../decisions.md) (census the mechanical classes; stop sampling them).

See [`overview.md`](overview.md) for the census sealed before any repair (R1–R6 + the stop condition,
commit `1f8c7ff`).

## What was measured

**R1 — CONFIRMED, and larger than stated.** `test_only_ONE_module_spells_the_corpus_source_construct`
enumerates by one literal string, `CONSTRUCT = 'rglob("*.md")'`, and its name promises a family census.
It sees **2 modules** (`fence_provenance`, `clone_drift_guard`). A discovery over collector *function
names* sees **9**. The five fences iter-211 routed forward derive the corpus through
`glob("corpus/**/*.md")` and `rglob("*")`+suffix and carry the literal string nowhere — **the fork
iter-210 wrote that arm to prevent was already present, five times over, in another spelling.**

**R2 — CONFIRMED.** Re-expanded from each module's OWN constants (not by calling its collector, so the
comparison was about the declaration): **six spellings, 114 documents each, all 15 pairwise symmetric
differences 0.**

**R3 — CONFIRMED.** `platform_predicate_guard`'s `SCAN_ROOTS = ("corpus", ".claude")` is a **strict
superset by rule** of the other four's `.claude/skills/**/*.md`, and equal to them **by absence**: this
repo's `.claude/` holds `settings.json`, `settings.local.json` and `skills/` — **0 markdown outside
`skills/`**. `.claude/agents/*.md` and `.claude/commands/*.md` are standard locations in this harness.

**R4 — CONFIRMED, and it is the finding worth the most.** `fence_provenance.py:267-270` still shipped
iter-210's rationale for keeping the five separate: *"Those answer a DIFFERENT question — they fence
those two documents' own prose."* **iter-211 retracted exactly that proposition** and the retraction
reached the milestone ledger, the journal and `progress.md` — **not the comment, which was the artifact
ACTING on the claim.** For one whole iter the false claim stayed load-bearing as a design justification.

**R5 — HELD; the stop condition did NOT fire.** Every live fence's verdict is byte-identical to its
pre-fold baseline: `markdown_structure` OK · `anchor_construct` OK · `claim_twin` OK (264 adjudicated
claims) · `corpus_citation` OK · `retracted_pin` OK · `repair_leak` `CANNOT RUN` (its declared
no-candidate-shingles refusal, unchanged) · `platform_predicate` OK against `stack-demo/platform`.

**R6 — CONFIRMED by construction.** The live tree cannot separate `.claude/**` from
`.claude/skills/**` (R3: the difference is 0 documents), so the separating control had to be **staged**.

## What was shipped

- **The retracted comment corrected in place**, quoting the retracted text rather than deleting it.
- **`fence_provenance.claude_docs_outside_skills()`** — the wider member's extra as a **named, callable,
  sized** derivation instead of a private constant. Size **printed** every run (0); **shape asserted**
  (`§5` — *print the SIZE, assert the SHAPE*).
- **Five fences folded onto one derivation.** `value_change_guard` shares at one hop (it delegates to
  `repair_leak_guard`). `repair_leak_guard` had written the exact hazard above its own declaration and
  kept a private copy anyway — `§5` iter-189, *a stated-but-unfenced rule is a comment* — now discharged.
- **A census by EFFECT** (`TheFamilyIsCensusedByEFFECTnotBySPELLING`, 4 arms): discover every collector
  from source, call it, require the returned set to be `corpus_sources()` or `corpus_sources() | its
  declared extra`. Reconciled both ways. Live reading: **8 collectors, 114 each, extra 0.**
- **The literal arm KEPT**, with its blind spot declared — see `D-M257x-212-3`.

## Close — 2026-08-09

**Outcome:** the family's five private source-set derivations are folded onto one, with the single
genuinely-wider member's extra kept as a named sized derivation rather than dissolved; and the arm
written to prevent exactly this fork was measured seeing **2 modules of 9** because it enumerates one
literal spelling. The rationale that kept the five separate was a proposition **iter-211 had already
retracted** and which never reached the code.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (forty-fourth consecutive `closed-fixed`;
**no `P`/`N` reading taken — UNMEASURED, not unmoved**, `§9`) — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n — **counted, not felt: iter-212 is tik 1 of this run against a cap of 5** —
(6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-212-1` … `D-M257x-212-3` (see [`decisions.md`](decisions.md))

**Audit:** `/usr/bin/python3 -m pytest` (**8.4.2** / CPython **3.9.6**), **Python**, `stack-core` only —
**416 passed / 0 failed** across nine affected modules (`test_corpus_citation_guard`,
`test_guard_family`, `test_platform_predicate_guard`, `test_repair_leak_guard`,
`test_repair_leak_guard_mutation_battery`, `test_claim_twin_guard`, `test_value_change_guard`,
`test_retracted_pin_guard`, `test_anchor_construct_denominator`), 292 s; **31 passed** on the changed
test module alone.
**Live fences on the real rosetta tree, before and after the fold:** identical — see R5 above.
**RED-proof battery, mtime-mitigated (`§5` r77), both restores sha-verified:**
`markdown_structure_guard` re-forked to a narrower private walk → **the census arm RED**;
`platform_predicate_guard` stripped of its declared extra → **the staged `.claude/agents` arm RED while
the live census arm stayed GREEN** — which is precisely why that control must be staged (R6).
*Scope, stated rather than implied (`§5` r60): `stack-core` only, Python only, changed-code reach. No
whole-section run — the tree was edited during the iter. No Go, no TypeScript. The four non-`stack-core`
Python sections were read at iter-208 and not re-read since.*

**Side-deliverables:** none this iter.

**Routes carried forward:**
- `SURVEY-M257x-iter211-A-and-B-still-spell-their-own-scope` — **CLOSED by this iter.** The five now
  share the derivation; the constants remain as declarations and stay fenced against it.
- `SURVEY-M257x-iter212-a-retraction-does-not-reach-the-code-that-acts-on-it` — **NEW, and general.**
  iter-211's retraction landed in three prose ledgers and not in the comment justifying the design. No
  instrument in this milestone connects a retracted route id to the artifacts citing its claim.
  **Sizeable mechanically** (route ids are greppable; `route_disposition_guard` already owns the
  dispositions) and squarely a `TOK-08` class. Wants its own iter.
- `SURVEY-M257x-iter210-clone-drift-reads-a-third-corpus` — unchanged; now also the declared blind spot
  of the behavioural census (it has no collector function), reconciled both ways.
- All routes from iters 207–209, unchanged, plus the standing queue.

**Lessons:**
- **A retraction must reach the artifact that ACTS on the claim, not only the ledger that records it.**
  Three prose ledgers carried iter-211's retraction; the comment it falsified kept justifying a design
  for a whole iter.
- **An enumerator's blind spot is a property of HOW it enumerates.** A string matcher and a name-based
  discovery over one class had disjoint misses; the population is their union, and the honest fence
  declares each one's gap rather than promising either is total.
- **Fold, do not narrow.** Four of the five were exactly equal and folding them was free; the fifth was
  wider by rule and equal only by absence. Dissolving that difference would have silently shrunk a
  shipped fence — the size is 0 today and the *shape* is what must be asserted.
- **This iter's own census arm asserted a superset relation and went RED on its author inside a
  minute** — kept visible in the arm's docstring rather than quietly corrected.
