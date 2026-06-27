# iter-09 progress

**Type:** tik (iter_shape: production-fix) -- under TOK-01, `sweep-then-route-by-leverage`. Per `coverage-protocol.md` "seedable structural row vs runtime-computed artifact" + the gate-honesty audit (`verify-gate.md`).

Corrects iter-08 D2: the `/sim/<slug>/start` skip was DISHONEST (the page is empty because Maya's org lacks the `FEATURE_JOB_SIMULATIONS` Sentinel `g3` entitlement -- a SEEDABLE deny-modal, not a runtime-computed surface).

## Phase B/C — the fix (in-rext, zero platform edit)
- **D1 — seed the Sentinel g3 feature grant.** Diagnosed the entitlement gate to its EXACT surface (read-only):
  the `/sim/<slug>/start` deny modal is gated by `userMembership.organizationFeatures`, which `app` resolves
  via **Sentinel's Casbin `g3` policy** (`g3 = _, _` → `casbin_rules` `p_type='g3', v0=org, v1=membership`), NOT
  the `app.organization_features` table (0-rows even normally — the symptom the verdict named). Live demo-3 had
  **0 g3 rows**. Added `seedCasbinFeatureGrants` to `identity.go` (reusing `resolveCasbinTable` + the idempotent
  `casbinGrantSQL`), wired a g3 grant per membership in `users.go` + for the demo identity in `identity.go`.
- **D2 — removed `SIM_START_SKIP`** from `coverage.spec.ts`: `/sim/<slug>/start` is back IN SCOPE + scored.
- Seeder tests updated (TestUsersSeeder 16→20 rows + a g3 assertion; TestIdentitySeeder 3→4 rows + a g3
  assertion); `go test ./...` GREEN across stack-seeding; the harness compiles (`playwright test --list`).

## Phase C — re-apply to the live demo
- Re-seeded demo-3 (`stackseed --stack demo-3 --seed stories.seed.yaml`): **341 g3 grants** landed (220 Cervato
  + 1 demo identity + 120 Solvantis), Maya Chen's membership confirmed; isolation **clean (no shared/external
  writes), prod=false**.
- **D3 — Sentinel policy reload.** The grant first stayed invisible (`organizationFeatures=null`, deny modal
  still showing) — Sentinel's enforcer `LoadPolicy()`s ONCE at startup with NO watcher, so a raw-INSERT into a
  LIVE stack isn't seen until reload. Restarted `demo-3-sentinel-1` → `organizationFeatures=["FEATURE_JOB_SIMULATIONS"]`,
  deny modal GONE, the **real launch UI renders** (`bodyTextLen=625`: "Welcome to your AI Simulation… Time
  available: 20 minutes…"). Probe-confirmed.

## Phase D — re-sweep (HONEST, /start in scope, frontier-EXHAUSTED)
- First re-sweep collapsed to **8 pages** (the sentinel cold-restart cleared the GraphQL caches; the 1.5s
  settle rendered only 1 of 22 skill-path cards → the BFS frontier collapsed). Diagnosed + fixed:
  - **D4 — `<main>`→`<body>` fallback** when `<main>` is below the floor (the `/start` launch UI renders
    OUTSIDE `<main>`: `mainLen=0`, content in `<body>`).
  - **D5 — settle ceiling 1.5s→4s** (a probe proved 4s renders the full 22-card grid; correct-over-fast).
- Authoritative re-sweep: **`reachable=95` (frontier EXHAUSTED, cappedAtFrontier=FALSE)** — DEEPER than
  iter-08's 93 (the `/start` page now in scope + scored).
- **HONEST residual = `(failing=0, escapes=1)`:**
  - **failing=0** — every reachable page renders, INCLUDING `/sim/get-ready-…/start` (`text=625`, the launch UI;
    the entitlement fix WORKS). The dishonest iter-08 skip is GONE.
  - **escapes=1** — a NEW honest escape the old truncated frontier hid: `/reimport-profile` →
    `linkedin.com/help/...` (a baked LinkedIn-help `<a href>` in platform onboarding UI). The 3 editorial
    citations remain correctly classified as presenter notes (not escapes).

## Phase E — Close

**Outcome:** The dishonest sim-start exclusion is FIXED HONESTLY — the Sentinel g3 entitlement is seeded, the
skip is removed, and `/start` now RENDERS + is SCORED (failing: 1→0 on the start cluster). But the deeper,
truly-exhausted frontier (95 pages) surfaced a previously-hidden **escapes=1** (`/reimport-profile` →
linkedin.com/help), so the HONEST residual is **`(failing=0, escapes=1)`** — the gate is **NOT (0,0)**. That
escape is a **hardcoded** platform-UI `<a href>` (not env-rewritable) → the **zero-edit-line Re-scope trigger**:
escalated to the user (3 options recorded; recommend the in-rext allow-rule extension mirroring the editorial-
citation disclose-not-hide design). iter-08's gate-MET claim is now corrected to the honest measurement.
**Type:** tik (iter_shape: production-fix)
**Status:** closed-fixed
**Gate:** NOT MET — honest residual `(failing=0, escapes=1)` with `/start` IN SCOPE + scored (was the dishonest `(0,0)`)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: **y** — (4) user-blocker: n — (5) cap-reached: n (tik #1 of session) — (6) protocol-stop: n — Outcome: exit-3 (re-scope-trigger)
**Decisions:** D1 (seed Sentinel g3 entitlement → /start renders), D2 (remove dishonest SIM_START_SKIP → /start scored), D3 (live-stack casbin seed needs a Sentinel reload — restart demo-3-sentinel), D4 (harness <main>→<body> fallback for out-of-<main> content), D5 (settle ceiling 1.5s→4s — under-settle collapses the BFS frontier) — see ./decisions.md; the re-scope-trigger record is in the milestone-root decisions.md
**Side-deliverables (if any):** D4 + D5 are reusable harness-fidelity fixes (folded into coverage-protocol.md): the out-of-<main> content fallback + the data-grid settle floor. They fixed a real measurement bug (the 8-page collapse) surfaced this iter.
**Routes carried forward:** the `/reimport-profile → linkedin.com/help` escape → the user's re-scope decision (allow-rule extension | platform PR | carve-out). Recorded as the milestone-root RE-SCOPE-TRIGGER entry.
**Lessons:** (1) an entitlement-gated empty page is SEEDABLE (a Sentinel g3 policy), not runtime-computed — the
correct fix is `stack-seeding`, not a crawl-scope skip; re-instating a skip on a seedable failure is the
gate-honesty failure mode. (2) A casbin seed on a LIVE stack needs a Sentinel policy reload (restart) — the
enforcer loads policy once at startup, no watcher. (3) Some pages render content OUTSIDE `<main>` — fall back to
`<body>` when `<main>` is below the floor. (4) The per-page settle must cover the heaviest DATA GRID, not the
first paint — under-settle collapses the BFS frontier (1.5s→4s). (5) Honestly putting a dishonestly-skipped
page back in scope + truly exhausting the frontier can SURFACE a previously-hidden escape — the deeper-crawl
discovery is the gate working, not regressing. All folded into coverage-protocol.md.
