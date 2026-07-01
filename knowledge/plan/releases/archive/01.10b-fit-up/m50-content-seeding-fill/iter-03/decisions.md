# iter-03 — decisions

## D1 — The re-survey was RIGHT: the manager sweep EXHAUSTS, it does not cap (iter-02's "will cap" was wrong)
**Context:** iter-02 stopped the manager sweep at page 39 (q=172) assuming it would hit the 150-page cap.
iter-03's re-survey of `crawl.ts` predicted the opposite (sampling caps VISITED pages, sampled-out paths only
inflate the queue, drained fast at dequeue).
**Result:** the cap=300 manager sweep **EXHAUSTED at reachable=68, cappedAtFrontier=false, frontierRemaining=0**
— a fully gate-VALID measurement. The sampling worked exactly as designed (sampledOut: 28 `/user`, 65 AD
drill-downs, 63 sims, 11 skill-paths). The cap=150 was simply too low for the manager's reachable set (≈68
visited + the sampled-out families); cap=300 exhausts. The iter-02 premature stop denied a valid baseline; no
SAMPLE_RULES tooling fix is needed (the iter-02-routed tooling-iter is CANCELLED).
**Choice:** the manager baseline is read at `COVERAGE_MAX_PAGES=300` (a safe margin above the ~68 exhaustion
point). Document this as the manager-sweep cap going forward.

## D2 — The manager gate has exactly ONE blocker: a prod-eject escape (content is fully populated)
**Manager baseline verdict (gate-valid):** `failingSections=0 personaFailures=0 crossPortFollowFails=0
notReachedPages=0 escapes=1 → GATE NOT MET`. So:
- **0 failing sections + 0 persona failures** — the manager CONTENT is fully populated (the member-field gap,
  languages, certs all pass the CURRENT manifest — confirming D4/F1: the manifest doesn't yet ASSERT those
  columns, so they don't fail; the believability fill is real but unmeasured by the gate).
- **cross-port studio follow OK** — the M49 #8 studio-url demopatch works on demo-1 (studio-desk home renders
  for Dan at `:19000`, HTTP 200, marker present — no eject, no login-loop).
- **The SOLE blocker:** `escapes=1` — the `/enterprise/activity-dashboard/ai-simulations/<simId>` drill-down
  renders a link to **prod** `https://anthropos.work/library/job-simulations/<slug>/` (the public sim page).

## D3 — Fix surface for the escape: a NEW demopatch (PUBLIC_WEBSITE_URL is a hardcoded constant, no env override)
**Context:** the escape host comes from `PUBLIC_WEBSITE_URL = 'https://anthropos.work'` —
`stack-demo/next-web-app/packages/core-js/src/constants/urls.ts:1`, a **hardcoded constant** with NO
`NEXT_PUBLIC_<thing>_URL` env override. The activity-dashboard sim drill-down combines it with the public
job-simulations route.
**Routing:** this is the routing table's **"Platform-bound escape (no per-URL override)"** class → the env-rewrite
row does NOT apply → the **demopatch** tool (source-patch the demo's EPHEMERAL clone to read an env var with a
behavior-identical fallback, trap-revert; the 6-guard demopatch contract). Well-precedented: the existing
`next-web-studio-url`, `next-web-members-pagination`, `app-targetrole-authz-skip` patches.
**Choice (routed to iter-04):** author a new `next-web-public-website-url` demopatch (patch `urls.ts`'s
`PUBLIC_WEBSITE_URL` to read `NEXT_PUBLIC_PUBLIC_WEBSITE_URL` with the prod value as fallback), wire it into
`up-injected.sh`'s inject loop, set the offset value in the `.env.local` overlay (point it at the demo's own
next-web offset so the public sim link stays demo-local), re-build the frontend, re-sweep manager → escapes=0 →
manager GATE MET. NOTE: this patch also fixes the OTHER `PUBLIC_WEBSITE_URL` links (the hiring/workforce-intel
modal links in `packages/ui`) that the crawl's allow-rule / sampling didn't surface but are the same class.

## D4 — iter-03 closes closed-fixed (planned scope = recover the gate-valid manager baseline + triage; both landed)
The iter's planned deliverable was the gate-valid manager baseline (the measurement iter-02 missed) + triage of
the real failing sections — NOT a fix landing (the overview made the fix conditional). Both landed: a valid
frontier-exhausted manager baseline + the single escape fully triaged to its fix surface. The fix (a new
demopatch + frontend rebuild + re-sweep) is iter-04. The member-field fix (committed iter-02) remains landed; its
gate-verification + the D4/F1 manifest-strengthening also ride forward (once the escape is fixed + the manifest
asserts the member columns, the member-field fill becomes gate-measured).
