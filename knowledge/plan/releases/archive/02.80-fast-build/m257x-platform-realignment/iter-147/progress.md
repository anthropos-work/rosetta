**Type:** tik — under [`TOK-08`](../decisions.md#tok-08-census-the-mechanical-classes-stop-sampling-them--2026-08-07),
per `§9`'s iter-type refinement (iters 135–147 took no reading; the metric is UNMEASURED, not unmoved, so
the 3-no-prog trigger cannot fire — and `TOK-08`'s sealed refutation branch bars an agent-authored
successor in any case).

# iter-147 — every path that CHOOSES a compose profile, censused

## Phase A — the census, with its denominator

The routed target (`SURVEY-M257x-iter146-other-retired-services-unaudited`) asked for iter-146's token
census over the other retired services. **Re-surveyed at open, that population is inert:** `cms` 8090,
`jobsimulation` 8400, `storage` 8300 and `roadrunner` 10400 occur on **74 lines** in
`rosetta-extensions` outside `.git`/`node_modules`/`tests/fixtures` — probe-registry rows, tests about
those rows, port-offset arithmetic, guard prose. **0 emitters.** A registry row is a *consumer*: scoped
out by `STACK_SERVICES`, and unscoped it warns loudly.

So the search was **inverted** (`D-M257x-147-1`), and the inverted population is the one with the defect.

> **Denominator: 7 profile-selecting compose sites across 3 entry points.**

| site | verb | how it obtains its profile | verdict |
|---|---|---|---|
| `demo-stack/up-injected.sh:2156` → `:2179` | up | **derived**, `die` on failure | correct |
| `demo-stack/rosetta-demo` `cmd_up` | up | **empty default** → no `--profile` passed | **DEFECT (live)** |
| `demo-stack/rosetta-demo` `gen` (verb `:506` → `gen()`) | gen | **empty default** → `--profiles ""` | **DEFECT (live)** |
| `demo-stack/rosetta-demo` `cmd_down:357` | down | **derived**, non-fatal fallback (iter-55, F-9) | correct |
| `dev-stack/dev-stack:235` | up | **derived**, `die` on failure (iter-85) | correct |
| `dev-stack/dev-stack:460` | gen | **derived**, `die` on failure (iter-85) | correct |
| `dev-stack/dev-stack:423` | down | project-scoped; the label sweep decides | correct |

**5 of 7 already correct. The 2 that were not are the two nothing exercises** — every documented
invocation (`demo-stack/GUIDE.md:69`, `demo-stack/README.md:35`) passed `--profile "$P"` with the
derivation done *by the reader*, so the default branch was reachable only by a hand-run
`rosetta-demo up N`. `D-M257x-146-2`, one iter old, predicted exactly this.

## Phase B — what the defect actually did

An empty compose profile is **not "the base profile". It is no profile.**

- `docker compose up -d` with no `--profile` selects only services declaring no `profiles:` key —
  on this platform postgresql, redis, sentinel — and **exits 0**.
- `gen_override.py --profiles ""` resolves the same reduced set. Its own help says so:
  *"empty = base/always-on only"* (`stack-core/gen_override.py:329`). So the generated override was
  hollow too, not merely the `up` call.
- And it was **announced three times as a success**: `profile='base'` printed to the operator,
  `"profile":"base"` written into the unified registry (which `/stack-list` reads), and
  `==> demo-N up.` **Compose has no `base` profile** — the word was this script's own name for *no
  profile at all* (`D-M257x-147-4`).

**Repair** (`rosetta-demo`, +34/−4 net of comments): a `derive_profile()` that asks
`stack-injection/platform_topology.py profile` — the same call `up-injected.sh` and `cmd_down` already
make, with **no literal profile token anywhere in the file** — wired into both `cmd_up` (before anything
is allocated or announced) and `gen()`. A derivation failure **refuses**, and the refusal names the
consequence rather than the symptom. The teardown keeps its non-fatal fallback and the asymmetry is
recorded, not left to be "cleaned up" (`D-M257x-147-2`). Both `${profile:-base}` substitutions deleted.

