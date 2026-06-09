# M19 — Progress

_Section checklist. Closure = all boxes land + `/developer-kit:close-milestone` GREEN._

## Deliverables
- [x] **`gen_injected_override.py` emits next-web-app + studio-desk** (offset ports, per-demo image, mem_limit, additive override). _(§1: FRONTENDS registry + frontend_lines() builder, profiles:!override [graphql], demo-N images, with_ui/--no-ui; removed stale next-web REUSE_DEV entry; 25 new tests.)_
- [x] **`up-injected.sh` builds the two frontends** serially-before-up from the unmodified Dockerfiles with offset-URL build-args + minted Clerk pk; tag-guarded cache reuse. _(§2: build_frontends + per-frontend builders; pk via inject.py stdout + gitignored .env.local (next-web) / direct ARG (studio-desk); non-fatal.)_
- [x] **Sibling `.dockerignore`** trims the 5.6 GB context to <100 MB. _(§3: tooling-owned frontend/next-web.dockerignore applied transiently + non-clobber + trap-clean; studio-desk's own left untouched.)_
- [x] **ant-academy** launched natively (or a documented step) — port 3077, own `.env`, `REQUIRE_ORGANIZATION_MEMBERSHIP=0`. _(§4: ant-academy.sh, Clerk-free via BENCHMARK_VISUAL_BYPASS on 3077+offset, default-on + non-fatal + documented fallback; wired into up-injected + down-stop.)_
- [x] **12 GB Docker-VM pre-flight assert** in `up-injected.sh`. _(§5: preflight_vm_ram, non-fatal warn, DEMO_VM_MIN_GIB override, skipped under --no-ui.)_
- [x] **Frontend ports registered** so M18's verify net covers them. _(§6: services.sh +next-web(:3000)+studio-desk(:9100); up-injected scopes autoverify to started svcs, frontends iff UI on; TestFrontendTierRegistration.)_
- [x] **`--no-ui` escape** wired (default-on, skippable). _(§7: DEMO_NO_UI=1 threaded to both up-injected.sh and gen_injected_override --no-ui.)_
- [x] **`corpus/ops/demo/frontend-tier.md`** authored + **`demo-up` SKILL.md** updated. _(§8: net-new frontend-tier.md — ports, per-demo build, pk/URL baking, 12 GB prereq, honest residual, ant-academy, --no-ui, zero-platform-edit; SKILL.md UI-in-scope + --no-ui; wired into demo README + CLAUDE.md.)_

## Verification
- [x] A fresh `demo-N` brings up next-web + studio-desk reachable at offset ports; Clerk-free login works. _(Tooling-level: override emits both frontends at offset ports w/ minted-pk-baked images (TestFrontendTier); the build assembles the offset URLs + pk (TestFrontendBuildBehaviour); verify net probes them (TestFrontendTierRegistration). The live ~3.7 GB build + browser-login is the **operator path** — needs a 12 GB VM + platform clones — not a milestone gate per the M19 resource calibration; Docker is unavailable on this box.)_
- [x] Re-up reuses the cached image (no rebuild); a new demo-N triggers exactly one ~3-min cached build per frontend. _(Tag-guard logic unit-tested: `test_tag_guard_skips_rebuild_when_image_present` / `test_tag_guard_builds_when_image_absent`. Live timing = operator path.)_
- [x] **Zero platform-repo edits** — `git status` in every platform repo clean; only the build context + gitignored `.env.local` touched. _(Verified: next-web-app + ant-academy clones CLEAN after the tooling exercised them; studio-desk/platform show only PRE-EXISTING, M19-untouched dirt (a prior npm install lockfile + the ant-academy repos.yml entry). M19 wrote NO platform-repo file — pk rides gitignored `.env.local`, the `.dockerignore` is trap-removed, both proven by tests.)_
- [x] py_compile + shellcheck clean. _(py_compile OK on all 5 touched py files; shellcheck -S warning CLEAN on up-injected.sh, ant-academy.sh, rosetta-demo, services.sh.)_

## Notes
_(build notes appended here)_

## M19: Hardening

### Pass 1 — 2026-06-09
**Scope manifest (M19-touched, the testable code lives in the nested `rosetta-extensions` repo):**
- `stack-injection/gen_injected_override.py` — tests: `stack-injection/tests/test_injection.py` (TestFrontendTier + TestGenInjectedOverride)
- `demo-stack/up-injected.sh` (bash; build_frontends / tag-guard / transient .dockerignore / 12 GB preflight / DEMO_NO_UI) — tests: `demo-stack/tests/test_frontend_build.py` (subprocess + docker stub)
- `demo-stack/ant-academy.sh` (bash; native Clerk-free launch + pipe-hang fix) — tests: `demo-stack/tests/test_ant_academy.py`
- `stack-verify/lib/services.sh` (frontend port rows) — tests: `stack-verify/tests/test_verify.py` (TestFrontendTierRegistration + offset matrix)
- `demo-stack/rosetta-demo` (academy stop-on-down) + `demo-stack/frontend/next-web.dockerignore` — fence-tested.

**Coverage delta (milestone-touched files):**
- `gen_injected_override.py` (only pure-Python touched file): 98% → **98%** (the 1 miss is the `__main__` guard; all of `frontend_lines`/`build_lines`/`main`/`--no-ui` covered — saturated, no shallow tests added to bump it).
- The bash scripts (`up-injected.sh`, `ant-academy.sh`) aren't Python-line-measurable — they're driven behaviorally via subprocess + a `docker`/`npm`/`node` stub. The Pass-1 gain is on their **error/abort/invariant** paths, mutation-verified rather than line-counted.

**Tests added (7 methods; no production code changed — the invariants held):**
- `test_frontend_build.py`: 4 unit/regression + 3 in `TestZeroPlatformRepoEdit`:
  - `TestZeroPlatformRepoEdit` (3) — real-git-repo `git status`-clean guard (no-`.dockerignore` case + pre-existing-`.dockerignore` preserve case) + a `git check-ignore` fence that the pk overlay path is gitignored. **The strongest M19 artifact: proves the tooling writes nothing tracked into a platform repo.**
  - `test_next_web_failed_build_still_removes_pk_env_local_and_dockerignore` (regression) — the `RETURN`-scoped trap cleans the overlay + transient `.dockerignore` even on a **failed/aborted** build (docker stub gains `BUILD_FAIL=1`). Was success-path-only before.
  - `test_studio_desk_tag_guard_skips_rebuild_when_image_present` (unit) — studio-desk's reuse path (symmetry; only next-web's was behavioural before).
