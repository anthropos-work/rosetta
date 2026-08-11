**Type:** tik — under `TOK-08` (*census the mechanical classes; stop sampling them*).

# iter-280 — the disposition registry's reasons were prose; they are graded now

## What the re-survey changed

The orchestrator pointed at `roadrunner:merged-into-app` as false-since-iter-137. **It was already
fixed** — harden pass 71 corrected it to `deleted-not-merged` and the tree was green. Working the named
target verbatim would have re-done landed work; the standing *route-lists-go-stale* rule caught it.

What pass 71 left is what it wrote down while fixing it, and the code still agreed:

> *"Nothing grades this token — `service_registry_guard` reads the NAME and stores the reason as free
> prose."*

So the target substituted, under the same strategy: **not the row — the population it sits in.**

## The population, enumerated and graded

`SERVICES_NOT_IN_PLATFORM_COMPOSE`, 5 rows. Before this iter: **1 graded factually** (by a hardcoded
per-row test naming one service and one token), **4 graded only against a comment block.**

Arm F now grades each token against the `app` clone, because that is what a token claims:

| row | token | predicate | measured |
|---|---|---|---|
| `jobsimulation` | `merged-into-app` | `app/internal/<pkg>/` EXISTS | PRESENT ✓ |
| `cms` | `merged-into-app` | EXISTS | PRESENT ✓ |
| `storage` | `merged-into-app` | EXISTS | PRESENT ✓ |
| `roadrunner` | `deleted-not-merged` | does NOT exist | ABSENT ✓ |
| `directus` | `rext-injected` | no claim about `app` | n/a |

**4 tokens graded, all true** — matching the verdict sealed in `d8e3c4b4` before the fence existed. The
deliverable is the fence, not a repair, which is TOK-08's shape.

The verdict line now states its own reach: *"…; 4 disposition token(s) graded against the app clone."*

## Fail-closed, proven rather than asserted

The `app` clone is **derived** from the platform clone (its sibling) — this family's thrice-stated rule,
so two guards cannot end up pointed at different clone sets. Unreachable clone → **CANNOT-MEASURE
(exit 2)**, never ALIGNED.

Proven with the precondition absent: `APP_CLONE=/nonexistent/app` → `rc=2`, arm F named, no `ALIGNED` in
the output.

## The battery (4 kills + 1 no-op that must survive)

- **the roadrunner regression replayed** — `deleted-not-merged` → `merged-into-app` fires
  `F/DISPOSITION-FALSE`, naming the package it looked for. The exact string that survived four months
  and pass 71's vocabulary check.
- **the opposite polarity** — `cms:merged-into-app` → `deleted-not-merged` fires. Deletion-mutants alone
  cannot tell a fence that reads the predicate from one that dislikes a token.
- **an unknown token** — fires `F/UNKNOWN-DISPOSITION` rather than being ignored.
- **the fail-closed control** above.
- **declared-GREEN no-op** — a token that grows an explanation (`merged-into-app - folded by …`) still
  grades. Without it, four REDs cannot distinguish a discriminating fence from a brittle one.

## What landing it broke, and why the fixture was the defect

Arm F turned **14 pre-existing arm A–D controls RED at once**. Their fixture wrote the compose to a bare
tmp root, so the sibling derivation walked out of the fixture and every control became CANNOT-MEASURE.

The tempting read — *"arm F is too strict, make it skip"* — is the fail-open trade, and it was refused.
**The fixture was modelling a directory shape production never has.** It now mirrors a real stack.

One control, `test_declaring_a_departure_is_how_you_go_green_again`, was **prescribing a repair that
asserts a false fact**: it cleared arm A by declaring the departed `gotenberg` as `merged-into-app` — a
fold that never happened. That is the roadrunner trade encoded in a test. Repaired to the true token.

## Side deliverables (separate facts; they do not move the close status)

1. **A fence citation naming a fence nobody built.** `services.sh` said an undeclared stack-injected
   service *"turns `service_registry_guard.py` **arm E** RED."* There is no arm E — `git log --all -S`
   over full history returns **0 commits** for both the array name and the phrase in that file. The check
   is real and lives in `stack-verify/tests/test_scope_union_m257x.py`. **Both halves landed in the same
   commit** (`83ada03`, iter-153): the citation described the plan, not the build, and the two diverged
   inside one commit. Unnoticed for 127 iters.
