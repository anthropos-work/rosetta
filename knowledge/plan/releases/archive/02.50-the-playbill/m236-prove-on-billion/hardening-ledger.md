# Hardening Ledger — M236 prove on billion

## Pass 1 — 2026-07-20 — final

**Iters hardened this pass:** all milestone-touched code (cumulative scope — final mode)
**Tiks covered since prior pass:** all iters in milestone (10 iters; no prior harden pass ran)

**Scope manifest.** 14 code files in `rosetta-extensions` (`git diff 60eff14..HEAD`, ~914 insertions)
+ 5 corpus docs in `rosetta`. Grouped:

| Subsystem | Files | Touching iters |
|---|---|---|
| `stack-verify/e2e` (harness) | `lib/content-result-page.ts`, `tests/content-stories.spec.ts`, `tests/coverage.spec.ts`, `tests/probe-navigation.spec.ts`, `run-content-stories.sh`, `run-latency.sh`, `aggregate-content.py`, `.gitignore` | 04, 06, 07, 08, 09 |
| `stack-seeding` (projection) | `seeders/content_manifest{,_test}.go`, `seeders/content_nonsim{,_test}.go`, `seeders/content_stories.go`, `presets/content-manifest.json` | 05, 07, 08 |
| corpus docs | `coverage-protocol.md`, `latency-budget.md`, `content-stories-spec.md`, `playthroughs.md`, `verification.md` | 04–10 |

**Cross-iter hot spots** (files touched by ≥2 iters — final mode's defining scan):
`lib/content-result-page.ts` (**4 iters**: 04, 06, 07, 08) · `presets/content-manifest.json` (05, 07, 08)
· `seeders/content_manifest.go` (05, 07) · `seeders/content_nonsim.go` (05, 08) ·
`tests/content-stories.spec.ts` (04, 07).

**Coverage delta on touched files:**
- `stack-seeding/seeders` (Go): 96.1% → 96.1% stmts — already at bar; touched seeders 81–94% per-func,
  residual uncovered lines are DB-error paths requiring a live DB. No shallow tests written to move it.
- `stack-verify/e2e` harness (TS/Py/sh): **0 unit tests → 68**. The grader, the denominator, and the
  reading had NO test that ran without a seeded demo stack on a remote VM.
- Suite collectability: **0 tests in 0 files → 139 in 21 files** (see bug 3).

**Tests added:**
- `tests/content-result-page.unit.spec.ts`: 37 unit/edge (all 6 render shapes × good-page/broken-page pairs,
  universal bounce+404 guards across every shape, portal-settle boundary)
- `tests/content-pairs.unit.spec.ts`: 9 unit (denominator: per-session manager views, presence-only vs
  dropped, missing seat/path, academy player-vs-manager origin split)
- `tests/aggregate-content.unit.spec.ts`: 10 integration (drives the real script as a subprocess, asserts
  on **exit code**)
- `tests/content-route-contract.unit.spec.ts`: 8 cross-module contract (pass 2)
- `stack-verify/tests/test_green_gate_age.py`: 8 regression + class-sweep (pass 2/3)

**Bugs surfaced + fixed inline:**

1. **The aggregator failed OPEN** (`aggregate-content.py`, commit `a0f7945`). Denominator was `len(rows)`
   and the script **always exited 0**, so a run in which nothing executed printed `LANDED 0 / 0` and
   reported success — arithmetically also 100%. This is the milestone's own signature defect (a check
   scoring green off a subject that proved nothing) sitting one layer **above** the six render shapes where
   it was found repeatedly. Now fail-closed on empty ledger / any non-landing pair / any dropped pair, with
   optional `EXPECTED_PAIRS` to pin the denominator itself.
2. **Dropped pairs were invisible to the machine-readable reading** (`a0f7945`). A dropped pair writes no
   ledger line, so the denominator shrank silently and the survivors landing read as a clean sweep. The
   spec's own comment asserted this "can never quietly shrink to flatter the score" — true of console
   output, never of `content-stories.json`. Spec now emits `dropped.jsonl`; aggregator counts and fails on it.
3. **Collection poisoning — a cross-iter regression** (`a0f7945`). `content-stories.spec.ts` threw at
   **module scope** when `CONTENT_MANIFEST` was unset, so a bare `npx playwright test` collected **0 tests
   in 0 files** for the whole project — silently taking the two pre-existing unit suites
   (`coverage-manifest`, `section-assert`, 61 tests) offline since iter-04. Unset (being collected) now
   yields no tests; set-but-missing still fails loud.
4. **`\bchapter\b` could not match "13 chapters"** (`a0f7945`). The skill-path shape's structure check used
   the singular with a trailing word boundary, rejecting the plural — the exact string its own
   live-calibration note quotes for a completed path. A false FAIL, hidden only because seeded paths also
   render singular headings. The academy shape already spelled it `chapters?`; the two now agree.
