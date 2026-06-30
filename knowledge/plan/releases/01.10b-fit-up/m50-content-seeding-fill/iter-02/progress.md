# iter-02 — progress

**Type:** tik (under TOK-01) — per `coverage-protocol.md` Phase A–E.

## Work
- **Phase A (baseline sweeps).** Ran the M42 coverage harness against demo-1 as both vantages:
  - **Employee (Maya): GATE MET** — `reachable=59 cappedAtFrontier=false (frontier EXHAUSTED) failingSections=0
    escapes=0 personaFailures=0 notReachedPages=0 gateMet=true` (Playwright 1 passed, 12.0m). A gate-valid
    measurement. The annotation's employee gaps were all already closed by M47/M48 + the seed.
  - **Manager (Dan): NOT MEASURABLE — frontier-cap.** All `/enterprise/*` pages + /profile render 200, 0 FAIL
    streamed, but the BFS frontier explodes past the 150-page cap on the per-member `/user/<uuid>` team-roster
    fan-out (221 Cervato members) + per-sim/per-skill-path `/enterprise/activity-dashboard/.../<uuid>`
    drill-downs (q observed at 172 by page 39). A capped sweep is gate-INVALID per the protocol's
    measurement convention → the manager gate can't be read until the SAMPLE_RULES collapse those
    template-identical fan-outs (routed to iter-03 tooling-iter). Sweep stopped (no value in a capped run +
    machine-aware).
- **Phase C (fix — member-field cluster).** The leading manager content gap (DB-confirmed: memberships
  joined_at / location / last_activity_date 0/221) is authored + unit-tested in rext `users.go` (the COPY +
  a NULL-only idempotent re-seed backfill). Full stack-seeding suite GREEN incl the M17 re-seed-idempotency
  contract. Committed to rext as `fix(M50/02)`.
- **Phase D (re-measure).** The member-field fix's gate effect on the manager vantage is NOT yet measurable
  (the manager sweep caps) — routed to iter-03 after the tooling fix. The employee re-measure is unchanged
  (gate already met; the fix is manager-vantage only).

## Close — 2026-06-30

**Outcome:** Employee gate MET on baseline (valid, frontier-exhausted) — the employee half of the M50 gate
needs no fix. Member-field fill (memberships join-date/location/last-activity) authored + unit-tested + committed
to rext. Manager baseline is blocked by a frontier-cap tooling gap (the `/user` + activity-dashboard drill-down
explosion) → routed to iter-03 as a tooling-iter; the member-field fix's manager-gate verification rides on it.
**Type:** tik
**Status:** closed-fixed-partial
**Gate:** NOT MET (employee half MET + valid; manager half unmeasured pending the iter-03 tooling fix)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (member-field targets memberships cols), D2 (NULL-only idempotency guard), D3 (close-partial; manager baseline → iter-03 tooling-iter), D4 (F1 manifest-strengthening reconciliation).
**Side-deliverables (if any):** none.
**Routes carried forward:**
- **iter-03 (tooling-iter):** tighten the manager SAMPLE_RULES so the `/user/<uuid>` + `/enterprise/activity-dashboard/.../<uuid>` template-identical fan-outs collapse to a representative+boundary sample → the manager frontier exhausts under the cap → a gate-valid manager baseline. THEN re-seed demo-1 + manager sweep to verify the member-field fix's gate delta.
- Manager manifest-strengthening (D4/F1): assert the `/enterprise/members` location/join/last-activity columns + the workforce tab contents (Growth/Verification/Talent) the annotation flagged — so the gate proves the gaps are closed (a future tik, paired with the languages + cert-coverage fills).
- Languages seeder (new `MemberLanguagesSeeder` + `world_languages` reference fill) + cert roster-coverage (the Talent-tab gaps) — future tiks, after the manager gate is measurable.
- AI-provider-keys policy (F7) + academy wiring (F6) — future tiks (decision deliverables).
**Lessons:** (1) The FRESH-demo re-diagnosis paid off massively — the employee vantage was ALREADY at the gate, so half the annotation's gaps were stale (observed on the pre-M47/M48 demo). Always re-diagnose before fixing. (2) The manager vantage's reachable set on a realistic-size org (221 members) explodes the crawl frontier via per-entity drill-downs; the SAMPLE_RULES calibrated against demo-3 don't collapse it enough — the manager gate isn't measurable without a sampling fix. A capped sweep is gate-invalid; never quote it. (3) A backfill of `time.Now()`-relative columns must use a NULL-only guard, not IS-DISTINCT-FROM, or it breaks re-seed idempotency.
