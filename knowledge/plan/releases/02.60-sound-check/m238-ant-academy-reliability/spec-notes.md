# M238 — Spec notes

Topic → doc → code triples + chapter-body / language-path findings accumulate here during build.

## Chapter-body demo path (#3 Start→404)
- Bodies are backend-authoritative, no FS fallback; the catalog demopatch covers only the catalog.
- Option: a chapter-body FS-fallback demopatch (analogous to `academy-fs-published-fallback`) vs wiring the academy backend for the demo.

## Language error (#2)
- Re-triaged in M237; likely the same backend-null path as #3.

## Academy presence/coverage sweep
- Extend the sweep to assert chapter-body render (not just catalog cards).

## LIVE VERIFY on billion (demo-1, 2026-07-21) — #3 + #2 PROVEN
Baseline reproduced + fix proven on the live billion demo-1 academy (`:13077`, tailnet
`https://billion.taildc510.ts.net:13077`). demo-1's academy `serverChapterBody.js` sha was already `c9ae7057`
(= the patch's `pre_sha256`; the M237 "5-behind" was other files) and the running process already carried
`ACADEMY_DEMO_FS_PUBLISHED=1` + `NODE_ENV=development` (from the catalog-patch launch) — so a hot-reload of the
patched file sufficed, no restart. Advanced demo-1's rext clone to tag `sound-check-m238-ant-academy-reliability`
(as `devops`), applied the chapter-body patch via the M238 helper.
- **#3 BEFORE:** `/chapters/a-day-in-the-new-job/` → **HTTP 404**, `<title>` "Not Found — AI Academy", visible
  `<h1>` none, visible ~3.8 KB.
- **#3 AFTER:** same URL → **HTTP 200**, `<title>` "A Day in the New Job — AI Academy by Anthropos", visible
  `<h1>` "A Day in the New Job", Playwright `innerText` 706 chars of real chapter content. FIXED.
- **#2 AFTER:** `/chapters/<slug>/?lang=it` → **HTTP 200**, real chapter innerText (705 chars) — the language
  switch on a chapter reader renders (was the same backend-null 404 pre-fix). `/it` → 308→404 (a non-route,
  expected — locale is `?lang=`). #2 confirmed fresh: not a code bug.
- **§3 assertion validated live:** the fixed chapter passes `ANT_ACADEMY_CHAPTER_SECTION` (innerText 706 ≥ 400,
  no "wandered"); a bogus slug → HTTP 404, innerText len 0 → FAILs on the floor. The mustNotInclude is the
  semantic belt-and-suspenders (a padded-404 variant); the floor is the primary catch — both fail the 404.
- **NB — "wandered off the trail" lives in Next's RSC flight `<script>` data even on a 200 render**, so a
  raw-HTML `grep` false-positives; the sweep reads Playwright `innerText` (visible only), which is clean.
- Full `coverage.spec.ts` billion sweep run (auth crawl → academy follow) deferred to **M244** (its exit gate
  (c) "the 39 live-browser specs execute green" already owns it — Fate 2). Probe logic is unit-proven + its live
  premise (catalog.json slug + chapter render + innerText distinguish) validated here.
- End state: demo-1 left FIXED (both academy patches applied; rext clone at the M238 tag; patches revert on
  teardown). billion `/tmp/m238-*` scratch cleaned.

_(will accumulate topic → doc → code triples during build)_

## Pre-flight audits — §1 #3 chapter-body demo path
- **Verdict:** YELLOW (`kb-fidelity-audit.md`, 2026-07-21). Central premise ALIGNED; 1 incidental gap (KB-1). Proceed.
- **Topic → doc → code triples:**
  - #3 chapter-body → `corpus/services/ant-academy.md` + `demopatch-spec.md` + `frontend-tier.md` → `ant-academy/code/src/lib/serverChapterBody.js` (`resolveServerChapterBody`), `backendChapterBody.js`, `chapterContent.js` (`loadChapterShape`), `draftMode.js` (`draftsEnabled`); rext `demo-stack/patches/academy-fs-published-fallback/*.yaml` + `stack-injection/apply-academy-fs-published.sh` + `demo-stack/ant-academy.sh`.
  - #2 language → `ant-academy/code/src/lib/` i18n (coerceLocale/translate/localizedCatalog) + `/it` route.
  - coverage sweep → `corpus/ops/demo/coverage-protocol.md` (`player-academy`) + rext `stack-verify/e2e/`.
- **Code-verified facts:** `getBackendChapterBody`-null → `notFound()` (no general FS fallback; the dev-only draft path needs `NODE_ENV=development` ∧ `ACADEMY_SHOW_DRAFTS` ∧ `isFsDraftSlug`). The catalog demopatch (`academy-fs-published-fallback`) patches `serverTenant.js::getServerCatalogView` (the `?? emptyCatalogView()` anchor), env-gated on `ACADEMY_DEMO_FS_PUBLISHED`, applied by a native-run shell helper.
