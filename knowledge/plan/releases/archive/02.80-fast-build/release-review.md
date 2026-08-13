# Release Review: v2.8 "fast build"

**Date:** 2026-08-12
**Milestones:** M255 · M256 · M257x · M257 · M258
**Reviewed at:** `release/02.80-fast-build` @ `c69ec062` · rext @ `d06a56d`

> **Phase 1b verdict: RED.** The release cannot tag until the items below receive explicit fates.
> **Phase 4b verdict: YELLOW.** No hard condition failed; the benchmark condition is a *refusal*, not a pass.
> **Zero platform-repo edits — VERIFIED HOLDING** across all 19 clones. **Zero net-new third-party deps.**

---

## Scope

- [ ] **[blocker] ~9 routings / ~20+ discrete items reached M258's plan and were never fated.** They appear in none of its 20 iters, `decisions.md`, `hardening-ledger.md`, `deferrals-audit.md` or `carry-forward.md`.
- [ ] **[blocker] `FIX-M256-studio-false-green` — VERIFIED STILL OPEN IN CODE.** `playthroughs/e2e/lib/studio-builder-page.ts:120` still matches the empty section scaffolding at +2.1 s, before the LLM draft populates it. **The Playthrough passes without the generation completing** — a false green inside the suite the release headline depends on.
- [ ] **[blocker] `BIND_HOST` / `D-M255-7` — unaccounted, and there is no later destination.** Routed at the M255 close, re-applied at M256, recorded In-scope at M258, never worked; `up-injected.sh:146` unchanged. M258 was the last milestone. Carries a disclosed safety consequence (`safety.md` §3.1 LAN exposure). Direct cause of the batch-gate skip below.
- [ ] **[blocker] Three items held under `LAND-NEXT → release close`; at release close that destination does not exist.** `F2` (ptvalidate unwired — verified: only occurrence outside its own package is a comment at `run-playthroughs.sh:370`), `PROFILE-M257-provisional-fields` (verified: no `provisional_fields` contract in the loader; provisional in 2 of 3 profiles, prose-only), and the literal ratchets.
- [ ] **[blocker] Zero `RELEASE-SCOPE-DEFER:` decisions exist release-wide, and `roadmap-vision.md` has no v2.8 section** (last updated 2026-07-23 — four days *before* v2.8 was designed). Nothing can escape-hatch until that destination exists.
- [ ] **[must-fix] M256's three `KEEP-DEFERRED-WITH-SIGNOFF` items were never signed** and appear nowhere downstream: `PERF-M256-parallel-lane`, `PT-M257-self-evaluation`, `PT-M257-talk-to-data`. A user signature was requested at a milestone close and never obtained.
- [ ] **[must-fix] The release's headline promise has no owner in its own terms.** `roadmap.md:140` promises **666 s → ≤360 s**. Achieved: **286.99 s on `macmini` against `macmini`'s own 449.51 s** — a 36.2 % cut, not 46 %. Legitimate on the gate's re-pointed terms; but at release scope the `666 → 360` framing is a two-host comparison this release forbids everywhere else. Re-word to the achieved form or record as not-measured-as-promised.
- [ ] **[must-fix] Four use cases re-reserved to M206/M207 — which are not milestones.** `roadmap-vision.md:312-318` explicitly says these were re-reserved across five consecutive releases and **"do not re-reserve them a sixth time."** They were.
- [ ] **[should-fix] M257x carry-forward clusters 1, 3, 5 + cluster 4's surviving half + the 215-token block fate have no row in M258's close.** Cluster 1 is the production-bucket pointer — graded by M257x as its highest-stakes open item.

## Code Quality

- [ ] **[must-fix] "A fence satisfied by its own comment" — third instance, LIVE.** `playthroughs/manifest/batch_gate_test.go:758` reads the raw script body while its sibling four lines below (`:762`) uses the exec-scoped reader. `resolve_stack_dir` occurs in `restore-presenter-world.sh` at `:83` (executed) **and** `:261` (comment) — delete the executed call and the assert stays GREEN. Same defect `dc31efc` fixed elsewhere in the same pass.
- [ ] **[must-fix] `clone_drift_guard` is RED at the release tip** — 18 findings, 3 advanced clones, **124 citing sites**: `ant-academy` +2 (30 sites), `next-web-app` +59 (44 sites), `rosetta-extensions` +20 (50 sites).
- [ ] **[should-fix] The batch gate never runs on the default `/demo-up` path.** `batch-gate.sh:156-173` skips when `STACK_PUBLIC_HOST` is set; remote reach is **default-on**. On any tailnet box a bare `/demo-up N` yields `verdict: skipped` — *"UP, and every journey verified"* does not hold in the mode a presenter uses. The skip is correct; the missing fix is `BIND_HOST` above.
- [ ] **[should-fix] Both teardowns print recovery advice that re-creates the leak M258 closed.** `rosetta-demo:443` and `dev-stack:502` print `docker rm -f …`; the *executed* sweeps were fixed to `-fv`, the *printed* ones were not.
- [ ] **[should-fix] `buildbench.reclaim()` returns two exit codes nothing reads** (`buildbench.py:1243`); a failed prune and a zero-eviction prune are the same value in the only consumed field.
- [ ] **[should-fix] Three stale cross-file anchors in M258's newest tooling** (`batch-gate.sh:207`, `:182-189`; `restore-presenter-world.sh:102`). **rext-internal comment anchors are outside every guard's reach** — that gap is the finding as much as the drifts.
- [ ] **[should-fix] Dev/demo teardown asymmetry with a justification that describes an impossible failure** (`dev-stack:459-490` argues about compose files that call never loads).
- [ ] **[should-fix] Two of three Dockerfiles invalidate their install layer on every source change** (`next-web`, `hiring`: `COPY . .` → `pnpm install`); `studio-desk.Dockerfile` in the same directory does it correctly. UI-tier builds are 65.5 % of the cycle.
- [ ] **[nice-to-have]** `batch-gate.sh:107` `|| true` lets a previous run's verdict survive · `--out` parsed but never passed (`:60,92`) · `next-web`/`hiring` Dockerfile twins unfenced · dead `VITE_*` ARGs in the studio runner stage · `run-latency.sh:171` prints "green gate: OK" on ungated branches · a fourth re-implementation of the pair-count rules · `CLAUDE.md` off-by-one on `app/.gitignore`.

