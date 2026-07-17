# M227 — Progress

Section checklist. Check off tasks as they land; commit after each section.

## Section 1 — Hiring-only content
- [ ] Trace + confirm the non-hiring-sim surface (recruiter AI-Sim list reads the mirror; PersonaSeeder pollutes it)
- [ ] Gate the generic workforce activity seeders on `st.IsHiringOrg()` (Persona + JobsimSessions + Skillpath + Assignments + Activity + HeroActivity; verify which write jobsim/mirror rows)
- [ ] Shared, documented guard helper (single source of the rule)
- [ ] Unit/regression tests: a hiring org writes ZERO non-hiring mirror/session rows; workforce orgs unchanged
- [ ] `go test ./...` (stack-seeding) + vet green
- [ ] Commit

## Section 2 — External candidate emails
- [ ] Key email domain on ROLE (candidate → external bank; else org domain) in `users.go` + a `userprofile.go` helper
- [ ] Preserve Clerkenstein roster email == seeded user email == login (trace roster.go + cockpit + preset `login:`)
- [ ] Cara/Cody heroes → external addresses, consistent end-to-end
- [ ] Tests: candidate emails external, employee emails org-domain, roster==user==login
- [ ] Commit

## Section 3 — 1 sim per candidate + gate retune
- [ ] `HiringFunnelSeeder`: each assessed candidate → exactly 1 position (even deterministic split); mirror pair kept
- [ ] Assessed candidate hero pinned to a known position (assignment + scored session agree)
- [ ] Compute exact per-position distribution (test) → set the floor with a safety margin
- [ ] Retune the gate EVERYWHERE `≥40` is asserted (D1, exit gates M226/M228, coverage manifest, playthrough, hiring.md, RENDER_GATE_FLOOR)
- [ ] Tests: 1 session/candidate, ~8/position, closure green
- [ ] Commit

## Section 4 — Gender-consistent avatars
- [ ] `assets/avatars.go`: gender partition of the 12 faces (F={0,2,4,6,9} M={1,3,5,7,8,10,11}) + a test
- [ ] `avatar.go`: `photoAvatarDataURIForName(seed, firstName)` + `inferGender` dictionary/heuristic
- [ ] Thread `firstName` at both call sites (`users.go`, `generated_batch.go`)
- [ ] Determinism/idempotency preserved; unknown-name fallback honest
- [ ] Tests: name→gender→face-subset; heroes matched; re-seed byte-identical
- [ ] Commit

## Section 5 — Local re-prove
- [ ] Tag the M227 rext; bump `.agentspace/rext.tag`; sync consumption
- [ ] Fresh LOCAL demo (offset ports), cold reset-to-seed, consuming the M227 tag
- [ ] Assert from this Mac: ≥floor per each of 5 positions; 1 sim/candidate; hiring-only; external emails; matched avatars
- [ ] Hiring coverage sweep GREEN + hiring playthrough GREEN
- [ ] Commit the re-prove record

## Cross-cutting
- [ ] Docs: `corpus/services/hiring.md` + specs updated (Phase 5)
- [ ] 0 platform-repo edits verified
- [ ] Pre-existing unrelated working-tree changes untouched
