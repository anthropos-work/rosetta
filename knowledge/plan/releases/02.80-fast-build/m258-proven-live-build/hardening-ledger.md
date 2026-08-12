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
