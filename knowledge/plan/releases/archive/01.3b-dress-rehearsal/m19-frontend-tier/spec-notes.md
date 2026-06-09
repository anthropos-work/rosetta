# M19 — Spec Notes

_Technical notes accumulated during build — mechanisms, file paths (with line cites), gotchas, and the concrete shape of each change. Populated by `/developer-kit:build-milestone`. The verified code locations from the design-time research are in the milestone `overview.md` and `.agentspace/demo-up-issue.md`._

## Pre-flight audits — section 1 (gen_injected_override.py emits frontends)

Phase 0b KB-fidelity verdict: **GREEN**. Report: `knowledge/plan/releases/01.3b-dress-rehearsal/m19-frontend-tier/kb-fidelity-audit.md` (sha f942933).

### topic → doc → code triples (verified against real platform clone present at `stack-dev/`)

| Topic | Knowledge doc | Code path | Status |
|---|---|---|---|
| next-web build/run + ARG contract | `corpus/services/next-web-app.md` | `stack-dev/next-web-app/Dockerfile.dev` | PAIRED, ALIGNED |
| studio-desk build/run + ARG contract | `corpus/services/studio-desk.md` | `stack-dev/studio-desk/Dockerfile.dev` | PAIRED, ALIGNED |
| ant-academy native run | `corpus/services/ant-academy.md` | `stack-dev/ant-academy/code/` | PAIRED, ALIGNED |
| demo override generator | `corpus/ops/rosetta_demo.md` | `.agentspace/rosetta-extensions/stack-injection/gen_injected_override.py` | PAIRED |
| per-demo build orchestration | `corpus/ops/rosetta_demo.md` | `.agentspace/rosetta-extensions/demo-stack/up-injected.sh` | PAIRED |
| verify service list | `corpus/ops/verification.md` | `.agentspace/rosetta-extensions/stack-verify/lib/services.sh` | PAIRED |
| tooling-only frontend plan | `.agentspace/demo-up-frontend-plan.md` | (contract) | DOC-ONLY (the plan) |

### Verified ARG/port contracts (the load-bearing facts the build relies on)

- **next-web `Dockerfile.dev`** declares ARGs `NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` (default `http://localhost:5050/graphql`), `NEXT_PUBLIC_BACKEND_API_URL` (`:8082`), `NEXT_PUBLIC_HOSTING_URL` (`:3000`) — and **NO pk ARG** → pk must be baked via gitignored `apps/web/.env.local` in the build context. EXPOSE 3000, `PORT=3000`. Builds `--filter=@anthropos/web-app --concurrency=1`, `NODE_OPTIONS=--max-old-space-size=4096` (the plan's no-op note: turbo's inline 8192 wins). Compose publishes `3000:3000`.
- **studio-desk `Dockerfile.dev`** declares ARGs `VITE_CLERK_PUBLISHABLE_KEY` (pk IS an ARG here — bake directly), `VITE_GRAPHQL_ENDPOINT`, `VITE_ENVIRONMENT`, `VITE_WEB_APP_URL`, `VERSION`. Dockerfile EXPOSE/`PORT=80`, but **compose publishes `9000:9000` + `9100:9100`** with `PORT=9000`/`FRONTEND_PORT=9100` env overrides → the offset targets are 9000 + 9100 (backend + frontend).
- **gen_injected_override.py** today: `INJECTED` (5 Go svcs → `:injected` image), `REUSE_DEV` (incl. a STALE `next-web-app: anthropos-next-web-app` reuse entry — see finding KB-1), per-service offset ports via `ports: !override`, `build: !reset null`, `pull_policy: never`. `build_lines()` is the pure builder (unit-testable). New blocks will follow the INJECTED image pattern (`demo-N-next-web` / `demo-N-studio-desk`, `build:!reset null`, `pull_policy:never`, `mem_limit:1g`, `profiles:!override [graphql]`).
- **stack-verify `lib/services.sh`** SERVICES array = base ports for project `anthropos`; `service_rows()` applies STACK_OFFSET + project prefix + STACK_SERVICES filter centrally. Frontend ports register here as new rows (next-web `localhost:3000 http /`, studio-desk-fe `localhost:9100 http /`).
- **up-injected.sh** builds the 5 Go svcs in a parallel fan-out + the 2 fake servers, waits, then gens override + brings up + migrates + autoverify. Frontend builds land **serially BEFORE the compose up** (kept out of the parallel fan-out), tag-guarded.

### Finding KB-1 (incidental, YELLOW-adjacent, fix in section 1)
`gen_injected_override.py` `REUSE_DEV` includes `"next-web-app": "anthropos-next-web-app"` — but next-web is NOT in the demo's graphql profile today, so it never emitted. M19 supersedes this: next-web gets a per-demo built image (`demo-N-next-web`), not a dev-image reuse. Remove the stale `next-web-app` REUSE_DEV entry when adding the frontend emission to avoid a contradictory double-path.
