# Release Retro — v2.0 "opening night" (the Playthroughs pillar)

**Shipped:** 2026-07-02 · tag `v2.0` · rext code-of-record rolled from `opening-night-m204` (@ `c81c6dd`) +
the close-release prune (`00eef00`).
**Milestones (4):** M201 manifest corpus ∥ M202 foundation → { M203 employee ∥ M204 manager }.
**Thesis delivered:** a **Playthrough is an automated actor that IS the user** — it logs in as a seeded hero,
plays a real journey end-to-end, and proves the platform delivered the outcome. Where the M42 coverage protocol
proves **presence** (every page shows real content), the Playthroughs pillar proves **function** (the hero can
*do* the thing). Corpus at ship: **10 live Playthroughs (6 employee + 4 manager) GREEN on cold reset-to-seed,
1 declared in-manifest TODO.** Tooling + docs only — **zero platform-repo edits, zero net-new third-party deps.**

## What shipped

- **M201 (corpus):** 9 products · 26 stories · 28 use-cases of prose-intent manifest, adversarially re-grounded
  (11-agent verify, ~1.3M tokens) + user-signed-off. The Product→Story→UseCase contract every coverage milestone
  implements against. Its adversarial verify **discovered the stale-clone drift** that triggered the interposed
  v1.10b "fit-up" backfill — a finding worth more than the milestone itself.
- **M202 (foundation):** the 6th rext module `playthroughs/` — manifest model + §5.3 validator (both-way id +
  precondition coverage, datadna-gated), the per-surface page-object/locator layer, the dedicated decoupled
  `pt-world` seed, the reset-to-seed serial runner, the 4-state reporting map, + the `playthroughs.md` runbook
  (which IS the M203/M204 `iteration_protocol_ref`). Built ON the M42 e2e foundation + the seeding machinery,
  **never forked.**
- **M203 (employee):** Maya's 3 core journeys GREEN — Profile (identity + verified-skill Spotlight + claimed-vs-
  verified gap + growth + timeline), Skill Paths (browse→open→start→progress), AI-Sims (chat launch at the §5.8
  boundary). 6 live Playthroughs, 5/5 cold-reset deterministic.
- **M204 (manager):** the manager vantage GREEN — Workforce funnel + roster, per-member activity-dashboard
  drill-down, succession/at-risk. 4 live Playthroughs, 5/5 deterministic.

## Metrics delta (aggregate, from `metrics.json`)

| Metric | v2.0 close | Note |
|---|---|---|
| playthroughs Go test/fuzz funcs | **105** (101 Test + 4 Fuzz) | new module: M202 96 → M203 103 → M204 105 |
| rext total Go funcs | **1745** across 6 modules | +105 (the new pillar); carried modules untouched |
| playthroughs TS unit specs | **58** (url-shapes 46 + stack-env 12) | M202 13 → M203 38 → M204 58 |
| live browser Playthroughs | **10** (6 employee + 4 manager) + 1 TODO | 5/5 cold reset-to-seed deterministic each vantage |
| module coverage | **94.8–100%** stmt | invariant-pinning at harden, not %-gains |
| flake | **0** | Go 3/3 & 5/5 shuffled; TS 3/3 & 5/5; browser 5/5 cold-reset |
| net-new third-party deps | **0** | `ai v1.40.1` stays confined to stack-seeding |

## Cross-milestone patterns

- **Reuse-not-fork held across all 4 milestones.** `hero-login.ts` imports (never forks) the M37 cockpit-login
  handshake; the seed consumes `stackseed`/`datadna` unchanged; each vantage's surfaces are a purely **additive**
  merge into the shared page-object layer (M203 employee routes `/skill-path`,`/sim`,`/profile` vs M204 manager
  `/enterprise/*` — disjoint namespaces, verified at close). The M203→M204 measure→declare→page-object→play→
  diagnose→re-measure loop transferred verbatim.
- **Recurring hazard class: green-but-wrong route/scope matching.** Three separate instances — the M203 `\b`-terminal
  route hazard (a bare word-boundary false-matches look-alike sibling segments), the M204 iter-03 SPA-URL race
  (`page.url()` lagging a client-side nav), and the M204 iter-03 out-of-`<main>` table scope (a surface that
  renders outside `<main>`). Common root: **an assertion that passes against the wrong target.** The counter-
  discipline (segment-anchored routes single-sourced in `url-shapes.ts`; `waitForURL` not `page.url()`; scope to
  the *surface* not reflexively to `<main>`) is now captured in the runbook (§locator discipline + the M204
  additions landed at this close).
- **The extract-to-browser-free-predicates move paid off twice.** Pulling route/landmark decision logic into
  `url-shapes.ts` made the M203 route bug a fast unit fix + pin (not a live-stack debug) and gave the manager
  predicates the same unit-testable single-source treatment.
