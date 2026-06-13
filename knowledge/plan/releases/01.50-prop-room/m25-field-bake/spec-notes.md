# M25 — Spec notes

_Technical detail that doesn't belong in `overview.md` (code maps, contracts, edge cases). Accumulates during build._

## Pre-flight audits — field-bake (Phase 0b, 2026-06-13)

KB-fidelity audit (3 parallel Explore sweeps; authoring copy `.agentspace/rosetta-extensions` @ `6a4749d`,
tags `prop-room-m21..m24`). **Verdict: YELLOW** (no blind areas, no stale load-bearing claim; two arg-hint
drifts fixed inline). Report: `kb-fidelity-audit.md`.

Topic → doc → code triples verified ALIGNED:
- local-Directus provisioning + exit-0/exit-4 split → `corpus/ops/directus-local.md`, `snapshot-spec.md`
  → `stack-snapshot/cmd/stacksnap/main.go` (`exitUnprovisioned=4`, `exitCacheMiss=5`), `autoprovision.go`.
- asset-plane-stays-prod (`DIRECTUS_PUBLIC_BASE_ADDR` left on prod; only `DIRECTUS_BASE_ADDR` re-pointed)
  → `stack-injection/gen_injected_override.py`, `stack-core/gen_override.py`.
- demo default-on / dev opt-in / N=0 `--force` guard → `dev-stack/dev-setdress.sh` (stack-type default +
  N=0 die-without-force).
- structure+rows captured together (M21 fold) → `stack-snapshot/directus/directus.go` (`CapturesStructure:true`),
  `capture/capture.go` (StructureCapturer phase 4b), `manifest/manifest.go` (`Structure` artifact).
- cold-start `--dsn` + AssertPublicOnly (two-phase: AssertPlan + AssertCaptured) + postgres-MCP-not-a-source
  (`source/source.go` Kind enum has no MCP variant; real `COPY` via `pg.Conn`) → `snapshot-cold-start.md`.
- new Directus verify probes → `stack-verify/lib/services.sh` (server/health row), `lib/readiness.sh`
  (`probe_directus_collections()`, container-gated schema expectation).
- teardown reclaims directus container + frees registry slot → `rosetta-demo` down (`ureg_release`),
  `dev-stack` down (`reg_release`).

Fixed inline (KB-1, KB-2): `demo-up` arg-hint advertised a fictional `--full` + `--no-ui`/`--no-setdress` as
`rosetta-demo` CLI flags (they're env vars on `up-injected.sh`: `DEMO_NO_UI` / `DEMO_NO_SETDRESS` /
`DEMO_NO_LOCAL_CONTENT`) → rewrote the hint to the env-var reality (body was already correct). `dev-up`
arg-hint omitted `--local-content` (in parser `dev-stack:88` + body) → added it.
