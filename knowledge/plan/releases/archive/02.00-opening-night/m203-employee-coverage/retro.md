# M203 Retro — Employee-vantage coverage

**Closed:** 2026-07-02 · `closed-on-gate` · `iterative` · complexity large.
**Outcome:** Maya's three core employee journeys play GREEN as Playthroughs on the M202 foundation — **Profile**
(identity + verified-skill Spotlight + claimed-vs-verified gap + growth trajectory + work/education timeline),
**Skill Paths** (browse → open → start → progress; verify-skill composes on the profile side, P7), and **AI
Simulations** (chat launch, NON-voice, asserted at the §5.8 launch boundary) — **6 live Playthroughs, 0 TODO**,
proven on a COLD reset-to-seed demo with **0 false-fails over 5 consecutive reset runs** (the exit gate).
Tooling + docs only, **zero platform edits, zero new deps**. rext authoring @ `fb94458`, tagged `opening-night-m203`.

## Summary
6 iters (1 bootstrap tok + 5 tiks, all closed-fixed), TOK-01 strategy = deterministic-read-first → mutating flows
→ integration boundary. Front-loading the Profile read surfaces (iter-02/03) bought the earliest greens against
the least new machinery; the mutating Skill Paths + AI-sim flows (iter-04/05) came once the read baseline was
stable; the 5-run determinism gate (iter-06) closed it. The gate enumerated the 3 CORE journeys — all GREEN; 4
non-gate edge UCs were routed forward (see Carried Forward).

## Metrics delta (from `metrics.json`)
- **Go (rext playthroughs):** 103 test/fuzz funcs (99 Test + 4 Fuzz) — +7 vs M202's 96 (the iter-05 Sentinel-reload
  drift guard from harden + the close @pt-tag lockstep ×2 packages + invalid-engine + read-error arms). Coverage:
  manifest 100 / report 100 / ptvalidate 98.8 / ptreport 94.8 (checkPreconditionCoverage + discoverRegistry raised
  to 100% at close).
- **TS:** 38 unit specs (stack-env 12 + url-shapes 26 route/landmark predicate cases; +25 vs M202's 13) + 6 browser
  Playthroughs.
- **Flake:** 0 (Go 5/5 -shuffle + TS unit 5/5 at close; browser 5/5 cold reset-to-seed at iter-06).
- **Close findings:** 11, all Fate-1 (1 must-fix code, 2 should-fix, 1 nice-to-have, 2 docs, 3 tests, 2
  decision-triage); 0 escape-hatch.

## Incidents this cycle
- **P3 — sim-launch deny (iter-05, diagnosed + fixed).** Launching a sim as `pt-employee` hit a Sentinel deny modal
  while a showcase hero launched fine. Root cause: the running casbin enforcer caches its policy in-memory; the
  seeded g3 feature grant isn't seen until an explicit Reload RPC. Fixed by folding a post-seed Sentinel Reload into
  `run-playthroughs.sh` (idempotent, non-fatal, zero platform edits) + a cross-iter drift guard. Not a shipped
  incident — surfaced and closed within the iter.
- **P3 — green-but-wrong `\b`-terminal route hazard (caught at close code-review F1, not shipped).** The harden's
  segment-anchor fix (`\b` → `(?:[/?#]|$)`) was applied only to the two `url-shapes.ts` exported patterns; two inline
  `\b`-terminal copies survived (one in a load-bearing `waitForURL`). A `\b` false-matches look-alike sibling
  segments (`/chapter-list`, `/profile/skills-summary`). Centralized into `SKILLS_TAB_URL` + reused `CHAPTER_PLAYER_URL`
  at close; the specs never resolved on a wrong route in practice (the look-alikes don't exist in the app), so no
  green-but-wrong escape actually occurred — the fix removes the latent trap.

## What went well
- The M202 foundation held across all 6 iters — no reset-lifecycle rework, no page-object base changes; the employee
  surfaces were purely additive. The determinism gate (the hard part) passed 5/5 first try.
- Extracting the route/landmark decision logic into a browser-free `url-shapes.ts` module paid off: the close's
  route-shape bug was a fast unit fix + pin, not a live-stack debug.
- The deterministic-read-first strategy (TOK-01) never needed a triggered tok — the risk order held.

## What didn't
- The harden ledger over-claimed the `\b`-terminal fix as complete (it covered only the exported patterns, missing
  two inline duplicates). Lesson: a "we fixed the hazard" claim must sweep EVERY instance of the pattern, not just
  the centralized copy — the close code-review caught it, but the harden should have.
- The `@pt:` tag grammar was duplicated verbatim across two packages with only a "change both" comment enforcing it.
  Lesson: a cross-package "must stay identical" invariant needs a test, not a comment (landed as the twin lockstep
  tests at close, TEST-G1).

## Carried forward (three-fate, none escape-hatch)
- **`ai-simulations.code.UC1`** (Judge0 code sim) · **`ai-simulations.interview.UC1`** (text interview) · the
  **Skill-Paths verify-skill end-to-end terminal** · **`profile.self-evaluation.UC1`** (self-rate WRITE) →
  **Fate-3 → M206** (roadmap-vision.md annotated; D-CLOSE-1). All non-gate edge UCs (the gate enumerated the 3 core
  journeys, all GREEN); Fate-1 infeasible at a docs-only close (each needs a live demo + browser drive; code-sim
  needs a live Judge0 host; self-eval needs live click-intercept iteration).
- **Academy skill-path UC — OUT by M201 design** (separate Vercel deployment, M207 surface); **voice sims → M206 by
  design.** Neither a deferral.
- **Consumption re-pin — none.** The playthroughs section runs from the authoring copy against demo-1;
  `.agentspace/rext.tag` stays `v1.10.1`. Origin push of `opening-night-m203` = push-gated KEEP (user's step).
