# M227 — Progress

Section checklist. Check off tasks as they land; commit after each section.

## Section 1 — Hiring-only content ✅
- [x] Trace + confirm the non-hiring-sim surface (recruiter AI-Sim list reads the mirror; PersonaSeeder pollutes it)
- [x] Gate the generic workforce activity seeders on `st.IsHiringOrg()` (Persona + JobsimSessions + Skillpath + Assignments + Activity + HeroActivity)
- [x] Shared, documented guard helper (`hiring_scope.go` `skipGenericActivityForHiringOrg`)
- [x] Regression tests: a hiring org writes ZERO generic-activity rows; workforce orgs unchanged
- [x] `go test ./...` (stack-seeding) + vet green
- [x] Commit (rext 8ed8791)

## Section 2 — External candidate emails ✅
- [x] Key email domain on ROLE via single-sourced `emailForMember` (candidate → external bank; else org domain)
- [x] Preserve Clerkenstein roster email == seeded user email == login (roster.go + users.go single-source; `login:` is a display hint, auth uses the key)
- [x] Cara/Cody heroes → external addresses; manifest regenerated (honesty gate green, header preserved)
- [x] Tests: candidate emails external, employee emails org-domain, roster==user email
- [x] Commit (rext 9aa024b)

## Section 3 — 1 sim per candidate + gate retune ✅
- [x] `HiringFunnelSeeder`: each assessed candidate → exactly 1 position (round-robin even split); mirror pair kept
- [x] Assessed candidate hero's assignment + scored session on the SAME position
- [x] Computed per-position distribution (43 assessed → min 8 / max 9) → floor 6 (margin 2)
- [x] Retuned the gate EVERYWHERE: `hiringComparableFloor`, RENDER_GATE_FLOOR, run-hiring-render.sh, playthrough note, D1, M226/M228 exit gates, hiring.md/seeding-spec/stories-spec, roadmap
- [x] Tests: 1 session/candidate, even split, minPos>=floor, closure green
- [x] Commit (rext 877b091 + rosetta docs)

## Section 4 — Gender-consistent avatars ✅
- [x] `assets/avatars.go`: gender partition of the 12 faces (F={0,2,4,6,9} M={1,3,5,7,8,10,11}) + build-time fence + test
- [x] `avatar.go`: `photoAvatarDataURIForName(seed, firstName)` + `gender.go` curated `inferGender` dictionary
- [x] Thread `firstName` at ALL 3 call sites (`users.go`, `generated_batch.go`, `roster.go`)
- [x] Determinism preserved; Unknown-name fallback byte-identical to the old full-pool pick
- [x] Tests: name→gender→face-subset (200 seeds); heroes matched; menu==profile; Unknown byte-identical
- [x] Commit (rext 63c3e8d)

## Section 5 — Local re-prove (partial — data proven; live render env-blocked → M228)
- [x] Tag the M227 rext (`casting-call-m227-sections`, pushed origin); bump `.agentspace/rext.tag`; sync consumption clone
- [x] **DATA correctness of all 4 fixes proven DETERMINISTICALLY** by the unit + regression suite (exact invariants):
      fix#1 `TestGenericActivitySeeders_SkipHiringOrg` · fix#2 `TestCandidateEmails_RosterMatchesSeed` ·
      fix#3 `TestHiringFunnelSeeder_Funnel` (1/candidate, ~8/position, min 8 >= floor 6) ·
      fix#4 `TestPhotoAvatarForName_GenderMatched` + `TestHeroAvatar_MenuMatchesProfileAndGender`
- [~] Fresh LOCAL demo bring-up — **BLOCKED by the local Docker environment** (`--purge` removed the working demo's
      images → cold full rebuild → ENOSPC → `builder prune` evicted the go-build cache → ~35-min cold recompiles +
      `buildx` wedges under host CPU contention). Root-caused + fixed the demopatch **G6 refusal** (consumption clone
      + registry `type:demo` row). Bring-up re-run in progress from the consumption clone @ M227.
- [ ] Hiring coverage sweep + playthrough GREEN — **routed to M228 "second night"** (the billion live-render
      re-prove on a clean VM; Fate-2, already planned). See decisions.md D5.
- [x] Commit the re-prove record (docs + decisions)

