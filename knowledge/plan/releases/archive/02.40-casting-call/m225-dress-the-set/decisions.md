# M225 — Decisions

_Implementation choices with rationale. One entry per decision; cite the code/doc it binds._

---

## KB-1 — Phase 0b KB-fidelity: reconcile S1's stale `job_position`-replay premise (M222 BA-6 / M223 D4)

**Finding (Phase 0b, 2026-07-16 — YELLOW).** The M225 scaffold (`overview.md` Scope.In #1, `progress.md` S1,
`spec-notes.md`) said S1 folds "the `directus.job_position` replay + the 5-sim capture" into auto-set-dress. That
premise was **refuted before M225 even opened**: M222 BA-6 measured `directus.job_position` = **0 rows captured**
(the prod "443" was never in a snapshot) and the recruiter scoreboard does **not** read `job_position`
(`JobSimulation.jobPosition` is optional/unused); M223 D4 formally **DROPPED** the `job_position` replay. The 5
"positions" ARE 5 real captured `SIMULATION_TYPE_HIRING` sims, resolved by `readHiringSimPool`.

**The corpus was already correct** — `snapshot-spec.md:386`, `seeding-spec.md:352`, `stories-spec.md:633`, and
`hiring.md:75` all document the drop. Only the M225 **plan docs** carried the stale framing.

**Decision:** reconcile the plan docs inline (this audit) to the real S1 scope — **fold the HIRING-sim
(`SIMULATION_TYPE_HIRING`) capture + replay into the default auto-set-dress** so the hiring org's positions +
content come up real with no manual steps; **no `job_position` replay / no `snapshot-spec.md` surface**. Not a
scope change (the release already decided this in closed M222/M223) — plan hygiene.

**Binds:** `overview.md` §Why-section + §Scope.In #1 · `progress.md` S1 · `spec-notes.md` §Auto-set-dress ·
rext `stack-seeding` `readHiringSimPool`/`HiringConfigSeeder` (M223) · `presets/stories.seed.yaml` hiring story.

---

## D1 — S1: the hiring org ALREADY comes up real on a default `/demo-up`; S1's genuine deliverable is the bring-up-tail GUARD + docs (there is no separate "5-sim capture" to fold in)

**Investigation (Phase 1).** Traced the full default `/demo-up` bring-up chain and found the hiring set-dress is
**already end-to-end default-on** — most of the scaffolded S1 work was delivered incidentally by M223/M224:
- The hiring org (Meridian Talent, `is_hiring=true`) is in the DEFAULT preset `presets/stories.seed.yaml` (M223),
  and a bare `/demo-up N` seeds the stories preset by default (M38; `up-injected.sh` `STORIES_PRESET`).
- `resolveContentRefs` runs `readHiringSimPool` UNCONDITIONALLY (`contentref.go:177`); `HiringConfigSeeder` +
  `HiringFunnelSeeder` are registered by default (`stackseed/main.go:430,451`).
- **The `SIMULATION_TYPE_HIRING` sims ride along in the STANDARD directus content-surface capture.** The directus
  surface captures the `simulations` table INCLUDING its `type` column (`stack-snapshot/directus/directus.go:190`)
  under the public firewall (`private=false AND tenant_id IS NULL AND status='published'` — M222 measured 87 public
  hiring sims). So the auto-set-dress directus replay lands them; `readHiringSimPool` resolves 5. **There is NO
  separate "5-sim capture" step, and NO `directus.job_position` table surface** (0 rows; `job_position` exists only
  as a *column* of `simulations`, unread by the scoreboard — M222 D4).
- The two-app hiring UI container (`build_frontend_hiring`) + its 4 demo-patches are default-on (M224;
  `build_frontends` at `up-injected.sh:1530`, only `DEMO_NO_UI=1`/`DEMO_NO_PATCH=1` opt out).

**So no hiring-specific manual step exists to REMOVE.** The one residual risk is the general **cold-cache** case: a
cold/empty snapshot cache (or a starved HIRING pool) leaves `readHiringSimPool` empty → the seeders HONESTLY
degrade to 0 positions / 0 sessions → the recruiter comparison renders EMPTY while the stack still says UP. M223's
adversarial review named the downstream M224/M226 render gate as the loud catch.

