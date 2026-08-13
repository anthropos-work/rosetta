# iter-280 — decisions

## `D-M257x-280-1` — a disposition token is a claim about the platform, so grade it against the platform

`SERVICES_NOT_IN_PLATFORM_COMPOSE` carries, per row, a token saying WHY that row is absent from the
platform compose. Arms A–D grade the row's NAME. The token was split out at parse time
(`service_registry_guard.py`, `name, why = _split_decl(e, 2)`), stored, and never read again.

So the array's MEMBERSHIP was fenced in both directions while its FACTS were fenced not at all — which is
how `roadrunner` sat there declared `merged-into-app` for four months after M257x iter-137 measured the
opposite.

**Arm F grades each token against the `app` clone**, because that is what the tokens are claims about:

| token | predicate |
|---|---|
| `merged-into-app` | `app/internal/<pkg>/` must EXIST |
| `deleted-not-merged` | it must NOT exist |
| `rext-injected` | makes no claim about `app`; arm D already holds it |

Measured at platform `0c91421` / `app` in `stack-demo`: **4 tokens graded, all true.** The 5th
(`directus:rext-injected`) asserts nothing about `app` by construction. The iter's deliverable is
therefore the fence, not a repair — TOK-08's shape exactly.

**The clone is DERIVED from the platform clone (its sibling), never a new flag.** This family states that
rule three times over its own guards: *"its reference is DERIVED from `--platform` rather than taken as
its own flag, so the three cannot end up pointed at different clones."* An `APP_CLONE` override exists
for a non-`make init` layout and is not the normal path.

## `D-M257x-280-2` — the arm fails CLOSED, and that is proven with the precondition absent

A capability probe that fails open disarms the check it guards. Arm F returns **CANNOT-MEASURE (exit 2)**
when the clone is unreachable — it does **not** skip, and the guard does **not** print ALIGNED.

Proven, not asserted: `APP_CLONE=/nonexistent/app` → `rc=2`, message naming arm F,
`ALIGNED` absent from the output. Pinned as `test_a_missing_app_clone_is_CANNOT_MEASURE_never_aligned`.

## `D-M257x-280-3` — the fail-closed arm broke 14 arm A–D controls, and the FIXTURE was the defect

Landing arm F turned 14 pre-existing tests RED at once. They build a synthetic compose in a tmpdir; the
derived sibling walked out of the fixture entirely, so every one of them became CANNOT-MEASURE.

The tempting reading is *"arm F is too strict, make it skip."* That is the fail-open trade, refused.
**The fixture was modelling a directory shape production never has** — compose at a bare tmp root, with
no clone set around it. It now mirrors a real stack (`<tmp>/platform/docker-compose.yml` beside
`<tmp>/app/`), and carries an `app/internal/jobsimulation/` because its own registry declares
`jobsimulation:merged-into-app` — a synthetic clone has to make the synthetic claim TRUE.

**This is the session's carried class firing again:** an instrument that lives inside its own subject.
The controls for arms A–D were only ever green because nothing downstream of them depended on the
fixture being shaped like the thing it modelled. Caught only by running the whole suite against the tree
about to be committed — iter-279's lesson 4, one iter later.

## `D-M257x-280-4` — a control was prescribing a repair that asserts a false fact

`test_declaring_a_departure_is_how_you_go_green_again` proves the arm-A remedy works: declare the
departed service and the fence clears. It declared the departed `gotenberg` as **`merged-into-app`** —
and arm F broke it, **correctly**. Gotenberg is a third-party Office-to-PDF image; it was never folded
into `app`, and `app/internal/gotenberg` does not exist.

So the control was demonstrating that you clear arm A by **writing down something false** — the exact
trade the `roadrunner` row made for four months, encoded in a test. Repaired to `deleted-not-merged`,
which is both true and green. The remedy still works; it has to name the RIGHT disposition, and that is
arm F's whole point.

## `D-M257x-280-5` — "Both directions" was implemented in one direction

Harden pass 71's `test_every_reason_is_one_of_the_documented_kinds` opens *"Both directions."* It
computes `undocumented` and asserts **nothing** about the opposite direction. It was vacuous-but-true
(all three kinds were in use), which is why it read as correct.

This is the milestone's own recurring shape — **a correct fix carrying a false stated reason** — sitting
inside the test written to catch that shape elsewhere. **Implemented rather than retracted**, and
against the guard's real vocabulary rather than more prose: `test_the_guards_vocabulary_and_the_prose_agree`
asserts every token in `DISPOSITION_PREDICATE` is defined in the `services.sh` comment block. Someone
choosing a token reads the prose, not the guard.

## `D-M257x-280-6` — side discovery: a fence citation that named a fence nobody built

