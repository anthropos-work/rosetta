# M39 Progress

Section checklist (built by `/developer-kit:build-milestone`). Scope detail in `overview.md`.

## Section checklist
- [x] **G1 — org name** — thread `st.Org.Name`/slug through the roster (`stack-seeding/seeders/roster.go`: `RosterIdentity` + `BuildRoster`) → clerkenstein `registry.go` (`RosterEntry`) + `resources.go` (`DemoUser`, `orgMemberships()` l.227); FAPI org resource carries the story org name; `"Clerkenstein Demo Org"` stays the no-roster default; paired `DisallowUnknownFields` struct change + re-tag BOTH repos. **DONE** (commit fb9e300): + single-sourced `orgSlugFor`; both modules green; multi-identity + JS/FAPI alignment 9/9 / 100%.
- [x] **G2 — role backfill** — backfill `public.user_basic_info.job_role_id` (+ `job_title`/`summary`/`location`) from the resolved hero role in `stack-seeding/seeders/users.go`; one UPDATE lights the `/profile` header (`infoResolver.JobRole`) + role-gap radar / role-readiness widgets (`jobRoleMatch`). **DONE** (commit 010f422): idempotent UPDATE (IS-DISTINCT-FROM guard, validated live on demo-3); every member backfilled; no-fabrication preserved; seeders suite green.
- [x] **G4 — real-face avatars** — replace the DiceBear initials SVG (`users.picture`) with a bundled, license-clean real-face set mapped deterministically by hero key (offline-safe + deterministic + license-clean). **DONE** (commit fc8a841): a self-authored parametric SVG **face generator** → offline base64 data URI (no fetch, ~1 KB, license-clean — no vendored asset at all); visually verified (clean varied faces); stdlib-only.

## M39: Hardening

