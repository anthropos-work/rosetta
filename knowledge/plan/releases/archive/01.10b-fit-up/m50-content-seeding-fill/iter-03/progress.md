# iter-03 — progress

**Type:** tik (under TOK-01) — Phase A (measure) + B (triage); recovers the manager baseline iter-02 missed.

## Work
- **Step 0 re-survey** corrected iter-02's "manager sweep will cap" assumption (inspected `crawl.ts`: sampling
  caps VISITED pages, sampled-out paths only inflate the queue + drain fast). Predicted the manager sweep
  EXHAUSTS.
- **Phase A (manager sweep to completion, `COVERAGE_MAX_PAGES=300`).** Result: **gate-VALID** —
  `reachable=68 cappedAtFrontier=false frontierRemaining=0 frontier=EXHAUSTED`. The re-survey was confirmed; no
  SAMPLE_RULES tooling fix needed (the iter-02-routed tooling-iter is cancelled). The 60s cold-grid warm
  (M46 perf signature) is a one-time warm cost at 221 members; noted for the M53 cold gate.
- **Phase B (triage).** Verdict: `failingSections=0 personaFailures=0 crossPortFollowFails=0 escapes=1 →
  GATE NOT MET`. The manager CONTENT is fully populated; the SOLE blocker is **escapes=1**: the
  `/enterprise/activity-dashboard/ai-simulations/<simId>` drill-down links to prod
  `https://anthropos.work/library/job-simulations/<slug>/`. Traced to the hardcoded
  `PUBLIC_WEBSITE_URL` constant (`packages/core-js/src/constants/urls.ts:1`, no env override) → the
  "platform-bound escape" routing class → a NEW demopatch (iter-04). Visual confirm earlier: the
  `/enterprise/members` Location column renders "–" for every row (the member-field gap my iter-02 fix fills),
  but the current manifest doesn't assert it (D4/F1).

## Close — 2026-06-30

**Outcome:** Recovered the gate-VALID manager baseline (frontier-exhausted at 68 pages, the measurement iter-02
prematurely stopped) — the manager gate has exactly ONE blocker (a prod-eject to `anthropos.work` on the
activity-dashboard sim drill-down); CONTENT is fully populated (failingSections=0, personaFailures=0,
cross-port studio OK). Triaged the escape to a new `next-web-public-website-url` demopatch (iter-04). The
iter-02 "will cap → tooling-iter" inference is cancelled (no SAMPLE_RULES fix needed).
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (employee MET+valid; manager 0 failing-sections / 0 persona / 1 escape — the single escape is the blocker)
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (re-survey right — sweep exhausts, tooling-iter cancelled), D2 (single escape blocker; content populated), D3 (fix surface = new demopatch), D4 (closed-fixed on the recovered baseline + triage).
**Side-deliverables (if any):** none (measurement + triage tik; 0 code changed).
**Routes carried forward:**
- **iter-04 (the fix):** author `next-web-public-website-url` demopatch (patch the hardcoded `PUBLIC_WEBSITE_URL` to read an env override, behavior-identical fallback, 6-guard contract) + wire into `up-injected.sh` + set the offset `.env.local` value + re-build the frontend + re-sweep manager → escapes=0 → manager GATE MET.
- Manager manifest-strengthening (D4/F1): assert the `/enterprise/members` Location column (+ join/last-activity behind the Columns toggle, + the workforce Growth/Verification/Talent tab contents) so the member-field fill + the annotation gaps are gate-PROVEN — pair with the re-seed of the member-field fix.
- Languages seeder + cert roster-coverage: only if the strengthened manifest surfaces them as failing (the current sweep shows the content passes the existing assertions — confirm against the strengthened manifest first; do NOT speculatively seed).
- AI-keys policy (F7) + academy (F6): decision deliverables, future tiks.
**Lessons:** (1) NEVER stop a sweep early on a "will cap" assumption — re-survey the crawl's sampling math first. The cap counts VISITED pages; a huge queue (q) of sampled-out template-identical paths drains fast and does NOT cap. iter-02 burned a premature stop on this; iter-03's re-survey + a cap=300 re-run recovered the valid baseline. (2) The manager gate's real shape on the fresh demo-1: content is fully populated (0 failing sections) — the only gap is a single prod-eject link, exactly the kind the gate exists to catch. The annotation's perceived "empty" manager surfaces were the stale-demo artifact (like the employee gaps). (3) The member-field believability fill is real but the CURRENT manifest doesn't assert it — a fill that passes the gate without being measured needs a manifest-strengthening to actually be PROVEN closed.
