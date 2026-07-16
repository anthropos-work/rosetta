# M225 — Progress

_Section checklist, derived from `overview.md` § Scope.In. To be worked by `/developer-kit:build-milestone`._

## Sections

- [x] **S0 — KB-fidelity gate** (pre-milestone; `/developer-kit:audit-kb-fidelity`) → **YELLOW**; KB-1
      reconciled the stale `job_position` premise (corpus was already correct). Report: `kb-fidelity-audit.md`.
- [x] **S1 — Fold the HIRING-sim (`SIMULATION_TYPE_HIRING`) capture + replay into the auto-set-dress pass**
      (default `/demo-up`) — **NO `job_position` replay** (0 rows captured, unread by the scoreboard; M222 BA-6 /
      M223 D4, reconciled at M225 KB-1); the 5 positions are 5 real captured HIRING sims via `readHiringSimPool`.
      **Finding (D1):** the hiring org already comes up real by default (M223+M224); S1's deliverable = the
      bring-up-tail GUARD (autoverify hiring cheap-win, `is_hiring`-gated, ≥5 positions + ≥40 sessions) + docs.
      rext `eee2113`; 6 new tests (120/120 + shellcheck). LIVE guard-green: at the bring-up.
- [x] **S2 — Hiring coverage manifest** wired into `manifestFor(vantage, expectedOrg, identityKey)` (persona
      self-consistency role↔skills↔score + the compare-surface sections + 0 prod-eject). Recruiter Rae
      (`MANAGER_MANIFEST_HIRING`, apps/hiring Results) + candidate self-views (Cara/Cody); org/identity dispatch
      (the AB4 precedent); `persona-assert` `profileGated` mode; `COVERAGE_APP_PORT_BASE=3001`. rext `88e6fb9`;
      43/43 unit tests. calibrated:false → live-calibrated at the bring-up.
- [x] **S3 — `playthroughs/manifest/hiring.yaml`** (recruiter compares candidates on a shared sim) + the hiring org
      into the decoupled `pt-world` seed → **one GREEN playthrough**. pt-world Org D "Kestrel Hiring Group"
      (distinct test data); recruiter `pt-recruiter`; `pt-hiring-recruiter-compare` on apps/hiring. rext `29ceae6`;
      ptvalidate GREEN (7 products, 15 live + 1 TODO). LIVE recruiter-green: at the bring-up.
- [x] **S4 — Docs:** the hiring sections of `coverage-protocol.md` (the hiring-vantage org/identity dispatch +
      apps/hiring targeting + profileGated persona) + `playthroughs.md` (the hiring product + Org D + count 14→15).

_rext tagged `casting-call-m225-sections` (29ceae6), pushed; consumption copy synced; `.agentspace/rext.tag`
bumped. Fresh LOCAL demo-1 bring-up in progress to prove the S1 guard + S2 coverage gate + S3 recruiter
playthrough live (LOCAL only, never billion)._