`/developer-kit:harden-milestone` pass (2026-06-24). 3 passes, stopped on stop-condition (full
six-dimension scan clean + coverage delta < 2% + zero flakes), well under the 5-pass cap. All work in the
**rext authoring copy** (the separate `rosetta-extensions` repo — M39's code lives there); the
`m39/profile-identity` rosetta branch carries only this progress note. Tag `method-acting-m39` moved to
the new rext HEAD (test-only delta; go.mod/go.sum byte-identical, supply-chain GREEN; all 3 offline
Clerkenstein alignment gates held 100%/100% — Go 22/22, JS 9/9, multi 9/9).

### Scope manifest (Phase 1 — milestone-touched code)
Two rext modules, 9 source files. Non-test source touched + its test file:

| Module | Source file | Test file | Baseline cov |
|--------|-------------|-----------|--------------|
| stack-seeding/seeders | `avatar.go` (new, G4) | `avatar_test.go` | 100% (all 3 fns) |
| stack-seeding/seeders | `users.go` (G2 backfill) | `seeders_test.go` + `multiorg_test.go` | `backfillUserBasicInfo` 77.8%, `Seed` 97.4% |
| stack-seeding/seeders | `userprofile.go` (G2 banks) | `userprofile_test.go` | 100% |
| stack-seeding/seeders | `roster.go` (G1) | `roster_test.go` | 100% |
| stack-seeding/seeders | `helpers.go` (G1 `orgSlugFor`) | `seeders_test.go` | 100% |
| stack-seeding/seeders | `org.go` (G1 `orgSlugFor` call) | `seeders_test.go` + `multiorg_test.go` | `Seed` 88.9% (pre-M39 error branch) |
| clerkenstein/clerk-frontend | `resources.go` (G1 `orgMemberships`) | `resources_test.go` | `orgMemberships` 90.0% |
| clerkenstein/clerk-frontend | `registry.go` (G1 `RosterEntry`) | `registry_test.go` | 100% |

Package baselines: seeders **95.9%**, clerk-frontend **86.3%**. No new-unit-without-handbook findings (no
new tool/package — `avatar.go`/`userprofile.go` are new files within the existing `seeders` package, which
is already documented in `seeding-spec.md`). No `fix`/`bug` commits in the M39 range → no build-phase-bug
regression-test gap (G1/G2/G4 are clean feature commits).

### Pass 1 — 2026-06-24 (rext `a990a7c`)
Closed the two real coverage candidates on M39-touched code + added G4 markup-robustness.
**Coverage delta (milestone-touched functions):**
- `backfillUserBasicInfo` (G2): 77.8% → **100%** (statements)
- `orgMemberships` (G1): 90.0% → **100%** (statements)
- Package: seeders 95.9% → 96.1%, clerk-frontend 86.3% → 86.7%

**Tests added:**
- `avatar_test.go`: 3 (`WellFormedXMLAcrossSeeds` — parse 256 generated SVGs through `encoding/xml`;
  `AllHairStylesWellFormed` — all 4 `style%4` arms + mod-wrap; `EmptySeedStillAFace` — degenerate edge)
- `multiorg_test.go`: 3 (`BackfillUserBasicInfo_EmptyInputIsNoOp` edge; `_ExecErrorPropagates` error path;
  `BasicInfoBackfillErrorBubblesThroughSeed` error-propagation-through-`Seed`) + `storyConn.basicInfoExecErr` injector
- `resources_test.go`: 2 (`orgMembershipsNoOrgClaim` empty-OrgAuthID guard; `orgMembershipsMixedFallbacks`
  — name-set/slug-empty + slug-set/name-empty isolated independently)

**Bugs fixed inline:** none — the build phase's implementation was correct; these closed test gaps only.
**Flakes stabilized:** none observed.
**Knowledge backfill:** none warranted (the behaviors pinned — offline-safe well-formed SVG, the no-role
NULL backfill, the no-org / org-name fallbacks — are already documented in `seeding-spec.md`,
`stories-spec.md`, and `clerkenstein.md`; these tests pin existing documented contracts, surfacing no new
truth about the system).

### Pass 2 — 2026-06-24 (rext `d6f3880`)
Integration-level **cross-write agreement** invariants the build phase pinned only one-side-at-a-time.
**Coverage delta:** flat (seeders 96.1%; these run both writers and assert agreement — not new lines).

**Tests added (`multiorg_test.go`, 3 integration):**
- `TestG2_BasicInfoRoleAgreesWithMembership` (M39-D6): per user, `user_basic_info` header role
  (`job_role_id`/`job_title`) == that user's `memberships` role (`job_role_id`/`job_role_name`) — the
  header and member row read the role from two tables; a drift would advertise different roles. Holds for
  the no-taxonomy NULL case (`nullIfEmpty` maps `""`→nil on both writes).
- `TestG1_SeededOrgSlugAndNameMatchRoster` (M39-D2): the org name+slug `OrgSeeder` writes to
  `public.organizations` == the name+slug `BuildRoster` threads into the FAPI org resource, per story (top
  bar vs DB can't disagree). Covers explicit-slug AND slug-from-name stories (`orgSlugFor` single-source proof).
- `TestG4_AvatarStableAcrossReSeed`: re-seeding the same blueprint yields byte-identical `picture` for all
  70 users (column-level avatar idempotency — the ON-CONFLICT no-op stays a true no-op).

**Bugs fixed inline:** none. **Flakes:** none. **Knowledge backfill:** none — D2/D6 already record these
invariants in `decisions.md`; the tests enforce them, surfacing nothing new.

### Pass 3 — 2026-06-24 (rext `c360b4e`)
Fuzzed the one M39 function that turns free-form input into browser-rendered markup.
**Coverage delta:** flat (96.1%; fuzz is robustness, not new lines).

**Tests added (`avatar_test.go`, 1 fuzz):**
- `FuzzAvatarDataURI`: for ANY seed the output must be an offline data URI (no http/https), well-formed XML
  (token-walked to EOF), still a face (no `<text>`), no non-finite coords from the float-formatted paths,
  and deterministic. Seed corpus: empty, single-char, real uuid, XML-hostile chars, unicode, NUL bytes, 4 KB.
- **Verified:** 714K execs / 10 workers / 12s, **0 crashes** — the generator is TOTAL over arbitrary input
  (the seed drives only the deterministic feature picks, never the markup envelope). Active-fuzz interesting
  inputs land in `GOCACHE`, not `testdata/` — no corpus artifact committed.

**Bugs fixed inline:** none. **Flakes:** none. **Knowledge backfill:** none warranted.

### Stop condition
Stopped after Pass 3 (cap is 5): the full Step-2b six-dimension scan found nothing new worth adding for
M39's narrow touched surface (all touched functions at 100%; error/edge paths closed; cross-write
invariants pinned; the one free-form-input function fuzzed clean), coverage delta is < 2% (and flat across
Passes 2–3 by design — integration + fuzz, not line-fillers), and zero flakes across the verification runs.
The residual package-level uncovered lines (`org.go:Seed` 88.9%, `users.go:Seed` 97.4%) are **pre-M39**
error-return branches in surrounding code — out of this milestone's harden scope.

**Net:** 12 tests added (8 unit/edge/error + 3 integration + 1 fuzz), 0 bugs (build phase was correct), 0
flakes. M39-touched functions all at 100%. Re-tag `method-acting-m39` → new rext HEAD.

## M39: Final Review

`/developer-kit:close-milestone` review (2026-06-24). A near-clean close: the build phase's implementation was
correct and the harden pass had already deepened tests + verified all 3 alignment gates, so the review found
**0 scope gaps, 0 code-quality must-fixes, 0 doc inaccuracies, 0 test/flake gaps**. The only fix-queue items
were decision-triage reference tags. Deferral re-audit **GREEN** (0 deferrals; the two `Out:` items are
Fate-2, already owned by M40/M41 — confirmed, no edit). Adversarial review (Phase 2c) weighed 2 scenarios
(AR-1 G2 silent-0-rows, AR-2 G1 slug-collision); both no-fix-needed, recorded in `decisions.md`. All 3 offline
Clerkenstein alignment gates re-verified at close HEAD `c360b4e`: **Go 22/22, JS 9/9, multi 9/9 — 100%/100%**.
go.mod/go.sum byte-identical (supply-chain GREEN). Flake gate: 5× shuffle, 0 flakes, both modules.

### Scope
- [x] All 3 sections (G1/G2/G4) delivered as Fate-1 in this milestone; `Out:` items confirmed Fate-2 (M40/M41) — no gap.

### Code Quality
- [x] [verified] Cross-cutting review clean: consistent patterns (single-sourced `orgSlugFor`, deterministic-by-seed
  helpers, `nullIfEmpty` no-fabrication, audit-on-write), no dead code (`avatarURL` removed in D8), boundary-safe
  (empty-OrgName fallback, empty-jobTitle generic blurb), no resource leaks, `go vet` clean both modules.

### Documentation
- [x] [verified] `stories-spec.md` + `clerkenstein.md` M39 sections accurate against the implemented code; all
  intra-/cross-doc anchors resolve; no new top-level unit (avatar.go/userprofile.go are new files in the existing,
  already-documented `seeders` package); no handbook/count drift.

### Tests & Benchmarks
- [x] [verified] Full suites green (seeders 209 funcs/214 PASS, clerk-frontend 51/86); flake gate 5× shuffle 0 flakes;
  no regression-test gap (0 build-phase bugs).

### Decision Triage
- [x] D2 → blend confirmed + tag added (#M39-D2) in clerkenstein.md (orgSlugFor single-source)
- [x] D3 → blend confirmed + tag added (#M39-D3) in clerkenstein.md (no-roster fallback)
- [x] D4/D5 → blend confirmed + tag added (#M39-D4/D5) in stories-spec.md (idempotent UPDATE + IS-DISTINCT-FROM)
- [x] D7 → blend confirmed + tag added (#M39-D7) in stories-spec.md (offline SVG face generator)
- [x] D1, D6, D8 → archive (maintainer-only: paired-change mechanics, backfill-every-member, dead-code removal)
- [x] AR-1, AR-2 → archive (adversarial scenarios; no-fix-needed, recorded in decisions.md)

## M39: Completeness Ledger (Phase 9 — section variant)

Every `overview.md` scope item placed in exactly one three-fate category. Cross-checked against
`overview.md` (In/Out lists), `progress.md` (all checkboxes checked), `spec-notes.md` (no unmet promise),
`decisions.md` (no defer language), and the Phase-1 code scan (0 TODO/FIXME/HACK in M39 source).

### Done (Fate 1 — delivered in this milestone)
- **G1 — org name.** `RosterIdentity`/`BuildRoster` carry `org_name`/`org_slug` → `RosterEntry` → `DemoUser` →
  `orgMemberships()`; the FAPI org resource reads the story org name; no-roster `"Clerkenstein Demo Org"`
  default preserved; paired `DisallowUnknownFields` change + single tag. ✅ (commit fb9e300)
- **G2 — role backfill.** `public.user_basic_info` backfilled (job_role_id + job_title + summary + location)
  via the idempotent IS-DISTINCT-FROM UPDATE; lights the /profile header + role-gap widgets; no-fabrication. ✅
  (commit 010f422)
- **G4 — real-face avatars.** Offline self-authored parametric SVG face generator → base64 data URI,
  deterministic by uuid, license-clean. ✅ (commit fc8a841)

### Confirmed-covered (Fate 2 — already owned by another milestone of this release; no edit)
- **Work / education history + skill depth** → **M41 (profile depth)** — its `overview.md` `In:` list owns it
  (G3 `user_experiences`/`user_educations` + G5 skill-depth bump). Confirmed.
- **The surfacing / serve-grant fix (library + activity feed)** → **M40 (directus serve-grant)** — its
  `overview.md` `In:` list owns it. Confirmed.

### Annotated (Fate 3 — attached to a release-milestone at close)
- None.

### Dropped
- None.

### Release-scope-breaking deferral (escape hatch — requires user sign-off)
- None.

**Verdict:** all 3 In-list scope items delivered as Fate 1; the 2 Out-list items are Fate-2 (already owned by
M40/M41, confirmed). Nothing annotated, dropped, or escape-hatch-deferred. **Zero items require sign-off.**
