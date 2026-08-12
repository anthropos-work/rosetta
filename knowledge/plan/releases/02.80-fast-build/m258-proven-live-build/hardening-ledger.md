# Hardening Ledger — M258 proven-live build

> M258 ran **20 iters and was never hardened**, so this file opens at the final pass and its scope is
> cumulative over the whole milestone.
>
> ⚠️ **THIS IS NOT A GATE-MET CLOSE.** Clauses 1, 2, 4 and 5 are proven (clause 5 and the batch
> re-proven on the final stack). **Clause 3 — composed p50 ≤ 480 s over 3 cold cycles — is NOT MET and
> must never be recorded as met.** The user ruled the goal achieved on the other clauses plus a ~402 s
> clean projection, having concluded the machine's CPU contention is not removable. The shape is
> *achieved by user ruling, timing clause unmeasured under load* — the M257x `TOK-09` shape. The 840 s
> contended figure stays **instrument-rejected**, and the last iter's refusal to bank a flattering
> ~290 s cycle stands: it was **warm-cache on a quiet box**, missing the export/unpack leg that is
> 46.2 % of a cold one. That refusal is the honest reading and is preserved here deliberately.

## Pass 1 — 2026-08-12 — final

**Iters hardened this pass:** all milestone-touched code (cumulative — iter-01 … iter-20).

**Tiks covered since prior pass:** all iters in milestone (first harden pass).

**Scope measured:** 31 `rosetta-extensions` files across 10 commits (`679a5f7..HEAD`) + 35 non-plan
`rosetta` files. Cross-iter hotspots identified by mapping file → touching commits:
`demo-stack/up-injected.sh` (4 iters), `stack-core/buildbench.py` (3), `playthroughs/e2e/restore-presenter-world.sh` (2),
`stack-core/tests/test_buildbench.py` (2). The final-mode integration work concentrated there.

**Coverage delta on touched files:** statement coverage is **not instrumented** for this surface and was
not fabricated — the milestone-touched code is shell (`up-injected.sh`, `dev-stack`, `rosetta-demo`,
`restore-presenter-world.sh`), fence Python, and Go *test* files that read those scripts as data. The
meaningful measure here is **mutation coverage of the fences**, which went from *unmeasured* to **18
mutants, each proven rejected**, plus 2 false-RED controls. Test counts on touched modules:
`playthroughs/manifest` +9 (3 fences + 6 mutants), `demo-stack` +9 (new module), `dev-stack` 6 → 11,
`stack-verify` 275 (1 failing → 0).

**Tests added:**
- iter-16/17 (sentinel reload) → `playthroughs/manifest/runner_safety_test.go`: re-armed fences (5)/(6)
  onto a pure predicate; +`TestRunnerSafety_RestoreLegPolicyInvalidationGate`;
  +`TestPolicyInvalidationFence_FailsRed` (6 mutants).
- iter-16/17 (bring-up reload site) → `demo-stack/tests/test_policy_invalidation_m258.py` (**new**):
  3 fences + 6 mutants, incl. rung-order and pre-v11.0-rung-retention.
- iter-14/17 (`down -v`) → `dev-stack/tests/test_dev_teardown_sweep_m258.py`:
  +`TestTheSafetyDerivationIsMEASURED` (2), +`TestTheDerivationFenceFiresRed` (5, incl. 2 false-RED controls).
- iter-06 (batch restore) → `playthroughs/manifest/batch_gate_test.go`:
  +`TestRestore_ResolvesItsBinaryFromTheLIVEStackToo` (behavioural, faked docker + stub binaries),
  +`TestRestore_ResolvesTheLiveStackExactlyOnce`.

**Bugs surfaced + fixed inline:**
- **The sentinel-reload fence was satisfied by its own comment** (`dc31efc`). Fences (5)/(6) were
  `strings.Contains` over the raw script while the channel name also appears in the explanatory comment
  above the call. Proven by mutation: corrupting the **executed** channel to one nothing subscribes to,
  comment intact, kept the gate **GREEN** — re-introducing, twelve lines below the note recording its
  earlier repair, the very defect class M257x harden pass 1 fixed for asserts (1)/(1b). Two of the three
  reload sites were unfenced entirely, including the **restore leg**, whose miss costs the *user* a
  working stack rather than a test batch.
