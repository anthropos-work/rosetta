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

## Live proof (fresh LOCAL demo-1, cached images; never billion)

All three sections proven live on a fresh `up-injected.sh 1 --no-public-host` (17 containers incl. the two-app
`demo-1-hiring-app`):

- **S1 guard LIVE-GREEN** — autoverify printed `✓ hiring org set-dressed: 5 shared positions + 294 candidate
  HIRING sessions`. (The autoverify fake-FAPI curl WARN is a macOS-curl-vs-Go-TLS false-negative — a browser
  reaches fake-FAPI `/v1/environment` 200 + hiring `/enterprise/activity-dashboard` 200/402ms.)
- **S2 FULLY LIVE-GREEN — all 3 hiring seats GATE MET ✅** — rae-recruiter (manager @ Meridian Talent,
  reachable=53/90, frontier EXHAUSTED) + cara-assessed + cody-assigned, each failingSections=0,
  personaFailures=0 (profileGated), escapes=0. Live sweep surfaced + fixed one real bug (the crawler landed on
  `/`, which apps/hiring's root won't route → land on the manifest's first seedPath). Manifests flipped
  calibrated:true.
- **S3 GATE MET ✅ (the milestone headline)** — `run-playthroughs.sh 1 --reset --grep pt-hiring-recruiter-compare`
  → reset-to-seed pt-world (Org D Kestrel Hiring) + roster refresh + sentinel reload → `✓
  @pt:pt-hiring-recruiter-compare … 1 passed (3.4s)`. The recruiter logs in → apps/hiring Results → the isHiring
  re-skin + shared positions render with a candidate cohort.

_rext tag `casting-call-m225-sections` (moved to `b17756f` after the S2 live-calibration fix), pushed;
consumption copy synced; `.agentspace/rext.tag` points at it. **ZERO platform-repo edits.**_