2. **Two corpus rows declaring a production service that production destroyed.** Grading dispositions on
   the corpus side — the orchestrator's *"expect disagreement in both directions"* — the fenced services
   table had `roadrunner` **and** `graphql-wundergraph` at prod = `live-standalone`, which §1 defines as
   *"its own process, still on the traffic path"*, while each row's own evidence column says the service
   is destroyed. iter-124 repaired the prose of both and left both tokens. **Enumerated the class rather
   than repairing the site** — which is how the second one was found. Corrected to `decommissioned`.
   **Scope stated precisely:** the cells were made to agree with measured prose already in-row;
   `infrastructure` is not in the clone set and was not re-read — a limit the alignment guard prints
   itself (*"UNCLONABLE head 'infrastructure' x10 … NOT read"*).
3. **`services.sh`'s own "nothing grades this token" comment**, falsified by this iter's code and
   rewritten in the same commit as it — shipping otherwise would have committed the defect the arm exists
   to fence, in the file it fences.
4. **Pass 71's "Both directions" docstring** implemented one direction. The second is implemented, against
   the guard's real vocabulary rather than against more prose.
5. **`FIX-M257x-278-census-substrate-tests-hardcode-an-absolute-ROOT`**, in the file this iter had to edit
   anyway: the two candidate compose paths were absolute, so every platform-coupled control could only run
   on one box and would silently skip elsewhere. Derived now. This does not close the route — other files
   still carry it — it stops this iter adding to it.

## The whole-section suite came back RED, and the triage is the interesting part

First full run against the final tree: **22 failed / 2219 passed (31m15s)**. Nothing was committed on it
and no green was claimed. Triaged against a **frozen clone of the pre-iter commit** rather than by
judgement — three distinct causes wearing one colour:

- **2 were mine, and both were a fence being right.** `graphql-wundergraph.md` had no status banner the
  fence accepts (the corpus correction propagating into a second document); and a dated literal in the
  arm F docstring breached `DOCSTRING_LITERAL_CEILING` 241 > 240. Fixed by rephrasing, **not** by bumping
  a ceiling — back to 240.
- **1 was NOT mine.** `TEST_MODULE_LITERAL_CEILING` was **already RED on the control tree** at 655 > 653:
  harden pass 71 added two `0 commits` literals and shipped the suite RED. Repaired — the same phrase,
  for the same reason, in the file this iter already owns — taking it to **653, at the ceiling, no bump**.
- **The rest are full-run interaction**, all GREEN standalone on **both** trees.

A single RED count would have read as *"iter-280 broke 22 tests"*, and two thirds of that would have
been wrong.

**After the three repairs: 13 failed / 2224 passed (31m39s).** The residual sits in four files —
`test_suite_census` (7), `test_repair_leak_guard_mutation_battery` (4), `test_suite_census_population`
(1), `test_battery_stage` (1) — every one of which passes standalone on both trees.

Root cause of the largest group: **`test_suite_census` writes its probe module into the very directory it
censuses**, then asserts the runner counted 3 tests; inside the full run it reads 1 and an error. The
census is a member of the population it measures.

**The attribution I wanted could not be taken, and that is itself a finding.** A full-suite run on the
frozen control clone would have settled *pre-existing vs mine* outright — it **aborts at collection**,
because the census substrate resolves an absolute ROOT and does not survive cloning. That is the open
`FIX-M257x-278`, biting for the third consecutive iter, and this time it **blocks the exact measurement
needed to triage a RED suite**. So: the residual 13 are **not attributable to iter-280** on the standalone
evidence and the 22→13 direction, and **not proven pre-existing** either. Both halves stated, neither
inflated into the other.

## Close — 2026-08-11

**Outcome:** the ungraded half of the disposition registry is a census: **4 of 5 rows went from
comment-block-checked to platform-checked**, fail-closed and RED-proven, with the historical
`roadrunner` defect now caught **by rule** instead of by one hardcoded test. Two false *production-live*
claims found and corrected in the corpus's own fenced map by running the same census the other way. A
RED whole-section run was triaged against a control tree, and it surfaced a **pre-existing** ratchet
breach that harden pass 71 had shipped unnoticed. **Clause 5 NOT re-measured and no `P` is claimed.**
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: **y** — (5) cap-reached: n (1 tik) — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **exit-4**

