# iter-05 — decisions

## D1 — The residual escape was REPLAYED CONTENT (Directus URL fields) — fixed by a post-replay content-URL rewrite
iter-04 proved the escape is `directus.simulations.public_landing_page_url` (+ `read_more_link`) carrying a
hardcoded prod `https://anthropos.work/library/job-simulations/<slug>` replayed verbatim. iter-05 adds a
post-replay content-URL rewrite step in `up-injected.sh`'s NO_SETDRESS block (the content-side analog of the
injection link-rewriting for app constants): an idempotent demo-local `UPDATE` that `regexp_replace`s any
`https?://<subdomain>anthropos.work` host → the demo's own next-web host (3000+OFFSET). Same mechanism class as
the M46 FK indexes / Directus column backfill (demo-local DDL on the per-stack Directus, idempotent — a re-run
matches 0 rows, non-fatal M18/M19, gated on local content, `DEMO_NO_CONTENT_URL_REWRITE` opt-out).

## D2 — Broadened twice during the iter (planned-scope, same escape class): the staging host + the skill_paths root
- **Staging host:** a bare `LIKE 'https://anthropos.work%'` missed 5 sims carrying `https://staging.anthropos.work`
  (also an off-demo eject). Switched to a REGEX `https?://[a-z0-9.-]*anthropos\.work` to catch prod + staging +
  www + http variants uniformly.
- **The skill_paths content root:** `directus.skill_paths` ALSO carries a prod `public_landing_page_url`
  (`/library/skill-paths/<slug>`, 1 row) the `/skill-path/<slug>` Library page would render — extended the
  rewrite to cover it (skill_paths has NO `read_more_link` column — verified against the live schema, so only
  `public_landing_page_url` there; `simulations` keeps both). These were the same escape class surfacing across
  both content roots; broadening within the iter is planned-scope (the escape-elimination is the deliverable).

## D3 — The cms Redis DB-5 cache is the M46 poison class — clear it after the rewrite (load-bearing)
The cms caches `GetJobSimulation` per-id in Redis DB 5 (24h TTL, cache-first). After the content rewrite, a
stale (prod-URL) cached entry can serve the OLD URL → the eject persists despite the DB being demo-local. The
first iter-05 re-sweep's mid-finalization read showed `escapes=1` from exactly this stale cache (the `1bc8e23c`
page). Clearing DB-5 `simulations_*` after the rewrite (the M46-documented step) made the FINAL verdict clean.
**Lesson:** a Directus content change is only live once the cms sim cache is cleared — pair the rewrite with a
cache flush. (A fresh COLD `/demo-up` has no poison — the replay precedes cms's first query — so this only bites
an in-place re-replay/rewrite, which is the iterative-measurement path.)

## D4 — iter-05 closes closed-fixed; the WARM both-vantage M42 gate is MET (the COLD reset proof is M53)
**Manager re-sweep FINAL verdict (gate-valid, frontier exhausted at 69):** `failingSections=0 escapes=0
personaFailures=0 crossPortFailures=0 notReachedPages=0 gateMet=True`. Combined with the employee gate
(iter-02, gateMet=True), **the M42 coverage gate is GREEN on BOTH vantages on the warm demo-1.** All three M50
fixes are reproducibly baked into the bring-up tooling (member-field seeder backfill `fix(M50/02)`,
next-web-public-website-url demopatch `fix(M50/04)`, content-URL rewrite `fix(M50/05)`) — so a COLD `/demo-up`
reproduces them.
**The milestone exit_gate is specifically "on a COLD reset-to-seed demo."** The warm both-vantage gate-met is
the protocol's primary metric `(failingSections, escapes)=(0,0)` reached on both vantages — a major milestone —
but the COLD reset-to-seed proof is the explicit exit-gate requirement, reserved for the heavy acceptance pass
(per the orchestration's machine-aware constraint + the v1.10b M53 "cold-rebuild acceptance" milestone). So
iter-05 closes `closed-fixed` with **Gate: MET (warm, both vantages)**; the COLD reset-to-seed acceptance is the
remaining exit-gate step — routed to the close/harden + M53. The member-field fill's gate-PROOF (the D4/F1
manifest-strengthening to assert the Location column) is the one carried-forward believability item: the data
renders (visually confirmed) but the current manifest doesn't ASSERT it, so it passes without being measured.