- **Deferral honesty end-to-end.** The one gap (assign-WRITE UC1) was declared OUT at M201, kept as a tracked
  `playthrough: TODO`, surfaced as `unimplemented` (a first-class state, never a silent drop), and presence-pinned
  by a harden test. The 4 M203 non-gate edge UCs routed Fate-3 → M206 (roadmap-vision). No repeat-defers.

## Incidents across the release

- **P3 (M202-D4) — seeding layering collision**, surfaced + fixed at build: seeding `pt-world` onto an already-
  seeded demo-1 duplicate-keyed on the single-tenant Clerkenstein default-org slot. Zero-platform-edit fix = a
  leading anchor story (size 0); blended into `stories-spec.md`.
- **P3 (M203 iter-05) — sim-launch Sentinel deny**, diagnosed + fixed in-iter: the casbin enforcer caches policy
  in-memory; the seeded g3 feature grant isn't seen until an explicit Reload. Fixed by folding a post-seed Sentinel
  Reload into `run-playthroughs.sh` (idempotent, non-fatal, zero platform edit) + a drift guard.
- **P3 (M203 close F1) — latent `\b`-terminal route hazard**, caught at close code-review (not shipped): the harden's
  segment-anchor fix missed two inline copies. No green-but-wrong escape actually occurred (the look-alikes don't
  exist in the app); the fix removed the latent trap. **Lesson: a "we fixed the hazard" claim must sweep EVERY
  instance of the pattern, not just the centralized copy.**
- **P3 (M204 iter-02) — runner reporter-override stale-JSON hazard**, caught + fixed in-iter, pinned by a harden
  drift guard: a Playwright CLI `--reporter=` flag REPLACES the config's entire reporter list, silently suppressing
  the JSON reporter `ptreport` reads. Also fixed a latent M202/M203 wiring defect.
- **No P0/P1/P2 across the entire release.** Zero shipped incidents; every P3 was surfaced and closed within its
  own cycle or at close.

## Carry-forward (with destinations)

- **`assignment-monitoring.assign-and-track.UC1`** (assign-WRITE half, two-backend org-admin write) → **Fate-2**,
  tracked in-manifest as `playthrough: TODO` (M204 D-CLOSE-1). A future manager-write tier is the natural home.
- **4 M203 non-gate employee edge UCs** (`ai-simulations.code.UC1` Judge0 · `ai-simulations.interview.UC1` text
  interview · Skill-Paths verify-skill terminal · `profile.self-evaluation.UC1`) → **Fate-3 → M206** (roadmap-
  vision, M203 D-CLOSE-1). Each needs a live demo + browser drive infeasible at a docs-only close.
- **Future v2 vantages (roadmap-vision):** M205 Hiring + tier gates · M206 AI-sim mirror tier + the carried edge
  UCs · M207 Academy coverage. None scheduled.
- **Standing (cross-release, unchanged):** DEF-M10-01 (cloud SnapshotStore) · DEF-M21-01 (`replayCmd` hermetic
  test) · M25-D9 (dev taxonomy `rc=4`) · M314b (prod frozen-read whole-org hydration — a prod-team follow-up).

## Process notes

- **Interposed backfill lineage.** v2.0 was cut, ran M201, then **paused** for the v1.10b "fit-up" backfill that
  M201's own adversarial verify triggered (stale-clone drift), then **resumed** with no re-design needed — the
  M202/M203/M204 scaffold was already in place. The resume fast-forwarded the release branch to catch up to shipped
  v1.10b. The lesson from M201 ("check the clone's distance from prod before trusting any does-it-run judgment") is
  now baked into the backfill's milestone-1 posture.
- **project-stats layout gap (Phase 8c).** The `stats.sh` auto-detect reported 0 source/docs lines because this
  repo's code lives in the gitignored `.agentspace/rosetta-extensions/` and its docs live in `corpus/` (not
  `knowledge/`). The git-side metrics (668 commits, 202 in the last 7 days, 21% churn, 161-day age) are accurate;
  the code/doc line counts are not meaningful for this layout. Snapshot: `knowledge/journal/stats/2026-07-02.json`
  (prev: `2026-07-01.json`, the v1.10b close). Not worth wiring a custom detector for a docs-corpus; noted so a
  future reader doesn't misread the zeros.
- **No CI wired** (docs/tooling corpus) — the Phase-8b triple-clean gate ran as 3× local random-order suites
  (Go 3/3 shuffled + TS 3/3, all clean). CI-wiring remains an unscheduled nicety, not a release blocker.
- **Close-release found the release clean.** All 9 review sweeps GREEN, no blockers; the only reconciliation was a
  stale M201 `status: planned → archived` label (closed-on-gate + merged before the v1.10b pause; the close never
  flipped it). 5 small land-now fixes (1 code prune + 4 docs + 2 optional runbook additions), all Fate-1.
