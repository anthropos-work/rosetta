# iter-74 — decisions

## `D-M257x-74-1` — the ambiguous class grew because the fence started seeing more, and it is provable by partition

`CHECK-M257x-iter73-ambiguous-grew` asked whether `ambiguous` going **12 → 39** meant *a
corpus-writing habit worth changing* or *a fence limitation*. It is neither, and the answer is not a
judgement: partitioning the class by **which regex alternative matched** — the only thing iter-73's
widening changed — gives `bare-code 27 · md 0 · path 12`. The `path` count is **iter-71's 12,
unchanged to the citation**; every one of the 27 belongs to a class that could not be counted at all
before iter-73, because a bare `<name>.<ext>:N` never reached `resolve()`.

**The rule this generalises:** when a counted class grows in the same run that reach grows, split it
by the reach dimension before touching anything. A class that grows *entirely inside the newly-reached
partition* is the instrument improving; a class that grows *inside the old partition* is the corpus
moving. Those need opposite responses, and the count alone cannot tell them apart. Recorded against
§5 rule 16 (*an UNREAD metric is indistinguishable from an UNMOVED one*) — this is its sibling on the
other side: **an unread metric that becomes read looks exactly like a metric that got worse.**

## `D-M257x-74-2` — in a table the block is the CELL, and the two guards now agree about it

`anchor_construct_guard._block_of` walked to the nearest blank line. **A markdown table has no blank
lines between its rows**, so a citation inside a table took the *whole table* as its window and every
sha named in any row chose the ref for the citations in every row.

That contradicts **§5 rule 33**, derived in iter-63 and implemented since then in
`platform_predicate_guard._pin_window`: *"a pin's scope is the claim's own block — a markdown CELL in
a table, a wrapped sentence in prose."* Two guards, two definitions of *block*.

Measured instance, before the edit — `external_services.md`'s provider table: rows 539 (**AWS
Bedrock**) and 540 (**Mistral**) carry no pin and were being read at `b948604`, a ref named by row
542 (**Anthropic Direct**). Rule 33's *"the pin crosses a ROW boundary"* mechanism, in the guard that
**chooses refs** rather than the one that grants exemptions — and the consequence is worse there: an
exemption silences a check, a mis-chosen ref makes the guard **adjudicate a citation against a file
the document never named**.

Landed: cell window in a table (`|`-delimited by the format itself, so the boundary is the table's,
not ours), blank-line block in prose **unchanged**, and a table row is a boundary for the prose walk.
`col=None` falls back to the whole row — the conservative direction, because a wider window can only
make a citation `ambiguous`, never make it read at a ref nobody wrote down. Self-references pass
`col=None` deliberately: their offsets are matched against a *substituted* copy of the line, so the
column would cut the wrong cell.

**Dry-run before landing** (iter-73 lesson 1): predicted `ambiguous 39 → 24, block-pinned 45 → 48,
default 63 → 75, 0 findings`; live it reproduced exactly.

## `D-M257x-74-3` — a blockquoted table row is still a table row, and the first draft of the fix contained the defect it removes

Reading the residual, `storage.md`'s v9.0 fold block is a **blockquoted** table (`> | side | … |`).
`_TABLE_ROW`, copied verbatim from the sibling guard, anchors at `|` — so every one of those rows
failed the row test, fell into the prose branch, and took the whole quoted table as its window. *The
exact defect being fixed, surviving inside the first draft of the fix for it* — the same shape as
iter-68's default-argument bug (iter-68's own warning reappearing as a scoping bug inside the fix for
it).

Admitted on a **count, not on the document**: `rg -c '^\s*>\s*\|.+\|\s*$'` finds **75 blockquoted
rows across 14 files** against 2197 plain rows in 86 files. A construct with 75 instances is a
construct.

Final: **ambiguous 39 → 20 · block-pinned 45 → 45 · default 63 → 82 · 177 resolved · 0 findings.**

## `D-M257x-74-4` — the residual 20 are adjudicated by reading, not closed by rule

An ambiguity only matters if it can change an answer. Classifying each ambiguous citation at **every
live sha its own window names** as well as at the default: **19 AGREE, 1 DIVERGES.**

The one is `platform-migration-status.md:71` → `app/internal/jobsimwiring/wiring.go:123`, whose cell
names three live shas — and the cell itself settles it: *"`app` @ `9d00a313` v1.367.0 — re-resolved
M257x iter-68; the first four stood at … at `b948604` v1.366.0 and at … `5ba17044` v1.363.2."* It is
contrasting them deliberately and names the one it asserts, which is the ref the default ladder
picks.

**No repair, and no new rule.** Picking a sha out of a contrast block would be §4 Trap A — a rule
fitted to a sentence — and `block_ref` already commits to falling back and **counting** instead. What
changed is that the fallback is now known to be **inert in 19 of 20 cases**, measured, rather than
assumed.

## `D-M257x-74-5` — a hand-off whose count was measured and whose head list was not

The orchestrator carried two numbers forward. **39 reproduces exactly.** For the other:
*"92 bare citations still unresolvable"* — the **count** reproduces (91 distinct bare-code citations
across 103 sites); the **head list does not**. iter-73 wrote *"`gen.py` x10, `intelligence.go` x8,
`main.go` x7"*; measured now: **`up-injected.sh` x32, `intelligence.go` x5, `20260722104506.sql` x5,
`main.go` x2, `gen.py` x0**. `gen.py` is cited eleven times and **never** in `file:N` form — every one
is a RANGE (`` `gen.py:484-492` ``), which the regex cannot match.

So a single routed item carried a count from one instrument and heads from another. §5 rule 12 says
*say which INVOCATION produced the number, not just which tool*; this adds the case where **one
sentence quotes two invocations** and reads as one measurement. Routed to iter-75 with the
derivation attached, so the class that gets repaired is the class that was measured.

Also derived, and carried with it: of the 239 sites the guard reports unresolvable, **88 are URLs**
(`http://backend:8083`, and the `ENV=http://host:port` form) — not citations at all. The
unresolvable denominator has never been net of them, and any repair target taken from it inherits
that.

## `D-M257x-74-6` — a routine `git fetch` moved the adjudication ref mid-iteration, and the delta cost nothing

P3 at close found `app`'s `origin/main` had moved from `9d00a313` (v1.367.0) to `7177374`
(`v1.367.0-4-g717737471`) — **four commits, during this iteration**, and the move happened because
the P3 check itself fetched. The `auto` ladder prefers `origin/main`, so **82 default-adjudicated
citations silently changed the file they are read at, with no diff anywhere** (§5 rule 26).

Both ref-aware guards were re-run at the new ref and both stayed GREEN, reach unchanged. The
citation delta (`D-M257x-59-3`, §7 rule 4's second half) over every corpus citation landing in the
`app` clone: **48 HELD · 1 MOVED · 0 DEAD** — and the one MOVED is `clerk-integration.md:103`'s
`` `app/go.mod:31` @ `5ba17044` ``, which is **true at the ref it names** (`clerk-sdk-go/v2 v2.7.0`)
and false at both newer ones. The guard reads it at its pin and is right to; **the delta script
forced both refs and manufactured a move that isn't one** — iter-69's lesson (*a pin is a date; do
not repair a dated claim onto a floating ref*) reappearing inside a measurement written the same
day. Net real breakage across the advance: **zero**.

Worth carrying: the milestone has now measured two app advances with opposite results — iter-58's
moved **22 of 23** `main.go:N` citations, this one moved **0 of 49**. The difference is not the size
of the advance; it is **whether the advance edited the cited files**. A count of commits predicts
nothing, which is the argument for measuring the citation delta rather than estimating it.
