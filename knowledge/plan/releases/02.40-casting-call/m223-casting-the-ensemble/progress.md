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