`services.sh` told the reader that an undeclared stack-injected service *"turns
`stack-core/service_registry_guard.py` **arm E** RED."*

**There is no arm E.** Measured over full history, not just at HEAD:

    git log --all -S "STACK_INJECTED_SERVICES_NOT_PROBED" -- stack-core/service_registry_guard.py  → 0 commits
    git log --all -S "arm E"                              -- stack-core/service_registry_guard.py  → 0 commits

The guard has never contained a single reference to that array. **The check is real** — it lives in
`stack-verify/tests/test_scope_union_m257x.py`, whose both-directions arm reads the array as the declared
side and a real generator emission as the measured side.

**Both halves landed in the SAME commit**, `83ada03` (iter-153, 2026-08-08): that commit wrote the array,
wrote the real grader as a test, and wrote a comment citing a guard arm it did not build. **The citation
described the plan, not the build, and the two diverged inside one commit.** Nothing caught it for 127
iters.

A wrong pointer is worse than no pointer: it sends whoever is debugging a RED to a file where the absence
of the arm reads as a *missing fence*, and invites re-implementing one that already exists. Recorded as a
side-deliverable — it does not upgrade this iter's close status.

## `D-M257x-280-7` — the comment this iter falsified, fixed in the same commit as the code

`services.sh` said *"Nothing grades this token — `service_registry_guard` reads the NAME and stores the
reason as free prose."* True when pass 71 wrote it; **false the moment arm F landed.** Shipping arm F
without rewriting that block would have committed the very defect the arm exists to fence, in the file it
fences. The block now states the predicate per token and that a false one turns the guard RED.

## `D-M257x-280-8` — two clock-discipline violations, and the second exposed the CAUSE

A journal heartbeat was stamped `02:57` when the clock read `02:49`. Corrected by an appended line (the
journal is append-only, so the wrong line stands with its correction beneath it) — then **it happened
again**: `02:56` written when the clock read `02:53`.

Twice is not carelessness, it is a method defect. **Both heartbeats were written in the SAME shell
command that read the clock** — `date …` and the `cat >> journal` heredoc in one invocation. The heredoc
body is composed before the command runs, so the timestamp is necessarily written from expectation; the
`date` output arrives too late to inform the line it was supposed to supply. The instrument was inside
its own subject, which is this session's carried class, arriving in the timekeeping.

Method changed: **read the clock in its own call, then write the heartbeat with the value just seen.** No
heartbeat after 02:54 was composed in the same command as its `date`.

Recorded rather than hidden because a timestamp nobody can trust is worse than no timestamp, and this
milestone's whole method rests on measurements being read rather than assumed.

## `D-M257x-280-9` — a malformed commit trailer, recorded and NOT force-fixed

The rext commit `38c0aba` carries `Co-Authored-By: Claude Opus 5 (1M context) <noreply@anthropos.com>`.
The correct domain is `anthropic.com`; `anthropos.com` is the platform's org name, typed by reflex in a
repo where that word appears on nearly every line.

Recorded rather than repaired: rewriting a published commit message is a history rewrite, and this
milestone's standing rule is that a malformed message is recorded, never force-fixed. iter-279 booked the
same class one iter ago (`D-M257x-279-4`, a message mangled by the shell). Two occurrences in two iters
is a pattern in the commit-message path, not two accidents — routed forward as
`ROUTE-M257x-280-commit-trailer-has-no-check`.

## `D-M257x-280-10` — the whole-section suite came back RED, and triaging it against a CONTROL TREE was the whole point

The first full run against the final tree returned **22 failed / 2219 passed** (31m15s). A green was
claimed nowhere; the run was triaged before anything was committed.

**The triage instrument was a frozen clone**, not judgement: `rext` cloned at `9933b6a` (the pre-iter
commit) and the failing files run there. That separates *mine* from *pre-existing* mechanically, which is
the only way to answer the question honestly — and this milestone has booked the frozen-clone technique
before (iters 249/250).

| failing file | control tree | verdict |
|---|---|---|
| `test_service_doc_status_fence` | GREEN | **mine** |
| `test_frozen_expectation_census_m257x` (docstring ceiling) | — | **mine**, +1 literal |
| `test_frozen_expectation_census_m257x` (test-module ceiling) | **already RED, 655 > 653** | **NOT mine — harden pass 71** |
| `test_suite_census` ×6, `test_suite_census_population` | GREEN standalone on BOTH trees | full-run interaction |
| `test_repair_leak_guard_mutation_battery` ×4, `test_battery_stage` | GREEN standalone (`RC=0`) | full-run interaction |

