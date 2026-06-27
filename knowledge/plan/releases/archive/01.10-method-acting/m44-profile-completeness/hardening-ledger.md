# M44 — Profile completeness · Hardening Ledger

The `/developer-kit:harden-milestone` pass over the M44 footprint (the believable-profile
DATA-DENSITY completeness milestone). Code-of-record lives in the `rosetta-extensions` authoring
copy (`stack-seeding/seeders/`); this corpus ledger is the audit trail. Tests added with **zero
source change** → rext tag bumped `…-fix1` → **`method-acting-m44-profile-completeness-fix2`**.

## Scope manifest (M44-touched, single Go stack — `stack-seeding`)

| Source file | Section | Existing tests | Harden coverage |
|---|---|---|---|
| `seeders/persona.go` + `persona_write.go` | §A trajectory-aware self-rating; §C manager verified-count/band | `trajectory_test.go`, `persona_*_test.go` | manager tail-distinctness (trajectoryVerifiedCount, not EffectiveVerified) |
| `blueprint/blueprint.go` `EffectiveSelfRated()` | §A self-rated split | `trajectory_test.go`, `blueprint`/`stories_test.go` | thriving self-rated vs struggling NULL already pinned |
| `seeders/certificates.go` (NEW) | §B1 CertificatesSeeder | `certificates_test.go` | role-coherence (all 3 families) + field-of-study + expiry mix + COPY-error + months-default + uint64-modulo regression |
| `seeders/projects.go` (NEW) | §B2 ProjectsSeeder | `projects_test.go` | role-coherence (all 3 families) + COPY-error + months-default |
| `seeders/profile.go` | §C manager profile; §D bulk-member depth | `profile_test.go`, `profile_harden_test.go` | manager cross-seeder skip consistency; managerTitle branches; member role↔membership coherence; member offset-0 tail; no-pool degradation |
| `seeders/users.go` | §D fix1 avatar column (`memberships.picture_url`) | `seeders_test.go` | heal-failure propagation (the §D fix1 backfill error path) |

No new-unit-without-handbook finding: `certificates`/`projects` are added seeders in the existing
`stack-seeding` tree (two new `reg.MustRegister` entries in `cmd/stackseed/main.go`), not new
top-level units; the surface family is already documented in `stack-seeding`'s README + the M44
spec docs. All 6 footprint sections already carried build-phase tests — hardening deepened the
thin dimensions (role-coherence branches, error paths, the cross-seeder + believability invariants).

