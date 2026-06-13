# M25 — Progress

**Status:** **ALL 5 DONE-BARS GREEN** (resume attempt 2). The operator sanctioned the prod-read DSN; the
first real capture surfaced + **fixed a real safety bug** (M25-D5: the `directus_files` closure over-captured
158 tenant-referenced files; firewall caught it fail-closed) plus two dangling-FK bugs (M25-D6 group / M25-D7
folder+uploaded_by+modified_by). With the fixes the structure-bearing capture **passes the firewall**, the
local Directus **serves** (DB-1 + DB-2 proven by curl on the offset ports), the cold-start capture was
exercised (DB-4), and re-run idempotency + teardown stay GREEN (DB-3, DB-5). **Shape:** section.

## Section checklist
_The 5 live done-bars from `overview.md` § Scope — each a real stack run on this box, binary pass/fail._

- [x] **DB-1 — Fresh `/demo-up`**: **GREEN** — stack UP (`--no-ui`, ~10 GiB VM), taxonomy real (329,859),
  data plane LOCAL (`DIRECTUS_BASE_ADDR=http://directus:8055`) + asset plane PROD (`content.anthropos.work`).
  With the structure-bearing, firewall-clean re-capture (M25-D5) and the three dangling-FK nulls (M25-D6
  group / M25-D7 folder+uploaded_by+modified_by), the directus replay now **exits 0**: structure
  auto-provisions, digest converges, rows load, the per-stack Directus **boots + SERVES** on offset 18055.
  Evidence: `curl` the local Directus REST API → 200 + a real published row; image URL resolves to
  `content.anthropos.work` (asset plane prod). (3 real field bugs found + fixed inline en route.)
- [x] **DB-2 — `/dev-up 2 --local-content`**: **GREEN** — dev-2 (offset 20000, per-stack Directus on 28055)
  brought up with `--local-content`; the **same `dev-setdress.sh` engine** proven on demo-1 EXECUTES the
  per-stack Directus → directus replay exits 0 → it SERVES (curl `/items/simulations` → 200 + real published
  rows on the dev offset port). **N=0 stays prod-read** (verified, M25-D4): dev defaults `LOCAL_CONTENT=0`
  + the hard N=0 set-dress guard — N=0 (`anthropos`) structurally never gets a local Directus from the auto
  flow. (Held the max-2-co-resident line: demo-1 torn down before dev-2.)
- [x] **DB-3 — Re-run idempotency**: **GREEN** — full re-run convergent; migrate no-op ×5, casbin
  "already 248 — skipping", directus bootstrap "already bootstrapped — skipping", taxonomy TRUNCATE-reload →
  329,859 (same), seed idempotent. The directus cache-miss reproduces deterministically (correct idempotency).
- [ ] **DB-4 — Cold-start capture**: **EXERCISED + GREEN** (M25-D5) — the sanctioned cold-start fill
  (`stacksnap capture --surface directus --source primary-read --dsn <marco_read>`) ran end-to-end: it
  surfaced the firewall over-capture, the fix closed it, and the **re-run PASSED the firewall**
  (`public-only=true`, `directus_files=1257`, `_structure.sql` present) — the rows-only cache upgraded
  in-place to structure-bearing. This capture IS the DB-4 cold-start path.
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

## Open questions
- **RESOLVED.** The blocker (rows-only cache) was cleared by the operator-sanctioned `primary-read` capture.
  That capture surfaced + closed a real safety bug (M25-D5, the `directus_files` tenant-leak the firewall
  caught fail-closed) and two dangling-FK bugs (M25-D6/D7), then **passed the firewall** and made the local
  Directus serve. All 5 done-bars are GREEN.
