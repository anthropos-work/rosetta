# M42e Decisions

Implementation decisions with rationale, recorded during the iteration loop (harness home, crawl strategy,
assertion shape, link-rewriting surface, escalations of platform-only blockers).

| ID | Decision | Rationale | Date |
|----|----------|-----------|------|

## RE-SCOPE-TRIGGER (iter-09): the `/reimport-profile` → linkedin.com/help escape is platform-only — 2026-06-25

**Trigger:** the zero-edit-line re-scope trigger (`coverage-protocol.md` "Re-scope trigger"). iter-09 honestly
seeded the sim-start entitlement (the dishonest iter-08 skip is gone — `/start` is scored + renders), driving
**failing → 0**. But putting `/start` in scope + the iter-09 settle fix deepened the frontier to a
TRUE-EXHAUSTED **95 pages**, which surfaced a previously-hidden **escapes=1**: the `/reimport-profile` page
links to `https://www.linkedin.com/help/linkedin/answer/a522735/...` (a "how to find your LinkedIn public
profile URL" help reference). The HONEST residual is therefore **`(failing=0, escapes=1)`**, NOT `(0,0)` —
the employee gate is **failing-1 on escapes**, not met.

**Why it's platform-only:** the URL is a **hardcoded `<a href>` string literal** in the platform UI
(`next-web-app/packages/ui/src/Onboarding/OnboardingImportOptions/OnboardingImportLinkedin.tsx:99` + 3 sibling
literals in `OnboardingImportLinkedin.tsx`/`ImportProfileFailed/index.tsx`) — it is NOT env-driven, so the demo
injection/env link-rewriting CANNOT touch it. The only complete fix is a **platform-repo edit** (rewrite/remove
the hardcoded LinkedIn help URLs), which the v1.10 zero-edit line forbids.

**Nature of the link:** it is a `<Typography.Link target='_blank'>` to LinkedIn's own help DOCUMENTATION inside
the profile-import form — a legitimate external help reference (analogous to the iter-08 editorial citations:
real, opens in a new tab, not a nav-chrome escape to a competing product surface like `studio.anthropos.work`).
It is NOT "valid filler" under the strict gate definition, but it is also not a degraded/broken surface.

**The user decision (per `coverage-protocol.md` "Re-scope trigger" + the skill's Phase 5 §3):** choose one of —
- **(a) Allow-rule extension (in-rext, no platform edit):** classify a baked external **help/documentation**
  link (e.g. `linkedin.com/help`) on a profile-import surface as a VALID presenter-note citation, NOT a gate
  escape — the same disclose-don't-hide design as the `/chapter` editorial-citation allow-rule (iter-08 D3),
  extended to onboarding help links. This makes the gate honestly `(0,0)` with the link DISCLOSED (not hidden).
  Needs the user's sign-off because it WIDENS the allow-rule's escape carve-out beyond `/chapter` content.
- **(b) Upstream platform PR (out-of-band):** rewrite/remove the hardcoded LinkedIn help URLs in next-web-app.
- **(c) Carve `/reimport-profile` out of the employee gate** with a documented rationale (it is a profile
  RE-import utility; the seeded hero Maya already has a full M39–M41 profile, so a presenter never needs it).

Recommendation: **(a)** — it mirrors the already-sanctioned editorial-citation disclose-not-hide design, keeps
the gate honest (the link is reported as a presenter note, not silently skipped), stays in-rext (zero platform
edit), and the allow-rule stays narrow (a nav-chrome/baked-app-URL escape to a competing product still fails).
Surfaced to the user for the decision; the iter does NOT unilaterally widen the gate's escape semantics.

## TOK-01: sweep-then-route-by-leverage — 2026-06-25

**Tok type:** bootstrap (iter-01)
**Initial strategy:** Run the Playwright coverage sweep as the employee hero (`maya-thriving`) against live
demo-3, then iterate **highest-leverage-cluster-first**: each tik runs the sweep (Phase A), triages the
failing pages + escapes by fix surface (Phase B, the routing table in `coverage-protocol.md`), lands the fix
in the routed rext surface (Phase C — `stack-seeding` for empty sections, `stack-snapshot` serve-grants for
content errors, the demo injection/env link-rewriting for escapes, roster/FAPI for identity gaps), re-applies
the affected stack step to the live demo, re-sweeps (Phase D), and closes on whether the targeted cluster
cleared. Drive toward `(failing-pages, escapes) = (0, 0)` over the employee vantage's reachable set.
**Rationale:** The page set + failure modes are discovered by the sweep, not enumerable up front (the reason
the milestone is iterative). M39-M41 already landed the known high-leverage fills (G1 org-name, G2 role
backfill, G3 work/education, G4 avatars, G5 skill depth, G6 library serve-grants), so the sweep chases the
**tail** — the residual empties, the under-investigated G7 activities feed, and any escape links the baked
URLs miss (no studio-host rewrite exists yet). Leverage-first ordering (most pages unblocked per fix) clears
the dominant clusters first; a single serve-grant or seed can light up many pages.
**Strategy class:** new-direction
**Distance-to-gate context:** Gate metric = the coverage report's `(failing-pages, escapes)`; gate = `(0, 0)`
over the employee vantage's reachable pages. Starting value UNMEASURED — the baseline sweep is iter-02 (the
first tik). Known risk areas from `.agentspace/profile_gaps.md`: G7 activities feed (under-investigated), and
escape links (no `NEXT_PUBLIC_STUDIO_URL`-style rewrite baked → a left-menu "Studio" likely escapes to prod).
**Next-tik direction:** iter-02 — run `run-coverage.sh 3 employee maya-thriving` against live demo-3; capture
the baseline `(reachable, failing, escapes)`; triage the highest-leverage failing cluster + pick it as iter-02's
target (or iter-03's if iter-02 establishes the baseline + the first fix in one tik).

