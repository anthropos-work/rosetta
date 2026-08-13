# iter-104 — decisions

## `D-M257x-104-1` — the tok is DELIBERATE, and the streak was checked before it was written

The 3-no-prog trigger does **not** fire. Last three tiks: iter-101 (`N` 28 → 24, progress), iter-102 (repair,
no reading), iter-103 (`N` 24 → 33). iter-101 moved the metric, so the window never reached three
consecutive no-progress tiks.

Recorded because the alternative — writing a tok and *calling* it triggered — would have been a claim about
the milestone's own history that its own ledger refutes, in a milestone whose whole subject is claims that
outrun their evidence. TOK-04 and TOK-05 set the precedent for author-initiated toks; this one states which
kind it is and shows the arithmetic.

## `D-M257x-104-2` — it does not terminate the call, and the rationale is bootstrap's, not convenience

A *triggered* tok terminates so a user can review a revision before the next tik commits to it. That value
exists when the user has watched a strategy stall and the revision is the new, unreviewed thing.

It does not exist here. **Every element of TOK-06's sequence is an item iter-103 already routed** —
`FIX-M257x-iter103-guard-tree-provenance`, `FIX-M257x-iter103-drift-fence-gap`, the two induction shapes
named in iter-103's band #10, and `FIX-M257x-iter103-read-union`. TOK-06 orders them; it does not open new
territory. So the loop continues, exactly as it does after a bootstrap tok.

**The honest limit:** this reasoning would be abusable as a general escape from the terminating rule, so it
is bounded here — a deliberate tok is non-terminating **only when it sequences already-routed work.** A
deliberate tok that proposes an unrouted direction should terminate like a triggered one.

## `D-M257x-104-3` — the Chapman sweep marks IN PLACE; it does not rewrite

Two standing point estimates survived iter-103's retirement, both in the milestone ledger:

| site | surviving claim |
|---|---|
| iter-101's entry | *"the residual is on the order of ~100, not ~45, and a zero reading is not near"* |
| iter-102's entry | *"the pool was probably always ~100"* |

`state.md` was already clean — iter-103's close landed the retirement there, including the explicit *"stop
quoting a point estimate from it."*

Both are marked with a **⚠ correction block in place**, following the convention the ledger already uses
(iter-102's own entry carries a `⚠ CORRECTED at iter-103 (DEF-4)` block for exactly this reason). Rewriting
them would erase what the milestone believed when it believed it — and this milestone's most useful records
are the ones where a conclusion and its later correction sit in the same paragraph.

**The corrections are asymmetric and are marked as such.** iter-101's *conclusion* (a zero reading is not
near) survives on the floors alone; only the point estimate goes. iter-102's second sentence (four
corrections to an underestimate, not a growing pool) survives intact; only *"probably always ~100"* goes.
A correction that took more than it had evidence for would be this milestone's own class, in this
milestone's own records, for the fourth time.

## `D-M257x-104-4` — the sequence puts provenance BEFORE the fences, and the reason is not tidiness

`FIX-M257x-iter103-guard-tree-provenance` is step 0 rather than step 3 because steps 1–2 **ship fences**, and
a fence's verdict is settled by the tree its configuration lives in. Ship the drift fence first and its first
green is provenance-unstated in exactly the way that produced iter-103's two false quotable conclusions —
and it would be a *new* fence, i.e. the one whose verdict nobody has any prior reason to doubt.

Cost of the ordering: one iter before the highest-value fence. Cost of the alternative: a fence whose founding
green cannot be re-checked.
