**Type:** tok (bootstrap) — authors the first strategy (TOK-01). Runs under the M226 iteration protocol
(`corpus/ops/verification.md` + `coverage-protocol.md` + `latency-budget.md`), the M215/M221 prove-on-billion lineage.

# iter-01 — work log

1. Confirmed milestone iterative; created branch `m226/opening-night` from `release/02.40-casting-call` @ 112c28c.
2. Ran Phase 0b KB-fidelity audit → **GREEN** (report `kb-fidelity-audit.md`; recorded in `spec-notes.md`).
   All 5 KB-dep docs PAIRED + aligned; the 4 hiring demo-patches present in the rext manifest + wired in
   `up-injected.sh`; tailscale-serve.md reflects the M220 default-on flip; the recruiter 3rd latency vantage is a
   declared deliverable; rext code-of-record `casting-call-m225-harden`.
3. Read-only recon of billion (peer-side + one `ssh devops@billion`) — see `overview.md` § Recon result.
4. Authored **TOK-01** in the milestone-root `decisions.md`.

## Close — 2026-07-17

**Outcome:** TOK-01 authored (initial strategy `reprove-hiring-on-billion`); billion recon captured (stale v2.3
demo up, host prereqs green, C-6 memory risk sharper for the 2-app hiring demo, rext cutover needed). No gate
progress (tok).
**Type:** tok (bootstrap)
**Status:** closed-fixed
**Gate:** N/A for tok
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (bootstrap does NOT exit) — (3) re-scope: n —
(4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue (into iter-02 tik under TOK-01)
**Decisions:** iter-local D1 (ssh user + cutover target); TOK-01 in milestone-root decisions.md.
**Side-deliverables:** none.
**Routes carried forward:**
- iter-02 (tik): the substrate cutover + first default `/demo-up 1` on billion + first 7-condition measurement
  from this Mac. Handler: `PROVE-M226-iter02-first-cold-bringup`.
**Lessons:** The last billion proof was v2.3 panorama; billion carries a stale demo + panorama rext tag, so the
first tik's real cost is a clean teardown + a full casting-call rebuild (2-app hiring image), not just a re-run.
Memory (7.3 GiB) is the sharpest live-only risk for the 2-app hiring demo.
