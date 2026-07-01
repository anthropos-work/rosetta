# iter-04 — decisions

## D1 — Built + verified the next-web-public-website-url demopatch end-to-end on demo-1
The demopatch (iter-03 D3 design) was authored, tested (demopatch suite 47→49 incl a chained-hashes-reproduce-
against-live-clone guard; frontend-build 55 GREEN), committed to rext (`fix(M50/04)`), then BUILT INTO demo-1:
the new `demo-1-next-web` image baked all 3 demopatches via `up-injected.sh build_frontend_next_web` (driven
through the lib-only seam from the consumption clone). `NEXT_PUBLIC_PUBLIC_WEBSITE_URL=http://localhost:13000`
is confirmed baked into the server bundle. The consumption clone was left git-CLEAN (the trap reverted all 3
patches). The container was recreated with `--no-deps` (postgres + the 221-member seed intact). The chained
apply→revert was proven clean on the real demo-1 clone (LIFO: pubweb reverts before studio).

## D2 — The demopatch is CORRECT but the manager gate did NOT move: the residual escape is a DIFFERENT class (replayed CONTENT, not a JS constant)
**Re-sweep verdict (gate-valid, frontier exhausted at 68):** `failingSections=0 personaFailures=0
crossPortFailures=0 escapes=1 → GATE NOT MET` — escapes UNCHANGED (1→1).
**Root cause (definitive):** the residual escape on `/enterprise/activity-dashboard/ai-simulations/1bc8e23c`
is `https://anthropos.work/library/job-simulations/dev-ops-conflict-.../` — NOT built from `PUBLIC_WEBSITE_URL`
(the env IS baked: `localhost:13000` is in the bundle, and the OTHER 6 AD sim drill-downs the crawl visited
showed **0 eject**, proving the constant-class links are now demo-local). The link is the
**`directus.simulations.public_landing_page_url`** field (+ `read_more_link`) — a hardcoded
`https://anthropos.work/library/job-simulations/<slug>/` value **REPLAYED FROM PROD** by the snapshot. The
activity-dashboard sim drill-down renders this content field as a link → the prod-eject. **28 sims carry a prod
`public_landing_page_url`, 14 a prod `read_more_link`** (the full content-escape class; the AD-drilldown sample
surfaced one — `1bc8e23c` — but the rest are the same class, sampled-out).
**Why the demopatch still matters:** it correctly fixes the JS-constant-built links (the studio link + any
`PUBLIC_WEBSITE_URL`-built in-app link) — a necessary, landed, tested fix. It is NOT reverted; it stays.

## D3 — Fix surface for the residual: a stack-snapshot content rewrite (replayed Directus URL fields), NOT a demopatch (→ iter-05)
**Routing:** this is the routing table's **"Federation / content"** class — the escape lives in REPLAYED
CONTENT (a Directus field), so the fix is in `stack-snapshot` (the replay layer), not next-web. Two options:
**(a)** a `stack-snapshot` content-rewrite step that rewrites `public_landing_page_url` / `read_more_link`
(+ any other `anthropos.work`-bearing content field) from the prod host to the demo-local host during replay
(the snapshot already does serve-grants + content transforms — add a URL rewrite); **(b)** a demo-local
post-replay `UPDATE` on `directus.simulations` (an idempotent rext-owned `up-injected.sh` step, like the FK
indexes / the Directus column backfill — demo-local DDL, the `cms`/Directus clones stay pristine).
**Choice (routed to iter-05):** option (a) — the `stack-snapshot` rewrite — because it's the snapshot layer's
job to make replayed content demo-local (the same firewall/transform discipline), it covers ALL content URL
fields uniformly, and it's reproducible on a fresh `/demo-up` (the cold M53 gate). Verify with another manager
re-sweep → escapes 0 → manager GATE MET → (with the already-met employee gate) the full M50 gate met on warm
demo-1 (the COLD reset proof remains M53).

## D4 — iter-04 closes closed-fixed-partial
The planned deliverable (the demopatch) LANDED + is verified-correct for its class (cleared the constant-built
ejects) + committed + built into demo-1 + the container recreated (seed intact, clone clean). But the planned
OUTCOME (escapes 1→0 → gate met) was NOT reached because the re-sweep revealed the residual escape is a
DIFFERENT, previously-conflated class (replayed content URL fields). Real, correct work landed (not a revert →
not closed-no-lift), but the gate metric didn't move (→ not closed-fixed). The residual (the content-field
escape) routes to iter-05 with a named fix surface (stack-snapshot rewrite). The member-field fill (iter-02)
also still rides forward for its gate-verification + the D4/F1 manifest-strengthening.
