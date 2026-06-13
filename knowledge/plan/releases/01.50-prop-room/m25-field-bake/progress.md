# M25 — Progress

**Status:** building — **3 done-bars resolved (DB-3, DB-5 GREEN; DB-1/DB-2 partial), 1 hard blocker
(DB-3-cache → DB-1/DB-2/DB-4)**. **Shape:** section.

## Section checklist
_The 5 live done-bars from `overview.md` § Scope — each a real stack run on this box, binary pass/fail._

- [~] **DB-1 — Fresh `/demo-up`**: stack UP, taxonomy real (329,859 rows), data plane wired LOCAL
  (`DIRECTUS_BASE_ADDR=http://directus:8055`) + asset plane PROD (`content.anthropos.work`), local Directus
  **boots + healthy on offset 18055**. **BLOCKED on content-serve** (0 collections) — rows-only cache predates
  M21 structure; needs the structure-bearing re-capture (M25-D3). The new Directus probes fire correctly; the
  serve probe is RED *because* of the cache gap, not a probe bug.
- [~] **DB-2 — `/dev-up 2 --local-content`**: content-serve half blocked by the same cache gap + VM budget.
  **N=0-stays-prod-read half VERIFIED** in code (M25-D4): dev defaults `LOCAL_CONTENT=0` + the hard N=0
  set-dress guard.
- [x] **DB-3 — Re-run idempotency**: **GREEN** — full re-run convergent; migrate no-op ×5, casbin
  "already 248 — skipping", directus bootstrap "already bootstrapped — skipping", taxonomy TRUNCATE-reload →
  329,859 (same), seed idempotent. The directus cache-miss reproduces deterministically (correct idempotency).
- [ ] **DB-4 — Cold-start capture**: **BLOCKED** (M25-D3) — needs an operator-sanctioned prod read or a
  staging dump; neither available autonomously, policy mandates operator confirmation.
- [x] **DB-5 — Clean teardown**: **GREEN** — `/demo-down 1 --purge` reclaims the directus container + frees
  offset port 18055 + releases the registry slot (no stacks registered; 0 demo-1 containers remain).

## Build log

### Session 2026-06-13 (field-bake)
- Phase 0b KB-fidelity: **YELLOW** — 3 parallel audits, all v1.5 behaviors ALIGNED; 2 arg-hint drifts fixed
  inline (`demo-up` fictional `--full`/env-var flags; `dev-up` missing `--local-content`). Report:
  `kb-fidelity-audit.md`.
- Box reconcile: torn down a stale `demo-1` + a crash-looping N=0 sentinel orphan; re-pinned the per-stack
  `stack-demo` ext clone from the v1.3b tag → `prop-room-m25` (the release-under-test tooling).
- **Bug fixed inline (M25-D1):** `/demo-up` build gate aborted deterministically — the offline clerkenstein
  cross-compile (`GOPROXY=off GOSUMDB=off`) couldn't fetch/verify M24's pinned `toolchain go1.25.11`. Added
  `GOTOOLCHAIN=local` (2 sites). Committed authoring `35180c0`; tagged `prop-room-m25`.
- **Resource findings (M25-D2):** the documented 12 GB Docker-VM prereq **fails to boot** on this 16 GB host
  (backed off to 10 GB); the VM **disk filled** (M3 precedent) — pruned 45 GB of build-cache/images to recover.
- **BLOCKER (M25-D3):** the local-Directus **content serve** + the cold-start capture (DB-1/DB-2/DB-4) are
  gated on a **structure-bearing directus capture** that needs an operator-confirmed prod read / staging dump
  — the cache on this box is rows-only (pre-M21). Surfaced to the user. See `decisions.md` M25-D3.

## Open questions (blocker — needs user input)
- **The directus snapshot cache is rows-only (pre-M21 structure).** To finish DB-1/DB-2/DB-4 the operator must
  run one sanctioned structure-bearing capture (`stacksnap capture --surface directus --dsn <restored-staging-
  dump | prod-read-DSN>`) — a privileged prod read this agent cannot self-authorize. Once the cache is
  structure-bearing, re-running `/demo-up` should make the local Directus serve (no code change expected).