**Why exit-4 and not exit-7:** the section-suite gate is **RED (13 failed)** at close. Everything
attributable to this iter was found and fixed (22 → 13), the iter's own gates are green — whole
`service_registry_guard` suite **38 OK**, guard family **30 GREEN · 0 RED**, both ratchets back at their
ceilings — and the residual is confined to four files that pass standalone on both trees. But a RED
protocol gate is a user-blocker by rule, and this one carries a real decision: the residual is a
**test-isolation** defect (`test_suite_census` censuses the suite it belongs to) whose triage is itself
blocked by `FIX-M257x-278`. Whether to spend an iter on that pair, or accept a RED section gate, is not
a call to make silently inside an iter that was about something else.

**Decisions:** `D-M257x-280-1` … `D-M257x-280-9` (see `decisions.md`), including two recorded
self-defects: a repeated clock-discipline violation whose structural cause was found and fixed, and a
malformed commit trailer recorded rather than force-fixed.

**Side-deliverables:** the five listed above.

**Routes carried forward:**
- **`ROUTE-M257x-280-map-state-tokens-are-graded-against-nothing`** — the corpus map's per-row STATE
  token has no fence against its own evidence column. `platform_alignment_guard` validates the token is
  in the vocabulary and grades *library* rows via assertion G, but a **service** row may claim
  `live-standalone` while its evidence says destroyed — which is exactly what two rows did. The fence
  is the natural sibling of arm F on the corpus side.
- **`ROUTE-M257x-280-commit-trailer-has-no-check`** — two malformed commit messages in two iters
  (`D-M257x-279-4`, `D-M257x-280-9`). Nothing checks the trailer.
- **`ROUTE-M257x-280-suite-census-is-a-member-of-its-own-population`** — `test_suite_census` writes its
  probe into the directory it censuses, so it is green alone and red in company. 7 of the residual 13.
- **`FIX-M257x-278-census-substrate-tests-hardcode-an-absolute-ROOT` — RAISED to a blocker of triage.**
  Third consecutive iter it has bitten, and this time it stopped the control-tree full run that would
  have settled attribution for a RED suite. It is no longer a nuisance.
- **`ROUTE-M257x-280-the-31-minute-gate-is-skipped-because-it-is-31-minutes`** — the whole-section suite
  costs 31 min on this host, so it gets skipped, so ratchets ship RED (pass 71 did; iter-279 booked the
  same shape). The fix is a fast subset that runs the ratchets and censuses alone — seconds, not half an
  hour — not more exhortation in a banner nobody can afford to obey.
- **`ROUTE-M257x-h70-corpus-and-code-prose-are-copies-with-no-fence`** — advanced on the tooling side
  (arm F fences the disposition copy) but the 172 verbatim `(module, doc)` pairs remain unfenced.
- Unchanged: `ROUTE-M257x-h70-quotation-verification-instrument-is-unreliable`,
  `ROUTE-M257x-279-durations-are-unclassified-measurement-nouns`,
  `ROUTE-M257x-278-thirteen-unpinned-rext-anchors-are-on-undecidable-clocks`,
  `FIX-M257x-278-census-substrate-tests-hardcode-an-absolute-ROOT` (reduced, not closed),
  `ROUTE-M257x-278-rext-tag-SoT-was-six-iters-stale-unnoticed`,
  `ROUTE-M257x-274-successor-half-is-uncovered`, `ROUTE-M257x-274-tie-order-is-unstable`,
  `FIX-M257x-269`, `ROUTE-M257x-270-directus-consumer-cms-key`, `FIX-M257x-266`, `FIX-M257x-265`,
  `ROUTE-M257x-h59`, `ROUTE-M257x-h65`, the fence half of `ROUTE-M257x-277`.
- **Clause 5's semantic reading is still unmeasured** (last: iter-131, `P = 29 / N = 47`, a floor).

**Lessons:**
1. **A route list goes stale exactly like prose** — the named target was already fixed. Re-survey is not
   ceremony; it is what stopped this iter re-doing landed work.
2. **A fence's remedy message can prescribe writing something false.** The arm-A repair path taught
   clearing one fence by asserting a fold that never happened. A remedy is a claim and needs grading too.
3. **When a new fence breaks 14 old tests, suspect the fixtures before weakening the fence.** They were
   modelling a shape production never has, and only a fail-closed arm could reveal it.
4. **Enumerate, then repair.** Pointed at one bad row, the census found a second instance in the corpus
   that four iters of site-repair had missed.
5. **An instrument inside its subject also applies to timekeeping** — a heartbeat composed in the same
   command that reads the clock is written from expectation, twice over.
