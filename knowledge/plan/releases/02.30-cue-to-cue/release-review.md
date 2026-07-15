# Release Review — v2.3 "cue to cue" (M217 → M221)

## Verdict

**YELLOW — proceed to close-release; discharge the non-blocking follow-ups before the merge→main + tag.**

No blockers, no RED dimension. The code / tests / supply-chain dimension is clean GREEN
(verified by real execution against `HEAD = cue-to-cue-m221-final`). What holds the aggregate at
YELLOW is (a) **2 must-fix documentation-staleness defects** — the safety contract still describes
the *pre-fix* academy bind that the release's own code-of-record already corrected — and (b) a cluster
of **close-release pipeline paperwork that is owed but not yet discharged** (v2.4 deferral landing spots,
the cross-release deferral audit + user sign-off, the release-level D17 synthesis, and the release
metrics rows). None of it is new engineering; all of it is the close-release pipeline doing its job
before the tag.

**Per-dimension verdicts:** Scope YELLOW · Documentation YELLOW · Code/Tests/Supply-chain GREEN · Decision Consolidation YELLOW.

**Finding tally:** 0 blocker · 2 must-fix · 9 should-fix · 3 nice-to-have · 9 info.

---

## Scope

*Dimension verdict: YELLOW. v2.3 delivered its scope honestly — all 5 milestones `done` + `--no-ff`
merged into `release/02.30-cue-to-cue`, the headline gate (click→ACCESS <5s) MET at M218 and RE-PROVEN
live at M221 (8/8, maya 2.11s / dan 1.31s over the tailnet, no flags), and all 6 original user asks
delivered and live-proven on `billion`. The rosetta release diff is 108 files, 100% docs/planning —
zero platform-repo edits. YELLOW only because the escape-hatch paperwork the audit-deferrals skill
MANDATES for the KEEP-DEFERRED carries is not yet discharged.*