## Pass 1 — 2026-06-26
**Coverage delta (seeders pkg statements):** 96.5% → 97.0% (+0.5).
**Tests added** (`profile_completeness_harden_test.go`, 10): §B `certBankForRole` /
`projectBankForRole` (all 3 role families + cross-family disjointness — a sales hero never draws
AWS), `certFieldOfStudy` (every keyword branch + the empty default), the **uint64-modulo
bank-index regression** (5000-key sweep proving `bank[hashInt(k) % uint64(len(bank))]` is never
out of range — the negative-int-index class the build commit's "modulo-in-uint64" guard closes),
cert date invariants (issued-in-past, expiry strictly after issue); §C the manager
cross-seeder-skip + populated-profile DAG invariant, `managerTitle` field branches (incl. the
default `<Role> Manager`), the manager **claimed-tail distinctness** (tail skips
`trajectoryVerifiedCount`, not the bare `EffectiveVerified()=0`); §D member role↔membership
draw coherence, member offset-0 whole-tail-claimed.
Functions lifted to 100%: `certBankForRole`, `projectBankForRole`, `certFieldOfStudy`,
`managerTitle`.

## Pass 2 — 2026-06-26
**Coverage delta:** 97.0% → 97.3% (+0.3).
**Tests added** (5, error paths + edge): cert/project COPY-failure propagation (wrapped, names
the seeder + the failing table, 0 rows — via the reused `failCopyConn`/`assertWrapped`); the **§D
fix1 `memberships.picture_url` heal-failure** propagation (both the `UsersSeeder` integration path
and the `backfillMembershipPictures` unit, naming the offending membership id) — so an avatar-heal
failure on a re-seed is loud, never silent; the `months<=0 → 6` span default (reached via a
declared `PassRate` with `Months==0`, the one config the blueprint resolver leaves unfilled).
Functions lifted to 100%: `CertificatesSeeder.Seed`, `ProjectsSeeder.Seed`, the span-default arms
in `certRowsForHero`/`projectRowsForHero`.

## Pass 3 — 2026-06-26 — final
**Coverage delta:** 97.3% → 97.5% (+0.2 — below the 2% stabilization threshold).
**Tests added** (2): the §B **cert-expiry believability mix** (a 10-hero roster's cert set must
carry BOTH expiring [`to` set] and perpetual [`to` NULL] credentials — a section where every cert
expires, or none, reads synthetic), and the §D **no-pool tail degradation** (an absent taxonomy →
a member's claimed tail is empty, no fabricated node-ids; the timeline still seeds — the
no-fabrication closure rule at the member level).
All **M44-introduced** seeder functions now at **100%**. The sole remaining sub-100% in a touched
file is `resolveHeroSkills` (95.7%, `persona.go:176` — the curated-pool blank/dup `continue`) and
two `users.go` casbin grant-error arms (`256`/`267`) — **all pre-existing M42e code**, not the M44
introduction; out of strict footprint scope.

### Stop condition
**Stabilized.** Coverage delta this pass +0.2% (< 2%); the full six-dimension Step-2b scan found no
remaining M44-footprint gap worth a non-shallow test (the residual uncovered blocks are pre-existing
defensive branches outside the M44 introduction); zero flakes across 3 sequential shuffled runs.

## Bugs fixed inline
**None.** No production bug surfaced — the build code (incl. the `fix1` avatar-column correction)
was correct; hardening pinned its invariants. The uint64-modulo bank-index pattern was already the
safe form (modulo taken in `uint64` before the `int` cast); the regression test now guards against
a future refactor reintroducing the negative-int-index class. One test self-correction during Pass 2
(the months-default test's precondition: the blueprint resolver only defaults `Activity.Months` when
BOTH `Months` and `PassRate` are 0 — fixed the fixture to declare a `PassRate` so `Months==0` reaches
the seeder's own guard).

## Flakes stabilized
**None** — zero flakes. New tests: 3 sequential shuffled `-count=1` runs clean. Full `seeders` suite:
3 sequential shuffled runs clean + `go test -race` clean.

## Knowledge backfill
**No KB-worthy new system behavior.** The hardening pinned invariants already documented in
`decisions.md` (D7 — `/enterprise/members` reads `memberships.picture_url`, not `users.picture`; the
COPY-fill + idempotent `backfillMembershipPictures` heal), `spec-notes.md` / the seeders' doc comments
(the LIVE-SCHEMA corrections — `user_certifications.certification`/no-created_at/no-organization_id;
`user_projects.title`/`to`/no-organization_id), and `corpus/ops/demo/profile-completeness-spec.md` +
`corpus/ops/seeding-spec.md` (the role-coherent cert/project banks, the trajectory-aware self-rating,
the manager unskip, the bulk-member depth). No new behavior was discovered; nothing to blend.

## Audits
- **Closure GREEN** — no fabrication; every test asserts real-taxonomy node-ids or empty (the no-pool
  degradation test pins the closure rule at the member level: absent taxonomy → empty tail, no
  fabricated ids).
- **Isolation PerStackIsolated** — every M44 write is per-stack Postgres; zero shared/prod writes
  (audit-recorded; the seeders' `Isolation()` contract unchanged).
- **Supply-chain GREEN** — zero `go.mod`/`go.sum`/`package.json`/lockfile change across the whole M44
  footprint (byte-identical to the build tag); no new deps.
- **Alignment N/A** — zero `clerkenstein/` change; all 5 Clerkenstein gates carry forward 100%/100%.
- **The avatar both-columns regression** — `memberships.picture_url == users.picture` pinned for every
  member (the cross-column match in `seeders_test.go`) + the backfill heal-once-then-zero idempotency
  + the documented fact that `/enterprise/members` reads the membership column (so future avatar work
  targets the right surface) + (new this pass) the heal-failure-propagates error path. The gotcha that
  bit the build is now bulletproof.

## Commits (rext authoring copy; corpus holds this ledger)
- `7425c4e` harden(M44): §B role-coherence/field-of-study + uint64-modulo bank-index regression; §C
  manager cross-seeder skip + managerTitle + tail distinctness; §D member role↔membership + offset-0 tail
- `01bb217` harden(M44): error-path + span-default — cert/project COPY-failure wraps,
  memberships.picture_url heal-failure propagation (§D fix1), months<=0 default
- `7e6629a` harden(M44): §B cert-expiry believability mix + §D no-pool tail degradation
- rext code-of-record @ tag **`method-acting-m44-profile-completeness-fix2`** (tests-only bump from `…-fix1`).
- Corpus: this ledger + the `## M44: Hardening` section in `progress.md` (`m44/profile-completeness` branch).

**Next:** `/developer-kit:close-milestone` (its Phase-4 audit runs independently as defense-in-depth).
