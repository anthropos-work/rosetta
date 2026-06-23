# Release Retro: v1.9 "storytelling"

**Shipped:** 2026-06-23 · tag `v1.9` · branch `release/01.90-storytelling` → `main`
**Milestones:** M34 (verified-skill-chain) · M35 (stories-multi-org) · M36 (dashboard-surfaces) · M37 (clerkenstein-multi-identity) · M38 (presenter-cockpit)
**Code-of-record:** `rosetta-extensions` @ tags `storytelling-m34` (`8eb603b`) · `m35` (`06d872c`) · `m36` (`11e15e3`) · `m37` (`52c1be0`) · `m38` (`237bede`)

## What v1.9 was

The **believable-demo-narrative release.** Through v1.8 the seeder produced a structurally-correct-but-hollow
world: every user "User N", binary 85/35 scores, **zero** verified skills (the core profile rendered empty),
and — recon finding G14 — invalid free-text enum/result sessions that were *inserted-but-invisible* dead rows.
v1.9 turned the placeholder seeder into a declarative **Stories & Heroes** engine: each *story* is one org with
a thriving/struggling/manager **hero** trio, seeded via the real **verified-skill 7-table chain** so the
individual **skill profile** (Must #1) + the org **Workforce dashboard** (Must #2) tell one coherent story (the
claimed-vs-verified gap is the "aha"), plus a standalone **presenter cockpit** (login-as a hero + jump-to the
right screen) wired on **Clerkenstein multi-identity**. **Tooling + docs only — zero platform-repo edits.**

## Incidents across the release

All incidents were **build-or-close-caught and fixed in-section** — none shipped, none reached a user, zero
post-tag incidents.

- **M34 · P1 (build-caught) — missing NOT-NULL `validation_attempt_result_id` FK.** The PersonaSeeder omitted
  an FK; the hermetic unit test passed (fake conn doesn't enforce NOT-NULL) but the integration test against
  the real schema rejected the insert. Fixed + the unit test now asserts the FK non-empty. **Lesson:** a
  fake-conn unit test can't substitute for one against-real-schema pass per schema-touching seeder.
- **M35 · P2 (harden) — supporting-population name collision** (two extra "Leah Donovan"s). Fixed
  deterministically (`nameForIndexAvoiding` re-rolls off reserved hero names; first roll byte-identical).
- **M35 · P2 (build) + P3 (close) — negative-modulo panic class.** `int(hashInt(...)) % len` can be negative.
  `jobRoleRefs.at()` got the guard at build; its structurally-identical sibling `skillPool.at()` did NOT and
  surfaced at close. **Lesson (the cross-milestone pattern):** when fixing a defensive-index class, sweep ALL
  structurally identical helpers in the same pass, not just the one that crashed.
- **M36 · none.** 3 harden passes surfaced zero bugs. One close finding (seeder error-prefix dash style) was
  investigated to a non-defect (a false-positive should-fix that cost one investigation cycle).
- **M37 · none.** 2 harden passes + `FuzzLoadRoster` (360K execs, 0 panics) surfaced zero bugs. One close
  finding was a stale "four DNAs" doc claim (fixed).
- **M38 · P2 (close-found) — three-write lockstep gap in the M38-D8 fix.** The close re-fated M38-D7
  (Fate-3→Fate-1). A **crashed prior close attempt** had begun the fix (added `roleForHero` + swapped the
  `users.go` call-site) but never finished: `roster.go` still called the OLD `roleForIndex`, so a manager hero
  would export `org_role=member` while the seeder wrote `admin` — re-introducing the exact divergence the fix
  exists to prevent. **The close-review's code-quality + adversarial scans BOTH independently caught it**, and
  the test was itself asserting the old function (masking the regression). Fixed: both call-sites on
  `roleForHero` + a dedicated lockstep regression. **Lesson:** a half-landed fix is more dangerous than no fix
  — it builds and the old test passes, so it looks done; the cross-cutting close scans are exactly what catch it.

## Cross-milestone patterns

- **The "port, don't reinvent" + vertical-slice-first discipline held all the way through.** M34 ported the
  chain from the proven `/seed-verified-skill` `seed.sql` (the only bug was a *missing* column from the port,
  not a logic error); the `personas` field is purely additive (empty list = byte-identical legacy seed).
- **The normalization seam (D-M35-1, `EffectiveStories()`) is the load-bearing regression-containment move.**
  One resolved-story code path for both legacy single-org and multi-story blueprints, with byte-identical
  legacy ids, meant the existing preset/dev-min suite stayed green unchanged across M35→M38 — the roadmap's #1
  worry (multi-org refactor regresses the single-org path) never materialized.
- **Single-source discipline, extended.** M37's roster-id contract (Clerkenstein owns the registry, the seeder
  owns id-derivation, fed a roster not imported) and M38's `roleForHero` single-source (one helper that BOTH
  the seeder's membership-row/casbin-grant AND the roster claim call) are the same discipline — agreement *by
  construction*, not by hope.
