# M47 — spec notes

_Technical notes accumulate here during build (file:line surfaces, rext tag, schema findings)._

## Pre-flight audits — S2 (cold-start MCP-DSN auto-capture)

**Phase 0b KB-fidelity verdict: YELLOW** (recorded 2026-06-29).

Basis: a comprehensive KB blind-area audit was just run for exactly these topics during the v1.10b
`/developer-kit:design-roadmap` session (the "Backfill scope + KB blind-area audit" research agent) — it
confirmed M47's contract docs are adequate, with one tracked load-bearing item:

- **Load-bearing stale claim (tracked → S6 resolves it):** `corpus/ops/snapshot-cold-start.md` (the M20-D4 spike
  result) currently states the wired `postgres` MCP is **NOT** a sanctioned capture source (it returns JSON rows,
  not COPY bytes). **M47 reverses this** — it makes the MCP DSN a sanctioned **primary-read** capture source under
  `AssertPublicOnly`. The doc update is M47's `Delivers →` (section **S6**); the claim is updated in the same
  milestone that changes the behavior, so this is YELLOW (tracked), not RED (blind).
- **Adequate contracts (read as-is):** `snapshot-spec.md` (capture/replay + firewall), `safety.md`
  (`AssertPublicOnly` + capture-source policy), `secrets-spec.md` (the secret source).

Tracked KB item: **KB-47-01** — snapshot-cold-start.md MCP-not-a-source claim → updated in S6. (decisions.md)

Proceeding to Phase 1 (S2) per the GREEN/YELLOW gate.

## S2 — file surfaces (from the design-session research, to verify against the authoring copy)
- `stack-snapshot/pg/pg.go` — `DSNForOffset` (~L54): parse DSN; does not normalize `sslmode`. pgx pool rejects
  `sslmode=no-verify` (the MCP DSN's mode).
- `stack-snapshot/source/source.go` (~L84-103): `Resolve()` picks dump-ingest vs primary-read; neither
  auto-reads the MCP DSN.
- `demo-stack/up-injected.sh` set-dress (~L567-593): the cold-cache path warns/prompts instead of auto-capturing.
