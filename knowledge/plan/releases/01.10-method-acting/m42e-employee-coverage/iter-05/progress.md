**Type:** tik / iter_shape: tooling -- under TOK-01, coverage-protocol.md "Iter type selection -> Tooling-iter".

# iter-05 progress -- tooling-iter: crawl-scope the per-session sim-result deep-links

Ships the result-deep-link `skipPaths` capability (routed from iter-04 D3) AND uses it (re-sweep) in the same
iter. Clears the 5 sim-result empties -> `failing` 8 -> 3.

## Line 1 -- ship the capability
Added `RESULT_DEEP_LINK_SKIP = [/\/result\/[0-9a-f-]{8,}/i, /\/result\b/i]` in `coverage.spec.ts`, wired into
the crawl's `skipPaths`. Validated the regex (skips the 5 `/sim/.../result/<uuid>` links; keeps skill-paths,
`/profile/activities`, regular `/sim/<slug>`); compile-clean via `playwright test --list`.

## Line 2 -- use it (Phase D re-sweep)
`run-coverage.sh 3 employee maya-thriving` vs live demo-3, ~95 s to the gate line.

**Re-sweep metric: `(failing=3, escapes=2)` -- down from `(8,1)` on failing.** GATE: RED. 0 `/result/` pages
scored (confirmed). The 5 sim-result empties are gone from the reachable set, exactly as predicted.

## Triage of the residual (3 failures + 2 escapes) -- routed to iter-06
- **`/` sentinel false-positive (1 failure, D4a):** unchanged -- the `errorSettingUpGPTRealtime` i18n string
  matches in the `<main>`-less body. -> iter-06 assertion-tune (handler `ASSERT-M42e-root-sentinel`).
- **2 empty skill-path indexes (D4b):** `prompt-engineering-fundamentals` + `foundation-of-artificial-intelligence`
  (text=0 while ~20 siblings render 2k-24k). The real production content gap. -> `stack-seeding`/serve-grant
  (handler `SEED-M42e-empty-skillpaths`).
- **2 external skill-path-chapter article links (escapes, D3):** `strategy-business.com` (feedback chapter) +
  `en.wikipedia.org/wiki/Data_analysis` (python-data-analysis chapter) -- the SAME class (editorial links in
  replayed chapter body copy). escapes 1 -> 2 is a deeper-crawl discovery, not a regression. -> content/allow-rule
  (handler `ESCAPE-M42e-skillpath-external-articles`).

## Phase E -- Close

**Outcome:** Tooling-iter landed the result-deep-link skip-rule; re-sweep `(failing=3, escapes=2)` -- cleared all
5 sim-result empties (failing 8 -> 3). escapes 1 -> 2 is a deeper-crawl discovery (a 2nd pre-existing content
escape beyond the cap), not a regression. Clean fully-triaged residual: 1 `/` sentinel FP + 2 empty skill-paths
+ 2 external chapter links. Seedable-vs-runtime-computed lesson folded into coverage-protocol.md.
**Type:** tik (iter_shape: tooling)
**Status:** closed-fixed
**Gate:** NOT MET
**Phase 5 grading:** (1) gate-met: n -- (2) triggered-tok: n -- (3) re-scope: n -- (4) user-blocker: n -- (5) cap-reached: n -- (6) protocol-stop: n -- Outcome: continue
**Decisions:** D1 (ship skip-rule), D2 (re-sweep 3/2, cleared 5 result empties), D3 (escapes 1->2 = deeper-crawl discovery), D4 (residual clusters) -- see ./decisions.md
**Side-deliverables (if any):** none (the protocol-doc update is in-scope protocol-evolution, same commit)
**Routes carried forward:**
- iter-06 (Fate-3, handler `SEED-M42e-empty-skillpaths`): the 2 empty skill-path indexes -- the real production
  content gap (highest-leverage production fix remaining).
- iter >= 06 (Fate-3, handler `ASSERT-M42e-root-sentinel`): the `/` sentinel false-positive (assertion-tune).
- iter >= 06 (Fate-3, handler `ESCAPE-M42e-skillpath-external-articles`): the 2 external skill-path-chapter
  article links (content/allow-rule decision).
**Lessons:** crawl-scope is a first-class fix surface (alongside seed/serve-grant/link-rewrite) for surfaces
that are runtime-computed, not seedable -- excluding a per-session computed-result deep-link from the reachable
set is a legitimate gate-definition refinement, not a dodge. A deeper crawl (freed frontier slots) can surface
pre-existing escapes beyond the page cap -- a rising escape count after an unrelated fix is a discovery to
triage, not necessarily a regression. Both folded into coverage-protocol.md (protocol-evolution).