- **The `-v` safety argument was fenced as prose, not as fact** (`b29fc1a`). The derivation was
  re-verified independently and holds at both refs on this box (`0c91421`, `766df6c`; and it survives
  `--local-content` — the per-stack Directus declares no volumes). But nothing *measured* it: the day
  the platform adds a named volume, `down -v` becomes destructive on both paths and every fence stays
  green. The argument also omitted the **injected override layer**, which the dev path actually runs.
- **The restore fixed *where* it writes and not *what* writes it** (`e429793`). iter-06 moved the
  manifest destination onto `resolve_stack_dir` and left the stackseed **binary** on `$EXT_ROOT`, so a
  run from the authoring copy exported the live stack's menu with a stale clone's binary — and on a
  clone that never brought demo-N up, refused with *"bring demo-N up first"* about a stack that is up.
  Reachable through the gate's **own** printed recovery command, which omits the `--stackseed` flag
  batch-gate itself always passes.
- **The milestone found sentinel-in-app and told only the reload sites** (`3b2ad38`). Pointed at the
  platform demo-4 was built from, `service_registry_guard` reported `A/DEPARTURE` for `sentinel`.
  iter-16 *discovered* the v11.0 fold and repaired three reload sites without asking what else named
  sentinel — repaired where found, not everywhere it exists. A coupled fence then caught the half still
  missing: the unscoped-run disclosure named four services when the answer had become five.

**Knowledge backfill:**
- `demo-stack/frontend/studio-desk.Dockerfile` (`9bf2697`) — the header pre-registered its prediction
  (`D62`) and the **outcome that refuted it lived only in a progress file**. Now recorded beside it:
  space **1.7 → 1.35 GB, 350 MB (20.6 %)**, roughly **one third** of the revised prediction, because the
  image is dominated by **production** dependencies (838 MB of 1.04 GB survives `--omit=dev`); and time
  **UNMEASURED** (`D75`) — the one cold attempt's ~1.5 s export was a BuildKit cache hit. The sentence
  most likely to mislead was *"both are measured by buildbench, never asserted"*: one axis was measured,
  the other never was, and the file said the same thing about both.

**Flakes stabilized:** none surfaced. No new test is timing- or network-dependent by construction — the
behavioural restore test fakes `docker` and stubs the binary, so it neither reaches a daemon nor a clone.

**Verification:** `playthroughs` 4/4 Go packages green · `dev-stack` 166 green · `demo-stack` 1109 run,
9 failing · `stack-verify` 275 green · `stack-core` coupled modules 243 green (service-registry 38,
bringup-verify-scope 15, platform-predicate 190) · **corpus fence set 15/15 `rc=0`** against
`stack-demo/platform` @ `766df6c`.

**Pre-existing failure attribution (re-derived, not inherited):** the 9 `demo-stack` failures were
attributed against a `git archive 5566538` pristine extract. **The first extract was invalid and said
so out loud** — sited at the wrong depth, its clone discovery missed `stack-demo/next-web-app` and six
of the tests **SKIPPED** rather than ran, each printing *"SKIPPED, NOT PASSED … Nothing else verifies
it"*. Re-sited so the discovery resolves, pristine reports **10** failures to HEAD's 9: **nothing is in
current-but-not-pristine, so zero failures were introduced.** All three live clones verified git-clean
afterwards — the demopatch G5 self-revert held even for the tests that failed.

**Fence-invocation finding (recorded, not fixed):** the 15 guards have **five different invocation
contracts** (`--repo-root`, positional root, `$ROSETTA_ROOT`, `argv[1]=map argv[2]=repos.yml`,
`argv[1]=services.sh argv[2]=compose`). Three deliberately **fail closed with `rc=2`** on a missing
platform reference — *"there is deliberately no default — a fidelity check against the wrong reference
passes"* — which is excellent design and **indistinguishable from a RED to any caller grading `rc==0`**.
A first naive sweep this pass read **5 green / 10 red on a completely healthy corpus**. Anyone quoting a
fence count must state the invocation and the platform ref with it.

**Stop condition:** continue-to-next-pass — four inline fixes and one backfill landed; the dimension scan
found new defects in every one of the five scoped areas, so coverage has not stabilised. Pass 2 to sweep
the cross-iter hotspots not yet reached (`up-injected.sh`'s other three iters, `buildbench.py`'s three)
and the corpus-sweep class.