Docs re-pointed the same commit: `demo-stack/README.md` and `GUIDE.md` stopped teaching the by-hand
derivation (a workaround for a defect that is now fixed reads as instruction), and
**`corpus/ops/demo/demo-up-defaults.md`'s `--profile` row gained the default it never had** — the
omission being the corollary in `§5` rule 69.

## Phase C — the fence

`demo-stack/tests/test_profile_derivation_m257x.py`, **8 tests**, behavioural where it can be:

- the shipped `derive_profile` + `gen` are **extracted from the real script with `awk`** (the
  `test_purge.py` pattern — fence the shipped code, never a copy that can drift) and run against a
  stubbed `python3`; the assertion is on the **argv `gen_override.py` actually receives**.
- **RED-proof:** the same assertion against a mutant with the derivation guard deleted — and the
  mutation asserts its target line is present first, so a rename fails loud instead of quietly
  producing a non-mutant.
- **Anti-vacuity:** both structural predicates are proven able to fire *and* not to over-fire, against
  synthetic content (`test_the_two_structural_predicates_CAN_fire`).
- **Twin-drift arm:** both entry points must derive on both their verbs, and the failure **names which
  side drifted** (`§5` rule 68's lesson — `12 != 13` said nothing for four months).

**Control run against the REAL pre-fix text recovered from `HEAD`, not a reconstruction:**

| | pre-fix (`HEAD:demo-stack/rosetta-demo`) | post-fix |
|---|---|---|
| `--profiles` reaching `gen_override.py` | **`''`** | **`'core'`** |
| announcements substituting a placeholder profile | **2** | **0** |

## Phase D — gates

| gate | result |
|---|---|
| new fence `test_profile_derivation_m257x.py` | **8 passed** |
| `demo-stack` section suite | **9 failed · 1,055 passed · 2 skipped** (213 s) — the 9 identical **by name** to iter-145's baseline set; +8 passed = this iter's fence. **0 regressions** |
| `dev-stack` section suite | **151 passed** — iter-145's figure exactly (the fence reads this section) |
| guard family (`--repo-root` rosetta, `--platform` stack-demo/platform @ `0c91421df`) | **19 GREEN · 0 RED · 0 could-not-check · 4 not-run** — identical to iter-144/146's close, so the delta is attributable. `demo_knob_guard` GREEN over the edited defaults row; `corpus_citation_guard` GREEN |
| `bash -n demo-stack/rosetta-demo` | OK |

**Mid-run correction, recorded (`D-M257x-147-3`):** the first `demo-stack` run came back **13 failed**.
Graded per-failure before the count was quoted (`D-M257x-144-2`): **9 pre-existing** (identical names to
iter-145's) and **4 mine** — four `test_tooling.py::RosettaDemoRegistry` tests whose stub platform's
`docker-compose.yml` was created with `open(...).close()`, i.e. **empty**, and which then asserted
`rosetta-demo up` returned **0**. That fixture was the defect encoded as expected behaviour. Repaired the
fixture (a faithful `backend` anchor row) rather than passing `--profile` around the derivation.

`stack-core`, `stack-injection` and `stack-verify` **not re-run, and saying so** (`§5` rule 60): this
iter touches **zero** files in them, which is the condition under which iter-145's/iter-146's runs still
cover the tree — and the corpus-side edits are covered by the guard-family run above.

## Close — 2026-08-08

**Outcome:** the profile-selection class censused — **7 sites, 5 correct, 2 live defects**, both in
`rosetta-demo`'s bring-up and gen verbs, where an omitted `--profile` became **no profile**: compose
exits 0, only the always-on floor starts, and the run announced `profile='base'` — a token compose does
not have — to the operator *and* to the stack registry. Repaired to derive-or-refuse, both announcements
de-euphemised, docs and the defaults contract re-pointed, and fenced over **both** entry points with a
RED-proof taken against the real pre-fix text. The finding that generalises: **a token census cannot see
an ABSENT value**, and **iter-85's own comment named this twin lag and left it for 62 iters**.
`§5` gains **rule 69**.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET — **4 of 5**, unchanged. **No `N` reading was taken, so no `N` movement is claimed**
(`§9`'s guard-rail 1, in its required words).
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (**iters 135–147 took no reading, so the metric is UNMEASURED not unmoved — `§9`'s iter-type refinement; and `TOK-08`'s sealed refutation branch bars an agent-authored successor in any case**) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (**1 tik this session**) — (6) protocol-stop: n — (7) budget-exhausted: n — Outcome: **continue**
**Decisions:** `D-M257x-147-1` (a token census finds a WRONG value and never an ABSENT one — invert the
search) · `D-M257x-147-2` (the bring-up refuses where the teardown falls through; the asymmetry is
deliberate and recorded so it is not "cleaned up") · `D-M257x-147-3` (the RED fixture was repaired, not
bypassed — an empty stub compose file was the defect written down as expected behaviour) ·
`D-M257x-147-4` (`base` is not a compose profile; the invented word is removed from the operator line
*and* from the registry).
**Side-deliverables:** none.
**Routes carried forward:**
- **`SURVEY-M257x-iter147-absent-value-class` (NEW)** — the inverted census run here covered exactly one
  choice-point (the compose profile). The same *absent-value* question is unasked for the tooling's other
  choice-points: `--services`, `--ref`, `--data-root`, and the `STACK_*` scope variables. Each has the
  same shape — an omitted value that compose or a downstream tool accepts silently. **Grade a population
  first** (iter-144), and do not widen the new fence un-audited.
- `SURVEY-M257x-iter146-other-retired-services-unaudited` — **answered with a measurement, not executed**:
  74 lines / 0 emitters for the four ported retired services. The narrower live question it was really
  reaching for is the new route above. The token arm stays open only for `skiller`/`skillpath`/`chronos`/
  `intelligence`, which carry no port at all.
- `FIX-M257x-iter145-sha-baseline-drift` · `-iter145-migrate-race-needs-a-host-postgres` ·
  `-iter145-green-but-stale-graphql-mentions` (the 9 standing `demo-stack` failures are these two
  classes; re-attested by name this iter rather than carried as "unchanged, not re-verified").
- `SURVEY-M257x-iter144-orphan-arm-is-the-residual` · `FIX-M257x-iter144-correction-vs-retraction-unfenced` ·
  `FIX-M257x-iter143-wrong-head-is-unfenced` · `-iter143-scope-derivation-by-grep` ·
  `-iter143-appending-to-the-protocol-doc-rots-the-ledger` (⚠️ **widened again by rule 69 — fifth
  consecutive iter to grow it and say so**) · `FIX-M257x-iter142-value-change-articles` ·
  `-iter142-path-arm-window` · `-iter142-tier-b-underflag` · `FIX-M257x-iter135-adjudicated-live-defects` ·
  `-iter140-receipts-not-checkable-here` · `-iter140-receipt-fence` · `-iter138-anchor-rot-fence` ·
  `-iter134-fence-family-has-no-shared-predicate-layer` · `-iter133-two-fives-need-a-fence` ·
  `-iter132-suite-walltime-is-not-a-measurement` · `-iter131-predicate-sets-not-enumerated`.
**Lessons:**
0. **Grep finds wrong values. It cannot find missing ones.** The whole emitter-census method inherited
   from iter-146 is blind to the class of defect where the bug *is* the empty string. Inverting the
   search — enumerate the choice-points, grade each — costs the same and reaches both.
1. **When a repair's own comment names a sibling that lagged, the sibling set is already enumerated.**
   iter-85 wrote *"the dev path kept the literal for four more releases"* and did not fence the demo
   path's remaining verbs. Write the fence over the SET at the moment the twin is first noticed; the
   observation is free and the next 62 iters are not.
2. **A missing row in a defaults contract is a claim.** The `--profile` row simply had no default, and
   that read as "nothing to say" rather than "the default is none" — inside the one document written to
   stop precisely this.
3. **A fixture that cannot reach the code path under test converts a fence into decoration.** Four tests
   asserted a successful bring-up from an EMPTY platform compose file. The RED was the fence working.
