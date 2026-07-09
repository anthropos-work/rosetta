**Type:** tik (v2.0 Playthroughs suite GREEN on merged demo-1). Under TOK-01 move (4).

# iter-16 — tik progress

## Execution log
1. **ptvalidate static PASS:** 5 products, 11 use cases, 10 live Playthroughs discovered, 1 TODO.
2. **First full run FAILED all 10** with ONE root cause: `cockpit-login: POST /v1/demo/select(pt-*) -> 400
   {"code":"unknown_identity"}`. The reset-to-seed swapped the DB world to `pt-world`, but the fake-FAPI/BAPI
   roster (`/roster/roster.json`, mounted) still held the STORIES heroes baked at bring-up (demo-1 was brought
   up for coverage). M204 masked this by bringing its demo up `pt-world`-native (per M204 iter-01: "the
   pt-manager seat … is already wired in demo-1's roster export").
3. **Proved the fix manually:** regenerated the roster from `pt-world` (`stackseed --roster-export --seed
   pt-world` → 3 seats pt-employee/pt-manager/pt-free) + restarted the fake services → the proof-of-life
   Playthrough (pt-profile-identity) PASSED (✓ hero logs in, sees her identity).
4. **Durable Fate-1 fix:** taught `run-playthroughs.sh --reset` to also refresh the roster (docker-inspect the
   mount path → `stackseed --roster-export --seed $SEED --roster-out <mount>`) + restart `demo-N-fake-{fapi,bapi}`
   + bounded FAPI-readiness wait — so reset-to-seed resets the WHOLE world (DB + identities) on ANY demo. Fixed a
   `set -e` curl-abort in the readiness loop (guard `|| echo 000`). bash -n + shellcheck clean. rext `e822c70`
   (tag `quick-change-m211` moved) + consumption clone re-synced.
5. **Full suite re-run (--reset, full lifecycle):** **68 passed (5.0m)** → **Playthroughs coverage 10/11
   passing (passing=10, failing=0, unimplemented=1, unimplementable=0)** — all 6 employee + 4 manager
   Playthroughs GREEN; the 1 unimplemented is the declared in-manifest TODO (assign-WRITE half → Fate-2). Gate
   `no-regressions` MET.
6. **ptvalidate closure gate vs demo-1 PASS:** "every seeded verified-skill node-id resolves in the replayed
   taxonomy" (datadna built + on PATH). playthroughs Go tests (manifest/validator/report/ptreport/ptvalidate)
   all GREEN.

## Re-measurement (gate sub-conditions)
| Sub-condition | Pre-iter | Post-iter |
|---|---|---|
| (e) M42 coverage (both vantages) | GREEN (iter-14/15) | GREEN |
| (e) **v2.0 Playthroughs** | NOT MET | **GREEN** (10/11, 0 failing, 1 declared TODO; closure PASS) |
**Metric:** Playthroughs gate `no-regressions` **MET** (10 live GREEN on cold reset-to-seed). **Sub-condition
(e) is now COMPLETE** (coverage both vantages + Playthroughs). Remaining composite-gate piece: the cold `/dev-up`.

## Close — 2026-07-08

**Outcome:** The v2.0 Playthroughs suite is GREEN on the merged demo-1 — 10/11 passing (0 failing, 1 declared
TODO), closure gate PASS. Root-caused the all-10-fail to a reset-to-seed gap (DB swapped, roster not) and
landed the durable fix (reset-to-seed now refreshes the Clerkenstein roster + restarts the fake services, so
Playthroughs run on ANY demo). Sub-condition (e) COMPLETE.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (sub-condition (e) complete, but the composite M211 gate still needs the cold `/dev-up`)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (tik with +progress) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n (3 tiks) — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (the all-10-fail is a reset-to-seed WORLD-completeness gap — DB reset but not the fake-FAPI/BAPI roster; M204 masked it with a pt-world-native demo; the fix belongs in `run-playthroughs.sh --reset`, not a per-run hack), D2 (roster refresh discovers the mount path via `docker inspect` so it works from any rext clone + restarts both fake services + waits for FAPI readiness; guard the readiness curl against set-e)
**Side-deliverables:** none (the runner fix is the load-bearing deliverable enabling the gate).
**Routes carried forward (Fate-3 → final iter):** cold `/dev-up` (the DEV half of the cold headline — verify extensions-bootstrap + casbin load before migrate).
**Lessons:** "reset-to-seed" must reset the WHOLE world the actor observes — DB **and** the identity roster the cockpit seat-switch resolves. A DB-only reset is silently incomplete on a demo not brought up seed-native (the login 400s before any assertion runs). This is the Playthroughs analog of the M42e "green-but-wrong" additive-reseed trap, one layer up (identities, not rows). Documented in playthroughs.md's reset-to-seed section.
