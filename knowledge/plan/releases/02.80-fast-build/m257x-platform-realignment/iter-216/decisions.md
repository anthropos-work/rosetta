# iter-216 — decisions

## `D-M257x-216-1` — the routed widening is sized at 9 claims, not 76 rows

**10** candidate tables · **76** rows · **9** rows yielding a matchable refuted form. On a 264
denominator that is **264 → ~273, +3.4 %**. The route stays queued and now carries its price.

## `D-M257x-216-2` — the binding constraint is QUOTATION, not the column name

The first mutation control — `MIN_FRAGMENT_CHARS` 30 → 1 — **did not fire**, and the null result is the
finding. A ledger row becomes a claim only if its claim cell **quotes** the offending sentence; the
candidate tables mostly do not quote at all, so no fragment floor can rescue them. The control that
does fire replaces the quote requirement with a whole-cell fallback: **9 → 69 of 76**.

So widening `_CLAIM_COL` alone buys ~9 claims. Getting the other 60 would mean matching **unquoted**
cell prose tree-wide — which is not a vocabulary change but a different, far less safe fence. Recorded
so the queued route is not later picked up as if the two were the same job.

## `D-M257x-216-3` — iter-215's headline is CORRECTED, appended, not substituted

iter-215's milestone-ledger line says *"a 44-row real ledger was hiding in them."* The table is real
and has 44 rows; it yields **2** matchable forms. Both are true and only the second is the size. The
correction is **appended** to the milestone ledger in this milestone's standing practice (never
substituted) and pinned by an executable arm.

## `D-M257x-216-4` — part of the residual is not this fence's to take

`claim's anchor`, `corpus line`, `doc says`, `what is actually there` are the **wrong-construct anchor
class**, which iter-42 assigned to a symbol-aware anchor check and `claim_ledger`'s own docstring
declines by name (*"they are somebody else's job"*). A widening that reached them would pull this fence
into another instrument's territory — the specificity its docstring warns a later reader not to lose.