- **Recurring friction: the `clerk-express-1` env-fragility.** It drives the genuine `@clerk/express` SDK and
  needs installed npm modules — unrunnable in the authoring copy. Surfaced at M37 + M38. **Not a regression**
  (never touched; its Go runner unit test passes), but the recurring "run-it-where-deps-are" friction the v1.1
  CI-wiring carry-forward already tracks.
- **Default-off blast-radius discipline.** The whole storytelling layer gates on `DEMO_STORIES=1`; every
  existing demo + all 5 Clerkenstein alignment gates stayed byte-identical through the release.

## Deferrals & carry-forward

- **Release deferral re-audit: GREEN** — 0 open, 0 repeat, 0 aged-out, **0 escape-hatch**. Both v1.9 items
  landed Fate-1 in-release: #M34-D7 → D-M35-4 (in full at M35 close), M38-D7 → M38-D8 (LAND-NOW at M38 close).
  Report: `audit-deferrals/deferral-audit-2026-06-23-release-close.md`.
- **Carry-forward (user-authorized, still open):** the live field-bake on a freshly-emptied `stack-demo/`; the
  literal browser-pixels `DEMO_STORIES=1` cockpit end-to-end (a deliberate demo step); **pushing the ext tags +
  `main` + the `v1.9` tag to `origin`** (the orchestrator's separate post-close step). These have no in-release
  home (they need a fresh box / a deliberate demo re-deploy / the orchestrator's push lane) and are tracked in
  `roadmap-vision.md` § Unscheduled backlog + `state.md`.
- **Inherited unscheduled backlog (unchanged, orthogonal):** M33 ant-academy liveness, DEF-M10-01 (cloud
  store / S3 blob bytes), DEF-M21-01..04 (prop-room residuals), M25-D9 (dev taxonomy rc=4).

## Metrics delta (aggregate — from `metrics.json`)

- **Go test funcs:** 1027 (v1.8) → **1248** (+221). stack-seeding **259 → 444** (+185); clerkenstein **223 →
  259** (+36, incl the 5th alignment DNA `clerk-multi-1` + FuzzLoadRoster). alignment/stack-snapshot/stack-secrets
  unchanged (untouched).
- **Python:** the two touched suites grew (demo-stack 138→166, stack-injection 113→117); no suite decreased.
- **Coverage:** no >2pp drop on any measured surface; the M38 producers + M37 registry ~100%.
- **Flake:** 0 (triple-clean 3/3). **Supply-chain:** GREEN (0 new deps, byte-identical to v1.8).
- **Alignment:** 100%/100% on all **5** Clerkenstein surfaces.
- **Regression verdict vs v1.8:** GREEN (no test-count decrease, no coverage drop, flake 0).

## Stats delta reference

Phase 8c snapshot: `knowledge/journal/stats/2026-06-23.json` (vs the prior `v1.5-close` snapshot). Git
velocity: **430 total commits** (+96 since v1.5-close), 44/44 milestones done, 30 commits in the last 7 days,
~2.8 commits/day over 152 project days, 62.5K lines added / 15K removed (24% churn). NOTE: the project-stats
code/docs LoC counters read 0 for rosetta because the code-of-record lives in the SEPARATE `rosetta-extensions`
repo (gitignored under `.agentspace/`) — rosetta itself is the documentation corpus; the git-velocity figures
are the meaningful project-level deltas here (same pattern as prior rosetta snapshots).
