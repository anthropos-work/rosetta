**Type:** tik (iter_shape: production-fix) -- under TOK-01, `sweep-then-route-by-leverage`.

# iter-08 progress -- land iter-07's three routed fixes; re-sweep to the post-fix residual

Run-2 continuation (recovery: the prior attempt stalled mid-sweep on a 600s no-output watchdog). iter-07's
close routed exactly three well-scoped Fate-3 fixes for its frontier-exhausted true residual
`(failing=3, escapes=3)`. This tik confirms each landed, then runs the AUTHORITATIVE frontier-exhausted
re-sweep to quote the post-fix residual.

## Phase A -- fast-confirm the previously-broken pages (cheap, streaming)
Re-probed the 2 empty skill-paths + the rendering sibling + the empty sim-start via `probe-empty.spec.ts`:
- **`/skill-path/prompt-engineering-fundamentals`** -- now RENDERS: `httpStatus=200, hasMain=true,
  innerTextLen=3456`, real content ("Prompt Engineering Fundamentals Intermediate 3h 29min 6 chapters 4
  simulations..."), `getSkillPath=obj` (was null), **NO federation error, NO failed requests**. FIXED by the
  taxonomy re-capture (D1).
- **`/skill-path/foundation-of-artificial-intelligence`** -- now RENDERS: `innerTextLen=2788`, real content,
  `getSkillPath=obj`, no federation error. FIXED (D1).
- **`/skill-path/learn-python-data-engineering`** (control sibling) -- still renders clean (10337 chars); no
  regression.
- **`/sim/.../start`** -- confirmed still `innerTextLen=0` (empty `<main>`) -- the runtime launch surface the
  `SIM_START_SKIP` rule excludes (D2). Expected.

## Phase D -- the authoritative frontier-exhausted re-sweep (streaming)
Ran the full coverage sweep against live demo-3 (employee/`maya-thriving`) at `COVERAGE_MAX_PAGES=150`,
STREAMING per-page `[crawl]` lines (D4 -- the observability fix that prevents the prior stall). The BFS
frontier **EXHAUSTED at reachable=93** (`cappedAtFrontier=FALSE, frontierRemaining=0`, well under the 150 cap)
-- DEEPER than iter-07's 87 because the citation allow-rule (D3) kept the `/skill-path/.../chapter` pages
PASSING (instead of the old escape-fail), so their outbound sim links got crawled + scored too (net +7
reachable vs the sim-start skip's -1).

**Post-fix authoritative residual = `(failing=0, escapes=0)`** over the 93-page frontier-exhausted reachable
set:
- **0 empty/error pages** -- both previously-empty skill-paths + their chapter pages now render full (the
  taxonomy fix); the sim-start runtime surface is correctly scoped out; all 93 reachable pages are `ok 200`.
- **0 escapes** -- the 3 editorial citations (strategy-business.com, en.wikipedia.org/wiki/Data_analysis,
  dremio.com) are now classified as VALID content (the allow-rule) and surfaced as **3 presenterNotes**, NOT
  counted as escapes. A nav-chrome/baked-URL escape would still fail (the rule is narrow to `/chapter` paths).

The spec's own gate assertion (`failing===0 && escapes===0`) PASSED (`1 passed`, 13.4 min single run).

## Phase E -- Close

**Outcome:** All three iter-07-routed fixes landed and verified by the authoritative frontier-exhausted
re-sweep: taxonomy re-capture cleared both empty skill-paths (federation error gone, `getSkillPath=obj`,
innerText 3456/2788), the sim-start skip scoped out the runtime launch surface, the citation allow-rule
reclassified the 3 escapes as valid presenter notes. **Post-fix residual = `(failing=0, escapes=0)` over the
93-page frontier-EXHAUSTED employee reachable set -- the gate is MET.** (3/3 -> 0/0.)
**Type:** tik (iter_shape: production-fix)
**Status:** closed-fixed
**Gate:** MET
**Phase 5 grading:** (1) gate-met: y -- (2) triggered-tok: n -- (3) re-scope: n -- (4) user-blocker: n -- (5) cap-reached: n (tik #2 of session) -- (6) protocol-stop: n -- Outcome: exit-1
**Decisions:** D1 (taxonomy re-capture landed -> skill-paths render), D2 (sim-start skip -> runtime surface scoped out), D3 (citation allow-rule -> 3 escapes reclassified as presenter notes), D4 (per-page `[crawl]` progress heartbeat, observability) -- see ./decisions.md
**Side-deliverables (if any):** the `crawl.ts` per-page progress heartbeat (D4) -- a reusable observability fix folded into the harness (prevents a long sweep from tripping a no-output watchdog). Folded into coverage-protocol.md as a measurement-convention note.
**Routes carried forward:** none -- the employee gate is MET. (The manager vantage is the separate milestone M42m, which reuses this harness.)
**Lessons:** (1) a stale-public-taxonomy cache silently breaks ANY page whose federated query touches a
non-nullable skiller field referencing a newly-added public skill -- the fix is a full re-capture, and the
re-sweep confirms it across the whole reachable set, not just the probed page. (2) An editorial citation
inside replayed course content is VALID content fidelity, not a gate escape -- disclose it (presenter note),
don't strip it; keep the allow-rule narrow (anchored to `/chapter` paths) so a real nav-chrome escape still
fails. (3) A long single-test sweep MUST emit per-page progress to stdout or it trips an output watchdog --
fold the heartbeat into the harness, not the runner. Lessons folded into coverage-protocol.md.
