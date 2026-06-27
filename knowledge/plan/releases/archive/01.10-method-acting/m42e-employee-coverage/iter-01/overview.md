---
iter: 01
milestone: M42e
iteration_type: tok
tok_flavor: bootstrap
status: closed-fixed
date: 2026-06-25
---

# iter-01 — bootstrap tok (author strategy + protocol + scaffold harness)

## Type
tok (bootstrap) — the unconditional iter-01 of an iterative milestone. Authors the FIRST strategy. Does NOT
terminate the call; the loop continues into iter-02 (a tik) under TOK-01.

## Inputs
- `overview.md` (exit gate + the four open questions + the per-failure fix-surface map + KB contract).
- `spec-notes.md` (the harness/login/crawl/assertion/report TODOs).
- `.agentspace/profile_gaps.md` (the v1.10 design input — G1-G7; G1/G2/G3/G4/G5/G6 addressed by M39-M41; G7
  activities feed flagged under-investigated).
- The KB-contract docs: `frontend-tier.md`, `verification.md`, `rosetta_demo.md`, `recipe-browser-login.md`,
  `stories-spec.md` — all PAIRED + ALIGNED (KB-fidelity GREEN).
- The protocol doc this iter authors: `corpus/ops/demo/coverage-protocol.md`.
- Live demo-3 (17 containers, M39+M40+M41 seeded) on offset ports — verified responding.

## Bootstrap deliverables (this iter)
1. **`corpus/ops/demo/coverage-protocol.md`** (NEW) — the sweep + triage + fix iteration protocol. DONE.
2. **Resolve the four open questions** (recorded in spec-notes + below). DONE.
3. **Scaffold the Playwright coverage harness** under `stack-verify/e2e/` (extends the existing smoke harness).
4. **Initial strategy (TOK-01)** + next-tik direction (the baseline sweep). Recorded in milestone `decisions.md`.

## Open-question resolutions (overview §"Open questions")
1. **Harness home** → **under `stack-verify`** (`stack-verify/e2e/coverage.spec.ts` + a runner), NOT a new
   `stack-coverage` section. Rationale: `stack-verify/e2e/` ALREADY pins `@playwright/test ^1.49.0` with a
   working `playwright.config.ts` (offset via `ROSETTA_E2E_BASE_URL`) + an unauthenticated smoke test. The
   coverage sweep is the authenticated multi-page evolution of that exact harness; reuses verify's
   offset/project/scope plumbing (`lib/target.sh`, `STACK_PROJECT`/`STACK_OFFSET`). No new section needed.
2. **Crawl strategy** → **pure in-app nav-link discovery (BFS)**, NOT a static route manifest. The gate
   requires escape-detection; a manifest can't see a nav that escapes the demo (external links are invisible
   to a route list). The crawl observes the actual rendered same-origin links + nav chrome.
3. **Non-empty assertion shape** → **two-tier**: Tier-1 generic text-density floor + error-sentinel (default,
   cheap, catches the dominant empty mode); Tier-2 per-section DOM selectors escalated per-page only when
   Tier-1 false-passes/-fails. The tier choice is an iter decision.
4. **Playwright wiring** → already pinned at `stack-verify/e2e/package.json` (`@playwright/test ^1.49.0`) with
   `npm`-managed lockfile + `install:browsers` script. The coverage spec is a sibling; no new dependency wiring
   beyond a sibling spec + a runner that injects the cockpit login + the offset baseURL.

## Login mechanism (resolved from code)
The cockpit `[Login as]` is a deep-link redirect through the fake-FAPI handshake:
`https://<fapi-host>/v1/client/handshake?…&__clerk_identity=<hero-key>` — `handleHandshake`
(`clerk-frontend/server.go:176`) selects the seat (`s.reg.Select(key)`) and establishes the RS256 session as
that hero. The employee/member vantage hero is **Maya Chen** (`stories.seed.yaml`, vantage `end-user`). The
harness drives next-web through this handshake with Maya's roster key, picking up the `__session` cookie, then
crawls. `ignoreHTTPSErrors: true` covers the openssl-fallback FAPI cert.

## Next-tik direction (iter-02, the first tik)
Run the **baseline sweep** as Maya (employee vantage) against live demo-3: crawl the reachable employee pages,
emit the coverage report `(failing-pages, escapes)`. This is the baseline metric. Then triage the highest-
leverage cluster.

## Phase plan
Bootstrap: author protocol (Phase A authoring) → scaffold harness → record TOK-01. No metric move expected
(bootstrap toks don't move the gate).
