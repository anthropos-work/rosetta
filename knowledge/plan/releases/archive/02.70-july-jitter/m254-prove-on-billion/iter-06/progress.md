**Type:** tik — TOK-01 cluster 4 (fix-forward). (f) studio session-carry: the linchpin (blocks (c)-studio + (e)).

# M254 · iter-06 — progress

## Close — 2026-07-25

**Outcome:** (f) session-carry is **NOT a defect** — it was a measurement artifact (the FCP probe defaulted to
`maya-thriving`, an employee that studio's `checkEnterpriseAndAdmin` bounces by design). Shipped the tooling
fix (default identity → `dan-manager`, a studio-eligible admin) — rext `cbe9256`, tag
`july-jitter-m254-studio-fcp-identity` **on origin**. Live-proven: admin heroes reach the studio shell
(reachedShell all true). This unblocked **(c)-studio render (now LIVE 3/4)** and **(e)** the builder Playthrough.
FCP shell paint p50 637-726 ms < 1 s on billion (the M253 shell fix HOLDS); p95 tailnet-jitter-bound.

**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET (overall) — (f) session-carry MET (fixed + live); (f) FCP <1 s p95 disposition-pending; (c) now 3/4 render LIVE; (e) unblocked, (h)-Playthroughs still pending.
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n — (3) re-scope: n — (4) user-blocker: n (the FCP-p95 disposition is a surface/route-forward item; the code fix is complete + shipped, so it changes no code that lands this iter) — (5) cap-reached: n (tik 2 of the session) — (6) protocol-stop: n — Outcome: continue (→ iter-07).
**Decisions:** D1 (f) root cause = wrong-default-hero measurement artifact (not session-carry); D2 fix = default identity maya→dan (rung-zero); D3 FCP p50 <1s / p95 tailnet-jitter → disposition-pending; D4 (c)-studio render LIVE-confirmed (unblocked).
**Side-deliverables:** none (the fix IS the planned scope; temp Playwright diagnostics used + removed, rext left clean).
**Routes carried forward:**
- **(f) FCP <1 s cold p95 on billion** — p50 637-726 ms < 1 s (shell fix holds), but p95 tailnet-jitter-bound (outliers 1443/2014/4943 ms; reachedShell always true). The roadmap's anticipated "environment-sensitive" (f). → **coordinator disposition** (analogous to the (b) presence-only disposition). Handler: `DISP-M254-studio-fcp-tailnet-p95`.
- **(c)-academy** durability fix (from iter-05); **(e)** builder Playthrough (now unblocked); **(g)** host-sensitive tests; **(h)-Playthroughs**.
**Lessons:** (1) A "session-carry / --public-host" symptom can be an AUTHZ artifact — studio's two-stage gate (`requireAuth` networkless-JWT PASSES, then `checkEnterpriseAndAdmin` does a per-request fake-bapi `getOrganizationMembershipList` round-trip) bounces a non-admin to the web app, which looks identical to "session didn't carry." Always distinguish the login-302 (→/login, unauth) from the eligibility-303 (→web-app /home, authed-but-ineligible). (2) A combined Playwright selector + `.first()` silently falls back to the first-card CTA for every hero — always pin the exact `__clerk_identity=<hero>` href, or you log in as the default seat and mis-diagnose. (3) The studio per-request server-side auth round-trip is IN the shell-paint path, so studio FCP is inherently more tailnet-jitter-sensitive than a static asset.
