# iter-17 decisions

| ID | Decision | Rationale | Date |
|----|----------|-----------|------|
| D1 | Use the `AuthorizationService/Reload` RPC as the PRIMARY reload (docker-restart fallback) | Read the sentinel handler (READ-ONLY): `rpcsrv/rpc.go:400` `Reload` calls `s.e.LoadPolicy()` directly and returns 200 only on success → a 2xx PROVES the policy re-loaded (self-proving, no separate probe). It's faster + lighter than a container restart (no service downtime). `docker restart demo-N-sentinel-1` is the fallback (re-runs LoadPolicy at startup) when the RPC is unreachable. | 2026-06-25 |
| D2 | Reload block lands in `demo-stack/up-injected.sh` after set-dress, before verify; non-fatal; opt-out env | up-injected.sh is the demo bring-up orchestrator; the reload belongs after the set-dress seed (which wrote the g3 grant) and before the auto-verify (so the casbin/sim probes see the reloaded policy). Gated on `$NO_SETDRESS != 1` (no seed → nothing to reload). NON-FATAL (M18/M19 pattern — a reload bug never blocks a good stack); `DEMO_NO_SENTINEL_RELOAD=1` opts out. The offset port is `8087 + N*10000` (the established demo port-offset math). | 2026-06-25 |
