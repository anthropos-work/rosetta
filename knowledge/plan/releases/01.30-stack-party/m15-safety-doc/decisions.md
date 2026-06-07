# M15 — Decisions

_Implementation decisions with rationale. ID scheme: M15-D1, M15-D2, …_

| ID | Decision | Rationale | Date |
|----|----------|-----------|------|
| M15-D1 | Doc home = `corpus/ops/safety.md` (resolves Q1). | Ops-adjacent to seeding-spec / snapshot-spec / db-access (the three docs it cross-links and consolidates from); a reader operating a stack looks under `corpus/ops/`, not `corpus/architecture/`. | 2026-06-07 |
| M15-D2 | safety.md names the REAL Go functions (`AssertPlan` + `AssertCaptured`) alongside the conceptual umbrella name `AssertPublicOnly`. | KB-fidelity finding 1: `AssertPublicOnly` is a conceptual name (the firewall "invoked twice, defense in depth"), not a Go symbol. The existing docs use the umbrella name; safety.md keeps that convention BUT names the two grep-able functions so a reader can find them. Accuracy over prose. | 2026-06-07 |
| M15-D3 | safety.md must NOT claim an offline pg_dump-FILE reader. | KB-fidelity finding 3: the direct file reader was DROPPED in M9b-D9. A dump is ingested by RESTORING into Postgres and read over a DSN. The capture is always a DSN read (dump-ingest default → throttled primary-read fallback, both `SET TRANSACTION READ ONLY`). | 2026-06-07 |
| M15-D4 | safety.md describes the n=0-dev guard PRECISELY: it gates (a) unsolicited auto-set-dressing (`dev-setdress.sh`) and (b) destructive `--reset` (`stackseed`) — NOT all tools. `stacksnap` replay has no N=0 guard (correctly — replay writes only public data into the stack's own isolated stores). | KB-fidelity finding 5: the `dev-setdress.sh:20` source comment over-states "stacksnap/stackseed independently refuse N=0 too" — stacksnap does not. The "doubled in M13" framing = two independent enforcement points (auto-set-dress + reset), which is accurate. Pre-existing comment inaccuracy flagged, not fixed here (out of M15's diff — it's M13 code; noted for a future touch). | 2026-06-07 |

## Open at design (RESOLVED)
- M15-Q1 → **resolved (M15-D1):** doc home is `corpus/ops/safety.md`.

## KB-fidelity (Phase 0b) verdict
- **GREEN** (2026-06-07). Report: `kb-fidelity-audit.md`. Every read-side/write-side safety claim ALIGNED with the actual extensions code. Two accuracy guardrails carried into the build → M15-D3 (no offline file reader) + M15-D4 (precise n=0 scope).
