# iter-210 — decisions

## `D-M257x-210-1` — the repair is ONE derivation, not two re-synced copies

The cheap fix was to paste iter-209's widening into `retracted_pin_guard` too. It would have made both
fences read 114 today and left the family in exactly the state that produced the defect — two copies
under a comment asserting they are kept identical.

`§5` iter-190 settled this on a different pair (`_SVC_KEY` / `_COMPOSE_SERVICE_KEY`, two regexes reading
one construct out of one file and disagreeing on 5 of 9 names): **the repair is one shared derivation,
never two matching literals.** So the derivation moved to `fence_provenance`, the module every fence in
the family already imports, and both callers now route through it.

`EXTRA_SOURCES` is re-exported from both guards rather than deleted, because tests and callers name it;
the single *definition* is what matters, not the single spelling.

## `D-M257x-210-2` — `clone_drift_guard` is DECLARED private, not folded in

The sharing arm found a third private corpus walk within a minute of existing. Folding it in was
tempting and would have been wrong: its subject is backticked **sha tokens**, and the shared set adds
22 documents carrying **73** of them — occurrences 1,454 → 1,527, **+5.0 %** — with `CLAUDE.md` the
densest of them. That changes what the fence *finds*, which the `overview.md` pre-registered as the
condition to stop on.

So it is recorded in `DECLARED_PRIVATE` with the reason **and the measurement**, and reconciled in both
directions: an undeclared speller fails, and a declared entry whose module stopped spelling the
construct fails too. A waiver that outlives its subject reads as coverage.

## `D-M257x-210-3` — the five `SCAN_FILES` fences are NOT part of this class

`markdown_structure_guard`, `anchor_construct_guard`, `claim_twin_guard`, `repair_leak_guard` and
`platform_predicate_guard` each declare `SCAN_FILES = ("CLAUDE.md", "README.md")`. It is tempting to
call that a sixth, seventh, … copy of the same registry.

It is not. Those fences' subject **is** those two documents' prose; they are not answering *"which
documents are the corpus?"* at all. Merging them would be precisely the conflation this iter is about,
one grain up — the same error as reading iter-208's *"ten non-`stack-core` sections"* as *"ten Python
sections"*.

What is true and is routed rather than fixed: **nothing grades whether each of those five still wants
exactly those two files.** `SURVEY-M257x-iter210-five-fences-scan-only-the-two-root-docs`.
