**Type:** tik (baseline sweep) -- under TOK-01, coverage-protocol.md Phase A.

# iter-02 progress -- baseline employee coverage sweep

The first tik of M42e: run the harness against live demo-3 as the employee hero (`maya-thriving`), capture the
baseline coverage metric, and triage the dominant failing cluster for iter-03.

## Phase A -- Sweep (baseline)
`run-coverage.sh 3 employee maya-thriving` against live demo-3 (17 containers, consumed tag `method-acting-m41`).
Foreground run, ~10 min (hit the test's 600s budget during the screenshot phase).

**Baseline metric: `(failing=44, escapes=1)` over `reachable=80`.** Report:
`stack-verify/e2e/coverage-out/employee/coverage-report.json`. GATE: RED.

- **36 pages OK** with real content -- `/profile` (145), `/profile/activities` (998), `/profile/skills` (1572),
  `/library` (259), `/library/skill-paths` (17288), and the full `/skill-path/*` index set (2k-24k textLen each),
  plus several `/sim/*` + `/skill-path/.../chapter` pages (600-946). The employee **core** is fully populated
  (M39-M41 + G6 serve-grants landed).
- **44 pages "FAIL"** -- but all with `http=- text=0` and a `page.goto` **TimeoutError** note. NOT empty pages.

## Phase B -- Triage
Classified the 44 failures + 1 escape (full breakdown in `decisions.md`):

- **Dominant cluster (44/44) = a harness capability gap (D2).** The crawl waits `waitUntil: 'networkidle'`,
  which **never settles** on this app (long-lived ws/polling connections). Every `page.goto` eats its full 30s
  timeout; the 600s test budget exhausts before later pages are reached -> they record as `empty/error`.
  **Probe-confirmed** (a direct Playwright probe outside the harness): `/home`, `/dashboard`,
  `/skill-path/.../chapter` all return **http=200 with real content** (textLen 534-885) under
  `domcontentloaded`, but `networkidle=NEVER-IDLE(6s)`. So the failures are a wait-strategy artifact, not real
  empties. -> the protocol's named **tooling-iter** trigger. Routed to **iter-03 (tooling-iter)**: swap the
  crawl wait `networkidle` -> `domcontentloaded` + a bounded settle, and re-tune `maxPages`/budget so the
  sweep finishes within budget. Only after that fix is the TRUE failing count knowable.
- **1 escape (real, D3):** `/skill-path/how-to-give-feedback-.../chapter` body copy links to
  `https://www.strategy-business.com/article/...` (an editorial external link inside replayed skill-path
  content). Routed forward (Fate-3, iter >= 04) -- a content/allow-rule decision, distinct from the harness cluster.
- **Crawl saturated at maxPages=80 (D4)** -- the BFS reached every library `/sim/*` + `/skill-path/.../chapter`.
  The genuine reachable set + true failing count are only knowable after the wait-strategy fix.

## Phase C -- Fix
None this tik. The dominant-cluster fix is a **harness capability change** the protocol explicitly routes to a
tooling-iter (iter-03); landing it mid-baseline-tik would blur this tik's measurement deliverable and open a
2nd line of investigation (scope-creep tripwire). Baseline + triage is the planned deliverable.

## Phase D -- Re-sweep
N/A (pure baseline; no fix landed this tik).

## Close -- 2026-06-25

**Outcome:** Baseline established: `(failing=44, escapes=1)` over `reachable=80` (employee/`maya-thriving`).
Metric moved UNMEASURED -> `(44, 1)`. Dominant cluster (all 44 failures) probe-confirmed as a `networkidle`
harness flake (pages are http=200 with real content) -> routed to iter-03 (tooling-iter). 1 real content
escape + the maxPages saturation routed forward.
**Type:** tik
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n -- (2) triggered-tok: n -- (3) re-scope: n -- (4) user-blocker: n -- (5) cap-reached: n -- (6) protocol-stop: n -- Outcome: continue
**Decisions:** D1 (baseline metric), D2 (networkidle harness flake -> tooling-iter), D3 (1 real content escape), D4 (maxPages saturation) -- see ./decisions.md
**Side-deliverables (if any):** none
**Routes carried forward:**
- iter-03 (tooling-iter, Fate-3, handler `TOOL-M42e-iter03-wait-strategy`): fix the crawl wait strategy
  (`networkidle` -> `domcontentloaded` + bounded settle) + re-tune `maxPages`/budget; re-sweep for the TRUE
  failing count. (coverage-protocol.md "Iter type selection -> Tooling-iter".)
- iter >= 04 (Fate-3, handler `ESCAPE-M42e-skillpath-external-article`): the 1 real escape -- the external
  `strategy-business.com` link in the feedback skill-path chapter body. Content/allow-rule decision.
**Lessons:** `waitUntil: 'networkidle'` is the wrong wait for this app -- next-web holds long-lived connections
so networkidle never fires; the baseline's "failures" were a wait-strategy artifact, not empty pages. ANY
coverage sweep against this platform must use `domcontentloaded` + a bounded settle, never bare `networkidle`.
This lesson generalizes to the manager vantage (M42m) -- folded into coverage-protocol.md's measurement
conventions in iter-03 (the tooling-iter that lands the fix), per the protocol-evolution guideline.
