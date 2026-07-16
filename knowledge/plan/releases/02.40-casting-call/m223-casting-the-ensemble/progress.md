# M223 — Progress

_Section checklist, derived from `overview.md` § Scope.In. To be worked by `/developer-kit:build-milestone`._

## Sections

- [x] **S0 — KB-fidelity gate** (pre-milestone; `/developer-kit:audit-kb-fidelity` — needs M222's `hiring.md`) — **GREEN** (`kb-fidelity-audit.md`; all topics PAIRED + ALIGNED)
- [x] **S1 — The 4th `stories[]` entry** (`narrative: hiring`, RoleMix ≈ 0.1 admin / 0.9 candidate, hero-trio
      placeholder) — Meridian Talent (id `hiring` → `HiringOrgID()`), size 50 → 5 admin + 45 candidate, 0
      heroes (M224 seats); manifest regenerated (honesty gate green); `TestStoriesPreset` → 4 stories.
- [x] **S2 — `HiringConfigSeeder`** (the 5 shared HIRING sims) + the **type-aware hiring-sim reader**
      + the disjoint reservation — `readHiringSimPool` (`type='SIMULATION_TYPE_HIRING'`), `hiringSimRefs`
      (first 5), `withoutIDs` withholds them from the generic pool; config writes 5
      `organization_sim_invitation_links` positions/org (folds S6). 7 tests green.
- [x] **S3 — DROPPED** (M222 D4/BA-6/Fate-3): NO `directus.job_position` replay — 0 rows captured (the
      "443" was a PROD count), and the scoreboard doesn't read `job_position`. The 5 "positions" ARE 5
      real captured `SIMULATION_TYPE_HIRING` sims (resolved by S2's reader). Recorded in `decisions.md`
      D4 + the `snapshot-spec.md` note (S7). Built nothing for S3.
- [x] **S4 — The `candidate`-assessment funnel seeder** — `HiringFunnelSeeder`: writes the
      `local_jobsimulation_sessions` MIRROR pair (the M219/M222 render-gate trap avoided — RED-proven);
      measured 43 assessed + 2 assigned-only (215 pairs), score spread [27,100] / 68 distinct (rankable);
      only candidates audition; 0 skill refs (closure green). 9 tests green.
- [x] **S5 — Wire hiring rows into `resetTables` + the closure gate + the isolation audit** —
      `organization_sim_invitation_links` added to resetTables (child-first); closure UNAFFECTED (funnel
      writes 0 skill refs); both seeders isolation-clean (per-stack postgres, `AssertClean` passes).
- [x] **S6 — SATISFIED by S2** (`organization_sim_invitation_links` — the 5 positions/org). NOT skipped:
      the config seeder writes them (bounded 5/org via the table's UNIQUE(sim,org)). The comparison reads
      sessions not links, so they are believability + auditability, not a gate. (decisions.md D3.)
- [x] **S7 — Docs:** `seeding-spec.md` (the recruiter-vantage section + the M223 status para),
      `stories-spec.md` (the M223 hiring chain — 2 seeders + the 4-org roster refs), `snapshot-spec.md` (the
      `job_position`-replay drop note — a type-filtered read of the captured `simulations`, no new surface),
      `hiring.md` (the M223-implemented pointer).

## M223: Final Review

_Close review (2026-07-16). M223 shipped essentially clean — the build was verified (17→19 hiring test funcs,
mirror-pair fence RED-proven, full suite green, 0 platform edits). The close re-ran every gate._

### Scope
- [x] All sections S0–S7 checked off. S3 = the pre-existing **M222 D4/Fate-3 drop** (no `job_position` replay —
      0 captured rows, unread by the scoreboard); S6 = folded into S2 (invitation links; done, not deferred).
      One forward-note recorded for **M224** (its scope, not a deferral): `roleForHero` returns `member` for
      end-user heroes, so the funnel's `roleForIndex` candidate-selection is correct only because M223's story
      is **heroless** — M224 must account for it when it adds the cockpit hero trio.

### Code Quality
- [x] The two seeders (`hiring_config.go`, `hiring_funnel.go`) follow the AI-readiness/persona seeder patterns
      (config-seeder + funnel-seeder). No dead code, no leaked handles. gofmt + go vet clean.

### Tests & Benchmarks
- [x] Full `stack-seeding` suite GREEN (13 packages, `-count=1`). gofmt + vet clean.
- [x] **Flake gate 5/5** (seeders pkg, random-order `-shuffle=on`) — 0 flakes.
- [x] **The mirror-pair fence** (`TestHiringFunnelSeeder_WritesMirrorPair`) present + RED-provable (disabling the
      `local_jobsimulation_sessions` mirror write fails it with the M219-trap message). The single most important
      guardrail — it stays.

### Adversarial review
- **Scenario: the captured snapshot has < 5 HIRING sims.** M222 measured 87, but a re-captured/thinner snapshot
      could drop below 5. The reader must **fail loud** (not pad from the generic pool — the M219 anti-pad
      discipline). Recorded in decisions.md; the reserved-disjoint + type-filtered reader means a short pool
      surfaces as a seed error, never as generic sessions masquerading as hiring assessments.

### Decision Triage
- [x] D1–D5 (M223 decisions): the seeder contract + the mirror-pair + the funnel-shape are already in
      `corpus/services/hiring.md` (M222) + `seeding-spec.md`/`stories-spec.md` (M223 S7). Archive the rest.

## M223: Completeness Ledger (section)

**All In-scope items delivered (Fate-1).**
- **Done (Fate 1):** S1 (4th story Meridian Talent) · S2 (HiringConfigSeeder + type-aware reader + disjoint
  reservation) · S4 (HiringFunnelSeeder — the mirror pair, 43-on-5 + 2 assigned-only, [27,100]/68-distinct
  spread, 0 junk) · S5 (resetTables + closure + isolation wiring) · S6 (invitation links, folded into S2) ·
  S7 (docs).
- **Confirmed-covered (Fate 2):** the render proof + cockpit heroes + Clerkenstein `isHiring` → **M224**;
  coverage/playthrough → **M225**; latency → **M226**. All in their `overview.md` In: lists.
- **Annotated (Fate 3):** the `job_position`-replay drop → **M223 overview** (the M222 BA-6 refinement, applied).
- **Dropped:** 0. **Escape-hatch:** 0 → no user sign-off needed.
