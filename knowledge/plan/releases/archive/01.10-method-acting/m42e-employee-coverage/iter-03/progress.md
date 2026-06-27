**Type:** tik / iter_shape: tooling -- under TOK-01, coverage-protocol.md "Iter type selection -> Tooling-iter".

# iter-03 progress -- tooling-iter: fix the crawl wait strategy

Ships the harness wait-strategy capability AND uses it (re-sweep) within the same iter. The prior tik (iter-02)
closed blocked on the harness `networkidle` flake (D2); this iter lands the fix + measures the true residual.

## Line 1 -- ship the capability
Adopted the prior-attempt rext WIP (`crawl.ts` + `coverage.spec.ts`, uncommitted -> now committed + tagged):
- `page.goto` `waitUntil: 'networkidle'` -> **`'domcontentloaded'`** + a **bounded** 4 s `networkidle` settle
  (`.catch(()=>{})`) -- loads + paints without ever blocking on a never-idle page.
- **Inline screenshots** via a new crawl `onPage` hook (captured while the page is already loaded) -- replaces
  the 2nd full re-navigation screenshot pass that re-triggered the timeout + exhausted the budget.
- **Seed-vs-discovered scoring:** a guessed seed that 404s or redirects away is dropped (not false-scored);
  only nav-discovered pages are coverage commitments. Seeds adjusted (`/skills` 404 removed; `/` is the
  post-login landing, redirect followed + scored).
- Compile-validated via `npx playwright test --list` (esbuild; 3 tests, no import errors). (`tsc` not installed
  -- pre-existing debt, D7.)

## Line 2 -- use it (Phase D re-sweep)
`run-coverage.sh 3 employee maya-thriving` vs live demo-3, ~105 s to the gate line (the sweep itself finishes
fast now; the residual ~11 min wall is Playwright trace teardown, not page timeouts).

**Re-sweep metric: `(failing=8, escapes=1)` over reachable=80 -- down from the baseline `(44,1)`.** GATE: RED.
The fix removed 36 of the 44 false failures.

## Triage of the TRUE residual (8 failures + 1 escape) -- routed to iter-04
- **1x harness false-positive (`/`, D3):** root has no `<main>`; the assertion read the whole `<body>` (189016
  chars incl. an inlined i18n JSON blob) and the `"Something went wrong"` sentinel matched a **translation
  string** (`errorSettingUpGPTRealtime`), not a rendered error. -> iter-04 assertion-tuning. Folded into
  `coverage-protocol.md` measurement conventions this iter (protocol-evolution).
- **5x `/sim/.../result/<uuid>` empty (text=0) (D4):** per-session sim result deep-links. -> iter-04: seed the
  result data (`stack-seeding`) OR exclude per-UUID result pages from the vantage's reachable set (crawl-scope).
- **2x empty skill-path indexes (D5):** `prompt-engineering-fundamentals` + `foundation-of-artificial-intelligence`
  (text=0 while ~20 siblings rendered 2k-24k). -> iter-04: `stack-seeding` / `stack-snapshot` serve-grant.
- **1x escape (D6):** the `strategy-business.com` link persists (content link, unaffected by the wait fix).

## Phase E -- Close

**Outcome:** Tooling-iter landed the wait-strategy fix (`networkidle`->`domcontentloaded` + bounded settle +
inline screenshots + seed-vs-discovered scoring); re-sweep `(failing=8, escapes=1)` down from `(44,1)` -- 36
false failures removed. TRUE residual (8 fails + 1 escape) triaged into 4 clusters, routed to iter-04. Lesson
folded into `coverage-protocol.md`.
**Type:** tik (iter_shape: tooling)
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n -- (2) triggered-tok: n -- (3) re-scope: n -- (4) user-blocker: n -- (5) cap-reached: n -- (6) protocol-stop: n -- Outcome: continue
**Decisions:** D1 (adopt WIP), D2 (re-sweep 8/1), D3 (`/` false-positive), D4 (5 sim-result empties), D5 (2 empty skill-paths), D6 (escape persists), D7 (tsc debt) -- see ./decisions.md
**Side-deliverables (if any):** none (the protocol-doc update is in-scope protocol-evolution, same commit)
**Routes carried forward:**
- iter-04 (Fate-3, handler `SEED-M42e-sim-result-pages`): the 5 `/sim/.../result/<uuid>` empties -- seed or crawl-scope.
- iter-04 (Fate-3, handler `SEED-M42e-empty-skillpaths`): the 2 empty skill-path indexes -- seed / serve-grant.
- iter-04 (Fate-3, handler `ASSERT-M42e-root-sentinel`): the `/` sentinel false-positive -- visible-main-region
  assertion / root normalization.
- iter >= 04 (Fate-3, handler `ESCAPE-M42e-skillpath-external-article`): the `strategy-business.com` escape.
**Lessons:** `networkidle` is fatal for crawling next-web (never idle -> every nav times out); use
`domcontentloaded` + a bounded settle + inline screenshots. Seed paths are guesses (drop 404/redirect-away, do
not score). The error-sentinel must read the visible `<main>`, not the raw `<body>` (an inlined i18n JSON error
table false-matches). All three folded into `coverage-protocol.md` measurement conventions (protocol-evolution)
-- they apply to the manager vantage (M42m) verbatim.
