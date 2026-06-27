---
iteration_type: tik
status: closed-fixed
---

# iter-04 — reconcile the manager route + populate the dashboard + sample the fan-outs + calibrate (TOK-01 lines 2-4)

**Active strategy reference:** TOK-01 (the bootstrap strategy — reconcile-route + clear-escape +
populate-dashboard + exhaust-frontier). This tik executes **lines 2-4** (line 1, the Studio escape, was
resolved by iter-03's demo-patch tool — escapes already 0).

**Cluster / target identified:** the iter-03 close routed forward TOK-01 lines 2-4 as the manager-gate
residual: `notReached=5` (the manifest still pointing at the wrong `/workforce/*` routes) +
`frontier=CAPPED` (the two manager fan-outs with no sample rules). Persona PASSes + escapes=0 (run 1/2). So
the residual is: reconcile the route, prove the dashboard renders real content, sample the fan-outs so the
frontier exhausts, and calibrate.

**Hypothesis:** (line 2) the notReached=5 is a route-model error, not a content gap — re-author `MANAGER_PAGES`
to the real `/enterprise/*` routes (the dashboard is ONE tabbed route) and the M36 dashboard already renders
real data (only re-confirm + seed what's proven empty). (line 3) adding the two manager fan-out sample rules
makes the frontier exhaust. (line 4) calibrate the manifest floors against the live render.

**Expected lift:** notReached 5→0, frontier CAPPED→EXHAUSTED, failingSections stays 0 (the dashboard renders);
manager gate → MET (escapes already 0, persona already PASS).

**Phase plan:** the coverage-protocol A–E loop. Phase A = the live diagnostic probes (route discovery) + the
baseline. Phase B = triage (route-model error + the one empty page). Phase C = the fixes (manifest reconcile +
the FeedbackSeeder mirror fix + sample rules) + re-seed demo-3. Phase D = the authoritative sweep. Phase E =
close + grade against the gate.

**Escalation conditions:** a dashboard surface that is empty for a platform-only reason (parallel to M42e's
sim-start) → seed the entitlement (reproducible) or, if platform-bound, a re-scope trigger. (None hit — the
one empty page was a seed-mirror gap, fixed in `stack-seeding`.)

**Acceptable close-no-lift outcomes:** if a dashboard surface proves platform-bound-empty (a re-scope trigger),
that falsification satisfies the protocol even without the gate moving.
