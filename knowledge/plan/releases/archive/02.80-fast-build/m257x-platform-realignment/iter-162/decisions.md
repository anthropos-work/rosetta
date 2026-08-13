# iter-162 — decisions

## `D-M257x-162-1` — fence the registry's COMPLETENESS, never its contents

The obvious repair for a 3-entry registry is a bigger registry. It is the wrong one: a 24-entry
hand-list is the same defect with a better number, and it rots the same way — silently, because
nothing compares it to the tree.

So `derivation_registry.DECISIONS` is still **declared** (contents are a judgement: some derivations
shell out to docker, some return a *verdict*, some would make the census match itself), but its
**completeness is derived**. `unclassified()` enumerates the population by AST and subtracts the
table; `stale_decisions()` goes the other way. Both are asserted empty, so a derivation added to rext
tomorrow turns a fence RED **with its own id in the failure message**, and the fix is one line.

**This is not theory — it fired twice inside its own iter.** `build_derivables` entered the
population the moment this iter rewrote it as a comprehension, and seven declines were caught reading
`"DECLINE:instance-state."` with nothing after the class. Neither would have survived a review; both
were caught by a mechanism, in the same commit that introduced them.

**Not claimed:** that the seven decline classes are the right partition, or that any individual
decline is correct. Those are judgements and they are on the record, per site, with reasons — which
is what makes them arguable later. What is *fenced* is that no site is missing.

## `D-M257x-162-2` — all four new candidates were graded at SOURCE, and all four are false positives

Reported as the headline rather than buried: widening 3 → 24 executed derivations bought **reach, not
defects**. That outcome was pre-registered in this iter's `overview.md` as acceptable *before* the
widening ran, so it is a result and not an excuse.

The grading matters more than the count. iter-158's rule — *a routed item's proposed repair is a
hypothesis, not a plan* — was applied to all four, and it is what stopped a mechanical "derive the
expectation" repair from landing on:

- a **golden over a synthetic compose** (`:182`), which would have made a parser test read its answer
  from the module it exists to cross-check;
- a **fuzz input** (`:1485`) and a **tokenizer contract** (`:1639`), where the literal is an
  *argument* to the code under test. Deriving those would test the platform instead of the tokenizer.

Each is exempted **at the site, with a reason**, adjacent to what it exempts (iter-161's tight-window
convention). Two of them name a limit the instrument did not previously have words for: it compares
**values**, and it cannot see a literal's **role**. iter-161 measured that it cannot see an input's
**provenance**; this is the sibling axis. Both are routed, neither is hidden.

## `D-M257x-162-3` — the `backend.md` anchor is repaired against its SUBJECT, and the general case is routed, not fixed

A one-line comment added to a test file turned `repair_postcondition` RED on a corpus citation five
lines away from its subject. The tempting reading is *"my edit broke a citation."* It is the opposite:
**the citation was already wrong, by 5 lines, at `c083819`, and green** — `:435` happened to be
`corpus = write_corpus(self.root, body)`, ordinary code, so the anchor guard's *closing-delimiter*
predicate had nothing to fire on. A **+1** shift put a `)` there and it went RED the same second.

The repair re-derives the **subject** (`§5` rule 22 / iter-22: re-derive the correction, not just the
anchor) — `CMS_RPC_ADDR=http://backend.internal.anthropos:8081` now lives at `:441`, and that is what
the citation names. **The offset was not simply bumped by one.**

**Deliberately NOT fixed here:** the detector. A rot-detector keyed on landing-on-a-delimiter measures
whether the rot was lucky, and widening it is a separate instrument with its own false-positive
budget. That is `FIX-M257x-iter138-anchor-rot-fence`, which now carries a measured instance instead of
a suspicion. Opening it in this iter would have been the 3rd unplanned line of investigation — the
scope-creep tripwire's own example.
