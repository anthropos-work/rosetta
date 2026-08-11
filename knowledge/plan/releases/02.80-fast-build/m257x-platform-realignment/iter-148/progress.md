**Type:** tik — under [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07).

# iter-148 — the absent value, one section over

## Phase A — the census, with its denominator

iter-147's route asked the *absent-value* question at the tooling's other choice-points.
`stack-verify/lib/services.sh:33` **discloses the answer in its own comment**: an unset `STACK_SERVICES`
makes verify *"probe everything in the table and false-`down` the merged-away rows."*

> **Denominator: 3 sites that invoke `live/verify.sh`.**

| site | scope | verdict |
|---|---|---|
| `stack-verify/live/autoverify.sh:204` | `--services` from the caller; `up-injected.sh:2681` and `dev-stack:347` both DERIVE it via `platform_topology.py services` and both carry a fallback | correct |
| `stack-verify/reports/generate.sh:63` — the **`/test-platform` driver** | **nothing at all** | **DEFECT (live)** |
| `stack-verify/live/autoverify.sh:699` — the operator hint string | unscoped by construction | follows the fix |

`rosetta-demo`'s remaining choice-points were graded in the same pass and are sound: `--ref`/`--only`
use `${var:+--flag "$var"}` (omitted when empty), and an empty `--services` genuinely means *every
service in the profile* — the distinction iter-147's fence encodes.

## Phase B — measured, both arms, against a live stack

Read-only probes against `demo-1` (11 containers up) — the same GETs `autoverify` runs at every bring-up.

| arm | invocation | result |
|---|---|---|
| **A — unscoped** | how `generate.sh` invoked it | **✗ 6 of 20 probes failed** |
| **B — scoped** | how `autoverify.sh` invokes it | **✗ 1 of 14** |

The five-probe delta is **entirely the merged-away rows**: `jobsimulation`, `cms`, `storage`,
`roadrunner` liveness — all `HTTP 000000`, no listener, and there never will be one — plus the
`storage-rpc` readiness probe on the same cause.

**So `/test-platform` — the tool whose entire purpose is *"what is actually working"* — reported four
services the platform deleted as DOWN, and exited 1.** The skill's own documented invocation
(`SKILL.md:76-78`) sets `STACK_ROOT` and `REPORT_DIR` and nothing else.

**The 1 remaining failure in arm B is NOT booked as a defect** (`D-M257x-148-4`): `postgres-schemas`
fails in *both* arms because its `repos.yml` candidate path resolves in a per-stack consumption copy and
not in the `.agentspace` authoring copy this iter measured from (verified present at
`stack-demo/rosetta-extensions/…/platform/repos.yml`). Substrate, not defect — `D-M257x-122-4`.

## Phase C — repair

`generate.sh` derives `STACK_SERVICES` from `$STACK_ROOT/platform` when unset, and **discloses** when it
cannot — it does not refuse (`D-M257x-148-1`: the disposition follows what the artifact is *for*). The
scope note rides **in the markdown report**, not only on stderr. A caller-supplied `STACK_SERVICES`
always wins. Both branches exercised end-to-end:

- derivable → `→ probe scope DERIVED from …/stack-demo/platform/docker-compose.yml: postgresql redis
  sentinel backend gotenberg`, and the note appears at line 15 of the generated report.
- underivable → the `⚠ probe scope is UNSCOPED …` disclosure, naming the four rows and why they read
  `down`.

`.claude/skills/test-platform/SKILL.md` gained the scope note **and** the `STACK_PROJECT`/`STACK_OFFSET`
requirement for a demo or `dev-N` target — without them the probes go to project `anthropos` on base
ports, which is a different stack than the one the operator meant.

## Phase D — the fence, and what it caught

`stack-verify/tests/test_probe_scope_m257x.py`, **6 tests**. The load-bearing arm is
**derived-vs-declared in both directions** (`D-M257x-148-2`): the registry rows the platform's compose no
longer declares, minus the rows this tooling itself injects, must **equal** the set the disclosure names
— so the warning goes RED at the next fold instead of quietly reassuring a reader about the wrong four
services.

⚠️ **The RED-proof caught a real defect in this fence** (`D-M257x-148-3`). The first draft's
"derivation precedes invocation" check used bare `src.find()` for both tokens, and **both bound to
comments** — `generate.sh`'s usage header names `verify.sh`, and the comment block explaining this very
fix names `platform_topology.py`. The check would have kept passing after the code it guards was
deleted. It was the **mutation control that found it, not review**: the de-scoped mutant retained the
comment and the assertion held. Now both bind to **executable content only**
(`_first_executable_index`). `§5` rule 67 / 68(d)'s axis, reproduced inside the fence written to apply
it — twice in one iter.

