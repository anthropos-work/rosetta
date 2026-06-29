# M47 — decisions

_Implementation decisions with rationale (one entry per decision: context → options → choice → why)._

## D1 — sslmode normalization lives at the `Connect` choke point (not in `DSNForOffset`)
**Context:** the MCP capture DSN carries `sslmode=no-verify`, which pgx rejects. **Options:** (a) normalize in
`DSNForOffset` (port helper); (b) normalize in `Connect`; (c) normalize in the capture caller. **Choice:** (b) — a
pure `NormalizeDSN` helper applied inside `Connect`, the single point every pooled connection passes through.
**Why:** capture and replay both go through `Connect`; local replay DSNs lack `no-verify` so the helper is a no-op
for them (guarded by a `Contains("no-verify")` early-return → byte-identical passthrough). `DSNForOffset` stays
single-purpose (port math).

## D2 — `source.go` needs no change; the MCP DSN is just a `primary-read` `--dsn`
**Context:** the plan listed "accept the MCP DSN as a sanctioned primary-read source in `source.go`." **Finding:**
`source.go`'s precedence already includes `KindPrimaryRead`, and `cmd/stacksnap/main.go:204` already marks
`primary-read` a candidate when `--dsn` is given. **Choice (Fate-2, already covered):** no `source.go` edit — the
only real blockers were the sslmode rejection (D1, fixed) and the orchestration (extract the DSN + auto-capture,
no prompt — the remaining S2 wiring). Recorded so the section checklist reflects reality.

## KB-47-01 (YELLOW, tracked) — snapshot-cold-start.md "MCP is NOT a capture source"
`corpus/ops/snapshot-cold-start.md` (M20-D4) states the wired `postgres` MCP is **not** a sanctioned capture
source (it returns JSON rows, not COPY bytes). M47 reverses this: it uses the MCP's *configured DSN* (not the MCP
*tool*) as a `primary-read` source via direct pgx COPY. The doc is updated in **S6** (M47's `Delivers →`), same
milestone that changes the behavior. → resolve in Phase 5 / S6.

## OPEN-Q1 — where does `up-injected.sh` read the prod capture DSN from? (checkpoint)
The cold-cache auto-capture needs a prod read DSN. The wired `postgres` MCP's DSN lives in `~/.claude.json` (the
only config that defines a postgres MCP server). **Options:** (a) `up-injected.sh` reads + `jq`s the DSN out of
`~/.claude.json` automatically; (b) a dedicated env var (e.g. `SNAPSHOT_CAPTURE_DSN`) the operator exports, with
`~/.claude.json` as a documented fallback; (c) documented manual extraction only. Secret-sensitivity: the DSN is a
prod credential — whichever path, it must stay **values-blind** (never logged/echoed), per `secrets-spec.md` /
`safety.md`. **Surfaced to the user before implementing the wiring.**
