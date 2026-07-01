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

## S2 — file surfaces (verified against the authoring copy)
- `stack-snapshot/pg/pg.go` — added `NormalizeDSN` + helpers, applied at `Connect` (was: raw `pgxpool.New(dsn)` at
  L188 which rejected `sslmode=no-verify`). 7 table tests in `pg_test.go`. rext commit `c5323a1`.
- `stack-snapshot/source/source.go` + `cmd/stacksnap/main.go:204` — already make `--dsn`→`primary-read` a
  candidate; **no change** (decision D2).
- **No `up-injected.sh` change** — per the user, capture stays the existing operator/agent step (bring-up is
  replay-only by design); no new auto-capture entry point. The fix is the sslmode normalization (so the wired DSN
  works) + the S6 doc.

## S3/S4/S5 — re-sync + recapture results (2026-06-29)
- **S3 lag (fetch+count):** next-web-app v2.88.0→**v2.89.0** (2 behind, ff'd by `make pull`); app v1.315, cms
  v0.254.1, jobsimulation v0.252.0, skiller v0.103.0, all others 0 behind. **Clones were already current.**
- **S4 recapture (primary-read over the wired MCP DSN, sslmode-normalized; public-only firewall):**
  - directus  @ `ea2e187a…` (UNCHANGED digest) — 14 tables / 11,986 rows + `_structure.sql` (425 stmts)
  - sim-embeddings @ `10146f28…` (UNCHANGED) — 4 tables / 1,490 rows
  - taxonomy @ `c75ce94d…` (UNCHANGED) — 10 tables / ~330k rows, ~1.4 GB (background recapture)
  - dry-run proved the sslmode fix end-to-end against the live wired DSN (DSN carries `sslmode=no-verify`).
- **S5:** `app` AI-readiness commits present (e.g. "AI-generated narratives for AI readiness diagnosis",
  "frequency field to AIReadinessInterviewFinding"); next-web `ai-readiness/` UI present → M201 false-negative resolved.
