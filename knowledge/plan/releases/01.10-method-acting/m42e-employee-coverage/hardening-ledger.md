# Hardening Ledger — M42e Employee believability coverage

## Pass 1 — 2026-06-25 — final

**Iters hardened this pass:** all milestone-touched code (cumulative final-mode scope across iters 01–23)
**Tiks covered since prior pass:** all iters in milestone (first harden pass — no prior ledger)
**Coverage delta on touched files:**
- stack-seeding/assets (avatars.go): 77.8% -> 83.3% stmts (AvatarJPEG 80% -> 100%)
- stack-seeding/seeders: 95.7% -> 96.4% stmts (curatedNamedFor 66.7%->100%, curatedSkillsFor 80%->100%, experienceTitle 66.7%->100%, isDevOpsRoleTitle 80%->100%, seedCasbinFeatureGrants 76.9%->100%, resolveCasbinTable 85.7%->100%)
- clerkenstein/clerk-frontend: 86.7% stmts held; the M42e image-threading invariant (image_url/has_image empty-vs-set JSON) newly PINNED by a behavioural test (was statement-covered transitively, never asserted)
**Tests added:**
- curated_pools_harden_test.go: 5 (zero-value guards, curatedSkillsFor all-categories+default, precedence-collision, case/whitespace, empty-names short-circuit)
- profile_harden_test.go: 2 (experienceTitle DevOps-arc, isDevOpsRoleTitle classifier)
- assets/avatars_test.go: 1 (AvatarJPEG negative-index normalization)
- identity_casbin_harden_test.go: 4 (seedCasbinFeatureGrants empty/g3-shape/idempotency/2 error paths)
- clerk-frontend/resources_test.go: 2 (avatar + org-logo image-threading empty-vs-set JSON serialization)
**Bugs surfaced + fixed inline:** none (all gaps were absence-of-coverage on edge/error branches; no behavioural defects)
**Flakes stabilized:** none observed (flake gate deferred to session-exit verification)
**Knowledge backfill:** none required (no new edge-case/precision semantics discovered beyond what spec-notes + iter progress.md already document)
**Alignment gates:** all 5 GREEN (Go 22/22, JS 9/9, multi 9/9, deploy 7/7, express 13/13) + drift-test 9/9. Supply-chain: 0 go.mod/go.sum, Playwright ^1.49.0 held.
**Stop condition:** continue-to-next-pass — org.go/hero_activity.go/identity.go Seed + simembeddings error paths not yet swept; need a 2nd pass to measure coverage delta

## Pass 2 — 2026-06-25 — final

**Iters hardened this pass:** all milestone-touched code (cumulative — the seeder write/error branches the happy paths skipped)
**Tiks covered since prior pass:** n/a (same final-mode session, 2nd pass)
**Coverage delta on touched files:**
- stack-seeding/seeders: 96.4% -> 97.0% stmts (pass-over-pass +0.3%; cumulative across both passes 95.7% -> 97.0%)
  - org.go Seed 88.9% -> 100%; hero_activity.go Seed 93.8% -> 98.4%; identity.go Seed 89.7% -> 96.6%; itoaInt 73.3% -> 100%; seedCasbinGrant 100%
**Tests added:**
- org_harden_test.go: 2 (logo_url per-org monogram thread + organizations COPY error path)
- hero_activity_test.go: 4 (3 per-table COPY error paths w/ partial-total assertions + months<=0 default-span)
- seeders_test.go: 2 (IdentitySeeder users + memberships COPY error paths)
- profile_harden_test.go: 1 (itoaInt incl. negative branch)
**Bugs surfaced + fixed inline:** none (all gaps were absence-of-coverage on write/error branches; no behavioural defects)
**Flakes stabilized:** none — flake gate 3/3 consecutive clean (shuffled, -race) across seeders + assets + clerk-frontend new tests
**Knowledge backfill:** none required (no new semantics surfaced; the error-wrap contract + monogram determinism are already documented in the seeder source + iter progress.md)
**Alignment gates:** all 5 GREEN (Go 22/22, JS 9/9, multi 9/9, deploy 7/7, express 13/13) + drift-test 9/9 (re-verified). Supply-chain: 0 go.mod/go.sum across the milestone, Playwright ^1.49.0 held. -race GREEN: stack-seeding 9 ok, stack-snapshot 13 ok, clerkenstein 14 ok. TS harness transpiles + playwright --list 7 tests/6 files. Shell scripts bash -n + shellcheck clean. datadna closure: no fabricated node-ids (replay-resolved real public node-ids only — proven per-iter on demo-3 + the curated drop-not-fabricate logic now branch-covered).
**Residual (accepted, not chased):** mustAvatarNames 75% + photoAvatarDataURI 75% (the embed-broken defensive panic/fallback — unreachable without breaking the compile-time go:embed); seedCasbinGrants 92.3% (a g2 bulk-loop error path structurally identical to the now-100% seedCasbinFeatureGrants); identity.go Seed 96.6%. These are intentionally-defensive / structurally-equivalent-to-covered branches; fault-injecting the embed is not cheaply testable and would add no real assurance.
**Stop condition:** stabilized — pass-over-pass coverage delta +0.3% (< 2%) AND the Phase 2 dimension scan found no new behavioural bugs (only intentionally-unreachable defensive residuals). Final-mode harden complete; ready for /developer-kit:close-milestone.
