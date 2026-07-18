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

## M225: Hardening

### Pass 1 — 2026-07-17 (`/developer-kit:harden-milestone`)

**Scope manifest (rext diff `66ed56d..HEAD`, 16 files):** the coverage manifest + `manifestFor` 3-arg
dispatch + `profileGated` (`stack-verify/e2e/`), the autoverify hiring cheap-win (`stack-verify/live/` +
`tests/test_verify.py`), the hiring playthrough (`playthroughs/manifest/hiring.yaml`, `HiringResultsPage`,
`hiring-recruiter.spec.ts`, `stack-env.ts`, pt-world Org D seed + `corpus_test.go`). Coverage assessed by
direct branch analysis (read every M225-touched fn + its tests) rather than an instrumented tool — the rext
e2e module carries no line-coverage harness and wiring one for non-navigating Playwright specs is out of
proportion for a section harden; the deterministic Go/Py/TS/bash suites already carry deep unit coverage
(43 coverage-manifest + 6 autoverify + corpus pins + stack-env). The four task-flagged seams were the finder.

**Tests added (10 fences, all Fate-1 LAND-NOW; 0 bugs surfaced — build shipped green):**
- `coverage-manifest.unit.spec.ts` (+5): (1) **every hiring manifest page is `calibrated:true`** — the
  pre-close "no untuned descriptor ships" gate covered only base+showcase manager manifests, leaving the
  M225 `calibrated:false→true` flip (D2) UNGUARDED; (2) the recruiter **Results scoreboard pins the
  5-shared-positions floor** (`reservedHiringSimRefs`; the generic loop only proved SOME floor); (3–5) the
  **`profileGated` apps/hiring wiring** static-fence chain — `run-coverage.sh` forwards
  `COVERAGE_PROFILE_GATED` + honors `COVERAGE_APP_PORT_BASE=3001`; `coverage.spec.ts` threads `profileGated`
  into `runPersonaChecks` + lands on `seedPaths[0]` when gated; `persona-assert` threads `profileGated` to the
  role-skills + avatar checks (the whole `/profile→/home` seam had zero deterministic coverage). Stale
  `calibrated:false` build comment corrected.
- `test_verify.py` (+4): the autoverify **(e) floor boundaries** — positions `==5` pass / `==4` warn, sessions
  `==40` pass / `==39` warn (the build bracketed but never pinned the `-ge` edges).
- `playthroughs/manifest/hiring_isolation_test.go` (new): the **pt-world Org D "test data ≠ demo data"
  invariant (D3)** — reads the real `pt-world.seed.yaml` and pins Org D's hiring shape (Kestrel Hiring Group /
  `narrative: hiring` / `is_hiring: true` / `pt-recruiter`) AND its distinctness (no pt-world org contains the
  demo `HIRING_ORG` "Meridian Talent"; the hiring org is not a "Meridian" variant; Org D is the sole hiring
  org in the world).

**Bugs fixed inline:** none — the build shipped green; this pass added regression fences only.
**Flakes stabilized:** none (all added tests are deterministic unit/static; no live-stack dependency).
**Knowledge backfill:** no KB-worthy findings — every fence pins an already-documented invariant (decisions
D2/D3; the autoverify (e) floors in `verification.md` / `demo-up-defaults.md`). No corpus edit warranted.

**Verify:** `coverage-manifest.unit` 48 pass, `stack-env.unit` 13 pass, `test_verify.py` **124 OK** (incl
`shellcheck`), playthroughs `go test ./...` all ok + `go vet` clean, both e2e `tsc --noEmit` clean.