## Documentation

- [ ] **[must-fix] The content-story denominator correction reached the tooling and none of the corpus.** `content-denominator.json` pins **45**; `CLAUDE.md:467,469,470,472`, `coverage-protocol.md:943`, `content-stories-spec.md:157`, `playthroughs.md:1513` still assert 47/49 — **seven sites**, plus `content-stories.spec.ts:35`. A number whose source of truth is a rext JSON pin and whose mirrors are corpus prose is **covered by no guard in the family**.
- [ ] **[verified clean, recorded to stop re-derivation]** A 33-file grep for "live sentinel" phrasing is a **false positive**. `CLAUDE.md` and `corpus/README.md` are both correct — they state the floor is **two** and record their own prior wordings with retraction points. *A token census finds a wrong value, never an absent one.*

## Tests & Benchmarks

- [ ] **[must-fix] TWO literal-ceiling ratchets are breached, not one.** `DOCSTRING` **249 vs 240** (known, recorded). `TEST_MODULE` **662 vs 653** — **recorded in no milestone's `metrics.json`.** Whether the +9 predates v2.8 or grew inside it **cannot be stated from the artifacts**; one `git archive` extract at the v2.7 close tag would settle it. Both owed an explicit fate.
- [ ] **[must-fix] The benchmark condition is a REFUSAL, not a pass.** p50 **840.01 s** with **all 3 reps instrument-rejected** (peak load1 40/75/52 vs a limit of 10). 401.60 s is a **projection**; the 290 s warm-cache cycle was deliberately not banked. **Benchmark regression is not gradeable at this close.**
- [ ] **[should-fix] The full `stack-core` sweep has never completed** — reached 892 of 2419 at ~2 tests/min, not deadlocked; the cost concentrates in corpus-walking fences. **13 failures had accrued by test 892 and the census module's own 11 sort after the slow region — the full-sweep total is ≥24 and has never been established.**
- [ ] **[should-fix] Python dependencies are unpinned and unhashed** — no manifest anywhere in the tracked tree. The surface is genuinely two packages, but *small is an argument, not a measurement*.
- [ ] **[should-fix] v2.7 is the only release from v1.00 onward with no `metrics.json`**, and its two surviving artifacts disagree with its own `release-review.md`.
- [ ] **[watch] `seeders` coverage has fallen three releases running:** 96.1 → 95.7 → 94.4 (−1.7pp cumulative; each step in tolerance).

## Supply Chain

- [ ] **[should-fix] `GO-2026-5970` / `CVE-2026-56852` in `golang.org/x/text` is CALLED** (symbol-reachable, not merely present) in `stack-seeding` @ v0.37.0 and `stack-snapshot` @ v0.29.0. Fixed in v0.39.0 — one line per module. Neither the Go vuln DB nor OSV publishes a CVSS; the MODERATE grading is the reviewer's, on the basis that inputs are operator-supplied local URLs and DSNs.
- [ ] **[info]** `GO-2026-5942` in `x/net` — present, **not called**. npm: **0 vulnerabilities**. Licenses: **zero copyleft**. Vendored deps: none.

## Decision Consolidation

- [ ] **[must-fix] M258's `carry-forward.md:52-53` re-asserts a claim M257's close retracted 14 hours earlier** — that `hostprofiles/` lacks a profile for this host. `macmini.json` was added at M257 iter-04 and its baseline filled at iter-08. **The release's most-repeated failure shape, landing in the artifact release-close reads first.**
- [ ] **[should-fix] "A routing is not a routing until the target's own doc says so" fired at least three times in one release** — `BIND_HOST`, M257x's carry-forward (zero hits at destination), M255's four items to M257. Each caught by the *next* close, never the routing one. **This has earned a mechanical check, not another prose restatement.**
- [ ] **[should-fix] "A fence satisfied by its own comment" is a named recurring class with three instances and no detector.** The `shellInvocationLines` remedy already exists; what is missing is a guard asserting every fence reading a script body for an *executed* token routes through it.

---

## Not covered — stated so no one infers coverage that was not had

- **`stack-core/tests`** (94 files, the largest suite) — never completed; **no verdict**, and its own family banner warns a green guard verdict says nothing about it.
- **11 of 36 guards NOT-RUN** (need `--platform` / `--range` / `--ledger`).
- **65 of 424 Playwright specs not executed** — they need a cold reset-to-seed stack, and `demo-4` is the user's.
- **`stack-seeding` (89 files) and `stack-snapshot` Go sources** — `go vet` only, no line-level read.
- All timings above were taken on a **contended** host with three review agents running concurrently. **No duration here is offered as a baseline.**
