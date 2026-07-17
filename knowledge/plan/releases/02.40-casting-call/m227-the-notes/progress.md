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
- [ ] Docs: `corpus/services/hiring.md` + specs updated (Phase 5)
- [ ] 0 platform-repo edits verified
- [ ] Pre-existing unrelated working-tree changes untouched

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
