# Hardening Ledger — M257x platform re-alignment

Two repos are in scope by construction (the milestone's exit gate splits 1/2/4 → `rosetta-extensions`,
3/5 → the `rosetta` corpus), so every pass below records work in both and names the repo per commit.

**Measured baselines for this milestone** (do not read a pre-existing failure as a harden regression):

| section | baseline |
|---|---|
| `stack-verify` | **RED — 11 failures + 1 error of 224.** Pre-existing; re-measured this pass on a pristine `git archive HEAD` control (which shows 12F+3E only because the archive has no `node_modules`, so the 3 `test_e2e_collection_integrity` cases degrade there and not in the real tree) |
| `stack-core` | 14 failures (m255/m220 batteries) |
| `demo-stack` | 7 failures of 1029 (`CHECK-M257x-live-clone-suites-red`) |
| `stack-injection` | OK (277, 1 skip) · `dev-stack` OK (122) · all Go sections green |

Host has no `pytest`; suites run as `cd <section> && python3 -m unittest discover -s tests -q`.

---

## Pass 1 — 2026-08-01 — incremental

**Iters hardened this pass:** iter-01 … iter-15 (the whole milestone — first harden pass, so the ledger's
"last terminating commit" is unset and scope is every closed iter to date)
**Tiks covered since prior pass:** 15
**Scope:** rext `31d2b5d..af00ad9` (77 files, +4165/−431) + rosetta `d2a6fbb..4a509a5` (10 corpus docs +
the milestone's own plan tree). No `iter_shape` frontmatter is declared on any iter of this milestone, so
all 15 were hardened as production-fix tiks (the default carve-out).

**Coverage delta on touched files:**
- `stack-snapshot/cmd/stacksnap` (pkg): 80.0% → 80.3% stmts
- `stack-snapshot/cmd/stacksnap/autoprovision.go::reconcileSequenceOwnership`: 88.9% → **100.0%**
- `stack-verify/lib/readiness.sh`: no line coverage tooling for bash in this repo; the new probe is
  covered by 13 behavioural tests + 6 mutants (the section's established discipline)

**Tests added:**
- iter-14/iter-15 → `stack-verify/tests/test_verify.py`: **13** (1 happy-path, 5 error-path, 4 edge-case,
  3 anti-self-satisfaction/wiring) in the net-new `TestDirectusServesContent`; 1 pre-existing assertion
  re-pointed (`TestDirectusCheapWins.test_registered_collections_nonzero_passes`)
- iter-15 → `stack-snapshot/cmd/stacksnap/autoprovision_test.go`: **3** (1 SQL plan-safety fence, 1
  fence-can-fail control, 1 error path) + 1 fake field

**Bugs surfaced + fixed inline:**
- **`autoverify` could grade `green:true` over a Directus serving nothing** (commit `1eb2922`, rext). The
  only Directus check in the whole verify path counted rows in `directus.directus_collections` — a
  REGISTRY table the STRUCTURE replay populates — and nothing anywhere asked the running Directus for an
  item. This is why clause 1's three checked-in verdicts are compromised: the content replay had failed on
  every cycle (0 of 11986 rows; every anon `/items` read 403). iter-15 fixed the replay; nothing had yet
  closed the hole that let the failure be graded green. Fixed by `probe_directus_serves_content` — target
  collection DERIVED from the stack's own catalog (never hardcoded), derivation FAIL-CLOSED (no collection
  with a row IS the defect, not a skip), exact counts via `query_to_xml` (`reltuples` is −1 on a
  freshly-replayed never-analyzed table, i.e. on exactly the stack shape being graded), and the read goes
  over HTTP to the running Directus on the stack's own offset port — an independent measurement from the
  Postgres count that selected the target. 403 / `data:[]` / no-data-array are each named distinctly.
  autoverify's ✓ line, which asserted *"the local Directus serves the captured catalog"* on all three green
  cycles, now claims REGISTRATION only.
- **The plan-safety property was a comment on one query and a check on its twin** (commit `a657ec1`, rext).
  iter-15 fixed `identitySeqColumnsSQL` with a MATERIALIZED barrier **and fenced the barrier**;
  `unownedSequenceOwnershipSQL`, added in the same commit, is safe for a different reason (both arguments
  derive from joined relation columns, so no single-relation pushdown is possible) — recorded only in
  prose. An edit "simplifying" `quote_ident(n.nspname)` back to `quote_ident($1)` would reintroduce the
  exact raising shape with every existing test still green. Now fenced, with balanced-paren argument
  extraction (a naive scan stops inside `quote_ident(` and inspects a truncated string that passes
  everything) and a control that runs the fence against the known-bad shape and requires it to REJECT.
- **An error path unreachable through its own fake.** `fakeProvisionConn.applyErr` failed *every*
  `ExecScript`, so it could only ever exercise the structure script — the sequence-ownership
  reconciliation's own apply-failure path had no way to be reached. Fake gained `applyErrOnCall`.

**Mutation results:** 7 mutants, **7 RED**, controls GREEN.
`fail-closed→pass` · `hardcoded collection` · `403 accepted` · `empty-data check dropped` ·
`base port not offset` · `probe unwired` · `predicate reverted to bind parameters`.
Every mutant was applied to source and observed failing its named test before being reverted.

**Flakes stabilized:** none observed. 4 `ResourceWarning`s from unclosed file handles in the new + adjacent
source-reading tests were closed at authoring time rather than shipped.

**Knowledge backfill:** deferred to Pass 2 (`corpus/ops/platform-alignment.md` §5 — the registry-vs-serving
distinction generalises past Directus and belongs in the protocol doc, not only in a probe comment).

**Stop condition:** continue-to-next-pass — the pass swept the two subsystems carrying the milestone's
load-bearing unhardened finding (`stack-verify` grading, `stack-snapshot` sequence provisioning); the
remaining iter-touched subsystems (`stack-core` fences, `stack-injection`, `demo-stack` / `dev-stack`
bring-up scripts, `stack-seeding` re-points) have not yet had a dimension scan, so no cross-pass coverage
delta exists to measure stabilization against.

---

## Pass 2 — 2026-08-01 — incremental

**Iters hardened this pass:** iter-01 … iter-15 (same scope; pass 1 covered `stack-verify` + `stack-snapshot`,
this pass swept the remaining iter-touched sections)
**Tiks covered since prior pass:** 0 new tiks — this is the second pass of one harden session
**Scope:** `stack-core`, `stack-injection`, `stack-seeding`, `demo-stack`, `dev-stack` — driven by two
parallel read-only dimension scans over the per-section iter diffs, hunting the milestone's dominant class
(*a check that reports a state without measuring it*). The scans returned **20 concrete findings**, each
backed by a source mutation the reporter had verified leaves the suite GREEN.

**Coverage delta on touched files:** stack-core 363 → 365 tests · stack-injection 277 → 286 ·
stack-seeding +3 helper controls. Line coverage is not the useful number this pass — every gap below was a
test that ran and asserted nothing, so it scored as *covered* throughout.

**Tests added:** 9 (`TestDependsOnResolvableFence`) + 2 (Directus parity presence + port-divergence) +
3 (collapsed-write helper controls); 3 tautological tests rewritten.

**Bugs surfaced + fixed inline** (commit `5889e92`, rext):
1. **`FENCE-M257x-iter13-compose-service-exists` had no tests at all.** iter-13 landed the fence and — in
   the same diff — enlarged `_main_cfg()` with `backend` + `postgresql` so the fence would stop failing the
   four `main()` tests. Nothing imported `assert_depends_on_resolvable`, called it, or asserted on its
   `SystemExit`; the return value it documents as existing "so a test can prove the fence actually looked at
   something" was consumed by nobody. Replacing the whole body with `return set()` left **156/156 GREEN**.
   *A fence adjusted-to until it stops complaining, with no test that it can complain.*
2. **The Directus parity fence passed on SYMMETRIC DELETION.** `_find` → `None` for an absent key, and the
   comparison only asserted `d == m`, so `None == None` passed. Dropping `image:` / `mem_limit:` /
   `DB_SEARCH_PATH` from **both** emitters retired three stated rules at once with the fence green.
3. **The same fence's exclusion list named a trigger it never measured.** Its docstring said "if the secret
   *and the port* ever became identical the exclusion would be over-broad" — the port condition was
   **already true** at the class's own `(1, 10000)` fixture, and the body asserted on the secret only.
4. **Three seeder tests became tautologies** when iter-06 collapsed the session mirrors: both lookups now
   return the same `copyCall`, so every comparison was an identity, under failure strings still describing
   the pre-M257 two-table world.

**A fence deliberately NOT built.** The first cut of the seeding work added a runtime
`assertNoDroppedMirrorWrites`. It was redundant — `stack-core/tests/test_dropped_mirror_fence.py` already
scores every quoted occurrence of a dropped relation across every rext artifact, statically and repo-wide —
**and that fence went RED on the new file**, catching the duplication itself. A weaker runtime twin of a
stronger static fence is a second thing to keep in step, which is the wrong-twin shape this milestone has
already paid for. Only the proposition the static fence *cannot* see was kept: **how many times** the
surviving canonical table is written (both writes name a live table, so a resurrected mirror step doubles
every row invisibly).

**Mutation results this pass:** 4 mutants, **4 RED**, controls GREEN — symmetric `mem_limit` deletion;
a write re-pointed onto a dropped mirror (3 seeder sites); a duplicated canonical write. Session total
**14 mutants, 14 RED**.

**Knowledge backfill:** `corpus/ops/platform-alignment.md` §5 **rule 14 — REGISTERED is not SERVED**
(commit `8f94463`), generalising pass 1's finding, with the same registry-vs-served split named in four
other systems. Plus **three corpus sites corrected that asserted the wrong side of it**:
`verification.md:114` ("Directus **serves** the catalog" over a registry count), `directus-local.md:247`,
and `snapshot-spec.md:419` ("the replay exits 0 and a booted Directus serves the captured catalog" — true
as design, false in fact from M21 until iter-15) — commits `8f94463`, `d1a924c`. Each retraction is fenced
in place rather than quietly rewritten.

**Flakes stabilized:** none observed; the newly added tests passed 3 consecutive runs (flake gate clean).

### Routed forward — findings NOT fixed this pass

The scans found more than one harden session can land well. These are recorded here with the evidence, so
they are a queue rather than a rediscovery. **Every one is backed by a named source mutation that leaves the
suite green today.** They are Fate-3 routed-forward, not dropped.

| id | file | what it claims vs measures | mutation that stays GREEN |
|---|---|---|---|
| **RF-1** | `dev-stack/migrate-dev.sh:95-106` | Prints *"done — the derived migration set applied"* and exits 0. Classifies **every** non-zero atlas exit as `"had migration warnings (non-fatal — see atlas output)"` — with the output discarded by `>/dev/null 2>&1`. **The demo twin (`migrate-demo.sh:149-177`) has `mig_fail` + output capture + `exit 1`; the dev twin never got the fix**, and `test_host_prereqs_m215.py:349` bans that exact string but is pointed only at the demo file. The wrong-twin class, live. | Replace the whole atlas block with `log "  $r ok"` — no dev-stack test executes the loop |
| **RF-2** | `demo-stack/ant-academy.sh:685-711` | The entire "SERVING is not RENDERING" check is executed by **zero** tests (all four test references are `assertNotIn`). Its message reports a `200` it never measured — `serving=1` is set on `2??\|3??`, and the render probe is a *separate* `curl -fsS … \|\| true` whose status is discarded (`-f` collapses a 5xx into silence — the defect iter-10 fixed 40 lines above, still live in the paired check). | `grep -qi ''` at :700, or delete :685-708 |
| **RF-3** | `demo-stack/up-injected.sh:1739-1749` | iter-04's FATAL apply-authn path has no test; `authn_out`/`authn_rc` appear in no test file. | Revert to the pre-iter-04 one-liner — the ordering assert at `test_frontend_build.py:1156` still passes |
| **RF-4** | `dev-stack/dev-setdress.sh:365-395,427` | The `*)` arm turns every unclassified rc into `skipped(error)`, returns 0, and the verdict line prints `set-dressed (…)` over it. **This is the exact M257x clause-1 signature, still in the changed file.** | Collapse the `case` so every rc lands in the `4)` arm — both surviving tests still pass |
| **RF-5** | `stack-core/tests/test_write_target_schema_fence.py:232-249` | Of 92 scored write targets, only the **41 COPY** hits are exercised end-to-end; the 19 steps-slice and **32 `resetTables`** entries depend on a `scan_text` state machine no test drives. The anti-vacuity floor (`found > 40`) is cleared by the 41 COPYs **alone**, so the margin is exactly 1. | Make `scan_text` never set `in_block = True` — 51 of 92 targets go dark, fence stays green |
| **RF-6** | `stack-core/tests/test_write_target_schema_fence.py:161-209`, `lib/repos_yml.sh:137-143` | `discover_repos_yml` takes the first glob hit and never reports which file it read; aimed at a stale pre-fold `repos.yml` it re-admits the two dead schemas and all 13 tests pass. Its own test's assertions are guaranteed by the hardcoded `ANCHOR_SCHEMA`, not by the derivation. `repos_yml.sh:140`'s `[ -f "$f" ]` is a silent skip of the whole declared-schema half. | Delete `repos_yml_declared_schemas` from `repos_yml.sh:140` |
| **RF-7** | `stack-seeding/seeders/ai_readiness_funnel.go:999` + `cmd/stackseed/main.go:62,104,105` | Four COPY targets with **no offline assertion of any kind**; `main_test.go`'s `mustCover` is a hand-maintained list — the shape iter-08 removed from the *other* fence and left here. | Rename `interview_aggregated_reports` to anything else under `public` |
| **RF-8** | `stack-seeding/cmd/stackseed/main.go:763-791` | `--reset` prints `reset complete` regardless of how many tables were silently skipped as absent. A mis-named target reports its own absence and the command still succeeds. | (no mutation needed — the grading is unchanged since iter-06) |
| **RF-9** | `stack-injection/tests/test_apply_authn.py:608` | The shellcheck guard still `skipTest`s when the binary is absent — **in the very commit whose message names that class**. Latent on this box (shellcheck present), live on a clean CI host. | (no mutation needed) |
| **RF-10** | `stack-injection/tests/test_apply_authn.py:534-542` | The class docstring declares "never a whole-file substring" and then uses three; `up-injected.sh` already carries comments mentioning `apply-authn`. | Set `authn_rc=1` while adding a comment containing `authn_rc=$?` |
| **RF-11** | `demo-stack/up-injected.sh:1119` | `browser_graphql_endpoint` is fenced by `count(...) >= 3` against **6** call sites, so two can be re-inlined; `build_frontend_hiring` has no stale-offset test, unlike its next-web and studio-desk siblings. | Re-inline the old endpoint at :1119 — count 5 ≥ 3 still clears |
| **RF-12** | `stack-core/tests/test_write_target_schema_fence.py:336-351` | The derived scope is **Go-bearing sections only**: `stack-verify`, `demo-stack`, `dev-stack`, `stack-core`, `stack-injection` are neither scored nor required to declare coverage. No live violation found — a scope gap, not a breach, but the same narrower-than-it-claims shape iter-08 exists to close. | (scope widening; no single mutation) |

**RF-1 and RF-4 are the two that would most plausibly cost a gate cycle**, and RF-4 is the clause-1
signature itself. They are source changes to bring-up scripts with real blast radius (RF-1 adds an `exit 1`
to the dev migration path), which is why they are routed to an iter rather than bundled into a harden pass —
per the fixable-inline boundary, a bundled fix of that size blurs the iter→harden→fix attribution.

**Stop condition:** continue-to-next-pass — pass 2 fixed 4 of 20 findings and routed 12; the remaining
queue is real, enumerated and evidence-backed, so coverage has demonstrably not stabilized.

---

## Pass 3 — 2026-08-01 — incremental

**Iters hardened this pass:** iter-01 … iter-15 (working the Pass-2 routed-forward queue; scope
`stack-core`'s iter-08 write-target fence — RF-5 and RF-6)
**Tiks covered since prior pass:** 0 new tiks — third pass of one harden session
**Pass cap:** 3 of 3 for incremental mode. This pass ends the session.

**Coverage delta on touched files:** `stack-core` 365 → 372 tests. Again not a line-coverage story: both
gaps were code that ran on every invocation and asserted nothing about itself.

**Tests added:** 7 — 4 driving `scan_text`'s block state machine, 1 per-construct vacuity floor,
2 proving the legal-schema derivation is actually read (both directions).

**Bugs surfaced + fixed inline** (commit `46f8cc3`, rext) — both in **the very fence this milestone built
because the class has now recurred three times**:

1. **51 of 92 write targets could go dark with the fence green (RF-5).** Measured breakdown of the scored
   targets: **41 COPY calls · 19 steps-slice entries · 32 `resetTables` entries**. Only the COPY rule was
   exercised end-to-end — the existing behaviour test calls `scan_line()` with its *default*
   `in_block=True`, so it never runs `scan_text` and never touches `RE_STEPS_TYPE`, `RE_RESET_OPEN`, or the
   state machine at all. Make `scan_text` never open a block and both block-scoped constructs vanish,
   including every `resetTables` entry — which is interpolated **raw into `TRUNCATE TABLE … CASCADE`**.
   And the anti-vacuity guard could not notice: `found > 40` is cleared by the 41 COPY hits **alone**, so
   the margin on the only vacuity check in the file was **exactly one**.
   Fixed on both axes: 4 tests that drive `scan_text` (never `scan_line`), including block *closing* — so a
   casbin-grant lookalike outside the block is still not re-scored, the false-RED that gets a fence disabled
   — and a **per-construct** vacuity floor. *An aggregate a single construct satisfies cannot detect a
   construct going blind, which is the one thing it exists to detect.*
2. **The legal-schema derivation was never proven read (RF-6).** Its test asserted `"public" in legal` —
   guaranteed by `out.add(ANCHOR_SCHEMA)`, not by any derivation — and `"skillpath" not in legal`, satisfied
   by two hardcoded strings that never mention it. Removing the declared-schema half of the derivation from
   `lib/repos_yml.sh` entirely returns a **byte-identical** set, because `app → public` is also what the
   anchor supplies. Now driven with a fixture declaring a schema no anchor or infra constant could supply,
   plus the other direction: a schema present in a **stale pre-fold** `repos.yml` must LEAVE the set when
   the platform stops declaring it — the wrong-vantage class (iter-11), which would otherwise silently
   re-admit the two dead schemas this fence exists to catch.

**Mutation results this pass:** 4 mutants, **4 RED**, controls GREEN — state machine never opens ·
`RE_RESET_OPEN` never matches · block never closes · declared-schema derivation removed.
**Session total: 18 mutants, 18 RED.**

**Flakes stabilized:** none observed. Flake gate clean — 3 consecutive green runs of every test added in
passes 1–3.

**Final suite state (both repos, all sections):**

| section | result | vs baseline |
|---|---|---|
| `stack-verify` | 11 failures + 1 error of **237** | baseline 11F+1E of 224 — **unchanged**, +13 tests |
| `stack-core` | 14 failures of **372** | baseline 14 of 363 — **unchanged**, +9 tests |
| `stack-injection` | **OK** (286, 1 skip) | baseline OK (277) — +9 tests |
| `dev-stack` | **OK** (122) | unchanged |
| `demo-stack` | 7 failures of 1029 | baseline 7 of 1029 — **unchanged** |
| Go sections | all **GREEN** | `stack-snapshot` `replay` 98.2% → **100.0%** |

No pre-existing failure was fixed and none was added.

**Live proof (pass 1's probe).** `probe_directus_serves_content` was run against the live `demo-1` stack
rather than only against stubs: the derivation resolved `task_sub_checks` out of the real catalog, the anon
`GET /items/task_sub_checks?limit=1` returned **200 with a real item**, and the probe exited 0. Aimed at an
offset with no Directus it went RED naming the state correctly (*"no response from the Directus on port
58055"*). Both verdicts observed on a real stack — a probe that has never executed is this milestone's own
defect class.

**Stop condition:** cap reached without stabilization — 3 incremental passes fired. Coverage has NOT
stabilized: passes 2–3 fixed 6 of the 20 scanned findings and **RF-1 through RF-4 and RF-7 through RF-12
remain open**, each with a named source mutation that leaves the suite green today. This is not a plea for
a fourth pass — the remaining queue is dominated by **source** changes to bring-up scripts (RF-1's `exit 1`
on the dev migration path, RF-4's `skipped(error)` verdict) whose blast radius belongs to an iter, not to a
harden pass. See §"Routed forward" in Pass 2 for the full table.

**Recommended disposition (for the orchestrator/user, not decided here):** route **RF-1** and **RF-4** into
the next iter — RF-4 is the clause-1 signature itself, still live in a file this milestone edited, and RF-1
is the wrong-twin class with the demo half already fixed. The remaining test-only items (RF-3, RF-7, RF-9,
RF-10, RF-11) are natural work for the **final** harden pass after the gate fires.

---

## Pass 4 — 2026-08-01 — incremental

**Iters hardened this pass:** iter-24, iter-25 (the two most recent production fixes in the
iter-16..26 window; iter-26 is a measurement tik with no runtime surface)
**Tiks covered since prior pass:** 11 (iters 16–26; Passes 1–3 terminated at `e028c77` / rext `46f8cc3`)

**Coverage delta on touched files:** no coverage instrumentation exists for these sections (stdlib
`unittest`, no pytest on this host), so the measurement is test count + mutation kill, as in Passes 1–3.

| section | before | after |
|---|---|---|
| `playthroughs` (Go) | **FAIL — 2 standing RED** | ok, +2 tests |
| `stack-injection` | OK 288 | OK 297 (+9) |

**Tests added:** iter-25 → `playthroughs/manifest/runner_safety_test.go`: 2 net-new (1 completeness
fence, 1 executable behaviour test) + 3 corrected asserts + 3 helpers · iter-24 →
`stack-injection/tests/test_directus_consumer_derivation.py`: 9 net-new (6 offline derivation, 2 twin,
1 live).

**Bugs surfaced + fixed inline:**

1. **Two Go tests have been RED since iter-25 — and the milestone's baseline records "Go sections
   green" (commit `8aad6ce`).** `TestRunnerSafety_ResetContract` and
   `TestRunnerSafety_RosterRefreshGate` in `playthroughs/manifest` fail on `m257x/platform-realignment`
   at `0ab2921`. The suite was not re-run after iter-25. Both are the same root cause and it is the
   milestone's own defect class in its purest form: the fence read the **raw file**, so when iter-25
   routed every call through a resolved `"$STACKSEED"` path, assert (1) — *"uses the real `stackseed
   --reset --stack` path"* — went on passing, **satisfied by the ECHO line printed beside the call**
   (`run-playthroughs.sh:126`) rather than by any call. The invariant it names (never an additive
   re-seed as the reset) has therefore been unenforced since iter-25 **while reporting enforced**,
   and its two siblings went RED against a runner that was correct. A fence that measures the
   narration of a command instead of the command is exactly what this milestone exists to find.
   Fixed by measuring EXECUTED lines only (`shellInvocationLines` — comments, blanks and
   `echo`/`printf`-leading lines dropped) and by **deriving** the tool reference from the runner's own
   `<VAR>="${PT_STACKSEED:-…}"` assignment instead of pinning the spelling `"$STACKSEED"`.
2. **iter-25's sibling-sweep is now testable, which it was not (commit `8aad6ce`).** There were three
   bare `stackseed` calls; the first fix swept one *while citing the sweep rule in its own commit
   message*, and the two it missed are `|| true`-shaped, so they degraded in silence — after a
   reset-to-seed swapped the world, the fake-FAPI kept serving the previous world's identities.
   `TestRunnerSafety_StackseedResolution` asserts **completeness, not presence**: zero bare
   invocations at *command position* anywhere in the file (string literals blanked; lines split on
   `;`/`&&`/`||`/`|`/`if`/`then`/`do`; only each segment's first token considered), plus an
   anti-vacuity floor — a "no occurrences" assert passes beautifully once the calls are *deleted*, so
   all four legs must still be present through the resolved reference — plus an `[ -x ]` check on the
   two non-fatal legs specifically. `TestRunnerSafety_RefusesUnresolvedStackseed` then **executes**
   the runner: `--reset` with an unresolvable binary must exit 2 saying *"REFUSING to continue"*, and
   its positive control demands a stub's own exit status **37** surface — proving the guard was passed
   *and* the resolved path actually invoked. `bash` absence is a `t.Fatal`, never a `t.Skip`.
3. **iter-24's consumer list is now DERIVED rather than asserted (commit `749ee0f`).** iter-24 fixed
   the value; the class was untouched — nothing could notice a *reader moving*.
   `derive_directus_readers()` computes the set from the platform's own Go source (`*.go` files naming
   `DIRECTUS_BASE_ADDR`, mapped through `INJECTED`). Dated and citable: `app`'s Go source first
   mentioned the var at platform **`38ee0c44` (2026-07-27**, "fold cms GraphQL subgraph into app"),
   zero before — **this fence would have gone RED that day**; a human found it on 2026-08-01 by
   reading 96 lines of 403 in a live backend log. `*.go` only is load-bearing: `app/terraform/main.tf`
   and eight `app/knowledge/plan/…` roadmap docs carry the string and say nothing about who *reads*,
   and counting them makes every repo a reader. It is also what makes the **pre-fold fixture derive
   the pre-fold answer** — the property that separates a derivation from a constant in a function's
   clothes. The twin check pins both emitters to the *derived* set rather than to each other, because
   bare parity **passes on symmetric deletion** (empty both and they agree perfectly while every stack
   reverts to reading production content — the hole Pass 2 found in another parity fence).

**Mutation results this pass:** 15 mutants, **15/15 matching declared expectation** (12 declared-RED
all killed, 3 declared-GREEN no-op controls all survived), controls GREEN before and after each
battery, every shell mutant `bash -n`-gated.

The three surviving controls are the point of the pass, not padding:
- stripping the `stackseed` prose out of **every echo line** must not be noticed — it *could not have
  passed before this pass*, because that prose was what the fence was reading;
- renaming `STACKSEED` → `SEEDBIN` throughout must not be noticed — the guard asserts the derivation,
  not the spelling;
- reordering both consumer tuples to `("backend", "cms")` must not be noticed — semantically identical
  for a membership test, so a test that reddens on it is asserting the literal.

**Knowledge backfill:** none this pass (both findings are tooling-internal; no corpus claim moved).

**Stop condition:** continue-to-next-pass — the iter-16..20 half of the window is unswept, and the
scan has already named its targets: `platform_alignment_guard.py`'s live positive control skips
itself when the git-ignored `stack-demo/platform/repos.yml` is absent, while
`corpus/ops/platform-alignment.md:616` asserts it runs "on every suite run"; and iter-18's Directus
bootstrap race has no test forcing both sides of the race.

---

## Pass 5 — 2026-08-01 — incremental

**Iters hardened this pass:** iter-16, iter-20 (the earlier half of the iter-16..26 window)
**Tiks covered since prior pass:** continuation of the same 11-tik window (Pass 4 covered iter-24/25)

**Coverage delta on touched files:** test count + mutation kill, as above.

| section | before | after |
|---|---|---|
| `stack-core` | 14F of 390 | 14F of **392** (+2; the standing 14 unchanged) |
| `dev-stack` | OK 132 | OK **138** (+6) |

**Tests added:** iter-20 → `stack-core/tests/test_platform_alignment_guard.py`: +2 (a whole new
unskippable class) · iter-16 → `dev-stack/tests/test_dev_stack.py`: +6 (`MigrateDevAtlasClassification`,
the first tests that EXECUTE `migrate-dev.sh`'s atlas loop).

**Bugs surfaced + fixed inline:**

1. **The milestone's flagship fence skipped its own live control on any box without a demo stack**
   (rext `79bbc0d`, corpus `e74dad6`). `platform_alignment_guard.py` is iter-20's durable deliverable
   and **nothing in CI or any bring-up path invokes it** — its only automatic execution is its test
   file. That file pointed at ONE hardcoded path, `stack-demo/platform/repos.yml`, and a platform
   clone is git-ignored and ephemeral (`/demo-down` removes it; a fresh checkout never had one). On
   such a box both real-artifact tests did not fail, they **skipped** — and a skip reports success.
   RF-9's class, in the milestone's headline fence. Meanwhile
   `corpus/ops/platform-alignment.md` §8 asserted the test runs the real map against the real
   `repos.yml` *"on every suite run"*: a claim about a measurement that was not being taken, written
   into the fence's own documentation.
   The substantive fix is a **split**, not a wider search. Assertions **C** (state vocabulary), **D**
   (every row cites evidence) and **E** (census/clone-set overlap) are properties of the **map
   alone**; only A and B need a clone set. They now run against a `repos.yml` synthesized from the
   map's own `yes` rows — A and B trivially satisfied by construction, deliberately — which makes
   C/D/E **unskippable everywhere**, floored against vacuity (≥ 5 synthesized rows or the parsing has
   broken). The A/B half then honours `PLATFORM_REPOS_YML` and searches every
   `stack-*/platform/repos.yml`, and names what it looked for when it does skip. §8 corrected to say
   which assertions actually run unconditionally.
2. **iter-16's migration classifier had never been executed** (rext `f05c1cb`). Every test of the
   RF-1 rewrite greps the source — `assertIn("mig_fail=1", guard[:600])`, `assertIn("FAILED (rc=",
   loop)`. **RF-1's own note said it**: *no dev-stack test executes the loop*, and that stayed true
   after the fix that note produced. The arm deciding whether a failed migration is FATAL is a
   `grep -qiE` over five benign phrases; three of the four outcomes can be produced by a source that
   passes every existing grep. `MigrateDevAtlasClassification` runs the real script against stubbed
   `docker` + `atlas` and asserts the **verdict and the exit code** — including a table of four
   plausible atlas errors that share vocabulary with the benign set without being benign (the
   term-scoped-audit risk as a test: a widened regex silently deletes the fatal arm), and that the
   closing *"the derived migration set applied"* line does **not** print over a failure.

**Mutation results this pass:** 13 mutants, **13/13 matching declared expectation** (9 declared-RED
killed, 4 declared-GREEN controls survived), controls GREEN before and after, shell mutants
`bash -n`-gated.

The two that carry the pass: **no platform clone on the box AND a broken map row still goes RED** —
on the vocabulary and on the citation, both of which were previously green-by-skip. And the atlas
battery **found a defect in this pass's own harness**: the stub printed its diagnosis on stdout, so
restoring `2>/dev/null` on the atlas call stayed GREEN because stdout was captured either way. A test
that cannot observe the defect it names is precisely what this milestone keeps finding; the stub now
writes to stderr, as atlas does, and that mutant dies.

**Knowledge backfill:** `corpus/ops/platform-alignment.md` §8 layer-1 row — corrected from "on every
suite run" to the actual split (C/D/E unconditional; A/B clone-dependent, searched then skipped by
name). The tooling change is what made the corrected claim *true* rather than merely weaker.

**Flakes stabilized:** none observed.

**Stop condition:** continue-to-next-pass — iter-17/18/19/21/22/23/26 have not been scanned to
completion, and the iter-18 review found its own tests unusually strong (a docker stub that already
forces both sides of the bootstrap race, `sys_tables` / `sys_tables_after` / `bootstrap_rc`), so the
next pass should look at the corpus-side iters (21–23) and at the `--local-content` verdict path
rather than re-covering iter-18.

---

## Pass 6 — 2026-08-01 — incremental

**Iters hardened this pass:** iter-21, iter-22, iter-23 (the corpus-audit half of the window)
**Tiks covered since prior pass:** continuation of the same 11-tik window

**Coverage delta on touched files:**

| section | before | after |
|---|---|---|
| `stack-core` | 14F of 392 | 14F of **396** (+4; the standing 14 unchanged) |

**Tests added:** iters 21–23 → `stack-core/tests/test_service_doc_status_fence.py`: 4 net-new.

**Bugs surfaced + fixed inline:**

1. **`corpus/services/studio-room.md` read as a live pipeline for five paragraphs** (rext `fa0028b`
   found it; corpus fix `75c9ba7`). The map calls `anthropos-studio-room` `merged-into-app` on **both**
   sides — pulled into the `app` image by CI, spawned as a subprocess from `app/internal/cms/studio/`,
   not a service, not a container, not in `repos.yml`. The doc said so **in prose, in paragraph six**,
   after five paragraphs introducing it as an AI content-generation engine and *"the manufacturing
   floor"*. Seven of the corpus's eight gone-service docs open with the standing ⚠ banner; this was the
   one that did not, and **three sweeps across iters 21–23 did not reach it** — because those sweeps
   were grep-vocabulary-bound and this doc never used the words being grepped.

**The fence, and why it is shaped this way.** iter-21's audit reported 11 residual claims, then 5,
then 2 — a curve that looks like convergence and was not: it was **exhausting its own grep
vocabulary**, and a full read found **53**. The lesson is not "grep harder". A vocabulary-based sweep
converges on itself, and nothing it produces can fence itself. What *can* be fenced is a
**derivation**: there is now exactly one place that says which services are merged —
`platform-migration-status.md`, itself machine-checked against the platform's own `repos.yml`
(layer 1, iter-20) — so `ServiceDocStatusFence` reads the map and holds every per-service doc to it.
A service that folds tomorrow fails on the day the map records it, whatever words its doc uses.

**One-way, deliberately.** Only *map-says-merged ⇒ doc-must-say-so*. The converse is not a finding,
because a correct row already looks like it: `roadrunner` is `live-standalone` in prod (still
deployed) while its doc rightly opens *"MERGED INTO `app` / ORPHANED"*. Reddening a nuanced, correct
doc is how a fence gets disabled — iter-16 paid for that lesson from the other side. A service with no
doc is likewise not a finding: `chromedp`/`simulator`/`realtime`/`web-app`/`nats` were never
documented, and demanding docs would invent work rather than pin truth.

**Mutation results this pass:** 8 mutants, **8/8 matching declared expectation**, control GREEN before
and after. 6 declared-RED — `cms.md`'s banner deleted; its banner pushed below the 20-line window; the
doc resolver blinded; `GONE_STATES` emptied; the detector widened to `return True`; and **a new row
(`sentinel`) flipped to `merged-into-app` against an un-bannered doc**, which is the arrival direction
and how this will actually fire. 2 declared-GREEN controls survive: `cms.md`'s banner **reworded**
(mark + status word kept — the fence must assert the proposition, not the sentence) and an unrelated
live row's whitespace.

**Knowledge backfill:** `corpus/services/studio-room.md` — the ⚠ status banner, in the same shape as
its six siblings, pointing at the map as the authoritative per-service state.

**Flakes stabilized:** none observed. Flake gate clean — 3 consecutive green runs of every test added
across passes 4–6.

**Final suite state (both repos, all sections):**

| section | result | vs baseline |
|---|---|---|
| `playthroughs` (Go) | **ok** | baseline recorded "green"; it was **FAIL (2)** — now genuinely ok, +2 tests |
| `stack-core` | 14F of **396** | baseline 14F of 390 — **unchanged**, +6 tests |
| `stack-injection` | **OK 297** (1 skip) | baseline OK 288 — +9 tests |
| `dev-stack` | **OK 138** | baseline OK 132 — +6 tests |
| `demo-stack` | 7F of 1030 | baseline 7 of 1030 — **unchanged** |
| `stack-verify` | 11F + 1E of 237 | baseline 11F+1E of 237 — **unchanged** |
| Go: alignment / stack-secrets / stack-seeding / stack-snapshot | all **green** | unchanged |
| Go: clerkenstein | **environmental** — `go: downloading colony v0.34.3` fails, no GH credentials in this sandbox. Pre-existing, unrelated to this pass | — |

No pre-existing failure was fixed and none was added. **Session total: 44 mutants, 44/44 matching
declared expectation** (35 declared-RED killed, 9 declared-GREEN no-op controls survived).

**Tag:** `fast-build-m257x-harden-p6` at rext `fa0028b`, **pushed and verified on origin**
(`git ls-remote --tags origin`). `.agentspace/rext.tag` and the `stack-demo` consumption clone both
still read `fast-build-m257x-iter-25b` — **deliberately not re-pinned**: this pass changed tests and one
corpus doc only, no tooling runtime source, and re-pinning a live `demo-1` mid-milestone would change
what the next gate measurement runs against.

**Stop condition:** cap reached without stabilization — 3 incremental passes fired (4, 5, 6) and pass
6 still surfaced a live drift on its first run, so the dimension scan is not clean. What remains:

* **iters 17, 19, 26 unscanned** (all measurement/withdrawal tiks with little runtime surface — the
  cheapest remaining work, and plausibly genuinely empty).
* **iter-18 was scanned and deliberately left alone.** Its own tests are unusually strong: the docker
  stub already forces both sides of the bootstrap race (`sys_tables` / `sys_tables_after` /
  `bootstrap_rc`), and `test_bootstrap_race_lost_but_schema_present_reads_as_provisioned` drives the
  race-lost branch end-to-end. The prompt's concern — *2 of 3 cold cycles never exercised the branch* —
  is about the LIVE runs, not the suite. No gap found; not re-covered.
* **The prior pass's routed-forward queue is only partly drained.** iter-16 landed RF-1 and RF-4;
  Pass 3 landed RF-5 and RF-6. Still open and **unchanged**: **RF-2** (verified still live this pass —
  `demo-stack/ant-academy.sh:700`'s "SERVES BUT DOES NOT RENDER" check is executed by zero tests),
  **RF-3**, **RF-7**, **RF-8**, **RF-9** (verified still live — `test_apply_authn.py:608` still
  `skipTest`s on absent shellcheck, in a test whose own comment names that class), **RF-10**, **RF-11**,
  **RF-12**. This pass adds **nothing** to that queue — every finding it made was fixed inline.

---

## Pass 7 — 2026-08-02 — incremental

**Iters hardened this pass:** iter-42, iter-43 (the tok + the tooling iter that TOK-02 stakes clause 5 on)
**Tiks covered since prior pass:** 15 (iters 27–43; Passes 4–6 terminated at `68eada7` / rext `fa0028b`)

**Coverage delta on touched files:**

| section | before | after |
|---|---|---|
| `stack-core` | 14F of 415 | 14F of **420** (+5; the standing 14 unchanged) |
| `claim_twin_guard.py` + `claim_ledger.py` mutants | 8 | **11** (2 of the 3 added are inversions) |

**Tests added:** iter-43 → `stack-core/tests/test_claim_twin_guard.py`: 5 net-new (`TestReachIsMeasured`);
`tests/test_m257x_claim_twin_mutation_battery.py`: 3 net-new mutants.

**Bugs surfaced + fixed inline:**

1. **The documented coverage-decay report did not exist** (rext `b7a97f4`). `claim_ledger.py`'s module
   docstring promises that an audit which *"invents a new shape is reported as contributing zero rows
   (see `coverage()`), which is a finding rather than silence (§5 rule 8 — a check that SKIPS reads
   exactly like one that PASSES)"*. It was never implemented. A `## BLOCKERS` table headed
   `| # | Where | Claim | What is true |` matches no column pattern, so `is_ledger_table` drops it, so
   `discover_ledgers` never returns the file, so it contributes to **no field at all** — `coverage()`
   came back **byte-identical** to a run in which that audit did not exist. Measured directly, in a
   temp milestone. **The honesty machinery committed the exact defect it cites**, which is the
   milestone's signature class found for the 7th time inside its own instrumentation. Now measured as
   `shape_misses` and named **with the header that defeated the derivation**, scoped to
   BLOCKERS-sectioned tables so it cannot redden on ordinary or deliberately-excluded `Minors` tables.

2. **The UNMATCHABLE report — the fence's declared 30-char fragment-floor boundary — had ZERO coverage.**
   Deleting `main`'s entire reporting loop left **15 of 15 green**. The property the fence's own charter
   rests on ("reports `UNMATCHABLE` **by name** rather than silently dropping") was itself droppable in
   silence. This is the specific hazard the harden brief named, and it was live.

3. **`--json` named none of the three fall-out classes.** It carried `coverage` (counts) and the hits.
   A wrapper — the only consumer a JSON mode has — could watch the fence's reach shrink and never learn
   what it lost. `unmatchable` / `unquoted_rows` / `shape_misses` now ship by name.

**iter-42's classification: corroborated far beyond its spot check, and the correction is upward.**
The brief flagged it as spot-verified on 4 of 18 rows, with TOK-02's whole sequence aimed by it. It is
in fact **machine-corroborated**, and by an instrument built independently of it. The partition is
exact — class A `{1..9,12,14,15,18}` (13) ∪ class B `{13,16,17}` (3) ∪ class C `{10,11}` (2) = `{1..18}`,
no gap, no overlap, every `n` matching its own enumeration. And iter-43's fence, which targets **only**
class A, went RED on **13 of 13** of it while its **two misses — #10 and #16 — both land outside it**
(#10 class C, #11's sibling scalar; #16 class B). A classification that was wrong would not produce
that alignment. Verified status: **15 of 18 rows machine-corroborated**, not 4. No defect; recorded so
the next reader does not re-litigate it.

**Mutation results this pass:** 11 mutants, **11/11 matching declared expectation**, no-op control GREEN
before *and* after the battery. The 3 added: `unmatchable-report-goes-silent` (RED — the boundary going
quiet), `shape-miss-detector-inverted` (RED — **an inversion**: it names the READABLE ledgers as misses
and lets the unreadable one through, which no removal mutant can see, §8 rule 5), and
`shape-miss-scoping-inverted` (RED — killed by the cry-wolf negative control, §8 rule 6).

**The perishable fixture is intact.** `git status --short stack-core/tests/fixtures/claim_twin/` is
empty. Nothing in `red/` was read-modified, normalised or repaired, and the live corpus's 18 blockers
are untouched — TOK-02 step 4 still owns that repair.

**Flakes stabilized:** none observed.

**Stop condition:** continue-to-next-pass — the pass swept the two most recent iters only. Unswept:
the iters 36–41 clause-2 fixes (including the open
`CHECK-M257x-iter35-negative-control-rests-on-the-same-tie`), the fifth-generation tenancy fence, and
the iters 27–35 seeder work.

---

## Pass 8 — 2026-08-02 — incremental

**Iters hardened this pass:** iter-31 (the tenancy scoping fix) + the multi-tenancy fence claim carried
through iters 33/34/38/39/41
**Tiks covered since prior pass:** continuation of the same 15-tik window (Pass 7 covered iters 42–43)

**Coverage delta on touched files:**

| section | before | after |
|---|---|---|
| `stack-seeding/seeders` | 0 fences on `memberRoleAt` call sites | **3** (+1 new file, 5-mutant battery) |

**Tests added:** iter-31 → `stack-seeding/seeders/role_tenancy_fence_test.go`: 3 net-new AST fences.

**Bugs surfaced + fixed inline:**

1. **The role-tenancy fix was proven in the function and unproven at its six call sites** (rext
   `c597dcf`). iter-31's first cut handed every org every story's hero roles — *"DevOps Engineer"*
   became a key-role card on the CONTRAST org's succession view and `negative-controls.spec.ts:429`
   caught it **live**, not in Go. `jobroleref_test.go` proves `orgRoleSet` scopes correctly **given**
   per-story input; nothing proved the **callers** supply it, and supplying the wrong thing IS the
   defect that shipped. Six hand-copied call sites (users, membership_skills, population_evidence,
   certificates, profile, target_roles) plus one forwarding helper. `memberRoleAt`'s own doc comment
   is about this exact hazard one layer down — *"the identical expression appeared in SIX seeders […]
   and the first sweep of this fix found only FOUR"* — so **the repetition moved up a level rather
   than going away**. Three fences now hold it: the argument shape at every call site, the forwarding
   helper's own callers (so the exemption cannot smuggle the aggregate out one frame), and
   `storyHeroRoleNames`'s arity (variadic/slice re-creates the aggregate while the other two stay
   green — §5 rule 17). Both scanners fail **closed** on zero sources / zero call sites / zero helper
   calls.

**Mutation results this pass:** 5 mutants, **5/5 matching declared expectation**, no-op control
**SURVIVES**. Recorded in the fence's own header so it is reproducible. **The first cut of this battery
had three mutants that did not COMPILE** — per §8 rule 5 that is not a kill, and taking it as one would
have certified a fence that was never exercised. All five now compile.

**The multi-tenancy corpus fence — measured, NOT repaired.** Re-derived against the live platform clone
(`stack-demo/app` @ `v1.363.2`): the doc's own shell derivation **reproduces exactly** — 135 `ent.Schema`
files of 139 `.go`, **30** `OrganizationMixin{}`, **7** `OrganizationIDMixin{}`, **18** plain
`organization_id`. And **only four files in the whole schema dir declare any `Policy()`**
(`mixin.go`, `org_membership.go`, `organization.go`, `user.go`), with `OrganizationIDMixin` declaring
none — so iter-41's blocker **#5 is confirmed live**: *"16 carry an `organization_id` with no policy of
any kind"* understates it, because the 7 `OrganizationIDMixin` users are unpoliced too (23). **Not
repaired here — it is one of the 18, and TOK-02 step 4 owns that repair, fence-assisted.** It is class
A, and iter-43's claim-twin fence already detects it.

**Routed forward — RF-13 (new; joins RF-2/RF-3/RF-7..RF-12, supersedes none).** *The derivation block
cannot produce the claim it supports.* `security_compliance.md`'s `Derivation:` snippet derives the
**18** mechanically and then instructs, in a comment, *"then subtract any that declare their own
`Policy()` or carry `UserMixin{}` -> 16"* — a **hand** step, un-derived, and **that is exactly where all
five failures have lived**. It subtracts down to 16 while never **adding** the 7 `OrganizationIDMixin`
users. A block that says *"Re-derive it; do not quote it"* while shipping a derivation that stops one
step short of the number in dispute is the §5 rule 17 shape in its purest form. **When TOK-02 step 4
repairs #5, the derivation must be made to derive the unpoliced set end-to-end** — otherwise the sixth
generation is another hand-subtraction, which is how the previous five were produced. Not fixed here:
the block is inside blocker #5's own blockquote, and editing it is repairing one of the 18.

**Flakes stabilized:** none observed.

**Stop condition:** continue-to-next-pass — iters 27–30, 32–41 have not been scanned to completion; the
open `CHECK-M257x-iter35-negative-control-rests-on-the-same-tie` is still unhardened, and the iters
36–41 page-object/seed changes are unswept.

---

## Pass 9 — 2026-08-02 — incremental

**Iters hardened this pass:** iter-35 (the drill-down tie) + its unrepaired twin
**Tiks covered since prior pass:** continuation of the same 15-tik window

**Coverage delta on touched files:**

| section | before | after |
|---|---|---|
| `playthroughs/e2e` positional drills in specs | **1** (unfenced) | **0** (+1 fence file, 3 assertions) |
| `stack-core` | 14F of 415 | 14F of **420** |
| `stack-seeding/seeders` fences on `memberRoleAt` | 0 | **3** |

**Tests added:** iter-35 → `playthroughs/e2e/tests/tie-ordering-fence.unit.spec.ts`: 3 net-new.

**Bugs surfaced + fixed inline:**

1. **`CHECK-M257x-iter35-negative-control-rests-on-the-same-tie` — CLOSED** (rext `e947401`). iter-35
   measured that the activity grid's *"first row"* premise was never true: **11 distinct contents share
   the same `max(started_at)` to the microsecond**, because the seeder stamps every backdated session
   at one instant, so *"most-recent activity"* does not order the grid at all and **2 of the 11 carry
   no hero session**. The two runs that "confirmed" the old selector were **two draws from a tie**. It
   fixed the Playthrough, and booked the twin — `negative-controls.spec.ts` was still drilling row 0
   one test over, *"currently green by luck"*. **It stayed that way for 8 iterations.** That is §5
   rule 19 (*repair by CLAIM, not by FILE*) observed on the suite rather than the corpus. Repaired with
   iter-35's own property-based selector; **strictly an improvement** — when row 0 does contain the
   member the scan returns 0 immediately and the path is byte-for-byte the old one, so it differs only
   in the runs the old form would have failed.

**The fence, and why it is shaped this way.** `tie-ordering-fence.unit.spec.ts` makes the rule
mechanical: a positional drill followed within its window by an assertion naming a seeded **person** is
RED; `drillIntoContentContaining(name)` is exempt; `lib/` is out of scope because the page objects OWN
the positional accessor. It **fails closed** on 0 specs or 0 sanctioned calls, and carries **its RED
watch inline** — the recogniser is asserted to fire on the exact pre-fix shape *and* to stay silent on
the sanctioned one, so "0 violations" cannot quietly mean "matches nothing". Measured: **30 specs, 0
positional drills, 2 property drills**; re-introducing the pre-fix shape turns it RED at
`negative-controls.spec.ts:577`.

**Stated limitation (§5 rule 8 — a skip must not read as a pass).** The `playthroughs/e2e` dir has **no
`node_modules` on this host**, so the fence was validated by running its scan logic under plain `node`
against the real tree, and by the same mutation. **It has not yet been executed by the Playwright
runner**, and the `negative-controls.spec.ts` edit has not been executed live. Both are covered by the
next Playthrough run; recorded here rather than left to be discovered.

**Flakes stabilized:** none observed. **Flake gate: 3 consecutive clean runs** of every test added
across passes 7–9.

**Final suite state (both repos, all sections):**

| section | result | vs baseline |
|---|---|---|
| `stack-core` | 14F of **420** | baseline 14F of 415 — **unchanged**, +5 tests |
| `stack-verify` | 11F + 1E of 237 | **unchanged** |
| `demo-stack` | 7F of 1030 | **unchanged** |
| `stack-injection` | **OK 299** (1 skip) | unchanged |
| `dev-stack` | **OK 138** | unchanged |
| Go: stack-seeding / stack-snapshot / stack-secrets / alignment / playthroughs | all **green** | unchanged |
| Go: clerkenstein | **environmental** — private-module fetch, no GitHub creds in this sandbox. Pre-existing | — |

No pre-existing failure was fixed and none was added. **Session total: 21 mutants, 21/21 matching
declared expectation** (11 claim-twin, 5 role-tenancy, 5 tie-ordering incl. its inline watch), with a
surviving no-op control in each battery.

**The perishable fixture is intact** across all three passes, and **none of the 18 live corpus blockers
was repaired** — TOK-02 step 4 still owns that, fence-assisted.

**Stop condition:** cap reached without stabilization — 3 incremental passes fired (7, 8, 9) and pass 9
still surfaced a live unrepaired twin on its first look, so the dimension scan is not clean. What
remains:

* **iters 27–30, 32–34, 36–41 not scanned to completion.** The seeder work (feedback hero policy,
  hiring funnel, assignment plans) and the iters 36–41 page-object changes are unswept.
* **`CHECK-M257x-iter35-seeder-writes-one-instant` is still open and is the ROOT of what pass 9 fixed.**
  Pass 9 repaired the two *selectors*; the seeder still stamps every backdated session with a single
  timestamp, which flattens **all** recency ordering in the product, not just these tests'. Every
  future assertion about "most recent" anything rests on the same tie. This is a believability defect
  as much as a test one and it is the highest-value item left in the queue.
* **The prior routed-forward queue is unchanged and now one longer.** Still open: **RF-2**, **RF-3**,
  **RF-7**, **RF-8**, **RF-9**, **RF-10**, **RF-11**, **RF-12**, and **RF-13** (new this session, pass
  8 — the tenancy derivation block that stops one step short of the number in dispute; **joins** the
  queue, supersedes nothing).

---

## Pass 10 — 2026-08-03 — incremental

**Iters hardened this pass:** iter-44 … iter-56 (13 iters: 11 tiks + 2 toks — iter-51 TOK-03, iter-54
TOK-04). First pass of the fourth invocation; the prior pass-9 terminating commit is `87a8948`.

**Tiks covered since prior pass:** 11.

**Sequencing note.** TOK-04 held this invocation back until *after* the `app` pin advanced
`v1.363.2 → v1.365.0`, on the reasoning that the residue is Playthrough- and seeder-heavy and hardening
against a version about to be replaced hardens the wrong thing. The advance landed at iter-56 and
clauses 1 and 2 were restored green against it, so this pass hardens the post-advance state. The risk
that motivated the wait — the advance breaking the seeders, as it did at v2.1 and v2.7 — did **not**
materialise.

**Coverage delta on touched files:**

| subject | before | after |
|---|---|---|
| `demo-stack/rosetta-demo` cmd_down teardown ORDER | 0 executable tests (8 source-string tests on the sweep) | **10**, the shipped block EXECUTED |
| `stack-injection/platform_topology.py` volume-syntax coverage | short syntax only | **+5**, long syntax refused |
| `stack-core` mutation batteries (M255 + M220) | **13 of 26 failures**, unattributable | **0**, all 21 mutants RED under distinct signatures |
| `corpus/ops/demo/demo-up-defaults.md` ↔ parsers | 27 disagreements, guard RED | **0**, guard OK both directions |
| `playthroughs` baseline-settle class | 0 | **3** (fence + 2 floors), pre-fix watched RED |

**The `stack-core` baseline this pass owns establishing.** Quoted at `14F/527` and unverified for four
rounds. **Measured: 26F of 585.** Repaired to **1F of 585** (below). The one that remains is
`test_claim_twin_guard_iter48_answer_key::test_02_the_green_twin_of_every_site_stays_SILENT` — a LIVE
corpus-state red against a deliberate answer-key fixture. It belongs to TOK-02 step 4's repair queue and
was **not** touched; the fixture is perishable and this pass did not spend it.

**Tests added:** 26 net-new across 4 files + 1 net-new guard.

- `demo-stack/tests/test_teardown_purge_order_m257x.py` — 10
- `stack-injection/tests/test_platform_topology.py` — 5
- `playthroughs/e2e/tests/baseline-settle-fence.unit.spec.ts` — 3
- (Pass 11: `stack-core/tests/test_evidence_visibility_guard.py` — 14)

**Bugs surfaced + fixed inline:**

1. **The teardown's fix stopped at the diagnosis; the write it was silent about still happened** (rext
   `458a9a4`). iter-55 made `sweep_project_containers` name the survivors and set `purge_failed`, then
   ran `purge_data_dir` on the very next line, unconditionally — so on the exact branch the fix exists
   to describe, `$stack/data` was still `rm -rf`'d from a root container **under a live postgres**. That
   is verbatim the damage iter-55's own commit message calls *"worse than the failure"*. The wipe is now
   GATED on the sweep; F-9 is preserved (no mid-flight death, slot and images still reclaimed, re-raise
   last), and the final `die` distinguishes *"the wipe failed"* from *"the containers are still up"*.
   The eight iter-55 tests could not have caught it: they are source-string assertions, and **two
   statements and a gate contain the same two names**. Hence 10 tests that EXTRACT and EXECUTE the
   shipped block against stubs, with an INVERTED mutant (a removal mutant cannot tell a gate from an
   inverted one), the literal pre-fix two-statement form, and a no-op control that survives.

2. **A volume spelling the parser cannot read graded GREEN** (rext `8e8ef5c`). iter-56's
   `check-host-mounts` reads compose's SHORT volume syntax. Against the LONG syntax the mount is silently
   dropped, and `check_host_mounts` grades the list it collected — so a file written that way yields zero
   findings and exits **0**, over exactly the missing-source state iter-56 measured costing a cold cycle
   and a misattributed `STORAGE_RPC_ADDR` diagnosis. Now a `TopologyError` naming the service. Latent,
   not live — which is the footing `INJECT_SVCS` was on the day before the prune.

3. **The anti-theatre battery ran on a runner this toolchain does not ship** (rext `22c1da8`, `bb474b7`).
   Both mutation batteries spawn `python -m pytest`; nothing else in rext needs pytest and this host has
   none. Every nested run exits 1 with no `FAILED` lines, so `test_00` fails with an EMPTY failure set —
   and, the part worth recording, **`test_01`'s PRIMARY assert (`assertNotEqual(rc, 0, "THEATRE: mutant
   left the suite GREEN")`) PASSED for all eleven mutants**, because rc was 1 whatever the mutation did.
   The battery would have certified *"every mutant went RED"* on a host where it never executed one test.
   It did not — solely because M255 wrote a SECOND assert demanding a **named** failing test rather than
   a non-zero exit code. That is §5 rule 7 doing its job, and it is the whole argument for the rule.
   Repaired in **both** batteries in one pass (§8 rule 7's recurrence corollary): M220's carried the
   identical dependency and its retry ladder was re-taking, three times at 900 s each, a measurement that
   could not come out differently.

4. **iter-55 added a module dependency the M255 battery's explicit list never learned** (rext `22c1da8`).
   `gen_injected_override.py` gained `import platform_topology`; `_COPY_FILES` did not. `_stage` asserts
   every NAMED file exists, which **cannot notice a dependency that was ADDED** — so the staged tree
   imported a module that was not there. Invisible for two reasons at once: the pytest-less runner could
   only report "RED, no attribution". Changing the runner named it on the first run.

5. **iter-56 shipped a knob no reader could find, and 22 anchors had rotted** (rosetta `d7440e8`).
   `demo_knob_guard` — the fence for exactly this — was RED and unread, carrying ten of the 26 stack-core
   failures. `DEMO_ALLOW_MISSING_HOST_MOUNTS` (up-injected.sh:2136) had no row anywhere, and it gates the
   only FATAL member of the host pre-flight: the one knob an operator reaches for when the pre-flight
   blocks them was the one the contract did not mention. Plus 22 stale `file:line` citations (iters 55–56
   added ~90 lines above them) and four count mirrors at 30-vs-31.

6. **The `pt-assignment-assign` flake, root-caused rather than re-run** (rext `8eb1fb0`). 29/1 then 30/0
   on an unchanged re-run. The spec read its baseline count from a members table that was **still
   filling**, and asserted a strict delta against a settled one. `waitForMembersTableSettled()` already
   existed — written at M256 iter-13 from a real trace whose own docstring measures *"2.2 s after the
   first row appeared"* — but its first caller ran AFTER the baseline. A short baseline makes `before - 1`
   unreachable and the Playthrough reports RED about a write that landed. iter-35's shape exactly: the
   rule applied at the ACTION, not at the MEASUREMENT one statement earlier. Fixed, and the CLASS fenced.

**Knowledge backfill:** `corpus/ops/demo/demo-up-defaults.md` (the missing knob row, stating the
measurement — no mount → starts; docker's auto-created empty DIRECTORY → exit 0 in 137 ms; a regular
empty file → starts and stays up — plus 22 re-pointed anchors), `corpus/ops/demo/README.md` and root
`CLAUDE.md` (count mirrors), and `rosetta-extensions/stack-core/README.md` (the guard registry row, Pass
11).

**Flakes stabilized:** 1 — `pt-assignment-assign`, root-caused to the pre-settle baseline read (above),
not papered over with a re-run. **Flake gate:** the tests added in this pass were run 3× consecutively
clean; the two mutation batteries were each run to completion twice after repair.

**Suite state (both repos):**

| section | result | vs baseline |
|---|---|---|
| `stack-core` | **1F of 585** | baseline was quoted `14F/527`, **measured 26F/585** — repaired to 1F |
| `demo-stack` | 7F of **1048** | baseline 7F/1038 — **unchanged**, +10 tests |
| `stack-injection` | **OK 331** | 316 (pre-iter-56) + 10 (iter-56) + 5 mine |
| `dev-stack` | **OK 138** | unchanged |
| `stack-verify` | not re-run this pass | baseline 11F+1E/237 |
| Go sections | unchanged | `clerkenstein` still environmental (no GitHub creds in this sandbox) |

shellcheck clean on `rosetta-demo`.

**Stop condition:** continue-to-next-pass — the dimension scan surfaced six live defects on its first
look, which is not a stabilizing signal. Still unswept: iters 27–30, 32–34, 36–41; `stack-verify` not
re-run; and `CHECK-M257x-iter35-seeder-writes-one-instant` (the seeder stamping every backdated session
at one timestamp) remains the root under the tie-ordering repairs and is untouched.

---

## Pass 11 — 2026-08-03 — incremental

**Iters hardened this pass:** iter-53, iter-54, iter-56 (the instrument-provenance thread) — the routed
`CHECK-M257x-iter54-gitignored-instrument-sweep`.

**Coverage delta on touched files:** `knowledge/plan/**/evidence/**` visibility: **0 checks → 1 guard +
14 tests**. The class had no instrument at all.

**Tests added:** 14 (`stack-core/tests/test_evidence_visibility_guard.py`).

**Bugs surfaced + fixed inline:**

1. **The sweep, taken — and `.agentspace/rext.tag` was not the only one** (rext `04f72af`, rosetta
   `04f59b1`). TOK-04 P2's corollary said to *assume it was not the only one until measured*. Measured:

   ```
   evidence/pt-run-iter26.log                 <- .gitignore  *.log
   iter-36/evidence/binding-run-report.json   <- .gitignore  knowledge/plan/**/*-report.json
   ```

   Both sat in this milestone's own evidence directories, on disk, in no diff. **The second is the run
   artifact of iter-36/37 — the clause-2 GATE-MEETING run TOK-04 re-opened for recording no platform
   ref.** Its evidence was never in the repository to re-read: the same defect one layer down. Neither
   ignore rule was wrong for its own purpose; what neither anticipated is that a run log and a Playwright
   report ARE the two artifacts a reading produces. iter-56 hit the identical trap two days later and
   worked around it with `.txt`, which fixes one commit and leaves the trap armed — which is why this is
   a guard, not a `.gitignore` edit. `evidence/` under `knowledge/plan` now always ships; a stray report
   in an iter-dir ROOT stays ignored, and a test pins that.

2. **The guard's own first cut was 4-of-5 false positive, and was narrowed before shipping.** Matching
   any path-shaped string containing `evidence/` fired on iter-39's plan sentence (*intent*, not a
   citation), on iter-56's decisions.md **quoting the two ignored filenames inside the finding that they
   were ignored**, and on two roadmap link LABELS whose targets resolve. Recogniser narrowed to markdown
   LINK TARGETS — someone wrote those expecting a reader to follow them. §8 rule 6, and one of the 14
   tests pins exactly those three shapes as NOT citations. Recorded here rather than quietly fixed,
   because a fence that cries wolf is disabled within a week and this one nearly shipped as one.

**Knowledge backfill:** `stack-core/README.md` guard-registry row; the `.gitignore` block carries the
measurement and its two offenders inline, so the next reader does not re-derive it.

**Flakes stabilized:** none observed.

**Stop condition:** continue-to-next-pass — the routed CHECK is closed and the class is fenced, but the
unswept-iter residue (27–30, 32–34, 36–41) and `stack-verify` are untouched, so the dimension scan is
not clean.

---

## Pass 12 — 2026-08-03 — incremental

**Iters hardened this pass:** the RECURRENCE dimension across iters 44–56 — every class this invocation
found, swept for its twin rather than closed at its site (§5 rule 19, §8 rule 7's corollary).

**Coverage delta on touched files:**

| subject | before | after |
|---|---|---|
| `dev-stack/dev-stack` teardown | 0 | **13** (incl. a 4-assert PAIR FENCE against the demo twin) |
| `dev-stack` verify scope | a 12-name hand tuple, 7 of them deleted services | derived |

**Tests added:** 13 (`dev-stack/tests/test_dev_teardown_sweep_m257x.py`).

**Bugs surfaced + fixed inline:**

1. **The dev teardown had the demo teardown's defect, and worse** (rext `18399c5`). Pass 10 fixed the
   demo purge-order; this pass asked where else the class lived. `dev-stack cmd_down` was
   `docker compose -p "dev-$n" down 2>/dev/null || true` followed by an unconditional
   `rm -rf "$STACKS_DIR/dev-$n"`. Three aggravations over the demo original: compose's stderr goes to
   **/dev/null**, so the "invalid compose project" line — the *only* signal on the demo side — is
   DISCARDED rather than swallowed; `--purge` removes the **entire** stack dir (data, env, override,
   certs), not just data; and `reg_release` frees the N while surviving containers still hold their
   ports, so the next `dev-stack up N` dies on a bind error with no clue. A `dev-N` override is generated
   from the same platform compose and goes stale the same way, so the refusal is reachable on this path
   **today**. Same remedy + the purge GATED on the sweep.

2. **`dev-stack`'s verify scope was still the hand tuple** — `postgresql redis sentinel backend skiller
   skillpath jobsimulation cms storage roadrunner graphql gotenberg`, **seven of twelve deleted**
   (skiller v2.1, skillpath M507, graphql 2adcf71, cms/jobsimulation/roadrunner ef32d4c, storage
   0dab54d). iter-55 derived the identical tuple away on the demo side and did not carry it here.

3. **`down${purge:+ (purged)}`** — `purge` holds the STRING `0` when the flag is absent and `:+` treats
   that as set, so **every** dev teardown has announced itself as *"(purged)"*. Now a function of what
   happened; a refused purge returns non-zero.

**The duplication is FENCED, not hidden.** Two byte-parallel sweeps in two rext sections is exactly §2's
defect. Extracting a shared shell library across sections is larger than a harden pass should make, so
`TestTheTwinsAgree` asserts both twins ask for the exact project label, RE-READ after removing, and fail
loud on a survivor — with a non-vacuity self-test on a mutated copy.

**Recorded because it is the class, and this pass walked into it too:** the new call-site assertions
FAILED on their first run, matching `rm -rf` and `${purge:+` inside the **comments** that quote the
pre-fix source to explain it. Comments are stripped before matching now (§8 rule 6), and a test pins that.

**Also swept, and CLEAN:** the pytest-dependency class has no remaining *executable* instance in rext —
the two batteries were the only ones; every other mention is a docstring or a doc line.

**Flakes stabilized:** none new. **Flake gate: 3 consecutive clean runs** of every test file added across
passes 10–12, plus 3 clean runs of the TS fence's shipped scan functions under node.

**Final suite state (both repos, all sections):**

| section | result | vs baseline |
|---|---|---|
| `stack-core` | **1F of 599** | baseline quoted `14F/527`; **measured 26F/585**; repaired to 1F, +14 tests |
| `demo-stack` | 7F of **1048** | 7F/1038 — **unchanged**, +10 tests |
| `stack-injection` | **OK 331** | unchanged + 5 |
| `dev-stack` | **OK 151** | OK 138 — **unchanged**, +13 tests |
| `stack-verify` | 11F + 1E of 237 | **unchanged** |
| Go sections | green | `clerkenstein` environmental (no GitHub creds in this sandbox), unchanged |

`shellcheck` clean on both edited shell scripts. **The live `demo-1` stack (11 containers) carrying
clauses 1 and 2's evidence was NOT touched** — verified up at the end of the session.

**Session totals:** 3 passes · 6 rext commits + 3 rosetta commits · **53 tests added** across 5 files +
1 net-new guard · 9 defects fixed inline · 1 flake root-caused · 1 routed CHECK closed.

**Stop condition:** cap reached without stabilization — 3 incremental passes fired (10, 11, 12) and pass
12 surfaced three more live defects on its first look, so the dimension scan is not clean.

### Does the third `HARDEN-CAP-ACCEPTED`'s conclusion still hold?

It was recorded as: *"the residue needs AST / call-site assertions, not another sweep of the same shape."*
**No — not in that form, and this invocation is the counter-evidence.**

Three defects here were found by *executing something that had never been executed on this host* (the
`stack-core` suite, quoted at `14F/527` and unmeasured for four rounds — it was **26F/585**), and three
more by *following a fix to its twin instead of closing at its site*. Neither needed a sharper
instrument. What the prior three passes had in common was the **surface**: all three swept corpus prose,
where the fixed point genuinely had been reached. The conclusion generalized from that surface to the
milestone.

What remains is a **coverage** gap, not a dimensional one, and it is still the same list:

* **iters 27–30, 32–34, 36–41 remain unscanned** — the seeder work (feedback hero policy, hiring funnel,
  assignment plans) and the iters 36–41 page-object changes. Now six passes old.
* **`CHECK-M257x-iter35-seeder-writes-one-instant` is still open** and is still the ROOT under every
  tie-ordering repair: the seeder stamps every backdated session at ONE timestamp, so "most recent"
  orders nothing anywhere in the product. Pass 9 called it the highest-value item left and it is
  untouched — it is a believability defect as much as a test one, and it is **Fate 3** (a distribution
  change to seeded data, on a box whose demo-1 stack is live gate evidence).
* **The routed-forward queue:** RF-2, RF-3, RF-7…RF-13 unchanged. **RF-14 is new this session** — the two
  byte-parallel container sweeps in `demo-stack` and `dev-stack` want one shared shell primitive; the
  pair fence holds them together but does not merge them. **Joins** the queue; supersedes nothing.
* **`stack-core`'s single remaining red** (`test_claim_twin_guard_iter48_answer_key::test_02`) is a LIVE
  corpus-state failure against a perishable answer-key fixture. It belongs to TOK-02 step 4 and was
  deliberately not spent.

---

## Pass 13 — 2026-08-03 — incremental

**Iters hardened this pass:** iter-27, iter-30 — the first two of the **unscanned window** (iters 27–30,
32–34, 36–41), which had been named as the outstanding coverage gap for six passes.

**The hypothesis under test.** The 4th cap was declined on the grounds that the residue is a COVERAGE gap,
not a dimensional one — that the unscanned window would yield defects at a rate comparable to fresh
material, rather than needing a sharper instrument. **It held.** Two iters into the window, on their first
look: **three findings, two of them live defects**, both in the *same function* iter-30 shipped, and both
of a class this milestone has already paid for twice.

**Coverage delta on touched files:**

| subject | before | after |
|---|---|---|
| `write_run_provenance` crash path (binding run, no results) | **0** — the iter-30 helper always pre-created the artifacts, so the branch never ran | 1 executable pair-invariant test + a survive-control |
| `write_run_provenance` JSON validity | 0 | 7 sub-cases, parsed + round-tripped |
| hero-share fence "must carry a reason" clause | 1 of 6 files (`feedback.go` only) | **6 of 6**, derived |

**Tests added:** 3 (`playthroughs/manifest/runner_safety_test.go` ×2 executable,
`stack-seeding/seeders/hero_share_policy_fence_test.go` ×1 AST-derived).

**Bugs surfaced + fixed inline** (rext `ff67f1c`):

1. **The guard and the action were not connected — the pass-10 class, again.** `write_run_provenance()`
   made both RESULT copies conditional on the artifacts existing and left the PROVENANCE copy
   **unconditional** beside them. The runner `rm -f`s `last-run.json` before Playwright starts, so a
   **crashed binding run** reaches the function with the file ABSENT and advanced
   `last-binding-run.provenance.json` to the new run while `last-binding-run.json` still held the old one.
   Measured: a sidecar reading `run_start_epoch: 2000, playwright_exit: 1` beside results from epoch 1000.
   **This is iter-30's own defect rotated.** It shipped so a reader could tell a binding verdict from an
   advisory probe; in this state a reader cannot tell a **CURRENT** binding verdict from a **STALE** one —
   and the fresh timestamp actively vouches for the old numbers, which is worse than no sidecar at all.
2. **The sidecar could be unparseable.** `grep_pattern` is operator argv and was interpolated RAW, so
   `--grep '@pt:a"b'` emitted `"grep_pattern": "@pt:a"b"` — rejected by every JSON parser, at READ time,
   far from the run that wrote it. A scoped diagnostic is precisely when an operator reaches for a quoted
   pattern, so it is reachable on the path the file exists for. Escaped **backslash-first** + control chars.
3. **iter-27's fence left its own clause unenforced.** The policy map's doc says `heroIndifferent` "must
   carry a reason"; only `feedback.go`'s reason was ever checked. A new seeder could satisfy every
   assertion with `{heroIndifferent, ""}` — **accidental indifference re-admitted through the front door**,
   the one state that fence exists to prevent.

**Twin sweep (§5 rule 19) — CLEAN, and recorded as such.** The
conditional-copy-beside-unconditional-sidecar shape has no other instance in rext; `last-binding` appears
in exactly this script and its guard. Pass 12's finding was that fixes travel in pairs; here the pair
genuinely does not exist, which is worth writing down so the next pass does not re-run the search.

**Mutation results — 6 mutants, all RED, four INVERTED rather than removals (§8 rule 5):**

| mutant | verdict |
|---|---|
| M1 unconditional sidecar (the original defect restored) | RED |
| **M2 sidecar never copied (freeze)** | **RED — caught by the positive control** |
| M3 raw JSON interpolation | RED |
| **M4 escape order inverted (quote before backslash)** | **RED — pins the ordering claim itself** |
| M5 reason blanked · M6 reason whitespace-only | RED |

**M2 is the one that matters.** The cheapest way to pass a staleness assert is to stop writing the sidecar
entirely, so the control asserts a later *healthy* binding run still advances **both** artifacts — the
no-op positive control that must SURVIVE. Without it, every assertion in that test would be green against
a function that does nothing. M4 is the second: the fix's own comment claims backslash must be escaped
before the quote, and that claim is now tested rather than asserted in prose.

**Flakes stabilized:** none new. **Flake gate: 3 consecutive clean runs** of all 3 added tests.

**Suites:** `playthroughs` green (4 pkgs), `stack-seeding` green (12 pkgs), `go vet` clean both,
`shellcheck` clean on the edited script. The live `demo-1` stack (11 containers) carrying the clause-1/2
evidence was **not touched**.

**Stop condition:** continue-to-next-pass — the window's remaining code carriers (**iter-36**'s 262-line
`assignment_plans.go` + `hiring_funnel`/`assignments` deltas, and **iter-37**'s `stack-injection` override
generator) are unscanned, as are the measurement iters 28/29/32/33/34/38/39/41 and the iter-40 cleanup.

---

## Pass 14 — 2026-08-03 — incremental

**Iters hardened this pass:** iter-36, iter-37 — the unscanned window's remaining CODE carriers
(iter-36's 262-line `assignment_plans.go` + the `assignments`/`hiring_funnel` deltas; iter-37's
`stack-injection` override generator).

**No live defect this pass — both are correct TODAY.** What was missing is the ability to NOTICE if they
stopped being, and in both cases the gap has one shape: **a property asserted at one of its two sites, or
asserted by its spelling rather than by its behaviour.**

**Coverage delta on touched files:**

| subject | before | after |
|---|---|---|
| plan-model FK write ORDER, hiring writer | **0** (fenced for the generic writer only) | shared helper, both writers |
| skiller Azure fallback RESOLUTION | 0 — spelling only | 4 env scenarios × 2 vars, resolved by a real shell |

**Tests added:** 2 (+1 shared helper extracted).

**Findings:**

1. **The hiring writer had no write-order fence.** iter-36 materializes the M7 plan model from **two**
   code paths — the `AssignmentsSeeder` tail and `hiringFunnelRows.flush`. The plan model is the FK
   **parent** of every assignment's four new FKs, so it must be COPIED first, and only the generic path
   was fenced for it. **Measured** by moving the hiring flush's plan block after its assignments copy:

   | test | verdict on the broken ordering |
   |---|---|
   | `TestHiringFunnelSeeder_PlanMaterialized` (rows fence) | **PASS — blind to it** |
   | `TestPlanModelWriteOrderIsFKSafe` (generic writer) | **PASS — does not cover the twin** |
   | the new twin fence | **FAIL** |

   The rows fence and the order fence are independent properties: the first proves the FKs point at rows
   the same run wrote, and stays green with those rows copied *afterwards* — which Postgres rejects
   outright. The order assert is now a shared helper both writers call.

2. **The skiller Azure fallback was asserted by SPELLING.** iter-37's two tests check that the emitted
   line CONTAINS `${SKILLER_…:-` and `${AZURE_…:-}`. A concatenating variant,
   `${SKILLER_AZURE_OPENAI_KEY:-}${AZURE_OPENAI_KEY:-}`, contains both and **passes both** — while
   resolving to `dedicated-real-keydemo-key` when both are set, i.e. **corrupting an operator's real
   dedicated production credential by appending the demo one to it**. That is precisely what the secret
   DNA's DISTINCT-SIMILAR rule exists to prevent, and it was the one claim in iter-37's own docstring that
   nothing tested. The new test resolves the expression with a real shell under four scenarios.

**Mutation results — 4 mutants on the fallback, all RED on the new test; two of them INVISIBLE to the
pre-existing pair:**

| mutant | new resolution test | pre-existing iter-37 tests |
|---|---|---|
| M8 concatenating (corrupts a real key) | RED | **still GREEN** |
| **M9 INVERTED precedence (shared wins)** | RED | **still GREEN** |
| M10 fallback dropped | RED | RED |
| M11 `:-` → `:+` (inverted operator) | RED | RED |

**M9 justifies the test on its own** — the demo's shared key silently overriding a real dedicated one is
invisible to every spelling assert. **M10 is what the positive control is aimed at:** dropping the
fallback satisfies "the operator's value wins" perfectly, so the only thing making the feature real is the
assert that the shared pair IS used when it is alone.

**A METHOD NOTE WORTH KEEPING (this pass nearly recorded a false RED).** The first mutation run was scored
with `python3 -m unittest discover -k skiller`, which reported `FAILED (errors=2)` — read as "caught". The
**pristine** tree reports the identical `FAILED (errors=2)` under that filter: the errors were unrelated
and the rc was 1 either way. This is the pass-10 pytest finding exactly (`rc 1 regardless`), arrived at
from the other direction, and the only thing that surfaced it was **baselining the mutation harness on
unmutated source before trusting a single verdict**. Every verdict in this ledger's mutation tables is now
scored per-test, by name, against a measured pristine baseline.

**Also checked and CLEAN** (recorded so a later pass need not redo them): `orgLess` members never reach
`pm.attach()` — the guard precedes it, so no member gains an enrollment in an org she is not in;
`membershipUUID` returns a `string`, so the enrollment's `%v` `membership_id` is identity today (a real
fragility only if that return type ever changes); iter-37's `if name == "backend"` gate names the service
that actually hosts the skills domain, so its guard and action are connected.

**Flakes stabilized:** none new. **Flake gate: 3 consecutive clean runs** of both added tests.

**Suites:** `stack-injection` **OK 332** (was 331), `stack-seeding` seeders green, `go vet` clean. The live
`demo-1` stack was not touched.

**Stop condition:** continue-to-next-pass — the window's MEASUREMENT iters (28, 29, 32, 33, 34, 38, 39, 41)
and the iter-40 cleanup are still unscanned, and pass 13's finding rate makes a recurrence sweep across
them worth one more pass.

---

## Pass 15 — 2026-08-03 — incremental

**Iters hardened this pass:** the window's MEASUREMENT iters (28, 29, 32, 33, 34, 38, 39, 41) + the
iter-40 cleanup — **and the RECURRENCE dimension for every class passes 13–14 found** (§5 rule 19, §8
rule 7's corollary).

**With this pass the unscanned window is CLOSED.** iters 27–30, 32–34, 36–41 — the gap named in four
consecutive cap-acceptance discussions and six passes old — is fully swept: 27/30 at pass 13, 36/37 at
pass 14, and the remainder here.

**What these iters are, and what that implies.** iters 28/29 touched **evidence only**; 32/33/34/38/39/41
are **corpus prose**; 40 is a skills-doc cleanup. This is the surface the third cap-acceptance correctly
identified as having reached its fixed point — so the dimension worth spending here is not the prose but
the **GUARDS that fence it**, which are code and can be blind (pass 12's `check-host-mounts` graded GREEN
over a syntax it could not parse).

**Findings: none. Every sweep a clean negative — established empirically, not by reading.**

1. **Guard vacuity (the pass-12 shape-4 class).** A crude grep suggested seven `stack-core` guards had no
   vacuity protection. That was the grep being wrong, and running them proved it. Each was pointed at (a)
   a missing tree, (b) a present-but-empty corpus, and (c) a corpus with zero matching files:

   | guard | missing tree | present-but-empty | zero files |
   |---|---|---|---|
   | `anchor_construct_guard` | exit 2 | **exit 2** — *"0 anchor(s) resolved; the resolver, not the corpus, is what this measures"* | — |
   | `derived_value_guard` | exit 2 | **exit 2** | — |
   | `markdown_structure_guard` | exit 2 | exit 0, honestly reporting *"scanned 2"* | **exit 2** — *"0 files to scan"* |

   All fail **closed**. `markdown_structure_guard`'s exit 0 on a 2-file clean corpus is correct behaviour
   (it is a per-file structural check) and it *states its denominator*, which is the property that matters.

2. **Raw JSON interpolation (pass 13's second defect), swept across ALL of rext.** Every `.sh` line
   emitting a JSON string field from a bare shell expansion: **2 hits, both the ones pass 13 fixed.** Every
   hand-built JSON string field in non-test Go/Python: **1 hit** —
   `ai_readiness_funnel.go:680`, which **already escapes** `"` and `\` correctly for its stated surface
   (code-owned constant prompts) and says so. **No twin.**

**Tests added:** 0 — deliberately. Nothing surfaced that a test would pin, and a test written to justify a
pass is the thing this milestone punishes.

**Baselines RE-MEASURED this pass (all five python sections + Go), and all unchanged:**

| section | measured | vs baseline |
|---|---|---|
| `stack-core` | **1F of 599** | unchanged — the single red is still the iter-48 perishable answer-key fixture (TOK-02 step 4, deliberately not spent) |
| `demo-stack` | 7F of 1048 | unchanged |
| `dev-stack` | **OK 151** | unchanged |
| `stack-verify` | 11F + 1E of 237 | unchanged |
| `stack-injection` | **OK 332** | 331 + the 1 test pass 14 added |
| Go sections | green | `clerkenstein` environmental (no GitHub creds here), unchanged |

**Flakes stabilized:** none new.

**Session totals (passes 13–15):** 2 rext test/fix commits + 3 rosetta ledger commits · **5 test functions
added** across 3 files (pass 13: 3, pass 14: 2, pass 15: 0) **+ 2 helpers** (`provRun`,
`assertPlanModelWrittenBeforeAssignments`) · **2 live defects fixed inline** · **2 untested-property gaps
closed** · **12 mutants run, 12 RED**, of which **4 were invisible to the pre-existing tests**.

*(This line first read "6 tests", which 3 + 2 + 0 does not equal. Corrected in the same pass that wrote it
— an unchecked count in the ledger of a milestone about claims-versus-measurement is the defect class
itself, at the smallest possible scale.)*

**Stop condition: stabilized.** Both conditions are met and this is the first pass in five invocations
where that is true rather than asserted: the coverage delta this pass is **0** (no test was needed), and
the dimension scan — two recurrence sweeps run to completion against measured baselines — **found nothing
new**. The gap that motivated four declined cap-acceptances is closed.

**What "stabilized" does NOT mean here.** It is a statement about the harden dimension, not about the
milestone: **`CHECK-M257x-iter35-seeder-writes-one-instant` remains open and Fate 3** (a seeded-data
distribution change against a live evidence stack), the routed-forward queue **RF-2, RF-3, RF-7…RF-14** is
unchanged, and `stack-core`'s single red belongs to TOK-02. None of those is a hardening deficit; all of
them are correctly-routed work with named handlers.

### Did the coverage-gap hypothesis hold?

**Yes, decisively, and the shape of the answer matters more than the verdict.** The 4th cap was declined on
the argument that the residue was a coverage gap rather than a dimensional one — that fresh material would
yield defects without needing a sharper instrument. On first contact with the window, pass 13 found **two
live defects and one unenforced clause**, and pass 14 found **two properties fenced at one of their two
sites**, one of which (an inverted-precedence secret expression) would corrupt a real dedicated production
credential while passing every test that existed. **None of the five needed AST or call-site analysis** —
they needed someone to *run the function under the state it was written for* and to *follow a property to
its second site*.

And the window is now genuinely exhausted: pass 15 swept its remainder and both recurrence dimensions and
found **nothing**. The hypothesis held, and it has now been spent.

## Pass 16 — 2026-08-04 — incremental

**Iters hardened this pass:** iter-58, iter-59 (tok), iter-60 … iter-68
**Tiks covered since prior pass:** 10 (iters 58, 60–68; iter-59 was a tok)

**Scope.** The window that built the milestone's primary deliverable — the `platform_predicate_guard`
predicate fence (G1–G7, ~1 360 lines net-new), assertion F in `platform_alignment_guard`, and iter-68's
ref-awareness across three guards. `stack-core` 610 → 682+ tests before this pass.

### The dimension this pass ran: REACH vs CLASS, measured per assertion

iter-61 stated the milestone's sharpest rule — *a fence whose reach is narrower than its class
over-reports its own GREEN, invisibly, because the fence is what you'd check with* — and iter-67
demonstrated it. This pass **applied that rule to the fence family itself**, by enumerating every
assertion's class and measuring what fraction of it the assertion can actually read. That is a different
instrument from a test, and it found what tests had not.

**Every assertion's reach, live, before → after:**

| assertion | class | reach BEFORE | reach AFTER |
|---|---|---|---|
| G1 profile tokens | documented profile tokens | 99 sites / 8 tokens | unchanged |
| G2 repo-count | repo-count claims in clone context | 3 (2 ref-pinned) | unchanged |
| **G3 default profile** | rows marking a default | **0 of 3 — 0 %** | **3 of 3 — 100 %** |
| G4 RPC address | local-topology address claims | 13 (2 ref-pinned) | unchanged |
| G5 migration target | `migrations: true\|false` sites | 1 enumerated / 21 free prose / 2 pinned **of 24** | unchanged, partition now stated |
| G5b sole-migrator | `_ONLY_MIGRATOR` sites | 4 checked (5 sites) — **separate universe** | unchanged, no longer summed into G5 |
| **G6 mid-fold** | RPC vars graded | **"measured"** — no count at all | **7 graded, 0 mid-fold @ origin/main; 1 @ `b948604`** |
| **G7 profile membership** | membership rows | **12 of 22 — 54.5 %** | **21 of 22 — 95.5 %** |
| F (alignment) | citations in the map | 74 resolved = 20 subject-checked + 53 range-only + 1 unattributable | unchanged — partition already closes |
| anchor-construct | resolvable anchors | 124 graded, **ref unnamed** | 94 at a ref + **30 `worktree(fallback)`**, named |

**Bugs surfaced + fixed inline (5):**

1. **G3's reach was ZERO on the corpus it ships to guard** (`247b847`). `_DEFAULT_MARK` was
   `\(default\)` — a BARE parenthesis, and the corpus writes none. All three of its rows attach the
   evidence to the mark: `*(default — `PROFILE ?= core`)*` (CLAUDE.md:314, platform_repo.md:77) and
   `(the Makefile default — `PROFILE ?= core`)` (service_taxonomy.md:428). `documented_default_profiles`
   returned `[]`, so the `wrong-default` loop never ran and the `undocumented-default` arm — itself gated
   on `marked` being non-empty — could not fire either. **G3 reported GREEN by never looking, on every
   corpus, for as long as it has existed.** The synthetic fixtures used the bare spelling, so the whole
   pre-existing G3 suite passed: *the fixture agreed with the regex instead of with the corpus.*
2. **The reach LINE made the error this milestone exists to punish** (`247b847`). It read
   `24 migration claim(s) of which 1 enumerated + 4 sole-migrator checked and 21 free prose UNREACHED`,
   and **1 + 4 + 21 = 26**. `of which` is a partition claim and it did not close: `_ONLY_MIGRATOR` matches
   a different line set (5 live sites, 3 of which also match the migration universe and **2 lie entirely
   outside the denominator they were reported against**). A reader deriving *"G5 reaches 5 of 24"* got a
   number wrong in both numerator and denominator.
3. **The third ref-aware guard could not say which file it read** (`a2f29a2`). iter-68 made three guards
   ref-aware because *three checkers were reading the wrong copy of the code*. Two came out of it naming
   their provenance and refusing an unresolvable named ref. `anchor_construct_guard.read_target` returned
   bare lines and **fell through to `target.read_text()` — the CHECKOUT — silently**, adjudicating every
   anchor against it. A typo'd `CITE_REF`, or a clone with no git dir, and the guard graded the worktree
   and printed OK. Two more found while testing it: `ref: str = CITE_REF` was a DEFAULT ARGUMENT (bound at
   import) while `resolve()` reads the module global at call time — the resolving half and the reading
   half took their ref from two different places; and the UNMEASURED refusal had to be ordered BEFORE the
   resolver positive control, because a bad ref drives `resolved` to 0 and the control then blamed *"the
   resolver, not the corpus"* for a ref the operator typo'd.
4. **The cell-scope rule was built at iter-63 and never wired into G2/G4/G5/G5b** (`63af69a`).
   `_pin_window(lines, i, col)` narrows a table row's pin scope to the CLAIM'S OWN CELL — half of
   `D-M257x-63-1`, built for a measured reason. Only G1's prose path ever passed `col`. **Live impact 0**
   (of the 6 pin-exempted claims, one is on a table row and its pin shares the cell) — latent, not live,
   which is why a reading found it and no counter could.
5. **G7 could not read 45 % of its class, and the corpus was right in all of it** (`52fb3fd`). All 10
   misses were the parser: 6 rows state the membership and explain it after an em dash
   (`storage — the rollback path only; …`), 2 name the service in display case (`**Storage**`) against
   compose's `storage`. **It had to be a CUT, not a widening:** `CLAUDE.md:319` is *"next-web-app — **but
   selecting it alone exits 1**: `next-web-app` declares `depends_on: backend`"*, and harvesting the whole
   cell yields `{next-web-app, backend}` → `[G7 wrong-membership] … NOT STARTED ['backend']`, a fence
   inventing a claim out of an explanation. **12/22 → 21/22, corpus GREEN** — every row that was
   unreadable was also correct, which is exactly why the gap could sit there: nothing it hid was wrong yet.

**The iter-68 `lru_cache`, pinned in both directions** (`a2f29a2`). Nothing tested it. Measured: keying
HOLDS (the clone root is in the key); **freshness does NOT — a ref that moves in-process is not seen
again**. That is the RIGHT behaviour and it is load-bearing rather than incidental: `read_target` and
`resolve` re-resolve the ref *per citation*, so without the cache a fetch landing mid-run would split one
verdict across two trees while the report named one ref. **The cache is what makes the provenance line
true.** Documented so nobody "fixes" it into freshness; `cache_clear()` exercised as the sanctioned escape.

**Tests added:** 26 across 2 files — `test_platform_predicate_guard.py` 85 → 102, `test_iter45_mechanical_fences.py` 35 → 44.

**Mutation batteries: 27 inverted mutants, 27 RED** (A 6, B 8, C 5, D 4, E 4).

**The battery caught FIVE weak tests of mine before it caught anything else** — the fourth pass running
where that is the headline, and the mechanism working as designed:

* the G5 cell-scope fixture used `(currently: …)`, but `_MIGRATION_ENUM` accepts `(currently|now|today: …)`
  and `_ASSERTS_CURRENCY` matches `currently` — so an enum spelled that way asserts currency BY
  CONSTRUCTION and can never be pin-exempted at any scope. It proved the currency rule while claiming to
  prove the cell rule, and passed identically with the fix reverted. (Corollary kept:
  `ref_pinned_skipped_migration` can only ever count NON-enumerated claims.)
* G5b had no neighbouring-cell test at all.
* `test_mid_fold_count_agrees_with_the_histogram` ran with `app_root=None`, so BOTH sides were 0 — an
  identity, not an assertion — and a mutant pinning `mid_fold_count = 0` SURVIVED it.
* its completed-fold twin read an INVENTED variable name, which does not model a completed fold: it just
  relocates the mid-fold onto the new name.
* `test_prose_still_contributes_nothing` asserted emptiness at the tokenizer, which reads by SHAPE; the
  compose-name filter lives in the caller. It pinned a rule that does not live in that function.
* and the anchor end-to-end test asserted only the exit code, and passed through the WRONG branch of two
  that both exit 2.

**Knowledge backfill:** none to the corpus this pass — every finding was in `rosetta-extensions` guard
code, and the reach numbers are now emitted by the guards themselves rather than written down anywhere
that could go stale. (That is the point: a reach figure in prose is the class of claim this milestone
exists to distrust.)

**Flakes stabilized:** none new. The `dev-stack` nested-run interference (151 OK solo vs 6 spurious
`test_dev_public_host` failures beside `stack-core`) was re-confirmed as environmental and is unchanged.

**Stop condition: continue-to-next-pass** — the reach dimension is not exhausted. G5's attribution reach
is **1 of 24** and G1's prose class is the known clause-5 residual; and the recurrence question this pass
raises (*which OTHER guards compute a reach counter and never print it?*) has not been swept.

## Pass 17 — 2026-08-04 — incremental

**Iters hardened this pass:** iter-58 … iter-68 (recurrence dimension, whole-section scope)
**Tiks covered since prior pass:** 0 (same window as pass 16; a different instrument on it)

**The dimension: RECURRENCE of pass 16's class, across all 19 `stack-core` guards.** Pass 16 found
the same defect three times inside one guard — *a verdict sentence whose denominator is not the
quantity the sentence is about* (G3 reading 0 of its 3 rows, G7 12 of 22, G6 saying "measured" with
no count). Pass 15 had already swept every guard for FAIL-CLOSED behaviour; this pass swept the
sharper property.

**It recurs exactly once.** `derived_value_guard` printed `N service doc(s) measured … M unmeasured`
and then `OK — every checkable scalar matches its source`. The reach line counts DOCS; the verdict
is about SCALARS, and nothing counted scalars. A silent narrowing of `_DOC_GO` or `_TF_CPU` would
drop the graded set to one comparison and print an identical run — positive control included,
because that control guards docs rather than comparisons. Live: **5 of 29 docs measured, 6 scalars
graded**; the verdict now states both.

It was the near-miss rather than the disaster because its 24 unmeasured docs were already NAMED
individually with a reason (`no clone of X` vs `clone present, no scalar this guard reads`). The
other 17 guards state a denominator on the line their verdict is about and hold.

**The pre-existing ratchet caught the caller I missed, and failed CLOSED** — `repair_postcondition`
exited 2 with `derived_value_guard.postcondition_sites raised: ValueError('too many values to
unpack')` rather than dropping a fence. iter-44's post-condition doing its job on the first change
to touch it since.

**Tests added:** 2 (`test_iter45_mechanical_fences.py` 44 → 46). **Battery F: 3 mutants, 2 RED + 1
EXPECTED SURVIVOR.**

**The expected survivor is the finding.** My `if not scalars` refusal is **unreachable by
construction** — `measured` is only incremented under `checked_here`, which is only ever set beside
a `scalars += 1`. My first test drove it through `main()` and passed via the *doc-level* control:
the same two-exit-2-paths trap pass 16 hit in `anchor_construct_guard`. Rather than reshape the test
to pretend, the branch is kept as a documented fail-closed backstop, **declared unreachable today**,
and the test asserts the INVARIANT that makes it unreachable. A mutant removing unreachable code
cannot be caught, and claiming otherwise is the theatre this milestone exists to punish.

**Stop condition: continue-to-next-pass** — the reach/denominator dimension is now exhausted across
the section, but dimensions 2/3/5 (edge cases, error paths, fuzzing) have not been run on the
iter-60..68 parsers at all.

## Pass 18 — 2026-08-04 — incremental

**Iters hardened this pass:** iter-60 … iter-68 (the new parser surface)
**Tiks covered since prior pass:** 0 (same window, third instrument)

**The dimension: error paths + boundary fuzzing (2/3/5), untried by passes 16-17.** ~40 pathological
cells × ~13 pathological documents against the iter-60..68 parsers — empty, delimiter-only,
unbalanced backticks/emphasis, unbalanced and NESTED parentheses, a 5 000-character cell, emoji,
tabs, NULs, CRLF, lowercase and emphasised headers, the dash lookalikes (‒ ― –), leading/trailing
separators.

**Zero raises, every reach difference correct.** Nothing needed changing, and that is the result:
fencing what already holds is as much the point as breaking it. CRLF, lowercase and emphasised
headers are asserted to FIND their row rather than merely not to crash — *"it did not crash"* is the
weakest fuzz assertion there is and passes against a parser that returns nothing.

**Two boundary facts recorded rather than fixed, because measuring said not to:**

1. `_pin_window` raises `IndexError` on out-of-range `i`. Every production caller is a
   `for i, line in enumerate(lines, 1)`, so `i` is in range by construction — defending it would
   MASK a caller bug rather than prevent one. Contract asserted instead.
2. **A UTF-8 BOM before a table header hides that whole table** from G1, G3 and G7 at once.
   Measured: **0 of 112 scanned files carry a BOM, 0 open with a table row** — the shape needs both
   coincidences. Patching an unobserved failure mode on speculation is the habit this milestone
   distrusts. The measurement is the deliverable; the test is a TRIPWIRE that fails if BOM handling
   ever silently changes.

**Tests added:** 6 (`test_platform_predicate_guard.py` 102 → 108).

**Session totals (passes 16–18):** 7 rext commits + 3 rosetta ledger commits · **34 tests added**
across 2 files (`test_platform_predicate_guard.py` 85 → 108, `test_iter45_mechanical_fences.py`
35 → 46) · **6 live defects fixed inline** · **30 mutants run, 29 RED + 1 declared-unreachable
survivor**.

**Seven weak tests of MINE were caught by the battery before it caught anything else** — more than
in any prior pass, and three of them share one root cause worth naming: **asserting a rule at the
layer that does not hold it.** `_cell_service_tokens` reads by SHAPE and the compose-name filter
lives in its CALLER, so `everything`, a 5 000-character run of `a`, and prose all come back as
tokens and are harmless only downstream. Three first drafts pinned emptiness at the tokenizer. The
other four: a fixture using `(currently: …)` where the construct asserts currency by definition and
can never be pin-exempted at any scope; a G5b case that did not exist; a mid-fold assertion whose
two sides were both zero (an identity, not an assertion); and two tests that asserted an exit code
where TWO paths return the same one.

**Flakes stabilized:** none new. `dev-stack`'s nested-run interference re-confirmed environmental
(151 OK solo; 6 spurious `test_dev_public_host` failures when run beside `stack-core`, whose m220
battery spawns nested `dev-stack` runs).

**Baselines re-measured, all unchanged:** `demo-stack` 7F/1048 · `stack-injection` OK 332 ·
`stack-verify` 11F+1E/237 · `stack-core` 1F (the perishable iter-48 answer-key fixture, TOK-02
step 4, deliberately not spent).

**Knowledge backfill:** none to the corpus. Every finding was in `rosetta-extensions` guard code,
and each reach figure is now EMITTED BY THE GUARD rather than written down somewhere that can go
stale — which is the point: a reach number in prose is exactly the class of claim this milestone
exists to distrust.

**Stop condition: cap reached without stabilization — the 3-pass incremental cap fired.** Coverage
delta did not fall below 2 %: pass 18 alone added 6 tests to a file that started the session at 85,
and each of the three passes ran a DIFFERENT dimension and each found something (reach → recurrence
→ fuzz). The dimension scan is not "found nothing new"; it is "has not run out of dimensions".

**This is cadence, not test debt** — and the distinction is measurable rather than asserted. The
window under test is 10 tiks that shipped the milestone's primary deliverable (~1 360 lines of
net-new fence in one file, plus assertion F and a ref-awareness change across three guards); pass 15
stabilized against a window a fraction of that size. Every defect found this session was **fixed
inline, with an inverted mutant proving the test can fail**; nothing was routed forward, nothing was
waived, and no gap is known-and-unclosed. What remains is that a batch this large has more than
three dimensions worth running — G5's attribution reach is **1 of 24** and G1's prose class is the
declared clause-5 residual, both of which are corpus-repair work with a named owner, not hardening
deficits.

### Pass 18 — amendment: the flake gate found a regression I had introduced

Written after the entry above, and it corrects it. The Phase-5 flake gate (3 consecutive full
`stack-core` runs) took `stack-core` from the expected **1F to 2F**, and the extra failure was the
**pre-existing mutation battery reporting THEATRE on one of its own mutants**:

```
FAIL: test_01_every_mutant_matches_its_DECLARED_verdict (mutant='value-no-clone-reads-as-clean')
AssertionError: THEATRE: mutant 'value-no-clone-reads-as-clean' left the suite GREEN.
  an unread source is not a matching source: on a fresh box with no clone this would report a
  clean corpus it never opened
```

That mutant flips `derived_value_guard`'s doc-level control to `if measured and False:` and
requires the suite to go RED. **Pass 17's `if not scalars: return 2` gate, added behind it as
belt-and-braces, absorbed the mutation** — the mutated guard still exited 2, the suite stayed
GREEN, and a mutant that had been catching a real property for three iterations stopped catching
anything.

**Defence-in-depth that blinds your own mutation battery is a net LOSS of detection.** The gate
bought a branch nothing can reach — pass 17 had already proved `measured > 0` implies
`scalars > 0` by construction, and said so in its own commit message — and it spent a working
mutant. Removed (`4c301dd`), with the reasoning left in place of the code so it is not re-added.
Pass 17's actual deliverable, the scalar DENOMINATOR in the verdict, is untouched.

**Two lessons, both general:**

1. **When two controls return the same exit code, adding one is never free.** This is the SECOND
   time in this session — `anchor_construct_guard` needed its UNMEASURED refusal ordered ahead of
   the resolver control (right code, wrong diagnosis), and this one needed removing outright. Both
   first drafts went through the wrong branch. You must say which control fires, and prove the
   other one still can.
2. **Only the full-suite gate could see it.** Every targeted run was GREEN, because the battery
   that caught it lives in a file that my change did not touch and that no per-symptom discipline
   would have selected. This is precisely the case for running the flake gate over the WHOLE
   section rather than the touched tests, and it is the first time in eighteen passes that the
   gate has paid for itself with a real defect rather than a stability measurement.

**Corrected session totals (passes 16–18):** 8 rext commits + 3 rosetta ledger commits ·
**34 tests added** · **7 live defects fixed inline** (6 in the passes + this one) · **30 mutants
run, 29 RED + 1 declared-unreachable survivor**, and **1 pre-existing mutant re-armed** after I
disarmed it.

**Flake gate:** re-run from scratch after the fix.

## Pass 19 — 2026-08-05 — incremental

**Iters hardened this pass:** iter-69 … iter-79 (11 tiks — the guard-assertion window: G9, G10, the
`ref_resolves_in` / `pin_dates_a_platform_claim` helper, and iter-79's `can_resolve_refs`)
**Tiks covered since prior pass:** 11 (pass 18 closed at iter-68)
**Deferred:** no. This pass had been deferred three consecutive runs; it ran.

**The dimension: the routed `CHECK-M257x-iter79-three-valued-discriminators`.** iter-79 closed a
three-valued hole in `pin_dates_a_platform_claim` and generalised the rule from its own defect —
***every derived discriminator has three outcomes: yes, no, and cannot-tell*** — then routed the
sweep for siblings to a harden pass rather than widening in place. This is that sweep.

**Swept surface.** Every subprocess-derived discriminator in the `stack-core` guard family
(`grep -rn returncode --include='*.py'`, 7 guards + `buildbench`, tests excluded). Most were already
three-valued and said so in their own prose — `repos_yml_history` returns `(set(), "UNMEASURED")`,
`app_rpc_reads` returns `None` for an absent clone with *"None is not zero"* in the docstring,
`compose_counts_at` returns `None`, `_reads_at_ref`'s neighbour at `:762` explicitly reasons about
rc 1. That is the pass's main finding and it is a good one: **iters 77–79 taught this module the
rule, and the module mostly learned it.**

**One sibling had not.** `_reads_at_ref` — the function whose own comment reasons about rc 1 —
published **two** outcomes where `git grep` has three:

```
if r.returncode not in (0, 1):        # 1 == no match, which is a real answer
    return out                        # <-- the SAME empty dict as rc 1
```

So a clone whose object store cannot be read reported **"the consumer side reads no `*_RPC_ADDR`"**
under a confident `origin/main@<sha>` provenance. `res.reach["app_consumer_side"]` read `measured`.
That is the `|| echo 0` signature M257 opened on, one level down, in the guard built to end it —
and it is *the same author, the same window, the same rule, missed in the adjacent branch of the
function that states the rule*. The **fifth** occurrence this milestone of *"the author of a newly
written rule violated it while writing it."*

**Reachability was PROVEN, not argued — and the existing upstream guard does not cover it.**
Pass 18's own lesson (*"when two controls return the same exit code, adding one is never free"*)
makes this the load-bearing question, so it was answered by construction before any code changed:

| condition | `rev-parse --verify` (what `app_ref_sha` runs) | `git grep <sha>` | covered? |
|---|---|---|---|
| shallow clone, sha absent | **fails** → `unresolvable:` | — | **already covered** |
| unreadable / corrupt tree object | **succeeds** | **rc 128** *"unable to read tree"* | **NOT covered** |

`rev-parse` answers from the **commit** object; `git grep` needs the **tree**. Reproduced with a
built repo + `chmod 000` on the loose tree object: sha resolves, grep returns 128, guard returned
`{}`. Same shape for a partial clone whose promisor is unreachable.

**Fixed** so the three outcomes are three values: rc 1 → `{}` (a real answer), rc ≥ 2 → `None`,
propagated by `app_rpc_reads` as provenance `unreadable:<ref>@<sha>`. Deliberately **not** a
fall-through to the next AUTO candidate — answering a question about `origin/main` with a reading of
`HEAD` is §5 rule 7.

**3 mutants, 3 kills, 3 distinct signatures**, collected before running (§8 rule 5):

| mutant | signature |
|---|---|
| restore the pre-fix two-valued collapse | `'unmeasured' != 'measured'` — *the defect itself, named* |
| over-correct: every non-zero rc → cannot-tell | `{} != None` — caught by the **control** test, which exists so the fix is not "treat rc 1 as unmeasurable" |
| verdict right, provenance silent about **why** | `False is not true : <sha>@<sha>` — pass 18's *"a reporting path with no mutant is a docstring"*, applied |

**+3 tests** (157 → **160** in `test_platform_predicate_guard.py`), including the rc-1 control, which
is the one that keeps the fix honest: without it, a change making the guard unmeasurable whenever
the fold *completed* would pass identically.

**Flake gate:** full `stack-core` run before and after; the pre-existing `1F`
(`test_claim_twin_guard_iter48_answer_key`, the known perishable iter-48 fixture) reproduces
unchanged and is the only non-green.

**Routed-forward queue:** RF-2, RF-3, RF-7…RF-14 unchanged — none is in this pass's scope
(all Playthrough/seeder surface; this window was guard code).
`CHECK-M257x-iter79-three-valued-discriminators` is **CLOSED** by this pass.

## Pass 20 — 2026-08-05 — incremental

**Iters hardened this pass:** iter-80 … iter-94 (15 tiks — the guard-family + demopatch window)
**Tiks covered since prior pass:** 15 (pass 19 closed at iter-80)
**Scope note:** this is the LAST pass permitted to touch the measuring instrument. The milestone's
open clause 5 is met only by a KB-fidelity reading of zero, the protocol forbids repair inside a
measuring pass, and the previous run declined to take the reading precisely because it had rewritten
the guards four times in the same run. So the plan is: harden the instrument here, freeze it, and let
the NEXT run take the reading as its whole job. **No clause-5 reading was taken in this pass.**

### The dimension: anti-vacuity, applied to the instrument itself

Three anti-vacuity defects were fixed in iters 91–94 (a control that could never fire; a guard passing
over zero corpus docs). The generalisation nobody had run is the obvious one: **the guards were audited
for vacuity; the guards' own TESTS were not.**

**The family runner's one load-bearing property was covered by nothing.** `guard_family.py` exists
because *"all six corpus guards exit 0"* was a statement about a list somebody remembered. Its single
value is that a guard which says it checked NOTHING must not read as GREEN. The test named for that
property was:

```python
for out in ("repair-leak: CANNOT RUN — no candidate shingles", ...):
    self.assertTrue("CANNOT RUN" in out or "Nothing was checked" in out)
```

Two string literals declared three lines above, asserted to contain substrings of themselves. It never
imported the behaviour and never called `run_one`. **Measured:** blinding the branch to `if False:`
left all 22 tests in the file GREEN.

Fourth occurrence of the pattern in this window, and the first one *inside the runner built to refuse
it*. Replaced with a real subprocess probe through the real `run_one`, plus the negative control (a
guard that really checked something stays GREEN — otherwise "grade every exit 0 as CANNOT-CHECK" would
pass) and the end-to-end half (a CANNOT-CHECK member takes the FAMILY out of green — a separate code
path from the member verdict).

### Two live defects fixed inline

**1. `--verify-remote` was a silent no-op without `--platform`** (`guard_family.py`). Its only subject
sat inside `if platform:`, so a run that ASKED for the freshness check and supplied no `--platform`
made no network call, said nothing about it, and could exit 0 — a transcript indistinguishable from a
genuinely remote-verified green. That flag is **the remedy this milestone shipped for the reported
stale-clone reading**, and `grep -rn 'verify.remote' tests/` returned nothing: the remedy was
unexercised and half of it did nothing. Now UNMEASURED (exit 2) — the module's own doctrine (`ab107b1`)
turned on the module. This is the regression test for the reported class the prior run could not
reproduce: the class was real, and its remedy was untested.

**2. demopatch's post-condition rollback abandoned the discipline it exists to enforce** — closes
`CHECK-M257x-iter91-g5xg7-journal-on-postcondition-failure` **by fault injection**, the one interacting
guard pair iter-91's grid left routed (G7 can only fail on a short write, which the harness could not
produce; the source-mutant vehicle the journal-blinding control already uses reaches it).

`cmd_apply`'s forward write is atomic and verified — tmp + fsync + `os.replace` + read-back — and its
comment says why (the apply_patch B12 lesson about half-written source). The rollback that runs when
that read-back **fails** was a bare truncating `open(target,"w")`: no fsync, no read-back, no
verification — and `_journal_drop` ran unconditionally on the next line. Under the one condition that
causes a G7 failure, the rollback is written the same way and is just as likely to be short, and the
tool then destroys the only exact copy of the pre-image it holds. `_revert_inner` already had the right
ordering forty lines below (restore → verify → drop). Fixed by extracting `_write_atomic` so both
writes share one implementation, and verifying before dropping; the journal is KEPT when the rollback
did not stand, because that is exactly when it is the recovery.

**A second G7 finding, pinned rather than fixed:** after a G7-refused apply that rolled back cleanly, a
later `revert` exits 1 — the journal is correctly gone, so the baseline path sees a target that is
neither `pre_sha256` nor `post_sha256`. Clone is clean, nothing harmed; what is lost is that
`demopatch.log` cannot distinguish *"never applied"* from *"failed to come off"*. That is
`CHECK-M257x-iter90-revert-idempotency`, and it is **deliberately not re-pinned** — a sha re-pin goes
stale on the next `make pull`. The test pins current behaviour and names what to flip.

### The mutation battery caught a defect in its own author's new test

Mutant M6 (*"stop printing the corpus sha"*) **survived**. Cause: two git repos built by one helper in
the same second are byte-identical commits — same tree, author, message, timestamp — and therefore share
a sha, so `assertIn(corpus_head, out)` was being satisfied by the **platform** reference line. The new
provenance test was passing while asserting nothing. Fixture now tags each repo distinctly and asserts
against the corpus LINE. **Recorded in `_repo`'s docstring**: the same defect class as the pass's own
subject, committed by the person fixing it, within the hour. This is the third time this milestone that
the author of a newly written rule violated it while writing it — and the first time a mutation battery,
rather than a reviewer, is what caught it.

**Coverage delta on touched files:** `test_guard_family.py` 22 → 31 tests; `test_demopatch.py` 60 → 62.
**Tests added:** iter-86/91 → `stack-core/tests/test_guard_family.py`: +9 (3 anti-vacuity incl. negative
control + end-to-end, 2 provenance, 4 verify-remote). iter-90/91 → `demo-stack/tests/test_demopatch.py`:
+2 (G5×G7 fault injection).
**Mutants:** 11 run, 11 killed after the M6 fixture repair (10 on `guard_family`, 1 re-mutating the
demopatch fix). Signatures distinct.
**Bugs surfaced + fixed inline:** 2 (`--verify-remote` no-op `f310cb7`; demopatch rollback `53264e4`).
**Flakes stabilized:** none surfaced.
**Knowledge backfill:** deferred to the pass that closes the session (the protocol doc §8 anti-vacuity
rule wants the whole session's findings, not this pass's alone).

**Routed-forward queue:** `CHECK-M257x-iter91-g5xg7-journal-on-postcondition-failure` **CLOSED** (fault
injection). `CHECK-M257x-iter90-revert-idempotency` pinned, still open, deliberately not re-pinned.
`CHECK-M257x-iter90-realmanifest-baseline` untouched by design. `CHECK-M257x-iter93-general-hedge-fence`
and `CHECK-M257x-iter91-claim-twin-answer-key-stale` not in this pass's dimension.

**Stop condition:** continue-to-next-pass — a full-family anti-vacuity scan surfaced three further
unearned-GREEN paths not yet closed: `repair_reach_guard` exits 0 printing *"every booked finding was
reached"* over an all-ungraded ledger; `unreadable_repo_claim_guard`'s "PREMISE LIFTED" is an exit 0 the
family reads as GREEN; and `guard_family --allow-not-run` prints *"OK — every member of the census was
run and returned green"* immediately after printing that N were not run.

## Pass 21 — 2026-08-05 — incremental

**Iters hardened this pass:** iter-80 … iter-94 (same window; the dimension is different)
**Tiks covered since prior pass:** 0 (second pass of the same session)

**The dimension: the family-wide anti-vacuity sweep.** iters 91–94 gave three guards a floor **one at a
time**, each as a reaction to a specific symptom. The sweep nobody had run is the one that asks the same
question of all seventeen. Three more paths where a guard publishes a verdict over a universe it never
examined — and the family records GREEN.

**1. `repair_reach_guard` printed *"every booked finding was reached or dispositioned"* over a graded set
of ZERO.** Controls 1 and 2 guard the INPUTS (empty ledger, empty diff). Nothing guarded the OUTPUT of
classification: `graded` is `TOUCHED + WAIVED + UNREACHED`, and `NO_ANCHOR` (the `**Anchor:**` bullet
would not parse) and `OUT_OF_TREE` are in **neither** `graded` nor `UNREACHED`. A ledger of 152 findings
whose anchor-bullet shape drifted classifies 152× `NO_ANCHOR`, the reach line is suppressed by
`if graded:`, `bad` is empty — and the guard exits 0 saying all 152 were reached. The parser is
demonstrably fallible: the shipped fixture already yields 4 no-anchor + 1 out-of-tree. **The same hole
swallows a shrinking denominator** — 151 no-anchor + 1 touched prints `reach 1/1 = 100.0%` — a hazard this
module's own tests *name in prose* (*"a silently shrinking denominator is how a reach number flatters
itself"*) and never assert.

**2. `repair_reach_guard`: `git ls-tree` failing silently DISABLED the out-of-tree classification.**
`if rc == 0: tree_files = …`, else left `None`, and `None` means *do not restrict* downstream. The
`_reads_at_ref` shape pass 19 fixed, one guard over, in the same module family. Now fatal, quoting git's
own stderr. Not constructible from a bad ref — the diff control refuses first, correctly — so the test
injects at the single call site.

**3. `unreadable_repo_claim_guard`: "PREMISE LIFTED" was an exit 0 the family read as GREEN.** Measured
through the real `guard_family.run_one` with a clones root containing `infrastructure`: **rc=0,
verdict=GREEN** — a guard that scanned zero sites counting toward the family's green total. Now exit 2 in
the family's own vocabulary, which also makes the tripwire loud: the day `infrastructure` joins a clone
set the family goes UNMEASURED until somebody measures those declarations and retires the fence, which is
what the guard's own class docstring says should happen. The existing test pinned the 0; it now pins both
halves with the measurement in its docstring. **Also:** its anti-vacuity floor counted a different
universe from its findings (`README.md` in the findings scan, absent from the denominator), so a
construct living only in `README.md` produced a real finding **and** `total == 0`, and the floor fired
first — exit 2 over a violation the guard was holding.

**4. `guard_family`'s summary sentence contradicted the line above it.** With `--allow-not-run` it printed
`N guard(s) NOT RUN and accepted` and then, unconditionally, **`OK — every member of the census was run
and returned green.`** The second is false whenever the first prints, and the second is the one that gets
quoted forward — this milestone has quoted a family green forward more than once. `--allow-not-run` is
documented to RECORD the gap, never hide it; the summary hid it in the same breath the line above
disclosed it.

**A survived mutant, and it changed the design.** A dedicated *"the tree listing was empty"* refusal could
not be killed: with an empty `tree_files`, `classify` marks every anchor out-of-tree, the graded set
empties, and control 3 refuses anyway — naming `out-of-tree=<n>` while it does. **Pass 18's lesson applied
rather than re-learned:** two controls on one exit code buy an unreachable branch and spend a working
mutant. The branch was removed, the reasoning left in its place, and the test now pins the OUTCOME and
names its enforcer. This is the second consecutive session in which that rule has changed a decision.

**Tests added:** +12 (`test_repair_reach_guard.py` 16 → 21, `test_unreadable_repo_claim_guard.py` 10 → 13,
`test_guard_family.py` 31 → 33).
**Mutants:** 8 run, 7 killed + 1 survivor that was correctly resolved by deleting the redundant control.
**Bugs surfaced + fixed inline:** 4, all in `83637c6`.
**Instrument end-to-end after the changes:** `14 GREEN · 0 RED · 0 could-not-check · 3 not-run` over 17
members, corpus `5a1646718`, platform `0c91421df` (origin/main in sync, fetched 92m ago), and the summary
line now reads *"OK with gaps — 14 of 17 … This is NOT a whole-family green."*

**Stop condition:** continue-to-next-pass — the UNMEASURED path `platform_alignment_guard` gained at
iter-91 still has no end-to-end `main() == 2` test, `guard_family`'s `fetch_age_min` provenance field is
unpinned, and the family's CANNOT-CHECK detection is a content sniff over merged stdout+stderr that has
never been tested against a guard which echoes corpus prose.

## Pass 22 — 2026-08-05 — incremental

**Iters hardened this pass:** iter-80 … iter-94 (same window; third and last dimension)
**Tiks covered since prior pass:** 0 (third pass of the same session; the incremental cap is 3)

**The dimension: the code iter-91 added, which is the newest in the window and the least exercised.**
Both defects live in the paths that gave the guards their third verdict.

**1. `ALIGNMENT_ALLOW_UNMEASURED=1` hid exactly what it promised to record.** The refusal offers the
variable as accepting the gap *"which RECORDS it rather than hiding it."* Measured with it set, against a
platform root that is not a git repo:

```
platform_alignment_guard: assertion F resolved 3 citation(s) — … 0 unresolvable
platform_alignment_guard: OK — map.md and repos.yml agree in both directions.
```

Nothing on either stream said that **3 citations had been read from the WORKTREE because no ref
resolved**. The always-printed reach line reported `unresolvable` and never the worktree-fallback count,
so on the accepted path there was no trace at all: a clean, unqualified green, reachable by one
environment variable, on the guard the exit gate leans on hardest. **That is the precise substitution the
refusal block exists to prevent, restored by its own escape hatch.** Fixed in three parts — the worktree
count joins the reach line on every run; an accepted gap is announced on the same stream as the refusals;
and the verdict becomes `OK WITH AN ACCEPTED GAP … this is NOT a whole-map green`.

Separately: **the UNMEASURED verdict itself had no end-to-end test.** Every test for it stopped at
`cited_text`; `main()`'s grading step (`no_ref_clones` → `blind` → `return 2`) was exercised by nothing. It
needed no fixture construction — a platform root that is not a git repo **is** the condition, so the
synthetic battery's existing fixture reaches it. Three tests now cover exit 2, the accepted path, and the
unqualified-OK control.

**2. `guard_family`'s CANNOT-CHECK sniff could DOWNGRADE a real RED.** It ran first and unconditionally
over merged stdout+stderr, so a guard that exited 1 while **echoing a corpus line** containing the phrase
was graded CANNOT-CHECK — which drops it from `red` and turns the family's report from `RED — <guard>` into
`could not check`, so the findings vanish from the one view that claims to summarise the family. Two guards
in this census echo corpus lines verbatim and both phrases occur in this repository's own prose. Scoped to
`rc == 0`, the only case it was ever for: a guard that exits non-zero has already said it is not green, and
its exit code is a better witness than a substring of its output.

**3. `fetch_age_min`** — the third provenance field — was still unpinned and could be deleted green. Both
directions now covered: reported when a `FETCH_HEAD` exists, and **absent rather than invented** when it
does not.

**Tests added:** +7 (`test_platform_alignment_guard.py` 42 → 45, `test_guard_family.py` 33 → 36).
**Mutants:** 8 run, 8 killed.
**Bugs surfaced + fixed inline:** 3, all in `6130bfd`.

**Knowledge backfill:** `corpus/ops/platform-alignment.md` §8 gains *"Then audit the guards' TESTS the same
way — that is where the next three were"* (`8ef3906`), generalising iter-94's rule: **the thing that
reports is not the thing that measured.** Names the three shapes (a flag or hatch that silently does
nothing; a verdict over a graded set of zero; a summary sentence contradicting the line above it) and the
two operational corollaries (an accepted gap is still a gap, and grade on the exit code rather than a
substring wherever the code exists).

### Flake gate — 3 consecutive full `stack-core` runs on the settled tree

| run | tests | failures |
|---|---|---|
| 1 | 910 | 1 — `test_claim_twin_guard_iter48_answer_key` |
| 2 | 910 | 1 — same |
| 3 | 910 | 1 — same |

Plus two earlier full runs at intermediate commits (894/1F, 904/1F) and a full `demo-stack` run:
**1058 tests, 6 failures — exactly the attributed set** (3 need a live container; 3 are stale live-clone
baselines across two independent patch vehicles). **Both known pre-existing failure sets reproduce
unchanged.** The one change in the set during this session was the collection fence going RED, which was a
real defect in iter-94 and is fixed (`7f64003`).

**Session totals (passes 20–22):** 5 rext commits + 4 rosetta commits · **+27 tests** across 5 files ·
**10 live defects fixed inline** · **26 mutants run, 25 killed + 1 survivor that correctly deleted a
redundant control**.

**Stop condition:** cap reached without stabilization — the 3-pass incremental cap fired, and the third
pass still surfaced three defects, so the dimension scan did not come up empty. That is the honest reading
and it is **not** a request for a fourth pass: **this was the last pass permitted to touch the measuring
instrument** (the next run takes the clause-5 reading and the protocol requires the instrument untouched
during it). The remainder is therefore **routed forward, disclosed, not silently dropped**:

- `platform_predicate_guard` records `app_consumer_side=unmeasured`, `repo_vocabulary_history=UNMEASURED`
  and `guard_platform_ref=UNRESOLVED` in its reach line and still returns 0 — the sibling of the fix
  iter-91 applied to `platform_alignment_guard`, left unapplied. **Not live on this box** (measured: the
  app consumer side reads `measured @ origin/main@2035f9a`), and grading it is a design decision about
  whether partial blindness should block the family, so it is routed rather than taken unilaterally.
- **Waivers are not reported and staleness is not detected.** `repair_leak_guard.py` says in its own words
  *"It can only ever make the fence quieter, so it is reported"* and `repair_leak_waivers.json` says
  *"every one is reported on each run"* — neither is true of the code; a leak waiver swallows a finding
  with zero trace. Neither waiver file has stale-entry detection, and `repair_reach_waivers.json`'s six
  entries are keyed to one specific ledger.
- **A crash is rendered as RED with its traceback itemised as findings.** `story_org_count_guard`,
  `platform_predicate_guard` and `unreadable_repo_claim_guard` each `read_text()` an input with no
  `is_file()` check, so a missing input is an uncaught exception (exit 1), and `headline()` then counts the
  indented traceback lines as a finding count.
- `demo_knob_guard` has no vacuity control at all.
- demopatch: a corrupt or truncated journal at revert falls back to the sha baseline **silently** — apply
  warns about a missing entry, revert never does — and an `OSError` inside the G7 failure path escapes
  untyped past `main`'s `except PatchError`, leaving a stale journal entry.

**The instrument as frozen**, run end-to-end after every change in this session:
`14 GREEN · 0 RED · 0 could-not-check · 3 not-run` over 17 members, corpus `8ef3906`, platform `0c91421df`
(origin/main in sync), summary line `OK with gaps — 14 of 17 … This is NOT a whole-family green.`

## Pass 23 — 2026-08-06 — incremental

**Iters hardened this pass:** iter-105 … iter-109 (the TOK-06 fence-build window)
**Tiks covered since prior pass:** 15

**The scope, stated before the numbers.** Pass 22 closed at iter-94. The unhardened *code* surface of
iters 105–109 is entirely in `rosetta-extensions/stack-core`: three brand-new guards
(`fence_provenance.py` 152 lines, `clone_drift_guard.py` 314, `anchor_offset_guard.py` 361), a provenance
banner threaded through 16 existing guards, and `guard_family`'s grading changes. **iter-109 shipped no
production code and no tooling** — it is the clause-5 reading — so it contributes no surface, exactly as
the milestone brief said.

**Coverage delta on touched files:** not a line-coverage story — there is no `coverage` module on this
host's Python 3.9.6 and none of these sections has ever been instrumented, so this ledger's established
measure applies: test count plus fault injection.

`python3 -m pytest tests/test_fence_provenance.py tests/test_clone_drift_guard.py
tests/test_anchor_offset_guard.py tests/test_guard_family.py tests/test_test_collection_fence.py -q`
in `stack-core`: **105 → 121 passed** (+16). The same five-module invocation is the one every count below
is taken with.

**Bugs surfaced + fixed inline: 4.**

**1. An ordinary APPEND was a false RED** (`68587e3`). `anchor_offset_guard.line_map` took a pure
insertion's shift boundary as `old_start + old_len` — which for a pure insertion *is* `old_start`. git
writes `@@ -10,0 +11,2 @@` for *insert AFTER old line 10*, so old line 10 does not move; the guard
declared the line immediately above every insertion to have moved. Measured: a citation to the last line
of a file, plus a two-line append touching nothing that already existed, graded `what was line 10 is now
line 12`.

This is a **false RED on the shape the guard exists for** — both booked incidents (iter-100, iter-102)
inserted prose — and it is precisely the failure `test_append_BELOW_every_citation_is_GREEN` names in its
own assertion message (*"cries wolf on ordinary appends and gets suppressed"*) while being unable to catch
it. **Off-by-one and correct agree everywhere except ON the boundary**, and not one of the four existing
fixtures cited there: the RED fixture inserts at the file TOP (`@@ -0,0 +1,2 @@`, where boundary 0 and 1
are indistinguishable), the GREEN fixture appends at the END while citing line 3. The pinned iter-102
answer key is unaffected for the same reason — its insertion sits above `:321`, never on it — and still
asserts `321 → 331` after the fix, which is the check that the correction did not move a real measurement.

**2. A clone whose HEAD could not be read was graded as DRIFT** (`be269b0`). `clone_drift_guard`'s
`reconciled` compares against `self.head`; an unreadable HEAD is `""`, which starts with nothing, so the
repo failed that test for a reason unrelated to drift and was printed as a verdict:
`[D1 advanced] beta is at , which the corpus never cites — None commit(s) past …`. An empty sha and a
`None` count, asserting an advance about a repo it could not read at all — the substitution `ab107b1`
fixed elsewhere in this family. **Reachable, not contrived:** it is the state `ensure-clones.sh` leaves a
bootstrap clone in (objects fetched, HEAD unborn), so `cat-file` answers — which is what puts the repo in
the cited set — while `rev-parse HEAD` does not. Now UNMEASURED on D2's existing three-verdicts footing,
named on its own line and **disclosed on the OK line**, since an accepted gap is still a gap.

**3. The provenance caveat reached the two OK lines and no others** (`0aac9fa`). iter-105 computed it
after the `red` / `blocked` / fatal-`not_run` branches had already returned. **The branch it missed is the
one the fence was written for: iter-103's incident was a misread RED.** A run from a dirty authoring copy
reproduces that condition exactly, and the line quoted forward out of a RED run said nothing about it.
Also named the third state — `dirty is None` (HEAD read, `status` did not), which fell through both
branches and let the summary imply a clean tree it had never checked.

**4. A test class hidden from direct execution since iter-107** (`be269b0`). `adcf689` appended
`TestKnownWeakness` **after** `if __name__ == "__main__"`, so `python3 test_clone_drift_guard.py` skipped
it and printed OK. `test_test_collection_fence` has been **RED on this file for three iters** with nobody
seeing it — because the full `stack-core` suite does not complete on this host and no scoped run in that
window included the fence. The hidden test is the one pinning the drift fence's documented weakness *"in
the suite rather than only in prose"*; it was in neither. Found because the fence fired on **this pass's
own append**, then named the older violation beside it.

**Mutation controls:** 3 mutants applied by fault injection into scratch copies (the live tree was never
modified) — the pre-fix boundary expression, the caveat dropped from the RED/could-not-check lines, and
the third-state branch disabled. **All 3 killed** (9 test failures across them). The surviving tests in
each case are the deliberate negative controls, which must hold under both versions: a REPLACEMENT hunk
unperturbed, a top-of-file insert still shifting line 1, a clean tree adding nothing to the RED line, and
a genuinely advanced clone still RED beside an unreadable one.

**Adjacent-module regression check:** `python3 -m pytest tests/test_repair_postcondition.py
tests/test_repair_postcondition_audit_mode.py tests/test_test_collection_fence.py
tests/test_m257x_mechanical_fences_mutation_battery.py tests/test_iter45_mechanical_fences.py -q` →
**147 passed, 1 failed in 481 s**, the single failure being finding 4 above; green after the fix.

**Stop condition:** continue-to-next-pass — four live defects in the first dimension pass is not a
stabilized surface, and `fence_provenance.py` itself plus the 16-guard stamp threading have not yet had a
dimension scan of their own.

## Pass 24 — 2026-08-06 — incremental

**Iters hardened this pass:** iter-105 … iter-109 (same window; second dimension)
**Tiks covered since prior pass:** 0 (second pass of the same session; the incremental cap is 3)

**The dimension: `fence_provenance.py` itself** — the primitive iter-105 built the other two guards'
provenance on, and the one file pass 23 did not scan. Its 152 lines carry the milestone's densest
concentration of *stated* safety properties, so the pass was run as an audit of those statements rather
than of the code: **take each claim the docstring makes and try to make it false.**

One of them was.

**5. `STAMPED_ENV` is a general quiet switch after all, on the one caller that must never be quiet**
(`9dfd717`). The module says, in its own words:

> It is deliberately NOT a general quiet switch: a suppressible stamp with no one else stating it would
> re-introduce the exact defect this module exists to close. `guard_family` sets it only AFTER printing
> the line itself.

The second sentence is true of `guard_family`. The first is not true of the variable, because
**nothing enforces the second sentence.** `stamp()` reads `STAMPED_ENV` from the AMBIENT environment, so
an operator, a wrapper script, or a shell that exported it once silences the stamp in every process
started thereafter — including the family runner, where by construction there is no one else to state
the tree.

Measured, with `FENCE_PROVENANCE_STAMPED=1` exported into a `guard_family` run:

```
guard-family: corpus /…/fam @ ?
guard-family: 1 GREEN · 0 RED · 0 could-not-check · 0 not-run
guard-family: OK — every member of the census (1) was run and returned green.
```

No `fence-tree:` line at all. **A clean, unqualified green stating no tree** — iter-103's condition
exactly, reachable by one environment variable, on the fence built to close it.

**Why the 19 existing tests could not see it.** `TestLiveBehaviour._run` **pops** `STAMPED_ENV` before
every invocation, so the ambient case was unreachable from the suite by construction; and the one
suppression test that does run drives a **member** (`corpus_index_guard`), where suppression is correct
and must stay correct. The suite tested the half of the contract that works.

`stamp()` gains `force`; `guard_family` forces. A member's suppression is untouched.

**Coverage delta on touched files:** `python3 -m pytest tests/test_fence_provenance.py
tests/test_clone_drift_guard.py tests/test_anchor_offset_guard.py tests/test_guard_family.py
tests/test_test_collection_fence.py -q` in `stack-core`: **121 → 125 passed** (+4;
`test_fence_provenance.py` 19 → 23).

**Mutation controls — both directions, because this fix has a symmetric failure mode.** A `force` that is
always on would make every member stamp, which is the noise the variable exists to prevent, and it would
pass a one-directional test. Applied **separately**, since applied together they cancel exactly (the
first run of this battery was a false green for that reason, and is recorded here because the shape
recurs): dropping `force=True` fails 2 tests; making `force` unconditional fails 1.

**Scanned and found clean / not worth a fix:**

- The docstring's two `headline()` claims (a stamp printed last would replace a green member's summary;
  a stamp cannot inflate a RED cardinality) are **defensive about a state that needs two independent
  regressions to reach** — members do not stamp in a family run at all, because `run_one` sets the
  variable in the child env. Left untested deliberately rather than pinned at cost.
- `citations()` resolves its file set from `ls-tree` at a rev but from `rglob("*.md")` in the working
  tree — **8830 files vs 90 corpus ones on the live tree, in 1.0 s**. A latent asymmetry, not a live
  defect: target resolution filters to `corpus/` afterwards, and `corpus_files` is built from the same
  prefix in both branches. Measured rather than assumed, and not repaired — the repair would be
  gold-plating a path with no reachable wrong answer.

**Stop condition:** continue-to-next-pass — one live defect on the second dimension is fewer than the
four on the first, but a scan that still surfaces a defect in the fence family's own primitive has not
come up empty.

## Pass 25 — 2026-08-06 — incremental

**Iters hardened this pass:** iter-105 … iter-109 (same window; third and last dimension — the incremental cap is 3)
**Tiks covered since prior pass:** 0 (third pass of the same session)

**The dimension: what the tooling SAYS about itself, graded against what it does.** Passes 23 and 24 read
the code; this one took every quotable number and stated property in the three new guards and re-measured
it. That is the milestone's own subject applied to the instrument, and it is where both findings came
from.

**6. The answer key's own rationale quoted a total under one part's name** (`7dd5148`).
`normalize_range`'s docstring — the rationale for iter-108's bare-rev fix — records
`cd16967^..cd16967` as *"53 changed files, 17 graded, **10 findings**."* Re-measured: **17 graded, 5
findings plus 5 CANNOT-TELL.** Ten is the **sum of the two classes this module goes to some length to
keep apart**: `findings`, which sets the exit code, and `review`, which is reported, counted, and
deliberately does not. Quoting the total under the name of one of its parts is the §5 rule 11/12
substitution — and it had landed *inside the rationale of a fix for the same class.*

The answer key could not catch it: `test_iter102s_commit_is_RED` asserts only that `findings` is
non-empty. Now pinned **per class** — 5 findings / 5 review / 17 graded / 53 changed.

**7. `--json` is unparseable across the family — ROUTED, not fixed**
(`FIX-M257x-harden23-json-polluted-by-provenance-stamp`). iter-105 put `stamp()` on **stdout** in every
guard's `__main__`; **12 of those guards also offer `--json`**, so the provenance line precedes the JSON
document and any machine consumer dies at char 0. **The suite does not see it because the suite works
around it** — every existing test that parses a guard's `--json` sets `FENCE_PROVENANCE_STAMPED=1` first
(`test_anchor_offset_guard.py:224`, `test_clone_drift_guard.py:265/276/284`). An undocumented,
load-bearing workaround is how a live defect stays invisible while the tests stay green.

**Why routed rather than taken.** The fix is one line — default `stamp()`'s stream to `sys.stderr` — but
it retires the *printed-FIRST-on-stdout* property iter-105 designed deliberately and documented in
`fence_provenance.py`'s docstring **plus sixteen per-guard `__main__` comments**, and it changes what
`test_a_standalone_guard_run_prints_the_tree` asserts. That is a design decision about where provenance
belongs, not a corollary of a regression test. **No production consumer of `--json` exists today** (12
guards offer it; the only parsers are the tests), which is why it is a routing and not an escalation.
**+2 known-limitation tests pin both the defect and the workaround**, on the `TestKnownWeakness` pattern
this suite already uses — the assertion has an expiry date, and when the fix lands those tests are the
ones to rewrite.

### Did the append fix change any real verdict? No — and this ledger says so.

The pass-23 boundary correction was re-run against **seven real ranges** (the five in-scope iter commits
plus `22eaac4`, `e6aed2e`, and the pinned iter-102 answer key `cd16967`), fixed versus a fault-injected
pre-fix copy:

| range | fixed (findings/graded) | pre-fix |
|---|---|---|
| the five iter-105…109 commits | 0/0 each | 0/0 each |
| `22eaac4`, `e6aed2e` | 0/0 | 0/0 |
| `cd16967^..cd16967` | **5/17** (+5 CANNOT-TELL) | **5/17** (+5 CANNOT-TELL) |

**Identical everywhere.** The false RED is real, reproducible, and pinned by six tests — and it **has not
yet fired on this corpus's history**. Recorded that way deliberately: the temptation is to report five
repaired REDs, and the honest reading is that the defect was caught before it cost anything. The same
table is the check that the correction moved **no real measurement**.

### The instrument, run end-to-end after every change this session

```
guard-family: 15 GREEN · 0 RED · 0 could-not-check · 4 not-run
guard-family: 4 guard(s) NOT RUN and accepted: anchor_offset_guard, repair_leak_guard,
              repair_reach_guard, value_change_guard
guard-family: OK with gaps — 15 of 19 member(s) ran and returned green; 4 was/were NOT RUN and
              accepted. This is NOT a whole-family green.
```

over **19 members** (17 at pass 22; iter-106 and iter-107 added two), corpus `4ede6495b`, platform
`0c91421df` (origin/main in sync), fence tree `9dfd717f2`. The four NOT-RUN are the commit- and
input-scoped members with no `--range` / `--ledger` supplied; **not a whole-family green, and the
summary line says so.**

**Coverage delta on touched files:** `python3 -m pytest tests/test_fence_provenance.py
tests/test_clone_drift_guard.py tests/test_anchor_offset_guard.py tests/test_guard_family.py
tests/test_test_collection_fence.py -q` in `stack-core`: **125 → 128 passed**. Session total across the
three passes: **105 → 128 (+23 tests)**, every count from that one invocation.

**Flake gate:** the 24 newly added tests, 3 consecutive runs — `24 passed` at 22.82 s / 23.07 s / 22.75 s.
Invocation: the four-module set with
`-k "PureInsertion or UnreadableHead or ProvenanceCaveatReaches or NotSuppressibleByAnybody or KnownWeaknessJson or CARDINALITY"`.

**Suite-completion gap, restated because it must not read as a pass:** `stack-core`'s full `pytest tests/`
**does not complete on this host** (`FIX-M257x-iter108-stackcore-suite-hangs`, open). **No whole-suite
total is quoted anywhere in this entry** — every number above names the invocation that produced it. The
one known pre-existing failure (`test_claim_twin_guard_iter48_answer_key::test_02`) was not run by any
scoped invocation this session and is therefore **not re-attested here** — unchanged, not re-verified.

**Scoped around, deliberately:** `repair_reach_guard`'s ledger-derivation path (`read_ledger()`) was left
untouched. iter-109 booked `FIX-M257x-iter109-repair-scope-is-detection-bounded`, whose binding change is
that the anchor set must in future be re-derived **from the corpus per predicate**, never from a
`raw/` ledger dir. Deepening tests against the current derivation would pin behaviour that is known to be
changing and would make the fix harder to land. No test was added to that path this session.

**Stop condition:** cap reached without stabilization — the 3-pass incremental cap fired and the third
pass still surfaced two items, so the dimension scan did not come up empty. That is the honest reading and
it is **not** a request for a fourth pass: the two remaining items are **one routed with a named handler
and two pinning tests, and one corrected in place**. What is left standing is the routed `--json` defect
plus pass 22's still-open queue (waiver staleness, the crash-rendered-as-RED class — to which
`clone_drift_guard` is now a **new member**, sharing the uncaught-`read_text()` shape, and
`demo_knob_guard`'s absent vacuity control).

## Pass 26 — 2026-08-07 — incremental

**Iters hardened this pass:** iter-111 … iter-119 (pass 25 terminated at `fab0e13`)
**Tiks covered since prior pass:** 9
**Mode note:** the invocation named `--final`. **This pass is recorded as INCREMENTAL, deliberately.**
The gate has NOT fired — clause 5 is open and the milestone is awaiting a user scope decision — and a
final-mode entry is exactly what `close-milestone` greps for to unblock a merge. Writing one here would
have made a pending scope decision look like a cleared gate. The trigger that was actually stated (9 tiks
since pass 25) and the cap that was actually referenced (3 passes) are both the incremental ones.

### The dimension scan, and what it found

Three passes, each attacking a young instrument rather than running it. **Four live defects**, two of
them in fences the milestone was relying on.

**Pass 1 — FENCE-M257x-iter117's controls could not be COLLECTED, and took the suite with them.**
`tests/test_corpus_citation_guard.py` annotates `-> Path | None` at module level with no
`from __future__ import annotations`. PEP 604 needs 3.10; the **only** interpreter on this host with
pytest is `/usr/bin/python3` == **3.9.6** (`python3` on this shell is homebrew 3.14 and has no pytest —
which is why every count in this entry names its interpreter).

```
BEFORE  /usr/bin/python3 -m pytest <every rext test_*.py>
        -> 2837 tests collected, 1 error ... Interrupted: 1 error during collection
AFTER   -> 2876 tests collected in 0.73s
```

Two harms, the second larger: the census fence's mutation **and** anti-vacuity controls had **never
executed** (the file landed at iter-117, after pass 25), and a collection error **aborts the run**, so one
module suppressed 2,836 other tests. `tests/test_test_collection_fence.py` was GREEN throughout — it
fences statement ORDER, which cannot see a module that never imports. It now carries an **importability**
arm (6 RED-proofs, mutation-verified). Once collectable, iter-117's own controls were **proven to fire**:
`return [], census` → 6 RED; `docs = []` → 9 RED including the anti-vacuity arm. **So iter-117 is
UNREACHABLE, not vacuous** — a distinction the eighth-vacuous-fence tally should not swallow.

**Pass 1 — G10 was a FALSE RED on a CORRECT corpus.**
`FIX-M257x-iter116-predicate-guard-takes-first-pin-in-block`, booked iter-116, carried through iter-119.
`sentinel.md:5` names `d11a403` in one clause and `0c91421d` in the next; G10 took the FIRST pin and
graded the second clause's count at the first clause's ref. **The RED appeared the moment iter-115's
repair added the second, MORE PRECISE ref** — the fence penalised the precision this milestone spent the
release adding. `_governing_pin` now takes the nearest pin at or before the claim. Reach moved the right
way: **3 → 0** claims graded at a weaker historical ref, i.e. all 6 real claims now checked at the live
checkout. Family **15 GREEN/1 RED → 16 GREEN/0 RED**. Only G10 is re-pointed; the token stays **OPEN**
for the other arms.

**Pass 2 — `fence_provenance` was blind to its own subject.** `sha = None` in `fence_tree()` → every
family verdict reads *"provenance unknown"* → **34 passed**. Its `TestAntiVacuity` is a good control
aimed at the wrong noun: it asserts *which modules exist*, when the subject is *the provenance value*.
**A fence can hold a correct anti-vacuity control and still be vacuous about its actual claim.** Closed
with 4 tests, both halves (real sha on a real tree; UNKNOWN still works on a non-git dir).
MUT `sha=None` → 3 RED · MUT `sha="0"*40` → 1 RED.

**Pass 3 — FENCE-M257x-iter118's mutation control could not isolate its own mechanism.**
`_NOT_A_CITATION` replaced with a never-matching pattern — the exact pre-iter118 behaviour — → **6
passed**. `is_not_a_citation` ORs three clauses and every URL the controls tested contains `//`, so the
broad clause satisfied every assertion alone. **A mutation control is only a control for the clause it can
ISOLATE.** Closed with clause-isolating cases (`redis:6379`, `grpc:9000` — regex-only;
`//cdn.example.com/x.js` — `//`-only), each asserted to be isolated so fixture drift cannot un-isolate it.
MUT-B → 2 RED (was 0) · MUT-C → 1 RED.

**Held under the same attack:** `clone_drift_guard` (D1 blinded → 6 RED; D1+D2 → 7 RED; its anti-vacuity
control correctly stays GREEN under blinding because it measures REACH, not findings — noted so a future
pass does not "fix" it) and `anchor_offset_guard` (2 sites blinded → 8 RED).

**Method defect found in my own instrument, and it is the reason three findings above are trustworthy:**
**three mutation attempts silently failed to apply** — a `.append` regex that missed multi-line calls, a
name set omitting `drifted`, and a single-quoted pattern against a double-quoted file — and each read as
*"the controls survive."* **A mutant that did not apply and a mutant that survived are indistinguishable
in the output, and only one of them is a finding.** Every mutation in this pass now carries
`assert count == 1, "MUTATION DID NOT APPLY"` before it runs.

### Coverage delta on touched files

| module | before | after |
|---|---|---|
| `tests/test_platform_predicate_guard.py` | 167 | **179** (0 skipped; the first anti-vacuity arm SKIPPED on a wrong attribute and was made fail-closed) |
| `tests/test_fence_provenance.py` | 34 | **38** |
| `tests/test_test_collection_fence.py` | 10 | **16** |
| `tests/test_anchor_construct_denominator.py` | 6 | **9** |
| `tests/test_corpus_citation_guard.py` | **0 collectable** | **19** |
| whole-tree collection | 2837 + 1 error (**run aborted**) | **2876, 0 errors** |

**Tests added:** +44 across 5 modules (12 · 4 · 6 · 3 + 19 recovered).
**Bugs surfaced + fixed inline:** 4 (`cbb60de`, `4bbe73e`, `b9bb2b6`; the corpus-side repairs `7f19846`,
`a5f0d81`, `f723101`).
**Flakes stabilized:** none found. Flake gate: 3 consecutive runs of the pass-26 tests → **35 passed** at
3.35 / 3.36 / 3.37 s.

**Knowledge backfill:** `corpus/ops/platform-alignment.md` — `FIX-M257x-iter120-anchor-guard-detects-blank-not-wrong`
(`anchor_construct_guard` detects *"resolves to blank"*, not *"resolves to the right construct"*, so
iter-119's 8-item wrong-construct class is itself a floor). `deferrals-audit.md` §11 — the re-derived
inventory. `progress.md` — the clause-5 measurement section.

**Suite-completion gap, restated so it cannot read as a pass:** `stack-core`'s full `pytest tests/` still
does not complete on this host (`FIX-M257x-iter108-stackcore-suite-hangs`, **open**). **No whole-suite
total is quoted anywhere in this entry** — every count names its invocation. The pre-existing
`test_claim_twin_guard_iter48_answer_key::test_02` red was not run by any scoped invocation here and is
**not re-attested** — unchanged, not re-verified.

**Stop condition: cap reached without stabilization** — the 3-pass incremental cap fired and pass 3 still
surfaced a live defect, so the dimension scan did not come up empty.

### The cap has now fired without stabilizing THREE consecutive times (22, 25, 26), and that is the finding

Recorded plainly rather than papered over, because the pattern says something the individual passes do not.

**It is not that the passes are failing.** Each of the three found real, live defects and closed them.
**It is that the supply has not thinned.** Three passes at the cap, three different young fences, and the
defect class is the same every time: *a fence that reports GREEN over something it never actually checked*
— unreachable controls (iter-117), the wrong noun (`fence_provenance`), the wrong clause (iter-118), the
wrong pin (G10). That is now **twelve** instruments in this milestone found green-over-nothing, and the
last four were found by **attacking** fences that had passed every prior harden pass.

**The honest reading: harden passes on this family are still finding first-order defects, so no pass has
yet measured a coverage plateau.** A fourth pass is not requested — the remedy is not more passes of the
same kind. What each of the four defects has in common is that the fence's control was aimed at something
adjacent to its claim, and no amount of running the suite surfaces that; only mutating the named mechanism
does. **That is an input to the scope decision the milestone is holding, not a request to continue.**

---

## Post-pass-26 note (iter-121, 2026-08-07) — NO fourth pass was requested, and the reason is now a standing rule

Pass 26's own stop condition already said it: *"the remedy is not more passes of the same kind … what each
of the four defects has in common is that the fence's control was aimed at something adjacent to its claim,
and no amount of running the suite surfaces that; only mutating the named mechanism does."* That reasoning
is **promoted out of this ledger** so it outlives the milestone:

- **`corpus/ops/platform-alignment.md` §5 rule 53** — three consecutive caps without stabilizing (22, 25,
  26) is a finding about the METHOD, and carries pass 26's load-bearing method defect: **three of its own
  mutations silently failed to apply and each read as *"the controls survive."*** Every mutation must now
  assert `count == 1` before its result is interpreted, **and** a mutation is only a control for the clause
  it can **isolate**.
- Both halves were applied to iter-121's own new controls: the beacon mutant removes **both** calls
  (removing one leaves the site still beaconing — a mutant that does not isolate the property), and the
  `blocking_state_guard` mutant is the historical bug itself (shrink `BLOCKING_FIELDS` to
  `("user-blocker",)` and the finding must be **lost**).

**Suite-completion gap — CLOSED as stated, and here is what replaces it.**
`FIX-M257x-iter108-stackcore-suite-hangs` was never a runtime defect (iter-111 measured the run at
**1090.88 s**, 1 failed / 1011 passed). It was an **ambiguity**: a running suite and a wedged suite emitted
the same thing, so three consecutive entries in this ledger had to say *"no whole-suite total is quoted
anywhere in this entry."* iter-121 landed `tests/progress_beacon.py`, wired into all 7 nested-run sites and
fenced by an AST walk. **The standing invocation, which any future whole-suite count in this ledger must
name:**

```
cd .agentspace/rosetta-extensions/stack-core
STACKCORE_PROGRESS_LOG=/tmp/m257x-beacon.log \
  /usr/bin/python3 -m pytest tests/ -q --tb=line -p no:cacheprovider --no-header --durations=5
```

`/usr/bin/python3` is **3.9.6** and is the only interpreter on this host with pytest (`python3` on this
shell is homebrew 3.14 and has none) — pass 26's rule, restated because the count depends on it. Expected
wall time is recorded with the measurement in `progress.md` § iter-121; **a run far off it is an
unexplained measurement, not a faster one.** Progress goes to `/dev/tty` live, and to
`$STACKCORE_PROGRESS_LOG` for `tail -f` or a no-tty host.

**THE MEASUREMENT, and it is the first whole-suite total this ledger has been able to quote:**

```
1 failed · 1125 passed  in  1032.57 s (0:17:12)      rext 1bb64c3 · 239 beacon lines
```

The one failure is `test_claim_twin_guard_iter48_answer_key::test_02` — the standing, documented RED,
**re-attested by a full run** rather than carried as *"unchanged, not re-verified"* for a fourth
consecutive pass. **Expected wall ~1030 s**, which reconciles with iter-111's 1090.88 s to within 6 %.

**The check rule 51(b) asks for did not pass on the first attempt, and that is the whole point of it.**
Today's FIRST run returned **632.20 s with 4 failures**, and the 460 s gap *was* the defect: the
mechanical-fences battery had been dying on its baseline since iter-117 (`corpus_citation_guard.py` in the
participating baseline but not in `_COPY_FILES`), and a battery that dies on its baseline never runs its
mutants — `test_01_every_mutant_matches_its_DECLARED_verdict` cost **45 s** broken and **405.69 s** fixed.
Proven pre-existing at `b9bb2b6` in a detached worktree, since removed. **A fast suite is not evidence of
a healthy one**, and a whole-suite claim that had only ever been quoted once would have banked the 632 s.

**Guard family, re-run after this iter's corpus edits: 21 members · 17 GREEN · 0 RED · 4 not-run** (the
commit-/input-scoped members). Still **not a whole-family green**, and the summary line still says so.

---

## Pass 27 — 2026-08-07 — incremental

**Iters hardened this pass:** iter-121 … iter-131 (11 closed iters; iter-132 was in-flight and excluded).
**Tiks covered since prior pass:** 12.
**Scope:** the 25 `rosetta-extensions` files changed since pass 26's terminating commit `b9bb2b6` —
3 net-new guards (`blocking_state_guard`, `claim_census_guard`, `tests/progress_beacon`), 6 modified
(`anchor_construct_guard`, `gen_override`, `guard_family`, `platform_alignment_guard`,
`unreadable_repo_claim_guard`, `gen_injected_override`) and one Go comment repair (`isolation.go`).

### Three defects, all live, none caught by the controls that shipped with them

**1. `blocking_state_guard` — "represented" was two independent whole-document substring searches.**
The fence exists because a sweep keyed on ONE grading field reported zero and was wrong. Its own
representation test was `it in body and field in body` over a 60 KB audit, so **the fence built to
close *green over something it never checked* had the same defect one level in**. Two false greens,
both constructed:

  * **association never checked** — an audit naming `iter-19` for an unrelated reason, with `re-scope`
    in a paragraph about `iter-07`, satisfied both conjuncts → `represented: True`, exit 0;
  * **substring iter ids** — `"iter-11" in body` is satisfied by `iter-119`. Measured on this
    milestone's own `deferrals-audit.md`: **`iter-10`, `iter-11` and `iter-12` are absent from the
    document and read as present**, purely by collision with `iter-10x`/`11x`/`12x`.

Fixed as block-scoped association (`represented_in`) over a word-boundaried id (`iter_token`). Block,
not line, so a heading-plus-prose entry still reads GREEN — **all 8 live blocking pairs associate at
both granularities**, so the rule costs no real green and the real tree stays exit 0. (`4ca7670`)

**2. `claim_census_guard` — the ratchet went RED for a corpus file with ZERO defects.**
The `[new-file]` arm of C2 had **no control at all** — present in the guard, absent from its 21 tests —
and was also wrong. `per_file` carries an entry for every in-scope file (`setdefault(rel, 0)`), so a
brand-new file whose every assertion is cited produced `[new-file] x.md: 0 unevidenced assertions` and
exit 1. C2 asserts a count *RISES*; 0 has not risen. **Worse than an ordinary false positive because of
what it teaches:** the operator's only remedy is `--update-baseline`, which re-seals every *other*
file's debt at its current level. A guard whose false alarm is cured by disarming the guard trains the
operator to disarm it — and iter-123 added two corpus files, so the next one would have fired it.
Latent, not live: the shipped baseline covers all 41 in-scope files, real tree exit 0 before and after.
(`c8ec339`)

**3. `gen_override` — the dev emitter deleted postgres's data mount, and the comment said it couldn't.**
The severe one. iter-129 fixed a real defect (the `$HOME/.aws` mitigation was keyed on the DELETED
`jobsimulation` literal and had gone dead on `backend`) and introduced a worse one doing it:

  * **the predicate was widened, not mirrored.** The demo twin matches the **unexpanded literal**
    `^($HOME|${HOME}|~)/` — *"Unexpanded on purpose"*, `platform_topology.py:84` — a precise proxy for
    *the compose author wrote a path into the operator's home directory*. The dev twin reads the
    **resolved** compose, where that literal is gone, and tested `src.startswith(home + "/")`. **That is
    a different predicate: *anything under the user's home directory*, which on a normal dev box is the
    entire workspace.** `common.yml` gives `postgresql` one volume, `./data/postgresql`, which resolves
    `$HOME`-rooted → `postgresql` was selected.
  * **the comment asserted a guard the code did not have.** *"(`postgresql` is handled above and never
    reaches here with a home bind.)"* — there was no `elif` and no `continue`. Step 2b overwrote the
    per-stack data bind step 2 had just repointed with `[]`, which `to_yaml` emits as
    `volumes: !override`: **the mount REMOVED.** Postgres runs on the container's writable layer, every
    byte lost on recreate, and the per-`dev-N` data-root repoint silently does nothing.

Measured: predicate selected `{postgresql, backend}`, override emitted `postgresql: {volumes: []}`.
Both halves fixed independently. Live on the `/dev-up N` path — but **`415240f` is in NO TAG**, so no
stack consumed it. **Caught before it shipped.** (`9ab7590`)

### Coverage delta on touched files

| module | before | after |
|---|---|---|
| `tests/test_blocking_state_guard.py` | 10 | **17** |
| `tests/test_claim_census_guard.py` | 21 | **24** |
| `tests/test_gen_override_home_binds.py` | 5 | **16** |
| in-scope scoped invocation | 139 | **160** |

**Tests added:** +21 across 3 modules (7 · 3 · 11).
**Bugs surfaced + fixed inline:** 3 (`4ca7670`, `c8ec339`, `9ab7590`).
**Flakes stabilized:** none found.
**Regression check:** `stack-core/tests/test_gen_override.py` 22 passed (the base override suite the
`elif` could have broken) and the whole `dev-stack` suite **151 passed in 100.46 s** — gen_override is
dev's shared engine, so its consumer suite is the one that had to stay green.

### Every fix carries a mutation control that ISOLATES it (§5 rule 53), verified by running the mutant

| mutant | kills | leaves standing |
|---|---|---|
| restore the whole-document conjunct | the association arm | the collision arm |
| boundary-free `iter_token` | the collision arm | the association arm |
| drop the `cnt > 0` guard | `test_15` clean-new-file | `test_16` dirty-new-file |
| delete the `[new-file]` branch | `test_16` dirty-new-file | `test_15` clean-new-file |
| predicate → *anywhere under $HOME* | the 4 predicate arms | the postgres end-to-end arm (the `elif` saves it — which is what makes the `elif` genuine defence in depth, not a duplicate) |
| `elif` → `if` | ONLY the defence-in-depth arm | everything else |

**`test_17` is the pass's other kind of finding, and it is a coverage one, not a bug:** the census
guard's 21 controls all fixture their baseline via `monkeypatch`, so **the one artifact the guard
actually ships — `claim_census_baseline.json` against this repo's own `corpus/` — was asserted by
nothing in the suite**, only by running `guard_family` by hand. Now pinned, with an anti-vacuity floor
(≥ 30 files, ≥ 500 tier-1 pairs) and a check that no in-scope file is missing from the baseline.

**Stop condition: continue-to-next-pass** — three first-order defects in one pass is not a plateau, and
the dimension scan has not yet come up empty on the remaining modified guards
(`unreadable_repo_claim_guard`, `platform_alignment_guard`, `anchor_construct_guard`, `progress_beacon`).

---

## Pass 28 — 2026-08-07 — incremental

**Iters hardened this pass:** iter-121 … iter-131 (same scope, second pass).
**Scope this pass:** the modified guards pass 27 had not yet attacked — `unreadable_repo_claim_guard`,
`platform_alignment_guard`, `anchor_construct_guard`, `tests/progress_beacon`.

### The finding: pass 27's defect shape was not a one-off

`unreadable_repo_claim_guard` gained a second satisfaction route at iter-123 — a paragraph may report a
**ref-pinned reading** of `infrastructure` instead of hedging about it. The test was

```python
any(p in block for p in MEASURED_PHRASES) and bool(_SHA.search(block))
```

**two independent whole-paragraph searches that never have to co-refer** — the *identical* shape
`blocking_state_guard.represented_in` carried, found one pass earlier in a different fence. Constructed
and confirmed: a paragraph naming `infrastructure` as an **English noun** (*"the observability
infrastructure around it is unchanged"*) and carrying the **platform's** sha about a different fact
(*"deleted at `838d907`"*) satisfies both halves and reads as a ref-pinned measurement — an unmeasured
claim through the very boundary the fence exists to hold.

**A character window would not have separated them, and would have been fitted rather than derived:**
the false green sits ~70 characters apart and the **widest true site at 61**. Any constant between the
two is tuned to the fixture — precisely the "control aimed at something adjacent to its claim" that §5
rule 53 names. What actually distinguishes them is that **a citation NAMES ITS SUBJECT**, and both real
idioms were measured off the live corpus rather than invented:

  * the repo as a **backticked code artifact** — 12 of the 13 live sites;
  * bound to its ref by **`@`** — the 13th, inside a mermaid diagram comment where backticks are not
    the idiom. (A backticks-only rule would have false-RED'd it; that is why the second clause exists
    and why a control fails for it alone.)

An English noun does neither. Live corpus unchanged: 23 mentions, 9 hedged + 13 measured, exit 0 before
and after. (`be5b21f`)

### The pattern was then SWEPT, not assumed absent

Two instances in two passes justified asking whether the shape was endemic. A tree-wide grep for
co-presence conjunctions (`X in blob and Y in blob`, `.search(…) and .search(…)`, `any(…) and any(…)`)
across every non-test `.py` in `rosetta-extensions` returns **only the two docstrings describing the
fixes**. So: two instances, both closed, no third — and the sweep is on the record so a future pass
does not re-derive it.

### The remaining in-scope guards were attacked and HELD

* **`tests/progress_beacon`** (net-new, iter-121) — already carries pass 26's two lessons: an
  anti-vacuity floor on the AST site-walk, and `assertEqual(applied, 2, "MUTATION DID NOT APPLY")`
  before its mutant is read. Nothing to add.
* **`platform_alignment_guard` assertion G** (net-new, iter-130) — checked for vacuity, and it is not:
  *"read 2 go.mod file(s), 5 org module require(s) (analytics-go, colony, proto, storage, taxonomy),
  and graded 4 library row(s)"*. Real subjects on both sides. The `unclonable` vs `unresolvable` split
  (iter-126) prints its 10 unread citations by head, which is disclosure rather than a silent pass.
* Bad and missing argv both exit **2**, not 0 — spot-checked because an earlier reading of `EXIT=0`
  turned out to be `tail`'s exit code, not the guard's. Worth recording: a piped exit code is not the
  guard's exit code.

### Coverage delta on touched files

| module | before | after |
|---|---|---|
| `tests/test_unreadable_repo_claim_guard.py` | 18 | **25** |

**Tests added:** +7. **Bugs surfaced + fixed inline:** 1 (`be5b21f`).
**Flakes stabilized:** none found. **Flake gate:** 3 consecutive runs of the pass-27+28 modules →
**104 passed** at 1.36 / 1.28 / 1.30 s.

**Stop condition: continue-to-next-pass** — the dimension scan still produced a first-order defect, so
no plateau has been measured. `guard_family` and `anchor_construct_guard` remain un-attacked, and no
whole-suite total has been taken since the fixes landed.

---

## Pass 29 — 2026-08-07 — incremental

**Iters hardened this pass:** iter-121 … iter-131 (same scope, third and final pass of this invocation).
**Scope this pass:** `guard_family`, `anchor_construct_guard`, the cross-section emitter fence in
`stack-injection`, the Go `isolation` package, and — the pass's main instrument — **a whole-suite run**,
which is the only thing that could have found either of this pass's two defects.

### Two more live defects, and BOTH were standing REDs nobody was told about

**4. `stack-injection` — a cross-section fence RED for three iters, asserting the literal iter-129 deleted.**
`TestDevEmitter::test_dev_emitter_resets_jobsimulation_volumes` asserted the DEV emitter's source still
contained `if name == "jobsimulation":`. iter-129 removed that literal — correctly — and did not update
the test. **Proven pre-existing at `f2ea567` in a detached worktree** (since removed), so it is
iter-129's, not this session's.

What makes it worth more than a line: **the demo twin's sibling test, three functions above in the same
file, was converted at iter-88 with the reason spelled out** — *"a test pinned to the implementation's
service LITERAL is a second copy of the thing that went stale … it would have kept the dead literal in
place by failing if anyone removed it."* This one was left behind, and then did **exactly that**. The
warning was already on the page; only one of the two twins acted on it. Now the same shape as its twin,
plus a net-new arm for pass 27's `elif` finding — which belongs here as well as in stack-core, because
this module is the fence whose job is keeping the two emitters in step. (`6ad8866`)

**5. `baseline_mirror_fence` — RED since iter-129, and it is the *provenance* rule that broke.**
iter-129 moved M255's provenance block out of `state.md` — correctly; *a measurement whose provenance
lives in an index that gets rewritten is one close away from being un-reproducible* — and in the move
left `666.29 s` and `658 / 666 / 672 s` in a bullet saying only *"the host"*. The section **heading**
names `billion`; the fence's naming lookback **resets at a blank line**, so the heading does not reach
the bullet. `bbdbd61` had already done this exact repair release-wide under the title *"name the host on
every baseline number"* — **the rule was in place, and a well-intentioned move re-opened it.**
(`b911b77`)

### THE MEASUREMENT — the whole-suite total, with its invocation (§5 rule 51(b))

```
cd .agentspace/rosetta-extensions/stack-core
STACKCORE_PROGRESS_LOG=/tmp/m257x-beacon-p29.log \
  /usr/bin/python3 -m pytest tests/ -q --tb=line -p no:cacheprovider --no-header --durations=5
```

```
2 failed · 1202 passed  in  1243.66 s (0:20:43)     rext be5b21f · 131+ beacon lines
```

**The wall time is 1243.66 s against an expected ~1030 s, and rule 51(b) says that is an unexplained
measurement until it is explained.** Two accountable causes, both real: (a) the suite **grew** — 1204
collected against iter-121's 1126, +78, of which 28 are this session's; (b) the run was **not alone on
the host** — `dev-stack` (100.5 s), `stack-injection` twice, `guard_family` and four scoped batteries ran
concurrently with it. The signature is in the durations: the mechanical-fences battery took **524.13 s
here against 405.69 s at iter-121**, +118 s on identical work. So: contention plus growth, not a
regression — but it is stated rather than waved past, and **the next whole-suite claim should be taken
on an otherwise-idle host** if the number is to be compared to iter-121's.

**The two failures were BOTH triaged, and neither is a mystery:** one is the standing, documented
`test_claim_twin_guard_iter48_answer_key::test_02`; the other is defect 5 above, **closed in this pass**.
`test_baseline_mirror_fence` re-run scoped after the fix: **28 passed**. **No post-fix whole-suite total
is quoted, because none was taken** — the scoped re-run is what this entry claims and all it claims.

### The remaining in-scope surfaces were attacked and HELD

* **`guard_family`** — the registry is DERIVED (`guard_dir.glob("*_guard.py")`) and `reconcile()` checks
  it in **both** directions, with `EXIT 2` teeth: a guard on disk with no invocation, or an invocation
  with no guard, stops the family. 22 members, all placed. A guard cannot land unregistered and be
  silently un-run.
* **`anchor_construct_guard`** — iter-121's `KNOWN_WEAKNESS` disclosure is genuinely **wired into
  `--json`**, verified by parsing the document rather than by reading the source.
* **`isolation.go`** (iter-130, comment-only) — `go vet` clean, `go test ./isolation/...` ok.
* **Guard family, whole-family run:** 22 members · **18 GREEN · 0 RED** · 4 not-run (the
  commit-/input-scoped members needing `--range`/`--ledger`). Still **not a whole-family green**, and the
  runner's own summary says so.

### Session totals (passes 27–29)

**Tests added:** **+29** — `test_gen_override_home_binds` 5 → 16, `test_blocking_state_guard` 10 → 17,
`test_unreadable_repo_claim_guard` 18 → 25, `test_claim_census_guard` 21 → 24, plus the
`stack-injection` cross-section arms.
**Bugs surfaced + fixed inline: 6** — `4ca7670`, `c8ec339`, `9ab7590`, `be5b21f`, `6ad8866`, `b911b77`.
**Flakes stabilized:** none found; flake gate 3/3 clean at 104 passed.
**Every behavioural fix carries a mutation control that ISOLATES it**, each verified by running the
mutant and confirming it kills that arm **and no other**.

**Stop condition: cap reached without stabilization** — the 3-pass incremental cap fired and pass 3 still
produced two first-order defects, so the dimension scan never came up empty.

### The cap has now fired without stabilizing FOUR consecutive times (22, 25, 26, 29) — and the supply CHANGED

Pass 26 recorded three consecutive caps and named the remedy: *"not more passes of the same kind."*
That was acted on — this session attacked mechanisms rather than re-running suites — and it worked, but
**what it found is not what the previous three found, and that difference is the finding.**

Passes 22/25/26 found **one** class: a fence green over something it never checked. This session found
that class **twice more** (defects 1 and 4 by shape — `blocking_state_guard` and
`unreadable_repo_claim_guard` carried the *identical* two-independent-whole-block-searches conjunction,
in different fences, found one pass apart). **That pattern was then SWEPT tree-wide rather than assumed
closed** — a grep for co-presence conjunctions across every non-test `.py` in `rosetta-extensions`
returns only the two docstrings describing the fixes. Two instances, both closed, no third.

But **three of the six defects are a different class entirely, and it is the more expensive one:**

  * `claim_census_guard` — a **false RED** whose only remedy disarms the guard;
  * `gen_override` — a comment asserting a guard the code did not have, costing the dev **database its
    data mount**;
  * `baseline_mirror_fence` + the `stack-injection` twin — **standing REDs that ran for three iters with
    nobody told.**

The first two are *fences that harm*, not fences that miss. The last two share a single root: **nothing
runs the whole suite per-iter**, so a fence can go RED and stay RED across iters while every scoped run
an iter takes is green. iter-121 built the beacon precisely so a whole-suite run is *watchable*; what
this pass shows is that being watchable is not the same as being **watched**. Both of pass 29's defects
were invisible to every scoped invocation and fell out of the first whole-suite run in eight iters.

**That is an input to the scope decision the milestone is holding, not a request for a fourth pass.**
The cheapest change is not another harden pass: it is that **an iter's close should run the whole suite,
or the milestone should say out loud that it does not.**

**Knowledge backfill (`a027bd9`).** All three classes are promoted out of this ledger into
`corpus/ops/platform-alignment.md` §5, following pass 26's precedent with rule 53 — the ledger outlives
the pass, not the milestone:

* **rule 58** — a conjunction of two whole-document predicates is not an association check, *plus the
  trap*: do not reach for a character window, it is the obvious fix and it is fitted. Carries the
  tree-wide sweep result so no future pass re-derives it.
* **rule 59** — a fence that FALSE-REDs is worse than one that misses when its only remedy disarms it.
  The sibling of rule 8 facing the other way: 8 is a green that checked nothing, 59 is a red that found
  nothing.
* **rule 60** — nothing runs the whole suite per-iter, so a fence can go RED and stay RED; iter-121 made
  a 20-minute run *watchable*, and watchable is not *watched*.

Guard family re-run after the corpus edit: 22 members, **18 GREEN · 0 RED**, unchanged.

## Pass 30 — 2026-08-08 — incremental

**Iters hardened this pass:** iter-132 … iter-142 (11 closed iters; of them only **iter-132** and
**iter-142** touched executable code — `unreadable_repo_claim_guard.py` and the net-new
`retracted_pin_guard.py`. iters 133–141 are corpus markdown, so dimensions 2/3/5/6 have no surface
there and the carve-out applies.)
**Tiks covered since prior pass:** 11
**Scope boundary:** rext `6ad8866` (pass 29's terminating commit) → `f493615`. Exactly **5 files**
changed in that span, all in `stack-core` — which is *itself* a measurement, and it is the one that
made the demo-stack / stack-verify triage decidable (below).

### THE OWED WHOLE-SUITE RUN — and "whole suite" has meant ONE SECTION OF FIVE

iter-142 closed with `FIX-M257x-iter142-whole-suite-owed`. Taken here, on an otherwise-idle host,
**before any edit** (the pass then did read-only analysis for 35 minutes rather than contaminate it):

| section | result |
|---|---|
| `stack-core` | **1 failed · 1229 passed** (23:34) |
| `demo-stack` | 9 failed · 1038 passed · 11 skipped (3:37) |
| `dev-stack` | 151 passed (1:38) |
| `stack-injection` | 335 passed (0:07) |
| `stack-verify` | 12 failed · 225 passed (6:04) |
| **total** | **2,978 passed · 22 failed · 11 skipped** |

**stack-core's single failure is the standing, documented `test_claim_twin_guard_iter48_answer_key::test_02`.**
So iters 132–142 introduced **no RED** — pass 29's two REDs stayed closed.

**The finding is the other 21.** Every prior "whole-suite" number in this ledger — including pass 29's
*"2 failed · 1202 passed"* — is **`stack-core` alone**: one section of five, 1,230 of 3,011 tests, where
*tests* means **executed = passed + failed + skipped**. The
milestone that exists to catch denominators stated a denominator that omitted **59 %** of its own
suite. The 21 failures live in sections nothing in this milestone has ever run.

> ⚠️ **This line read `1,230 of 2,989` until harden pass 39, and the way it was wrong is the finding
> iter-173 booked.** `2,989 = 2,978 passed + 11 skipped` — assembled from the five-section table three
> lines above, **dropping that same table's 22 failures**, so the denominator silently changed unit from
> *executed* to *passed-and-skipped* inside the entry whose subject is denominators. Re-derived from the
> table: `2,978 + 22 + 11 = 3,011`, of which `stack-core` is `1,229 + 1 + 0 = 1,230`. The **59 %** below
> is unaffected (`1 − 1230/3011 = 59.2 %`; the old operands gave 58.9 %) — which is exactly why it
> survived: *a percentage can survive an error its operands do not.* Repaired by
> `FIX-M257x-iter173-ledger-denominator`, owed to a harden pass because the iter that found it is
> contractually barred from editing this file.

**They are provably not ours.** `git diff --name-only 6ad8866..HEAD` returns 5 files, all `stack-core`;
`git log 6ad8866..HEAD -- demo-stack/` is empty. The failures are live-clone / live-container
assertions (demopatch sha pins against the mutable `stack-demo` working clones; `pg_isready` against a
running container) on a box with a `demo-1` stack up. **Pre-existing, environment-coupled, and out of
this pass's iter-diff scope** — routed, not fixed, as `FIX-M257x-h30-nonstackcore-suite`.

### Defects surfaced + fixed inline — 3, all in iter-142's net-new fence

**1. The census was LINE-scoped, and this corpus HARD-WRAPS — 10 live class members were never in the
denominator.** (`ccfd575`) iter-142 published *"50 live instances … the class stands at **0** and the
fence holds it there."* The 0 is real over the population it enumerated; that population excluded
every retraction clause straddling a soft line break, and at ~100 columns that is a large share of
them. Joining each line to the one above surfaces **10 more**, **hand-read 8 true / 2 false** (the two
falses are the same shape as iter-142's own two — a marker governing something other than the pin).
Every one still live *after* the repair that reported zero.

**The family had already learned this three times, and the document already said it in general form.**
`platform_predicate_guard._pin_window` has joined `line[i-1] + line[i]` since **iter-63** (and learned
the table-row exclusion there); `_NEGATED` needed the same widening at **iter-68**; a third predicate
records *"line-scoped it reached only 2 of the 4 — two of the live sites wrap"*; and **§7 rule 4 of
`platform-alignment.md` states it outright — *"the paragraph is the unit of publication."*** The
newest fence in the family was written after all four and shipped line-scoped anyway. The table-row
exclusion was re-derived here independently and landed exactly where iter-63 put it.

**2. The disclosed tuned constant was LOAD-BEARING, and the comment beside it said it was not.**
(`ccfd575`) `PATH_TIER_A_REACH`'s comment promised *"the path arm is SURVEY, never a gate — see
`main()`."* `main()` returned 1 on any finding at all, so a window fitted to **five sites** decided the
verdict while the code beside it denied it. **Not one of the 21 tests touched `main()`** — every one
drove `run()` — which is exactly how a verdict contract stays false in writing for a whole iter.
`Finding.gating` now implements the promise: the 44/44 bare single-line arm sets the exit code, the
path and wrapped arms print in full under a SURVEY heading. **Detection unchanged; only belief
narrowed.**

**3. A docstring asserting an exclusion the code never had.** (`ccfd575`) *"Struck-through /
code-fenced text is not prose"* — the fenced half is real, the struck-through half was never written.
Measured **0** sites corpus-wide, so it cost nothing; recorded because it is this milestone's **fourth**
catch of *comment asserts a guard the code does not have*, and corrected rather than quietly
implemented, with `DocstringAccuracy` pinning the real behaviour.

**Corpus side (`95983fd`):** rule 63(c‴) now says which population its 0 is over, and **§5 gains
rule 64** — *a fence over wrapped prose must state its line reach*, with the counter-half (the lesson
is not "always join"; the reach is a design parameter that must be CHOSEN AND STATED) and the reason
it recurred: **the remedy existed only as source comments inside one guard**, so nothing a new fence's
author reads carried it. *A defect class solved in code but not in the rulebook is unsolved.*

**Tests:** 21 → 45 on `retracted_pin_guard` (+24) — the whole `main()` verdict surface, the wrapped
arm's finds and its four must-not-join boundaries, the two measured falses as named regressions
asserting they can never reach the VERDICT, a meta-mutation proving the join can be switched off, and
a **structural** anti-vacuity floor (`joins_evaluated`, measured 1,451) that repairing prose cannot
drive to zero — deliberately not a finding-count floor, which is the ratchet-against-its-own-repair
defect iter-132 fixed elsewhere.

**Routed forward:** `FIX-M257x-h30-crossline-repair` — the 8 true sites across 6 files. Editorial prose
judgement in six documents is an iter's work, not a harden inline fix (Fate 3).

**Stop condition: continue-to-next-pass** — three first-order defects in one fence; the dimension scan
was nowhere near empty, and the obvious next question (how many sibling fences share the hole) was
unanswered.

## Pass 31 — 2026-08-08 — incremental

**Scope:** the rule-64 family sweep, plus the self-correction it forced.

### The sweep — all 22 guards classified by scanning unit

`unreadable_repo_claim_guard` already works in blank-line-delimited **paragraphs** and is immune.
`platform_predicate_guard` has joined since iter-63. The rest either do not associate two things
across prose or do so within a single token. **One real second instance: `clone_drift_guard`** — its
D2 rule wanted the `go.mod` citation and the module pin on **one line**, while the `staging-sync.md`
colony-requires sentence wraps between them. The site fell through a `continue` **silently**, three
lines below a comment reading *"Named, not silently dropped."* Fixing it moves that guard **3 → 4
graded sites**. (`0f446cf`)

**`prose_reach.py` is the fix to the fourth-encounter problem itself** — `continues_paragraph()` /
`join_prev()` as ONE unit with both consumers asserted (`SingleSource`) to hold the *same object*
rather than two copies that agree today, every iter-63 exclusion carried as a named test, and
deliberately not named `*_guard.py` so the family's derived registry excludes it by construction
rather than by a maintained exclusion list.

### ⚠️ THE SELF-CORRECTION — 4 of pass 31's 5 findings were PHANTOMS of pass 31's own bug

(`0a4fe0e`, corpus `9e5b207`.) The first cut reported four *further* ref-pinned sites and a relabelled
fifth. **All five were manufactured by the change.** The join fired whenever the current line lacked a
citation — including when it carried **no module pin at all**, including when it was **blank** — so a
citation was glued onto the empty line beneath it and reported as a site at a coordinate holding
nothing. Four of those, in four different documents, each with a plausible ref-pin reason attached.
They were written into a commit message *and* into rule 64 before anything questioned them.

**What caught it could not have been this guard.** Its mutation controls fired correctly and its own
tests passed — the code did exactly what it was told. **`anchor_construct_guard` caught it**, by
resolving two of the published coordinates and finding blank lines. A fence on a *different axis*
catching a fence-widening — the same shape as `repair_leak_guard` catching iter-142's path-arm gap,
and the argument for running the family rather than the change's own scoped suite.

**Corrected delta: exactly ONE site** (graded 3 → 4); unmeasured unchanged at 14. Two rules fall out,
both now in rule 64 and in the source:

  * `continues_paragraph` answers whether **PREV** continues; only the **caller** knows whether the
    current line has anything to continue INTO. **A widened reach must not invent subjects.**
  * **A measurement taken with a just-changed instrument is a claim about the instrument until
    something independent confirms it.** The phantoms had file names, line numbers and a mechanism.
    Nothing inside the changed guard could have told them from real findings.

Every docstring claiming the ref-pin-on-the-citation-line case was *measured* is corrected to say it
is guarded **prospectively** by a synthetic test — *"we handled a case"* and *"we found the case"* are
different claims and only one was earned.

**Tests:** +20 (14 `prose_reach`, 6 `clone_drift`), then +1 named regression for the phantom bug
(`test_a_join_never_fires_on_a_line_with_nothing_to_attribute`).

**Stop condition: continue-to-next-pass** — the in-scope iter-132 surface had still not been attacked
directly.

## Pass 32 — 2026-08-08 — incremental

**Scope:** iter-132's own change, attacked directly; then the gaps this session created.

### iter-132's `unreadable_repo_claim_guard` change HELD under every angle

The three-bucket tally (`hedged` / `mixed` / `measured`), the re-cut anti-vacuity floor and the
print order were all attacked and none moved:

* **The buckets are disclosed, not guessed.** `mixed = marked AND measured` is a conjunction of two
  whole-paragraph predicates — the shape §5 rule 58 warns about — but it is *reported as its own
  bucket with a `KNOWN_WEAKNESS` line saying the guard cannot tell a quoted retraction from a live
  hedge*, which is the sanctioned disposition (`D-M257x-121-4`: when a distinction is not mechanical,
  disclose it in the instrument rather than guess it). Not a defect.
* **The floor spans all three buckets** (`hedged + mixed + measured >= 22`, actual 27), so a
  legitimate repair that retires a hedge cannot force a re-cut — the *ratchet-arguing-against-its-own-repair*
  defect iter-132 fixed. Correctly generalised.
* **The NOTE is gated on `measured and hedged`, NOT nested under `if mixed:`** — which its diff hunk's
  indentation makes it look like it is. Read the file, not the hunk. Verdict prints last, per
  `guard_family.run_one` reporting `lines[-1]`. Correct.
* **§5 rule 64 does not reach it**: it already works in blank-line-delimited paragraphs.

**No finding.** The first empty in-scope dimension scan in five harden sessions.

### What pass 32 DID find is in code THIS SESSION wrote (`7d986f6`)

Pass 30 split the gate from the survey and added `gating_findings` / `survey_findings` / the wrapped
census keys to `--json` — **with no test whatsoever**. That is the *identical* omission that let the
text-mode contract be false in writing for a whole iter: 21 tests, all against `run()`, none against a
verdict path. A harden pass reproducing the defect it just fixed, one layer over.

Six tests. The one that matters is `test_a_survey_only_run_is_exit_0_with_a_NON_EMPTY_findings_list`
— **findings present, verdict green**. A consumer reading `findings` as the gate would call that run
RED while the exit code says otherwise, and nothing before now would have caught the disagreement.
Plus a partition assert (the halves sum to `findings` — no double-count, no loss), the arms named in
the document rather than only in the source, and `wrapped` on every finding so a survey hit is
attributable to the arm that produced it.

### THE MEASUREMENTS — counts, not wall-time (§5 rule 51's timing leg is unusable on this host)

**Post-change `stack-core` whole suite:**

```
1 failed · 1280 passed  in  1275.99 s (0:21:15)
```

versus the pre-change baseline of **1 failed · 1229 passed** — **+51 tests, the same single failure**,
the standing documented `test_claim_twin_guard_iter48_answer_key::test_02`. **No cross-test breakage
from four changed source files.**

⚠️ **Disclosed confound, and its resolution.** A ledger append landed ~1 minute into that run, so
tests reading `knowledge/plan/**` saw a mixed tree — a confounded measurement by this milestone's own
standard. Resolved rather than waved past: the ten `knowledge/plan`-reading test files were re-run on
a **fully stable tree** and reproduce **1 failed · 162 passed** — the same standing failure, nothing
else. The whole-suite number above stands for the surfaces it is quoted for.

**Flake gate:** 3 consecutive clean runs, **98 passed** each, across all three touched test files.
**Guard family:** 23 members — **17 GREEN · 0 RED** · 6 not-run (commit-/input-scoped, needing
`--range`/`--ledger`/`--platform`). Still not a whole-family green, and the runner's own summary says
so.

### Session totals (passes 30–32)

**Tests added: +51** — `retracted_pin_guard` 21 → 51, `prose_reach` 0 → 14 (net-new), `clone_drift_guard`
26 → 33.
**Bugs surfaced + fixed inline: 5** — `ccfd575` (×3: the line-scoped census, the load-bearing "survey"
constant, the unimplemented docstring exclusion), `0f446cf` (clone_drift's silently-dropped wrapped
sites), `0a4fe0e` (the phantom-site bug this session introduced and retracted).
**Flakes stabilized:** none found; gate 3/3 clean.
**Knowledge backfill:** `95983fd` + `9e5b207` — rule 63(c‴) corrected to state its population, **§5
rule 64** net-new with the sweep result, the counter-half, and the two sub-rules the self-correction
produced.
**Routed forward:** `FIX-M257x-h30-crossline-repair` (8 true wrapped sites, 6 files);
`FIX-M257x-h30-nonstackcore-suite` (21 pre-existing failures in demo-stack + stack-verify).

**Stop condition: cap reached without stabilization** — the 3-pass incremental cap fired. Coverage
delta is not under 2 % (a previously-0 %-covered verdict path went to six tests), so the mechanical
condition is not met.

### The cap has now fired without stabilizing FIVE consecutive times (22, 25, 26, 29, 32) — but the STREAM changed, and that is the finding

Passes 22/25/26/29 each found first-order defects in **shipped** code and kept finding them to the
last pass. This session did too — in passes 30 and 31. **Pass 32 did not.** The in-scope iter surface
(iter-132's guard, iter-142's fence as repaired) came up **empty** for the first time; the only gap
pass 32 found was one *pass 30 had created three hours earlier*.

That is a different failure mode from "the iters need rework", and it names the real bottleneck:
**this session's dominant defect source was the harden pass itself.** Two of the five bugs
(`0a4fe0e`'s phantom sites, `7d986f6`'s untested verdict path) were introduced by passes 30 and 31,
and both are the *same shape as the defect being fixed* — a widened reach that invented subjects while
fixing a reach that missed them; an untested verdict path added while fixing an untested verdict path.

**The transferable rule is already booked as rule 64's second sub-rule** — *a measurement taken with a
just-changed instrument is a claim about the instrument until something independent confirms it* — and
the mechanism that saved both was **another fence on a different axis** (`anchor_construct_guard`
resolving published coordinates to blank lines), not the changed guard's own controls, which passed
throughout.

**So the recommendation is unchanged from pass 29 and now better evidenced: not a fourth pass.** Pass
29 asked that an iter's close run the whole suite, or that the milestone say out loud that it does
not. This session adds the sharper half: **"the whole suite" in this ledger has always meant
`stack-core` alone — one section of five, 1,281 of 3,062 tests** (*tests* = executed = passed + failed
+ skipped; this line read `1,280 of 3,040` until harden pass 39 — it carried pass 30's dropped-failures
hole forward as `2,989 + 51 = 3,040` and dropped this section's own single failure from the numerator.
`1,280/3,040` and `1,281/3,062` are **both 42 %**, which is how it survived 28 iters). 21 failures sit in sections no
harden pass or iter close has ever executed, and they were invisible for the whole milestone. Deciding
what "the suite" means is a scope call for the milestone, not something a fourth harden pass can fix.

---

## Pass 33 — 2026-08-08 — incremental

**Iters hardened this pass:** iter-143 … iter-152 (scope shared across passes 33–35)
**Tiks covered since prior pass:** 10

**Scope manifest.** 23 files across five rext sections, plus 7 corpus/skill files in `rosetta`. The
production surface is small and new: one net-new 354-line fence (`stack-core/service_registry_guard.py`,
iter-152), two amended guards, two shell entry points (`demo-stack/rosetta-demo`, `dev-stack/dev-stack`,
iter-147), and `stack-verify`'s report driver (iter-148). Six net-new test files. Everything else is
milestone bookkeeping.

**Target: iter-152's `service_registry_guard.py`** — the newest fence, and the one with the largest
untested surface. It grades `stack-verify/lib/services.sh` (the table that decides what every stack is
probed on) against the platform's own compose, in four arms and both directions.

**Finding — the fence blamed the registry for its OWN parse failure.**

It reaches `postgresql` and `redis` only by following compose's `include:` one level into `common.yml`.
The include parser understood exactly one shape, scanned from line 0, and `break`ed on the first
non-indented line. Measured against the real platform compose at `0c91421`, **four realistic file shapes
made it resolve zero includes**:

| shape | pre-fix | post-fix |
|---|---|---|
| a comment line above `include:` | exit 1, 2 × A/DEPARTURE | exit 0, aligned |
| the flow list `include: [common.yml]` | exit 1, 2 × A/DEPARTURE | exit 0, aligned |
| the long form `- path: common.yml` | exit 1, 2 × A/DEPARTURE | exit 0, aligned |
| an included file not on disk | exit 1, 2 × A/DEPARTURE | **exit 2, cannot-measure** |
| `include:` with no entries | exit 1, 2 × A/DEPARTURE | **exit 2, cannot-measure** |
| no `include:` at all | exit 1 (honest) | exit 1 (unchanged) |

In none of the four did it go quiet. It printed, with full confidence, *"A/DEPARTURE: registry row
'postgresql' is not a service in the platform compose"* — and its remedy, *"add it to
`SERVICES_NOT_IN_PLATFORM_COMPOSE`"*. **Following that remedy would permanently stop grading two rows
that were never wrong, in response to a bug in the guard.** This milestone's founding class run
backwards: an INSTRUMENT failure presented as a SUBJECT finding, carrying an instruction to disarm the
subject.

**Three false claims removed, all about this same leg.** The module docstring said a failed include
"would silently miss both and its arm-**B** denominator would be wrong by two" — the miss is loud and it
lands on arm **A**. `TestIncludeIsFollowed`'s docstring said the same miss would "still report ALIGNED",
refuted from birth by the test directly beneath it asserting `rc == 1`. And `parse_compose`'s closing
comment described a "drop anything with no ports AND no build/image" filter that **was never written**,
solving a problem that **does not exist** — the top-level reset already keeps `networks:`'s children out.
The real mechanism is now pinned by a test so it cannot be deleted on the strength of a comment that
never matched it.

**Coverage delta:** `test_service_registry_guard.py` 18 → 28 tests; the include leg 2 → 9.
**Tests added:** iter-152 → `stack-core/tests/test_service_registry_guard.py`: 10 (7 edge-case, 3
error-path).
**Bugs surfaced + fixed inline:** 1 behavioural + 3 documentation (`e5c0dda`).
**Controls:** the 10 net-new tests were run against the PRE-FIX guard — **8 of 8 executed cases fail**
(7 failures + 1 error). `§5` rule 64's second sub-rule satisfied: the instrument is proven independently
of the change.

**Stop condition: continue-to-next-pass** — a first-order defect in shipped code on the first pass; the
in-scope surface is not exhausted.

---

## Pass 34 — 2026-08-08 — incremental

**Iters hardened this pass:** iter-143 … iter-152 (target: iter-148 + iter-151)

**Method change that produced the finding.** Every existing arm of `test_probe_scope_m257x.py` reads
`generate.sh`'s **source**. Nothing had ever read the **artifact** — the markdown a human actually grades
a stack on. Running it takes 0.05–0.24 s per report; the gap was never cost, only method.

**Finding 1 — the report never named the stack it graded.**

`STACK_PROJECT` and `STACK_OFFSET` are the two stack-TARGETING variables (iter-151 measured them as
read-side only). `target.sh` defaults them to project `anthropos` at offset 0 — the main dev stack.
iter-148's own SKILL.md note tells the operator to pass them for a demo or a `dev-N`, and warns that
without them "the probes go to project `anthropos` on base ports, which is the main dev stack, not the
one you meant."

**Neither variable appeared anywhere in the report.** Measured: a run against `demo-1 @ 10000` and a run
against `demo-2 @ 20000` produce markdown that is **byte-identical apart from the timestamp**. The
failure output of the new control is the finding — two complete reports, character for character the
same, about two different stacks. A forgotten `STACK_PROJECT` was unrecoverable from the artifact.

Now a `**Target**:` line in the header; forgetting **both** variables says `DEFAULTED` out loud rather
than reading like a deliberate main-dev-stack run.

**Finding 2 — the scope line covered two branches of three.**

`scope_note` was set on the DERIVED branch and the UNDERIVABLE branch, and left **empty** when the caller
supplied `STACK_SERVICES` — so no scope line reached the report at all. That is the branch the skill
actively recommends ("an explicit `STACK_SERVICES` always wins") and the branch where the scope is
**arbitrary**: `✓ pass` off a hand-narrowed one-probe run was indistinguishable from a full sweep. The
site's own comment claimed *"the probe SCOPE rides in the report"*; it rode in 2 of 3.

**A fence of iter-148's own went RED for a restructure that preserved its property exactly.**
`test_a_caller_supplied_scope_is_never_overridden` required the LITERAL `[ -z "${STACK_SERVICES:-}" ]`.
Making the caller-supplied branch disclose turned it into `if [ -n … ]; then <caller> else <derive> fi` —
same property, different spelling. That is `§5` rule 67/68(d)'s axis **pointed at itself**, inside the
fence written to apply it. It now asserts the ORDER (an emptiness test, then the derivation after it),
and the property gained a behavioural arm: a caller's scope survives even with a real platform clone
under `STACK_ROOT` that the derivation would otherwise have replaced it from.

**Coverage delta:** `test_probe_scope_m257x.py` 6 → 14 tests; behavioural (artifact-reading) arms 0 → 8.
**Tests added:** iter-148 → `stack-verify/tests/test_probe_scope_m257x.py`: 8 integration.
**Bugs surfaced + fixed inline:** 2 (`ccffc69`) + the corpus-side doc promise (`420137c` in `rosetta`).
**Controls:** the 8 net-new tests run against the PRE-FIX `generate.sh` — **6 of 8 fail**. The 2 that
pass are the derived/underivable disclosures, which are regression pins, not gap proofs; recorded as
such rather than counted toward the delta.

**Stop condition: continue-to-next-pass** — two more first-order defects in shipped code; the surface is
still producing.

---

## Pass 35 — 2026-08-08 — incremental

**Iters hardened this pass:** iter-143 … iter-152 (target: iter-148's derivation + verification)

**Finding — iter-148 fixed the OVER-broad scope and left the UNDER-broad one.**

`generate.sh` derives its probe scope from `$STACK_ROOT/platform/docker-compose.yml` — the platform's
**unmodified** compose. Measured at platform `0c91421`:

```
platform_topology.py services --platform-dir stack-demo/platform
-> postgresql redis sentinel backend gotenberg          (five)
```

A default demo runs **eight** of the registry's rows: those five plus `next-web-app`, `studio-desk` and
`directus`, which `up-injected.sh` adds through the stack's own **generated override** — a file the
derivation never reads. The UI tier is opt-**OUT** on a demo (`DEMO_NO_UI`), so this is the default
shape, not a corner. Three services a presenter cares most about sit outside the scope while the report
reads `✓ pass`.

Under-broad is the quieter of the two failures iter-148 sits between. Over-broad printed four false
`down`s — loud, and it got fixed within one iter. Under-broad prints **nothing at all**, and had it not
been measured here it would have read as a clean bill of health indefinitely. `§5` rule 60/66 exactly: a
scoped green is evidence about its scope alone.

The derived disclosure now names what it leaves out. **Unioning the stack's generated override into the
derivation is the real fix and was NOT done** — it reaches from `stack-verify` into `demo-stack`'s
per-stack layout and cannot be verified without a live demo, which is past the inline-fix boundary.

**Coverage delta:** `test_probe_scope_m257x.py` 14 → 16.
**Tests added:** 2 — one of which **asserts the measurement itself** rather than trusting the commit
message: it runs the real derivation against a real platform clone and fails if any of the three UI rows
ever enters the derived set, so the disclosure cannot outlive the gap it describes.
**Bugs surfaced + fixed inline:** 1 (`95ef640`).
**Routed forward:** `FIX-M257x-h33-derive-includes-stack-override` — union the stack's generated override
compose into `generate.sh`'s derivation, so a demo's UI tier is probed rather than disclosed away.
Cross-section; needs a live demo to verify. Fate 3.

### Two self-inflicted defects, and what caught them

**(a) This harden pass hid 10 of its own tests behind the `__main__` guard.** Both new classes were
appended with `cat >>`, landing *below* the `if __name__ == "__main__"` block already at the end of the
file. Under pytest they ran — 16 passed, the number quoted in two commit messages. Run the file the way
its own docstring says to (`python3 tests/test_probe_scope_m257x.py`) and it collected **6**, printed
**OK**, and said nothing about the 10 it skipped. A green that silently drops most of its subjects, in
the tests written to stop exactly that.

**(b) A test of mine graded the invocation, not the code.** `TestNetworkKeysAreNotServices` did a bare
`import service_registry_guard`, which resolves only when pytest's rootdir puts `stack-core` on
`sys.path`. It passed the 3× flake gate **and** the full 1,338-test `stack-core` suite — both of which
run from inside `stack-core` — and failed the first time anything ran it from one directory up. Now
imported by path; verified green from the rext root, from `stack-core`, and by direct execution.

**Neither was caught by any control this session wrote.** (a) was caught by
`test_test_collection_fence.py` — a **different fence, on a different axis, in a different section**.
(b) was caught by **changing the working directory**. All three of this session's passes were green on
their own terms throughout, exactly as passes 30/31 were when they introduced their own.

**Flake gate:** 3 consecutive clean runs of both new test files — `test_probe_scope_m257x` 16 passed ×3,
`test_service_registry_guard` 28 passed ×3.

**Knowledge backfill:** `.claude/skills/test-platform/SKILL.md` — the iter-148 note promised
*"prints the scope into the report"* unqualified while recommending the branch that printed nothing, and
warned about a forgotten `STACK_PROJECT` while the report it produced named no stack. Both now true in
the code and the doc states which branches it covers.

### Suite results (counts, never wall-time — `§5` rule 51's timing leg fails on this host)

| suite | result | baseline |
|---|---|---|
| `stack-core` | **2 failed · 1338 passed** | pass 32 recorded 1 failed · 1229 passed |
| `stack-verify` | **0 failed · 252 passed** | iter-148 recorded 244 passed |
| targeted (registry guard + probe scope + collection fence) | 0 failed · 60 passed | — |

`stack-core`'s two failures are the standing `test_claim_twin_guard_iter48_answer_key::test_02` **and
defect (a) above**, which the run caught and which is now fixed — so the residual is the one standing
failure, unchanged. `stack-verify`'s 252 = 244 baseline + this session's 8; **zero regressions**.

**Guard family** (`--platform stack-demo/platform`): **20 GREEN · 0 RED · 4 not-run**
(`anchor_offset_guard`, `repair_leak_guard`, `repair_reach_guard`, `value_change_guard` — all
commit-/input-scoped, needing `--range` or `--ledger`). `service_registry_guard` itself: **ALIGNED**, 12
registry rows (7 graded, 5 declared absent) vs 7 compose services publishing 10 host ports.

### Session totals (passes 33–35)

**Tests added: +20** — `test_service_registry_guard` 18 → 28, `test_probe_scope_m257x` 6 → 16.
**Bugs surfaced + fixed inline: 4 behavioural + 4 documentation**, across `e5c0dda`, `ccffc69`,
`95ef640`, `420137c` (rosetta), plus `622b1cf` and `f7c7ace` repairing this session's own two.
**Flakes stabilized:** none found; gate 3/3 clean on both files.
**Routed forward:** `FIX-M257x-h33-derive-includes-stack-override`.

**Stop condition: cap reached without stabilization** — the 3-pass incremental cap fired. Coverage delta
is nowhere near under 2 % (the include leg went 2 → 9 tests; the report driver gained 8 artifact-reading
arms where it had none), and the Phase 2 dimension scan found something first-order on **every** pass.

### The sixth consecutive cap-without-stabilization — and the stream is back to what 30/31 saw

Pass 32 reported the cap firing for the fifth time and observed something new: its in-scope iter surface
came up **empty**, and the only defect it found was one an earlier pass in the same session had created.
It framed that as the stream changing.

**This session does not reproduce that.** All three passes found first-order defects in **shipped** code
from the in-scope iters — a fence that mis-attributes its own parse failure to its subject, a report that
cannot say which stack it is about, a probe scope covering five of eight running services. None of the
three needed a prior pass to create it. The two defects this session did create are recorded above as
its own, and neither was found by any control it wrote.

**Per the user's standing ruling, this is routed, not re-litigated, and no machinery is built for it.**
One observation is worth carrying to the close, and it is narrower than pass 32's: **the three findings
share a shape.** In each, a mechanism reported a confident verdict about a subject it had not actually
read — the include leg naming rows it never reached, the report naming a pass it could not attribute to a
stack, the scope naming five services as though they were all of them. That is not "the iters need
rework". It is the milestone's own subject, appearing in the milestone's own tooling, which is where a
harden pass is supposed to find it.

⚠️ **Pass 32's retrospective was shown wrong by iter-145** for characterising a set of failures instead of
grading them individually. This entry grades each of the three findings separately above and states the
shared shape only as an observation drawn from three graded cases, not as a property asserted over a set.

## Pass 36 — 2026-08-08 — incremental

**Iters hardened this pass:** iter-153 … iter-165 (target: the census family iters 159–165 built)

**Finding — the census said PAIR in every claim it makes and graded one literal per LINE.**

`anchor_subject_census` (iter-163) opens *"A `(citation, literal)` pair is adjudicated iff…"*, counts
`adjudicable pairs`, and waives per pair. `run()` took `min(cands)` — the literal nearest the citation —
and dropped every other literal beside it. The unit it graded was the **line**, and that narrowing is
declared nowhere in a file that declares three other blindnesses in advance.

Measured on the live corpus: **226 candidate pairs over 193 lines**. Of the 33 discarded, **20 were
adjudicable** and **one was a real unexempted finding** — `demo-up-defaults.md:77` quoting
`` `@clerk/backend` `` against `up-injected.sh:43`, where the dotless-host refusal that clause describes
lives at `:181`. The census printed `0 unexempted finding(s)` and exited 0.

The coupled half is in the **waiver layer**, and it is `§5` r70/71 one storey up from the code that rule
usually describes: `EXEMPT` was keyed `doc:line`, so the second pair at `:77` would have been absorbed by
an exemption adjudicated for `assertValidPublishableKey` — a different claim, read on a different day.
**A waiver pinned to a POSITION is not pinned to the subject somebody read.** Keys are now
`doc:line:literal`, sourced from a single `Finding.key`.

**Coverage delta:** `test_anchor_subject_census_m257x.py` 21 → 27; adjudicable population 137 → 157.
**Tests added:** 6 — the regression fixture puts a CORRECT literal beside the citation and a WRONG one
further along the same line (under `min()` the correct one wins and the wrong one is never looked at),
**verified RED against the pre-fix code by mutation**, not by argument; plus an anti-vacuity control that
`pairs > paired` on the live tree, a key-form conformance check between `EXEMPT` and `Finding.key`, the
two-independently-graded-pairs proof at `:77`, and the property that exempting one literal leaves its
siblings RED.
**Bugs surfaced + fixed inline:** 1 (`d4f208a`).

### Carried into pass 2 — a fence that went RED and was shipped over three times

`derivation_registry.unclassified()` (iter-162, *"the registry can no longer silently fall behind the
tree"*) is **RED at HEAD** and has been since **iter-163**, the very next iter, which added
`anchor_subject_census.py::Census.unexempt` and classified it nowhere.
`test_frozen_expectation_census_m257x.py::test_every_executable_derivation_is_classified` fails on this
tree. The fence did its job; iters 163/164/165 each ran a change-derived scoped suite and none of them
included the module that grades them — the standing `FIX-M257x-iter142-whole-suite-owed` gap, observed
firing. `§5` r60/66 in its operational form: **a scoped green is evidence about its scope, and the thing
outside the scope was already RED.**

Also measured, for pass 2: `PATHLIKE_ARGS` is a set of argument **names**. **13 required arguments across
9 functions are annotated `Path` / `Path | None` / `list[Path]` / `Path | str` and fall outside it**, so
**8 sites** are filtered out of the executable-here sub-population — the population `unclassified()` is
the completeness fence *for*. r70/71 again, in the newest registry.

**Stop condition:** continue-to-next-pass — the derivation registry is RED at HEAD and its
executable-here filter is pinned to a spelling; both land in pass 2.

## Pass 37 — 2026-08-08 — incremental

**Iters hardened this pass:** iter-153 … iter-165 (target: iter-162's derivation registry)

**Finding — two fences went RED exactly as designed, and three iters shipped over them.**

`stack-core` was carrying **two RED fences at HEAD**, both authored by the iters this pass hardens:

1. `derivation_registry.unclassified()` — built at iter-162 so *"the registry can no longer silently fall
   behind the tree"* — went RED at **iter-163**, the very next iter, which added
   `anchor_subject_census.py::Census.unexempt` and classified it nowhere.
2. `test_test_collection_fence` (the M256 fence for classes defined below `if __name__ == "__main__"`)
   named `test_frozen_expectation_census_m257x.py:229` (iter-162) and
   `test_anchor_construct_denominator.py:340` (iter-164).

**They are the same event seen twice.** The 12 tests direct execution could not see in the frozen-
expectation module **include the registry-completeness fence itself** — `python3 <that file>` collected
20 and printed OK. Iters 163/164/165 each ran a change-derived scoped suite that did not include the
module grading them: `FIX-M257x-iter142-whole-suite-owed`, observed firing. `§5` r60/66 operationally —
**a scoped green is evidence about its scope, and the thing outside the scope was already RED.**

**Second finding — the registry's executable-here filter was pinned to argument SPELLINGS.**
`PATHLIKE_ARGS` is a set of parameter *names* answering a question about parameter *types*. Measured:
**13 required arguments across 9 functions are annotated `Path` / `Path | None` / `list[Path]` /
`Path | str` and named none of them**, so **8 sites** sat outside the sub-population the completeness
fence exists *for* — a derivation could be added, be perfectly path-satisfiable, and never need a
decision, because its author wrote `target` instead of `path`. `§5` r70/71, in the registry built to end
that class. Third: the population glob was `*/*.py`, section roots only — widened; **measured, no verdict
moves today** (3 deeper modules, 2 derivations, neither executable-here), and fixed anyway, because *the
fence happens to be right* is not the property being claimed.

**Coverage delta:** `test_frozen_expectation_census_m257x.py` 26 → 32; direct-execution collection
20 → 32 and 26 → 30 in the two guard-misplaced files. Population 131 → 133, executable-here 60 → 68,
unclassified 1 → 0.
**Tests added:** 6 — the type-side admission proven in BOTH directions (without the negative arm the
widening could be `return True`), an anti-vacuity control that ≥ 8 live sites are admitted by TYPE ALONE
**and every one of them is graded** (RED against the pre-fix filter, verified by mutation), the sub-root
reach, and that the widened glob still excludes `tests/` at depth.
**Bugs surfaced + fixed inline:** 3 (`2e2b135`). All nine newly-visible sites adjudicated at source.
**Stop condition:** continue-to-next-pass — the shipping probe-scope path has the same unread-subject
shape and is measured but not yet repaired.

## Pass 38 — 2026-08-08 — incremental

**Iters hardened this pass:** iter-153 … iter-165 (target: iter-153/154's probe-scope helper)

**Finding — "I could not look" and "there is nothing to see" were the same bytes.**

`scope-union.sh` opens by explaining that its line 3 exists so *"the override adds nothing probeable"*
and *"there is no override"* cannot collapse — *"the exact shape harden pass 35 booked as a defect one
level up (a mechanism reporting a confident verdict about a subject it never read)"*. Four lines below
that paragraph: `. "$HERE/services.sh" … || { echo; echo; echo; exit 0; }`.

Measured: a missing `services.sh` and a project with no stack dir produce **byte-identical** output, and
`generate.sh` renders both as *"no generated override found … this is the PLATFORM's service set alone"*.
An operator reading `/test-platform` is told a fact the tooling never established. The sibling test class
is literally named `TestNoOverrideIsDistinguishableFromAnEmptyOne`, and its three covered cases do not
include *the instrument could not run*.

Line 4 is now a STATUS (`ok` / `unreadable:services-lib` / `unreadable:override-file`); it still exits 0
and never aborts (`D-M257x-148-1`). **All three consumers read it and all three fail CLOSED** — an EMPTY
line 4, which is what a missing or older `scope-union.sh` yields, counts as unreadable, not as `ok`.
`dev-stack` and `up-injected.sh` read only lines 1-2, so the contract change is safe for them, and both
carried the same blind spot; both now disclose it.

**Coverage delta:** `test_scope_union_m257x.py` 16 → 21 (+1 existing test updated to the four-line
contract). **Tests added:** 5, including **the measurement asserted directly** — the two outputs must no
longer be equal — and `generate.sh`'s real block executed against a chmod-000 override.
**Bugs surfaced + fixed inline:** 1 (`757730b`).

**Routed forward — `FIX-M257x-h36-labeled-prover-denominator` (Fate 3).** `labeled_spelling_pins.py`
**contradicts itself inside one report**: `BLIND SPOT 1 instance(s) confirmed structurally invisible`
three lines above `3 of 7 instances` for the same property. `expect_blind` is set on **1 of the 3**
instances its own taxonomy calls structurally invisible (*"No haystack clause can see these — nothing
reads a source file"*), so the published `RECALL 4/6 = 67%` is an artifact of applying the module's own
exclusion rule to one third of the sites that qualify for it; applied consistently it is 4/4, applied to
none, 4/7. The pre-registered 50 % refutation floor is computed over that denominator. Routed rather than
fixed inline because **choosing between the two consistent readings changes a figure the milestone
quotes** — a design decision, not a corollary of a test. Two adjacent latent defects to fold in: a stale
`expect_blind` declaration prints `⚠ … the declaration is stale` and still **exits 0**, and an instance
whose commit is unreadable is `NOT COUNTED`, **silently shrinking the recall denominator**.

### Suite results (counts, never wall-time — `§5` rule 51's timing leg fails on this host)

| suite | result | baseline |
|---|---|---|
| `stack-core` | **1 failed · 1478 passed** | pass 35: 2 failed · 1338 passed |
| `stack-verify` | **0 failed · 275 passed** | pass 35: 252 passed |
| `dev-stack` | **0 failed · 151 passed** | — |
| `demo-stack` | **9 failed · 1055 passed · 2 skipped** | — |

`stack-core`'s single failure is the standing `test_claim_twin_guard_iter48_answer_key::test_02`,
unchanged; pass 35's second failure was its own defect and is gone. **The 9 `demo-stack` failures are
graded individually, never as a set** (pass 32 characterised 21 failures and iter-145 proved 57 % of that
characterisation false): 3 `test_migrate_race_live` reproduce at the pre-harden commit `5385390` in a
fresh clone with no live container; the other 6 each assert a **sha256 of a file in a live clone on this
box** — `ant-academy/code/next.config.js`, next-web's `urls.ts` — and **neither those files nor the
manifests they are compared against appear in this pass's 10-file diff**, so both operands of every one
of those six comparisons are byte-unchanged by this pass and the verdict is identical at both commits.
That is a per-test proof, not a property asserted over a group.

**Flake gate:** 3 consecutive clean runs of all five touched test files — 126 passed ×3.
**Invocation-parity check** (the pass-33/35 self-inflicted class): all four changed test files collect
identically under pytest and by direct execution — 27, 32, 30, 21 — from `stack-core`, from
`stack-verify`, and from the rext root.

**Knowledge backfill:** none needed — every finding this session is a defect in an instrument, and each
instrument's own docstring now carries its retraction and the measurement behind it.

**Stop condition:** cap reached without stabilization — three passes, five defects fixed inline, one
routed. Each pass found a real defect in the previous iters' newest instruments, so coverage has not
stabilized; the residual is the routed `FIX-M257x-h36-labeled-prover-denominator` plus the standing
process gap the two RED fences exposed (a change-derived scoped suite cannot see the fence that grades
it). Per the user's standing ruling, this is routed and NOT met with new machinery.

---

## Pass 39 — 2026-08-09 — incremental

**Iters hardened this pass:** iter-166 … iter-176 (11 tiks; target: the registry/enumeration machinery
iters 174–176 built, and the ledger figure iter-173 was barred from repairing)
**Tiks covered since prior pass:** 11
**Runner named on every count below** (§5, iter-170): `/usr/bin/python3` **3.9.6**, the only interpreter
on this box with pytest. Direct-execution checks additionally on `python3` **3.14.6**.

### The owed item, discharged — `FIX-M257x-iter173-ledger-denominator` (`c7234fd`)

Two sites in this file published a figure the file's own table refutes. `1,230 of **2,989**` was
assembled as `2,978 passed + 11 skipped`, **dropping the same table's 22 failures**, so the denominator
changed unit from *executed* to *passed-and-skipped* inside the entry whose subject is denominators; pass
32 carried the hole forward as `2,989 + 51 = 3,040` and additionally dropped its own section's single
failure from the numerator. Re-derived here **independently of iter-173's pre-computation**, because
`1,280/3,040` and `1,281/3,062` are both 42 % and a careless repair would inherit the same protection
the defect had: `2,978 + 22 + 11 = 3,011`, `stack-core` `1,229 + 1 + 0 = 1,230`; after the `+51`,
**1,281 of 3,062**. Both sites now name the unit and carry the retraction. The `59 %` on the following
line is unaffected (59.2 % vs 58.9 %) and is left as written.

### Finding 1 — the enumeration was blind to a registry its own docstring names twice (`6b83e61`)

iter-176 shipped `test_fence_registry_population_m257x.py` to close
`FIX-M257x-iter174-accept-registers-one-registry-of-two` *"at its population rather than at its last
member"*. It reported **5** sites and classified 5. **Its own history table names six.**
`derivation_registry` appears there twice — found at iter-173 by grep, again at iter-175 by hand — and
appeared in the fence's output never.

Cause: `_norm()` normalised a string constant by stripping a trailing `.py`.
`derivation_registry.DECISIONS` is keyed `"stack-core/derived_count_guard.py::postcondition_sites"` — a
**qualified id** — so it matched no fence name and scored 0 against a floor of 2. **Structurally
invisible: no number of runs could have surfaced it**, the same shape as iter-174's fail-open probe, and
the **fourth** enrolment-by-spelling defect in M257x after iters 157, 162 and 175.

That it is a registry is not an opinion: `3fb1d98` — iter-173's fence commit — adds that key **in the
same commit that ships the fence**, and `test_frozen_expectation_census_m257x.py:239` asserts
`unclassified() == []`, so the obligation is enforced. Widened `_norm` measured **5 → 7 sites, 0 lost**:
`derivation_registry.py` (REGISTRY:reconciled) and `test_m255_mutation_battery.py` (DECLINE:subject — it
stages M255's own three subjects and asserts no superset over the fence population, unlike the M257x
battery whose `test_000` does). Two additions, both real — a widening, not a volume change.
Control: `test_mutation_control_the_QUALIFIED_KEY_reader_is_load_bearing` restores the shipped
bare-name `_norm` and **requires `derivation_registry.py` to be LOST** — written as a mutation rather
than an `assertIn`, which would still pass in a future that re-keys `DECISIONS` by bare name and leaves
the `::` reader dead.

**The registry population is now 6 + 2 declines.**
`FIX-M257x-iter174-accept-registers-one-registry-of-two` stays open (unchanged by this pass).

### Finding 2 — the disclosed limit said `16 of 27`; it was `15 of 26` the day it was written (`d24132f`)

> ## ⚠️ THIS FINDING IS RETRACTED — measured and refuted by iter-177, re-derived independently by harden pass 42
>
> **`16 of 27` was never wrong.** It is an **exact** reading of `guard_family.union()`. `15 of 26` is an
> equally exact reading of a **third** population — `repair_postcondition.declared_kind` — that this
> finding did not know existed. The sentence below, *"both operands were wrong at publication"*, is
> **false in both operands**, and the heading above is false as written.
>
> One population, **three live derivations**, re-derived from the owners by pass 42 in a scratch
> `git archive` copy at **both** refs this finding used (never in the tree):
>
> | derivation | owner | rule | at `5b108d0` (iter-175's own commit) | at `c7f4c3d` | at rext HEAD `08ad440` |
> |---|---|---|---|---|---|
> | `union` | `guard_family.union` | spelled ∪ declared ∪ `EXTRA_CENSUS_MEMBERS` | **16 of 27** | **16 of 27** | 17 of 28 |
> | `census` | `guard_family.census` | `union` − `CENSUS_EXCLUSIONS` | **16 of 26** | **16 of 26** | 17 of 27 |
> | `declaring` | `repair_postcondition.declared_kind` | declares `FENCE_KIND` | **15 of 26** | **15 of 26** | 16 of 27 |
>
> **`census` and `declaring` both return 26 over different members** — they differ by exactly one member
> in *each* direction, each for a correct, documented reason: `guard_family` is IN `declaring` (it
> declares a `FENCE_KIND`) and OUT of `census` (running the family runner inside itself recurses);
> `repair_postcondition` is the mirror. `union == census | declaring` holds at all three refs. **So every
> count-based comparison of the two 26-member sets reads GREEN**, which is how one population came to
> publish `15 of 26`, `16 of 26` and `16 of 27` at the same time with nothing going RED.
>
> The defect was never arithmetic — it was a **missing population label, in the claim *and* in this
> retraction of it**. This finding corrected a number onto a different population without noticing there
> were three, so it did not fix the error; it reproduced it and added authority. **A retraction inherits
> every weakness of the claim it retracts.**
>
> **Consequently `FIX-M257x-h39-survey-id-embeds-retracted-figure` is CLOSED AS REFUTED** — the routed id
> `SURVEY-M257x-iter175-readme-fence-index-is-16-of-27` embeds a **correct** figure and needs no rename.
> What iter-175 actually got wrong was calling a `union` reading *"the census"*. The live instrument
> (`tests/test_fence_registry_population_m257x.py`) now publishes one **labelled** figure per derivation
> and asserts each against its own owner; the two 26-member sets are compared by **membership**, not
> count. Note the population **grows** — 27/26/26 at iter-177, 28/27/27 at HEAD — so quote these figures
> with their ref, never as standing facts.

Same file, one level down. `test_the_disclosed_limit_is_STATED_not_assumed` exists, in its own words, to
keep the limit *"honest by measuring it rather than asserting its absence"* — and asserted
`len(named) >= 2`. Measured now: **15 of 26**. Re-derived at **`5b108d0`, iter-175's own commit**,
reconstructed with `git archive`: the tree said 15 of 26 **there too**.

So it is **not drift — both operands were wrong at publication**, by an instrument iter-175 ran once and
did not check in, and iter-176 quoted it twice. A floor of two cannot detect an error of one in a
numerator or one in a denominator. The assert is now the claim: both docstrings' `**N of M**`
disclosures are parsed as a construct (§8) and must equal the live measurement — **no hand-maintained
constant**, the published prose is the subject (§2, derive at the point of use). Mutation-proved:
restoring `16 of 27` FAILS naming both figures, deleting the disclosure FAILS naming the missing
construct, baseline and restore PASS.

**Routed forward — `FIX-M257x-h39-survey-id-embeds-retracted-figure` (Fate 3).** The routed id
`SURVEY-M257x-iter175-readme-fence-index-is-16-of-27` carries the retracted figure **inside the
identifier**, and iter-175's and iter-176's `progress.md` both cite it. A harden pass does not rewrite an
iter's own record, so the rename is routed rather than taken.

### Suite results (counts, never wall-time — `§5` rule 51's timing leg fails on this host)

| suite | runner | result |
|---|---|---|
| 20 in-scope + consuming test files | pytest 3.9.6 | **693 passed · 1 skipped · 0 failed** |
| guard family (`--repo-root` = rosetta) | 3.9.6 | **18 GREEN · 0 RED · 0 could-not-check · 8 not-run** (each not-run names the input it lacks; exit 2, which is the runner refusing to call an unsupplied input a pass) |

`test_claim_twin_guard_iter48_answer_key::test_02` — the failure standing since pass 29 — is **GREEN**,
as iter-167 reported; this pass confirms it independently.

**NOT COVERED by this pass, stated rather than implied (§5 rule 60):** the six mutation batteries
(`test_m255_*`, `test_m257x_claim_twin_*`, `test_m257x_mechanical_fences_*`,
`test_m257x_repair_postcondition_*`, `test_m257x_repair_reach_*`, `test_repair_leak_guard_*`) and the
`dev-stack` / `stack-injection` / `stack-verify` sections. They run at pass close, not per pass.

**Knowledge backfill:** none as a corpus edit — both findings are defects in instruments, and each
instrument's own docstring now carries its retraction, its measurement and the ref the measurement was
taken at.

**Stop condition:** continue-to-next-pass — two real defects in the newest fence, both fixed inline; the
`FIVE registries` structural lead is now enumerated at six but its *sibling* obligation
(`FIX-M257x-iter174-accept-registers-one-registry-of-two`) is untouched, and the batteries are unrun.

---

## Pass 40 — 2026-08-09 — incremental

**Iters hardened this pass:** iter-166 … iter-176 (same batch; target: the two instruments iters 172 and
174 shipped, on the axis pass 39 opened — *does the control test the claim, or something near it*)
**Tiks covered since prior pass:** 11 (shared scope with pass 39)
**Runners:** `/usr/bin/python3` **3.9.6** (pytest) and `python3` **3.14.6** (direct). Both named per count.

### Finding 3 — the repaired stdlib set still missed every C extension (`dc2e677`)

iter-174 repaired a capability probe that failed OPEN and stated the asymmetry that makes it serious:
*an over-broad stdlib set refuses a legitimate file LOUDLY; an empty one permits a shadow SILENTLY.*
**The repair sat on the wrong side of that asymmetry for one class of module.** It excluded
`lib-dynload` as a **directory name** and never listed inside it, so every C-extension module was absent
— measured **232 names on 3.9.6**, carrying `json`/`os`/`re` and **not** `math`, `array`, `binascii`,
`select`, `zlib`, `cmath`, `mmap`, `unicodedata`, `termios`, `resource`, `readline`, `pyexpat`, `fcntl`,
`grp`, `syslog`. A repo module named `math.py` was **still silently staged over the stdlib**, on the only
interpreter this milestone's suite counts are taken on.

Why iter-174's four net-new controls could not see it — **and the fourth one says so in its own
docstring**: the floor is `> 100` and 232 clears it; the membership control names six modules, all
pure-Python; and `test_it_agrees_with_the_native_attribute_where_the_interpreter_HAS_one` **skips on 3.9
and is a TAUTOLOGY on 3.14**, where the fast path returns the very attribute it is compared with. Its
docstring reads *"where it is [used], nothing else can cross-check it."* That sentence was the finding.

Fix: walk `stdlib` **and** `platstdlib`, and list the platform extension directory (`lib-dynload` POSIX /
`DLLs` Windows) by module stem — **306 names on 3.9.6, 303 on 3.14.6**. And the cross-check iter-174 said
could not exist, does: on an interpreter that HAS the attribute, **patch it away, force the derivation,
compare**. `native − derived` must be empty against an allowance **enumerated by class and named member
by member** — 12 names: 7 windows-only, 2 optional C extensions this build lacks, 2 frozen bootstrap
modules, 1 excluded on purpose — with a staleness check in the other direction, and an assert that the
patch TOOK so the control cannot decay back into the tautology it replaces. Over-claim measured too:
**18 extras on 3.14.6**, all CPython test/build artifacts, **0 collisions against the repo's 171 module
names**.

### Finding 4 — the census's completeness check compared a set with a subset of itself (`f0ac50e`)

`test_the_module_population_is_enumerated_not_globbed_from_one_section` asserts
`{m.split("/")[0] for m in modules(REXT)} == set(SECTIONS)`. **`modules()` iterates `SECTIONS`**, so the
left side is a subset of the right *by construction*: the equality can only catch a **declared** section
that contributes nothing. It cannot catch a **Python-bearing section that was never declared** — the
direction the word *population* is about, and the direction the test's own name promises.
`suite_census.SECTIONS` is a hand-written tuple of five, and the census publishes its result as the whole
population (iter-172: *112 modules*, *3350 tests*).

**No published number is wrong.** Measured: the repo has **12** top-level directories, **5** hold Python
at all, and they are exactly the five declared — the tuple is complete today. It is not *derived*, which
is what `§2` prescribes, and the risk is concrete for this repo: `CLAUDE.md`'s own section enumeration
omitted `stack-secrets` and `playthroughs` until iter-129, one of which a documented skill already
depended on. Added the missing direction, derived by walking the tree independently of `SECTIONS`, with
an anti-vacuity floor. **Mutation-proved:** dropping `dev-stack` from `SECTIONS` leaves the OLD test
**GREEN** and turns the NEW one **RED** naming `['dev-stack']`.

### Suite results (counts, never wall-time)

| suite | runner | result |
|---|---|---|
| 9 touched + consuming test files | pytest 3.9.6 | **248 passed · 2 skipped · 0 failed** |
| `test_battery_stage` | unittest 3.14.6 | **Ran 19, OK** — all 19 execute there, incl. the new cross-check |
| `test_battery_stage` | unittest 3.9.6 | **Ran 19, OK** (skipped=2 — the two that need a native attribute) |
| `test_suite_census` | pytest 3.9.6 / unittest 3.14.6 | **12 passed / Ran 12, OK** |

**NOT COVERED by this pass (§5 rule 60):** the six mutation batteries and the four non-`stack-core`
sections [**left as written, and it is the exhibit of pass 51's finding**: this entry carries the right
NUMBER with an ambiguous noun — there are **ten** non-`stack-core` sections and **four** of them carry
Python — while passes 45–50 carried the right noun with the wrong number. Neither pass ever held both
halves, and both halves were in this one file the whole time]. `TestTheRealBatteriesStillDerive` is green, so no battery's staged set moved, but the batteries
themselves run at pass close.

**Knowledge backfill:** none as a corpus edit; both repairs carry their measurement, their allowance and
their runner in the instrument's own docstring.

**Stop condition:** continue-to-next-pass — two further defects, both of the same shape as pass 39's
(a control that measures something adjacent to its claim). Four findings in two passes is not a
stabilizing series; one more pass is owed before the cap.

---

## Pass 41 — 2026-08-09 — incremental

**Iters hardened this pass:** iter-166 … iter-176 (same batch; target: iter-171's monorepo-wide server
population scan — the one instrument in the batch whose *scope* claim is larger than its section)
**Tiks covered since prior pass:** 11 (shared scope with passes 39–40)
**Runners:** `/usr/bin/python3` **3.9.6** (pytest) · `python3` **3.14.6** (direct). Named per count.

### Finding 5 — the monorepo-wide scan skipped every basename collision, silently (`c7f4c3d`)

iter-171's `TestSectionPopulationHasNoUnfixedServer` is explicit that its scope is the whole monorepo:
*"would let the next HTTP server — in `stack-verify`, say — bind through CPython's resolver unnoticed."*
Its loader cached by **basename**: `name = "_iter171_" + splitext(basename(path))[0]`, with a
`sys.modules` hit in front of it.

Two candidate files sharing a basename therefore **collapse onto one module**. Demonstrated with a clean
`a/srv.py` and a stock-class `b/srv.py`: the second import returns the first's module, `OffenderServer`
is never read, and **`checked` increments anyway** — counting the first module's classes twice. Result
`checked = 2, offenders = []`: **a missed offender AND a reach metric that hides the miss**, with no
assert able to fire, because the anti-vacuity floor `checked >= 3` is satisfied by the duplicate.

Latent today only because all four candidates live in `demo-stack/` with distinct names — and the names
that repeat across sections are exactly `cockpit.py`, `server.py`, `test_cockpit.py`, the ones this scan
exists to reach. Keyed on the repo-relative path now. **Mutation-proved:** restoring the shipped
basename key FAILS with `['_iter171_srv', '_iter171_srv']`.

**And the class is closed at its population, not at its last member** (iter-169's rule; §8 iter-168's
denominator). Censused: **16 `spec_from_file_location` sites in the repo; exactly 2 insert into
`sys.modules`** — this one, and `test_predicate_enumerator.py:255`, which uses a constant name over a
single per-test temp path and cannot collide. The other 14 name a constant and load one fixed path with
no cache, so re-importing merely re-executes. **1 of 16 had the hazard. Not systemic.**

### The whole population, one runner, unit named

**pytest 3.9.6**, five sections, taken on an otherwise-idle box **before** the pass's last edit. *tests*
below = **collected = passed + failed + skipped**.

| section | passed | failed | skipped | collected |
|---|---|---|---|---|
| `stack-core` | 1545 | 0 | 2 | 1547 |
| `demo-stack` | 1063 | 9 | 2 | 1074 |
| `stack-injection` | 335 | 0 | 0 | 335 |
| `stack-verify` | 275 | 0 | 0 | 275 |
| `dev-stack` | 151 | 0 | 0 | 151 |
| **total** | **3369** | **9** | **4** | **3382** |

`stack-core` — the section this ledger's older entries meant by *"the whole suite"* — is
**1,547 of 3,382 collected, 46 %**.
*(That triple is on ONE line deliberately: `derived_count_guard`'s arm C matches per line, so the first
draft — which wrapped between `of` and `3,382` — published a percent-triple the fence could not see.
Caught by reading the guard's arm counts before and after, `C 4 → 5`. The same defect this pass spent
three findings on, committed while writing them up.)*
**`stack-core` is 0 failed**: the long-standing
`test_claim_twin_guard_iter48_answer_key::test_02` is GREEN (iter-167), and no pass 39–41 change
introduced a RED. `stack-verify` was **12 failed** at pass 30 and is **0** here.

**The 9 `demo-stack` failures, graded ONE BY ONE** — never as a set (pass 32 characterised 21 failures
and iter-145 proved 57 % of that characterisation false):

* **3 × `test_migrate_race_live`** — and *not* for the reason a set-level reading would have given.
  A demo stack **is** up on this box (`demo-1-backend-1`, `demo-1-directus-1`, …), so "no container" is
  refuted. The actual messages: `pg_isready` probes `/var/run/postgresql:5432 — no response` (a local
  socket, not the demo's port); `migrate-demo.sh` aborts with *"missing …/stack-core/lib/repos_yml.sh —
  this rext checkout is incomplete"* from its **own** temp staging (a shell path, **not** the
  `battery_stage.local_deps` helper this pass edited); the third fails downstream of the second
  (`first run should seed; got ''`).
* **6 × sha-pin against a live clone** (`test_ant_academy` 1, `test_demopatch` 2, `test_ssr_origin_chain`
  3) — each compares a file in `/Users/marco/workspace/anthropos/rosetta/stack-demo/…`, **outside this
  repo entirely**, against a manifest under `demo-stack/patches/`. `git log 95e174a..HEAD --
  demo-stack/patches/` is **empty**, and the clone is not in the repo at all, so **both operands of every
  one of the six are byte-unchanged by this session**.
* **The consumption check, per module:** none of the four failing modules imports `battery_stage`,
  `suite_census`, the fence-registry test or the bind test. They import `manifest_loader`,
  `gen_injected_override` and `uuid` — none touched. That is a per-test proof, not a property asserted
  over a group.

**Flake gate:** 3 consecutive clean runs of the four touched test files — **45 passed · 2 skipped**, ×3.
**Fence check:** `derived_count_guard` GREEN over the edited ledger — **28 sites, 0 findings**.
**Invocation parity** (the pass-33/35 self-inflicted class): the changed files were run under pytest
3.9.6, unittest 3.14.6 **and** unittest 3.9.6, from `stack-core`, the rext root and `stack-core/tests`.

**Knowledge backfill:** none as a corpus edit. All five findings are defects in instruments; each
instrument now carries its own retraction, its measurement, and the ref the measurement was taken at.

**Routed forward (Fate 3), both open:**
* ~~`FIX-M257x-h39-survey-id-embeds-retracted-figure`~~ — **CLOSED AS REFUTED** (iter-177; re-derived
  independently by harden pass 42). This route existed only because pass 39 believed `16 of 27` was
  retracted. It is not: it is an exact reading of `guard_family.union()` at both refs, and the id
  `SURVEY-M257x-iter175-readme-fence-index-is-16-of-27` therefore embeds a **correct** figure. **Only
  one of the two routes below was ever open.** See the retraction banner on pass 39 Finding 2 for the
  three-derivation measurement.
* `FIX-M257x-iter174-accept-registers-one-registry-of-two` — untouched. Pass 39 *enumerated* the
  registry population (3 → **6 registries + 2 declines**); it did not make `--accept` write the second
  one.

**Stop condition:** cap reached without stabilization — three passes, **five defects fixed inline, two
routed**. Every pass found a real defect in the previous iters' newest instruments, and all five are the
same shape: *the control measured something adjacent to its claim* — a registry reader blind to a key
spelling, a disclosed limit asserted as a floor, a stdlib set that excluded a directory instead of
reading it, a completeness check derived from the thing it checks, and a population scan keyed on
filenames. Coverage has not stabilized. Per the user's standing ruling this is **routed and NOT met with
new machinery**; the eighth cap-without-stabilization in this milestone (22, 25, 26, 29, 32, 35, 38, 41).

---

## Pass 42 — 2026-08-09 — incremental

**Iters hardened this pass:** iter-177 … iter-186 (10 tiks; target: the instrument's own
population/backlog/grammar, which is what those ten iters were about)
**Tiks covered since prior pass:** 10
**Runner named on every count below** (`§5` rules 75/76): `/usr/bin/python3` **3.9.6**, the only
interpreter on this box with pytest. Every new module additionally run under `python3` **3.14.6**
(unittest) and `/usr/bin/python3` unittest, from **two** directories.
**Section scope named on every count** (iter-186's finding — a "whole-population" figure that does not
name its sections describes 5 of 11): this pass is **`stack-core` only**. The other ten sections were
not run and nothing is claimed about them.

### The owed item, discharged — `FIX-M257x-iter177-ledger-carries-a-retracted-retraction`

Pass 39 Finding 2 now carries a retraction banner, and pass 41's routed-forward list is corrected.
**Re-derived here independently before writing it**, because the whole class exists because equal-sized
sets compare green — reading iter-177's table and copying it would have reproduced exactly the error
pass 39 made. Read through the owners (`guard_family.union` / `.census`, `repair_postcondition.
declared_kind`) in a scratch `git archive` copy at both refs; the tree was never edited to take a
measurement:

| derivation | at `5b108d0` (iter-175's own commit) | at `c7f4c3d` | at rext HEAD |
|---|---|---|---|
| `union` | **16 of 27** | **16 of 27** | 17 of 28 |
| `census` | 16 of 26 | 16 of 26 | 17 of 27 |
| `declaring` | **15 of 26** | **15 of 26** | 16 of 27 |

`16 of 27` is an **exact** reading of `union`; `15 of 26` is an **exact** reading of `declaring`. Pass
39's *"both operands were wrong at publication"* is **false in both operands**. `census` and `declaring`
differ by exactly one member in each direction (`repair_postcondition` / `guard_family`), both for
documented reasons, and `union == census | declaring` holds at all three refs.
`FIX-M257x-h39-survey-id-embeds-retracted-figure` is **closed as refuted**. Note the population
**grows** — quote these with a ref, never as standing facts.

### Finding 1 — harden passes route into a file the backlog fence cannot read (`5228b2d`)

`route_disposition_guard` exists to audit *"the queue every sub-agent on this milestone is briefed
from"*. Its population is one line — `milestone_dir.glob("iter-*")` (`:207`) — and **`hardening-ledger.md`
is not in it**. But harden passes route too, under their own origin token `h{K}`, and the ledger *is*
briefed from: this pass's own prompt lifted its owed item out of it.

Measured — **5** harden-origin routes; **4 are reachable only because some later iter happened to cite
them**, which is luck, not a rule:

| route | reachable? | via |
|---|---|---|
| `FIX-M257x-h30-crossline-repair` | yes | iter-144 |
| `FIX-M257x-h30-nonstackcore-suite` | yes | iter-144, iter-145 |
| `FIX-M257x-h33-derive-includes-stack-override` | yes | iter-153 |
| `FIX-M257x-h39-survey-id-embeds-retracted-figure` | yes | iter-177 |
| **`FIX-M257x-h36-labeled-prover-denominator`** | **NO** | **nothing, ever** |

That last one has been **open and structurally invisible since pass 36** — through passes 37, 38, 39, 40
and 41. It is not trivia: `labeled_spelling_pins.py` contradicts itself inside one report, and the
published `RECALL 4/6 = 67 %` is an artifact of applying the module's own exclusion rule to one third of
the sites that qualify for it (consistently: 4/4, or 4/7).

`test_harden_origin_route_visibility_m257x.py` makes the exclusion **loud** rather than widening the
guard's population — the ledger uses a different disposition grammar (`**Routed forward — … (Fate 3).**`,
not `**Routes carried forward:**`), and choosing one is a design decision, not a corollary of a test
(Fate-3 boundary). The registry is checked **both ways**: an undispositioned invisible route is RED, and
so is a stale entry — a "known exceptions" list that only grows in one direction is how iters 157, 162,
174 and 175 each failed. Readers are **imported from the guard that owns them**; a private `ID_RE` would
be a second derivation of the very thing being compared (iter-177's rule).

**This module shipped a fail-open green in its own first run and its own baseline caught it:**
`REPO_ROOT` was `parents[3]` (correct from `stack-core/`, off by one from `stack-core/tests/`) and the
missing-subject path raised `SkipTest` — together they read **`1 passed, 4 skipped`**, green in CI-speak,
with all four load-bearing asserts unrun. The probe now **FAILS** on a missing subject. *A capability
probe that fails OPEN disarms the check it guards*, this time inside the pass that wrote it.

### Finding 2 — a size-preserving mutation can be invisible to the runner (`5228b2d`)

Found while RED-proving Finding 1: one mutation read **GREEN batched and RED alone**. Not flake — in the
green run **the interpreter never loaded the edit**.

CPython invalidates a cached `.pyc` on **(source mtime in whole seconds, source size)**. A
**size-preserving** edit landing **within the same clock second** changes neither operand, so stale
bytecode is reused. Demonstrated end-to-end: `VALUE = "CCC"` on disk, `import mod; mod.VALUE` → `AAA`.

**Two hypotheses were refuted by their own controls before the real one** — recorded because the
refutations are the finding's evidence, not detours:

| attempted mitigation | result |
|---|---|
| `rm -rf __pycache__` | **no effect** — the directory is genuinely ABSENT while stale bytecode is served |
| `PYTHONDONTWRITEBYTECODE=1` | **no effect** — suppresses *writes*; reads still consult the cache |
| bump the source mtime | **works** |

The host-specific half: macOS system Python caches **out of tree**, at
`~/Library/Caches/com.apple.python/<abs source path>/`. Scope was corrected **by the control's own
3.14.6 run**, not assumed — the first draft claimed 3.14.6 "is not affected" and it is: staleness is
CPython's rule, not Apple's, and only the cache LOCATION differs (in-tree there, so clearing helps
there and not here). **The mtime bump is the mitigation that holds on both.**

**Why this milestone is exposed:** rule 75 established `/usr/bin/python3` 3.9.6 is the only interpreter
here with pytest — the runner named in nearly every mutation proof in this file — and the exposed class
is exactly the shape these fences are mutated with: `16 of 27` → `15 of 26`, `>=` → `<=`, a single-digit
count, a one-character path index. **A mutation of that shape reported as "did not fire" may never have
been loaded.** `test_mutation_proof_cache_hazard_m257x.py` reproduces the masking constructively and
asserts the mitigation defeats it.

**This pass re-ran its own mutation battery under the mitigation rather than trusting the first run:**
4/4 RED (including the previously-masked size-preserving one), baseline and restore green.

**Coverage delta on touched files:** 2 net-new control modules, 8 tests, over machinery that had 0
coverage for these two properties (harden-origin route visibility; mutation-proof validity).
**Tests added:** `stack-core/tests/test_harden_origin_route_visibility_m257x.py` 5 ·
`stack-core/tests/test_mutation_proof_cache_hazard_m257x.py` 3.
**Bugs surfaced + fixed inline:** 2 — the fail-open `parents[3]`/`SkipTest` pair, and the first draft's
wrong interpreter-scope claim (both caught by this pass's own controls, both fixed before commit).
**Flakes stabilized:** none — but one apparent flake was **diagnosed as not a flake**: the batched-vs-
alone mutation disagreement was a stale-bytecode artifact with a deterministic cause.

**Suite results (counts, never wall-time — `§5` rule 51's timing leg fails on this host; section scope
named per iter-186):**

| suite | runner | section scope | result |
|---|---|---|---|
| 2 new modules + 5 consuming/adjacent modules | pytest 3.9.6 | `stack-core` only | **106 passed · 0 failed** |
| both new modules | unittest 3.14.6 | `stack-core` only | **OK** (8 tests) |
| both new modules | unittest 3.9.6 | `stack-core` only | **OK** (8 tests) |
| `guard_family` runner (`--repo-root` = rosetta) | 3.9.6 | fence family | **19 GREEN · 0 RED · 0 could-not-check · 8 not-run** (each names the input it lacks; exit 2 = the runner refusing to call an unsupplied input a pass) |
| `derived_count_guard` over the edited ledger + corpus | 3.9.6 | milestone + corpus | **OK — 42 sites, 0 findings** |
| `route_disposition_guard` | 3.9.6 | 3 milestones | **OK — 0 contradictions** |

**NOT COVERED by this pass, stated rather than implied (`§5` rule 60, sharpened by iter-186):** the
whole-section `stack-core` run (~1,550 collected) was **not** taken this pass — it is reserved for pass
close with the tree frozen, since editing during a run has confounded nine runs in this milestone. The
other **ten** rext sections (`stack-verify`, `dev-stack`, `stack-injection`, `demo-stack`,
`stack-seeding`, `stack-secrets`, `stack-snapshot`, `alignment`, `playthroughs`, `clerkenstein`) and the
**Go** suites were not run. Nothing here is a whole-population claim.

**Knowledge backfill:** `corpus/ops/platform-alignment.md` **§5 rule 77** — *a size-preserving mutation
can be invisible to the runner*, with the two failed mitigations, the per-interpreter cache-location
table, and the instruction to re-run any size-preserving mutation proof under a forced mtime bump before
quoting it. Placed as rule 76's twin one layer down: rule 76 says an unexplained runner disagreement is
evidence about the **code**; rule 77 says an unexplained mutation result is evidence about the
**toolchain**.

**Routed forward (Fate 3):**
* `FIX-M257x-h36-labeled-prover-denominator` — **unchanged and still open**, but no longer silent: it
  now carries a written disposition inside a fence that goes RED if it is dropped. Choosing between the
  two consistent readings (4/4 vs 4/7) moves a figure the milestone quotes, so it stays a design
  decision.
* `SURVEY-M257x-h42-size-preserving-mutation-proofs-unaudited` — **NEW.** Rule 77 says a size-preserving
  mutation proof may have been vacuous; this pass proved the mechanism and fixed its **own** battery, but
  did **not** audit the milestone's prior mutation proofs for the pattern. That is a census across ~40
  ledger passes and iter records — out of inline scope, and explicitly **not** a claim that any prior
  proof was vacuous (grade findings individually — pass 32 characterised a set and iter-145 proved 57 %
  of it false).

**Stop condition:** continue-to-next-pass — two real defects surfaced in the newest instruments, both
fixed inline, plus a toolchain hazard that undermines the *method* the previous passes used. Coverage
has not stabilized: no delta has been measured across two passes yet.

---

## Pass 43 — 2026-08-09 — incremental

**Iters hardened this pass:** iter-177 … iter-186 (same batch; target: the population machinery iters
185 and 186 shipped, and — the productive move — their **consumers** rather than only the modules they
changed)
**Tiks covered since prior pass:** 10 (same batch as pass 42)
**Runner named on every count** (`§5` rules 75/76): `/usr/bin/python3` **3.9.6** (pytest); every changed
module additionally under `python3` **3.14.6** and `/usr/bin/python3` unittest.
**Section scope:** **`stack-core` only.** Ten sections and both Go/TS suites unrun; nothing here is a
whole-population claim (iter-186's rule).

### Finding 1 — a PARTITION cannot detect a MISASSIGNMENT (`961a774`)

iter-186 fixed *"Go sections silently ABSENT from the stated denominator"*. The inverse was still open:
a Go-only section silently **PRESENT in the Python denominator**. `LANGUAGE_EXCLUDED_SECTIONS` is keyed
by **name**, so `derive_sections` collects anything not spelled in it.

Measured, not argued: dropping a new Go-only section into the repo (a `go.mod` + one `*_test.go`) moves
`SECTIONS` **5 → 6** and leaves **all six** population arms GREEN. The census then runs a Python runner
over it, finds nothing, and folds a **silent zero** into a total this milestone quotes as a
whole-population baseline.

Structurally invisible to every arm iter-186 wrote, because **a misassignment is still a partition** —
`len(SECTIONS) + len(EXCLUDED) == len(disk)` holds either way. `§5` r70/71: pinned to a **spelling**, not
to the **property** that justifies it. New arm checks the assignment **by property** in the missing
direction: excluded sections non-empty in their own language (existing) · collected sections non-empty
in **Python** (new). iter-150's declared-partition/derived-completeness split preserved.

### Finding 2 — the neighbouring registry had rotted, in a file that states the rule (`961a774`)

`CITATION_EXTS` has a both-ways fence since iter-185, and the file says it outright: *"a registry rots in
both directions or it is not fenced."* Its immediate neighbour `CITATION_NON_FILE_TAILS` was fenced for
**overlap** and for **size**, never for **occurrence** — while its own comment claimed *"each is a live
corpus token."*

Measured over the same 93-file surface the sibling arms use: **4 of 9 never occur** — `org` 0, `work` 0,
`internal` 0, `io` 0, against `anthropos` 16, `net` 10, `com` 4, `local` 2, `de` 1. (`internal` reads
live only inside `backend.internal.anthropos:8083`, where the **tail** is `anthropos`.)

The four are **dropped, not excused**: a dead carve-out is not inert — its only reachable effect is to
**hide** a citation whose tail it names, which is the `go.mod` miss iter-185 paid **51 citations** for,
pointed the other way. Dropping restores the designed workflow (a future `anthropos.work:8080` turns the
both-ways `CITATION_EXTS` arm RED and a human buckets it). Behaviour re-checked: `app/go.mod:14-18` and
`page.jsx:28` still match; all five live authorities still excluded.

### Finding 3 — two fences were RED at HEAD, each naming the iter that broke it (`07035ca`)

Surfaced by running the **consumers** of the in-scope iters. Both are **byte-identical** at `08ad440`
and in the working tree — pre-existing at HEAD, not introduced here. Graded **per-test**, never as a set
(pass 32 characterised 21 failures as *"provably not ours"* and iter-145 proved 57 % of that false).

* **iter-185** defined `class TheCitationExtensionClassIsARegistry` (**5 tests**) **after** the
  `__main__` guard (`test_predicate_enumerator.py:582`), so `python3 test_predicate_enumerator.py` did
  not collect them **and still printed OK** — **iter-182's finding exactly, committed three iters after
  the fence for it shipped.** The 5 hidden tests are the citation-registry arms, including Finding 2's
  new rot arm. Guard moved to the end.
* **iter-186** added the `derive_sections` derivation without classifying it, so
  `derivation_registry`'s completeness fence was RED at HEAD. Classified **REGISTERED**.

**The pattern is the finding, and the registry already documents it.** Its DECLINED block records that
`anchor_subject_census` went RED at iter-163 *"and stayed RED, unlooked-at, through iters 163/164/165,
because each ran a change-derived scoped suite that did not include the module grading it."* Iters 185
and 186 then did exactly that. `§5` rule 60, sharpened: **the module that GRADES your change is rarely
in the scope your change derives.**

### The instrument caught its own author

Pass 42's harden-origin-route fence went RED on its **first run against pass 42's own ledger entry** —
the new `SURVEY-M257x-h42-size-preserving-mutation-proofs-unaudited` route it had just created was
invisible to the backlog fence and undispositioned. Now dispositioned. A fence that fires on the pass
that wrote it is the cheapest possible evidence it is not vacuous.

> **The id is spelled in full above on purpose.** It was first written abbreviated — the id cut short
> after its `h42-` segment by a prose ellipsis — and an ellipsis is not an abbreviation to a reader that
> enumerates: `route_disposition_guard.ID_RE` stops at the ellipsis and yields a **truncated id**, the
> same producer iter-183 and iter-184 documented for line wraps, manufactured here by prose instead.
> It cost a wrong count within the hour (next paragraph). Note this note does not reproduce the broken
> token: quoting a malformed id to explain it re-emits it, which is iter-98's rule — *write the
> retraction in the vocabulary the fence enumerates.*

> **Count reconciliation — pass 42's table of `5` is a snapshot, not a standing fact.** Pass 42
> measured **5** harden-origin routes *before its own entry existed*; that entry routed the 6th. **Any
> figure written here is stale the moment a later pass routes anything** — pass 44 then routed the 7th,
> which is the regress, and the reason no standing number belongs in this prose. The live value has
> exactly one reader: `_harden_origin_routes()` in
> `stack-core/tests/test_harden_origin_route_visibility_m257x.py`, which is also the fence that fails if
> a route goes undispositioned. Quote **it**, at a ref. Recorded
> because this pass then made the milestone's own signature error while checking it: a probe written
> here with a **private** prefix regex — instead of the module's `HARDEN_ORIGIN_RE` — counted **7**, by
> admitting the truncated fragment above. The module's own reader says 6. *A count about a population
> is unreadable until it names the derivation that produced it* (iter-177), and a harden pass is not
> exempt from its own rule.

**Coverage delta on touched files:** 2 net-new arms (population-by-property; carve-out rot) + 2
pre-existing REDs cleared. **Tests added:** `test_suite_census_population.py` +1 ·
`test_predicate_enumerator.py` +1.
**Bugs surfaced + fixed inline:** 4 — the census misassignment, the rotted carve-out (4 dead members),
the hidden test class, the unclassified derivation.
**Flakes stabilized:** none.

**A vacuous RED-proof, caught and redone — recorded because it is this milestone's own class:** the
first RED-proof of Finding 2 ran in a scratch clone with **no corpus beside it**, so `setUpClass` set
`root = None` and the arm **skipped** — `32 passed, 3 skipped`, which reads green. Re-done in a mirror
where the corpus is reachable: baseline **35 run · 0 skipped**, each of the four dead tails re-added
fails the rot arm, restore green. *A skipped arm proves nothing, and a skip inside a green line is
invisible.*

**Suite results (counts, never wall-time; section scope named):**

| suite | runner | section scope | result |
|---|---|---|---|
| 10 touched + consuming modules | pytest 3.9.6 | `stack-core` only | **174 passed · 0 failed** |
| the two previously-RED fences | pytest 3.9.6 | `stack-core` | **57 passed · 0 failed** |
| `test_predicate_enumerator` | unittest 3.14.6 / 3.9.6 | `stack-core` | **OK** (35 each) |
| RED-proof battery, mtime-mitigated | pytest 3.9.6 | `stack-core` | 4/4 RED, baseline + restore green |

**NOT COVERED (`§5` rule 60):** the whole-section `stack-core` run is reserved for pass close with the
tree frozen. The other **ten** sections and the Go/TS suites were not run.

**Knowledge backfill:** none as a corpus edit this pass — all four defects are in instruments, and each
instrument now carries its own measurement and the ref it was taken at. Pass 42's `§5` rule 77 stands as
this batch's corpus contribution.

**Routed forward (Fate 3):** both pass-42 routes unchanged and still open
(`FIX-M257x-h36-labeled-prover-denominator`,
`SURVEY-M257x-h42-size-preserving-mutation-proofs-unaudited`), each now carrying a written disposition
inside a fence that goes RED if it is dropped.

**Stop condition:** continue-to-next-pass — four more real defects, two of them RED **at HEAD** and
unnoticed because the iters that caused them ran scoped suites excluding their graders. Coverage has not
stabilized.

---

## Pass 44 — 2026-08-09 — incremental

**Iters hardened this pass:** iter-177 … iter-186 (same batch; target: the whole-population claim
itself, taken with the tree **frozen**)
**Tiks covered since prior pass:** 10 (same batch as passes 42–43)
**Runner named on every count** (`§5` rules 75/76): `/usr/bin/python3` **3.9.6** (pytest); the new arm
additionally under `python3` **3.14.6**, `/usr/bin/python3` unittest, **and direct execution**.
**Section scope:** **`stack-core` only** — this is the whole-*section* run, not a whole-*population*
run. Ten sections and both non-Python toolchains were not run (iter-186's rule).

### Finding — the FOURTH way a test file hides its tests (`cf0ba62`)

`test_test_collection_fence.py` already fences three shapes: a `TestCase` below the `__main__` guard
(statement order), a module that does not import at all, and iter-182's needs-a-runner shape. The
fourth is the quietest, and it was **inside the fence's own population the whole time**.

A module whose tests are all module-level `def test_*` functions defines nothing `unittest` can
collect. The stdlib runner then reports:

```
Ran 0 tests in 0.000s
OK
```

**A pass that executed nothing** — this milestone's defect shape verbatim, one layer below where the
file was already looking.

Measured statically over the fence's own population (no imports; the section run was live):

| section | test modules | `TestCase` tests | module-level fns |
|---|---|---|---|
| stack-core | 67 | 1,530 | **25** |
| demo-stack | 35 | 959 | 0 |
| stack-injection | 7 | 311 | 0 |
| stack-verify | 5 | 239 | 0 |
| dev-stack | 5 | 118 | 0 |
| **total** | **119** | **3,157** | **25** |

**Exactly one module** has the shape — `stack-core/tests/test_claim_census_guard.py`, 25 tests —
confirmed on all three runners: pytest 3.9.6 collects **25**; unittest 3.9.6 collects **0 and prints
OK**; unittest 3.14.6 **cannot import it at all** (no pytest on that interpreter). One of those 25 is
the `REXT_SECTION_NAMES` disk-drift fence iter-149 built, so it currently runs under **one** runner.

The arm does not forbid the shape — it forbids the shape being **UNDECLARED**: iter-186's rule (*a
correct exclusion is still a defect while it is silent*) applied to a **runner** instead of a language,
fenced both ways so the declaration cannot rot. **Conversion of the one member is ROUTED, not taken** —
iter-182 measured that the obvious translation silently takes a module from 16 collected to 12.

### Two self-inflicted defects, both caught by this pass's own instruments

* **A private regex, in the pass that published iter-177's rule.** A probe written here with an ad-hoc
  prefix pattern — instead of the module's own `HARDEN_ORIGIN_RE` — counted **7** harden-origin routes.
  The module's reader says **6**. The spurious member was a truncated id.
* **The truncated id was manufactured by this ledger's own prose**: pass 43 abbreviated a route id with
  an ellipsis, and `ID_RE` stops there. Both are corrected in place above, and the correction
  deliberately does **not** reproduce the broken token (iter-98).

Neither reached a commit. Both are recorded because the milestone's standing caution is that passes
33–35 *"created two defects no control they wrote caught"* — these were caught, by controls written
in this session, within minutes.

**Coverage delta on touched files:** 3 net-new arms on a fence that had 3 shapes and now has 4; the
`stack-core` section run moved **1,548 → 1,621 passed** across this batch (+73: iters 178–186 plus this
session's 10 tests), **0 failed** both times.
**Tests added:** `test_test_collection_fence.py` +3.
**Bugs surfaced + fixed inline:** 1 (the undeclared unittest-invisible module) + the 2 self-inflicted
above.
**Flakes stabilized:** none. **Flake gate:** 3 consecutive runs of this session's new modules —
**73 passed** ×3.

**Suite results (counts, never wall-time; scope named):**

| suite | runner | section scope | result |
|---|---|---|---|
| **whole section, tree FROZEN** | pytest 3.9.6 | **`stack-core` only** | **1,621 passed · 2 skipped · 0 failed** |
| 9 touched + consuming modules (post-ledger-edit) | pytest 3.9.6 | `stack-core` | **168 passed · 0 failed** |
| new arm | unittest 3.14.6 / 3.9.6 / direct exec | `stack-core` | **OK** ×3 |
| RED-proof battery, mtime-mitigated | pytest 3.9.6 | `stack-core` | **5/5 RED**, baseline + restore green |
| `guard_family` · `derived_count_guard` · `route_disposition_guard` | 3.9.6 | fence family · corpus · 3 milestones | **19 GREEN · 0 RED** · **OK, 42 sites** · **OK, 0 contradictions** |

**NOT COVERED, stated rather than implied (`§5` rule 60 + iter-186):** the ten non-`stack-core` sections
and the **264 Go test files / 45 TypeScript specs** they carry. **1,621 is a section number, not a
population number** — the exact conflation iter-186 was about, so it is not restated here.

**Knowledge backfill:** none this pass; `§5` rule 77 (pass 42) is this batch's corpus contribution.

**Routed forward (Fate 3), three open:**
* `FIX-M257x-h36-labeled-prover-denominator` — unchanged, open, now dispositioned inside a fence.
* `SURVEY-M257x-h42-size-preserving-mutation-proofs-unaudited` — unchanged, open.
* `FIX-M257x-h44-claim-census-guard-is-single-runner` — **NEW.** Convert
  `test_claim_census_guard.py`'s 25 pytest-style functions to `TestCase` so both runners collect them,
  **carefully**: iter-182 measured that the obvious translation loses tests silently. Declared and
  fenced meanwhile.

**Stop condition:** **cap reached without stabilization** — three passes, **seven defects fixed inline
and three routed**, and every pass found real defects in the previous iters' newest instruments.
Coverage has not stabilized. The batch's shape is one sentence: *the instruments audit their subject
and not themselves* — a backlog fence blind to the file harden passes route into, a partition that
cannot see a misassignment, a registry fenced for overlap but not occurrence, two fences RED at HEAD
because the iters that broke them ran suites excluding their graders, a runner-invisible module, and a
toolchain that can hide a mutation from the proof of the fence that mutation is testing. Per the
user's standing ruling this is **routed and NOT met with new machinery**; the **ninth**
cap-without-stabilization in this milestone (22, 25, 26, 29, 32, 35, 38, 41, 44).

## Pass 45 — 2026-08-09 — incremental

**Iters hardened this pass:** iter-187 … iter-196 (the batch's own thesis, turned on the batch:
*an instrument's declared population is a registry, and it rots at whatever grain nobody checked*)
**Tiks covered since prior pass:** 10
**Runner named on every count** (`§5` rules 75/76): pytest **3.9.6**, unittest **3.9.6**, unittest
**3.14.6**; the Go arm `go1.26.5 darwin/arm64`.
**Section scope:** **`stack-core` only** for the Python arms. The Go figures are **all six Go
sections**; the TypeScript figures are **enumerated, never run** (iter-196's vocabulary holds here).

### Baseline, taken with the tree FROZEN before any edit

`suite_census.py` whole-population run, both runners: **122 modules · GREEN 118 · ENV-GATED 4 ·
RED 0 · TIMEOUT 0**, **0 modules on which the runners disagree**, **0 actionable RED**.
Tests: **3,502 (unittest) / 3,527 (pytest)** — the 25-test gap is exactly
`FIX-M257x-h44-claim-census-guard-is-single-runner`, still open, and it reconciles.

### Finding 1 — the memberships were derived and the SIZES were literals, and one of them was a VERDICT

Iters 187–196 derived every population they **listed** and hand-wrote every population they
**counted**: eight per-section file counts inside `LANGUAGE_EXCLUDED_SECTIONS` / `UNREAD_IN_COLLECTED`,
`264 Go + 75 TS` repo-wide in the header, and on the **always-printed** scope line
`2,714 passed · 0 failed` — a verdict the printing invocation did not measure. Measured this pass: all
eight agree with disk, **and nothing read any of them**. Every membership arm fires when a section
enters or leaves a bucket; a section that keeps its name and loses two hundred test files was invisible
to all of them. That is the live rule *print the SIZE, assert the SHAPE* with its halves swapped, and
iter-192's *an agreeing reconstruction is indistinguishable from a reading* one grain further down.

Two components, both proven rather than argued:

* **A repo-wide arm that could not fail.** `test_the_repo_wide_unread_population_is_STATED_not_implied`
  asserted `sum(A) + sum(B) == sum(A|B)` — an **identity** whenever A and B are disjoint, which
  `test_the_two_languages_PARTITION_the_repo` already guarantees. It could not fail for **any** file
  count, and the figures its name promises to keep readable — **264 and 75** — were asserted **nowhere
  in the class**; they appeared only in its docstring.
* **The fencing ran opposite to the strength of the claim.**

  | published claim | strength | live arms on the REAL repo | ratchet |
  |---|---|---|---|
  | TypeScript 424 tests / 75 files | POPULATION | **3** | yes |
  | Go 264 `*_test.go` | POPULATION | non-empty only | no |
  | Go **2,714 passed · 0 failed** | **VERDICT** | **0** | no |

  `go_census(REPO)` was called by **zero** arms — its only two callers pass synthetic one-test temp
  dirs. **Cost cannot explain it**: measured here the six-section Go census re-runs in **23 s** and the
  TypeScript enumeration that *is* ratcheted costs **0.59 s**. Both affordable; only the weaker claim
  was fenced.

Repair: published figures become named constants the fences read (`GO_FIRST_READING`,
`TS_FIRST_ENUMERATION`, `NON_PYTHON_FILE_FLOOR`) so a printed number and its ratchet are the same
object; `declared_file_counts()` / `derived_file_counts()` give **sizes** the declared/derived split
iter-150 gave **memberships**, with an unreadable declaration a REFUSAL rather than a pass; the identity
becomes two real repo-wide floors; and the Go verdict gets its first arm — live behind
`SUITE_CENSUS_GO_LIVE=1`, whose **skip names the exact command**, because *a NOT-REACHED clause is a
measurement or it is a mood*. Live re-derivation: **2,714 pass · 0 fail, 24.7 s**.

### Finding 2 — the prover called its own recall unreadable and exited 0

`labeled_spelling_pins.prove()` prints, on a broken bucket partition, *"The RECALL above is unreadable
as a rate; fix the accounting before quoting it"* — and then fell through to **`return 0`** and
published that rate. `spelling_pin_census.py --labeled-set` returns the value straight to the shell, so
**every exit-code-keyed caller read a pass** over a denominator the instrument had just disowned.

Reachable and not contrived: `declared_blind` is derived over the whole registry while `unreadable`
counts every unreadable instance, so the buckets over-count by exactly `#(blind AND unreadable)` — the
day a declared-blind instance's file moves. Proven by monkeypatching `lines_at` for the one blind
instance: **PARTITION BROKEN printed, rc=0**.

**iter-193 built the control for this branch and asserted on the printed MESSAGE**, because its `_run`
harness **discarded `prove`'s return value**. The check stayed decoration for exactly as long as the
control written to prove it wasn't. *Assert the verdict, not the message.*

### The repair that was written, measured, and NOT taken

Netting the blind-and-unreadable overlap out of `declared_blind` makes the three buckets sum to
`len(LABELED_SET)` **identically, for every possible input** — the branch becomes **unreachable** and
iter-193's arm goes RED for the right reason. A tidier ratio bought by deleting the check (`§9`: *a good
repair can destroy the proof the instrument fires* — and it is Finding 1's identity, re-created by my
own hand two hours after finding it). Written, measured, reverted, and now pinned by a reachability arm.

**Coverage delta on touched files:** `test_suite_census_population.py` **30 → 35** arms (+1 rewritten
from an identity to two floors); `test_spelling_pin_census_m257x.py` **22 → 23** (+4 arms upgraded from
message-assertions to verdict-assertions).
**Tests added:** +5 / +1, and 5 existing arms strengthened.
**Bugs surfaced + fixed inline:** 2 (both above). **Self-inflicted and caught by this pass's own
controls:** 2 — the netting repair above, and a `pgrep -f suite_census.py` wait-loop that **matched its
own command line** and would never have exited.
**Flakes stabilized:** none. **Flake gate:** 3 consecutive runs of the two modules — **57 passed ×3**.

**Suite results (counts, never wall-time; runner + scope + language named):**

| suite | runner | section scope | language | result |
|---|---|---|---|---|
| whole population, tree FROZEN (baseline) | pytest + unittest 3.9.6 | 5 collected | Python | **122 modules · 118 GREEN · 0 RED · 0 disagreements** |
| 12 touched + consuming modules | pytest 3.9.6 | `stack-core` | Python | **240 passed · 1 skipped · 0 failed** |
| the 2 changed modules | unittest 3.9.6 / 3.14.6 | `stack-core` | Python | **OK ×2** |
| live Go verdict ratchet | `go1.26.5 darwin/arm64` | all 6 Go sections | **Go** | **2,714 pass · 0 fail** (24.7 s) |
| RED-proof battery, mtime-mitigated (`§5` r77) | unittest 3.9.6 | `stack-core` | Python | **8/8 RED**, restore green |
| `story_org_count_guard` · `route_disposition_guard` · `derivation_registry` | 3.9.6 | corpus · 3 milestones · rext | Python | **OK, 164 files** · **OK, 0 contradictions** · **0 unclassified** |

**Graded clean, recorded so it is not re-derived** (`§5` — grade findings individually, never
characterise a set): `story_org_count_guard`'s printed denominator and its violation scanner are **two
readers** (`rglob` + per-file filter vs `os.walk` + dir-pruning). Enumerated both: **164 = 164, zero
divergence**, and the traversals are equivalent on this interpreter (`rglob` does not skip dot-dirs).
**No defect.** iter-191's repair is sound.

**NOT COVERED, stated rather than implied (`§5` rule 60):** the four non-`stack-core` **Python** sections [**four**, CORRECTED at pass 51 — this entry read *ten* as written; the four are `demo-stack`, `dev-stack`, `stack-injection`, `stack-verify`];
the **424 TypeScript tests**, which remain **ENUMERATED and never executed** — iter-196's vocabulary is
in force and no count here is a TypeScript verdict.

**Knowledge backfill:** none this pass; the two rules it contributes (*assert the verdict, not the
message*; *a fence's ratchet should track the STRENGTH of the claim, not its convenience*) are recorded
here and in the commit bodies.

**Routed forward (Fate 3), four open:**
* `FIX-M257x-h36-labeled-prover-denominator` — **CLOSED by Finding 2's repair** (the verdict half was
  the remainder).
* `SURVEY-M257x-h42-size-preserving-mutation-proofs-unaudited` — unchanged, open.
* `FIX-M257x-h44-claim-census-guard-is-single-runner` — unchanged, open; the 3,502/3,527 baseline gap
  above is its live size.
* `SURVEY-M257x-h45-printed-measurement-literals-uncensused` — **NEW.** iter-193's
  `printed_arithmetic_totals` censuses a printed total assembled by ARITHMETIC over a registry. The
  strictly simpler sibling — a printed total that is a **LITERAL of a past measurement** — has no
  census at all, and this pass fixed its three instances **by hand, in one module**. Per the user's
  standing ruling the general census is routed, not built.

**Stop condition:** **continue-to-next-pass** — two defects found and fixed in the first pass over this
batch, both inside instruments shipped in the last ten iters, so the coverage delta has not been
measured against a second pass yet.

## Pass 46 — 2026-08-09 — incremental

**Iters hardened this pass:** iter-187 … iter-196 (same batch; dimension 3 — **error paths and refusal
gates**, asked of the whole population rather than one module)
**Tiks covered since prior pass:** 10 (same batch as pass 45)
**Runner named on every count** (`§5` rules 75/76): pytest **3.9.6**, unittest **3.9.6**, unittest
**3.14.6**.
**Section scope:** **`stack-core` only**; **Python only** — no Go or TypeScript figure is quoted here.

### The sweep, and why it was worth doing as a population

Pass 45's second finding was a refusal that printed and returned 0. Rather than trust that it was
singular, this pass **enumerated the class**: every `print` in the batch's seven instruments whose text
speaks about the *instrument's own validity* (`CANNOT RUN`, `⚠`, `BROKEN`, `UNRUNNABLE`, `STALE`,
`vacuous`, `not GREEN`, `UNDECLARED`, `REFUS…`) — **16 sites** — each graded against what its enclosing
function can still return.

### Finding — the census reports two kinds of registry rot and grades only one

`suite_census.main`'s exit expression was `1 if actionable or stale else 0`.

**In it:** `stale_declarations` — a declared TEST that no longer exists.

**Printed with a ⚠ and returned 0:**

* `stale_excluded_sections` — a declared language-exclusion naming no directory on this tree. **This is
  the rot iter-192 wrote that function to expose.**
* a section on disk in **neither** bucket — the partition, broken on the real tree rather than in a unit
  test.
* `SECTIONS_ARE_DERIVED is False`, where the scope line prints, verbatim: *"scope: UNMEASURED — the
  section list is the frozen iter-186 literal … The figures below describe that declaration, not this
  tree."*

Two rot reporters, two consequences, and **nothing declaring the difference** — the module states no
exit-code contract anywhere. Three sibling guards in this repo (`story_org_count_guard`,
`claim_census_guard`, and `labeled_spelling_pins` since pass 45) already answer this situation with
CANNOT RUN. A census that can print a total over a denominator it has just disowned, and exit green, is
this batch's defect **inside the machinery built to end it**.

Repair: the three breaches are collected and returned as **2** (CANNOT RUN — the code this module
already uses when the pytest half is unavailable), outranking the **1** it returns for findings. A RED
module is a finding *about* the population; a breach means there is no population to have findings
about. The list also lands in `--json`.

**Folded into the final return rather than failing fast, on purpose.** Failing fast at the detection
point would leave the three ⚠ prints unreachable — a check deleted rather than enforced, which is
exactly the trap pass 45 walked into one module over and reverted. Stated cost: a breach is learned
after the run, not before.

### Graded individually, found clean — recorded so they are not re-derived

`§5` — grade findings individually, never characterise a set (pass 32 characterised 21 failures as
"provably not ours" and iter-145 proved 57 % of that false).

* `claim_census_guard` `COULD NOT RUN` ×3 and `no baseline on disk` → **`return 2`**. Correct.
* `claim_census_guard` `RATCHET BROKEN` → **`return 1`**. Correct.
* `claim_census_guard` **`STALE SUBSTRATE`** → non-fatal, and the message itself says a stale substrate
  *"produces evidence AGAINST a true one"*. Graded: the substrate feeds `_live_names`, so it can move
  the tier-2 count — but **only upward**, so its error direction is toward a false RED, never a hidden
  GREEN. Conservative, therefore defensible; the **direction and the verdict consequence are
  undeclared**, which is iter-186's rule. Routed, not fixed — and it is outside this batch's diff
  scope (iter-188 touched this module's `SKIP_DIRS`, not its substrate logic).
* `route_disposition_guard` **`N ambiguous segments REFUSED`** (3 live) → non-fatal. Graded: refusing to
  interpret an ambiguous segment is a declined reading, not a claimed one; it is named, counted in the
  summary line, and the guard carries anti-vacuity floors elsewhere. **No defect.**
* `labeled_spelling_pins` `⚠ N instance(s) could not be read` → non-fatal **by design**: those instances
  are accounted, named, and removed from the denominator with the direction stated. Correct as it
  stands (its sibling, the partition breach, was pass 45's finding).

**Coverage delta on touched files:** `test_suite_census_population.py` **35 → 39** arms. Across both
passes this batch's two most-modified test modules moved **30 → 39** and **22 → 23**.
**Tests added:** +4 (three breach arms driving the real `main()` end-to-end, narrowed to one module so a
full census is not needed — ~3 s each — plus the anti-vacuity control that a healthy narrowed census
still exits 0). Anything less would have fenced the message again instead of the verdict.
**Bugs surfaced + fixed inline:** 1.
**Flakes stabilized:** none. **Flake gate:** 3 consecutive runs — **38 passed ×3**.

**Suite results (counts, never wall-time; runner + scope + language named):**

| suite | runner | section scope | language | result |
|---|---|---|---|---|
| the changed module | unittest 3.9.6 / 3.14.6 | `stack-core` | Python | **39 passed · 1 skipped · 0 failed**, both |
| 5 consuming modules | pytest 3.9.6 | `stack-core` | Python | **131 passed · 1 skipped · 0 failed** |
| RED-proof battery, mtime-mitigated (`§5` r77) | unittest 3.9.6 | `stack-core` | Python | **4/4 RED**, restore green |
| `derivation_registry` completeness | 3.9.6 | rext | Python | **0 unclassified · 0 printed-arithmetic totals** |

**NOT COVERED, stated rather than implied (`§5` rule 60):** the four non-`stack-core` Python sections [**four**, CORRECTED at pass 51 — this entry read *ten* as written; the four are `demo-stack`, `dev-stack`, `stack-injection`, `stack-verify`]; Go
and TypeScript were not re-read this pass and no figure here describes them.

**Knowledge backfill:** none; the rule this pass contributes — *two rot reporters in one function must
either share a consequence or declare why they do not* — is recorded here and in the commit body.

**Routed forward (Fate 3), five open:**
* `SURVEY-M257x-h42-size-preserving-mutation-proofs-unaudited` — unchanged, open.
* `FIX-M257x-h44-claim-census-guard-is-single-runner` — unchanged, open.
* `SURVEY-M257x-h45-printed-measurement-literals-uncensused` — unchanged, open.
* `SURVEY-M257x-h46-stale-substrate-direction-undeclared` — **NEW.** `claim_census_guard`'s
  STALE SUBSTRATE warning is conservative (false-RED only) and says neither that nor its verdict
  consequence. Declare the direction, or grade it.
* `FIX-M257x-h36-labeled-prover-denominator` — closed at pass 45; listed here only so the transition is
  legible.

**Stop condition:** **continue-to-next-pass** — a second consecutive pass found a real defect of the
same class in a different module of the same batch, so the class is confirmed repeating and the
coverage delta has not settled. One pass remains before the incremental cap.

## Pass 47 — 2026-08-09 — incremental

**Iters hardened this pass:** iter-187 … iter-196 (same batch; dimension 2/5 — **boundary and input
shape**, aimed at the three parsers this batch shipped)
**Tiks covered since prior pass:** 10 (same batch as passes 45–46)
**Runner named on every count** (`§5` rules 75/76): pytest **3.9.6**, unittest **3.9.6**, unittest
**3.14.6**; the Go arm `go1.26.5 darwin/arm64`.
**Section scope:** the Python figures are **`stack-core` only**; the Go figures are **all six Go
sections**; TypeScript was not re-read.

### The whole-section measurement, tree FROZEN

**1,699 passed · 3 skipped · 0 failed** — pytest 3.9.6, `stack-core` only, Python only, 23m11s, taken
with no edit in flight. Pass 44's comparable reading was **1,621 · 2 · 0**: **+78**, being iters
187–196's own arms plus this session's **+10**. The third skip is this session's opt-in live Go arm.
**1,699 is a SECTION number**, not a population number.

### Finding — `2,714` counts test CASES, and it was published three times as test FUNCTIONS

This batch's headline figure — iter-195's *"the first-ever reading"* — was labelled **test FUNCTIONS**
at three sites, each contrasting it with iter-186's **264 FILES** under the heading *"Name the unit
(`§5`, iter-177)"*. Measured this pass:

| measurement | value | instrument |
|---|---|---|
| test **CASES** (what `go_census` tallies) | **2,714** | `go test -json`, `Test` field present |
| top-level test functions | **2,204** | same run, `Test` name without `/` |
| **subtests** | **510** (18.8 %) | same run, `Test` name with `/` |
| `func Test…` declarations | **2,186** | static scan of 264 `*_test.go` |
| `*_test.go` files | **264** | disk (agrees with iter-186 exactly) |

`go test -json` puts a `Test` field on **subtest** events too, so the tally admits them. One population,
three measurements, and the corpus published the largest under the name of the smallest — **the rule
against unit errors, containing one, in the number it was written to protect.** A reader checking 2,714
against a `func Test` grep would find 2,186 and conclude the count had rotted.

**Graded fairly:** the parser was already careful in the direction that had been considered —
`if ev.get("Test")` correctly excludes package-level `pass` events, which would have inflated the total
by one per package. Subtests are the case nobody asked about, and the one the unit word is wrong for.

Repair: `go_census` derives the split per section (`subtests`); `GO_FIRST_READING` carries
`funcs` / `subtests` / `static_decls` so the three measurements can never again be quoted as one; the
scope line and the `--go` TOTAL say **CASES** and print the split. Verified live — the tool now reports
**2,714 = 2,204 + 510**, matching the independent measurement exactly.

**Coverage delta on touched files:** `test_suite_census_population.py` **39 → 40** arms, plus two new
assertions inside the live ratchet (a subtest floor and an anti-vacuity arm: zero subtests across six Go
sections is a broken split, not a repo without table-driven tests). Across the three passes this module
moved **30 → 40**.
**Tests added:** +1 arm, +3 assertions.
**Bugs surfaced + fixed inline:** 1.
**Flakes stabilized:** none. **Flake gate:** 3 consecutive runs — **39 passed ×3**.

**Suite results (counts, never wall-time; runner + scope + language named):**

| suite | runner | section scope | language | result |
|---|---|---|---|---|
| **whole section, tree FROZEN** | pytest 3.9.6 | **`stack-core` only** | Python | **1,699 passed · 3 skipped · 0 failed** |
| the changed module | unittest 3.9.6 / 3.14.6 | `stack-core` | Python | **40 passed · 1 skipped**, both |
| live Go ratchet (`SUITE_CENSUS_GO_LIVE=1`) | `go1.26.5 darwin/arm64` | all 6 Go sections | **Go** | **40 passed · 0 skipped** |
| RED-proof battery, mtime-mitigated (`§5` r77) | unittest 3.9.6 | `stack-core` | Python | **4/4 RED**, restore green |
| `derivation_registry` completeness | 3.9.6 | rext | Python | **0 unclassified** |

**NOT COVERED, stated rather than implied (`§5` rule 60):** the four non-`stack-core` Python sections [**four**, CORRECTED at pass 51 — this entry read *ten* as written; the four are `demo-stack`, `dev-stack`, `stack-injection`, `stack-verify`];
the **424 TypeScript tests**, still **ENUMERATED and never executed**.

**Knowledge backfill:** none; the rule this pass contributes — *a unit is a claim, and the sentence that
names the unit is not exempt from checking it* — is recorded here and in the commit body.

**Routed forward (Fate 3), four open (unchanged this pass):**
`SURVEY-M257x-h42-size-preserving-mutation-proofs-unaudited` ·
`FIX-M257x-h44-claim-census-guard-is-single-runner` ·
`SURVEY-M257x-h45-printed-measurement-literals-uncensused` ·
`SURVEY-M257x-h46-stale-substrate-direction-undeclared`.

**Stop condition:** **cap reached without stabilization** — three passes, **four defects fixed inline**
and four routed, and **every pass found a real defect in the previous iters' newest instruments**.
Coverage has not stabilized. The batch's shape in one sentence: *the instruments were built to grade
their subject's populations, and nobody graded the instruments' own **sizes**, **verdicts** or
**units*** — a repo-wide arm that was an identity, a Go verdict fenced by nothing while the weaker
TypeScript population was ratcheted twice, a prover that called its own recall unreadable and exited 0,
a census that could disown its denominator and exit 0, and a headline figure carrying the wrong unit in
the sentence that names units. Per the user's standing ruling this is **routed and NOT met with new
machinery**; the **tenth** cap-without-stabilization in this milestone (22, 25, 26, 29, 32, 35, 38, 41,
44, 47).

## Pass 48 — 2026-08-09 — incremental

**Iters hardened this pass:** iter-197 … iter-206 (dimension 1/3 — **the batch's own derived figures**,
re-derived rather than read, because iter-206 measured this batch's error rate on freshly-written
derived figures at **4 of 6**)
**Tiks covered since prior pass:** 11
**Runner named on every count** (`§5` rules 75/76): unittest **3.9.6**, `stack-core` section only,
**Python** only. No Go, no TypeScript arm this pass.

### The scope hole this pass opened on

iter-206 derived **six** standing figures in `claim_census_guard.py`'s comments and found **four**
stale, **two** of them written by the same run. It graded exactly one module. The module the batch
actually spent itself on — `derivation_registry.py`, **530 net-new lines across iters 199/200/203/204/205**,
and the module that *implements* the comment and docstring censuses — was never turned on itself.

### Finding — one shared classifier, two unequal windows (22 of 121, 18.2 %)

iter-205 factored `_classify_measurement` out precisely so the docstring census and the comment census
*"cannot drift apart on the classification while sharing a name for it."* **They drifted on the
ARGUMENT.** A `#` comment is tokenised per LINE, so the comment caller handed the classifier
`tok.string` while the docstring caller hands it the whole string — and the classifier's whole basis
for `dated` is a ±120-char context window it could therefore never fill from a comment:

```
# M257x iter-197 wrote this and it rotted:
# the repo has 121 modules today      → standing  (as a comment)
                                      → dated     (the same sentence, as a docstring)
```

| measurement | value |
|---|---|
| comment matches misclassified | **22 of 121 (18.2 %)** |
| rows moved `standing` → `dated` | **21** |
| `standing` bucket | **95 → 75** |
| population SIZE change | **0** — a classification defect, never a reach defect |

The published reading that comments carry *"a far higher standing share than either sibling, which
fits what comments are for"* is therefore **part artefact of the window**, not a fact about comments.

**The fence for this exact property passed throughout.** `test_the_classifier_is_SHARED_so_the_two_cannot_drift`
greps both callers' *source* for the callee's NAME — **a name check cannot see what is passed to the
name.** New rule: *two callers of one rule are only as shared as the argument they hand it.*

**Fixed as a pair, because the first alone imports a second defect.** `_DOC_RELATIVE` was searched over
the same ±120 neighbourhood, so a number could borrow an unrelated *"65 lines away"* beside it. Latent
while comments were read one line at a time (**measured: 0 of 7** doc-relative docstring rows borrowed)
and live the instant the window widens — this module's own `205 of 695` specimen sits two lines under
two genuine ones and flipped. Now decided on the match's own phrase.

### Four more, every one a figure this batch wrote

1. `comment_measurement_literals`'s docstring published **"101 measurement-shaped numbers across 89
   sites"** as a standing size. Live: **117 across 101** — both operands stale one iter after writing,
   and **101 had become the SITE count**, so a reader checking the number would have been *reassured*.
2. `DOCSTRING_LITERAL_CEILING`'s own block narrated *"RE-BASELINED to 160"*, then *"162, not the 160"*,
   beside a constant reading **164**. Three numbers for one value; iter-205's reason lived only in a
   commit message, which is not where *"may not grow without a recorded reason"* can be read.
3. *"Reporting the split with its own miss rate stated"* was a **MOOD** — no miss rate had ever been
   measured, while the sibling making the same promise names the audit that measures its own
   (`claim_census_guard.py:101`). Measured now, and the first reading is the 18.2 % above.
4. The `205 of 695` specimen is marked **REFUTED** (iter-202 derived **292 of 704**) — an unmarked
   refuted figure inside the module that finds refuted figures is the class illustrating itself.

### The pass committed the error it was auditing, three times, and that is recorded

The comment-ceiling paragraph named **126** (a prediction), then **120** (measured before the note
explaining it existed), then **124** (measured after) — because **a provenance note joins the
population it explains**, exactly as `noun_vocabulary_reach`'s docstring did at iter-204. `D-M257x-203-2`,
committed while writing the paragraph that warns about it. Stable form: that paragraph carries **no
figure**, and its sibling's arrow chain is **fenced** instead. Ceilings re-taken from the census after
the last edit — comments **118 → 121**, docstrings **164 → 165**; `matches` joined `_MEASURED_NOUNS`,
forced by iter-204's residual arm going RED on this pass's own sentence.

**Coverage delta on touched files:** `test_frozen_expectation_census_m257x.py` **76 → 83** arms
(2 new classes, 7 arms).
**Tests added:** +7 arms.
**Bugs surfaced + fixed inline:** 6 (commit `92b25c1`).
**Flakes stabilized:** none.

**Suite results (counts, never wall-time; runner + scope + language named):**

| suite | runner | section scope | language | result |
|---|---|---|---|---|
| the changed module | unittest 3.9.6 | `stack-core` | Python | **83 passed** (76 before) |
| the four sibling batch modules | unittest 3.9.6 | `stack-core` | Python | **88 passed** |
| RED-proof battery, mtime-mitigated (`§5` r77) | unittest 3.9.6 | `stack-core` | Python | **4/4 RED**, restore sha-verified, post-restore green |

**NOT COVERED, stated rather than implied (`§5` rule 60):** the whole-section pytest reading is **not
taken at this pass** — the tree is still being edited, and nine runs in this milestone have been
discarded as confounded for exactly that; it is taken once at session end with the tree frozen. The ten
non-`stack-core` Python sections, the six Go sections and the **424 TypeScript tests** are untouched
this pass.

**Knowledge backfill:** the rule this pass contributes — *two callers of one rule are only as shared as
the ARGUMENT they hand it* — is recorded here and in the commit body.

**Routed forward (Fate 3), four open (unchanged) + one new:**
`SURVEY-M257x-h42-size-preserving-mutation-proofs-unaudited` ·
`FIX-M257x-h44-claim-census-guard-is-single-runner` ·
`SURVEY-M257x-h45-printed-measurement-literals-uncensused` ·
`SURVEY-M257x-h46-stale-substrate-direction-undeclared` ·
**`SURVEY-M257x-h48-the-censuses-cannot-see-a-bold-wrapped-operand`** — `_MEASURED_RE` needs
whitespace after the closing operand, so this repo's own emphasis idiom (`**292 of 704**`) is invisible
to all three censuses. Measured incidentally while writing this pass's own notes; sizing it moves all
three ceilings again, so it is routed rather than bundled.

**Stop condition:** `continue-to-next-pass` — the classifier/window defect is closed and fenced, but the
sweep of the batch's derived figures has covered one module of the three the batch touched
(`suite_census.py` and `claim_census_guard.py`'s non-comment figures are unread).

## Pass 49 — 2026-08-09 — incremental

**Iters hardened this pass:** iter-197 … iter-206 (same batch; dimension 1/3 continued — the module
pass 48 could not reach, `suite_census.py`)
**Tiks covered since prior pass:** 11 (same batch as pass 48)
**Runner named on every count:** unittest **3.9.6**, `stack-core` only, **Python**; plus one **Go**
re-derivation, `go1.26.5 darwin/arm64`, all **6** Go sections.

### Finding — iter-197's module count was wrong on the day it was typed

`suite_census.py` said *"measured 23 s serial for **122 modules** × 2 runners"* and its fence's own
docstring said *"the whole **122-module** × 2-runner census."* Derived at every relevant ref:

| ref | live `modules()` | what the prose said |
|---|---|---|
| `1d7e5cf` (iter-197, the iter that wrote it) | **123** | 122 |
| `e1b7345` (iter-201) | **124** | 122 |
| `ede026d` (iter-206 / batch HEAD) | **124** | 122 |

**Wrong when written, then wrong again** — and iter-197's *own prose* reads *"1 of 123"* in the same
iter. This is the second instance of the defect that iter's progress notes describe as *"a size
literal that rotted inside its own iter"*; the first was found, these two were not.

### Finding — a THIRD copy that the census could not see: the separator was `\s+`

`suite_census.py:889` carried `122-module` — **invisible to all three measurement-literal censuses**,
because `_MEASURED_RE` required whitespace between number and noun. It flagged the space-separated
twin **two lines away** and was blind to the hyphenated one. Sized before widening: **21 attributive
compounds across 13 modules** (`3-service floor`, `98-site sweep`, `93-repo register`, `60-line`, …).
The same shape as iter-205's case-sensitivity find, **one character over** — and it was hiding a
defect at the time it was measured, which the case-sensitivity find also was.

### Finding — the ratchet was anti-vacuity wearing a ratchet's badge

`MIN_MODULES = 100` against a live population of **124**: **two dozen** test modules could have left
this repo without one arm moving. That is precisely how three wrong prose copies survived ten iters.
Replaced by `MODULE_FLOOR = 124`, taken from the census, plus an arm requiring the collection census to
cover the same population the floor is taken over.

### Finding — the fencing register had gone stale in the safe direction

`suite_census.py:130`'s table still read `Go 2,714 passed · 0 failed | **none — 0 arms** | no`. That
column is the **pass-45 diagnosis**; **harden pass 47 gave the Go verdict its live arm.** Wrong in the
harmless direction and still costly: a register of *which claims are fenced* is the one table that,
stale, points the next reader at work already done. Now labelled *"live arm AT PASS 45"* with pass 47's
outcome recorded beneath it.

### What re-derived CLEAN, stated because a sweep that only reports defects is not a measurement

* the **Go verdict**: `2,714 pass · 0 fail · 510 subtests · 2,204 top-level`, six sections — **exact**;
* `424 tests / 75 files` TypeScript — exact (`215/45` + `209/30`);
* `5 of 11 sections`, three sites — exact;
* `35 modules` for demo-stack — exact, **and now derived rather than merely correct**: `DERIVED_PROSE_COUNTS`
  exempts a prose count from the new ban only because the same entry recomputes it from the live census
  and requires the rendered text verbatim. **The exemption IS the derivation**; a count that is merely
  correct gets none;
* `claim_census_guard.py` — **zero** ungraded standing docstring figures. iter-206's fence covers its
  six comment figures, and this pass found nothing left in it. The batch's last iter did its job on the
  module it chose.

**Coverage delta on touched files:** `test_suite_census_collection.py` **16 → 19** arms;
`test_frozen_expectation_census_m257x.py` **83 → 84**.
**Tests added:** +4 arms.
**Bugs surfaced + fixed inline:** 5 (commit `7b73cff`).
**Flakes stabilized:** none.

**Suite results (counts, never wall-time; runner + scope + language named):**

| suite | runner | section scope | language | result |
|---|---|---|---|---|
| the two changed fence modules | unittest 3.9.6 | `stack-core` | Python | **100 passed** |
| five sibling batch modules | unittest 3.9.6 | `stack-core` | Python | **143 passed · 1 skipped** |
| Go verdict re-derivation | `go1.26.5 darwin/arm64` | all 6 Go sections | **Go** | **2,714 pass · 0 fail** |
| RED-proof battery, mtime-mitigated (`§5` r77) | unittest 3.9.6 | `stack-core` | Python | **4/4 RED**, restore sha-verified over 3 files |

**NOT COVERED (`§5` rule 60):** the whole-section pytest reading — taken once at session end, tree
frozen. The four non-`stack-core` Python sections [**four**, CORRECTED at pass 51 — this entry read *ten* as written; the four are `demo-stack`, `dev-stack`, `stack-injection`, `stack-verify`] and the **424 TypeScript tests** are unrun as ever.

**Knowledge backfill:** *the exemption IS the derivation* — a prose figure earns its place by being
recomputed, not by being right. Recorded here and in the commit body.

**Routed forward (Fate 3):** the four standing entries, plus
`SURVEY-M257x-h48-the-censuses-cannot-see-a-bold-wrapped-operand` — **now the only open reach hole of
its family**, since this pass closed the hyphen one. `**292 of 704**` is still invisible: the closing
operand needs whitespace after it, and this repo bolds its important figures.

**Stop condition:** `continue-to-next-pass` — the three modules the batch touched are swept, but the
batch also published derived figures in **iter progress/decisions markdown**, which no census reads at
all, and pass 48's own headline (`22 of 121`) has not been re-derived since the regex widened.

## Pass 50 — 2026-08-09 — incremental

**Iters hardened this pass:** iter-197 … iter-206 (same batch; dimension 1/4 — **re-derive what passes
48 and 49 themselves published**, plus the milestone-level figure they invalidate)
**Tiks covered since prior pass:** 11 (same batch as passes 48–49)
**Runner named on every count:** unittest **3.9.6** and pytest **3.9.6**, `stack-core` only, **Python**.

### CORRECTION to iters 205/206 — the standing class is 147, not 168

Pass 48 fixed the classifier's window. This pass asked what that does to the number the **milestone**
published. Re-derived at **iter-206's own tree** (`ede026d`), with **iter-206's own vocabulary**, and the
window fixed — so the only variable is the defect:

| site-kind | published | re-derived at `ede026d` |
|---|---|---|
| comments | 118 rows · **95 standing** | 117 rows · **74 standing** |
| docstrings | 164 rows · 73 standing | 164 · **73** — correct, unaffected |
| **the standing class** | **168 sized · 7 derived · 161 unverified** | **147 · 7 · 140** |

**Inflated by 21.** iter-205's headline — comments carry *"a far higher standing share than either
sibling"* — survives as a **fact** but not at the size it was given: **63 % against 45 %**. And iters
205 and 206 published **95** and **96** for the same reading in adjacent progress lines. Recorded as a
marked correction in the milestone's own `progress.md`, appended rather than substituted.

### Three of THIS SESSION'S derived figures were wrong, and that is the finding

The batch's disclosed error rate reproduced itself inside the hardening of it:

1. *"22 of 121 matches (18.2 %)"* was written into **two docstrings two hundred lines apart with two
   DIFFERENT `standing` figures** — `95 → 73` in the module, `95 → 75` in the fence. Live at the time:
   **75**. One pass later the hyphen widening moved every operand again (**31 of 137** today).
2. *"21 sites across 13 modules"* for the hyphen gain was a **pre-widening dry run reported under the
   WRONG UNIT** — they were rows. Live: **25 rows at 22 sites across 14 modules**. `§5` r75, *name the
   unit*, missed inside the paragraph announcing a reach fix.
3. Pass 48's comment-ceiling paragraph named **126 → 120 → 124** across three drafts (already recorded).

**The conclusion is not "be more careful."** Across passes 48–50 every single derived figure written
into prose went stale or wrong within one pass, including the ones written by the pass that was
auditing exactly that. The class is not a property of iters 197–206; it is a property of **carrying a
derived figure in prose at all**. So the repair is iter-206's, turned on our own writing:
`classifier_window_miss_rate()` derives the rate on every run, **no docstring carries it**, and an arm
greps both files for the shape. The hyphen figure is stated **once**, in `_MEASURED_RE`'s own note,
with its unit and its wrong first draft beside it.

**Coverage delta on touched files:** `test_frozen_expectation_census_m257x.py` **84 → 86** arms
(the live-rate arm now reads the derived helper **and** requires the per-line window to over-report
`standing` — the defect's *direction*, not merely its existence).
**Tests added:** +2 arms, +1 derived production helper.
**Bugs surfaced + fixed inline:** 3 (commit `5f4b779`).
**Flakes stabilized:** none.

**Suite results (counts, never wall-time; runner + scope + language named):**

| suite | runner | section scope | language | result |
|---|---|---|---|---|
| **whole section, tree FROZEN** | pytest **3.9.6** | **`stack-core` only** | Python | **1,799 passed · 3 skipped · 0 failed** |
| the two changed fence modules | unittest 3.9.6 | `stack-core` | Python | **102 passed** |
| all seven modules this batch touches | unittest 3.9.6 | `stack-core` | Python | **245 passed · 1 skipped** |
| RED-proof battery, mtime-mitigated (`§5` r77) | unittest 3.9.6 | `stack-core` | Python | **2/2 RED**, restore sha-verified |
| **flake gate** | unittest 3.9.6 | `stack-core` | Python | **12 passed ×3 consecutive** |

**1,799 is a SECTION number, not a population number.** Pass 47's comparable reading was **1,699**:
**+100**, of which **+13** are this session's arms (passes 48/49/50 = 7 + 4 + 2) and **+87** are iters
197–206's own. Taken with **no edit in flight** — the tree was frozen from the last commit of pass 50
through the run, per the nine runs this milestone has discarded as confounded.

**NOT COVERED, stated rather than implied (`§5` rule 60):** the four non-`stack-core` Python sections [**four**, CORRECTED at pass 51 — this entry read *ten* as written; the four are `demo-stack`, `dev-stack`, `stack-injection`, `stack-verify`];
the six Go sections beyond the one verdict re-derivation recorded at pass 49; the **424 TypeScript
tests**, still **ENUMERATED and never executed** — **no TypeScript verdict is claimed by any of passes
48–50.**

**Knowledge backfill:** two rules, recorded here and in the commit bodies —
*two callers of one rule are only as shared as the ARGUMENT they hand it* (pass 48), and
*a derived figure carried in prose goes stale between the paragraph announcing a fix and the paragraph
implementing it — the repair is not to carry it* (pass 50, earned three times in three passes).

**Routed forward (Fate 3), four standing + one new:**
`SURVEY-M257x-h42-size-preserving-mutation-proofs-unaudited` ·
`FIX-M257x-h44-claim-census-guard-is-single-runner` — **note: iter-201 closed the runner gap 25 → 0;
this route's title now describes a closed condition and should be retired or re-scoped at close** ·
`SURVEY-M257x-h45-printed-measurement-literals-uncensused` ·
`SURVEY-M257x-h46-stale-substrate-direction-undeclared` ·
`SURVEY-M257x-h48-the-censuses-cannot-see-a-bold-wrapped-operand` (new at pass 48; the hyphen sibling
was closed at pass 49, this one is not — `**292 of 704**` needs whitespace after the closing operand
and this repo bolds its important figures).

**Stop condition:** **cap reached without stabilization** — three passes, **14 defects fixed inline**
and one routed, and **every pass found a real defect in the previous iters' newest instruments — and
in its own.** Coverage has not stabilized. The batch's shape in one sentence: *the instruments were
built to grade prose figures, and the figures the instruments themselves publish were graded by
nobody* — a shared classifier fed two unequal windows, a module count wrong on the day it was typed
and wrong twice more since, a ceiling narrated at three values at once, a stated miss rate that was a
mood, two reach holes (case at iter-205, separator here) each of which was hiding a live defect at the
moment it was measured, and a milestone-level denominator inflated by 21. Per the user's standing
ruling this is **routed and NOT met with new machinery**; the **eleventh** cap-without-stabilization in
this milestone (22, 25, 26, 29, 32, 35, 38, 41, 44, 47, 50).

**The pass's own honesty note.** Passes 48–50 wrote **three** wrong derived figures of their own — a
miss rate published with two different operands two hundred lines apart, a reach gain reported under
the wrong unit, and a ceiling paragraph that named three values across three drafts. Each was caught
by re-derivation inside this session, not by review. That is the same **~1-in-3** rate iter-206
measured for the batch, reproduced by the pass auditing it, which is the strongest available evidence
that the class is structural rather than a lapse of the ten iters under audit.

## Pass 51 — 2026-08-09 — incremental

**Iters hardened this pass:** iter-207 … iter-216
**Tiks covered since prior pass:** 10
**Runner named on every count:** unittest **3.9.6** and pytest **8.4.2 / CPython 3.9.6**, section scope
stated per row, **Python**. No Go and no TypeScript verdict is claimed anywhere in this entry.
**Diff scope:** 20 files, all `rosetta-extensions/stack-core`, `5f4b779..7ba6054`.

### THE HEADLINE: the milestone's own frozen-expectation fence was RED at HEAD for seven iters

Measured, not inferred. Each of the batch's own commits was unpacked with `git archive` and censused by
**that tree's own** `derivation_registry`:

| commit | docstring | comment | test-module | |
|---|---|---|---|---|
| iter-207 `f6c6f27` | 181 / 181 | 141 / 141 | 462 / 462 | exact, on the day it was taken |
| iter-210 `0b17938` | 184 / 181 | 144 / 141 | 465 / 462 | **BREACHED** |
| iter-216 `7ba6054` | 188 / 181 | 148 / 141 | 477 / 462 | **BREACHED, and shipped** |

Running `test_frozen_expectation_census_m257x.py` at `7ba6054` reproduces **9 genuine failures** (plus
5 artifacts of an archived tree not sitting inside the rosetta worktree). **Every iter close in this
batch reported `N passed / 0 failed`** over a scoped module set, and **not one of those sets contained
this fence.** `§5` rule 60 — *a scoped green is evidence about its scope alone* — at the highest price
this milestone has paid for it: the file passes 48, 49 and 50 were spent deepening was RED at HEAD
throughout, including while those passes were deepening it.

The 9 are three independent defects wearing seven faces, all repaired inline:

1. **All three literal ratchets breached.** Re-pinned **181 → 195**, **141 → 153**, **462 → 487**, each
   with the recorded reason its contract demands, written where the ratchet lives.
2. **Four orphaned decisions + two unclassified derivations — one defect, two faces.** iter-212's fold
   turned `scan_targets` in four guards into a one-line delegation: **the NAME survives, so any check
   looking for the name sees nothing**, while the derivation SITE left the registry's population. Four
   `DECLINE:tree-scan` decisions were orphaned and the two sites that inherited the work
   (`corpus_sources`, `claude_docs_outside_skills`) arrived unclassified. iter-214's rule (*a waiver
   outliving its subject reads as coverage*) meeting iter-212's own (*two readers of one construct must
   SHARE the derivation*) **in a single commit**: the sharing was done correctly and the registry of
   what-derives-what was left describing the tree before it.
3. **Four undecided measurement nouns** — `collectors`, `pins`, `duplicates` admitted; **`iters` sent to
   `_NOT_NOUNS`**, because `iter` is already an ORDINAL word and admitting its plural would have made
   every *"iter 210"* reference a measurement literal. The vocabulary's reach closing on the sentence
   that widened it, for the fourth time.

**Repair beyond the re-pin.** A ratchet that says only `477 > 462` makes a **blind bump** the cheapest
available response — and a ratchet answered with a blind bump has become the chore its own contract
warns against. `rows_by_file` now folds the heaviest contributing modules into all three breach
messages, so the recorded reason is *writable*.

### The two items owed to this pass, both re-derived here and neither carried

**1 — the rule-60 clause: there are FOUR, not ten.** Derived with `suite_census.python_sections()`: 11
sections on disk, **10 non-`stack-core`**, of which **4 carry Python** — `demo-stack`, `dev-stack`,
`stack-injection`, `stack-verify`. All five wrong clauses (passes **45, 46, 47, 49, 50**) corrected in
place with a provenance marker so no reader mistakes the correction for what that pass wrote.

**And the correct value was already in this file.** **Pass 40** wrote *"the four non-`stack-core`
sections"* — the right NUMBER with an ambiguous noun — thirteen hundred lines above five entries
carrying the right noun with the wrong number. **Neither pass ever held both halves, and both halves
were in this one file the whole time.** Pass 40's line is annotated, not corrected: it is the exhibit.

Re-read with the tree frozen (pytest **8.4.2 / CPython 3.9.6**, each section its own invocation):
`demo-stack` **1,063 passed · 9 failed · 2 skipped** · `dev-stack` **151** · `stack-injection` **335** ·
`stack-verify` **275** — **1,824 passed**, and the 9 failures are **exactly** the 9 declared `ENV_GATED`
entries, **0 undeclared**.

**2 — the standing class, re-derived at a STATED ref rather than copied.** At **`7ba6054`** (the batch's
closing commit): docstrings **188 rows / 85 standing**, comments **148 / 85**, **union 336 distinct rows
/ 170 standing**, overlap 0. **Fourth value for one class: 168 → 147 → 157 → 170.**

**And the drift is the smaller half of the story.** All four values describe the **non-test** population
only, while the population iter-207 exposed — `test_*.py` modules — measures **477 rows / 323 standing**
at the same commit. **The milestone-level "standing class" has always named ~41 % of the repo's
measurement-literal rows and has never said so.** Post-pass at `0248f38` the censused union reads **348
/ 179** and the excluded population **487 / 332**; both moved because this pass wrote prose, which is
the mechanism, not an accident.

### The pass's own three findings, beyond the two owed items

**F1 — the env-gated bucket's key cannot identify its subject** (commit `f2a9b34`). `ENV_GATED` is keyed
`module-relpath::test-name` and **both** readers drop the CLASS. Demonstrated on a staged tree, **not
argued**: declare `NeedsTheWorld::test_x`, let an undeclared `PlainLogic::test_x` fail for a real reason,
and `run_one` returns **ENV-GATED under both runners** — the `failures` list carrying the same key twice,
the instrument holding the evidence of its own ambiguity and discarding it. A genuinely actionable RED
reported as the platform's absence is the exact inversion `§5` r73 exists to prevent, committed by the
bucket written to honour it. Exposure **LIVE and sized**: **15 test methods across 5 files** share a
(file, name) key, three pairs of them inside `test_frozen_expectation_census_m257x.py` itself. None of
the nine declarations sits in such a file today, so **nothing is mis-absorbed right now** — iter-213's
discipline, report the reach and never a false-GREEN it did not have. Repair is **fail-closed** and
**deliberately does not parse the class out of runner output**: measured here, the two interpreters this
module runs both of disagree about the failure line (3.9.6 `FAIL: test_x (mod.Alpha)`, 3.14.6 `FAIL:
test_x (mod.Alpha.test_x)`), so keying on that shape would make the bucket's correctness a property of
which python ran it. ⚠ The **sibling-survives** staleness direction is **structurally out of reach at
this key grain** and is routed, not claimed; the ast read did widen the other half — four staged shapes
where the old `f"def {test}(" in source` reported a **deleted** test as present.

**F2 — the family's denominator is a `len` and every check is a `set`** (commit `9cb833d`). iters
210–212 folded five private derivations onto `fence_provenance.corpus_sources` and fenced **the
sharing** — correctly. Every one of those arms, **including iter-212's census-by-effect**, compares a
SET; the number the family publishes (*"1,801 citations over 114 sources"*, *"8 collectors, 114 each"*)
is the **`len` of a concatenated list**. Staged: an `EXTRA_SOURCES` entry inside `corpus/` made the
denominator read **4** where every arm saw **3**, all fifteen symmetric differences still zero. The
batch's own class **one grain below where it was fixed**. Order-preserving dedup; **114 before, 114
after**, which is the half a de-duplication must always prove.

**F3 — three self-corrections inside this pass, left visible.** (a) The `claude_docs_outside_skills`
rationale's first draft named its sibling in backticks and **correctly went RED** on
`ARationaleThatAssertsASetRelationIsGRADED`: the true claim is a **DIFFERENCE**, which the `RELATION:`
grammar cannot express and whose operands `_resolve_operand` would resolve empty regardless (it passes
the guard dir, never a repo root). Routed rather than smuggled. (b) That fix's own pointer **named an arm
that does not exist**. (c) The mutation control **resurrected one orphan and instantly armed the new
rationale against itself** through a backticked name — de-armed in place with the reason. Three in one
pass, all caught by re-derivation rather than review, which is the same **~1-in-3** rate iters 205–206
measured and passes 48–50 reproduced.

**Coverage delta on touched files:** `test_suite_census.py` **11 → 19** arms; `test_corpus_citation_guard.py`
**31 → 35**; `test_frozen_expectation_census_m257x.py` **91 → 95**.
**Tests added:** +16 arms (8 + 4 + 4), +2 production helpers (`defining_classes`/`ambiguous_declarations`,
`rows_by_file`), 1 dedup, 3 ceiling re-pins, 6 registry entries changed, 4 nouns decided.
**Bugs surfaced + fixed inline:** **12** — 9 pre-existing REDs at HEAD (`0248f38`), the ENV_GATED
absorption + staleness pair (`f2a9b34`), the `len`-vs-`set` denominator (`9cb833d`).
**Flakes stabilized:** none surfaced.

**Suite results (counts only; runner + section scope + language named on every row):**

| suite | runner | section scope | language | result |
|---|---|---|---|---|
| the frozen-expectation fence, before | unittest 3.9.6 | `stack-core` | Python | **9 genuine failures at `7ba6054`** |
| the frozen-expectation fence, after | unittest 3.9.6 | `stack-core` | Python | **95 passed · 0 failed** |
| `test_suite_census` | unittest 3.9.6 | `stack-core` | Python | **19 passed** |
| the `suite_census` family (4 modules) | unittest 3.9.6 | `stack-core` | Python | **98 passed · 1 skipped** |
| the `corpus_sources` consumers (4 modules) | unittest 3.9.6 | `stack-core` | Python | **155 passed** |
| the four non-`stack-core` Python sections | pytest 8.4.2 | `demo-stack`, `dev-stack`, `stack-injection`, `stack-verify` | Python | **1,824 passed · 9 failed · 2 skipped** — all 9 declared `ENV_GATED`, 0 undeclared |
| RED-proof battery, mtime-mitigated (`§5` r77) | unittest 3.9.6 | `stack-core` | Python | **3/3 mutations fired**, all restores sha-verified |

**NOT COVERED, stated rather than implied (`§5` rule 60), and this clause is DERIVED this pass rather
than copied:** the **six** non-`stack-core` sections that carry no Python at all (`alignment`,
`clerkenstein`, `playthroughs`, `stack-secrets`, `stack-seeding`, `stack-snapshot`) — their Go suites
were not re-read here and **no Go verdict is claimed**; the **424 TypeScript tests**, still
**ENUMERATED and never executed**. The four Python sections above *were* read this pass and are
reported, not disclaimed — which is exactly the confusion iter-208 found in the five clauses this pass
corrected.

**Knowledge backfill:** three rules, recorded here and in the commit bodies —
*a fold that keeps the NAME and moves the DERIVATION is invisible to every check that looks for the
name* (F-headline 2); *the published figure is a `len` over a list and every check is over a set — put
distinctness in the derivation, not in each reader* (F2); and *a ratchet whose breach message names no
file has already chosen the blind bump for you* (the repair beyond the re-pin).

**Routed forward (Fate 3), two new:**
`SURVEY-M257x-h51-env-gated-key-drops-the-class` — re-key `ENV_GATED` to carry the class so the
sibling-survives staleness direction becomes reachable; pinned by an executable arm that goes RED the
day someone does it ·
`SURVEY-M257x-h51-relation-grammar-cannot-express-a-difference` — `RELATION: A == B [| C]` plus a
resolver that passes the guard dir make a difference-claim ungradeable; both halves need work before a
rationale can state one.

**Stop condition:** **continue-to-next-pass** — the pass repaired 12 defects and every one of them was
in an instrument this milestone built to catch that exact class, so coverage has not stabilized; the
whole-section reading is still owed.

## Pass 52 — 2026-08-09 — incremental

**Iters hardened this pass:** iter-207 … iter-216 (same batch as pass 51; dimension 3 — error paths,
and the reach of the instruments that grade this milestone's own record)
**Tiks covered since prior pass:** 10 (same batch as pass 51)
**Runner named on every count:** unittest **3.9.6**, `stack-core` only, **Python**; the guard-family
readings are `guard_family.py` over the **rosetta** corpus at `001d1b8`, also Python.

### The whole guard family, read — and it is green

The reading nobody in this batch took. `guard_family.py --repo-root <rosetta> --allow-not-run`:
**19 GREEN · 0 RED · 0 could-not-check · 8 not-run**, and with `--range` supplied for the batch's own
rosetta commits, **20 GREEN · 0 RED · 2 could-not-check · 5 not-run**. The two that could-not-check —
`repair_leak_guard`, `value_change_guard` — are **correct refusals**: the batch's rosetta diff added
prose and removed none, so there were no candidate shingles and no replaced token run. **A refusal that
fires for the right reason is the control this family was built to have**, and it is worth recording as
the one place this pass looked and found nothing wrong.

### THE FINDING: the ledger records every measurement, and 46 of its 51 entries are graded by nothing

`derived_count_guard` prints `OK — 57 site(s): 48 derivable count(s) equal their derivation` over a
subject of **829 files**. Attributed at this pass:

| | |
|---|---|
| sites | **57** |
| files carrying at least one site | **30** of 829 |
| sites inside `hardening-ledger.md` | **11** — the heaviest single contributor |
| ledger pass entries carrying a site | **5** of **51** — passes 23, 25, 30, 41, 44 |
| ledger pass entries carrying NONE | **46**, including every entry since pass 44 |

So the document that records every measurement this milestone takes is graded almost nowhere — and the
three entries that were *explicitly auditing their own derived figures* (passes 48–50) are among the 46.
iter-215's rule, *a check that SKIPS reads exactly like one that PASSES*, on the highest-value subject
available. It is also the **structural explanation** for the standing observation that derived figures
here are wrong at roughly one in three: **nothing grades them.**

**The reach is NOT widened, deliberately.** Arms A/B/C/D grade table totals, explicit deltas, percent
triples and clause-5 dispositions — not prose figures — and that is correct as designed. Widening moves
a live RED surface and would need iter-209's zero-false-REDs precondition paid first. What is repaired
is the **silence**: the report now derives and states how many scanned files contributed no site, and
what a zero there means (`NOT CHECKED — it was not cleared`). Printed above the verdict line, so
`guard_family.run_one`'s `lines[-1]` reporting is untouched — family view re-run, unchanged.

### The pass's own control that did not fire, and the arm was the reason

Worth more than the fix. `test_the_disclosure_is_DERIVED_from_the_sites_never_restated` searched for a
digit run preceded by a **quote** (`'"829'`), while a hard-coded figure inside an f-string is preceded
by a **space** — so replacing `{len(_site_files)}` with a literal `30` left the arm **green**. iter-205's
case-sensitivity class and iter-210's `corpus`-vs-`CORPUS_DIR` class for a third time, and it surfaced
only because the mutation was **run** rather than assumed — harden pass 45's caution, *check your own
new arms can fail*, earned rather than quoted. The arm now reads the REACH line, strips interpolations,
and refuses any remaining digit; re-mutated after strengthening, it fires. Kept in the docstring as an
executable record rather than deleted.

**Coverage delta on touched files:** `test_derived_count_guard.py` **17 → 21** arms.
**Tests added:** +4 arms, +1 derived disclosure line.
**Bugs surfaced + fixed inline:** 2 — the undisclosed silence (commit `2aa68cc`) and this pass's own
non-firing control, fixed in the same commit.
**Flakes stabilized:** none surfaced.

**Suite results (counts only; runner + section scope + language named on every row):**

| suite | runner | section scope | language | result |
|---|---|---|---|---|
| `test_derived_count_guard` | unittest 3.9.6 | `stack-core` | Python | **21 passed** |
| guard family, tree-scoped | `guard_family.py` | rosetta corpus @ `001d1b8` | Python | **19 GREEN · 0 RED · 8 not-run** |
| guard family, `--range d3f1c64..062df38` | `guard_family.py` | rosetta corpus | Python | **20 GREEN · 0 RED · 2 correct refusals** |
| RED-proof battery, mtime-mitigated (`§5` r77) | unittest 3.9.6 | `stack-core` | Python | **2 mutations, 1 fired then 1 after the arm was fixed**; restores sha-verified |

**NOT COVERED, stated rather than implied (`§5` rule 60), derived this pass:** the **six** non-`stack-core`
sections carrying no Python (`alignment`, `clerkenstein`, `playthroughs`, `stack-secrets`,
`stack-seeding`, `stack-snapshot`) — **no Go verdict is claimed**; the **424 TypeScript tests**, still
**ENUMERATED and never executed**. The four Python sections were read at pass 51 and are not re-claimed
here. The whole-`stack-core` section reading is owed and is taken at pass 53, tree frozen.

**Knowledge backfill:** one rule — *a site count without its silent remainder reads as coverage of the
whole subject; derive the remainder and say what a zero there MEANS.*

**Routed forward (Fate 3), one new:**
`SURVEY-M257x-h52-derived-count-guard-reaches-no-prose-figure-in-the-ledger` — the ledger's pass entries
state their figures as prose, which arms A–D do not reach by design; widening needs the zero-false-REDs
precondition paid first, and until it is, the ledger's own numbers are checked by re-derivation alone.

**Stop condition:** **continue-to-next-pass** — the silence is disclosed but not closed, and the
whole-section reading this batch has never had is still outstanding.

## Pass 53 — 2026-08-09 — incremental

**Iters hardened this pass:** iter-207 … iter-216 (same batch as passes 51–52; the whole-section
reading this batch never had, plus the structural half of pass 51's headline)
**Tiks covered since prior pass:** 10 (same batch as passes 51–52)
**Runner named on every count:** pytest **8.4.2 / CPython 3.9.6** for the section reading, unittest
**3.9.6** for the fences; **`stack-core` only**, **Python**. No Go and no TypeScript verdict is claimed.

### THE WHOLE-SECTION READING, tree frozen — and it came back RED

Taken at `2aa68cc` with the rext tree frozen from pass 52's last commit through the run (three earlier
attempts this session were **discarded as confounded** the moment an edit landed mid-run, per the nine
this milestone has already thrown away):

> **`stack-core` — 4 failed · 1,850 passed · 3 skipped**, pytest 8.4.2 / CPython 3.9.6, Python.

For scale: pass 50's comparable reading was **1,799 passed · 3 skipped · 0 failed**. The **+51** is this
batch's arms plus passes 51–53's own; the **4 failed** is the finding.

### THE FINDING: the defect reproduced itself one pass after being documented

All four failures are ceiling breaches — **the same three ratchets pass 51 had just re-pinned with
recorded reasons.** The cause is pass 52's **own** four arms and its **own** comment block, written one
pass later, in the same session, by the same author, with the recorded reason sitting directly above the
constant it invalidated.

So pass 51's headline was right about the facts and **wrong about the class**. It is not that iters
207–216 forgot: **any commit that writes prose moves these populations**, and the only reading available
cost a 95-arm import of a fence nothing in the iter loop runs. *A check whose cheapest invocation is a
suite is a check that gets skipped* — and then the ratchet is the thing that rots, which is precisely
what the seven-iter history shows.

**Repair — the reading gets a one-command form.** `derivation_registry.py --ceilings` prints live vs
ceiling for all three, names the heaviest contributing modules on a breach, and exits non-zero. It
asserts nothing the arms do not; it makes the cost a second instead of a suite, so an iter's close step
can carry it. A bare invocation **errors** rather than reading — a `__main__` that acts by default is a
footgun, and a silent default read would look like a green. Re-pinned **COMMENT 153 → 159** and
**TEST_MODULE 487 → 492**; **DOCSTRING held exact at 195**.

The fence then caught this pass's own `ceiling_report` as an unclassified derivation — **the same
registry arm that caught iter-212's fold at pass 51, now catching its author.** Registered.

### CORRECTION to pass 51's own coverage figures — 1 of 4 was wrong

Appended, not substituted, per this milestone's standing practice. Pass 51 published
`test_suite_census.py` **11 → 19** arms. Re-derived by running the module at `7ba6054`: **Ran 12 tests**.
The correct arrow is **12 → 19**. The other three published pairs are correct —
`test_corpus_citation_guard` **31 → 35**, `test_frozen_expectation_census_m257x` **91 → 95**,
`test_derived_count_guard` **17 → 21**.

The wrong operand came from counting `def test_` with `grep`, which also matches the **staged test
modules written inside fixture strings** — a count of a different population wearing the arm count's
name. `§5` r75, *name the unit*. **One in four**, which is the rate the brief predicted, the rate
iter-206 measured, the rate passes 48–50 reproduced, and now the rate passes 51–53 reproduced too.

### One clean zero, and it is a real zero

Pass 52 found `CHANGELOG.md` — a git-tracked root document, 639 lines — is in **no** fence's population:
`EXTRA_SOURCES` is `("README.md", "CLAUDE.md")` and there is nothing declaring the third root document
either in or out. Sized before judging: **17 distinct backticked `corpus/`/`.claude/` paths and 2
markdown links**, against README's 5/12 and CLAUDE.md's 88/8. **Every one of the 18 references resolves
on this tree**, so there is no live defect — only an undeclared population boundary, and a plausible
reason to exclude (a changelog describes past states, so a stale path in it may be correct history
rather than a defect). **Routed rather than landed**, because declaring it is a population change and
iter-209's zero-false-REDs precondition governs those.

**Coverage delta on touched files:** `test_frozen_expectation_census_m257x.py` **95 → 99** arms.
**Tests added:** +4 arms, +1 derived reader (`ceiling_report`), +1 CLI verb, 2 ceiling re-pins,
1 registry entry.
**Bugs surfaced + fixed inline:** 3 — the two re-breached ratchets and the unclassified
`ceiling_report` (commit `3965790`).
**Flakes stabilized:** none surfaced; **flake gate 3/3 clean** on the new arms.

**Suite results (counts only; runner + section scope + language named on every row):**

| suite | runner | section scope | language | result |
|---|---|---|---|---|
| **whole section, tree FROZEN at `2aa68cc`** | pytest 8.4.2 / CPython 3.9.6 | **`stack-core` only** | Python | **1,850 passed · 4 failed · 3 skipped** |
| the frozen-expectation fence, after repair | unittest 3.9.6 | `stack-core` | Python | **99 passed · 0 failed** |
| `--ceilings` one-command form | `derivation_registry.py` | `stack-core` | Python | **3 ratchets exact, exit 0** |
| RED-proof battery, mtime-mitigated (`§5` r77) | unittest 3.9.6 | `stack-core` | Python | **2/2 mutations fired**, restores sha-verified |
| flake gate | unittest 3.9.6 | `stack-core` | Python | **4 passed ×3 consecutive** |

**NOT COVERED, stated rather than implied (`§5` rule 60), derived this pass:** the **six**
non-`stack-core` sections carrying no Python (`alignment`, `clerkenstein`, `playthroughs`,
`stack-secrets`, `stack-seeding`, `stack-snapshot`) — **no Go verdict is claimed by any of passes
51–53**; the **424 TypeScript tests**, still **ENUMERATED and never executed**. The four
non-`stack-core` **Python** sections were read at pass 51 (**1,824 passed**) and are not re-claimed here.
⚠ **The 4 failed above is a reading of `stack-core` at `2aa68cc` and is REPAIRED at `3965790`** — the
post-repair whole-section re-read is **not taken**, and this entry does not claim one.

**Knowledge backfill:** two rules —
*a check whose cheapest invocation is a suite is a check that gets skipped; give the reading a
one-command form or expect the ratchet to rot* (this pass, earned by breaking it ourselves), and
*a `grep -c "def test_"` is a count of a different population when the module stages fixtures — name the
unit or run the runner* (the corrected 11 → 12).

**Routed forward (Fate 3), one new, plus the two from pass 51 and one from pass 52:**
`SURVEY-M257x-h53-changelog-is-a-tracked-root-doc-in-no-fence-population` — declare it in or out with a
reason and reconcile both ways; needs the zero-false-REDs precondition paid first.

**Stop condition:** **cap reached without stabilization** — three passes, **17 defects fixed inline**
and four routed, and **every pass found a real defect in the previous passes' newest instruments,
including its own.** Coverage has not stabilized: the pass that documented a ratchet's seven-iter rot
broke that same ratchet one pass later, which is the strongest available evidence that the class is
structural rather than a lapse. Per the user's standing ruling this is **routed and NOT met with new
machinery**; the **twelfth** cap-without-stabilization in this milestone (22, 25, 26, 29, 32, 35, 38,
41, 44, 47, 50, 53).

**The pass's own honesty note.** Passes 51–53 published **one wrong derived figure of four** and were
caught by re-derivation, not review — the same ~1-in-3 rate the batch under audit exhibits. Three
sessions of harden passes have now each reproduced it while auditing it. That is no longer evidence
about iters 207–216; it is a measured property of writing a derived figure into prose at all, and the
only repair this milestone has found that holds is the one both pass 50 and this pass reached
independently: **do not carry the figure — derive it where it is printed.**

## Pass 54 — 2026-08-09 — incremental

**Iters hardened this pass:** iter-217 … iter-228 (12 iters; 5 pre-redirect instrument, 7 post-redirect
platform/corpus). Scope weighted to the post-redirect half per the user's 2026-08-09 redirect — *"the goal
remains alignment and be able to build a working stack with the new platform repos"* — so the two brand-new
instruments that stand between the tooling and a buildable stack were taken first.

**Tiks covered since prior pass:** 12.

**Subject:** `stack-core/clone_pin_guard.py` (iter-222) and `stack-core/patch_anchor_guard.py` (iter-223) —
the artifact that declares *which platform repos a demo clones* and the one that answers *would the demo's
patch layer still apply*. Both shipped in the last two iters; neither had been hardened.

**Bugs surfaced + fixed inline (4, commit `8345c1d`, pushed to `rosetta-extensions` origin):**

- **`clone_pin_guard` arm C compared the pin ref LITERALLY** to `{main, master, HEAD}`, so it caught three
  spellings and no others. `ensure-clones.sh` phase (d3) runs `git checkout -f "$_pref"`, which resolves
  `origin/main`, `refs/heads/main` and `refs/remotes/origin/main` to the same moving branch — **all read
  GREEN**. The "reproducibility BARRIER" could therefore name a different tree every day while the fence
  written to refuse exactly that agreed with it. `moving_branch_name()` normalises first.
  **The shipped test iterated `sorted(G.MOVING_BRANCHES)` — an identity test against the guard's own
  constant** — which is how a gap in the arm survived its own battery. (Pass 45 of this milestone found
  the same shape in a repo-wide arm; that makes twice.)
- **`clone_pin_guard` arm B never REQUIRED either sanctioned extra**, only allowed them, because it is
  derived from `repos.yml` and neither extra is in it. **A pin with no `platform` key read GREEN** — the
  largest possible instance of the exact hole arm B exists to catch, since `platform` is the clone whose
  `repos.yml` defines the topology every other entry is derived against. `platform` is now required;
  `ant-academy` stays optional **and the docstring now says so**, with a test pinning the statement, so the
  remaining hole is declared rather than invisible (`§5` rule 60).
- **`patch_anchor_guard` answered from the cache.** It resolves `--ref` with one command, `git show
  <ref>:<path>`, which reads the local object store — so `--ref origin/main` answered *"would the patch
  layer apply at the last `origin/main` this clone happened to fetch"* and printed, flatly, `OK at
  origin/main`. **This is the class iter-222 measured ONE ITER EARLIER** (nobody had fetched; `app` was 28
  commits stale behind it and 17 corpus anchors were rotten behind that). Adds `--fetch` (rc-checked,
  stderr never suppressed; a **failed** fetch at a remote-tracking ref is **exit 2**, never a verdict) and
  puts the ref's authority *in the verdict line*.
- **`guard_family`'s RED summary named its reds and not its reach.** Measured, same command, two working
  directories: **24 GREEN · 0 RED · 0 could-not-check** from the rosetta root against **4 GREEN · 3 RED ·
  17 could-not-check** from `stack-core/` — and the line an operator quotes forward was, in full,
  `guard-family: RED — demo_knob_guard, dev_flag_guard, platform_predicate_guard`. All three reds are
  artifacts of the wrong corpus root, the same misconfiguration that blocked the other 17, and the quoted
  sentence carried no trace of either. **That is iter-103's misread-RED incident — a RED read against a
  sheet of greens, two false conclusions drafted from the difference — reachable today by running the
  family from the directory the guards live in.** `reach_caveat()` now rides on the summary line, scoped
  to the RED branch because the could-not-check and not-run branches name their own members.
  This is the **third** pass to fix this one sentence (pass-20 the `OK` line, pass-23 the provenance
  caveat, pass-54 the reach) — the pattern is that each fix covered the branch that pass was looking at.

**Plus one in the test layer:** `TestCollectionParity`'s locator, `src.find('if __name__ == "__main__"')`,
**matched its own quoted literal** four lines below. It was checking *"is every class above
`TestCollectionParity`"*, not *"above the `__main__` guard"* — so a class appended in the correct place,
at the foot of the file above the real guard, went RED **naming a cause that was not true**. Anchored at
column zero and taken as the last occurrence, with a regression test and an anti-identity control.

**Tests added:** +21 across 3 modules —
`tests/test_clone_pin_guard.py` 16 → **23** (7: a 9-spelling moving-ref battery, a 7-spelling
false-RED control, the finding's naming, the normaliser as a unit, and 3 required-extra),
`tests/test_patch_anchor_guard.py` 15 → **22** (7: a real-remote fixture whose upstream moves without the
clone fetching, both fetch dispositions, the local-ref control, the predicate as a unit),
`tests/test_guard_family_verdict_line_m257x.py` 30 → **37** (7: 5 reach-caveat incl. the wired-call-site
arm, 2 parity-locator).

**Mutation controls (`§9` — a green must prove its instrument):** 5 pin spellings flipped GREEN → RED with
3 controls holding GREEN and the live canonical pin unaffected; blinding `is_remote_tracking` to the
pre-fix world fired **4 of 7** cache tests (the other 3 do not depend on the predicate, and say so).

**Flakes stabilized:** none surfaced.

**Suites (runner · section scope · language, per the standing rule):**

| suite | runner | section scope | language | result |
|---|---|---|---|---|
| the 3 touched modules | unittest, CPython 3.9.6 | `stack-core` | Python | **45 + 37 passed · 0 failed** |
| collection parity, one module | pytest 8.4.2 / CPython 3.9.6 | `stack-core` | Python | **37 passed** — same 37 |
| `--ceilings` one-command form | `derivation_registry.py` | `stack-core` | Python | **3 ratchets exact, exit 0** |
| `guard_family --platform`, correct root | `guard_family.py` | `stack-core` | Python | **24 GREEN · 0 RED · 5 not-run** |

**NOT COVERED, stated rather than implied:** no whole-`stack-core` pytest reading is claimed by this entry
(one is taken at session end); the four non-`stack-core` **Python** sections were last read at pass 51
(**1,824 passed · 9 failed, all ENV_GATED**) and are not re-claimed; **no Go verdict**; the **424
TypeScript tests** remain ENUMERATED and never executed.

**Knowledge backfill:** two rules, both earned here —
*a fence that compares a ref to a literal set has the reach of that set, not of the property it names —
`git` resolves more spellings than any list you will write, so normalise before you compare*; and
*a summary line must state its REACH as well as its verdict, because a misconfigured reference blocks some
members and reddens others from the same cause, and only the verdict gets quoted forward.*

**Stop condition:** continue-to-next-pass — four defects in the newest instruments, all found by the first
dimension scan of the pass; the corpus-prose half of the batch (iters 224–228) has not been swept.

## Pass 55 — 2026-08-09 — incremental

**Iters hardened this pass:** iter-217 … iter-228 (same batch; this pass swept the **corpus-prose** half
the pass-54 entry recorded as unswept).

**Subject:** the fence between the corpus's prose and `repos.yml` — i.e. whether the corpus can go on
naming the wrong platform repos. Directly the redirect's target: *"be able to build a working stack with
the new platform repos (only the remaining ones that are still part of it)."*

**Method:** by-effect census, the method iter-219 established — mutate the reference, ask which guards
notice, rather than reading code and inferring. Two mutations of a scratch `repos.yml` copy:

| mutation | `clone_pin_guard` | `platform_alignment_guard` | `platform_predicate_guard` |
|---|---|---|---|
| ARRIVAL (`brand-new-svc` added, 4 → 5) | RED | RED | RED — 10 G2 sites incl. `CLAUDE.md:360` |
| **SWAP** (`studio-desk` → `brand-new-svc`, count stays 4) | RED | RED | **GREEN** |

**Bugs surfaced + fixed inline (2, commit `91c926e`, pushed to `rosetta-extensions` origin):**

- **G2 grades the COUNT and reads as grading the MEMBERSHIP.** On the swap, ten corpus sites go on naming
  `studio-desk` — `CLAUDE.md:360` among them, **inside a runnable `make init` block, in the file every
  session loads** — and G2 is green because 4 is still 4. **A consolidation program retiring one repo and
  adding another is exactly what produces a count-preserving swap**, which is the program this milestone
  exists to track. New arm **G2b**: where a repo-count claim also ENUMERATES the set on its own line, the
  names must be exactly `repos.yml`'s. **5 enumerations reached live, 0 findings; 5 of 5 fire on the
  swap.** Conservative by construction (only on a line G2 already matched, only after a list delimiter,
  every token must be a known repo name, and the list must be CLOSED on the line).
  **The closed-list clause is not decoration.** The first draft went RED on `staging-bringup.md:118`,
  where the list continues on the next line inside a bash comment —
  a **false RED in the corpus's own onboarding path**, in a module whose stated landing precondition is
  *zero false REDs*. It held only because the arm was run against the live tree before it was believed.
- **The repo-wide test-collection fence was blind to the pattern it is mostly written in.** Both of its
  predicates matched only a base spelled `…TestCase`, so a class inheriting from a **local fixture base**
  — `Baseline(Fixture)`, `MutationBattery(Fixture)`, `AntiVacuity(Fixture)`, … — was invisible.
  Demonstrated, not argued: **pass 54 appended `RemoteTrackingIsACache(Fixture)`, 7 tests, below the guard
  in `test_patch_anchor_guard.py`, and this fence stayed GREEN** while a direct run skipped all 7 and
  printed OK. Bases now resolve transitively, with a mutant reproducing the blind spot and a
  false-positive control (a plain helper class must stay unswept).

**Three of this session's OWN regressions, each caught by a shipped repo-wide fence rather than by
review** — which is the fences working, and is worth recording as such: two hidden test classes (moved
above the guard), two fixtures duplicating the derived clone set (rewritten synthetic — the frozen-
expectation census's preferred repair is *stop duplicating*, not *claim an exemption*), and a stale
derived share in `claim_census_guard`'s docstring, **re-derived at this tree: 293 of 706 distinct-grain,
410 of 981 pair-grain** (it read 704 / 979).

**Tests added:** +8 `tests/test_platform_predicate_guard.py` (`TestG2bRepoSetMembership`: the three
delimiter shapes, the half-list refusal, incidental-mention and unknown-token refusals, the swap
comparison, an anti-identity control, and a **live-corpus reach control** that fails if G2b reaches
nothing); +5 `tests/test_test_collection_fence.py` (`TheFenceSeesLocalFixtureBases`).

**Suites (runner · section scope · language):**

| suite | runner | section scope | language | result |
|---|---|---|---|---|
| the 7 touched modules | pytest 8.4.2 / CPython 3.9.6 | `stack-core` | Python | **430 passed · 0 failed** |
| `--ceilings` one-command form | `derivation_registry.py` | `stack-core` | Python | **3 ratchets exact, exit 0** |
| `platform_predicate_guard`, live | itself | `stack-core` | Python | **exit 0**, G2b reach **5** |
| `platform_predicate_guard`, swap probe | itself | `stack-core` | Python | **exit 1**, 5 G2b findings |

**Knowledge backfill:** one rule — *a fence that grades a COUNT reads as grading the SET; a swap that
preserves the count passes it, and consolidation programs produce swaps.* Filed against the same family
as pass 54's *normalise before you compare*: both are the shape where the check is a proxy for the
property and the proxy is not said out loud.

**Stop condition:** continue-to-next-pass — the count-vs-membership class was found in the first
by-effect probe of the pass, and its sibling question (whether the anchor instruments still read a
FETCHED ref after iter-222's *"a remote-tracking ref is a cache"* finding) is not yet asked.

## Pass 56 — 2026-08-09 — incremental

**Iters hardened this pass:** iter-217 … iter-228. **No instrument was changed this pass** — deliberately.
The user's redirect says the target is the platform and a buildable stack, not deeper tooling, so this
pass is a **reading**: it asks the one question iters 222–228 opened and did not close, using instruments
that already exist.

**The question.** iter-222 established *"a remote-tracking ref is a cache, not a remote"* and re-derived
17 rotted anchors. iter-224 fetched and advanced. But `anchor_construct_guard` — the instrument that
resolves the corpus's anchors — was **not touched by either iter**, and its `auto` ladder prefers
`origin/main`. So: **is the corpus aligned to the tree a stack actually BUILDS, or to `origin/main`?**
Those are not the same tree. `clones.pin.json` pins `app` at `ad9f3c49`; `origin/main` is `3eaadae6`.

**Measured — every clone's currency, from `.git/FETCH_HEAD` mtime, no network:** all **13** clones
fetched **1.0 h ago** (iter-224's fetch). Three are behind their own `origin/main` — `app` **28**,
`next-web-app` **12**, `ant-academy` **9** — plus the per-stack `rosetta-extensions` at its pin (159),
which is by design. So the readings taken this session are current; what is unasserted is that they stay
that way.

**Measured — the same corpus, the same instrument, three refs:**

| `CITE_REF` | verdict | reach |
|---|---|---|
| `origin/main` | **exit 0 — OK** | 881/1475 = 59.7 % |
| `auto` (the family default) | **exit 0 — OK** | 881/1475 = 59.7 % |
| `HEAD` (= the pinned tree a demo builds) | **exit 1 — RED, 1 anchor** | 881/1475 = 59.7 % |

The reach is identical at all three and **the verdict is not** — which is worth stating on its own,
because a reader comparing reach numbers across refs would conclude the ref made no difference.

**The one RED, and it is a real corpus defect:** `corpus/ops/observability.md:28` cites
`app/main.go:278` for `colony.WithLoggingTracing(0.15, 0.15)`. At `ad9f3c49` line 278 is `)`; the
construct is at **`:277`**. And the document **already knows this** — its own preamble (`:12-23`) says,
in as many words, *"Every `app` anchor in the table below is pinned to `app` `ad9f3c49`, and to that ref
ALONE"*, and then names the correct pair at that ref: **`:273`/`:277`**. The table cites `:273` and
**`:278`** — the first from the declared pin, the second from `3eaadae6`. **The cell contradicts its own
preamble**, in a block whose preamble exists because *"an earlier revision cited `main.go:278` @
`3eaadae6` and named `ad9f3c49` in the same cell"* and cost exactly one RED to learn. The prose repair
landed; the cell was not re-derived with it.

**And the default masks it.** The family runs `auto`, `auto` prefers `origin/main`, and at `origin/main`
`:278` is correct — so **the guard reads GREEN by resolving at a ref the document explicitly
disclaims.** The two `go.mod` anchors in the same table (`:200`, `:232`) are identical at both refs and
are not implicated.

**Disposition: ROUTED FORWARD (Fate 3), not fixed inline.** Correcting the cell to `:277` makes the
document right at the ref it declares and would turn the **default** family run RED, because the block
pin the preamble states is not the scope the resolver honours for that row. Which of those two moves
first — restructure the document so its pin is inside the block, or widen the resolver's block-pin scope
— is a design decision, which is the inline boundary's own named exit.
`ROUTE-M257x-h56-observability-cell-contradicts-its-own-declared-pin`.

**Also routed, from the same thread:** `anchor_construct_guard` and `clone_drift_guard` run exactly one
git command each and never fetch, so both answer at whatever the local cache holds — **the identical
shape pass 54 fixed in `patch_anchor_guard`**, whose repair (disclose the ref's authority in the verdict;
opt-in `--fetch`; exit 2 on a failed fetch at a remote-tracking ref) is a ready template.
`ROUTE-M257x-h56-anchor-guards-answer-from-the-cache-like-patch-anchor-guard-did`.

**Tests added:** none — by design; this pass changed no code.

**Suites (runner · section scope · language) — the session's whole-section reading, taken at the FINAL
tree (`rosetta-extensions` `91c926e`), after every pass-54/55 edit had landed:**

| suite | runner | section scope | language | result |
|---|---|---|---|---|
| **whole section, final tree** | pytest 8.4.2 / CPython 3.9.6 | **`stack-core` only** | Python | **1,945 passed · 0 failed · 3 skipped** (28 m 45 s) |
| flake gate, the 32 net-new tests | pytest 8.4.2 / CPython 3.9.6 | `stack-core` | Python | **32 passed ×3 consecutive** |
| `guard_family --platform`, correct root | `guard_family.py` | `stack-core` | Python | **24 GREEN · 0 RED · 0 could-not-check · 5 not-run** |
| `--ceilings` one-command form | `derivation_registry.py` | `stack-core` | Python | **3 ratchets exact, exit 0** |

**This is the first whole-section reading in this milestone's recent harden history with ZERO failures**
— pass 53 read **1,850 passed · 4 failed** at a frozen tree and explicitly did not claim a post-repair
re-read; the batch's own mid-session reading was **1,922 passed · 3 failed**, and all three of those were
this session's own regressions, each caught by a shipped repo-wide fence and repaired in pass 55.

**NOT COVERED, stated rather than implied (`§5` rule 60):** the four non-`stack-core` **Python** sections
were last read at pass 51 (**1,824 passed · 9 failed, all declared ENV_GATED**) and are **not** re-claimed
here; **no Go verdict** is claimed by any of passes 54–56; the **424 TypeScript tests** remain ENUMERATED
and never executed. The 3 skips are unchanged from the prior reading and are not new.

**Knowledge backfill:** one rule — *a reach number is not a verdict; the same census can return an
identical denominator at two refs and a different answer at each, so a claim that "the ref made no
difference" must compare verdicts, not reaches.*

**Stop condition:** cap reached without stabilization — three passes, **7 defects fixed inline** and
**2 routed**, and pass 56 found a corpus defect that the family's own default ref selection hides.

> **The 7, enumerated — because this entry first said SIX.** (1) `clone_pin_guard` arm C compared the ref
> literally; (2) arm B never required `platform`; (3) `patch_anchor_guard` answered from the local cache;
> (4) `guard_family`'s RED summary omitted its reach; (5) `TestCollectionParity`'s locator matched its own
> quoted literal; (6) `platform_predicate_guard` G2 graded the count and read as membership; (7) the
> repo-wide collection fence was blind to local fixture bases. The wrong figure was **carried** from a
> commit message rather than derived from the list — the ~1-in-3 rate this milestone has measured four
> times, reproduced once more by the pass that was auditing for it, and caught the way the ledger says to
> catch it: **enumerate where you print, never carry.** Per
the user's standing ruling this is routed and NOT met with new machinery; the **thirteenth**
cap-without-stabilization in this milestone (22, 25, 26, 29, 32, 35, 38, 41, 44, 47, 50, 53, 56).

## Pass 57 — 2026-08-10 — incremental

**Iters hardened this pass:** iter-229 … iter-238 (the ten tiks since pass 56).

**Tiks covered since prior pass:** 10.

**Scope, and the shape of it — stated first because it decided the whole pass.** Only **one** of the ten
iters landed executable code: **iter-229** (`rext` `da093f1`, +203 `buildbench.py` / +189 its tests). Iters
230–238 were **census tiks** whose instruments lived in the scratchpad and were never committed, so their
per-iter diff footprint in either repo is **documentation only** (8 corpus/CLAUDE.md files, 254 insertions).
A six-dimension sweep against nine doc-only diffs would have produced nothing; the pass went instead at the
one production surface and at the **substrate the ten censuses ran on** — which the batch's own §9 corrective
names as the thing to distrust first.

**Bugs surfaced + fixed inline: 3.** Enumerated here rather than counted, because this ledger has twice
recorded a *carried* figure that its own list refuted.

**(1) The identity gate reached the exit code and nothing else** (`rext` `dd5ba84`). iter-229's
`profile_describes_host()` is sound — 6/6 sealed predictions held and its three arms each catch a case a
cheaper arm passes. It was wired into `run_campaign` **only**: a mismatch exits with the reserved
`EXIT_HOST_IDENTITY`, and `BUILDBENCH_ALLOW_HOST_MISMATCH=1` lets the campaign measure while keeping the
exit code non-zero. The comment at the return statement states the intent exactly — *"a campaign run on a
host its profile does not describe must never exit 0, or the hatch converts an honest UNMEASURED into a
quotable green"* — and **that is the entire disclosure.** Measured: `build_report` never reads
`host_identity`; `campaign.json` records `"ok": true, "gateable": true` beside the p50; `print_report` prints
a clean summary; and **`buildbench.py report <dir>` re-aggregates the same rep ledgers in a different process
that never sees the run's exit code**, returning 0.

> **The corpus made it worse rather than better, and that is why this counts as a platform-alignment
> defect rather than a tooling nit.** `corpus/ops/demo/build-budget.md` instructs the reader, in bold,
> **"Read `gateable`, not the exit code, before quoting any number."** The rule is *correct* about `--reps
> 1` — and it routed a reader past the one channel that carried the host gap. Following the documented rule
> exactly produced a quoted p50 measured on a machine the profile does not describe.

Repaired with `_identity_rollup()`: every rep's verdict folded **worst-first** into `report["host_identity"]`
— `mismatch` → RED (any single rep is enough: a p50 aggregated across two machines measures neither);
`unmeasured` and `absent` → not RED but **not `gateable`**, with the reason stated; `match` → still gateable.
**`absent` is a verdict, not a default**: defaulting a missing field to `match` would have awarded a
gate-quality green to every campaign directory written before iter-229 — including the ones this release's
baseline numbers came from.

**(1b) The enabler, which is the more useful finding.** The test fixture `_ledger` carries the docstring
*"a rep ledger in the shape `run_campaign` really writes"* and **omitted `host_identity` entirely**, so every
report test in the file was asserting against a ledger shape the harness does not produce. That is how a
field could be written into every rep ledger by one function and read by no other for a whole iter. Fixed at
the fixture, which is what made the new control arm (`match` is *still* gateable) meaningful instead of
vacuous.

**(2) The same rext file resolved to two different checkouts depending on how it was spelled**
(`rext` `dd5ba84`). `anchor_construct_guard.resolve()` walks the clone roots before falling back to the
authoring rext copy. `stack-demo/rosetta-extensions` **is a real directory**, so a citation headed
`rosetta-extensions/…` matched the loop's `(root / head).is_dir()` test and returned the **tag-pinned
per-stack clone**; the fallback at the bottom was unreachable for that spelling. `tracked_basenames`
excludes that clone on the stated grounds that *"the authoring copy is the current one, and `resolve()`'s
existing rext fallback already prefers it"* — **false for one of the three spellings, and the docstring is
where the belief was recorded.**

| spelling | resolved to (measured, before) |
|---|---|
| `rosetta-extensions/stack-core/buildbench.py` | `stack-demo/rosetta-extensions/…` — **pinned `09d0607`** |
| `stack-core/buildbench.py` | `.agentspace/rosetta-extensions/…` — authoring `da093f1` |
| `.agentspace/rosetta-extensions/stack-core/buildbench.py` | `.agentspace/rosetta-extensions/…` — authoring |

The two checkouts are **162 commits apart**, with **66 of 1,106** common tracked files differing (both
measured this pass). So this was not one guard with a spelling quirk — it was **two guards**, and nothing in
the output said which one had answered.

**(3) A carry-forward glob that reads as live backlog** (`9bc8aef`). `route_disposition_guard` was **RED at
`907cbc3`**, the tree this pass inherited: iter-238 carried `ROUTE-M257x-236-*` / `ROUTE-M257x-235-*` as
wildcards, and `§5` rule 73 refuses them — a glob leaves a truncated id stem behind, and that stem reads as
an open route in every brief quoting the queue. Enumerated the four real ids. **The first repair attempt
stayed RED**, because the explanatory note quoted the offending spellings — the `retracted_pin_guard` class
reproduced one document over, by the repair for a sibling rule. Guard: `EXIT 1 → OK`, malformed **2 → 0**.

**Coverage delta on touched files:** not computed as a percentage this pass, and stated rather than implied.
The in-scope production surface is a single file (`buildbench.py`) whose new arms were reached by 8 net-new
tests written to fail first (**8/8 RED before the fix, 8/8 GREEN after**) — a line-coverage percentage over a
1,400-line module would move ~1 % and describe nothing. The measurement that mattered was the *reach* one:
which artifacts carry the verdict (4 checked: `campaign.json`, `print_report` stdout, the `report`
subcommand's exit code, the rep ledgers — **1 of 4 carried it before, 4 of 4 after**).

**Tests added: 12.**
- iter-229 → `stack-core/tests/test_buildbench.py`: **8** (`TestIdentityReachesTheQUOTABLEArtifact`) — RED
  before the fix, including the control arm that asserts a matched campaign is **still** gateable.
- substrate → `stack-core/tests/test_anchor_construct_denominator.py`: **4**
  (`TestRextCitationsResolveToOneCheckout`) — three-spelling convergence, a fixture-non-vacuity control, an
  unresolvable-stays-unresolvable arm, and a scope control proving non-rext citations are untouched.

**Knowledge backfill:** `corpus/ops/demo/build-budget.md` — the `gateable` second clause, the four-verdict
table, and why `absent` is a verdict. The edit also **dropped three `buildbench.py:NN` line-anchors** in
favour of the symbol names they cite: this pass's own edit had just moved them, and re-deriving a line
anchor only re-arms the rot iter-234 measured (5 of 5 hand-checked "mismatches" were the corpus being right
and the instrument wrong).

**Flakes stabilized:** none surfaced.

**Stop condition:** continue-to-next-pass — three defects in the first dimension scan and no coverage delta
measured across passes yet; the census substrate is the named next target.

## Pass 58 — 2026-08-10 — incremental

**Iters hardened this pass:** iter-229 … iter-238 (same batch; this pass went at the **substrate** and at
the runnable surface the batch repaired).

**Tiks covered since prior pass:** 10 (unchanged — pass 57 and 58 are two passes over one batch).

**Bugs surfaced + fixed inline: 2.** Enumerated, not counted.

**(1) `CLAUDE.md`'s "Go Services" section described `app` and named two repos it does not** (`bfc4bd9`).
The heading read *"Go Services (Backend, CMS, Sentinel, etc.)"* and offered one *"common development
pattern"*. Measured against the **two** Go repos a current stack builds (`stack-demo/{app,sentinel}`,
`platform` at `0c91421` = `origin/main`):

| documented step | `app` | `sentinel` |
|---|---|---|
| `make setup` | ✅ — and it installs **five** tools, not the three the doc listed | ❌ **no such target** (`initdb`, `proto`) |
| `make gen` | ✅ `go generate ./...` | ❌ **no such target** — it is `make proto` |
| `atlas migrate apply --env local` | ✅ `atlas.hcl` declares `env "local"` | ❌ **no `atlas.hcl` at all** |
| the 5 listed "key directories in Go services" | **5 of 5** | **0 of 5** |

`repos.yml` independently agrees on the last row — `sentinel` carries `migrations: false`. And the heading
named **`cms`**, a repo this milestone established is decommissioned. A developer following the section
into `sentinel` collects two *"No rule to make target"* errors and then hunts for an `rpc.go` that has
never existed there. **This is `ROUTE-M257x-238-claude-md-fences-are-unmaintained` firing again** — the
prose one screen above is current about the merges and the fence below it is four releases old, which is
exactly what iter-238 predicted the next defect would be.

**(2) `guard_family`: "the clone has NO origin/main" was three states wearing one sentence, and the remedy
it printed is impossible for two of them** (`rext` `8fbf5b9`). The refusal is correct and stays — without
`origin/main` a guard asked to read the ref the exit gate names silently reads the checkout. The *advice*
was *"Fetch the clone"*, universally. Measured on this box:

- **`absent`** — `--platform stack-dev/platform` names a directory that **does not exist** (the platform
  clone lives under `stack-demo/`). The runner printed `@ ? (origin/main ABSENT, DIVERGED)` for a path with
  no repository at all and then told the operator to fetch it.
- **`bundle`** — `stack-dev/studio-desk`, the clone **iter-233 measured**: a correct-looking GitHub fetch
  URL on `origin` and **not one ref under `refs/remotes/origin/`**; all 281 refs came from the bundle
  remote, so `git fetch origin` cannot create what is missing without the real network remote.
- **`unfetched`** — the only state for which the original advice was ever right.

Classified into five kinds; the print site keys a table with **no default branch**, so an unclassified
state `KeyError`s rather than inheriting the fetch advice that caused this.

> **The rule the first attempt got wrong, recorded because it is the generalisable part:** *the presence of
> an `origin` **remote** is not evidence the clone came from origin.* The first classifier read
> `git remote` and graded the real bundle clone `unfetched`. It was caught by **exercising the new arm
> against the actual clone rather than only the fixture** — the fixture would have passed either way.

**Measured and NOT changed, because the corpus was right — stated rather than omitted (`§5` rule 60).**
Three of this pass's first-look "findings" were the instrument, which is the batch's own §9 caution
reproduced a third time in one session:

| checked | reading | verdict |
|---|---|---|
| all 7 `CLAUDE.md` → `docker-compose.yml:NN` anchors | **7/7** resolve to the text claimed | corpus right |
| the profile table vs compose | correct **once `include: common.yml` is read** — `postgresql`/`redis` live there, both with no `profiles:` key. A first parse of `docker-compose.yml` alone said "5 services, no database" | **instrument wrong** |
| `make setup`/`make gen` | present in `app/Makefile`; absent from `platform/Makefile` — which is not where the doc says to run them | **instrument wrong** (the real defect was `sentinel`, found second) |
| `make force-gen` (`shared_libraries.md`) | belongs to the `proto` repo, which is not in the clone set | **unmeasurable**, not wrong |
| `repos.yml` vs the 4 repos `CLAUDE.md` names | exact match | corpus right |
| `app/.gitignore:78-79` (iter-236's studio cite) | lands on the studio-ignore comment + `studio/*` | corpus right |
| `stack-demo/app/studio` populated (18 entries) | iter-236's repaired `cd` target is real | corpus right |

**Coverage delta on touched files:** the two touched modules gained arms that were **unreachable before**
(`_ref_state`'s classifier did not exist; `CLAUDE.md` has no executable coverage). Reach measured instead:
`_ref_state`'s missing-`origin/main` path had **1** outcome and now has **5**, each with a distinct remedy
and each exercised — plus a non-vacuity arm asserting a healthy clone carries **no** kind, and a control
keeping `unfetched` saying "fetch".

**Tests added: 7** — `stack-core/tests/test_guard_family.py::TestMissingOriginMainIsCLASSIFIEDNotJustReported`.

**Knowledge backfill:** `CLAUDE.md` § Working with Service Code — replaced the one-pattern block with a
per-repo table + a runnable block per repo + a pointer to `corpus/services/sentinel.md`.

**Flakes stabilized:** none surfaced.

**Stop condition:** continue-to-next-pass — two more defects, both in the batch's own repaired surface, and
no cross-pass coverage delta measured yet; pass 59 takes the whole-section reading at the final tree.

## Pass 59 — 2026-08-10 — incremental

**Iters hardened this pass:** iter-229 … iter-238 (third and final pass over the batch).

**Tiks covered since prior pass:** 10 (one batch, three passes).

**Bugs surfaced + fixed inline: 2.** Enumerated. The first is the session's most consequential finding and
the second is the session's own defect, found by the corrective this batch wrote.

**(1) A third of the corpus's qualified anchors were not *unresolvable* — they were UNSEEN**
(`rext` `cb8bcb3`). `anchor_construct_guard._QUALIFIED` requires a closing backtick immediately after the
line number, so a **range** citation — `` `foo.py:1396-1433` `` — matches nothing the guard scans. Range
anchors appear in **neither** `resolved` **nor** `unresolvable`: the reach ratio is computed over a
denominator that silently excludes them, and the guard's green reads as coverage it does not have.

Measured on the fenced surface (`corpus/**` + `CLAUDE.md` + `README.md` + `.claude/skills/**`, 114 docs):

| population | count |
|---|---|
| single-line `` `path:NN` `` citations — what `_QUALIFIED` matches | **975** |
| **RANGE `` `path:NN-MM` `` citations — invisible to it** | **491** across **67** files |
| range share of the qualified population | **33.5 %** (the guard prints **24.9 %**, against the wider `denom + total` it can see — both stated with their denominators) |

The top file is **`corpus/architecture/platform-migration-status.md` (36)** — the map this corpus calls
authoritative for per-service claims — then `org-repos.md` (29), `ai_architecture.md` (21),
`backend.md` (21), `safety.md` (19).

> **This is the structural explanation for iter-229's "22 anchors rotted, the instruments caught 3."**
> The batch read that as a tuning gap. It is a population gap: the guard's floor already discloses that it
> cannot detect *"resolves to the WRONG construct"*, and this adds the sharper one — a third of the
> anchors are not in its subject at all.

**Grading them is ROUTED, not done** — which line of a range carries the claim is a design decision with a
491-anchor blast radius. `ROUTE-M257x-h59-range-anchors-are-ungraded`. **Counting them is free**, and an
unmeasured population that is STATED is a different object from a silent one: the disclosure prints on
every run **including when the count is zero**, names the route, lists the top five files, and rides in
`--json` beside `reach` so a machine consumer cannot take the partial denominator for the whole.

**(2) This session rotted a corpus anchor and shipped it past three green fences** (`735c8ea`).
`build-budget.md`'s `buildbench.py:1396-1433` was correct when written. Pass 57's edits to that module
moved it: `1396` now lands on `_reclaim_attribution`, and the argparse the sentence claims to verify
against begins at `1464`. **Nothing caught it** — `anchor_construct_guard` read GREEN (finding 1 is why),
the pre-commit hook passed, and the out-of-range check *cannot* fire on an insertion, which only grows the
file. It surfaced because the anchor re-derivation was **deliberately deferred to after the session's last
edit** to those modules rather than run when the first one landed.

> **§9's corrective, working exactly as written, on the pass that was auditing for it.** *"Re-derive after
> the LAST edit, not the first."* The batch's own caution earned its place. Repaired to the **single-line**
> form `buildbench.py:1464` rather than to the corrected range — restoring a correct number into a slot
> finding (1) proves is ungraded would have been the weaker fix.

**Coverage delta on touched files:** the cross-pass reading the stop condition asks for. Passes 57 → 58 →
59 surfaced **3 → 2 → 2** inline-fixable defects; the dimension scan has **not** gone quiet, so the delta
condition is not met and this is a cap, not a stabilization.

**Tests added: 5** — `test_anchor_construct_denominator.py::TestRangeCitationsAreCOUNTEDEvenThoughUngraded`,
including a **partition control** (the two patterns must not both match one citation — so the counter
cannot become a double-count if `_QUALIFIED` ever learns ranges), a **zero-count** arm, and an arm pinning
the share's denominator.

**Suites (runner · section scope · language) — taken at the FINAL tree** (`rosetta` `735c8ea`,
`rext` `cb8bcb3`, pushed to origin), after the last edit of the session:

| suite | runner | section scope | language | result |
|---|---|---|---|---|
| flake gate, the **24 net-new** tests | pytest 8.4.2 / CPython 3.9.6 | `stack-core` | Python | **24 passed ×3 consecutive** |
| `guard_family --platform stack-demo/platform` | `guard_family.py` | `stack-core` | Python | **24 GREEN · 0 RED · 0 could-not-check · 5 not-run** |
| whole section, **first** reading at the final tree | pytest 8.4.2 / CPython 3.9.6 | **`stack-core` only** | Python | **1,977 passed · 8 FAILED · 3 skipped** (24 m 48 s) |
| whole section, **after the repair below** | pytest 8.4.2 / CPython 3.9.6 | **`stack-core` only** | Python | **1,985 passed · 0 failed · 3 skipped** (24 m 50 s) |
| `derivation_registry.py --ceilings` | `derivation_registry.py` | `stack-core` | Python | **3 ratchets exact +0, exit 0** |

> **THE 8 WERE THIS SESSION'S OWN, AND THE FENCE CAUGHT ALL OF THEM.** Every failure was in
> `test_frozen_expectation_census_m257x.py` — the milestone's frozen-expectation census — and every one
> was a literal ratchet breached by passes 57–59's own prose. Repaired, not accepted; **`1,977 + 8 =
> 1,985`**, so the re-read is the same population with the 8 flipped and nothing lost.
>
> **Repaired in the order the ratchets require, because the first step moves the population the other
> three bound.** The residual arm went first: `refs?` added to `_MEASURED_NOUNS`, because pass 58's
> `guard_family` fix wrote *"all 281 refs under the bundle"* into an advice string and `refs` is
> unambiguously a count this repo had never written after a number. **Seventh consecutive time the
> vocabulary's reach has closed on the sentence that widened it**, and the seventh time the residual arm
> surfaced it rather than anyone reading the list.
>
> Then the three ceilings, each isolated **the iter-223 way — tree held fixed, matcher varied** — so each
> arrow states which half is prose and which is reach, which is what its own message means by *"never a
> blind bump"*:
>
> | ratchet | arrow | prose | reach |
> |---|---|---|---|
> | `DOCSTRING_LITERAL_CEILING` | **219 → 221** | +1 — `_print_range_disclosure` / `_rext_authoring` docstrings state the measurements that motivate them | +1 — the `refs?` widening (pre-widening matcher on this tree returns 220, post 221) |
> | `COMMENT_LITERAL_CEILING` | **190 → 196** | +6 — the counting-site block, then **the recorded reasons themselves** | 0 (pre == post both times) |
> | `TEST_MODULE_LITERAL_CEILING` | **597 → 601** | +4 — the measurements the new tests exist to pin | 0 (pre == post) |
>
> **The comment ceiling's +5 second arrow is the fixpoint this ceiling's own top paragraph describes,
> hit for the THIRD consecutive harden session**: re-pinning a ratchet demands a recorded reason, a
> recorded reason is made of figures, and those figures live in `#:` comments — which is that ratchet's
> population. It converged in one iteration only because the matcher requires *number + measurement
> noun*, and an arrow like `191 → 196` is followed by no noun. `derivation_registry.py` is now the
> heaviest contributor at **63** for exactly this reason and will be for every future re-pinning pass.
>
> **Grade this as a finding, not as friction.** The section reading was taken **once, after the last
> edit** — the §9 corrective — and it is the only thing in this session that caught these 8. Three green
> per-file runs, three green `guard_family` runs, a green pre-commit hook and a passing flake gate all
> preceded it and none of them saw the breach: **the ratchets are repo-wide, and nothing scoped to the
> files a pass touched can read them.**

**NOT COVERED, stated rather than implied (`§5` rule 60):** the four non-`stack-core` **Python** sections
were last read at pass 51 (1,824 passed · 9 failed, all declared ENV_GATED) and are **not** re-claimed
here; **no Go verdict** is claimed by passes 57–59; the **424 TypeScript tests** remain enumerated and
never executed. The `guard_family` line above is **not a whole-family green** — it says so itself, and 5
members had no input supplied.

**Knowledge backfill:** two rules, both generalisable beyond their sites.
1. *The presence of an `origin` **remote** is not evidence a clone came from origin.* Only the refs say who
   populated it — the measured bundle clone carries a correct GitHub fetch URL and zero refs under it.
2. *A citation shape the scanner's regex cannot match is not a low-confidence citation; it is an absent
   one.* Reach must be reported over the population, and the shapes excluded by **construction** are the
   ones no reach ratio will ever reveal.

**Routes carried forward from this session:**
- `ROUTE-M257x-h59-range-anchors-are-ungraded` → **new.** 491 range anchors are counted and not graded;
  deciding which line of a range carries the claim is the design call.
- `ROUTE-M257x-h59-rext-edits-fire-no-fence-anywhere` → **new.** `.agentspace/` is git-ignored by `rosetta`
  and `rext` has no hooks and no CI, so an edit to a rext module can never appear in `git diff --cached` on
  the corpus side. The pre-commit runner is gated on staged `corpus/`, `.claude/skills/`, `CLAUDE.md`,
  `README.md` paths. A rext edit that rots corpus anchors fires **nothing, anywhere, by construction** —
  which is the enabler behind finding (2) and behind iter-229's 19 silent.

**Flakes stabilized:** none surfaced; the 24 net-new tests passed 3×3 consecutive. **And the flake gate is
exactly what did NOT catch the 8 ratchet breaches** — it is scoped to the new tests by construction, so it
is structurally incapable of reading a repo-wide ceiling. Recorded because a green flake gate beside a RED
section is otherwise read as a contradiction rather than as two instruments with different reaches.

**Stop condition:** cap reached without stabilization — three passes, **7 defects fixed inline** and
**2 routed**, and the third pass found the largest reach gap of the session rather than going quiet.

> **The 7, enumerated where they are printed rather than carried** (`§5`, and the rule this ledger has
> broken twice — **it very nearly broke it a third time here**, see the note under the list):
>
> | # | pass | defect |
> |---|---|---|
> | 1 | 57 | the identity gate disclosed **only** through the exit code, so `campaign.json`, `print_report` and the offline `report` subcommand all read green on a mismatched host — **and `_ledger`, the fixture whose docstring promises the real ledger shape, omitted `host_identity`**, which is how the field went a whole iter unread |
> | 2 | 57 | rext anchors resolved to two checkouts **162 commits apart** depending on how the citation was spelled |
> | 3 | 57 | iter-238's route glob left a truncated id stem reading as live backlog; `route_disposition_guard` was **RED on the inherited tree** |
> | 4 | 58 | `CLAUDE.md`'s "Go Services" block described `app` while naming `sentinel` (**0 of 5** directories, **0 of 2** make targets, no `atlas.hcl`) and decommissioned `cms` |
> | 5 | 58 | `guard_family` told two **unfetchable** clone states to fetch |
> | 6 | 59 | **491** range anchors sat outside the reach denominator that reported on them |
> | 7 | 59 | this session's **own** rotted anchor — `build-budget.md` → `buildbench.py:1396-1433`, moved by pass 57 |
>
> **The near-miss, recorded because it is the point.** The first draft of this list ran to seven items by
> counting the `_ledger` fixture as its own row and **dropping row 7 entirely** — the session's own defect,
> the one whose whole lesson is that a repair pass rots what it does not re-derive. It still summed to 7,
> so the arithmetic would not have caught it; only reading the list against the per-pass counts (**3 + 2 +
> 2**) did. The fixture is folded into row 1 because that is how pass 57 counted it. *The enumeration is
> the source and the total is the derivative* — a total that agrees with a wrong list is the failure mode.
>
> Per the user's standing ruling the two routes are recorded and NOT met with new machinery; the
> **fourteenth** cap-without-stabilization in this milestone (22, 25, 26, 29, 32, 35, 38, 41, 44, 47, 50,
> 53, 56, 59).

## Pass 60 — 2026-08-10 — incremental

**Iters hardened this pass:** iter-239 … iter-248 (first pass over the batch).

**Tiks covered since prior pass:** 10.

**Scope manifest.** 20 `rosetta-extensions` files (5 net-new guards — `skill_invocation_guard`,
`toolchain_floor_guard`, `rext_path_guard`, `fence_command_guard`, `env_absence_guard`; 3 modified —
`platform_alignment_guard`, `guard_family`, `anchor_construct_guard`; 1 Go working-stack fix —
`stack-seeding/cmd/stackseed/main.go`) + 10 corpus files, across 30 commits (`907cbc3..8a9b321`).
`rext` `17e0795..7d75c52`, 3,324 insertions. No iter declares an `iter_shape`; the batch reads as 9
tooling-iters + 1 production fix, so dimensions 1 (test depth) and 5 (input fuzzing) carry the weight.

**Bugs surfaced + fixed inline: 2.** Both are the same defect in two iters, and it is the defect this
milestone keeps re-finding: **a fence whose denominator is drawn around the spelling that happened to be
in front of the iter that wrote it.**

**(1) `skill_invocation_guard` graded FENCED lines; the corpus writes invocations in PROSE**
(`rext` `cca7938`, corpus `c61c2c5`). iter-239's headline was **8 of 8 target-bearing invocations
wrong**. It repaired those 8 and shipped a fence scoped to fenced blocks. Measured on the tree it
shipped: **14 inline backticked invocations name a target for one of the 3 slot skills, and 12 disagree
with that skill's own contract** — 10 a bare `N`/`1`, 2 inverting the order — while the fenced arm read
a confident green. *The repair went to the instances, not to the class* — which is the pathology
`fence_command_guard`'s own docstring names one file over, committed six iters later.

> **The control that settles the obvious objection.** *Is a bare `N` inline just a placeholder
> convention?* No: **2 inline sites already spelled it qualified** — `.claude/skills/dev-up/reference.md:146`
> (`/stack-snapshot dev-N …`) and `corpus/ops/demo/tailscale-serve.md:252` (`/stack-secrets demo-1 …`).
> The corpus is inconsistent inline, not conventional.

Three of the 12 are worth naming rather than counting:

| site | what it is |
|---|---|
| `.claude/skills/stack-snapshot/SKILL.md:32` | **the contract document.** The skill's own flow line — the one iter-239 *quoted as evidence* for target-then-verb — while its target was spelled `N`. The line carries four invocations and a bare `N` is **right for two** (`/dev-up N`, `/demo-up N` declare `[N]`) and **wrong for two**. That is iter-239's own `P-239-5` refutation, sitting in the document that declares the split |
| `corpus/ops/demo/recipe-enterprise-onboarding.md:73` | a **literal `/stack-seed 1`**, not a placeholder. `dev-1` and `demo-1` are both real stacks. iter-239 repaired the identical shape in the sibling `recipe-skill-progression.md` and missed this one |
| `corpus/ops/demo/recipe-snapshot-world.md:38` | the document whose `:28` states replay works on `dev-N\|demo-N`. iter-239 **recorded this document's `:28`-vs-`:44` self-contradiction** and repaired `:44` only |

Arm D now grades both populations, tagged by surface (`(inline, in prose)`). **Arm A deliberately does
NOT reach inline text** and the refusal is tested: `` `/profile` ``, `` `/home` ``,
`` `/sim/<slug>/result/<sessionId>` `` are backticked slash-tokens and none is an invocation — that is
the **152-URL-path instrument** the guard's own denominator section measures. Arm D cannot reach them
because it never asks *"is this a skill?"*, only *"given that this IS one of the 3 slot skills, how is
its target spelled?"*.

**(2) `rext_path_guard` excluded the most EXPLICIT spelling of a rext path** — the one that names the
repo outright: `.agentspace/rosetta-extensions/…`, `stack-dev/rosetta-extensions/…`. iter-244's
tail-match boundary (*"a path preceded by any path segment is not a reference to rext"*) was written
against `app/knowledge/…` and swept this up by accident. Measured: **27 occurrences of 25 distinct
paths, 16 of them in NO other spelling** — including two operator-facing
`stack-dev/rosetta-extensions/dev-stack/migrate-dev.sh` sites, in `setup_guide.md` and the `/dev-up`
skill. Reach **301 → 329** occurrences, **145 → 161** distinct. **All 16 resolve**, so the gap hid no
live defect on this tree — it would have hidden the next one silently. The allowance is keyed on the
**literal `rosetta-extensions` token**, never on "any leading segment": `some-repo/stack-core/nope.py`
and `.claude/skills/stack-secrets/SKILL.md` stay excluded and a test pins that, because widening it
further would re-open the class iter-244 measured at **23 false findings**.

**Tests added: 8** (`stack-core`, Python/unittest-under-pytest) — 5 on `skill_invocation_guard` (inline
bare-`N`, inline verb-first, inline qualified accepted, fenced-is-not-double-counted, and the arm-A
refusal over URL paths), 3 on `rext_path_guard` (prefixed-path RED, prefixed-path reverts GREEN, and
the token-keyed boundary preservation). The `test_rext_path_guard.py` file goes 17 → 20 and
`test_skill_invocation_guard.py` 17 → 22; both files green, **42 passed** together.

**Mutation control on this pass's OWN new arms** (`§5`, and pass 54's caution that a new arm can be a
self-matching identity): each new test was re-run against the **pre-fix guards at `7d75c52`** in a temp
checkout. **4 of 5** inline tests and **2 of 3** reach tests go RED there and GREEN here. The 2 that
pass on both are the two that assert an **exclusion** both versions honour — recorded as passing-on-both
rather than counted as controls.

**Answer key repartitioned rather than re-baselined.** iter-239's `test_answer_key_…_eight_findings`
reconstructs the whole live tree at `2a0a939` and asserted a flat 8. The widened arm reports **15**
there. The assertion is now **partitioned by surface** — 8 fenced (5 bare + 3 verb, iter-239's published
number, **unchanged**) and 7 inline (5 bare + 2 verb). **7, not 12**, and the difference is a property
of that reconstruction, not a disagreement: the test overwrites `.claude/skills/*/SKILL.md` from the
LIVE tree, so the 5 inline defects in `demo-up/SKILL.md` (3) and `stack-snapshot/SKILL.md` (2) arrive
already repaired. **7 + 5 = 12**, and the test says so in place rather than leaving two numbers in one
ledger to look like a contradiction.

**Knowledge backfill:** one rule, stated where it is enforced (both guard docstrings) rather than only
here. *A fence inherits the reach of the iter that wrote it.* An iter censuses the surface it can see,
repairs what it finds, and then draws the fence around **that same selector** — so the fence is green by
construction over exactly the population already repaired, and blind to the rest of the class by the
same construction. Both of this pass's findings have that shape, three iters apart, and neither guard
was wrong about anything it graded. **Grade a new fence on its DENOMINATOR before its verdict.**

**Flakes stabilized:** none surfaced.

**Stop condition:** continue-to-next-pass — the dimension scan covered 2 of the 5 net-new guards; the 3
modified guards (`platform_alignment_guard`, `guard_family`, `anchor_construct_guard`), the remaining
new ones (`toolchain_floor_guard`, `fence_command_guard`, `env_absence_guard`) and the batch's single
Go production fix (`--reload-sentinel`, iter-243) are unscanned.

## Pass 61 — 2026-08-10 — incremental

**Iters hardened this pass:** iter-239 … iter-248 (second pass over the batch).

**Tiks covered since prior pass:** 10 (one batch, second of the session's passes).

**Bugs surfaced + fixed inline: 7.** Enumerated below, and the enumeration is the source — the total is
its derivative (the rule this ledger has broken twice and near-missed a third time). **All seven
pre-date this harden pass**; none was introduced by pass 60.

> **The finding that frames the other six: pass 59 predicted this in the words it is now recorded in.**
> *"The ratchets are repo-wide, and nothing scoped to the files a pass touched can read them."* It routed
> the cause as `ROUTE-M257x-h59-rext-edits-fire-no-fence-anywhere` — `.agentspace/` is git-ignored by
> rosetta and `rext` has no hooks and no CI, so a rext edit that rots a repo-wide ratchet **fires
> nothing, anywhere, by construction**. Ten iters later the prediction had come true five times over.
> **A whole-section run over the tree iters 239–248 shipped reports `11 failed / 2,073 passed / 3
> skipped`** (`stack-core`, pytest 8.4.2 / CPython 3.9.6, Python — 27 m 43 s), and every one of the 11
> was RED at HEAD while each of those ten iters closed on a green per-file run **and** a green
> `guard_family` line, neither of which reaches those files. iter-248's own scope disclaimer — *"this
> runs the GUARDS, not their test suite"* — is the correct diagnosis, written by the batch that was
> demonstrating it.

| # | what | why it survived |
|---|---|---|
| 1 | **NINE unregistered derivations** across all five new guards — `env_absence_guard::live_markdown`, `fence_command_guard::{live_markdown,make_targets,package_scripts}`, `rext_path_guard::{live_markdown,sections}`, `skill_invocation_guard::live_markdown`, `toolchain_floor_guard::{required_go,required_node}` | the classification arm is repo-wide; **RED at HEAD for all ten iters** |
| 2 | **`DOCSTRING_LITERAL_CEILING` 221 → 234** | same |
| 3 | **`COMMENT_LITERAL_CEILING` 196 → 213** | same |
| 4 | **`TEST_MODULE_LITERAL_CEILING` 601 → 622** | same |
| 5 | **three undecided measurement nouns** — `variables`, `skills`, `series` | same |
| 6 | **`basename_index`'s docstring rotted in BOTH grains** — `293 of 706` / `410 of 981` stated, `291 of 706` / `407 of 981` live | see the note below |
| 7 | **`sentinelContainerStack` guessed `demo-N` for any unrecognised family** + the doc comment the same insertion orphaned | inside the batch's only production fix; the gate was tested, the *general case* was not |

**(1) is a RE-FRAGMENTATION, not an omission, and that is the part worth reading.** Four of the nine are
`live_markdown`. **iter-212 DELETED four per-guard tree-scan declines** and consolidated them into ONE
shared derivation (`fence_provenance::corpus_sources`) under its own rule — *two readers of one construct
must SHARE the derivation*. Iters **239, 244, 246 and 247** then each wrote a private `live_markdown`
with **byte-identical** `LIVE_PREFIXES`/`LIVE_FILES` — the exact shape iter-212 removed, re-grown one
guard at a time — while the registry that exists to notice it was RED and therefore silent. All four
report the same **114** documents today. All nine are now registered; **no `RELATION:` clause**, for the
reason `claude_docs_outside_skills` already records (the grammar's `_resolve_operand` hands each operand
the GUARD DIR, so both would resolve empty against a repo root they never receive). **The consolidation
itself is a five-module refactor and is ROUTED, not smuggled into a harden pass.**

**(6) is a test that hid what it had not yet checked.** `test_the_docstring_carries_the_numbers_this_tree
_derives` asserts the *distinct* grain first and the *pair* grain second. The distinct grain failed, so
the pair assertion never ran — and it was **also** stale. Both are now current and the docstring says so
in place. Recorded because "1 failing assertion" and "1 stale figure" are not the same claim.

**(7) is the same lesson learned twice in one batch, three iters apart.** iter-243 fixed
`--reload-sentinel` restarting `demo-N-sentinel-1` on every `dev-N` stack — then collapsed *"carries no
family"* and *"carries a family I do not recognise"* into a single `demo-N` default, which leaves the
**general form of the bug it just fixed**: `--stack devv-2` (a typo) or `--stack stage-2` (CLAUDE.md
declares `stack-stage/` and `stack-tests/` as future members of this family) restarts
`demo-2-sentinel-1` — **a container belonging to a different stack**. Nothing upstream rejects those
names: `ParseStackN` accepts every one of them. Now **three** cases — known family → the name; **no**
family (`""`, `anthropos`, a bare offset) → the historical `demo-N`, **byte-identical**; unrecognised
family → `""` and the caller skips the restart **loudly**. This is precisely the
pardon-that-travels-with-the-name that **iter-246's `fence_command_guard::workspace_substitute` already
refuses**, in the same batch, three iters later.

**Plus the reach defect this pass's scan was looking for, the THIRD of the session** (the first two are
pass 60's): **`toolchain_floor_guard` graded `corpus/ops/setup_guide.md` and nothing else.** The corpus
states the same floors a **fourth** time in **`corpus/tools/toolchain_overview.md`** — the *Development
tools registry*, the document whose entire job is to say what you must install — where they read
**`Go (v1.23+)`** (the identical defect iter-240 repaired, still live) and **`Node.js (v20+)`** (four
major versions low; a v20 host cannot install `next-web-app`, which declares `">=24.0.0"`). Both
repaired, with the derivation stated in place; the Go bullet also named `cms` and `jobsimulation` as
service code you compile locally — both merged into `app`, neither repo in the clone set — corrected in
the same line rather than left as a second wrong claim inside a repaired sentence.

> **The subject was widened to a NAMED PAIR, deliberately not to a sweep.** Measured: **13** further
> live-corpus sites match the floor patterns and **11 are correct**. The two classes a sweep would
> falsely redden are stated in the guard: **per-app floors** (`frontend-tier.md:149` — *"ant-academy
> needs Node ≥ 22"*, true of ant-academy, which the global floor of 24 already covers) and **narrative
> quotes** (`frontend-tier.md:155`, `tailscale-serve.md:168` — a document reporting that ANOTHER
> document claims Go 1.26). *A fence that cannot tell a claim from a quotation of a claim manufactures
> findings.* The guard now reads **4** Go floors and **3** Node floors across the two documents and
> names them in its verdict.

**Tests added: 3 Python + 4 Go table cases.** `test_toolchain_floor_guard.py` **14 → 17** (the registry
graded ALONE against a correct guide, so a green guide cannot mask it; a missing subject document failing
CLOSED; and an answer key running the REAL pre-repair registry from `8a9b321`). `main_test.go`'s
`TestSentinelContainerStack` table **5 → 9** cases (`anthropos`, surrounding whitespace, and the two
unrecognised-family refusals). Both new Python tests were **mutation-controlled** against the pre-fix
guard at `7d75c52` and go RED there. The three-file total for the session is **59 passed**
(`test_skill_invocation_guard` 17 → 22, `test_rext_path_guard` 17 → 20, `test_toolchain_floor_guard`
14 → 17).

**The ceiling fixpoint fired for the FOURTH consecutive harden session, exactly as its own prose
predicted** — re-pinning a ceiling needs a recorded reason, a recorded reason is made of figures, and
those figures live in `#:` comments, which is the comment ceiling's own population. It converged in one
step here because `derivation_registry.py --ceilings` (the one-command form pass 59 built for this) was
read *after* the reason prose was written rather than before. **`TheCeilingProseDoesNotContradictTheCeiling`
still caught a defect in this pass's own repair:** all three re-pin blocks opened *"196 → re-pinned at
harden pass 61"* with no arrow TARGET, and the arm parses the last arrow and compares it to the constant.
A recorded reason that does not state the number it justifies is not a recorded reason.

**Knowledge backfill:** the pass-60 rule, now with its counterexample. *A fence inherits the reach of the
iter that wrote it* — three instances this session (`skill_invocation_guard`, `rext_path_guard`,
`toolchain_floor_guard`). The counterexample is `fence_command_guard`: its fenced-block scope is
**correct**, because a `make` target needs a directory context and prose does not establish one. **The
test is not "is the scope narrow" but "is the narrowness a property of the CLAIM or of the iter".**

**Flakes stabilized:** none surfaced.

**Stop condition:** continue-to-next-pass — `guard_family` and `anchor_construct_guard` (iters 242, 245,
248) have not been dimension-scanned, `fence_command_guard` and `env_absence_guard` were scanned but not
fuzzed, and the whole-section re-run proving the 11 closed is still in flight.

## Pass 62 — 2026-08-10 — incremental

**Iters hardened this pass:** iter-239 … iter-248 (third and final pass over the batch).

**Tiks covered since prior pass:** 10 (one batch, three passes).

**Bugs surfaced + fixed inline: 3.** Enumerated; the total is the derivative of the list.

**(1) `anchor_construct_guard`'s range accumulators never reset — while their two siblings, in the same
function, do** (`rext` `1bf6fc6`). `run()` clears `RESOLVE_ROUTES` and `BARE_REFUSALS` at its top. It did
not clear `RANGE_CITATIONS` (added by **harden pass 59**) nor iter-245's three range buckets. Measured on
the live tree, in one process:

| call | RANGE_CITATIONS | RANGE_RESOLVED | RANGE_UNRESOLVED |
|---|---|---|---|
| 1 | **490** | 349 | 141 |
| 2 | **980** | 698 | 282 |
| 3 | **1,470** | 1,047 | 423 |

**The disclosure line pass 59 built specifically to make the ungraded range population VISIBLE published
a number that doubles on every repeat invocation**, and `RANGE_FINDINGS` — a *list* — would report a real
out-of-bounds range N times with the RED message's own `len()` N× the truth.

> **Invisible from two directions at once, which is why nine iters did not see it.** The **shipped** path
> was never wrong: `guard_family` runs each member as a **subprocess**, so every published verdict came
> from a fresh interpreter. The **tests** worked around it — four `.clear()` calls in their own `setUp`.
> *A test that compensates for a defect cannot detect it.* This is `§5` rule 77's second clause
> (in-process state defeating a control) with the roles reversed: the state did not defeat a mutation
> control, it was *hidden by* a fixture that tidied up after it.

Fixed where the siblings already reset. The fence is a test that runs the guard **twice in-process
without clearing between** — clearing there would reproduce the workaround — asserting an identical
census, plus an anti-vacuity pin on the exact tuple so equal-and-empty cannot pass for the right answer.
Mutation-controlled: `(3,2,1,1)` becomes `(6,4,2,2)` at `7d75c52`.

**(2) `env_absence_guard` sized its own reach with a quantifier it does not ship** (`rext` `7276c97`).
Every figure in two of its docstring sections — *41* family-absence sites, *23* naming a platform file,
*18* not, *15* documents, `*_RPC_ADDR` at *23* sites — is exactly what the draft measured while
`ABSENCE_RE` still admitted a bare `no`. That admission was then **removed**, and the comment on
`ABSENCE_RE` records why (it produced the guard's only two findings and **both were false**, one on a
sentence asserting the exact opposite). Narrowing the instrument moved the population; the census was
never re-taken. Measured on one tree with the guard's own shipped constants:

| quantifier | sites | name a platform file | name none | documents |
|---|---|---|---|---|
| **shipped** (`zero`/`nowhere`/`none`/…) | **18** | **17** | **1** | **11** |
| pre-narrowing draft (+ bare `no`) | 41 | 23 | 18 | 15 |

So the docstring sized the reach at **more than twice** what the guard holds, and its **printed** refusal
count (**1**) contradicted its **stated** one (**18**) on every run. **The number was right about a
population the code no longer has** — this milestone's most-repeated defect, turned on one of its own
guards. Both readings are now stated and labelled. The single-variable figures are pre-narrowing too and
are **left unreplaced rather than re-derived**: reproducing them needs iter-247's own `UPPER_SNAKE`
selector, which is not a shipped constant of that module, and a number derived with a different
instrument is not the same measurement. *Stated, not carried.*

**(3) this pass's own two defects, both caught by fences this milestone built.**
`TheCeilingProseDoesNotContradictTheCeiling` went RED because all three of pass 61's re-pin blocks opened
*"196 → re-pinned at harden pass 61"* with **no arrow target** — *a recorded reason that does not state
the number it justifies is not a recorded reason*. And `TheNounVocabularyIsMeasuredNotAssumed` went RED
on this pass's **own comment** (*"call 1 reports 490 … and call 2 reports 980"*), where `reports` is a
verb; rephrased so no number precedes it, rather than forcing an ambiguous word into a vocabulary bucket
other sites use the other way. Separately, `DOCSTRING_LITERAL_CEILING` was re-pinned **DOWN 234 → 232**:
replacing five prose figures with a two-row table took the live count below the ceiling set earlier in
the session, and two numbers of slack is room for two undocumented literals. **A ratchet that only ever
moves up stops bounding anything.** *(And this ledger's own attribution was wrong before it was written:
the toolchain-registry repair was captioned "harden pass 60" in five places across the corpus and the
guard; it landed in **61**. Corrected — `49b167c`, `dfb3fb6`.)*

**Also scanned, and clean — recorded because a scan that finds nothing is only evidence if it is stated:**
* **Dimension 5, fuzzing** — 4,000 random printable inputs × **12** regexes / pure functions across the
  five new guards, plus 6 family-glob shapes through `glob_to_re`: **0 exceptions**.
* `fence_command_guard`'s `--report` returns 0 on findings **by design**; `guard_family` invokes it
  **without** that flag, so a RED cannot be masked there.
* iter-242's symlink-disclosure condition read correct against four resolution cases (relative-real,
  absolute-real, symlinked-leaf, symlinked-parent).
* iter-248's SCOPE line: the **84** is `len(_test_files())`, derived from disk, and matches; and
  `_print_scope_disclosure()` is called **above every terminating branch**, so it rides a RED summary
  too — the rule pass-23 already established for the dirty-tree caveat.
* **Two pre-registrations on `env_absence_guard` REFUTED**, both optimistic: **0** live sites name
  `.env_example` without also naming compose (so grading against all three platform files is never wrong
  today), and compose carries **0** pass-through `- NAME` env entries (so requiring `=`/`:` misses
  nothing). Recorded per run-29's caution to predict honestly in both directions.
* **Dimension 6, benchmarks: no-op, stated.** The batch made nothing performance-sensitive.

**Tests added: 1 Python** (the repeat-run census pin; `test_anchor_construct_denominator.py` **45 → 46**).
**Session total: 12 Python test methods + 4 Go table cases.**

**VERIFICATION — runner, section scope and language on every figure.**
* **`stack-core`** (pytest 8.4.2 / CPython 3.9.6, **Python**): **2,096 passed · 0 failed · 3 skipped**
  (28 m 41 s) on the final tree — from **2,073 passed · 11 failed · 3 skipped** at session open. All 11
  closed; the arithmetic reconciles exactly (2,084 collected + 12 net-new = 2,096).
* **Four non-core Python sections**, re-measured rather than carried: `demo-stack` 1,063 passed / **9
  failed** / 2 skipped, `dev-stack` 151, `stack-injection` 335, `stack-verify` 275 → **1,824 passed · 9
  failed · 2 skipped**. The 9 were checked **name-for-name** against `suite_census.ENV_GATED` and are
  **all 9 of the 9 declared entries** — not "9 failures that resemble the declared ones".
* **Go**: `stack-seeding` **16 packages ok**, `go build ./...` + `go vet ./cmd/...` clean.
* **`guard_family`** (`--platform`, repo root): **29 GREEN · 0 RED · 0 could-not-check · 5 not-run** on
  the final tree. **Not a whole-family green** — it says so itself, and 5 members had no input supplied.
* **Flake gate**: the four touched Python test files **3× consecutive** (105 passed each) and the Go
  `TestSentinel*` **3× consecutive**. Clean.
* **NOT COVERED, stated rather than implied:** the **424 TypeScript tests** remain enumerated and never
  executed. No live stack was brought up, so gate clause 1 is untouched by this pass.

**Knowledge backfill:** one rule, and it is the session's. *A fence inherits the reach of the iter that
wrote it* — **four** instances across three passes (`skill_invocation_guard` fenced-only,
`rext_path_guard` bare-paths-only, `toolchain_floor_guard` one-document, `env_absence_guard`'s
superseded census). The counterexample keeps it honest: `fence_command_guard`'s fenced-block scope is
**correct**, because a `make` target needs a directory context and prose does not establish one. **The
test is not "is the scope narrow" but "is the narrowness a property of the CLAIM or of the iter."**

**Flakes stabilized:** none surfaced.

**Stop condition:** cap reached without stabilization — three passes, **12 defects fixed inline** (2 + 7
+ 3), and the third pass found the session's fourth instance of its own headline class rather than going
quiet. Per the user's standing ruling the routes below are recorded and NOT met with new machinery; the
**fifteenth** cap-without-stabilization in this milestone (22, 25, 26, 29, 32, 35, 38, 41, 44, 47, 50,
53, 56, 59, 62).

**Routes carried forward:**
- `ROUTE-M257x-h59-rext-edits-fire-no-fence-anywhere` → **re-affirmed, and now with a measured cost.**
  Pass 59 routed it; ten iters later it had produced **11 repo-wide RED assertions** nobody could see.
  It is the single highest-leverage open item in this milestone's tooling.
- `ROUTE-M257x-h62-live-markdown-refragmented` → **new.** Four private `live_markdown` copies re-grew
  the derivation iter-212 consolidated. Registered, not consolidated — the refactor spans five modules.
- `ROUTE-M257x-h59-range-anchors-are-ungraded` → **still open on its undecidable half** (which line of a
  range carries the claim); iter-245 closed the decidable half and pass 62 fixed its accounting.

## Pass 63 — 2026-08-10 — incremental

**Iters hardened this pass:** iter-249 … iter-257 (first pass over the batch).

**Tiks covered since prior pass:** 9 — the threshold is 10; run at 9 on the user's explicit go-ahead,
recorded as a deviation rather than a defect.

### ⚠️ TWO RECORDS THIS PASS CORRECTS — read these before quoting any prior figure

**(a) THE FRESH-CHECKOUT CLASS WAS NEVER ZERO. IT IS 3, AND THE `0` WAS WRONG WHEN IT WAS WRITTEN.**

iter-255's close and the pass-62 ledger both record the fresh-checkout-hostility class closing at
**zero**, and that figure was relayed onward as a closed class. Measured this pass on a frozen clone
pair — `suite_census.py --fresh-checkout --runner pytest`, **138 modules censused**, live control
re-running the failing modules — the class is **3**:

* `stack-verify/tests/test_e2e_collection_integrity.py::test_collection_succeeds_at_all`
* `stack-verify/tests/test_e2e_collection_integrity.py::test_every_spec_file_still_registers`
* `stack-verify/tests/test_e2e_collection_integrity.py::test_the_suite_collects_a_real_number_of_tests`

**This is NOT a regression.** Nothing broke after iter-255. The module was authored at v2.5 (M236
close), it predates the whole 249→255 repair, and it was never among the 22 tests that repair declared.
The class had a live residue of 3 on the day the `0` was published, and the `0` was **wrong when
written** — it was a statement about the 22 members that had been enumerated, published as a statement
about the class.

That is exactly what this pass's knowledge backfill says: **a repair closes its MEMBERS, not its
class.** iters 249–255 closed 22 named failures, correctly and well. Only a *derived* census can say
whether the class is closed, and the census had not been re-run since the repair — because it freezes a
clone pair and runs the whole suite twice, so nobody runs it.

Why the three matter rather than being a technicality: they shell out to `npx playwright test --list`,
which needs an installed dependency tree. On a clean clone they do not say *"this box has no
`node_modules`"* — they say **a spec is throwing at module scope** and **the suite has collapsed below
its floor.** Both sentences accuse THE CORPUS. A new reader cloning both repos is told the corpus is
broken, which is the precise failure mode the whole 249→255 arc existed to eliminate.

Now declared with the canonical mechanism (imported from `suite_census`, never re-spelled, so it cannot
drift from the other 20), raised in `setUpClass` so one statement covers all three arms.

**(b) THE COLLECTION FENCE GOVERNED SIX OF ELEVEN REXT SECTIONS — AND MISSED ONE OF THE FIVE THAT
`D-M257x-145-3` CALLS "THE SUITE".**

`tests/test_test_collection_fence.py` enforces *no test module may hide tests from the runner that reads
it*. Its `_SECTIONS` tuple is **enumerated, not globbed** — for a good reason, preserved: a glob from the
repo root would sweep in the vendored platform clones under `demo-stack/stacks/*/clones/**`, and a fence
that starts failing because someone brought a demo up gets turned off rather than fixed.

Measured at this pass, the three declared section populations in this repo do not agree:

| declaration | n | members |
|---|---|---|
| `claim_census_guard.REXT_SECTION_NAMES` | 11 | every non-`knowledge` dir |
| `suite_census.SECTIONS` (the "five sections" of `D-M257x-145-3`) | 5 | demo-stack · **dev-stack** · stack-core · stack-injection · stack-verify |
| `test_test_collection_fence._SECTIONS` **before this pass** | 6 | clerkenstein · demo-stack · stack-core · stack-injection · stack-seeding · stack-verify |

**The five sections the fence did NOT govern were `dev-stack`, `stack-secrets`, `stack-snapshot`,
`alignment` and `playthroughs`.** Four of those hold zero Python test files, so their absence cost
nothing today. **`dev-stack` is the exception: 5 modules, 146 `TestCase` tests, 151 collected by
pytest** — and it is a *member of the five*. So the fence covered **4 of the 5** sections that
`D-M257x-145-3` defines as "the suite", while including two (`clerkenstein`, `stack-seeding`) that carry
no Python tests at all and are outside that five.

The hole was **LATENT, not live** — all 5 `dev-stack` modules read clean on all three of the fence's
predicates when found, and the section's own suite is green at 151. It is recorded because it is the
same shape as the two defects the in-scope iters actually hit (iter-255's unreadable mutation battery;
iter-257's six tests below the `__main__` guard), one level up: **the fence that exists to catch tests no
runner reaches was itself unable to reach a whole section.**

Fixed by governing `dev-stack` — `_SECTIONS` is now **7**, a superset of the five — and by fencing the
enumeration in **both directions** (`TheSectionListReachesEveryRextSection`), so a section that grows its
first Python test can no longer stay ungoverned in silence.

> **⚠️ THE DISAGREEMENT IS RECORDED, NOT RESOLVED.** `D-M257x-145-3` — *"the suite" means all five
> sections* — **remains the user's to rule**, exactly as iter-186 left it. This pass adds one fact that
> ruling did not have: a **third** declared population existed alongside the other two, it disagreed with
> both, and it omitted a member of the very five the assumption names. Nothing here decides what "the
> suite" means; `_SECTIONS` was widened to cover the five **and** the two extras it already had, which is
> a superset under either reading and therefore does not pre-empt the call.

**Scope.** 30 files in `rosetta-extensions` (10 source modules + 20 test modules, +2,370 lines) and 11
corpus files. All nine are tiks; no tok in the batch.

**Bugs surfaced + fixed inline: 2.**

**(1) The test-collection fence governed SIX of ELEVEN rext sections, and `dev-stack` — 5 modules,
146 tests — was in neither the list nor any exclusion** (`rext` `bc2d5c6`). `_SECTIONS` is *enumerated*
rather than globbed, and the reason at its definition is good: a glob from the repo root would sweep in
the vendored platform clones under `demo-stack/stacks/*/clones/**`, and a fence that starts failing
because someone brought a demo up gets turned off rather than fixed. What the enumeration lacked was a
reader that notices when it rots. Every arm of that file — the `__main__`-guard arm, the 3.9-syntax arm,
the pytest-only arm — was structurally unable to reach `dev-stack`.

The hole was **LATENT, not live**: all 5 modules read clean on all three predicates when found, and
`dev-stack`'s own suite is green at 151. That is *why* it was worth catching rather than a reason to
shrug — it is the same shape as the two defects the in-scope iters actually hit, one level up:

* **iter-255** — the mutation battery carrying ~30 of this milestone's mutation proofs raised
  `ModuleNotFoundError` under `suite_census.collected_by_pytest`'s own invocation and was recorded
  `UNREADABLE (-1)` for **four iters**.
* **iter-257** — arm D's six tests were appended BELOW the `__main__` guard, so direct execution skipped
  them and printed OK. Caught by this very file, in a section it does govern.

Fixed by governing `dev-stack` **and** by fencing the enumeration in both directions
(`TheSectionListReachesEveryRextSection`): a section holding python tests must be in `_SECTIONS` or
declared in `UNGOVERNED_SECTIONS` with a reason, and a declaration that outlives its subject fails. The
original reason for enumerating is preserved and separately pinned.

**Watched RED → GREEN, on the real tree:** appending a 2-test class below the guard in
`dev-stack/tests/test_aws_heal.py` now names the file, line, class and test count. Before the fix that
mutation was invisible. Applied and reverted atomically; the file is byte-identical afterwards.

**(2) The fresh-checkout repair had NO watcher.** iters 249–255 drove the class 29 → 0, and the repair
is not one code change — it is **20 declaration sites over 12 modules**. The census that measured the
class freezes a clone pair and runs the whole suite twice, so nobody runs it, and both silent failure
modes were therefore unobserved: a declaration **deleted** (the arm goes back to reporting the CORPUS as
wrong on a clean clone), or a declaration that **keeps its guard and loses its REASON** — which still
skips, so the suite stays green and the census still reads zero while the operator loses the one
sentence naming what to provision. The second is what `D-M257x-249-2` is actually about, and it is how a
repair degrades into its own cargo cult with every automated reading saying fine.

`TheREPAIRIsRatcheted` watches the **repair** rather than re-measuring the class: pure AST, no
subprocess, every invocation. Floor on the site count; every site must skip AND cite
`census.CLONE_SET_REASON` / `NODE_MODULES_REASON`; both predicates must stay in use (the `node_modules`
half is what refuted iter-254's `PR-1`). It does **not** catch a brand-new hostile test — that still
needs `suite_census.py --fresh-checkout`, and the docstring says so.

**Also scanned, and clean — recorded because a scan that finds nothing is only evidence if it is stated:**
* The **mutation battery's iter-255 fix HOLDS**: 6 tests collected under *both* the rext-root invocation
  (the census's own) and the `stack-core` invocation, re-verified again at the end of the session after
  `suite_census` had been modified.
* **Dimension 3, error paths** on the two most-reworked in-scope guards: `rext_path_guard`'s four
  refusal paths and `corpus_citation_guard`'s two are **all pinned by exit code**. An initial reading
  that called them untested was grepping message TEXT rather than behaviour, and was wrong in the
  optimistic direction — recorded because the optimistic direction is the one worth confessing.

**Tests added: 12** (`test_test_collection_fence.py` 44 → 50; `test_fresh_checkout_census_m257x.py`
24 → 30 at this pass).

**Stop condition:** continue-to-next-pass — the reach question this pass opened (does an enumerated
repair close its class?) is measurable and was not yet measured.

## Pass 64 — 2026-08-10 — incremental

**Iters hardened this pass:** iter-249 … iter-257 (second pass over the batch).

**Dimensions:** 2 (edge cases), 3 (error paths), 5 (fuzzing), against the iter-diff scope.

**Bugs surfaced + fixed inline: 4.** One theme found three times independently, plus a grading gap.

**(1)+(2) `test_fence_command_guard`, TWO fail-OPEN rungs** (`rext` `ffa4238`). Both live arms skipped on
`code == 2` / `total == 0` with the reason *"no clone set on this host"* — a condition the
`census.clone_set_present` guard **four lines above** has already excluded (iter-254 added it). Exit 2 has
three causes (`fence_command_guard.py:351/359/364`): no `CLAUDE.md`, ZERO fenced blocks, or zero gradeable
commands. Past that guard the clone-set explanation is gone, so what remains is **the fence regex or the
corpus scope having drifted — i.e. the guard itself being broken.** That is the one outcome meaning *this
fence stopped working*, and it was reported as a benign environmental skip with a reason that was
definitionally false. Now an `assertNotEqual`, so it is a RED. iter-254's own comment three lines up
already recorded that a bare checkout grades **1** command and not 0, so `total == 0` was
documented-unreachable on the very tree the rung claimed to excuse.

**(3) `advance_impact_census.classify_pair` reported a non-positive line number as an over-run.** `n < 1`
shared a branch with `n > len(old)` and inherited its message, so `classify_pair(["a","b"], …, -1)`
returned *"line -1 > 2 at old ref"* — arithmetically false. Surfaced by a **6,000-input fuzz** (None /
empty / short line lists, `n` drawn from -3..9): **0 exceptions, 0 out-of-vocabulary buckets.** The logic
was never in doubt; the SENTENCE was. It matters because this census exists to tell a reader why a
citation did not survive an advance — that message names a truncation, and the reader goes looking at the
new ref for a file that got shorter, when the real condition is a malformed anchor that never had a valid
line number at either ref. Split, with the genuine over-run message kept exact.

**(4) `clone_pin_guard` arm D read the workspace copy's EXTRA keys and its differing VALUES, and was
SILENT on its MISSING ones.** Not symmetry-bookkeeping: **arm B already treats a hole in the canonical as
a finding**, on the stated ground that `pinned` *"leaves an unpinned repo UNTOUCHED, so the reproducibility
barrier silently does not cover it"* — and arm D's entire premise is that the mechanism reads the **COPY**,
not the canonical. So the sentence justifying arm B applies with more force one file over, and nothing was
checking it.

The direction is the one that will occur. iter-257 found the copy as a **SUPERSET** (11 names to the
canonical's 6) because entries had been REMOVED. The mirror arrives the first time the canonical **GAINS**
one: `ensure-clones.sh` seeds the copy copy-if-absent (`:204`) and never reconciles, so every workspace
that already exists keeps the shorter file indefinitely — the new repo is unpinned on precisely the
longest-lived stacks, which are the ones a presenter is most likely to be holding.

**DISCLOSED, not failed**, and that is a deliberate limit rather than timidity: a missing key is the one
divergence an operator could plausibly mean (*"do not touch this repo on this stack"*). Ruling it a finding
would answer a question iter-257 did not ask; ending the silence needs no ruling. The summary noun moved
`value drift` → `divergence` so a reader of that line alone cannot miss the second shape.

**Tests added: 10** (`test_advance_impact_census.py` 24 → 29, incl. the fuzz + unicode/oversized grid;
`test_clone_pin_guard.py` 29 → 34, incl. `D11`, a mutation control proving the pre-pass reading sees
NOTHING on a holed copy).

**Stop condition:** continue-to-next-pass — the class-closure question was still unmeasured, and the
integration reading (guard family + whole suite) had not been taken.

## Pass 65 — 2026-08-10 — incremental

**Iters hardened this pass:** iter-249 … iter-257 (third and final pass over the batch).

**The pass that measured the batch's headline claim, and refuted it.**

**Bugs surfaced + fixed inline: 4** (plus 3 ratchet re-pins and one self-review correction).

**(1) THE FRESH-CHECKOUT CLASS WAS NOT ZERO. IT WAS 3.** (`rext` `adc2848`.) iter-255 and the pass-62
ledger both record the class closing at **zero**. Measured this pass with the tool built for it —
`suite_census.py --fresh-checkout --runner pytest`, a frozen clone pair, **138 modules censused**, live
control re-running the failing modules — the reading is:

| bucket | n | node-ids |
|---|---|---|
| DECLARED (already honest) | 3 | `demo-stack/tests/test_migrate_race_live.py` ×3 |
| **BOX (fresh-checkout-hostile)** | **3** | `stack-verify/tests/test_e2e_collection_integrity.py` ×3 |
| REAL | 0 | — |

All three shell out to `npx playwright test --list`, which needs an installed dependency tree. On a clean
clone they do not report a missing precondition — they report **that a spec is throwing at module scope
and that the suite has collapsed below its floor.** Both sentences accuse THE CORPUS. The module predates
iters 249–255 and was simply not among the 22 they declared.

Nothing was wrong with the repair; the error is in what was concluded from it. Declared with the canonical
mechanism (imported, never re-spelled, so it cannot drift from the other 20), raised in `setUpClass` so one
statement covers all three arms and a bare box pays none of the 300 s timeout.

**(2) `node_modules_present` could not see `root/node_modules`.** The glob ran depth 1..5, so the one tree
invisible to it was the one whose own child is `node_modules` — what a caller gets passing a project
directory rather than a root above it. The first cut of the declaration above passed `E2E` and **got a SKIP
on a box with a fully installed e2e tree.** This is the **second** time this one function has answered
False on a machine that plainly has one, and the second time a live control caught it rather than review —
so depth 0 is now covered and the boundary is pinned by 5 arms, including a file-named-`node_modules` case.

**(3) The iter-257 divergence disclosure fired UNCONDITIONALLY.** `ensure-clones.sh` compared with
`! cmp -s`, and `cmp` is not a bash builtin. Absent from `PATH` it exits 127; `!` inverts that to TRUE; the
disclosure fires on **every** run with both files present, including byte-identical ones. A warning that
always fires is not a weaker warning — it is one an operator learns to scroll past, which returns the seam
to **the exact silence iter-257 fixed**. Now `[ "$(<a)" != "$(<b)" ]`, bash's own redirection, with no
subprocess to be missing. Found by the **pass-direction** test written for the disclosure — the test that
had no reason to exist except that a disclosure needs one.

The disclosure had also landed with **no assertion of its own**: it executes inside
`test_canonical_pin_never_clobbers_operator_workspace_pin`, whose fixture already differs, so the branch was
run by the suite and read by nothing. That is the weaker half of this milestone's defect class — not a check
that never runs, a check nobody reads.

**(4) THIS PASS'S OWN RATCHET HAD THE DEFECT IT WAS WRITTEN TO PREVENT.** `TheREPAIRIsRatcheted` (pass 63)
globbed `stack-core/tests/` and matched only `if not <pred>(...)` + `self.skipTest`. The **twenty-first**
declaration site — added by this same pass, in `stack-verify`, spelled
`if census is not None and not …: raise SkipTest(…)` inside a `setUpClass` where `self` does not exist —
was structurally invisible to it. It would have reported *"20, unchanged"* while the repair grew underneath
it. Now repo-wide, condition searched as a **subtree**, both skip spellings recognised: **21 sites over 13
modules**, floor re-pinned 20 → 21.

**Self-review correction** (`rext` `669ce4e`): pass 63's `dev-stack` arm asserted `len(swept) == 5` two
lines below a docstring explaining why the test count is a FLOOR — *"a ratchet that fights growth gets
deleted rather than fixed."* A sixth module would have turned it RED and printed a message reading as a
defect. Now a floor.

**Integration reading — the guard family, and §7 rule 4 applied to our own repo.** The 8-line
`ensure-clones.sh` fix moved cited lines and turned **three** guards RED from one cause — the identical
side effect iter-257 recorded from the identical file:

| | before | after |
|---|---|---|
| `guard_family --platform` | 26 GREEN · **3 RED** · 0 could-not-check · 5 not-run | **29 GREEN · 0 RED** · 0 could-not-check · 5 not-run |

RED: `anchor_construct_guard`, `demo_knob_guard`, `repair_postcondition`. Re-pointed in
`corpus/ops/demo/demo-up-defaults.md`: `DEMO_ADVANCE_CLONES` `ensure-clones.sh:220 → :228`,
`DEMO_FRESHNESS_STRICT` `:475 → :483`. **Read off the file, not derived by adding 8** — the arithmetic
happened to agree, but a shift is not a guarantee and the guard exists precisely because nobody should
trust arithmetic here. Only the LIVE corpus doc was re-pointed; the `knowledge/` hits are dated statements
by closed iters, and re-pointing those would falsify a record rather than repair a reference. The family was
re-run **again** after the knowledge backfill, because that append adds numeric claims and numeric claims are
what `claim_census_guard` and `derived_count_guard` grade: still 29 GREEN / 0 RED.

**Three literal ratchets re-pinned** (`rext` `99c4872`): `DOCSTRING` 238 → 238 (untouched), `COMMENT`
222 → 226, `TEST_MODULE` **637 → 648** — the largest single move that block has taken. The size is the
point rather than an apology: the pass added 31 arms and the ones that moved it are the ones whose reach is
a NUMBER. `_MEASURED_NOUNS` also widened by `exceptions`, surfaced by the residual arm and not by reading
the list — the **sixth** time the vocabulary's reach has closed on the sentence that widened it. All three
ceilings now sit at **slack 0**. The two-pass convergence iters 254–255 recorded was, this time,
**anticipated in advance rather than discovered by a second RED**.

**VERIFICATION — runner, scope and language on every figure.** pytest 8.4.2 / CPython 3.9.6
(`/usr/bin/python3`, the only interpreter on this host carrying pytest — `§9`'s measurement-preconditions
block).

| section | result |
|---|---|
| `stack-core` — **the 17 consumers of the two shared modules this pass changed** (`suite_census.py`, `derivation_registry.py`), derived by import-grep, not chosen | **542 passed · 0 failed · 1 skipped** (420 s) |
| `stack-core` — the 6 modules this pass touched, **3× consecutive** (flake gate) | **163 passed** each run |
| `stack-core` — `test_frozen_expectation_census_m257x` (the three literal ratchets) | **99 passed** |
| `demo-stack` | **1,067 passed · 9 failed · 2 skipped** (254 s) |
| `dev-stack` — newly governed by the collection fence this pass | **151 passed** (99 s) |
| `stack-injection` | **335 passed** (7 s) |
| `stack-verify` — includes the module this pass declared | **275 passed** (420 s) |
| Go (`stack-seeding`) | `go build ./...` **rc=0**, `go vet ./cmd/...` **rc=0** |
| `guard_family --platform` | **29 GREEN · 0 RED · 0 could-not-check · 5 not-run** |

**The 9 `demo-stack` failures were checked NAME-FOR-NAME against `suite_census.ENV_GATED`, not merely
counted:** normalising the class segment out of the pytest node-ids, **0 undeclared failures and 0
declared entries that failed to fire** — all 9 are exactly the 9 declared entries. Same reading pass 62
took, re-taken rather than carried.

**THE WHOLE-SECTION `stack-core` RUN LANDED: 2,191 passed · 0 failed · 3 skipped, in 2,234.93 s
(37 m 14 s).** Started 13:05:04Z, finished 13:42:19Z. Executed population **2,194**, against pass 62's
**2,099** (2,096 passed + 3 skipped) — the growth is iters 249–257's own additions plus this pass's
**+27** stack-core arms.

> **⚠️ ONE COMMIT IS OUTSIDE THAT RUN'S COLLECTED TREE, and the boundary is stated rather than
> smoothed over.** pytest imports at collection, so the run reflects the tree at **13:05:04Z**.
> `99c4872` (the three ratchet re-pins) landed at **13:04:59Z — five seconds before**, so it IS covered.
> `669ce4e` (the `dev-stack` arm's `assertEqual` → `assertGreaterEqual` floor fix) landed at
> **13:16:44Z**, eleven minutes *after*, so it is **NOT**. That one module was therefore re-run at
> current HEAD, **3× consecutive: 35 passed each**. Nothing else in the pass falls outside the window.
>
> An earlier draft of this entry recorded the run as *"did not land, no number claimed"* and offered the
> import-derived consumer sweep in its place. The run then completed. **The sweep is kept below rather
> than deleted** — it was the right instrument for the question *"what could these two shared-module
> edits break?"*, it is derived rather than chosen, and a whole-section green does not retroactively make
> a scoped reading worthless. What is withdrawn is only the *caveat*, not the evidence.
>
> **Timing remains CONTENDED and is not a baseline.** 37 m 14 s against pass 62's 28 m 41 s for the same
> section, at a measured ≈9 % CPU utilisation under permanent third-party host load (VS Code helpers and
> an unrelated `cart-runner` were resident throughout). The gap is contention, not a regression, and the
> same 6-module set varied 10 s → 21 s between flake-gate repetitions.

**Knowledge backfill** (`corpus/ops/platform-alignment.md`, this milestone's `iteration_protocol_ref`) —
one rule, and it earned its place by firing three times inside one pass:

> **A repair CLOSES ITS MEMBERS, not its class — and the watcher inherits the same limit.**
> An enumerated repair closes the members you enumerated; only a **derived** census can say whether it
> closed the class, and a class measured once, at the moment it was closed, is a dated reading rather
> than a property.

The three instances: the collection fence governing 6 of 11 sections; the fresh-checkout repair closing
22 named tests while the class kept a residue of 3; and this pass's own ratchet, blind to the site the
same pass created. That is **pass 62's own rule — *a fence inherits the reach of the iter that wrote
it*** — one level up, and the sharper form is the test to apply: **ask whether the narrowness is a
property of the CLAIM or of where its author was standing.**

The entry also records *why* the watcher ratchets the repair rather than re-running the census: the
census freezes a clone pair and runs the whole suite twice, so nobody runs it — and a watcher nobody
runs is the thing being guarded against.

**Flakes stabilized:** none surfaced. **Flake gate clean:** the six touched modules **3× consecutive**
(163 passed each) and the four new shell-driven divergence tests **3× consecutive** (6 passed each).
Wall-clock on identical sets varied 10 s → 21 s, which is the contention disclosure below, not a signal.

**Timing is CONTENDED and is not a baseline.** The host runs other work permanently and this session ran
up to four suites concurrently; every duration here is wall-clock under contention and must not be quoted
as a measurement.

**NOT COVERED, stated rather than implied:**
* The **TypeScript** suites remain enumerated and never executed; no live stack was brought up, so gate
  clause 1 is untouched by this pass.
* The **mutation battery was collect-verified, not executed** (6 tests under both invocations, twice —
  once at session open and again after `suite_census` was modified). It was deliberately NOT run: it
  mutates source files in place, and the whole-section suite was in flight, so executing it would have
  corrupted the very reading this pass depends on.
* `guard_family` reports **29 of 34** members; **5 are NOT-RUN** for want of `--range`/`--ledger`. That is
  not a whole-family green and the runner says so itself.

**Routes carried forward:**
- `ROUTE-M257x-h59-rext-edits-fire-no-fence-anywhere` → **re-affirmed, and this pass paid its cost in
  cash.** Pass 59 routed it; pass 62 measured 11 repo-wide RED assertions nobody could see. This pass
  added the direct evidence: an 8-line edit to `ensure-clones.sh` turned **three** guards RED, and the
  only reason it was caught is that a human-driven harden pass happened to run `guard_family` by hand.
  Still the single highest-leverage open item in this milestone's tooling.
- `ROUTE-M257x-h62-live-markdown-refragmented` → **still open**, untouched by this pass.
- `ROUTE-M257x-h59-range-anchors-are-ungraded` → **still open on its undecidable half** (which line of a
  range carries the claim). `anchor_construct_guard` continues to report 490 range citations, 349
  bounds-checked, 141 refused.
- `ROUTE-M257x-h65-fresh-checkout-class-needs-a-scheduled-remeasure` → **NEW, and it is the honest
  residue of this pass's headline finding.** The pass-63 ratchet watches the *repair* (cheap, every
  invocation) and provably does **not** catch a brand-new hostile test written from scratch — which is
  exactly how the residue of 3 arose in the first place, from a module that predated the repair. Closing
  that needs `suite_census.py --fresh-checkout` on a cadence, and the census costs a frozen clone pair
  plus two whole-suite runs. Registered rather than solved: making it cheap enough to run per-batch is a
  tooling change, not a harden-pass change, and the standing ruling is to route rather than build new
  machinery inside a harden pass.

**Stop condition:** cap reached without stabilization — three passes, **10 defects fixed inline**
(2 + 4 + 4), three literal ratchets re-pinned, one self-review correction, and the third pass **refuted
the batch's own headline claim** rather than going quiet. Per the user's standing ruling the routes above
are recorded and NOT met with new machinery; the **sixteenth** cap-without-stabilization in this
milestone (22, 25, 26, 29, 32, 35, 38, 41, 44, 47, 50, 53, 56, 59, 62, 65).

## Pass 66 — 2026-08-10 — incremental

**Iters hardened this pass:** iter-258 … iter-269 (12 tiks, the milestone's payload range — the demo and
dev halves of the closing condition, then the corpus-accuracy limb). The production-code footprint of
that range is small and concentrated: **one new 508-line guard plus its registry enrolments** (rext
`25cdb84`, `450c57d`, `cf4da42`) and **nine corpus documents**. Everything else is milestone artifacts.

> **Pass-numbering disclosure.** This session ran passes **66, 67, 68**. In-code annotations from passes
> 66 and 67 are labelled *"harden pass 66"* uniformly: they were one contiguous stream of work over one
> batch, sharing a ratchet re-pin, and the 66/67 split is a **ledger-level grading of stop conditions**,
> not two code epochs. Recorded rather than tidied, because a reader grepping `pass 67` in rext will
> find nothing and should know why.

**Subject:** `decommissioned_instruction_guard.py` — iter-265's fence, and specifically its **assertion
D**. The orchestrator's stress note was exact: iter-265 *measured* that its own assertions A–C miss the
defect the guard exists for (they fired on **1 site of 17**, and on **none of the 3 that mattered**,
because the obsolete remedy's own words — *"has been removed from"* — are marker vocabulary). D is the
only assertion that hits, so D is the guard.

**Bugs surfaced + fixed inline (3, all in D, commit `b2a0c74`):**

1. **D FAILED OPEN — `run()` returned 0, a full GREEN, on any tree with no clone set**, while printing
   `assertion D could not run` four lines above the verdict. So a fresh checkout — the exact class
   passes 63–65 spent three passes on — ran only the assertions *known to miss*, skipped the one that
   hits, and reported OK. Now exits **2** (could-not-check).
   **The convention was already written down twice, in the `INVOCATIONS` entries either side of this
   guard's own**: `rext_path_guard` and `toolchain_floor_guard` each *"exit 2 (never 0) when [the tree]
   is absent, so a host without the clone reports could-not-check rather than an unearned green."* This
   member was enrolled **between them** and did not do it. **Printing your reach is not the same as
   refusing a verdict you did not earn.**
2. **A true outcome published with a reason that can be false.** `NO clone root provisioned` was printed
   for *both* empty-index states — roots absent, and roots present carrying no build input, which is
   what a half-cloned stack actually produces. Three states now, named apart. Same class harden pass 64
   fixed in three sibling checks.
3. **An undisclosed reach floor.** The 20-char index minimum was an unnamed literal in two places and
   disclosed in neither. A live line below it is invisible to D — `COPY studio /build` (18 chars)
   indexes as nothing, so a page telling an operator to delete it reads green. Named `MIN_LIVE_LINE`
   and printed on every run.

**The test that had the finding in its own docstring.** `test_reach_is_a_property_of_the_clone_set` is
documented *"with no clone root, D cannot run, and must say so rather than pass quietly"* — and graded
only that the **index** was empty, which is a fact about the indexer, not about the verdict. It measured
green while `run()` returned 0 on the same input. **A test can state the property it does not assert**,
and this is the second time in three passes that the docstring was ahead of the body.

**Tests added:** `tests/test_decommissioned_instruction_guard.py` **12 → 17** — the could-not-check exit
(both reach states), the reason-not-just-outcome arm, A–C-RED-outranks-could-not-check precedence, and
the `MIN_LIVE_LINE` disclosure (which names the 18-char line, because a test asserting "D has a floor"
would pass against a floor that reaches nothing).

**Coverage delta on touched files:** `decommissioned_instruction_guard.py` — the three `run()` exit
branches were **1 of 3** covered (only the RED path); now **3 of 3**. The `d_reach` message had **1 of
2** states reachable by any test; now **3 of 3**.

**Literal ratchets re-pinned, with recorded reasons:** `COMMENT` 228 → 233, `TEST_MODULE` 648 → 650;
`DOCSTRING` 239, untouched. **The comment ratchet converged in two passes and the ANTICIPATION WAS STILL
SHORT** — 231 was pinned on the arithmetic *"the reason adds about one"* and re-measured at **233**,
because this block's reason-comment is the densest carrier of numbers in the file. The 231 is left
standing in the prose as the record of a bump that was **reasoned rather than read**, and the block now
says: **budget the second pass, do not predict its size.**

**Flakes stabilized:** none surfaced. **Flake gate:** the touched module **3× consecutive**, 22 passed
each (12.62 s / 12.11 s / 12.17 s).

**Stop condition:** continue-to-next-pass — the guard is repaired but the enrolment shape that let
iter-265 ship it half-wired is untouched, and iters 267/268/269's stress targets are not yet reached.

## Pass 67 — 2026-08-10 — incremental

**Iters hardened this pass:** iter-265, iter-268, iter-269 (three arms, one per stress target).
Commit `abadff7`; the stack-secrets arm in the same commit.

### Arm 1 — make the enrolment omission IMPOSSIBLE, not detectable

iter-265 shipped a guard without enrolling it; iter-269 found it by a **37-minute whole-suite run, four
iters late** — 24 tests RED across four registry fences. The author had run the guard, its own 12 tests,
and the four fences they believed adjacent, all green. *A new member is not tested by the tests it ships
with.* This is the **second** time in the milestone a new fence was not wired into the family.

**Measured what is UNIVERSAL before building anything, and only two of the four surfaces are.**
`derivation_registry.DECISIONS` is keyed per *derivation function* (a guard with none needs no entry) and
the literal ratchets are measured counts enrolled by nothing — **a blanket "must appear in all four"
would be false for most members**, and would have been the obvious wrong arm.

Of the two universal surfaces, **`INVOCATIONS` was already fail-closed**: `reconcile()` refuses a family
verdict outright (EXIT 2, *"a guard that cannot be selected is not a guard"*). **The provenance stamp was
not** — its predicate lived in `tests/test_fence_provenance.py`, so it could only be found by whoever ran
the suite, which for four iters was nobody. That is the whole difference between *detectable* and
*impossible*, and it is why the two omissions travelled together.

**Fix:** promoted `stamps_on_direct_execution()` out of the test into `fence_provenance` itself and read
it from `guard_family.reconcile()`. Both universal surfaces now fail closed at the moment anyone runs the
runner. Proven by **mutation against a real member with its stamp stripped**, not a synthetic stub.

**What the arm then cost, and why the fixtures were fixed rather than the arm narrowed:** six staged
fixtures across two modules stopped looking like real members and the family exited 2 before reaching the
behaviour they test. **A fixture that a real fence's requirement would reject cannot stand in for a real
fence** — so they were made compliant, through **one shared `stamped()` helper** reused by the sibling
module. That is the same cure iter-175's note on that very fixture already prescribed for the same shape;
its comment calls itself *"the SIXTH fixture of this shape."* It is now the seventh occasion, and the
seventh was avoided by reusing the one definition instead of growing a local copy.

**One branch of the new arm is UNREACHABLE and is pinned as such rather than deleted.** An unparseable
member is fail-closed *harder* one layer up: `repair_postcondition.declared_kind` raises `CouldNotRun`
before `reconcile()` is reached. The test was rewritten to assert the property that actually holds, and
says why the unreachable `except SyntaxError` stays (defence in depth for a direct `reconcile()` call —
which is how this pass's own mutation control invoked it).

### Arm 2 — what else is a list nobody fenced? (iter-268)

iter-268 found **one** named-consumer list hardcoding a decommissioned repo (`_studio_repos="cms"`,
dormant on a fresh box and live on any box carrying the fossil), closed that member, routed the fix, and
recorded that **nothing derived the SET of registries**. Asked the generalising question. **There are
three more, in two files, two of them on the demo bring-up path:**

| site | corpses named | disposition |
|---|---|---|
| `demo-stack/up-injected.sh:216` `INJECT_CANDIDATES` | `cms`, `jobsimulation` | disclosed **and filtered** — but the filter **FAILS OPEN** |
| `stack-injection/gen_injected_override.py:52` `INJECTED` | `cms`, `jobsimulation` | its own comment still asserts a live set of 3 that `ef32d4c` falsified |
| `stack-injection/gen_injected_override.py:153` `REUSE_DEV` | `storage`, `roadrunner` | dormant (opt-in `--reuse-dev-images`, default off) |

**The dispositions differ and that is the point — this is not three copies of one bug.** The sharpest is
the first: `derive_inject_svcs` filters the candidates against the platform compose's own build set, but
on an empty derivation it logs a warning, `return 0`s, and **leaves `INJECT_SVCS` holding the unfiltered
list** (`up-injected.sh:1691-1694`). So the derivation's failure mode is to inject two decommissioned
services. **That is the same shape as this session's assertion-D finding, one directory over**: a check
that could not run, reporting as though it had. Two independent instances in one pass is what makes it a
class rather than a coincidence.

**Pinned as a CHARACTERISATION with a disclosed reach.** The predicate reads assignments enumerating
**≥2** repo-ish tokens, so **iter-268's own one-token `_studio_repos` is NOT reached by it** — stated in
the test rather than papered over by widening the predicate until the known answer appears, because *a
predicate tuned until it returns what you already knew measures nothing*. **Known population 4, graded
3.** Fixing any of them needs a tag + pin bump (spending `D-M257x-258-1`'s frozen-pin control) → routed.

### Arm 3 — the safety property iter-269 refuted a routed fix with

iter-269 rejected iter-262's *"make the writer replace-or-skip"* because replace-in-place would break one
of two properties: **values-blindness** (*"an existing line is never re-read for its value or rewritten"*)
or the **trailing blank** that disarms `DIRECTUS_TOKEN` by last-wins. Checked whether those are
**asserted** or only reasoned about:

* the blanking half **is** asserted, and well — `TestProvision_PreexistingArmedDirectus_ForceBlanks`
  explicitly reads the **last** occurrence;
* **the values-blindness half lived in a comment at `io.go:173-175` and nothing else.**
  `TestProvision_Idempotent` covers only the **non-force** path; `TestProvision_Force` asserts the new
  line arrives, never that the old one survives.
* and **no test applied `--force` more than once**, while the behaviour the demo depends on is what
  happens on the **31st** bring-up.

So a refutation that decided the shape of a routed fix rested, in half, on unasserted reasoning.
`TestProvision_RepeatedForce_PreservesEveryPriorLineAndKeepsTheBlankWinning` now pins both halves under
repetition: prior content must remain a literal **prefix** (which catches rewrite, reorder and deletion
alike, where *"the old line is still in there somewhere"* would pass over all three), the last
`DIRECTUS_TOKEN` must still be blank after **every** run, and the per-run growth is pinned as a
characterisation so a future pruning fix cannot land without confronting both properties. Anti-vacuity
counter on the prefix check (it is skipped on run 1 by construction); **mutation-controlled** — a
`writeTargetFile` that rewrites instead of appending fails it at run 1.

**Tests added:** +5 `test_decommissioned_instruction_guard.py` (17 → 22, the rext named-consumer census)
· +3 `test_guard_family.py` (fire / accept / reach-boundary) · 1 Go test in `stack-secrets/provision`.

**Coverage delta on touched files:** `guard_family.reconcile()` — **4 → 6** complaint branches, all
covered in both directions. `fence_provenance.stamps_on_direct_execution` — unchanged coverage, **two**
call sites instead of one (test-time *and* run-time).

**Ratchets re-pinned:** `COMMENT` 233 → 236, `TEST_MODULE` 650 → 653. **The comment ratchet converged in
two passes for the third time in its history**, so the block stopped calling it a surprise.

**Flakes stabilized:** none surfaced. **Flake gate:** `stack-secrets/provision` **3× consecutive**
(`-count=1`, no cache); `test_decommissioned_instruction_guard.py` **3× consecutive**, 22 passed each.

**Stop condition:** continue-to-next-pass — the arms are landed and verified against their own modules,
but iter-269's lesson is precisely that *targeted modules are not the instrument that catches this
class*, so the whole-section run and the knowledge backfill are still owed.
