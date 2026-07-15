# M223 — Progress

_Section checklist, derived from `overview.md` § Scope.In. To be worked by `/developer-kit:build-milestone`._

## Sections

- [ ] **S0 — KB-fidelity gate** (pre-milestone; `/developer-kit:audit-kb-fidelity` — needs M222's `hiring.md`)
- [ ] **S1 — The 4th `stories[]` entry** (`narrative: hiring`, RoleMix ≈ 0.1 admin / 0.9 candidate, hero-trio
      placeholder)
- [ ] **S2 — `HiringConfigSeeder`** (the 5 shared HIRING sims) + the **type-aware hiring-sim reader**
      (`type=HIRING AND job_position NOT NULL`) + the disjoint reserved-tail reservation
- [ ] **S3 — Snapshot extension: replay `directus.job_position` rows** (all 443 public) + pin the 5 chosen HIRING
      sims (REAL replayed content per D-DESIGN-2)
- [ ] **S4 — The `candidate`-assessment funnel seeder** (5 admins + 45 candidates; MOST on all 5 shared sims, SOME
      assigned-not-taken; a differentiated non-degenerate score distribution; the M219 ladder + closure green)
- [ ] **S5 — Wire hiring rows into `resetTables` + the closure gate + the isolation audit**
- [ ] **S6 — (optional) `organization_sim_invitation_links`** (faithful; not a gate — comparison reads sessions)
- [ ] **S7 — Docs:** `snapshot-spec.md` (job_position replay) + `seeding-spec.md`/`stories-spec.md` (the 4th story +
      the funnel)
