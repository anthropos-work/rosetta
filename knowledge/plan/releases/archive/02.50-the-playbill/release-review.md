# Release Review: v2.5 "the playbill"

**Date:** 2026-07-20
**Milestones:** M229, M230, M231, M232, M233, M234, M235, M236 (8, all `archived`)
**Release diff:** rosetta `main..release/02.50-the-playbill` — 85 commits / 159 files (~all markdown);
rosetta-extensions `1d97861..main` — 34 commits / 94 files / +20,354 −108

## Gate verdicts

| Phase | Gate | Verdict |
|---|---|---|
| 0 — supply chain | — | **GREEN** — 0 CVEs (Go, npm), 0 GPL/AGPL, lockfile committed `b855987` |
| 1 — release scope | — | **GREEN** — 11 of 12 in-release Fate-3 routings verified landed |
| 1b — deferral re-audit | blocking | **RED** — 22 items · 8 aged-out (7 newly detected) · 3 unhomed · 5 blocking |
| 2 — code quality | — | 4 must-fix · 15 should-fix · 10 nice-to-have; toolchain clean |
| 3 — documentation | — | 11 must-fix · 6 should-fix |
| 3b — KB consolidation | blocking | **3 blocking** structural-debt items |
| 4 — tests | — | Go GREEN (2461/0) · TS GREEN (196/0) · Python YELLOW (8 fail, 2 flake) |
| 4b — metrics regression | blocking | **RED** — `flake_count` 0 → 2 |
| 5 — decisions | — | 9 unblended · 7 conflicting · 5 release-level patterns |

**Three blocking gates are RED.** The release cannot tag until each converts to Fate 1 or receives
explicit escape-hatch sign-off with a per-item "why Fate 1/2/3 failed concretely" rationale.

## The release's own thesis, recurring

v2.5's most transferable finding is *a check can report success while proving nothing* — nine instances
across M235–M236 (a scrub that removed zero names with every test green; three tests asserting the defect
they should catch; a `0/0` aggregator exiting 0; a suite passing by collecting 0 tests; a grader with no
negative tests; a regression test that was a self-consistent tautology).

**This review found the class alive inside the close that named it.** The headline `29/29` has three
independent integrity problems (§CQ-1, §CQ-2, §CQ-6, §T-4). Four more tautological tests and two
vacuous-pass tests survive (§CQ-5, §CQ-7). And the propagation rule the release itself authored —
*prose does not propagate; only a shared definition or an executable fence does* — was violated by the
correction sweep that was supposed to enforce it (§D-1..§D-6, one root cause).

---

## Scope

- [~] **[must-fix] S-1** — **DEFERRED → M237** (KEEP-DEFERRED-WITH-SIGNOFF; argued down from the audit's
      LAND-NOW: landing the fix without a live re-prove would ship an *unverified fix to a render defect*,
      the exact class this release is named for). Recorded in `release-deferrals.md`.
      ORIGINAL FINDING: `DEF-M230-03` → `ACADEMY-M236-iter08-public-catalog-twin` is the release's only
      Fate-3-UNDELIVERED item: anonymous academy routes (`/library`, `/free`) render 0 cards. Chain
      M230→M235→M236→release close; carried in the `In:` list at every hop, so disclosed non-delivery,
      not a silent drop. Aged out across 3 milestones.