## Pass 2 — 2026-08-12 — final

**Iters hardened this pass:** the cross-iter hotspots pass 1 identified but did not reach —
`stack-core/buildbench.py` (iter-03, iter-06, iter-09), `stack-injection/inject.py` (iter-03), and the
`demo-stack/up-injected.sh` batch-gate hook (iter-06).

**Tiks covered since prior pass:** n/a (same final-mode session, cumulative scope).

**Coverage delta on touched files:** `stack-core` isolation 47 → 53 · `stack-injection` 341 → 342 ·
`demo-stack` 1109 → 1119. Mutation coverage +7 (the batch-gate hook, previously **zero** tests of any
kind). Two surfaces were examined and found **already sufficient**, which is also a result:
`parse_setdress_attribution` (iter-09) already tests both failure modes its own docstring names —
out-of-span cuts in both directions and segment overlap — and `write_injection_block`'s three existing
tests cover non-accumulation, foreign-content survival and byte-idempotence.

**Tests added:**
- `stack-core/tests/test_isolation_assert_m257.py` +6 (3 unparseable-last-value shapes, the
  fallback-cannot-undo-the-refusal case, and 2 narrowing controls).
- `stack-injection/tests/test_injection.py` +1 (collapse from the measured 24-block state).
- `demo-stack/tests/test_batch_gate_hook_m258.py` (**new**) 3 contract + 7 mutants.

**Bugs surfaced + fixed inline:**
- **Last-wins won the assignment but not the decision** (`7d58990`). iter-03 made a dotenv read
  last-wins; the loop still only overwrote `found` when the value *parsed*, so a later assignment with
  an unparseable value left the earlier key standing. Measured on the shipped function, three shapes did
  exactly that and all three returned the **foreign** key: a later blank (`KEY=`), a trailing inline
  comment (compose takes it literally), and a placeholder written over the real value. That is
  first-wins returning by another door, inside the one function written to end it, with the same
  consequence — the false RED that cost iter-02 its measurement. The subtler half was one line lower:
  `_key_of` read `_dotenv_key(text) or _first_key(text)`, so a refusal **fell through to the first token
  in the file**, reinstating the original bug beneath its own fix. The last assignment now always
  decides, including when it decides "unparseable"; the fallback survives only where no such assignment
  exists at all.
