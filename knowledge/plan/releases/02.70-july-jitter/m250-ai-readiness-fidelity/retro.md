# M250 — Retro (AI-readiness fidelity)

## Summary
Brought the Northwind AI-readiness demo from an **invented 8-skill / 6.5-weight** seed to the platform's **real
31-skill / 25.0-weight default** + **3 real track-keyed named sims** + a **net-new Directus set-dress**
(`AIReadinessSimSkillsSeeder` → `directus.simulations.skills` → the step-2 evaluated-skills list) + a **net-new
evidence-distribution** fan-out (`ai_readiness_evidence.go` → `validation_attempt_results` 5→345, 897 skill-results,
787 session-backed verified `user_skill_evidences`). Iterative: 1 tok (TOK-01) + 6 tiks, single-day serial render
loop. **Core gate parts 1/2/3/5 + the core part-4 fidelity sections proven LIVE-GREEN both vantages** (employee
`aria-completed` + manager `dana-manager`, Northwind, demo-2, cold reset-to-seed, escapes=0); arithmetic re-derived
green (Champion 30/30, started hero 9 core → 11/30; the double-round divergence made a live invariant, unreachable at
25.0). Closed-incomplete on a **user pragmatic-close mandate**: 3 adjacent manager-dashboard sections
(`by-tag`/`interview-findings`/`handled-for-you`) were **post-M246 platform drift**, fixed + data-confirmed +
unit-green, their live sweep routed to M254. rext code-of-record **`july-jitter-m250-iter07` @ 584f1fe** (on origin;
seeders + dna closure unit-GREEN). Delivers landed: `ai-readiness.md` + `seeding-spec.md`. Deferral audit GREEN.
**0 platform-repo edits.**

## Incidents This Cycle
- **P3 (no shipped bug) — three "failing" manager sections were platform DRIFT, not seeding gaps (D17).**
  iter-06's part-4 gate distance looked like a data gap; iter-07's re-survey-at-source found all three were
  **post-M246 vocabulary/KPI-id drift** the coverage manifest (and, for one, the M219 seeder) still asserted against:
  `by-tag` renamed "…by Team", `handled-for-you` renders "Hours saved", and `interview-findings`'s `usageDimSpecs`
  keys were renamed (`avg_frequency`/`avg_breadth`/`avg_context_fit` → `avg_adoption`/`avg_transformation`/
  `avg_originality`/`avg_depth`/`avg_ownership`, only `avg_depth` surviving). Fixed at the manifest + seeder;
  data-confirmed at the demo-2 DB; unit-green. Not a defect in shipped M250 code — a reminder that a fidelity milestone
  built on a post-barrier platform inherits the barrier's drift surface.
- **P3 — a stale band-label doc-comment in the rext funnel seeder** (`aiReadinessHeldSkills` lines 439-443 still show
  pre-M250 8-skill Step-1 labels; the constants + line 517-519 are correctly re-derived). NOT re-tagged at close —
  disproportionate to churn the M254 re-pin target for a comment; reconciles when M254 next touches the funnel seeder.

## What Went Well
- **Measure-first (iter-04) made the build surgical.** Reading the platform's own SQL read-paths as a faithful render
  proxy BEFORE spending an expensive live reset-to-seed pinpointed one precise gap (the validation fan-out) — one file
  (iter-05), and the re-measure flipped two gate parts at once (part-3 profile + part-4 dots share the
  `validation_attempt_skill_results` source).
- **No-fabrication moved from by-starvation to by-construction.** The 31 defaults ARE the platform's own real
  node-ids, so the config seeder writes them unconditionally (mirroring `provision.go`) and closure stays green by
  construction — cleaner than the old pool-derived "drop-if-unresolved" contract, and it inverted three starvation-
  premised fences into "writes the real defaults" (the iter-02 lesson).
- **The atomic-edit ordering held.** Landing the 8→31 arithmetic spine as ONE compile-and-fence-green commit (config +
  funnel + all fence files) before anything rendered kept a half-applied change from breaking compilation + every
  fence at once.
- **Zero platform edits.** The whole fidelity upgrade rides seeders + a Directus set-dress; the canonical repos are
  untouched.

## What Didn't
- **The live manager sweep doesn't fit the local box.** The manager coverage crawl (~150 pages) times out locally, so
  the 3 drift-fix sections' LIVE confirmation couldn't be observed on demo-2 — routed to M254 (its exit gate re-runs
  the same sweep on billion by design). The fixes are data-confirmed + unit-green, so this is an observation residual,
  not a fidelity gap — but it is the reason the milestone closes incomplete rather than gate-fired.
- **Drift ate a whole tik.** iter-07 existed only because a post-M246 platform rename made three correct-in-data
  sections read as failing. A tighter M246→M250 manifest re-sync would have caught it at build time.

## Carried Forward
- **CARRY-M250-01 → M254 (Fate 2, confirmed-covered).** The LIVE manager-coverage-sweep confirmation of the 3
  adjacent drift-fix sections (`by-tag`/`interview-findings`/`handled-for-you`). M254 `depends_on` M250; its exit gate
  **(d)** "AI-readiness page faithful per M250 gate, live, both vantages" + **(h)** the live-browser sweeps re-run it
  on billion. No sibling `overview.md` edit; recorded in `carry-forward.md` + the deferral audit.
- **DEF-M250-01 → DROPPED (D18).** `participants_filter` track-tagging + per-member business-sim session routing — a
  non-gate believability nicety whose only gate-relevant suspicion (empty `by-tag`) was falsified at iter-07 as a copy
  drift; the track label rides the landed name-heuristic set-dress.

## Metrics Delta
- **rext AI-readiness suite:** 52 test funcs across 7 files (net-new: `ai_readiness_sim_skills_test.go` +4 iter-03,
  `ai_readiness_evidence_test.go` +4 iter-05; `ai_readiness_interview_report_test.go` 7 with the iter-07 KPI-id fix;
  `m219`/`harden` fences RE-DERIVED at 25.0). Seeders pkg + dna closure **GREEN** at the code-of-record tag.
- **Live gate (demo-2, cold reset-to-seed):** parts 1/2/3/5 + core part-4 GREEN both vantages, **escapes=0**,
  **closure 31/31**. `validation_attempt_results` (AI sims) 5 → 345; verified `user_skill_evidences` 787.
- **Flake:** 0. **Platform-repo edits:** 0. **Deferral audit:** GREEN.
- Full machine-readable delta: `metrics.json`.
