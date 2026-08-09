# iter-194 — decisions

## `D-M257x-194-1` — the routed framing was wrong, and re-survey is why that mattered

iter-193 routed this as *"harden-routed items are still invisible to the backlog fence"* and called the
root cause unrepaired. Re-survey found **harden pass 42 had already found it**, and had shipped
`tests/test_harden_origin_route_visibility_m257x.py` — a fence whose first arm **pins the exclusion**:

> *"Anti-vacuity, written against the SUBJECT (iter-94): if the guard ever starts reading the ledger,
> every assert below becomes vacuous while still passing. Pin the exclusion itself."*

So the exclusion was **disclosed by an instrument**, not silent — iter-186's rule had already been
applied. What was missing was the repair, and pass 42 said exactly why it declined it:

> *"It does not widen the guard's population (that needs a disposition grammar the ledger does not use
> — `**Routed forward — … (Fate 3).**`, not `**Routes carried forward:**` — and choosing one is a
> design decision, not a corollary of a test)."*

**That analysis was correct in every part**, including the part that turned out to be the whole
difficulty. This iter makes the decision rather than re-discovering the problem. *A routed item's
framing is a hypothesis too* — the fourth time that has been vindicated in this milestone, and the
first time the correction has been that the prior work was **better** than the route implied.

## `D-M257x-194-2` — the measurement: 7 harden-origin routes, 2 unreachable

| | |
|---|---|
| distinct harden-origin (`h{K}`) route ids | **7** |
| present in some `iter-*/progress.md` — the guard's whole population | **5** |
| **LEDGER-ONLY, unreachable by the guard** | **2** |

The two: `FIX-M257x-h44-claim-census-guard-is-single-runner` and
`SURVEY-M257x-h42-size-preserving-mutation-proofs-unaudited`.

**Both are on the open-routes list in this session's brief — and the same brief says to trust the fenced
registry *over* any list in it where they differ.** Following that instruction drops precisely the two
items the registry cannot read. A registry that is a strict subset of the list it supersedes has to say
so or stop superseding it.

It is also iter-193's measured root cause: `FIX-M257x-h36-labeled-prover-denominator` was reachable
**only by luck** — an iter had happened to re-type it. Visibility that depends on somebody re-typing an
id is not visibility.

## `D-M257x-194-3` — the first repair manufactured a FALSE RED, and that is the entry worth keeping

Ledger dispositions were first parked at a constant key **above every iter** (`HARDEN_TIMELINE_BASE =
100000`), on the reasoning that a pass number is not an iter number. That made the ledger the *latest
word about every route*, and the live registry immediately reported a contradiction:

```
RED FIX-M257x-h36-labeled-prover-denominator
    CLOSED at iter-193, asserted open at iter-100000
```

Open in pass 36, closed by iter-193 — chronologically correct, reported as a contradiction because the
ordering was invented. **A reach repair that manufactures a false RED in a registry every brief quotes
is worse than the silence it replaced.**

Repaired by **deriving** the position, because the ledger states it: each `## Pass NN` section names the
iters it covered (`iters 36–41`, `iter-01 … iter-15`). A pass sits at `max(iter) + 0.5` — strictly
between two iters, never colliding with one. Sections naming no iter fall back to `0.0`, before every
iter: the direction that can only ever **suppress** a contradiction, never invent one — and the fallback
count is **printed**, because an undated pass must be visible rather than assumed. Measured live:
**0 undated passes of 44.**

## `D-M257x-194-4` — reading a source with another source's grammar reports as an EMPTY BACKLOG

The ledger read was first wired to `BLOCK_RE` — `**Routes carried forward:**`, which *iters* write. It
extracted **2 dispositions from 45 passes**, and 2 is a plausible-looking number: it does not read as a
bug.

Measured, the ledger writes **`**Routed forward:**` 12 times** and the iter spelling **once**. And the
regex written for it had its own defect — `Routes?` matches `Route` and `Routes`, **not `Routed`** — so
the first fix changed nothing at all and the count stayed at exactly 2. Two different failures, one
symptom, and the symptom was a number that looked like a result.

With the ledger's own grammar:

| | before | after |
|---|---|---|
| ledger dispositions | 2 | **15** |
| routes with a disposition (m257x) | 327 | **329** |
| total dispositions (m257x) | 1,301 | **1,316** |
| distinct route ids (m257x) | 374 | **379** |
| closures | 41 | **42** |
| contradictions | 0 | **0** |

**Both formerly-invisible routes now carry real dispositions**, not merely membership. The intermediate
state — present in the population, no disposition — is the half-repair this iter had to walk past, and
it is named in a test so it cannot be mistaken for the finished thing.

## `D-M257x-194-5` — pass 42's compensating machinery is RETIRED, on its own instruction

Its pin fired the moment the guard learned to read the ledger, and its failure message says what to do:

> *"route_disposition_guard now names hardening-ledger.md. If it reads it, this whole module should be
> replaced by the guard's own disposition check rather than left passing beside it."*

Done. `LEDGER_ONLY_DISPOSITIONS` — a hand-written registry of routes the fence could not reach — is
emptied and **fenced empty**, because with the reach repaired every possible entry would be a second
registry for a route the first one already holds (its own author's warning). The five arms are
**re-based, not deleted**: the subject that must exist, the census that must not be empty, and the
membership-not-count comparison all still have work to do. The exclusion-pin is inverted into a
**reach pin** — if the ledger read is ever removed, the arms would go quietly vacuous, which is pass
42's own anti-vacuity reasoning pointed the other way.

## `D-M257x-194-6` — a malformed id in this session's own iter-193, caught by the guard

The widened run reported `'FIX-M257x-h36-' is not a route id — carried at iter-193`. It is **my own
prose**: iter-193's routes block abbreviated the id as `` `FIX-M257x-h36-…` ``, and a truncated stem
reads as live backlog in every brief that quotes the queue (`§5` rule 73). Spelled in full.

Worth recording because of *when* it was caught: the id was abbreviated for readability in a routes
block **one iter after** iter-193 wrote the lesson about denominators being read rather than
reconstructed. The registry's grammar is not a style preference, and the guard is the only thing that
enforces it.