- [ ] **[should-fix]** The 4 v2.4 tail carries (F4 academy-grid-render, BURNIN-M221-dev-public-host, F-M220-4, PROBE-M218-c3-rerun) have **NO `roadmap-vision.md` landing spot**. The audit-deferrals skill Phase 5 mandates, for every KEEP-DEFERRED-WITH-SIGNOFF, "add the item to roadmap-vision.md under the target release with a backref". `grep` for v2.4|BURNIN|F-M220-4|c3-rerun|F4 → 0 hits; the file has no v2.4 section at all. The carries currently live ONLY in `state.md:138-141` + the M221 audit; on the release→main merge + state rotation the canonical futures doc would carry no v2.4 entry (the untracked-void / silent-scope-erosion this skill exists to prevent). *Fix:* close-release Phase 1b authors the v2.4 vision entry with per-item backrefs to the M221 audit BEFORE the merge (matching the DEF-M21-01 "landed at close-release so it survives the merge" precedent). *(`knowledge/plan/roadmap-vision.md` — no v2.4 section; contrast `state.md:138-141`)*
- [ ] **[should-fix]** The blocking **cross-release deferral audit + per-item user sign-off** for the 4 tail carries have NOT been executed. No release-scope audit file exists (`02.30-cue-to-cue/` has only the 4 milestone-scope `audit-deferrals/` dirs, no release-root one). The M221 milestone audit explicitly hands the RELEASE-SCOPE-DEFER sign-off to close-release Phase 1b. *Fix:* close-release Phase 1b runs `/developer-kit:audit-deferrals --scope=release` and obtains explicit LAND-NOW/DROP/KEEP-DEFERRED-WITH-SIGNOFF from the user for each carry (BURNIN, F-M220-4, PROBE-c3 span ≥2 milestones → require fresh sign-off per the Aging Policy). *(`02.30-cue-to-cue/` no release-scope `audit-deferrals/`; m221 `deferral-audit-2026-07-15-m221-close.md:75-80`)*
- [ ] **[should-fix]** One Fate-3 routed M218→M221 (PROBE-c3-rerun, C-3 router-403 re-check) did NOT deliver in-release and escaped to v2.4 — the exact class flagged as the headline failure mode. Three more (F4=F-M220-2 academy-empty-catalog, F-M220-4, BURNIN) routed M220→M221 likewise appear in M221's In-list but did not deliver. This is the HONEST version (disclosed, fresh-reasoned, non-gate, routed to v2.4), **not** a silent drop — the 8/8 gate is independent of all 4, each has a structural non-"no-time" reason, none is Fate-1-able now (platform-repo or live-infra). *Fix:* no new work owed beyond the two items above; confirm at close-release all 4 are recorded in `roadmap-vision.md` so the M221→v2.4 chain is auditable end-to-end. *(m221 `overview.md:76` + `progress.md:88-98`; m220 `progress.md:126,442`)*
- [ ] **[nice-to-have]** `metrics-history.md` is missing MORE rows than the tracked note admits. Newest row is v2.1 (line 9); line 10 jumps to v1.10 — so **v2.0 AND v2.2 are both absent**, and v2.3 is not yet written. `state.md:148` records only "no v2.2 row", understating the gap. *Fix:* close-release adds the v2.3 row and backfills v2.0 + v2.2 (also note/discharge the missing v2.1 + v2.2 release-scope deferral audits per `state.md:149`). *(`metrics-history.md:9-10`; `state.md:148`)*
- [ ] **[info]** Naming drift on the C-3 carry: M218 close renamed it `PROBE-M221-c3-rerun` (`m218 progress.md:134`), but M221's overview/progress/audit + `state.md` all call it `PROBE-M218-c3-rerun`. Same item (router-403 re-check); cosmetic, but could trip a reader tracing the carry across the merge. *Fix:* pick one id (`PROBE-M218-c3-rerun` is the majority form) when authoring the v2.4 entry.
- [x] **[info]** STRENGTH (audit trail): the D17 spine is honestly and thoroughly told, not sanded down — ~24 self-incriminating instances verified across all 5 progress ledgers + `state.md:61-83` (M217 unsourced `reap.sh` that never ran; M218 `--purge`-that-purged-nothing + SKIPPED-patch-graded-green; M219 `run-coverage.sh` printing the previous run's numbers + 17 tests guarding the junk-fallback bug; M220 mutation-batteries-claimed-but-never-committed + "2 orgs"/11-sites/4-releases; M221 `test_reap.py` "Ran 21 tests OK"/omitted-20 + wrong-root cache shadow). The release actively refuses to round up (M219 two disclosed battery asterisks; M221 "reproducibly = 2 cold cycles, disclosed, not 5"). No remediation.

---

## Code Quality

*From the Code/Tests/Supply-chain dimension: GREEN. Verified by real execution against
`HEAD = cue-to-cue-m221-final`.*

- [x] **[info]** `go vet` clean on all 6 modules; `py_compile` clean on all 53 `.py`. Dead code from iteration is genuinely cleaned (M217's "drift injector" has zero remaining references in prod or test). Integration seams are coherent and load-bearing — the F1 store-root SHADOW fence drives the real replay path and asserts a loud wrong-root message; the demopatch G7 apply-post-condition now re-hashes ON DISK so it can actually fire; the coverage harness `rm -f`'s + freshness-gates its report; the exposure guard is one file *extended* (not forked) across M220→M221. No remediation.
- [ ] **[info]** M221 close claims "shellcheck clean on all 7 M221-touched shell scripts", but at shellcheck's DEFAULT severity `hostlock.sh` emits SC2064 (lines 74-81) and `reap.sh`/`ant-academy.sh` emit SC2086/SC2015/SC2012 info. All are benign/intentional — the `hostlock.sh` SC2064 trap bodies MUST expand `$n`/`$prev` at install time (correct). The claim is true only at *error* severity; code is correct, only the claim's precision is loose. *Fix:* qualify the claim as "shellcheck clean at error severity" (or add inline `# shellcheck disable=SC2064` on the intentional trap lines). *(`hostlock.sh:74-81`, `reap.sh:105,329`, `ant-academy.sh:91,149` vs m221 `metrics.json close_review.code_quality`)*

---

## Documentation

*Dimension verdict: YELLOW. The 21-doc v2.3 diff is unusually coherent and reads as a deliberate corpus.
Task-item-2 fully passes (the M43-D5 "~2-5s we can't shorten" lie is corrected everywhere; `latency-budget.md`
states the real p95<5s ACCESS gate). All relative links resolve; org counts consistently "3 orgs"; `state.md`
correctly shows v2.3 code-complete awaiting close-release. BUT there is one material coherence defect —
itself a textbook D17 instance — where a status artifact (the M220 exposure disclosure) outlived the thing
it described (the M221 fix landed one iteration later) and was never re-synced.*

- [x] **[must-fix]** *(FIXED at close-release — `safety.md` §3.1 + §3.6 now state the fix LANDED in M221, M220 S3 kept as a dated measurement; container `0.0.0.0` claim untouched.)* The M221 academy-loopback-bind fix (F-M220-5) **SHIPPED in the release's code-of-record** but the safety contract still says the opposite. Final tag `cue-to-cue-m221-final` has `ant-academy.sh:330` binding `-H 127.0.0.1` by default (comment L315: "M221 (F-M220-5): bind next dev LOOPBACK on the localhost path"); iter-03 confirms it landed + RED-proven. Yet `safety.md §3.1` still asserts as current fact: `ant-academy (3077+off) | *:13077 | YES` (L433), "the academy is world-published on **every** demo" (L436), and "(Routed: FIX-M221-academy-loopback-bind …)" (L442) as if pending. This is the release's own D17 signature embedded in the very safety doc that names the hazard; it also undercuts §3.6/frontend-tier's "the BAPI is the demo's FIRST loopback-bound port … because every other port is world-published" reasoning (there are now two loopback host-native ports). *Fix:* add a landed-in-M221 note to §3.1 (academy binds 127.0.0.1 on a localhost demo, fix shipped not pending); keep the M220-S3 measurement as an explicitly-dated historical artifact; reconcile §3.6. *(`safety.md:433,436,442` + §3.6 L651-655 vs rext `cue-to-cue-m221-final:demo-stack/ant-academy.sh:330`)*
- [x] **[must-fix]** *(FIXED at close-release — `tailscale-serve.md` L462/L465 flip table + L630-633 and `demo-up-defaults.md` exposure box now show `127.0.0.1:13077` / "landed M221 F-M220-5"; container ports left as `0.0.0.0`.)* Same academy staleness in the remote-access runbook + defaults contract. `tailscale-serve.md` flip table describes the localhost-demo academy bind as `*:13077`/`0.0.0.0` (L462, L465) and states "the academy is world-published on `*:13077` on every demo, flag or no flag … FIX-M221-academy-loopback-bind … routed to M221" (L630-633) — present tense, fix pending. `demo-up-defaults.md`'s exposure box says "ant-academy binds `*:13077` and answers 200 from the tailnet IP ❌" (L124-126). All three contradict the shipped `-H 127.0.0.1` default. *Fix:* update the localhost-demo column to `127.0.0.1:13077` and change "routed to M221"/"world-published on every demo" to "landed in M221 (fix shipped)"; keep the historical M220 exposure story as dated. *(`tailscale-serve.md:462,465,630-633`; `demo-up-defaults.md:124-126`)*
- [x] **[should-fix]** *(FIXED at close-release — `CLAUDE.md:305` and the demo family index `corpus/ops/demo/README.md` tailscale blurb both rewritten to v2.3 M220 state: default-on for demo / opt-out `--no-public-host`, `/dev-up` opt-in; now consistent with `CLAUDE.md:282`.)* Both parent indexes still describe `tailscale-serve.md` with the v2.2 "opt-in, default-off" framing that M220 flipped to default-on and explicitly RETRACTED AS FALSE. `CLAUDE.md:305` ("opt-in, default-off … an opt-in flag only") and `README.md:149-155`. This makes CLAUDE.md internally contradictory: L282 (safety.md entry) says the exposure was denied "until M220; the retraction is in place and fenced", while L305 re-asserts the retracted "opt-in, default-off" summary — the D17 pattern sitting in the top-level project index. *Fix:* rewrite both blurbs to v2.3 M220 state (remote reach DEFAULT-ON, opt-out `--no-public-host` for `/demo-up`, opt-in for `/dev-up`; drop "opt-in, default-off"; bump the "(v2.2 M212-M214)" tag to note the M220 flip). *(`CLAUDE.md:305` vs `:282`; `README.md:149-155`)*
- [ ] **[should-fix]** `safety.md` Part 3 (the new v2.3 M220 exposure axis) is not surfaced in the README parent-index blurb — README's safety.md entry still reads "read-side … + write-side … (v1.3 M15)" with no mention of the third axis, while CLAUDE.md's entry (L282) WAS updated — so the two parents disagree on what `safety.md` now contains. Task-item-6 asked to confirm Part 3 is discoverable from parents; it is from CLAUDE.md but not from the demo family index that says "Read this first". *Fix:* append a one-line Part 3 mention (the exposure side / disclosure-not-a-third-never, v2.3 M220) to README's safety.md entry. *(`README.md:95-97`)*
- [ ] **[nice-to-have]** CLI-flag count drift: `demo-up-defaults.md` header is "## CLI flags — all 10" (L151, 10-row table incl. `--no-public-host`), but `README.md:140` and `CLAUDE.md:284` both summarize it as "9 CLI flags". Env-knob count (25) matches; only the CLI count is stale. *Fix:* change "9 CLI flags" → "10 CLI flags" in both index entries. *(`README.md:140`; `CLAUDE.md:284` vs `demo-up-defaults.md:151`)*
- [ ] **[info]** Minor imprecision (not a contradiction): `README.md:137` and CLAUDE.md's frontend-tier entry still describe ant-academy as "natively (Clerk-free)" while `frontend-tier.md:24` now documents it as Clerkenstein-wired sharing the demo session (v2.3 M220 S5). "Clerk-free" remains defensible (no real Clerk) but the one-liners predate the M220 rewiring. *Fix (optional):* note "Clerkenstein-wired (M220), shares the demo session".

---

## Tests & Benchmarks

*From the Code/Tests/Supply-chain dimension: GREEN — tests actually run, not trusted. Go = 1831 test
funcs, all modules `ok`. Python full tree = 1341 collected / 1325 passed / **0 FAILURES / 0 ERRORS** /
16 skipped across all 5 sections (demo-stack 663, stack-injection 260, stack-core 182, stack-verify 114,
dev-stack 122) — exactly the milestone-close claim. The dev-stack "spins forever" ghost is genuinely
DISCHARGED (123s, 118 passed / 4 skipped, no hang). Suite honesty confirmed: `test_reap.py` now runs all
41 tests on BOTH paths (was 21/41 + false "OK"). Flake 0 every milestone. The two findings below are
metrics-RECORDING hygiene for the pending close-release — neither touches code, tests, or dependency safety.*

- [ ] **[should-fix]** TS e2e count is not recorded by the canonical `playwright test --list` method across v2.3, and the one place it IS recorded reads as a FALSE regression. M219 `metrics.json tests.ts` records `stack-verify_e2e=38 + playthroughs_e2e=56 = 94` (an "unit/typecheck" measure) against v2.2's recorded 124 (via `playwright test --list`). If close-release copies M219's 94 it registers a **124→94 DECREASE — a blocker under the regression rules**. Re-running the authoritative like-for-like method at HEAD: `stack-verify/e2e = 69` + `playthroughs/e2e = 82` = **151, i.e. +27, an INCREASE**. (Itself a live D17 instance — a count artifact read as evidence.) *Fix:* close-release records TS e2e = 151 (69 + 82) via `npx playwright test --list`; note the 94 was a different measure so no false-regression is filed. *(m219 `metrics.json tests.ts`; release-level `metrics.json` not yet emitted)*
- [ ] **[should-fix]** Cross-milestone Python counting-method drift produces two apparent per-milestone DECREASES that are real method-changes, not shrinking suites (both documented in-place, but they muddy a naive diff): (a) stack-injection 194 (M217, total incl. skipped) → 186 (M218, passing-only); (b) stack-core 152 (M220, passing) → 182 (M221, JUnit collected incl. 30 subtests). Underlying suites only grew (stack-injection 260 collected / 252 pass / 8 skip; stack-core 182 collected / 182 pass at HEAD). *Fix:* standardize release-level metrics on JUnit-XML COLLECTED counts (M220/M221 already do); state the like-for-like release delta explicitly (v2.2 668 across 3 sections → v2.3 1341 across 5; the 3 shared sections 668 → 1105). *(m218 `metrics.json python_delta_vs_m217`; m221 `metrics.json python_delta_vs_m220`)*
- [ ] **[info]** No release-level `metrics.json` / `release-review.md` / `release-retro.md` / `dependencies.lock` existed for v2.3 as of review time (v2.2 has all four at `archive/02.20-panorama/`). Expected — M221 defers the release→main merge + tag to close-release, and this review is part of that close. *Fix:* emit the release metrics recording the measured aggregate (Go 1831, Python 1341/0-fail, TS 151, flake 0, 0 net-new deps) + release-retro + `dependencies.lock`; bake in the two should-fix items above. *(`02.30-cue-to-cue/` — no release-level artifacts present at review time)*
- [x] **[info]** Metrics regression GREEN: no test-count decrease on the like-for-like methods (Go 1749→1831 +82; Python 668→1105 on shared sections; TS 124→151 +27 via `playwright test --list`), flake 0 every milestone, 0 net-new deps, 0 platform-repo edits (rosetta diff = 108 files, all docs/planning). The metrics.json files disclose caveats rather than round them away. No remediation.

---

## Supply chain

*From the Code/Tests/Supply-chain dimension: GREEN. `govulncheck` = "No vulnerabilities found" on all
6 modules. No findings — recorded clean.*

- [x] **[info]** The ONLY dep change in the entire v2.3 range is `golang.org/x/crypto v0.51.0 → v0.52.0` (indirect, behaviour-neutral, patched). Zero net-new direct deps; zero `package.json`/`requirements` changes. No remediation.
- [x] **[info]** All licenses permissive (MIT/BSD/Apache, no copyleft). The corpus-flagged deliberate dep `anthropos-work/ai v1.40.1` is a v1.10/M45 addition — present, unchanged, org-internal, clean. rext ships zero Python third-party deps (the `pip-audit` "300 vulns" is the host's global env, not the project). No remediation.

---

## Decision Consolidation

*Dimension verdict: YELLOW. v2.3's decision record is unusually coherent and honest — **no cross-milestone
contradictions**, the D-DESIGN-1 ID collision is handled by an explicit written supersession (M220 D7)
reproduced and disambiguated in five places, and the D17 signature is told honestly per-milestone.
Nothing central lives only in git-log. YELLOW because two consolidation items are owed and at risk of being
dropped at merge (KB Consolidation Phase 3b has not yet run).*

- [ ] **[should-fix]** The release's SIGNATURE finding — the D17 thread — has **NO release-level synthesis**. The corpus captures the HAZARD (`verification.md` "THE STALE-VERDICT HAZARD") but scoped to M218's FIVE instances only. The deepest meta-lesson — that **naming the class did NOT inoculate against it** (it recurred ~24× across all 5 milestones, INCLUDING inside the very passes convened to fix it: M218 D18 "the sixth instance… the first to appear inside the pass convened to name it"; M219 self-counts D-M219-14/17/18/20/21/25/26/27; M220 D30 "the third time this milestone… reproduced inside the battery built to enforce D17"; M221 harden's "Ran 21 tests OK" false-green) — lives ONLY scattered across per-milestone files. The takeaway "**a named hazard is not a fence; only an executable probe binds**" (M218 D18) has no durable cross-milestone home. *Fix:* at close-release Phase 3b, write the cross-milestone D17 arc (the ~24-instance count, recurrence-after-naming, the probe rule) — extend `verification.md`'s hazard section to release scope, or add a release-retro synthesis. *(`verification.md:261-301` scoped to M218's 5, vs scattered m218 D18 / m219 `decisions.md` D-M219-26 / m220 D30 / m221 `retro.md:32-35`)*
- [ ] **[should-fix]** M218 D18 — the reusable close-discipline rule "**resolve BASELINE to the milestone START ref** (release-branch merge-base or the prior milestone's rext tag), NEVER a mid-milestone commit; a mid-milestone baseline converts a regression into a clean bill of health (D17 itself)" — lives ONLY in `m218/decisions.md`. It already caught one misfiled regression (`test_next_web_block_shape` passed at `cue-to-cue-m217` but M218 broke it; the harden pass mis-declared it "pre-existing" by baselining on M218's OWN iter-05 commit `f296e5e`). No durable process home; will be lost when this milestone's decisions.md ages out. *Fix:* surface D18 in the release consolidation as a standing close-discipline rule (state the baseline ref explicitly in every ledger). *(`m218-seat-change/decisions.md:178-211`)*
- [ ] **[nice-to-have]** `safety.md §3.5` uses BARE "D-DESIGN-1" at lines 520-521 — violating the section's OWN rule stated 11 lines later at :532 ("Cite it as 'v2.2's D-DESIGN-1', never bare"). Contextually safe (the header scopes to v2.2's D-DESIGN-1) so it does not misresolve in practice, but it is a self-inconsistency in the one doc that legislates the disambiguation. *Fix:* prefix both with "v2.2's". *(`safety.md:520-521`, rule at `:532`)*
- [x] **[info]** SUPERSESSION CONFIRMED EXPLICIT (PASS, not silent). M220 superseded v2.2's D-DESIGN-1 ("public reach never default-on") for the demo path and wrote it in FIVE cross-referenced, always-disambiguated places: `safety.md §3.5`, `demo-up/SKILL.md:102-108`, `tailscale-serve.md:19`, `roadmap.md:129` ("SUPERSEDES v2.2's D-DESIGN-1"), `state.md:91`. A written supersession, not a contradiction. No remediation.
- [x] **[info]** NO cross-milestone CONFLICTS found (PASS). The five milestones' decisions form a coherent forward-chained record: M218 D13 honoured by M219 D-M219-12/15 + M221 TOK-01; M219 D-M219-23 routes host-isolation forward, M221 iter-02 lands it; M220 D23 routes the academy-loopback fix forward, M221 D-M221-03a lands it. The only ID collision (v2.2's vs v2.3's D-DESIGN-1) is a naming hazard handled by M220 D7, not a decision conflict. No remediation.
- [x] **[info]** BOOKKEEPING-corruption lesson CONFIRMED recorded beyond git-log (PASS). The M221 iter-02 host-isolation lock prevented DATA corruption but not BOOKKEEPING corruption (the orchestrator prematurely committed the iter-05 close `766c029` telling a stale story while the agent was correcting r3→r4). Captured in `m221 retro.md:29-31`, commit `3c64af1`'s PROCESS NOTE, and iter-05 `decisions.md D-M221-05f`. Recorded at INCIDENT granularity; the generalizable rule "isolating a shared RESOURCE does not isolate the shared RECORD" is stated in the commit but not elevated to a reusable process rule (its natural home is the developer-kit workflow, outside this repo). *Optional:* generalize to a one-line reusable rule in the release consolidation.
- [x] **[info]** UNBLENDED-decision sweep CONFIRMED essentially clean (PASS). M220's close `progress.md:403-407` tracks every blend-triaged decision as landed; corpus-side durable homes verified present (M218 F1-3 capability-coverage check at `alignment_testing.md:174-202`; M219 D-M219-7 UNMEASURABLE exit-3; the D-DESIGN-4 "3 orgs" fix now across 9 corpus/skill files; the loopback/0.0.0.0 retraction at `tailscale-serve.md:592-632`). CAVEAT: the rext-side blend targets (`stack-core/README.md`, `stack-injection/README.md`) live in the SEPARATE `rosetta-extensions` repo, marked `[x]` in the M220 blend log but **unverifiable from this rosetta worktree**. *Optional:* spot-check the two rext README blends in `.agentspace/rosetta-extensions @ cue-to-cue-m221-final`.

---

## Bottom line

v2.3 "cue to cue" is **code-complete and its headline gate is live-proven** — all 5 milestones are `--no-ff`
merged into the release branch, click→ACCESS <5s was re-proven 8/8 over the tailnet on `billion` with no
flags, and the code / tests / supply-chain dimension is **clean GREEN with zero blockers** (Go 1831 + Python
1341/0-fail + TS 151 all up, `govulncheck` clean, one indirect patch bump, zero platform-repo edits). It is
fundamentally sound and **ready to close**.

Before the merge→main + tag, close-release must (1) fix the **2 must-fix documentation-staleness defects**
where `safety.md` / `tailscale-serve.md` / `demo-up-defaults.md` still describe the *pre-fix* academy bind
(the shipped code binds `127.0.0.1`), and (2) discharge the tracked close-release paperwork: author the
**v2.4 `roadmap-vision.md` landing spots + obtain user sign-off** for the 4 tail carries (the mandated
deferral-audit escape-hatch), write the **release-level D17 synthesis** (the ~24-instance, 5-milestone arc),
and emit the **release metrics** — recording TS e2e as **151** via `playwright test --list`, not M219's
differently-measured 94, so no false 124→94 regression is filed. None of this is new engineering; it is the
close-release pipeline doing its job.