5. **The runner swallowed its own verdict** (`a0f7945`). `run-content-stories.sh` ended on the aggregator
   with the result discarded → exit 0 whether 29/29 or 0/29 landed. Now `exec`s it. Notable:
   `run-coverage.sh` already carries this exact lesson in a comment ("swallowing it with `|| true` is what
   let a failed sweep exit 0 for four releases") and the newer runner reintroduced it anyway.

**Answering the milestone's signature question — "does any test pass against a broken subject?"**
Yes: **three** were found, all in the harness rather than the product. (1) the aggregator, which passed
against an empty run; (2) the whole `stack-verify/e2e` suite, which "passed" by collecting nothing; (3) the
grader itself had no negative tests at all, so it could not have distinguished a working implementation from
one returning `ok: true` unconditionally. The three defect-defending *unit* tests the iters found (iters 05,
07, 08) were correctly inverted by those iters and verified still-inverted this pass.

**Flakes stabilized:** none observed.
**Stop condition:** continue-to-next-pass — the Go↔TS route coupling and the iter-09 timezone fix are
untested; error paths on the touched runners unswept.

---

## Pass 2 — 2026-07-20 — final

**Scope:** cross-iter integration (the defining work of final mode) + the untested iter-09 fix.

**Coverage delta on touched files:** +8 contract tests, +6 regression tests. No source change beyond pass 1.

**Cross-iter integration finding — the route contract nobody was testing.**
The Go projection (`content_manifest.go`, `content_nonsim.go` — iters **05, 07, 08**) and the TS grader
(`shapeFor` in `content-result-page.ts` — iters **04, 06, 07, 08**) agree on the content-story routes by
**bare string prefix, across two languages**. Four iters touched each side; no test touched the join. The
failure mode is silent in the worst way: `shapeFor` **falls through to `player-scored`** for any unrecognised
prefix, so a Go-side rename of `/courses/` throws nothing, fails no Go test and no TS test, and grades every
academy page against a scored-report shape — reporting a correct render as broken. **That is precisely the
iter-08 defect, and after iter-08 nothing prevented its return.**

`tests/content-route-contract.unit.spec.ts` reads the **checked-in canonical manifest** (the same
honesty-gated preset the Go projection is pinned to) and asserts the grader understands every route in it:
per-product expected shape · no fall-through to `player-scored` unless the path really is `/sim/` ·
interview vs simulation manager surfaces kept distinct · manager paths still uuid-terminated · ai-labs
player-routeless · skill-path manager-routeless. It also **independently pins the denominator at 29** from
the canonical projection — both of the milestone's counting defects (31→18 per-product read, 31→29 gate
correction) were arithmetic a human had to notice, and neither ever went red.

**The iter-09 timezone fix had no test.** `test_green_gate_age.py` extracts the shipped `v_epoch=` line and
evaluates it under five zones spanning both sides of UTC — including a **half-hour offset**, which a
whole-hours patch would still get wrong — asserting the epoch is identical **and** equals the true UTC
instant (zone-independence alone would also be satisfied by a consistently wrong constant). Plus a
class-level sweep for any unpinned `date -jf` anywhere in `e2e/`, because the bug is a class, not an
instance. Verified on this platform that the fix genuinely engages: BSD `date -d` fails with "illegal
option" so the GNU branch cannot mask the fallback, and `TZ=UTC` shifts the parse by exactly the 7200 s
CEST offset.

**Mutation verification (break the subject, confirm red — then restore):**
- remove `TZ=UTC` from `run-latency.sh` → **5 of 6** green-gate guards go red. Restored; diff confirmed empty.
- rename `/courses/` → `/course/` in `shapeFor` → **2 of 8** route-contract tests go red. Restored.

**Bugs surfaced + fixed inline:** none new (both gaps were missing tests, not defects).
**Flakes stabilized:** none observed.
**Stop condition:** continue-to-next-pass — error paths / boundary inputs on the touched runners unswept.

---

## Pass 3 — 2026-07-20 — final

**Scope:** dimension 3 (error paths) + dimension 2 (boundary inputs) over the touched runners + aggregator.

**Coverage delta on touched files:** +4 tests (2 aggregator error-path, 2 runner-input guards). Cumulative
harness unit tests: **0 → 72**.

**Bugs surfaced + fixed inline:**

6. **A mistyped stack number silently swept the DEV stack** (`01cb153`). Every runner computes
   `OFFSET=$(( N * 10000 ))`, and bash evaluates a non-numeric `N` to **0 with no error** — so
   `./run-latency.sh abc` pointed every probe at offset 0, the ports the dev stack sits on, and would have
   reported those timings as demo-N's. `run-latency.sh`'s own stated premise is that it *refuses to measure
   a stack that is not what it claims to be*; measuring an entirely different stack is that failure wearing
   a different hat. Both milestone-touched runners now reject a non-integer `N` with rc=2.
