**Type:** tik · **Protocol:** `corpus/ops/platform-alignment.md` · **Strategy:** `TOK-08`

# iter-194 — the registry every brief is told to trust could not read where harden passes route

## Re-survey changed the iter (`D-M257x-194-1`)

iter-193 routed this as an unrepaired, unfenced root cause. **Harden pass 42 had already found it** and
shipped a dedicated fence whose first arm *pins the exclusion* — so it was disclosed, not silent — and
declined the repair for a stated reason: it needs *"a disposition grammar the ledger does not use …
a design decision, not a corollary of a test."* That analysis was correct in every part. This iter makes
the decision instead of re-discovering the problem.

## The measurement (`D-M257x-194-2`)

**7** harden-origin route ids · **5** reachable · **2 LEDGER-ONLY**:
`FIX-M257x-h44-claim-census-guard-is-single-runner` and
`SURVEY-M257x-h42-size-preserving-mutation-proofs-unaudited`. Both are on this session brief's
open-routes list — and the same brief says to trust the fenced registry **over** any list in it.
Following that instruction drops exactly the two the registry cannot read.

## The repair, and the false RED it first produced (`D-M257x-194-3`, `D-M257x-194-4`)

The first cut keyed ledger dispositions above every iter, making the ledger the latest word about every
route, and the live registry reported `FIX-M257x-h36-labeled-prover-denominator` — open in pass 36,
closed by iter-193 — as a **CONTRADICTION**. A reach repair that manufactures a false RED is worse than
the silence it replaced. Position is now **derived** from the iters each pass says it covered
(`max(iter) + 0.5`), with an undated pass falling back to before-every-iter and that fallback **counted**
(measured: 0 of 44).

Then the reader was using the **iter** block grammar against the **ledger** and pulled 2 dispositions
from 45 passes — a plausible-looking number that does not read as a bug. The ledger writes
`**Routed forward:**` 12 times to the iter spelling's 1, and the first fix for that changed nothing
because `Routes?` does not match `Routed`.

| | before | after |
|---|---|---|
| ledger dispositions | 2 | **15** |
| routes with a disposition | 327 | **329** |
| total dispositions | 1,301 | **1,316** |
| distinct route ids | 374 | **379** |
| contradictions | 0 | **0** |

## Retiring pass 42's compensating machinery (`D-M257x-194-5`)

Its pin fired on cue and its message named the remedy. `LEDGER_ONLY_DISPOSITIONS` is emptied and
**fenced empty**; the five arms are re-based rather than deleted, and the exclusion-pin is inverted into
a **reach pin** so the arms cannot go vacuous if the ledger read is ever removed.

## Close — 2026-08-09

**Outcome:** the backlog registry every session brief is told to trust **over hand-written lists** could
not read the file harden passes route into — **7 harden-origin routes, 2 unreachable**, and both of the
unreachable ones sit on this brief's own open list. Repaired by reading the ledger with **the ledger's
own measured grammar**, at a position **derived** from the iters each pass declares it covered. Both
formerly-invisible routes now carry real dispositions; the registry reads **379 ids · 329 with a
disposition · 1,316 dispositions · 42 closures · 0 malformed · 0 contradictions**. Pass 42's
compensating registry is retired on its own written instruction and its exclusion-pin inverted into a
reach pin. A malformed id in **this session's own iter-193** was caught by the widened run and spelled
in full.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (twenty-sixth consecutive `closed-fixed`;
**no `P`/`N` reading taken — UNMEASURED, not unmoved**, `§9`) — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: n (third tik of this invocation) — (6) protocol-stop: n — (7) budget-exhausted: n —
Outcome: **continue**
**Decisions:** `D-M257x-194-1` … `D-M257x-194-6` (see [`decisions.md`](decisions.md))

**Audit:** `stack-core`, `/usr/bin/python3 -m pytest` (3.9.6) — **72 passed** across the **6** test
modules naming `route_disposition_guard` / `hardening-ledger` / `harden_origin`; and **118 passed**
across the wider route+family+census set. *Scope: `stack-core` only, Python only, changed-code reach
only (`§5` r60) — not the whole-section figure, and the other 10 sections remain unread.*

**Side-deliverables:**
- `iter-193/progress.md` — the abbreviated `` `FIX-M257x-h36-…` `` spelled in full
  (`D-M257x-194-6`); a truncated stem reads as live backlog in every brief quoting the queue.

**Routes carried forward:**
- `SURVEY-M257x-iter193-harden-routed-items-are-still-invisible-to-the-backlog-fence` — **CLOSED.**
  Both formerly-unreachable routes are in the population with dispositions; pass 42's compensating
  registry retired and fenced empty.
- `FIX-M257x-h44-claim-census-guard-is-single-runner` and
  `SURVEY-M257x-h42-size-preserving-mutation-proofs-unaudited` — **unchanged and still open**, but no
  longer invisible: both now carry a written disposition inside the fence, which is the whole point of
  the repair rather than a resolution of either.
- `SURVEY-M257x-iter194-other-milestones-ledgers-are-unaudited` — **NEW.** The ledger read is generic,
  but only `m257x` has a substantial hardening ledger; `m256` and `m257` were not measured for
  harden-origin routes at all. The reach claim above is about **one milestone**.
- `SURVEY-M257x-iter194-the-pass-position-derivation-is-untested-against-a-real-multi-range-pass` —
  **NEW.** `_pass_positions` takes `max()` over every iter number in a pass section, which is right for
  the shapes measured here (`iters 36–41`, `iter-01 … iter-15`) and unproven against a section that
  mentions an unrelated iter in prose. Conservative in the safe direction; not verified.
- `SURVEY-M257x-iter193-the-arithmetic-census-is-python-only` ·
  `SURVEY-M257x-iter192-printed-cardinality-census-is-one-section-of-eleven` ·
  `SURVEY-M257x-iter190-one-construct-two-regexes-is-unenumerated` ·
  `SURVEY-M257x-iter190-the-dual-reader-census-covers-one-section-of-eleven` ·
  `SURVEY-M257x-iter187-the-grain-question-is-unasked-elsewhere` ·
  `SURVEY-M257x-iter186-264-go-tests-have-never-been-read` ·
  `SURVEY-M257x-iter185-other-declared-populations-unaudited` · `D-M257x-145-3` (the user's to rule) ·
  `SURVEY-M257x-iter185-reach-percentages-predate-the-citation-widening` ·
  `SURVEY-M257x-iter184-the-standing-queue-wildcard-is-unbounded` ·
  `SURVEY-M257x-iter183-segment-grammar-refuses-28-multi-id-segments` ·
  `SURVEY-M257x-iter179-thirty-battery-tests-unrun` ·
  `SURVEY-M257x-iter180-relation-grammar-supports-only-equality` · `FIX-M257x-iter173-ledger-denominator` ·
  `SURVEY-M257x-iter175-readme-fence-index-is-16-of-27` ·
  `SURVEY-M257x-iter172-published-counts-predate-the-unit-fix` (observed half) ·
  `SURVEY-M257x-iter178-n-of-m-outside-clause-5-is-52-sites`
  — unchanged; open. Standing queue unchanged.

**Lessons:** **a registry that supersedes a list must reach everything the list does** — this one was
quoted as authoritative *over* the brief while being a strict subset of it for a whole class of route.
And two about repairing reach: **a reach repair that manufactures a false RED is worse than the silence
it replaced** (derive the ordering; never invent it), and **reading a source with another source's
grammar reports as an empty backlog** — 2 dispositions from 45 passes is a plausible-looking number, not
an obvious bug. Written into `platform-alignment.md` in this iter's commit.
