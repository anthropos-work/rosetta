# M253 — Retro (studio-desk first-paint)

## Summary
Cut studio-desk **first-meaningful-paint from ~4669 ms to p95 817 ms** (p50 743, max 817; 5/5 cold loads painted the
`.page-skeleton` shell, 0 login bounces) on demo-2 (**local laptop**) — **~5.7×** faster, decisively under the < 1000 ms
gate with no blank > 1 s. The fix is a **pure, zero-platform-edit paint-ordering** demopatch pair on the M249
`build_frontend_studio_desk` ladder: `studio-desk-shell-first-paint` (inject the `.page-skeleton` DOM synchronously
after `preloadCriticalCSS()`, **before** `Sentry.init`/`posthog.init`/`clerk.load`/`l12nService.init`/
`userService.canAccess`; de-dup automatic via `PageWrapper#init`'s `document.body.innerHTML` wipe) **chained** with
`studio-desk-no-thirdparty` (no-op `Sentry.init` + `posthog.init` on the demo host), the 5-manifest patch-set
fingerprint forcing a studio rebuild — plus a **net-new studio-FCP runner** (`stack-verify/e2e/run-studio-fcp.sh`).
Iterative: 1 tok (bootstrap TOK-01) + 2 tiks (iter-02 the fix + runner, iter-03 the 3 docs Delivers). **Closed
on-gate** (numerical, local-bootstrap charter); the fully-green COLD-p95 confirmation on billion is the deliberate
coordination-rule-9 split → M254 gate (f). Code-of-record **`july-jitter-m253-studio-first-paint` @ b8969c0** (on
origin; rung-zero verified). Delivers landed: `latency-budget.md` + `demopatch-spec.md` (21→23) + `studio-desk.md`.
Deferral audit GREEN. **0 platform-repo edits.**

## Incidents This Cycle
- **P3 (no shipped bug) — the milestone's own root-cause hypothesis was a red herring, corrected by measurement
  (iter-01 D2).** The overview's open question named `clerk.load`'s **10 s timeout** as the suspected dominant boot
  leg. The bootstrap per-leg instrumentation refuted it: clerk.load was **140 ms** (not 10 s), l12n 12 ms — the blank
  is a **~3.9 s `userService.canAccess()`** GraphQL org-memberships **404 → 3-attempt retry ladder** (1776 + 2102 ms).
  This turned the fix from a speculative timeout-tuning into a precise paint-ordering change that paints the shell
  **ahead** of the awaits, independent of the 404 (out of M253 scope). Measure-first earned its keep — a pre-designed
  fix would have chased the wrong leg.

## What Went Well
- **Measure-first (the bootstrap tok) pinned the exact leg before pinning any patch sha.** The per-leg baseline made
  the fix minimal and correct on the first tik — one paint-ordering demopatch pair, no timeout guesswork.
- **Rode M249's ladder, didn't fork it.** M253 extended the existing `build_frontend_studio_desk` patch-set
  fingerprint (3 → 5 manifests) so a pre-M253 studio image is detected stale + rebuilt — the additive-UI injection
  pattern M249 authored, applied to a new `main.ts` anchor. No new machinery.
- **De-dup was free by construction.** `PageWrapper#init` wipes `document.body.innerHTML` then rebuilds its own
  skeleton, so the early-injected shell is seamlessly replaced — no double-skeleton flash, no extra code.
- **The number landed with margin.** 817 ms p95 vs a 1000 ms gate (5/5 cold, 0 bounce) — not a squeaker.
- **Zero platform edits.** The whole first-paint fix is demopatches on the demo's own ephemeral clone + a docs pass;
  the canonical `studio-desk` repo is untouched.

## What Didn't
- **The fresh-green cold clause can't be satisfied on a warm local demo.** demo-2 is warm/partially-set-dressed, so a
  fully-green `autoverify.json` there is unrelated to studio first-paint and unachievable — the number is met, but the
  formal fresh-green COLD confirmation genuinely belongs to a cold billion bring-up (M254 gate (f), by coordination
  rule 9). This is why the milestone closes on-gate with a carry, not fully self-contained.
- **The FCP runner's secondary `browser FCP` field reads `n/a`.** Chromium paint-timing isn't populated in the
  measure path; the PRIMARY gate metric (skeleton-visible wall-clock) works and gates. Candidate for a
  `PerformanceObserver` harden (Fate-3 / future harden — recorded in decisions.md; not blocking).

## Carried Forward
- **CARRY-M253-01 → M254 (Fate 2, confirmed-covered).** The fresh-green COLD-p95 confirmation on billion — re-measure
  the studio FCP gate on a freshly brought-up, fully set-dressed cold demo with a green `autoverify.json`. M254
  `depends_on` M253; its exit gate **(f)** "studio first-paint < 1 s cold p95 (← M253)" already owns it verbatim. No
  sibling `overview.md` edit; recorded in `carry-forward.md` + the deferral audit. The deliberate coordination-rule-9
  split (two live-measured iteratives can't share billion RAM), NOT a gate miss.
- **FCP-runner `browser FCP` field → `PerformanceObserver` harden (Fate-3 / future).** The primary skeleton-visible
  wall-clock metric works; the secondary field is a nice-to-have hardening candidate. Not owed to any current
  milestone.

## Metrics Delta
- **Benchmark (demo-2, local laptop, 5 cold loads):** skeleton-visible **4669 ms → p95 817 ms** (p50 743, max 817;
  samples 817/795/480/539/743 ms), 5/5 reached the shell, 0 login bounces. **PASS** (p95 < 1000 ms, max ≤ 1000 ms).
- **Demopatch inventory:** 21 → 23 (studio-desk 3 → 5). `TestPatchInventory` re-pinned in rext at the tag.
- **rext:** net-new `run-studio-fcp.sh` + `tests/studio-fcp.spec.ts` + `lib/studio-fcp.ts`; `demopatch check` PASS on
  both new manifests; `bash -n` clean on the extended ladder.
- **Rosetta:** docs corpus — no test suite; 0 broken markdown links across the 3 touched corpus docs.
- **Flake:** 0. **Platform-repo edits:** 0. **Deferral audit:** GREEN.
- Full machine-readable delta: `metrics.json`.
