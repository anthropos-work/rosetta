**Type:** tok (bootstrap)

# iter-01 progress — bootstrap tok

The unconditional iter-01 of the M42e iterative milestone. Authors the iteration protocol, scaffolds the
Playwright coverage harness, resolves the overview's four open questions, and records the initial strategy
(TOK-01). Does NOT terminate the call — the loop continues into iter-02 (the baseline-sweep tik) under TOK-01.

## Phase 0b — pre-flight KB-fidelity gate
`/developer-kit:audit-kb-fidelity --milestone=M42e` → **GREEN** (report: `../kb-fidelity-audit.md`). 5 fidelity
findings all ALIGNED, 0 blind areas. Two DOC-ONLY items (`coverage-protocol.md`, the link-rewriting fix
surface) are declared M42e deliverables. Strategy authored against verified docs.

## What landed (bootstrap deliverables)
1. **`corpus/ops/demo/coverage-protocol.md`** (NEW, corpus) — the sweep + triage + fix iteration protocol: the
   gate, the harness home, the BFS crawl strategy, the two-tier non-emptiness assertion, the fix-surface
   routing table, the re-scope/zero-edit line, the Phase A–E loop, the tooling-iter refinement, measurement
   conventions.
2. **The Playwright coverage harness** (rext, `stack-verify/e2e/`) — extends the existing smoke harness:
   - `lib/cockpit-login.ts` — log in as a roster hero via the fake-FAPI handshake (`POST /v1/demo/select` →
     drive a protected route → middleware handshake establishes the RS256 session as that seat). `demoLocalHosts`
     builds the offset-port authority set for escape-detection.
   - `lib/crawl.ts` — the in-app nav-link BFS crawler + per-page assertions (Tier-1 density floor +
     error-sentinel; every-link-host-demo-local escape check) + the `CoverageReport` roll-up.
   - `tests/coverage.spec.ts` — wires login + crawl + per-page screenshots + report emission; asserts the gate
     (`failing===0 && escapes===0 && reachable>0`).
   - `run-coverage.sh` — resolves offset/ports from N, installs deps+chromium (idempotent), runs the spec,
     prints the GATE roll-up.
   - `playwright.config.ts` — `ignoreHTTPSErrors: true` for the openssl-fallback FAPI cert.
   - `package.json` — `test:coverage` + `typecheck` scripts (Playwright `^1.49.0` already pinned).
   - **Validated:** `npx playwright test --list` compiles all specs (esbuild) → 3 tests in 2 files, no import
     errors. (The harness RUNS in iter-02; the bootstrap tok only scaffolds + validates compile.)

## Open-question resolutions
1. Harness home → **under `stack-verify/e2e/`** (reuses Playwright pin + offset plumbing; no new section).
2. Crawl → **pure in-app nav-link BFS** (a manifest can't catch escapes).
3. Assertion → **two-tier** (density floor default; per-section selectors escalated per-page).
4. Playwright wiring → **already pinned** at `stack-verify/e2e/package.json`; coverage spec is a sibling.

## Login mechanism (resolved from code)
Employee/member vantage seat = **`maya-thriving`** (live roster: `org_role: member`, `auth_id
user_seed_demo-3_1`, the default seat). Manager = `dan-manager` (`admin`). The handshake deep-link
`__clerk_identity=<key>` (or `POST /v1/demo/select`) selects the seat; `handleHandshake` establishes the
session as that hero.

## Close — 2026-06-25

**Outcome:** Bootstrap strategy (TOK-01) authored; coverage protocol + Playwright harness scaffolded + compile-validated; 4 open questions resolved. No gate metric move (bootstrap toks don't move the gate — baseline is iter-02).
**Type:** tok (bootstrap)
**Status:** closed-fixed
**Gate:** N/A for tok
**Phase 5 grading:** (1) gate-met: n — (2) triggered-tok: n (bootstrap, not triggered) — (3) re-scope: n — (4) user-blocker: n — (5) cap-reached: n — (6) protocol-stop: n — Outcome: continue
**Decisions:** D1 (harness home), D2 (crawl strategy), D3 (assertion shape), D4 (login seat) — see ./decisions.md; TOK-01 in ../decisions.md
**Side-deliverables (if any):** none
**Routes carried forward:** iter-02 — run the baseline employee sweep as `maya-thriving` against live demo-3; emit the coverage report; triage the highest-leverage failing cluster.
**Lessons:** Playwright was already pinned under `stack-verify/e2e/` (an existing unauthenticated smoke harness) — the milestone's "first non-Go rext dev/test dependency" was already present, so the harness is an extension, not a from-scratch add. Any coverage milestone (M42e, M42m) reuses this exact harness with a different vantage seed + seat key.