**Decision:** S1's genuine, non-redundant deliverable is (a) a **bring-up-tail GUARD** that brings that catch
FORWARD to the default `/demo-up` tail, and (b) the documentation fold-in. Built the `autoverify.sh` demo-only
cheap-win (e) — the exact shape of the ISSUE-7 casbin assert: gated on a hiring org existing (`is_hiring=true`, so a
hiring-less demo SKIPS, no false-warn), it asserts ≥5 `organization_sim_invitation_links` positions + ≥40
`local_jobsimulation_sessions` candidate sessions for the hiring org; else WARNS (non-fatal) that the comparison is
empty (cold cache / starved pool). So "the hiring org comes up real with no manual steps" is now a CHECKED property
of the default bring-up, not an assumed one. 6 new tests (`test_verify.py::TestAutoVerify`), 120/120 green +
shellcheck-clean.

**Binds:** rext `stack-verify/live/autoverify.sh` (e) · `stack-verify/tests/test_verify.py` · doc fold-in
(`snapshot-spec.md` / `recipe-snapshot-world.md` / `verification.md` / `demo-up-defaults.md`, Phase 5).

**Not gold-plating:** no invented plumbing — the finding that S1 was largely pre-delivered is a real result; the
value added is making it self-verifying (the guard) + documented.

---

## D2 — S2: the hiring coverage manifest reuses `manifestFor` org/identity dispatch (the AB4 precedent); a `profileGated` persona mode adapts to apps/hiring (not a fork)

**Decision:** extend the M42 coverage machinery (`coverage-manifest.ts` + `manifestFor` + `coverage.spec.ts` +
`persona-assert.ts` + `run-coverage.sh`) with a HIRING vantage — **never forked**, only extended:
- `MANAGER_MANIFEST_HIRING` (recruiter Rae): the compare surface `/enterprise/activity-dashboard` — the 5 shared
  positions render as custom-tanstack-table rows (`tbody.tbody > tr.tr`, the M224 R4 selector, NOT AntD) + the
  `isHiring` "Results" re-skin. The per-sim ranked-candidate DRAWER (0-junk + non-degenerate distribution — the
  cohort-level role↔skills↔score self-consistency) stays with `render-hiring-comparison.spec.ts` (M224).
- `EMPLOYEE_MANIFEST_HIRING_ASSESSED`/`_ASSIGNED` (Cara/Cody): the candidate `/home` self-views (apps/hiring
  /profile is admin-gated → redirect to /home). Per-candidate role↔score self-consistency: assessed → a
  completed+scored position; assigned-only → a pending position, no score.
- `manifestFor` is now **manager-org-conditional** (Meridian Talent → hiring, else showcase/base) **and
  employee-identity-conditional** (hiring candidate seats → their self-views), hiring checked FIRST. `HIRING_ORG`
  substring-matched case-insensitively (the AB4 convention). No false-promotion: "Meridian Labs" (pt-world Org A) +
  the showcase org + base orgs all stay put (unit-tested).

**`profileGated` persona mode (not a fork).** `coverage.spec.ts` gates on `personaFailures===0` and runs
`runPersonaChecks` unconditionally; persona-assert targets next-web `/profile/skills`, which apps/hiring
admin-redirects. So `runPersonaChecks(…, {profileGated})` adapts: `roleSkillsCoherence` + `avatarConsistency` read
the `/home` self-view (assert no-junk + real-photo) instead of the next-web `/profile*` pages. Wired via
`COVERAGE_PROFILE_GATED` (`coverage.spec.ts`) + `COVERAGE_APP_PORT_BASE=3001` (`run-coverage.sh` → the hiring app).

**calibrated:false during build.** The section floors/copy are AUTHORED from the M224 render evidence + the seed
contract, and CALIBRATED against the live apps/hiring render at the M225 shared bring-up (the coverage-protocol
discipline — never assert an unrendered string; the M219 "wrote-and-never-ran" trap). The pre-close gate flips them
to true.

**Binds:** rext `stack-verify/e2e/lib/coverage-manifest.ts` (`manifestFor` + hiring manifests) ·
`coverage-manifest.unit.spec.ts` (dispatch tests) · `persona-assert.ts` (profileGated) · `coverage.spec.ts` ·
`run-coverage.sh` · doc: `corpus/ops/demo/coverage-protocol.md` hiring section (S4).

---

## D3 — S3: the hiring playthrough reuses the M202 machinery; a DISTINCT pt-world hiring org (Org D); the recruiter surface is apps/hiring

