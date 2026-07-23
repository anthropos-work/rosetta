---
title: "KB Fidelity Audit — M238 ant-academy reliability"
date: 2026-07-21
scope: milestone:M238
invoked-by: build-milestone
---

## Verdict
YELLOW — proceed with tracking (KB-1 recorded in decisions.md).

## Topic Inventory

| Topic | Knowledge doc | Code paths | Status |
|---|---|---|---|
| #3 chapter-body demo path (Start→404) | `corpus/services/ant-academy.md` (catalog side only), `corpus/ops/demo/demopatch-spec.md`, `corpus/ops/demo/frontend-tier.md` | `ant-academy/code/src/lib/serverChapterBody.js`, `backendChapterBody.js`, `chapterContent.js`, `draftMode.js`; rext `demo-stack/patches/academy-fs-published-fallback/`, `stack-injection/apply-academy-fs-published.sh`, `demo-stack/ant-academy.sh` | PAIRED (catalog) + BLIND-AREA (chapter-body demo path is undocumented — **covered by M238 Delivers**) |
| #2 language switch error | `corpus/services/ant-academy.md` (no explicit i18n demo claim) | `ant-academy/code/src/lib/` i18n (coerceLocale/translate/localizedCatalog), `/it` route | DOC-ONLY thin — investigation item; no stale load-bearing claim |
| Academy presence/coverage sweep | `corpus/ops/demo/coverage-protocol.md` (`player-academy` → `/courses/<slug>`), `corpus/ops/demo/frontend-tier.md` | rext `stack-verify/e2e/` | PAIRED (course-page presence) + BLIND-AREA (chapter-body render assertion — **covered by M238 Delivers/In-scope**) |

## Fidelity Findings

### #3 — chapter-body is backend-authoritative with no general FS fallback → 404
- **Source:** milestone premise (roadmap M238 In-list) vs `serverChapterBody.js`.
- **Expected:** bodies backend-authoritative; a backend-null → `notFound()` → the "You wandered off the trail" 404; the existing catalog demopatch does NOT cover the body.
- **Actual:** `resolveServerChapterBody(slug, locale)` = `getBackendChapterBody(slug, locale)`; if null → `maybeResolveDraftBody` (dev + `ACADEMY_SHOW_DRAFTS` only, gated on `isFsDraftSlug`) → else `{ notFound: true }`. `academy-fs-published-fallback.yaml` `path: code/src/lib/serverTenant.js` (the `getServerCatalogView` catalog fallback) — NOT the body.
- **Verdict:** ALIGNED. The root-cause is code-accurate.

### demopatch apply-vehicle for native ant-academy patches
- **Source:** `demopatch-spec.md` §4 "Three apply vehicles".
- **Expected/Actual:** natively-run ant-academy patches apply via a `stack-injection/apply-*.sh` shell helper (apply-before-launch, revert-on-stop), NOT the `demopatch` tool. Confirmed: `apply-ant-academy-dev-origins.sh` documented; `apply-academy-fs-published.sh` exists + is wired in `ant-academy.sh`.
- **Verdict:** ALIGNED (pattern documented). M238's chapter-body patch follows the same vehicle.

## Completeness Gaps

1. **(KB-1, incidental)** `demopatch-spec.md` patch inventory is stale: it states "11 patches … 1 × ant-academy" and inventories only `ant-academy-dev-origins`, **omitting `academy-fs-published-fallback`** (a real, wired native-run patch). 14 patch dirs on disk. → Fix the **academy portion** in M238 Phase 5 (add `academy-fs-published-fallback` + the new chapter-body patch); the non-academy count drift is out of M238's subject area (not M238's to fix).

## Applied Fixes
None inline (the one gap is a Phase-5 doc update tied to M238's own new patch, not a pre-existing stale load-bearing claim).

## Open Items (require user decision)
None.

## Gate Result
YELLOW: proceed. Every blind area is covered by M238's `Delivers →` line (`ant-academy.md` + `frontend-tier.md`) and the In-scope "extend the academy coverage sweep" item. No blind area un-owned; no stale claim the implementation would read as truth. KB-1 tracked for Phase 5.
