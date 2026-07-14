# M219 — Progress

_Section checklist. Populated from `overview.md` § Scope.In at build time; closed by `/developer-kit:close-milestone`._

## Sections

- [x] **S1 — The surface split: current vs legacy, and every pointer repointed.** *(overview items 9 + 4)*
  - [x] The split established **in code** for both vantages and **documented** in `corpus/services/ai-readiness.md`
        (§ Surfaces — routes named; the legacy orphan named; the employee no-route-of-its-own fact named).
  - [x] KB-1 (the misattributed `?cycle=` caveat), KB-2 (the falsified perf claim), KB-3, and the missing
        `ai_readiness_recommendations` table — all corrected.
  - [x] Cockpit deep-link catalog → `/ai-readiness`; the **missing** `end-user` readiness entry added.
  - [x] Dana's `jump_to` → `/ai-readiness`; Aria's + Ben's → `/home`.
  - [x] Coverage manifest repointed; sections re-derived from the CURRENT page's real i18n strings.
  - [x] The stale ACTIVE-vs-CLOSED comment fixed (item 4).
  - [x] Regression tests: a legacy pointer is a **hard failure** (`LegacyReadinessPaths` +
        `ValidateCockpitManifest`, enforced in `WriteCockpitManifest`). **Proven RED against the pre-M219
        shipped preset, GREEN after.**
  - [x] **VERIFIED LIVE on the cold-seeded cockpit manifest:** `dana-manager → /ai-readiness`,
        `aria-completed → /home`, `ben-started → /home`; catalog carries both `ai-readiness` (manager) and the
        new `ai-readiness-member` (end-user); **zero** occurrences of the legacy `/enterprise/workforce/ai-readiness`.
  - [x] Bonus: `run-coverage.sh`'s out-dir keyed on vantage **and seat** (it silently clobbered across orgs).

- [x] **S2 — Every element and sub-section FILLED, on both vantages.** *(overview items 1, 2, 3, 8)*
  - [x] The **ACTIVE cycle** seeded (F-6). Ben's funnel renders (it rendered **nothing**); Aria promoted from
        the compact archived card to the **full done-hero**; the manager's `interview` / `diagnosis` / `sources`
        sections all **PRESENT** (they were NULL/absent).
  - [x] The closed cycle retained (cycle history; the frozen rows stay live).
  - [x] The manager load re-measured and **reported**: cold **2.09 s** (was 24 ms frozen). D-DESIGN-1 / D-M219-4.
  - [x] Per-section manifest for BOTH readiness surfaces (manager: 8 sections; employee: per-hero mode).
  - [x] **The coverage sweep EXECUTED against a live demo** (it never had been — the harness could not even be
        pointed at a remote demo; **F-14**, D-M219-9). From the presenter's true vantage (a tailnet peer):
        **manager `failingSections=0` / all 8 readiness sections PASS**; **employee both heroes
        `failingSections=0`**, `ai-readiness-member-done` + `ai-readiness-member-progress` PASS. 0 escapes,
        0 not-reached. The zero-cell fence is a genuine **value** assert (`textMatch` on a leading non-zero
        digit, `minMatches: 3`) — it passes, so all three counters are non-zero.
  - [x] **R-1** interview turns — **1346 interactions / 156 sessions** (== the stage-3 count exactly), from cold.
  - [x] **R-2** Step-1 scores — Aria **30/30** (8/8 skills, weight 6.5/6.5), total **89 "champion"**; Ben **14/30**
        (3/8, weight 3.0/6.5). Arithmetic matches the platform's `computeTier1` exactly. Was 5/30 and 2/30.
  - [x] **R-3** the structural disjointness fence — Ben has **0** sessions on a readiness sim; **0** stage-1
        members have any. Sessions per sim: interview **156**, simulation **177** (= stage-2 21 + stage-3 156).
  - [x] **R-4** ≡ the ACTIVE-cycle **SEEDER** proven from a **COLD reset-to-seed** (round 2's proof inserted the
        row by hand). `ai_readiness_cycles` = **1 active + 1 closed**, out of the seeder, from zero.
  - [x] **R-5** (NEW, found by the sweep) — the readiness heroes' skills were the taxonomy's **alphabetical
        head**. Two defects: no curated category for either readiness role (Aria the *Data Analyst* claimed
        "24-hour dietary recall" / "3D Bioprinting in Dentistry"), and **"Operations Specialist" is not a public
        `job_role` at all** (Ben had no role ⇒ even his **verified** skills came off the flat head — he was
        "verified" in `15Five`/`17Track`). Fixed + fenced over **every shipped preset role** (RED-proven).
  - [x] **R-5b** (found by the cold battery) — the curated allow-list was a **NO-OP for the role it was written
        for**: `skillsForRole`'s ladder was role → **FLAT**, and `combinedNamedPool` draws its role tier from it,
        so tier 1 *was already the junk* and the curated pool was never read. R-5b made the ladder
        role → **CURATED** → flat. RED-proven. **Two unit tests were green throughout** — they proved the pool
        *resolved*, never that anything *read* it (D-M219-14).
        > ⚠️ **SUPERSEDED by R-8:** that ladder still shipped the junk, because `flat` **fired whenever the
        > curated family ran dry before `want`**. The flat tier is now **DELETED** — ladder =
        > role → curated → **general** → **STOP** (exhausted ⇒ *fewer* skills, never padded ones). See § R-8.
  - [x] The stale `app-aireadiness-snapshot-loadmembers` manifest header comment — **N/A**: F-7 refuted the
        "dead patch" premise; the patch self-heals and needs no re-pin. Recorded instead of "fixed".

