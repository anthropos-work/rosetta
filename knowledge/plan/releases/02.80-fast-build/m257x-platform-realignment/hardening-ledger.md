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
