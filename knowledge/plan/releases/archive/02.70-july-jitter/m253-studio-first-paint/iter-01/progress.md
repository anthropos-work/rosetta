**Type:** tok (bootstrap)

# M253 iter-01 — progress

Bootstrap tok: verified the boot chain against the KB, took the baseline, authored TOK-01.

## Pre-flight (Phase 0b — KB-fidelity, milestone-once)
Verdict: **GREEN (inline)**. The milestone is tightly specified with a pre-identified root cause; I verified the
load-bearing KB claims directly against code rather than spawning the heavyweight audit sub-agent (recorded as a
decision, D1). Evidence: (a) `latency-budget.md`'s never-gate-on-networkidle / always-gate-on-fresh-green rules
apply verbatim; (b) `main.ts` boot order (preloadCriticalCSS L97 → Sentry/posthog → clerk.load L181 → l12n L191
→ canAccess L199 → PageWrapper L206) matches the milestone's stated boot model; (c) `pageWrapper.js#init` wipes
`document.body.innerHTML` (L122) then rebuilds the skeleton — so the de-dup open question is already answered by
the source (an early-injected skeleton is auto-replaced, no double shell); (d) the M249 studio patch ladder +
patch-set fingerprint exist at `up-injected.sh:806` exactly as `demopatch-spec.md` §5-bis describes.

## Baseline
See overview.md — **skeleton-visible 4669 ms** on demo-2 (local laptop), dominated by `userService.canAccess()`'s
~3.9 s GraphQL retry ladder; clerk.load 140 ms, l12n 12 ms.

## Close — 2026-07-24

**Outcome:** authored TOK-01 (shell-before-awaits + no-thirdparty demopatches on the M249 ladder + net-new FCP
runner). Baseline established: skeleton-visible 4669 ms vs < 1000 ms gate. Dominant await identified =
`userService.canAccess()` (not clerk.load).
**Type:** tok (bootstrap)
**Status:** closed-fixed
**Gate:** N/A for tok
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (bootstrap does not exit) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue (bootstrap tok → loop into iter-02 tik)
**Decisions:** D1 (inline KB-fidelity verdict), D2 (dominant-await = canAccess, open question resolved)
**Side-deliverables:** none
**Routes carried forward:** iter-02 — author the 2 patches + FCP runner + extend the ladder + rebuild + measure.
**Lessons:** the 10 s clerk.load timeout was a red herring on this stack (140 ms actual); the blank is a GraphQL
retry ladder inside canAccess. Confirms the fix is purely a paint-ordering change, independent of the canAccess
404 (which is out of M253 scope).
