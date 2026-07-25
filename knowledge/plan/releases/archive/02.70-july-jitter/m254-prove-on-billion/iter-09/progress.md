**Type:** tik (5th — the cap) — TOK-01 cluster 4 (mutating tail): (e) builder Playthrough + (h) Playthroughs.

# M254 · iter-09 — progress

## Close — 2026-07-25

**Outcome:** pt-world reset-to-seed LANDED on billion; the browser Playthrough suite is PROVEN to run + pass
from this peer (117 tests serial, first 6 green incl a manager Playthrough). The full-suite completion + the two
dispositions route to a fresh agent (the prove-on-billion "fresh agent per run" discipline — context saturated).

**Type:** tik
**Status:** closed-fixed-partial
**Gate:** (e)+(h)-Playthroughs proven-to-run (6/6 green foreground); full 117-test completion routed. Overall gate ~6/8 MET (a,b,c,d,f-session-carry) + dispositions pending.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: **y** (5th tik of the session — iter-05..09; big Playthrough-suite completion + dispositions best done by a fresh agent per the prove-on-billion discipline) — (6) protocol-stop: n — Outcome: exit-5 (cap-reached).
**Decisions:** D1 pt-world reset landed; D2 browser suite proven-to-run (detached-log plumbing needs a fix); D3 full-suite + dispositions routed to a fresh agent.
**Side-deliverables:** none (0 rext/platform edits this iter — reset + measurement).
**Routes carried forward (→ fresh agent, iter-10+):**
- **(e)+(h)** — run the full browser Playthrough suite to completion (proven-to-run; the demo is pt-world-seeded now): `cd .agentspace/rosetta-extensions/playthroughs/e2e && PT_HOST=billion.taildc510.ts.net PT_APP_SCHEME=https ./run-playthroughs.sh 1` (foreground works; the detached redirect needs a plumbing fix or use the Bash long-timeout foreground pattern). Assert all green (16 live Playthroughs + specs) + the builder Playthrough (e).
- **(f)-FCP-p95** — coordinator disposition (p50 <1s shell fix holds; p95 tailnet-jitter-bound).
- **(g)-testhealth** — the 6 routed host-sensitive tests (`FIX-M254-g-testhealth`) → coordinator fate + the fix batch.
- **(c)-academy-durability** — native academy loses the Back-to-Cockpit item on a long-running demo (fresh bring-up renders it) → academy-durable follow-up.
**Lessons:** (1) The Playthrough runner splits reset (on-host, docker+DB) from browse (tailnet peer) — `--reset-only` first, then the peer run with NO reset flag. (2) A `setsid` fully-detached Playwright run silently failed to create its redirect log here; the reliable pattern is the Bash long-timeout foreground (which the harness backgrounds cleanly to a task file) — the same pattern the billion reset-to-seed used.
