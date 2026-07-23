# iter-20 — decisions

## D1 — the academy course-start proof MUST be a browser render-assert, never a curl-grep
When freshly re-confirming the M238 chapter body live, a raw `curl` of
`https://billion.taildc510.ts.net:13077/chapters/accessing-open-models/` returned HTTP 200 + a 386KB body +
the correct chapter `<title>How to Access Open Models — Literacy & Landscape — AI Academy</title>` — real
chapter content — yet `grep -c 'wandered off the trail'` (the 404 page's copy) hit **2**. That is a FALSE
POSITIVE: the "You wandered off the trail" string ships inside the Next.js error-boundary/RSC bundle on
EVERY academy page regardless of what renders, so a raw-HTML grep cannot distinguish a rendered chapter from
the 404 page. This is exactly why the shipped harness (`probeAcademyChapter` → `assertSection` on the
`main, body` region with a real browser + `mustNotInclude: ['You wandered off the trail']` on the RENDERED
text) is the authoritative fence. Ran the real browser sweep instead: fresh employee `run-coverage.sh 1
employee` against billion → **GATE MET**, cross-port follow "chapter body (slug accessing-open-models) +
?lang=it OK, HTTP 200." The gate-(h) academy course-start + academy language proofs are graded on THAT, not
the curl. (Recorded as the iter's Lesson + folds into the coverage-protocol discipline.)

## Evidence index (all live on billion, this tailnet peer, 2026-07-23)
- autoverify.json refreshed green, 0 warnings, ts 2026-07-23T00:34:22Z (STACK_DIR set, run as devops).
- p95<5s: employee 1.46s / manager 1.31s (run-latency.sh, https, LATENCY_GATE_MS=5000, fresh green-gate).
- academy: fresh coverage sweep GATE MET + probeAcademyChapter EN + ?lang=it HTTP 200.
- M241 EN/IT: content-manifest 21 bilingual tuples (EN10/IT13); cockpit :17700 renders "Content stories" +
  "EN/IT" + 28 `data-lang` cells.
- 0 platform edits; 0 rext code changes (pure live verification — harness run from the local authoring copy).
