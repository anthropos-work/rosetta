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

## OPEN-Q1 — where does the cold-start capture read the prod DSN from? → RESOLVED (user)
The cold-cache capture needs a prod read DSN. The wired `postgres` MCP's DSN lives in `~/.claude.json`.
**Resolution (user, 2026-06-29):** reuse the **existing wired DSN** — **no new process or entry point.** The
capture stays the existing operator/agent step (per `snapshot-cold-start.md`; `/demo-up` is replay-only by design).
The agent extracts the wired DSN **values-blind** and passes it to `stacksnap capture --source primary-read --dsn`.
The S6 doc documents this (Option 2b). No env var, no `up-injected.sh` reader. (See D3.)

## D3 — the `up-injected.sh` cold-cache auto-capture wiring was DESCOPED (per user)
**Context:** the original S2 plan listed "drop the cold-cache prompt → auto-capture" + "wire the auto-capture into
`up-injected.sh`." **User direction (2026-06-29):** *no new process or entry point — use the same process the
demo-up flow already uses.* **Choice:** do NOT add an auto-capture hook to `up-injected.sh`. The bring-up stays
replay-only (the existing design); cold-start capture remains the existing operator/agent step, now **turnkey**
because (a) the sslmode fix (D1) makes the wired DSN connect, and (b) S6 documents the values-blind wired-DSN
invocation. **Why:** the only real blockers were the sslmode rejection (fixed) + a documentation gap (fixed) — not
a missing mechanism. Adding a new shell entry point would duplicate the existing capture process. This is a
conscious scope revision, not a dropped item — the *capability* (turnkey cold-start capture) is delivered.

## Adversarial review (Phase 2c)
- **Scenario — `no-verify` in a non-sslmode position** (password / userinfo / `application_name` / dbname): could
  `NormalizeDSN`'s `Contains("no-verify")` early-enter cause a spurious rewrite? **Verified handled:** the form-
  specific code only rewrites the **sslmode** query-param (URL) or **sslmode** keyword token; all other text
  (incl. the userinfo/password preserved verbatim via `dsn[:q+1]`) is untouched. **Two regression tests added**
  (rext `f30194c`): `url no-verify in password is untouched`, `keyword no-verify in a non-sslmode field is untouched`.
- **Scenario — quoted keyword value** (`sslmode='no-verify'`): NOT normalized (the quotes make `EqualFold` miss).
  **Accepted as a known limitation:** the real capture DSN (the MCP `marco_read`) is **URL-form**, and libpq
  keyword DSNs effectively never quote `sslmode`. Documented here; not a fix-now (no real input hits it, and the
  live dry-run proved the real DSN connects).

## Phase 1b — deferral re-audit verdict: GREEN (inline)
M47's deferral surface is tiny + conscious: OPEN-Q1 RESOLVED (above); the consumption-clone **re-pin** is a
push-gated KEEP (tracked with the release's other pending origin pushes — not a repeat-deferral); D3 is a scope
revision, not a deferral. No repeat-deferral pattern. Ran inline rather than spawning `/developer-kit:audit-deferrals`
given the trivial surface (one pure function + docs + an operational recapture).
