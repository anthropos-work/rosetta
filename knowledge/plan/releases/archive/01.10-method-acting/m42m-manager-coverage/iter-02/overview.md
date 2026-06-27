---
iteration_type: tik
status: closed-no-lift
---

# iter-02 — tik (TOK-01 line 1): the Studio-link escape — DIAGNOSE + decide rewrite vs re-scope

**Type:** tik under TOK-01 (line 1, the highest-leverage residual: escapes=139, all `studio.anthropos.work`).

## Step 0 — Re-survey (mandatory)
TOK-01 named the Studio-link escape as the iter-02 target. Re-confirmed: escapes=139 (iter-23 smoke-sweep,
all `studio.anthropos.work`) is the largest single residual + the user's explicit complaint ("clicking Studio
goes to production"). Target stands.

## Active strategy reference
TOK-01 (bootstrap), line 1.

## Cluster / target
The baked `studio.anthropos.work` left-nav "Studio" link (the manager/enterprise nav renders it on every
authenticated page; the employee nav omits it → employee had 0 escapes). One root cause → all 139.

## Hypothesis
The Studio link host is a `NEXT_PUBLIC_*_URL` the demo injection can rewrite to the demo-local studio-desk
offset port (`:39000`). If so → rewrite in the demo injection (rext, zero platform edit); if hardcoded →
re-scope trigger.

## Phase plan
Phase A/B diagnose (read-only, against the stack-demo next-web source + the live container) → decide:
land the rewrite (Phase C) OR escalate as re-scope trigger.

## Outcome: RE-SCOPE TRIGGER (hypothesis FALSIFIED — the host is NOT env-rewritable)

The diagnosis (decisions.md D1) proves the Studio link is **NOT cleanly env-rewritable to the demo-local
offset port without a platform-repo edit**:

- `STUDIO_URL` (`packages/core-js/src/constants/urls.ts:12-15`) is a **`NEXT_PUBLIC_NODE_ENV === 'development'`
  ternary** → `http://localhost:9000` (dev) **or** `https://studio.anthropos.work` (else). There is **NO
  `NEXT_PUBLIC_STUDIO_URL` override** in the source (unlike `ACADEMY_URL` on line 16-17, which DOES read
  `NEXT_PUBLIC_ACADEMY_URL` — and which is exactly why ant-academy's demo can rewrite its Studio link but
  next-web cannot).
- The Studio nav item (`useNavbarSections.tsx:283`, `studioMenuItem`, `key: STUDIO_URL`,
  `type: MenuType.outbound`) is the rendered left-nav link.
- **Live confirmation (demo-3 container):** `NEXT_PUBLIC_NODE_ENV=[]` (unset), `NODE_ENV=production` →
  `STUDIO_URL` resolves to `https://studio.anthropos.work`, baked into the client bundle (grep-confirmed in
  `/app/apps/web/.next`).
- The demo next-web build (`up-injected.sh::build_frontend_next_web`) bakes only 3 build-args
  (`NEXT_PUBLIC_WUNDERGRAPH_ENDPOINT` / `_BACKEND_API_URL` / `_HOSTING_URL` — the only ARGs the Dockerfile
  declares) and writes a gitignored `apps/web/.env.local` overlay (the pk). It does NOT set
  `NEXT_PUBLIC_NODE_ENV`.
- The **only available knob** — adding `NEXT_PUBLIC_NODE_ENV=development` to the gitignored `.env.local`
  overlay — is **broad and wrong**: (1) it points Studio at `:9000`, but the demo studio-desk is on the
  **offset** port `:39000` (`docker ps`: `0.0.0.0:39000->9000/tcp`), so it would still be a broken link, not a
  demo-local one; and (2) `NEXT_PUBLIC_NODE_ENV` ALSO gates `WEB_APP_URL` (→`:3000`), `HIRING_APP_URL`
  (→`:3001`), and `TablePagination` — a global dev-flip would create NEW wrong-port links across other
  manager-reachable surfaces (`PublicFooter` privacy/terms, `skillPathRoute` public skill-path URLs,
  `ShareAISim`, PDF download), i.e. it would introduce new escapes/broken-links, not remove one.

There is no sanctioned response-rewriting proxy in the demo stack (the injection operates at the
docker-compose/env layer, not at the HTML layer). So the only clean fix is a **platform-source edit** to
`urls.ts` (add a `process.env.NEXT_PUBLIC_STUDIO_URL ||` read, mirroring `ACADEMY_URL`) — **forbidden** by the
v1.10 zero-platform-edit line.

**This is the milestone's Re-scope trigger** (coverage-protocol.md §"Re-scope trigger (the zero-edit line)" +
the milestone overview's re-scope clause). Recorded in the milestone-root `decisions.md` (RESCOPE-1 + D1).
Exit `EXIT_REASON: re-scope-trigger`. No platform edit made; no fix landed (hypothesis falsified — the
complete iter outcome is the characterization).