**Three distinct causes wearing one colour.** A single RED count would have been read as "iter-280 broke
22 tests", and two thirds of that would have been wrong.

### The two that were mine, and both were the fence being RIGHT

1. **`graphql-wundergraph.md` carried no status banner the fence accepts.** Flipping its prod cell to
   `decommissioned` put it in the gone-set, and the fence requires one line carrying **both** `⚠` and a
   disposition word within the first 20 lines. Its banner headline said *"GONE FROM LOCAL DEV"* — which
   is not one of them, and was also no longer the whole truth. Rewritten to lead with **DECOMMISSIONED —
   destroyed in production, archived as a repo, gone from local dev.** A corpus correction propagating
   into a second document is the fence working, not a cost.
2. **A dated measurement literal in the arm F docstring** (`→ 0 commits, ever`) breached
   `DOCSTRING_LITERAL_CEILING` 241 > 240. Removed by rephrasing (*"returns no adds at all"*), not by
   bumping the ceiling: **240 again.**

### The one that was NOT mine, fixed anyway because it is the same phrase

`TEST_MODULE_LITERAL_CEILING` was **already breached at 9933b6a** — pass 70 sat exactly at 653, pass 71
took it to 655 and **shipped the section suite RED without noticing.** Measured precisely, not guessed:
the two added literals are `test_service_registry_guard.py:467` and `:529`, both the string
`0 commits` — **the identical phrase this iter had just removed from its own docstring for the identical
reason**, in the file this iter already owns.

So it was repaired the same way rather than re-pinned: rephrased to *"returns no adds at all"*, taking
the census to **653, at the ceiling, with no bump.** The fence's own instruction — *"re-pin WITH A
RECORDED REASON naming these, never a blind bump"* — is satisfied more cleanly by deleting the growth
than by ratifying it.

**Attribution stated plainly: this breach was pass 71's, not iter-280's**, and it is repaired here only
because the offending lines sit in this iter's file and are this iter's own subject.

## `D-M257x-280-11` — a harden pass shipped the suite RED, one pass after telling everyone to run it

`guard_family` prints, every run: *"a green above is a statement about guard verdicts alone… Run the
tests before closing work that touched a guard, a fixture, or a cited corpus line."* Pass 71 touched a
guard's tests and left a ratchet RED.

That is not a scolding, it is the measurement this milestone keeps producing: **the whole-section suite
is 31 minutes on this host, and a gate nobody can afford to run is a gate that drifts.** iter-279 booked
the same shape (*"a green from before the last edit grades a tree that no longer exists"*). Two
consecutive units of work have now been caught by the same 31-minute gap. Routed forward as
`ROUTE-M257x-280-the-31-minute-gate-is-skipped-because-it-is-31-minutes` — the fix is a fast subset that
runs the ratchets and censuses alone, not exhortation.

## `D-M257x-280-12` — the residual 13, and the measurement that was BLOCKED rather than skipped

After the three repairs the section suite reads **13 failed / 2224 passed** (31m39s), down from 22/2219.
The residual sits in exactly four files:

    test_suite_census (7) · test_repair_leak_guard_mutation_battery (4)
    test_suite_census_population (1) · test_battery_stage (1)

**All four pass standalone on BOTH the control tree and the final tree.** They fail only inside the full
run.

**Root cause, for the largest group:** `test_suite_census` writes its probe module
`stack-core/tests/test_unit_probe.py` **into the very directory it is censusing**, then asserts the
runner counted 3 tests. Inside the full run it reads `tests: 1` with `FAILED (errors=1)`. **The census
is a member of the population it measures** — the session's carried class in its purest form, and the
reason the file is green alone and red in company.

**The attribution I wanted, I could not take — and the reason is a route already open.** A full-suite run
on the frozen control clone would have settled *pre-existing vs mine* outright. It **cannot be run**:
collection aborts immediately with `ERROR collecting tests/test_anchor_subject_census_m257x.py`, because
the census substrate resolves an absolute ROOT and does not survive being cloned. That is
`FIX-M257x-278-census-substrate-tests-hardcode-an-absolute-ROOT`, biting for the **third** consecutive
iter (278, 279, now 280) — and this occurrence is the one that shows it is not cosmetic: **it blocks the
exact measurement needed to triage a RED suite.**

So the honest statement, and no stronger one: the residual 13 are **not attributable to iter-280** on the
standalone evidence and on the 22→13 direction of travel, and they are **not proven pre-existing**
either, because the instrument for that proof is broken. Both halves are stated; neither is inflated
into the other.

Routed forward as `ROUTE-M257x-280-suite-census-is-a-member-of-its-own-population`, and it raises
`FIX-M257x-278` from a nuisance to a blocker of triage.