**Decision:** add the FOURTH product (Hiring) to the Playthroughs corpus — **never forked**, reusing
`hero-login`/`resolveStackEnv`/`PageObject`/`ptvalidate`/`ptreport`/`run-playthroughs.sh`:
- **pt-world Org D "Kestrel Hiring Group"** (`narrative: hiring` + `is_hiring: true`, size 40 → 4 admin + 36
  candidates). **Deliberately distinct** from the demo's "Meridian Talent" AND this world's Org A "Meridian
  Labs" (test data ≠ demo data; and no "Meridian" prefix collision with the S2 coverage `HIRING_ORG` gate). The
  two `narrative`/`is_hiring` flags trigger the SAME `HiringConfig`/`HiringFunnel` seeders (no bespoke pt-only
  seeder — the Org C ai-readiness precedent). Recruiter hero `pt-recruiter` (Quinn, **Talent Acquisition
  Specialist** — a resolvable role with role-skills, the M224 iter-04 non-resolving-role lesson). Seed
  validates: 5 orgs / 7 heroes / pop 190; cockpit export emits `is_hiring:true` → the two-app hiring-base route.
- **The recruiter surface is apps/hiring, not next-web** (the M224 two-app finding). Added
  `resolveStackEnv().hiringAppBaseUrl` (3001+offset, `PT_HIRING_BASE_URL` override) + `run-playthroughs.sh`
  exports it; `HiringResultsPage` reuses the M224 render-probe's calibrated tanstack anchor
  (`tbody.tbody > tr.tr` — NOT AntD); `hiring-recruiter.spec.ts` logs in on the hiring base + asserts the
  isHiring "Results" re-skin + the shared positions render with a candidate cohort (an empty grid = a cold
  cache / starved pool, a FAILURE).
- **Scope: recruiter only** (one GREEN playthrough = the gate). The candidate is "optional" per the overview,
  and is covered on the PRESENCE side by S2's candidate coverage manifests — a clean pillar split (S2 =
  presence, S3 = function).

**Deterministic gates all GREEN:** ptvalidate (7 products, 15 live Playthroughs + 1 TODO, both-way integrity +
precondition-coverage) · `go test ./...` · tsc · 69 e2e unit tests · shellcheck · `stackseed --validate`. The
one GREEN live recruiter run lands at the shared bring-up.

**Binds:** rext `playthroughs/seed/pt-world.seed.yaml` (Org D) · `seed/seed-worlds.yaml` (roster + caps) ·
`manifest/hiring.yaml` · `manifest/corpus_test.go` (M225 pin) · `e2e/lib/{stack-env,hiring-results-page}.ts` ·
`e2e/tests/hiring-recruiter.spec.ts` · `e2e/run-playthroughs.sh` · doc: `corpus/ops/demo/playthroughs.md` (S4).

---

## D-AUDIT — Close Phase 1b deferral re-audit: YELLOW (2 inherited carries, 0 new; both re-fated fresh, routed to release close)

**Finding (close Phase 1b, 2026-07-17 — YELLOW).** The `/developer-kit:audit-deferrals --scope=milestone` pass
found **M225 introduced ZERO new deferrals** (every Scope.In item landed Fate-1; the Out: live proof is Fate-2 to
M226; S3's candidate-playthrough is a conscious pillar split per D3, not a defer). Two INHERITED carries remain, both
consciously tracked with destinations + standing sign-off:

- **DEF-CARRY-A — 8 pre-existing demo-stack test failures** (6× `test_cockpit.py` + `test_purge` + `test_reap`).
  Inherited-failure carry from **M224 D6**; HEAD-identical, in files M225 never touched, outside the hiring domain.
  **Fresh fate (today): KEEP-DEFERRED (carry)** → standing test-debt backlog / a future demo-stack test-debt harden
  pass. Re-fate explicitly at v2.4 release close.
- **DEF-CARRY-B — the M204 `assign-and-track.UC1` assign-WRITE TODO** (a declared in-manifest `unimplemented`
  build-reference gap, carried since v2.0). Surfaced in M225's `playthroughs.md` count only (14→15 live, the same
  1 TODO). **Fresh fate (today): KEEP-DEFERRED (declared TODO)** → its declared-TODO fate is a v2.4 **release-close**
  decision.

**No blocking items** (no repeat-deferral of promised milestone work). Verdict YELLOW → the milestone close proceeds;
both carries inherit into the v2.4 close-release Phase 1b audit (release scope, extra-scrutiny).

**Binds:** `audit-deferrals/deferral-audit-2026-07-17-m225-close.md` · M224 `decisions.md` D6 · `state.md` standing
backlog · `corpus/ops/demo/playthroughs.md` (the "15 live, 1 TODO" line).

---