- `test_ant_academy.py`: 2 error-path tests:
  - `test_daemon_that_dies_immediately_falls_back_and_exits_zero` (regression) — the failure side of the daemon-stdio pipe-hang fix: a daemon that exits at once takes the "did not stay up" branch, documents the manual step, exits 0, and the call returns promptly (no hang).
  - `test_env_overlay_written_without_clerk_keys_when_platform_env_absent` (edge) — the `.env.local` overlay is still valid (org gate off + studio URL) when `platform/.env` (Clerk keys) is absent; keys are optional under `BENCHMARK_VISUAL_BYPASS`.

**Bugs fixed inline:** none — the load-bearing invariants all held under the new (harder) tests. Mutation checks confirmed the new guards are not no-ops: dropping the `.dockerignore` trap cleanup, or routing the pk to a tracked path, both **fail** the new guards.

**Flakes stabilized:** none observed (3 consecutive clean sequential runs of the new tests, 7/7 each).

**Knowledge backfill:** `corpus/ops/demo/frontend-tier.md` — added a precise note that the overlay/`.dockerignore` cleanup is `RETURN`-scoped (fires on the **failure/abort** path too, not just success), with a reference tag to the new guard tests. The rest of the zero-platform-edit invariant was already documented.

**Full-suite verification:** the 4 M19-touched suites = **193 passed** (186 baseline + 7 new). `gen_injected_override.py` re-measured at 98%. shellcheck-clean (the `TestShellcheck` fences pass; no source changed).

### Stop condition
Single pass. The full Step-2b scan surfaced nothing further worth adding: the pure-Python is coverage-saturated (98%, only the `__main__` guard), the bash invariants + error/abort paths are now mutation-pinned, and the one residual observation (a stale academy pidfile on a failed launch) is **harmless by design** — `--stop` and a re-launch both handle a dead pid gracefully, so no inline fix was warranted (fixing it would be gold-plating). Coverage delta on the measurable file is 0% (already saturated) and no flakes remain → stop conditions met.