7. **A truncated ledger line cost the entire reading** (`01cb153`). The spec appends one line per completed
   pair, so a run killed mid-append leaves a partial last line — realistic, not hypothetical. Bare
   `json.loads` raised before `content-stories.json` was written at all, so the operator got a traceback and
   **no reading**, instead of "2 rows read, 1 truncated". A row missing `ok` (schema drift) `KeyError`'d the
   same way. Exit code was already correctly nonzero; what was lost was the diagnosis and every row that
   *did* parse. Bad lines are now counted, carried into the report as `malformed`, and named in `problems`.

**Deliberately NOT fixed (scope boundary, routed forward — see Deferrals below):** `run-coverage.sh` and
`run-hiring-render.sh` carry the identical `OFFSET=$(( N * 10000 ))` hazard but were **not touched by this
milestone**. Fixing them would grow the harden footprint into files no iter touched. Recorded for
close-milestone to route.

**Flakes stabilized:** none observed.
**Stop condition:** continue-to-next-pass — flake gate + full-suite verification not yet run.

---

## Pass 4 — 2026-07-20 — final

**Scope:** flake gate + full-suite verification + knowledge backfill. No new dimension findings.

**Flake gate:** 3 consecutive clean runs of the full TS unit battery — **125 passed** each run, identical
(9.0s / 9.1s / 9.0s). 3 consecutive clean runs of `test_green_gate_age.py` (8 tests). Go suite re-verified:
`stack-seeding/seeders` **96.1%**, ok. **0 flakes.**

**Coverage delta vs pass 3:** 0% (no new tests). Dimension scan found nothing new.

**Pre-existing failures observed — NOT introduced by this milestone, NOT re-routed** (`stack-core` is
untouched by M236: `git diff 60eff14..HEAD -- stack-core` is empty):
- `test_story_org_count_guard.py` — `stories.seed.yaml` ships **4 orgs** (Meridian Talent, the hiring org,
  landed at M224/M226 — **v2.4**) while **18 doc lines** across `corpus/` and two shipped skills still say
  3. A v2.4-era doc drift with a named handler.
- `test_m220_mutation_battery.py` — the harness reports its own **unmutated** subject failing, so it cannot
  produce a green and no mutant result is readable. Pre-existing; belongs to the standing demo-stack
  failure set.
- `test_stack_registry.py` — 62 tests OK (the stray `1` in output is not a failure).

**Knowledge backfill:**
- `corpus/ops/demo/coverage-protocol.md` → two new subsections in the M236 content-stories section: **"The
  reading must be fail-CLOSED"** (the 0/0 rule, dropped-pairs-in-denominator, `EXPECTED_PAIRS`, the
  exit-code rule + the observation that `run-coverage.sh`'s own written-down lesson did not propagate to its
  sibling, truncated-ledger legibility, the wrong-stack hazard) and **"Prove the test fails (mutation, not
  coverage)"** (the good/broken pair discipline, the two mutation verifications, the cross-language route
  contract). Carries the generalisation: *ask of every layer that reports a number — what does it print
  when nothing happened?*
- `corpus/ops/demo/latency-budget.md` → the green-gate age check is now regression-tested (incl. the
  half-hour-offset case and the class sweep) + the new non-integer-`N` refusal + the two unguarded siblings.
- `corpus/ops/demo/content-stories-spec.md` → §4 gains the **cross-language route-prefix contract** as an
  explicit, pinned invariant, with the instruction that changing a route in `contentProductRegistry` is
  expected to fail that spec.

**Bugs surfaced + fixed inline:** none.
**Flakes stabilized:** none needed.
**Stop condition:** **stabilized** — coverage delta 0% (< 2%) across passes 3→4 AND the Phase 2 dimension
scan found nothing new. 4 passes of the 5-pass final-mode cap used.

---

## Deferrals routed forward (Fate 3, named handler: `/developer-kit:close-milestone`)

- **`run-coverage.sh` + `run-hiring-render.sh` non-integer-`N` guard.** Identical
  `OFFSET=$(( N * 10000 ))` hazard to the two runners fixed in pass 3 (a mistyped stack number silently
  sweeps the dev stack). Not fixed here because neither file is in the milestone-touched manifest — fixing
  them would grow the harden footprint into files no iter touched. ~6 lines each, no design decision.

*(The milestone's existing Fate-3 items — the anonymous academy catalog twin, the client GraphQL endpoint on
a non-offset port, remaining v2.4-era method docs, and the 14 standing pre-existing demo-stack test failures
— already carry named handlers and are deliberately not re-routed here.)*
