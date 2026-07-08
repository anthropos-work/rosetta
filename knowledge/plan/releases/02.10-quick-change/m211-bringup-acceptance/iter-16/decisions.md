# iter-16 — Decisions

**D1 — The all-10-fail is a reset-to-seed WORLD-completeness gap, not 10 defects.** Every Playthrough failed
identically at `cockpit-login: POST /v1/demo/select(pt-*) -> 400 unknown_identity`. reset-to-seed swapped the
DB to `pt-world` but the fake-FAPI/BAPI resolve identities from a mounted `/roster/roster.json` baked at
BRING-UP from the demo's preset (demo-1 came up for coverage → stories heroes). M204 never hit this because it
brought its demo up `pt-world`-native (M204 iter-01: "the pt-manager seat is already wired in demo-1's roster
export"). So the fix belongs in the reset-to-seed lifecycle (`run-playthroughs.sh --reset`), completing it to
reset the WHOLE world (DB + identities), NOT a per-run manual roster hack — otherwise the Playthroughs can only
run on a seed-native demo, defeating the "run on any demo-N" contract.

**D2 — Refresh mechanics: discover the mount via docker-inspect, restart both fake services, wait for FAPI.**
The roster is a pure function of the seed (`stackseed --roster-export --seed $SEED`, no DB), so re-export it to
the ACTUAL mount path (`docker inspect demo-N-fake-fapi-1 … /roster/roster.json` — robust regardless of which
rext clone drives the run), then restart `demo-N-fake-{fapi,bapi}` (they read the file at startup), then a
bounded readiness poll on the FAPI. The readiness `curl` must be guarded (`|| echo 000`) — under the runner's
`set -e`, a curl SSL/connect failure (rc 35, the FAPI mid-restart) would abort the whole runner (the first
symptom: DONE-rc=35 with the log ending at "waiting for the fake-FAPI").

**D3 — Validation is belt-and-suspenders.** Beyond the suite's own `no-regressions` gate (10/11, 0 failing),
ran `ptvalidate --stack demo-1` (the datadna closure gate: every seeded verified-skill node-id resolves in the
replayed taxonomy — PASS) + the playthroughs Go tests (manifest/validator/report — PASS). The 1 unimplemented
(assign-and-track.UC1, the assign-WRITE half) is the declared in-manifest TODO (Fate-2), NOT a regression — the
`no-regressions` gate (nothing failing) is the coverage-milestone gate, and `AllGreen` is the foundation gate.
