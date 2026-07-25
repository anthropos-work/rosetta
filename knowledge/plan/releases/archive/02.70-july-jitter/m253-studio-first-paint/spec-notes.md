# M253 — Spec notes

Per-lane build notes for the studio-desk first-paint milestone — accumulated during the measure→patch→re-measure loop.

## Pre-flight audits — iter-01
KB-fidelity: **GREEN (inline, iter-01 D1)**. Verified `latency-budget.md` rules + the `main.ts` boot order +
`pageWrapper.js#init` body-wipe de-dup + the M249 ladder/fingerprint at `up-injected.sh:806` directly against
code; all hold. Baseline: skeleton-visible 4669 ms (demo-2, laptop).

## Shell-before-awaits demopatch (core/main.ts reorder)
Baseline timeline (demo-2, authenticated as maya-thriving, t0 = studio goto):
`clerk.load` ~140 ms · `l12nService.init` ~12 ms · **`userService.canAccess` ~3.9 s** (GraphQL org-memberships
404 → 3-attempt retry ladder 1776 + 2102 ms) · `new PageWrapper()` builds the skeleton only after → skeleton
paints at 4669 ms. Fix: inject `.page-skeleton` DOM synchronously after `preloadCriticalCSS()` (L97), before the
awaits. De-dup: `PageWrapper#init` wipes `document.body.innerHTML` (L122) then rebuilds — early shell is
auto-replaced.

## studio-desk-no-thirdparty twin
_(TBD during build.)_

## Net-new studio-desk FCP runner (stack-verify/e2e)
_(TBD during build.)_

## Docs
_(TBD during build.)_
