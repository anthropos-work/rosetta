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
        so tier 1 *was already the junk* and the curated pool was never read. Ladder is now
        role → **CURATED** → flat. RED-proven. **Two unit tests were green throughout** — they proved the pool
        *resolved*, never that anything *read* it (D-M219-14).
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

## Battery — 5 cold reset-to-seed cycles at final code (rext `cue-to-cue-m219-r6`)

Each cycle: `down --purge` (data dir + images) → cold `up` (initdb → migrate → replay → seed) → autoverify.
Coldness is *proven per cycle*, not assumed: `data dir GONE (purged)`, `images remaining: 0`, and a **new
`PG_VERSION` mtime** (initdb re-ran).

| Cycle | up rc | purge | images | initdb re-ran | autoverify |
|---|---|---|---|---|---|
| 1 | _pending_ | | | | |
| 2 | _pending_ | | | | |
| 3 | _pending_ | | | | |
| 4 | _pending_ | | | | |
| 5 | _pending_ | | | | |

## Notes

- **Phase 0b — KB-fidelity: YELLOW** (satisfied by the census; D-M219-1; reused across sections per the skill's
  audit-reuse rule — same subsystem, knowledge docs unchanged but for this milestone's own output).
- **Two of the overview's premises were REFUTED by measurement** (F-2, F-7). The planned **new demo-patch is
  WITHDRAWN** — the non-patch fix (point the demo at the *current* surface) was available, which is the correct
  order of preference per `demopatch-spec.md §1`. **Zero platform-repo edits; zero new demo-patches.**
- The user's kickoff report is **confirmed in code**: every demo pointer targeted the **legacy** page (F-1).
- **D17 bit FIVE more times inside this milestone** — see `spec-notes.md` § The D17 tally. The sharpest:
  R-5b, where **two green unit tests proved a pool *resolved* while nothing ever *read* it**. Only a cold
  reset-to-seed exposed it.
