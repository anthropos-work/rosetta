# Hardening Ledger — M50 Content & seeding fill

## Pass 1 — 2026-06-30 — final

**Iters hardened this pass:** all milestone-touched code (final cumulative scope) — focused on the Go/testable surface (the seeders) per the milestone's static-only harden constraint.
**Tiks covered since prior pass:** all iters in milestone (first harden pass; ledger did not exist).
**Coverage delta on touched files:**
- `stack-seeding/seeders/` package: 97.1% -> 97.4% stmts (set mode)
- `member_languages.go`: `Seed` 84.0% -> 100%; `languageRowsForMember` 96.7% (one residual defensive arm — see below)
- `certificates.go`: `memberRoleName` 66.7% -> 100%
- `users.go`: `Seed` 97.1% -> 98.1%; `backfillMembershipFields` 88.9% -> 100%
**Tests added:**
- M50 -> `seeders/m50_errorpath_harden_test.go`: 7 tests — 2 COPY-error (member_languages catalog + user_languages, with partial-total), 1 zero-population early-return, 1 native-English no-duplicate invariant, 1 no-taxonomy cert member-coverage (closure / empty-skills-envelope), 1 membership-fields backfill Exec-error, 1 empty-rows no-Exec guard.
**Bugs surfaced + fixed inline:** none — every error arm propagated wrapped (seeder name + surface) with correct partial-total accounting; every edge branch behaved to spec.
**Flakes stabilized:** none (no flakes; flake gate 3/3 clean on the new tests).
**Coverage notes (uncovered-by-design):**
- `member_languages.go:199` — the `add()` closure's `code == "" || seen[code]` dedup short-circuit is defensive depth-in-depth: every call site (`add(native)`, the guarded `add("en")`, the `if !seen[cand]`-gated third-language `add`) already pre-filters, so the guard is unreachable by construction. It protects the `(user, world_language)` unique-key invariant against a future caller that forgets to pre-filter; the whole-member no-duplicate invariant IS asserted (TestLanguageRowsForMember_NativeEnglish + TestMemberLanguagesSeeder_WritesCatalogAndPerMember). Kept the guard, not a contrived private-state poke.
- `users.go:300/311` — the casbin g2/g3 grant Exec-error arms are PRE-M50 (M42e/earlier) shared `Seed` code, outside the M50 cumulative footprint; not in scope for this M50-specific final harden.
**Cross-iter integration check:** the M50 footprint's cross-iter interactions are (a) the iter-02 member-field fill + iter-06 MemberLanguagesSeeder + cert-coverage all sharing the SAME deterministic per-member uuid space (`%s:user:%d`) UsersSeeder writes — to be pinned in Pass 2; (b) the iter-04/05 demopatch + content-URL rewrite are Python/shell (demo-stack/stack-snapshot), non-Go, exercised live by the gate (no standalone harness in the authoring copy per the static-only constraint).
**Stop condition:** continue-to-next-pass — final-mode cross-iter integration sweep (the shared per-member uuid-space lockstep across the M50 seeders) not yet pinned; delta-across-passes not yet computable.

## Pass 2 — 2026-06-30 — final

**Iters hardened this pass:** all milestone-touched code (final cumulative scope) — the cross-iter integration sweep, final-mode's defining work.
**Tiks covered since prior pass:** all iters in milestone.
**Coverage delta on touched files:**
- `stack-seeding/seeders/` package: 97.4% -> 97.4% stmts (stable — Pass 2 is integration tests pinning seeder INTERACTIONS, not new-statement coverage; delta across passes 1->2 = 0.0%, < 2%)
**Tests added:**
- M50 -> `seeders/m50_crossiter_harden_test.go`: 3 cross-iter integration tests — MemberIdentityLockstep (every languages/certs row references a UsersSeeder-created member uuid; no orphans; languages whole-roster; cert coverage a non-empty proper subset), HeroIsTheSamePersonEverywhere (the hero's membership-fields + languages + §B certs all on her ONE uuid), JointReseedIdempotent (re-running all three adds 0 rows — the reset-to-seed M17 contract composes).
**Bugs surfaced + fixed inline:** none — the shared per-member uuid-space lockstep holds across UsersSeeder + MemberLanguagesSeeder + CertificatesSeeder; the joint re-seed is idempotent.
**Flakes stabilized:** none (flake gate 3/3 clean across ALL new M50 harden tests — pass 1 + pass 2).
**Cross-iter integration findings:** the load-bearing invariant is `deterministicUUID("%s:user:%d", storyKeyPrefix(stack, story), i)` — every member-facing seeder derives a member's uuid identically, so the iter-02 member-field fill + the iter-06 languages + the iter-06 cert member-coverage all attach to the SAME person. A drift would silently break believability with every per-seeder test green. Now explicitly pinned. The iter-04/05 demopatch + content-URL rewrite (demo-stack/stack-snapshot, Python/shell) carry no standalone Go harness per the milestone's static-only constraint and were proved by the live M42 gate both vantages.
**Knowledge backfill:** `corpus/ops/demo/profile-completeness-spec.md` — added §"The shared per-member uuid-space (the cross-seeder believability spine)" documenting the shared-uuid-space invariant + the cross-iter sweep that pins it (the harden-discovered cross-cutting truth, previously implicit in the code).
**Stop condition:** stabilized — coverage delta across passes 1->2 = 0.0% (< 2%) AND the Phase 2 dimension scan (final-mode cross-iter integration) found nothing new (lockstep + joint idempotency hold, no bugs). All reachable M50-touched lines covered; the two residual uncovered blocks are documented uncovered-by-design (the defensive `add()` dedup short-circuit + the pre-M50 casbin arms). This is the M50 final-mode harden pass close-milestone's iterative gate requires.