- **The batch GATE was fenced; the HOOK that carries its verdict was not** (`f398c1b`). Three contract
  properties live only in `up-injected.sh` and had no test of any kind. The sharpest: replacing the
  `if/else BATCH_RC=$?` capture with `|| true` — the idiom the surrounding script uses for genuinely
  optional steps — discards the deliverable, so a bring-up with a **fully red batch exits 0 under an UP
  banner**, inverting the milestone's central claim with no failing signal anywhere. Also fenced: a bare
  call (which `set -e` turns into clause 5's worst violation, a test bug costing a good demo), the UP
  banner's ordering, the verdict reaching `exit "$BATCH_RC"`, no teardown between gate and exit, and the
  opt-out's polarity.

**Test-only gap closed:** *"never accumulates" is not "repairs an accumulation"* (`60bb4f6`). iter-03's
three tests all start from a clean file and prove the injection block never stacks **going forward** —
but `--purge` does not clear the stack dir, so every stack dir predating the fix still carries its stack
of blocks today, and the ISOLATION false RED persists there until a re-up collapses them. Added at the
measured size (24 blocks): 123 lines → 8, one header, no previous key surviving, fixed point in ONE
pass. The fixture carries a blank `DIRECTUS_TOKEN=` deliberately — it ships blank and filling it is the
classic *stack boots, catalog empty* fix, so a rewrite dropping empty-valued foreign lines would
silently undo a user's repair.

**Flakes stabilized:** none surfaced.

**Verification:** `stack-injection` 342 green · `stack-core` isolation 53 + buildbench 152 green ·
`demo-stack` 1119 run, the same **9** pre-existing failures as pass 1 — **nothing introduced**.

**Stop condition:** continue-to-next-pass — two production defects and one coverage gap landed, and two
surfaces came back already-sufficient. The finding rate is falling (5 → 3) but the scan is still
producing, so it has not stabilised. Pass 3 to sweep the last unexamined milestone code: iter-20's G1
repair in `platform_predicate_guard.py` (the newest fence code in the milestone) and iter-06's
`buildbench.py` batch wiring.

## Pass 3 — 2026-08-12 — final

**Iters hardened this pass:** the milestone's newest fence code — iter-20's G1 `(c) DOMAIN`
discriminator (`stack-core/platform_predicate_guard.py`) and iter-06's batch-applicability wiring in
`stack-core/buildbench.py`.

**Coverage delta on touched files:** `platform_predicate` 190 → 191 · `buildbench` 152 → 159.

**Tests added:** +1 G1 reach boundary (two directions, RED-proven by widening the window to the whole
file) · +7 batch-applicability derivation.

**Bugs surfaced:** **none.** Both probes came back sound, which is a result and is recorded as one — a
harden pass that reports only what it changed hides how much of its scan came back clean.
- G1's `(c)` discriminator: I expected a fail-open where a real dead compose token sits near a genuine
  `hostprofiles/` mention, and **measured that it is not one**. The exemption reaches exactly
  `_pin_window`'s two lines — the claim's own plus the preceding one, because prose WRAPS — and stops:
  a mention two lines away, or across a blank line, still reports the token. A wrapped claim is
  exempted; a neighbouring claim is not. That is `D-M257x-63-1` (*a pin bought silence for its
  neighbour*) already correctly applied. Its battery pinned the discriminator's VOCABULARY and not its
  REACH, so the reach is now pinned in both directions.

**Testability gap closed (`5b86007`):** iter-06 gave the batch anchor five tests and every one passes
`batch=` directly, while the derivation computing it from a rep's env snapshot lived as two inline
expressions inside the 200-line campaign `run()` — exercisable only by running a real campaign. **A
tested consumer fed by an untested producer is a covered call site and an uncovered decision:** get the
derivation backwards and all five keep passing while every campaign silently excuses a missing batch
phase. Extracted (semantics unchanged) to `knob_is_on` + `batch_is_expected` and pinned across all four
`(public_host × knob)` combinations, plus absent-reads-as-ON (the `DEMO_*` family is opt-OUT, default
`0`, so absent means ON — the other reading would excuse the anchor on every ledger predating the knob),
a non-string value (a JSON ledger can carry `1` as an int, and `1 != "1"` reads as ON), and only an exact
`"1"` switching a knob off, so `DEMO_NO_BATCH=""` does not excuse an anchor. Two mutants, 2 and 1 fails.

**Stop condition:** continue-to-next-pass — zero defects, but two coverage gaps closed and the sweep had
not yet reached `check-cockpit-roster.py`, `rosetta-demo`, `derivation_registry.py` or `stack-snapshot`.

## Pass 4 — 2026-08-12 — final

**Iters hardened this pass:** the remainder of the cumulative footprint — `check-cockpit-roster.py`
(iter-06), `rosetta-demo`'s `down -v` (iter-14), `derivation_registry.py` (iter-10), `stack-snapshot`
replay timings (iter-09).

**Coverage delta on touched files:** `playthroughs/manifest` +4 subtests · census module +1.

**Bugs surfaced + fixed inline:**
- **A keyless seat was satisfied by a keyless identity** (`07177c1`). `check-cockpit-roster.py` is the
  post-condition that stops a restored world advertising seats the roster cannot serve — it matters
  because the fake FAPI signs an unknown identity in ANYWAY, as whoever was last active, so the failure
  is a silent WRONG login rather than a broken button. Both sides read their key with
  `dict.get("key")`, which yields `None` for a malformed entry, and `None in {None}` is True: a hero
  with no `key` was satisfied by a roster entry with no `key`, and the check printed *"✓ all 2 cockpit
  seats resolve"*. With a well-formed roster the same input **crashed** instead — `sorted()` over mixed
  `None`/`str` raises TypeError, which escapes `main`'s try (it wraps `load`, not `check`) and prints a
  traceback where a diagnosis belongs. Fail-open in one direction, unreadable in the other; neither is
  a verdict. Keyless roster entries are now dropped (they can satisfy nothing) and a keyless hero is
  named as the orphan it is. Two controls keep the fix from being blunt: a well-formed pair still
  passes, and the LIVE defect's own shape is still named.

**Claim converted to a measurement (`e96c437`):** `_CENSUS_SKIP`'s `stacks` entry is component-EXACT,
justified by the comment *"these are the only two directories named `stacks` in the tree"* — true today
(`git ls-files` shows no tracked path with that component; both are gitignored per-stack workspaces) and
asserted nowhere. It fails silently: a section that ever adds tracked source under a `stacks/` component
drops out of the census, the denominator shrinks, and every ratchet built on it keeps passing against a
smaller subject. Now measured, with the module's required anti-vacuity floor.

**Examined and found sufficient:** `rosetta-demo`'s `-v` (already covered by pass 1's fact-based fence,
which reads every `stack-*/platform` clone and so covers the demo path too, plus the existing twin
fence) · `stack-snapshot` replay timings (4 tests over 232 lines; whole section green) ·
`derivation_registry`'s skip rules (behavioural venv/dot-rule tests already in place).

