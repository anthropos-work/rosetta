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
- [ ] A fresh `demo-N` brings up next-web + studio-desk reachable at offset ports; Clerk-free login works.
- [ ] Re-up reuses the cached image (no rebuild); a new demo-N triggers exactly one ~3-min cached build per frontend.
- [ ] **Zero platform-repo edits** — `git status` in every platform repo clean; only the build context + gitignored `.env.local` touched.
- [ ] py_compile + shellcheck clean.

## Notes
_(build notes appended here)_
