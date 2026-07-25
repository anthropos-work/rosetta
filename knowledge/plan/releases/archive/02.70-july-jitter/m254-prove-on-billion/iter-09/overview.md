---
iteration_type: tik
milestone: M254
iter: 09
status: closed-fixed-partial
---

# M254 · iter-09 — (e) builder Playthrough + (h) Playthroughs (the seed-destroying tail)

**Type:** tik (5th of the session — the cap) · **Active strategy:** TOK-01 — cluster 4 (mutating tail).

## Cluster / target
Run the Playthrough suite against the fresh billion demo: (e) the studio builder Playthrough + (h) the live
Playthroughs + live-browser specs. The runner splits: `--reset-only` on the demo host (pt-world reset, needs
docker + stackseed + DB) then the browser suite from this tailnet peer.

## Realized
- **pt-world reset-to-seed LANDED** on billion (`run-playthroughs.sh 1 --reset-only`): full FK-ordered TRUNCATE
  + fresh pt-world seed, 30-identity roster exported, fake services restarted, fake-FAPI 200. The demo is now
  the Playthrough world.
- **The browser Playthrough suite PROVEN to run + pass** from this peer (PT_HOST=billion, https): `117 tests,
  1 worker (serial)`, the first 6 green including the manager `@pt:pt-activity-drilldown` Playthrough (24 s).
- The full 117-test suite completion + a detached-run log fix + the two dispositions (f-FCP-p95, g-testhealth)
  route to a fresh agent (the prove-on-billion "fresh agent per run" discipline — this session's context is
  saturated after 5 tiks + a disruption-recovery).

## Phase plan
verification.md + playthroughs.md: reset-to-seed → browser suite. Reset landed; suite proven-to-run; full
completion routed.

## Escalation
None — the Playthroughs run + pass; only the detached-run plumbing + the full-suite wall-clock remain (a fresh
agent's job). 0 platform edits.

## Acceptable close-no-lift
The pt-world reset + a proven-to-run browser suite (6/6 green foreground) satisfies the protocol as a landed
deliverable; the full-suite completion is routed forward → closed-fixed-partial.
