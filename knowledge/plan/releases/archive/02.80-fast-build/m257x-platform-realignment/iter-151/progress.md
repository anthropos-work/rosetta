**Type:** tik — under `TOK-08`.

# iter-151 — an absent-value default is only as dangerous as the side that reads it

## The census

Every non-comment, non-prose line in the monorepo naming `STACK_PROJECT` or `STACK_OFFSET`, partitioned
by section.

| where | reads |
|---|---|
| `stack-verify/` — the read-only probe section (`lib/target.sh`, `lib/services.sh`, `lib/readiness.sh`, `live/autoverify.sh`, its own tests) | **all of them** |
| any write-side section — `stack-seeding` · `stack-snapshot` · `stack-injection` · `stack-secrets` · `demo-stack` · `dev-stack` · `playthroughs` · `clerkenstein` · `alignment` | **0** |

`autoverify.sh` always `export`s `STACK_PROJECT` explicitly before delegating, so even inside the probe
section the unset path is reached only by a hand-run `verify.sh`.

## The grading, which is the point

iter-147's worry — *an unset default that silently targets the main dev stack* — is **conditional, not
live**. Unset, `STACK_PROJECT` resolves to `anthropos` (the developer's own stack) and `STACK_OFFSET`
derives from it. Today the worst that can produce is **a probe pointed at the wrong stack**: iter-148's
false-report class, loud and recoverable, and the class iter-148 already repaired at the caller. What it
is **not** is a seeder, a snapshot replay or an injection writing into someone's dev stack because a
variable was missing from an environment — because **no writer reads either variable**.

That is a property of the current code and nothing was holding it.

## What landed

`stack-core/tests/test_stack_target_vars_are_read_side_only.py` — the partition, fenced:

- **the write-side assertion**, whose failure message says the repair explicitly: *do not add the section
  to the read-side list; make its resolution REFUSE an absent value* (iter-147's own rule). The day a
  write path starts resolving its target from these variables is the day the unset default becomes a
  data-loss default, and it should fail in a test rather than on a dev stack;
- **anti-vacuity on the subject** (`§8`, iter-94's shape) — the variables must still be read *somewhere*,
  and the probe side must still read them. A rename would otherwise empty the census and leave every
  assertion trivially true;
- **a RED-proof**, because the census returns zero and `D-M257x-149-3` says a zero must prove its
  instrument: the write-side predicate is shown matching a synthetic
  `stack-seeding/cmd/fake.go` reading `STACK_PROJECT`.

## Tests

`test_stack_target_vars_are_read_side_only.py` — **3/3 green**. **Not run:** the rest of `stack-core`
(~20–35 min) and the other four sections — `§5` rule 60's scoped default; this iter adds one file and
changes no existing behaviour.

## Close — 2026-08-08

**Outcome:** the last open arm of the absent-value class closes with a **falsification**: every read of
`STACK_PROJECT` / `STACK_OFFSET` is probe-side, so the unset-defaults-to-the-main-dev-stack hazard is a
wrong-answer hazard and not a wrong-write hazard. Fenced so it stays that way, with the repair for the
day it changes written into the failure message.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — **4 of 5**, unchanged. **No `N` reading was taken, so no `N` movement is claimed**
(`§9`'s guard-rail 1, in its required words).
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (**iters 135–151 took no reading, so the metric is UNMEASURED not unmoved — `§9`'s iter-type refinement; `TOK-08`'s sealed refutation branch bars an agent-authored successor in any case**) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (**3 tiks this session**) — (6) protocol-stop: n — (7) budget-exhausted: **y** — Outcome: **exit-7**
**Decisions:** `D-M257x-151-1` (grade an absent-value default by the SIDE that reads it — read-side is a
wrong answer, write-side is a wrong write; fence the partition, not the default).
**Side-deliverables:** none.
**Routes carried forward:**
- `SURVEY-M257x-iter147-absent-value-class` — **CLOSED by this iter** (`STACK_SERVICES` at iter-148,
  `rosetta-demo`'s flags graded sound there, these two here).
- `SURVEY-M257x-iter150-partition-completeness-elsewhere` — **narrowed, not closed.** A first look says
  the two sibling guards do not hold the same closed-world assumption (`markdown_structure_guard`'s
  `_DOUBLE_WORDS` is a heuristic word list, not a partition of an owned vocabulary;
  `evidence_visibility_guard` has none). Grade the rest of the family before believing that.
- `SURVEY-M257x-iter149-declared-lists-unfenced-against-layout` — closed at iter-150 ·
  `SURVEY-M257x-iter148-registry-is-hand-maintained` · `FIX-M257x-iter145-sha-baseline-drift` ·
  `-iter145-migrate-race-needs-a-host-postgres` · `-iter145-green-but-stale-graphql-mentions` ·
  `SURVEY-M257x-iter144-orphan-arm-is-the-residual` · `FIX-M257x-iter144-correction-vs-retraction-unfenced` ·
  `FIX-M257x-iter143-wrong-head-is-unfenced` · `-iter143-scope-derivation-by-grep` ·
  `-iter143-appending-to-the-protocol-doc-rots-the-ledger` · `FIX-M257x-iter142-value-change-articles` ·
  `-iter142-path-arm-window` · `-iter142-tier-b-underflag` · `FIX-M257x-iter135-adjudicated-live-defects` ·
  `-iter140-receipts-not-checkable-here` · `-iter140-receipt-fence` · `-iter138-anchor-rot-fence` ·
  `-iter134-fence-family-has-no-shared-predicate-layer` · `-iter133-two-fives-need-a-fence` ·
  `-iter132-suite-walltime-is-not-a-measurement` · `-iter131-predicate-sets-not-enumerated`.
**Lessons:**
0. **An absent-value default is not a defect on its own — the reader's side decides.** iter-147 booked
   these two as an open hazard on the shape of the default alone. The census that settles it is not
   "where is this read" but "is any reader a writer", and the answer changed the finding from a defect
   into a property worth fencing.
1. **Write the repair into the failure message when the fence guards a boundary that will legitimately
   move.** Someone will one day have a good reason to give a write path a target variable. The message
   tells them the right move (refuse an absent value) instead of the easy one (add the section to the
   allowlist) — which is the only difference between a fence and an obstacle.
