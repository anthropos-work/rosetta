---
iteration_type: tok
tok_flavor: bootstrap
status: closed-fixed
milestone: M253
iter: 1
created: 2026-07-24
---

# M253 iter-01 — bootstrap tok (initial strategy + baseline)

**Type:** tok (bootstrap) — authors the first strategy (TOK-01). Does NOT terminate the call; iter-02 (tik)
continues under TOK-01 within the same invocation.

## Inputs
- `overview.md` (exit_gate, In/Out, KB deps), `spec-notes.md`, protocol `corpus/ops/demo/latency-budget.md`.
- Direct code verification of the studio boot chain (`stack-demo/studio-desk/app/core/main.ts`) +
  `pageWrapper.js` (the de-dup mechanism) + the M249 studio patch ladder in
  `rosetta-extensions/demo-stack/up-injected.sh:806` (`build_frontend_studio_desk`).

## Baseline measurement (distance-to-gate)
Environment: **local laptop, demo-2** (studio-desk :29000), single confirming cold load, authenticated as
`maya-thriving` via the real cockpit handshake CTA.

- **skeleton-visible = 4669 ms** (gate: < 1000 ms). Blank lasts ~4.7 s.
- Boot timeline (console, t0 = studio goto):
  - `clerk.load()` → ~140 ms (NOT the 10 s timeout — cheap vs Clerkenstein)
  - `l12nService.init()` → ~12 ms (cheap)
  - **`userService.canAccess()` → ~3.9 s** — its GraphQL org-memberships check **404s** and burns a
    3-attempt retry ladder (1776 ms + 2102 ms backoff) before returning.
  - `new PageWrapper()` (which builds the `.page-skeleton` DOM) runs only AFTER those three awaits → the
    skeleton paints at ~4669 ms.

**Open-question resolution (overview.md):** the dominant await is **`userService.canAccess()`**, not
`clerk.load`'s 10 s timeout (which was 140 ms here) and not `l12nService.init()`. The fix does not need to touch
canAccess — it needs to paint the shell BEFORE it.

## Initial strategy → see TOK-01 in the milestone-root decisions.md
Inject `.page-skeleton` DOM synchronously right after `preloadCriticalCSS()` (main.ts L97), before
Sentry/posthog/clerk.load/l12n/canAccess, via a sha-pinned demopatch on the M249 studio ladder; de-dup is
automatic (`PageWrapper#init` wipes `document.body.innerHTML`). Plus a `studio-desk-no-thirdparty` twin patch.
Author a net-new studio-FCP runner in `stack-verify/e2e/`. Iterate on demo-2.

## Next-tik direction (iter-02)
Author the 2 demopatch manifests (shell-before-awaits + no-thirdparty) + the FCP runner, extend the
`build_frontend_studio_desk` ladder + patch-set fingerprint, rebuild the studio image on demo-2, and measure
5-cold-load p95 FCP.
