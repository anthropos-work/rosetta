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
