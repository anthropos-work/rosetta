**Type:** tik · **Protocol:** `corpus/ops/platform-alignment.md` · **Strategy:** `TOK-08`

# iter-196 — the last unread language, and the sentence that would have kept it unread

## What this iter actually caught (`D-M257x-196-1`)

iter-195 closed Go by testing an implication nobody had tested. The tempting next sentence —
*"Playwright needs a live stack, therefore unreadable"* — is the **identical shape, one language over**,
and it would have closed the route by assertion minutes after the rule against doing that was written.

It is **half-true**, and the false half is worth **424 tests**: `playwright test --list` enumerates with
no stack, no browser and no server.

| section | tests | files | notes |
|---|---:|---:|---|
| `playthroughs` | 215 | 45 | deps **borrowed** — ships no `node_modules` |
| `stack-verify` | 209 | 30 | own install |
| **TOTAL** | **424** | **75** | |

The 75 files corroborate iter-186/187's counts (45 + 30) exactly, independently derived.

**The generalisation:** an *unreadable-because* clause names a **capability**, and capabilities are
**per-verb**. Playwright cannot *run* without a stack and *can* list without one; the exclusion was
written against the strongest verb and silently inherited by every weaker one.

## The risk this iter creates, and the fence for it (`D-M257x-196-2`)

`424` now sits one line from Go's `2,714`, and they are **different kinds of statement**: one is a
population, the other a verdict. A listed test parses and registers; it has not passed.

So the fence is on the **vocabulary**, not only the count — `ts_census` returns `tests` / `files` /
`listed` / `borrowed_deps`, and an arm asserts there is **no `pass` or `fail` key**. If the two
vocabularies ever converge a reader can quote one as the other and be right about the words.

## Disclosure and controls (`D-M257x-196-3`, `-4`, `-5`)

A cold checkout **cannot** enumerate `playthroughs` unaided — it borrowed `stack-verify`'s install via
`NODE_PATH`, while the six Go sections needed nothing. Disclosed as `borrowed_deps`, and fenced so that
`playthroughs` shipping its own install later **fails loudly as good news** rather than passing
silently. A missing summary line is `listed: 0`, never `tests: 0` — the third instance in three iters of
one rule. The population is a **floor**, not a pin.

## Close — 2026-08-09

**Outcome:** the last unread language is enumerated — **424 TypeScript tests in 75 files** (`playthroughs`
215/45, `stack-verify` 209/30) — and the iter's real work was refusing the half-true sentence that would
have closed the route by assertion one iter after the rule against it was written. **Enumeration needs no
stack; execution does**, so the tool reports a POPULATION here and a VERDICT for Go and **fences the
vocabulary** (no `pass`/`fail` key) so the two can never be quoted as the same thing. The borrowed
dependency tree `playthroughs` required is disclosed and fenced; a missing summary line records as
NOT-LISTED rather than zero.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (twenty-eighth consecutive `closed-fixed`;
**no `P`/`N` reading taken — UNMEASURED, not unmoved**, `§9`) — (3) re-scope: n — (4) user-blocker: n —
(5) cap-reached: **y** — **and this time it is counted, not felt**: iters 192, 193, 194, 195, 196 =
**five** tiks (iter-195 mis-graded itself the fifth and was corrected in place at `9c0291f`) —
(6) protocol-stop: n — (7) budget-exhausted: n —
Outcome: **exit-5**
**Decisions:** `D-M257x-196-1` … `D-M257x-196-5` (see [`decisions.md`](decisions.md))

**Audit:** `stack-core`, `/usr/bin/python3 -m pytest` (3.9.6) — **100 passed** across the **7** modules
naming `suite_census` / `derivation_registry`. *Scope: `stack-core` only, Python only, changed-code
reach (`§5` r60).* The TypeScript figure in this iter is an **enumeration, not a run** — no TS test was
executed and none may be reported as passing.

**Side-deliverables:** none.

**Routes carried forward:**
- `SURVEY-M257x-iter195-typescript-is-now-the-only-unread-language` — **CLOSED for the ENUMERATION
  half only.** The population is measured and fenced; **no TypeScript test has been RUN**, and the
  execution half is re-routed below rather than folded into this closure.
- `SURVEY-M257x-iter196-no-typescript-test-has-ever-been-EXECUTED` — **NEW.** 424 e2e tests requiring a
  live demo stack. This is a genuine environment-gated population, and unlike the Go case that claim is
  now **measured** (they list; they do not run here) rather than assumed.
- `SURVEY-M257x-iter196-playthroughs-ships-no-node-modules` — **NEW.** A cold checkout cannot enumerate
  it unaided. Disclosed, not repaired.
- `SURVEY-M257x-iter195-the-go-reading-is-a-single-host-single-toolchain-sample` ·
  `SURVEY-M257x-iter194-other-milestones-ledgers-are-unaudited` ·
  `SURVEY-M257x-iter194-the-pass-position-derivation-is-untested-against-a-real-multi-range-pass` ·
  `FIX-M257x-h44-claim-census-guard-is-single-runner` ·
  `SURVEY-M257x-h42-size-preserving-mutation-proofs-unaudited` ·
  `SURVEY-M257x-iter193-the-arithmetic-census-is-python-only` ·
  `SURVEY-M257x-iter192-printed-cardinality-census-is-one-section-of-eleven` ·
  `SURVEY-M257x-iter190-one-construct-two-regexes-is-unenumerated` ·
  `SURVEY-M257x-iter190-the-dual-reader-census-covers-one-section-of-eleven` ·
  `SURVEY-M257x-iter187-the-grain-question-is-unasked-elsewhere` ·
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

**Lessons:** **an "unreadable because X" clause names a CAPABILITY, and capabilities are PER-VERB** — the
Playwright exclusion was written against *run* and silently inherited by *list*, which is the same
structure that kept Go unread for nine iters and was avoided here only because iter-195 had just written
the rule down. And the companion this iter had to invent: **when a new number is a different KIND of
statement from the one beside it, fence the VOCABULARY** — `424` enumerated and `2,714` verdicted are one
misquote apart, so the census asserts it has no `pass` key at all. Both written into
`platform-alignment.md` in this iter's commit.
