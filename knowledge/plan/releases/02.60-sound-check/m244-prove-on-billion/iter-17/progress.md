**Type:** tik (run 6, tik 5 — the run's final tik). Active strategy: TOK-02.

# iter-17 — progress

## Discrete stack-verify specs — driveable run + landscape map
Assessed each discrete spec's host/env contract, then ran the remote-driveable ones against billion (demo seed, no reset):

- **persona-check** (`COVERAGE_APP_BASE_URL`/`COVERAGE_FAPI_BASE_URL`/`COVERAGE_IDENTITY_KEY`/`COVERAGE_EXPECTED_ORG`) — **PASS ✅** on billion (employee maya-thriving): role-skills coherence PASS ("Maya Chen DevOps Engineer", 0 junk-pool tokens), avatar-consistency PASS (real photo, menu==profile), org-identity PASS ("Cervato Systems" + real logo).
- **m224-candidate-heroes** (`CANDIDATE_HOST`/`CANDIDATE_APP_SCHEME`/`CANDIDATE_OFFSET`, the hiring app :13001) — **PASS ✅** on billion (3/3): Cody (candidate) home renders real hiring assignments; Rae (recruiter) enterprise activity-dashboard renders the results scoreboard (5 sims, real scores/attempts); recruiter Results comparison reachable (no regression).

## Landscape map (the finding — the gate-(c) discrete half needs a REMOTE-CAPABILITY RETROFIT)
Of the ~13 discrete stack-verify specs, only THREE carry a host env (persona-check ✓, m224-candidate-heroes ✓, probe-navigation [a skip-unless-`PROBE_PATH` diagnostic, not a gate spec]). The rest are NOT remote-capable as written:
- **offset-only** (read `DEMO_OFFSET` for the port but hardcode localhost scheme/host): `talk-to-data-m239`, `enterprise-surfaces-m239` — need a host+scheme var to drive against billion's tailnet HTTPS.
- **no host/base env** (localhost-only): `smoke`, `verify-members-B`, `verify-activity-dashboard-servegrant`, `render-hiring-comparison`, `cockpit-overlay-return`, `m220-session-and-egress`, `calibrate` / `calibrate-avatar` / `calibrate-manager` / `calibrate-talent`, `probe-empty`.

⇒ Running the FULL discrete-spec half green on billion is blocked on retrofitting the `COVERAGE_HOST`/`RENDER_HOST`/`CANDIDATE_HOST`-style host+scheme pattern into these localhost-only specs (the same remote-capability retrofit M219 did for `run-coverage.sh` + `run-playthroughs.sh` + m224). A rext tooling effort — routed to run 7. This is why the coverage sweep + these 3 host-capable specs were the only stack-verify live-browser proofs that ran remotely before now.

## Gate (c) status
Coverage sweep GREEN (iter-16) + persona-check + m224 green (iter-17). Gate (c) as a whole ("40 = 24 stack-verify + 16 Playthroughs") NOT ticked: the localhost-only discrete specs need the remote retrofit, and the 16 Playthroughs run LAST (pt-world reset). Metric stays **5/8**.

## Close — 2026-07-23

**Outcome:** 2 remote-driveable discrete stack-verify specs (persona-check + m224-candidate-heroes) PROVEN green on billion; mapped the remaining ~11 as needing a remote-capability retrofit (host+scheme env) before they can run against the tailnet. Gate (c) not ticked. Metric stays **5/8**.
**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET (gate part (c) — coverage + 2 discrete specs green; the localhost-only discrete specs + Playthroughs remain)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: **y** (tik 5/5) — (6) protocol-stop: n — Outcome: **exit-5 (cap-reached)**
**Decisions:** D1 (iter-17/decisions.md).
**Side-deliverables:** none.
**Routes carried forward (run 7):**
- **gate (c) — the remote-capability retrofit:** add a host+scheme env (COVERAGE_HOST-style) to the localhost-only discrete specs (calibrate*/verify-*/talk-to-data-m239/enterprise-surfaces-m239/render-hiring-comparison/cockpit-overlay-return/m220/smoke) so they drive against billion; run them green. THEN the 16 Playthroughs LAST (pt-world reset).
- gate (f) 3 drift-carries · (h) v2.6 fixes + p95<5s · DEF-M239-01 — demo-seed gates.
**Lessons:** "the specs execute green on billion" silently assumes the specs CAN target billion. Most stack-verify specs were authored localhost-only; only the ones M219 retrofitted (coverage/playthroughs/m224) + persona-check carry a host var. The gate-(c) count (40) is real, but a third of it needs a remote-capability retrofit before it can even be attempted remotely — a tooling precondition the gate wording hides.
