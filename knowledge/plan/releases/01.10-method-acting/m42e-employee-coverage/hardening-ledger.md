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
