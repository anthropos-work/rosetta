# M19 — Progress

_Section checklist. Closure = all boxes land + `/developer-kit:close-milestone` GREEN._

## Deliverables
- [ ] **`gen_injected_override.py` emits next-web-app + studio-desk** (offset ports, per-demo image, mem_limit, additive override).
- [ ] **`up-injected.sh` builds the two frontends** serially-before-up from the unmodified Dockerfiles with offset-URL build-args + minted Clerk pk; tag-guarded cache reuse.
- [ ] **Sibling `.dockerignore`** trims the 5.6 GB context to <100 MB.
- [ ] **ant-academy** launched natively (or a documented step) — port 3077, own `.env`, `REQUIRE_ORGANIZATION_MEMBERSHIP=0`.
- [ ] **12 GB Docker-VM pre-flight assert** in `up-injected.sh`.
- [ ] **Frontend ports registered** so M18's verify net covers them.
- [ ] **`--no-ui` escape** wired (default-on, skippable).
- [ ] **`corpus/ops/demo/frontend-tier.md`** authored + **`demo-up` SKILL.md** updated.

## Verification
- [ ] A fresh `demo-N` brings up next-web + studio-desk reachable at offset ports; Clerk-free login works.
- [ ] Re-up reuses the cached image (no rebuild); a new demo-N triggers exactly one ~3-min cached build per frontend.
- [ ] **Zero platform-repo edits** — `git status` in every platform repo clean; only the build context + gitignored `.env.local` touched.
- [ ] py_compile + shellcheck clean.

## Notes
_(build notes appended here)_
