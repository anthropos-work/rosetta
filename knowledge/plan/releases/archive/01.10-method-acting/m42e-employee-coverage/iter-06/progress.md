**Type:** tik / iter_shape: tooling -- under TOK-01, coverage-protocol.md "Iter type selection -> Tooling-iter".

# iter-06 progress -- tooling-iter: the `/` sentinel false-positive (innerText + seed-drop)

Clears the last harness false-positive in the residual (the root `/` flagged `error` by a sentinel match in an
inlined i18n JSON blob). Ships the assertion fix + the `/`-seed correction + re-sweeps. tik #5 -> the 5-tik cap
fires after this iter closes.

## Line 1 -- ship the capability
1. **`crawl.ts`: `textContent()` -> `innerText()`** (D1) -- the non-emptiness/sentinel assertion now reads only
   VISIBLE text, structurally excluding next-web's inlined i18n translation table (whose "Something went wrong"
   value false-matched the sentinel on the fully-rendered root `/`).
2. **`coverage.spec.ts`: dropped `/` from `EMPLOYEE_SEED_PATHS`, replaced with `/home`** (D3) -- the innerText
   fix exposed `/` as a `<main>`-less, visibly-empty redirect/loading shell (text=0), not a real content page;
   `/home` is the real landing (text=347 ok), also reached via nav. Login is unaffected (`loginAs` still uses
   landingPath `/`).
Compile-clean via `playwright test --list`.

## Line 2 -- use it (Phase D re-sweep, twice)
- **Re-sweep #1 (innerText only):** `(failing=3, escapes=2)` -- `/` flipped `error,189016` -> `empty,0` (D2):
  false-positive GONE, but `/` now reads its real (empty) shell text -> still counted.
- **Re-sweep #2 (after dropping `/` seed):** **`(failing=2, escapes=3)`** -- `/` removed from the report (D4).
  `failing` 3 -> 2. GATE: RED.

## Triage of the final residual (2 failures + 3 escapes) -- routed forward
- **2 empty skill-path indexes (D4):** `prompt-engineering-fundamentals` + `foundation-of-artificial-intelligence`
  (text=0 while ~20 siblings render 2k-24k). The single real PRODUCTION content gap remaining. ->
  `stack-seeding`/`stack-snapshot` serve-grant (handler `SEED-M42e-empty-skillpaths`). Next session's iter-07.
- **3 external skill-path-chapter article links (escapes, D5):** strategy-business.com + en.wikipedia.org +
  dremio.com -- all editorial links in replayed chapter body copy. escapes 1->2->3 is the deeper-crawl-discovery
  pattern, not a regression. -> content/allow-rule (handler `ESCAPE-M42e-skillpath-external-articles`).

## Phase E -- Close

**Outcome:** Tooling-iter cleared the `/` sentinel false-positive via `innerText` (excludes the inlined i18n
table) + dropping the empty `/` redirect-shell seed for `/home`. `failing` 3 -> 2 (the 2 remaining are genuine
empty skill-paths). escapes 2 -> 3 is a deeper-crawl discovery (3rd external chapter link, same class). The
innerText/seed-drop lesson folded into coverage-protocol.md.
**Type:** tik (iter_shape: tooling)
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n -- (2) triggered-tok: n -- (3) re-scope: n -- (4) user-blocker: n -- (5) cap-reached: **y** (tik #5) -- (6) protocol-stop: n -- Outcome: exit-5
**Decisions:** D1 (innerText), D2 (FP cleared, `/` is empty shell), D3 (drop `/` seed -> `/home`), D4 (final 2/3), D5 (escapes->3 deeper-crawl) -- see ./decisions.md
**Side-deliverables (if any):** none
**Routes carried forward:**
- iter-07 (Fate-3, handler `SEED-M42e-empty-skillpaths`): the 2 empty skill-path indexes -- the real production
  content gap (the next session's highest-leverage target; expected failing 2 -> 0).
- iter >= 07 (Fate-3, handler `ESCAPE-M42e-skillpath-external-articles`): the 3 external skill-path-chapter
  article links -- a content/allow-rule decision (rewrite/strip in replayed chapter content, or an explicit
  allow-rule if external citations are acceptable demo content -- a user-facing gate-definition call).
**Lessons:** `innerText` not `textContent` for both density + sentinel reads (excludes inlined SSR JSON);
a `<main>`-less redirect-shell route (`/`) is not a real content page -- drop it as a seed, score the real
landing (`/home`). Both fold into coverage-protocol.md. The escape-count-rising-with-crawl-depth pattern is now
3-for-3 confirmed across iters 02/05/06 -- a deeper crawl surfaces pre-existing content escapes beyond the cap;
always triage a rising escape as a discovery before assuming regression.
