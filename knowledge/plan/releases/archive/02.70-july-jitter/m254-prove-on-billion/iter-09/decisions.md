# M254 · iter-09 — decisions

## D1 — pt-world reset-to-seed landed (the demo is now the Playthrough world)
`run-playthroughs.sh 1 --reset-only` on billion: the REAL `stackseed --reset` (full FK-ordered TRUNCATE,
per-stack, N=0-guarded) then a fresh `pt-world.seed.yaml` seed, the 30-identity roster re-exported to
`fake-fapi-roster.json`, the fake services restarted, fake-FAPI HTTP 200. DONE_rc=0. The billion demo now
carries the decoupled Playthrough seed (test data ≠ demo data), ready for the browser suite.

## D2 — the browser Playthrough suite runs + passes from this peer (proven-to-run)
`PT_HOST=billion.taildc510.ts.net PT_APP_SCHEME=https ./run-playthroughs.sh 1` from this workstation:
`Running 117 tests using 1 worker` (serial per spec §5.7), the first 6 GREEN including the manager
`@pt:pt-activity-drilldown` Playthrough (24 s) + the ai-readiness locator unit specs. So the harness + the
demo + the seed all cohere for the Playthroughs. A fully-detached `setsid` invocation silently failed to
create its redirect log (a plumbing quirk in this environment); the foreground run is the reliable driver.

## D3 — full-suite completion + dispositions routed to a fresh agent (cap-reached, tik 5)
This is the 5th tik of the session (iter-05..09), following an eventful disruption-recovery. Per the
prove-on-billion "fresh agent per run" discipline (context saturates on these long live-prove sessions), the
full 117-test suite wall-clock + the (f)-FCP-p95 and (g)-testhealth coordinator dispositions + the
harden/close are the next agent's work, on the pt-world-seeded billion demo. Gate stands at ~6/8 MET
(a, b-disp, c[4/4 render], d, f-session-carry) with e+h proven-to-run.