**Stop condition:** loop terminated after pass 1 — the six-dimension scan found nothing else worth adding
(the browser-only `profileGated` runtime is now statically fenced; no untested public fn remains cheaply
testable; no perf path — that's M226; no flake surface). Coverage delta negligible beyond the targeted fences.

_rext harden commit `be431c3`, tag `casting-call-m225-harden`, pushed (authoring `main`). `.agentspace/rext.tag`
+ the consumption copy left at `casting-call-m225-sections` — the fences are test-only (no runtime tooling
change), so the live demo-1 needs no re-sync. **ZERO platform-repo edits.**_

## M225: Final Review (`/developer-kit:close-milestone`, 2026-07-17)

Cross-cutting close review of the whole milestone. **0 findings requiring a fix** — the build shipped green, the
harden pass added 10 regression fences (0 bugs), and this close's re-verification reproduced every deterministic
suite GREEN at the merge base.

### Scope — all Scope.In delivered Fate-1; 0 gaps
- S1 bring-up guard + docs · S2 coverage manifest (3 seats GATE MET) + `manifestFor` 3-arg + `profileGated` ·
  S3 `hiring.yaml` + pt-world Org D + one GREEN playthrough · S4 3 corpus docs. `overview.md` Out: (the live
  cross-machine proof) is Fate-2 → M226. S3's candidate playthrough is a conscious pillar split (D3), not a gap.

### Code Quality — N/A in this repo (docs only); rext reviewed at build+harden
- The rosetta branch diff is docs only (3 corpus + plan docs). The tooling code lives in the rext authoring clone
  (`be431c3`), reviewed per-section at build + across the harden pass; `go vet` clean, `tsc` clean, shellcheck clean.

### Adversarial — covered by the D1/D2/D3 investigations + the harden fences
- The cold-cache / starved-HIRING-pool silent-empty scenario (the milestone's real failure mode) is fenced at THREE
  layers: the S1 autoverify bring-up guard (warns at the tail), the S2 coverage sweep (empty grid = FAIL), and the
  S3 playthrough (empty scoreboard = FAIL). The pt-world Org D "test data ≠ demo data" invariant is pinned by
  `hiring_isolation_test.go`.

### Documentation — the S4 deliverable; accurate + house-style-consistent
- The 3 corpus edits reviewed against D1/D2/D3 + KB-1: accurate, cross-referenced, and citation-consistent with
  each doc's established style (inline milestone citation in coverage-protocol.md / playthroughs.md; the cheap-win
  heading form in verification.md). No new top-level unit introduced in this repo → no per-unit handbook gap.

### Tests & Benchmarks — re-verified GREEN at merge base
- Go 1885→**1887** (+2) + vet clean · `test_verify.py` **124** (incl shellcheck) · TS unit **61** (stack-verify/e2e)
  + **69** (playthroughs/e2e) · both `tsc` clean · **flake 5/5 clean** on the M225 deterministic fences · the live
  hiring playthrough independently orchestrator-verified GREEN on a clean pt-world reset-to-seed. No benchmark
  surface (perf is M226).

### Decision Triage
- **KB-1** → archive (maintainer-only plan-doc reconciliation).
- **D1** → already blended into `verification.md` (the hiring set-dress cheap-win, S4) — accurate, cited inline.
- **D2** → already blended into `coverage-protocol.md` (the hiring vantage dispatch + profileGated, S4) — accurate.
- **D3** → already blended into `playthroughs.md` (the hiring product + Org D, S4) — accurate.
- **D-AUDIT** (new this close) → archive (records the Phase-1b re-fate of the two inherited carries).

## Completeness Ledger

### Done (Fate 1 — landed in this milestone, properly and completely)
- **S0** KB-fidelity gate (YELLOW; KB-1 reconciled the stale `job_position` premise) — `kb-fidelity-audit.md`
- **S1** auto-set-dress hiring bring-up GUARD (autoverify cheap-win (e)) + docs — verified: `test_verify.py` 124 + live guard-green (`5 positions + 294 sessions`)
- **S2** hiring coverage manifest (`manifestFor` 3-arg org/identity dispatch + `profileGated`) — verified: `coverage-manifest.unit` 48 + all 3 hiring seats GATE MET live (rae/cara/cody, failingSections=0, personaFailures=0, escapes=0)
- **S3** `playthroughs/manifest/hiring.yaml` + pt-world Org D "Kestrel Hiring Group" + `pt-recruiter` — verified: `hiring_isolation_test.go` + `ptvalidate` GREEN + live `pt-hiring-recruiter-compare` GREEN (orchestrator-reverified on clean reset-to-seed)
- **S4** the hiring sections of `coverage-protocol.md` + `playthroughs.md` + the hiring cheap-win in `verification.md` — verified in the branch diff

### Confirmed-covered by another milestone in this release (Fate 2)
- The **live cross-machine (billion/tailnet) proof** of the hiring vantage + the **recruiter p95 click→ACCESS < 5 s** 3rd vantage — **Owned by:** M226 opening-night → `roadmap.md` M226 exit-gate conditions (2)+(5) — **Decision:** `overview.md` §Scope.Out + roadmap M226. (No roadmap edit needed; M226's gate already owns it.)

### Annotated to a milestone of this release (Fate 3)
- None.

### Dropped (cut from roadmap entirely)
- None. (S3's candidate playthrough was non-committed "optional" scope covered on the presence side by S2 — a conscious pillar split per **D3**, not a drop.)

### Carried known-issues (pre-existing / declared — flagged per Phase 4, not M225's work)
- **8 pre-existing demo-stack test failures** (6× `test_cockpit.py` + `test_purge` + `test_reap`) — HEAD-identical, untouched by M225, outside the hiring domain → standing test-debt backlog / a future demo-stack test-debt harden pass — **Decision:** M224 D6 + M225 D-AUDIT; re-fate at v2.4 release close.
- **M204 `assign-and-track.UC1` assign-WRITE TODO** — a declared in-manifest `unimplemented` build-reference gap (surfaced only in the `playthroughs.md` 14→15-live count) → its declared-TODO fate is a v2.4 release-close decision — **Decision:** M225 D-AUDIT.

### Release-scope-breaking deferral (escape hatch — requires user sign-off)
- None.