- [x] **S3 — `FIX-M219-bapi-org-eid` (F-11).** *(overview item 5 — inherited from M218)*
  - [x] `SeedOrgIdentity` / `LookupOrgEid` on the BAPI store; the roster's real `org_eid` wired through
        `seedRosterMemberships` → `organizationWithEid`, behind a 3-tier ladder (roster → demo-org → stub) so
        the alignment runner stays byte-identical and **exactly one gene moves**.
  - [x] The gene is **GREEN**: Go surface **97.2% → 100.0% overall / 100% critical, 27/27, no divergences.**
        The gene is **retained** in the DNA as a permanent fence.
  - [x] Regression tests **proven RED pre-fix** and GREEN after; the no-roster fallback test passes on **both**.
  - [x] **The 5-cycle cold battery RUN** — at **final** code (per M218 D13 the count restarts on a demo-runtime
        change; F-13 + R-5 + R-5b are all on the demo's runtime/seed path). See § Battery below.

- [x] **S4 — The two absence-read-as-success gates.** *(overview items 6 + 7 — inherited from M218)*
  - [x] `TEST-M219-expressrun-dep-gate` — `alignctl` exit codes **split**: `3` = UNMEASURABLE vs `2` = REGRESSED,
        with a banner that cannot be mistaken for a pass. **Verified live.**
  - [x] `TEST-M219-freshness-gate-skips` — the preflight's silent skip now **speaks** (it printed *nothing*).
        Plus: three unit tests deferred to a *"live-verify gate"* that **does not exist** — they now report
        themselves as **coverage holes, not passes**.
  - [x] Both proven against the pre-fix behavior.

- [x] **S5 — The `ai-readiness` playthrough.** *(overview § Delivers — was a BLIND AREA)*
  - [x] `pt-world` gains **Org C** (`narrative: ai-readiness`) + the four `seed-worlds.yaml` capabilities + the
        `ai-readiness.yaml` product manifest + page objects + specs — all three artifacts in lockstep.
  - [x] `ptvalidate` green: **6 products, 15 use cases, 14 live Playthroughs, 1 TODO.**
  - [x] **EXECUTED, not merely validated** (the same D17 rule): **82 passed / 0 failed**; ptreport **14/15
        passing (93.3%)** — failing=0, unimplemented=1 (the declared assign-WRITE TODO), unimplementable=0.
        **All four AI-readiness Playthroughs PASS** (`manager-dashboard.UC1/UC2`, `member-funnel.UC1/UC2`).
  - [x] Its section in `corpus/ops/demo/playthroughs.md`.

- [x] **F-13 (net-new, found on the live host)** — the bring-up reported the academy **"started"** while it was
      dying, and served a bare **502** for the life of every stack. The node check tested **existence**, not the
      `engines: >=22` requirement (the box ran Node 18; `next dev` died importing `node:util`'s `styleText`), and
      the liveness probe polled `kill -0 $pid` for 3 s — *a probe that cannot outlive the thing it probes.* Now:
      resolve a satisfying node under `~/.nvm` or **fail loud**; poll the **port**. **Verified on a cold
      bring-up:** `started + SERVING on :13077`, and the academy answers **HTTP 200** from a tailnet peer.

## R-8 — the r7 battery's cycle 1 came back RED, and the fixes that answer it

The r7 5-cycle battery ran **cycle 1** on a **proven-cold, proven-green** stack and returned **RED**. It stopped
rather than retrying — correctly. Four defects, all fixed; **the battery restarts from zero (D13 — these are
seed-path changes).**

| # | Defect | Root cause (as MEASURED, not as guessed) | Fix | Fence (RED @ `ffc6ffe` → GREEN) |
|---|---|---|---|---|
| **D1** | Aria claims `15Five` / `17Track` / `24-hour dietary recall`; **8 other members claim "24-hour dietary recall"** — org-wide, not hero-only | **Pool SIZE, not resolution.** `want`=28; `data` shipped **28 names, 23 resolved** (5 dead), **~8 deduped** vs the role's 10 → **25 usable** → 3 off the flat head. Ben clean only because his `want` (16) was covered — **that asymmetry is the proof** | **Flat tier DELETED.** Ladder = role → curated → **general** → STOP. New `general` family (33 verified names); `data` 28→50, `operations` 30→45 | `TestSeededMembers_NeverDrawFromFlatPool` (poisoned flat, models attrition + role-overlap) + `TestCuratedLadder_CoversLargestWant` |
| **D2** | Ben renders **ROLE-LESS** | **"Operations Analyst" (`J-OPERAT-3566`) has ZERO `job_role_skills`.** The seeder's own resolver requires `EXISTS(job_role_skills)` → **rejected** → `job_role_id` NULL → `user_basic_info.job_role_id` NULL → no title. **"It resolves" ≠ "it has skills"** | → **Business Operations Analyst** (`J-BUSOPE-38C4`, 10 role-skills, same curated family) in all 3 presets + **`assertHeroRolesResolve`** = HARD seed failure | `TestAssertHeroRolesResolve` + `TestShippedPresets_EveryHeroRoleIsSeedable` |
| **D3** | The manager's 4 interview-findings sub-sections are **EMPTY**, and the gate **passes them under a disclosed exception** | **The brief's DB corroboration was a RED HERRING.** Not `conversation_extractions`; `computeInterviewInsightsV2` reads **exactly one** table — **`jobsimulation.interview_aggregated_reports`** — which **no seeder wrote** (`git grep` @ `ffc6ffe`: 0 refs) | New seeder (KPIs derived from the org's real Step-3 scores; quotes pinned to **real seeded session ids**) + **the exception DELETED** + a **900-char** content floor | `TestInterviewReport_FillsAllFourFindingsBlocks` (decodes through the platform's own contract) |
| **D4** | studio-desk `:19000` → 302 → `:13000/login` (presenter **leaves** Studio); ant-academy serves 200 but **renders nothing** (Clerk keyless bounce) | Demo-up **wiring**, needs live browser iteration | **Fate 3 → M220** (which already carries item **(i)** for the academy's Clerk wiring). **Evidence handed over:** `platform/.env` has **11** matching Clerk-key lines; the last wins and it is not `PK_DEMO` | **Landed here:** the launcher now reads the **body** (fails loud on a keyless bounce) and `ANT_ACADEMY_HOME_SECTION` drops its meaningless 40-char floor for an `AI Academy` marker + 400. **Both RED until M220 — intended.** |

**The 17-test reckoning.** The D1 fix turned **17 existing seeder tests red**, and every one of them asserted the
flat fallback as the **contract** (*"unmatched role must fall back to flat pool"*; an expected value with
`K-JUNK-1` in it). They were the bug, pinned. All **inverted**. Several fixtures modelled a taxonomy containing
**none** of the general/curated skills — a taxonomy that does not exist — which is *why* they "needed" flat to
fill; they were given the real families.

**And the fence caught itself.** The first cut of the poisoned-flat fence **PASSED against the broken ladder**
(it modelled no attrition, so Aria's 28 was exactly covered and flat never fired). It was strengthened before
shipping. See **D-M219-18**.

## Battery — 5 cold reset-to-seed cycles at final code (rext `cue-to-cue-m219-r8`)

Each cycle: `down --purge` (data dir + images) → cold `up` (initdb → migrate → replay → seed) → autoverify.
Coldness is *proven per cycle*, not assumed: `data dir GONE (purged)`, `images remaining: 0`, and a **new
`PG_VERSION` mtime** (initdb re-ran).

> **RESTARTED from zero at R-8** (D13): D1/D2/D3 are all **seed-path** changes. The r7 run's cycle 1 is
> **void as a grade** — it is retained above as the *evidence* that produced R-8.

**RESULT — 5 cold reset-to-seed rebuilds at `cue-to-cue-m219-r8` passed all five graded gates.** Each cycle is
**independently evidenced**, not asserted: a `PG_VERSION` mtime read from **inside** the container and compared
against the cycle-start fence (initdb provably re-ran), a **fresh** `autoverify.json` timestamp, per-section char
counts, **all 132 distinct claimed skill names enumerated and read** (the D1 gate), the hero role titles (D2),
`ai_readiness_cycles == 2` (1 active + 1 closed), and `interview_aggregated_reports` non-empty (D3).

| Cycle | up rc | purge | images | initdb re-ran | autoverify | graded gates |
|---|---|---|---|---|---|---|
| 1 | 0 | data dir GONE | 0 | ✅ new PG_VERSION mtime | fresh ts | 5/5 |
| 2 | 0 | data dir GONE | 0 | ✅ new PG_VERSION mtime | fresh ts | 5/5 |
| 3 | 0 | data dir GONE | 0 | ✅ new PG_VERSION mtime | fresh ts | 5/5 |
| 4 | 0 | data dir GONE | 0 | ✅ new PG_VERSION mtime | fresh ts | 5/5 |
| 5 | 0 | data dir GONE | 0 | ✅ new PG_VERSION mtime | fresh ts | 5/5 |

### ⚠️ DISCLOSED CAVEAT — the battery is not a single uncontested consecutive run

**Two further runs were graded RED for BATTERY-INTEGRITY reasons caused by an ORCHESTRATION error — not by a demo
defect.** The orchestrator ran **two batteries concurrently against the single demo host**, and one of them
**purged the stack mid-measurement** while the other was reading it. Cycle 5's `no-junk-skills` gate consequently
went **UNEXECUTED** — and *an unexecuted gate is a **FINDING**, not a pass* (D17: absence must never read as
success).

**No demo defect was observed in anything that was measured.** The five greens above are each individually
evidenced. But they were not produced by one clean consecutive sweep, and this milestone will not pretend
otherwise: **the audit trail carries a disclosed asterisk.** Recording it as a clean 5/5 would be precisely the
"status artifact that outlives the thing it describes" pattern this release exists to kill.

**Route forward (Fate 3 → M221):** `GUARD-M221-host-isolation` — a **host lock** (or a per-cycle stack `N`) so two
agents can never run cycles against one demo host. Attached to `m221-prove-on-billion/overview.md` § *Inherited
from M219*. It is a **prerequisite for M221's own gate**, which is itself a multi-cycle battery on that same
single host.

> **Provenance:** the battery was run by the orchestrating session at `cue-to-cue-m219-r8`. The **harden pass did
> not re-run it** and did not touch the demo host (`billion` is pinned, idle, clones clean). This table records
> the reported result and its caveat verbatim; it does not add an independent verification the harden pass did not
> perform.

## Notes

- **Phase 0b — KB-fidelity: YELLOW** (satisfied by the census; D-M219-1; reused across sections per the skill's
  audit-reuse rule — same subsystem, knowledge docs unchanged but for this milestone's own output).
- **Two of the overview's premises were REFUTED by measurement** (F-2, F-7). The **planned `CycleID == nil`
  demo-patch is WITHDRAWN** — the non-patch fix (point the demo at the *current* surface) was available, which is
  the correct order of preference per `demopatch-spec.md §1` (D-M219-2). **Zero platform-repo edits.**
- **ONE new demo-patch WAS authored, and it is not the withdrawn one:** **`next-web-aireadiness-flag-gate`**
  (`useAiReadinessActive.ts`). The MEMBER surface never mounts on a demo — a demo bakes no PostHog, so
  `useFeatureFlagEnabled()` resolves `undefined` **forever** and the code demands `=== true`; the org-enablement
  query is never even fired. No env/config/compose seam exists (the flag's value is not an env var), so it takes
  the sanctioned hatch. Behaviour-identical wherever PostHog **is** configured. Registered in
  `corpus/ops/demo/demopatch-spec.md` § 5 (the roster now carries **7** patches, four of them `next-web-app`).
  > ⚠️ **This bullet previously read "zero new demo-patches."** That was false from the moment the patch landed,
  > and it is the release's own signature hazard (M218 **D17**) inside the milestone's own status artifact —
  > caught by the harden pass, not by the close. The withdrawn patch and the authored patch are different patches.
- The user's kickoff report is **confirmed in code**: every demo pointer targeted the **legacy** page (F-1).
- **D17 bit FIVE more times inside this milestone** — see `spec-notes.md` § The D17 tally. The sharpest:
  R-5b, where **two green unit tests proved a pool *resolved* while nothing ever *read* it**. Only a cold
  reset-to-seed exposed it.

## M219: Hardening

### Pass 1 — 2026-07-14

**Scope manifest.** `rosetta` = docs only (13 files). The code is in `rosetta-extensions`
(`cue-to-cue-m218..HEAD`, 12 commits, 64 files / +5,659−332): Go (`stack-seeding/seeders`, `clerkenstein`,
`alignment/cmd/alignctl`), Python (`demo-stack/tests`), TS (`stack-verify/e2e`, `playthroughs/e2e`), shell
(`ant-academy.sh`, `up-injected.sh`, `run-coverage.sh`, `run-playthroughs.sh`).

**Baseline (all GREEN at HEAD):** stack-seeding 13/13 pkgs · clerkenstein 15/15 pkgs · demo-stack 506 passed/4
skipped · stack-verify TS 38 + typecheck · playthroughs TS 56 + typecheck · stack-verify pytest 109.

#### The fence audit — MUTATION-TESTED, not re-run

The release's rule is that *a fence which passes against both the pre- and post-fix code is theatre*. So each
M219 fence was graded by **surgically re-introducing the defect it claims to catch** into HEAD's production code
(a mutation), rather than by re-running it green. **7 mutations, 7 REDs — the fence set is load-bearing:**

| # | Mutation (the defect, restored) | Fence | Verdict |
|---|---|---|---|
| D1 | the `flat` tier re-appended to `combinedNamedPool` | `TestSeededMembers_NeverDrawFromFlatPool` | **RED** — names `ZZZ-FLAT-POOL-JUNK-0` |
| D2 | hero role reverted to `Operations Analyst` (0 job_role_skills) | `TestShippedPresets_EveryHeroRoleIsSeedable` | **RED** — names `J-OPERAT-3566` vs `J-BUSOPE-38C4` |
| D3 | the four findings blocks emptied | `TestInterviewReport_FillsAllFourFindingsBlocks` | **RED** — all four blocks fail *independently* |
| S1 | Dana repointed at `/enterprise/workforce/ai-readiness` | `TestStoriesPreset_ReadinessHeroesPointAtCurrentSurfaces` | **RED** — diagnostic cites the corpus doc |
| S3 | the roster eid tier dropped from `organizationWithEid` | `TestNewServer_RosterOrgEid_MembershipReportsRealEid` | **RED** — names the fabricated `org_eid_org_seed_northwind` |
| S4a | `ExitUnmeasurable` collapsed back into `ExitRegressed` | `TestExitCodes_UnmeasurableIsDistinctFromRegressed` | **RED** — *"collapsing them IS the bug"* |
| F-13 | the pid-only liveness check restored | `test_daemon_alive_but_never_serving_is_NOT_reported_started` | **RED** — mutant prints *"started + SERVING"* over a dead academy |

#### Bugs found and fixed inline (commit `b5bf65b`)

**1. `run-coverage.sh` reported the PREVIOUS run's numbers as the current run's.** The release's signature hazard
(D17), live on the coverage harness's **own reporting path**, and worse than first described:
`tests/coverage.spec.ts` writes `coverage-report.json` as its **last** statement, so a spec that throws first
(login failure, poisoned session, unreachable demo) writes **nothing** — while the script swallowed playwright's
rc with `|| true`, printed *"coverage report written to …"* **unconditionally**, and summarized whatever JSON was
on disk: the **previous run's, "GATE: MET ✅" and all**, exiting **0**. It nearly graded an M219 rebuild on
hours-old numbers from the **old, broken** stack.
- **Fix** (no clock, no `stat(1)`): **delete the report first**, so its presence afterwards *proves this run wrote
  it*. Absence ⇒ a loud `THIS IS NOT A PASS` banner + non-zero rc, never a carry-forward. `generatedAt` is
  additionally fenced against the run-start timestamp — which catches the **second** shape: a **concurrent** run
  writing the same out-dir (exactly how this milestone's battery got corrupted). playwright's rc is now
  propagated.
- **Fenced** by `TestCoverageReportFreshness` (4 tests) + the shellcheck gate the script was never in.
  **RED-proven: 3 of the 4 FAIL against the pre-fix script** — including `0 == 0 : a sweep that produced NO report
  exited 0` — while the success-path test passes on **both**, which is correct: that path was never broken.
- **`run-playthroughs.sh` already got this right.** M204 iter-02 hit the identical class (a `--reporter` flag
  suppressing the json reporter, leaving `last-run.json` stale and decoupling `ptreport` from the actual run) and
  fixed it *there*. **The lesson was never propagated one directory over.**

**2. `resolveClaimedSkillNames` — the flat pool's OTHER consumer — was entirely untested.** R-8 deleted the flat
tier from the claimed-tail ladder, but the pool is **retained** for two legitimate callers, and one of them (the
AI-generated-batch path, `generated_batch.go:254`) takes it as a parameter with **zero tests**. It is safe today
because it uses the pool as a **name→node-id INDEX** and **drops** names it cannot resolve — the
CODE-owns-structure / AI-owns-content boundary. That boundary is load-bearing and was **unpinned**: a future
author "helpfully" padding a short generated tail from `flatPool` would reintroduce **D1 on a path D1's own fence
does not watch**. Fenced (4 subtests); **RED-proven** against a pad-from-flat mutation (`K-JUNK-0`/`K-JUNK-1`
enter the result).

**3. `namedSkillRefs.flat` still documented itself as *"the universal fallback pool"*.** The comment described the
mechanism R-8 **deleted**, and would have walked the next author straight back into it. Corrected to state what
`flat` is now for (filter + lookup, **never** fill) and what it must never become again.

#### Documentation defects found (the same D17 shape, in the docs)

**4. `progress.md` claimed *"zero new demo-patches"* — false.** M219 **did** author one:
**`next-web-aireadiness-flag-gate`**. The *withdrawn* patch (`CycleID == nil`, D-M219-2) and the *authored* patch
are **different patches**, and the note conflated them. Corrected, with the distinction spelled out.

**5. `demopatch-spec.md`'s patch roster was missing that patch** — the doc CLAUDE.md calls *"the contract to read
before adding or re-pinning any patch"* did not list the patch this milestone added, and still said *"the **three**
`next-web-app` patches"* (now four). Roster now carries **7**.

**6. `ai-readiness.md` asserted the exact claim the new patch FALSIFIES** — and the patch manifest already said
*"the doc is corrected in the same milestone"*, **which it was not**. The doc said absence of PostHog *defaults the
flag through*; in fact it makes the flag `undefined` **forever** while the code demands `=== true`, so the **member
surface never mounts on a demo**. **Why the false claim survived: it was proven against the wrong vantage** — the
cited "empirical proof" (M53 **AB5**) renders the **manager** dashboard, but `flag_ai_readiness` gates the
**employee side only**, so AB5 was never evidence about that gate at all. The doc's own parenthetical conceded the
mechanism was *"inferred … not separately traced"*. **This is the same wrong-vantage trap that made two of M219's
opening premises false** (F-1/F-2). Corrected in full, with the mechanism traced in code.

**Coverage delta:** stack-verify pytest **109 → 114** (+5); stack-seeding seeders **+1 fence, 4 subtests**. Go
coverage is a *finder* here, not a goal — the milestone's real defects (R-5b, D1, D3) were **invisible to a green
unit suite** and only a cold reset-to-seed exposed them (**D-M219-14**), which is why this pass graded fences by
**mutation** rather than by chasing line coverage.

**Flakes stabilized:** none observed. 3 consecutive sequential runs of the new fences: clean.

**Knowledge backfill:** `corpus/services/ai-readiness.md` (the falsified PostHog gate-2 claim, traced in code) ·
`corpus/ops/demo/demopatch-spec.md` (the 7th patch + the corrected `next-web-app` count) ·
`m220-cue-sheet/overview.md` (the academy **session-poisoning** escalation) · `m221-prove-on-billion/overview.md`
(the `GUARD-M221-host-isolation` + native-academy-reap Fate-3 items).

### Stop condition
**Pass 1, stopped.** The Step-2b scan surfaced nothing further worth adding: the fence set is mutation-proven
load-bearing (7/7), the two real code gaps found (the coverage-report staleness path + the untested flat-pool
consumer) are both fixed and RED-proven, and the remaining findings are documentation defects — now corrected. The
known REDs (academy session-poisoning, studio-desk 302) are **routed to M220 by design** and their fences
**deliberately report RED** until it lands; deepening tests against them here would be testing a defect this
milestone does not own.
