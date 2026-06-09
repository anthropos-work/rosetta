---
title: "KB Fidelity Audit — M19 frontend-tier"
date: 2026-06-09
scope: milestone:M19
invoked-by: build-milestone
---

## Verdict
GREEN

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| next-web build/run + ARG contract | `corpus/services/next-web-app.md` | `stack-dev/next-web-app/Dockerfile.dev`, platform `docker-compose.yml` next-web-app block | PAIRED |
| studio-desk build/run + ARG contract | `corpus/services/studio-desk.md` | `stack-dev/studio-desk/Dockerfile.dev`, platform compose studio-desk block | PAIRED |
| ant-academy native run | `corpus/services/ant-academy.md` | `stack-dev/ant-academy/code/` | PAIRED |
| demo override generator | `corpus/ops/rosetta_demo.md` | `stack-injection/gen_injected_override.py` | PAIRED |
| per-demo build orchestration | `corpus/ops/rosetta_demo.md` | `demo-stack/up-injected.sh` | PAIRED |
| verify service list / net | `corpus/ops/verification.md` | `stack-verify/lib/services.sh`, `stack-verify/live/autoverify.sh` | PAIRED |
| tooling-only frontend plan (the contract) | `.agentspace/demo-up-frontend-plan.md` | — | DOC-ONLY (plan; implemented this milestone) |

## Fidelity Findings

1. **next-web Dockerfile ARGs** — Doc (`next-web-app.md` env table) claims `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT`, `NEXT_PUBLIC_BACKEND_API_URL`, `NEXT_PUBLIC_HOSTING_URL` are baked at build. **Actual:** `Dockerfile.dev` declares exactly those three ARGs (lines 23-28). **ALIGNED.** Also confirms the plan's key fact: NO pk ARG → pk must be baked via gitignored `.env.local`.
2. **studio-desk Dockerfile ARGs** — Doc claims `VITE_CLERK_PUBLISHABLE_KEY` + `VITE_GRAPHQL_ENDPOINT`. **Actual:** `Dockerfile.dev` declares `VITE_CLERK_PUBLISHABLE_KEY`, `VITE_GRAPHQL_ENDPOINT`, `VITE_ENVIRONMENT`, `VITE_WEB_APP_URL`, `VERSION`. **ALIGNED** (pk is a direct ARG for studio-desk — simpler than next-web).
3. **studio-desk ports** — Doc claims 9100 (fe) / 9000 (be). **Actual:** Dockerfile EXPOSE 80 / PORT=80, but compose publishes `9000:9000` + `9100:9100` with `PORT=9000`/`FRONTEND_PORT=9100` env overrides. **ALIGNED** (doc describes the effective runtime ports; the offset targets are 9000+9100).
4. **gen_injected_override.py shape** — `build_lines()` pure builder emits per-service offset ports + injected/reuse image strategy. **ALIGNED** with the plan's intended additive-override extension point.
5. **stack-verify service list is the registration point** — `lib/services.sh` SERVICES array + central offset/project rewrite in `service_rows()`. M18 explicitly left the frontend-port hook for M19. **ALIGNED.**

## Completeness Gaps

- **KB-1 (incidental):** `gen_injected_override.py` `REUSE_DEV` carries a stale `next-web-app: anthropos-next-web-app` entry that never emits (next-web isn't in the demo graphql profile today). M19 supersedes it with a per-demo built image. Fix inline in section 1 (remove the stale entry). Recorded in `spec-notes.md`.

## Applied Fixes
- Recorded the verified topic→doc→code triples + ARG/port contracts in `spec-notes.md` (fast-start for future audits).
- No doc edits needed at audit time — the service docs are accurate. KB-1 is a code fix owned by section 1, not a doc staleness.

## Open Items (require user decision)
None.

## Gate Result
GREEN: proceed to Phase 1. All milestone topics PAIRED; all load-bearing ARG/port claims verified ALIGNED against the real platform clone at `stack-dev/`. KB-1 (stale REUSE_DEV entry) is a code fix scheduled for section 1, not a doc blocker.
