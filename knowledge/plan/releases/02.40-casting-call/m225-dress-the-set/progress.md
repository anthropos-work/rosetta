# M225 — Progress

_Section checklist, derived from `overview.md` § Scope.In. To be worked by `/developer-kit:build-milestone`._

## Sections

- [ ] **S0 — KB-fidelity gate** (pre-milestone; `/developer-kit:audit-kb-fidelity`)
- [ ] **S1 — Fold the HIRING-sim (`SIMULATION_TYPE_HIRING`) capture + replay into the auto-set-dress pass**
      (default `/demo-up`) — **NO `job_position` replay** (0 rows captured, unread by the scoreboard; M222 BA-6 /
      M223 D4, reconciled at M225 KB-1); the 5 positions are 5 real captured HIRING sims via `readHiringSimPool`
- [ ] **S2 — Hiring coverage manifest** wired into `manifestFor(vantage, expectedOrg, identityKey)` (persona
      self-consistency role↔skills↔score + the compare-surface sections + 0 prod-eject)
- [ ] **S3 — `playthroughs/manifest/hiring.yaml`** (recruiter compares candidates on a shared sim) + the hiring org
      into the decoupled `pt-world` seed → **one GREEN playthrough**
- [ ] **S4 — Docs:** the hiring sections of `coverage-protocol.md` + `playthroughs.md`