## Cross-cutting
- [x] Docs: `corpus/services/hiring.md` + specs updated — fix #3 (gate retune) landed during build; **fixes #1/#2/#4
      mechanisms blended at close Phase 5** (hiring-only #M227-D1 + external emails #M227-D2 into `stories-spec.md`
      + `seeding-spec.md` + `hiring.md`; gender avatars #M227-D4 into `profile-completeness-spec.md`)
- [x] 0 platform-repo edits verified (all branch changes under `corpus/` or `knowledge/`; rext tooling in its own repo)
- [x] Pre-existing unrelated working-tree changes untouched (`.gitignore` + `_email-assets/*.png` never staged)

## M227: Final Review

_Close-milestone review (2026-07-17). Deterministic suites GREEN; build+harden already did the deep code/test review._

### Scope
- [x] Section 5 live re-prove → **Fate 2 → M228 "second night"** (already planned; M228 exit gate covers the
      corrected-data live render). Deferral audit **YELLOW** — `audit-deferrals/deferral-audit-2026-07-17-m227-close.md`.

### Code Quality
- [x] go vet clean (stack-seeding); 13 packages green; no dead code / boundary issues beyond harden's documented residuals.

### Documentation (Phase 3/5 — the one real finding)
- [x] Fix #1 hiring-only scoping (#M227-D1) blended → `stories-spec.md` + `seeding-spec.md` + `hiring.md`
- [x] Fix #2 external candidate emails (#M227-D2) blended → `stories-spec.md` + `seeding-spec.md` + `hiring.md`
- [x] Fix #4 gender-consistent avatars (#M227-D4) blended → `profile-completeness-spec.md` (+ pointers)
- [x] Cross-ref anchors verified resolvable

### Tests & Benchmarks
- [x] Flake gate 5/5 (seeders pkg, shuffle, `-count=1`); tsc `stack-verify/e2e` exit 0. Harden added 4 RED-proven fences.

### Decision Triage
- [x] D1 → blend into `stories-spec.md`/`seeding-spec.md`/`hiring.md` (hiring-only mechanism) — DONE
- [x] D2 → blend into `stories-spec.md`/`seeding-spec.md`/`hiring.md` (role-keyed external email) — DONE
- [x] D3 → already blended during build (gate retune, all specs) — verified accurate
- [x] D4 → blend into `profile-completeness-spec.md` (gender avatars) — DONE
- [x] D5 (open) → archive as the Fate-2 record (live re-prove → M228); recorded in the deferral audit

## Completeness Ledger (section variant)

Every scope item placed into exactly one three-fate category. No "backlog / follow-up / later" — those aren't fates.

### Done (Fate 1 — landed in this milestone)
- **S1 — Hiring-only content** (#M227-D1): 6 generic activity seeders skip a hiring org (`hiring_scope.go`);
  regression-fenced (ZERO generic-activity rows for a hiring org). rext `8ed8791`.
- **S2 — External candidate emails** (#M227-D2): role-keyed `emailForMember`; login == `public.users` == roster;
  Cara/Cody external + manifest regenerated. rext `9aa024b`.
- **S3 — 1 sim per candidate + gate retune** (#M227-D3): round-robin 1 position/candidate (~8/position); compare
  gate retuned `≥40 → ≥6` EVERYWHERE (Go const · render-probe · playthrough · D1 · M226/M228 gates · 3 corpus docs ·
  roadmap); mirror pair + closure kept green. rext `877b091` + rosetta docs.
- **S4 — Gender-consistent avatars** (#M227-D4): name→gender→face-subset, all orgs; Unknown → byte-identical old pick.
  rext `63c3e8d`.
- **Close docs (Phase 5)**: fix #1/#2/#4 mechanisms blended into `stories-spec.md` + `seeding-spec.md` + `hiring.md`
  + `profile-completeness-spec.md` (fix #3 was already documented; verified accurate).
- **Harden**: 4 RED-proven fences (the load-bearing UsersSeeder population write-path + score clamps + Go↔probe
  floor-drift). rext `78a3cb2` (test-only).

### Confirmed-covered (Fate 2 — already owned by another milestone of this release)
- **Section 5 — the LOCAL live full-stack render/coverage/playthrough re-prove on the corrected data → M228
  "second night"** (already planned in `roadmap.md`; M228's exit gate explicitly covers the corrected-data live
  render — 1-sim/candidate ≥ floor 6, external emails, matched avatars, hiring-only content). The DATA correctness
  of all 4 fixes is proven DETERMINISTICALLY here; only the live full-stack render is routed (env-blocked local box).
  No plan edit needed — M228 owns it. See `audit-deferrals/deferral-audit-2026-07-17-m227-close.md` (DEF-M227-01).

### Annotated (Fate 3 — attached to a release-milestone at close)
- None new. (Inherited DEF-M226-01 — pre-bind serve reap — was already annotated to the next prove-on-VM run at M226.)

### Dropped
- None.

### Release-scope-breaking deferral (escape hatch — requires user sign-off)
- None.

### Carried known issues (inherited, out of M227 scope → v2.4 RELEASE close)
- 8 pre-existing demo-stack test failures (6× `test_cockpit.py` + `test_purge` + `test_reap`) — HEAD-identical,
  predate v2.4, in files M222–M227 never touched → the release close (test-debt harden pass). Re-confirmed fresh.
- M204 `assign-and-track.UC1` assign-WRITE declared in-manifest TODO → the release close (declared-TODO fate).

> **Ledger verdict: clean.** All in-scope items landed Fate 1; Section 5's live render is a clean Fate 2 to an
> already-planned M228. **Zero escape-hatch deferrals → no user sign-off required.** Deferral audit YELLOW.

## M227: Hardening

### Pass 1 — 2026-07-17
**Scope manifest (M227-touched code, `7032aea..HEAD`):** 26 files. Source touched (all rext
`stack-seeding/`): fix#1 `hiring_scope.go` + 6 gated seeders (persona/activity/assignments/hero_activity/
jobsim_sessions/skillpath_sessions); fix#2 `userprofile.go` (emailForMember/externalCandidateDomain),
`users.go`, `roster.go`; fix#3 `hiring_funnel.go`, `contentref.go`; fix#4 `gender.go`, `avatar.go`,
`assets/avatars.go`, `generated_batch.go`. Existing tests: `hiring_scope_test.go`, `gender_test.go`,
`assets/gender_test.go`, `candidate_email_test.go`, `hiring_funnel_test.go`.

**Coverage delta (milestone-touched files, `go test -covermode=set`):**
- `seeders` pkg: 96.7% → 96.8% statements (already high — harden value is behavioral, not line count).
- `hiringSessionScore`: 71.4% → **100%** (clamp branches closed).
- Residual uncovered: embed-error fallbacks in `photoAvatarDataURIForName` (85.7%) / `mustAvatarNames` /
  `init` panic — unreachable at runtime unless the go:embed is broken; not forced (build-time bugs).

**The gap hardening found (the load-bearing one):** the build-phase fix#2/#4 tests RE-DERIVE the helper
(`emailForMember` / `photoAvatarDataURIForName`) and compare it to the Clerkenstein ROSTER — **neither runs
the `UsersSeeder`**. So a revert of the two `users.go` call sites would ship SILENTLY: candidates flip to
the org domain and a "Sara" to a man's photo, with every existing test green (roster.go is a separate write
path). The whole seeded population (60 users) had ZERO write-path fence for fix#2/#4.

**Tests added (4 new, all RED-proven then reverted):**
- `users_m227_test.go`
  - `TestUsersSeeder_PopulationEmailRoleKeyedAndAvatarGenderMatched` — drives the REAL UsersSeeder over the
    hiring (5 admin + 45 candidate) + workforce population; asserts on the actual `public.users` /
    `public.memberships` COPY rows: `isExternalDomain==isCandidate`, avatar face ∈ inferred-gender set,
    `membership.email==users.email`, `membership.picture_url==users.picture`. (1 regression, whole-pop.)
  - `TestUsersSeeder_HeroSeededEmailAndAvatarEqualRoster` — heroes' SEEDED `public.users` email/picture ==
    the ACTUAL roster entry (the byte-identical login==postgres==menu triple, both real write paths — not
    a re-derivation). (1 regression.)
  - `TestHiringComparableFloor_MatchesRenderProbeDefault` — cross-surface drift fence: reads the render-probe
    spec (cwd-independent via `runtime.Caller`) and asserts its `RENDER_GATE_FLOOR` default ==
    `hiringComparableFloor` (Go const + TS spec have no shared source). (1 drift fence.)
- `hiring_funnel_test.go`
  - `TestHiringSessionScore_ClampsBounds` — the floor/ceiling clamp branches (defensive; the [30,95]±jitter
    band never fires them) swept over 12×5 (candidate,position) coords. (1 edge/error-path.)
  - Fixed the stale `TestHiringFunnelSeeder_Funnel` header comment ("all 5 positions" → the M227 1-position
    reality).

**RED-prove evidence (each fence catches its re-break, then reverted):**
- Go const `6→7`: drift fence FAILED ("floor DRIFTED: spec=6, Go=7"). ✓
- floor clamp disabled: clamp fence FAILED (`apt=0 → -4/-1/3, want 15`). ✓
- `emailForMember → emailFor(...,orgDomain,i)`: population fence FAILED
  (`liam...@meridian-talent.com … want external`) + hero fence FAILED (roster `cara...@hotmail.com` !=
  seeded `cara...@meridian-talent.com`) — the exact silent divergence. ✓
- `photoAvatarDataURIForName → photoAvatarDataURI`: population fence FAILED ("Felix, inferred Male, non-male
  face") + hero menu==profile FAILED. ✓

**Bugs fixed inline:** none — no production bug surfaced; the 4 fixes are correct. Hardening added fences
only (+ a stale-comment fix).

**Flakes stabilized:** none — 3 consecutive clean sequential runs of the new tests (deterministic seeders).

**Knowledge backfill:** no new KB doc warranted. The floor-surface list (Go const · render probe ·
playthrough · `corpus/services/hiring.md`) already lives in `spec-notes.md` fix#3 + `hiring.md`; the new
drift fence now *enforces* the Go↔render-probe pair the docs describe. The "two write paths must agree"
invariant is already in decisions D2/D4 + the code comments.

**rext commit:** `78a3cb2` (authoring clone). Test-only additions — no stack-consumed tooling changed, so
the sections tag `casting-call-m227-sections` stays the consumption point (no new tag). 13 pkgs green.

### Stop condition
Loop stopped after **pass 1**: the full Step-2b six-dimension scan found the one load-bearing gap (the
whole-population write-path) + the clamp edge-path + the cross-surface drift and closed all three; the
remaining uncovered lines are unreachable embed-error fallbacks (Fate-1-inappropriate to force). Coverage
delta < 2% (already saturated), no flakes. The **LOCAL live-render re-prove stays deferred to M228** (Fate-2,
the billion venue) — harden worked the deterministic surface only, per its scope.