- [x] **[must-fix] S-2** — **LANDED, but the finding was partly WRONG.** KB-8 (`internal/labsession`)
      and KB-1 landed. **KB-2 REFUTED**: 8400/8401 is correct for anything driven through `platform`
      (`docker-compose.yml`); 8080/8081 is only the unset-env fallback (`cmd/root.go:77-78`) — both true
      in context, so both are now documented rather than one "corrected". **KB-6 "~6 docs wider" REFUTED**
      (two agents independently): the caller claim exists at exactly ONE site; every other hit is a plain
      enumeration. The core KB-6 claim resolved *harder* than filed — roadrunner has **NO caller at all**
      (jobsimulation's codebase has one hit, a comment reading "formerly the standalone 'roadrunner'
      service"; it is now a direct Judge0 HTTP client), so `ROADRUNNER_RPC_ADDR` is dead config and the
      service is marked orphaned pending retirement. KB-5 moot confirmed. ORIGINAL FINDING: M231's **KB-2 / KB-6 / KB-8** are *unhomed* — classed "Fate-3, arch-doc pass",
      which is not a milestone. In no `In:` list, no carry-forward, no backlog. All three verified still
      stale today: `corpus/services/jobsimulation.md:17,44` (ports 8400/8401, actual 8080/8081);
      `corpus/services/roadrunner.md:63`; `corpus/services/backend.md:25,109` (`internal/labsession`,
      actual `internal/labs/session`). KB-6 is ~6 docs wider than M231 recorded. KB-5 is moot.
- [x] **[must-fix] S-3** **v2.4 shipped with no release-scope deferral audit at all** —
      `releases/archive/02.40-casting-call/` has four milestone-scope audit dirs and none at release root.
      Every item destined for "the v2.4 release close" fired unchecked. This is the structural cause of
      the 8 aged-out triggers and the first release-scope audit since v2.3.
- [~] **[must-fix] S-4** — **fresh dated KEEP-DEFERRED-WITH-SIGNOFF → M238**, with a drop-expiry.
      ORIGINAL FINDING: `DEF-M235-03` / M204 **assign-WRITE** — ~10 routings across 5 releases. Declared
      destination was the v2.4 close, which fired 2026-07-18 with no fate taken; it then silently
      re-anchored on v2.5. **Both v2.5 milestone audits recorded it as "correctly routed"**
      (`m236/decisions.md:210`, `m236/carry-forward.md:91`). Verbatim AGED_OUT; warrants the same footing
      the standing test-debt carry got.
- [~] **[must-fix] S-5** — **KEEP-DEFERRED → M237 WITH TEETH**: M237 must *test* the never-tested
      "self-resolves in the default flow" claim, else DROP. ORIGINAL FINDING: `DEF-M226-01` (pre-bind serve reap) **aged out twice** — M228 fired 2026-07-18,
      then M236 (the next prove-on-VM) fired 2026-07-20. `state.md:137-139` still names M228, stale for
      two releases. Its standing justification ("self-resolves in the default flow") has never been tested.
- [~] **[should-fix] S-6** — **KEEP-DEFERRED → M237, reclassified `DRIFT_DEFER`** (the "needs live infra"
      rationale is now false). ORIGINAL FINDING: Four v2.3 tail carries (F4 · `BURNIN-M221-dev-public-host` · `F-M220-4` ·
      `PROBE-M218-c3-rerun`) were signed off KEEP-DEFERRED → v2.4; v2.4 shipped without them.
      `roadmap-vision.md` still reads "v2.4: investigate…" and has **no v2.5 carry section**. Their
      "needs live infra" rationale is now false — billion is up and ran 10 live iters this release.
      Reclassify as `DRIFT_DEFER`.
- [x] **[should-fix] S-7** — **DISCHARGE RECORDED**; `F4` retired as an id, remainder is S-1.
      ORIGINAL FINDING: F4 is now *partly* discharged — the signed-in half landed at M230/M236; only
      the anonymous half survives, and it **is** S-1. `state.md:140-142` lists F4 whole, over-stating debt.
- [~] **[should-fix] S-8/S-9** — one item, not two (the double-identity is *why* it vanished from every
      ledger). KEEP-DEFERRED → M237 as an explicit added assertion. Class named. ORIGINAL FINDING: M232's interview plan-section render fidelity was closed Fate-2 "covered by
      M235's gate"; M235 closed `closed-incomplete` and that gate never ran. Absent from M235's
      carry-forward and from `DEF-M235-01..04`. Where the gate did run (M236) it measured a weaker
      property. **"Confirmed-covered" is the deferral form that leaves no ledger entry** — name the class.
- [x] **[should-fix] S-9** M232 routed one item to two destinations (`m232/spec-notes.md:40` → M234;
      `m232/progress.md:49` → M235). Neither ran it.
- [~] **[should-fix] S-10** — **KEEP-DEFERRED-WITH-SIGNOFF**, now written down by id with a
      drop-if-it-survives-another-release condition. ORIGINAL FINDING: `F11` / `DEF-M215-03(a)` routed to "standing backlog"; never arrived. Absent
      from `state.md:120-147`.

## Code Quality

- [x] **[must-fix] CQ-1** `shapeFor` misgrades a scored assessment as an interview — **a live false PASS
      in the shipped canonical manifest.** `stack-verify/e2e/lib/content-result-page.ts:352-353` builds a
      keyword haystack from `sim_type + key + path` then matches `/interview/`. `asmt-voice-pass` is an
      ASSESSMENT whose sim slug contains "interview", so it is graded by the `player-interview` branch
      (`:192-198`) — **no length floor, no score check, no feedback check**. That pair cannot verify the
      scored report rendered. The file's own header warns against exactly this. Fix: discriminate on
      `sim_type === 'SIMULATION_TYPE_INTERVIEW'` before any keyword fallback.
- [x] **[must-fix] CQ-2** **The four new `*.unit.spec.ts` files — including M236's route-contract pin —
      are executed by no runner.** Each runner names exactly one live spec
      (`run-content-stories.sh:121`, `run-coverage.sh:139`, `run-latency.sh:154`, `run-hiring-render.sh:78`);
      `grep -rn 'unit.spec' run-*.sh package.json ../tests/*.py` returns nothing.
      `stack-verify/tests/test_e2e_collection_integrity.py:38` *collects* them (they count toward
      `MIN_TESTS=120`) but never runs them. M236's harden raised the collection floor without raising the
      assertion floor. **A pin nobody runs is not a pin.**
- [x] **[must-fix] CQ-3** **The org-name scrub arm has never fired, fails open, and sits outside the leak
      gate.** `stack-seeding/cmd/content-capture/main.go:98-101` swallows the query error
      (`err == nil && org != ""`) while the sibling owner-identity query at `:110-115` returns
      `fmt.Errorf("owner identity: %w", err)` — two PII sources, opposite failure postures, same function.
      The org is inserted into `repl` directly, bypassing `addName`/`nameTokens`, so
      `leakCheck`/`SurvivingToken` (`:337`) never checks whether it survived. Measured across all 13
      fixtures: **0 `<<ORG>>` placeholders against 840 `<<ACTOR_i>>`**. Either the source org never appears
      in the free text, or it was copied verbatim — **the pipeline cannot tell you which.** This is real
      customer session content and the one PII class the gate structurally cannot catch. Given the
      release's central decision (accepted residual re-identification risk), close this before tag.
- [x] **[must-fix] CQ-4** The auditable disclosure artifact contradicts itself on PII posture.
      `stack-seeding/manifest/manifest.go:87` — "provably PII-free — every free-text facet is synthesized
      at seed time, not stored"; `:239-240` — "NOT guaranteed clean … residual re-identification risk is
      real and ACCEPTED". The first is stale pre-`c575c0e` prose (the M232 COPY-real rework superseded it).
      Same stale claims at `contentsession/contentsession.go:7,66-67,96`. The new code is honest
      everywhere; only the surface an auditor reads asserts the retracted claim.
- [x] **[should-fix] CQ-5** Four more tautological assertions — direct siblings of the one M236 found.
      `seeders/roster_test.go:333-339` and `:354-356` (`want := len(BuildRoster(bp).Identities)` vs a
      `WriteRoster` that *is* that expression — `len(x) != len(x)`; the `want <= 5` floor is the only
      teeth); `content_manifest_test.go:281-286` (same shape, **no floor**);
      `content_stories_test.go:163-171` (computes `want := fillPlaceholders(...)` with the seeder's own
      transform — substituting the wrong actor's name still passes). M236's close added a proper golden
      for `membershipUUID` (`:316-329`) and didn't sweep the siblings.
- [x] **[should-fix] CQ-6** `EXPECTED_PAIRS` is not an external expectation —
      `run-content-stories.sh:88-106` computes it from the same served manifest the sweep reads, via a
      fourth inline reimplementation of `buildPairs`' counting rules. If the Go projection dropped 3
      sessions the sweep reports a flawless 26/26. The only literal pin of `29` is
      `content-route-contract.unit.spec.ts:130` — which per CQ-2 never runs.
- [x] **[should-fix] CQ-7** Two tests in `test_green_gate_age.py` can pass having asserted nothing
      (`:97-101`, `:113-115` — `if not got: continue` with no non-empty precondition). The sibling
      `test_epoch_is_identical_in_every_timezone` guards this correctly at `:83`.
- [x] **[should-fix] CQ-8** Manager-route fall-through unpinned. `content-result-page.ts:343` returns
      `manager-dashboard` as a fall-through and `content-route-contract.unit.spec.ts:102-104` asserts
      exactly that for every non-interview manager route — so renaming
      `/enterprise/activity-dashboard/ai-simulations/` in Go stays green. No manager twin of the `:69`
      anti-fall-through assertion. `EXPECTED_PLAYER_SHAPE:42` allows `'player-scored|player-interview'`
      for `simulation`, which is what lets CQ-1 through.
- [x] **[should-fix] CQ-9** Seeder/projection drop asymmetry: `resolveNonSimSession` fail-closes on empty
      `SkillPathID` (`content_nonsim.go:258-260`); `ContentStoryNonSimSeeder.Seed` (`:348-350`) has no
      equivalent guard, so it would write a row the manifest drops. Unreachable today; the registry is
      exactly what a future milestone edits.
- [x] **[should-fix] CQ-10** Asserted-but-unenforced actor-name uniqueness —
      `content_stories_write.go:171-184` claims distinctness satisfies `actors.unique(username, session)`,
      but `syntheticStakeholderNames` has 6 entries indexed `% 6`, so slots `i` and `i+6` collide and
      `CopyRowsIdempotent` guards only `ON CONFLICT (id)`. A session with ≥8 actors errors the whole seed.
      Latent (max today 4). The same PR guards this hazard class deliberately elsewhere.
- [x] **[should-fix] CQ-11** Uncalibrated schema assumption that hard-fails the seed —
      `content_nonsim.go:447-457`'s comment still calls the `lab_sessions` NOT-NULL column set "an M236
      live-seed-calibration item", M236 being the milestone now closing, with no DDL evidence in-repo. A
      wrong column set aborts the entire seed (`:380-383`), inverting the file's stated degrade-to-zero-rows
      posture (`:300-302`).
- [x] **[should-fix] CQ-12** Dead manager branch — M236 set `skill-path-legacy`'s `managerKind` to `""`
      (`content_manifest.go:147`), making `content_nonsim.go:264-269` unreachable. Defensible to keep
      (restore when the platform ships the surface) but needs a comment saying so. Stale prose at
      `content_manifest.go:330`.
- [x] **[should-fix] CQ-13** — **KEPT AND STRENGTHENED, not removed** (argued): the 105-vs-12 gap *is* the
      disclosure (a demo shows a bounded slice of a longer real session). Renamed to `Source*Count`, and a
      new `ValidateAgainstContent` now cross-checks the pins against captured content. ORIGINAL FINDING: Obsolete + actively wrong fixture fields —
      `contentsession.go:95-98`'s `ActorCount`/`InteractionCount` are validated (`:209`) and read by nothing
      post-rework, and misreport: `asmt-doc-fail` declares 105 interactions against a `--transcript-cap`
      default of 12; also `train-chat-fail` 51→12, `asmt-voice-pass-2` 44→12, `asmt-code-pass` 25→12.
- [x] **[should-fix] CQ-14** `test_academy_fs_published.py:152-172` retypes the strip expression inline
      rather than extracting it from `m["replacement"]`; the drift guard (`:174-182`) only substring-checks.
- [x] **[should-fix] CQ-15** `content-result-page.unit.spec.ts:253` pins a known-bad behaviour (asserts
      `player-scored` *does* false-pass an 11k-char skill path), so tightening `player-scored` — e.g.
      actually using `hasScore`, computed at `:204` then never used in the verdict — turns it red for an
      improvement.
- [x] **[should-fix] CQ-16** Stale denominator in the file a triager reads first —
      `content-stories.spec.ts:17-21,71,91` still teaches "31 landable"; `content-result-page.ts:100` same.
      `lib/content-pairs.ts:12` records the correction.
- [x] **[should-fix] CQ-17** Playwright harness resources — `tests/aggregate-content.unit.spec.ts:36,145,164,180,197`
      has five `mkdtempSync` calls and zero cleanup (0 hits for `rmSync|afterEach|finally`).
      `content-result-page.unit.spec.ts:414-421` hot-spins ~8 real seconds inside a "unit" test.
      *(Credit: the sweep itself is resource-clean — zero listeners, fixture-owned pages, `mktemp` + `EXIT`
      trap in all four runners.)*
- [x] **[should-fix] CQ-18** `content-pairs.ts:111` — the manager branch skips the fail-closed base-URL
      guard the player branch applies at `:83-85`. Fails loudly rather than false-passing; low impact.
- [x] **[should-fix] CQ-19** `readableLength` contradicts its own comment eight lines above —
      `content-result-page.ts:170-175` explains that `main + body` "silently HALVES the effective floor",
      and `:151` returns exactly `main.length + body.length`. `MIN_SETTLED_CHARS = 200` is a ~100-char floor.
- [x] **[nice-to-have] CQ-20..29** `gofmt -l` on 3 release-introduced files (comment alignment only);
      `helpers.go:82-107` orphaned godoc; `content_stories.go:174-175` backwards flush-order comment;
      five `err != pgx.ErrNoRows` → `errors.Is` (`main.go:113,165,294,310,327`); `run-coverage.sh:168`
      unclosed handle; `cockpit.py:711` vs `:631` empty-products 200; `cleanliness_test.go:102,153`
      no-precondition range loops; `test_tooling.py:1286-1301` self-including doc-truth count;
      skipped-by-default behavioural halves; `content-stories.spec.ts:74` append-only `pairs.jsonl`.

## Documentation

**Root cause of D-1..D-6:** commit `302a32e` ("correct the claims M236 refuted") touched
`corpus/ops/**` + `CLAUDE.md` + `.claude/skills/**` — **16 files, zero under `corpus/services/`**.
The service-doc tier was never swept. Fix as one scoped sweep.

- [x] **[must-fix] D-1 — BLOCKING** `corpus/ops/safety.md` — **four unqualified absolutes that §3.8 makes
      false.** `:6-7` ("It never reads private/customer data"); `:36-38` ("no customer's private data can
      be copied"); `:65` ("A snapshot **capture** is the only operation that reads production" — but
      `cmd/content-capture` is a second, deliberately customer-scoped prod read outside `AssertPublicOnly`);
      `:496-501` §3.3 ("**There is nothing behind the door** … not 'should not', *cannot*"). The v2.5
      exception was threaded into `:56-59`, §3.7 `:678` and `:639` but **not** the top-of-file guarantee or
      §3.3 — and `:496-501` is the load-bearing argument justifying default-on remote reach (D-DESIGN-3).
      **The safety contract cannot ship asserting a guarantee its own §3.8 retracts.**
- [x] **[must-fix] D-2 — BLOCKING** `knowledge/plan/state.md` — frontmatter says merged/unblocked; body
      says BLOCKED in 8 places (`:12, :23, :26, :29, :47, :111, :121, :149`). `git show abdbf25` changed
      **only** the two frontmatter fields. Currently 14,622 / 15,360 bytes (95.2%) — the rewrite must not grow it.
- [x] **[must-fix] D-3 — BLOCKING** `knowledge/plan/context.md` is two releases stale — `:7` and `:85`
      "Status (2026-07-15): v2.4 IN DEVELOPMENT"; `:97` "Next: build-milestone → M222"; `:60` "the ACTIVE
      v2.3 dirs". The plan cluster's orientation doc.
- [x] **[must-fix] D-4** `corpus/ops/demo/session-clone-spec.md:66-68` says **9** pinned sessions; the
      fixture has **13** (assessment 7 = 3 voice / 2 code / 2 doc, training 2, hiring 2, interview 2).
      Grew at rext `590082a` (M235). This is the only doc stating the pin inventory; it breaks the
      18-session chain at `coverage-protocol.md:804` and the `13/13` at `:442`.
- [x] **[must-fix] D-5** — **CORRECTED, but the finding was IMPRECISE and was not propagated as filed.**
      Two different routes: the *cohort* scoreboard genuinely renders and genuinely reads the mirror — that
      guidance is correct and was kept. Only the *per-member* drill-down is the dead "Coming soon" surface.
      (M236's denominator drop remains correct — those pairs were the per-member route.) ORIGINAL FINDING: `corpus/services/skillpath.md:84-93` still describes the skill-path manager
      surface as working ("the mirror row must be co-written"). M236 iter-07 proved it renders literal
      "Coming soon". A reader here would build a mirror seeder for a page that cannot render.
- [x] **[must-fix] D-6** `corpus/services/ant-academy.md:63` still routes the academy through
      `app/cmd/academy-seed` unqualified; M236 iter-08 established it is moot on a demo.
- [x] **[must-fix] D-7** `corpus/ops/demo/content-stories-routes.md` — the academy refutation is struck
      through in **1 of 5** places. Corrected at `:310`; NOT at `:11-12` (verdict headline "Academy is IN"),
      `:24-25` (the *For PMs* paragraph), `:49` (route matrix), `:293` (§6 heading). The correction block
      sits at `:354-375`, after every one of them — a PM reading only the PM section gets the exactly-wrong
      answer.
- [x] **[must-fix] D-8** `corpus/ops/demo/content-stories-spec.md` — two internal contradictions: `:287`
      ("no academy session; lights up when M235 adds the fixture") vs `:101`/`:233` (live since M236
      iter-08); `:313` ("proving every CTA lands is M235") vs `:241` ("is M236") — `:241` is correct.
- [x] **[must-fix] D-9** `CLAUDE.md:312` and `:427` reference `corpus/ops/update_checklist.md` — **the file
      does not exist.** Two dead links in the file every new agent reads first.
- [x] **[must-fix] D-10** `corpus/ops/demo/tailscale-serve.md` — `:319` runbook step 6 says the cockpit is
      "deliberately **not** fronted by `tailscale serve`" vs `:479-486` "the last plain-HTTP surface is
      gone". An operator following the runbook verbatim uses the wrong scheme on the demo's entry point.
      Patch-tail count: heading `:530` says THREE, `:539`/`:84-86`/`:648-650` still say two.
- [x] **[must-fix] D-11** `corpus/services/backend.md:25,109` — `internal/labsession`, actual
      `internal/labs/session`. `content-stories-routes.md:291` **documents the error rather than fixing it**.
      Shipping a doc that annotates another doc's bug is a close-phase anti-pattern.
- [x] **[must-fix] D-12** Playthrough count stated three ways: `playthroughs.md:105` **15** (authoritative,
      v2.4 M225 added `pt-hiring-recruiter-compare`); `CLAUDE.md:299` **14**; `demo/README.md:205` **10**.
- [x] **[should-fix] D-13** `latency-budget.md` never records the release's own gate outcome — the
      measured-baseline table stops at M226. Hero p95 1.22 / 1.51 s appears corpus-wide only at
      `tailscale-serve.md:682`. Also `:159-160` lists two vantages in the harness recipe; the gate at
      `:28-30` says three.
- [x] **[should-fix] D-14** `coverage-protocol.md` — the M236 content-stories sweep (`:702-921`, **24% of
      the file**) is bolted on after the terminal `## Related` block at `:696`; the header `:1-9` still
      scopes the doc to v1.10 M42 and never mentions the second sweep.
- [x] **[should-fix] D-15** `safety.md:512-516` §3.4 residual #1 describes exactly what §3.8 deliberately
      did, with no forward pointer. The two sections read as unaware of each other.
- [x] **[should-fix] D-16** Milestone-attribution drift — `content-stories-routes.md:148` ("M232 …
      forthcoming"; it shipped); `session-clone-spec.md:197` and `content-stories-spec.md:286` credit M235
      for the render proof that M236 did.
- [x] **[should-fix] D-17** `CLAUDE.md:145-157` — the **normative** rext policy still reads "built and
      tested in the authoring copy, **tagged**, then consumed per-stack" with no origin-push step.
      `verification.md:330-351` names this exact wording as the root cause of M236's lost first iteration
      ("tagging is not publishing"). The fix landed in the doc-index blurb at `:294` but not in the policy
      text 150 lines above, where an agent actually reads it.
- [x] **[should-fix] D-18** `knowledge/plan/roadmap.md:108` — the v2.4 row asserts both "🟢 CODE-COMPLETE …
      awaiting close-release" and "✅ SHIPPED 2026-07-18 (tag v2.4)" in the same cell; `:95` still says v2.1
      "is now IN DEVELOPMENT".

### Knowledge-base consolidation (Phase 3b)

- [x] **[must-fix] KB-A — BLOCKING** `F-M236-CLOSE-1` (demo clone staleness) has **no corpus anchor**, and
      three docs actively *reassure* about the broken mechanism: `corpus/ops/rosetta_demo.md:70` describes
      skip-if-present approvingly as "the platform's own idempotent clone loop" with no staleness
      consequence; `:71` presents `clones.lock.json` as `{ref,sha}` provenance without noting `ref` is the
      literal `"HEAD"` for every detached clone (so it structurally cannot distinguish pinned from stale);
      `frontend-tier.md:444` documents a *different* skip-if-present failure as **fixed**, so the class
      reads closed while the staleness half is wide open. Measured: `app` **249** commits behind `main`,
      `next-web-app` **202**, identically on both boxes. **This was the user-reported stale left menu and
      the upstream generator of the entire pin-drift-looking class.** → new `## Clone freshness` section in
      `rosetta_demo.md` after `:71`, cross-linked from `verification.md`'s pre-flight rung.
- [x] **[must-fix] KB-B — BLOCKING** `F-M236-CLOSE-2` — the R1 pristine-revert rung covers **3 of ~15**
      patch manifests, and `demopatch-spec.md:52` presents G5 content-anchored self-revert as complete with
      no mention of the gap. Both boxes measured carrying leftovers in disjoint sets (5 in local
      `next-web-app`, 2 in billion's `ant-academy`). `coverage-protocol.md:315` documents `--force-pristine`
      as a working safety net.
- [x] **[should-fix] KB-C** `knowledge/plan/roadmap.md` = 2079 lines / 203 KB. Split precedent exists
      (`roadmap-legacy.md`). Natural seam: cut shipped v2.0–v2.4 into `roadmap-archive-v2.0-v2.4.md` at the
      `## Active — v2.5` boundary (`:125`).
- [x] **[should-fix] KB-D** `metrics-history.md` main table is missing **v2.0, v2.2 and v2.4** rows (v2.4
      sits in a separate 5-column second table at the file bottom, contradicting the file's "newest first"
      convention). `state.md:143` under-reports this as "v2.0 + v2.2".
- [x] **[should-fix] KB-E** No index for `corpus/services/` (24 docs) or `corpus/tools/` (2 docs); combined
      with `CLAUDE.md:325-329` naming only 2 of 24, **there is no enumerated list of service docs anywhere.**
- [x] **[should-fix] KB-F** Ten corpus docs `CLAUDE.md` has never mentioned. Largest:
      `corpus/ops/directus-local.md` (316 lines, 28 inbound refs, owns the per-stack Directus model
      `CLAUDE.md:281` depends on). Also the staging family — and `corpus/ops/README.md:11` calls
      `staging-bringup.md` "New engineer (or AI agent) joining the team — start here", so the doc that
      self-declares as the agent entry point is absent from the agent entry file.
- [x] **[should-fix] KB-G** — orphans given real references. **`intelligence.md` NOT merged** (argued): the
      `skiller.md` precedent applies — a reader following an old link should land on an explanation of the
      fate, not a 404. ORIGINAL FINDING: Effective orphans (one inbound, a bare taxonomy row):
      `corpus/tools/anthropos-labs.md`, `services/chronos.md`, `customerio-sync.md`, `messenger.md`,
      `storage.md`. Merge candidate: `services/intelligence.md` (18-line tombstone) → `service_taxonomy.md`.
      **Do not merge `services/skiller.md`** — 55-line tombstone but 25 inbound refs, a working redirect.
- [x] **[should-fix] KB-H** `corpus/README.md` and `corpus/ops/README.md` mention none of the v2.5 trio
      while `ops/README.md` gives individual rows to `demo/frontend-tier.md` and `demo/tailscale-serve.md`.
      Reach is intact via explicit delegation — precedent inconsistency, not a broken chain. Pick one rule.
- [x] **[should-fix] KB-I** `releases/02.50-the-playbill/` lacks release-level artifacts its archived
      predecessor has (`design-notes.md`). Design provenance points at `.agentspace/scratch/…`, outside the KB.
- [x] **KB-J** *(landed this phase)* `dependencies.lock` (Phase 0, `b855987`) and `metrics.json` +
      `metrics-history.md` row (Phase 4b).

## Tests & Benchmarks

**Counts are from parsed JUnit XML, never piped stdout.**

| Stack | Result |
|---|---|
| Go | **GREEN** — 2461 testcases / 0 failed, 6 modules; reproducible func count **1976** |
| TypeScript | **GREEN** — 196 unit specs / 0 failed (e2e 127 + playthroughs 69) |
| Python | **YELLOW** — 1409 testcases, 8 deterministic failures, 2 flakes, 8 skipped, 5 sections |

- [x] **[must-fix] T-1 — BLOCKING (Phase 4b)** `flake_count` **0 → 2**. Per the regression rule any flake
      is blocking. Both are pre-existing test-infrastructure timing flakes, neither touches v2.5 code:
      `test_host_isolation.py::test_released_on_sigterm` (1 fail / 3 isolated runs — races a 2 s poll budget
      for the SIGTERM trap to unlink the lockfile; **not named in any prior ledger**) and
      `test_m220_mutation_battery.py::test_the_dev_fences_are_red_proven` (1 fail / 3 full-suite runs but
      6/6 green isolated — a nested-pytest self-check tripping under load). **Record correction:** M236
      logged these as "2 pre-existing *failures*"; they are flakes — measured 152/152 on two subsequent runs.
      Decide: test-side fix (retry/backoff on two poll budgets) or `KEEP-REGRESSION` with sign-off.
- [x] **[must-fix] T-2** **Standing-failure count: 8 observed** — exactly the documented clean-clone macOS
      set (4 M53 academy-link, 2 M218 overlay, 1 M215 browser-port list, 1 macOS-conditional purge). The 6
      clone-dependent failures did not reproduce, so **the `stack-demo` clone set is pristine**. M236's own
      `metrics.json` still records **716/14** — the dirty-clone reading taken before the re-baseline. Not
      propagated into the aggregate; the milestone artifact still needs correcting.
- [~] **[must-fix] T-3** — **DEFERRED → M237** (KEEP-DEFERRED-WITH-SIGNOFF). ORIGINAL FINDING: **39 live-browser specs unexecuted at close** (24 stack-verify + 15 playthroughs) —
      they need a running demo stack. The entire prove-by-render layer is unverified in this close.
- [~] **[must-fix] T-4 / CLOSE-D3** — **ESCAPE-HATCH, user-authorized 2026-07-20 → M237.** The release
      tags on the unit-proven number; the live re-prove is v2.6's opening work. Per-item why-Fate-1/2/3-failed
      rationale + acceptance condition in `release-deferrals.md` §A. **Fate 1 did NOT fail on capability —
      it failed on an explicit user choice, and is recorded as an escape hatch, not an impossibility.**
      ORIGINAL FINDING: **`CLOSE-D3` — the shipped harness has no live green.** The 29/29 headline was
      measured on `playbill-m236-hardened`; the close then fixed **10 harness defects**. Unit-proven, never
      re-run live. Tagging on a number whose measuring instrument changed afterwards is the D17 hazard this
      project has repeatedly named. Compounded by CQ-1 (one of the 29 graded by the wrong check) and CQ-6
      (the target count is self-derived).
- [x] **[should-fix] T-5** Coverage gaps: `cmd/content-capture` **28.7%** (`captureSession`/`main` are the
      DB-bound prod-read surface — acknowledged at M235); `cmd/stackseed` 64.6%; `cmd/datadna` 51.0%.
- [x] **[should-fix] T-6** 9 demo-stack pytest failures reproduce identically at the base revision
      `1d97861` (3 of 4 failing files untouched this release) — pre-existing, but the release shipped on a
      red suite.

**Metrics deltas (Phase 4b):** Go funcs 1879 → **1976** (+97) · TS collected 151 → **235** (+84) ·
seeders coverage 96.8% → **96.1%** (−0.7pp, within tolerance) · p95 1270 ms → 1220/1510 ms (flat, far
inside the 5000 ms gate) · supply chain GREEN → GREEN (0 new deps) · **flakes 0 → 2 (BLOCKER)**.

**Two count corrections recorded, not propagated:** v2.4's Go baseline is **1879** by the reproducible
git-grep method, not the 1902 M233 quoted non-reproducibly. v2.3's Python figures folded pytest *subtests*
into totals (stack-core 182 = 152 + 30), so the apparent 182 → 152 is a method artifact, not a decrease.

## Decision Consolidation

**Verified landed and coherent:** the data-controller acceptance of real-prod-session sourcing
(`safety.md:686-726` §3.8 ↔ `session-clone-spec.md:11-18, 87-102, 172-183`, clause-for-clause, no drift);
the 31 → 29 denominator correction (landed three times — `coverage-protocol.md:788-827`,
`content-stories-spec.md:100-110`, `content-stories-routes.md:352-360`, including the rule *"a resolver
existing is not a surface existing"*); the demo-reach supersession chain (D-DESIGN-1 → D-DESIGN-3 at
`safety.md:524-550`, correctly scoped, no doc asserts the retracted position).

- [x] **[must-fix] DC-1** **M236's gate re-scope never landed in knowledge** — the two decisions the user
      made. B2 (p95 hero-only) exists only in `m236/overview.md:38-39` + `decisions.md:126-134`;
      `latency-budget.md:27-30` still declares the gate as "**three** vantages" over "**5 consecutive cold
      reset-to-seed runs**". B1's rationale ("restricting reach is the VM + VPN's job, not the demo
      stack's") has **zero hits in `corpus/`**.
- [x] **[must-fix] DC-2** **The p95 gate contradicts itself, doc vs claim** — the doc requires 3 vantages /
      5 cold cycles; M236 claimed MET (`m236/metrics.json:34-39`, `overview.md:112`) on **2 vantages / 1
      cold cycle** (5 login runs within it). Neither side reconciles it. The sharpest live conflict in the
      release.
- [x] **[must-fix] DC-3** **`31` still stands as the machine-readable exit gate** —
      `m236/overview.md:10` (frontmatter `exit_gate`) and `m236/carry-forward.md:6` (`gate_target`) both say
      "all 31 landable pairs"; `:23` repeats it in prose; only `:25-33` corrects it. **Tooling reads the
      frontmatter.** (The corpus is correct at 29 throughout.)
- [x] **[must-fix] DC-4** **The standing-failure count is asserted three ways** — `state.md:121` = 14 ·
      `roadmap.md:397` = 19 · `m236/carry-forward.md:34-47` = 14 reproduced → 8 macOS / 7 Linux. Exactly the
      drift-under-a-fixed-label that CLOSE-D2 escalated.
- [x] **[must-fix] DC-5** **A dangling cross-reference to content that was never written** —
      `coverage-protocol.md:778` asserts "which is exactly the per-item fan-out signature
      `latency-budget.md` teaches"; that content (`m236/iter-06/decisions.md:38-41`) was never added.
- [x] **[must-fix] DC-6** **The login-shell finding never reached its declared home** —
      `m236/iter-03/decisions.md:19-22` routed it to `tailscale-serve.md`'s F-ledger with handler
      `DOC-M236-iterTBD-protocol-backfill`; the ledger at `:583-594` has no such entry, and `:124` (F2)
      still tells the operator to add `/usr/local/go/bin` to PATH — leading them into the exact false
      "Go is missing" diagnosis iter-03 disproved.
- [x] **[should-fix] DC-7** Cold-vs-warm p95 contrast unrecorded — warm 3.15 / 2.71 s appears nowhere in
      `corpus/`; the baseline table `latency-budget.md:85-102` still ends at M226. This contrast *is* the
      one place v2.5 demonstrates the doc's own "state the environment with every number" rule (`:127-129`).
- [x] **[should-fix] DC-8** *"A fix for a false-FAIL that creates a false-PASS is a net loss"*
      (`m236/decisions.md:170`) — not landed; no `net loss` string in `corpus/`. The one directional rule an
      author needs when tuning a grader.
- [x] **[should-fix] DC-9** **The knob count did not propagate — the same pattern, inside the close that
      named it.** `demo-up-defaults.md:37` says **27**, `CLAUDE.md:284` says **27 + 10 flags**, but
      `demo/README.md:141` still says "all **25** env knobs + 9 CLI flags" and
      `.claude/skills/demo-up/SKILL.md:4` still says "all **26**".
- [x] **[should-fix] DC-10** `latency-budget.md:21` cites a **bare `D-DESIGN-1`** — precisely what
      `safety.md:544` forbids ("the ids collide across releases; a bare reference resolves to the wrong
      decision"), and glosses it differently from `safety.md:544`. Same decision, two glosses, no canonical
      statement.

### Release-level patterns worth promoting to knowledge

- [x] **[should-fix] DC-P1** **The release's true thesis: "a check can report success while proving
      nothing."** Nine instances across M235–M236. `coverage-protocol.md:866-921` captures it well but frames
      it as a coverage-protocol concern. It is a **release-level epistemics finding**; its crisp form —
      *"ask of every layer that reports a number: what does it print when nothing happened?"* — deserves
      promotion above one doc.
- [x] **[should-fix] DC-P2** **"Prose does not propagate; only a shared definition or an executable fence
      does"** (`m236/retro.md:58-70`) — landed only as a parenthetical at `coverage-protocol.md:886`, and it
      recurred **during the close that named it** (DC-9, D-1..D-6). The release's most reliable predictor of
      future defects. A doc-truth guard already exists (`stack-core/demo_knob_guard.py`); the knob-count
      claims in `README.md` and the skill are not fenced by it.
- [x] **[should-fix] DC-P3** **"Offline-authored, never driven" as a defect class.** Four artifacts shipped
      wrong for exactly this reason, each defended by a green test: the manager route built from a user id;
      `/library/<slug>` which 404s; `managerKind: skill-paths` onto "Coming soon"; the skill-path version
      `"2"` guess. Recorded individually; the class — **unit-proven ≠ route-proven** — is never stated.
- [x] **[should-fix] DC-P4** **A measurement-hygiene cluster.** Four instances of a measurement lying: the
      suppressed-stderr `git fetch` measuring stale-vs-stale; the TZ age-check failing **open** west of UTC;
      a denominator counted from survivors; a probe hand-built from the artifact under suspicion (*"a probe
      intended to discriminate between two hypotheses must not be constructed from the artifact under
      suspicion"*). Individually landed; collectively unrecorded.
- [x] **[must-fix] DC-P5** **This phase's own template can issue a false green.** v2.4's
      `release-review.md:37` asserted "✅ Per-milestone decisions blended into corpus at each close" **and**
      "✅ No cross-milestone decision conflicts" while six items were aging out; `:16` asserted "✅ no
      undelivered Fate-3 items" while `release-retro.md:41` on the same date listed two as inherited
      standing carries; and `:34` points readers to a "**Completeness Ledger below**" — **the file ends at
      line 42 with no such section.** That dangling pointer is the concrete mechanism by which the standing
      test-debt carry and `DEF-M226-01` left v2.4 with no landing record. **v2.5's review must carry a real
      ledger and a per-item destination-still-valid check.**

---

## Recommended fix ordering (Phase 7)

1. **KB-B** (R1 pristine sweep 3/15) — first, or clone-dependent failures return on the next interrupted build
2. **CQ-3 + CQ-4 + D-1** — the PII/safety cluster; the release's central decision rests on it
3. **KB-A** — the clone-staleness anchor (the upstream generator of the whole class)
4. **CQ-1 + CQ-2 + CQ-6** — the `29/29` integrity cluster; then **T-4** (live re-run) to earn the headline
5. **T-1** — flake gate: fix or `KEEP-REGRESSION` with sign-off
6. **D-2, D-3, DC-3, DC-4** — the status/record contradictions
7. **D-4..D-12** — one scoped `corpus/services/` sweep (single root cause) + the 4 unstruck sites
8. **S-1..S-5** — deferral fates (user decision at Phase 9)
9. **DC-1, DC-2, DC-5, DC-6** — the unblended decisions
10. Everything tagged should-fix / nice-to-have

---

## Phase 7 disposition summary

**Legend:** `[x]` landed in this release · `[~]` routed forward with a recorded fate + named destination.

**Landed:** 6 concurrent agents, ~30 commits (25 in `rosetta`, ~7 in `rosetta-extensions`, tagged
`playbill-v25-close`). All 4 code must-fixes, all 11 doc must-fixes, all 3 KB blockers, all 6 decision
must-fixes, the flake gate, and the full should-fix/nice-to-have tail.

**Findings this review got WRONG** (recorded because a review that only lists its hits is the same
false-green shape the release is named for):

1. **KB-2 (ports)** — 8400/8401 is correct in the platform-driven context; 8080/8081 is the unset-env
   fallback. Both true. Documented rather than "fixed".
2. **KB-6 "~6 docs wider"** — refuted independently by two agents; one site, not six. The core claim
   resolved harder than filed: roadrunner has *no* caller at all.
3. **KB-6 spillover to `platform_repo.md` / `staging-sync.md`** — no caller claims in either file; no
   edits made.
4. **D-5 (skill-path manager)** — over-broad. The cohort scoreboard works and reads the mirror; only the
   per-member drill-down is dead. Retracting the mirror guidance wholesale would have destroyed correct
   advice.
5. **KB-A (`clones.lock.json` records `"HEAD"`)** — true only for detached clones; every observed row reads
   `"main"`. The real defect is *stronger*: both `ref` and `sha` are purely local and never compared to the
   remote, so a 249-behind clone is structurally indistinguishable from a fresh one. The healthy-looking
   branch name is what misleads.
6. **KB-B (~15 manifests)** — 14, not ~15. R1 covers 3 of 14.

**Findings the review MISSED, found during the fix pass:**

- Two further unqualified absolutes in `safety.md`, including the one in the first 30 lines — the sentence
  a reader is most likely to quote.
- The consequence nobody had drawn: **§3.1's "the delta cuts in the flip's favour" does not transfer**,
  because the always-open ports now carry scrubbed-real customer content. The reasoning that justified
  default-on remote reach is narrower than it reads.
- The org-scrub arm had a *second* silent-no-op mode (an empty org name treated as "nothing to do",
  indistinguishable from "never ran").
- `corpus/ops/README.md:33` described remote demo reach as "opt-in, default-off" — stale since v2.3 M220.
- A fifth false green in v2.4's archived review (its 644/14 figure is the same dirty-clone reading as T-2).
- `content-stories-routes.md`'s skill-path-legacy manager cell was also unstruck (same refutation, same
  table, not in D-7's list).

**Open item requiring user input (see `release-deferrals.md` §A and the Phase 9 ledger):**

- **CQ-3 residual — the 13 exhibits cannot be proven org-clean offline.** The source org name is never
  persisted (by design), so there is no in-repo artifact to run `SurvivingToken` against. The
  0-`<<ORG>>`-vs-840-`<<ACTOR_i>>` asymmetry is consistent with *both* "never appeared" and "copied
  verbatim". The recurrence is now structurally prevented (fail-closed + the leak-check bypass made
  unexpressible), but **the existing fixtures' status is unknown**. Settling it needs one read-only prod
  query resolving 13 org names, then an offline leak-check over the already-committed fixtures — no
  re-capture, and no customer names need enter an agent transcript.

**Process finding carried forward (DC-P5):** v2.4's review asserted four false greens and pointed at a
completeness ledger that does not exist in the file. That dangling pointer is the mechanism by which the
standing test-debt carry and `DEF-M226-01` left v2.4 with no landing record. v2.4's review is **annotated
in place, not corrected** — how a close issued them is itself the finding. Suggested fence for the next
close: refuse to emit a review containing a forward reference to a section heading absent from the file.
