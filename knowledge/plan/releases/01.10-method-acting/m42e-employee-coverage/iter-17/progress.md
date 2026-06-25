# iter-17 progress

**Type:** tik (production-fix). Active strategy: **TOK-10**. P5 (FATAL) — post-set-dress Sentinel-reload
reproducibility fix.

## Phase B — confirm the Reload RPC re-runs LoadPolicy (sentinel source, READ-ONLY)
- `sentinel/internal/rpcsrv/rpc.go:400` `func (s *RPCHandler) Reload(...)` calls `s.e.LoadPolicy()` directly
  and returns `connect.NewError(CodeInternal, ...)` on failure, else a 200 `ReloadResponse{}`. So a **2xx
  PROVES the in-memory enforcer re-read the casbin policy** — not just a liveness probe. D1.
- The Connect service path is `/sentinel.authorization.v1.AuthorizationService/Reload`; the sentinel
  container's internal `PORT=8087` → demo-N external `8087 + N*10000` (demo-3 = 38087, confirmed
  `docker ps`). The no-watcher property is the root: `internal/authorization/casbin.go:75` calls
  `LoadPolicy()` only at enforcer construction.

## Phase C — fix (rext bring-up script + corpus doc, zero platform edit)
- `demo-stack/up-injected.sh`: NEW post-set-dress Sentinel-reload block — inserted AFTER the set-dress
  block (which wrote the g3 grant) and BEFORE the auto-verify (so the casbin/sim probes see the reloaded
  policy). Prefers the `AuthorizationService/Reload` RPC (`curl -fsS -m 10 -X POST …:$((8087+OFFSET))/…
  /Reload`); falls back to `docker restart demo-$N-sentinel-1` when the RPC is unreachable. Runs only when
  the set-dress actually ran (`$NO_SETDRESS != 1`); NON-FATAL (the M18/M19 pattern — a reload miss WARNS,
  never aborts); `DEMO_NO_SENTINEL_RELOAD=1` opts out. `bash -n` + shellcheck clean. D2.
- `corpus/ops/rosetta_demo.md`: NEW "Post-set-dress Sentinel policy reload (M42e P5)" section — documents
  the no-watcher root cause, the migrate-before-seed ordering, the RPC-2xx-proves-LoadPolicy property, the
  fallback, and that it precedes the verify.

## Phase D — prove BOTH reload paths on demo-3 (the exact bug state)
- The bug state was LIVE: a reset+re-seed wrote 341 g3 grants into `sentinel.casbin_rules` (Maya's
  membership `957cc282…` among them) but the running enforcer was stale (no reload since the seed).
- **RPC path:** the EXACT command the script runs (offset port 38087) → `curl … /Reload` returned
  **HTTP 200** (exit 0). Sentinel log shows `path=/sentinel.authorization.v1.AuthorizationService/Reload`.
  Since the handler returns 200 only after `LoadPolicy()` succeeds, the 341 g3 grants (incl. Maya's) are
  now in the in-memory enforcer.
- **Fallback path:** `docker restart demo-3-sentinel-1` → sentinel came back up + answered the Reload RPC
  200 within ~4s, with all 341 g3 grants present in the DB (re-loaded at startup). Both paths verified.
- Downstream effect (established in iter-09, re-confirmed by the mechanism): with the g3 grant in the
  enforcer, `organizationFeatures` resolves `FEATURE_JOB_SIMULATIONS` → the `/sim/<slug>/start` deny modal
  is gone, the launch UI renders. No manual restart needed.

## Close — 2026-06-25
**Outcome:** P5 landed. A fresh demo-up now reloads Sentinel's Casbin policy after the set-dress seed
(RPC-preferred, docker-restart fallback, non-fatal) — the g3 sim entitlement reaches the running enforcer
with ZERO manual step. Both reload paths proven on demo-3 (RPC 200 = LoadPolicy re-ran; restart reloads 341
grants).
**Type:** tik (production-fix)
**Status:** closed-fixed
**Gate:** NOT MET (P5 of P0–P8; the believability gate needs the avatar decision + P6–P8 + the P7 semantic harness)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (3rd tik) — (6) protocol-stop: n — Outcome: continue (run scope T1→P4→P5 complete; natural phase boundary)
**Decisions:** D1 (Reload RPC re-runs LoadPolicy — 2xx proves it), D2 (reload block in up-injected.sh, RPC+restart, non-fatal) — see ./decisions.md.
**Routes carried forward:** P4 avatar (user-blocker), P6 (library capture-path), P7 (semantic harness),
P8 (fresh-demo-up acceptance) — later runs; the AUTHORITATIVE fresh-demo-up proof of the reload-in-sequence
is part of P8 acceptance (this iter proved the mechanism + both paths on the live stack).
**Lessons:** a Connect-unary RPC whose handler returns 2xx ONLY on a successful side-effect (here
LoadPolicy) is a self-proving reload — a 200 is the proof, no separate "did it reload?" probe needed.
Confirm the handler body (not just the endpoint name) before relying on an RPC as a reload mechanism.
**rext:** commit `d945821`, tag `method-acting-m42e-iter17`.
