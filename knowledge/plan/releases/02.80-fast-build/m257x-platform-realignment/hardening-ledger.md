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