| gate | result |
|---|---|
| new fence | **6 passed**, RED-proof load-bearing (a de-scoped `generate.sh` fails all three checks) |
| `stack-verify` section suite | **244 passed · 0 failed** (238 at iter-145's close + 6 new) |
| guard family (`--repo-root` rosetta, `--platform` stack-demo/platform) | **19 GREEN · 0 RED · 0 could-not-check · 4 not-run** — identical to iter-147's close |
| `bash -n stack-verify/reports/generate.sh` | OK |

`stack-core`, `demo-stack`, `dev-stack` and `stack-injection` **not re-run, and saying so** (`§5` rule
60): this iter touches zero files in them — and `demo-stack`/`dev-stack` were run in full one iter ago.

## Close — 2026-08-08

**Outcome:** the verify entry points censused — **3 sites, 2 correct, 1 unscoped** — and the unscoped one
is the `/test-platform` driver. Measured on a live stack: **unscoped 6 of 20 probes failed vs scoped 1 of
14**, the delta being four services the platform merged into `app` plus one readiness probe on the same
cause, so the report said the platform was broken and exited 1. Repaired to derive-or-disclose, the scope
note carried into the report itself, the skill's invocation documented, and fenced derived-vs-declared so
the warning cannot rot. **The RED-proof caught a comment-bound predicate in this iter's own fence.**
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — **4 of 5**, unchanged. **No `N` reading was taken, so no `N` movement is claimed**
(`§9`'s guard-rail 1, in its required words).
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (**iters 135–148 took no reading, so the metric is UNMEASURED not unmoved — `§9`'s iter-type refinement; `TOK-08`'s sealed refutation branch bars an agent-authored successor in any case**) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (**2 tiks this session**) — (6) protocol-stop: n — (7) budget-exhausted: **y** — Outcome: **exit-7**
**Decisions:** `D-M257x-148-1` (a report DISCLOSES where a bring-up REFUSES — the disposition follows
what the artifact is for) · `D-M257x-148-2` (the disclosure is fenced derived-vs-declared; a warning
naming the wrong services is worse than none) · `D-M257x-148-3` (the RED-proof caught a comment-bound
predicate in this fence — a form-matching fence must bind to executable content by construction) ·
`D-M257x-148-4` (`postgres-schemas` is a substrate artifact of the authoring copy, not a defect).
**Side-deliverables:** none.
**Routes carried forward:**
- **`SURVEY-M257x-iter148-registry-is-hand-maintained` (NEW)** — the deeper defect is that
  `stack-verify/lib/services.sh` is a **hand-maintained** service table with a base-port literal per row,
  which is §2's tuple defect in the one section that grades every stack. iter-145 fenced its *test-side
  copy* against it; nothing derives the table itself from the platform's compose. **Grade the cost first**
  — iter-145's `REGISTRY_BASES` is a deliberate anti-vacuity control and a naive derivation would delete
  it.
- `SURVEY-M257x-iter147-absent-value-class` — **partially closed**: `STACK_SERVICES` done here,
  `rosetta-demo`'s `--ref`/`--only`/`--services` graded sound. Still open for `STACK_PROJECT` /
  `STACK_OFFSET`, whose unset defaults silently target the main dev stack.
- `SURVEY-M257x-iter146-other-retired-services-unaudited` (token arm, `skiller`/`skillpath`/`chronos`/
  `intelligence` only) · `FIX-M257x-iter145-sha-baseline-drift` ·
  `-iter145-migrate-race-needs-a-host-postgres` · `-iter145-green-but-stale-graphql-mentions` ·
  `SURVEY-M257x-iter144-orphan-arm-is-the-residual` · `FIX-M257x-iter144-correction-vs-retraction-unfenced` ·
  `FIX-M257x-iter143-wrong-head-is-unfenced` · `-iter143-scope-derivation-by-grep` ·
  `-iter143-appending-to-the-protocol-doc-rots-the-ledger` · `FIX-M257x-iter142-value-change-articles` ·
  `-iter142-path-arm-window` · `-iter142-tier-b-underflag` · `FIX-M257x-iter135-adjudicated-live-defects` ·
  `-iter140-receipts-not-checkable-here` · `-iter140-receipt-fence` · `-iter138-anchor-rot-fence` ·
  `-iter134-fence-family-has-no-shared-predicate-layer` · `-iter133-two-fives-need-a-fence` ·
  `-iter132-suite-walltime-is-not-a-measurement` · `-iter131-predicate-sets-not-enumerated`.
**Lessons:**
0. **A file that discloses a hazard in a comment is a file nobody fixed.** `services.sh:33` described
   this defect precisely, and had for as long as the rows existed. iter-146 found the same shape
   (`gen_tailscale_serve.py` naming the hazard `dev-stack` then committed). **Treat a self-disclosing
   comment as an open finding with a written repro, not as documentation.**
1. **A mutation control is not a formality — it is the only thing that catches a fence bound to its own
   comment.** Two form-matching predicates in this iter's fence looked right, passed against the real
   file, and would have passed against a file with the guarded code deleted.
2. **State the substrate before booking a failure.** One of the two arms' residual failures is an
   artifact of measuring from the authoring copy rather than a per-stack copy. Read the path in the
   error message before believing the verdict.