**Flakes stabilized:** none. **Flake gate: PASSED** — 3 consecutive clean runs of every test added this
session, across Go and Python, run sequentially.

**Stop condition:** continue-to-next-pass — one real fail-open closed, and the cumulative scope is now
fully swept, so pass 5 is the cross-iter composition check rather than new ground.

## Pass 5 — 2026-08-12 — final

**Iters hardened this pass:** none — final mode's defining work, the **cross-iter composition check**.
The cumulative scope was fully swept by pass 4, so this pass asks the only question left: do this
session's seven changes, which touch four sections and two files twice each, compose?

**The interactions checked, not assumed:**
- `restore-presenter-world.sh` was edited twice (the casbin-invalidation fence in pass 1, the
  `STACK_BIN` live-resolution fix in pass 1) — both fences green together.
- `buildbench.py` was edited twice (the dotenv decision in pass 2, the batch-applicability extraction in
  pass 3) — 159 buildbench + 57 isolation green together.
- `check-cockpit-roster.py` was made STRICTER in pass 4 while `restore-presenter-world.sh` invokes it as
  a post-condition — the restore-leg fences and the roster fences pass together, and the live defect's
  own shape is still reported.
- `services.sh` and the unscoped-run disclosure were changed as a pair — caught by a coupled fence at
  the time, and still agreeing.

**Verification (all re-run after every fix was in):** corpus fence set **15/15 `rc=0`** ·
`playthroughs` 4/4 Go packages · `stack-injection` 342 · `dev-stack` 166 · `stack-verify` 275 ·
`stack-snapshot` 6/6 packages · `stack-core` scoped modules green (buildbench 159, isolation 57,
platform-predicate 191, service-registry 38, bringup-verify-scope 15, emphasis-reach) ·
`demo-stack` 1119 run with **the same 9 failures, byte-identical to the pass-1 list** — `diff` of the
two failure sets is empty.

**⚠️ ONE RATCHET GREW, and it is recorded rather than absorbed.** `test_the_population_does_not_GROW`
counts measurement literals outside a `print()`. Measured on both trees: **pristine `5566538` = 248
(dated 120), HEAD = 249 (dated 121)** against a ceiling of **240**. So the ratchet was **already
breached by 8 before this session** — it is part of the pre-existing census family iter-18 measured and
iter-20 routed forward — and this session added **exactly one** literal: the `range(24)` in the
injection-collapse test, reconstructing the measured 24-block shape. The delta is **entirely in the
`dated` class** (120 → 121; `doc-relative` 7 and `standing` 121 are unchanged), which is the instrument's
own sanctioned remedy — *"derive it, date it, or raise the ceiling with a reason"*. **No ceiling was
raised.** The literal is load-bearing: 24 is the count measured on the live `demo-1`, and the test exists
to prove the collapse works at that size.

**Pre-existing failure attribution, re-derived for the census family:** `frozen_expectation` reports
**11** failures at HEAD and **11 at pristine `5566538`** — the **same eleven by name**. Checked
deliberately, because this session added ~600 lines of heavily-documented test code and the ceiling
tests are exactly the shape that would notice: nothing was introduced. My own added test
(`test_no_TRACKED_file_lives_under_a_stacks_component`) is not among them.

**Flakes stabilized:** none. Flake gate re-confirmed in pass 4 (3 consecutive clean runs, both languages).

**Stop condition:** stabilized — the cumulative scope is fully swept, this pass introduced no new tests
and found no new defects, every touched suite is green, and the `demo-stack` failure set is
byte-identical to where the session started.
