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
  - [x] Bonus: `run-coverage.sh`'s out-dir keyed on vantage **and seat** (it silently clobbered across orgs).

- [x] **S2 — Every element and sub-section FILLED, on both vantages.** *(overview items 1, 2, 3, 8)*
  - [x] The **ACTIVE cycle** seeded (F-6). Live-verified: Ben's funnel renders (it rendered **nothing**); Aria
        promoted from the compact archived card to the **full done-hero**; the manager's `interview` /
        `diagnosis` / `sources` sections all **PRESENT** (they were NULL/absent).
  - [x] The closed cycle retained (cycle history; the 199 frozen rows stay live).
  - [x] The manager load re-measured and **reported**: cold **2.09 s** (was 24 ms frozen). D-DESIGN-1.
  - [x] Per-section manifest for BOTH readiness surfaces (manager: 7 sections; employee: per-hero mode).
  - [ ] ⚠️ **The coverage sweep has NOT been EXECUTED against a live demo.** The asserts are written but
        unrun — and **an unexecuted assert is not evidence** (D17). → PENDING.
  - [ ] **R-1** `howWeMeasure.cycleTotals.interviewQuestions = 0` — a genuine zero cell, on both read paths.
  - [ ] **R-2** Step-1 scores are weak (Aria 5/30, Ben 2/30 — few `user_skill_evidences` match the org's 8
        `ai_readiness_skills`). A believability weakness, not an empty section.
  - [ ] **R-3** Ben carries interview/sim *signals* while his step-progress row says `not_started` — the
        session seeder and the funnel seeder disagree.
  - [ ] **R-4** The **seeder** path for the active cycle is unit-tested but **not yet proven on a cold
        reset-to-seed** (the live proof applied the row by hand).
  - [x] The stale `app-aireadiness-snapshot-loadmembers` manifest header comment — **N/A**: F-7 refuted the
        "dead patch" premise; the patch self-heals and needs no re-pin. Recorded instead of "fixed".

- [x] **S3 — `FIX-M219-bapi-org-eid` (F-11).** *(overview item 5 — inherited from M218)*
  - [x] `SeedOrgIdentity` / `LookupOrgEid` on the BAPI store; the roster's real `org_eid` wired through
        `seedRosterMemberships` → `organizationWithEid`, behind a 3-tier ladder (roster → demo-org → stub) so
        the alignment runner stays byte-identical and **exactly one gene moves**.
  - [x] The gene is **GREEN**: Go surface **97.2% → 100.0% overall / 100% critical, 27/27, no divergences.**
        The gene is **retained** in the DNA as a permanent fence.
  - [x] Regression tests **proven RED pre-fix** (they fail with the exact fabrication
        `org_eid_org_seed_northwind`) and GREEN after; the no-roster fallback test passes on **both**.
  - [ ] ⚠️ **The fresh 5-cycle cold battery has NOT been run.** It is on the demo runtime path. → PENDING.

- [x] **S4 — The two absence-read-as-success gates.** *(overview items 6 + 7 — inherited from M218)*
  - [x] `TEST-M219-expressrun-dep-gate` — `alignctl` exit codes **split**: `3` = UNMEASURABLE (no genes ran,
        no score exists) vs `2` = REGRESSED (a *measured* score below the gate), with a banner that cannot be
        mistaken for a pass. **Verified live:** expressrun without `@clerk/express` now exits **3**; the Go
        gate still exits **0**.
  - [x] `TEST-M219-freshness-gate-skips` — the preflight's silent skip now **speaks** (it printed *nothing*).
        Plus: three unit tests deferred to a *"live-verify gate"* that **does not exist** — they now report
        themselves as **coverage holes, not passes**.
  - [x] Both proven against the pre-fix behavior.

- [ ] **S5 — The `ai-readiness` playthrough.** *(overview § Delivers — a BLIND AREA)* — **NOT STARTED.**
  - [ ] `pt-world` has **no** `ai-readiness` org (confirmed); `seed-worlds.yaml` has none of the capabilities,
        so every new use-case would hard-fail `checkPreconditionCoverage`. All three must land in lockstep.
  - [ ] The `ai-readiness` product manifest + page objects + specs; `ptvalidate` green.
  - [ ] Its section in `corpus/ops/demo/playthroughs.md`.

## Notes

- **Phase 0b — KB-fidelity: YELLOW** (satisfied by the census; D-M219-1). KB-1..KB-3 fixed in S1.
- **Two of the overview's premises were REFUTED by measurement** (F-2, F-7). The planned **new demo-patch is
  WITHDRAWN** — the non-patch fix (point the demo at the *current* surface) was available, which is the correct
  order of preference per `demopatch-spec.md §1`. **Zero platform-repo edits; zero new demo-patches.**
- The user's kickoff report is **confirmed in code**: every demo pointer targeted the **legacy** page (F-1).
- rext tagged **`cue-to-cue-m219`**. Suites: clerkenstein 14/14 · alignment 7/7 · stack-seeding 13/13 ·
  demo-stack 64/64 · coverage-manifest 33/33. gofmt clean.
