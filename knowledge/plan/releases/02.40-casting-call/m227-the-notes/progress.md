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
