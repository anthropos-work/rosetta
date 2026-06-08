---
milestone: M19
slug: frontend-tier
version: v1.3b "dress rehearsal"
milestone_shape: section
status: planned
created: 2026-06-08
last_updated: 2026-06-08
complexity: large
delivers: corpus/ops/demo/frontend-tier.md (net-new) + updated demo-up skill + frontend-emitting override generator + per-demo build in up-injected.sh
issues: ISSUE-6, ISSUE-8, ISSUE-9
---

# M19 — The demo-up frontend tier

## Goal
`/demo-up` brings up the full UI — **next-web + studio-desk** at offset ports (per-demo **cached** image from the
**unmodified** platform Dockerfile) **+ ant-academy** natively — so a demo is actually demoable, on a 16 GB Mac.

## Why section
The build-cost question is already investigated end-to-end and the tooling-only solution is **verified with real
Docker builds** ([`.agentspace/demo-up-frontend-plan.md`](../../../../.agentspace/demo-up-frontend-plan.md)): the
override blocks, the per-demo build-args, the Clerk-pk baking, the `.dockerignore`, the VM size. The `In:` list is
writable with confidence. (The largest milestone of v1.3b, but not uncertain.)

## Repo split
- **`rosetta-extensions`** (authoring → tag `dress-rehearsal-m19` → consume): the override generator + the per-demo
  frontend build in `up-injected.sh` + the VM preflight + the sibling `.dockerignore`.
- **`rosetta`**: the net-new `corpus/ops/demo/frontend-tier.md` + the `demo-up` SKILL.md update.

## Scope
- **In:**
  - **`rosetta-extensions`:**
    - Extend `stack-injection/gen_injected_override.py` to **emit `next-web-app` + `studio-desk`** —
      `profiles:!override [graphql]`, `ports:!override` (offset), `image: demo-N-*` + `build:!reset null` +
      `pull_policy:never`, `mem_limit:1g` (today it emits backend-only). (ISSUE-8.)
    - `up-injected.sh` **builds the two frontends serially, before compose up**, from the **unmodified** Dockerfiles
      with **offset-URL build-args** (`NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` etc. / `VITE_GRAPHQL_ENDPOINT`) **+ the
      minted Clerk pk** via a gitignored `.env.local`/BuildKit overlay; **tag-guarded for cache reuse**. (ISSUE-9.)
    - Ship a sibling `next-web.Dockerfile.dockerignore` (5.6 GB context → <100 MB). (ISSUE-9.)
    - **Launch (or document) ant-academy natively** — port 3077, its own `ant-academy/code/.env`,
      `REQUIRE_ORGANIZATION_MEMBERSHIP=0`. (ISSUE-8.)
    - A **12 GB Docker-VM pre-flight assert** in `up-injected.sh`. (ISSUE-6/ISSUE-9.)
    - **Register the frontend ports** so M18's verify net covers them.
    - **Default-on + skippable** (`--no-ui`).
  - **`rosetta`:** update the `demo-up` SKILL.md (UI is now in scope) + author `corpus/ops/demo/frontend-tier.md`.
- **Out:** the **optional upstream platform PR** for *true* zero-rebuild (runtime rewrites + `__env.js` +
  `output:standalone`) — it edits platform repos → **forbidden / user-owned**, documented as a follow-up (the v1.4
  deploy-CI precedent), not built here.

## Depends on
**M18** (so the verify net covers the new frontend services). **Parallel with:** none.

## Open questions (resolve during build)
- ant-academy native-launch *by* the tool vs documented-manual — lean: launch it, fall back to a documented step if
  the native run proves fiddly.
- Pre-warm the frontend image cache as part of `dev-up` too — defer (demo-first).

## KB dependencies (read as contract)
- [`.agentspace/demo-up-frontend-plan.md`](../../../../.agentspace/demo-up-frontend-plan.md) (the verified tooling-only plan)
- `corpus/services/next-web-app.md`, `corpus/services/ant-academy.md`, `corpus/services/studio-desk.md` *(if present)*
- `corpus/ops/rosetta_demo.md` (the demo lifecycle + ports)

## Delivers
- **→ rosetta:** `corpus/ops/demo/frontend-tier.md` — net-new: the ports, the per-demo build, the Clerk-pk baking,
  the 12 GB VM prereq, the honest "one ~3-min cached build per new demo-N" residual; + the updated `demo-up` skill.
- **→ rosetta-extensions:** the frontend-emitting `gen_injected_override.py` + the per-demo frontend build in
  `up-injected.sh` + the VM preflight + the `.dockerignore`.

## Risk
**(scope+resource, blocks-quality)** the ~3.7 GB / ~3 min per-frontend build swap-thrashes on an undersized VM (the
original "hour"). Mitigate: 12 GB Docker-VM preflight assert, serial cached builds, `mem_limit`, the sibling
`.dockerignore`. **Hard line: zero platform-repo edits — the repo is a build *context* only, the Dockerfile is
unmodified (verified achievable).**
