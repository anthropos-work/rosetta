# M50 — progress

## Running ledger
_Appended after each iter (tik/tok). Each entry: what was attempted, what moved, gate distance._

- iter-01 (tok·bootstrap): authored TOK-01 + FRESH-demo-1 re-diagnosis (4 genuine seed-gap clusters; has-data surfaces flagged for sweep) — see iter-01/progress.md
- iter-02 (tik, closed-fixed-partial): EMPLOYEE gate MET on baseline (valid, frontier-exhausted, 59 pages, all 0) → employee half needs no fix; member-field fill (memberships join-date/location/last-activity) authored+tested+committed to rext (`fix(M50/02)`); MANAGER baseline (prematurely) read as frontier-capped → iter-03 — see iter-02/progress.md
- iter-03 (tik, closed-fixed): re-survey CORRECTED iter-02 — the manager sweep EXHAUSTS, not caps (cap=300 → reachable=68, cappedAtFrontier=false, gate-VALID); tooling-iter CANCELLED. Manager verdict: failingSections=0 personaFailures=0 crossPortFollowFails=0 escapes=1 → the SOLE blocker is a prod-eject to `anthropos.work` on the activity-dashboard sim drill-down (then-attributed to hardcoded `PUBLIC_WEBSITE_URL`). Content fully populated. — see iter-03/progress.md
- iter-04 (tik, closed-fixed-partial): built+tested+verified the `next-web-public-website-url` demopatch on demo-1 (cleared the JS-constant-built ejects — 6 AD drill-downs clean) BUT the re-sweep shows escapes still 1: the residual is a DIFFERENT class — replayed Directus content `public_landing_page_url`/`read_more_link` carrying a prod `anthropos.work` URL (28/14 sims). Routes to iter-05 (stack-snapshot content rewrite). — see iter-04/progress.md
- iter-05 (tik, closed-fixed, **GATE: MET**): post-replay Directus content-URL rewrite (simulations + skill_paths, regex over any anthropos.work subdomain → demo host) + cms cache clear killed the residual escape. **Manager re-sweep FINAL: reachable=69 failingSections=0 escapes=0 personaFailures=0 crossPortFailures=0 frontier=EXHAUSTED gateMet=True.** With employee (iter-02) → **M42 gate GREEN BOTH vantages on warm demo-1.** — see iter-05/progress.md

## GATE STATUS
**M42 coverage gate MET on BOTH vantages on the WARM demo-1** (employee iter-02 + manager iter-05; both frontier-exhausted, (failingSections,escapes)=(0,0), 0 persona, 0 cross-port). The 3 M50 fixes are reproducibly baked into the bring-up tooling. **The explicit milestone exit_gate ("on a COLD reset-to-seed demo") is the remaining acceptance step** — reserved for the heavy COLD pass (close/harden + the v1.10b M53 cold-rebuild milestone).

## Next-iter queue (routes carried forward — to close/harden + M53, NOT blocking the warm metric)
- **COLD reset-to-seed acceptance** (the explicit exit_gate): fresh `/demo-up` (all 3 fixes reproduce from tooling) + both-vantage sweeps → confirm (0,0) both vantages on COLD. The v1.10b M53 "cold-rebuild acceptance" milestone owns this; close-milestone runs it or surfaces to the orchestrator.
- Manager manifest-strengthening (D4/F1): assert the `/enterprise/members` Location column (+ workforce tab contents) so the member-field fill (iter-02) is gate-PROVEN (renders, but the manifest doesn't yet assert it).
- Languages seeder + cert roster-coverage: ONLY if the strengthened manifest surfaces them failing (the current sweep passes existing assertions — confirm first; do NOT speculatively seed).
- AI-provider-keys policy (F7) + academy menu-link/non-anonymous-session (F6): decision deliverables → secrets-spec.md; academy AI chat documented-as-absent (not a gate blocker) — for close/M51.
- Re-pin the consumption clone (`stack-demo/rosetta-extensions`) to the `fit-up-m50` tag at close (it carries the iter-04/05 fixes synced for live verification).
